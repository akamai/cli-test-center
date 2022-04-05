package main

import (
	"github.com/akamai/cli-test-center/cmd"
	"github.com/akamai/cli-test-center/internal"
)

var (
	VERSION string = "0.1.0"
)

func main() {
	internal.InitLoggingConfig()
	cmd.Execute(VERSION)
}
