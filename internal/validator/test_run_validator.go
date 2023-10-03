package validator

import (
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/util"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

//This validator class will have the test run related validate methods.

func (validator Validator) ValidateGetTestRunFlag(testRunId string) {

	if testRunId == internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing,
			util.GetJsonKeyForFlag(externalconstant.FlagTestRunId)), internalconstant.ExitStatusCode2)
	}

	id, err := strconv.Atoi(testRunId)
	if testRunId != internalconstant.Empty && (err != nil || id == 0) {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid,
			util.GetJsonKeyForFlag(externalconstant.FlagTestRunId)), internalconstant.ExitStatusCode2)
	}
}

func (validator Validator) ValidateTestRunFlagsAndGetRunEnum(testSuiteId, testSuiteName, propertyId, propertyName, propVersion, url, condition, ipVersion,
	targetEnvironment, client, location, requestMethod, requestBody string, addHeader, modifyHeader []string, jsonData []byte, testRunRequest *model.TestRun,
	isStandardInputAvailable, encodeRequestBody bool) string {

	var runTestUsing = make([]string, 0)

	isJsonInput := util.CheckIfBothJsonAndFlagAreSetForCommand(validator.cmd, jsonData, isStandardInputAvailable)

	if isJsonInput {
		log.Debug("Test Run Using JSON Input!!!!!")
		util.ByteArrayToStruct(validator.cmd, validator.jsonData, testRunRequest)
		runTestUsing = append(runTestUsing, internalconstant.RunTestUsingJsonInput)
		return runTestUsing[0]
	}

	// Check if target environment flag value is from given enums only.
	environmentFlagCheck(validator.cmd, strings.ToUpper(targetEnvironment))

	if testSuiteName != internalconstant.Empty {
		runTestUsing = append(runTestUsing, internalconstant.RunTestUsingTestSuiteName)
	}

	if testSuiteId != internalconstant.Empty {
		// validate if other flag is not already set to run test
		checkExclusiveTestRunFlags(validator.cmd, runTestUsing)

		validator.EditTestSuiteIdFlagCheck(testSuiteId)
		runTestUsing = append(runTestUsing, internalconstant.RunTestUsingTestSuiteId)
	}

	if propertyName != internalconstant.Empty || propVersion != internalconstant.Empty {
		// validate if other flag is not already set to run test
		checkExclusiveTestRunFlags(validator.cmd, runTestUsing)

		validator.PropertyAndVersionFlagCheck(propertyId, propertyName, propVersion, true)
		runTestUsing = append(runTestUsing, internalconstant.RunTestUsingPropertyVersion)
	}

	if url != internalconstant.Empty || condition != internalconstant.Empty {
		// validate if other flag is not already set to run test
		checkExclusiveTestRunFlags(validator.cmd, runTestUsing)

		if url == internalconstant.Empty {
			util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, "url"), internalconstant.ExitStatusCode2)
		}

		if condition == internalconstant.Empty {
			util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, "condition"), internalconstant.ExitStatusCode2)
		}

		//Check if ip version is set properly
		ipVersionFlagCheck(validator.cmd, strings.ToUpper(ipVersion))

		//Check if headers are sent properly
		headerFlagCheck(validator.cmd, addHeader, modifyHeader)

		// Check if client flag value is from given enums only.
		clientTypeFlagCheck(validator.cmd, strings.ToUpper(client))

		// Check if location flag value is from given enums only.
		locationFlagCheck(validator.cmd, strings.ToUpper(location))

		// Check if requestMethod flag value is from given enums only.
		methodFlagCheck(validator.cmd, strings.ToUpper(requestMethod), strings.ToUpper(client))

		// Check if requestBody flag value is given for post request only.
		requestBodyFlagCheck(validator.cmd, requestBody, strings.ToUpper(requestMethod), strings.ToUpper(client))

		// Check if encodeRequestBody flag value is given for POST request and client type CURL only.
		encodeRequestBodyFlagCheck(validator.cmd, encodeRequestBody, strings.ToUpper(requestMethod), strings.ToUpper(client))

		runTestUsing = append(runTestUsing, internalconstant.RunTestUsingSingleTestCase)
	}

	switch len(runTestUsing) {
	case 0:
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.Any), internalconstant.ExitStatusCode2)
		return internalconstant.Empty
	case 1:
		return runTestUsing[0]
	default:
		// validate if other flag is not already set to run test
		checkExclusiveTestRunFlags(validator.cmd, runTestUsing)
		return internalconstant.Empty
	}
}

// Check if length of run test using list has more than one different king of flag.
func checkExclusiveTestRunFlags(cmd *cobra.Command, runTestUsing []string) {
	if len(runTestUsing) != 0 {
		util.AbortWithUsageAndMessageAndCode(cmd, util.GetErrorMessageForFlag(cmd, internalconstant.Invalid, internalconstant.Exclusive), internalconstant.ExitStatusCode2)
	}
}

func (validator Validator) ValidateGetRawRequestResponseFlag(testRunId, testCaseExecutionId string) {

	if testRunId == internalconstant.Empty && testCaseExecutionId == internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing,
			util.GetJsonKeyForFlag(internalconstant.Any)), internalconstant.ExitStatusCode2)
	}

	if testRunId != internalconstant.Empty && testCaseExecutionId != internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid,
			util.GetJsonKeyForFlag(internalconstant.OneOf)), internalconstant.ExitStatusCode2)
	}

	id, err := strconv.Atoi(testRunId)
	if testRunId != internalconstant.Empty && (err != nil || id == 0) {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid,
			util.GetJsonKeyForFlag(externalconstant.FlagTestRunId)), internalconstant.ExitStatusCode2)
	}

	tcxId, err := strconv.Atoi(testCaseExecutionId)
	if testCaseExecutionId != internalconstant.Empty && (err != nil || tcxId == 0) {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid,
			util.GetJsonKeyForFlag(externalconstant.FlagTestCaseExecId)), internalconstant.ExitStatusCode2)
	}
}

func (validator Validator) ValidateGetLogLinesFlag(tcxId string) {

	if tcxId == internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing,
			util.GetJsonKeyForFlag(externalconstant.FlagTestCaseExecId)), internalconstant.ExitStatusCode2)
	}

	id, err := strconv.Atoi(tcxId)
	if tcxId != internalconstant.Empty && (err != nil || id == 0) {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid,
			util.GetJsonKeyForFlag(externalconstant.FlagTestCaseExecId)), internalconstant.ExitStatusCode2)
	}

}

func locationFlagCheck(cmd *cobra.Command, location string) {
	if location != internalconstant.Empty && !(location == internalconstant.DefaultLocation) {
		util.AbortWithUsageAndMessageAndCode(cmd, util.GetErrorMessageForFlag(cmd, internalconstant.Invalid, externalconstant.FlagLocation), internalconstant.ExitStatusCode2)
	}
}

func environmentFlagCheck(cmd *cobra.Command, environment string) {

	if environment != internalconstant.Empty && !(environment == internalconstant.Staging || environment == internalconstant.Production) {
		util.AbortWithUsageAndMessageAndCode(cmd, util.GetErrorMessageForFlag(cmd, internalconstant.Invalid, internalconstant.TargetEnvironment), internalconstant.ExitStatusCode2)
	}
}
