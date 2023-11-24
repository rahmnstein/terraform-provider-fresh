package provider

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"terraform-provider-fresh/internal/freshclient"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure AssetResource satisfies various resource interfaces.
var _ resource.Resource = &AssetResource{}
var _ resource.ResourceWithImportState = &AssetResource{}

// NewAssetResource returns a new resource.
func NewAssetResource() resource.Resource {
	return &AssetResource{}
}

// AssetResource defines the resource implementation.
type AssetResource struct {
	client *freshclient.Client
}

// AssetResourceModel describes the resource data model.
type AssetResourceModel struct {
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

func (m AssetResourceModel) fromFreshAsset(assetDetails freshclient.AssetDetails) AssetResourceModel {
	return AssetResourceModel{
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

func (m AssetResourceModel) toFreshAsset() freshclient.AssetDetails {
	return freshclient.AssetDetails{
		AgentID:      m.AssetTypeID.ValueInt64(), // Set to your actual value or leave it as 0 based on your logic
		AssetTag:     m.AssetTag.ValueString(),
		AssetTypeID:  m.AssetTypeID.ValueInt64(),
		AssignedOn:   m.AssignedOn.ValueString(),
		AuthorType:   m.AuthorType.ValueString(),
		CreatedAt:    m.CreatedAt.ValueString(),
		DepartmentID: m.DepartmentID.ValueInt64(),
		Description:  m.Description.ValueString(),
		DisplayID:    m.DisplayID.ValueInt64(),
		EndOfLife:    m.EndOfLife.ValueString(),
		GroupID:      m.GroupID.ValueInt64(),
		ID:           m.ID.ValueInt64(),
		Impact:       m.Impact.ValueString(),
		LocationID:   m.LocationID.ValueInt64(),
		Name:         m.Name.ValueString(),
		UpdatedAt:    m.UpdatedAt.ValueString(),
		UsageType:    m.UsageType.ValueString(),
		UserID:       m.UserID.ValueInt64(),
	}
}

// Metadata returns the metadata for the resource.
func (r *AssetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_asset"
}

// Schema returns the schema for the resource.
func (r *AssetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Asset Resource",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the asset type",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"asset_tag": schema.StringAttribute{
				MarkdownDescription: "Asset tag of the asset type",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"asset_type_id": schema.Int64Attribute{
				MarkdownDescription: "ID of the asset type",
				Required:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"assigned_on": schema.StringAttribute{
				MarkdownDescription: "Date and time of assignment",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"author_type": schema.StringAttribute{
				MarkdownDescription: "Type of author",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Date and time of creation",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"department_id": schema.Int64Attribute{
				MarkdownDescription: "ID of the department",
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the asset type",
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"display_id": schema.Int64Attribute{
				MarkdownDescription: "Display ID of the asset type",
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"end_of_life": schema.StringAttribute{
				MarkdownDescription: "Date and time of end of life",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"group_id": schema.Int64Attribute{
				MarkdownDescription: "ID of the group",
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.Int64Attribute{
				MarkdownDescription: "Unique ID of the asset type",
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"impact": schema.StringAttribute{
				MarkdownDescription: "Impact of the asset type",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"location_id": schema.Int64Attribute{
				MarkdownDescription: "ID of the location",
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "Date and time of last update",
				Computed:            true,
			},
			"usage_type": schema.StringAttribute{
				MarkdownDescription: "Usage type of the asset type",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"user_id": schema.Int64Attribute{
				MarkdownDescription: "ID of the user",
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

// Configure configures the resource.
func (r *AssetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*freshclient.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"invalid provider data",
			"the provider data was not the expected type",
		)
		return
	}

	r.client = client
}

// Create the resource.
func (r *AssetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AssetResourceModel

	// Read the resource data from Terraform into the data model.
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	newAssetDetail := freshclient.AssetDetails{
		Name:        data.Name.ValueString(),
		AssetTypeID: data.AssetTypeID.ValueInt64(),
		Description: data.Description.ValueString(),
	}

	// Create the resource.
	var assetDetails *freshclient.AssetDetails
	if r.client != nil {
		var err error
		assetDetails, err = r.client.CreateAsset(newAssetDetail)

		if err != nil {
			resp.Diagnostics.AddError("Error creating asset", err.Error())
			return
		}
	} else {
		resp.Diagnostics.AddError("Error creating asset", "Client is nil")
		return
	}

	// Save data into Terraform state
	data = data.fromFreshAsset(*assetDetails)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read the resource and convert it into a resource object.
func (d *AssetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AssetDataSourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, strconv.FormatInt(data.DisplayID.ValueInt64(), 10))
	assetDetails, err := d.client.GetAsset(data.DisplayID.ValueInt64())

	if err != nil {
		resp.Diagnostics.AddError("Error getting asset", err.Error())
		return
	}

	// Save data into Terraform state
	data = data.fromFreshAsset(*assetDetails)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update the resource.
func (r *AssetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AssetResourceModel

	// Read the resource data from Terraform into the data model.
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	data_byes, _ := json.Marshal(data)
	log.Println(string(data_byes))
	fresh_byes, _ := json.Marshal(data.toFreshAsset())
	log.Println(string(fresh_byes))

	// Create the resource.
	assetDetails, err := r.client.UpdateAsset(data.toFreshAsset())
	log.Println("DOINGUPDATE")
	log.Println(json.Marshal(assetDetails))
	if err != nil {
		resp.Diagnostics.AddError("Error updating asset", err.Error())
		return
	}

	// Save data into Terraform state
	data = data.fromFreshAsset(*assetDetails)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete the resource.
func (r *AssetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AssetResourceModel

	// Read the resource data from Terraform into the data model.
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create the resource.
	err := r.client.DeleteAsset(data.toFreshAsset())

	if err != nil {
		resp.Diagnostics.AddError("Error deleting asset", err.Error())
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AssetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("display_id"), req, resp)
}
