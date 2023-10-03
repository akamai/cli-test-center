package constant

// command usage and example
const (
	RootCommandUse = "test-center"

	TestUse          = "test"
	TestCommandAlias = "t"

	TestRunUse = `run  [--test-suite-id ID] | 
			[--test-suite-name 'NAME'] |
			[--property-name 'PROPERTY NAME' --property-version 'PROPERTY VERSION'] |
			[--property-id ID --property-version 'PROPERTY VERSION'] |
			[-u URL -c CONDITION --ip-version V4|V6 [--add-header 'name: value' ...] [--modify-header 'name: value' ...] [--filter-header name ...]]
			--env STAGING|PRODUCTION`
	TestRunExample = `  $ akamai test-center test run --test-suite-id 2500
  $ akamai test-center test run --test-suite-name 'Regression test cases for example.com'
  $ akamai test-center test run --property-name 'example.com' --property-version '26'
  $ akamai test-center test run --property-id '438285' --property-version '26'
  $ akamai test-center test run --url 'https://example.com/' --condition 'Response code is one of "200"' --ip-version 'V6' --modify-header 'Accept: application/json'
  $ akamai test-center t run -u 'https://example.com/' -c 'Response code is one of "200"' --ip-version 'V6' -m 'Accept: application/json'
  $ akamai test-center test run < {FILE_PATH}/FILE_NAME.json
  $ echo '{"functional":{"testSuiteExecutions":[{"testSuiteId":123}]},"targetEnvironment":"STAGING","purgeOnStaging":false,"note":"Testing 1 test suites on staging.","sendEmailOnCompletion":false}' | akamai test-center test run
  $ echo '{"functional":{"propertyManagerExecution":{"propertyId":4567,"propertyVersion":1}},"targetEnvironment":"STAGING","purgeOnStaging":true,"sendEmailOnCompletion":true,"note":"Testing the example.com v1 property"}' | akamai test-center test run
  $ echo '{"functional":{"propertyManagerExecution":{"propertyName":"example.com","propertyVersion":1}},"targetEnvironment":"STAGING","purgeOnStaging":true,"sendEmailOnCompletion":true,"note":"Testing the example.com v1 property"}' | akamai test-center test run
  $ echo '{"functional":{"testCaseExecution":{"testRequest":{"testRequestUrl":"https://example.com.com","requestHeaders":[{"headerName":"Accept","headerValue":"application/json","headerAction":"ADD"}],"requestBody":"{\"name\": \"akamai\"}","encodeRequestBody":true,"requestMethod":"POST"},"condition":{"conditionExpression":"Response code is one of \"200\""},"clientProfile":{"ipVersion":"IPV4","client":"CURL","geoLocation":"US"}}},"targetEnvironment":"STAGING","purgeOnStaging":true,"note":"Testing a simple test case on staging.","sendEmailOnCompletion":true}' | akamai test-center test run`

	TestListUse     = "list"
	TestListAlias   = "ls"
	TestListExample = `  $ akamai test-center test list`

	TestGetUse     = "get"
	TestGetAlias   = "view"
	TestGetExample = `  $ akamai test-center test get --test-run-id 1
  $ akamai test-center test get -i 1`

	TestSuiteUse          = "test-suite"
	TestSuiteCommandAlias = "ts"

	TestSuiteAddUse     = "create --test-suite-name NAME [--description DESCRIPTION] [--unlocked] [--stateful]  [--property-name 'PROPERTY NAME' --property-version 'PROPERTY VERSION'] "
	TestSuiteAddExample = `  $ akamai test-center test-suite create --test-suite-name 'Example TS'
  $ akamai test-center test-suite create --test-suite-name 'Example TS' --description 'TS for example.com' --unlocked --stateful --property-name 'example.com' --property-version '4'
  $ akamai test-center ts add -n 'Example TS' -d 'TS for example.com' --unlocked --stateful -p 'example.com' -v '4'
  $ akamai test-center test-suite create < {filepath}/filename.json`

	TestSuiteAddCommandAlias = "add"

	TestSuiteGenerateDefaultUse     = "generate-default --property-name 'PROPERTY NAME' --property-version 'PROPERTY VERSION' --url URL ... "
	TestSuiteGenerateDefaultExample = `  $ akamai test-center test-suite generate-default --property-name 'example.com' --property-version '4' --url "https://www.example.com/" -u "https://www.example.com/index/"
  $ akamai test-center ts template -p 'example.com' -v '4' -u "https://www.example.com/" -u "https://www.example.com/index/"  
  $ echo '{"configs":{"propertyManager":{"propertyName":"atc_test_config","propertyVersion":1}},"testRequestUrls":["http://www.example.com/"]}' | akamai test-center test-suite generate-default
  $ akamai test-center test-suite generate-default < {filepath}/filename.json`
	TestSuiteGenerateDefaultCommandAlias = "gd"

	TestSuiteEditUse     = "update --test-suite-id ID [--test-suite-name NAME] [--description DESCRIPTION] [--unlocked | --locked] [--stateful | --stateless] [--property-name 'PROPERTY NAME' --property-version 'PROPERTY VERSION' | --remove-property]"
	TestSuiteEditExample = `  $ akamai test-center test-suite update --test-suite-id 1001 --test-suite-name 'Updated Example TS'
  $ akamai test-center test-suite update --test-suite-id 1001 --test-suite-name 'Updated Example TS' --description 'TS for example.com' --property-name 'example.com' --property-version '4' --unlocked
  $ akamai test-center test-suite update --test-suite-id 1001 --stateful --remove-property
  $ akamai test-center test-suite update -i 1001 --stateful
  $ akamai test-center test-suite update < {filepath}/filename.json`
	TestSuiteEditCommandAlias = "edit"

	TestSuiteCreateWithChildObjects        = "create-with-child-objects"
	TestSuiteCreateWithChildObjectsExample = `  $ akamai test-center test-suite create-with-child-objects < {FILE_PATH}/FILE_NAME.json
  $ echo '{"testSuiteName":"ts1","testSuiteDescription":"ts1 description.","isLocked":true,"isStateful":false,"configs":{"propertyManager":{"propertyId":4567,"propertyVersion":1}},"variables":[{"variableName":"host","variableValue":"www.akamai.com"}],"testCases":[]}' | akamai test-center test-suite import`
	TestSuiteCreateWithChildObjectsCommandAlias = "import"

	TestSuiteListUse     = "list [--property-name 'PROPERTY NAME'] [--property-version 'PROPERTY VERSION'] [-u 'USERNAME'] [--search 'SEARCH STRING']"
	TestSuiteListExample = `  $ akamai test-center test-suite list
  $ akamai test-center test-suite list --property-name 'example.com' --property-version '4'
  $ akamai test-center test-suite list -u 'johndoe' --search 'regression'`
	TestSuiteListCommandAlias = "ls"

	TestSuiteUpdateWithChildObjectsUse     = "update-with-child-objects"
	TestSuiteUpdateWithChildObjectsExample = `  $ akamai test-center test-suite manage < {FILE_PATH}/FILE_NAME.json
  $ echo '{"testSuiteId":1,"testSuiteName":"ts1","testSuiteDescription":"ts1 description.","isLocked":true,"isStateful":false,"configs":{"propertyManager":{"propertyId":4567,"propertyVersion":1}},"variables":[{"variableName":"host","variableValue":"www.akamai.com"}],"testCases":[]}' | akamai test-center test-suite manage`
	TestSuiteUpdateWithChildObjectsCommandAlias = "manage"

	TestSuiteRemoveUse     = "remove [--test-suite-id ID | --test-suite-name NAME]"
	TestSuiteRemoveExample = `  $ akamai test-center test-suite remove --test-suite-name "Test suite name"
  $ akamai test-center test-suite remove --test-suite-id 12345`

	TestSuiteRestoreUse     = "restore [--test-suite-id ID | --test-suite-name NAME]"
	TestSuiteRestoreExample = `  $ akamai test-center test-suite restore --test-suite-name "Test suite name"
  $ akamai test-center test-suite restore --test-suite-id 12345`

	TestSuiteGetUse     = "get [--test-suite-id ID | --test-suite-name NAME]"
	TestSuiteGetExample = `  $ akamai test-center test-suite get --test-suite-id 1001
  $ akamai test-center test-suite get --test-suite-name 'Example TS'`
	TestSuiteGetCommandAliases = "view"

	TestSuiteGetWithChildObjectsUse     = "get-with-child-objects [--test-suite-id ID | --test-suite-name NAME] [--group-by test-request | condition | client-profile]"
	TestSuiteGetWithChildObjectsExample = `  $ akamai test-center test-suite get-with-child-objects --test-suite-id 1001
  $ akamai test-center test-suite get-with-child-objects --test-suite-name 'Example TS' --group-by test-request`
	TestSuiteGetWithChildObjectsCommandAliases = "export"

	ConditionUse            = "condition"
	ConditionCommandAliases = "c"

	ConditionTemplateUse     = "template"
	ConditionTemplateExample = `  $ akamai test-center condition template`

	ConditionListUse            = "list"
	ConditionListExample        = `  $ akamai test-center condition list`
	ConditionListCommandAliases = "ls"

	TestRequestUse            = "test-request"
	TestRequestCommandAliases = "tr"

	TestRequestListUse            = "list"
	TestRequestListExample        = `  $ akamai test-center test-request list`
	TestRequestListCommandAliases = "ls"

	TestCaseUse          = "test-case"
	TestCaseCommandAlias = "tc"

	CreateTestCaseCommandAlias = "add"
	CreateTestCaseUse          = "create [--test-suite-id ID | --test-suite-name NAME]  -u URL -c CONDITION [--ip-version V4|V6] [-a header ...] [-m header ...] [-f header ...] [-C client ...] [-M request-method ...] [-E -encode-request-body] [-b request-body ...] [-S set-variables ...]"
	CreateTestCaseExample      = `  $ akamai test-center test-case create --test-suite-id 1001 --url 'https://example.com/' --condition 'Response code is one of "200,201"'
  $ akamai test-center test-case create --test-suite-id 1001 -u 'https://example.com/' -c 'Response code is one of "200"' -a 'Accept: text/html' -a 'X-Custom: 123' -m 'User-Agent: Mozilla' -f 'Accept-Language' -C curl -M POST -S 'varName: varValue'
  $ akamai test-center test-case create --test-suite-id 1001 -u 'https://example.com/' -c 'Response code is one of "{{variableName}}"' -a 'Accept: text/html' -a 'X-Custom: 123' -m 'User-Agent: Mozilla' -f 'Accept-Language' -C curl -M POST -S 'varName: varValue' -E`
	ListTestCasesExample = `  $ akamai test-center test-case list --test-suite-id 1001 --resolve-variables --group-by test-request
  $ akamai test-center test-case list --test-suite-id 1001 --resolve-variables --group-by condition
  $ akamai test-center test-case list --test-suite-id 1001 --resolve-variables --group-by client-profile
  $ akamai test-center test-case ls -i 1001 --resolve-variables -g test-request`
	ListTestCasesUse          = "list [--test-suite-id ID | --test-suite-name name] --resolve-variables [--group-by ...]"
	ListTestCasesCommandAlias = "ls"

	GetTestCaseExample      = `  $ akamai test-center test-case get --test-suite-id 1001 --test-case-id 101 --resolve-variables`
	GetTestCaseUse          = "get [--test-suite-id ID | --test-suite-name name] [--test-case-id TEST_CASE_ID] --resolve-variables"
	GetTestCaseCommandAlias = "view"

	RemoveTestCaseUse     = "remove --test-suite-id ID [--order-num ORDER_NUMBER | --test-case-id TEST_CASE_ID]"
	RemoveTestCaseExample = `  $ akamai test-center test-case remove --test-suite-id 1001 --order-num 6
  $ akamai test-center test-case remove --test-suite-id 1001 --test-case-id 101`

	UpdateTestCaseCommandAlias = "edit"
	UpdateTestCaseUse          = "update [--test-suite-id ID | --test-suite-name NAME] [--test-case-id TEST_CASE_ID] -u URL -c CONDITION [--ip-version V4|V6] [-a header ...] [-m header ...] [-f header ...] [-C client ...] [-M request-method ...] [-E -encode-request-body] [-b request-body ...] [-S set-variables ...]"
	UpdateTestCaseExample      = `  $ akamai test-center test-case update --test-suite-id 1001 --test-case-id 101 --url 'https://example.com/' --condition 'Response code is one of "200,201"'
  $ akamai test-center test-case update --test-suite-id 1001 --test-case-id 101 -u 'https://example.com/' -c 'Response code is one of "200"' -a 'Accept: text/html' -a 'X-Custom: 123' -m 'User-Agent: Mozilla' -f 'Accept-Language' -C curl -M POST -S 'varName: varValue'
  $ akamai test-center test-case update --test-suite-id 1001 --test-case-id 101 -u 'https://example.com/' -c 'Response code is one of "{{variableName}}"' -a 'Accept: text/html' -a 'X-Custom: 123' -m 'User-Agent: Mozilla' -f 'Accept-Language' -C curl -M POST -S 'varName: varValue' -E`

	VariableUse            = "variable"
	VariableCommandAliases = "var"

	VariableCreateUse     = "create --test-suite-id ID --name NAME [--value VALUE | --group-value H1: value1, value2 --group-value H2: value3, value4]"
	VariableCreateExample = `  $ akamai test-center variable create --test-suite-id 1001 --name url --value 'https://example.com/' 
  $ akamai test-center variable create --test-suite-id 1001 --name url --group-value hostName: https://example.com/,https://example.com/123 --group-value ResponseCodes: 200,300`
	VariableCreateCommandAliases = "add"

	VariablesListUse            = "list --test-suite-id ID"
	VariablesListExample        = "  $ akamai test-center variable list --test-suite-id 1"
	VariablesListCommandAliases = "ls"

	VariableGetUse            = "get --test-suite-id ID --variable-id VARIABLE_ID"
	VariableGetExample        = "  $ akamai test-center variable get --test-suite-id 1 --variable-id 1"
	VariableGetCommandAliases = "view"

	VariableUpdateUse     = "update --test-suite-id ID --variable-id VARIABLE_ID --name NAME [--value VALUE | --group-value H1: value1, value2 --group-value H2: value3, value4]"
	VariableUpdateExample = `  $ akamai test-center variable update --test-suite-id 1001 --variable-id 1 --name url --value 'https://example.com/'
  $ akamai test-center variable update --test-suite-id 1001 --name url --variable-id 1 --group-value hostName: https://example.com/,https://example.com/123 --group-value ResponseCodes: 200,300`
	VariableUpdateCommandAliases = "edit"

	VariableRemoveUse     = "remove --test-suite-id ID --variable-id VARIABLE_ID"
	VariableRemoveExample = "  $ akamai test-center variable remove --test-suite-id 1 --variable-id 1"

	FunctionUse     = "function"
	FunctionAliases = "fn"

	TryItFunctionUse     = "try-it"
	TryItFunctionAliases = "try"
	TryItFunctionExample = `  $ akamai test-center function try-it < {filepath}/filename.json
  $ echo '{"functionExpression": "fn_getResponseHeaderValue(headerName, regex)","responseData": {"response": {"status": 200,"statusText": "OK","httpVersion": "HTTP/1.1","headers": [{"name": "server", "value":"Apache/2.2.15 (CentOS)"}]}}}' | akamai test-center function try-it`

	TestRawReqResUse        = "raw-request-response [--test-run-id ID | --tcx-id ID]"
	TestRawReqResExampleUse = `  $ akamai test-center test raw-request-response --test-run-id 1
  $ akamai test-center test raw-request-response --tcx-id 2`
	TestRawReqResCommandAliases = "rr"

	TestLogLinesUse            = "log-lines --tcx-id 1"
	TestLogLinesExampleUse     = "  $ akamai test-center test log-lines --tcx-id 1"
	TestLogLinesCommandAliases = "ll"
)

