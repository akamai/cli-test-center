package internal

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Service struct {
	api        ApiClient
	cmd        *cobra.Command
	jsonOutput bool
}

func NewService(api ApiClient, cmd *cobra.Command, jsonOutput bool) *Service {
	return &Service{api, cmd, jsonOutput}
}

func (svc Service) GetTestSuites(propertyName, propVersion, user, searchString string) []TestSuiteV3 {

	spinner := NewSpinner(GetServiceMessage(svc.cmd, MessageTypeSpinner, "", "getTestSuite"), !svc.jsonOutput).Start()
	testSuitesV3, err := svc.api.GetTestSuitesV3(propertyName, propVersion, user)
	if err != nil {
		spinner.StopWithFailure()
		AbortForCommand(svc.cmd, err)
	}
	spinner.StopWithSuccess()

	filteredTestSuites := filterTestSuitesByString(searchString, false, true, testSuitesV3)

	if svc.jsonOutput {
		PrintJsonAndExit(filteredTestSuites)
	}

	return filteredTestSuites
}

func (svc Service) ImportTestSuites(testSuiteImport TestSuiteDetailsWithChildObjects) {

	spinner := NewSpinner(GetServiceMessage(svc.cmd, MessageTypeSpinner, "", "importTestSuite"), !svc.jsonOutput).Start()
	testSuiteImportResponseV3, err := svc.api.ImportTestSuite(testSuiteImport)

	if err != nil {
		spinner.StopWithFailure()
		AbortForCommand(svc.cmd, err)
	}
	spinner.StopWithSuccess()

	if svc.jsonOutput {
		PrintJsonAndExit(testSuiteImportResponseV3)
	}

	PrintViewTestSuite(svc.cmd, []TestSuiteV3{testSuiteImportResponseV3.Success.TestSuite}, testSuiteImportResponseV3.Success.TestSuite.TestSuiteName)
	printLabelAndValue(GetServiceMessage(svc.cmd, MessageTypeDisplay, "", "testCasesAdded"), testSuiteImportResponseV3.Success.TestSuite.TestCaseCount)

	if testSuiteImportResponseV3.Failure.TestCases != nil {
		PrintError("\n" + GetServiceMessage(svc.cmd, MessageTypeDisplay, "", "importTSTestCaseFailed") + "\n")
		printErrorMessages(GetApiSubErrorMessagesForCommand(svc.cmd, testSuiteImportResponseV3.Failure.TestCases, "", Empty, Empty))
	}
	if testSuiteImportResponseV3.Failure.Variables != nil {
		PrintError("\n" + GetServiceMessage(svc.cmd, MessageTypeDisplay, "", "importTSVariableFailed") + "\n")
		printErrorMessages(GetApiSubErrorMessagesForCommand(svc.cmd, testSuiteImportResponseV3.Failure.Variables, "", Empty, Empty))
	}
}

func (svc Service) ManageTestSuites(testSuiteManage TestSuiteDetailsWithChildObjects) {

	spinner := NewSpinner(GetServiceMessage(svc.cmd, MessageTypeSpinner, "", "manageTestSuite"), !svc.jsonOutput).Start()
	testSuiteManageResponseV3, err := svc.api.ManageTestSuite(testSuiteManage, testSuiteManage.TestSuite.TestSuiteId)

	if err != nil {
		spinner.StopWithFailure()
		AbortForCommand(svc.cmd, err)
	}
	spinner.StopWithSuccess()

	if svc.jsonOutput {
		PrintJsonAndExit(testSuiteManageResponseV3)
	}
	PrintSuccess(GetServiceMessage(svc.cmd, MessageTypeDisplay, "", "manageTSSuccess")+"\n", testSuiteManageResponseV3.TestSuite.TestSuiteName)
}

func (svc Service) AddTestSuite(name, description, propertyName string, propVersion int, unlocked, stateful bool) *TestSuiteV3 {

	spinner := NewSpinner(GetServiceMessage(svc.cmd, MessageTypeSpinner, "", "addTestSuite"), !svc.jsonOutput).Start()

	testSuite := TestSuiteV3{
		TestSuiteName:        name,
		TestSuiteDescription: description,
		Locked:               !unlocked,
		Stateful:             stateful,
		Configs: AkamaiConfigs{
			PropertyManager: PropertyManager{
				PropertyName:    propertyName,
				PropertyVersion: propVersion,
			}},
	}

	testSuitesV3, err := svc.api.AddTestSuitesV3(testSuite)
	if err != nil {
		spinner.StopWithFailure()
		AbortForCommand(svc.cmd, err)
	}
	spinner.StopWithSuccess()

	if svc.jsonOutput {
		PrintJsonAndExit(testSuitesV3)
	}

	return testSuitesV3
}

