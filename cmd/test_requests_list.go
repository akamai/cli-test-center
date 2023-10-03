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

var testReqListCmd = &cobra.Command{
	Use:     externalconstant.TestRequestListUse,
	Example: externalconstant.TestRequestListExample,
	Aliases: []string{externalconstant.TestRequestListCommandAliases},
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		svc := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, false)

		svc.GetTestRequestAndPrintResult(cmd)
	},
}

func init() {
	testReqCmd.AddCommand(testReqListCmd)

	testReqListCmd.Short = util.GetMessageForKey(testReqListCmd, internalconstant.Short)
	testReqListCmd.Long = util.GetMessageForKey(testReqListCmd, internalconstant.Long)
}
