package validator

import (
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/util"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Validator struct {
	cmd      *cobra.Command
	jsonData []byte
}

// This validator class will have generic validation methods.

func NewValidator(cmd *cobra.Command, jsonData []byte) *Validator {
	return &Validator{cmd, jsonData}
}

func ipVersionFlagCheck(cmd *cobra.Command, ipVersion string) {

	if ipVersion != internalconstant.Empty && !(ipVersion == "v6" || ipVersion == "v4" || ipVersion == "V6" || ipVersion == "V4") {
		util.AbortWithUsageAndMessageAndCode(cmd, util.GetErrorMessageForFlag(cmd, internalconstant.Invalid, "ipVersion"), internalconstant.ExitStatusCode2)
	}
}

func clientTypeFlagCheck(cmd *cobra.Command, client string) {
	if client != internalconstant.Empty && !(client == internalconstant.Chrome || client == internalconstant.Curl) {
		util.AbortWithUsageAndMessageAndCode(cmd, util.GetErrorMessageForFlag(cmd, internalconstant.Invalid, externalconstant.FlagClient), internalconstant.ExitStatusCode2)
	}
}

func methodFlagCheck(cmd *cobra.Command, method string, client string) {
	if method != internalconstant.Empty && !(method == internalconstant.GetRequestMethod || method == internalconstant.HeadRequestMethod || method == internalconstant.PostRequestMethod) {
		util.AbortWithUsageAndMessageAndCode(cmd, util.GetErrorMessageForFlag(cmd, internalconstant.Invalid, util.GetJsonKeyForFlag(externalconstant.FlagRequestMethod)), internalconstant.ExitStatusCode2)
	}

	if client == internalconstant.Chrome && (method == internalconstant.HeadRequestMethod || method == internalconstant.PostRequestMethod) {
		util.AbortWithUsageAndMessageAndCode(cmd, util.GetErrorMessageForFlag(cmd, internalconstant.Invalid, internalconstant.RequestMethodWithClient), internalconstant.ExitStatusCode2)
	}
}

func requestBodyFlagCheck(cmd *cobra.Command, requestBody string, method string, client string) {
	if requestBody != internalconstant.Empty && method != internalconstant.PostRequestMethod && client != internalconstant.Curl {
		util.AbortWithUsageAndMessageAndCode(cmd, util.GetErrorMessageForFlag(cmd, internalconstant.Invalid, util.GetJsonKeyForFlag(externalconstant.FlagRequestBody)), internalconstant.ExitStatusCode2)
	}
}

func encodeRequestBodyFlagCheck(cmd *cobra.Command, encodeRequestBody bool, method string, client string) {
	if encodeRequestBody == true && (method != internalconstant.PostRequestMethod || client != internalconstant.Curl) {
		util.AbortWithUsageAndMessageAndCode(cmd, util.GetErrorMessageForFlag(cmd, internalconstant.Invalid, util.GetJsonKeyForFlag(externalconstant.FlagEncodeRequestBody)), internalconstant.ExitStatusCode2)
	}
}

func headerFlagCheck(cmd *cobra.Command, addHeader, modifyHeader []string) {

	if len(addHeader) > 0 {
		for _, header := range addHeader {
			headerComponents := strings.Split(header, externalconstant.Colon)
			if len(headerComponents) < 2 {
				util.AbortWithUsageAndMessageAndCode(cmd, util.GetErrorMessageForFlag(cmd, internalconstant.Invalid, internalconstant.AddHeader), internalconstant.ExitStatusCode2)
			}
		}
	}

	if len(modifyHeader) > 0 {
		for _, header := range modifyHeader {
			headerComponents := strings.Split(header, externalconstant.Colon)
			if len(headerComponents) < 2 {
				util.AbortWithUsageAndMessageAndCode(cmd, util.GetErrorMessageForFlag(cmd, internalconstant.Invalid, internalconstant.ModifyHeader), internalconstant.ExitStatusCode2)
			}
		}
	}
}

func (validator Validator) PropertyAndVersionFlagCheck(propertyId string, propertyString string, versionString string, isOptional bool) {

	if propertyId != internalconstant.Empty && propertyString != internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.PropertyCombinationErrorKey), internalconstant.ExitStatusCode2)
	}

	if propertyId != internalconstant.Empty && versionString == internalconstant.Empty && propertyString == internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.MissingVersionWithId), internalconstant.ExitStatusCode2)
	}

	if propertyString != internalconstant.Empty && versionString == internalconstant.Empty && propertyId == internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.MissingVersionWithName), internalconstant.ExitStatusCode2)
	}

	if versionString != internalconstant.Empty && propertyId == internalconstant.Empty && propertyString == internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.MissingIdOrNameWithVersion), internalconstant.ExitStatusCode2)
	}

	if propertyId != internalconstant.Empty || propertyString != internalconstant.Empty {
		if propertyId != internalconstant.Empty {
			_, err := strconv.Atoi(propertyId)
			if err != nil {
				util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.PropertyIdErrorKey), internalconstant.ExitStatusCode2)
			}
		}

		_, err := strconv.Atoi(versionString)
		if err != nil {
			util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.PropertyVersionErrorKey), internalconstant.ExitStatusCode2)
		}

		// return if everything is valid and property flags are optional.
		return
	}

	if !isOptional {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.PropertyVersionKey), internalconstant.ExitStatusCode2)
	}
}

func (validator Validator) EditTestSuiteIdFlagCheck(testSuiteId string) {

	if testSuiteId == internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.Id), internalconstant.ExitStatusCode2)
	}

	id, err := strconv.Atoi(testSuiteId)
	if err != nil || id == 0 {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.Id), internalconstant.ExitStatusCode2)
	}
}

func (validator Validator) ValidateGetTestSuiteWithChildObjectsFlags(testSuiteId, testSuiteName, groupBy string) {

	validator.TestSuiteIdAndNameFlagCheck(testSuiteId, testSuiteName)

	// validate if group by flag value is - condition, test-request, client-profile only
	validator.GroupByFlagCheck(groupBy)
}

func (validator Validator) TestSuiteIdAndNameFlagCheck(testSuiteId, testSuiteName string) {

	if testSuiteName == internalconstant.Empty && testSuiteId == internalconstant.Empty {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Missing, internalconstant.Any), internalconstant.ExitStatusCode2)
	}

	id, err := strconv.Atoi(testSuiteId)
	if testSuiteId != internalconstant.Empty && (err != nil || id == 0) {
		util.AbortWithUsageAndMessageAndCode(validator.cmd, util.GetErrorMessageForFlag(validator.cmd, internalconstant.Invalid, internalconstant.Id), internalconstant.ExitStatusCode2)
	}

}
