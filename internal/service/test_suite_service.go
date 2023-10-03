package service

import (
	"fmt"
	"github.com/akamai/cli-test-center/internal/api"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/print"
	"github.com/akamai/cli-test-center/internal/util"
	"github.com/akamai/cli-test-center/internal/validator"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

type Service struct {
	api        api.ApiClient
	cmd        *cobra.Command
	jsonOutput bool
}

func NewService(api api.ApiClient, cmd *cobra.Command, jsonOutput bool) *Service {
	return &Service{api, cmd, jsonOutput}
}

func (svc Service) GetTestSuitesAndPrint(cmd *cobra.Command, propertyId, propertyName, propVersion, user, searchString string) {

	testSuites := svc.GetTestSuites(propertyId, propertyName, propVersion, user, searchString)
	print.PrintTestSuitesTable(cmd, testSuites)
}

func (svc Service) GetTestSuites(propertyId, propertyName, propVersion, user, searchString string) []model.TestSuite {

	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.GetTestSuite), !svc.jsonOutput).Start()
	testSuitesV3, err := svc.api.GetTestSuites(propertyId, propertyName, propVersion, user, true)
	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}
	spinner.StopWithSuccess()

	filteredTestSuites := filterTestSuitesByString(searchString, false, true, testSuitesV3)

	if svc.jsonOutput {
		util.PrintJsonAndExit(filteredTestSuites)
	}

	return filteredTestSuites
}

func (svc Service) ImportTestSuites(testSuiteImport model.TestSuite) {

	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.ImportTestSuite), !svc.jsonOutput).Start()
	testSuiteImportResponseV3, err := svc.api.ImportTestSuite(testSuiteImport)

	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}
	spinner.StopWithSuccess()

	if svc.jsonOutput {
		util.PrintJsonAndExit(testSuiteImportResponseV3)
	}

	print.PrintTestSuitesResult(svc.cmd, []model.TestSuite{testSuiteImportResponseV3.Success}, testSuiteImportResponseV3.Success.TestSuiteName)
	util.PrintLabelAndValue(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.TestCasesAdded), testSuiteImportResponseV3.Success.ExecutableTestCaseCount)

	if testSuiteImportResponseV3.Failure.TestCases != nil {
		util.PrintError("\n" + util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.ImportTSTestCaseFailed) + "\n")
		util.PrintErrorMessages(util.GetApiSubErrorMessagesForCommand(svc.cmd, testSuiteImportResponseV3.Failure.TestCases, internalconstant.Empty, internalconstant.Empty, internalconstant.Empty))
	}
	if testSuiteImportResponseV3.Failure.Variables != nil {
		util.PrintError("\n" + util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.ImportTSVariableFailed) + "\n")
		util.PrintErrorMessages(util.GetApiSubErrorMessagesForCommand(svc.cmd, testSuiteImportResponseV3.Failure.Variables, internalconstant.Empty, internalconstant.Empty, internalconstant.Empty))
	}
}

func (svc Service) ManageTestSuites(testSuiteManage model.TestSuite) {

	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.ManageTestSuite), !svc.jsonOutput).Start()
	testSuiteManageResponseV3, err := svc.api.ManageTestSuite(testSuiteManage, testSuiteManage.TestSuiteId)

	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}
	spinner.StopWithSuccess()

	if svc.jsonOutput {
		util.PrintJsonAndExit(testSuiteManageResponseV3)
	}
	util.PrintSuccess(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.ManageTSSuccess)+"\n", testSuiteManageResponseV3.TestSuiteName)
}

func (svc Service) AddTestSuiteAndPrint(cmd *cobra.Command, name, description, propertyId, propertyName, propVersion string, unlocked, stateful, isStandardInputAvailable bool, jsonInputTestSuite model.TestSuite) {

	testSuite := svc.AddTestSuite(name, description, propertyId, propertyName, propVersion, unlocked, stateful, isStandardInputAvailable, jsonInputTestSuite)

	util.PrintSuccess(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.AddTestSuiteSuccess) + "\n")
	print.PrintTestSuite(cmd, *testSuite)
}

