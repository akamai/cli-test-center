package cmd

import (
	"github.com/akamai/cli-test-center/internal/api"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/service"
	"github.com/akamai/cli-test-center/internal/util"
	"github.com/akamai/cli-test-center/internal/validator"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/spf13/cobra"
)

var testSuiteEditCmd = &cobra.Command{
	Use:     externalconstant.TestSuiteEditUse,
	Example: externalconstant.TestSuiteEditExample,
	Aliases: []string{externalconstant.TestSuiteEditCommandAlias},
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		svc := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, true)

		validator := validator.NewValidator(cmd, jsonData)

		// validate json input or flags provided by user and assign json to test suite if provided
		validator.ValidateUpdateTestSuiteFields(&testSuite, isStandardInputAvailable, id, propertyId, propertyName, propertyVersion,
			locked, unlocked, stateful, stateless, removeProperty)

		svc.EditTestSuiteAndPrint(cmd, id, name, description, propertyId, propertyName, propertyVersion, unlocked, stateful,
			removeProperty, locked, stateless, isStandardInputAvailable, testSuite)
	},
}

func init() {

	testSuiteCmd.AddCommand(testSuiteEditCmd)
	testSuiteEditCmd.Flags().SortFlags = false

	testSuiteEditCmd.Short = util.GetMessageForKey(testSuiteEditCmd, internalconstant.Short)
	testSuiteEditCmd.Long = util.GetMessageForKey(testSuiteEditCmd, internalconstant.Long)

	testSuiteEditCmd.Flags().StringVarP(&id, externalconstant.FlagTestSuiteId, externalconstant.FlagTestSuiteIdShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteEditCmd, externalconstant.FlagTestSuiteId))
	testSuiteEditCmd.Flags().StringVarP(&name, externalconstant.FlagTestSuiteName, externalconstant.FlagTestSuiteNameShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteEditCmd, externalconstant.FlagTestSuiteName))
	testSuiteEditCmd.Flags().StringVarP(&description, externalconstant.FlagDescription, externalconstant.FlagDescriptionShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteEditCmd, externalconstant.FlagDescription))
	testSuiteEditCmd.Flags().BoolVar(&unlocked, externalconstant.FlagUnlocked, false, util.GetMessageForKey(testSuiteEditCmd, externalconstant.FlagUnlocked))
	testSuiteEditCmd.Flags().BoolVar(&stateful, externalconstant.FlagStateFul, false, util.GetMessageForKey(testSuiteEditCmd, externalconstant.FlagStateFul))
	testSuiteEditCmd.Flags().BoolVar(&locked, externalconstant.FlagLocked, false, util.GetMessageForKey(testSuiteEditCmd, externalconstant.FlagLocked))
	testSuiteEditCmd.Flags().BoolVar(&stateless, externalconstant.FlagStateless, false, util.GetMessageForKey(testSuiteEditCmd, externalconstant.FlagStateless))
	testSuiteEditCmd.Flags().StringVarP(&propertyId, externalconstant.FlagPropertyId, internalconstant.Empty, internalconstant.Empty, util.GetMessageForKey(testSuiteEditCmd, externalconstant.FlagPropertyId))
	testSuiteEditCmd.Flags().StringVarP(&propertyName, externalconstant.FlagPropertyName, externalconstant.FlagPropertyShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteEditCmd, externalconstant.FlagPropertyName))
	testSuiteEditCmd.Flags().StringVarP(&propertyVersion, externalconstant.FlagPropertyVersion, externalconstant.FlagPropertyVersionShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteEditCmd, externalconstant.FlagPropertyVersion))
	testSuiteEditCmd.Flags().BoolVar(&removeProperty, externalconstant.FlagRemoveProperty, false, util.GetMessageForKey(testSuiteEditCmd, externalconstant.FlagRemoveProperty))

}
