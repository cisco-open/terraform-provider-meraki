// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// DATA SOURCE NORMAL
import (
	"context"
	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksSmProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSmProfilesDataSource{}
)

func NewNetworksSmProfilesDataSource() datasource.DataSource {
	return &NetworksSmProfilesDataSource{}
}

type NetworksSmProfilesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSmProfilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSmProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_profiles"
}

func (d *NetworksSmProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"payload_types": schema.ListAttribute{
				MarkdownDescription: `payloadTypes query parameter. Filter by payload types`,
				Optional:            true,
				ElementType:         types.StringType,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSmGetNetworkSmProfiles`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"description": schema.StringAttribute{
							MarkdownDescription: `Description of a profile.`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `ID of a profile.`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of a profile.`,
							Computed:            true,
						},
						"payload_types": schema.ListAttribute{
							MarkdownDescription: `Payloads in the profile.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"scope": schema.StringAttribute{
							MarkdownDescription: `Scope of a profile.`,
							Computed:            true,
						},
						"tags": schema.ListAttribute{
							MarkdownDescription: `Tags of a profile.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSmProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSmProfiles NetworksSmProfiles
	diags := req.Config.Get(ctx, &networksSmProfiles)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmProfiles")
		vvNetworkID := networksSmProfiles.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkSmProfilesQueryParams{}

		queryParams1.PayloadTypes = elementsToStrings(ctx, networksSmProfiles.PayloadTypes)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sm.GetNetworkSmProfiles(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmProfiles",
				err.Error(),
			)
			return
		}

		networksSmProfiles = ResponseSmGetNetworkSmProfilesItemsToBody(networksSmProfiles, response1)
		diags = resp.State.Set(ctx, &networksSmProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSmProfiles struct {
	NetworkID    types.String                          `tfsdk:"network_id"`
	PayloadTypes types.List                            `tfsdk:"payload_types"`
	Items        *[]ResponseItemSmGetNetworkSmProfiles `tfsdk:"items"`
}

type ResponseItemSmGetNetworkSmProfiles struct {
	Description  types.String `tfsdk:"description"`
	ID           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	PayloadTypes types.List   `tfsdk:"payload_types"`
	Scope        types.String `tfsdk:"scope"`
	Tags         types.List   `tfsdk:"tags"`
}

// ToBody
func ResponseSmGetNetworkSmProfilesItemsToBody(state NetworksSmProfiles, response *merakigosdk.ResponseSmGetNetworkSmProfiles) NetworksSmProfiles {
	var items []ResponseItemSmGetNetworkSmProfiles
	for _, item := range *response {
		itemState := ResponseItemSmGetNetworkSmProfiles{
			Description:  types.StringValue(item.Description),
			ID:           types.StringValue(item.ID),
			Name:         types.StringValue(item.Name),
			PayloadTypes: StringSliceToList(item.PayloadTypes),
			Scope:        types.StringValue(item.Scope),
			Tags:         StringSliceToList(item.Tags),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
