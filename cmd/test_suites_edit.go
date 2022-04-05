package cmd

import (
	"github.com/akamai/cli-test-center/internal"
	"github.com/spf13/cobra"
)

var testSuitesEditCmd = &cobra.Command{
	Use:     TestSuiteEditUse,
	Example: TestSuiteEditExample,
	Run: func(cmd *cobra.Command, args []string) {
		eghc := internal.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, jsonOutput)
		validator := internal.NewValidator(cmd, []byte{})

		// validate subcommand
		validator.ValidateSubcommandsNoArgCheck(cmd, args)

		// validate id flag check.
		validator.EditTestSuiteIdFlagCheck(id)

		// validate propertyVersion usage.
		propName, propVersion := validator.PropertyAndVersionFlagCheck(propertyName, propertyVersion, true)

		//Remove config flag check
		validator.RemoveConfigFlagCheck(propName, removeProperty)

		testSuite := svc.EditTestSuite(id, name, description, propName, propVersion, unlocked, stateful, removeProperty)
		internal.PrintSuccess(internal.GetServiceMessage(cmd, internal.MessageTypeDisplay, "", "editTSSuccess") + "\n")
		internal.PrintTestSuite(*testSuite)
	},
}

func init() {

	testSuitesCmd.AddCommand(testSuitesEditCmd)
	testSuitesEditCmd.Flags().SortFlags = false

	testSuitesEditCmd.Short = internal.GetMessageForKey(testSuitesEditCmd, internal.Short)
	testSuitesEditCmd.Long = internal.GetMessageForKey(testSuitesEditCmd, internal.Long)

	testSuitesEditCmd.Flags().StringVar(&id, FlagId, "", internal.GetMessageForKey(testSuitesEditCmd, FlagId))
	testSuitesEditCmd.Flags().StringVar(&name, FlagName, "", internal.GetMessageForKey(testSuitesEditCmd, FlagName))
	testSuitesEditCmd.Flags().StringVar(&description, FlagDescription, "", internal.GetMessageForKey(testSuitesEditCmd, FlagDescription))
	testSuitesEditCmd.Flags().BoolVar(&unlocked, FlagUnlocked, false, internal.GetMessageForKey(testSuitesEditCmd, FlagUnlocked))
	testSuitesEditCmd.Flags().BoolVar(&stateful, FlagStateFul, false, internal.GetMessageForKey(testSuitesEditCmd, FlagStateFul))
	testSuitesEditCmd.Flags().StringVar(&propertyName, FlagProperty, "", internal.GetMessageForKey(testSuitesEditCmd, FlagProperty))
	testSuitesEditCmd.Flags().StringVar(&propertyVersion, FlagPropver, "", internal.GetMessageForKey(testSuitesEditCmd, FlagPropver))
	testSuitesEditCmd.Flags().BoolVar(&removeProperty, FlagRemoveProperty, false, internal.GetMessageForKey(testSuitesEditCmd, FlagRemoveProperty))

}
