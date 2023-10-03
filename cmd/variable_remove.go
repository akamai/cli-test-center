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

var variableRemoveCmd = &cobra.Command{
	Use:     externalconstant.VariableRemoveUse,
	Example: externalconstant.VariableRemoveExample,
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		svc := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, false)

		validator := validator.NewValidator(cmd, []byte{})
		// validate flag checks
		validator.ValidateVariableFlagCheck(testSuiteId, variableId)

		svc.RemoveVariablesAndPrintResult(cmd, testSuiteId, variableId)
	},
}

func init() {
	variableCmd.AddCommand(variableRemoveCmd)
	variableRemoveCmd.Flags().SortFlags = false

	variableRemoveCmd.Short = util.GetMessageForKey(variableRemoveCmd, internalconstant.Short)
	variableRemoveCmd.Long = util.GetMessageForKey(variableRemoveCmd, internalconstant.Long)

	variableRemoveCmd.Flags().StringVarP(&testSuiteId, externalconstant.FlagTestSuiteId, externalconstant.FlagTestSuiteIdShortHand, internalconstant.Empty, util.GetMessageForKey(variableRemoveCmd, externalconstant.FlagTestSuiteId))
	variableRemoveCmd.Flags().StringVarP(&variableId, externalconstant.FlagVariableId, internalconstant.Empty, internalconstant.Empty, util.GetMessageForKey(variableRemoveCmd, externalconstant.FlagVariableId))

}
