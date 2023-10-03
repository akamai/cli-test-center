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

var getTestCaseCmd = &cobra.Command{
	Use:     externalconstant.GetTestCaseUse,
	Example: externalconstant.GetTestCaseExample,
	Aliases: []string{externalconstant.GetTestCaseCommandAlias},
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		testCaseService := service.NewService(*api, cmd, jsonOutput)

		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, false)

		validator := validator.NewValidator(cmd, []byte{})

		//Check if required test suite flag is present.
		validator.TestSuiteIdAndNameFlagCheck(testSuiteIdStr, testSuiteName)

		//Check if required test case flag is present.
		validator.ValidateTestCaseFlagCheck(testCaseIdStr)

		testCaseService.GetTestCaseByIdToTestSuite(cmd, testSuiteIdStr, testSuiteName, testCaseIdStr, resolveVariables)
	},
}

func init() {

	testCaseCmd.AddCommand(getTestCaseCmd)
	getTestCaseCmd.Flags().SortFlags = false

	getTestCaseCmd.Short = util.GetMessageForKey(getTestCaseCmd, internalconstant.Short)
	getTestCaseCmd.Long = util.GetMessageForKey(getTestCaseCmd, internalconstant.Long)

	getTestCaseCmd.Flags().StringVarP(&testSuiteIdStr, externalconstant.FlagTestSuiteId, externalconstant.FlagTestSuiteIdShortHand, internalconstant.Empty, util.GetMessageForKey(getTestCaseCmd, externalconstant.FlagTestSuiteId))
	getTestCaseCmd.Flags().StringVarP(&testCaseIdStr, externalconstant.FlagTestCaseId, externalconstant.FlagTestCaseIdShortHand, internalconstant.Empty, util.GetMessageForKey(getTestCaseCmd, externalconstant.FlagTestCaseId))
	getTestCaseCmd.Flags().BoolVar(&resolveVariables, externalconstant.FlagResolveVariables, false, util.GetMessageForKey(getTestCaseCmd, externalconstant.FlagResolveVariables))
}
