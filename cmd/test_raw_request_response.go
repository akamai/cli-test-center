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

var testRawReqResCmd = &cobra.Command{
	Use:     externalconstant.TestRawReqResUse,
	Aliases: []string{externalconstant.TestRawReqResCommandAliases},
	Example: externalconstant.TestRawReqResExampleUse,
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		testRunService := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, false)

		validator := validator.NewValidator(cmd, []byte{})
		// validate ids flag.
		validator.ValidateGetRawRequestResponseFlag(testRunId, testCaseExecutionId)

		testRunService.GetRawRequestResponseAndPrintResult(testRunId, testCaseExecutionId)

	},
}

func init() {
	testCmd.AddCommand(testRawReqResCmd)
	testRawReqResCmd.Flags().SortFlags = false

	testRawReqResCmd.Short = util.GetMessageForKey(testRawReqResCmd, internalconstant.Short)
	testRawReqResCmd.Long = util.GetMessageForKey(testRawReqResCmd, internalconstant.Long)

	testRawReqResCmd.Flags().StringVarP(&testRunId, externalconstant.FlagTestRunId, externalconstant.FlagTestRunIdShortHand, internalconstant.Empty, util.GetMessageForKey(testRawReqResCmd, externalconstant.FlagTestRunId))
	testRawReqResCmd.Flags().StringVarP(&testCaseExecutionId, externalconstant.FlagTestCaseExecId, externalconstant.FlagTestCaseExecShortHand, internalconstant.Empty, util.GetMessageForKey(testRawReqResCmd, externalconstant.FlagTestCaseExecId))

}
