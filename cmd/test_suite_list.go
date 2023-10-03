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

var testSuiteListCmd = &cobra.Command{
	Use:     externalconstant.TestSuiteListUse,
	Example: externalconstant.TestSuiteListExample,
	Aliases: []string{externalconstant.TestSuiteListCommandAlias},
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		svc := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, false)

		validator := validator.NewValidator(cmd, []byte{})

		validator.PropertyAndVersionFlagCheck(propertyId, propertyName, propertyVersion, true)

		svc.GetTestSuitesAndPrint(cmd, propertyId, propertyName, propertyVersion, user, search)
	},
}

func init() {

	testSuiteCmd.AddCommand(testSuiteListCmd)
	testSuiteListCmd.Flags().SortFlags = false

	testSuiteListCmd.Short = util.GetMessageForKey(testSuiteListCmd, internalconstant.Short)
	testSuiteListCmd.Long = util.GetMessageForKey(testSuiteListCmd, internalconstant.Long)

	testSuiteListCmd.Flags().StringVarP(&propertyId, externalconstant.FlagPropertyId, internalconstant.Empty, internalconstant.Empty, util.GetMessageForKey(testSuiteListCmd, externalconstant.FlagPropertyId))
	testSuiteListCmd.Flags().StringVarP(&propertyName, externalconstant.FlagPropertyName, externalconstant.FlagPropertyShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteListCmd, externalconstant.FlagPropertyName))
	testSuiteListCmd.Flags().StringVarP(&propertyVersion, externalconstant.FlagPropertyVersion, externalconstant.FlagPropertyVersionShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteListCmd, externalconstant.FlagPropertyVersion))
	testSuiteListCmd.Flags().StringVarP(&user, externalconstant.FlagUser, externalconstant.FlagUserShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteListCmd, externalconstant.FlagUser))
	testSuiteListCmd.Flags().StringVar(&search, externalconstant.FlagSearch, internalconstant.Empty, util.GetMessageForKey(testSuiteListCmd, externalconstant.FlagSearch))

}
