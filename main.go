package main

import (
	"github.com/akamai/cli-test-center/cmd"
	"github.com/akamai/cli-test-center/logger"
)

var VERSION = "1.0.0"

func main() {
	logger.InitLoggingConfig()
	cmd.Execute(VERSION)
}
