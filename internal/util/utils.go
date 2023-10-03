package util

import (
	"bufio"
	"fmt"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/clarketm/json"
	"github.com/fatih/camelcase"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

func CamelToTitle(inp string) string {
	inp = strings.Title(inp)
	arr := camelcase.Split(inp)
	out := strings.Join(arr, " ")
	out = strings.Replace(out, "Url", "URL", -1)
	return out
}

type Spinner struct {
	_spinner *spinner.Spinner
	enabled  bool
}

func NewSpinner(message string, enabled bool) Spinner {
	s := Spinner{enabled: enabled}
	s._spinner = spinner.New([]string{".     ", "..    ", "...   ", "....  ", "..... ", "......"}, 300*time.Millisecond)
	s._spinner.Writer = os.Stderr // Out-of-band info like progress info should go to stderr
	s._spinner.Prefix = message
	return s
}

func (s Spinner) Start() Spinner {
	if s.enabled {
		s._spinner.Start()
	}
	return s
}

func (s Spinner) StopWithSuccess() {
	if s.enabled {
		stopSpinner(s._spinner, color.GreenString("[OK]"))
	}
}

func (s Spinner) StopWithFailure() {
	if s.enabled {
		stopSpinner(s._spinner, color.RedString("[FAIL]"))
	}
}

func stopSpinner(s *spinner.Spinner, terminalMessage string) {
	s.FinalMSG = s.Prefix + "...... " + terminalMessage
	s.Stop()
	_, _ = fmt.Fprintln(s.Writer)
}

// Check contains and irrespective of case.
func ContainsIgnoreCase(a string, b string) bool {
	return strings.Contains(strings.ToLower(a), strings.ToLower(b))
}

// Check contains and irrespective of case.
func ContainsInArray(array []string, inputString string) bool {
	var result = false
	for _, x := range array {
		if x == strings.ToLower(inputString) {
			result = true
			break
		}
	}

	return result
}

func ClientProfileInCLIOutputFormat(clientProfile model.ClientProfile) string {

	if clientProfile.IpVersion == internalconstant.Ipv6 {
		return clientProfile.Client + externalconstant.PlusSign + "IPv6"
	} else {
		return clientProfile.Client + externalconstant.PlusSign + "IPv4"
	}
}

func RequestHeaderInCLIOutputFormat(headerName, headerAction, headerValue string) string {

	switch strings.ToUpper(headerAction) {
	case internalconstant.Add:
		return fmt.Sprintf("(%s): %s: %s", "Added", headerName, headerValue)
	case internalconstant.Modify:
		return fmt.Sprintf("(%s): %s: %s", "Modified", headerName, headerValue)
	case internalconstant.Filter:
		return fmt.Sprintf("(%s): %s: %s", "Filtered", headerName, "N/A")
	}
	return ""
}

func ConvertBooleanToYesOrNo(input bool) string {
	if input {
		return "Yes"
	}
	return "No"
}

func FormatTime(inputTime string) string {
	layout := "2006-01-02T15:04:05+0000"
	myDate, err := time.Parse(layout, inputTime)
	if err != nil {
		fmt.Println(err)
	}
	// convert this date to desired format when decided.
	return myDate.Format("01/02/2006, 15:04 PM -07:00")
}

// Get all placeholders in string inside {{}}
func GetPlaceHoldersInString(errorMessage, regex string) []string {

	r := regexp.MustCompile(regex)
	matches := r.FindAllStringSubmatch(errorMessage, -1)
	var placeHolders = make([]string, len(matches))
	for i, v := range matches {
		placeHolders[i] = v[1]
	}

	return placeHolders
}

func ReadStdin(cmd *cobra.Command) (bool, []byte) {
	file := os.Stdin
	fi, err := file.Stat()
	if err != nil {
		AbortWithExitCode(fmt.Sprintf(GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, "", "standardInputErrorMsg"), err), internalconstant.ExitStatusCode1)
	}
	size := fi.Size()
	isStandardInputAvailable := false

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		isStandardInputAvailable = true
	}

	if size > 0 {
		log.Debug("%v bytes available in Stdin\n", size)
		jsonData, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
		if err != nil {
			AbortWithExitCode(fmt.Sprintf(GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, "", "standardInputErrorMsg"), err), internalconstant.ExitStatusCode1)
		}

		return isStandardInputAvailable, jsonData
	}
	log.Debug("Stdin is empty")
	return isStandardInputAvailable, nil
}

