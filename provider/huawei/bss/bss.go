package bss

import (
	v2 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/bss/v2/model"

	"github.com/infraboard/cmdb/app/bill"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewBssOperater(client *v2.BssClient) *BssOperater {
	return &BssOperater{
		client: client,
		log:    zap.L().Named("Huawei BSS"),
	}
}

type BssOperater struct {
	client *v2.BssClient
	log    logger.Logger
}

func (o *BssOperater) transferSet(list *[]model.NvlCostAnalysedBillDetail) *bill.BillSet {
	set := bill.NewBillSet()
	items := *list
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *BssOperater) transferOne(ins model.NvlCostAnalysedBillDetail) *bill.Bill {
	h := bill.NewDefaultBill()
	return h
}
