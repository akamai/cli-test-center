package print

import (
	"fmt"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/util"
	"github.com/spf13/cobra"
	"strings"
)

func PrintConditionsTemplate(cmd *cobra.Command, condTemplate model.ConditionTemplate) {
	fmt.Println()
	util.PrintHeader(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty,
		internalconstant.ConditionsTemplatesText))
	fmt.Println()
	fmt.Println()
	for _, condType := range condTemplate.ConditionTypes {
		fmt.Println(util.Bold(condType.Label))
		for _, condExpression := range condType.ConditionExpressions {
			fmt.Println("  ", condExpression.ConditionExpression)
		}
		fmt.Println()
	}
}

func PrintConditions(cmd *cobra.Command, conditionList []model.Condition) {
	if len(conditionList) <= 0 {
		util.PrintWarning(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty,
			internalconstant.ConditionNotFoundError))
		fmt.Println()
		return
	}

	util.PrintHeader(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.Conditions))
	fmt.Println()
	fmt.Println()
	for _, condition := range conditionList {
		conditionExpression := strings.TrimSpace(condition.ConditionExpression)
		fmt.Println(conditionExpression)
	}
	fmt.Println()
	fmt.Print(len(conditionList))
	fmt.Println(" items")

}
