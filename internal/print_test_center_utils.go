package internal

// This class will have all test center related print common methods
import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type GroupedTestCases struct {
	Key   string
	Order int
	Value []TestCase
}

func GetFunctionalContextMap(testRunContext *TestRunContext) (*FunctionalContextMap, *CliError) {
	functionalContext := testRunContext.Functional
	var functionalContextMap = FunctionalContextMap{}
	functionalContextMap.ConfigVersionsMap = make(map[int]ConfigVersionContextMap)

	for _, configVersionContext := range functionalContext.ConfigVersions {
		var configVersionContextMap ConfigVersionContextMap
		err := copier.Copy(&configVersionContextMap, &configVersionContext)
		if err != nil {
			log.Errorln(err)
			return nil, CliErrorWithMessage(CliErrorMessageTestRunContext)
		}
		configVersionContextMap.TestSuitesMap = testSuiteContextListToMap(configVersionContext.TestSuites)
		functionalContextMap.ConfigVersionsMap[configVersionContext.ConfigVersionId] = configVersionContextMap
	}

	functionalContextMap.TestSuitesMap = testSuiteContextListToMap(functionalContext.TestSuites)

	return &functionalContextMap, nil
}

func testSuiteContextListToMap(testSuiteContexts []TestSuiteContext) map[int]TestSuiteContextMap {
	testSuitesContextMap := make(map[int]TestSuiteContextMap)
	for _, testSuiteContext := range testSuiteContexts {
		var testSuiteContextMap TestSuiteContextMap
		err := copier.Copy(&testSuiteContextMap, &testSuiteContext)
		if err != nil {
			return nil
		}
		testSuiteContextMap.TestCasesMap = testCasesListToMap(testSuiteContext.TestCases)
		testSuitesContextMap[testSuiteContext.TestSuiteId] = testSuiteContextMap
	}
	return testSuitesContextMap
}

func testCasesListToMap(testCases []TestCase) map[int]TestCase {
	testCasesMap := make(map[int]TestCase)
	for _, testCase := range testCases {
		testCasesMap[testCase.TestCaseId] = testCase
	}
	return testCasesMap
}

func GetTestRunStats(testRun TestRun) (ResultStats, map[int]ResultStats) {
	resultStats := ResultStats{}
	tsxResultStatsMap := make(map[int]ResultStats)

	for _, cvx := range testRun.Functional.ConfigVersionExecutions {
		cvxResultStats := getConfigVersionExecutionStats(cvx, tsxResultStatsMap)
		addResultStats(&resultStats, &cvxResultStats)

	}

	for _, tsx := range testRun.Functional.TestSuiteExecutions {
		tsxResultStats := getTestSuiteExecutionStats(tsx)
		addResultStats(&resultStats, &tsxResultStats)
		tsxResultStatsMap[tsx.TestSuiteExecutionId] = tsxResultStats
	}

	return resultStats, tsxResultStatsMap
}

func getTestCaseExecutionsStats(testCaseExecutions []TestCaseExecutionV2) ResultStats {
	resultStats := ResultStats{}
	for _, tcx := range testCaseExecutions {
		resultStats.TotalTestCasesCount += 1
		if tcx.Status == Completed {
			if tcx.ConditionEvaluationResult.Result == Passed {
				resultStats.PassedTestCasesCount += 1
			} else {
				resultStats.FailedTestCasesCount += 1
			}
		} else {
			resultStats.FailedTestCasesCount += 1
		}
	}
	return resultStats
}

func getTestSuiteExecutionStats(testSuiteExecution TestSuiteExecution) ResultStats {
	resultStats := getTestCaseExecutionsStats(testSuiteExecution.TestCaseExecutionV2)
	return resultStats
}

func getConfigVersionExecutionStats(configVersionExecution ConfigVersionExecution, tsxResultStatsMap map[int]ResultStats) ResultStats {
	resultStats := ResultStats{}
	for _, tsx := range configVersionExecution.TestSuiteExecutions {
		tsxResultStats := getTestSuiteExecutionStats(tsx)
		addResultStats(&resultStats, &tsxResultStats)
		tsxResultStatsMap[tsx.TestSuiteExecutionId] = tsxResultStats
	}
	return resultStats
}

