package constant

// HTTP headers
const (
	ContentType      = "Content-Type"
	IsRequestFromCli = "X-ATC-IS-CLI-REQUEST"
)

// HTTP header values
const (
	ApplicationJson  = "application/json"
	RequestIsFromCli = "YES"
)

// HTTP request methods
const (
	Get    = "GET"
	Post   = "POST"
	Put    = "PUT"
	Delete = "DELETE"
)

// Request header actions
const (
	Add    = "ADD"
	Modify = "MODIFY"
	Filter = "FILTER"
)

// Test Center execution / evaluation status
const (
	CompletedEnum                      = "COMPLETED"
	CompletedWithUnexpectedResultsEnum = "COMPLETED_WITH_UNEXPECTED_RESULTS"
	InProgressEnum                     = "IN_PROGRESS"
	PassedEnum                         = "PASSED"
	Inconclusive                       = "INCONCLUSIVE"
	FailedEnum                         = "FAILED"
)

// IP versions
const (
	Ipv4 = "IPV4"
	Ipv6 = "IPV6"
)

// Supported client types
const (
	Chrome = "CHROME"
	Curl   = "CURL"
)

// DefaultLocation Supported locations
const DefaultLocation = "US"

// Supported request methods
const (
	GetRequestMethod  = "GET"
	HeadRequestMethod = "HEAD"
	PostRequestMethod = "POST"
)

// Environments
const (
	Staging    = "STAGING"
	Production = "PRODUCTION"
)

// Group by Test cases constants
const (
	GroupByTestRequest   = "test-request"
	GroupByCondition     = "condition"
	GroupByClientProfile = "client-profile"
)

// Run Test Using
const (
	RunTestUsingTestSuiteId     = "testSuiteId"
	RunTestUsingTestSuiteName   = "testSuiteName"
	RunTestUsingPropertyVersion = "propertyVersion"
	RunTestUsingSingleTestCase  = "singleTestCase"
	RunTestUsingJsonInput       = "jsonInput"
)

// Raw Request response using
const (
	RawRequestUsingTestRunId           = "testRun"
	RawRequestUsingTestCaseExecutionId = "testCaseExecution"
)

// Exit Status Codes
const (
	ExitStatusCode0 = 0   // Success 1xx 2xx 3xx
	ExitStatusCode1 = 1   // Generic CLI Exception
	ExitStatusCode2 = 2   // Command arguments and flag missing/mismatch exception
	ExitStatusCode3 = 3   // Parsing Error
	BaseSubtractor  = 300 // This number is subtracted from 4xx and 5xx status code API responses to get exit code. e.g. For response code 404, Exit code will be - 404-300 = 104
)

// Default Values
const (
	EdgercFileNameDefaultValue = ".edgerc"
	SectionDefaultValue        = "default"
	IpVersionDefaultValue      = "V4"
)

// Environment variable Constants
const (
	DefaultEdgercPathKey    = "AKAMAI_EDGERC"
	DefaultEdgercSectionKey = "AKAMAI_EDGERC_SECTION"
	DefaultJsonOutputKey    = "AKAMAI_OUTPUT_JSON"
)

const (
	Empty      = ""
	Space      = " "
	TabSpace   = "\t"
	EmptyArray = "[]"
)
