package cmd

import (
	"github.com/akamai/cli-test-center/internal"
	"github.com/spf13/cobra"
)

var condTemplateCmd = &cobra.Command{
	Use:     ConditionTemplateUse,
	Example: ConditionTemplateExample,
	Run: func(cmd *cobra.Command, args []string) {
		eghc := internal.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := internal.NewApiClient(*eghc)
		svc := internal.NewService(*api, cmd, jsonOutput)
		validator := internal.NewValidator(cmd, []byte{})

		// validate subcommand
		validator.ValidateSubcommandsNoArgCheck(cmd, args)
		condTemplate := svc.GetConditionTemplate()
		internal.PrintConditions(cmd, *condTemplate)
	},
}

func init() {
	rootCmd.AddCommand(condTemplateCmd)

	condTemplateCmd.Short = internal.GetMessageForKey(condTemplateCmd, internal.Short)
	condTemplateCmd.Long = internal.GetMessageForKey(condTemplateCmd, internal.Long)
}
