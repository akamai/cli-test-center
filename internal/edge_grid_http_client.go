package internal

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	log "github.com/sirupsen/logrus"
)

type EdgeGridHttpClient struct {
	config           edgegrid.Config
	accountSwitchKey string
}

func NewEdgeGridHttpClient(config edgegrid.Config, accountSwitchKey string) *EdgeGridHttpClient {
	return &EdgeGridHttpClient{config, accountSwitchKey}
}

func (h EdgeGridHttpClient) request(method string, path string, payload *[]byte, headers http.Header) (*http.Response, *[]byte) {
	var (
		err    error
		req    *http.Request
		client = http.Client{}
	)

	var protocol = "https://"
	if strings.Contains(h.config.Host, "http") {
		protocol = "" // For mocking API calls locally
	}

	parsedPath, _ := url.Parse(path)
	if h.accountSwitchKey != "" {
		log.Debugf("Account switch key present :: %s. Adding to URL.", h.accountSwitchKey)
		query := parsedPath.Query()
		query.Set("accountSwitchKey", h.accountSwitchKey)
		parsedPath.RawQuery = query.Encode()
	}

	log.Debugf("Sending request:: %s %s, Headers: %v, Body: %s\n", method, parsedPath, headers, payload)

	if payload != nil {
		req, err = http.NewRequest(method, protocol+h.config.Host+parsedPath.String(), bytes.NewBuffer(*payload))
		if err != nil {
			AbortWithExitCode(err.Error(), ExitStatusCode1)
		}
	} else {
		req, err = http.NewRequest(method, protocol+h.config.Host+parsedPath.String(), nil)
		if err != nil {
			AbortWithExitCode(err.Error(), ExitStatusCode1)
		}
	}
	if headers != nil {
		req.Header = headers
	}

	req = edgegrid.AddRequestHeader(h.config, req)
	// adding this custom header for POST/HEAD data filtering done in TMF, to be removed later once we support POST/HEAD
	req.Header.Add(IsRequestFromCli, RequestIsFromCli)

	resp, er := client.Do(req)
	if er != nil {
		PrintError("\n" + GetEdgeGridErrorMessage("invalidHost") + "\n")
		PrintError(GetGlobalErrorMessage("initEdgeRc") + "\n")
		os.Exit(ExitStatusCode1)
	}
	defer resp.Body.Close()
	byt, err := io.ReadAll(resp.Body)

	if err != nil {
		AbortWithExitCode(err.Error(), ExitStatusCode1)
	}

	log.Debugf("Received response:: Status: %d\n", resp.StatusCode)
	log.Tracef("Response body: %s\n", byt)

	return resp, &byt
}
