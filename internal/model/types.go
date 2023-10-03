package model

// BrowserInfo object represents a browser with its name and version
type BrowserInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// ClientProfile object represents a client to be used for making requests
type ClientProfile struct {
	ClientProfileId int         `json:"clientProfileId,omitempty"`
	GeoLocation     string      `json:"geoLocation,omitempty"`
	IpVersion       string      `json:"ipVersion"`
	Browser         BrowserInfo `json:"browser,omitempty"`
	Client          string      `json:"client,omitempty"`
	ClientVersion   string      `json:"clientVersion,omitempty"`
}

type RequestHeader struct {
	HeaderName          string `json:"headerName"`
	HeaderNameResolved  string `json:"headerNameResolved,omitempty"`
	HeaderValue         string `json:"headerValue"`
	HeaderValueResolved string `json:"headerValueResolved,omitempty"`
	HeaderAction        string `json:"headerAction,omitempty"`
}

type TestRequest struct {
	TestRequestId          int             `json:"testRequestId,omitempty"`
	TestRequestUrl         string          `json:"testRequestUrl"`
	TestRequestUrlResolved string          `json:"testRequestUrlResolved,omitempty"`
	RequestMethod          string          `json:"requestMethod,omitempty"`
	RequestHeaders         []RequestHeader `json:"requestHeaders,omitempty"`
	Tags                   []string        `json:"tags,omitempty"`
	RequestBody            string          `json:"requestBody,omitempty"`
	RequestBodyResolved    string          `json:"requestBodyResolved,omitempty"`
	EncodeRequestBody      *bool           `json:"encodeRequestBody,omitempty"`
}

type Condition struct {
	ConditionId                 int    `json:"conditionId,omitempty"`
	ConditionExpression         string `json:"conditionExpression"`
	ConditionExpressionResolved string `json:"conditionExpressionResolved,omitempty"`
}

type AssociatedTestCases struct {
	AreAllTestCasesIncluded bool       `json:"areAllTestCasesIncluded"`
	TestCases               []TestCase `json:"testCases"`
}

type TestCase struct {
	TestCaseId       int               `json:"testCaseId,omitempty"`
	Order            int               `json:"order,omitempty"`
	ParentOrder      int               `json:"parentOrder,omitempty"`
	TestRequest      TestRequest       `json:"testRequest"`
	ClientProfile    ClientProfile     `json:"clientProfile,omitempty"`
	ClientProfileId  int               `json:"clientProfileId,omitempty"`
	Condition        Condition         `json:"condition"`
	SetVariables     []DynamicVariable `json:"setVariables,omitempty"`
	DerivedTestCases []TestCase        `json:"derivedTestCases,omitempty"`
	Warnings         []ApiSubError     `json:"warnings,omitempty"`
	CreatedBy        string            `json:"createdBy,omitempty"`
	CreatedDate      string            `json:"createdDate,omitempty"`
	ModifiedBy       string            `json:"modifiedBy,omitempty"`
	ModifiedDate     string            `json:"modifiedDate,omitempty"`
}
type TestCaseBulkResponse struct {
	Successes []TestCase    `json:"successes"`
	Failures  []ApiSubError `json:"failures"`
}

type BulkResponse struct {
	Successes []int         `json:"successes"`
	Failures  []ApiSubError `json:"failures"`
}

type VariableBulkResponse struct {
	Successes []Variable    `json:"successes"`
	Failures  []ApiSubError `json:"failures"`
}

// ListResponse We may add list responses for different objects here if needed in future e.g. variables, test-cases
type ListResponse struct {
	TestSuites         []TestSuite          `json:"testSuites,omitempty"`
	TestRuns           []TestRun            `json:"testRuns,omitempty"`
	Conditions         []Condition          `json:"conditions,omitempty"`
	TestRequests       []TestRequest        `json:"testRequests,omitempty"`
	Variables          []Variable           `json:"variables,omitempty"`
	RawRequestResponse []RawRequestResponse `json:"functionalRequestResponse,omitempty"`
}

type TestSuiteImportResponse struct {
	Success TestSuite              `json:"success,omitempty"`
	Failure TestSuiteImportFailure `json:"failure,omitempty"`
}

type TestSuiteImportFailure struct {
	Variables []ApiSubError `json:"variables,omitempty"`
	TestCases []ApiSubError `json:"testCases,omitempty"`
}

type TestSuite struct {
	CreatedBy               string        `json:"createdBy,omitempty"`
	CreatedDate             string        `json:"createdDate,omitempty"`
	ModifiedBy              string        `json:"modifiedBy,omitempty"`
	ModifiedDate            string        `json:"modifiedDate,omitempty"`
	DeletedBy               string        `json:"deletedBy,omitempty"`
	DeletedDate             string        `json:"deletedDate,omitempty"`
	TestSuiteId             int           `json:"testSuiteId,omitempty"`
	TestSuiteName           string        `json:"testSuiteName"`
	TestSuiteDescription    string        `json:"testSuiteDescription,omitempty"`
	IsLocked                bool          `json:"isLocked"`
	IsStateful              bool          `json:"isStateful"`
	ExecutableTestCaseCount int           `json:"executableTestCaseCount"`
	Configs                 AkamaiConfigs `json:"configs,omitempty"`
	TestCases               []TestCase    `json:"testCases,omitempty"`
	Variables               []Variable    `json:"variables,omitempty"`
}

