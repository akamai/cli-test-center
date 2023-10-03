package api

import (
	"fmt"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/util"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"net/http"
	"net/url"
	"strconv"

	"github.com/clarketm/json"
	log "github.com/sirupsen/logrus"
)

type ApiClient struct {
	client EdgeGridHttpClient
}

func NewApiClient(client EdgeGridHttpClient) *ApiClient {
	return &ApiClient{client}
}

// Adds query parameters to url.
// QueryMap entries with empty values are ignored
func addQueryParams(url *url.URL, queryMap map[string]string) {
	queryParams := url.Query()

	log.Debug("Adding query parameters to url %s", url)
	for k, v := range queryMap {
		log.Tracef("Processing query parameter - [%s]:[%s]", k, v)

		if v != "" {
			queryParams.Set(k, v)
		}
	}

	url.RawQuery = queryParams.Encode()
	log.Tracef("Url with query parameters: %s", url.String())
}

func (api ApiClient) SubmitTestRun(testRun model.TestRun) (*model.TestRun, *model.CliError) {
	testRunBytes, err := json.Marshal(testRun)
	if err != nil {
		log.Error(err)
		util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.RequestParsingError), internalconstant.ExitStatusCode3)
	}

	var requestHeaders = make(http.Header)
	requestHeaders.Add(internalconstant.ContentType, internalconstant.ApplicationJson)

	resp, byt := api.client.request(internalconstant.Post, "/test-management/v3/test-runs", &testRunBytes, requestHeaders)
	if resp.StatusCode == 202 {
		var testRunResponse model.TestRun
		err := json.Unmarshal(*byt, &testRunResponse)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return &testRunResponse, nil
	}
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorSubmitTestRunPostCall)
}

func (api ApiClient) GetTestRun(testRunId int) (*model.TestRun, *model.CliError) {
	resp, byt := api.client.request(http.MethodGet, fmt.Sprintf("/test-management/v3/test-runs/%d", testRunId), nil, nil)
	if resp.StatusCode == 200 {
		var testRun model.TestRun
		err := json.Unmarshal(*byt, &testRun)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return &testRun, nil
	}
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorTestRunGetCall)
}

func (api ApiClient) GetTestRunContext(testRunId int) (*model.TestRunContext, *model.CliError) {
	resp, byt := api.client.request(http.MethodGet, fmt.Sprintf("/test-management/v3/test-runs/%d/context", testRunId), nil, nil)
	if resp.StatusCode == 200 {
		var testRunContext model.TestRunContext
		err := json.Unmarshal(*byt, &testRunContext)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return &testRunContext, nil
	}
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorTestRunContextGetCall)
}

func (api ApiClient) GetTestSuites(propertyId, propertyName, propVersion, user string, includeDeleted bool) ([]model.TestSuite, *model.CliError) {

	v3Path := "/test-management/v3/functional/test-suites?includeRecentlyDeleted=" + strconv.FormatBool(includeDeleted)
	tsV3Url, _ := url.Parse(v3Path)

	// add optional query parameters
	queryMap := map[string]string{
		"propertyId":      propertyId,
		"propertyName":    propertyName,
		"propertyVersion": propVersion,
		"user":            user,
	}
	addQueryParams(tsV3Url, queryMap)

	// get response
	resp, byt := api.client.request(internalconstant.Get, tsV3Url.String(), nil, nil)

	// parse result
	if resp.StatusCode == 200 {
		var testSuites model.ListResponse
		err := json.Unmarshal(*byt, &testSuites)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}

		log.Infof("GetTestSuites [%s] returned %d items", tsV3Url.String(), len(testSuites.TestSuites))
		return testSuites.TestSuites, nil
	}

	// if not 200 response
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorTestSuiteV3GetCall)
}

