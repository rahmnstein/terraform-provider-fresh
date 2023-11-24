package freshclient

import (
	"encoding/json"
	"strconv"
)

// CreateAsset creates an asset in the FreshService API.
func (client *Client) CreateAsset(assetDetails AssetDetails) (*AssetDetails, error) {
	// Make the request
	resp, err := client.MakeRequest("POST", *client.APIEndpoint+"/assets", assetDetails)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var newAsset Asset
	if err := json.NewDecoder(resp.Body).Decode(&newAsset); err != nil {
		return nil, err
	}

	return &newAsset.AssetDetails, nil
}

// GetAsset gets an asset from the FreshService API.
func (client *Client) GetAsset(assetDipslayID int64) (*AssetDetails, error) {
	// Make the request
	resp, err := client.MakeRequest("GET", *client.APIEndpoint+"/assets/"+strconv.FormatInt(assetDipslayID, 10), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var asset Asset
	if err := json.NewDecoder(resp.Body).Decode(&asset); err != nil {
		return nil, err
	}

	return &asset.AssetDetails, nil
}

// UpdateAsset updates an asset in the FreshService API.
func (client *Client) UpdateAsset(assetDetails AssetDetails) (*AssetDetails, error) {

	// Make the request
	resp, err := client.MakeRequest("PUT", *client.APIEndpoint+"/assets/"+strconv.FormatInt(assetDetails.DisplayID, 10), assetDetails.ToAssetDetailsUpdate())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var updatedAsset Asset
	if err := json.NewDecoder(resp.Body).Decode(&updatedAsset); err != nil {
		return nil, err
	}

	return &updatedAsset.AssetDetails, nil
}

// DeleteAsset deletes an asset from the FreshService API.
func (client *Client) DeleteAsset(assetDetails AssetDetails) error {
	// Make the request
	_, err := client.MakeRequest("DELETE", *client.APIEndpoint+"/assets/"+strconv.FormatInt(assetDetails.DisplayID, 10), nil)
	if err != nil {
		return err
	}

	return nil
}