// Flag Names
const (
	FlagEdgerc          = "edgerc"
	FlagSection         = "section"
	FlagAccountKey      = "account-key"
	FlagJson            = "json"
	FlagForceColor      = "force-color"
	FlagPropertyId      = "property-id"
	FlagPropertyName    = "property-name"
	FlagPropertyVersion = "property-version"
	FlagUrl             = "url"
	FlagEnv             = "env"
	FlagAddHeader       = "add-header"
	FlagModifyHeader    = "modify-header"
	FlagFilterHeader    = "filter-header"
	FlagTestSuiteId     = "test-suite-id"
	FlagTestRunId       = "test-run-id"
	FlagTestSuiteName   = "test-suite-name"
	FlagIpVersion       = "ip-version"
	FlagCondition       = "condition"
	FlagDescription     = "description"
	FlagStateFul        = "stateful"
	FlagUnlocked        = "unlocked"
	FlagStateless       = "stateless"
	FlagLocked          = "locked"
	FlagRemoveProperty  = "remove-property"
	FlagUser            = "user"
	FlagSearch          = "search"
	FlagOrderNumber     = "order-num"
	FlagGroupBy         = "group-by"
	FlagHelp            = "help"
	FlagVersion         = "version"
	FlagVariableName    = "name"
	FlagVariableValue   = "value"
	FlagVarGroupValue   = "group-value"
	FlagVariableId      = "variable-id"
	FlagTestCaseExecId  = "tcx-id"

	// FlagClient etc...( new test run flags, can be used in test case as well.)
	FlagClient            = "client"
	FlagLocation          = "location"
	FlagRequestMethod     = "request-method"
	FlagEncodeRequestBody = "encode-request-body"
	FlagRequestBody       = "request-body"
	FlagSetVariables      = "set-variables"
	FlagTestCaseId        = "test-case-id"
	FlagResolveVariables  = "resolve-variables"
)

