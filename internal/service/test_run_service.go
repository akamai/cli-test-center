package service

import (
	"fmt"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/print"
	"github.com/akamai/cli-test-center/internal/util"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

func (svc Service) GetTestRunsAndPrint() {
	testRuns := svc.GetTestRuns()
	print.FormatAndPrintTestRunsTable(svc.cmd, testRuns)
}

func (svc Service) GetTestRuns() []model.TestRun {
	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.GetTestRuns), !svc.jsonOutput).Start()
	testRuns, err := svc.api.GetTestRuns()
	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}
	spinner.StopWithSuccess()

	if svc.jsonOutput {
		util.PrintJsonAndExit(testRuns)
	}

	return testRuns
}

func (svc Service) GetTestRunAndPrintResult(testRunId string) {

	testRunResult := svc.GetTestRun(testRunId)
	// Print test run result
	print.FormatAndPrintTestResult(svc.cmd, testRunResult)
}

func (svc Service) GetTestRun(testRunId string) *model.TestRun {
	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.GetTestRun), !svc.jsonOutput).Start()
	id, _ := strconv.Atoi(testRunId)
	testRun, err := svc.api.GetTestRun(id)

	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}

	spinner.StopWithSuccess()

	if svc.jsonOutput {
		util.PrintJsonAndExit(testRun)
	}

	return testRun
}

func (svc Service) RunTest(runTestUsing, testSuiteId, testSuiteName, propertyIdStr, propertyName, propVersion, url, condition, ipVersion,
	targetEnvironment, client, location, method, requestBody string, addHeader, modifyHeader, filterHeader []string,
	testRunRequestFromJson model.TestRun, encodeRequestBody bool) {

	targetEnvironment = strings.ToUpper(targetEnvironment)
	var testRun *model.TestRun

	switch runTestUsing {
	case internalconstant.RunTestUsingTestSuiteId:
		testSuiteId, _ := strconv.Atoi(testSuiteId)
		testRun = svc.runTestSuiteWithId(testSuiteId, targetEnvironment, runTestUsing)
	case internalconstant.RunTestUsingTestSuiteName:
		testRun = svc.findAndRunTestSuite(testSuiteName, targetEnvironment, runTestUsing)
	case internalconstant.RunTestUsingPropertyVersion:
		testRun = svc.runPropertyVersion(propertyIdStr, propertyName, propVersion, targetEnvironment, runTestUsing)
	case internalconstant.RunTestUsingSingleTestCase:
		testRun = svc.runTestCase(url, condition, ipVersion, addHeader, modifyHeader, filterHeader, targetEnvironment, runTestUsing,
			client, location, method, requestBody, encodeRequestBody)
	case internalconstant.RunTestUsingJsonInput:
		testRun = svc.startTestRun(testRunRequestFromJson, runTestUsing)
	}

	// Wait for test run results
	testResult, err := svc.waitForTestRunCompletion(testRun.TestRunId, runTestUsing)
	if err != nil {
		util.AbortForCommandWithSubResource(svc.cmd, err, internalconstant.TestRunResource, internalconstant.Read)
	}

	// Print test run result
	print.FormatAndPrintTestResult(svc.cmd, testResult)

}

func (svc Service) findAndRunTestSuite(testSuiteName, targetEnvironment string, runTestUsing string) *model.TestRun {
	testSuite := svc.GetSingleTestSuiteByIdOrName(internalconstant.Empty, testSuiteName, internalconstant.TestSuiteResource, false)
	if testSuite == nil {
		util.AbortWithExitCode(fmt.Sprintf(util.GetMessageForKey(svc.cmd, internalconstant.TestSuiteNameNotFound)+"\n", testSuiteName), internalconstant.ExitStatusCode0)
	}

	return svc.runTestSuiteWithId(testSuite.TestSuiteId, targetEnvironment, runTestUsing)
}

func (svc Service) runTestSuiteWithId(testSuiteId int, environment string, runTestUsing string) *model.TestRun {

	testRun := model.TestRun{
		SendEmailOnCompletion: false,
		TargetEnvironment:     environment,
		Functional: model.FunctionalTestRun{
			TestSuiteExecutions: []model.TestSuiteExecutions{
				{
					TestSuiteId: testSuiteId,
				},
			},
		},
	}

	return svc.startTestRun(testRun, runTestUsing)
}

func (svc Service) runPropertyVersion(propertyIdStr, propertyName, propVersion, environment string, runTestUsing string) *model.TestRun {
	testRun := model.TestRun{
		SendEmailOnCompletion: false,
		TargetEnvironment:     environment,
		Functional: model.FunctionalTestRun{
			PropertyManagerExecution: model.PropertyManagerExecution{
				PropertyId:      util.GetConvertedInteger(propertyIdStr),
				PropertyName:    propertyName,
				PropertyVersion: util.GetConvertedInteger(propVersion),
			},
		},
	}

	return svc.startTestRun(testRun, runTestUsing)
}

