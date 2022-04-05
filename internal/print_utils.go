package internal

// This class will have only akamai/global cli standard print functions

import (
	"fmt"
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

func bold(a ...interface{}) string {
	c := color.New(color.Bold)
	return c.Sprint(a...)
}

func italic(a ...interface{}) string {
	c := color.New(color.Italic)
	return c.Sprint(a...)
}

func printLabelAndValue(label string, value interface{}) {
	c := color.New(color.Bold)
	_, _ = c.Printf(label + ": ")
	fmt.Printf("%v\n", value)
}

func PrintJsonAndExit(data interface{}) {
	// This uses a third-party replacement (github.com/clarketm/json) for Go's default JSON encoder.
	// It considers empty structs (as opposed to just nil pointer to struct) for omitempty.
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
		AbortWithExitCode(GetGlobalErrorMessage("jsonOutputFailed"), ExitStatusCode1)
	}

	if string(b) != "null" {
		fmt.Println(string(b) + "\n")
	} else {
		fmt.Println("[]\n")
	}
	os.Exit(ExitStatusCode0)
}

func AbortWithExitCode(message string, code int) {
	PrintError(message + "\n")
	os.Exit(code)
}

func AbortWithUsageAndMessageAndCode(cmd *cobra.Command, message string, code int) {
	if message != Empty {
		PrintError(message + "\n\n")
	}
	err := cmd.Usage()
	fmt.Println()
	if err != nil {
		return
	}
	os.Exit(code)
}
func AbortForCommand(cmd *cobra.Command, cliError *CliError) {
	AbortForCommandWithSubResource(cmd, cliError, Empty, Empty)
}

func AbortForCommandWithSubResource(cmd *cobra.Command, cliError *CliError, subResource, operation string) {

	responseCode := strconv.Itoa(cliError.responseCode)
	if len(responseCode) == 3 && strings.Contains("500,502,503,504,405", responseCode) {
		PrintError(GetGlobalErrorMessage(responseCode) + "\n\n")
	} else if cliError.apiError != nil {
		printErrorMessages(GetApiErrorMessagesForCommand(cmd, *cliError.apiError, subResource, operation, responseCode))
		println()
	} else if len(cliError.apiSubErrors) != 0 {
		printErrorMessages(GetApiSubErrorMessagesForCommand(cmd, cliError.apiSubErrors, "", subResource, operation))
		println()
	} else {
		PrintError(cliError.errorMessage + "\n\n")
	}
	// Get equivalent exit status code for corresponding http status code
	statusCode := GetHttpExitCode(responseCode)
	os.Exit(statusCode)
}

func printErrorMessages(errorMessages []string) {
	for _, message := range errorMessages {
		PrintError(message + "\n")
	}
}

// ShowTable Standard function to show table in same format across Test Center CLI
func ShowTable(tableHeaders []string, tableContents [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)

	table.SetHeader(tableHeaders)
	table.AppendBulk(tableContents)
	table.Render()
	fmt.Printf("\n"+LabelTotalItem+": %d\n", len(tableContents))
}

func printTemplate(templateToParse string, data interface{}) {

	// standard functions for templates
	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"printHeader":        RequestHeaderInCLIOutputFormat,
		"printClientProfile": ClientProfileInCLIOutputFormat,
		"bold":               bold,
		"join": func(elements []string) string {
			return strings.Join(elements, ",")
		},
	}

	tmp, err := template.New("template").Funcs(funcMap).Parse(templateToParse)

	// if there is error,
	if err != nil {
		log.Error(err)
		PrintError(CliErrorMessageTemplateOutputError + "\n")
		os.Exit(ExitStatusCode1)
	} else {
		// standard output to print merged data
		err = tmp.Execute(os.Stdout, data)
		println()
		if err != nil {
			log.Error(err)
			PrintError(CliErrorMessageTemplateOutputError + "\n")
			os.Exit(ExitStatusCode1)
		}
	}
}