func ByteArrayToStruct(cmd *cobra.Command, byt []byte, payloadObject interface{}) {
	if err := json.Unmarshal(byt, payloadObject); err != nil {
		log.Debug(err)
		AbortWithUsageAndMessageAndCode(cmd, GetErrorMessageForFlag(cmd, internalconstant.Invalid, internalconstant.Json), internalconstant.ExitStatusCode3)
	}
}

// CheckIfBothJsonAndFlagAreSetForCommand method returns true if json input is set, otherwise false
func CheckIfBothJsonAndFlagAreSetForCommand(cmd *cobra.Command, jsonData []byte, isStandardInputAvailable bool) bool {
	checkIfSubCommandFlagsAreNonEmpty := checkIfAnySubCommandFlagIsSet(cmd)
	if checkIfSubCommandFlagsAreNonEmpty && jsonData != nil {
		AbortWithUsageAndMessageAndCode(cmd, GetGlobalErrorMessage("invalidCommandInput"), internalconstant.ExitStatusCode2)
	} else if !checkIfSubCommandFlagsAreNonEmpty && !isStandardInputAvailable {
		// If command is provided without JSON input and flags, we simply through flag missing error and abort with usage
		AbortWithUsageAndMessageAndCode(cmd, GetErrorMessageForFlag(cmd, internalconstant.Missing, "flagOrJsonImport"), internalconstant.ExitStatusCode2)
	} else if !checkIfSubCommandFlagsAreNonEmpty && jsonData == nil {
		AbortWithUsageAndMessageAndCode(cmd, GetGlobalErrorMessage("invalidJsonInput"), internalconstant.ExitStatusCode3)
	}

	return !checkIfSubCommandFlagsAreNonEmpty
}

/*
	Returns true/false if any flag has been set/modified for the child level sub-command. e.g. akamai test-center ts

generate-default, here flags for generate-default will be checked and global flags will be not be included in check.
*/
func checkIfAnySubCommandFlagIsSet(cmd *cobra.Command) bool {
	isSubCommandFlagSet := false

	// set boolean to true if any non-inherited flag is passed as a command line flag
	cmd.NonInheritedFlags().VisitAll(func(flag *pflag.Flag) {
		if cmd.Flags().Changed(flag.Name) {
			isSubCommandFlagSet = true
		}
	})

	return isSubCommandFlagSet
}

// LegacyArgs is used to invalidate unknown subcommands took reference from cobra.args legacyArgs() library
func LegacyArgs(cmd *cobra.Command, args []string) error {
	// we will not show error when no subcommand is given for parent command, but we will show help section.
	if len(args) <= 0 {
		return nil
	}

	// no subcommand, always take args
	if cmd.HasSubCommands() {
		return fmt.Errorf(GetErrorMessageForSubArgument(cmd, internalconstant.Invalid, internalconstant.SubCommandWrongArgumentPassed), args[0], cmd.Name(), FindSuggestions(cmd, args[0]))
	}

	// root command with subcommands, do subcommand checking.
	if !cmd.HasParent() && len(args) > 0 {
		return fmt.Errorf(GetErrorMessageForSubArgument(cmd, internalconstant.Invalid, internalconstant.SubCommandWrongArgumentPassed), args[0], cmd.Name(), FindSuggestions(cmd, args[0]))
	}
	return nil
}

// FindSuggestions returns a list possible subcommands referenced cobra.command findSuggestions() library
func FindSuggestions(cmd *cobra.Command, arg string) string {
	if cmd.DisableSuggestions {
		return internalconstant.Empty
	}
	if cmd.SuggestionsMinimumDistance <= 0 {
		cmd.SuggestionsMinimumDistance = 2
	}
	suggestionsString := internalconstant.Empty
	if suggestions := cmd.SuggestionsFor(arg); len(suggestions) > 0 {
		suggestionsString += "\n\nDid you mean this?\n"
		for _, s := range suggestions {
			suggestionsString += fmt.Sprintf("\t%v\n", s)
		}
	}
	return suggestionsString
}

