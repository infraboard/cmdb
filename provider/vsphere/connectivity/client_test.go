package connectivity_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/infraboard/cmdb/provider/vsphere/connectivity"
)

func TestClient(t *testing.T) {
	var host, username, password string
	if host = os.Getenv("VS_HOST"); host == "" {
		t.Fatal("empty VS_HOST")
	}

	if username = os.Getenv("VS_USERNAME"); username == "" {
		t.Fatal("empty VS_USERNAME")
	}

	if password = os.Getenv("VS_PASSWORD"); password == "" {
		t.Fatal("empty VS_PASSWORD")
	}

	client := connectivity.NewVsphereClient(host, username, password)
	vim, err := client.VimClient()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(vim.Client.URL().Host)
}