const (
	FlagUrlShortHand             = "u"
	FlagUserShortHand            = "u"
	FlagSectionShortHand         = "s"
	FlagEdgercShortHand          = "e"
	FlagAddHeaderShortHand       = "a"
	FlagModifyHeaderShortHand    = "m"
	FlagFilterHeaderShortHand    = "f"
	FlagConditionShortHand       = "c"
	FlagHelpShortHand            = "h"
	FlagTestSuiteIdShortHand     = "i"
	FlagTestRunIdShortHand       = "i"
	FlagTestSuiteNameShortHand   = "n"
	FlagDescriptionShortHand     = "d"
	FlagOrderNumberShortHand     = "o"
	FlagGroupShortHand           = "g"
	FlagPropertyShortHand        = "p"
	FlagPropertyVersionShortHand = "v"
	FlagVariableNameShortHand    = "n"
	FlagVariableValueShortHand   = "v"
	FlagVariableGroupShortHand   = "g"
	FlagTestCaseExecShortHand    = "x"

	// FlagClientShortHand etc ... (new test run flag short hands, can be used in test case as well.)
	FlagClientShortHand            = "C"
	FlagRequestMethodShortHand     = "M"
	FlagEncodeRequestBodyShortHand = "E"
	FlagRequestBodyShortHand       = "b"
	FlagSetVariablesShortHand      = "S"
	FlagTestCaseIdShortHand        = "I"
)

