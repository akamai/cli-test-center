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

var testListCmd = &cobra.Command{
	Use:     externalconstant.TestListUse,
	Aliases: []string{externalconstant.TestListAlias},
	Example: externalconstant.TestListExample,
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		testRunService := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, false)

		testRunService.GetTestRunsAndPrint()

	},
}

func init() {
	testCmd.AddCommand(testListCmd)
	testListCmd.Flags().SortFlags = false

	testListCmd.Short = util.GetMessageForKey(testListCmd, internalconstant.Short)
	testListCmd.Long = util.GetMessageForKey(testListCmd, internalconstant.Long)
}
