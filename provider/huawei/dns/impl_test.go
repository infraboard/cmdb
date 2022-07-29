package dns_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/cmdb/apps/domain"
	"github.com/infraboard/cmdb/provider/huawei/connectivity"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/provider"
	op "github.com/infraboard/cmdb/provider/huawei/dns"
)

var (
	operator provider.DnsOperator
)

func TestQuery(t *testing.T) {
	req := provider.NewQueryDomainRequest()
	pager := operator.QueryDomain(req)

	for pager.Next() {
		set := domain.NewDomainSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func init() {
	zap.DevelopmentSetup()
	err := connectivity.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}

	ec, err := connectivity.C().DnsClient()
	if err != nil {
		panic(err)
	}
	operator = op.NewDnsOperator(ec)
}