func (api ApiClient) GetTestRuns() ([]model.TestRun, *model.CliError) {

	v3Path := "/test-management/v3/test-runs"
	testRunsV3Url, _ := url.Parse(v3Path)

	// get response
	resp, byt := api.client.request(internalconstant.Get, testRunsV3Url.String(), nil, nil)

	// parse result
	if resp.StatusCode == 200 {
		var listResponse model.ListResponse
		err := json.Unmarshal(*byt, &listResponse)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}

		log.Infof("GetTestRuns: [%s] returned %d items", testRunsV3Url.String(), len(listResponse.TestRuns))
		return listResponse.TestRuns, nil
	}

	// if not 200 response
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorTestRunsGetCall)
}

func (api ApiClient) GetTestSuiteV3(id string) (*model.TestSuite, *model.CliError) {

	v3Path := "/test-management/v3/functional/test-suites/" + id

	// get response
	resp, byt := api.client.request(internalconstant.Get, v3Path, nil, nil)

	// parse result
	if resp.StatusCode == 200 {
		var testSuite model.TestSuite
		err := json.Unmarshal(*byt, &testSuite)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}

		return &testSuite, nil
	}

	// if not 200 response
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorTestSuiteV3GetCall)
}

func (api ApiClient) RemoveTestSuite(id string) *model.CliError {

	v3Path := "/test-management/v3/functional/test-suites/" + id

	// get response
	resp, byt := api.client.request(internalconstant.Delete, v3Path, nil, nil)

	// parse result
	if resp.StatusCode == 204 {
		return nil
	}

	// if not 204 response
	return model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorRemoveTestSuiteDeleteCall)
}

func (api ApiClient) RestoreTestSuite(id string) (*model.TestSuite, *model.CliError) {

	v3Path := "/test-management/v3/functional/test-suites/" + id + "/restore"

	// get response
	resp, byt := api.client.request(internalconstant.Post, v3Path, nil, nil)

	// parse result
	if resp.StatusCode == 200 {
		var testSuite model.TestSuite
		err := json.Unmarshal(*byt, &testSuite)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}

		return &testSuite, nil
	}

	// if not 200 response
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorRestoreTestSuitePostCall)
}

func (api ApiClient) AddTestSuitesV3(testSuiteV3 model.TestSuite) (*model.TestSuite, *model.CliError) {
	testSuiteBytes, err := json.Marshal(testSuiteV3)
	if err != nil {
		log.Error(err)
		util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.RequestParsingError), internalconstant.ExitStatusCode3)
	}

	var requestHeaders = make(http.Header)
	requestHeaders.Add(internalconstant.ContentType, internalconstant.ApplicationJson)

	v3Path := "/test-management/v3/functional/test-suites"

	// get response
	resp, byt := api.client.request(internalconstant.Post, v3Path, &testSuiteBytes, requestHeaders)

	// parse result
	if resp.StatusCode == 201 {
		var testSuite model.TestSuite
		err := json.Unmarshal(*byt, &testSuite)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}

		return &testSuite, nil
	}

	// if not 201 response
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorTestSuiteV3PostCall)
}

func (api ApiClient) EditTestSuitesV3(testSuiteV3 model.TestSuite, id string) (*model.TestSuite, *model.CliError) {
	testSuiteBytes, err := json.Marshal(testSuiteV3)
	if err != nil {
		log.Error(err)
		util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.RequestParsingError), internalconstant.ExitStatusCode3)
	}

	var requestHeaders = make(http.Header)
	requestHeaders.Add(internalconstant.ContentType, internalconstant.ApplicationJson)

	v3Path := "/test-management/v3/functional/test-suites/" + id

	// get response
	resp, byt := api.client.request(internalconstant.Put, v3Path, &testSuiteBytes, requestHeaders)

	// parse result
	if resp.StatusCode == 200 {
		var testSuite model.TestSuite
		err := json.Unmarshal(*byt, &testSuite)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}

		return &testSuite, nil
	}

	// if not 200 response
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorEditTestSuiteV3PostCall)
}

