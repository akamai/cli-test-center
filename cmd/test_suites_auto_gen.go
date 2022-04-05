package cmd

import (
	"github.com/akamai/cli-test-center/internal"
	"github.com/spf13/cobra"
)

var (
	urls             []string
	defaultTestSuite internal.DefaultTestSuiteRequest
)

var testSuitesDefaultCmd = &cobra.Command{
	Use:     TestSuiteAutoGenerationUse,
	Example: TestSuiteAutoGenerationExample,
	Run: func(cmd *cobra.Command, args []string) {
		eghc := internal.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, jsonOutput)
		validator := internal.NewValidator(cmd, jsonData)

		// validate subcommand
		validator.ValidateSubcommandsNoArgCheck(cmd, args)

		// validate configVersion and urls for flag usage OR construct payload json from input file
		propName, propVersion := validator.ValidateDefaultTestSuiteFields(propertyName, propertyVersion, urls, jsonData, &defaultTestSuite,
			isStandardInputAvailable)

		// generate default ts
		svc.GenerateTestSuite(propName, propVersion, urls, defaultTestSuite, isStandardInputAvailable)
	},
}

func init() {
	testSuitesCmd.AddCommand(testSuitesDefaultCmd)
	testSuitesDefaultCmd.Flags().SortFlags = false

	testSuitesDefaultCmd.Short = internal.GetMessageForKey(testSuitesDefaultCmd, internal.Short)
	testSuitesDefaultCmd.Long = internal.GetMessageForKey(testSuitesDefaultCmd, internal.Long)
	testSuitesDefaultCmd.Flags().StringVar(&propertyName, FlagProperty, "", internal.GetMessageForKey(testSuitesDefaultCmd, FlagProperty))
	testSuitesDefaultCmd.Flags().StringVar(&propertyVersion, FlagPropver, "", internal.GetMessageForKey(testSuitesDefaultCmd, FlagPropver))
	testSuitesDefaultCmd.Flags().StringArrayVarP(&urls, FlagUrl, FlagUrlShortHand, []string{}, internal.GetMessageForKey(testSuitesDefaultCmd, FlagUrl))
}
