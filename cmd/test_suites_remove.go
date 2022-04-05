package cmd

import (
	"github.com/akamai/cli-test-center/internal"
	"github.com/spf13/cobra"
)

var testSuitesRemoveCmd = &cobra.Command{
	Use:     TestSuiteRemoveUse,
	Example: TestSuiteRemoveExample,
	Run: func(cmd *cobra.Command, args []string) {
		eghc := internal.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, jsonOutput)
		validator := internal.NewValidator(cmd, []byte{})

		// validate subcommand
		validator.ValidateSubcommandsNoArgCheck(cmd, args)

		//Check if all required flag are present.
		validator.TestSuiteIdAndNameFlagCheck(id, name)

		svc.RemoveTestSuiteByIdOrName(id, name)
	},
}

func init() {

	testSuitesCmd.AddCommand(testSuitesRemoveCmd)
	testSuitesRemoveCmd.Flags().SortFlags = false

	testSuitesRemoveCmd.Short = internal.GetMessageForKey(testSuitesRemoveCmd, internal.Short)
	testSuitesRemoveCmd.Long = internal.GetMessageForKey(testSuitesRemoveCmd, internal.Long)

	testSuitesRemoveCmd.Flags().StringVar(&id, FlagId, "", internal.GetMessageForKey(testSuitesRemoveCmd, FlagId))
	testSuitesRemoveCmd.Flags().StringVar(&name, FlagName, "", internal.GetMessageForKey(testSuitesRemoveCmd, FlagName))
}
