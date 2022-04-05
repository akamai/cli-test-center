package cmd

import (
	"github.com/akamai/cli-test-center/internal"
	"github.com/spf13/cobra"
)

var (
	id              string
	name            string
	description     string
	unlocked        bool
	stateful        bool
	propertyName    string
	propertyVersion string
	removeProperty  bool
	search          string
	user            string
	orderNumber     string
	groupBy         string
)

var testSuitesCmd = &cobra.Command{
	Use:     TestSuiteUse,
	Aliases: []string{TestSuiteCommandAlias},
	Run: func(cmd *cobra.Command, args []string) {
		validator := internal.NewValidator(cmd, []byte{})

		// validate subcommand for no arguments
		validator.NotValidSubcommandCheck(cmd, args)

		// validate subcommand
		validator.ValidSubcommandLegacyArgsCheck(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(testSuitesCmd)

	testSuitesCmd.Short = internal.GetMessageForKey(testSuitesCmd, internal.Short)
	testSuitesCmd.Long = internal.GetMessageForKey(testSuitesCmd, internal.Long)
}