func (api ApiClient) GetTestSuitesWithChildObjects(testSuiteId int, resolveVariables bool) (*model.TestSuite, *model.CliError) {

	v3Path := fmt.Sprintf("/test-management/v3/functional/test-suites/%d/with-child-objects?resolveVariables="+strconv.FormatBool(resolveVariables), testSuiteId)
	resp, byt := api.client.request(internalconstant.Get, v3Path, nil, nil)

	if resp.StatusCode == 200 {
		var testSuite model.TestSuite
		err := json.Unmarshal(*byt, &testSuite)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return &testSuite, nil
	}
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorTestSuiteWithChildObjectGetCall)
}

func (api ApiClient) AddTestCaseToTestSuite(testSuiteId int, testCases []model.TestCase) (*model.TestCaseBulkResponse, *model.CliError) {

	testCasesBytes, err := json.Marshal(testCases)
	if err != nil {
		log.Error(err)
		util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.RequestParsingError), internalconstant.ExitStatusCode3)
	}

	path := fmt.Sprintf("/test-management/v3/functional/test-suites/%d/test-cases", testSuiteId)
	var requestHeaders = make(http.Header)
	requestHeaders.Add(internalconstant.ContentType, internalconstant.ApplicationJson)

	resp, byt := api.client.request(internalconstant.Post, path, &testCasesBytes, requestHeaders)
	if resp.StatusCode == 207 {
		var createTestCaseResponse model.TestCaseBulkResponse
		err := json.Unmarshal(*byt, &createTestCaseResponse)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return &createTestCaseResponse, nil
	}
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorAddTestCasesToTestSuitPostCall)
}

func (api ApiClient) EditTestCaseToTestSuite(testSuiteId int, testCases []model.TestCase) (*model.TestCaseBulkResponse, *model.CliError) {

	testCasesBytes, err := json.Marshal(testCases)
	if err != nil {
		log.Error(err)
		util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.RequestParsingError), internalconstant.ExitStatusCode3)
	}

	path := fmt.Sprintf("/test-management/v3/functional/test-suites/%d/test-cases", testSuiteId)
	var requestHeaders = make(http.Header)
	requestHeaders.Add(internalconstant.ContentType, internalconstant.ApplicationJson)

	resp, byt := api.client.request(internalconstant.Put, path, &testCasesBytes, requestHeaders)
	if resp.StatusCode == 207 {
		var createTestCaseResponse model.TestCaseBulkResponse
		err := json.Unmarshal(*byt, &createTestCaseResponse)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return &createTestCaseResponse, nil
	}
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorUpdateTestCasesToTestSuitPutCall)
}

func (api ApiClient) ImportTestSuite(testSuiteImport model.TestSuite) (*model.TestSuiteImportResponse, *model.CliError) {

	testSuiteImportJson, err := json.Marshal(testSuiteImport)
	if err != nil {
		log.Error(err)
		util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.RequestParsingError), internalconstant.ExitStatusCode3)
	}

	path := "/test-management/v3/functional/test-suites/with-child-objects"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(internalconstant.ContentType, internalconstant.ApplicationJson)

	resp, byt := api.client.request(internalconstant.Post, path, &testSuiteImportJson, requestHeaders)
	if resp.StatusCode == 207 {
		var testSuiteImportResponseV3 model.TestSuiteImportResponse
		err := json.Unmarshal(*byt, &testSuiteImportResponseV3)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return &testSuiteImportResponseV3, nil
	}
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorImportTestSuiteV3PostCall)
}

func (api ApiClient) ManageTestSuite(testSuiteManage model.TestSuite, testSuiteId int) (*model.TestSuite, *model.CliError) {

	testSuiteManageJson, err := json.Marshal(testSuiteManage)
	if err != nil {
		log.Error(err)
		util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.RequestParsingError), internalconstant.ExitStatusCode3)
	}

	path := fmt.Sprintf("/test-management/v3/functional/test-suites/%d/with-child-objects", testSuiteId)
	var requestHeaders = make(http.Header)
	requestHeaders.Add(internalconstant.ContentType, internalconstant.ApplicationJson)

	resp, byt := api.client.request(internalconstant.Put, path, &testSuiteManageJson, requestHeaders)
	if resp.StatusCode == 200 {
		var testSuiteManageResponseV3 model.TestSuite
		err := json.Unmarshal(*byt, &testSuiteManageResponseV3)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return &testSuiteManageResponseV3, nil
	}
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorManageTestSuiteV3PutCall)
}

