package print

import (
	"fmt"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/util"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/spf13/cobra"
	"strings"
)

func PrintVariablesResult(cmd *cobra.Command, variables []model.Variable, forTestSuite bool) {

	size := len(variables)
	switch size {
	case 0:
		util.PrintWarning(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.NoVariablesFoundWarning) + "\n")
	default:
		PrintVariables(variables, forTestSuite)
	}
}

func PrintVariables(variables []model.Variable, forTestSuite bool) {

	var variableGroup []model.Variable
	for _, variable := range variables {
		if variable.VariableValue != internalconstant.Empty {
			if !forTestSuite {
				util.PrintLabelAndValue(externalconstant.VariableId, variable.VariableId)
				util.PrintLabelAndValue(externalconstant.VariableName, variable.VariableName)
				util.PrintLabelAndValue(externalconstant.VariableValue, variable.VariableValue)
				util.PrintLabelAndValue(externalconstant.LabelCreated, util.FormatTime(variable.CreatedDate)+externalconstant.SeparateBy+variable.CreatedBy)
				util.PrintLabelAndValue(externalconstant.LabelLastModified, util.FormatTime(variable.ModifiedDate)+externalconstant.SeparateBy+variable.ModifiedBy)
				fmt.Println()
			} else {
				fmt.Printf("%v=%v\n", variable.VariableName, variable.VariableValue)
			}
		} else {
			variableGroup = append(variableGroup, variable)
		}
	}
	fmt.Println()

	if len(variableGroup) > 0 {
		printVariablesGroupTables(variableGroup, forTestSuite)
	}
}

func printVariablesGroupTables(varGroup []model.Variable, forTestSuite bool) {

	for _, variableGroup := range varGroup {
		if !forTestSuite {
			util.PrintLabelAndValue(externalconstant.VariableId, variableGroup.VariableId)
			util.PrintLabelAndValue(externalconstant.VariableName, variableGroup.VariableName)
			util.PrintLabelAndValue(externalconstant.LabelCreated, util.FormatTime(variableGroup.CreatedDate)+externalconstant.SeparateBy+variableGroup.CreatedBy)
			util.PrintLabelAndValue(externalconstant.LabelLastModified, util.FormatTime(variableGroup.ModifiedDate)+externalconstant.SeparateBy+variableGroup.ModifiedBy)
		} else {
			fmt.Printf("%v\n", variableGroup.VariableName)
		}

		//printing table for particular variableGroup
		printVariableGroupTable(&variableGroup)
	}
}

func printVariableGroupTable(variableGroup *model.Variable) {

	var columnHeaders []string
	// Assigning length of first columnValues of group value
	var columnValues = make([][]string, len(variableGroup.VariableGroupValue[0].ColumnValues))

	for _, variableGroupValue := range variableGroup.VariableGroupValue {
		// Setting headers formatting to Uppercase
		columnHeader := strings.ToUpper(variableGroupValue.ColumnHeader)

		columnHeaders = append(columnHeaders, columnHeader)
		for i, columnValue := range variableGroupValue.ColumnValues {
			columnValues[i] = append(columnValues[i], columnValue)
		}
	}
	util.ShowTable(columnHeaders, columnValues, false)
	fmt.Println()
}