func (svc Service) EditTestSuite(id, name, description, propertyName string, propVersion int, unlocked, stateful, removeProperty bool) *TestSuiteV3 {

	spinner := NewSpinner(GetServiceMessage(svc.cmd, MessageTypeSpinner, "", "editTestSuite"), !svc.jsonOutput).Start()
	getTestSuite, err := svc.api.GetTestSuiteV3(id)
	if err != nil {
		spinner.StopWithFailure()
		AbortForCommandWithSubResource(svc.cmd, err, Empty, Read)
	}

	var updatedTestSuite, isChanged = updateModifiedTestSuiteFields(*getTestSuite, name, description, propertyName, propVersion, unlocked, stateful, removeProperty)

	// Check if at least value changed to edit test suite.
	if !isChanged {
		spinner.StopWithFailure()
		NewValidator(svc.cmd, []byte{}).EditTestSuiteAllFlagCheck()
	}

	testSuitesV3, err := svc.api.EditTestSuitesV3(*updatedTestSuite, id)
	if err != nil {
		spinner.StopWithFailure()
		AbortForCommandWithSubResource(svc.cmd, err, Empty, Update)
	}

	spinner.StopWithSuccess()

	if svc.jsonOutput {
		PrintJsonAndExit(testSuitesV3)
	}

	return testSuitesV3
}

func (svc Service) RemoveTestSuiteById(id string) {

	spinner := NewSpinner(GetServiceMessage(svc.cmd, MessageTypeSpinner, "", "removeTestSuite"), !svc.jsonOutput).Start()
	err := svc.api.RemoveTestSuite(id)
	if err != nil {
		spinner.StopWithFailure()
		AbortForCommandWithSubResource(svc.cmd, err, Empty, Read)
	}

}

func (svc Service) RestoreTestSuiteById(id string) *TestSuiteV3 {

	spinner := NewSpinner(GetServiceMessage(svc.cmd, MessageTypeSpinner, "", "restoreTestSuite"), !svc.jsonOutput).Start()
	testSuitesV3, err := svc.api.RestoreTestSuite(id)
	if err != nil {
		spinner.StopWithFailure()
		AbortForCommandWithSubResource(svc.cmd, err, Empty, Read)
	}
	return testSuitesV3

}

func (svc Service) GetTestSuitesByIdOrName(id, name, subResource string, exactMatch, shouldMatchDescription bool) []TestSuiteV3 {

	spinner := NewSpinner(GetServiceMessage(svc.cmd, MessageTypeSpinner, "", "getTestSuite"), !svc.jsonOutput).Start()

	if id != "" {

		testSuite, err := svc.api.GetTestSuiteV3(id)
		if err != nil {
			spinner.StopWithFailure()
			AbortForCommandWithSubResource(svc.cmd, err, subResource, Read)
		}
		spinner.StopWithSuccess()

		return []TestSuiteV3{*testSuite}
	} else {

		testSuitesV3, err := svc.api.GetTestSuitesV3("", "", "")
		if err != nil {
			spinner.StopWithFailure()
			AbortForCommandWithSubResource(svc.cmd, err, subResource, Read)
		}
		spinner.StopWithSuccess()

		return filterTestSuitesByString(name, exactMatch, shouldMatchDescription, testSuitesV3)
	}
}

func (svc Service) GetSingleTestSuiteByIdOrName(id, name, subResource string) *TestSuiteV3 {

	testSuites := svc.GetTestSuitesByIdOrName(id, name, subResource, true, false)

	if len(testSuites) == 1 {
		return &testSuites[0]
	}

	return nil
}

func (svc Service) GetV3AssociatedTestCasesForTestSuite(testSuiteId int) ([]TestCase, bool) {

	spinner := NewSpinner(GetServiceMessage(svc.cmd, MessageTypeSpinner, "", "getTestCases"), !svc.jsonOutput).Start()
	associatedTestCases, err := svc.api.GetV3AssociatedTestCasesForTestSuite(testSuiteId)
	if err != nil {
		spinner.StopWithFailure()
		AbortForCommandWithSubResource(svc.cmd, err, Empty, Read)
	}

	spinner.StopWithSuccess()
	return associatedTestCases.TestCases, associatedTestCases.AllTestCasesIncluded
}

