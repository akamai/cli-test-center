package internal

import (
	log "github.com/sirupsen/logrus"
	"net/url"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Validator struct {
	cmd      *cobra.Command
	jsonData []byte
}

func NewValidator(cmd *cobra.Command, jsonData []byte) *Validator {
	return &Validator{cmd, jsonData}
}

func (validator Validator) ConfigFlagCheck(configVersionString string, isOptional bool) (string, int) {

	if configVersionString != "" {
		configVersionSplit := strings.Split(configVersionString, " ")

		if len(configVersionSplit) != 2 {
			AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "configVersion"), ExitStatusCode2)
		}

		configName, version := configVersionSplit[0], configVersionSplit[1]

		versionNumber, err := strconv.Atoi(version)
		if err != nil || configName == "" {
			AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "configVersion"), ExitStatusCode2)
		}

		return configName, versionNumber
	}

	if !isOptional {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, "configVersion"), ExitStatusCode2)
	}

	return "", 0
}

func (validator Validator) ConfigFlagCheckForListTestSuites(propVersion string, propertyName string) {
	// validate configVersion usage
	if propVersion != Empty {
		version, err := strconv.Atoi(propVersion)

		if err != nil || propertyName == Empty || version == 0 {
			AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, PropertyVersionFlagKey), ExitStatusCode2)
		}
	}
}

func (validator Validator) RemoveConfigFlagCheck(property string, removeProperty bool) {
	if removeProperty && property != "" {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "removeProperty"), ExitStatusCode2)
	}
}

func (validator Validator) EditTestSuiteAllFlagCheck() {
	AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "all"), ExitStatusCode2)
}

func (validator Validator) AddTestSuiteNameFlagCheck(name string) {
	if name == "" {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, "name"), ExitStatusCode2)
	}
}

func (validator Validator) EditTestSuiteIdFlagCheck(testSuiteId string) {

	if testSuiteId == "" {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, "id"), ExitStatusCode2)
	}

	id, err := strconv.Atoi(testSuiteId)
	if err != nil || id == 0 {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "id"), ExitStatusCode2)
	}
}

func (validator Validator) ValidateViewTestSuiteFlags(testSuiteId, testSuiteName, groupBy string) {

	validator.TestSuiteIdAndNameFlagCheck(testSuiteId, testSuiteName)

	// validate if group by flag value is - condition, test-request, ipversion only
	validator.GroupByFlagCheck(groupBy)
}

func (validator Validator) GroupByFlagCheck(groupBy string) {

	if groupBy != "" && !ContainsInArray([]string{GroupByTestRequest, GroupByCondition, GroupByIpVersion}, groupBy) {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "groupBy"), ExitStatusCode2)
	}
}

func (validator Validator) ValidateImportFields(testSuitesImport *TestSuiteV3) {

	if len(validator.jsonData) == 0 {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, Json), ExitStatusCode2)
	}

	if validator.jsonData != nil {
		ByteArrayToStruct(validator.cmd, validator.jsonData, &testSuitesImport)
		return
	}
}

func (validator Validator) ValidateManageFields(testSuitesManage *TestSuiteV3) {

	if len(validator.jsonData) == 0 {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, Json), ExitStatusCode2)
	}

	if validator.jsonData != nil {
		ByteArrayToStruct(validator.cmd, validator.jsonData, &testSuitesManage)
		testSuiteId := strconv.Itoa(testSuitesManage.TestSuiteId)
		if testSuiteId == "0" {
			AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, "testSuiteId"), ExitStatusCode2)
		}
		return
	}
}

func (validator Validator) ValidateDefaultTestSuiteFields(propertyName string, propertyVersion string, urls []string, jsonData []byte,
	defaultTestSuite *DefaultTestSuiteRequest, isStandardInputAvailable bool) (string, int) {
	// throw invalid or missing errors if any and abort
	checkIfSubCommandFlagsAreNonEmpty := CheckIfBothJsonAndFlagAreSetForCommand(validator.cmd, jsonData, isStandardInputAvailable)

	if !checkIfSubCommandFlagsAreNonEmpty {
		log.Debug("Generating default test suite with input json data!!!!")
		ByteArrayToStruct(validator.cmd, validator.jsonData, defaultTestSuite)
		// the below returned values are not used anywhere for now as JSON payload is passed directly to the api
		return defaultTestSuite.Configs.PropertyManager.PropertyName, defaultTestSuite.Configs.PropertyManager.PropertyVersion
	} else {
		log.Debug("Generating default test suite with flags!!!!")
		// validate configVersion usage.
		name, version := validator.PropertyAndVersionFlagCheck(propertyName, propertyVersion, false)
		validator.UrlsFlagCheck(urls)
		return name, version
	}
}

