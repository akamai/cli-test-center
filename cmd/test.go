package cmd

import (
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/util"
	"github.com/akamai/cli-test-center/internal/validator"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/spf13/cobra"
)

var testCaseExecutionId string

var testCmd = &cobra.Command{
	Use:     externalconstant.TestUse,
	Aliases: []string{externalconstant.TestCommandAlias},
	Run: func(cmd *cobra.Command, args []string) {
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateParentSubCommands(cmd, args, false)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	testCmd.Short = util.GetMessageForKey(testCmd, internalconstant.Short)
	testCmd.Long = util.GetMessageForKey(testCmd, internalconstant.Long)
}