func (svc Service) GetTestSuiteWithChildObjects(testSuiteId int) TestSuiteDetailsWithChildObjects {

	spinner := NewSpinner(GetServiceMessage(svc.cmd, MessageTypeSpinner, "", "testSuiteWithChildObject"), !svc.jsonOutput).Start()
	testSuiteDetailsWithChildObjects, err := svc.api.GetTestSuitesWithChildObjects(testSuiteId)
	if err != nil {
		spinner.StopWithFailure()
		AbortForCommandWithSubResource(svc.cmd, err, Empty, Read)
	}

	spinner.StopWithSuccess()
	return *testSuiteDetailsWithChildObjects
}

func (svc Service) AddTestCaseToTestSuite(testSuiteId int, url, condition, ipVersion string, addHeader, modifyHeader, filterHeader []string) ([]TestCase, *CliError) {
	spinner := NewSpinner(GetServiceMessage(svc.cmd, MessageTypeSpinner, "", "addTestCase"), !svc.jsonOutput).Start()

	var testCase = constructTestCase(url, addHeader, modifyHeader, filterHeader, condition, ipVersion)
	log.Debugf("Add test case [%+v] with the test suite id [%d]\n", testCase, testSuiteId)

	testCases, err := svc.api.AddTestCaseToTestSuite(testSuiteId, []TestCase{testCase})
	if err != nil {
		spinner.StopWithFailure()
		AbortForCommand(svc.cmd, err)
	}

	if len(testCases.Failures) != 0 {
		spinner.StopWithFailure()
		return nil, CliErrorFromPulsarProblemObject(nil, testCases.Failures, 207, ApiErrorAddTestCasesToTestSuitPostCall)
	}

	spinner.StopWithSuccess()
	return testCases.Successes, nil
}

func (svc Service) RemoveTestCasesFromTestSuite(testSuiteId int, testCaseIds []int) (*BulkResponse, *CliError) {
	spinner := NewSpinner(GetServiceMessage(svc.cmd, MessageTypeSpinner, "", "removeTestCase"), !svc.jsonOutput).Start()
	log.Debugf("Removing test cases [%d] from the test suite id [%d]\n", testCaseIds, testSuiteId)

	removeTestCaseResponse, err := svc.api.RemoveTestCasesFromTestSuite(testSuiteId, testCaseIds)
	if err != nil {
		spinner.StopWithFailure()
		AbortForCommand(svc.cmd, err)
	}

	if len(removeTestCaseResponse.Failures) != 0 {
		spinner.StopWithFailure()
		return nil, CliErrorFromPulsarProblemObject(nil, removeTestCaseResponse.Failures, 207, ApiErrorRemoveTestCasesPostCall)
	}

	spinner.StopWithSuccess()
	return removeTestCaseResponse, nil
}

func (svc Service) GetTestSuiteOrTestSuiteDetailsWithChildObjects(id, name string) ([]TestSuiteV3, TestSuiteDetailsWithChildObjects) {

	var testSuiteDetailsWithChildObjects = TestSuiteDetailsWithChildObjects{}
	if name != "" {
		testSuites := svc.GetTestSuitesByIdOrName(id, name, Empty, false, false)
		if len(testSuites) > 1 {
			return testSuites, testSuiteDetailsWithChildObjects
		}
		testSuiteDetailsWithChildObjects = svc.GetTestSuiteWithChildObjects(testSuites[0].TestSuiteId)
	} else {
		testSuiteId, _ := strconv.Atoi(id)
		testSuiteDetailsWithChildObjects = svc.GetTestSuiteWithChildObjects(testSuiteId)
	}
	return nil, testSuiteDetailsWithChildObjects
}

func (svc Service) ViewTestSuite(id string, name string, groupBy string) {

	var testSuites = []TestSuiteV3{}
	var testSuiteDetailsWithChildObjects = TestSuiteDetailsWithChildObjects{}
	allTestCasesIncluded := false
	testSuites, testSuiteDetailsWithChildObjects = svc.GetTestSuiteOrTestSuiteDetailsWithChildObjects(id, name)

	if len(testSuites) > 1 {
		if svc.jsonOutput {
			PrintJsonAndExit(testSuites)
		}
		PrintViewTestSuite(svc.cmd, testSuites, name)
	} else {
		if svc.jsonOutput {
			PrintJsonAndExit(testSuiteDetailsWithChildObjects)
		} else {
			PrintViewTestSuite(svc.cmd, []TestSuiteV3{testSuiteDetailsWithChildObjects.TestSuite}, name)
			if testSuiteDetailsWithChildObjects.TestSuite.TestCaseCount == len(testSuiteDetailsWithChildObjects.TestSuite.TestCases) {
				allTestCasesIncluded = true
			}
			PrintTestCases(svc.cmd, testSuiteDetailsWithChildObjects.TestSuite.TestCases, allTestCasesIncluded, groupBy)
		}
	}
}

