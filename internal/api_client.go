package internal

import (
	"fmt"
	"net/http"
	"net/url"

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

func (api ApiClient) GetConfigVersions() ([]ConfigVersion, *CliError) {
	resp, byt := api.client.request(Get, "/test-management/v2/functional/config-versions", nil, nil)
	if resp.StatusCode == 200 {
		var configVersions []ConfigVersion
		err := json.Unmarshal(*byt, &configVersions)
		if err != nil {
			log.Error(err)
			AbortWithExitCode(GetGlobalErrorMessage(ResponseParsingError), ExitStatusCode3)
		}
		return configVersions, nil
	}
	return nil, CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, ApiErrorConfigVersionGetCall)
}

func (api ApiClient) SubmitTestRun(testRun TestRun) (*TestRun, *CliError) {
	testRunBytes, err := json.Marshal(testRun)
	if err != nil {
		log.Error(err)
		AbortWithExitCode(GetGlobalErrorMessage(RequestParsingError), ExitStatusCode3)
	}

	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Post, "/test-management/v3/test-runs", &testRunBytes, requestHeaders)
	if resp.StatusCode == 202 {
		var testRunResponse TestRun
		err := json.Unmarshal(*byt, &testRunResponse)
		if err != nil {
			log.Error(err)
			AbortWithExitCode(GetGlobalErrorMessage(ResponseParsingError), ExitStatusCode3)
		}
		return &testRunResponse, nil
	}
	return nil, CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, ApiErrorSubmitTestRunPostCall)
}

func (api ApiClient) GetTestRun(testRunId int) (*TestRun, *CliError) {
	resp, byt := api.client.request(http.MethodGet, fmt.Sprintf("/test-management/v3/test-runs/%d", testRunId), nil, nil)
	if resp.StatusCode == 200 {
		var testRun TestRun
		err := json.Unmarshal(*byt, &testRun)
		if err != nil {
			log.Error(err)
			AbortWithExitCode(GetGlobalErrorMessage(ResponseParsingError), ExitStatusCode3)
		}
		return &testRun, nil
	}
	return nil, CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, ApiErrorTestRunGetCall)
}

func (api ApiClient) GetTestRunContext(testRunId int) (*TestRunContext, *CliError) {
	resp, byt := api.client.request(http.MethodGet, fmt.Sprintf("/test-management/v3/test-runs/%d/context", testRunId), nil, nil)
	if resp.StatusCode == 200 {
		var testRunContext TestRunContext
		err := json.Unmarshal(*byt, &testRunContext)
		if err != nil {
			log.Error(err)
			AbortWithExitCode(GetGlobalErrorMessage(ResponseParsingError), ExitStatusCode3)
		}
		return &testRunContext, nil
	}
	return nil, CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, ApiErrorTestRunContextGetCall)
}

func (api ApiClient) GetTestSuitesV3(propertyName, propVersion, user string) ([]TestSuiteV3, *CliError) {

	v3Path := "/test-management/v3/functional/test-suites?includeRecentlyDeleted=true"
	tsV3Url, _ := url.Parse(v3Path)

	// add optional query parameters
	queryMap := map[string]string{
		"propertyName":    propertyName,
		"propertyVersion": propVersion,
		"user":            user,
	}
	addQueryParams(tsV3Url, queryMap)

	// get response
	resp, byt := api.client.request(Get, tsV3Url.String(), nil, nil)

	// parse result
	if resp.StatusCode == 200 {
		var testSuites ListResponse
		err := json.Unmarshal(*byt, &testSuites)
		if err != nil {
			log.Error(err)
			AbortWithExitCode(GetGlobalErrorMessage(ResponseParsingError), ExitStatusCode3)
		}

		log.Infof("GetTestSuitesV3 [%s] returned %d items", tsV3Url.String(), len(testSuites.TestSuites))
		return testSuites.TestSuites, nil
	}

	// if not 200 response
	return nil, CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, ApiErrorTestSuiteV3GetCall)
}

func (api ApiClient) GetTestSuiteV3(id string) (*TestSuiteV3, *CliError) {

	v3Path := "/test-management/v3/functional/test-suites/" + id

	// get response
	resp, byt := api.client.request(Get, v3Path, nil, nil)

	// parse result
	if resp.StatusCode == 200 {
		var testSuite TestSuiteV3
		err := json.Unmarshal(*byt, &testSuite)
		if err != nil {
			log.Error(err)
			AbortWithExitCode(GetGlobalErrorMessage(ResponseParsingError), ExitStatusCode3)
		}

		return &testSuite, nil
	}

	// if not 200 response
	return nil, CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, ApiErrorTestSuiteV3GetCall)
}

func (api ApiClient) RemoveTestSuite(id string) *CliError {

	v3Path := "/test-management/v3/functional/test-suites/" + id

	// get response
	resp, byt := api.client.request(Delete, v3Path, nil, nil)

	// parse result
	if resp.StatusCode == 204 {
		return nil
	}

	// if not 204 response
	return CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, ApiErrorRemoveTestSuiteDeleteCall)
}