func (svc Service) runTestCase(url, condition, ipVersion string, addHeader, modifyHeader, filterHeader []string, targetEnvironment,
	runTestUsing, client, location, method, requestBody string, encodeRequestBody bool) *model.TestRun {

	var clientProfile = model.ClientProfile{IpVersion: internalconstant.Ipv4, Client: client, GeoLocation: location}
	if strings.ToUpper(ipVersion) == "V6" {
		clientProfile = model.ClientProfile{IpVersion: internalconstant.Ipv6, Client: client, GeoLocation: location}
	}

	testRun := model.TestRun{
		SendEmailOnCompletion: false,
		TargetEnvironment:     targetEnvironment,
		Functional: model.FunctionalTestRun{
			TestCaseExecution: model.TestCaseExecution{
				TestRequest: model.TestRequest{
					TestRequestUrl:    url,
					RequestMethod:     method,
					RequestHeaders:    getRequestHeaders(addHeader, modifyHeader, filterHeader),
					RequestBody:       requestBody,
					EncodeRequestBody: &encodeRequestBody,
				},
				Condition: model.Condition{
					ConditionExpression: strings.TrimSpace(condition),
				},
				ClientProfile: clientProfile,
			},
		},
	}

	return svc.startTestRun(testRun, runTestUsing)
}

func (svc Service) startTestRun(testRun model.TestRun, runTestUsing string) *model.TestRun {
	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeTestCmdSpinner, runTestUsing, internalconstant.StartTestRun), !svc.jsonOutput).Start()
	createdTestRun, err := svc.api.SubmitTestRun(testRun)
	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommandWithSubResource(svc.cmd, err, internalconstant.TestRunResource, internalconstant.Create)
		return nil
	}

	spinner.StopWithSuccess()

	if svc.jsonOutput {
		util.PrintJsonAndExit(createdTestRun)
	}

	successMessage := util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.TestRunStart)
	util.PrintSuccess(successMessage+"\n\n", createdTestRun.TestRunId)
	return createdTestRun
}

func (svc Service) waitForTestRunCompletion(testRunId int, runTestUsing string) (*model.TestRun, *model.CliError) {
	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeTestCmdSpinner, runTestUsing, internalconstant.RunTests), !svc.jsonOutput).Start()
	ticker := time.NewTicker(15 * time.Second)
	failureCount := 0

	for range ticker.C {
		log.Debugln("Polling test run status...")
		testRun, err := svc.api.GetTestRun(testRunId)
		if err != nil {
			failureCount++
			if failureCount > 3 {
				spinner.StopWithFailure()
				return nil, err
			}
			continue
		}

		// Reset failure count on success, so we only stop when 3 successive tries fail
		failureCount = 0

		if testRun != nil && testRun.Status != internalconstant.InProgressEnum {
			spinner.StopWithSuccess()
			ticker.Stop()

			return testRun, nil
		}
	}

	return nil, model.CliErrorWithMessage(externalconstant.CliErrorMessageTestRunStatus)
}

func (svc Service) GetRawRequestResponseAndPrintResult(testRunId, tcxId string) {

	if testRunId != internalconstant.Empty {
		rawReqResArray := svc.GetRawRequestResponseForRunId(testRunId)
		util.PrintHeader(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.GetRawReqResForTestRunSuccess) + "\n")
		fmt.Println()
		print.PrintRawRequestResponse(rawReqResArray, false)
	} else {
		rawRequestResponse := svc.GetRawRequestResponseForTcxId(tcxId)
		rawReqResArray := []model.RawRequestResponse{*rawRequestResponse}
		util.PrintHeader(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.GetRawReqResForTcxsSuccesss) + "\n")
		fmt.Println()
		print.PrintRawRequestResponse(rawReqResArray, true)
	}
}

func (svc Service) GetRawRequestResponseForRunId(testRunId string) []model.RawRequestResponse {
	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.GetRawReqResForTestRun), !svc.jsonOutput).Start()
	rawReqRes, err := svc.api.GetRawReqResForRunId(testRunId)

	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommandWithSubResource(svc.cmd, err, internalconstant.RawRequestUsingTestRunId, internalconstant.Empty)
	}

	spinner.StopWithSuccess()

	if svc.jsonOutput {
		util.PrintJsonAndExit(rawReqRes)
	}

	return rawReqRes
}

func (svc Service) GetRawRequestResponseForTcxId(testCaseExecutionId string) *model.RawRequestResponse {
	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.GetRawReqResForTcxs), !svc.jsonOutput).Start()
	rawRequestResponse, err := svc.api.GetRawReqResForTcxId(testCaseExecutionId)

	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommandWithSubResource(svc.cmd, err, internalconstant.RawRequestUsingTestCaseExecutionId, internalconstant.Empty)
	}

	spinner.StopWithSuccess()

	if svc.jsonOutput {
		util.PrintJsonAndExit(rawRequestResponse)
	}

	return rawRequestResponse
}

func (svc Service) GetTestLogLinesAndPrintJson(tcxId string) {
	id, _ := strconv.Atoi(tcxId)
	logLines, err := svc.api.GetLogLines(id)

	if err != nil {
		util.AbortForCommand(svc.cmd, err)
	}

	util.PrintJsonAndExit(logLines)
}
