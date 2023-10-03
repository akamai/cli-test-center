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

var testLogLinesCmd = &cobra.Command{
	Use:     externalconstant.TestLogLinesUse,
	Aliases: []string{externalconstant.TestLogLinesCommandAliases},
	Example: externalconstant.TestLogLinesExampleUse,
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		testRunService := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, false)

		validator := validator.NewValidator(cmd, []byte{})
		// validate tcx-id flag.
		validator.ValidateGetLogLinesFlag(testCaseExecutionId)

		testRunService.GetTestLogLinesAndPrintJson(testCaseExecutionId)

	},
}

func init() {
	testCmd.AddCommand(testLogLinesCmd)
	testLogLinesCmd.Flags().SortFlags = false

	testLogLinesCmd.Short = util.GetMessageForKey(testLogLinesCmd, internalconstant.Short)
	testLogLinesCmd.Long = util.GetMessageForKey(testLogLinesCmd, internalconstant.Long)

	testLogLinesCmd.Flags().StringVarP(&testCaseExecutionId, externalconstant.FlagTestCaseExecId, externalconstant.FlagTestCaseExecShortHand, internalconstant.Empty, util.GetMessageForKey(testLogLinesCmd, externalconstant.FlagTestCaseExecId))

}
