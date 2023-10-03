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

var removeTestCaseCmd = &cobra.Command{
	Use:     externalconstant.RemoveTestCaseUse,
	Example: externalconstant.RemoveTestCaseExample,
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		testCaseService := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, false)

		validator := validator.NewValidator(cmd, []byte{})

		//Check if all required flag are present.
		validator.RemoveTestCaseFromTestSuiteFlagCheck(testSuiteIdStr, orderNumber, testCaseIdStr)

		testCaseService.RemoveTestCaseFromTestSuiteUsingOrderNumberOrTestCaseId(testSuiteIdStr, orderNumber, testCaseIdStr)
	},
}

func init() {

	testCaseCmd.AddCommand(removeTestCaseCmd)
	removeTestCaseCmd.Flags().SortFlags = false

	removeTestCaseCmd.Short = util.GetMessageForKey(removeTestCaseCmd, internalconstant.Short)
	removeTestCaseCmd.Long = util.GetMessageForKey(removeTestCaseCmd, internalconstant.Long)

	removeTestCaseCmd.Flags().StringVarP(&testSuiteIdStr, externalconstant.FlagTestSuiteId, externalconstant.FlagTestSuiteIdShortHand, internalconstant.Empty, util.GetMessageForKey(removeTestCaseCmd, externalconstant.FlagTestSuiteId))
	removeTestCaseCmd.Flags().StringVarP(&orderNumber, externalconstant.FlagOrderNumber, externalconstant.FlagOrderNumberShortHand, internalconstant.Empty, util.GetMessageForKey(removeTestCaseCmd, externalconstant.FlagOrderNumber))
	removeTestCaseCmd.Flags().StringVarP(&testCaseIdStr, externalconstant.FlagTestCaseId, externalconstant.FlagTestCaseIdShortHand, internalconstant.Empty, util.GetMessageForKey(removeTestCaseCmd, externalconstant.FlagTestCaseId))

}
