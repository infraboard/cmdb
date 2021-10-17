package cvm

import (
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func NewCVMOperater(client *cvm.Client) *CVMOperater {
	return &CVMOperater{
		client: client,
	}
}

type CVMOperater struct {
	client *cvm.Client
}
