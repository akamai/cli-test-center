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

var getAllTestCaseCmd = &cobra.Command{
	Use:     externalconstant.ListTestCasesUse,
	Example: externalconstant.ListTestCasesExample,
	Aliases: []string{externalconstant.ListTestCasesCommandAlias},
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		testCaseService := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, false)

		validator := validator.NewValidator(cmd, []byte{})

		// validate testSuiteId, testSuiteName and groupBy flags.
		validator.ValidateGetTestSuiteWithChildObjectsFlags(testSuiteIdStr, testSuiteName, groupBy)

		testCaseService.GetTestCasesWithTestSuite(testSuiteIdStr, testSuiteName, groupBy, resolveVariables)
	},
}

func init() {

	testCaseCmd.AddCommand(getAllTestCaseCmd)
	getAllTestCaseCmd.Flags().SortFlags = false

	getAllTestCaseCmd.Short = util.GetMessageForKey(getAllTestCaseCmd, internalconstant.Short)
	getAllTestCaseCmd.Long = util.GetMessageForKey(getAllTestCaseCmd, internalconstant.Long)

	getAllTestCaseCmd.Flags().StringVarP(&testSuiteIdStr, externalconstant.FlagTestSuiteId, externalconstant.FlagTestSuiteIdShortHand, internalconstant.Empty, util.GetMessageForKey(getAllTestCaseCmd, externalconstant.FlagTestSuiteId))
	getAllTestCaseCmd.Flags().StringVarP(&testSuiteName, externalconstant.FlagTestSuiteName, externalconstant.FlagTestSuiteNameShortHand, internalconstant.Empty, util.GetMessageForKey(getAllTestCaseCmd, externalconstant.FlagTestSuiteName))
	getAllTestCaseCmd.Flags().BoolVar(&resolveVariables, externalconstant.FlagResolveVariables, false, util.GetMessageForKey(getAllTestCaseCmd, externalconstant.FlagResolveVariables))
	getAllTestCaseCmd.Flags().StringVarP(&groupBy, externalconstant.FlagGroupBy, externalconstant.FlagGroupShortHand, internalconstant.Empty, util.GetMessageForKey(getAllTestCaseCmd, externalconstant.FlagGroupBy))
}
