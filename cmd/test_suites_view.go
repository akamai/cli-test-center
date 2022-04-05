package cmd

import (
	"github.com/akamai/cli-test-center/internal"
	"github.com/spf13/cobra"
)

var testSuitesViewCmd = &cobra.Command{
	Use:     TestSuiteViewUse,
	Example: TestSuiteViewExample,
	Aliases: []string{TestSuiteViewCommandAliases},
	Run: func(cmd *cobra.Command, args []string) {
		eghc := internal.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, jsonOutput)
		validator := internal.NewValidator(cmd, []byte{})

		// validate subcommand
		validator.ValidateSubcommandsNoArgCheck(cmd, args)

		// validate id, name and groupBy flags.
		validator.ValidateViewTestSuiteFlags(id, name, groupBy)

		//Get and print test suite, test cases
		svc.ViewTestSuite(id, name, groupBy)
	},
}

func init() {

	testSuitesCmd.AddCommand(testSuitesViewCmd)
	testSuitesViewCmd.Flags().SortFlags = false

	testSuitesViewCmd.Short = internal.GetMessageForKey(testSuitesViewCmd, internal.Short)
	testSuitesViewCmd.Long = internal.GetMessageForKey(testSuitesViewCmd, internal.Long)

	testSuitesViewCmd.Flags().StringVar(&id, FlagId, "", internal.GetMessageForKey(testSuitesViewCmd, FlagId))
	testSuitesViewCmd.Flags().StringVar(&name, FlagName, "", internal.GetMessageForKey(testSuitesViewCmd, FlagName))
	testSuitesViewCmd.Flags().StringVar(&groupBy, FlagGroupBy, "", internal.GetMessageForKey(testSuitesViewCmd, FlagGroupBy))

}
