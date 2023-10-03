package cmd

import (
	"github.com/akamai/cli-test-center/internal/api"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/service"
	"github.com/akamai/cli-test-center/internal/util"
	"github.com/akamai/cli-test-center/internal/validator"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/spf13/cobra"
	"strings"
)

var createTestCaseCmd = &cobra.Command{
	Use:     externalconstant.CreateTestCaseUse,
	Example: externalconstant.CreateTestCaseExample,
	Aliases: []string{externalconstant.CreateTestCaseCommandAlias},
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		testCaseService := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, false)

		validator := validator.NewValidator(cmd, []byte{})
		//Check if all required flag are present.
		validator.AddTestCaseToTestSuiteFlagCheck(testSuiteIdStr, testSuiteName, url, condition, ipVersion, addHeader, modifyHeader, client, requestMethod, requestBody, encodeRequestBody, setVariables)

		// common method for create and update test case under TS. As part of create, testCaseId will be always empty.
		testCaseService.AddTestCaseWithTestSuite(cmd, testSuiteIdStr, testSuiteName, url, condition, ipVersion, addHeader, modifyHeader, filterHeader, internalconstant.Empty, strings.ToUpper(client), strings.ToUpper(requestMethod), requestBody, encodeRequestBody, setVariables)
	},
}

func init() {

	testCaseCmd.AddCommand(createTestCaseCmd)
	createTestCaseCmd.Flags().SortFlags = false

	createTestCaseCmd.Short = util.GetMessageForKey(createTestCaseCmd, internalconstant.Short)
	createTestCaseCmd.Long = util.GetMessageForKey(createTestCaseCmd, internalconstant.Long)

	createTestCaseCmd.Flags().StringVarP(&testSuiteIdStr, externalconstant.FlagTestSuiteId, externalconstant.FlagTestSuiteIdShortHand, internalconstant.Empty, util.GetMessageForKey(createTestCaseCmd, externalconstant.FlagTestSuiteId))
	createTestCaseCmd.Flags().StringVarP(&testSuiteName, externalconstant.FlagTestSuiteName, externalconstant.FlagTestSuiteNameShortHand, internalconstant.Empty, util.GetMessageForKey(createTestCaseCmd, externalconstant.FlagTestSuiteName))
	createTestCaseCmd.Flags().StringVarP(&url, externalconstant.FlagUrl, externalconstant.FlagUrlShortHand, internalconstant.Empty, util.GetMessageForKey(createTestCaseCmd, externalconstant.FlagUrl))
	createTestCaseCmd.Flags().StringArrayVarP(&addHeader, externalconstant.FlagAddHeader, externalconstant.FlagAddHeaderShortHand, []string{}, util.GetMessageForKey(createTestCaseCmd, externalconstant.FlagAddHeader))
	createTestCaseCmd.Flags().StringArrayVarP(&modifyHeader, externalconstant.FlagModifyHeader, externalconstant.FlagModifyHeaderShortHand, []string{}, util.GetMessageForKey(createTestCaseCmd, externalconstant.FlagModifyHeader))
	createTestCaseCmd.Flags().StringArrayVarP(&filterHeader, externalconstant.FlagFilterHeader, externalconstant.FlagFilterHeaderShortHand, []string{}, util.GetMessageForKey(createTestCaseCmd, externalconstant.FlagFilterHeader))
	createTestCaseCmd.Flags().StringVarP(&condition, externalconstant.FlagCondition, externalconstant.FlagConditionShortHand, internalconstant.Empty, util.GetMessageForKey(createTestCaseCmd, externalconstant.FlagCondition))
	createTestCaseCmd.Flags().StringVar(&ipVersion, externalconstant.FlagIpVersion, internalconstant.IpVersionDefaultValue, util.GetMessageForKey(createTestCaseCmd, externalconstant.FlagIpVersion))
	createTestCaseCmd.Flags().StringVarP(&client, externalconstant.FlagClient, externalconstant.FlagClientShortHand, internalconstant.Curl, util.GetMessageForKey(createTestCaseCmd, externalconstant.FlagClient))
	createTestCaseCmd.Flags().StringVarP(&requestMethod, externalconstant.FlagRequestMethod, externalconstant.FlagRequestMethodShortHand, internalconstant.GetRequestMethod, util.GetMessageForKey(createTestCaseCmd, externalconstant.FlagRequestMethod))
	createTestCaseCmd.Flags().StringVarP(&requestBody, externalconstant.FlagRequestBody, externalconstant.FlagRequestBodyShortHand, internalconstant.Empty, util.GetMessageForKey(createTestCaseCmd, externalconstant.FlagRequestBody))
	createTestCaseCmd.Flags().BoolVarP(&encodeRequestBody, externalconstant.FlagEncodeRequestBody, externalconstant.FlagEncodeRequestBodyShortHand, false, util.GetMessageForKey(createTestCaseCmd, externalconstant.FlagEncodeRequestBody))
	createTestCaseCmd.Flags().StringArrayVarP(&setVariables, externalconstant.FlagSetVariables, externalconstant.FlagSetVariablesShortHand, []string{}, util.GetMessageForKey(createTestCaseCmd, externalconstant.FlagSetVariables))
}