func addResultStats(destResultStats *ResultStats, resultStats *ResultStats) {
	destResultStats.TotalTestCasesCount += resultStats.TotalTestCasesCount
	destResultStats.PassedTestCasesCount += resultStats.PassedTestCasesCount
	destResultStats.FailedTestCasesCount += resultStats.FailedTestCasesCount
	destResultStats.SystemErrorTestCasesCount += resultStats.SystemErrorTestCasesCount
}

func PrintTestResult(cmd *cobra.Command, testRun *TestRun, testRunContext *TestRunContext) {
	testRunStats, tsxResultStatsMap := GetTestRunStats(*testRun)
	functionalContextMap, err := GetFunctionalContextMap(testRunContext)
	if err != nil {
		AbortForCommand(cmd, err)
	}

	var testRunStatus string
	if testRunStats.PassedTestCasesCount == testRunStats.TotalTestCasesCount {
		testRunStatus = color.GreenString(GetServiceMessage(cmd, MessageTypeDisplay, "", "completed"))
	} else if testRunStats.PassedTestCasesCount < testRunStats.TotalTestCasesCount {
		testRunStatus = color.YellowString(GetServiceMessage(cmd, MessageTypeDisplay, "", "completedNotAsExpected"))
	} else if testRunStats.FailedTestCasesCount == testRunStats.TotalTestCasesCount {
		testRunStatus = color.RedString(GetServiceMessage(cmd, MessageTypeDisplay, "", "failed"))
	}
	PrintHeader("\n" + GetServiceMessage(cmd, MessageTypeDisplay, "", "testRunHeader") + "\n\n")
	printLabelAndValue(LabelStatus, testRunStatus)
	printLabelAndValue(LabelTargetEnvironment, CamelToTitle(testRun.TargetEnvironment))
	fmt.Println()

	for i, configVersionExecution := range testRun.Functional.ConfigVersionExecutions {
		PrintConfigVersionExecution(cmd, configVersionExecution, functionalContextMap.ConfigVersionsMap, tsxResultStatsMap)
		if i < len(testRun.Functional.ConfigVersionExecutions)-1 {
			fmt.Println(SeparateLine)
		}
	}

	for i, testSuiteExecution := range testRun.Functional.TestSuiteExecutions {
		PrintTestSuiteExecution(cmd, testSuiteExecution, functionalContextMap.TestSuitesMap, tsxResultStatsMap)
		if i < len(testRun.Functional.TestSuiteExecutions)-1 {
			fmt.Println(SeparateLine)
		}
	}

	if testRun.Functional.TestCaseExecutionV3.TestRequest.TestRequestUrl != Empty {
		log.Debug("Printing result for single test case execution!!!")
		PrintTestCaseExecution(testRun)
	}

}

func PrintConfigVersionExecution(cmd *cobra.Command, configVersionExecution ConfigVersionExecution, configVersionsContextMap map[int]ConfigVersionContextMap, tsxResultStatsMap map[int]ResultStats) {
	propertyVersion := configVersionsContextMap[configVersionExecution.ConfigVersionId]
	fmt.Printf("%s: %s v%d\n\n", bold(LabelPropertyVersion), propertyVersion.PropertyName, propertyVersion.PropertyVersion)

	for i, tsx := range configVersionExecution.TestSuiteExecutions {
		PrintTestSuiteExecution(cmd, tsx, propertyVersion.TestSuitesMap, tsxResultStatsMap)
		if i < len(configVersionExecution.TestSuiteExecutions)-1 {
			fmt.Println(SeparateLine)
		}
	}
}

