package cmd

import (
	"github.com/akamai/cli-test-center/internal/api"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/service"
	"github.com/akamai/cli-test-center/internal/util"
	"github.com/akamai/cli-test-center/internal/validator"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/spf13/cobra"
)

var tryItCmd = &cobra.Command{
	Use:     externalconstant.TryItFunctionUse,
	Example: externalconstant.TryItFunctionExample,
	Aliases: []string{externalconstant.TryItFunctionAliases},
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		functionService := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, true)

		validator := validator.NewValidator(cmd, jsonData)

		validator.ValidateTryItFunctionInputFields(&tryFunction)

		functionService.EvaluateFunction(&tryFunction)

	},
}

func init() {

	functionCmd.AddCommand(tryItCmd)
	tryItCmd.Flags().SortFlags = false

	tryItCmd.Short = util.GetMessageForKey(tryItCmd, internalconstant.Short)
	tryItCmd.Long = util.GetMessageForKey(tryItCmd, internalconstant.Long)

}