func (api ApiClient) RestoreTestSuite(id string) (*TestSuiteV3, *CliError) {

	v3Path := "/test-management/v3/functional/test-suites/" + id + "/restore"

	// get response
	resp, byt := api.client.request(Post, v3Path, nil, nil)

	// parse result
	if resp.StatusCode == 200 {
		var testSuite TestSuiteV3
		err := json.Unmarshal(*byt, &testSuite)
		if err != nil {
			log.Error(err)
			AbortWithExitCode(GetGlobalErrorMessage(ResponseParsingError), ExitStatusCode3)
		}

		return &testSuite, nil
	}

	// if not 200 response
	return nil, CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, ApiErrorRestoreTestSuitePostCall)
}

func (api ApiClient) AddTestSuitesV3(testSuiteV3 TestSuiteV3) (*TestSuiteV3, *CliError) {
	testSuiteBytes, err := json.Marshal(testSuiteV3)
	if err != nil {
		log.Error(err)
		AbortWithExitCode(GetGlobalErrorMessage(RequestParsingError), ExitStatusCode3)
	}

	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	v3Path := "/test-management/v3/functional/test-suites"

	// get response
	resp, byt := api.client.request(Post, v3Path, &testSuiteBytes, requestHeaders)

	// parse result
	if resp.StatusCode == 201 {
		var testSuite TestSuiteV3
		err := json.Unmarshal(*byt, &testSuite)
		if err != nil {
			log.Error(err)
			AbortWithExitCode(GetGlobalErrorMessage(ResponseParsingError), ExitStatusCode3)
		}

		return &testSuite, nil
	}

	// if not 201 response
	return nil, CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, ApiErrorTestSuiteV3PostCall)
}

func (api ApiClient) EditTestSuitesV3(testSuiteV3 TestSuiteV3, id string) (*TestSuiteV3, *CliError) {
	testSuiteBytes, err := json.Marshal(testSuiteV3)
	if err != nil {
		log.Error(err)
		AbortWithExitCode(GetGlobalErrorMessage(RequestParsingError), ExitStatusCode3)
	}

	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	v3Path := "/test-management/v3/functional/test-suites/" + id

	// get response
	resp, byt := api.client.request(Put, v3Path, &testSuiteBytes, requestHeaders)

	// parse result
	if resp.StatusCode == 200 {
		var testSuite TestSuiteV3
		err := json.Unmarshal(*byt, &testSuite)
		if err != nil {
			log.Error(err)
			AbortWithExitCode(GetGlobalErrorMessage(ResponseParsingError), ExitStatusCode3)
		}

		return &testSuite, nil
	}

	// if not 200 response
	return nil, CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, ApiErrorEditTestSuiteV3PostCall)
}

func (api ApiClient) GetTestSuitesWithChildObjects(testSuiteId int) (*TestSuiteV3, *CliError) {

	v3Path := fmt.Sprintf("/test-management/v3/functional/test-suites/%d/with-child-objects", testSuiteId)
	resp, byt := api.client.request(Get, v3Path, nil, nil)

	if resp.StatusCode == 200 {
		var testSuite TestSuiteV3
		err := json.Unmarshal(*byt, &testSuite)
		if err != nil {
			log.Error(err)
			AbortWithExitCode(GetGlobalErrorMessage(ResponseParsingError), ExitStatusCode3)
		}
		return &testSuite, nil
	}
	return nil, CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, ApiErrorTestSuiteWithChildObjectGetCall)
}

func (api ApiClient) GetV3AssociatedTestCasesForTestSuite(testSuiteId int) (*AssociatedTestCases, *CliError) {

	v3Path := fmt.Sprintf("/test-management/v3/functional/test-suites/%d/test-cases", testSuiteId)
	resp, byt := api.client.request(Get, v3Path, nil, nil)

	if resp.StatusCode == 200 {
		var testCases AssociatedTestCases
		err := json.Unmarshal(*byt, &testCases)
		if err != nil {
			log.Error(err)
			AbortWithExitCode(GetGlobalErrorMessage(ResponseParsingError), ExitStatusCode3)
		}
		return &testCases, nil
	}
	return nil, CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, ApiErrorTestCasesAssociatedForTestSuitGetCall)
}

func (api ApiClient) AddTestCaseToTestSuite(testSuiteId int, testCases []TestCase) (*TestCaseBulkResponse, *CliError) {

	testCasesBytes, err := json.Marshal(testCases)
	if err != nil {
		log.Error(err)
		AbortWithExitCode(GetGlobalErrorMessage(RequestParsingError), ExitStatusCode3)
	}

	path := fmt.Sprintf("/test-management/v3/functional/test-suites/%d/test-cases", testSuiteId)
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Post, path, &testCasesBytes, requestHeaders)
	if resp.StatusCode == 207 {
		var createTestCaseResponse TestCaseBulkResponse
		err := json.Unmarshal(*byt, &createTestCaseResponse)
		if err != nil {
			log.Error(err)
			AbortWithExitCode(GetGlobalErrorMessage(ResponseParsingError), ExitStatusCode3)
		}
		return &createTestCaseResponse, nil
	}
	return nil, CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, ApiErrorAddTestCasesToTestSuitPostCall)
}

