package controller

import (
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/dns_api_ops/cloudprovider"
	"github.com/dns_api_ops/config"
	"github.com/dns_api_ops/modules"
)

func Controller() {
	accessKeyId := config.C.GetString("accessKeyId")
	accessSecret := config.C.GetString("accessSecret")
	regionId := config.C.GetString("regionId")
	alidns, err := cloudprovider.NewAliDNS(&accessKeyId, &accessSecret, &regionId)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	dnsOPS := modules.NewDNSOPS(alidns)
	ipPorts := config.C.GetStringSlice("value")
	ipDisabled := false
	for {
		nowDNSIP := dnsOPS.DescribeIP()
		for _, ipPort := range ipPorts {
			ip := strings.Split(ipPort, ":")

			if ipDisabled {
				ok := modules.CheckPort(ipPort)
				if !ok {
					slog.Warn("存在服务失联！", "ip_port", ipPort)
					continue
				}

				dnsOPS.DescribeIP()
				dnsOPS.ChangeIP(&ip[0])
				ipDisabled = false
			}

			if ip[0] != nowDNSIP {
				continue
			} else {
				i := 0
				for {
					ok := modules.CheckPort(ipPort)
					// 如果5次内重新连接上了当前dns所设置的ip端口，则计数器重置，重新累计5次
					if ok {
						i = 0
					}

					// 判断连续5次以上无法连接当前dns所设置的ip端口，则退出循环，切换dns的IP
					if i >= 5 {
						ipDisabled = true
						break
					}

					// 如果出现无法连接当前dns所设置的ip端口，则计数器+1
					if !ok {
						slog.Warn("存在服务失联！", "ip_port", ipPort)
						i++
						continue
					}

					time.Sleep(15 * time.Second)
				}
			}
		}
	}
}