func (api ApiClient) RemoveTestCasesFromTestSuite(testSuiteId int, testCaseIds []int) (*model.BulkResponse, *model.CliError) {

	testCasesIdsBytes, err := json.Marshal(testCaseIds)
	if err != nil {
		log.Error(err)
		util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.RequestParsingError), internalconstant.ExitStatusCode3)
	}

	path := fmt.Sprintf("/test-management/v3/functional/test-suites/%d/test-cases/remove", testSuiteId)
	var requestHeaders = make(http.Header)
	requestHeaders.Add(internalconstant.ContentType, internalconstant.ApplicationJson)

	resp, byt := api.client.request(internalconstant.Post, path, &testCasesIdsBytes, requestHeaders)
	if resp.StatusCode == 207 {
		var removeTestCaseResponse model.BulkResponse
		err := json.Unmarshal(*byt, &removeTestCaseResponse)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return &removeTestCaseResponse, nil
	}

	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorRemoveTestCasesPostCall)
}

func (api ApiClient) GetConditionTemplate() (*model.ConditionTemplate, *model.CliError) {

	v3Path := "/test-management/v3/functional/test-catalog/template"
	condV3Url, _ := url.Parse(v3Path)

	// get response
	resp, byt := api.client.request(internalconstant.Get, condV3Url.String(), nil, nil)

	// parse result
	if resp.StatusCode == 200 {
		var condTemplate model.ConditionTemplate
		err := json.Unmarshal(*byt, &condTemplate)
		if err != nil {
			log.Error(err)
			fmt.Println(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}

		return &condTemplate, nil
	}

	// if not 200 response
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorConditionTemplateGetCall)
}

func (api ApiClient) GenerateTestSuite(defaultTsRequest model.DefaultTestSuiteRequest) (*model.TestSuite, *model.CliError) {

	defaultTestSuiteReqJson, err := json.Marshal(defaultTsRequest)
	if err != nil {
		log.Error(err)
		util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.RequestParsingError), internalconstant.ExitStatusCode3)
	}

	path := "/test-management/v3/functional/test-suites/auto-generate"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(internalconstant.ContentType, internalconstant.ApplicationJson)

	resp, byt := api.client.request(internalconstant.Post, path, &defaultTestSuiteReqJson, requestHeaders)

	if resp.StatusCode == 200 {
		var defaultTsResponse model.TestSuite
		err := json.Unmarshal(*byt, &defaultTsResponse)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return &defaultTsResponse, nil
	}
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorAutoGeneratePostCall)
}

func (api ApiClient) GetV3AssociatedTestCasesForTestSuite(testSuiteId int, resolveVariables bool) (*model.AssociatedTestCases, *model.CliError) {

	path := fmt.Sprintf("/test-management/v3/functional/test-suites/%d/test-cases?resolveVariables="+strconv.FormatBool(resolveVariables), testSuiteId)

	resp, byt := api.client.request(internalconstant.Get, path, nil, nil)
	if resp.StatusCode == 200 {
		var associatedTestCases model.AssociatedTestCases
		err := json.Unmarshal(*byt, &associatedTestCases)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return &associatedTestCases, nil
	}
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorTestCasesAssociatedForTestSuitGetCall)
}

func (api ApiClient) GetTestCaseById(testCaseId int, testSuiteId int, resolveVariables bool) (*model.TestCase, *model.CliError) {

	v3Path := fmt.Sprintf("/test-management/v3/functional/test-suites/%d/test-cases/%d?resolveVariables="+strconv.FormatBool(resolveVariables), testSuiteId, testCaseId)

	// get response
	resp, byt := api.client.request(internalconstant.Get, v3Path, nil, nil)

	// parse result
	if resp.StatusCode == 200 {
		var testCase model.TestCase
		err := json.Unmarshal(*byt, &testCase)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}

		return &testCase, nil
	}

	// if not 200 response
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorTestCaseForTestCaseGetCall)
}

