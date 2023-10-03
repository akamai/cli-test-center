package validator

import (
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/util"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

//This validator class will have the test case related validate methods.

func (validator Validator) AddTestCaseToTestSuiteFlagCheck(testSuiteId, testSuiteName, url, condition, ipVersion string, addHeader,
	modifyHeader []string, client string, method string, requestBody string, encodeRequestBody bool, setVariables []string) {

	if url == internalconstant.Empty || condition == internalconstant.Empty || (testSuiteName == internalconstant.Empty && testSuiteId == internalconstant.Empty) {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.Any), internalconstant.ExitStatusCode2)
	}

	id, err := strconv.Atoi(testSuiteId)
	if testSuiteId != internalconstant.Empty && (err != nil || id == 0) {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.TestSuiteId), internalconstant.ExitStatusCode2)
	}

	//Check if ip version is sent properly
	ipVersionFlagCheck(validator.cmd, ipVersion)

	//Check if headers are sent properly
	headerFlagCheck(validator.cmd, addHeader, modifyHeader)

	// Check if client flag value is from given enums only.
	clientTypeFlagCheck(validator.cmd, strings.ToUpper(client))

	// Check if requestMethod flag value is from given enums only.
	methodFlagCheck(validator.cmd, strings.ToUpper(method), strings.ToUpper(client))

	// Check if requestBody flag value is given for post request only.
	requestBodyFlagCheck(validator.cmd, requestBody, strings.ToUpper(method), strings.ToUpper(client))

	// Check if encodeRequestBody flag value is given for POST request and client type CURL only.
	encodeRequestBodyFlagCheck(validator.cmd, encodeRequestBody, strings.ToUpper(method), strings.ToUpper(client))

	// Check if setVariables values are sent properly
	setVariablesFlagCheck(validator.cmd, setVariables)

}

func (validator Validator) RemoveTestCaseFromTestSuiteFlagCheck(testSuiteId, orderNumber, testCaseIdStr string) {

	if testSuiteId == internalconstant.Empty || (orderNumber == internalconstant.Empty && testCaseIdStr == internalconstant.Empty) {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.Any), internalconstant.ExitStatusCode2)
	}

	id, err := strconv.Atoi(testSuiteId)
	if testSuiteId != internalconstant.Empty && (err != nil || id == 0) {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.TestSuiteId), internalconstant.ExitStatusCode2)
	}

	if orderNumber != internalconstant.Empty && testCaseIdStr != internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.OrderNumTestCaseId), internalconstant.ExitStatusCode2)
	}

	orderNum, orderError := strconv.Atoi(orderNumber)
	if orderNumber != internalconstant.Empty && (orderError != nil || orderNum == 0) {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.OrderNum), internalconstant.ExitStatusCode2)
	}

	testCaseId, err := strconv.Atoi(testCaseIdStr)
	if testCaseIdStr != internalconstant.Empty && (err != nil || testCaseId == 0) {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.TestCaseId), internalconstant.ExitStatusCode2)
	}
}

func (validator Validator) ValidateTestCaseFlagCheck(testCaseId string) {
	if testCaseId == internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.TestCaseId), internalconstant.ExitStatusCode2)
	}

	tcId, err := strconv.Atoi(testCaseId)
	if testCaseId != internalconstant.Empty && (err != nil || tcId == 0) {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.TestCaseId), internalconstant.ExitStatusCode2)
	}
}

func setVariablesFlagCheck(cmd *cobra.Command, setVariables []string) {

	if len(setVariables) > 0 {
		for _, variable := range setVariables {
			setVariableComponents := strings.Split(variable, externalconstant.Colon)
			if len(setVariableComponents) < 2 {
				util.AbortWithUsageAndMessageAndCode(cmd, util.GetErrorMessageForFlag(cmd, internalconstant.Invalid, internalconstant.SetVariables), internalconstant.ExitStatusCode2)
			}
		}
	}
}