func (svc Service) RemoveTestSuiteByIdOrName(id string, name string) {

	testSuites := svc.GetTestSuitesByIdOrName(id, name, Empty, false, false)

	if len(testSuites) == 1 {
		testSuiteId := strconv.Itoa(testSuites[0].TestSuiteId)
		svc.RemoveTestSuiteById(testSuiteId)
		if !svc.jsonOutput {
			PrintSuccess("\n" + GetServiceMessage(svc.cmd, MessageTypeDisplay, "", "removeTSSuccess"))
			PrintSuccessInBold(GetServiceMessage(svc.cmd, MessageTypeDisplay, "", "removeTSSuccessInBold"), testSuiteId)
			PrintSuccess(GetServiceMessage(svc.cmd, MessageTypeDisplay, "", "removeTSSuccessEnd") + "\n")
			PrintTestSuiteForRemovedTestSuite(testSuites[0])
		}
	} else {
		if svc.jsonOutput {
			PrintJsonAndExit(testSuites)
		}
		PrintViewTestSuite(svc.cmd, testSuites, name)
	}

}

func (svc Service) RestoreTestSuiteByIdOrName(id string, name string) {

	testSuites := svc.GetTestSuitesByIdOrName(id, name, Empty, false, false)
	if len(testSuites) == 1 {
		testSuiteId := strconv.Itoa(testSuites[0].TestSuiteId)
		testSuite := svc.RestoreTestSuiteById(testSuiteId)
		if svc.jsonOutput {
			PrintJsonAndExit(testSuite)
		}
		PrintSuccess("\n" + GetServiceMessage(svc.cmd, MessageTypeDisplay, "", "restoreTSSuccess") + "\n")
		PrintTestSuiteForRemovedTestSuite(testSuites[0])
	} else {
		if svc.jsonOutput {
			PrintJsonAndExit(testSuites)
		}
		PrintViewTestSuite(svc.cmd, testSuites, name)
	}
}

func (svc Service) AddTestCaseWithTestSuite(testSuites []TestSuiteV3, searchedName, url, condition, ipVersion string, addHeader, modifyHeader, filterHeader []string) {

	if len(testSuites) == 1 {
		testCases, err := svc.AddTestCaseToTestSuite(testSuites[0].TestSuiteId, url, condition, ipVersion, addHeader, modifyHeader, filterHeader)

		if testCases != nil {
			if svc.jsonOutput {
				PrintJsonAndExit(testCases)
			}

			PrintSuccess(fmt.Sprintf(GetServiceMessage(svc.cmd, MessageTypeDisplay, "", "addTestCaseSuccess")+"\n\n", testSuites[0].TestSuiteName))
		} else {
			PrintError(fmt.Sprintf(GetServiceMessage(svc.cmd, MessageTypeDisplay, "", "addTestCaseFail")+"\n\n", testSuites[0].TestSuiteName))
			AbortForCommand(svc.cmd, err)
		}
	} else {
		PrintError(fmt.Sprintf(GetServiceMessage(svc.cmd, MessageTypeDisplay, "", "addTestCaseNoTestSuite")+"\n\n", searchedName))
	}
}

func (svc Service) RemoveTestCaseFromTestSuiteUsingOrderNumber(testSuite *TestSuiteV3, testCases []TestCase, orderNumber string) {

	if testSuite != nil {

		orderNum, _ := strconv.Atoi(orderNumber)
		testCase := filterTestCaseUsingOrderNumber(testCases, orderNum)

		if testCase != nil {
			removeTestCaseResponse, err := svc.RemoveTestCasesFromTestSuite(testSuite.TestSuiteId, []int{testCase.TestCaseId})
			if err != nil {
				PrintError(fmt.Sprintf(GetMessageForKey(svc.cmd, "failed")+"\n\n", testSuite.TestSuiteName))
				AbortForCommand(svc.cmd, err)
			} else {
				if svc.jsonOutput {
					PrintJsonAndExit(removeTestCaseResponse)
				}

				PrintSuccess(fmt.Sprintf(GetMessageForKey(svc.cmd, "success")+"\n\n", testSuite.TestSuiteName))
			}
		} else {
			PrintError(fmt.Sprintf(GetMessageForKey(svc.cmd, "notPresent")+"\n\n", orderNumber))
		}
	}
}

