package connectivity_test

import (
	"os"
	"testing"

	"github.com/infraboard/cmdb/provider/aws/connectivity"
)

func TestClient(t *testing.T) {
	var secretID, secretKey string
	if secretID = os.Getenv("AMAZON_CLOUD_ACCESS_KEY"); secretID == "" {
		t.Fatal("empty AMAZON_CLOUD_ACCESS_KEY")
	}

	if secretKey = os.Getenv("AMAZON_CLOUD_ACCESS_SECRET"); secretKey == "" {
		t.Fatal("empty AMAZON_CLOUD_ACCESS_SECRET")
	}

	client := connectivity.NewAwsCloudClient(secretID, secretKey, "ap-southeast-1")
	client.Ec2Client()
}
