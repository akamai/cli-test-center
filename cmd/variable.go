package cmd

import (
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/util"
	"github.com/akamai/cli-test-center/internal/validator"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/spf13/cobra"
)

var (
	testSuiteId         string
	variableName        string
	variableValue       string
	variableGroupValues []string
	variableId          string
)

var variableCmd = &cobra.Command{
	Use:     externalconstant.VariableUse,
	Aliases: []string{externalconstant.VariableCommandAliases},
	Run: func(cmd *cobra.Command, args []string) {
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateParentSubCommands(cmd, args, false)

	},
}

func init() {
	rootCmd.AddCommand(variableCmd)

	variableCmd.Short = util.GetMessageForKey(variableCmd, internalconstant.Short)
	variableCmd.Long = util.GetMessageForKey(variableCmd, internalconstant.Long)
}