func (svc Service) RunTest(runTestUsing, testSuiteId, testSuiteName, propertyName, propVersion,
	url, condition, ipVersion, targetEnvironment string, addHeader, modifyHeader, filterHeader []string,
	testRunRequestFromJson TestRun) {

	var testRun *TestRun

	switch runTestUsing {
	case RunTestUsingTestSuiteId:
		testSuiteId, _ := strconv.Atoi(testSuiteId)
		testRun = svc.runTestSuiteWithId(testSuiteId, targetEnvironment, runTestUsing)
	case RunTestUsingTestSuiteName:
		testRun = svc.findAndRunTestSuite(testSuiteName, targetEnvironment, runTestUsing)
	case RunTestUsingPropertyVersion:
		testRun = svc.findAndRunPropertyVersion(propertyName, propVersion, targetEnvironment, runTestUsing)
	case RunTestUsingSingleTestCase:
		testRun = svc.runTestCase(url, condition, ipVersion, addHeader, modifyHeader, filterHeader, targetEnvironment, runTestUsing)
	case RunTestUsingJsonInput:
		testRun = svc.startTestRun(testRunRequestFromJson, runTestUsing)
	}

	// Fetch test run context
	// TODO: Add retries for all API calls
	var testRunContext *TestRunContext
	go func() {
		trContext, err := svc.api.GetTestRunContext(testRun.TestRunId)
		if err != nil {
			AbortForCommandWithSubResource(svc.cmd, err, TestRunResource, Read)
			return
		}
		testRunContext = trContext
	}()

	// Wait for test run results
	testResult, err := svc.waitForTestRunCompletion(testRun.TestRunId, runTestUsing)
	if err != nil {
		AbortForCommandWithSubResource(svc.cmd, err, TestRunResource, Read)
	}

	// Print test run result
	PrintTestResult(svc.cmd, testResult, testRunContext)
}

func (svc Service) findAndRunTestSuite(testSuiteName, targetEnvironment string, runTestUsing string) *TestRun {
	testSuite := svc.GetSingleTestSuiteByIdOrName("", testSuiteName, TestSuiteResource)
	if testSuite == nil {
		AbortWithExitCode(fmt.Sprintf(GetMessageForKey(svc.cmd, "testSuiteNameNotFound")+"\n", testSuiteName), ExitStatusCode0)
	}

	return svc.runTestSuiteWithId(testSuite.TestSuiteId, targetEnvironment, runTestUsing)
}

func (svc Service) runTestSuiteWithId(testSuiteId int, targetEnvironment string, runTestUsing string) *TestRun {
	testCaseIds := svc.getAssociatedTestCaseIdsForTestSuiteId(testSuiteId, runTestUsing)

	if len(testCaseIds) == 0 {
		AbortWithExitCode(GetMessageForKey(svc.cmd, "testCasesNotFound")+"\n", ExitStatusCode0)
	}

	return svc.runTestSuite(testSuiteId, testCaseIds, targetEnvironment, runTestUsing)
}

func (svc Service) runTestSuite(testSuiteId int, testCaseIds []int, environment string, runTestUsing string) *TestRun {

	testRun := TestRun{
		SendEmailOnCompletion: false,
		TargetEnvironment:     environment,
		Functional: FunctionalTestRun{
			TestSuiteExecutions: []TestSuiteExecution{
				{
					TestSuiteId:         testSuiteId,
					TestCaseExecutionV2: getTestCaseExecutions(testCaseIds),
				},
			},
		},
	}

	return svc.startTestRun(testRun, runTestUsing)
}

