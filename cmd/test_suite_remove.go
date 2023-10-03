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

var testSuiteRemoveCmd = &cobra.Command{
	Use:     externalconstant.TestSuiteRemoveUse,
	Example: externalconstant.TestSuiteRemoveExample,
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

		svc.RemoveTestSuiteByIdOrName(id, name)
	},
}

func init() {

	testSuiteCmd.AddCommand(testSuiteRemoveCmd)
	testSuiteRemoveCmd.Flags().SortFlags = false

	testSuiteRemoveCmd.Short = util.GetMessageForKey(testSuiteRemoveCmd, internalconstant.Short)
	testSuiteRemoveCmd.Long = util.GetMessageForKey(testSuiteRemoveCmd, internalconstant.Long)

	testSuiteRemoveCmd.Flags().StringVarP(&id, externalconstant.FlagTestSuiteId, externalconstant.FlagTestSuiteIdShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteRemoveCmd, externalconstant.FlagTestSuiteId))
	testSuiteRemoveCmd.Flags().StringVarP(&name, externalconstant.FlagTestSuiteName, externalconstant.FlagTestSuiteNameShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteRemoveCmd, externalconstant.FlagTestSuiteName))
}
