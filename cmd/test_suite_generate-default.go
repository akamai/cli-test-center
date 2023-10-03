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

var (
	urls             []string
	defaultTestSuite model.DefaultTestSuiteRequest
)

var testSuiteDefaultCmd = &cobra.Command{
	Use:     externalconstant.TestSuiteGenerateDefaultUse,
	Example: externalconstant.TestSuiteGenerateDefaultExample,
	Aliases: []string{externalconstant.TestSuiteGenerateDefaultCommandAlias},
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		svc := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, true)

		validator := validator.NewValidator(cmd, jsonData)

		// validate configVersion and urls for flag usage OR construct payload json from input file
		validator.ValidateDefaultTestSuiteFields(propertyId, propertyName, propertyVersion, urls, jsonData, &defaultTestSuite,
			isStandardInputAvailable)

		// generate default ts
		svc.GenerateTestSuite(propertyId, propertyName, propertyVersion, urls, defaultTestSuite, isStandardInputAvailable)
	},
}

func init() {
	testSuiteCmd.AddCommand(testSuiteDefaultCmd)
	testSuiteDefaultCmd.Flags().SortFlags = false

	testSuiteDefaultCmd.Short = util.GetMessageForKey(testSuiteDefaultCmd, internalconstant.Short)
	testSuiteDefaultCmd.Long = util.GetMessageForKey(testSuiteDefaultCmd, internalconstant.Long)
	testSuiteDefaultCmd.Flags().StringVarP(&propertyId, externalconstant.FlagPropertyId, internalconstant.Empty, internalconstant.Empty, util.GetMessageForKey(testSuiteDefaultCmd, externalconstant.FlagPropertyId))
	testSuiteDefaultCmd.Flags().StringVarP(&propertyName, externalconstant.FlagPropertyName, externalconstant.FlagPropertyShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteDefaultCmd, externalconstant.FlagPropertyName))
	testSuiteDefaultCmd.Flags().StringVarP(&propertyVersion, externalconstant.FlagPropertyVersion, externalconstant.FlagPropertyVersionShortHand, internalconstant.Empty, util.GetMessageForKey(testSuiteDefaultCmd, externalconstant.FlagPropertyVersion))
	testSuiteDefaultCmd.Flags().StringArrayVarP(&urls, externalconstant.FlagUrl, externalconstant.FlagUrlShortHand, []string{}, util.GetMessageForKey(testSuiteDefaultCmd, externalconstant.FlagUrl))
}
