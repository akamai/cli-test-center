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

var testSuiteImport model.TestSuite

var testSuiteImportCmd = &cobra.Command{
	Use:     externalconstant.TestSuiteCreateWithChildObjects,
	Example: externalconstant.TestSuiteCreateWithChildObjectsExample,
	Aliases: []string{externalconstant.TestSuiteCreateWithChildObjectsCommandAlias},
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		svc := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, true)

		validator := validator.NewValidator(cmd, jsonData)

		validator.ValidateImportFields(&testSuiteImport)
		svc.ImportTestSuites(testSuiteImport)
	},
}

func init() {

	testSuiteCmd.AddCommand(testSuiteImportCmd)
	testSuiteImportCmd.Flags().SortFlags = false

	testSuiteImportCmd.Short = util.GetMessageForKey(testSuiteImportCmd, internalconstant.Short)
	testSuiteImportCmd.Long = util.GetMessageForKey(testSuiteImportCmd, internalconstant.Long)

}