func PrintTestSuiteExecution(cmd *cobra.Command, testSuiteExecution TestSuiteExecution, testSuitesContextMap map[int]TestSuiteContextMap, tsxResultStatsMap map[int]ResultStats) {
	testSuite := testSuitesContextMap[testSuiteExecution.TestSuiteId]
	tsxResultStats := tsxResultStatsMap[testSuiteExecution.TestSuiteExecutionId]

	testSuiteHeader := bold(GetServiceMessage(cmd, MessageTypeDisplay, "", "testSuiteText")) + testSuite.TestSuiteName + SeparatePipe
	testSuiteHeader += fmt.Sprintf(GetServiceMessage(cmd, MessageTypeDisplay, "", "testCases"), tsxResultStats.TotalTestCasesCount)

	var testCaseResults []string
	if tsxResultStats.PassedTestCasesCount > 0 {
		testCaseResults = append(testCaseResults, fmt.Sprintf(GetServiceMessage(cmd, MessageTypeDisplay, "", "passedText"), tsxResultStats.PassedTestCasesCount))
	}
	if tsxResultStats.FailedTestCasesCount > 0 {
		testCaseResults = append(testCaseResults, color.RedString(fmt.Sprintf(GetServiceMessage(cmd, MessageTypeDisplay, "", "failedText"), tsxResultStats.FailedTestCasesCount)))
	}
	if tsxResultStats.SystemErrorTestCasesCount > 0 {
		testCaseResults = append(testCaseResults, color.YellowString(fmt.Sprintf(GetServiceMessage(cmd, MessageTypeDisplay, "", "systemErrorText"), tsxResultStats.SystemErrorTestCasesCount)))
	}

	testSuiteHeader += strings.Join(testCaseResults, ", ")
	fmt.Println(testSuiteHeader + "\n")

	for _, tcx := range testSuiteExecution.TestCaseExecutionV2 {
		PrintTestCasesExecution(tcx, testSuite.TestCasesMap)
	}

	fmt.Println()
}

func PrintTestCasesExecution(testCaseExecution TestCaseExecutionV2, testCasesMap map[int]TestCase) {
	testCase := testCasesMap[testCaseExecution.TestCaseId]

	if testCase.ClientProfile.IpVersion == Ipv4 {
		fmt.Println(testCase.TestRequest.TestRequestUrl + SeparatePipe + bold(LabelIPv4))
	} else {
		fmt.Println(testCase.TestRequest.TestRequestUrl + SeparatePipe + bold(LabelIPv6))
	}

	printLabelAndValue(LabelExpected, italic(testCase.Condition.ConditionExpression))

	var actualDataList []string
	for _, actualDataItem := range testCaseExecution.ConditionEvaluationResult.ActualConditionData {
		keyValuePair := CamelToTitle(actualDataItem.Name) + ": "
		if actualDataItem.Value != "" {
			keyValuePair += actualDataItem.Value
		} else {
			keyValuePair += italic(LabelNotFound)
		}
		actualDataList = append(actualDataList, keyValuePair)
	}
	printLabelAndValue(LabelActual, strings.Join(actualDataList, ", "))

	if testCaseExecution.Status == Completed {
		if testCaseExecution.ConditionEvaluationResult.Result == Passed {
			PrintSuccess(LabelPassed + "\n")
		} else {
			PrintError(LabelFailed + "\n")
		}
	} else {
		var errorMessages []string
		for _, errorObject := range testCaseExecution.Errors {
			errorMessages = append(errorMessages, errorObject.Title)
		}
		PrintError(LabelNotRunError + strings.Join(errorMessages, ", ") + "\n")
	}

	fmt.Println()
}

