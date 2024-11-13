package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/dns_api_ops/common"
	"github.com/dns_api_ops/config"
	"github.com/dns_api_ops/controller"
	"github.com/dns_api_ops/logging"
)

func main() {
	isPrintVersion := flag.Bool("version", false, "[--version]")
	flag.Parse()
	if *isPrintVersion {
		common.GetVersion()
		os.Exit(0)
	}

	execPath := common.GetExecPath()
	err := config.SetConfig(execPath + "/conf/config.yml")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	logging.NewLogger(config.C.GetString("log.level"))
	defer logging.LogFileClose()

	controller.Controller()
}
