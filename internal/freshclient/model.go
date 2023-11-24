package freshclient

// Asset represents a FreshService asset
// asset
type Asset struct {
	// AssetDetails
	AssetDetails AssetDetails `json:"asset"`
}

// AssetDetails represents a FreshService asset field
// agent_id
// asset_tag
// asset_type_id
// assigned_on
// author_type
// created_at
// department_id
// description
// discovery_enabled
// display_id
// end_of_life
// group_id
// id
// impact
// location_id
// name
// updated_at
// usage_type
// user_id
type AssetDetails struct {
	AgentID      int64  `json:"agent_id,omitempty"`
	AssetTag     string `json:"asset_tag,omitempty"`
	AssetTypeID  int64  `json:"asset_type_id"`
	AssignedOn   string `json:"assigned_on,omitempty"`
	AuthorType   string `json:"author_type,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	DepartmentID int64  `json:"department_id,omitempty"`
	Description  string `json:"description,omitempty"`
	DisplayID    int64  `json:"display_id,omitempty"`
	EndOfLife    string `json:"end_of_life,omitempty"`
	GroupID      int64  `json:"group_id,omitempty"`
	ID           int64  `json:"id,omitempty"`
	Impact       string `json:"impact,omitempty"`
	LocationID   int64  `json:"location_id,omitempty"`
	Name         string `json:"name"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	UsageType    string `json:"usage_type,omitempty"`
	UserID       int64  `json:"user_id,omitempty"`
}

// ToAssetDetailsUpdate converts an AssetDetails to AssetDetailsUpdate
func (details AssetDetails) ToAssetDetailsUpdate() AssetDetailsUpdate {
	return AssetDetailsUpdate{
		AssetTag:     details.AssetTag,
		AssetTypeID:  details.AssetTypeID,
		AssignedOn:   details.AssignedOn,
		DepartmentID: details.DepartmentID,
		Description:  details.Description,
		EndOfLife:    details.EndOfLife,
		GroupID:      details.GroupID,
		Impact:       details.Impact,
		LocationID:   details.LocationID,
		Name:         details.Name,
		UsageType:    details.UsageType,
		UserID:       details.UserID,
	}
}

type AssetDetailsUpdate struct {
	AssetTag         string `json:"asset_tag,omitempty"`
	AssetTypeID      int64  `json:"asset_type_id"`
	AssignedOn       string `json:"assigned_on,omitempty"`
	DepartmentID     int64  `json:"department_id,omitempty"`
	Description      string `json:"description,omitempty"`
	DiscoveryEnabled bool   `json:"discovery_enabled,omitempty"`
	EndOfLife        string `json:"end_of_life,omitempty"`
	GroupID          int64  `json:"group_id,omitempty"`
	Impact           string `json:"impact,omitempty"`
	LocationID       int64  `json:"location_id,omitempty"`
	Name             string `json:"name"`
	UsageType        string `json:"usage_type,omitempty"`
	UserID           int64  `json:"user_id,omitempty"`
}

type User struct {
	ID        int64  `json:"id,omitempty"`
	UserName  string `json:"user_name,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

// AssetType represents a FreshService asset type
// asset_type
type AssetType struct {
	// AssetTypeDetails
	AssetTypeDetails AssetTypeDetails `json:"asset_type"`
}

// AssetTypeDetails represents a FreshService asset field
// created_at
// description
// id
// name
// parent_asset_type_id
// updated_at
// visible
type AssetTypeDetails struct {
	CreatedAt         string `json:"created_at"`
	Description       string `json:"description"`
	ID                int64  `json:"id"`
	Name              string `json:"name"`
	ParentAssetTypeID int64  `json:"parent_asset_type_id"`
	UpdatedAt         string `json:"updated_at"`
	Visible           bool   `json:"visible"`
}

// AssetTypes represents a FreshService asset type
// JSON Example:
/*
{
    "asset_types": [
        {
            "id": 1,
            "name": "Services",
            "parent_asset_type_id": null,
            "description": "",
            "visible": true
            "created_at": "2019-02-14T10:08:02Z",
            "updated_at": "2019-02-14T10:08:02Z"
        },
        {
            "id": 2,
            "name": "Cloud",
            "parent_asset_type_id": null,
            "description": "",
            "visible": true
            "created_at": "2019-02-14T10:03:02Z",
            "updated_at": "2019-02-14T10:03:02Z"
        }
    ]
}
*/
type AssetTypes struct {
	AssetTypes []AssetTypeDetails `json:"asset_types"`
}
