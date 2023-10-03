package cmd

import (
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/util"
	"github.com/akamai/cli-test-center/internal/validator"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/spf13/cobra"
)

var condCmd = &cobra.Command{
	Use:     externalconstant.ConditionUse,
	Aliases: []string{externalconstant.ConditionCommandAliases},
	Run: func(cmd *cobra.Command, args []string) {
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateParentSubCommands(cmd, args, false)
	},
}

func init() {
	rootCmd.AddCommand(condCmd)

	condCmd.Short = util.GetMessageForKey(condCmd, internalconstant.Short)
	condCmd.Long = util.GetMessageForKey(condCmd, internalconstant.Long)
}
