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

var testSuiteManage model.TestSuite

var testSuiteManageCmd = &cobra.Command{
	Use:     externalconstant.TestSuiteUpdateWithChildObjectsUse,
	Example: externalconstant.TestSuiteUpdateWithChildObjectsExample,
	Aliases: []string{externalconstant.TestSuiteUpdateWithChildObjectsCommandAlias},
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		svc := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, true)

		validator := validator.NewValidator(cmd, jsonData)

		validator.ValidateManageFields(&testSuiteManage)
		svc.ManageTestSuites(testSuiteManage)
	},
}

func init() {

	testSuiteCmd.AddCommand(testSuiteManageCmd)
	testSuiteManageCmd.Flags().SortFlags = false

	testSuiteManageCmd.Short = util.GetMessageForKey(testSuiteManageCmd, internalconstant.Short)
	testSuiteManageCmd.Long = util.GetMessageForKey(testSuiteManageCmd, internalconstant.Long)

}
