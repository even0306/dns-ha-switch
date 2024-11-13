package cloudprovider

import (
	dns "github.com/alibabacloud-go/alidns-20150109/v2/client"
)

type Aliyun struct {
	Client *dns.Client
}

// 更新域名解析记录
func (d *Aliyun) UpdateDomainRecord(recordId *string, rr *string, recordType *string, Value *string) (*dns.UpdateDomainRecordResponse, error) {
	request := &dns.UpdateDomainRecordRequest{}
	request.Value = Value
	request.RR = rr
	request.RecordId = recordId
	request.Type = recordType
	response, err := d.Client.UpdateDomainRecord(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// 查询域名解析记录
func (d *Aliyun) DescribeDomainRecords(domainName *string) (*dns.DescribeDomainRecordsResponse, error) {
	request := &dns.DescribeDomainRecordsRequest{}
	request.DomainName = domainName
	response, err := d.Client.DescribeDomainRecords(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
