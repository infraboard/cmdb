package cvm

import (
	"encoding/json"
	"fmt"

	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

// 续费实例询价
// 创建实例询价 https://cloud.tencent.com/document/api/213/15726
// API 参考文档: https://console.cloud.tencent.com/api/explorer?Product=cvm&Version=2017-03-12&Action=InquiryPriceRenewInstances&SignVersion=

func (o *CVMOperator) InquiryPrice(req *cvm.InquiryPriceRunInstancesRequest) error {
	resp, err := o.client.InquiryPriceRunInstances(req)
	if err != nil {
		return err
	}

	v, _ := json.Marshal(resp)
	fmt.Println(string(v))
	return nil
}
