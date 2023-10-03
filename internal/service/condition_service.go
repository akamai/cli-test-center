package service

import (
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/print"
	"github.com/akamai/cli-test-center/internal/util"
	"github.com/spf13/cobra"
)

func (svc Service) GetConditionTemplateAndPrintResult(cmd *cobra.Command) {
	condTemplate := svc.GetConditionTemplate()
	print.PrintConditionsTemplate(cmd, *condTemplate)
}

func (svc Service) GetConditionTemplate() *model.ConditionTemplate {

	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner,
		internalconstant.Empty, internalconstant.GetConditionTemplate), !svc.jsonOutput).Start()
	condTemplate, err := svc.api.GetConditionTemplate()
	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}
	spinner.StopWithSuccess()

	if svc.jsonOutput {
		util.PrintJsonAndExit(condTemplate)
	}

	return condTemplate
}

func (svc Service) GetConditionListAndPrintResult(cmd *cobra.Command) {
	condList := svc.GetConditionList()
	print.PrintConditions(cmd, condList)
}

func (svc Service) GetConditionList() []model.Condition {
	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner, internalconstant.Empty,
		internalconstant.GetConditions), !svc.jsonOutput).Start()
	conditions, err := svc.api.GetConditionsV3()
	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}

	spinner.StopWithSuccess()

	if svc.jsonOutput {
		util.PrintJsonAndExit(conditions)
	}
	return conditions

}
