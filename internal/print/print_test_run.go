package print

import (
	"fmt"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/util"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"sort"
	"strconv"
	"strings"
)

func FormatAndPrintTestRunsTable(cmd *cobra.Command, testRuns []model.TestRun) {

	if len(testRuns) <= 0 {
		util.PrintWarning(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.NoTestRunsFoundWarning))
		fmt.Println()
		return
	}

	// sort given items by created date / ID
	sort.SliceStable(testRuns, func(i, j int) bool {
		return testRuns[i].SubmittedDate < testRuns[j].SubmittedDate
	})

	tableHeaders := []string{"ID", "ENVIRONMENT", "STATUS", "SUBMITTED BY", "SUBMITTED DATE", "COMPLETED DATE", "NOTE"}
	var contents = make([][]string, len(testRuns))

	// prepare data for table
	for i, testRun := range testRuns {
		tsId := strconv.Itoa(testRun.TestRunId)
		row := []string{tsId, testRun.TargetEnvironment, testRun.Status, testRun.SubmittedBy, testRun.SubmittedDate, testRun.CompletedDate,
			testRun.Note}
		contents[i] = row
	}

	// show table
	util.ShowTable(tableHeaders, contents, true)

	fmt.Println()
}

func FormatAndPrintTestResult(cmd *cobra.Command, testRun *model.TestRun) {
	_, tsxResultStatsMap := GetTestRunStats(*testRun)

	util.PrintHeader("\n" + util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.TestRunHeader) + "\n\n")

	statusKey, statusColour := util.GetStatusKeyAndColour(testRun.Status)
	message := util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, statusKey)

	util.PrintLabelValueWithColour(externalconstant.LabelStatus, &statusColour, message)
	util.PrintLabelAndValue(externalconstant.LabelTargetEnvironment,
		util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, strings.ToLower(testRun.TargetEnvironment)))

	// show purge warning if present
	printPurgeWarning(cmd, testRun)
	// show allExecutionObjectsIncluded waring if false
	printAllExecutionObjectsIncludedWarning(cmd, testRun)
	fmt.Println()

	propertyManagerExecution := testRun.Functional.PropertyManagerExecution
	printPropertyManagerExecution(cmd, propertyManagerExecution, tsxResultStatsMap)

	for i, testSuiteExecution := range testRun.Functional.TestSuiteExecutions {
		printTestSuiteExecution(cmd, testSuiteExecution, tsxResultStatsMap)
		if i < len(testRun.Functional.TestSuiteExecutions)-1 {
			fmt.Println(externalconstant.SeparateLine)
		}
	}

	if testRun.Functional.TestCaseExecution.TestRequest.TestRequestUrl != internalconstant.Empty {
		logrus.Debug("Printing result for single test case execution!!!")
		printTestCaseExecution(cmd, testRun)
	}

}

func printPurgeWarning(cmd *cobra.Command, testRun *model.TestRun) {
	if testRun.PurgeInfo.Status == internalconstant.FailedEnum {
		util.PrintWarning(externalconstant.Star + util.GetApiSubErrorMessagesForCommand(cmd, testRun.PurgeInfo.Errors,
			internalconstant.Empty, internalconstant.Empty, internalconstant.PurgeOperation)[0])
		fmt.Println()
	}
}

func printAllExecutionObjectsIncludedWarning(cmd *cobra.Command, testRun *model.TestRun) {
	if testRun.Functional.AllExecutionObjectsIncluded != nil && *testRun.Functional.AllExecutionObjectsIncluded == false {
		util.PrintWarning(externalconstant.Star + util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay,
			internalconstant.Empty, internalconstant.AllExecutionObjectsIncluded))
		fmt.Println()
	}
}

