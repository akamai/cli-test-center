package validator

import (
	"fmt"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/util"
	"github.com/spf13/cobra"
)

type GlobalValidator struct {
	cmd      *cobra.Command
	jsonData []byte
}

func NewGlobalValidator(cmd *cobra.Command, jsonData []byte) *GlobalValidator {
	return &GlobalValidator{cmd, jsonData}
}

// Use this function to validate commands where json input is not allowed. Keep the order always one for parent commands
func (globalValidator GlobalValidator) ValidateJsonInputNotAllowed(cmd *cobra.Command, jsonData []byte) {
	if jsonData != nil {
		util.AbortWithUsageAndMessageAndCode(cmd, util.GetGlobalErrorMessage(internalconstant.JsonInputNotAllowed),
			internalconstant.ExitStatusCode2)
	}
}

// Use this function to validate root/parent subcommands/argument, gives suggestion as well if keywords match.
func (globalValidator GlobalValidator) ValidateParentSubCommands(cmd *cobra.Command, args []string, isJsonInputAllowed bool) {

	if !isJsonInputAllowed {
		globalValidator.ValidateJsonInputNotAllowed(cmd, globalValidator.jsonData)
	}

	err := util.LegacyArgs(cmd, args)
	if err != nil {
		util.AbortWithUsageAndMessageAndCode(globalValidator.cmd, err.Error(), internalconstant.ExitStatusCode2)
	} else {
		util.AbortWithUsageAndMessageAndCode(globalValidator.cmd, fmt.Errorf(internalconstant.Empty).Error(),
			internalconstant.ExitStatusCode0)
	}
}

// Use this function to validate commands where no subcommands are allowed, mostly used in child level commands.
func (globalValidator GlobalValidator) ValidateSubCommandsNotAllowed(cmd *cobra.Command, args []string, isJsonInputAllowed bool) {
	if !isJsonInputAllowed {
		globalValidator.ValidateJsonInputNotAllowed(cmd, globalValidator.jsonData)
	}

	if err := util.NoArgsCheck(cmd, args); err != nil {
		util.AbortWithUsageAndMessageAndCode(globalValidator.cmd, err.Error(), internalconstant.ExitStatusCode2)
	}
}
