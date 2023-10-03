package cmd

import (
	"github.com/akamai/cli-test-center/internal/api"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/service"
	"github.com/akamai/cli-test-center/internal/util"
	"github.com/akamai/cli-test-center/internal/validator"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/spf13/cobra"
)

var (
	url                string
	addHeader          []string
	modifyHeader       []string
	filterHeader       []string
	condition          string
	ipVersion          string
	testSuiteName      string
	testSuiteIdStr     string
	propertyIdStr      string
	propertyNameStr    string
	propertyVersionStr string
	targetEnvironment  string

	// new flag variables
	client            string
	location          string
	requestMethod     string
	encodeRequestBody bool
	requestBody       string
)

var testRunRequest model.TestRun

var runCmd = &cobra.Command{
	Use:     externalconstant.TestRunUse,
	Example: externalconstant.TestRunExample,
	Run: func(cmd *cobra.Command, args []string) {
		eghc := api.NewEdgeGridHttpClient(config, accountSwitchKey)
		api := api.NewApiClient(*eghc)
		testRunService := service.NewService(*api, cmd, jsonOutput)
		globalValidator := validator.NewGlobalValidator(cmd, jsonData)
		// validate subcommand
		globalValidator.ValidateSubCommandsNotAllowed(cmd, args, true)

		validator := validator.NewValidator(cmd, jsonData)

		// validate different flags and combination to run test
		runTestUsing := validator.ValidateTestRunFlagsAndGetRunEnum(testSuiteIdStr, testSuiteName, propertyIdStr, propertyNameStr, propertyVersionStr,
			url, condition, ipVersion, targetEnvironment, client, location, requestMethod, requestBody, addHeader, modifyHeader,
			jsonData, &testRunRequest, isStandardInputAvailable, encodeRequestBody)

		//Run test
		testRunService.RunTest(runTestUsing, testSuiteIdStr, testSuiteName, propertyIdStr, propertyNameStr, propertyVersionStr,
			url, condition, ipVersion, targetEnvironment, client, location, requestMethod, requestBody, addHeader, modifyHeader,
			filterHeader, testRunRequest, encodeRequestBody)
	},
}

func init() {
	testCmd.AddCommand(runCmd)
	runCmd.Flags().SortFlags = false

	runCmd.Short = util.GetMessageForKey(runCmd, internalconstant.Short)
	runCmd.Long = util.GetMessageForKey(runCmd, internalconstant.Long)

	runCmd.Flags().StringVarP(&url, externalconstant.FlagUrl, externalconstant.FlagUrlShortHand, internalconstant.Empty, util.GetMessageForKey(runCmd, externalconstant.FlagUrl))
	runCmd.Flags().StringVarP(&testSuiteIdStr, externalconstant.FlagTestSuiteId, externalconstant.FlagTestSuiteIdShortHand, internalconstant.Empty, util.GetMessageForKey(runCmd, externalconstant.FlagTestSuiteId))
	runCmd.Flags().StringVarP(&testSuiteName, externalconstant.FlagTestSuiteName, externalconstant.FlagTestSuiteNameShortHand, internalconstant.Empty, util.GetMessageForKey(runCmd, externalconstant.FlagTestSuiteName))
	runCmd.Flags().StringVarP(&propertyIdStr, externalconstant.FlagPropertyId, internalconstant.Empty, internalconstant.Empty, util.GetMessageForKey(runCmd, externalconstant.FlagPropertyId))
	runCmd.Flags().StringVarP(&propertyNameStr, externalconstant.FlagPropertyName, externalconstant.FlagPropertyShortHand, internalconstant.Empty, util.GetMessageForKey(runCmd, externalconstant.FlagPropertyName))
	runCmd.Flags().StringVarP(&propertyVersionStr, externalconstant.FlagPropertyVersion, externalconstant.FlagPropertyVersionShortHand, internalconstant.Empty, util.GetMessageForKey(runCmd, externalconstant.FlagPropertyVersion))
	runCmd.Flags().StringVar(&ipVersion, externalconstant.FlagIpVersion, internalconstant.IpVersionDefaultValue, util.GetMessageForKey(runCmd, externalconstant.FlagIpVersion))
	runCmd.Flags().StringArrayVarP(&addHeader, externalconstant.FlagAddHeader, externalconstant.FlagAddHeaderShortHand, []string{}, util.GetMessageForKey(runCmd, externalconstant.FlagAddHeader))
	runCmd.Flags().StringArrayVarP(&modifyHeader, externalconstant.FlagModifyHeader, externalconstant.FlagModifyHeaderShortHand, []string{}, util.GetMessageForKey(runCmd, externalconstant.FlagModifyHeader))
	runCmd.Flags().StringArrayVarP(&filterHeader, externalconstant.FlagFilterHeader, externalconstant.FlagFilterHeaderShortHand, []string{}, util.GetMessageForKey(runCmd, externalconstant.FlagFilterHeader))
	runCmd.Flags().StringVarP(&condition, externalconstant.FlagCondition, externalconstant.FlagConditionShortHand, internalconstant.Empty, util.GetMessageForKey(runCmd, externalconstant.FlagCondition))
	runCmd.Flags().StringVar(&targetEnvironment, externalconstant.FlagEnv, internalconstant.Staging, util.GetMessageForKey(runCmd, externalconstant.FlagEnv))

	// new Flags
	runCmd.Flags().StringVarP(&client, externalconstant.FlagClient, externalconstant.FlagClientShortHand, internalconstant.Curl, util.GetMessageForKey(runCmd, externalconstant.FlagClient))

	// we are disabling this flag for now as we support only one location, we may enable it in future.
	//runCmd.Flags().StringVarP(&location, externalconstant.FlagLocation, externalconstant.FlagLocationShortHand, internalconstant.DefaultLocation, util.GetMessageForKey(runCmd, externalconstant.FlagLocation))

	runCmd.Flags().StringVarP(&requestMethod, externalconstant.FlagRequestMethod, externalconstant.FlagRequestMethodShortHand, internalconstant.GetRequestMethod, util.GetMessageForKey(runCmd, externalconstant.FlagRequestMethod))
	runCmd.Flags().BoolVarP(&encodeRequestBody, externalconstant.FlagEncodeRequestBody, externalconstant.FlagEncodeRequestBodyShortHand, false, util.GetMessageForKey(runCmd, externalconstant.FlagEncodeRequestBody))
	runCmd.Flags().StringVarP(&requestBody, externalconstant.FlagRequestBody, externalconstant.FlagRequestBodyShortHand, internalconstant.Empty, util.GetMessageForKey(runCmd, externalconstant.FlagRequestBody))
}