func (svc Service) findAndRunPropertyVersion(propertyName, propVersion, targetEnvironment string, runTestUsing string) *TestRun {

	config := svc.findPropertyVersion(propertyName, propVersion, runTestUsing)

	testSuites := svc.findTestSuitesForPropertyVersion(propertyName, propVersion, runTestUsing)
	if len(testSuites) == 0 {
		AbortWithExitCode(GetMessageForKey(svc.cmd, PropertyVersionTestSuitesNotFound)+"\n", ExitStatusCode0)
	}

	var testSuiteIds []int
	for _, testSuite := range testSuites {
		testSuiteIds = append(testSuiteIds, testSuite.TestSuiteId)
	}

	testSuiteExecutionsChan := make(chan TestSuiteExecution, len(testSuiteIds))

	var wg sync.WaitGroup
	wg.Add(len(testSuiteIds))
	for _, testSuiteId := range testSuiteIds {
		go func(testSuiteId int) {
			defer wg.Done()
			testCaseIds := svc.getAssociatedTestCaseIdsForTestSuiteId(testSuiteId, runTestUsing)
			testCaseExecutions := getTestCaseExecutions(testCaseIds)
			testSuiteExecution := TestSuiteExecution{
				TestSuiteId:         testSuiteId,
				TestCaseExecutionV2: testCaseExecutions,
			}

			// Ignore test suites that don't have any test cases
			if len(testCaseExecutions) > 0 {
				testSuiteExecutionsChan <- testSuiteExecution
			}
		}(testSuiteId) // Fetch test suites' associated test cases in parallel
	}

	wg.Wait()
	close(testSuiteExecutionsChan)

	var testSuiteExecutions []TestSuiteExecution
	for testSuiteExecution := range testSuiteExecutionsChan {
		testSuiteExecutions = append(testSuiteExecutions, testSuiteExecution)
	}

	if len(testSuiteExecutions) == 0 {
		AbortWithExitCode(GetMessageForKey(svc.cmd, "noTestCases")+"\n", ExitStatusCode0)
	}

	return svc.runConfigVersion(config.ConfigVersionId, testSuiteExecutions, targetEnvironment, runTestUsing)
}

func (svc Service) findPropertyVersion(propertyName string, propVersion string, runTestUsing string) *ConfigVersion {

	spinner := NewSpinner(GetServiceMessage(svc.cmd, MessageTypeTestCmdSpinner, runTestUsing, "findPropertyVersion"), !svc.jsonOutput).Start()
	propertyVersions, err := svc.api.GetConfigVersions()
	if err != nil {
		spinner.StopWithFailure()
		AbortForCommandWithSubResource(svc.cmd, err, PropertyVersionsResource, Read)
	}

	cleanedLowerCasePropertyName := strings.TrimSpace(strings.ToLower(propertyName))
	versionNumber, _ := strconv.Atoi(propVersion)
	for _, propertyVersion := range propertyVersions {
		if cleanedLowerCasePropertyName == strings.TrimSpace(strings.ToLower(propertyVersion.PropertyName)) && versionNumber == propertyVersion.PropertyVersion {
			spinner.StopWithSuccess()
			return &propertyVersion
		}
	}

	spinner.StopWithFailure()
	AbortWithExitCode(GetMessageForKey(svc.cmd, PropertyVersionNotFound)+"\n", ExitStatusCode0)
	return nil
}

func (svc Service) findTestSuitesForPropertyVersion(propertyName string, propVersion string, runTestUsing string) []TestSuiteV3 {
	spinner := NewSpinner(GetServiceMessage(svc.cmd, MessageTypeTestCmdSpinner, runTestUsing, "findTestSuite"), !svc.jsonOutput).Start()
	testSuites, err := svc.api.GetTestSuitesV3(propertyName, propVersion, "")
	if err != nil {
		spinner.StopWithFailure()
		AbortForCommandWithSubResource(svc.cmd, err, TestSuiteResource, Read)
	}

	spinner.StopWithSuccess()
	return testSuites
}

func (svc Service) runConfigVersion(configVersionId int, testSuiteExecutions []TestSuiteExecution, environment string, runTestUsing string) *TestRun {

	testRun := TestRun{
		SendEmailOnCompletion: false,
		TargetEnvironment:     environment,
		Functional: FunctionalTestRun{
			ConfigVersionExecutions: []ConfigVersionExecution{
				{
					ConfigVersionId:     configVersionId,
					TestSuiteExecutions: testSuiteExecutions,
				},
			},
		},
	}

	return svc.startTestRun(testRun, runTestUsing)
}

func (svc Service) runTestCase(url, condition, ipVersion string, addHeader, modifyHeader,
	filterHeader []string, targetEnvironment string, runTestUsing string) *TestRun {

	var clientProfile = ClientProfile{IpVersion: Ipv4}
	if strings.ToLower(ipVersion) == "v6" {
		clientProfile = ClientProfile{IpVersion: Ipv6}
	}

	testRun := TestRun{
		SendEmailOnCompletion: false,
		TargetEnvironment:     targetEnvironment,
		Functional: FunctionalTestRun{
			TestCaseExecutionV3: TestCaseExecutionV3{
				TestRequest: TestRequest{
					TestRequestUrl: url,
					RequestHeaders: getRequestHeaders(addHeader, modifyHeader, filterHeader),
				},
				Condition: Condition{
					ConditionExpression: strings.TrimSpace(condition),
				},
				ClientProfile: clientProfile,
			},
		},
	}

	return svc.startTestRun(testRun, runTestUsing)
}

