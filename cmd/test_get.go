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

var testRunId string

var testGetCmd = &cobra.Command{
	Use:     externalconstant.TestGetUse,
	Aliases: []string{externalconstant.TestGetAlias},
	Example: externalconstant.TestGetExample,
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		testRunService := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, false)

		validator := validator.NewValidator(cmd, []byte{})
		// validate id flag.
		validator.ValidateGetTestRunFlag(testRunId)

		testRunService.GetTestRunAndPrintResult(testRunId)

	},
}

func init() {
	testCmd.AddCommand(testGetCmd)
	testGetCmd.Flags().SortFlags = false

	testGetCmd.Short = util.GetMessageForKey(testGetCmd, internalconstant.Short)
	testGetCmd.Long = util.GetMessageForKey(testGetCmd, internalconstant.Long)

	testGetCmd.Flags().StringVarP(&testRunId, externalconstant.FlagTestRunId, externalconstant.FlagTestRunIdShortHand, internalconstant.Empty, util.GetMessageForKey(testGetCmd, externalconstant.FlagTestRunId))
}
