package validator

import (
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/util"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"strconv"
	"strings"
)

//This validator class will have the variables related validate methods.

func (validator Validator) ValidateVariableCreateFlagCheck(testSuiteId, name, value string, group []string) {

	if testSuiteId == internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.Id), internalconstant.ExitStatusCode2)
	}

	id, err := strconv.Atoi(testSuiteId)
	if testSuiteId != internalconstant.Empty && (err != nil || id == 0) {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.Id), internalconstant.ExitStatusCode2)
	}

	if name == internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.VariableName), internalconstant.ExitStatusCode2)
	}

	if value == internalconstant.Empty && group == nil {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.Any), internalconstant.ExitStatusCode2)
	}

	if group != nil {
		for _, groupString := range group {
			parts := strings.Split(groupString, externalconstant.Colon)
			if len(parts) != 2 {
				util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.VariableGroup), internalconstant.ExitStatusCode2)
			}
		}
	}
}

func (validator Validator) ValidateVariablesListFlagCheck(testSuiteId string) {

	if testSuiteId == internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.Id), internalconstant.ExitStatusCode2)
	}

	id, err := strconv.Atoi(testSuiteId)
	if testSuiteId != internalconstant.Empty && (err != nil || id == 0) {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.Id), internalconstant.ExitStatusCode2)
	}

}

func (validator Validator) ValidateVariableFlagCheck(testSuiteId, variableId string) {

	if testSuiteId == internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.Id), internalconstant.ExitStatusCode2)
	}

	id, err := strconv.Atoi(testSuiteId)
	if testSuiteId != internalconstant.Empty && (err != nil || id == 0) {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.Id), internalconstant.ExitStatusCode2)
	}

	if variableId == internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.VariableId), internalconstant.ExitStatusCode2)
	}

	varId, err := strconv.Atoi(variableId)
	if variableId != internalconstant.Empty && (err != nil || varId == 0) {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.VariableId), internalconstant.ExitStatusCode2)
	}
}

func (validator Validator) ValidateVariableEditFlagCheck(testSuiteId, name, value, variableId string, group []string) {

	validator.ValidateVariableCreateFlagCheck(testSuiteId, name, value, group)

	if variableId == internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.VariableId), internalconstant.ExitStatusCode2)
	}

	varId, err := strconv.Atoi(variableId)
	if variableId != internalconstant.Empty && (err != nil || varId == 0) {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.VariableId), internalconstant.ExitStatusCode2)
	}
}