func (svc Service) getAssociatedTestCaseIdsForTestSuiteId(testSuiteId int, runTestUsing string) []int {
	spinner := NewSpinner(GetServiceMessage(svc.cmd, MessageTypeTestCmdSpinner, runTestUsing, "findTestCases"), !svc.jsonOutput).Start()
	testCases, err := svc.api.GetV3AssociatedTestCasesForTestSuite(testSuiteId)
	if err != nil {
		spinner.StopWithFailure()
		AbortForCommandWithSubResource(svc.cmd, err, TestCaseResource, Read)
	}

	spinner.StopWithSuccess()

	var testCaseIds []int
	for _, testCase := range testCases.TestCases {
		testCaseIds = append(testCaseIds, testCase.TestCaseId)
	}

	return testCaseIds
}

func (svc Service) startTestRun(testRun TestRun, runTestUsing string) *TestRun {
	spinner := NewSpinner(GetServiceMessage(svc.cmd, MessageTypeTestCmdSpinner, runTestUsing, "startTestRun"), !svc.jsonOutput).Start()
	createdTestRun, err := svc.api.SubmitTestRun(testRun)
	if err != nil {
		spinner.StopWithFailure()
		AbortForCommandWithSubResource(svc.cmd, err, TestRunResource, Create)
		return nil
	}

	spinner.StopWithSuccess()

	if svc.jsonOutput {
		PrintJsonAndExit(createdTestRun)
	}

	PrintSuccess(GetServiceMessage(svc.cmd, MessageTypeDisplay, "", "testRunStart") + "\n\n")
	return createdTestRun
}

func (svc Service) waitForTestRunCompletion(testRunId int, runTestUsing string) (*TestRun, *CliError) {
	spinner := NewSpinner(GetServiceMessage(svc.cmd, MessageTypeTestCmdSpinner, runTestUsing, "runTests"), !svc.jsonOutput).Start()
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

		if testRun != nil && testRun.Status != InProgress {
			spinner.StopWithSuccess()
			ticker.Stop()

			return testRun, nil
		}
	}

	return nil, CliErrorWithMessage(CliErrorMessageTestRunStatus)
}

func getTestCaseExecutions(testCaseIds []int) []TestCaseExecutionV2 {
	var testCaseExecutions []TestCaseExecutionV2

	for _, testCaseId := range testCaseIds {
		testCaseExecutions = append(testCaseExecutions, TestCaseExecutionV2{
			TestCaseId: testCaseId,
		})
	}

	return testCaseExecutions
}

func constructTestCase(url string, addHeader, modifyHeader, filterHeader []string, condition, ipVersion string) TestCase {
	var (
		testCase        TestCase
		testRequest     TestRequest
		conditionObject Condition
		clientProfileId int
		clientProfile   ClientProfile
	)

	testRequest = TestRequest{
		TestRequestUrl: url,
		RequestHeaders: getRequestHeaders(addHeader, modifyHeader, filterHeader),
	}

	conditionObject = Condition{
		ConditionExpression: strings.TrimSpace(condition),
	}

	clientProfileId = 2 // Use default IPv4
	clientProfile = ClientProfile{IpVersion: Ipv4}
	if ipVersion == "v6" || ipVersion == "V6" {
		clientProfileId = 1
		clientProfile = ClientProfile{IpVersion: Ipv6}
	}

	testCase = TestCase{
		TestRequest:     testRequest,
		Condition:       conditionObject,
		ClientProfileId: clientProfileId,
		ClientProfile:   clientProfile,
	}

	return testCase
}

func filterTestCaseUsingOrderNumber(testCases []TestCase, orderNumber int) *TestCase {

	for _, testCase := range testCases {
		if testCase.Order == orderNumber {
			return &testCase
		}
	}

	return nil
}

func updateModifiedTestSuiteFields(testSuiteV3 TestSuiteV3, name, description, propertyName string, propVersion int, unlocked, stateful, removeProperty bool) (*TestSuiteV3, bool) {

	var isChanged = false
	if name != "" && testSuiteV3.TestSuiteName != name {
		testSuiteV3.TestSuiteName = name
		isChanged = true
	}

	if description != "" && testSuiteV3.TestSuiteDescription != description {
		testSuiteV3.TestSuiteDescription = description
		isChanged = true
	}

	if testSuiteV3.Locked != !unlocked {
		testSuiteV3.Locked = !unlocked
		isChanged = true
	}

	if testSuiteV3.Stateful != stateful {
		testSuiteV3.Stateful = stateful
		isChanged = true
	}

	if testSuiteV3.Configs.PropertyManager.PropertyName != "" && removeProperty {
		testSuiteV3.Configs = AkamaiConfigs{}
		isChanged = true
	}

	if propertyName != Empty && (testSuiteV3.Configs.PropertyManager.PropertyName != propertyName ||
		testSuiteV3.Configs.PropertyManager.PropertyVersion != propVersion) {

		testSuiteV3.Configs.PropertyManager.PropertyName = propertyName
		testSuiteV3.Configs.PropertyManager.PropertyVersion = propVersion
		isChanged = true
	}

	return &testSuiteV3, isChanged
}

