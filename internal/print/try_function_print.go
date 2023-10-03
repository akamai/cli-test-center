package print

import (
	"fmt"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/util"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
)

func PrintTryFunctionResponse(tryFunction *model.TryFunction) {
	util.PrintLabelAndValue(externalconstant.LabelFunctionExpression, tryFunction.FunctionExpression)
	var result = tryFunction.Result
	if result == internalconstant.Empty {
		result = externalconstant.TryFunctionEmptyResult
	}
	util.PrintLabelAndValue(externalconstant.LabelFunctionResult, result)
	fmt.Println()
}