const (
	LabelId                       = "ID"
	LabelName                     = "Name"
	LabelDescription              = "Description"
	LabelStateful                 = "Stateful"
	LabelLocked                   = "Locked"
	LabelAssociatedProperty       = "Associated property version"
	LabelCreated                  = "Created"
	LabelLastModified             = "Last modified"
	LabelDeleted                  = "Deleted"
	LabelExpected                 = "Expected"
	LabelTestCaseExecutionId      = "Test case execution id"
	LabelRequestHeaders           = "Customized headers"
	LabelRequestBody              = "Request body"
	LabelEncodeRequestBody        = "URL encode"
	LabelNotRunError              = "Not run - Error: "
	LabelPassed                   = "Passed"
	LabelInProgress               = "In Progress"
	LabelInconclusive             = "Inconclusive"
	LabelFailed                   = "Failed"
	LabelActual                   = "Actual"
	LabelIsReevaluationInProgress = "Reevaluation In Progress"
	LabelIsReevaluationCompleted  = "Reevaluation Completed"
	LabelNotFound                 = "None found"
	LabelTotalItem                = "Total items"
	LabelStatus                   = "Status"
	LabelTargetEnvironment        = "Target Environment"
	LabelPropertyVersion          = "Property version"
	LabelIPv4                     = "IPv4"
	LabelIPv6                     = "IPv6"
	LabelTSDeleteState            = "Test suites with the '*' prefix in their names are in the deleted state. You can restore them for 30 days since their removal."
	LabelOrder                    = "Order"
	LabelTRUrl                    = "URL"
	LabelRequestMethod            = "Method"
	LabelSetVariables             = "Set variables"
	LabelCondition                = "Condition"
	LabelTags                     = "Keywords"
	VariableId                    = "Id"
	VariableName                  = "Name"
	VariableValue                 = "Value"
	TestCaseExecutionIds          = "Raw request response for test case execution ids"
	Request                       = "Request"
	Response                      = "Response"
	LabelFunctionExpression       = "Function expression"
	LabelFunctionResult           = "Result for expression"
	LabelClientProfile            = "Client Profile"
)

