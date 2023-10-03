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

var variablesListCmd = &cobra.Command{
	Use:     externalconstant.VariablesListUse,
	Example: externalconstant.VariablesListExample,
	Aliases: []string{externalconstant.VariablesListCommandAliases},
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		svc := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, false)

		validator := validator.NewValidator(cmd, []byte{})
		// validate flag checks.
		validator.ValidateVariablesListFlagCheck(testSuiteId)

		svc.GetVariablesAndPrintResult(cmd, testSuiteId)
	},
}

func init() {
	variableCmd.AddCommand(variablesListCmd)

	variablesListCmd.Short = util.GetMessageForKey(variablesListCmd, internalconstant.Short)
	variablesListCmd.Long = util.GetMessageForKey(variablesListCmd, internalconstant.Long)

	variablesListCmd.Flags().StringVarP(&testSuiteId, externalconstant.FlagTestSuiteId, externalconstant.FlagTestSuiteIdShortHand, internalconstant.Empty, util.GetMessageForKey(variablesListCmd, externalconstant.FlagTestSuiteId))

}
