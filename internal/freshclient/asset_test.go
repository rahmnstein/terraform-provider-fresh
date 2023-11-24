package freshclient

import (
	"os"
	"testing"
)

// TestAsset tests the UpdateAsset function.
func TestAsset(t *testing.T) {
	// Check if env vars are set
	if os.Getenv("FRESHDESK_API_KEY_TEST") == "" || os.Getenv("FRESHDESK_API_ENDPOINT_TEST") == "" {
		t.Errorf("TestAsset() error = %v, want %v", "Please set FRESHDESK_API_KEY_TEST and FRESHDESK_API_ENDPOINT_TEST", nil)
		t.FailNow()
	}

	// Retrieve API credentials from environment variables
	client := NewClient(os.Getenv("FRESHDESK_API_KEY_TEST"), os.Getenv("FRESHDESK_API_ENDPOINT_TEST"))

	// Create an asset for testing update
	assetDetails := AssetDetails{
		Name:        "TestGolangAsset",
		AssetTypeID: 50000240147,
	}

	createdAsset, err := client.CreateAsset(assetDetails)
	if err != nil {
		t.Errorf("freshclient.CreateAsset() error = %v, want %v", err.Error(), nil)
		t.FailNow()
	}

	getAsset, err := client.GetAsset(createdAsset.DisplayID)
	if err != nil {
		t.Errorf("freshclient.GetAsset() error = %v, want %v", err.Error(), nil)
		t.FailNow()
	}

	// Check if the createdAsset have same id as getAsset
	if createdAsset.DisplayID != getAsset.DisplayID {
		t.Errorf("freshclient.GetAsset() error = %v, want %v", getAsset.DisplayID, createdAsset.DisplayID)
		t.FailNow()
	}

	// Modify the createdAsset or create a new AssetDetailsUpdate object with changes
	createdAsset.Description = "TestAssetUpdate"

	// Update the asset
	updatedAsset, err := client.UpdateAsset(*createdAsset)
	if err != nil {
		t.Errorf("freshclient.UpdateAsset() error = %v, want %v", err, nil)
		t.FailNow()
	}

	// Add assertions to validate the update if needed
	if updatedAsset.Description != "TestAssetUpdate" {
		t.Errorf("freshclient.UpdateAsset() error = %v, want %v", updatedAsset.Description, "TestAssetUpdate")
		t.FailNow()
	}

	// Cleanup: Delete the asset created for testing update
	err = client.DeleteAsset(*updatedAsset)
	if err != nil {
		t.Errorf("freshclient.DeleteAsset() error = %v, want %v", err, nil)
	}
}
