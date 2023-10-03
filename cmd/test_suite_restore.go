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

var testSuiteRestoreCmd = &cobra.Command{
	Use:     externalconstant.TestSuiteRestoreUse,
	Example: externalconstant.TestSuiteRestoreExample,
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		svc := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, false)

		validator := validator.NewValidator(cmd, []byte{})

		//Check if all required flag are present.
		validator.TestSuiteIdAndNameFlagCheck(id, name)

		svc.RestoreTestSuiteByIdOrName(id, name)
	},
}

func init() {

	testSuiteCmd.AddCommand(testSuiteRestoreCmd)
	testSuiteRestoreCmd.Flags().SortFlags = false

	testSuiteRestoreCmd.Short = util.GetMessageForKey(testSuiteRestoreCmd, internalconstant.Short)
	testSuiteRestoreCmd.Long = util.GetMessageForKey(testSuiteRestoreCmd, internalconstant.Long)

	testSuiteRestoreCmd.Flags().StringVarP(&id, externalconstant.FlagTestSuiteId, externalconstant.FlagTestSuiteIdShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteRestoreCmd, externalconstant.FlagTestSuiteId))
	testSuiteRestoreCmd.Flags().StringVarP(&name, externalconstant.FlagTestSuiteName, externalconstant.FlagTestSuiteNameShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteRestoreCmd, externalconstant.FlagTestSuiteName))
}
