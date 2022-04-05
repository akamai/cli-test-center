#Test Center CLI

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
  - [List test suites](#list-test-suites)
  - [Create a test suite](#create-a-test-suite)
  - [Import a test suite](#import-a-test-suite)
  - [Add a test case to a test suite](#add-a-test-case-to-a-test-suite)
  - [Generate a default test suite for a property](#generate-a-default-test-suite-for-a-property)
  - [Edit a test suite](#edit-a-test-suite)
  - [Edit a test suite using JSON](#edit-a-test-suite-using-json)
  - [Remove a test suite](#remove-a-test-suite)
  - [Remove a test case from a test suite](#remove-a-test-case-from-a-test-suite)
  - [Restore a test suite](#restore-a-test-suite)
  - [Get a test suite's details](#get-a-test-suites-details)
  - [List supported conditions](#list-supported-conditions)
  - [Run a test](#run-a-test-1)
  - [Run a test using JSON](#run-a-test-using-json)
- [Available flags](#available-flags)
  - [edgerc](#edgerc)
  - [section](#section)
  - [account-key](#account-key)
  - [help](#help-1)
  - [version](#version)
  - [json](#json)
- [Windows 10 2018 version](#windows-10-2018-version)
- [Notice](#notice)



# Get started with the Test Center CLI

Test Center is a testing tool that checks the effect of configuration changes on your web property. Use this CLI as part of your testing protocol to increase your confidence in the safety and accuracy of your configuration changes.

## Install Test Center CLI

To install this CLI, you need the [Akamai CLI](https://github.com/akamai/cli) package manager. Once you install the Akamai CLI, run this command:

`akamai install test-center` 

## Stay up to date

To make sure you always use the latest version of the CLI, run this command:  

`akamai update test-center`  

## API credentials
Akamai-branded packages use a `.edgerc` file for standard EdgeGrid authentication. By default, CLI looks for credentials in your `$HOME` directory.

You can override both the file location and the credentials section by passing the [--edgerc](#edgerc) or [--section](#section) flags to each command.

To set up your `.edgerc` file, see [Get started with APIs](https://techdocs.akamai.com/developer/docs/set-up-authentication-credentials#add-credential-to-edgerc-file).

# Concepts

- **Test suite**. Test suites act as containers for test cases. You can add a name and description to a test suite to provide more details about the test suite and included test cases. You can also set if the test suite needs to be locked or stateful. Test suites can be tested as test objects associated with a property version or on their own. 
  
- **Locked test suite**. Locked test suites can be modified only by their editors and owners. Test Center users who create locked test suites automatically become their owners. Owners can designate owners or editors and other users can request edit access. To learn how, see [Give the edit access to a locked test suite](https://techdocs.akamai.com/test-ctr/docs/test-suite-edit-access).

- **Stateful test suite**. Stateful test suites are test suites within which test cases are executed based on the order number defined for each test case. Cookies and session information are retained for subsequent test cases.

- **Functional test case**. A test case in functional testing is the smallest unit of testing. It includes all settings for which the test needs to be run: conditions, test requests, and IP versions.
  
- **Property version**. Property version refers to a Property Manager property version. 

- **Test results**. A test result for functional testing is a comparison of the expected value with the actual value. It can be either *Passed* or *Failed*. *Passed* means that the *Expected* result of the test was the same as the *Actual* result. *Failed* means that the *Expected* result of the test was different from the *Actual* result. To learn more about Functional testing results, see [Test results concepts](https://techdocs.akamai.com/test-ctr/docs/results-concepts) and [How to read test run results](https://techdocs.akamai.com/test-ctr/docs/view-results#functional-testing).

# CLI workflows

There are three ways you can use this CLI.

## Run a test
You can run a test for a property version, test suite, or test case using a [CLI command](#run-a-test) or [JSON input](#run-a-test-using-json). 

## Manage Test Center objects using JSON 
Here are the commands supporting JSON input:
  - [Generate a default test suite for a property](#generate-a-default-test-suite-for-a-property). Generates a default test suite for a property version for you to [import](#import-a-test-suite).
  - [Import a test suite](#import-a-test-suite). Imports to Test Center a test suite with test cases, variables, and property version association. 
  - [Get a test suite's details](#get-a-test-suites-details). Fetches, or exports, test suite's details that you can save and import on a different account or [edit](#edit-a-test-suite-using-json). You can also use this operation to clone test suites within your account.
  - [Edit a test suite using JSON](#edit-a-test-suite-using-json). 
  - [Run a test using JSON](#run-a-test-using-json).

## Manage Test Center objects using granular commands
Here are the commands you can use to manage Test Center objects from the CLI:

  - [List test suites](#list-test-suites).
  - [Create a test suite](#create-a-test-suite).
  - [Add a test case to a test suite](#add-a-test-case-to-a-test-suite).
  - [Edit a test suite](#edit-a-test-suite).
  - [Remove a test suite](#remove-a-test-suite).
  - [Remove a test case from a test suite](#remove-a-test-case-from-a-test-suite).
  - [Restore a test suite](#restore-a-test-suite).
  - [Get a test suite's details](#get-a-test-suites-details).
  - [List supported conditions](#list-supported-conditions).
  - [Run a test](#run-a-test).

# Available operations and commands

## Help
The `help` command returns an overview of available commands and flags.

## List test suites
The `test-suite list` command lists all test suites. You can filter the results for test suites created by a user, a property the test suite is associated with, or a string from the test suite's name or description.
The list also includes also the recently deleted test suites that you can restore.


**Command**: `test-suite list [--property 'PROPERTY NAME'] [--propver 'PROPERTY VERSION'] [-user USERNAME] [--search STRING]`, where:

- the `--property` flag filters the results for specific `PROPERTY NAME` which the test suite is associated with. The `--propver` flag filters the results further for specific `PROPERTY VERSION`.
- the `--user` flag filters the results for a `USERNAME` who created, edited, or deleted the test suite.
- the `--search` flag filters the results for a specific `STRING` in the test suite's name or description.

You can combine multiple flags to narrow the list.

**Examples**
- `akamai test-center test-suite list`
- `akamai test-center test-suite list --property 'example.com' --propver '4'`
- `akamai test-center test-suite list -u 'johndoe' -s 'regression'`

**Expected output**: The response lists all test suites matching the requested filters. For more details, you can check the [API response](ref:https://techdocs.akamai.com/test-ctr/v3/reference/get-test-suites) description.

## Create a test suite
The `test-suite add` command creates a new test suite.

**Command**: `test-suite add --name NAME [--description DESCRIPTION] [--unlocked] [--stateful]  [--property 'PROPERTY NAME' --propver 'PROPERTY VERSION']`, where:

- `NAME` is the name of the test suite. 
- `DESCRIPTION` is the description for the test suite. The `--description` flag is optional.
- the `--unlocked` flag unlocks the test suite. This flag is optional. By default all test suites are [locked](#concepts).
- the `--stateful` flag makes the test suite stateful. This flag is optional. By default all test suites are [stateless](#concepts). 
- the `--property` and `--propver` flags associate the test suite with a specific property version. `PROPERTY NAME` is the name of the property in Property Manager and `PROPERTY VERSION`, its appropriate version. You can use [Property Manager CLI](https://github.com/akamai/cli-property-manager) and the `list-properties|lpr` operation to get these values. These flags are optional, but they need to be used together. 

**Examples**:
- `akamai test-center test-suite add --name 'new test suite'`
- `akamai test-center test-suite add --name 'new test suite' --description 'TS for example.com' --unlocked --stateful --property 'example.com' --propver '4'`

**Expected output**: The response includes details about the created test suite. It also includes the ID of the test suite that you can use in other operations — for example, to [add a test case to a test suite](#add-a-test-case-to-a-test-suite). For more details, you can check the [API response](ref:https://techdocs.akamai.com/test-ctr/v3/reference/post-test-suites) description.

## Import a test suite
The `test-suite import` command imports a test suite from a JSON file or standard JSON input. You can use the [API documentation](ref:https://techdocs.akamai.com/test-ctr/v3/reference/post-test-suites-with-child-objects) to create the JSON file. Add your values to `BODY PARAMS` fields, copy the body of your request from the CURL code sample, and save it as a JSON file. 

**Command**: 
- To import a specific file from your computer: `test-center test-suite import < {FILE_PATH}/FILE_NAME.json`, where `FILE_PATH` and `FILE_NAME` are respectively location and name of the file to import.
- To import an outputted string: `echo '{"testSuite":{"testSuiteName":"TEST_SUITE_NAME","testSuiteDescription":"TEST_SUITE_DESCRIPTION","locked":true | false,"stateful":true | false,"variables":[{"variableName":"VARIABLE_NAME","variableValue":"VARIABLE_VALUE"}],"testCases":[]}}' | akamai test-center test-suite import`, where `TEST_SUITE_NAME`,`TEST_SUITE_DESCRIPTION`, `VARIABLE_NAME`, and `VARIABLE_VALUE` are your values for the test suite.

**Examples**:
- `akamai test-center test-suite import < ./users/johndoe/documents/test_suite_prop19.json`
- `echo '{"testSuite":{"testSuiteName":"test_suite_prop19","testSuiteDescription":"test suite for property version 19","locked":true,"stateful":false,"variables":[{"variableName":"host","variableValue":"www.akamai.com"}],"testCases":[]}}' | akamai test-center test-suite import`

**Expected output**: The response includes details about the imported test suite. It also includes the ID of the test suite that you can use in other operations— for example, to [edit the test suite using JSON](#edit-a-test-suite-using-json). For more details, you can check the [API response](ref:https://techdocs.akamai.com/test-ctr/v3/reference/post-test-suites-with-child-objects) description.

## Add a test case to a test suite
The `test-suite add-test-case` command adds a functional test case to a specific test suite.

**Command**: `test-suite add-test-case [--test-suite-id ID | --test-suite-name NAME] --url URL --condition CONDITION_STATEMENT [--ip-version v4|v6] [-a header ...] [-m header ...] [-f header ...]`, where: 

- `ID` or `NAME` specify the test suite you want to add the test case to. To get these values, run the [List test suites](#list-test-suites) operation. You need to provide either of these flags: `--test-suite-id` or `--test-suite-name`.
- `URL` is the fully qualified URL of the resource to test. It needs to contain the protocol, hostname, path, and any applicable string parameters. For example *https://www.example.com*.
- `CONDITION_STATEMENT` is one of the condition statements from the list of supported conditions with entered required values. To get the available condition statements, run the [List supported conditions](#list-supported-conditions) operation. Make sure to substitute default values in `" "` with your own.
- the `--ip-version` flag specifies the IP version to execute the test suite over, either `v4` or `v6`. It's set to `v4` by default. This flag is optional.
- the `--add-header` and `--modify-header` flags specify the request headers to respectively add or modify by the request. Headers should follow the format `name: value`. These flags are optional. You can also use these flags to provide Pragma headers. See [Pragma headers](https://techdocs.akamai.com/edge-diagnostics/docs/pragma-headers) for the list of supported values.
- the `-f` flag filters the header from the request. Provide only the `name` of the header. This flag is optional. 

**Examples**:
- `akamai test-center test-suite add-test-case --test-suite-id 1001 --url 'https://example.com/' --condition 'Response code is one of "200,201"'`
- `akamai test-center test-suite add-test-case --test-suite-name 'Example TS' --u 'https://example.com/' -c 'Response code is one of "200"' -a 'Accept: text/html' -a 'X-Custom: 123' -m 'User-Agent: Mozilla' -f 'Accept-Language'`

**Expected output**: The response includes details of the test case created for the test suite. For more details, you can check the [API response](ref:https://techdocs.akamai.com/test-ctr/v3/reference/post-test-cases) description.

## Generate a default test suite for a property
The `test-suite generate-default` command generates a default test suite with test cases for a specific property. Based on property settings and its behaviors and the `--url` flag value, Test Center generates a test suite object with test cases and variables for you to modify and add to Test Center using the [Import a test suite](#import-a-test-suite) operation.

**Command**: `test-suite generate-default --property 'PROPERTY NAME' --propver 'PROPERTY VERSION' --url URL ...`, where: 

- `PROPERTY NAME` and `PROPERTY VERSION` specify the property in Property Manager you want to generate a test suite for. To get these values you can use the [Property Manager CLI](https://github.com/akamai/cli-property-manager) and the `list-properties|lpr` operation. 
- `URL` is the fully qualified URL of the property hostname. The `--URL` flag can be used multiple times.


**Examples**:
- `akamai test-center test-suite generate-default --property 'example.com' --propver '4' --url "https://www.example.com/" -u "https://www.example.com/index/"`

**Expected output**: The response includes details about the generated response and included test cases. For more details, you can check the [API response](ref:https://techdocs.akamai.com/test-ctr/v3/reference/auto-generate-test-suite) description.

## Edit a test suite
The `test-suite edit` command edits basic data of a specific test suite. Provide only data you want to edit in the original test suite.

**Command**: `test-suite edit --id ID [--name NAME] [--description DESCRIPTION] [--unlocked] [--stateful] [--property 'PROPERTY NAME' --propver 'PROPERTY VERSION' | --remove-property ]`, where:

- `ID` is the identifier of the test suite you want to edit. To get this value, run the [List test suites](#list-test-suites) operation. You need to provide either of these flags: `--id` or `--name`.
- `NAME` is the test suite's name. To get this value, run the [List test suites](#list-test-suites) operation. You need to provide either of these flags: `--name` or `--id`.
- `DESCRIPTION` is the new description for the test suite. 
- `--unlocked` and `--stateful` flags update the status of the test suite. 
- `--property` and `--propver` update the test suite's association to the property with a specific `PROPERTY NAME` and `PROPERTY VERSION`. If applicable, provide either these flags or `--remove-property`.
- the `--remove-property` removes the association to a property version. If applicable, provide either this flag or `--property` and `--propver`.

**Examples**:
- `akamai test-center test-suite edit --id 1001 --name 'Updated test suite'`
- `akamai test-center test-suite edit --id 1001 --name 'Updated test suite' --description 'Test suite for example.com' --property 'example.com' --propver '4' --unlocked`
- `akamai test-center test-suite edit --id 1001 --stateful --remove-property`
  
**Expected output**: The response returns the updated test suite. For more details, you can check the [API response](ref:https://techdocs.akamai.com/test-ctr/v3/reference/put-test-suite) description.

## Edit a test suite using JSON
The `test-suite manage` command uses the JSON input to edit a specific test suite. Provide the whole test suite object, together with test cases and variables, to include in Test Center. Only data provided in the latest JSON input will be saved. You can use the [API documentation](ref:https://techdocs.akamai.com/test-ctr/v3/reference/put-test-suites-with-child-objects) to create the JSON file. Add your values to `BODY PARAMS` fields, copy the body of your request from the CURL code sample, and save it as a JSON file. 

**Command**:
- To edit the test suite with a specific file from your computer: `test-center test-suite manage < {FILE_PATH}/FILE_NAME.json`, where `FILE_PATH` and `FILE_NAME` are respectively location and name of the file to import.
- To edit the test suite with outputted string: `echo '{"testSuite":{"testSuiteId":ID,"testSuiteName":"TEST_SUITE_NAME","testSuiteDescription":"TEST_SUITE_DESCRIPTION","locked":true | false,"stateful":true | false,"variables":[{"variableName":"VARIABLE_NAME","variableValue":"VARIABLE_VALUE"}],"testCases":[]}}' | akamai test-center test-suite manage`, where:
  - `ID` is the unique identifier of the test suite you want to edit.
  - `TEST_SUITE_NAME`,`TEST_SUITE_DESCRIPTION`, `VARIABLE_NAME`, and `VARIABLE_VALUE` are your values for the edited test suite.

**Examples**:
- `akamai test-center test-suite manage < ./users/johndoe/documents/test_suite_prop19.json`
- `echo '{"testSuite":{"testSuiteName":"test_suite_prop19","testSuiteDescription":"test suite for property version 19","locked":true,"stateful":false,"variables":[{"variableName":"host","variableValue":"www.akamai.com"}],"testCases":[]}}' | akamai test-center test-suite manage`

**Expected output**: The response returns the updated test suite. For more details, you can check the [API response](ref:https://techdocs.akamai.com/test-ctr/v3/reference/put-test-suites-with-child-objects) description.

## Remove a test suite
The `test-suite remove` command removes a specific test suite from Test Center. Test suites can be [restored](#restore-a-test-suite) for 30 days since their removal. 

**Command**: `test-suite remove [--id ID | --name NAME]`, where `ID` and `NAME` specify the test suite you want to remove. To get these values, run the [List test suites](#list-test-suites) operation. You need to provide either of these flags: `--id` or `--name`.

**Example**: 
- `akamai test-center test-suite remove --name "Test suite name"`
- `akamai test-center test-suite remove --id 12345`

**Expected output**: Successful operation confirmed.

## Remove a test case from a test suite
The `test suite remove-test-case` command removes a test case with a specific order number from a test suite. Removed test cases can't be restored.

**Command**: `test-suite remove-test-case --test-suite-id ID --order-num ORDER_NUMBER`, where:

- `ID` is the unique identifier of the test suite you want to remove the test case from. To get this values, run the [List test suites](#list-test-suites) operation.
- `ORDER_NUMBER` is the order number of the test case you want to remove. To get this value, run the [Get a test suite's details](#get-a-test-suites-details) operation.

**Examples**: `akamai test-center test-suite remove-test-case --test-suite-id 1001 --order-num 6`

**Expected output**: Successful operation confirmed.

## Restore a test suite
The `test-suite restore` command restores a removed test suites and included test cases. Test suites can be restored for 30 days since their removal.

**Command**: `test-suite restore [--id ID | --name NAME]`, where `ID` and `NAME` specify the test suite you want to restore. To get these values, run the [List test suites](#list-test-suites) operation. You need to provide either of these flags: `--id` or `--name`.

**Example**:
- `akamai test-center test-suite restore --name "test_suite_prop19"`
- `akamai test-center test-suite restore --id 12345`

**Expected output**: Successful operation confirmed.

## Get a test suite's details
The `test-suites view` command exports details of a specific test suite, including its test cases. You can group the included test cases by a test request, condition, or IP version.

**Command**: `test-suite view [--id ID | --name NAME] [--group-by test-request | condition | ipversion]`, where:

- `ID` and `NAME` specify the test suite you want to get the details of. To get these values, run the [List test suites](#list-test-suites) operation. You need to provide either of these flags: `--id` or `--name`
- the `--group-by` flag specifies grouping of the included test cases, by `test-request`, `condition`, or `ipversion`. This flag is optional.

**Examples**:
- `akamai test-center test-suite view --id 1001`
- `akamai test-center test-suite view --name 'test_suite_prop19' --group-by test-request`

**Expected output**: The response includes details about the test suite. You can save the returned object and import it on a different account or [edit](#edit-a-test-suite-using-json) it. You can also use this operation to clone test suites within your account. For more details, you can check the [API response](ref:https://techdocs.akamai.com/test-ctr/v3/reference/get-test-suites-with-child-objects) description.

## List supported conditions
The `conditions` command returns the list of all conditions you can use when creating a test case for a test suite. Note that the statements contain default values in `" "`.  You need to replace them with your own values before creating the test case.

**Command**: `conditions`

**Expected output**: List of supported condition statements. For more details, you can check the [API response](ref:https://techdocs.akamai.com/test-ctr/v3/reference/get-test-catalog-template) description.

## Run a test 
The `test` command runs a test for a specific test suite, single test case, or a property version. 

**Command**: 
- To run a test for a test suite: `test [--test-suite-id ID] | [--test-suite-name 'NAME'] -e staging|production`
- To run a test for a property version: `test [--property 'PROPERTY NAME' --propver 'PROPERTY VERSION']  -e staging|production`
- To run a test for a single test case: `test [-u URL -c CONDITION STATEMENT -i v4|v6 [--add-header 'name: value' ...] [--modify-header 'name: value' ...] [--filter-header name ...]] -e staging|production`, where: 

  - `ID` and `NAME` specify the test suite you want to run the test for. To get these values, run the [List test suites](#list-test-suites) operation. You need to provide either of these flags: `--id` or `--name`
  - `PROPERTY NAME` and `PROPERTY VERSION` specify the property in Property Manager you want to run the test for. To get these values you can use the [Property Manager CLI](https://github.com/akamai/cli-property-manager) and the `list-properties|lpr` operation. 
  - `URL` is the fully qualified URL of the resource to test. It needs to contain the protocol, hostname, path, and any applicable string parameters — for example, *https://www.example.com*.
  - `CONDITION STATEMENT` is one of the condition statements from the list of supported conditions with entered required values. To get the list of supported conditions, run the [List supported conditions](#list-supported-conditions) operation. Make sure to replace default values in `" "` with your own.
  - the `-i` flag specifies the IP version to execute the test case over, either `v4` or `v6`. This flag is optional, set to `v4` by default.
  - the `--add-header` and `--modify-header` flags specify the request headers to respectively added or modify by the request. Headers should follow the format `name: value`. These flags are optional and accept multiple values. You can also use these flags to provide Pragma headers. See [Pragma headers](https://techdocs.akamai.com/edge-diagnostics/docs/pragma-headers) for the list of supported values.
  - the `--filter-header` flag filters the header from the request. Provide only the `name` of the header. This flag is optional. 
  - the `-e` flag specifies the environment you want to run the test on, either `staging` or `production`.

**Expected output**: Once you submit the test run, it may take few minutes for Test Center to execute the test. To learn more about returned test results, check [How to read test run results - Functional testing](https://techdocs.akamai.com/test-ctr/docs/view-results#functional-testing). For more details, you can check the [API response](ref:https://techdocs.akamai.com/test-ctr/v3/reference/post-test-runs) description.

## Run a test using JSON
The `test` command runs a test for a specific test suite, single test case, or a property version using a JSON file or JSON input. You can use the [API documentation](ref:https://techdocs.akamai.com/test-ctr/v3/reference/post-test-runs) to create the JSON file. Add your values to `BODY PARAMS` fields, copy the body of your request from the CURL code sample, and save it as a JSON file. 

**Command**: 
To import a specific file from your computer: `test-center test < {FILE_PATH}/FILE_NAME.json`, where `FILE_PATH` and `FILE_NAME` are respectively location and name of the file to import.

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
   "note":"Test run for a property version",
   "sendEmailOnCompletion":true
}
```

JSON example to run a test for a test suite:
```
{
   "comparative":{
      "testDefinitionExecutions":[
         {
            "ipVersions":[
               
            ]
         }
      ]
   },
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
   "note":"Test run for a test suite",
   "sendEmailOnCompletion":true
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
            "ipVersion":"ipv4"
         }
      }
   },
   "note":"Test run for a test case",
   "sendEmailOnCompletion":true,
   "targetEnvironment":"staging",
   "purgeOnStaging":true
}
```

**Expected output**: Once you submit the test run, it may take a few minutes for Test Center to execute the test. To learn more about returned test results, check [How to read test run results - Functional testing](https://techdocs.akamai.com/test-ctr/docs/view-results#functional-testing). For more details, you can check the [API response](ref:https://techdocs.akamai.com/test-ctr/v3/reference/post-test-runs) description.

# Available flags

You can use the following flags with all the listed commands.

## edgerc 
The `--edgerc` flag changes the default path to the .edgerc file. This file contains the API credentials required to run all commands.
Without this flag, the user's home directory is used by default.

**Command**: `$akamai test-center --edgerc EDGERC_PATH [command]`

**Example**: `$akamai test-center --edgerc C:/users/johndoe/.edgerc test-suite list`

## section
The `--section` flag changes the default section name. The section name specifies which section of API credentials to read from the .edgerc file. Without this flag, the default `test-center` section name is used by default.

**Command**: `$akamai test-center --section SECTION_NAME [command]`

**Examples**: 
- `$akamai test-center --section default test-suite list`
- `$set AKAMAI_EDGERC_SECTION=default`

## account-key 
The `--account-key` flag changes your account. When testing your configuration, you may need to switch between different accounts. To do this, run the required operation with the `--account-key` flag followed by the account ID of your choice. 

**Command**: `$akamai test-center ----account-key ACCOUNT KEY [command]`

**Example**: `akamai test-center --account-key 1-1TJZFB test-suite list`

##force-color
The `--force-color` flag forces color to non-TTY output.

## help
The `--help` flag returns help for a command. 

## version
The `--version` flag returns the version.

## json
The `--json` flag returns the information in JSON format. 


# Windows 10 2018 version
If you're using Windows 10, 2018 version and you're having problems running the Test Center 
CLI, we recommend you try the following workaround. In the downloaded repository, add the `.exe` 
suffix to the `akamai-test-center` executable file.

# Notice

Copyright © 2022 Akamai Technologies, Inc.

Your use of Akamai's products and services is subject to the terms and provisions outlined in [Akamai's legal policies](https://www.akamai.com/us/en/privacy-policies/).