func (svc Service) AddTestSuite(name, description, propertyId, propertyName, propVersion string, unlocked, stateful, isStandardInputAvailable bool, jsonInputTestSuite model.TestSuite) *model.TestSuite {

	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.AddTestSuite), !svc.jsonOutput).Start()

	testSuite := model.TestSuite{
		TestSuiteName:        name,
		TestSuiteDescription: description,
		IsLocked:             !unlocked,
		IsStateful:           stateful,
		Configs: model.AkamaiConfigs{
			PropertyManager: model.PropertyManager{
				PropertyId:      util.GetConvertedInteger(propertyId),
				PropertyName:    propertyName,
				PropertyVersion: util.GetConvertedInteger(propVersion),
			}},
	}

	// assign json input provided by user
	if isStandardInputAvailable {
		testSuite = jsonInputTestSuite
	}

	testSuitesV3, err := svc.api.AddTestSuitesV3(testSuite)
	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}
	spinner.StopWithSuccess()

	if svc.jsonOutput {
		util.PrintJsonAndExit(testSuitesV3)
	}

	return testSuitesV3
}

func (svc Service) EditTestSuiteAndPrint(cmd *cobra.Command, id, name, description, propertyId, propertyName, propVersion string,
	unlocked, stateful, removeProperty, locked, stateless, isStandardInputAvailable bool, jsonInputTestSuite model.TestSuite) {
	versionNumber, _ := strconv.Atoi(propVersion)
	testSuite := svc.EditTestSuite(id, name, description, propertyId, propertyName, versionNumber, unlocked, stateful, removeProperty, locked, stateless, isStandardInputAvailable, jsonInputTestSuite)

	util.PrintSuccess(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.EditTSSuccess) + "\n")
	print.PrintTestSuite(cmd, *testSuite)
}

func (svc Service) EditTestSuite(id, name, description, propertyId, propertyName string, propVersion int, unlocked, stateful,
	removeProperty, locked, stateless, isStandardInputAvailable bool, jsonInputTestSuite model.TestSuite) *model.TestSuite {

	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.EditTestSuite), !svc.jsonOutput).Start()

	var updatedTestSuite *model.TestSuite
	var isChanged = false
	if !isStandardInputAvailable {

		getTestSuite, err := svc.api.GetTestSuiteV3(id)
		if err != nil {
			spinner.StopWithFailure()
			util.AbortForCommandWithSubResource(svc.cmd, err, internalconstant.Empty, internalconstant.Read)
		}

		updatedTestSuite, isChanged = updateModifiedTestSuiteFields(*getTestSuite, name, description, propertyId, propertyName, propVersion, unlocked, stateful, removeProperty, locked, stateless)

		// Check if at least value changed to edit test suite.
		if !isChanged {
			spinner.StopWithFailure()
			validator.NewValidator(svc.cmd, []byte{}).EditTestSuiteAllFlagCheck()
		}
	} else {
		updatedTestSuite = &jsonInputTestSuite
		id = strconv.Itoa(jsonInputTestSuite.TestSuiteId)
	}

	testSuitesV3, err := svc.api.EditTestSuitesV3(*updatedTestSuite, id)
	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommandWithSubResource(svc.cmd, err, internalconstant.Empty, internalconstant.Update)
	}

	spinner.StopWithSuccess()

	if svc.jsonOutput {
		util.PrintJsonAndExit(testSuitesV3)
	}

	return testSuitesV3
}

func (svc Service) RemoveTestSuiteById(id string) {

	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.RemoveTestSuite), !svc.jsonOutput).Start()
	err := svc.api.RemoveTestSuite(id)
	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommandWithSubResource(svc.cmd, err, internalconstant.Empty, internalconstant.Read)
	}

}

func (svc Service) RestoreTestSuiteById(id string) *model.TestSuite {

	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.RestoreTestSuite), !svc.jsonOutput).Start()
	testSuitesV3, err := svc.api.RestoreTestSuite(id)
	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommandWithSubResource(svc.cmd, err, internalconstant.Empty, internalconstant.Read)
	}
	return testSuitesV3

}