func PrintTestCaseExecution(testRun *TestRun) {

	testCaseExecution := &testRun.Functional.TestCaseExecutionV3
	if testCaseExecution.ClientProfile.IpVersion == Ipv4 {
		fmt.Println(testCaseExecution.TestRequest.TestRequestUrl + SeparatePipe + bold(LabelIPv4))
	} else {
		fmt.Println(testCaseExecution.TestRequest.TestRequestUrl + SeparatePipe + bold(LabelIPv6))
	}

	printLabelAndValue(LabelExpected, italic(testCaseExecution.Condition.ConditionExpression))

	var actualDataList []string
	for _, actualDataItem := range testCaseExecution.ConditionEvaluationResult.ActualConditionData {
		keyValuePair := CamelToTitle(actualDataItem.Name) + ": "
		if actualDataItem.Value != "" {
			keyValuePair += actualDataItem.Value
		} else {
			keyValuePair += italic(LabelNotFound)
		}
		actualDataList = append(actualDataList, keyValuePair)
	}
	printLabelAndValue(LabelActual, strings.Join(actualDataList, ", "))

	if testCaseExecution.Status == Completed {
		if testCaseExecution.ConditionEvaluationResult.Result == Passed {
			PrintSuccess(LabelPassed + "\n")
		} else {
			PrintError(LabelFailed + "\n")
		}
	} else {
		var errorMessages []string
		for _, errorObject := range testCaseExecution.Errors {
			errorMessages = append(errorMessages, errorObject.Title)
		}
		PrintError(LabelNotRunError + strings.Join(errorMessages, ", ") + "\n")
	}

	fmt.Println()
}

func PrintTestSuite(testSuiteV3 TestSuiteV3) {

	printLabelAndValue(LabelId, testSuiteV3.TestSuiteId)
	printLabelAndValue(LabelName, testSuiteV3.TestSuiteName)
	printLabelAndValue(LabelDescription, testSuiteV3.TestSuiteDescription)
	printLabelAndValue(LabelStateful, ConvertBooleanToYesOrNo(testSuiteV3.Stateful))
	printLabelAndValue(LabelLocked, ConvertBooleanToYesOrNo(testSuiteV3.Locked))
	printLabelAndValue(LabelVariables, "")
	printVariables(testSuiteV3.Variables)
	if testSuiteV3.Configs.PropertyManager.PropertyVersion != 0 {
		printLabelAndValue(LabelAssociatedPropertyVersion, fmt.Sprintf("%s %d",
			testSuiteV3.Configs.PropertyManager.PropertyName, testSuiteV3.Configs.PropertyManager.PropertyVersion))
	} else {
		printLabelAndValue(LabelAssociatedPropertyVersion, fmt.Sprintf("%s", testSuiteV3.Configs.PropertyManager.PropertyName))
	}

	printLabelAndValue(LabelCreated, FormatTime(testSuiteV3.CreatedDate)+SeparateBy+testSuiteV3.CreatedBy)
	printLabelAndValue(LabelLastModified, FormatTime(testSuiteV3.ModifiedDate)+SeparateBy+testSuiteV3.ModifiedBy)

	if testSuiteV3.DeletedBy != "" {
		printLabelAndValue(LabelDeleted, FormatTime(testSuiteV3.DeletedDate)+SeparateBy+testSuiteV3.DeletedBy)
	}

	fmt.Println()
}

func PrintTestSuiteForRemovedTestSuite(testSuiteV3 TestSuiteV3) {

	printLabelAndValue(LabelId, testSuiteV3.TestSuiteId)
	printLabelAndValue(LabelName, testSuiteV3.TestSuiteName)
	printLabelAndValue(LabelDescription, testSuiteV3.TestSuiteDescription)
	var configVersion string
	if testSuiteV3.Configs.PropertyManager.PropertyName != "" {
		configVersion = testSuiteV3.Configs.PropertyManager.PropertyName + " " + strconv.Itoa(testSuiteV3.Configs.PropertyManager.PropertyVersion)
	}
	printLabelAndValue(LabelPropertyVersion, configVersion)
}

// PrintViewTestSuite details and return test suiteId to print test cases only if single test suite is present in array else return 0
func PrintViewTestSuite(cmd *cobra.Command, testSuiteV3 []TestSuiteV3, name string) {

	PrintHeader("\n" + GetServiceMessage(cmd, MessageTypeDisplay, "", "testSuiteDetails") + "\n")

	size := len(testSuiteV3)
	switch size {
	case 0:
		fmt.Printf(GetServiceMessage(cmd, MessageTypeDisplay, "", "noTestSuiteFoundWithName")+"\n\n", name)
	case 1:
		PrintTestSuite(testSuiteV3[0])
	default:
		PrintTestSuitesTable(cmd, testSuiteV3)
	}
}

