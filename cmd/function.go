package cmd

import (
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/util"
	"github.com/akamai/cli-test-center/internal/validator"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/spf13/cobra"
)

var tryFunction model.TryFunction

var functionCmd = &cobra.Command{
	Use:     externalconstant.FunctionUse,
	Aliases: []string{externalconstant.FunctionAliases},
	Run: func(cmd *cobra.Command, args []string) {
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		globalValidator.ValidateParentSubCommands(cmd, args, false)
	},
}

func init() {

	rootCmd.AddCommand(functionCmd)
	createTestCaseCmd.Flags().SortFlags = false

	functionCmd.Short = util.GetMessageForKey(functionCmd, internalconstant.Short)
	functionCmd.Long = util.GetMessageForKey(functionCmd, internalconstant.Long)
}
