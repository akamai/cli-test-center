package cmd

// command usage and example
const (
	RootCommandUse = "test-center"

	TestUse     = "test [--test-suite-id ID | --test-suite-name 'NAME'] | [--property 'PROPERTY NAME' --propver 'PROPERTY VERSION'] |\n\t\t   [-u URL -c CONDITION -i V4|V6 [--add-header 'name: value' ...] [--modify-header 'name: value' ...] [--filter-header name ...]] \n\t\t   -e STAGING|PRODUCTION"
	TestExample = `  $ akamai test-center test --test-suite-id 2500
  $ akamai test-center test --test-suite-name 'Regression test cases for example.com'
  $ akamai test-center test --property 'example.com' --propver '26'
  $ akamai test-center test --url 'https://example.com/' --condition 'Response code is one of "200"' --ip-version 'V6' --modify-header 'Accept: application/json'
  $ akamai test-center test < {FILE_PATH}/FILE_NAME.json`
	TestCommandAlias = "t"

	TestSuiteUse          = "test-suite"
	TestSuiteCommandAlias = "ts"

	TestSuiteAddUse     = "add --name NAME [--description DESCRIPTION] [--unlocked] [--stateful]  [--property 'PROPERTY NAME' --propver 'PROPERTY VERSION'] "
	TestSuiteAddExample = `  $ akamai test-center test-suite add --name 'Example TS'
  $ akamai test-center test-suite add --name 'Example TS' --description 'TS for example.com' --unlocked --stateful --property 'example.com' --propver '4'`

	TestSuiteAddTestCaseUse     = "add-test-case [--test-suite-id ID | --test-suite-name NAME] -u URL -c CONDITION [-i V4|V6] [-a header ...] [-m header ...] [-f header ...]"
	TestSuiteAddTestCaseExample = `  $ akamai test-center test-suite add-test-case --test-suite-id 1001 --url 'https://example.com/' --condition 'Response code is one of "200,201"'
  $ akamai test-center test-suite add-test-case --test-suite-name 'Example TS' -u 'https://example.com/' -c 'Response code is one of "200"' -a 'Accept: text/html' -a 'X-Custom: 123' -m 'User-Agent: Mozilla' -f 'Accept-Language'`

	TestSuiteAutoGenerationUse     = "generate-default --property 'PROPERTY NAME' --propver 'PROPERTY VERSION' --url URL ... "
	TestSuiteAutoGenerationExample = `  $ akamai test-center test-suite generate-default --property 'example.com' --propver '4' --url "https://www.example.com/" -u "https://www.example.com/index/"
  $ akamai test-center test-suite generate-default < {filepath}/filename.json`

	TestSuiteEditUse     = "edit --id ID [--name NAME] [--description DESCRIPTION] [--unlocked | --locked] [--stateful | --stateless] [--property 'PROPERTY NAME' --propver 'PROPERTY VERSION' | --remove-property]"
	TestSuiteEditExample = `  $ akamai test-center test-suite edit --id 1001 --name 'Updated Example TS'
  $ akamai test-center test-suite edit --id 1001 --name 'Updated Example TS' --description 'TS for example.com' --property 'example.com' --propver '4' --unlocked
  $ akamai test-center test-suite edit --id 1001 --stateful --remove-property`

	TestSuiteImportUse     = "import"
	TestSuiteImportExample = `  $ akamai test-center test-suite import < {FILE_PATH}/FILE_NAME.json
  $ echo '{"testSuiteName":"ts1","testSuiteDescription":"ts1 description.","isLocked":true,"isStateful":false,"variables":[{"variableName":"host","variableValue":"www.akamai.com"}],"testCases":[]}' | akamai test-center test-suite import`

	TestSuiteListUse     = "list [--property 'PROPERTY NAME'] [--propver 'PROPERTY VERSION'] [-u 'USERNAME'] [-s 'SEARCH STRING']"
	TestSuiteListExample = `  $ akamai test-center test-suite list
  $ akamai test-center test-suite list --property 'example.com' --propver '4'
  $ akamai test-center test-suite list -u 'johndoe' -s 'regression'`
	TestSuiteListCommandAlias = "ls"

	TestSuiteManageUse     = "manage"
	TestSuiteManageExample = `  $ akamai test-center test-suite manage < {FILE_PATH}/FILE_NAME.json
  $ echo '{"testSuiteId":1,"testSuiteName":"ts1","testSuiteDescription":"ts1 description.","isLocked":true,"isStateful":false,"variables":[{"variableName":"host","variableValue":"www.akamai.com"}],"testCases":[]}' | akamai test-center test-suite manage`

	TestSuiteRemoveUse     = "remove [--id ID | --name NAME]"
	TestSuiteRemoveExample = `  $ akamai test-center test-suite remove --name "Test suite name"
  $ akamai test-center test-suite remove --id 12345`

	TestSuiteRemoveTestCaseUse     = "remove-test-case --test-suite-id ID --order-num ORDER_NUMBER"
	TestSuiteRemoveTestCaseExample = `  $ akamai test-center test-suite remove-test-case --test-suite-id 1001 --order-num 6`

	TestSuiteRestoreUse     = "restore [--id ID | --name NAME]"
	TestSuiteRestoreExample = `  $ akamai test-center test-suite restore --name "Test suite name"
  $ akamai test-center test-suite restore --id 12345`

	TestSuiteViewUse     = "view [--id ID | --name NAME] [--group-by test-request | condition | ip-version]"
	TestSuiteViewExample = `  $ akamai test-center test-suite view --id 1001
  $ akamai test-center test-suite view --name 'Example TS' --group-by test-request`
	TestSuiteViewCommandAliases = "export"

	ConditionTemplateUse     = "conditions"
	ConditionTemplateExample = `  $ akamai test-center conditions`
)