func (api ApiClient) GetConditionsV3() ([]model.Condition, *model.CliError) {

	v3Path := "/test-management/v3/functional/test-catalog/conditions"
	condV3Url, _ := url.Parse(v3Path)

	// get response
	resp, byt := api.client.request(internalconstant.Get, condV3Url.String(), nil, nil)

	// parse result
	if resp.StatusCode == 200 {

		var listResponse model.ListResponse
		err := json.Unmarshal(*byt, &listResponse)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}

		log.Infof("GetConditions: [%s] returned %d items", condV3Url.String(), len(listResponse.Conditions))
		return listResponse.Conditions, nil
	}

	// if not 200 response
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorGetConditions)
}

func (api ApiClient) GetTestRequestsV3() ([]model.TestRequest, *model.CliError) {

	v3Path := "/test-management/v3/functional/test-requests"
	testReqV3Url, _ := url.Parse(v3Path)

	// get response
	resp, byt := api.client.request(internalconstant.Get, testReqV3Url.String(), nil, nil)

	if resp.StatusCode == 200 {
		var listResponse model.ListResponse
		err := json.Unmarshal(*byt, &listResponse)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}

		log.Infof("GetTestRequests: [%s] returned %d items", testReqV3Url.String(), len(listResponse.TestRequests))
		return listResponse.TestRequests, nil
	}

	// if not 200 response
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorGetTestRequests)
}

func (api ApiClient) CreateVariables(variable []model.Variable, testSuiteId string) (*model.VariableBulkResponse, *model.CliError) {
	variables, err := json.Marshal(variable)
	if err != nil {
		log.Error(err)
		util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.RequestParsingError), internalconstant.ExitStatusCode3)
	}

	var requestHeaders = make(http.Header)
	requestHeaders.Add(internalconstant.ContentType, internalconstant.ApplicationJson)

	v3Path := "/test-management/v3/functional/test-suites/" + testSuiteId + "/variables"

	// get response
	resp, byt := api.client.request(internalconstant.Post, v3Path, &variables, requestHeaders)

	if resp.StatusCode == 207 {
		var br model.VariableBulkResponse
		err := json.Unmarshal(*byt, &br)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return &br, nil
	}

	// if not 207 response
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorVariablePostCall)
}

func (api ApiClient) GetVariables(testSuiteId string) ([]model.Variable, *model.CliError) {

	v3Path := "/test-management/v3/functional/test-suites/" + testSuiteId + "/variables"

	v3Url, _ := url.Parse(v3Path)

	resp, byt := api.client.request(internalconstant.Get, v3Url.String(), nil, nil)
	if resp.StatusCode == 200 {
		//var variables []model.Variable
		var listResponse model.ListResponse
		err := json.Unmarshal(*byt, &listResponse)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return listResponse.Variables, nil
	}
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorVariablesGetCall)
}

func (api ApiClient) GetVariable(testSuiteId, variableId string) (*model.Variable, *model.CliError) {

	v3Path := "/test-management/v3/functional/test-suites/" + testSuiteId + "/variables/" + variableId

	v3Url, _ := url.Parse(v3Path)

	resp, byt := api.client.request(internalconstant.Get, v3Url.String(), nil, nil)
	if resp.StatusCode == 200 {
		//var variables []model.Variable
		var variable model.Variable
		err := json.Unmarshal(*byt, &variable)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return &variable, nil
	}
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorVariableGetCall)
}

