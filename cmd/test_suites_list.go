package cmd

import (
	"github.com/akamai/cli-test-center/internal"
	"github.com/spf13/cobra"
)

var testSuitesListCmd = &cobra.Command{
	Use:     TestSuiteListUse,
	Example: TestSuiteListExample,
	Aliases: []string{TestSuiteListCommandAlias},
	Run: func(cmd *cobra.Command, args []string) {
		eghc := internal.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, jsonOutput)
		validator := internal.NewValidator(cmd, []byte{})

		// validate subcommand
		validator.ValidateSubcommandsNoArgCheck(cmd, args)

		validator.ConfigFlagCheckForListTestSuites(propertyVersion, propertyName)

		testSuites := svc.GetTestSuites(propertyName, propertyVersion, user, search)
		internal.PrintTestSuitesTable(cmd, testSuites)
	},
}

func init() {

	testSuitesCmd.AddCommand(testSuitesListCmd)
	testSuitesListCmd.Flags().SortFlags = false

	testSuitesListCmd.Short = internal.GetMessageForKey(testSuitesListCmd, internal.Short)
	testSuitesListCmd.Long = internal.GetMessageForKey(testSuitesListCmd, internal.Long)

	testSuitesListCmd.Flags().StringVar(&propertyName, FlagProperty, "", internal.GetMessageForKey(testSuitesListCmd, FlagProperty))
	testSuitesListCmd.Flags().StringVar(&propertyVersion, FlagPropver, "", internal.GetMessageForKey(testSuitesListCmd, FlagPropver))
	testSuitesListCmd.Flags().StringVarP(&user, FlagUser, FlagUserShortHand, "", internal.GetMessageForKey(testSuitesListCmd, FlagUser))
	testSuitesListCmd.Flags().StringVarP(&search, FlagSearch, FlagSearchShortHand, "", internal.GetMessageForKey(testSuitesListCmd, FlagSearch))

}
