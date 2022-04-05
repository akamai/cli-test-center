package cmd

import (
	"github.com/akamai/cli-test-center/internal"
	"github.com/spf13/cobra"
)

var (
	url               string
	addHeader         []string
	modifyHeader      []string
	filterHeader      []string
	condition         string
	ipVersion         string
	testSuiteName     string
	testSuiteIdStr    string
	propertyStr       string
	propverStr        string
	targetEnvironment string
)

var testRunRequest internal.TestRun

var testCmd = &cobra.Command{
	Use:     TestUse,
	Aliases: []string{TestCommandAlias},
	Example: TestExample,
	Run: func(cmd *cobra.Command, args []string) {
		eghc := internal.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, jsonOutput)
		validator := internal.NewValidator(cmd, jsonData)

		// validate subcommand
		validator.ValidateSubcommandsNoArgCheck(cmd, args)

		// validate different flags and combination to run test
		runTestUsing := validator.ValidateTestRunFlagsAndGetRunEnum(testSuiteIdStr, testSuiteName, propertyStr, propverStr,
			url, condition, ipVersion, targetEnvironment, addHeader, modifyHeader, jsonData, &testRunRequest,
			isStandardInputAvailable)

		//Run test
		svc.RunTest(runTestUsing, testSuiteIdStr, testSuiteName, propertyStr, propverStr,
			url, condition, ipVersion, targetEnvironment, addHeader, modifyHeader, filterHeader, testRunRequest)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().SortFlags = false

	testCmd.Short = internal.GetMessageForKey(testCmd, internal.Short)
	testCmd.Long = internal.GetMessageForKey(testCmd, internal.Long)

	testCmd.Flags().StringVarP(&url, FlagUrl, FlagUrlShortHand, "", internal.GetMessageForKey(testCmd, FlagUrl))
	testCmd.Flags().StringArrayVarP(&addHeader, FlagAddHeader, FlagAddHeaderShortHand, []string{}, internal.GetMessageForKey(testCmd, FlagAddHeader))
	testCmd.Flags().StringArrayVarP(&modifyHeader, FlagModifyHeader, FlagModifyHeaderShortHand, []string{}, internal.GetMessageForKey(testCmd, FlagModifyHeader))
	testCmd.Flags().StringArrayVarP(&filterHeader, FlagFilterHeader, FlagFilterHeaderShortHand, []string{}, internal.GetMessageForKey(testCmd, FlagFilterHeader))
	testCmd.Flags().StringVarP(&condition, FlagCondition, FlagConditionShortHand, "", internal.GetMessageForKey(testCmd, FlagCondition))
	testCmd.Flags().StringVarP(&ipVersion, FlagIpVersion, FlagIpVersionShortHand, FlagIpVersionDefaultValue, internal.GetMessageForKey(testCmd, FlagIpVersion))
	testCmd.Flags().StringVar(&testSuiteName, FlagTestSuiteName, "", internal.GetMessageForKey(testCmd, FlagTestSuiteName))
	testCmd.Flags().StringVar(&testSuiteIdStr, FlagTestSuiteId, "", internal.GetMessageForKey(testCmd, FlagTestSuiteId))
	testCmd.Flags().StringVar(&propertyStr, FlagProperty, "", internal.GetMessageForKey(testCmd, FlagProperty))
	testCmd.Flags().StringVar(&propverStr, FlagPropver, "", internal.GetMessageForKey(testCmd, FlagPropver))
	testCmd.Flags().StringVarP(&targetEnvironment, FlagEnv, FlagEnvShortHand, internal.Staging, internal.GetMessageForKey(testCmd, FlagEnv))
}
