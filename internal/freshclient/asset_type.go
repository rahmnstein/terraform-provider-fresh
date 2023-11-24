package freshclient

import (
	"encoding/json"
	"fmt"
)

// GetAssetType gets an asset type from the FreshService API.
func (client *Client) GetAssetType(name string) (*AssetTypeDetails, error) {
	resp, err := client.MakeRequest("GET", *client.APIEndpoint+"/asset_types/?per_page=600", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var assetTypes AssetTypes
	if err := json.NewDecoder(resp.Body).Decode(&assetTypes); err != nil {
		return nil, err
	}

	for _, assetType := range assetTypes.AssetTypes {
		if assetType.Name == name {
			return &assetType, nil
		}
	}

	return nil, fmt.Errorf("asset type %s not found", name)
}
