package model

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
	MaxLimit         int           `json:"maxLimit,omitempty"`
	RequestField     string        `json:"requestField,omitempty"`
	RequestValues    []interface{} `json:"requestValues,omitempty"`
	ExistingEntities []interface{} `json:"existingEntities,omitempty"`
	RequirementId    int           `json:"requirementId,omitempty"`
	ConfigVersionId  int           `json:"configVersionId,omitempty"`
	TestSuiteId      int           `json:"testSuiteId,omitempty"`
	TestCaseId       int           `json:"testCaseId,omitempty"`
	Hostname         string        `json:"hostname,omitempty"`
	IpVersion        string        `json:"ipVersion,omitempty"`
	RequestObjects   []interface{} `json:"requestObjects,omitempty"`
	Errors           []ApiSubError `json:"errors,omitempty"`
}

// CliError is used to transmit errors across the app
type CliError struct {
	ApiError     *ApiError
	ApiSubErrors []ApiSubError
	ErrorMessage string
	ResponseCode int
}

func CliErrorWithMessage(message string) *CliError {
	return &CliError{ErrorMessage: message}
}

func CliErrorFromPulsarProblemObject(apiErrorByte []byte, apiSubError []ApiSubError, responseCode int, fallbackMessage string) *CliError {
	var cliError CliError
	cliError.ResponseCode = responseCode

	if responseCode == 207 {
		cliError.ApiSubErrors = apiSubError
	} else {
		var apiError ApiError
		apiParsingError := json.Unmarshal(apiErrorByte, &apiError)
		if apiParsingError != nil {
			cliError.ErrorMessage = fallbackMessage
			log.Debugf("Error while parsing api error response: [%s]", apiParsingError)
		} else {
			cliError.ApiError = &apiError
		}
	}

	return &cliError
}
