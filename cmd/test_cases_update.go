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

var updateTestCaseCmd = &cobra.Command{
	Use:     externalconstant.UpdateTestCaseUse,
	Example: externalconstant.UpdateTestCaseExample,
	Aliases: []string{externalconstant.UpdateTestCaseCommandAlias},
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

		//Check if required test case flag is present.
		validator.ValidateTestCaseFlagCheck(testCaseIdStr)

		// common method for create and update test case under TS. As part of update, testCaseId should be valid.
		testCaseService.EditTestCaseWithTestSuite(cmd, testSuiteIdStr, testSuiteName, url, condition, ipVersion, addHeader, modifyHeader, filterHeader, testCaseIdStr, strings.ToUpper(client), strings.ToUpper(requestMethod), requestBody, encodeRequestBody, setVariables)
	},
}

func init() {

	testCaseCmd.AddCommand(updateTestCaseCmd)
	updateTestCaseCmd.Flags().SortFlags = false

	updateTestCaseCmd.Short = util.GetMessageForKey(updateTestCaseCmd, internalconstant.Short)
	updateTestCaseCmd.Long = util.GetMessageForKey(updateTestCaseCmd, internalconstant.Long)

	updateTestCaseCmd.Flags().StringVarP(&testSuiteIdStr, externalconstant.FlagTestSuiteId, externalconstant.FlagTestSuiteIdShortHand, internalconstant.Empty, util.GetMessageForKey(updateTestCaseCmd, externalconstant.FlagTestSuiteId))
	updateTestCaseCmd.Flags().StringVarP(&testSuiteName, externalconstant.FlagTestSuiteName, externalconstant.FlagTestSuiteNameShortHand, internalconstant.Empty, util.GetMessageForKey(updateTestCaseCmd, externalconstant.FlagTestSuiteName))
	updateTestCaseCmd.Flags().StringVarP(&testCaseIdStr, externalconstant.FlagTestCaseId, externalconstant.FlagTestCaseIdShortHand, internalconstant.Empty, util.GetMessageForKey(updateTestCaseCmd, externalconstant.FlagTestCaseId))
	updateTestCaseCmd.Flags().StringVarP(&url, externalconstant.FlagUrl, externalconstant.FlagUrlShortHand, internalconstant.Empty, util.GetMessageForKey(updateTestCaseCmd, externalconstant.FlagUrl))
	updateTestCaseCmd.Flags().StringArrayVarP(&addHeader, externalconstant.FlagAddHeader, externalconstant.FlagAddHeaderShortHand, []string{}, util.GetMessageForKey(updateTestCaseCmd, externalconstant.FlagAddHeader))
	updateTestCaseCmd.Flags().StringArrayVarP(&modifyHeader, externalconstant.FlagModifyHeader, externalconstant.FlagModifyHeaderShortHand, []string{}, util.GetMessageForKey(updateTestCaseCmd, externalconstant.FlagModifyHeader))
	updateTestCaseCmd.Flags().StringArrayVarP(&filterHeader, externalconstant.FlagFilterHeader, externalconstant.FlagFilterHeaderShortHand, []string{}, util.GetMessageForKey(updateTestCaseCmd, externalconstant.FlagFilterHeader))
	updateTestCaseCmd.Flags().StringVarP(&condition, externalconstant.FlagCondition, externalconstant.FlagConditionShortHand, internalconstant.Empty, util.GetMessageForKey(updateTestCaseCmd, externalconstant.FlagCondition))
	updateTestCaseCmd.Flags().StringVar(&ipVersion, externalconstant.FlagIpVersion, internalconstant.IpVersionDefaultValue, util.GetMessageForKey(updateTestCaseCmd, externalconstant.FlagIpVersion))
	updateTestCaseCmd.Flags().StringVarP(&client, externalconstant.FlagClient, externalconstant.FlagClientShortHand, internalconstant.Curl, util.GetMessageForKey(updateTestCaseCmd, externalconstant.FlagClient))
	updateTestCaseCmd.Flags().StringVarP(&requestMethod, externalconstant.FlagRequestMethod, externalconstant.FlagRequestMethodShortHand, internalconstant.GetRequestMethod, util.GetMessageForKey(updateTestCaseCmd, externalconstant.FlagRequestMethod))
	updateTestCaseCmd.Flags().StringVarP(&requestBody, externalconstant.FlagRequestBody, externalconstant.FlagRequestBodyShortHand, internalconstant.Empty, util.GetMessageForKey(updateTestCaseCmd, externalconstant.FlagRequestBody))
	updateTestCaseCmd.Flags().BoolVarP(&encodeRequestBody, externalconstant.FlagEncodeRequestBody, externalconstant.FlagEncodeRequestBodyShortHand, false, util.GetMessageForKey(updateTestCaseCmd, externalconstant.FlagEncodeRequestBody))
	updateTestCaseCmd.Flags().StringArrayVarP(&setVariables, externalconstant.FlagSetVariables, externalconstant.FlagSetVariablesShortHand, []string{}, util.GetMessageForKey(updateTestCaseCmd, externalconstant.FlagSetVariables))
}
