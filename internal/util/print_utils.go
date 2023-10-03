package util

// This class will have only akamai/global cli standard print functions

import (
	"fmt"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/clarketm/json"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// PrintError Start Standard Print Functions
func PrintError(message string, args ...interface{}) {
	c := color.New(color.FgRed)
	_, _ = c.Fprintf(os.Stderr, message, args...)
}

func PrintWarning(message string, args ...interface{}) {
	c := color.New(color.FgCyan)
	_, _ = c.Fprintf(os.Stderr, message, args...)
}

func PrintHeader(message string, args ...interface{}) {
	c := color.New(color.FgYellow).Add(color.Bold)
	_, _ = c.Printf(message, args...)
}

func PrintSuccess(message string, args ...interface{}) {
	c := color.New(color.FgGreen)
	_, _ = c.Printf(message, args...)
}

func PrintSuccessInBold(message string, args ...interface{}) {
	c := color.New(color.FgGreen).Add(color.Bold)
	_, _ = c.Printf(message, args...)
}

func Bold(a ...interface{}) string {
	c := color.New(color.Bold)
	return c.Sprint(a...)
}

func Italic(a ...interface{}) string {
	c := color.New(color.Italic)
	return c.Sprint(a...)
}

func PrintLabelAndValue(label string, value interface{}) {
	c := color.New(color.Bold)
	_, _ = c.Printf(label + externalconstant.Colon + internalconstant.Space)
	fmt.Printf("%v\n", value)
}

func PrintLabelValueWithColour(label string, clr *color.Color, value interface{}) {
	c := color.New(color.Bold)
	_, _ = c.Printf(label + externalconstant.Colon + internalconstant.Space)
	_, _ = clr.Printf("%v\n", value)
}

// print json output if flag is passed
func CheckAndPrintJson(jsonOutput bool, data interface{}) {
	// print json output if flag is passed
	if jsonOutput {
		PrintJsonAndExit(data)
	}
}

func PrintJsonAndExit(data interface{}) {
	// This uses a third-party replacement (github.com/clarketm/json) for Go's default JSON encoder.
	// It considers empty structs (as opposed to just nil pointer to struct) for omitempty.
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
		AbortWithExitCode(GetGlobalErrorMessage("jsonOutputFailed"), internalconstant.ExitStatusCode1)
	}

	if string(b) != "null" {
		fmt.Println(string(b) + "\n")
	} else {
		fmt.Println("[]\n")
	}
	os.Exit(internalconstant.ExitStatusCode0)
}

func AbortWithExitCode(message string, code int) {
	PrintError(message + "\n")
	os.Exit(code)
}

func AbortWithUsageAndMessageAndCode(cmd *cobra.Command, message string, code int) {
	if message != internalconstant.Empty {
		PrintError(message + "\n\n")
	}
	err := cmd.Usage()
	fmt.Println()
	if err != nil {
		return
	}
	os.Exit(code)
}
func AbortForCommand(cmd *cobra.Command, cliError *model.CliError) {
	AbortForCommandWithSubResource(cmd, cliError, internalconstant.Empty, internalconstant.Empty)
}

func AbortForCommandWithSubResource(cmd *cobra.Command, cliError *model.CliError, subResource, operation string) {

	responseCode := strconv.Itoa(cliError.ResponseCode)
	if len(responseCode) == 3 && strings.Contains("500,502,503,504,405", responseCode) {
		PrintError(GetGlobalErrorMessage(responseCode) + "\n\n")
	} else if cliError.ApiError != nil {
		PrintErrorMessages(GetApiErrorMessagesForCommand(cmd, *cliError.ApiError, subResource, operation, responseCode))
		println()
	} else if len(cliError.ApiSubErrors) != 0 {
		PrintErrorMessages(GetApiSubErrorMessagesForCommand(cmd, cliError.ApiSubErrors, "", subResource, operation))
		println()
	} else {
		PrintError(cliError.ErrorMessage + "\n\n")
	}
	// Get equivalent exit status code for corresponding http status code
	statusCode := GetHttpExitCode(responseCode)
	os.Exit(statusCode)
}

func PrintErrorMessages(errorMessages []string) {
	for _, message := range errorMessages {
		PrintError(message + "\n")
	}
}

func PrintWarnings(waringMessages []string) {
	for _, message := range waringMessages {
		PrintWarning(message + "\n")
	}
}

// ShowTable Standard function to show table in same format across Test Center CLI
func ShowTable(tableHeaders []string, tableContents [][]string, showTotal bool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)

	table.SetHeader(tableHeaders)
	table.AppendBulk(tableContents)
	table.SetAutoFormatHeaders(false)
	table.Render()

	if showTotal {
		PrintTotalItems(len(tableContents))
	}
}

func PrintTemplate(templateToParse string, data interface{}) {

	// standard functions for templates
	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"dec": func(i int) int {
			return i - 1
		},
		"printRequestURL":    GetResolvedOrUnResolvedRequestURL,
		"printHeader":        GetResolvedOrUnResolvedHeaders,
		"printRequestBody":   GetResolvedOrUnResolvedRequestBody,
		"printClientProfile": ClientProfileInCLIOutputFormat,
		"printCondition":     GetResolvedOrUnResolvedCondition,
		"printSetVariables":  SetVariablesInCLIOutputFormat,
		"bold":               Bold,
		"join": func(elements []string) string {
			return strings.Join(elements, externalconstant.Comma)
		},
	}

	tmp, err := template.New("template").Funcs(funcMap).Parse(templateToParse)

	// if there is error,
	if err != nil {
		log.Error(err)
		PrintError(externalconstant.CliErrorMessageTemplateOutputError + "\n")
		os.Exit(internalconstant.ExitStatusCode1)
	} else {
		// standard output to print merged data
		err = tmp.Execute(os.Stdout, data)
		println()
		if err != nil {
			log.Error(err)
			PrintError(externalconstant.CliErrorMessageTemplateOutputError + "\n")
			os.Exit(internalconstant.ExitStatusCode1)
		}
	}
}
func PrintTotalItems(count int) {
	fmt.Printf("\n"+externalconstant.LabelTotalItem+": %d\n", count)
}

func PrintRawRequestResponseHeaders(requestResponseHeaders []model.Header) {
	for _, header := range requestResponseHeaders {
		fmt.Println(header.Name + externalconstant.Colon + header.Value)
	}
}