func (validator Validator) AddTestCaseToTestSuiteFlagCheck(testSuiteId, testSuiteName, url, condition, ipVersion string, addHeader, modifyHeader []string) {

	if url == "" || condition == "" || (testSuiteName == "" && testSuiteId == "") {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, "any"), ExitStatusCode2)
	}

	id, err := strconv.Atoi(testSuiteId)
	if testSuiteId != "" && (err != nil || id == 0) {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "testSuiteId"), ExitStatusCode2)
	}

	//Check if ip version is sent properly
	ipVersionFlagCheck(validator.cmd, ipVersion)

	//Check if headers are sent properly
	headerFlagCheck(validator.cmd, addHeader, modifyHeader)
}

func (validator Validator) RemoveTestCaseFromTestSuiteFlagCheck(testSuiteId, orderNumber string) {

	if testSuiteId == "" || orderNumber == "" {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, "any"), ExitStatusCode2)
	}

	id, err := strconv.Atoi(testSuiteId)
	if testSuiteId != "" && (err != nil || id == 0) {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "testSuiteId"), ExitStatusCode2)
	}

	orderNum, orderError := strconv.Atoi(orderNumber)
	if orderNumber != "" && (orderError != nil || orderNum == 0) {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "orderNum"), ExitStatusCode2)
	}
}

func (validator Validator) TestSuiteIdAndNameFlagCheck(testSuiteId, testSuiteName string) {

	if testSuiteName == "" && testSuiteId == "" {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, "any"), ExitStatusCode2)
	}

	id, err := strconv.Atoi(testSuiteId)
	if testSuiteId != "" && (err != nil || id == 0) {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "id"), ExitStatusCode2)
	}

}

func (validator Validator) ValidateTestRunFlagsAndGetRunEnum(testSuiteId, testSuiteName, propertyName, propVersion,
	url, condition, ipVersion, targetEnvironment string, addHeader, modifyHeader []string, jsonData []byte,
	testRunRequest *TestRun, isStandardInputAvailable bool) string {

	var runTestUsing = make([]string, 0)

	checkIfSubCommandFlagsAreNonEmpty := CheckIfBothJsonAndFlagAreSetForCommand(validator.cmd, jsonData, isStandardInputAvailable)

	if !checkIfSubCommandFlagsAreNonEmpty {
		log.Debug("Test Run Using JSON Input!!!!!")
		ByteArrayToStruct(validator.cmd, validator.jsonData, testRunRequest)
		runTestUsing = append(runTestUsing, RunTestUsingJsonInput)
		return runTestUsing[0]
	}

	// Check if target environment flag value is from given enums only.
	environmentFlagCheck(validator.cmd, strings.ToUpper(targetEnvironment))

	if testSuiteName != "" {
		runTestUsing = append(runTestUsing, RunTestUsingTestSuiteName)
	}

	if testSuiteId != "" {
		// validate if other flag is not already set to run test
		checkExclusiveTestRunFlags(validator.cmd, runTestUsing)

		validator.EditTestSuiteIdFlagCheck(testSuiteId)
		runTestUsing = append(runTestUsing, RunTestUsingTestSuiteId)
	}

	if propertyName != "" || propVersion != "" {
		// validate if other flag is not already set to run test
		checkExclusiveTestRunFlags(validator.cmd, runTestUsing)

		validator.PropertyAndVersionFlagCheck(propertyName, propVersion, true)
		runTestUsing = append(runTestUsing, RunTestUsingPropertyVersion)
	}

	if url != "" || condition != "" {
		// validate if other flag is not already set to run test
		checkExclusiveTestRunFlags(validator.cmd, runTestUsing)

		if url == "" {
			AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, "url"), ExitStatusCode2)
		}

		if condition == "" {
			AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, "condition"), ExitStatusCode2)
		}

		//Check if ip version is set properly
		ipVersionFlagCheck(validator.cmd, strings.ToUpper(ipVersion))

		//Check if headers are sent properly
		headerFlagCheck(validator.cmd, addHeader, modifyHeader)

		runTestUsing = append(runTestUsing, RunTestUsingSingleTestCase)
	}

	switch len(runTestUsing) {
	case 0:
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, "any"), ExitStatusCode2)
		return ""
	case 1:
		return runTestUsing[0]
	default:
		// validate if other flag is not already set to run test
		checkExclusiveTestRunFlags(validator.cmd, runTestUsing)
		return ""
	}
}