func printPropertyManagerExecution(cmd *cobra.Command, propertyManagerExecution model.PropertyManagerExecution, tsxResultStatsMap map[int]model.ResultStats) {

	for i, tsx := range propertyManagerExecution.TestSuiteExecutions {
		printTestSuiteExecution(cmd, tsx, tsxResultStatsMap)
		if i < len(propertyManagerExecution.TestSuiteExecutions)-1 {
			fmt.Println(externalconstant.SeparateLine)
		}
	}
}

func printTestSuiteExecution(cmd *cobra.Command, testSuiteExecution model.TestSuiteExecutions, tsxResultStatsMap map[int]model.ResultStats) {
	tsxResultStats := tsxResultStatsMap[testSuiteExecution.TestSuiteExecutionId]

	testSuiteHeader := util.Bold(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.TestSuiteText)+externalconstant.Colon+internalconstant.Space) + testSuiteExecution.TestSuiteContext.TestSuiteName + externalconstant.SeparatePipe
	testSuiteHeader += fmt.Sprintf(util.Bold(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.TestCasesText)), testSuiteExecution.TestSuiteContext.ExecutableTestCaseCount)

	var testCaseResults []string
	if tsxResultStats.InProgressTestCasesCount > 0 {
		testCaseResults = append(testCaseResults, fmt.Sprintf(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.InProgressText), tsxResultStats.InProgressTestCasesCount))
	}
	if tsxResultStats.PassedTestCasesCount > 0 {
		testCaseResults = append(testCaseResults, color.GreenString(fmt.Sprintf(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.PassedText), tsxResultStats.PassedTestCasesCount)))
	}
	if tsxResultStats.FailedTestCasesCount > 0 {
		testCaseResults = append(testCaseResults, color.RedString(fmt.Sprintf(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.FailedText), tsxResultStats.FailedTestCasesCount)))
	}
	if tsxResultStats.InconclusiveTestCasesCount > 0 {
		testCaseResults = append(testCaseResults, color.MagentaString(fmt.Sprintf(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.SystemErrorText), tsxResultStats.InconclusiveTestCasesCount)))
	}

	testSuiteHeader += strings.Join(testCaseResults, ", ")
	fmt.Println(testSuiteHeader + "\n")

	for _, tcx := range testSuiteExecution.TestCaseExecutions {
		if tcx.DerivedTestCaseExecutions != nil {
			for _, dtcx := range tcx.DerivedTestCaseExecutions {
				printTestCaseExecutions(cmd, dtcx, strconv.Itoa(tcx.TestCaseContext.Order), strconv.Itoa(dtcx.TestCaseContext.Order))
			}
		} else {
			printTestCaseExecutions(cmd, tcx, strconv.Itoa(tcx.TestCaseContext.Order), internalconstant.Empty)
		}
	}

	fmt.Println()
}

func printTestCaseExecutions(cmd *cobra.Command, testCaseExecutions model.TestCaseExecutions, parentOrder string, childOrder string) {

	// print sort test request and client profile details
	printExecutedTestCaseDetails(testCaseExecutions.TestCaseContext.TestRequest, testCaseExecutions.TestCaseContext.ClientProfile,
		testCaseExecutions.TestCaseContext.Condition, parentOrder, childOrder, internalconstant.TabSpace, testCaseExecutions.TestCaseExecutionId)

	// print evaluation details
	printEvaluationData(cmd, testCaseExecutions.ConditionEvaluationResult.ActualConditionData, testCaseExecutions.Status,
		testCaseExecutions.ConditionEvaluationResult.Result, testCaseExecutions.Errors, testCaseExecutions.IsReevaluationInProgress,
		internalconstant.TabSpace)
}

func printTestCaseExecution(cmd *cobra.Command, testRun *model.TestRun) {

	testCaseExecution := &testRun.Functional.TestCaseExecution

	// print sort test request and client profile details
	printExecutedTestCaseDetails(testCaseExecution.TestRequest, testCaseExecution.ClientProfile, testCaseExecution.Condition,
		internalconstant.Empty, internalconstant.Empty, internalconstant.Empty, testCaseExecution.TestCaseExecutionId)

	// print evaluation details
	printEvaluationData(cmd, testCaseExecution.ConditionEvaluationResult.ActualConditionData, testCaseExecution.Status,
		testCaseExecution.ConditionEvaluationResult.Result, testCaseExecution.Errors, testCaseExecution.IsReevaluationInProgress,
		internalconstant.Empty)
}