type Variable struct {
	VariableId         int                  `json:"variableId,omitempty"`
	VariableName       string               `json:"variableName"`
	VariableValue      string               `json:"variableValue,omitempty"`
	VariableGroupValue []VariableGroupValue `json:"variableGroupValue,omitempty"`
	IsDynamicallyUsed  *bool                `json:"isDynamicallyUsed,omitempty"`
	CreatedBy          string               `json:"createdBy,omitempty"`
	CreatedDate        string               `json:"createdDate,omitempty"`
	ModifiedBy         string               `json:"modifiedBy,omitempty"`
	ModifiedDate       string               `json:"modifiedDate,omitempty"`
}

type VariableGroupValue struct {
	ColumnHeader string   `json:"columnHeader,omitempty"`
	ColumnValues []string `json:"columnValues,omitempty"`
}

type AkamaiConfigs struct {
	PropertyManager PropertyManager `json:"propertyManager,omitempty"`
}

type PropertyManager struct {
	PropertyId      int    `json:"propertyId,omitempty"`
	PropertyName    string `json:"propertyName,omitempty"`
	PropertyVersion int    `json:"propertyVersion"`
}

type PurgeInfo struct {
	Status string        `json:"status,omitempty"`
	Errors []ApiSubError `json:"errors,omitempty"`
}

type TestRun struct {
	TestRunId             int               `json:"testRunId,omitempty"`
	Status                string            `json:"status,omitempty"`
	TargetEnvironment     string            `json:"targetEnvironment"`
	SendEmailOnCompletion bool              `json:"sendEmailOnCompletion,omitempty"`
	Note                  string            `json:"note,omitempty"`
	Functional            FunctionalTestRun `json:"functional"`
	SubmittedBy           string            `json:"submittedBy,omitempty"`
	SubmittedDate         string            `json:"submittedDate,omitempty"`
	CompletedDate         string            `json:"completedDate,omitempty"`
	PurgeInfo             PurgeInfo         `json:"purgeInfo,omitempty"`
}

type FunctionalTestRun struct {
	Status                         string                   `json:"status,omitempty"`
	TestSuiteExecutions            []TestSuiteExecutions    `json:"testSuiteExecutions,omitempty"`
	PropertyManagerExecution       PropertyManagerExecution `json:"propertyManagerExecution,omitempty"`
	TestCaseExecution              TestCaseExecution        `json:"testCaseExecution,omitempty"`
	AllExecutionObjectsIncluded    *bool                    `json:"allExecutionObjectsIncluded,omitempty"`
	IsReevaluationInProgress       *bool                    `json:"isReevaluationInProgress,omitempty"`
	NextReevaluationCompletionTime string                   `json:"nextReevaluationCompletionTime,omitempty"`
	MaxReevaluationCompletionTime  string                   `json:"maxReevaluationCompletionTime,omitempty"`
}

type PropertyManagerExecution struct {
	PropertyId          int                   `json:"propertyId,omitempty"`
	PropertyName        string                `json:"propertyName,omitempty"`
	PropertyVersion     int                   `json:"propertyVersion"`
	TestSuiteExecutions []TestSuiteExecutions `json:"testSuiteExecutions,omitempty"`
}

type TestSuiteExecutions struct {
	TestSuiteId          int                  `json:"testSuiteId"`
	Status               string               `json:"status,omitempty"`
	TestCaseExecutions   []TestCaseExecutions `json:"testCaseExecutions"`
	TestSuiteContext     TestSuite            `json:"testSuiteContext,omitempty"`
	TestSuiteExecutionId int                  `json:"testSuiteExecutionId,omitempty"`
	SubmittedBy          string               `json:"submittedBy,omitempty"`
	SubmittedDate        string               `json:"submittedDate,omitempty"`
	CompletedDate        string               `json:"completedDate,omitempty"`
}

type TestCaseExecutions struct {
	TestCaseId                int                       `json:"testCaseId"`
	TestCaseExecutionId       int                       `json:"testCaseExecutionId,omitempty"`
	Status                    string                    `json:"status,omitempty"`
	ConditionEvaluationResult ConditionEvaluationResult `json:"conditionEvaluationResult,omitempty"`
	ResolvedSetVariables      interface{}               `json:"resolvedSetVariables,omitempty"`
	TestCaseContext           TestCase                  `json:"testCaseContext,omitempty"`
	DerivedTestCaseExecutions []TestCaseExecutions      `json:"derivedTestCaseExecutions,omitempty"`
	Errors                    []ApiSubError             `json:"errors,omitempty"`
	SubmittedBy               string                    `json:"submittedBy,omitempty"`
	SubmittedDate             string                    `json:"submittedDate,omitempty"`
	CompletedDate             string                    `json:"completedDate,omitempty"`
	IsReevaluationInProgress  *bool                     `json:"isReevaluationInProgress,omitempty"`
}

