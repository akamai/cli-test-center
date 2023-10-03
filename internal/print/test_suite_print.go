package print

// This class will have all test center related print common methods
import (
	"fmt"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/util"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type GroupedTestCases struct {
	Key   string
	Order int
	Value []model.TestCase
}

func PrintTestSuite(cmd *cobra.Command, testSuiteV3 model.TestSuite) {

	util.PrintLabelAndValue(externalconstant.LabelId, testSuiteV3.TestSuiteId)
	util.PrintLabelAndValue(externalconstant.LabelName, testSuiteV3.TestSuiteName)
	util.PrintLabelAndValue(externalconstant.LabelDescription, testSuiteV3.TestSuiteDescription)
	util.PrintLabelAndValue(externalconstant.LabelStateful, util.ConvertBooleanToYesOrNo(testSuiteV3.IsStateful))
	util.PrintLabelAndValue(externalconstant.LabelLocked, util.ConvertBooleanToYesOrNo(testSuiteV3.IsLocked))

	if testSuiteV3.Configs.PropertyManager.PropertyId != 0 {
		util.PrintLabelAndValue(externalconstant.LabelAssociatedProperty, fmt.Sprintf("(Id=%d) %s v%d",
			testSuiteV3.Configs.PropertyManager.PropertyId,
			testSuiteV3.Configs.PropertyManager.PropertyName,
			testSuiteV3.Configs.PropertyManager.PropertyVersion))
	}

	util.PrintLabelAndValue(externalconstant.LabelCreated, util.FormatTime(testSuiteV3.CreatedDate)+externalconstant.SeparateBy+testSuiteV3.CreatedBy)
	util.PrintLabelAndValue(externalconstant.LabelLastModified, util.FormatTime(testSuiteV3.ModifiedDate)+externalconstant.SeparateBy+testSuiteV3.ModifiedBy)

	if testSuiteV3.DeletedBy != internalconstant.Empty {
		util.PrintLabelAndValue(externalconstant.LabelDeleted, util.FormatTime(testSuiteV3.DeletedDate)+externalconstant.SeparateBy+testSuiteV3.DeletedBy)
	}

	// print variables if available
	if len(testSuiteV3.Variables) > 0 {
		fmt.Println()
		util.PrintHeader(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.VariableInTestSuiteHeader) + "\n")
		PrintVariablesResult(cmd, testSuiteV3.Variables, true)
	}

	fmt.Println()
}

func PrintTestSuiteForRemovedTestSuite(testSuiteV3 model.TestSuite) {

	util.PrintLabelAndValue(externalconstant.LabelId, testSuiteV3.TestSuiteId)
	util.PrintLabelAndValue(externalconstant.LabelName, testSuiteV3.TestSuiteName)
	util.PrintLabelAndValue(externalconstant.LabelDescription, testSuiteV3.TestSuiteDescription)
	var configVersion string
	if testSuiteV3.Configs.PropertyManager.PropertyName != internalconstant.Empty {
		configVersion = testSuiteV3.Configs.PropertyManager.PropertyName + " " + strconv.Itoa(testSuiteV3.Configs.PropertyManager.PropertyVersion)
	}
	util.PrintLabelAndValue(externalconstant.LabelPropertyVersion, configVersion)
}

// PrintTestSuitesResult details and return test suiteId to print test cases only if single test suite is present in array else return 0
func PrintTestSuitesResult(cmd *cobra.Command, testSuiteV3 []model.TestSuite, name string) {

	util.PrintHeader("\n" + util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.TestSuiteDetails) + "\n")

	size := len(testSuiteV3)
	switch size {
	case 0:
		fmt.Printf(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.NoTestSuiteFoundWithName)+"\n\n", name)
	case 1:
		PrintTestSuite(cmd, testSuiteV3[0])
	default:
		PrintTestSuitesTable(cmd, testSuiteV3)
	}
}

func PrintTestSuitesTable(cmd *cobra.Command, testSuites []model.TestSuite) {

	if len(testSuites) <= 0 {
		util.PrintWarning(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.NoTestSuiteFoundWarning))
		fmt.Println()
		return
	}

	// sort given items by created date / ID
	sort.SliceStable(testSuites, func(i, j int) bool {
		return testSuites[i].TestSuiteId < testSuites[j].TestSuiteId
	})

	tableHeaders := []string{"ID", "NAME", "DESCRIPTION", "PROPERTY  VERSION", "TEST CASES"}
	var contents = make([][]string, len(testSuites))

	var isDeletedPresent = false

	// prepare data for table
	for i, ts := range testSuites {

		// if deleted test-suite, add '*' at the beginning
		if ts.DeletedBy != internalconstant.Empty {
			ts.TestSuiteName = "* " + ts.TestSuiteName
			isDeletedPresent = true
		}

		// form property-version string
		propertyVersionStr := internalconstant.Empty
		if ts.Configs.PropertyManager.PropertyVersion != 0 {
			propertyVersionStr = fmt.Sprintf("%s v%d", ts.Configs.PropertyManager.PropertyName, ts.Configs.PropertyManager.PropertyVersion)
		}
		tsId := strconv.Itoa(ts.TestSuiteId)
		tcCount := strconv.Itoa(ts.ExecutableTestCaseCount)

		row := []string{tsId, ts.TestSuiteName, ts.TestSuiteDescription, propertyVersionStr, tcCount}
		contents[i] = row
	}

	// show table
	util.ShowTable(tableHeaders, contents, true)

	if isDeletedPresent {
		fmt.Println(externalconstant.LabelTSDeleteState)
	}
	fmt.Println()
}