func (validator Validator) UrlsFlagCheck(urls []string) {

	if len(urls) < 1 {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, "url"), ExitStatusCode2)
	}

	for _, fullUrl := range urls {

		u, err := url.Parse(fullUrl)
		if err != nil {
			AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "url"), ExitStatusCode2)
		}

		proto := u.Scheme
		hn := u.Host
		if proto == "" || hn == "" {
			AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "url"), ExitStatusCode2)
		}
	}
}

// Check if length of run test using list has more than one different king of flag.
func checkExclusiveTestRunFlags(cmd *cobra.Command, runTestUsing []string) {
	if len(runTestUsing) != 0 {
		AbortWithUsageAndMessageAndCode(cmd, GetErrorMessageForFlag(cmd, Invalid, "exclusive"), ExitStatusCode2)
	}
}

func environmentFlagCheck(cmd *cobra.Command, environment string) {

	if environment != "" && !(environment == Staging || environment == Production) {
		AbortWithUsageAndMessageAndCode(cmd, GetErrorMessageForFlag(cmd, Invalid, "targetEnvironment"), ExitStatusCode2)
	}
}

func ipVersionFlagCheck(cmd *cobra.Command, ipVersion string) {

	if ipVersion != "" && !(ipVersion == "v6" || ipVersion == "v4" || ipVersion == "V6" || ipVersion == "V4") {
		AbortWithUsageAndMessageAndCode(cmd, GetErrorMessageForFlag(cmd, Invalid, "ipVersion"), ExitStatusCode2)
	}
}

func headerFlagCheck(cmd *cobra.Command, addHeader, modifyHeader []string) {

	if len(addHeader) > 0 {
		for _, header := range addHeader {
			headerComponents := strings.Split(header, ":")
			if len(headerComponents) < 2 {
				AbortWithUsageAndMessageAndCode(cmd, GetErrorMessageForFlag(cmd, Invalid, "addHeader"), ExitStatusCode2)
			}
		}
	}

	if len(modifyHeader) > 0 {
		for _, header := range modifyHeader {
			headerComponents := strings.Split(header, ":")
			if len(headerComponents) < 2 {
				AbortWithUsageAndMessageAndCode(cmd, GetErrorMessageForFlag(cmd, Invalid, "modifyHeader"), ExitStatusCode2)
			}
		}
	}
}

func (validator Validator) PropertyAndVersionFlagCheck(propertyString string, versionString string, isOptional bool) (string, int) {
	if (propertyString != Empty && versionString == Empty) || (propertyString == Empty && versionString != Empty) {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, PropertyVersionFlagKey), ExitStatusCode2)
	}
	if propertyString != "" {
		versionNumber, err := strconv.Atoi(versionString)
		if err != nil {
			AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, PropertyVersionFlagKey), ExitStatusCode2)
		}
		return propertyString, versionNumber
	}
	if !isOptional {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Missing, PropertyVersionFlagKey), ExitStatusCode2)
	}
	return "", 0
}

func (validator Validator) ValidateSubcommandsNoArgCheck(cmd *cobra.Command, args []string) {
	if err := NoArgsCheck(cmd, args); err != nil {
		AbortWithUsageAndMessageAndCode(validator.cmd, err.Error(), ExitStatusCode2)
	}
}

func (validator Validator) NotValidSubcommandCheck(cmd *cobra.Command, args []string) {
	if err := NotValidSubcommandCheck(cmd, args); err != nil {
		AbortWithUsageAndMessageAndCode(validator.cmd, err.Error(), ExitStatusCode2)
	}
}

func (validator Validator) ValidSubcommandLegacyArgsCheck(cmd *cobra.Command, args []string) {
	if err := LegacyArgs(cmd, args); err != nil {
		AbortWithUsageAndMessageAndCode(validator.cmd, err.Error(), ExitStatusCode2)
	}
}

func (validator Validator) LockedAndStatefulFlagCheck(locked bool, unlocked bool, stateful bool, stateless bool) {
	if locked && unlocked {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "lockedUnlocked"), ExitStatusCode2)
	}
	if stateful && stateless {
		AbortWithUsageAndMessageAndCode(validator.cmd, GetErrorMessageForFlag(validator.cmd, Invalid, "statefulStateless"), ExitStatusCode2)
	}
}
