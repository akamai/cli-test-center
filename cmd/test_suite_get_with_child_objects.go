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

var testSuiteExportCmd = &cobra.Command{
	Use:     externalconstant.TestSuiteGetWithChildObjectsUse,
	Example: externalconstant.TestSuiteGetWithChildObjectsExample,
	Aliases: []string{externalconstant.TestSuiteGetWithChildObjectsCommandAliases},
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		svc := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, false)

		validator := validator.NewValidator(cmd, []byte{})

		// validate id, name and groupBy flags.
		validator.ValidateGetTestSuiteWithChildObjectsFlags(id, name, groupBy)

		//Get and print test suite, test cases
		svc.GetTestSuiteWithChildObjectsAndPrint(id, name, groupBy, resolveVariables)
	},
}

func init() {

	testSuiteCmd.AddCommand(testSuiteExportCmd)
	testSuiteExportCmd.Flags().SortFlags = false

	testSuiteExportCmd.Short = util.GetMessageForKey(testSuiteExportCmd, internalconstant.Short)
	testSuiteExportCmd.Long = util.GetMessageForKey(testSuiteExportCmd, internalconstant.Long)

	testSuiteExportCmd.Flags().StringVarP(&id, externalconstant.FlagTestSuiteId, externalconstant.FlagTestSuiteIdShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteExportCmd, externalconstant.FlagTestSuiteId))
	testSuiteExportCmd.Flags().StringVarP(&name, externalconstant.FlagTestSuiteName, externalconstant.FlagTestSuiteNameShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteExportCmd, externalconstant.FlagTestSuiteName))
	testSuiteExportCmd.Flags().StringVarP(&groupBy, externalconstant.FlagGroupBy, externalconstant.FlagGroupShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteExportCmd, externalconstant.FlagGroupBy))
	testSuiteExportCmd.Flags().BoolVar(&resolveVariables, externalconstant.FlagResolveVariables, false, util.GetMessageForKey(testSuiteExportCmd, externalconstant.FlagResolveVariables))

}