func (svc Service) GetTestSuitesByIdOrName(id, name, subResource string, exactMatch, shouldMatchDescription, includeDeleted bool) []model.TestSuite {

	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.GetTestSuite), !svc.jsonOutput).Start()

	if id != internalconstant.Empty {

		testSuite, err := svc.api.GetTestSuiteV3(id)
		if err != nil {
			spinner.StopWithFailure()
			util.AbortForCommandWithSubResource(svc.cmd, err, subResource, internalconstant.Read)
		}
		spinner.StopWithSuccess()

		return []model.TestSuite{*testSuite}
	} else {

		testSuitesV3, err := svc.api.GetTestSuites(internalconstant.Empty, internalconstant.Empty, internalconstant.Empty, internalconstant.Empty, includeDeleted)
		if err != nil {
			spinner.StopWithFailure()
			util.AbortForCommandWithSubResource(svc.cmd, err, subResource, internalconstant.Read)
		}
		spinner.StopWithSuccess()

		return filterTestSuitesByString(name, exactMatch, shouldMatchDescription, testSuitesV3)
	}
}

func (svc Service) GetSingleTestSuiteByIdOrName(id, name, subResource string, includeDeleted bool) *model.TestSuite {

	testSuites := svc.GetTestSuitesByIdOrName(id, name, subResource, true, false, includeDeleted)

	if len(testSuites) == 1 {
		return &testSuites[0]
	}

	return nil
}

func (svc Service) GetTestSuiteWithChildObjects(testSuiteId int, resolveVariables bool) model.TestSuite {

	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.TestSuiteWithChildObject), !svc.jsonOutput).Start()
	testSuiteDetailsWithChildObjects, err := svc.api.GetTestSuitesWithChildObjects(testSuiteId, resolveVariables)
	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommandWithSubResource(svc.cmd, err, internalconstant.Empty, internalconstant.Read)
	}

	spinner.StopWithSuccess()
	return *testSuiteDetailsWithChildObjects
}
func (svc Service) RemoveTestCasesFromTestSuite(testSuiteId int, testCaseIds []int) (*model.BulkResponse, *model.CliError) {
	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.RemoveTestCase), !svc.jsonOutput).Start()
	log.Debugf("Removing test cases [%d] from the test suite id [%d]\n", testCaseIds, testSuiteId)

	removeTestCaseResponse, err := svc.api.RemoveTestCasesFromTestSuite(testSuiteId, testCaseIds)
	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}

	if len(removeTestCaseResponse.Failures) != 0 {
		spinner.StopWithFailure()
		return nil, model.CliErrorFromPulsarProblemObject(nil, removeTestCaseResponse.Failures, 207, externalconstant.ApiErrorRemoveTestCasesPostCall)
	}

	spinner.StopWithSuccess()
	return removeTestCaseResponse, nil
}

func (svc Service) GetTestSuiteOrTestSuiteDetailsWithChildObjects(id, name string, resolveVariables bool) ([]model.TestSuite, *model.TestSuite) {
	if name != internalconstant.Empty {
		testSuites := svc.GetTestSuitesByIdOrName(id, name, internalconstant.Empty, false, false, true)
		if len(testSuites) == 1 {
			var testSuiteDetailsWithChildObjects = svc.GetTestSuiteWithChildObjects(testSuites[0].TestSuiteId, resolveVariables)
			return nil, &testSuiteDetailsWithChildObjects
		}
		return testSuites, nil
	} else {
		testSuiteId, _ := strconv.Atoi(id)
		var testSuiteDetailsWithChildObjects = svc.GetTestSuiteWithChildObjects(testSuiteId, resolveVariables)
		return nil, &testSuiteDetailsWithChildObjects
	}
}

func (svc Service) GetTestSuiteAndPrint(id, name string) {

	testSuites := svc.GetTestSuitesByIdOrName(id, name, internalconstant.Empty, false, false, false)

	// print json output if flag is passed
	if svc.jsonOutput {
		if len(testSuites) == 1 {
			util.PrintJsonAndExit(testSuites[0])
		} else {
			util.PrintJsonAndExit(testSuites)
		}
	}

	// valid test suite if returned 1 else either no test suite found or multiple test suites found
	print.PrintTestSuitesResult(svc.cmd, testSuites, name)

}