func (api ApiClient) ImportTestSuite(testSuiteImport TestSuiteV3) (*TestSuiteImportResponseV3, *CliError) {

	testSuiteImportJson, err := json.Marshal(testSuiteImport)
	if err != nil {
		log.Error(err)
		AbortWithExitCode(GetGlobalErrorMessage(RequestParsingError), ExitStatusCode3)
	}

	path := "/test-management/v3/functional/test-suites/with-child-objects"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Post, path, &testSuiteImportJson, requestHeaders)
	if resp.StatusCode == 207 {
		var testSuiteImportResponseV3 TestSuiteImportResponseV3
		err := json.Unmarshal(*byt, &testSuiteImportResponseV3)
		if err != nil {
			log.Error(err)
			AbortWithExitCode(GetGlobalErrorMessage(ResponseParsingError), ExitStatusCode3)
		}
		return &testSuiteImportResponseV3, nil
	}
	return nil, CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, ApiErrorImportTestSuiteV3PostCall)
}

func (api ApiClient) ManageTestSuite(testSuiteManage TestSuiteV3, testSuiteId int) (*TestSuiteV3, *CliError) {

	testSuiteManageJson, err := json.Marshal(testSuiteManage)
	if err != nil {
		log.Error(err)
		AbortWithExitCode(GetGlobalErrorMessage(RequestParsingError), ExitStatusCode3)
	}

	path := fmt.Sprintf("/test-management/v3/functional/test-suites/%d/with-child-objects", testSuiteId)
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Put, path, &testSuiteManageJson, requestHeaders)
	if resp.StatusCode == 200 {
		var testSuiteManageResponseV3 TestSuiteV3
		err := json.Unmarshal(*byt, &testSuiteManageResponseV3)
		if err != nil {
			log.Error(err)
			AbortWithExitCode(GetGlobalErrorMessage(ResponseParsingError), ExitStatusCode3)
		}
		return &testSuiteManageResponseV3, nil
	}
	return nil, CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, ApiErrorManageTestSuiteV3PutCall)
}

func (api ApiClient) RemoveTestCasesFromTestSuite(testSuiteId int, testCaseIds []int) (*BulkResponse, *CliError) {

	testCasesIdsBytes, err := json.Marshal(testCaseIds)
	if err != nil {
		log.Error(err)
		AbortWithExitCode(GetGlobalErrorMessage(RequestParsingError), ExitStatusCode3)
	}

	path := fmt.Sprintf("/test-management/v3/functional/test-suites/%d/test-cases/remove", testSuiteId)
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Post, path, &testCasesIdsBytes, requestHeaders)
	if resp.StatusCode == 207 {
		var removeTestCaseResponse BulkResponse
		err := json.Unmarshal(*byt, &removeTestCaseResponse)
		if err != nil {
			log.Error(err)
			AbortWithExitCode(GetGlobalErrorMessage(ResponseParsingError), ExitStatusCode3)
		}
		return &removeTestCaseResponse, nil
	}

	return nil, CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, ApiErrorRemoveTestCasesPostCall)
}

func (api ApiClient) GetConditionTemplate() (*ConditionTemplate, *CliError) {

	v3Path := "/test-management/v3/functional/test-catalog/template"
	condV3Url, _ := url.Parse(v3Path)

	// get response
	resp, byt := api.client.request(Get, condV3Url.String(), nil, nil)

	// parse result
	if resp.StatusCode == 200 {
		var condTemplate ConditionTemplate
		err := json.Unmarshal(*byt, &condTemplate)
		if err != nil {
			log.Error(err)
			fmt.Println(err)
			AbortWithExitCode(GetGlobalErrorMessage(ResponseParsingError), ExitStatusCode3)
		}

		return &condTemplate, nil
	}

	// if not 200 response
	return nil, CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, ApiErrorConditionTemplateGetCall)
}

func (api ApiClient) GenerateTestSuite(defaultTsRequest DefaultTestSuiteRequest) (*TestSuiteV3, *CliError) {

	defaultTestSuiteReqJson, err := json.Marshal(defaultTsRequest)
	if err != nil {
		log.Error(err)
		AbortWithExitCode(GetGlobalErrorMessage(RequestParsingError), ExitStatusCode3)
	}

	path := "/test-management/v3/functional/test-suites/auto-generate"
	var requestHeaders = make(http.Header)
	requestHeaders.Add(ContentType, ApplicationJson)

	resp, byt := api.client.request(Post, path, &defaultTestSuiteReqJson, requestHeaders)

	if resp.StatusCode == 200 {
		var defaultTsResponse TestSuiteV3
		err := json.Unmarshal(*byt, &defaultTsResponse)
		if err != nil {
			log.Error(err)
			AbortWithExitCode(GetGlobalErrorMessage(ResponseParsingError), ExitStatusCode3)
		}
		return &defaultTsResponse, nil
	}
	return nil, CliErrorFromPulsarProblemObject(*byt, nil, resp.StatusCode, ApiErrorAutoGeneratePostCall)
}
