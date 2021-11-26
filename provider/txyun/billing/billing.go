package billing

import (
	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"

	"github.com/infraboard/cmdb/app/bill"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewBillingOperater(client *billing.Client) *BillingOperater {
	return &BillingOperater{
		client: client,
		log:    zap.L().Named("Tx Billing"),
	}
}

type BillingOperater struct {
	client *billing.Client
	log    logger.Logger
}

func (o *BillingOperater) transferSet(items []*billing.BillResourceSummary) *bill.BillSet {
	set := bill.NewBillSet()
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *BillingOperater) transferOne(ins *billing.BillResourceSummary) *bill.Bill {
	b := bill.NewDefaultBill()
	return b
}
