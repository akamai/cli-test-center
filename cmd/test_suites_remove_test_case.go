package cmd

import (
	"strconv"

	"github.com/akamai/cli-test-center/internal"
	"github.com/spf13/cobra"
)

var testSuitesRemoveTestCaseCmd = &cobra.Command{
	Use:     TestSuiteRemoveTestCaseUse,
	Example: TestSuiteRemoveTestCaseExample,
	Run: func(cmd *cobra.Command, args []string) {
		eghc := internal.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, jsonOutput)
		validator := internal.NewValidator(cmd, []byte{})

		// validate subcommand
		validator.ValidateSubcommandsNoArgCheck(cmd, args)

		//Check if all required flag are present.
		validator.RemoveTestCaseFromTestSuiteFlagCheck(testSuiteIdStr, orderNumber)

		testSuiteId, _ := strconv.Atoi(testSuiteIdStr)
		testSuite := svc.GetSingleTestSuiteByIdOrName(testSuiteIdStr, "", internal.Empty, true)
		testCases, _ := svc.GetV3AssociatedTestCasesForTestSuite(testSuiteId)
		svc.RemoveTestCaseFromTestSuiteUsingOrderNumber(testSuite, testCases, orderNumber)
	},
}

func init() {

	testSuitesCmd.AddCommand(testSuitesRemoveTestCaseCmd)
	testSuitesRemoveTestCaseCmd.Flags().SortFlags = false

	testSuitesRemoveTestCaseCmd.Short = internal.GetMessageForKey(testSuitesRemoveTestCaseCmd, internal.Short)
	testSuitesRemoveTestCaseCmd.Long = internal.GetMessageForKey(testSuitesRemoveTestCaseCmd, internal.Long)

	testSuitesRemoveTestCaseCmd.Flags().StringVar(&testSuiteIdStr, FlagTestSuiteId, "", internal.GetMessageForKey(testSuitesRemoveTestCaseCmd, FlagTestSuiteId))
	testSuitesRemoveTestCaseCmd.Flags().StringVar(&orderNumber, FlagOrderNumber, "", internal.GetMessageForKey(testSuitesRemoveTestCaseCmd, FlagOrderNumber))

}
