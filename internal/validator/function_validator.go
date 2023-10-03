package validator

import (
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/util"
)

//This validator class will have the test center functions related validate methods.

func (validator Validator) ValidateTryItFunctionInputFields(tryFunction *model.TryFunction) {

	if len(validator.jsonData) == 0 {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.Json), internalconstant.ExitStatusCode2)
	}

	if validator.jsonData != nil {
		util.ByteArrayToStruct(validator.cmd, validator.jsonData, &tryFunction)
		return
	}
}
