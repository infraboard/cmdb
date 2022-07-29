package dns_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/infraboard/cmdb/apps/dns"
	"github.com/infraboard/cmdb/provider"
	"github.com/infraboard/cmdb/provider/txyun/connectivity"
	op "github.com/infraboard/cmdb/provider/txyun/dns"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	operator *op.DnsOperator
)

func TestQueryDomain(t *testing.T) {
	req := provider.NewQueryDomainRequest()
	pager := operator.QueryDomain(req)

	for pager.Next() {
		set := dns.NewDomainSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		for i := range set.Items {
			fmt.Println(set.Items[i])
		}
	}
}

func TestQueryReord(t *testing.T) {
	req := provider.NewQueryRecordRequest(os.Getenv("TX_DNS_DOMAIN"))
	pager := operator.QueryRecord(req)

	for pager.Next() {
		set := dns.NewRecordSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		for i := range set.Items {
			fmt.Println(set.Items[i])
		}
	}
}

func init() {
	zap.DevelopmentSetup()
	err := connectivity.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}

	client := connectivity.C()
	operator = op.NewDnsOperator(client.DnsClient())
}