func printVariables(variables []Variable) {

	for _, variable := range variables {
		fmt.Printf("%v= %v\n", variable.VariableName, variable.VariableValue)
	}
}

func PrintTestSuitesTable(cmd *cobra.Command, testSuites []TestSuiteV3) {

	if len(testSuites) <= 0 {
		PrintWarning(GetServiceMessage(cmd, MessageTypeDisplay, "", "noTestSuiteFoundWarning"))
		fmt.Println()
		return
	}

	// sort given items by created date / ID
	sort.SliceStable(testSuites, func(i, j int) bool {
		return testSuites[i].TestSuiteId < testSuites[j].TestSuiteId
	})

	tableHeaders := []string{"ID", "Name", "Description", "Property Version", "Test Cases"}
	var contents = make([][]string, len(testSuites))

	var isDeletedPresent = false

	// prepare data for table
	for i, ts := range testSuites {

		// if deleted test-suite, add '*' at the beginning
		if ts.DeletedBy != "" {
			ts.TestSuiteName = "* " + ts.TestSuiteName
			isDeletedPresent = true
		}

		// form property-version string
		propertyVersionStr := ""
		if ts.Configs.PropertyManager.PropertyVersion != 0 {
			propertyVersionStr = fmt.Sprintf("%s v%d", ts.Configs.PropertyManager.PropertyName, ts.Configs.PropertyManager.PropertyVersion)
		}
		tsId := strconv.Itoa(ts.TestSuiteId)
		tcCount := strconv.Itoa(ts.TestCaseCount)

		row := []string{tsId, ts.TestSuiteName, ts.TestSuiteDescription, propertyVersionStr, tcCount}
		contents[i] = row
	}

	// show table
	ShowTable(tableHeaders, contents)

	if isDeletedPresent {
		fmt.Println(LabelTSDeleteState)
	}
	fmt.Println()
}

func PrintTestCases(cmd *cobra.Command, testCases []TestCase, allTestCasesIncluded bool, groupBy string) {

	//Print All associated Test Cases
	PrintHeader(GetServiceMessage(cmd, MessageTypeDisplay, "", "testCaseHeader") + "\n")

	if len(testCases) <= 0 {
		PrintWarning(GetServiceMessage(cmd, MessageTypeDisplay, "", "noTestCaseWarning"))
		//Printing warning here only and returning from method so that empty table does not print
		printMissingTestCasesWarning(cmd, allTestCasesIncluded)
		fmt.Println()
		fmt.Println()
		return
	}

	if groupBy == "" {
		printTestCasesTable(testCases)
	} else {
		printGroupedTestCases(testCases, groupBy)
	}

	printMissingTestCasesWarning(cmd, allTestCasesIncluded)
	fmt.Println()
	fmt.Println()
}

func printTestCasesTable(testCases []TestCase) {
	// sort given items by orderId
	sort.SliceStable(testCases, func(i, j int) bool {
		return testCases[i].Order < testCases[j].Order
	})

	tableHeaders := []string{"Order Number", "Test Request", "Condition", "Ip Version"}
	var contents = make([][]string, len(testCases))

	// prepare data for table
	for i, tc := range testCases {

		var testRequest strings.Builder
		testRequest.WriteString(tc.TestRequest.TestRequestUrl)
		testRequest.WriteString("\n")
		if len(tc.TestRequest.Tags) != 0 {
			testRequest.WriteString("Keywords: ")
			testRequest.WriteString(strings.Join(tc.TestRequest.Tags, ","))
			testRequest.WriteString("\n")
		}

		if len(tc.TestRequest.RequestHeaders) != 0 {
			testRequest.WriteString("Request Headers: ")
			for _, header := range tc.TestRequest.RequestHeaders {
				testRequest.WriteString(RequestHeaderInCLIOutputFormat(header.HeaderName, header.HeaderAction, header.HeaderValue))
				testRequest.WriteString("\n")
			}
		}

		var clientProfile = ClientProfileInCLIOutputFormat(tc.ClientProfile.IpVersion)
		row := []string{strconv.Itoa(tc.Order), testRequest.String(), tc.Condition.ConditionExpression, clientProfile}
		contents[i] = row
	}

	// show table
	ShowTable(tableHeaders, contents)
}

