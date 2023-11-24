// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"
	"terraform-provider-fresh/internal/freshclient"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure ScaffoldingProvider satisfies various provider interfaces.
var _ provider.Provider = &FreshProvider{}

// FreshProvider defines the provider implementation.
type FreshProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// FreshProviderModel describes the provider data model.
type FreshProviderModel struct {
	Address types.String `tfsdk:"address"`
	ApiKey  types.String `tfsdk:"api_key"`
}

func (p *FreshProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "fresh"
	resp.Version = p.version
}

func (p *FreshProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"address": schema.StringAttribute{
				Description: "Address for fresh",
				Optional:    true,
			},
			"api_key": schema.StringAttribute{
				Description: "API Key for fresh",
				Optional:    true,
				Sensitive:   true,
			},
		},
	}
}

func (p *FreshProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Fresh provider")

	// Check environment variables
	address := os.Getenv("FRESH_ADDRESS")
	apiKey := os.Getenv("FRESH_API_KEY")

	var data FreshProviderModel

	// Retrieve provider data from configuration
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write data to file

	if data.Address.ValueString() != "" {
		address = data.Address.ValueString()
	}

	if data.ApiKey.ValueString() != "" {
		apiKey = data.ApiKey.ValueString()
	}

	if apiKey == "" {
		resp.Diagnostics.AddError("api_key is required", "Details")
		return
	}

	if address == "" {
		resp.Diagnostics.AddError("address is required", "Exmaple address: https://example.freshservice.com, cannot use CNAMES")
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
	// Example client configuration for data sources and resources
	client := freshclient.NewClient(apiKey, address)
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *FreshProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAssetResource,
	}
}

func (p *FreshProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAssetTypeDataSource,
		NewAssetDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &FreshProvider{
			version: version,
		}
	}
}
