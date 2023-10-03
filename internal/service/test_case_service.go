package service

import (
	"fmt"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/print"
	"github.com/akamai/cli-test-center/internal/util"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func (svc Service) AddTestCaseWithTestSuite(cmd *cobra.Command, testSuiteIdStr, testSuiteName, url, condition, ipVersion string, addHeader, modifyHeader, filterHeader []string, testCaseId, client, requestMethod, requestBody string, encodeRequestBody bool, setVariables []string) {

	testSuites := svc.GetTestSuitesByIdOrName(testSuiteIdStr, testSuiteName, internalconstant.Empty, true, false, true)

	if len(testSuites) == 1 {
		testCases, err := svc.ConstructRequestAndPrintResponseForAddTestCase(testSuites[0].TestSuiteId, url, condition, ipVersion, addHeader, modifyHeader, filterHeader, testCaseId, client, requestMethod, requestBody, encodeRequestBody, setVariables)
		if testCases != nil {
			if svc.jsonOutput {
				util.PrintJsonAndExit(testCases)
			}
			util.PrintSuccess(fmt.Sprintf(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.AddTestCaseSuccess)+"\n\n", testSuites[0].TestSuiteName))
			if len(testCases.Successes[0].Warnings) != 0 {
				util.PrintWarnings(util.GetApiSubErrorMessagesForCommand(cmd, testCases.Successes[0].Warnings, internalconstant.Empty, internalconstant.Warnings, internalconstant.Empty))
			}
		} else {
			util.PrintError(fmt.Sprintf(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.AddTestCaseFail)+"\n\n", testSuites[0].TestSuiteName))
			util.AbortForCommand(svc.cmd, err)
		}
	} else {
		util.PrintError(fmt.Sprintf(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.AddTestCaseNoTestSuite)+"\n\n", testSuiteName))
	}
}

func (svc Service) EditTestCaseWithTestSuite(cmd *cobra.Command, testSuiteIdStr, testSuiteName, url, condition, ipVersion string, addHeader, modifyHeader, filterHeader []string, testCaseId, client, requestMethod, requestBody string, encodeRequestBody bool, setVariables []string) {

	testSuites := svc.GetTestSuitesByIdOrName(testSuiteIdStr, testSuiteName, internalconstant.Empty, true, false, true)

	if len(testSuites) == 1 {
		testCases, err := svc.ConstructRequestAndPrintResponseForEditTestCase(testSuites[0].TestSuiteId, url, condition, ipVersion, addHeader, modifyHeader, filterHeader, testCaseId, client, requestMethod, requestBody, encodeRequestBody, setVariables)
		if testCases != nil {
			if svc.jsonOutput {
				util.PrintJsonAndExit(testCases)
			}
			util.PrintSuccess(fmt.Sprintf(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.UpdateTestCaseSuccess)+"\n\n", testSuites[0].TestSuiteName))
			if len(testCases.Successes[0].Warnings) != 0 {
				util.PrintWarnings(util.GetApiSubErrorMessagesForCommand(cmd, testCases.Successes[0].Warnings, internalconstant.Empty, internalconstant.Warnings, internalconstant.Empty))
			}
		} else {
			util.PrintError(fmt.Sprintf(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.UpdateTestCaseFail)+"\n\n", testSuites[0].TestSuiteName))
			util.AbortForCommand(svc.cmd, err)
		}
	} else {
		util.PrintError(fmt.Sprintf(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.UpdateTestCaseNoTestSuite)+"\n\n", testSuiteName))
	}
}

func (svc Service) ConstructRequestAndPrintResponseForAddTestCase(testSuiteId int, url, condition, ipVersion string, addHeader, modifyHeader, filterHeader []string, testCaseId, client, requestMethod, requestBody string, encodeRequestBody bool, setVariables []string) (*model.TestCaseBulkResponse, *model.CliError) {
	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.AddTestCase), !svc.jsonOutput).Start()

	var testCase = constructTestCase(url, addHeader, modifyHeader, filterHeader, condition, ipVersion, testCaseId, client, requestMethod, requestBody, encodeRequestBody, setVariables)
	log.Debugf("Add test case [%+v] with the test suite id [%d]\n", testCase, testSuiteId)

	testCases, err := svc.api.AddTestCaseToTestSuite(testSuiteId, []model.TestCase{testCase})

	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}

	if len(testCases.Failures) != 0 {
		spinner.StopWithFailure()
		return nil, model.CliErrorFromPulsarProblemObject(nil, testCases.Failures, 207, externalconstant.ApiErrorAddTestCasesToTestSuitPostCall)
	}

	spinner.StopWithSuccess()
	return testCases, nil
}