// NoArgsCheck returns an error if any args are included.
func NoArgsCheck(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf(GetErrorMessageForSubArgument(cmd, internalconstant.Invalid, internalconstant.SubCommandWithArgumentNotAllowed), args[0], cmd.Name())
	}
	return nil
}

// GetHttpExitCode get equivalent exit status code for corresponding http status code
func GetHttpExitCode(responseCode string) int {
	log.Debugf("Response Status Code [%q] ", responseCode)
	statusCode, err := strconv.Atoi(responseCode)
	if err != nil {
		statusCode = internalconstant.ExitStatusCode1
	}
	if statusCode >= 400 && statusCode <= 550 { // Difference used for 4xx and 5xx errors
		statusCode = statusCode - internalconstant.BaseSubtractor
	} else if statusCode >= 100 && statusCode <= 399 { // 1xx 2xx & 3xx errors are treated as success
		statusCode = internalconstant.ExitStatusCode0
	} else {
		statusCode = internalconstant.ExitStatusCode1
	}
	log.Debugf("Exit Status Code [%v] ", statusCode)
	return statusCode
}

func GetStatusKeyAndColour(status string) (string, color.Color) {
	if status == internalconstant.CompletedEnum {
		c := color.New(color.FgGreen)
		return internalconstant.CompletedKey, *c

	} else if status == internalconstant.CompletedWithUnexpectedResultsEnum {
		c := color.New(color.FgYellow)
		return internalconstant.CompletedWithUnexpectedResultsKey, *c

	} else if status == internalconstant.InProgressEnum {
		return internalconstant.InProgressKey, *color.New()

	} else {
		c := color.New(color.FgRed)
		return internalconstant.FailedKey, *c
	}
}

func GetColourForEnum(enum string, bold bool) color.Color {
	c := color.New()

	if bold {
		c.Add(color.Bold)
	}

	if enum == internalconstant.GetRequestMethod {
		c.Add(color.FgBlue)

	} else if enum == internalconstant.HeadRequestMethod {
		c.Add(color.FgYellow)

	} else if enum == internalconstant.PostRequestMethod {
		c.Add(color.FgGreen)
	}

	return *c
}
func GetResolvedOrUnResolvedHeaders(header model.RequestHeader) string {
	var headerName = header.HeaderName
	var headerAction = header.HeaderAction
	var headerValue = header.HeaderValue
	if header.HeaderNameResolved != internalconstant.Empty {
		headerName = header.HeaderNameResolved
	}
	if header.HeaderValueResolved != internalconstant.Empty {
		headerValue = header.HeaderValueResolved
	}
	return RequestHeaderInCLIOutputFormat(headerName, headerAction, headerValue)
}

func GetResolvedOrUnResolvedCondition(condition model.Condition) string {
	var cond = condition.ConditionExpression
	if condition.ConditionExpressionResolved != internalconstant.Empty {
		cond = condition.ConditionExpressionResolved
	}
	return cond
}

func GetResolvedOrUnResolvedRequestBody(testRequest model.TestRequest) string {
	var requestBody = testRequest.RequestBody
	if testRequest.RequestBodyResolved != internalconstant.Empty {
		requestBody = testRequest.RequestBodyResolved
	}
	return requestBody
}

func GetResolvedOrUnResolvedRequestURL(testRequest model.TestRequest) string {
	var url = testRequest.TestRequestUrl
	if testRequest.TestRequestUrlResolved != internalconstant.Empty {
		url = testRequest.TestRequestUrlResolved
	}
	return url
}

func SetVariablesInCLIOutputFormat(variable model.DynamicVariable) string {
	return variable.VariableName + externalconstant.Equals + variable.VariableValue
}

func GetConvertedInteger(number string) int {
	num, _ := strconv.Atoi(number)
	return num
}

func GetTotalTestCasesCount(testCases []model.TestCase) int {
	count := len(testCases)
	for _, testcase := range testCases {
		count += len(testcase.DerivedTestCases)
	}
	return count
}
