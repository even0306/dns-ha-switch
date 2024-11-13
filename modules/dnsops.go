package modules

import (
	"log/slog"
	"os"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/dns_api_ops/cloudprovider"
	"github.com/dns_api_ops/config"
)

type DNSOPS interface {
	// 修改dns解析ip
	ChangeIP(value *string)
	// 查询当前dns解析ip
	DescribeIP()
}

type DNSINFO struct {
	recordId   string
	rr         string
	recordType string
	ipv4addr   string
	c          *cloudprovider.Aliyun
}

func NewDNSOPS(c *cloudprovider.Aliyun) *DNSINFO {
	return &DNSINFO{
		recordId:   "",
		rr:         "",
		recordType: "",
		c:          c,
	}
}

func (dnsInfo *DNSINFO) ChangeIP(value *string) {
	resp2, err := dnsInfo.c.UpdateDomainRecord(&dnsInfo.recordId, &dnsInfo.rr, &dnsInfo.recordType, value)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info(resp2.Body.GoString())
}

func (dnsInfo *DNSINFO) DescribeIP() string {
	domainName := config.C.GetString("domainName")
	resp1, err := dnsInfo.c.DescribeDomainRecords(&domainName)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	dnsInfo.rr = config.C.GetString("rr")
	dnsInfo.recordType = config.C.GetString("recordType")

	for _, v := range resp1.Body.DomainRecords.Record {
		m := tea.ToMap(v)
		if m["RR"] == dnsInfo.rr {
			dnsInfo.recordId = m["RecordId"].(string)
			dnsInfo.ipv4addr = m["Value"].(string)
			break
		}
	}
	return dnsInfo.ipv4addr
}
