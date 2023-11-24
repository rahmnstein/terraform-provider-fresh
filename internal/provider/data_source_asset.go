package provider

import (
	"context"
	"terraform-provider-fresh/internal/freshclient"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &AssetDataSource{}

func NewAssetDataSource() datasource.DataSource {
	return &AssetDataSource{}
}

type AssetDataSource struct {
	client *freshclient.Client
}

type AssetDataSourceModel struct {
	Name         types.String `tfsdk:"name"`
	AssetTag     types.String `tfsdk:"asset_tag"`
	AssetTypeID  types.Int64  `tfsdk:"asset_type_id"`
	AssignedOn   types.String `tfsdk:"assigned_on"`
	AuthorType   types.String `tfsdk:"author_type"`
	CreatedAt    types.String `tfsdk:"created_at"`
	DepartmentID types.Int64  `tfsdk:"department_id"`
	Description  types.String `tfsdk:"description"`
	DisplayID    types.Int64  `tfsdk:"display_id"`
	EndOfLife    types.String `tfsdk:"end_of_life"`
	GroupID      types.Int64  `tfsdk:"group_id"`
	ID           types.Int64  `tfsdk:"id"`
	Impact       types.String `tfsdk:"impact"`
	LocationID   types.Int64  `tfsdk:"location_id"`
	UpdatedAt    types.String `tfsdk:"updated_at"`
	UsageType    types.String `tfsdk:"usage_type"`
	UserID       types.Int64  `tfsdk:"user_id"`
}

// Metadata returns the metadata for the data source.
func (d *AssetDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_asset"
}

func (m AssetDataSourceModel) fromFreshAsset(assetDetails freshclient.AssetDetails) AssetDataSourceModel {
	return AssetDataSourceModel{
		Name:         types.StringValue(assetDetails.Name),
		AssetTag:     types.StringValue(assetDetails.AssetTag),
		AssetTypeID:  types.Int64Value(assetDetails.AssetTypeID),
		AssignedOn:   types.StringValue(assetDetails.AssignedOn),
		AuthorType:   types.StringValue(assetDetails.AuthorType),
		CreatedAt:    types.StringValue(assetDetails.CreatedAt),
		DepartmentID: types.Int64Value(assetDetails.DepartmentID),
		Description:  types.StringValue(assetDetails.Description),
		DisplayID:    types.Int64Value(assetDetails.DisplayID),
		EndOfLife:    types.StringValue(assetDetails.EndOfLife),
		GroupID:      types.Int64Value(assetDetails.GroupID),
		ID:           types.Int64Value(assetDetails.ID),
		Impact:       types.StringValue(assetDetails.Impact),
		LocationID:   types.Int64Value(assetDetails.LocationID),
		UpdatedAt:    types.StringValue(assetDetails.UpdatedAt),
		UsageType:    types.StringValue(assetDetails.UsageType),
		UserID:       types.Int64Value(assetDetails.UserID),
	}
}

func (d *AssetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Asset Type Data Source",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the asset type",
				Computed:            true,
			},
			"asset_tag": schema.StringAttribute{
				MarkdownDescription: "Asset tag of the asset type",
				Computed:            true,
			},
			"asset_type_id": schema.Int64Attribute{
				MarkdownDescription: "ID of the asset type",
				Computed:            true,
			},
			"assigned_on": schema.StringAttribute{
				MarkdownDescription: "Date and time of assignment",
				Computed:            true,
			},
			"author_type": schema.StringAttribute{
				MarkdownDescription: "Type of author",
				Computed:            true,
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Date and time of creation",
				Computed:            true,
			},
			"department_id": schema.Int64Attribute{
				MarkdownDescription: "ID of the department",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the asset type",
				Computed:            true,
			},
			"display_id": schema.Int64Attribute{
				MarkdownDescription: "Display ID of the asset type",
				Required:            true,
			},
			"end_of_life": schema.StringAttribute{
				MarkdownDescription: "Date and time of end of life",
				Computed:            true,
			},
			"group_id": schema.Int64Attribute{
				MarkdownDescription: "ID of the group",
				Computed:            true,
			},
			"id": schema.Int64Attribute{
				MarkdownDescription: "Unique ID of the asset type",
				Computed:            true,
			},
			"impact": schema.StringAttribute{
				MarkdownDescription: "Impact of the asset type",
				Computed:            true,
			},
			"location_id": schema.Int64Attribute{
				MarkdownDescription: "ID of the location",
				Computed:            true,
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "Date and time of last update",
				Computed:            true,
			},
			"usage_type": schema.StringAttribute{
				MarkdownDescription: "Usage type of the asset type",
				Computed:            true,
			},
			"user_id": schema.Int64Attribute{
				MarkdownDescription: "ID of the user",
				Computed:            true,
			},
		},
	}
}

func (d *AssetDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*freshclient.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			"Expected *freshclient.Client, got: %T. Please report this issue to the provider developers.",
		)

		return
	}

	d.client = client
}

// Read the data source and convert it into a resource object.
func (d *AssetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AssetDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// tflog.Info(ctx, data.DisplayID.)
	assetDetails, err := d.client.GetAsset(data.DisplayID.ValueInt64())

	if err != nil {
		resp.Diagnostics.AddError("Error getting asset", err.Error())
		return
	}

	// Save data into Terraform state
	data = data.fromFreshAsset(*assetDetails)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}
