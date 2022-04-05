package cmd

import (
	"github.com/akamai/cli-test-center/internal"
	"github.com/spf13/cobra"
)

var testSuiteImport internal.TestSuiteDetailsWithChildObjects

var testSuitesImportCmd = &cobra.Command{
	Use:     TestSuiteImportUse,
	Example: TestSuiteImportExample,
	Run: func(cmd *cobra.Command, args []string) {
		eghc := internal.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, jsonOutput)
		validator := internal.NewValidator(cmd, jsonData)

		// validate subcommand
		validator.ValidateSubcommandsNoArgCheck(cmd, args)

		validator.ValidateImportFields(&testSuiteImport)
		svc.ImportTestSuites(testSuiteImport)
	},
}

func init() {

	testSuitesCmd.AddCommand(testSuitesImportCmd)
	testSuitesImportCmd.Flags().SortFlags = false

	testSuitesImportCmd.Short = internal.GetMessageForKey(testSuitesImportCmd, internal.Short)
	testSuitesImportCmd.Long = internal.GetMessageForKey(testSuitesImportCmd, internal.Long)

}