func printGroupedTestCases(testCases []model.TestCase, groupBy string) {
	var groupedTestCases = make(map[string][]model.TestCase)

	switch strings.ToLower(groupBy) {
	case internalconstant.GroupByTestRequest:
		for _, testCase := range testCases {
			//Get test request hash
			if len(testCase.DerivedTestCases) != 0 {
				for _, dtc := range testCase.DerivedTestCases {
					testRequestHash := getTestRequestHashCode(dtc.TestRequest)
					tcs, _ := groupedTestCases[testRequestHash]
					dtc.ParentOrder = testCase.Order
					groupedTestCases[testRequestHash] = append(tcs, dtc)
				}
			} else {
				testRequestHash := getTestRequestHashCode(testCase.TestRequest)
				tcs, _ := groupedTestCases[testRequestHash]
				groupedTestCases[testRequestHash] = append(tcs, testCase)
			}

		}
		util.PrintTemplate(externalconstant.GroupByTestRequest, sortTestCasesByOrder(groupedTestCases))

	case internalconstant.GroupByCondition:
		for _, testCase := range testCases {

			if len(testCase.DerivedTestCases) != 0 {
				for _, dtc := range testCase.DerivedTestCases {
					var condition = util.GetResolvedOrUnResolvedCondition(dtc.Condition)
					tcs, _ := groupedTestCases[condition]
					dtc.ParentOrder = testCase.Order
					groupedTestCases[condition] = append(tcs, dtc)
				}
			} else {
				var condition = util.GetResolvedOrUnResolvedCondition(testCase.Condition)
				tcs, _ := groupedTestCases[condition]
				groupedTestCases[condition] = append(tcs, testCase)
			}
		}
		util.PrintTemplate(externalconstant.GroupByCondition, sortTestCasesByOrder(groupedTestCases))

	case internalconstant.GroupByClientProfile:
		for _, testCase := range testCases {
			if len(testCase.DerivedTestCases) != 0 {
				for _, dtc := range testCase.DerivedTestCases {
					clientProfileHash := util.ClientProfileInCLIOutputFormat(dtc.ClientProfile)
					tcs, _ := groupedTestCases[clientProfileHash]
					dtc.ParentOrder = testCase.Order
					groupedTestCases[clientProfileHash] = append(tcs, dtc)
				}
			} else {
				clientProfileHash := util.ClientProfileInCLIOutputFormat(testCase.ClientProfile)
				tcs, _ := groupedTestCases[clientProfileHash]
				groupedTestCases[clientProfileHash] = append(tcs, testCase)
			}
		}

		util.PrintTemplate(externalconstant.GroupByClientProfile, sortTestCasesByOrder(groupedTestCases))
	}
}

func getTestRequestHashCode(testRequest model.TestRequest) string {

	//sort tags first so that sequence of tags is always matching in different test requests.
	sort.Strings(testRequest.Tags)
	tags := strings.Join(testRequest.Tags, externalconstant.Comma)

	sort.Slice(testRequest.RequestHeaders, func(i, j int) bool {
		return testRequest.RequestHeaders[i].HeaderName < testRequest.RequestHeaders[j].HeaderName
	})

	var headersArray []string
	for _, v := range testRequest.RequestHeaders {
		var resolvedHeaders = util.GetResolvedOrUnResolvedHeaders(v)
		headersArray = append(headersArray, resolvedHeaders)
	}
	headers := strings.Join(headersArray, externalconstant.Comma)

	var requestBody = internalconstant.Empty
	var encodeRequestBody = internalconstant.Empty
	if testRequest.RequestMethod == internalconstant.PostRequestMethod {
		requestBody = util.GetResolvedOrUnResolvedRequestBody(testRequest)
		encodeRequestBody = strconv.FormatBool(*testRequest.EncodeRequestBody)
	}
	var testRequestURL = util.GetResolvedOrUnResolvedRequestURL(testRequest)

	return fmt.Sprintf("%s/%s/%s/%s/%s/%s", testRequestURL, testRequest.RequestMethod, tags, headers, requestBody, encodeRequestBody)

}

// Sort test cases values of map by order number and also get map of the lowest order number from test cases to values in map
// The user should be able to view the test cases sorted by the order number in the list view and group by view. Grouped by object with the smallest order number test case will be at the top.
func sortTestCasesByOrder(testCaseMap map[string][]model.TestCase) []GroupedTestCases {

	var groupedTestCasesMap = make([]GroupedTestCases, 0)
	for key, testCases := range testCaseMap {
		sort.Slice(testCases, func(i, j int) bool {
			// considered derived test cases also for sorting.
			var order1 = testCases[i].Order
			var order2 = testCases[j].Order
			if testCases[i].ParentOrder != 0 {
				order1 = testCases[i].ParentOrder
			}
			if testCases[j].ParentOrder != 0 {
				order2 = testCases[j].ParentOrder
			}
			return order1 < order2
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

func printMissingTestCasesWarning(cmd *cobra.Command, areAllTestCasesIncluded bool) {
	if !areAllTestCasesIncluded {
		fmt.Println()
		fmt.Println()
		util.PrintWarning(util.GetMessageForKey(cmd, "hostnameAccessMissing"))
	}
}
