package cloudprovider

import (
	dns "github.com/alibabacloud-go/alidns-20150109/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
)

type DnsAPI interface {
	UpdateDomainRecord(recordId *string, RR *string, recordType *string, Value *string) (*dns.UpdateDomainRecordResponse, error)
	DescribeDomainRecords(domainName *string) error
}

func NewAliDNS(accessKeyId *string, accessKeySecret *string, regionId *string) (*Aliyun, error) {
	config := &openapi.Config{}
	// 传AccessKey ID入config
	config.AccessKeyId = accessKeyId
	// 传AccessKey Secret入config
	config.AccessKeySecret = accessKeySecret
	config.RegionId = regionId
	result := &dns.Client{}
	result, err := dns.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Aliyun{
		Client: result,
	}, nil
}
