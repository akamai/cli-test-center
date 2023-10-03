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

var testSuiteGetCmd = &cobra.Command{
	Use:     externalconstant.TestSuiteGetUse,
	Example: externalconstant.TestSuiteGetExample,
	Aliases: []string{externalconstant.TestSuiteGetCommandAliases},
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		svc := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, false)

		validator := validator.NewValidator(cmd, []byte{})

		// validate id, name flags.
		validator.TestSuiteIdAndNameFlagCheck(id, name)

		//Get and print test suite, test cases
		svc.GetTestSuiteAndPrint(id, name)
	},
}

func init() {

	testSuiteCmd.AddCommand(testSuiteGetCmd)
	testSuiteGetCmd.Flags().SortFlags = false

	testSuiteGetCmd.Short = util.GetMessageForKey(testSuiteGetCmd, internalconstant.Short)
	testSuiteGetCmd.Long = util.GetMessageForKey(testSuiteGetCmd, internalconstant.Long)

	testSuiteGetCmd.Flags().StringVarP(&id, externalconstant.FlagTestSuiteId, externalconstant.FlagTestSuiteIdShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteGetCmd, externalconstant.FlagTestSuiteId))
	testSuiteGetCmd.Flags().StringVarP(&name, externalconstant.FlagTestSuiteName, externalconstant.FlagTestSuiteNameShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteGetCmd, externalconstant.FlagTestSuiteName))

}
