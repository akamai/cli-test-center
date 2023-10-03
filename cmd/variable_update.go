package cmd

import (
	"github.com/akamai/cli-test-center/internal/api"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/service"
	"github.com/akamai/cli-test-center/internal/util"
	"github.com/akamai/cli-test-center/internal/validator"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/spf13/cobra"
)

var variableUpdateCmd = &cobra.Command{
	Use:     externalconstant.VariableUpdateUse,
	Example: externalconstant.VariableUpdateExample,
	Aliases: []string{externalconstant.VariableUpdateCommandAliases},
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		svc := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, false)

		validator := validator.NewValidator(cmd, []byte{})
		// validate flag checks
		validator.ValidateVariableEditFlagCheck(testSuiteId, variableName, variableValue, variableId, variableGroupValues)

		svc.UpdateVariablesAndPrintResult(cmd, testSuiteId, variableId, variableName, variableValue, variableGroupValues)
	},
}

func init() {
	variableCmd.AddCommand(variableUpdateCmd)
	variableUpdateCmd.Flags().SortFlags = false

	variableUpdateCmd.Short = util.GetMessageForKey(variableUpdateCmd, internalconstant.Short)
	variableUpdateCmd.Long = util.GetMessageForKey(variableUpdateCmd, internalconstant.Long)

	variableUpdateCmd.Flags().StringVarP(&testSuiteId, externalconstant.FlagTestSuiteId, externalconstant.FlagTestSuiteIdShortHand, internalconstant.Empty, util.GetMessageForKey(variableUpdateCmd, externalconstant.FlagTestSuiteId))
	variableUpdateCmd.Flags().StringVarP(&variableId, externalconstant.FlagVariableId, internalconstant.Empty, internalconstant.Empty, util.GetMessageForKey(variableUpdateCmd, externalconstant.FlagVariableId))
	variableUpdateCmd.Flags().StringVarP(&variableName, externalconstant.FlagVariableName, externalconstant.FlagVariableNameShortHand, internalconstant.Empty, util.GetMessageForKey(variableUpdateCmd, externalconstant.FlagVariableName))
	variableUpdateCmd.Flags().StringVarP(&variableValue, externalconstant.FlagVariableValue, externalconstant.FlagVariableValueShortHand, internalconstant.Empty, util.GetMessageForKey(variableUpdateCmd, externalconstant.FlagVariableValue))
	variableUpdateCmd.Flags().StringArrayVarP(&variableGroupValues, externalconstant.FlagVarGroupValue, externalconstant.FlagVariableGroupShortHand, []string{}, util.GetMessageForKey(variableUpdateCmd, externalconstant.FlagVarGroupValue))

}
