package freshclient

import (
	"os"
	"testing"
)

func TestGetAssetType(t *testing.T) {
	// Check if env vars are set
	if os.Getenv("FRESHDESK_API_KEY_TEST") == "" || os.Getenv("FRESHDESK_API_ENDPOINT_TEST") == "" {
		t.Errorf("TestAsset() error = %v, want %v", "Please set FRESHDESK_API_KEY_TEST and FRESHDESK_API_ENDPOINT_TEST", nil)
		t.FailNow()
	}

	// Retrieve API credentials from environment variables
	client := NewClient(os.Getenv("FRESHDESK_API_KEY_TEST"), os.Getenv("FRESHDESK_API_ENDPOINT_TEST"))

	got, err := client.GetAssetType("VMware VCenter VM")
	if err != nil {
		t.Errorf("freshclient.GetAssetType() error = %v, want %v", err, nil)
		t.FailNow()
	}
	if got == nil {
		t.Errorf("freshclient.GetAssetType() = %v, want %v", got, nil)
		t.FailNow()
	}
}
