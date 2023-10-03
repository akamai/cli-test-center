package cmd

import (
	"github.com/akamai/cli-test-center/internal/api"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/service"
	"github.com/akamai/cli-test-center/internal/util"
	"github.com/akamai/cli-test-center/internal/validator"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/spf13/cobra"
)

var testSuite model.TestSuite

var testSuiteAddCmd = &cobra.Command{
	Use:     externalconstant.TestSuiteAddUse,
	Example: externalconstant.TestSuiteAddExample,
	Aliases: []string{externalconstant.TestSuiteAddCommandAlias},
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		svc := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, true)

		validator := validator.NewValidator(cmd, jsonData)

		// validate json input or flags provided by user and assign json to test suite if provided
		validator.ValidateCreateTestSuiteFields(&testSuite, isStandardInputAvailable, name, propertyId, propertyName, propertyVersion)

		svc.AddTestSuiteAndPrint(cmd, name, description, propertyId, propertyName, propertyVersion, unlocked, stateful, isStandardInputAvailable, testSuite)
	},
}

func init() {

	testSuiteCmd.AddCommand(testSuiteAddCmd)
	testSuiteAddCmd.Flags().SortFlags = false

	testSuiteAddCmd.Short = util.GetMessageForKey(testSuiteAddCmd, internalconstant.Short)
	testSuiteAddCmd.Long = util.GetMessageForKey(testSuiteAddCmd, internalconstant.Long)

	testSuiteAddCmd.Flags().StringVarP(&name, externalconstant.FlagTestSuiteName, externalconstant.FlagTestSuiteNameShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteAddCmd, externalconstant.FlagTestSuiteName))
	testSuiteAddCmd.Flags().StringVarP(&description, externalconstant.FlagDescription, externalconstant.FlagDescriptionShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteAddCmd, externalconstant.FlagDescription))
	testSuiteAddCmd.Flags().BoolVar(&unlocked, externalconstant.FlagUnlocked, false, util.GetMessageForKey(testSuiteAddCmd, externalconstant.FlagUnlocked))
	testSuiteAddCmd.Flags().BoolVar(&stateful, externalconstant.FlagStateFul, false, util.GetMessageForKey(testSuiteAddCmd, externalconstant.FlagStateFul))
	testSuiteAddCmd.Flags().StringVarP(&propertyId, externalconstant.FlagPropertyId, internalconstant.Empty, internalconstant.Empty, util.GetMessageForKey(testSuiteAddCmd, externalconstant.FlagPropertyId))
	testSuiteAddCmd.Flags().StringVarP(&propertyName, externalconstant.FlagPropertyName, externalconstant.FlagPropertyShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteAddCmd, externalconstant.FlagPropertyName))
	testSuiteAddCmd.Flags().StringVarP(&propertyVersion, externalconstant.FlagPropertyVersion, externalconstant.FlagPropertyVersionShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteAddCmd, externalconstant.FlagPropertyVersion))

}