func (api ApiClient) UpdateVariables(variable []model.Variable, testSuiteId string) (*model.VariableBulkResponse, *model.CliError) {
	variables, err := json.Marshal(variable)
	if err != nil {
		log.Error(err)
		util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.RequestParsingError), internalconstant.ExitStatusCode3)
	}

	var requestHeaders = make(http.Header)
	requestHeaders.Add(internalconstant.ContentType, internalconstant.ApplicationJson)

	v3Path := "/test-management/v3/functional/test-suites/" + testSuiteId + "/variables"

	// get response
	resp, byt := api.client.request(internalconstant.Put, v3Path, &variables, requestHeaders)

	if resp.StatusCode == 207 {
		var br model.VariableBulkResponse
		err := json.Unmarshal(*byt, &br)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return &br, nil
	}

	// if not 207 response
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorVariablePutCall)
}

func (api ApiClient) RemoveVariableFromTestSuite(testSuiteId, variableId string) (*model.BulkResponse, *model.CliError) {

	varId, _ := strconv.Atoi(variableId)
	variableIdArray := []int{varId}

	variableIdBytes, err := json.Marshal(variableIdArray)
	if err != nil {
		log.Error(err)
		util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.RequestParsingError), internalconstant.ExitStatusCode3)
	}

	path := "/test-management/v3/functional/test-suites/" + testSuiteId + "/variables/remove"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(internalconstant.ContentType, internalconstant.ApplicationJson)

	resp, byt := api.client.request(internalconstant.Post, path, &variableIdBytes, requestHeaders)
	if resp.StatusCode == 207 {
		var removeVariablesResponse model.BulkResponse
		err := json.Unmarshal(*byt, &removeVariablesResponse)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return &removeVariablesResponse, nil
	}

	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorRemoveVariablePostCall)
}

func (api ApiClient) GetLogLines(tcxId int) (interface{}, *model.CliError) {
	resp, byt := api.client.request(http.MethodGet, fmt.Sprintf("/test-management/v3/functional/test-case-executions/%d/log-lines", tcxId), nil, nil)
	if resp.StatusCode == 200 {
		var loglines interface{}
		err := json.Unmarshal(*byt, &loglines)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return loglines, nil
	}
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorTestLogLinesGetCall)
}

func (api ApiClient) GetRawReqResForRunId(testRunId string) ([]model.RawRequestResponse, *model.CliError) {
	path := "/test-management/v3/test-runs/" + testRunId + "/raw-request-response"

	resp, byt := api.client.request(http.MethodGet, path, nil, nil)
	if resp.StatusCode == 200 {
		var listResponse model.ListResponse
		err := json.Unmarshal(*byt, &listResponse)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return listResponse.RawRequestResponse, nil
	}
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorRawRequestResponseGetCall)
}

func (api ApiClient) GetRawReqResForTcxId(testCaseExecutionId string) (*model.RawRequestResponse, *model.CliError) {
	path := "/test-management/v3/functional/test-case-executions/" + testCaseExecutionId + "/raw-request-response"

	resp, byt := api.client.request(http.MethodGet, path, nil, nil)
	if resp.StatusCode == 200 {
		var rawRequestResponse model.RawRequestResponse
		err := json.Unmarshal(*byt, &rawRequestResponse)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return &rawRequestResponse, nil
	}
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorRawRequestResponseGetCall)
}

func (api ApiClient) EvaluateFunction(function *model.TryFunction) (*model.TryFunction, *model.CliError) {
	tryFunction, err := json.Marshal(function)
	if err != nil {
		log.Error(err)
		util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.RequestParsingError), internalconstant.ExitStatusCode3)
	}

	path := "/test-management/v3/functional/functions/try-it"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(internalconstant.ContentType, internalconstant.ApplicationJson)

	resp, byt := api.client.request(internalconstant.Post, path, &tryFunction, requestHeaders)
	if resp.StatusCode == 200 {
		var tryFunctionResponse model.TryFunction
		err := json.Unmarshal(*byt, &tryFunctionResponse)
		if err != nil {
			log.Error(err)
			util.AbortWithExitCode(util.GetGlobalErrorMessage(internalconstant.ResponseParsingError), internalconstant.ExitStatusCode3)
		}
		return &tryFunctionResponse, nil
	}
	return nil, model.CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, externalconstant.ApiErrorTryFunctionPostCall)
}
