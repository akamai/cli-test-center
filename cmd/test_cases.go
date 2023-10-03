package cmd

import (
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/util"
	"github.com/akamai/cli-test-center/internal/validator"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/spf13/cobra"
)

var (
	testCaseIdStr    string
	resolveVariables bool
	setVariables     []string
)

var testCaseCmd = &cobra.Command{
	Use:     externalconstant.TestCaseUse,
	Aliases: []string{externalconstant.TestCaseCommandAlias},
	Run: func(cmd *cobra.Command, args []string) {
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)

		globalValidator.ValidateParentSubCommands(cmd, args, false)
	},
}

func init() {

	rootCmd.AddCommand(testCaseCmd)
	testCaseCmd.Flags().SortFlags = false

	testCaseCmd.Short = util.GetMessageForKey(testCaseCmd, internalconstant.Short)
	testCaseCmd.Long = util.GetMessageForKey(testCaseCmd, internalconstant.Long)

}
