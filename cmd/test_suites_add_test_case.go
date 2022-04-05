package cmd

import (
	"github.com/akamai/cli-test-center/internal"
	"github.com/spf13/cobra"
)

var testSuitesAddTestCaseCmd = &cobra.Command{
	Use:     TestSuiteAddTestCaseUse,
	Example: TestSuiteAddTestCaseExample,

	Run: func(cmd *cobra.Command, args []string) {
		eghc := internal.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, jsonOutput)
		validator := internal.NewValidator(cmd, []byte{})

		// validate subcommand
		validator.ValidateSubcommandsNoArgCheck(cmd, args)

		//Check if all required flag are present.
		validator.AddTestCaseToTestSuiteFlagCheck(testSuiteIdStr, testSuiteName, url, condition, ipVersion, addHeader, modifyHeader)

		testSuites := svc.GetTestSuitesByIdOrName(testSuiteIdStr, testSuiteName, internal.Empty, true, false)
		svc.AddTestCaseWithTestSuite(testSuites, testSuiteName, url, condition, ipVersion, addHeader, modifyHeader, filterHeader)
	},
}

func init() {

	testSuitesCmd.AddCommand(testSuitesAddTestCaseCmd)
	testSuitesAddTestCaseCmd.Flags().SortFlags = false

	testSuitesAddTestCaseCmd.Short = internal.GetMessageForKey(testSuitesAddTestCaseCmd, internal.Short)
	testSuitesAddTestCaseCmd.Long = internal.GetMessageForKey(testSuitesAddTestCaseCmd, internal.Long)

	testSuitesAddTestCaseCmd.Flags().StringVar(&testSuiteIdStr, FlagTestSuiteId, "", internal.GetMessageForKey(testSuitesAddTestCaseCmd, FlagTestSuiteId))
	testSuitesAddTestCaseCmd.Flags().StringVar(&testSuiteName, FlagTestSuiteName, "", internal.GetMessageForKey(testSuitesAddTestCaseCmd, FlagTestSuiteName))
	testSuitesAddTestCaseCmd.Flags().StringVarP(&url, FlagUrl, FlagUrlShortHand, "", internal.GetMessageForKey(testSuitesAddTestCaseCmd, FlagUrl))
	testSuitesAddTestCaseCmd.Flags().StringArrayVarP(&addHeader, FlagAddHeader, FlagAddHeaderShortHand, []string{}, internal.GetMessageForKey(testSuitesAddTestCaseCmd, FlagAddHeader))
	testSuitesAddTestCaseCmd.Flags().StringArrayVarP(&modifyHeader, FlagModifyHeader, FlagModifyHeaderShortHand, []string{}, internal.GetMessageForKey(testSuitesAddTestCaseCmd, FlagModifyHeader))
	testSuitesAddTestCaseCmd.Flags().StringArrayVarP(&filterHeader, FlagFilterHeader, FlagFilterHeaderShortHand, []string{}, internal.GetMessageForKey(testSuitesAddTestCaseCmd, FlagFilterHeader))
	testSuitesAddTestCaseCmd.Flags().StringVarP(&condition, FlagCondition, FlagConditionShortHand, "", internal.GetMessageForKey(testSuitesAddTestCaseCmd, FlagCondition))
	testSuitesAddTestCaseCmd.Flags().StringVarP(&ipVersion, FlagIpVersion, FlagIpVersionShortHand, FlagIpVersionDefaultValue, internal.GetMessageForKey(testSuitesAddTestCaseCmd, FlagIpVersion))

}
