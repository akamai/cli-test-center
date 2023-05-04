package internal

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
}

type RequestHeader struct {
	HeaderName   string `json:"headerName"`
	HeaderValue  string `json:"headerValue"`
	HeaderAction string `json:"headerAction,omitempty"`
}

type TestRequest struct {
	TestRequestId  int             `json:"testRequestId,omitempty"`
	TestRequestUrl string          `json:"testRequestUrl"`
	RequestMethod  string          `json:"requestMethod,omitempty"`
	RequestHeaders []RequestHeader `json:"requestHeaders,omitempty"`
	Tags           []string        `json:"tags,omitempty"`
}

type Condition struct {
	ConditionId         int    `json:"conditionId,omitempty"`
	ConditionExpression string `json:"conditionExpression"`
}

type AssociatedTestCases struct {
	AreAllTestCasesIncluded bool       `json:"areAllTestCasesIncluded"`
	TestCases               []TestCase `json:"testCases"`
}

type TestCase struct {
	TestCaseId      int           `json:"testCaseId,omitempty"`
	Order           int           `json:"order,omitempty"`
	TestRequest     TestRequest   `json:"testRequest"`
	ClientProfile   ClientProfile `json:"clientProfile,omitempty"`
	ClientProfileId int           `json:"clientProfileId,omitempty"`
	Condition       Condition     `json:"condition"`
}

type TestCaseBulkResponse struct {
	Successes []TestCase    `json:"successes"`
	Failures  []ApiSubError `json:"failures"`
}

type BulkResponse struct {
	Successes []int         `json:"successes"`
	Failures  []ApiSubError `json:"failures"`
}

type TestSuite struct {
	CreatedBy            string `json:"createdBy,omitempty"`
	CreatedDate          string `json:"createdDate,omitempty"`
	ModifiedBy           string `json:"modifiedBy,omitempty"`
	ModifiedDate         string `json:"modifiedDate,omitempty"`
	TestSuiteId          int    `json:"testSuiteId,omitempty"`
	TestSuiteName        string `json:"testSuiteName"`
	TestSuiteDescription string `json:"testSuiteDescription,omitempty"`
	IsLocked             bool   `json:"isLocked"`
	IsStateful           bool   `json:"isStateful"`
}

// We may add list responses for different objects here if needed in future e.g. variables, test-cases
type ListResponse struct {
	TestSuites []TestSuiteV3 `json:"testSuites,omitempty"`
}

type TestSuiteImportResponseV3 struct {
	Success TestSuiteV3            `json:"success,omitempty"`
	Failure TestSuiteImportFailure `json:"failure,omitempty"`
}

type TestSuiteImportFailure struct {
	Variables []ApiSubError `json:"variables,omitempty"`
	TestCases []ApiSubError `json:"testCases,omitempty"`
}

type TestSuiteV3 struct {
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
	VariableId    int    `json:"variableId,omitempty"`
	VariableName  string `json:"variableName"`
	VariableValue string `json:"variableValue,omitempty"`
}

type AkamaiConfigs struct {
	PropertyManager PropertyManager `json:"propertyManager,omitempty"`
}

type PropertyManager struct {
	PropertyId      int    `json:"propertyId,omitempty"`
	PropertyName    string `json:"propertyName,omitempty"`
	PropertyVersion int    `json:"propertyVersion"`
}

type ConfigVersion struct {
	ModifiedBy      string `json:"modifiedBy,omitempty"`
	ModifiedDate    string `json:"modifiedDate,omitempty"`
	ConfigVersionId int    `json:"configVersionId,omitempty"`
	ArlFileId       int    `json:"arlFileId"`
	PropertyName    string `json:"propertyName"`
	PropertyVersion int    `json:"propertyVersion"`
	LastSync        string `json:"lastSync,omitempty"`
}

type TestRun struct {
	TestRunId             int               `json:"testRunId,omitempty"`
	Status                string            `json:"status,omitempty"`
	TargetEnvironment     string            `json:"targetEnvironment"`
	SendEmailOnCompletion bool              `json:"sendEmailOnCompletion"`
	Note                  string            `json:"note,omitempty"`
	Functional            FunctionalTestRun `json:"functional"`
	SubmittedBy           string            `json:"submittedBy,omitempty"`
	SubmittedDate         string            `json:"submittedDate,omitempty"`
	CompletedDate         string            `json:"completedDate,omitempty"`
}

type FunctionalTestRun struct {
	Status                   string                   `json:"status,omitempty"`
	TestSuiteExecutionsV3    []TestSuiteExecutionV3   `json:"testSuiteExecutions,omitempty"`
	PropertyManagerExecution PropertyManagerExecution `json:"propertyManagerExecution,omitempty"`
	TestCaseExecution        TestCaseExecution        `json:"testCaseExecution,omitempty"`
}

type PropertyManagerExecution struct {
	PropertyId            int                    `json:"propertyId,omitempty"`
	PropertyName          string                 `json:"propertyName"`
	PropertyVersion       int                    `json:"propertyVersion"`
	TestSuiteExecutionsV3 []TestSuiteExecutionV3 `json:"testSuiteExecutions,omitempty"`
}

type TestSuiteExecutionV3 struct {
	TestSuiteId          int                   `json:"testSuiteId"`
	Status               string                `json:"status,omitempty"`
	TestCaseExecutionsV3 []TestCaseExecutionV3 `json:"testCaseExecutions"`
	TestSuiteContext     TestSuiteV3
	TestSuiteExecutionId int    `json:"testSuiteExecutionId,omitempty"`
	SubmittedBy          string `json:"submittedBy,omitempty"`
	SubmittedDate        string `json:"submittedDate,omitempty"`
	CompletedDate        string `json:"completedDate,omitempty"`
}

type TestCaseExecutionV3 struct {
	TestCaseId                int                       `json:"testCaseId"`
	TestCaseExecutionId       int                       `json:"testCaseExecutionId,omitempty"`
	Status                    string                    `json:"status,omitempty"`
	ConditionEvaluationResult ConditionEvaluationResult `json:"conditionEvaluationResult,omitempty"`
	TestCaseContext           TestCase                  `json:"testCaseContext,omitempty"`
	Errors                    []ApiSubError             `json:"errors,omitempty"`
	SubmittedBy               string                    `json:"submittedBy,omitempty"`
	SubmittedDate             string                    `json:"submittedDate,omitempty"`
	CompletedDate             string                    `json:"completedDate,omitempty"`
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
}

type ConditionEvaluationResult struct {
	ActualConditionData []ActualConditionData `json:"actualConditionData,omitempty"`
	Result              string                `json:"result,omitempty"`
}

type ActualConditionData struct {
	Name  string `json:"name"`
	Value string `json:"value"`
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

type FunctionalContextMap struct {
	TestSuitesMap       map[int]TestSuiteContextMap
	TestCasesMap        map[int]TestCase
	PropertyManagersMap map[int]PropertyManagerContextMap
}

type TestSuiteContext struct {
	*TestSuite
	TestCases []TestCase `json:"testCases,omitempty"`
}

type TestSuiteContextMap struct {
	*TestSuite
	TestCasesMap map[int]TestCase
}
type PropertyManagerContext struct {
	*PropertyManager
	TestSuites []TestSuiteContext `json:"testSuites,omitempty"`
}

type PropertyManagerContextMap struct {
	*PropertyManager
	TestSuitesMap map[int]TestSuiteContextMap
}

type ResultStats struct {
	TotalTestCasesCount       int
	PassedTestCasesCount      int
	FailedTestCasesCount      int
	SystemErrorTestCasesCount int
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