type TestCaseExecution struct {
	TestRequest               TestRequest               `json:"testRequest"`
	ClientProfile             ClientProfile             `json:"clientProfile,omitempty"`
	Condition                 Condition                 `json:"condition"`
	TestCaseExecutionId       int                       `json:"testCaseExecutionId,omitempty"`
	Status                    string                    `json:"status,omitempty"`
	ConditionEvaluationResult ConditionEvaluationResult `json:"conditionEvaluationResult,omitempty"`
	Errors                    []ApiSubError             `json:"errors,omitempty"`
	SubmittedBy               string                    `json:"submittedBy,omitempty"`
	SubmittedDate             string                    `json:"submittedDate,omitempty"`
	CompletedDate             string                    `json:"completedDate,omitempty"`
	IsReevaluationInProgress  *bool                     `json:"isReevaluationInProgress,omitempty"`
}

type ConditionEvaluationResult struct {
	ActualConditionData []ActualConditionData `json:"actualConditionData,omitempty"`
	Result              string                `json:"result,omitempty"`
}

type ActualConditionData struct {
	Name  string `json:"name"`
	Value string `json:"value,omitempty"`
}

type TestRunContext struct {
	TestRunId  int               `json:"testRunId"`
	Functional FunctionalContext `json:"functional"`
}

type FunctionalContext struct {
	TestSuites       []TestSuiteContext       `json:"testSuites,omitempty"`
	TestCases        []TestCase               `json:"testCases,omitempty"`
	PropertyManagers []PropertyManagerContext `json:"propertyManagers,omitempty"`
}

type TestSuiteContext struct {
	*TestSuite
	TestCases []TestCase `json:"testCases,omitempty"`
}

type PropertyManagerContext struct {
	*PropertyManager
	TestSuites []TestSuiteContext `json:"testSuites,omitempty"`
}

type ResultStats struct {
	PassedTestCasesCount       int
	FailedTestCasesCount       int
	InProgressTestCasesCount   int
	InconclusiveTestCasesCount int
}

//Condition template types

type ConditionExpression struct {
	ConditionExpressionId int      `json:"conditionExpressionId,omitempty"`
	ConditionExpression   string   `json:"conditionExpression,omitempty"`
	Examples              []string `json:"examples,omitempty"`
}

type PlaceHolder struct {
	PlaceHolder            string        `json:"placeHolder,omitempty"`
	ValueInputType         string        `json:"valueInputType,omitempty"`
	ValueDataType          string        `json:"valueDataType,omitempty"`
	IsCustomValueSupported bool          `json:"isCustomValueSupported,omitempty"`
	ValueSeparator         string        `json:"valueSeparator,omitempty"`
	AvailableValues        []interface{} `json:"availableValues"`
}

type ConditionType struct {
	ConditionType        string                `json:"conditionType,omitempty"`
	Label                string                `json:"label,omitempty"`
	ConditionExpressions []ConditionExpression `json:"conditionExpressions,omitempty"`
	PlaceHolders         []PlaceHolder         `json:"placeHolders"`
}

type ConditionTemplate struct {
	ConditionTypes []ConditionType `json:"conditionTypes,omitempty"`
}

type DefaultTestSuiteRequest struct {
	Configs        AkamaiConfigs `json:"configs"`
	TestRequestUrl []string      `json:"testRequestUrls"`
}

type DynamicVariable struct {
	VariableId    int    `json:"variableId,omitempty"`
	VariableName  string `json:"variableName"`
	VariableValue string `json:"variableValue"`
}

type RawRequestResponse struct {
	Request              Request  `json:"request,omitempty"`
	Response             Response `json:"response,omitempty"`
	TestCaseExecutionIds []int    `json:"testCaseExecutionIds,omitempty"`
}

type Request struct {
	Method      string        `json:"method,omitempty"`
	Url         string        `json:"url,omitempty"`
	HttpVersion string        `json:"httpVersion,omitempty"`
	QueryString []interface{} `json:"queryString,omitempty"`
	HeadersSize int           `json:"headersSize,omitempty"`
	BodySize    int           `json:"bodySize,omitempty"`
	Comment     string        `json:"comment,omitempty"`
	Cookies     []interface{} `json:"cookies,omitempty"`
	Headers     []Header      `json:"headers,omitempty"`
}

type TryFunction struct {
	FunctionExpression      string                  `json:"functionExpression"`
	Variables               []Variable              `json:"variables,omitempty"`
	TryFunctionResponseData TryFunctionResponseData `json:"responseData,omitempty"`
	Result                  string                  `json:"result"`
}

type TryFunctionResponseData struct {
	Response Response `json:"response,omitempty"`
}

type Response struct {
	Status      int      `json:"status"`
	StatusText  string   `json:"statusText,omitempty"`
	HttpVersion string   `json:"httpVersion"`
	Headers     []Header `json:"headers"`
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
