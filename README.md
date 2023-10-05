<!-- TOC ignore:true --> 
# Test Center CLI

<!-- TOC -->

- [Get started with the Test Center CLI](#get-started-with-the-test-center-cli)
    - [Install Test Center CLI](#install-test-center-cli)
    - [Stay up to date](#stay-up-to-date)
    - [API credentials](#api-credentials)
- [Concepts](#concepts)
- [CLI workflows](#cli-workflows)
    - [Run a test](#run-a-test)
    - [Manage Test Center objects using JSON](#manage-test-center-objects-using-json)
    - [Manage Test Center objects using granular commands](#manage-test-center-objects-using-granular-commands)
- [Available operations and commands](#available-operations-and-commands)
    - [Help](#help)
    - [Run a test](#run-a-test)
    - [Run a test using JSON](#run-a-test-using-json)
    - [List test runs](#list-test-runs)
    - [Get a test run](#get-a-test-run)
    - [List test suites](#list-test-suites)
    - [Get raw request and response for a test run or test case execution](#get-raw-request-and-response-for-a-test-run-or-test-case-execution)
    - [Get log lines for a test case execution](#get-log-lines-for-a-test-case-execution)
    - [Create a test suite](#create-a-test-suite)
    - [Import a test suite](#import-a-test-suite)
    - [Generate a default test suite for a property](#generate-a-default-test-suite-for-a-property)
    - [Edit a test suite](#edit-a-test-suite)
    - [Edit a test suite using JSON](#edit-a-test-suite-using-json)
    - [Remove a test suite](#remove-a-test-suite)
    - [Restore a test suite](#restore-a-test-suite)
    - [Get a test suite's overview](#get-a-test-suites-overview)
    - [Get a test suite with child objects](#get-a-test-suite-with-child-objects)
    - [Add a test case to a test suite](#add-a-test-case-to-a-test-suite)
    - [Remove a test case from a test suite](#remove-a-test-case-from-a-test-suite)
    - [Edit a test case in a test suite](#edit-a-test-case-in-a-test-suite)
    - [List test cases](#list-test-cases)
    - [Get a specific test case](#get-a-specific-test-case)
    - [Create variables](#create-variables)
    - [Edit a variable](#edit-a-variable)
    - [List variables](#list-variables)
    - [Get a specific variable](#get-a-specific-variable)
    - [Remove a specific variable](#remove-a-specific-variable)
    - [Check how functions work using JSON](#check-how-functions-work-using-json)
    - [List created test requests](#list-created-test-requests)
    - [List supported conditions](#list-supported-conditions)
    - [List created conditions](#list-created-conditions)
- [Available flags](#available-flags)
    - [edgerc](#edgerc)
    - [section](#section)
    - [account-key](#account-key)
    - [force-color](#force-color)
    - [help](#help)
    - [version](#version)
    - [json](#json)
- [Exit codes](#exit-codes)
- [Windows 10 2018 version](#windows-10-2018-version)
- [Notice](#notice)

<!-- /TOC -->

# Get started with the Test Center CLI

Test Center is a testing tool that checks the effect of configuration changes on your web property. Use this CLI as part of your testing protocol to increase your confidence in the safety and accuracy of your configuration changes.

> **_Breaking changes_**: Due to the active development of the new Test Center experience, we implemented the following changes to the CLI: <br> 1. The `--section` flag default value changed from `test-center` to `default`. <br> 2. The `conditions` command changes into `condition`; the parent command for all operations related to conditions. To get the template, you need to run the `condition template` command. <br> 3. The `test` command changes into the parent command for all operations related to test runs. You can now [get](#get-a-test-run), [list](#list-test-runs), or [create](#run-a-test) a test run. <br> 4. The subcommands for test suites changed their names: <br> - The `add` command got renamed to `create`. <br> -  The `edit` command got renamed to `update`. <br> - The `view` command got renamed to `get-with-child-objects`.  <br> -  The `import` command got renamed to `create-with-child-objects`. <br> -  The `manage` command got renamed to `update-with-child-objects`. <br> 5. The `--group-by` flag for test suite now accepts the `test-request`, `condition`, or `client-profile` value. <br> 6. The `add-test-case` and `remove-test-case` commands got moved under the `test-case` command and renamed to `create` and `remove`.<br> 7. The `-v` shorthand for the `--version` flag got removed. <br> 8. The `-i` shorthand for the `--ip-version` flag got removed.  <br> 9. The `--property` flag got renamed to `--property-name` with shorthand `-p`. <br> 10. The `--propver` flag got renamed to `--property-version` with shorthand `-v`.


## Install Test Center CLI

To install this CLI, you need the [Akamai CLI](https://github.com/akamai/cli) package manager. Once you install the Akamai CLI, run this command:

`akamai install test-center`

## Stay up to date

To make sure you always use the latest version of the CLI, run this command:  

`akamai update test-center`  

## API credentials
Akamai-branded packages use a `.edgerc` file for standard EdgeGrid authentication. By default, CLI looks for credentials in your `$HOME` directory.

You can override both the file location and the credential section by passing the [--edgerc](#edgerc) or [--section](#section) flags to each command.

To set up your `.edgerc` file, see [Get started with APIs](https://techdocs.akamai.com/developer/docs/set-up-authentication-credentials#add-credential-to-edgerc-file).

# Concepts

- **Test suite**. Test suites act as containers for test cases. You can add a name and description to a test suite to provide more details about the test suite and included test cases. You can also set if the test suite needs to be locked or stateful. Test suites can be tested as test objects associated with a property version or on their own.

- **Locked test suite**. Locked test suites can be modified only by their editors and owners. Test Center users who create locked test suites automatically become their owners. Owners can designate owners or editors and other users can request edit access. To learn how, see [Give the edit access to a locked test suite](https://techdocs.akamai.com/test-ctr/docs/functional-objects#give-the-edit-access-to-a-locked-test-suite).

- **Stateful test suite**. Stateful test suites are test suites within which test cases are executed based on the order number defined for each test case. Cookies and session information are retained for subsequent test cases.

- **Functional test case**. A test case in functional testing is the smallest unit of testing. It includes all settings for which the test needs to be run: conditions, test requests (combination of URL and headers), and IP versions.

- **Property version**. Property version refers to a Property Manager property version.

- **Test results**. A test result for functional testing is a comparison of the expected value with the actual value. It can be either *Passed* or *Failed*. *Passed* means that the *Expected* result of the test was the same as the *Actual* result. *Failed* means that the *Expected* result of the test was different from the *Actual* result. To learn more about Functional testing results, see [Test results concepts](https://techdocs.akamai.com/test-ctr/docs/glossary) and [Functional testing results](https://techdocs.akamai.com/test-ctr/docs/test-run-results#functional-testing-results).

- **Variables and functions**. Variables allow you to reuse specific values in test cases' input fields. They enable you to create test cases with complex metadata and run very specific tests. Variables can be assigned statically or dynamically. To extract the value from the test case response and assign it to a variable dynamically, you need to use functions. To learn more, see [Variables](https://techdocs.akamai.com/test-ctr/reference/variables-overview).

# CLI workflows

There are three ways you can use this CLI.

## Run a test
You can run a test for a property version, test suite, or test case using a [CLI command](#run-a-test-1) or [JSON input](#run-a-test-using-json).

## Manage Test Center objects using JSON
Here are the commands supporting JSON input:
  - [Generate a default test suite for a property](#generate-a-default-test-suite-for-a-property). Generates a default test suite for a property version for you to [import](#import-a-test-suite).
  - [Import a test suite](#import-a-test-suite). Imports to Test Center a test suite with test cases, variables, and property version association.
  - [Get a test suite with child objects](#get-a-test-suite-with-child-objects). Fetches, or exports, test suite's details that you can save and import on a different account or [edit](#edit-a-test-suite-using-json). You can also use this operation to clone test suites within your account.
  - [Edit a test suite using JSON](#edit-a-test-suite-using-json).
  - [Run a test using JSON](#run-a-test-using-json).

## Manage Test Center objects using granular commands
Here are the commands you can use to manage Test Center objects from the CLI:

  - [Run a test](#run-a-test)
  - [List test runs](#list-test-runs)
  - [Get a test run](#get-a-test-run)
  - [Get raw request and response for a test run or test case execution](#get-raw-request-and-response-for-a-test-run-or-test-case-execution)
  - [Get log lines for a test case execution](#get-log-lines-for-a-test-case-execution)
  - [List test suites](#list-test-suites)
  - [Create a test suite](#create-a-test-suite)
  - [Edit a test suite](#edit-a-test-suite)
  - [Remove a test suite](#remove-a-test-suite)
  - [Restore a test suite](#restore-a-test-suite)
  - [Get a test suite's overview](#get-a-test-suites-overview)
  - [Get a test suite with child objects](#get-a-test-suite-with-child-objects)
  - [Add a test case to a test suite](#add-a-test-case-to-a-test-suite)
  - [Remove a test case from a test suite](#remove-a-test-case-from-a-test-suite)
  - [Edit a test case in a test suite](#edit-a-test-case-in-a-test-suite)
  - [List test cases](#list-test-cases)
  - [Get a specific test case](#get-a-specific-test-case)
  - [Create variables](#create-variables)
  - [Edit a variable](#edit-a-variable)
  - [List variables](#list-variables)
  - [Get a specific variable](#get-a-specific-variable)
  - [Remove a specific variable](#remove-a-specific-variable)
  - [List created test requests](#list-created-test-requests)
  - [List supported conditions](#list-supported-conditions)
  - [List created conditions](#list-created-conditions)

# Available operations and commands

## Help
The `help` command returns an overview of available commands and flags.

## Run a test
The `test run` command runs a test for a specific test suite, single test case, or a property version.

**Command**:
- To run a test for a test suite: `test run [--test-suite-id ID | --test-suite-name 'NAME'] --env STAGING|PRODUCTION`
- To run a test for a property version: `test run [--property-name 'PROPERTY NAME' | --property-id 'ID' --property-version 'PROPERTY VERSION']  --env STAGING|PRODUCTION`
- To run a test for a simple test case: `test run [-url URL -condition CONDITION STATEMENT --ip-version v4|v6 [--client CURL| CHROME] [--request-method GET|HEAD|POST] [--request-body REQUEST_BODY] [--encode-request-body] [--add-header 'name: value' ...] [--modify-header 'name: value' ...] [--filter-header name ...]] --env STAGING|PRODUCTION`, where:

  - `ID` and `NAME` specify the test suite you want to run the test for. To get these values, run the [List test suites](#list-test-suites) operation. You need to provide either of these flags: `--test-suite-id` or `--test-suite-name`.
  - `PROPERTY NAME` or `ID` and `PROPERTY VERSION` specify the property in Property Manager you want to run the test for. To get these values you can use the [Property Manager CLI](https://github.com/akamai/cli-property-manager) and the `list-properties|lpr` operation. Provide either the `--property-name` or `--property-id` flag.
  - `URL` is the fully qualified URL of the resource to test. It needs to contain the protocol, hostname, path, and any applicable string parameters — for example, *https://www.example.com*. This flag needs to be combined with `--ip-version` and `--condition` flags.
  - `CONDITION STATEMENT` is a condition statement you want to run the test for. To get the list of created conditions to reuse, run the [List created conditions](#list-created-conditions) operation and to create a new condition, run the [List supported conditions](#list-supported-conditions) operation. Make sure to replace default values in `" "` with your own. The `--condition` flag needs to be combined with `--url` and `--ip-version` flags.
  - the `--ip-version` flag specifies the IP version to execute the test case over, either `v4` or `v6`. This flag is optional, set to `v4` by default. This flag needs to be combined with `--url` and `--condition` flags.
  - the `--client` flag specifies the client profile to execute the test case over, either `CURL` or `CHROME`. It's set to `CURL` by default. This flag is optional.
  - the `--request-method` flag specifies the request method for the test case, either `GET`, `HEAD` or `POST`. It's set to `GET` by default. This flag is optional.
  - the `--request-body` flag adds `REQUEST_BODY` to the request. This flag is optional and applicable only if `--client` is set to `CURL` and `--request-method` to `POST`.
  - the `--encode-request-body` flag encodes `REQUEST BODY`. This flag is optional and applicable only if `--client` is set to `CURL` and `--request-method` to `POST`.
  - the `--add-header` and `--modify-header` flags specify the request headers to respectively added or modify by the request. Headers should follow the format `name: value`. These flags are optional and accept multiple values. You can also use these flags to provide Pragma headers. See [Pragma headers](https://techdocs.akamai.com/edge-diagnostics/docs/pragma-headers) for the list of supported values. These flag needs to be combined with `--url`, `--ip-version`, and `--condition` flags.
  - the `--filter-header` flag filters the header from the request. Provide only the `name` of the header. This flag is optional. This flag needs to be combined with `--url`, `--ip-version`, and `--condition` flags.
  - the `--env` flag specifies the environment you want to run the test on, either `STAGING` or `PRODUCTION`. This flag is optional, set to `STAGING` by default.

**Examples**: 
- `akamai test-center test run --test-suite-name 'Regression test cases for example.com'`
- `akamai test-center test run --property-name 'example.com' --property-version '26'`
- `akamai test-center test run --property-id '438285' --property-version '26' --env PRODUCTION`
- `akamai test-center test run --url 'https://example.com/' --condition 'Response code is one of "200"' --ip-version 'V6' --modify-header 'Accept: application/json'`

**Expected output**: Once you submit the test run, it may take few minutes for Test Center to execute the test. To learn more about returned test results, check [Functional testing results](https://techdocs.akamai.com/test-ctr/docs/test-run-results#functional-testing-results). For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/post-test-runs) description.

## Run a test using JSON
The `test` command runs a test for a specific test suite, single test case, or a property version using a JSON file or JSON input. You can use the [API documentation](https://techdocs.akamai.com/test-ctr/reference/post-test-runs) to create the JSON file. Add your values to `BODY PARAMS` fields, copy the body of your request from the CURL code sample, and save it as a JSON file. To get `testCaseId` values run the `test-suite view` command for a specific test suite with the `--json` flag.

**Command**:

To import a specific file from your computer: `test-center test run < {FILE_PATH}/FILE_NAME.json`, where `FILE_PATH` and `FILE_NAME` are respectively location and name of the file to import.

JSON example to run a test for a property version:
```
{
   "functional":{
      "configVersionExecutions":[
         {
            "testSuiteExecutions":[
               {
                  "testSuiteId":31306
               },
               {
                  "testCaseExecutions":[
                     {
                        "testCaseId":41
                     }
                  ],
                  "testSuiteId":87
               }
            ],
            "configVersionId":22556
         }
      ]
   },
   "targetEnvironment": "STAGING"
}
```

JSON example to run a test for a test suite:
```
{
   "functional":{
      "testSuiteExecutions":[
         {
            "testCaseExecutions":[
               {
                  "testCaseId":65
               },
               {
                  "testCaseId":71
               },
               {
                  "testCaseId":98
               }
            ],
            "testSuiteId":281
         }
      ]
   },
   "targetEnvironment": "STAGING"
}
```
JSON example to run a test for a single test case:
```
{
   "functional":{
      "testCaseExecution":{
         "testRequest":{
            "testRequestUrl":"https://www.example.com"
         },
         "condition":{
            "conditionExpression":"Log request details - Referrer header is not logged"
         },
         "clientProfile":{
            "ipVersion":"IPV4"
         }
      }
   },
   "targetEnvironment":"STAGING"
}
```

**Expected output**: Once you submit the test run, it may take a few minutes for Test Center to execute the test. To learn more about returned test results, check [Functional testing results](https://techdocs.akamai.com/test-ctr/docs/test-run-results#functional-testing-results). For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/post-test-runs) description.

## List test runs
The `test list` command returns test runs created by users of your account.

**Command**: `test list`

**Expected output**: List of created test runs. You can use a returned `testRunId` to [get a specific test run results](#get-a-test-run). For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/get-test-runs) description.

## Get a test run
The `test get` command returns details and results of a specific test run.

**Command**: `test get --test-run-id ID`, where `ID` is the `testRunId` of the test run you want to get the details of. You can get this value with the [List test runs](#list-test-runs) operation.  

**Example**: `test get --test-run-id 2500`

**Expected output**: Returns the results of a test run. To learn more about returned test results, check [Functional testing results](https://techdocs.akamai.com/test-ctr/docs/test-run-results#functional-testing-results). For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/get-test-run) description.

## List test suites
The `test-suite list` command lists all test suites. You can filter the results for test suites created by a user, a property the test suite is associated with, or a string from the test suite's name or description.
The list also includes the recently deleted test suites that you can [restore](#restore-a-test-suite).


**Command**: `test-suite list [--property-name 'PROPERTY NAME'| --property-id 'ID'] [--property-version 'PROPERTY VERSION'] [-user 'USERNAME'] [--search STRING]`, where:

- the `--property-name` or `--property-id` flags filter the results for a specific `PROPERTY NAME` or `ID` which the test suite is associated with. The `--property-version` flag filters the results further for a specific `PROPERTY VERSION`. If applicable, provide either the `--property-name` or `--property-id` flag.
- the `--user` flag filters the results for a `USERNAME` who created, edited, or deleted the test suite.
- the `--search` flag filters the results for a specific `STRING` in the test suite's name or description.

You can combine multiple flags to narrow the list.

**Examples**
- `akamai test-center test-suite list`
- `akamai test-center test-suite list --property-name 'example.com' --property-version '4'`
- `akamai test-center test-suite list -u 'johndoe' --search 'regression'`

**Expected output**: The response lists all test suites matching the requested filters. You can use the returned test suite ID to [get a test suite's overview](#get-a-test-suites-overview) or [get a test suite with child objects](#get-a-test-suite-with-child-objects). For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/get-test-suites) description.

## Get raw request and response for a test run or test case execution
The `raw-request-response` command returns a raw request and response for a test run or a test case execution.

**Command**: `raw-request-response [--test-run-id TEST_RUN_ID | --tcx-id TEST_CASE_EXECUTION_ID]`, where:

- `TEST_RUN_ID` specifies the test run ID you want to get the transaction details of. To get this value, run the [List test runs](#list-test-runs) operation.
- `TEST_CASE_EXECUTION_ID` specifies the test case execution you want to get the details of. To get this value, run the [Get a test run](#get-a-test-run) operation. 

Provide either the `--tcx-id` or `--test-run-id` flag.

**Example**: `akamai test-center raw-request-response --test-run-id 2009`

**Expected output**: The response provides the raw request and response.

## Get log lines for a test case execution
The `log-lines` command returns log lines for a specific test case execution.

**Command**: `log-lines --tcx-id TEST_CASE_EXECUTION_ID`, where the `TEST_CASE_EXECUTION_ID` specifies the execution you want to get the logs for. To get this value, run the [Get a test run](#get-a-test-run) operation. 

**Example**: `akamai test-center test log-lines --tcx-id 1`

**Expected output**: The response provides log lines for the execution.

## Create a test suite
The `test-suite create` command creates a new test suite.

**Command**: `test-suite create --name 'NAME' [--description 'DESCRIPTION'] [--unlocked] [--stateful]  [--property-name 'PROPERTY NAME' | --property-id 'ID' --property-version 'PROPERTY VERSION']`, where:

- `NAME` is the name of the test suite.
- `DESCRIPTION` is the description for the test suite. The `--description` flag is optional.
- the `--unlocked` flag unlocks the test suite. This flag is optional. By default, all test suites are [locked](#concepts).
- the `--stateful` flag makes the test suite stateful. This flag is optional. By default, all test suites are [stateless](#concepts).
- the `--property-name` or `--property-id` and `--property-version` flags associate the test suite with a specific property version. `PROPERTY NAME` is the name of the property in Property Manager, ID its unique identifier, and `PROPERTY VERSION`, its appropriate version. You can use [Property Manager CLI](https://github.com/akamai/cli-property-manager) and the `list-properties|lpr` operation to get these values. These flags are optional. If appropriate, provide either `--property-name` or  `--property-id` and combine it with `--property-version`.

**Examples**:
- `akamai test-center test-suite create --name 'new test suite'`
- `akamai test-center test-suite create --name 'new test suite' --description 'TS for example.com' --unlocked --stateful --property-name 'example.com' --property-version '4'`

**Expected output**: The response includes details about the created test suite. It also includes the ID of the test suite that you can use in other operations — for example, to [add a test case to a test suite](#add-a-test-case-to-a-test-suite). For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/post-test-suites) description.

## Import a test suite
The `test-suite create-with-child-objects` command imports a test suite from a JSON file or standard JSON input. You can use the [Generate a default test suite for a property](#generate-a-default-test-suite-for-a-property) operation or the [API documentation](https://techdocs.akamai.com/test-ctr/reference/post-test-suites-with-child-objects) to create the JSON file. Add your values to `BODY PARAMS` fields, copy the body of your request from the CURL code sample, and save it as a JSON file.

**Command**:
- To import a specific file from your computer: `test-center test-suite create-with-child-objects < {FILE_PATH}/FILE_NAME.json`, where `FILE_PATH` and `FILE_NAME` are respectively location and name of the file to import.
- To import an outputted string: `echo '{"testSuiteName":"TEST_SUITE_NAME","testSuiteDescription":"TEST_SUITE_DESCRIPTION","isLocked":true|false,"isStateful":true|false,"variables":[{"variableName":"VARIABLE_NAME","variableValue":"VARIABLE_VALUE"}],"testCases":[]}' | akamai test-center test-suite create-with-child-objects`, where `TEST_SUITE_NAME`,`TEST_SUITE_DESCRIPTION`, `VARIABLE_NAME`, and `VARIABLE_VALUE` are your values for the test suite.

**Examples**:
- `akamai test-center test-suite create-with-child-objects < ./users/johndoe/documents/test_suite_prop19.json`
- `echo '{"testSuiteName":"test_suite_prop19","testSuiteDescription":"test suite for property version 19","isLocked":true,"isStateful":false,"variables":[{"variableName":"host","variableValue":"www.akamai.com"}],"testCases":[]}' | akamai test-center test-suite create-with-child-objects`

**Expected output**: The response includes details about the imported test suite. It also includes the ID of the test suite that you can use in other operations — for example, to [edit the test suite using JSON](#edit-a-test-suite-using-json). For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/post-test-suites-with-child-objects) description.


## Generate a default test suite for a property
The `test-suite generate-default` command generates a default test suite with test cases for a specific property. Based on property settings and its behaviors and the `--url` flag value, Test Center generates a test suite object with test cases and variables for you to modify and add to Test Center using the [Import a test suite](#import-a-test-suite) operation.

**Command**: `test-suite generate-default --property-name 'PROPERTY NAME'| --property-id 'ID' --property-version 'PROPERTY VERSION' --url URL ...`, where:

- `PROPERTY NAME` or `ID` and `PROPERTY VERSION` specify the property in Property Manager you want to generate a test suite for. To get these values you can use the [Property Manager CLI](https://github.com/akamai/cli-property-manager) and the `list-properties|lpr` operation. Provide either the `--property-name` or `--property-id` flag.
- `URL` is the fully qualified URL of the property hostname. The `--url` flag can be used multiple times.


**Examples**:
- `akamai test-center test-suite generate-default --property-name 'example.com' --property-version '4' --url "https://www.example.com/" -u "https://www.example.com/index/"`

**Expected output**: The response includes details about the generated response and included test cases. For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/post-auto-generate-test-suite) description.

## Edit a test suite
The `test-suite update` command edits basic data of a specific test suite. Provide only data you want to edit in the original test suite.

**Command**: `test-suite update --id ID [--name NAME] [--description DESCRIPTION] [--unlocked | --locked] [--stateful | --stateless] [--property-name 'PROPERTY NAME' | --property-id 'ID' --property-version 'PROPERTY VERSION' | --remove-property ]`, where:

- `ID` is the identifier of the test suite you want to edit. To get this value, run the [List test suites](#list-test-suites) operation. You need to provide either of these flags: `--id` or `--name`.
- `NAME` is the new test suite's name. 
- `DESCRIPTION` is the new description for the test suite.
- `unlocked`, `locked`, `stateful`, `stateless` flags update the status of the test suite.
- `--property-name` or `--property-id` and `--property-version` update the test suite's association to the property with a specific `PROPERTY NAME` or `ID` and `PROPERTY VERSION`. If applicable, provide either `--property-name` or `--property-id` or `--remove-property`. 
- the `--remove-property` removes the association to a property version. If applicable, provide either this flag or `--property-name` or `--property-id` with `--property-version`.

**Examples**:
- `akamai test-center test-suite update --id 1001 --name 'Updated test suite'`
- `akamai test-center test-suite update --id 1001 --name 'Updated test suite' --description 'Test suite for example.com' --property-name 'example.com' --property-version '4' --unlocked`
- `akamai test-center test-suite update --id 1001 --stateful --remove-property`

**Expected output**: The response returns the updated test suite. For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/put-test-suite) description.

## Edit a test suite using JSON
The `test-suite update-with-child-objects` command uses the JSON input to edit a specific test suite. Provide the whole test suite object, together with test cases and variables, to include in Test Center. Only data provided in the latest JSON input will be saved. You can use the [API documentation](https://techdocs.akamai.com/test-ctr/reference/put-test-suites-with-child-objects) to create the JSON file. Add your values to `BODY PARAMS` fields, copy the body of your request from the CURL code sample, and save it as a JSON file.

**Command**:
- To edit the test suite with a specific file from your computer: `test-center test-suite update-with-child-objects < {FILE_PATH}/FILE_NAME.json`, where `FILE_PATH` and `FILE_NAME` are respectively location and name of the file to import.
- To edit the test suite with outputted string: `echo '{"testSuiteId":ID,"testSuiteName":"TEST_SUITE_NAME","testSuiteDescription":"TEST_SUITE_DESCRIPTION","isLocked":true|false,"isStateful":true|false,"variables":[{"variableName":"VARIABLE_NAME","variableValue":"VARIABLE_VALUE"}],"testCases":[]}' | akamai test-center test-suite update-with-child-objects`, where:
  - `ID` is the unique identifier of the test suite you want to edit.
  - `TEST_SUITE_NAME`,`TEST_SUITE_DESCRIPTION`, `VARIABLE_NAME`, and `VARIABLE_VALUE` are your values for the edited test suite.

**Examples**:
- `akamai test-center test-suite update-with-child-objects < ./users/johndoe/documents/test_suite_prop19.json`
- `echo '{"testSuiteName":"test_suite_prop19","testSuiteDescription":"test suite for property version 19","isLocked":true,"isStateful":false,"variables":[{"variableName":"host","variableValue":"www.akamai.com"}],"testCases":[]}' | akamai test-center test-suite update-with-child-objects`

**Expected output**: Successful operation confirmed.

## Remove a test suite
The `test-suite remove` command removes a specific test suite from Test Center. Test suites can be [restored](#restore-a-test-suite) for 30 days since their removal.

**Command**: `test-suite remove [--id ID | --name "NAME"]`, where `ID` and `NAME` specify the test suite you want to remove. To get these values, run the [List test suites](#list-test-suites) operation. You need to provide either of these flags: `--id` or `--name`.

**Example**:
- `akamai test-center test-suite remove --name "Test suite name"`
- `akamai test-center test-suite remove --id 12345`

**Expected output**: Successful operation confirmed.

## Restore a test suite
The `test-suite restore` command restores a removed test suites and included test cases. Test suites can be restored for 30 days since their removal.

**Command**: `test-suite restore [--id ID | --name "NAME"]`, where `ID` and `NAME` specify the test suite you want to restore. To get these values, run the [List test suites](#list-test-suites) operation. You need to provide either of these flags: `--id` or `--name`.

**Example**:
- `akamai test-center test-suite restore --name "test_suite_prop19"`
- `akamai test-center test-suite restore --id 12345`

**Expected output**: Successful operation confirmed.

## Get a test suite's overview
The `test-suites get` command returns basic data about a test suite. You can group the included test cases by a test request, condition, or IP version.

**Command**: `test-suite get [--test-suite-id ID | --test-suite-name 'NAME']`, where the `ID` and `NAME` specify the test suite you want to get the details of. To get these values, run the [List test suites](#list-test-suites) operation. You need to provide either of these flags: `--id` or `--name`.

**Example**: `akamai test-center test-suite get --test-suite-id 1001`

**Expected output**: The response includes overview of the test suite. To get the test suite with included objects, test cases and variables, run the [Get a test suite with child objects](#get-a-test-suite-with-child-objects) operation. For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/get-test-suite) description.

## Get a test suite with child objects
The `test-suites get-with-child-objects` command returns details of a test suite and all included objects, test cases and variables. You can group the included test cases by a test request, condition, or IP version.

**Command**: `test-suite get-with-child-objects [--test-suite-id ID | --test-suite-name 'NAME'] [--group-by test-request | condition | client-profile]`, where:

- `ID` and `NAME` specify the test suite you want to get the details of. To get these values, run the [List test suites](#list-test-suites) operation. You need to provide either of these flags: `--id` or `--name`
- the `--group-by` flag specifies grouping of the included test cases, by `test-request`, `condition`, or `client-profile`. This flag is optional.

**Examples**:
- `akamai test-center test-suite get-with-child-objects --test-suite-id 1001`
- `akamai test-center test-suite get-with-child-objects --test-suite-name 'test_suite_prop19' --group-by test-request`

**Expected output**: The response includes details of the test suite. Run this operation with the `--json` flag so that you can save the returned object and [import it](#import-a-test-suite) on a different account or [edit it](#edit-a-test-suite-using-json). You can also use the response to clone test suites within your account. For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/get-test-suites-with-child-objects) description.

## Add a test case to a test suite
The `test-case create` command adds a functional test case to a specific test suite.

**Command**: `test-case create [--test-suite-id ID | --test-suite-name 'NAME'] --url URL --condition CONDITION_STATEMENT [--ip-version v4|v6] [--client CURL| CHROME] [--request-method GET|HEAD|POST] [--request-body REQUEST_BODY] [--encode-request-body] [--set-variable VARIABLE_NAME:VARIABLE_VALUE ...] [--add-header header ...] [--modify-header header ...] [--filter-header header ...]`, where:

- `ID` or `NAME` specify the test suite you want to add the test case to. To get these values, run the [List test suites](#list-test-suites) operation. You need to provide either of these flags: `--test-suite-id` or `--test-suite-name`.
- `URL` is the fully qualified URL of the resource to test. It needs to contain the protocol, hostname, path, and any applicable string parameters. For example *https://www.example.com*.
- `CONDITION_STATEMENT` is one of the condition statements from the list of supported conditions with entered required values. To get the available condition statements, run the [List created conditions](#list-created-conditions) operation and to create a new condition, run the [List supported conditions](#list-supported-conditions) operation. Make sure to substitute default values in `" "` with your own.
- the `--ip-version` flag specifies the IP version to execute the test suite over, either `v4` or `v6`. It's set to `v4` by default. This flag is optional.
- the `--client` flag specifies the client profile to execute the test suite over, either `CURL` or `CHROME`. It's set to `CURL` by default. This flag is optional.
- the `--request-method` flag specifies the request method for the test case, either `GET`, `HEAD` or `POST`. It's set to `GET` by default. This flag is optional.
- the `--request-body` flag adds `REQUEST_BODY` to the request. This flag is optional and applicable only if `--client` is set to `CURL` and `--request-method` to `POST`.
- the `--encode-request-body` flag encodes `REQUEST BODY`. This flag is optional and applicable only if `--client` is set to `CURL` and `--request-method` to `POST`.
- the `--set-variable` flag defines variables for the test case, provide in the format: `VARIABLE_NAME:VARIABLE_VALUE`. This flag is optional. To learn more about variables, see [Variables](https://techdocs.akamai.com/test-ctr/reference/variables-overview).
- the `--add-header` and `--modify-header` flags specify the request headers to respectively add or modify by the request. Headers should follow the format `name: value`. These flags are optional. You can also use these flags to provide Pragma headers. See [Pragma headers](https://techdocs.akamai.com/edge-diagnostics/docs/pragma-headers) for the list of supported values.
- the `--filter-header` flag filters the header from the request. Provide only the `name` of the header. This flag is optional.

**Examples**:
- `akamai test-center test-case create --test-suite-id 1001 --url 'https://example.com/' --condition 'Response code is one of "200,201"'`
- `akamai test-center test-case create --test-suite-name 'Example TS' --u 'https://example.com/' -c 'Response code is one of "200"' -a 'Accept: text/html' -a 'X-Custom: 123' -m 'User-Agent: Mozilla' -f 'Accept-Language'`

**Expected output**: The response includes details of the test case created for the test suite. For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/post-test-cases) description.

## Remove a test case from a test suite
The `test case remove` command removes a test case from a test suite. Removed test cases can't be restored.

**Command**: `test-case remove --test-suite-id ID [--order-num ORDER_NUMBER | --test-case-id TEST_CASE_ID] `, where:

- `ID` is the unique identifier of the test suite you want to remove the test case from. To get this values, run the [List test suites](#list-test-suites) operation.
- `ORDER_NUMBER` is the order number of the test case you want to remove. To get this value, run the [Get a test suite's overview](#get-a-test-suites-overview) operation.
- `TEST_CASE_ID` specifies the test case to be removed. To get this value, run the [Get a test suite's overview](#get-a-test-suites-overview) operation. Provide either the `--test-case-id` or `--order-num` flag.

**Examples**: `akamai test-center test-suite remove --test-suite-id 1001 --order-num 6`

**Expected output**: Successful operation confirmed.

## Edit a test case in a test suite
The `test case update` command edits a test case with a specific test case in a test suite. Provide only data you want to edit in the original test case.

**Command**: `test-case update [--test-suite-id ID | --test-suite-name 'NAME'] --url URL --condition CONDITION_STATEMENT [--ip-version v4|v6] [--client CURL| CHROME] [--request-method GET|HEAD|POST] [--request-body REQUEST_BODY] [--encode-request-body] [--set-variable VARIABLE_NAME:VARIABLE_VALUE ...] [--add-header header ...] [--modify-header header ...] [--filter-header header ...]`, where:

- `ID` or `NAME` specify the test suite you want to add the test case to. To get these values, run the [List test suites](#list-test-suites) operation. You need to provide either of these flags: `--test-suite-id` or `--test-suite-name`.
- `URL` is the fully qualified URL of the resource to test. It needs to contain the protocol, hostname, path, and any applicable string parameters. For example *https://www.example.com*.
- `CONDITION_STATEMENT` is one of the condition statements from the list of supported conditions with entered required values. To get the list of created conditions to reuse, run the [List created conditions](#list-created-conditions) operation and to create a new condition, run the [List supported conditions](#list-supported-conditions) operation. Make sure to substitute default values in `" "` with your own.
- the `--ip-version` flag specifies the IP version to execute the test suite over, either `v4` or `v6`. It's set to `v4` by default. This flag is optional.
- the `--client` flag specifies the client profile to execute the test suite over, either `CURL` or `CHROME`. It's set to `CURL` by default. This flag is optional.
- the `--request-method` flag specifies the request method for the test case, either `GET`, `HEAD` or `POST`. It's set to `GET` by default. This flag is optional.
- the `--request-body` flag adds `REQUEST_BODY` to the request. This flag is optional and applicable only if `--client` is set to `CURL` and `--request-method` to `POST`.
- the `--encode-request-body` flag encodes `REQUEST BODY`. This flag is optional and applicable only if `--client` is set to `CURL` and `--request-method` to `POST`.
- the `--set-variable` flag defines variables for the test case, provide in the format: `VARIABLE_NAME:VARIABLE_VALUE`. This flag is optional. To learn more about variables, see [Variables](https://techdocs.akamai.com/test-ctr/reference/variables-overview).
- the `--add-header` and `--modify-header` flags specify the request headers to respectively add or modify by the request. Headers should follow the format `name: value`. These flags are optional. You can also use these flags to provide Pragma headers. See [Pragma headers](https://techdocs.akamai.com/edge-diagnostics/docs/pragma-headers) for the list of supported values.
- the `--filter-header` flag filters the header from the request. Provide only the `name` of the header. This flag is optional.

**Examples**: 
- `akamai test-center test-case update --test-suite-id 1001 --test-case-id 101 --url 'https://example.com/' --condition 'Response code is one of "200,201"`
- `akamai test-center test-case update --test-suite-id 1001 --test-case-id 101 -u 'https://example.com/' -c 'Response code is one of "200"' -a 'Accept: text/html' -a 'X-Custom: 123' -m 'User-Agent: Mozilla' -f 'Accept-Language' -C curl -M POST --set-variable 'varName: varValue'`

**Expected output**: Successful operation confirmed.

## List test cases
The `test-case list` command lists all test cases. 

**Command**: `test-case list [--test-suite-name 'NAME'| --test-suite-id 'ID'] [--group-by test-request | condition | client-profile]`, where:

- the `--test-suite-name` or `--test-suite-id` flags filter the results for a specific test suite. Provide either the `--test-suite-name` or `--test-suite-id` flag.
- the `--group-by` flag specifies grouping of the included test cases, by `test-request`, `condition`, or `client-profile`. This flag is optional.

**Example**: `akamai test-center test-case list --test-suite-name 'Test suite for example.com' --group-by test-request`

**Expected output**: The response lists all test cases. You can use the returned test case ID to run the [Get a specific test case](#get-a-specific-test-case) operation. For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/get-test-cases) description. 

## Get a specific test case
The `test-case get` command returns details a specific test case. 

**Command**: `test-case get [--test-suite-id ID | --test-suite-name 'NAME'] --test-case-id TEST_CASE_ID`, where:

- `ID` and `NAME` specify the test suite you want to get the details of. To get these values, run the [List test suites](#list-test-suites) operation. You need to provide either of these flags: `--id` or `--name`
- `TEST_CASE_ID` specifies the test case you want to get the details of. To get these values, run the [Get a test suite's overview](#get-a-test-suites-overview) operation.

**Examples**: `akamai test-center test-suite get --test-suite-id 1001 --test-case-id 3`

**Expected output**: The response includes details of the test case. For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/get-test-case) description.

## Create variables
The `variable create` command creates variables for a test suite. You can use variables in a test request's URL or request headers and in condition expression, as a substitute of placeholders. To learn more, see [Variables](https://techdocs.akamai.com/test-ctr/reference/variables-overview).

**Command**: `variable create --test-suite-id ID --name VARIABLE_NAME [--value VALUE | --group-value H1: value1, value2 --group-value H2: value3, value4]`, where:

- `ID` specifies the test suite you want to create the variable for. To get this value, run the [List test suites](#list-test-suites) operation.
- `VARIABLE_NAME` is the name of the variable.
- For the value, you can either use the `--value` flag and provide a single `VALUE` or the `--group-value` flag to create variable groups. To learn more, see [Variable groups](https://techdocs.akamai.com/test-ctr/reference/variables-overview#variable-groups).

**Examples**:
- `akamai test-center variable create --test-suite-id 1001 --name url --value 'https://example.com/`
- `akamai test-center variable create --test-suite-id 1001 --name url --group-value hostName: https://example.com/,https://example.com/123 --group-value ResponseCodes: 200,300`

**Expected output**: The response includes details about the created variable. You can now reuse the variable by entering it in *{{ }}*. For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/post-variables) description.

## Edit a variable
The `variable update` command edits a specific variable. Provide only data you want to edit in the original variable.

**Command**: `update --test-suite-id ID --variable-id VARIABLE_ID --name NAME [--value VALUE | --group-value H1: value1, value2 --group-value H2: value3, value4]`, where:

- `ID` specifies the test suite you want to edit the variable for. To get this value, run the [List test suites](#list-test-suites) operation.
- `VARIABLE_ID` specifies the variable to edit. To get this value, run the [List variables](#list-variables) operation.
- `VARIABLE_NAME` is the name of the variable.
- For the value, you can either use the `--value` flag and provide a single `VALUE` or the `--group-value` flag to create variable groups. To learn more, see [Variable groups](https://techdocs.akamai.com/test-ctr/reference/variables-overview#variable-groups).

**Examples**:
- `akamai test-center variable update --test-suite-id 1001 --variable-id 1 --name url --value 'https://example.com/`
- `akamai test-center variable update --test-suite-id 1001 --name url --variable-id 1 --group-value hostName: https://example.com/,https://example.com/123 --group-value ResponseCodes: 200,300`

**Expected output**: The response returns the updated variable. You can now reuse the variable by entering it in *{{ }}*. For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/put-variables) description.

## List variables
The `variable list` command list all variables created for a test suite that you can further reuse.

**Command**: `variable list --test-suite-id ID`, where `ID` specifies the test suite you want to get the variables for. To get this value, run the [List test suites](#list-test-suites) operation.

**Example**: `akamai test-center --test-suite-id 2005`

**Expected output**: The response lists test suite's variables. For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/get-variables) description.

## Get a specific variable
The `variable get` command returns details for a specific variable. 

**Command**: `variable get --test-suite-id ID --variable-id VARIABLE_ID`, where:

- `ID` specifies the test suite you want to get the variable for. To get this value, run the [List test suites](#list-test-suites) operation.
- `VARIABLE_ID` specifies the variable you want to get the details of. To get this value, run the [List variables](#list-variables) operation.

**Examples**: `akamai test-center variable get --test-suite-id 1001 --variable-id 3`

**Expected output**: The response includes details of the variable. For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/get-variable) description.

## Remove a specific variable
The `variable remove` command removes a variable from a test suite. 

**Command**: `variable remove --test-suite-id ID --variable-id VARIABLE_ID`, where:

- `ID` specifies the test suite you want to remove the variable from. To get this value, run the [List test suites](#list-test-suites) operation.
- `VARIABLE_ID` specifies the variable you want to remove. To get this value, run the [List variables](#list-variables) operation.

**Examples**: `akamai test-center variable remove --test-suite-id 1001 --variable-id 4`

**Expected output**: Successful operation confirmed.

## Check how functions work using JSON
The `function try-it` command uses the JSON input to run a created function on sample data to check whether it returns the expected value. A function is valid for use in a test case if the response's `results` value returns only one value. To learn more about functions and variables, see [Variables](https://techdocs.akamai.com/test-ctr/reference/variables-overview).

**Command**: 
- To check a function with a specific file from your computer: `function try-it < {FILE_PATH}/FILE_NAME.json`, where `FILE_PATH` and `FILE_NAME` are respectively location and name of the file to import.
- To check a function with outputted string: `echo '"functionExpression": "FUNCTION_EXPRESSION","responseData": {SAMPLE_RESPONSE_DATA}' |  akamai test-center function try-it`, where:
  - `FUNCTION_EXPRESSION` is the function you want to test.
  - `SAMPLE_RESPONSE_DATA` is the sample response for Test Center to run the function evaluation on.

**Examples**:
- `akamai test-center function try-it < ./users/johndoe/documents/function.json`
- `echo '{"functionExpression": "fn_getResponseHeaderValue(headerName, regex)","responseData": {"response": {"status": 200,"statusText": "OK","httpVersion": "HTTP/1.1","headers": [{"name": "server", "value":"Apache/2.2.15 (CentOS)"}]}}}' | akamai test-center function try-it`

**Expected output**: The response is the matching result for the input function expression. For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/post-try-it) description.


## List created test requests
The `test-request list` command returns all test requests created by users of your account while creating test cases. You can further reuse these test requests in test cases.

**Command**: `test-request list`

**Expected output**: The response lists created test requests. For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/get-test-requests) description.

## List supported conditions
The `condition template` command returns the supported condition expressions you can use when creating test cases. Note that the statements contain default values in `" "`.  You need to replace them with your own values before creating the test case.

**Command**: `condition template`

**Expected output**: The response lists supported condition statements. For more details, you can check the [API response](https://techdocs.akamai.com/test-ctr/reference/get-test-catalog-template) description.

## List created conditions
The `condition list` command returns all conditions created by users of your account while creating test cases. You can further reuse these conditions in test suites or test cases.

**Command**: `condition list`

**Expected output**: The response lists created condition statements. 



# Available flags

You can use the following flags with all the listed commands.

## edgerc
The `--edgerc` flag changes the default path to the .edgerc file. This file contains the API credentials required to run all commands.
Without this flag, the user's home directory is used by default.

**Command**: `$akamai test-center --edgerc EDGERC_PATH [command]`

**Example**: `$akamai test-center --edgerc C:/users/johndoe/.edgerc test-suite list`

## section
The `--section` flag changes the default section name. The section name specifies which section of API credentials to read from the .edgerc file. Without this flag, the `default` section name is used by default.

**Command**: `$akamai test-center --section SECTION_NAME [command]`

**Examples**:
- `$akamai test-center --section default test-suite list`
- `$set AKAMAI_EDGERC_SECTION=default`

## account-key
The `--account-key` flag changes your account. When testing your configuration, you may need to switch between different accounts. To do this, run the required operation with the `--account-key` flag followed by the account ID of your choice.

**Command**: `$akamai test-center ----account-key ACCOUNT KEY [command]`

**Example**: `akamai test-center --account-key 1-1TJZFB test-suite list`

## force-color
The `--force-color` flag forces color to non-TTY output.

## help
The `--help` flag returns help for a command.

## version
The `--version` flag returns the version.

## json
The `--json` flag returns the information in JSON format.

# Exit codes
When you complete an operation, the CLI generates one of these exit codes:

- `0` (Success) - Indicates that the latest command or script executed successfully.
- `1` (Configuration error) - Indicates an error while loading the CLI.
- `2` - Indicates an error related to command arguments, missing flags, or mismatch exception.
- `3` - Indicates a parsing error in API request and response.
- `100-199` - Indicates a 4xx HTTP error. To get the HTTP code value add 300 to the returned exit code. 
- `200-255` - Indicates a 5xx HTTP error. To get the HTTP code value add 300 to the returned exit code.

# Windows 10 2018 version
If you're using Windows 10, 2018 version and you're having problems running the Test Center
CLI, we recommend you try the following workaround. In the downloaded repository, add the `.exe`
suffix to the `akamai-test-center` executable file.

# Notice

Copyright © 2022 Akamai Technologies, Inc.

Your use of Akamai's products and services is subject to the terms and provisions outlined in [Akamai's legal policies](https://www.akamai.com/us/en/privacy-policies/).
