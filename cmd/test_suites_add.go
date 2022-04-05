package cmd

import (
	"github.com/akamai/cli-test-center/internal"
	"github.com/spf13/cobra"
)

var testSuitesAddCmd = &cobra.Command{
	Use:     TestSuiteAddUse,
	Example: TestSuiteAddExample,
	Run: func(cmd *cobra.Command, args []string) {
		eghc := internal.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, jsonOutput)
		validator := internal.NewValidator(cmd, []byte{})

		// validate subcommand
		validator.ValidateSubcommandsNoArgCheck(cmd, args)

		// validate name flag check.
		validator.AddTestSuiteNameFlagCheck(name)

		// validate propertyVersion usage.
		propName, propVersion := validator.PropertyAndVersionFlagCheck(propertyName, propertyVersion, true)

		testSuite := svc.AddTestSuite(name, description, propName, propVersion, unlocked, stateful)
		internal.PrintSuccess(internal.GetServiceMessage(cmd, internal.MessageTypeDisplay, "", "addTestSuiteSuccess") + "\n")
		internal.PrintTestSuite(*testSuite)
	},
}

func init() {

	testSuitesCmd.AddCommand(testSuitesAddCmd)
	testSuitesAddCmd.Flags().SortFlags = false

	testSuitesAddCmd.Short = internal.GetMessageForKey(testSuitesAddCmd, internal.Short)
	testSuitesAddCmd.Long = internal.GetMessageForKey(testSuitesAddCmd, internal.Long)

	testSuitesAddCmd.Flags().StringVar(&name, FlagName, "", internal.GetMessageForKey(testSuitesAddCmd, FlagName))
	testSuitesAddCmd.Flags().StringVar(&description, FlagDescription, "", internal.GetMessageForKey(testSuitesAddCmd, FlagDescription))
	testSuitesAddCmd.Flags().BoolVar(&unlocked, FlagUnlocked, false, internal.GetMessageForKey(testSuitesAddCmd, FlagUnlocked))
	testSuitesAddCmd.Flags().BoolVar(&stateful, FlagStateFul, false, internal.GetMessageForKey(testSuitesAddCmd, FlagStateFul))
	testSuitesAddCmd.Flags().StringVar(&propertyName, FlagProperty, "", internal.GetMessageForKey(testSuitesAddCmd, FlagProperty))
	testSuitesAddCmd.Flags().StringVar(&propertyVersion, FlagPropver, "", internal.GetMessageForKey(testSuitesAddCmd, FlagPropver))

}