const (
	SeparatePipe     = " | "
	Dot              = "."
	Colon            = ":"
	PlusSign         = " + "
	SeparateLine     = "======================================================================================================================================================"
	SeparateBy       = " by "
	Star             = "*"
	Quote            = "\""
	Comma            = ","
	Equals           = "="
	SeparateLineStar = "**********************************************************************************************************************************************"
	IndentationSpace = "        "
)

// API CLinet Error Messages
const (
	ApiErrorAutoGeneratePostCall                  = "Failed to generate the test suite. Try again later."
	ApiErrorConditionTemplateGetCall              = "Failed to get the condition list. Try again later."
	ApiErrorRemoveTestCasesPostCall               = "Failed to remove test cases from the test suite. Try again later."
	ApiErrorSubmitTestRunPostCall                 = "Failed to submit the test run. Try again later."
	ApiErrorTestRunGetCall                        = "Failed to get the test run. Try again later."
	ApiErrorTestRunsGetCall                       = "Failed to get test runs. Try again later."
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
	ApiErrorGetConditions                         = "Failed to get Conditions. Try again later."
	ApiErrorGetTestRequests                       = "Failed to get TestRequests. Try again later."
	ApiErrorTestCaseForTestCaseGetCall            = "Failed to get test case for test case ID. Try again later."
	ApiErrorUpdateTestCasesToTestSuitPutCall      = "Failed to update test cases of the test suite. Try again later."
	ApiErrorVariablePostCall                      = "Failed to create the variable. Try again later"
	ApiErrorVariablesGetCall                      = "Failed to get the variables. Try again later"
	ApiErrorVariableGetCall                       = "Failed to get the variable. Try again later"
	ApiErrorVariablePutCall                       = "Failed to update the variable. Try again later"
	ApiErrorRemoveVariablePostCall                = "Failed to remove variable from the test suite. Try again later."
	ApiErrorTryFunctionPostCall                   = "Failed to evaluate the function expression. Try again later."
	ApiErrorRawRequestResponseGetCall             = "Failed to get the raw-request-response. Try again later."
	ApiErrorTestLogLinesGetCall                   = "Failed to get the log lines. Try again later."
)

const (
	CliErrorMessageTestRunStatus       = "Failed to check the test run status. Try again later."
	CliErrorMessageTemplateOutputError = "Output cannot be shown because of an internal CLI error. Try again later."
)

const TryFunctionEmptyResult = "No results to display. Modify the function or response sample and try again."
