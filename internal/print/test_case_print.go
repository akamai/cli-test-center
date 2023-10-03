package print

import (
	"fmt"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/util"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func PrintTestCases(cmd *cobra.Command, testCases []model.TestCase, areAllTestCasesIncluded bool, groupBy string) {

	//Print All associated Test Cases
	util.PrintHeader(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.TestCaseHeader) + "\n")

	if len(testCases) <= 0 {
		util.PrintWarning(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.NoTestCaseWarning))
		//Printing warning here only and returning from method so that empty table does not print
		printMissingTestCasesWarning(cmd, areAllTestCasesIncluded)
		fmt.Println()
		fmt.Println()
		return
	}

	if groupBy == internalconstant.Empty {
		PrintTestCaseDetails(cmd, testCases, false, internalconstant.Empty, 0)
		fmt.Println()
		util.PrintTotalItems(util.GetTotalTestCasesCount(testCases))
	} else {
		printGroupedTestCases(testCases, groupBy)
	}

	printMissingTestCasesWarning(cmd, areAllTestCasesIncluded)
	fmt.Println()
	fmt.Println()
}

func PrintTestCaseDetails(cmd *cobra.Command, testCases []model.TestCase, areDerivedTestCases bool, indentationSpace string, parentOrder int) {

	for i, testCase := range testCases {
		// do not print id for derived test cases.
		if parentOrder == 0 {
			util.PrintLabelAndValue(externalconstant.LabelId, testCase.TestCaseId)
			util.PrintLabelAndValue(indentationSpace+externalconstant.LabelOrder, testCase.Order)
		} else {
			util.PrintLabelAndValue(indentationSpace+externalconstant.LabelOrder, strconv.Itoa(parentOrder)+externalconstant.Dot+strconv.Itoa(testCase.Order))
		}
		//test request
		PrintTestRequestObjects([]model.TestRequest{testCase.TestRequest}, indentationSpace)
		// condition
		util.PrintLabelAndValue(indentationSpace+externalconstant.LabelCondition, util.GetResolvedOrUnResolvedCondition(testCase.Condition))
		// client profile
		util.PrintLabelAndValue(indentationSpace+externalconstant.LabelClientProfile, util.ClientProfileInCLIOutputFormat(testCase.ClientProfile))
		// set variables
		PrintSetVariables(testCase, indentationSpace)
		// audit info
		if !areDerivedTestCases {
			PrintAuditInfo(testCase, indentationSpace)
		}

		if len(testCase.DerivedTestCases) != 0 {
			fmt.Println()
			fmt.Print(externalconstant.IndentationSpace)
			util.PrintHeader(indentationSpace + util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.DerivedTestCaseHeader) + "\n")
			PrintTestCaseDetails(cmd, testCase.DerivedTestCases, true, externalconstant.IndentationSpace, testCase.Order)
		}

		if i != len(testCases)-1 {
			fmt.Println()
			if areDerivedTestCases {
				fmt.Println(indentationSpace + externalconstant.SeparateLineStar)
			} else {
				fmt.Println(externalconstant.SeparateLine)
			}
			fmt.Println()
		}
	}
	areDerivedTestCases = false
	parentOrder = 0
}

func PrintSetVariables(testCase model.TestCase, indentationSpace string) {
	if len(testCase.SetVariables) != 0 {
		var setVariables strings.Builder
		for i, variable := range testCase.SetVariables {
			setVariables.WriteString(util.SetVariablesInCLIOutputFormat(variable))
			if i != len(testCase.SetVariables)-1 {
				setVariables.WriteString("\n")
				setVariables.WriteString(indentationSpace + "               ")
			}
		}
		util.PrintLabelAndValue(indentationSpace+externalconstant.LabelSetVariables, setVariables.String())
	}
}

func PrintAuditInfo(testCase model.TestCase, indentationSpace string) {
	util.PrintLabelAndValue(indentationSpace+externalconstant.LabelCreated, util.FormatTime(testCase.CreatedDate)+externalconstant.SeparateBy+testCase.CreatedBy)
	util.PrintLabelAndValue(indentationSpace+externalconstant.LabelLastModified, util.FormatTime(testCase.ModifiedDate)+externalconstant.SeparateBy+testCase.ModifiedBy)
}
