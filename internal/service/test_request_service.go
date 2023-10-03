package service

import (
	internalconstant "github.com/akamai/cli-test-center/internal/constant"
	"github.com/akamai/cli-test-center/internal/model"
	"github.com/akamai/cli-test-center/internal/print"
	"github.com/akamai/cli-test-center/internal/util"
	"github.com/spf13/cobra"
)

func (svc Service) GetTestRequestAndPrintResult(cmd *cobra.Command) {
	testReqList := svc.GetTestRequests()
	print.PrintTestRequests(cmd, testReqList)
}

func (svc Service) GetTestRequests() []model.TestRequest {
	spinner := util.NewSpinner(util.GetServiceMessage(svc.cmd, internalconstant.MessageTypeSpinner,
		internalconstant.Empty, internalconstant.GetTestRequests), !svc.jsonOutput).Start()

	testRequests, err := svc.api.GetTestRequestsV3()
	if err != nil {
		spinner.StopWithFailure()
		util.AbortForCommand(svc.cmd, err)
	}

	spinner.StopWithSuccess()

	if svc.jsonOutput {
		util.PrintJsonAndExit(testRequests)
	}
	return testRequests

}
