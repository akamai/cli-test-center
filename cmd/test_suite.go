package cmd

import (
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/util"
	"github.com/akamai/cli-test-center/internal/validator"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/spf13/cobra"
)

var (
	id              string
	name            string
	description     string
	unlocked        bool
	locked          bool
	stateful        bool
	stateless       bool
	propertyId      string
	propertyName    string
	propertyVersion string
	removeProperty  bool
	search          string
	user            string
	orderNumber     string
	groupBy         string
)

var testSuiteCmd = &cobra.Command{
	Use:     externalconstant.TestSuiteUse,
	Aliases: []string{externalconstant.TestSuiteCommandAlias},
	Run: func(cmd *cobra.Command, args []string) {
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateParentSubCommands(cmd, args, false)

	},
}

func init() {
	rootCmd.AddCommand(testSuiteCmd)

	testSuiteCmd.Short = util.GetMessageForKey(testSuiteCmd, internalconstant.Short)
	testSuiteCmd.Long = util.GetMessageForKey(testSuiteCmd, internalconstant.Long)
}