func printExecutedTestCaseDetails(testRequest model.TestRequest, clientProfile model.ClientProfile, condition model.Condition,
	parentOrder, childOrder, tabSpaces string, testCaseExecutionId int) {

	// print order number
	if parentOrder != internalconstant.Empty {
		if childOrder != internalconstant.Empty {
			childOrder = externalconstant.Dot + childOrder
		}
		fmt.Print(tabSpaces + parentOrder + childOrder + internalconstant.Space)
	}

	// print url and client profiles
	printRequestUrlAndClientProfile(testRequest, clientProfile)

	// print test case execution id
	util.PrintLabelAndValue(tabSpaces+externalconstant.LabelTestCaseExecutionId, testCaseExecutionId)

	// print request headers if available
	printRequestHeaders(testRequest.RequestHeaders, tabSpaces)

	// print the encode request body and request body if applicable
	printRequestBody(testRequest, tabSpaces)

	// print condition
	printCondition(condition, tabSpaces)

}

func printRequestUrlAndClientProfile(testRequest model.TestRequest, clientProfile model.ClientProfile) {
	methodColour := util.GetColourForEnum(testRequest.RequestMethod, true)
	_, _ = methodColour.Printf(testRequest.RequestMethod)

	testRequestUrl := testRequest.TestRequestUrl
	if testRequest.TestRequestUrlResolved != internalconstant.Empty {
		testRequestUrl = testRequest.TestRequestUrlResolved
	}

	if clientProfile.ClientVersion != internalconstant.Empty {
		fmt.Print(internalconstant.Space + testRequestUrl + externalconstant.SeparatePipe +
			util.Bold(clientProfile.Client+externalconstant.Colon+clientProfile.ClientVersion) +
			externalconstant.PlusSign)
	} else {
		fmt.Print(internalconstant.Space + testRequestUrl + externalconstant.SeparatePipe +
			util.Bold(clientProfile.Client) + externalconstant.PlusSign)
	}

	if clientProfile.IpVersion == internalconstant.Ipv4 {
		fmt.Print(util.Bold(externalconstant.LabelIPv4))
	} else {
		fmt.Print(util.Bold(externalconstant.LabelIPv6))
	}
	fmt.Println()
}

func printRequestHeaders(headers []model.RequestHeader, tabSpaces string) {
	if headers != nil {
		var formattedRequestHeaders []string
		for _, header := range headers {

			// get header name
			var headerName = header.HeaderName
			if header.HeaderNameResolved != internalconstant.Empty {
				headerName = header.HeaderNameResolved
			}

			// get header value
			var headerValue = header.HeaderValue
			if header.HeaderValueResolved != internalconstant.Empty {
				headerValue = header.HeaderValueResolved
			}

			formattedRequestHeader := util.RequestHeaderInCLIOutputFormat(headerName, header.HeaderAction, headerValue)
			formattedRequestHeaders = append(formattedRequestHeaders, formattedRequestHeader)
		}
		util.PrintLabelAndValue(tabSpaces+externalconstant.LabelRequestHeaders, strings.Join(formattedRequestHeaders, ", "))
	}
}

func printRequestBody(testRequest model.TestRequest, tabSpaces string) {
	requestBody := testRequest.RequestBody
	if requestBody != internalconstant.Empty {
		if testRequest.RequestBodyResolved != internalconstant.Empty {
			requestBody = testRequest.RequestBodyResolved
		}
		// print condition details
		if testRequest.EncodeRequestBody != nil {
			util.PrintLabelAndValue(tabSpaces+externalconstant.LabelEncodeRequestBody, util.Italic(*testRequest.EncodeRequestBody))
		}
		util.PrintLabelAndValue(tabSpaces+externalconstant.LabelRequestBody, util.Italic(requestBody))
	}
}

