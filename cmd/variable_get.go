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

var variableGetCmd = &cobra.Command{
	Use:     externalconstant.VariableGetUse,
	Example: externalconstant.VariableGetExample,
	Aliases: []string{externalconstant.VariableGetCommandAliases},
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		svc := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, false)

		validator := validator.NewValidator(cmd, []byte{})
		// validate flag checks.
		validator.ValidateVariableFlagCheck(testSuiteId, variableId)

		svc.GetVariableAndPrintResult(cmd, testSuiteId, variableId)
	},
}

func init() {
	variableCmd.AddCommand(variableGetCmd)
	variableGetCmd.Flags().SortFlags = false

	variableGetCmd.Short = util.GetMessageForKey(variableGetCmd, internalconstant.Short)
	variableGetCmd.Long = util.GetMessageForKey(variableGetCmd, internalconstant.Long)

	variableGetCmd.Flags().StringVarP(&testSuiteId, externalconstant.FlagTestSuiteId, externalconstant.FlagTestSuiteIdShortHand, internalconstant.Empty, util.GetMessageForKey(variableGetCmd, externalconstant.FlagTestSuiteId))
	variableGetCmd.Flags().StringVarP(&variableId, externalconstant.FlagVariableId, internalconstant.Empty, internalconstant.Empty, util.GetMessageForKey(variableGetCmd, externalconstant.FlagVariableId))
}
