package internal

const (
	LabelId                        = "Id"
	LabelName                      = "Name"
	LabelDescription               = "Description"
	LabelStateful                  = "Stateful"
	LabelLocked                    = "Locked"
	LabelVariables                 = "Variables"
	LabelAssociatedPropertyVersion = "Associated property version"
	LabelCreated                   = "Created"
	LabelLastModified              = "Last modified"
	LabelDeleted                   = "Deleted"
	LabelExpected                  = "Expected"
	LabelNotRunError               = "Not run - Error: "
	LabelPassed                    = "Passed"
	LabelFailed                    = "Failed"
	LabelActual                    = "Actual"
	LabelNotFound                  = "None found"
	LabelTotalItem                 = "Total items"
	LabelStatus                    = "Status"
	LabelTargetEnvironment         = "Target Environment"
	LabelPropertyVersion           = "Property version"
	LabelIPv4                      = "IPv4"
	LabelIPv6                      = "IPv6"
	LabelTSDeleteState             = "Test suite names prefixed with '*' are in deleted state"
)

const (
	SeparatePipe = " | "
	SeparateLine = "=========================================="
	SeparateBy   = " by "
)

// API CLinet Error Messages
const (
	ApiErrorAutoGeneratePostCall                  = "Failed to generate the test suite. Try again later."
	ApiErrorConditionTemplateGetCall              = "Failed to get the condition list. Try again later."
	ApiErrorRemoveTestCasesPostCall               = "Failed to remove test cases from the test suite. Try again later."
	ApiErrorConfigVersionGetCall                  = "Failed to get property versions. Try again later."
	ApiErrorSubmitTestRunPostCall                 = "Failed to submit the test run. Try again later."
	ApiErrorTestRunGetCall                        = "Failed to get the test run. Try again later."
	ApiErrorTestRunContextGetCall                 = "Failed to get test run context. Try again later."
	ApiErrorTestSuiteV3GetCall                    = "Failed to get test suites. Try again later."
	ApiErrorTestSuiteV3PostCall                   = "Failed to create the test suite. Try again later."
	ApiErrorEditTestSuiteV3PostCall               = "Failed to edit the test suite. Try again later."
	ApiErrorImportTestSuiteV3PostCall             = "Failed to import the test suite. Try again later."
	ApiErrorManageTestSuiteV3PutCall              = "Failed to manage the test suite. Try again later."
	ApiErrorRemoveTestSuiteDeleteCall             = "Failed to remove test suite. Try again later."
	ApiErrorRestoreTestSuitePostCall              = "Failed to restore the test suite. Try again later."
	ApiErrorTestSuiteWithChildObjectGetCall       = "Failed to get test suites with child objects. Try again later."
	ApiErrorTestCasesAssociatedForTestSuitGetCall = "Failed to get associated test cases for test suite. Try again later."
	ApiErrorAddTestCasesToTestSuitPostCall        = "Failed to add test cases to the test suite. Try again later."
)

const (
	CliErrorMessageTestRunStatus       = "Failed to check the test run status. Try again later."
	CliErrorMessageTestRunContext      = "Failed to transform the test run context. Try again later."
	CliErrorMessageTemplateOutputError = "Output cannot be shown because of an internal CLI error. Try again later."
)
