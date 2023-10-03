package service

import (
	"fmt"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/print"
	"github.com/akamai/cli-test-center/internal/util"
)

func (svc Service) EvaluateFunction(tryFunction *model.TryFunction) {
	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty, internalconstant.GetFunctionResult), !svc.jsonOutput).Start()
	tryFunctionResponse, err := svc.api.EvaluateFunction(tryFunction)

	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}
	spinner.StopWithSuccess()
	fmt.Println()
	if svc.jsonOutput {
		util.PrintJsonAndExit(tryFunctionResponse)
	}
	print.PrintTryFunctionResponse(tryFunctionResponse)
}