func (svc Service) ConstructRequestAndPrintResponseForEditTestCase(testSuiteId int, url, condition, ipVersion string, addHeader, modifyHeader, filterHeader []string, testCaseId, client, requestMethod, requestBody string, encodeRequestBody bool, setVariables []string) (*model.TestCaseBulkResponse, *model.CliError) {
	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.UpdateTestCase), !svc.jsonOutput).Start()

	var testCase = constructTestCase(url, addHeader, modifyHeader, filterHeader, condition, ipVersion, testCaseId, client, requestMethod, requestBody, encodeRequestBody, setVariables)
	log.Debugf("Update test case [%+v] with the test suite id [%d]\n", testCase, testSuiteId)

	testCases, err := svc.api.EditTestCaseToTestSuite(testSuiteId, []model.TestCase{testCase})

	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}

	if len(testCases.Failures) != 0 {
		spinner.StopWithFailure()
		return nil, model.CliErrorFromPulsarProblemObject(nil, testCases.Failures, 207, externalconstant.ApiErrorAddTestCasesToTestSuitPostCall)
	}

	spinner.StopWithSuccess()
	return testCases, nil
}

func constructTestCase(url string, addHeader, modifyHeader, filterHeader []string, condition, ipVersion, testCaseId, client, requestMethod, requestBody string, encodeRequestBody bool, dynamicVariables []string) model.TestCase {
	var (
		testCase        model.TestCase
		testRequest     model.TestRequest
		conditionObject model.Condition
		clientProfileId int
		clientProfile   model.ClientProfile
	)

	testRequest = model.TestRequest{
		TestRequestUrl:    url,
		RequestMethod:     requestMethod,
		RequestHeaders:    getRequestHeaders(addHeader, modifyHeader, filterHeader),
		RequestBody:       requestBody,
		EncodeRequestBody: &encodeRequestBody,
	}

	conditionObject = model.Condition{
		ConditionExpression: strings.TrimSpace(condition),
	}

	// Use default IPv4
	clientProfile = model.ClientProfile{IpVersion: internalconstant.Ipv4, Client: client}
	if strings.ToUpper(ipVersion) == "V6" {
		clientProfile = model.ClientProfile{IpVersion: internalconstant.Ipv6, Client: client}
	}
	testCase = model.TestCase{
		TestRequest:     testRequest,
		Condition:       conditionObject,
		ClientProfileId: clientProfileId,
		ClientProfile:   clientProfile,
		SetVariables:    getSetVariables(dynamicVariables),
	}

	// test case id should be present for update test case request body.
	if testCaseId != internalconstant.Empty {
		tcId, _ := strconv.Atoi(testCaseId)
		testCase.TestCaseId = tcId
	}

	return testCase
}

func (svc Service) GetTestCaseByIdToTestSuite(cmd *cobra.Command, testSuiteIdStr, testSuiteName string, testCaseId string, resolveVariables bool) {

	testSuites := svc.GetTestSuitesByIdOrName(testSuiteIdStr, testSuiteName, internalconstant.Empty, true, false, true)

	if len(testSuites) == 1 {
		tcId, _ := strconv.Atoi(testCaseId)
		testCase, err := svc.GetTestCaseById(testSuites, tcId, resolveVariables)
		if testCase != nil {
			if svc.jsonOutput {
				util.PrintJsonAndExit(testCase)
			}
			fmt.Println()
			util.PrintHeader(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, "testCaseHeader") + "\n")
			print.PrintTestCaseDetails(cmd, []model.TestCase{*testCase}, false, internalconstant.Empty, 0)
			fmt.Println()
		} else {
			util.PrintError(fmt.Sprintf(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, "getTestCaseFail")+"\n\n", testSuites[0].TestSuiteName))
			util.AbortForCommand(svc.cmd, err)
		}
	} else {
		util.PrintError(fmt.Sprintf(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, "getTestCaseNoTestSuite")+"\n\n", testSuiteName))
	}

}

func (svc Service) GetTestCaseById(testSuites []model.TestSuite, testCaseId int, resolveVariables bool) (*model.TestCase, *model.CliError) {
	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, "getTestCase"), !svc.jsonOutput).Start()

	testCase, err := svc.api.GetTestCaseById(testCaseId, testSuites[0].TestSuiteId, resolveVariables)

	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}
	spinner.StopWithSuccess()
	return testCase, nil
}