// Flag Names
const (
	FlagEdgerc         = "edgerc"
	FlagSection        = "section"
	FlagAccountKey     = "account-key"
	FlagJson           = "json"
	FlagForceColor     = "force-color"
	FlagProperty       = "property"
	FlagPropver        = "propver"
	FlagUrl            = "url"
	FlagEnv            = "env"
	FlagAddHeader      = "add-header"
	FlagModifyHeader   = "modify-header"
	FlagFilterHeader   = "filter-header"
	FlagTestSuiteId    = "test-suite-id"
	FlagTestSuiteName  = "test-suite-name"
	FlagIpVersion      = "ip-version"
	FlagCondition      = "condition"
	FlagId             = "id"
	FlagName           = "name"
	FlagDescription    = "description"
	FlagStateFul       = "stateful"
	FlagUnlocked       = "unlocked"
	FlagStateless      = "stateless"
	FlagLocked         = "locked"
	FlagRemoveProperty = "remove-property"
	FlagUser           = "user"
	FlagSearch         = "search"
	FlagOrderNumber    = "order-num"
	FlagGroupBy        = "group-by"
	FlagHelp           = "help"
	FlagVersion        = "version"
)

const (
	FlagUrlShortHand          = "u"
	FlagUserShortHand         = "u"
	FlagSearchShortHand       = "s"
	FlagEnvShortHand          = "e"
	FlagAddHeaderShortHand    = "a"
	FlagModifyHeaderShortHand = "m"
	FlagFilterHeaderShortHand = "f"
	FlagIpVersionShortHand    = "i"
	FlagConditionShortHand    = "c"
	FlagHelpShortHand         = "h"
	FlagVersionShortHand      = "v"
)

// Flag Values
const (
	FlagEdgercDefaultValue    = "~/.edgerc"
	FlagSectionDefaultValue   = "test-center"
	FlagIpVersionDefaultValue = "V4"
)

// Environment variable Constants
const (
	DefaultEdgercPathKey    = "AKAMAI_EDGERC"
	DefaultEdgercSectionKey = "AKAMAI_EDGERC_SECTION"
	DefaultJsonOutputKey    = "AKAMAI_OUTPUT_JSON"
)
