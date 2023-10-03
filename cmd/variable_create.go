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

var variableCreateCmd = &cobra.Command{
	Use:     externalconstant.VariableCreateUse,
	Example: externalconstant.VariableCreateExample,
	Aliases: []string{externalconstant.VariableCreateCommandAliases},
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		svc := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, false)

		validator := validator.NewValidator(cmd, []byte{})
		// validate flag checks.
		validator.ValidateVariableCreateFlagCheck(testSuiteId, variableName, variableValue, variableGroupValues)

		svc.CreateVariablesAndPrintResult(cmd, variableName, variableValue, testSuiteId, variableGroupValues)
	},
}

func init() {
	variableCmd.AddCommand(variableCreateCmd)
	variableCreateCmd.Flags().SortFlags = false

	variableCreateCmd.Short = util.GetMessageForKey(variableCreateCmd, internalconstant.Short)
	variableCreateCmd.Long = util.GetMessageForKey(variableCreateCmd, internalconstant.Long)

	variableCreateCmd.Flags().StringVarP(&variableName, externalconstant.FlagVariableName, externalconstant.FlagVariableNameShortHand, internalconstant.Empty, util.GetMessageForKey(variableCreateCmd, externalconstant.FlagVariableName))
	variableCreateCmd.Flags().StringVarP(&variableValue, externalconstant.FlagVariableValue, externalconstant.FlagVariableValueShortHand, internalconstant.Empty, util.GetMessageForKey(variableCreateCmd, externalconstant.FlagVariableValue))
	variableCreateCmd.Flags().StringVarP(&testSuiteId, externalconstant.FlagTestSuiteId, externalconstant.FlagTestSuiteIdShortHand, internalconstant.Empty, util.GetMessageForKey(variableCreateCmd, externalconstant.FlagTestSuiteId))
	variableCreateCmd.Flags().StringArrayVarP(&variableGroupValues, externalconstant.FlagVarGroupValue, externalconstant.FlagVariableGroupShortHand, []string{}, util.GetMessageForKey(variableCreateCmd, externalconstant.FlagVarGroupValue))

}