func (svc Service) GetTestSuiteWithChildObjectsAndPrint(id string, name string, groupBy string, resolveVariables bool) {

	var testSuites []model.TestSuite
	var testSuiteDetailsWithChildObjects *model.TestSuite
	areAllTestCasesIncluded := false
	testSuites, testSuiteDetailsWithChildObjects = svc.GetTestSuiteOrTestSuiteDetailsWithChildObjects(id, name, resolveVariables)

	if testSuiteDetailsWithChildObjects != nil {
		if svc.jsonOutput {
			util.PrintJsonAndExit(testSuiteDetailsWithChildObjects)
		} else {
			print.PrintTestSuitesResult(svc.cmd, []model.TestSuite{*testSuiteDetailsWithChildObjects}, name)
			if testSuiteDetailsWithChildObjects.ExecutableTestCaseCount == len(testSuiteDetailsWithChildObjects.TestCases) {
				areAllTestCasesIncluded = true
			}
			print.PrintTestCases(svc.cmd, testSuiteDetailsWithChildObjects.TestCases, areAllTestCasesIncluded, groupBy)
		}
	} else {
		if svc.jsonOutput {
			util.PrintJsonAndExit(testSuites)
		}
		print.PrintTestSuitesResult(svc.cmd, testSuites, name)
	}
}

func (svc Service) RemoveTestSuiteByIdOrName(id string, name string) {

	testSuites := svc.GetTestSuitesByIdOrName(id, name, internalconstant.Empty, false, false, true)

	if len(testSuites) == 1 {
		testSuiteId := strconv.Itoa(testSuites[0].TestSuiteId)
		svc.RemoveTestSuiteById(testSuiteId)
		if !svc.jsonOutput {
			util.PrintSuccess("\n" + util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.RemoveTSSuccess))
			util.PrintSuccessInBold(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.RemoveTSSuccessInBold), testSuiteId)
			util.PrintSuccess(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, "removeTSSuccessEnd") + "\n")
			print.PrintTestSuiteForRemovedTestSuite(testSuites[0])
		}
	} else {
		if svc.jsonOutput {
			util.PrintJsonAndExit(testSuites)
		}
		print.PrintTestSuitesResult(svc.cmd, testSuites, name)
	}

}

func (svc Service) RestoreTestSuiteByIdOrName(id string, name string) {

	testSuites := svc.GetTestSuitesByIdOrName(id, name, internalconstant.Empty, false, false, true)
	if len(testSuites) == 1 {
		testSuiteId := strconv.Itoa(testSuites[0].TestSuiteId)
		testSuite := svc.RestoreTestSuiteById(testSuiteId)
		if svc.jsonOutput {
			util.PrintJsonAndExit(testSuite)
		}
		util.PrintSuccess("\n" + util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.RestoreTSSuccess) + "\n")
		print.PrintTestSuiteForRemovedTestSuite(testSuites[0])
	} else {
		if svc.jsonOutput {
			util.PrintJsonAndExit(testSuites)
		}
		print.PrintTestSuitesResult(svc.cmd, testSuites, name)
	}
}

func updateModifiedTestSuiteFields(testSuiteV3 model.TestSuite, name, description, propertyId, propertyName string, propVersion int, unlocked, stateful, removeProperty bool, locked bool, stateless bool) (*model.TestSuite, bool) {

	var isChanged = false
	if name != internalconstant.Empty && testSuiteV3.TestSuiteName != name {
		testSuiteV3.TestSuiteName = name
		isChanged = true
	}

	if description != internalconstant.Empty && testSuiteV3.TestSuiteDescription != description {
		testSuiteV3.TestSuiteDescription = description
		isChanged = true
	}

	if locked || unlocked {
		if locked {
			testSuiteV3.IsLocked = locked
		} else {
			testSuiteV3.IsLocked = !unlocked
		}
		isChanged = true
	}

	if stateful || stateless {
		if stateful {
			testSuiteV3.IsStateful = stateful
		} else {
			testSuiteV3.IsStateful = !stateless
		}
		isChanged = true
	}

	if testSuiteV3.Configs.PropertyManager.PropertyName != internalconstant.Empty && removeProperty {
		testSuiteV3.Configs = model.AkamaiConfigs{}
		isChanged = true
	}

	if (propertyName != internalconstant.Empty || propertyId != internalconstant.Empty) &&
		(testSuiteV3.Configs.PropertyManager.PropertyName != propertyName ||
			testSuiteV3.Configs.PropertyManager.PropertyId != util.GetConvertedInteger(propertyId) ||
			testSuiteV3.Configs.PropertyManager.PropertyVersion != propVersion) {

		testSuiteV3.Configs.PropertyManager.PropertyName = propertyName
		testSuiteV3.Configs.PropertyManager.PropertyId = util.GetConvertedInteger(propertyId)
		testSuiteV3.Configs.PropertyManager.PropertyVersion = propVersion
		isChanged = true
	}

	return &testSuiteV3, isChanged
}