func printCondition(condition model.Condition, tabSpaces string) {
	conditionExpression := condition.ConditionExpression
	if condition.ConditionExpressionResolved != internalconstant.Empty {
		conditionExpression = condition.ConditionExpressionResolved
	}

	// print condition details
	util.PrintLabelAndValue(tabSpaces+externalconstant.LabelExpected, util.Italic(conditionExpression))
}

func printEvaluationData(cmd *cobra.Command, actualConditionData []model.ActualConditionData, testCaseExecutionStatus string,
	testCaseExecutionResultStatus string, testCaseExecutionErrors []model.ApiSubError, isReevaluationInProgress *bool, tabSpaces string) {
	var actualDataList []string
	for _, actualDataItem := range actualConditionData {
		keyValuePair := actualDataItem.Name + externalconstant.Colon + internalconstant.Space
		if actualDataItem.Value != internalconstant.Empty {
			keyValuePair += externalconstant.Quote + actualDataItem.Value + externalconstant.Quote
		} else {
			keyValuePair += util.Italic(externalconstant.LabelNotFound)
		}
		actualDataList = append(actualDataList, keyValuePair)
	}

	if actualDataList != nil {
		util.PrintLabelAndValue(tabSpaces+externalconstant.LabelActual, strings.Join(actualDataList, ", "))
	}

	fmt.Print(tabSpaces + util.Bold(externalconstant.LabelStatus+externalconstant.Colon+internalconstant.Space))
	if testCaseExecutionStatus == internalconstant.CompletedEnum {
		if testCaseExecutionResultStatus == internalconstant.PassedEnum {
			util.PrintSuccess(externalconstant.LabelPassed + "\n")
		} else if testCaseExecutionResultStatus == internalconstant.Inconclusive {
			fmt.Print(color.MagentaString(externalconstant.LabelInconclusive))
			if isReevaluationInProgress != nil {
				if *isReevaluationInProgress {
					fmt.Printf(" ( %s )", util.Italic(externalconstant.LabelIsReevaluationInProgress))
				} else {
					fmt.Printf(" ( %s )", util.Italic(externalconstant.LabelIsReevaluationCompleted))
				}
			}
			fmt.Println()
		} else {
			util.PrintError(externalconstant.LabelFailed + "\n")
		}
	} else {
		if testCaseExecutionStatus == internalconstant.InProgressEnum {
			fmt.Println(externalconstant.LabelInProgress)
		} else {
			fmt.Println(color.MagentaString(externalconstant.LabelInconclusive))
			errorMessages := util.GetApiSubErrorMessagesForCommand(cmd, testCaseExecutionErrors, internalconstant.Empty,
				internalconstant.Empty, internalconstant.EvaluationErrors)
			for _, errorMessage := range errorMessages {
				util.PrintError(tabSpaces + externalconstant.LabelNotRunError + errorMessage + "\n")
			}
		}
	}

	fmt.Println()
}

func GetTestRunStats(testRun model.TestRun) (model.ResultStats, map[int]model.ResultStats) {
	resultStats := model.ResultStats{}
	tsxResultStatsMap := make(map[int]model.ResultStats)

	pmx := testRun.Functional.PropertyManagerExecution
	pmxResultStats := getPropertyManagerExecutionStats(pmx, tsxResultStatsMap)
	addResultStats(&resultStats, &pmxResultStats)

	for _, tsx := range testRun.Functional.TestSuiteExecutions {
		tsxResultStats := getTestSuiteExecutionStats(tsx)
		addResultStats(&resultStats, &tsxResultStats)
		tsxResultStatsMap[tsx.TestSuiteExecutionId] = tsxResultStats
	}

	return resultStats, tsxResultStatsMap
}