// Returns testSuites containing searchString in either name or description
func filterTestSuitesByString(searchString string, exactMatch, shouldMatchDescription bool, testSuites []TestSuiteV3) []TestSuiteV3 {
	var filteredItems []TestSuiteV3

	for _, testSuite := range testSuites {
		if exactMatch {
			if strings.EqualFold(testSuite.TestSuiteName, searchString) ||
				(shouldMatchDescription && strings.EqualFold(testSuite.TestSuiteDescription, searchString)) {
				filteredItems = append(filteredItems, testSuite)
			}
		} else {
			if ContainsIgnoreCase(testSuite.TestSuiteName, searchString) ||
				(shouldMatchDescription && ContainsIgnoreCase(testSuite.TestSuiteDescription, searchString)) {
				filteredItems = append(filteredItems, testSuite)
			}
		}
	}

	return filteredItems
}

func getRequestHeaders(headerAdd []string, headerModify []string, headerFilter []string) []RequestHeader {
	var requestHeaders []RequestHeader

	for _, header := range headerAdd {
		headerComponents := strings.Split(header, ":")
		requestHeaders = append(requestHeaders, RequestHeader{
			HeaderName:   strings.TrimSpace(headerComponents[0]),
			HeaderValue:  strings.TrimSpace(headerComponents[1]),
			HeaderAction: Add,
		})
	}

	for _, header := range headerModify {
		headerComponents := strings.Split(header, ":")
		requestHeaders = append(requestHeaders, RequestHeader{
			HeaderName:   strings.TrimSpace(headerComponents[0]),
			HeaderValue:  strings.TrimSpace(headerComponents[1]),
			HeaderAction: Modify,
		})
	}

	for _, header := range headerFilter {
		requestHeaders = append(requestHeaders, RequestHeader{
			HeaderName:   strings.TrimSpace(header),
			HeaderAction: Filter,
		})
	}

	return requestHeaders
}

func (svc Service) GetConditionTemplate() *ConditionTemplate {

	spinner := NewSpinner(GetServiceMessage(svc.cmd, MessageTypeSpinner, "", "getConditionTemplate"), !svc.jsonOutput).Start()
	condTemplate, err := svc.api.GetConditionTemplate()
	if err != nil {
		spinner.StopWithFailure()
		AbortForCommand(svc.cmd, err)
	}
	spinner.StopWithSuccess()

	if svc.jsonOutput {
		PrintJsonAndExit(condTemplate)
	}

	return condTemplate
}

func (svc Service) GenerateTestSuite(propertyName string, propVersion int, urls []string,
	defaultTestSuiteRequest DefaultTestSuiteRequest, isJsonInputPresent bool) {
	spinner := NewSpinner(GetServiceMessage(svc.cmd, MessageTypeSpinner, "", "generateDefaultTestSuite"), !svc.jsonOutput).Start()

	defaultTsReq := DefaultTestSuiteRequest{}

	// form request
	if isJsonInputPresent {
		defaultTsReq = defaultTestSuiteRequest
	} else {
		defaultTsReq = DefaultTestSuiteRequest{
			TestRequestUrl: urls,
			Configs: AkamaiConfigs{
				PropertyManager: PropertyManager{
					PropertyName:    propertyName,
					PropertyVersion: propVersion,
				}},
		}
	}

	// get default ts
	generatedTs, err := svc.api.GenerateTestSuite(defaultTsReq)

	if err != nil {
		spinner.StopWithFailure()
		AbortForCommand(svc.cmd, err)
	}
	spinner.StopWithSuccess()

	// print output
	if !svc.jsonOutput {
		PrintSuccess(fmt.Sprintf(GetServiceMessage(svc.cmd, MessageTypeDisplay, "", "generateDefaultTSSuccess")+"\n\n", propertyName, propVersion))
	}
	PrintJsonAndExit(generatedTs)
}
