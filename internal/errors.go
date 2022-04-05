package internal

import (
	"github.com/clarketm/json"
	log "github.com/sirupsen/logrus"
)

// ApiError object for API error payloads
type ApiError struct {
	Type          string        `json:"type"`
	Title         string        `json:"title"`
	Status        int           `json:"status"`
	Detail        string        `json:"detail,omitempty"`
	Code          string        `json:"code,omitempty"`
	Instance      string        `json:"instance"`
	Method        string        `json:"method,omitempty"`
	ServerIp      string        `json:"serverIp,omitempty"`
	ClientIp      string        `json:"clientIp,omitempty"`
	RequestId     string        `json:"requestId,omitempty"`
	RequestTime   string        `json:"requestTime,omitempty"`
	RequestField  string        `json:"requestField,omitempty"`
	RequestValues []interface{} `json:"requestValues,omitempty"`
	Errors        []ApiSubError `json:"errors"`
}

// ApiSubError object represents sub-errors of an error payload or error response in 207
type ApiSubError struct {
	Type             string        `json:"type"`
	Title            string        `json:"title"`
	MaxLimit         int           `json:"maxLimit"`
	RequestField     string        `json:"requestField"`
	RequestValues    []interface{} `json:"requestValues"`
	ExistingEntities []interface{} `json:"existingEntities"`
	RequirementId    int           `json:"requirementId,omitempty"`
	ConfigVersionId  int           `json:"configVersionId,omitempty"`
	TestSuiteId      int           `json:"testSuiteId,omitempty"`
	TestCaseId       int           `json:"testCaseId,omitempty"`
}

// CliError is used to transmit errors across the app
type CliError struct {
	apiError     *ApiError
	apiSubErrors []ApiSubError
	errorMessage string
	responseCode int
}

func CliErrorWithMessage(message string) *CliError {
	return &CliError{errorMessage: message}
}

func CliErrorFromPulsarProblemObject(apiErrorByte []byte, apiSubError []ApiSubError, responseCode int, fallbackMessage string) *CliError {
	var cliError CliError
	cliError.responseCode = responseCode

	if responseCode == 207 {
		cliError.apiSubErrors = apiSubError
	} else {
		var apiError ApiError
		apiParsingError := json.Unmarshal(apiErrorByte, &apiError)
		if apiParsingError != nil {
			cliError.errorMessage = fallbackMessage
			log.Debugf("Error while parsing api error response: [%s]", apiParsingError)
		} else {
			cliError.apiError = &apiError
		}
	}

	return &cliError
}
