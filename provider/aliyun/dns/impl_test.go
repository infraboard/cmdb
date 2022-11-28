package dns_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/infraboard/cmdb/apps/dns"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/provider/aliyun"
)

var (
	operator provider.DnsOperator
	ctx      = context.Background()
)

func TestQueryDomain(t *testing.T) {
	req := provider.NewQueryDomainRequest()
	pager := operator.QueryDomain(req)
	for pager.Next() {
		set := dns.NewDomainSet()
		if err := pager.Scan(ctx, set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestQueryRecord(t *testing.T) {
	req := provider.NewQueryRecordRequest(os.Getenv("AL_DNS_DOMAIN"))
	pager := operator.QueryRecord(req)
	for pager.Next() {
		set := dns.NewRecordSet()
		if err := pager.Scan(ctx, set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func init() {
	zap.DevelopmentSetup()
	err := aliyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}
	operator = aliyun.O().DnsOperator()
}