func (svc Service) GetTestCasesWithTestSuite(testSuiteIdStr, testSuiteName, groupBy string, resolveVariables bool) {

	testSuites := svc.GetTestSuitesByIdOrName(testSuiteIdStr, testSuiteName, internalconstant.Empty, true, false, true)

	if len(testSuites) == 1 {
		associatedTestCases, err := svc.GetTestCasesToTestSuite(testSuites[0].TestSuiteId, resolveVariables)

		if associatedTestCases != nil {
			if svc.jsonOutput {
				util.PrintJsonAndExit(associatedTestCases)
			}
			print.PrintTestCases(svc.cmd, associatedTestCases.TestCases, associatedTestCases.AreAllTestCasesIncluded, groupBy)
		} else {
			util.PrintError(fmt.Sprintf(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, "getTestCaseFail")+"\n\n", testSuites[0].TestSuiteName))
			util.AbortForCommand(svc.cmd, err)
		}
	} else {
		util.PrintError(fmt.Sprintf(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, "getTestCaseNoTestSuite")+"\n\n", testSuiteName))
	}
}

func (svc Service) GetTestCasesToTestSuite(testSuiteId int, resolveVariables bool) (*model.AssociatedTestCases, *model.CliError) {
	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.GetTestCase), !svc.jsonOutput).Start()

	log.Debugf("Get test cases for the test suite id [%d]\n", testSuiteId)

	associatedTestCases, err := svc.api.GetV3AssociatedTestCasesForTestSuite(testSuiteId, resolveVariables)
	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}
	spinner.StopWithSuccess()
	fmt.Println()
	return associatedTestCases, nil
}

func (svc Service) RemoveTestCaseFromTestSuiteUsingOrderNumberOrTestCaseId(testSuiteIdStr, orderNumber, testCaseIdStr string) {

	testSuiteId, _ := strconv.Atoi(testSuiteIdStr)
	testSuite := svc.GetSingleTestSuiteByIdOrName(testSuiteIdStr, internalconstant.Empty, internalconstant.Empty, true)

	if testSuite != nil {

		testCases, _ := svc.GetV3AssociatedTestCasesForTestSuite(testSuiteId)

		testCase := svc.filterTestCaseUsingOrderNumberOrTestCaseId(testCases, orderNumber, testCaseIdStr)

		if testCase != nil {
			removeTestCaseResponse, err := svc.RemoveTestCasesFromTestSuite(testSuite.TestSuiteId, []int{testCase.TestCaseId})
			if err != nil {
				util.PrintError(fmt.Sprintf(util.GetMessageForKey(svc.cmd, internalconstant.FailedKey)+"\n\n", testSuite.TestSuiteName))
				util.AbortForCommand(svc.cmd, err)
			} else {
				if svc.jsonOutput {
					util.PrintJsonAndExit(removeTestCaseResponse)
				}

				util.PrintSuccess(fmt.Sprintf(util.GetMessageForKey(svc.cmd, "success")+"\n\n", testSuite.TestSuiteName))
			}
		} else {
			util.PrintError(fmt.Sprintf(util.GetMessageForKey(svc.cmd, "notPresent")+"\n\n", orderNumber))
		}
	}
}

func (svc Service) GetV3AssociatedTestCasesForTestSuite(testSuiteId int) ([]model.TestCase, bool) {

	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, "getTestCases"), !svc.jsonOutput).Start()
	associatedTestCases, err := svc.api.GetV3AssociatedTestCasesForTestSuite(testSuiteId, false)
	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommandWithSubResource(svc.cmd, err, internalconstant.Empty, internalconstant.Read)
	}

	spinner.StopWithSuccess()
	return associatedTestCases.TestCases, associatedTestCases.AreAllTestCasesIncluded
}

func (svc Service) filterTestCaseUsingOrderNumberOrTestCaseId(testCases []model.TestCase, orderNumber, testCaseId string) *model.TestCase {

	if orderNumber != internalconstant.Empty {
		orderNum, _ := strconv.Atoi(orderNumber)
		for _, testCase := range testCases {
			if testCase.Order == orderNum {
				return &testCase
			}
		}
	}
	tcId, _ := strconv.Atoi(testCaseId)
	for _, testCase := range testCases {
		if testCase.TestCaseId == tcId {
			return &testCase
		}
	}
	return nil
}
