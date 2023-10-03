package print

import (
	"fmt"
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/util"
	externalconstant "github.com/akamai/cli-test-center/user/constant"
	"github.com/spf13/cobra"
	"strings"
)

func PrintTestRequests(cmd *cobra.Command, testRequests []model.TestRequest) {
	util.PrintHeader(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.TestRequests))
	fmt.Println()
	fmt.Println()

	if len(testRequests) <= 0 {
		util.PrintWarning(util.GetServiceMessage(cmd, internalconstant.MessageTypeDisplay, internalconstant.Empty, internalconstant.TestRequestNotFoundError))
		fmt.Println()
		return
	}

	PrintTestRequestObjects(testRequests, internalconstant.Empty)
	fmt.Println()
	util.PrintTotalItems(len(testRequests))
	fmt.Println()

}

func PrintTestRequestObjects(testRequests []model.TestRequest, indentationSpaces string) {

	for i, tr := range testRequests {
		util.PrintLabelAndValue(indentationSpaces+externalconstant.LabelRequestMethod, tr.RequestMethod)
		util.PrintLabelAndValue(indentationSpaces+externalconstant.LabelTRUrl, util.GetResolvedOrUnResolvedRequestURL(tr))

		//tags
		if len(tr.Tags) > 0 {
			util.PrintLabelAndValue(indentationSpaces+externalconstant.LabelTags, strings.Join(tr.Tags, externalconstant.Comma))
		}

		//request headers
		if len(tr.RequestHeaders) > 0 {
			var headers = strings.Builder{}
			for j, header := range tr.RequestHeaders {
				headers.WriteString(util.GetResolvedOrUnResolvedHeaders(header))
				if j != len(tr.RequestHeaders)-1 {
					headers.WriteString("\n")
					headers.WriteString(indentationSpaces + "                    ")
				}
			}
			util.PrintLabelAndValue(indentationSpaces+externalconstant.LabelRequestHeaders, headers.String())
		}

		if tr.RequestMethod == internalconstant.Post {
			util.PrintLabelAndValue(indentationSpaces+externalconstant.LabelRequestBody, util.GetResolvedOrUnResolvedRequestBody(tr))
			util.PrintLabelAndValue(indentationSpaces+externalconstant.LabelEncodeRequestBody, *tr.EncodeRequestBody)
		}

		if i != len(testRequests)-1 {
			fmt.Println()
			fmt.Println(externalconstant.SeparateLine)
			fmt.Println()
		}
	}
}
