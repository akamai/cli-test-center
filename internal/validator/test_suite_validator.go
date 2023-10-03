package validator

import (
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/util"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	log "github.com/sirupsen/logrus"
	"net/url"
	"strconv"
)

//This validator class will have the test suite related validate methods.

func (validator Validator) EditTestSuiteAllFlagCheck() {
	util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.All), internalconstant.ExitStatusCode2)
}

func (validator Validator) AddTestSuiteNameFlagCheck(name string) {
	if name == internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, externalconstant.FlagVariableName), internalconstant.ExitStatusCode2)
	}
}

func (validator Validator) ValidateCreateTestSuiteFields(testSuite *model.TestSuite, isStandardInputAvailable bool, name, propertyId, propName, propVersion string) {

	// throw invalid or missing errors if any and abort
	isJsonInput := util.CheckIfBothJsonAndFlagAreSetForCommand(validator.cmd, validator.jsonData, isStandardInputAvailable)

	if isJsonInput {
		if len(validator.jsonData) == 0 {
			util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.Json), internalconstant.ExitStatusCode2)
		}

		if validator.jsonData != nil {
			util.ByteArrayToStruct(validator.cmd, validator.jsonData, &testSuite)
		}
	} else {
		// validate name flag check.
		validator.AddTestSuiteNameFlagCheck(name)

		// validate propertyVersion usage.
		validator.PropertyAndVersionFlagCheck(propertyId, propName, propVersion, true)
	}
}

func (validator Validator) ValidateUpdateTestSuiteFields(testSuite *model.TestSuite, isStandardInputAvailable bool, id,
	propertyId, propertyName, propertyVersion string, locked, unlocked, stateful, stateless, removeProperty bool) {

	// throw invalid or missing errors if any and abort
	isJsonInput := util.CheckIfBothJsonAndFlagAreSetForCommand(validator.cmd, validator.jsonData, isStandardInputAvailable)

	if isJsonInput {
		validator.ValidateManageFields(testSuite)
	} else {
		// validate id flag check.
		validator.EditTestSuiteIdFlagCheck(id)

		// validate propertyVersion usage.
		validator.PropertyAndVersionFlagCheck(propertyId, propertyName, propertyVersion, true)

		//Remove config flag check
		validator.RemoveConfigFlagCheck(propertyName, removeProperty)
		validator.LockedAndStatefulFlagCheck(locked, unlocked, stateful, stateless)
	}
}

func (validator Validator) RemoveConfigFlagCheck(propertyName string, removeProperty bool) {
	if removeProperty && propertyName != internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.RemoveProperty), internalconstant.ExitStatusCode2)
	}
}

func (validator Validator) ValidateImportFields(testSuitesImport *model.TestSuite) {

	if len(validator.jsonData) == 0 {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.Json), internalconstant.ExitStatusCode2)
	}

	if validator.jsonData != nil {
		util.ByteArrayToStruct(validator.cmd, validator.jsonData, &testSuitesImport)
		return
	}
}

func (validator Validator) ValidateManageFields(testSuitesManage *model.TestSuite) {

	if len(validator.jsonData) == 0 {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.Json), internalconstant.ExitStatusCode2)
	}

	if validator.jsonData != nil {
		util.ByteArrayToStruct(validator.cmd, validator.jsonData, &testSuitesManage)
		testSuiteId := strconv.Itoa(testSuitesManage.TestSuiteId)
		if testSuiteId == "0" {
			util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.TestSuiteId), internalconstant.ExitStatusCode2)
		}
		return
	}
}

func (validator Validator) ValidateDefaultTestSuiteFields(propertyId, propertyName, propertyVersion string, urls []string, jsonData []byte,
	defaultTestSuite *model.DefaultTestSuiteRequest, isStandardInputAvailable bool) {
	// throw invalid or missing errors if any and abort
	isJsonInput := util.CheckIfBothJsonAndFlagAreSetForCommand(validator.cmd, jsonData, isStandardInputAvailable)

	if isJsonInput {
		log.Debug("Generating default test suite with input json data!!!!")
		util.ByteArrayToStruct(validator.cmd, validator.jsonData, defaultTestSuite)
		// the below returned values are not used anywhere for now as JSON payload is passed directly to the api
	} else {
		log.Debug("Generating default test suite with flags!!!!")
		// validate configVersion usage.
		validator.PropertyAndVersionFlagCheck(propertyId, propertyName, propertyVersion, false)
		validator.UrlsFlagCheck(urls)
	}
}

func (validator Validator) LockedAndStatefulFlagCheck(locked bool, unlocked bool, stateful bool, stateless bool) {
	if locked && unlocked {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.LockedUnlocked), internalconstant.ExitStatusCode2)
	}
	if stateful && stateless {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.StatefulStateless), internalconstant.ExitStatusCode2)
	}
}

func (validator Validator) UrlsFlagCheck(urls []string) {

	if len(urls) < 1 {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, externalconstant.FlagUrl), internalconstant.ExitStatusCode2)
	}

	for _, fullUrl := range urls {

		u, err := url.Parse(fullUrl)
		if err != nil {
			util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, externalconstant.FlagUrl), internalconstant.ExitStatusCode2)
		}

		proto := u.Scheme
		hn := u.Host
		if proto == internalconstant.Empty || hn == internalconstant.Empty {
			util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, externalconstant.FlagUrl), internalconstant.ExitStatusCode2)
		}
	}
}

func (validator Validator) GroupByFlagCheck(groupBy string) {

	if groupBy != internalconstant.Empty && !util.ContainsInArray([]string{internalconstant.GroupByTestRequest, internalconstant.GroupByCondition, internalconstant.GroupByClientProfile}, groupBy) {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.GroupBy), internalconstant.ExitStatusCode2)
	}
}
