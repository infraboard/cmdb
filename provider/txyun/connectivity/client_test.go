package connectivity_test

import (
	"os"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"

	"github.com/infraboard/cmdb/provider/txyun/connectivity"
)

func TestClient(t *testing.T) {
	var secretID, secretKey string
	if secretID = os.Getenv("TCCLOUD_SECRET_ID"); secretID == "" {
		t.Fatal("empty TCCLOUD_SECRET_ID")
	}

	if secretKey = os.Getenv("TCCLOUD_SECRET_KEY"); secretKey == "" {
		t.Fatal("empty TCCLOUD_SECRET_KEY")
	}

	client := connectivity.NewTencentCloudClient(secretID, secretKey, regions.Beijing)
	client.CvmClient()
}