// Returns testSuites containing searchString in either name or description
func filterTestSuitesByString(searchString string, exactMatch, shouldMatchDescription bool, testSuites []model.TestSuite) []model.TestSuite {
	var filteredItems []model.TestSuite

	for _, testSuite := range testSuites {
		if exactMatch {
			if strings.EqualFold(testSuite.TestSuiteName, searchString) ||
				(shouldMatchDescription && strings.EqualFold(testSuite.TestSuiteDescription, searchString)) {
				filteredItems = append(filteredItems, testSuite)
			}
		} else {
			if util.ContainsIgnoreCase(testSuite.TestSuiteName, searchString) ||
				(shouldMatchDescription && util.ContainsIgnoreCase(testSuite.TestSuiteDescription, searchString)) {
				filteredItems = append(filteredItems, testSuite)
			}
		}
	}

	return filteredItems
}

func getRequestHeaders(headerAdd []string, headerModify []string, headerFilter []string) []model.RequestHeader {
	var requestHeaders []model.RequestHeader

	for _, header := range headerAdd {
		headerComponents := strings.Split(header, ":")
		requestHeaders = append(requestHeaders, model.RequestHeader{
			HeaderName:   strings.TrimSpace(headerComponents[0]),
			HeaderValue:  strings.TrimSpace(headerComponents[1]),
			HeaderAction: internalconstant.Add,
		})
	}

	for _, header := range headerModify {
		headerComponents := strings.Split(header, ":")
		requestHeaders = append(requestHeaders, model.RequestHeader{
			HeaderName:   strings.TrimSpace(headerComponents[0]),
			HeaderValue:  strings.TrimSpace(headerComponents[1]),
			HeaderAction: internalconstant.Modify,
		})
	}

	for _, header := range headerFilter {
		requestHeaders = append(requestHeaders, model.RequestHeader{
			HeaderName:   strings.TrimSpace(header),
			HeaderAction: internalconstant.Filter,
		})
	}

	return requestHeaders
}

func getSetVariables(setVariables []string) []model.DynamicVariable {
	var variables []model.DynamicVariable

	for _, variable := range setVariables {
		variableComponent := strings.Split(variable, ":")
		variables = append(variables, model.DynamicVariable{
			VariableName:  strings.TrimSpace(variableComponent[0]),
			VariableValue: strings.TrimSpace(variableComponent[1]),
		})
	}
	return variables
}

func (svc Service) GenerateTestSuite(propertyId, propertyName, propVersion string, urls []string,
	defaultTestSuiteRequest model.DefaultTestSuiteRequest, isJsonInputPresent bool) {
	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.GenerateDefaultTestSuite), !svc.jsonOutput).Start()

	defaultTsReq := model.DefaultTestSuiteRequest{}

	// form request
	if isJsonInputPresent {
		defaultTsReq = defaultTestSuiteRequest
	} else {
		defaultTsReq = model.DefaultTestSuiteRequest{
			TestRequestUrl: urls,
			Configs: model.AkamaiConfigs{
				PropertyManager: model.PropertyManager{
					PropertyId:      util.GetConvertedInteger(propertyId),
					PropertyName:    propertyName,
					PropertyVersion: util.GetConvertedInteger(propVersion),
				}},
		}
	}

	// get default ts
	generatedTs, err := svc.api.GenerateTestSuite(defaultTsReq)

	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}
	spinner.StopWithSuccess()

	// print output
	if !svc.jsonOutput {
		util.PrintSuccess(fmt.Sprintf(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.GenerateDefaultTSSuccess)+"\n\n", generatedTs.Configs.PropertyManager.PropertyName, generatedTs.Configs.PropertyManager.PropertyVersion))
	}
	util.PrintJsonAndExit(generatedTs)
}
