package provider

import (
	"context"
	"terraform-provider-fresh/internal/freshclient"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &AssetTypeDataSource{}

func NewAssetTypeDataSource() datasource.DataSource {
	return &AssetTypeDataSource{}
}

type AssetTypeDataSource struct {
	client *freshclient.Client
}

type AssetTypeDataSourceModel struct {
	CreatedAt         types.String `tfsdk:"created_at"`
	Description       types.String `tfsdk:"description"`
	ID                types.Int64  `tfsdk:"id"`
	Name              types.String `tfsdk:"name"`
	ParentAssetTypeID types.Int64  `tfsdk:"parent_asset_type_id"`
	UpdatedAt         types.String `tfsdk:"updated_at"`
	Visible           types.Bool   `tfsdk:"visible"`
}

// Metadata returns the metadata for the data source.
func (d *AssetTypeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_asset_type"
}

func (m AssetTypeDataSourceModel) fromFreshAssetType(assetType freshclient.AssetTypeDetails) AssetTypeDataSourceModel {
	return AssetTypeDataSourceModel{
		CreatedAt:         types.StringValue(assetType.CreatedAt),
		Description:       types.StringValue(assetType.Description),
		ID:                types.Int64Value(assetType.ID),
		Name:              types.StringValue(assetType.Name),
		ParentAssetTypeID: types.Int64Value(assetType.ParentAssetTypeID),
		UpdatedAt:         types.StringValue(assetType.UpdatedAt),
		Visible:           types.BoolValue(assetType.Visible),
	}
}

func (d *AssetTypeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Asset Type Data Source",

		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Date and time of creation",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the asset type",
				Computed:            true,
				Optional:            true,
				Required:            false,
			},
			"id": schema.Int64Attribute{
				MarkdownDescription: "Unique ID of the asset type",
				Computed:            true,
				Optional:            true,
				Required:            false,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the asset type",
				Required:            true,
			},
			"parent_asset_type_id": schema.Int64Attribute{
				MarkdownDescription: "ID of the parent asset type",
				Computed:            true,
				Optional:            true,
				Required:            false,
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "Date and time of last update",
				Computed:            true,
				Optional:            true,
				Required:            false,
			},
			"visible": schema.BoolAttribute{
				MarkdownDescription: "Whether the asset type is visible",
				Computed:            true,
				Optional:            true,
				Required:            false,
			},
		},
	}
}

func (d *AssetTypeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

// Read the data source and convert it into a resource object
func (d *AssetTypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AssetTypeDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	assetTypeDetails, err := d.client.GetAssetType(data.Name.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("Error getting asset type", err.Error())
		return
	}

	// Save data into Terraform state
	data = data.fromFreshAssetType(*assetTypeDetails)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}
