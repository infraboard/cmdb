package connectivity_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/infraboard/cmdb/provider/huawei/connectivity"
)

func TestClient(t *testing.T) {
	var secretID, secretKey string
	if secretID = os.Getenv("HW_CLOUD_ACCESS_KEY"); secretID == "" {
		t.Fatal("empty HW_CLOUD_ACCESS_KEY")
	}

	if secretKey = os.Getenv("HW_CLOUD_ACCESS_SECRET"); secretKey == "" {
		t.Fatal("empty HW_CLOUD_ACCESS_SECRET")
	}

	client := connectivity.NewHuaweiCloudClient(secretID, secretKey, "cn-north-4")
	ec, _ := client.BssClient()
	fmt.Println(ec.HcClient)
}