func getPropertyManagerExecutionStats(pmExecution model.PropertyManagerExecution, tsxResultStatsMap map[int]model.ResultStats) model.ResultStats {
	resultStats := model.ResultStats{}
	for _, tsx := range pmExecution.TestSuiteExecutions {
		tsxResultStats := getTestSuiteExecutionStats(tsx)
		addResultStats(&resultStats, &tsxResultStats)
		tsxResultStatsMap[tsx.TestSuiteExecutionId] = tsxResultStats
	}
	return resultStats
}

func getTestSuiteExecutionStats(testSuiteExecution model.TestSuiteExecutions) model.ResultStats {
	resultStats := getTestCaseExecutionsStats(testSuiteExecution.TestCaseExecutions)
	return resultStats
}

func getTestCaseExecutionsStats(testCaseExecutions []model.TestCaseExecutions) model.ResultStats {
	resultStats := model.ResultStats{}
	for _, tcx := range testCaseExecutions {
		if tcx.DerivedTestCaseExecutions != nil {
			for _, dtcx := range tcx.DerivedTestCaseExecutions {
				resultStats = getResultStats(dtcx, resultStats)
			}
		} else {
			resultStats = getResultStats(tcx, resultStats)
		}
	}
	return resultStats
}

func getResultStats(tcx model.TestCaseExecutions, resultStats model.ResultStats) model.ResultStats {
	if tcx.Status == internalconstant.CompletedEnum {
		if tcx.ConditionEvaluationResult.Result == internalconstant.PassedEnum {
			resultStats.PassedTestCasesCount += 1
		} else if tcx.ConditionEvaluationResult.Result == internalconstant.Inconclusive {
			resultStats.InconclusiveTestCasesCount += 1
		} else {
			resultStats.FailedTestCasesCount += 1
		}
	} else {
		if tcx.Status == internalconstant.InProgressEnum {
			resultStats.InProgressTestCasesCount += 1
		} else {
			resultStats.InconclusiveTestCasesCount += 1
		}
	}

	return resultStats
}

func addResultStats(destResultStats *model.ResultStats, resultStats *model.ResultStats) {
	destResultStats.PassedTestCasesCount += resultStats.PassedTestCasesCount
	destResultStats.FailedTestCasesCount += resultStats.FailedTestCasesCount
	destResultStats.InProgressTestCasesCount += resultStats.InProgressTestCasesCount
	destResultStats.InconclusiveTestCasesCount += resultStats.InconclusiveTestCasesCount
}

func PrintRawRequestResponse(RawReqResArray []model.RawRequestResponse, isTcx bool) {

	for i, tcx := range RawReqResArray {
		//print text case execution ids only for raw request response of test run
		if !isTcx {
			ids := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(tcx.TestCaseExecutionIds)), externalconstant.Comma), internalconstant.EmptyArray)
			util.PrintLabelAndValue(externalconstant.TestCaseExecutionIds, ids)
			fmt.Println()
		}

		fmt.Println(util.Bold(externalconstant.Request))
		fmt.Print(tcx.Request.Method + internalconstant.Space)
		fmt.Print(tcx.Request.Url + internalconstant.Space)
		fmt.Println(tcx.Request.HttpVersion)
		util.PrintRawRequestResponseHeaders(tcx.Request.Headers)
		fmt.Println()

		fmt.Println(util.Bold(externalconstant.Response))
		fmt.Print(tcx.Response.HttpVersion + internalconstant.Space)
		fmt.Print(tcx.Response.Status)
		fmt.Println(internalconstant.Space + tcx.Response.StatusText)
		util.PrintRawRequestResponseHeaders(tcx.Response.Headers)
		fmt.Println()

		//print dash lines only for test case execution from index 0 to len-2
		lastIndex := len(RawReqResArray) - 1
		if !isTcx && i != lastIndex {
			fmt.Println(externalconstant.SeparateLine)
			fmt.Println()
		}
	}

	if !isTcx {
		util.PrintTotalItems(len(RawReqResArray))
	}

}
