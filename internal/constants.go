package internal

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
	Completed  = "COMPLETED"
	InProgress = "IN_PROGRESS"
	Passed     = "PASSED"
)

// IP versions
const (
	Ipv4 = "IPV4"
	Ipv6 = "IPV6"
)

// Chrome Client
const (
	Chrome = "CHROME"
)

// Environments
const (
	Staging    = "STAGING"
	Production = "PRODUCTION"
)

// constant message keys
const (
	Short                             = "short"
	Long                              = "long"
	Json                              = "json"
	Example                           = "example"
	Missing                           = "missing"
	Invalid                           = "invalid"
	Global                            = "global"
	RequestParsingError               = "requestParsingError"
	ResponseParsingError              = "responseParsingError"
	Empty                             = ""
	PropertyVersionsResource          = "propertyVersions"
	TestSuiteResource                 = "testSuite"
	TestCaseResource                  = "testCase"
	TestRunResource                   = "testRun"
	Read                              = "read"
	Create                            = "create"
	Update                            = "update"
	PropertyVersionNotFound           = "propertyVersionNotFound"
	PropertyVersionTestSuitesNotFound = "propertyVersionTestSuitesNotFound"
	SubCommandNoArgumentPassed        = "noArgumentPassed"
	SubCommandWrongArgumentPassed     = "wrongArgumentPassed"
	SubCommandWithArgumentPassed      = "argumentPassed"
	PropertyVersionFlagKey            = "propertyVersion"
	MessageTypeSpinner                = "spinner"
	MessageTypeTestCmdSpinner         = "testCmdSpinner"
	MessageTypeDisplay                = "display"
)

// Group by Test cases constants
const (
	GroupByTestRequest = "test-request"
	GroupByCondition   = "condition"
	GroupByIpVersion   = "ipversion"
)

// Run Test Using
const (
	RunTestUsingTestSuiteId     = "testSuiteId"
	RunTestUsingTestSuiteName   = "testSuiteName"
	RunTestUsingPropertyVersion = "propertyVersion"
	RunTestUsingSingleTestCase  = "singleTestCase"
	RunTestUsingJsonInput       = "jsonInput"
)

// Exit Status Codes
const (
	ExitStatusCode0 = 0   // Success 1xx 2xx 3xx
	ExitStatusCode1 = 1   // Generic CLI Exception
	ExitStatusCode2 = 2   // Command arguments and flag missing/mismatch exception
	ExitStatusCode3 = 3   // Parsing Error
	BaseSubtractor  = 300 // This number is subtracted from 4xx and 5xx status code API responses to get exit code. e.g. For response code 404, Exit code will be - 404-300 = 104
)
