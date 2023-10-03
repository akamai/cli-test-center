package cmd

import (
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/util"
	"github.com/akamai/cli-test-center/internal/validator"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/spf13/cobra"
)

var testReqCmd = &cobra.Command{
	Use:     externalconstant.TestRequestUse,
	Aliases: []string{externalconstant.TestRequestCommandAliases},
	Run: func(cmd *cobra.Command, args []string) {
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateParentSubCommands(cmd, args, false)
	},
}

func init() {
	rootCmd.AddCommand(testReqCmd)

	testReqCmd.Short = util.GetMessageForKey(testReqCmd, internalconstant.Short)
	testReqCmd.Long = util.GetMessageForKey(testReqCmd, internalconstant.Long)
}