func printGroupedTestCases(testCases []TestCase, groupBy string) {
	var groupedTestCases = make(map[string][]TestCase)

	switch strings.ToLower(groupBy) {
	case GroupByTestRequest:
		for _, testCase := range testCases {
			//Get test request hash
			testRequestHash := getTestRequestHashCode(testCase.TestRequest)
			tcs, _ := groupedTestCases[testRequestHash]
			groupedTestCases[testRequestHash] = append(tcs, testCase)
		}
		printTemplate(groupByTestRequest, sortTestCasesByOrder(groupedTestCases))

	case GroupByCondition:
		for _, testCase := range testCases {
			tcs, _ := groupedTestCases[testCase.Condition.ConditionExpression]
			groupedTestCases[testCase.Condition.ConditionExpression] = append(tcs, testCase)
		}
		printTemplate(groupByCondition, sortTestCasesByOrder(groupedTestCases))

	case GroupByIpVersion:
		for _, testCase := range testCases {
			tcs, _ := groupedTestCases[testCase.ClientProfile.IpVersion]
			groupedTestCases[testCase.ClientProfile.IpVersion] = append(tcs, testCase)
		}

		printTemplate(groupByIpVersion, sortTestCasesByOrder(groupedTestCases))
	}
}

func getTestRequestHashCode(testRequest TestRequest) string {

	//sort tags first so that sequence of tags is always matching in different test requests.
	sort.Strings(testRequest.Tags)
	tags := strings.Join(testRequest.Tags, ",")

	sort.Slice(testRequest.RequestHeaders, func(i, j int) bool {
		return testRequest.RequestHeaders[i].HeaderName < testRequest.RequestHeaders[j].HeaderName
	})

	var headersArray []string
	for _, v := range testRequest.RequestHeaders {
		headersArray = append(headersArray, v.HeaderName, v.HeaderValue, v.HeaderAction)
	}
	headers := strings.Join(headersArray, ",")

	return fmt.Sprintf("%s/%s/%s", testRequest.TestRequestUrl, tags, headers)
}

// Sort test cases values of map by order number and also get map of the lowest order number from test cases to values in map
// The user should be able to view the test cases sorted by the order number in the list view and group by view. Grouped by object with the smallest order number test case will be at the top.
func sortTestCasesByOrder(testCaseMap map[string][]TestCase) []GroupedTestCases {

	var groupedTestCasesMap = make([]GroupedTestCases, 0)
	for key, testCases := range testCaseMap {
		sort.Slice(testCases, func(i, j int) bool {
			return testCases[i].Order < testCases[j].Order
		})

		groupedTestCasesMap = append(groupedTestCasesMap, GroupedTestCases{
			Key:   key,
			Order: testCases[0].Order,
			Value: testCases,
		})
	}

	sort.Slice(groupedTestCasesMap, func(i, j int) bool {
		return groupedTestCasesMap[i].Order < groupedTestCasesMap[j].Order
	})

	//use new key and values list in struct to print values
	return groupedTestCasesMap
}

func printMissingTestCasesWarning(cmd *cobra.Command, allTestCasesIncluded bool) {
	if !allTestCasesIncluded {
		fmt.Println()
		fmt.Println()
		PrintWarning(GetMessageForKey(cmd, "hostnameAccessMissing"))
	}
}

func PrintConditions(cmd *cobra.Command, condTemplate ConditionTemplate) {
	fmt.Println()
	PrintHeader(GetServiceMessage(cmd, MessageTypeDisplay, "", "conditionsHeader"))
	fmt.Println()
	fmt.Println()
	for _, condType := range condTemplate.ConditionTypes {
		fmt.Println(bold(condType.Label))
		for _, condExpression := range condType.ConditionExpressions {
			fmt.Println("  ", condExpression.ConditionExpression)
		}
		fmt.Println()
	}
}
