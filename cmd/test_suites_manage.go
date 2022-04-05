package cmd

import (
	"github.com/akamai/cli-test-center/internal"
	"github.com/spf13/cobra"
)

var testSuiteManage internal.TestSuiteDetailsWithChildObjects

var testSuitesManageCmd = &cobra.Command{
	Use:     TestSuiteManageUse,
	Example: TestSuiteManageExample,
	Run: func(cmd *cobra.Command, args []string) {
		eghc := internal.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, jsonOutput)
		validator := internal.NewValidator(cmd, jsonData)

		// validate subcommand
		validator.ValidateSubcommandsNoArgCheck(cmd, args)

		validator.ValidateManageFields(&testSuiteManage)
		svc.ManageTestSuites(testSuiteManage)
	},
}

func init() {

	testSuitesCmd.AddCommand(testSuitesManageCmd)
	testSuitesManageCmd.Flags().SortFlags = false

	testSuitesManageCmd.Short = internal.GetMessageForKey(testSuitesManageCmd, internal.Short)
	testSuitesManageCmd.Long = internal.GetMessageForKey(testSuitesManageCmd, internal.Long)

}
