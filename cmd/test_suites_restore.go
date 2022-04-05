package cmd

import (
	"github.com/akamai/cli-test-center/internal"
	"github.com/spf13/cobra"
)

var testSuitesRestoreCmd = &cobra.Command{
	Use:     TestSuiteRestoreUse,
	Example: TestSuiteRestoreExample,
	Run: func(cmd *cobra.Command, args []string) {
		eghc := internal.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, jsonOutput)
		validator := internal.NewValidator(cmd, []byte{})

		// validate subcommand
		validator.ValidateSubcommandsNoArgCheck(cmd, args)

		//Check if all required flag are present.
		validator.TestSuiteIdAndNameFlagCheck(id, name)

		svc.RestoreTestSuiteByIdOrName(id, name)
	},
}

func init() {

	testSuitesCmd.AddCommand(testSuitesRestoreCmd)
	testSuitesRestoreCmd.Flags().SortFlags = false

	testSuitesRestoreCmd.Short = internal.GetMessageForKey(testSuitesRestoreCmd, internal.Short)
	testSuitesRestoreCmd.Long = internal.GetMessageForKey(testSuitesRestoreCmd, internal.Long)

	testSuitesRestoreCmd.Flags().StringVar(&id, FlagId, "", internal.GetMessageForKey(testSuitesRestoreCmd, FlagId))
	testSuitesRestoreCmd.Flags().StringVar(&name, FlagName, "", internal.GetMessageForKey(testSuitesRestoreCmd, FlagName))
}
