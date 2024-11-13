package modules_test

import (
	"fmt"
	"log/slog"
	"os"
	"testing"

	"github.com/dns_api_ops/cloudprovider"
	"github.com/dns_api_ops/common"
	"github.com/dns_api_ops/config"
	"github.com/dns_api_ops/logging"
	"github.com/dns_api_ops/modules"
)

func TestDnsOps(t *testing.T) {
	execPath := common.GetExecPath()
	err := config.SetConfig(execPath + "../conf/config.yaml")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	logging.NewLogger(config.C.GetString("log.level"))
	defer logging.LogFileClose()

	accessKeyId := config.C.GetString("accessKeyId")
	accessSecret := config.C.GetString("accessSecret")
	regionId := config.C.GetString("regionId")
	alidns, err := cloudprovider.NewAliDNS(&accessKeyId, &accessSecret, &regionId)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	dnsOps := modules.NewDNSOPS(alidns)
	ip := dnsOps.DescribeIP()
	fmt.Print(ip)
}
