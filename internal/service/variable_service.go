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

func (svc Service) CreateVariablesAndPrintResult(cmd *cobra.Command, varname, value, testSuiteId string, varGroupValue []string) {
	variables, err := svc.CreateVariable(varname, value, testSuiteId, varGroupValue)

	if err != nil {
		util.AbortForCommand(svc.cmd, err)
	}

	if len(variables.Successes) != 0 {
		util.PrintSuccess(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.AddVariablesSuccess) + "\n")
		fmt.Println()
		print.PrintVariables(variables.Successes, false)
	}

	if len(variables.Failures) != 0 {
		util.PrintErrorMessages(util.GetApiSubErrorMessagesForCommand(svc.cmd, variables.Failures, internalconstant.Empty, internalconstant.Empty, internalconstant.Empty))
	}

}

func (svc Service) CreateVariable(varname, value, testSuiteId string, varGroupValue []string) (*model.VariableBulkResponse, *model.CliError) {

	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.AddVariables), !svc.jsonOutput).Start()

	variables := []model.Variable{
		{
			VariableName:       varname,
			VariableValue:      value,
			VariableGroupValue: GetVarGroupArray(varGroupValue),
		},
	}

	varBulkResp, err := svc.api.CreateVariables(variables, testSuiteId)

	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}

	spinner.StopWithSuccess()
	if svc.jsonOutput {
		util.PrintJsonAndExit(varBulkResp)
	}

	return varBulkResp, nil

}

func GetVarGroupArray(varGroup []string) []model.VariableGroupValue {
	var varGroupArray []model.VariableGroupValue
	for _, userProvidedGroupValue := range varGroup {
		var vg model.VariableGroupValue

		// Get the header and values parts of the variable group
		parts := strings.SplitN(userProvidedGroupValue, externalconstant.Colon, 2)
		vg.ColumnHeader = parts[0]
		vg.ColumnValues = strings.Split(parts[1], externalconstant.Comma)

		// Assign the header and values parts of the variable group in request body format
		varGroupArray = append(varGroupArray, vg)
	}

	return varGroupArray
}

func (svc Service) GetVariablesAndPrintResult(cmd *cobra.Command, testSuiteId string) {
	variableList := svc.GetVariablesList(testSuiteId)
	util.PrintHeader(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.GetVariablesSuccess) + "\n\n")

	print.PrintVariablesResult(cmd, variableList, false)
	fmt.Printf("\n"+externalconstant.LabelTotalItem+": %d\n", len(variableList))

}

func (svc Service) GetVariablesList(testSuiteId string) []model.Variable {
	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty,
		internalconstant.GetVariables), !svc.jsonOutput).Start()
	variables, err := svc.api.GetVariables(testSuiteId)
	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}

	spinner.StopWithSuccess()

	if svc.jsonOutput {
		util.PrintJsonAndExit(variables)
	}
	return variables
}

func (svc Service) GetVariableAndPrintResult(cmd *cobra.Command, testSuiteId, variableId string) {
	variable := svc.GetVariable(testSuiteId, variableId)
	var variables []model.Variable
	variables = append(variables, *variable)

	util.PrintHeader(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.GetVariableSuccess) + "\n\n")
	print.PrintVariablesResult(cmd, variables, false)
}

func (svc Service) GetVariable(testSuiteId, variableId string) *model.Variable {
	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty,
		internalconstant.GetVariable), !svc.jsonOutput).Start()

	variable, err := svc.api.GetVariable(testSuiteId, variableId)
	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}

	spinner.StopWithSuccess()

	if svc.jsonOutput {
		util.PrintJsonAndExit(variable)
	}
	return variable
}

func (svc Service) UpdateVariablesAndPrintResult(cmd *cobra.Command, testSuiteId, variableId, variableName, value string, variableGroupValue []string) {
	variables, err := svc.UpdateVariable(testSuiteId, variableId, variableName, value, variableGroupValue)

	if err != nil {
		util.AbortForCommand(svc.cmd, err)
	}

	if len(variables.Successes) != 0 {
		util.PrintSuccess(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.UpdateVariablesSuccess) + "\n")
		fmt.Println()
		print.PrintVariables(variables.Successes, false)
	}

	if len(variables.Failures) != 0 {
		util.PrintErrorMessages(util.GetApiSubErrorMessagesForCommand(svc.cmd, variables.Failures, internalconstant.Empty, internalconstant.Empty, internalconstant.Empty))
	}

}

func (svc Service) UpdateVariable(testSuiteId, variableId, variableName, value string, variableGroupValue []string) (*model.VariableBulkResponse, *model.CliError) {

	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.UpdateVariables), !svc.jsonOutput).Start()

	variables := constructVariablesForUpdate(variableId, variableName, value, variableGroupValue)

	varBulkResp, err := svc.api.UpdateVariables(variables, testSuiteId)

	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}

	spinner.StopWithSuccess()
	if svc.jsonOutput {
		util.PrintJsonAndExit(varBulkResp)
	}

	return varBulkResp, nil

}

func constructVariablesForUpdate(variableId, variablename, value string, varGroupValue []string) []model.Variable {

	varId, _ := strconv.Atoi(variableId)
	variables := []model.Variable{
		{
			VariableId:         varId,
			VariableName:       variablename,
			VariableValue:      value,
			VariableGroupValue: GetVarGroupArray(varGroupValue),
		},
	}
	return variables
}

func (svc Service) RemoveVariablesAndPrintResult(cmd *cobra.Command, testSuiteId, variableId string) {
	variables, err := svc.RemoveVariable(testSuiteId, variableId)

	if err != nil {
		util.AbortForCommand(svc.cmd, err)
	}

	if len(variables.Successes) != 0 {
		util.PrintSuccess(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.RemoveVariableSuccess) + "\n")
	}

	if len(variables.Failures) != 0 {
		util.PrintErrorMessages(util.GetApiSubErrorMessagesForCommand(svc.cmd, variables.Failures, internalconstant.Empty, internalconstant.Empty, internalconstant.Empty))
	}

}

func (svc Service) RemoveVariable(testSuiteId, variableId string) (*model.BulkResponse, *model.CliError) {
	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.RemoveVariable), !svc.jsonOutput).Start()
	log.Debugf("Removing test cases [%d] from the test suite id [%d]\n", variableId, testSuiteId)

	removeVariableResponse, err := svc.api.RemoveVariableFromTestSuite(testSuiteId, variableId)
	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}

	spinner.StopWithSuccess()
	if svc.jsonOutput {
		util.PrintJsonAndExit(removeVariableResponse)
	}

	return removeVariableResponse, nil
}
