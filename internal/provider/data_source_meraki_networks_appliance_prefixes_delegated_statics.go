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
	_ datasource.DataSource              = &NetworksAppliancePrefixesDelegatedStaticsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksAppliancePrefixesDelegatedStaticsDataSource{}
)

func NewNetworksAppliancePrefixesDelegatedStaticsDataSource() datasource.DataSource {
	return &NetworksAppliancePrefixesDelegatedStaticsDataSource{}
}

type NetworksAppliancePrefixesDelegatedStaticsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksAppliancePrefixesDelegatedStaticsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksAppliancePrefixesDelegatedStaticsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_prefixes_delegated_statics"
}

func (d *NetworksAppliancePrefixesDelegatedStaticsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"static_delegated_prefix_id": schema.StringAttribute{
				MarkdownDescription: `staticDelegatedPrefixId path parameter. Static delegated prefix ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"created_at": schema.StringAttribute{
						MarkdownDescription: `Prefix creation time.`,
						Computed:            true,
					},
					"description": schema.StringAttribute{
						MarkdownDescription: `Identifying description for the prefix.`,
						Computed:            true,
					},
					"origin": schema.SingleNestedAttribute{
						MarkdownDescription: `WAN1/WAN2/Independent prefix.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"interfaces": schema.ListAttribute{
								MarkdownDescription: `Uplink provided or independent`,
								Computed:            true,
								ElementType:         types.StringType,
							},
							"type": schema.StringAttribute{
								MarkdownDescription: `Origin type`,
								Computed:            true,
							},
						},
					},
					"prefix": schema.StringAttribute{
						MarkdownDescription: `IPv6 prefix/prefix length.`,
						Computed:            true,
					},
					"static_delegated_prefix_id": schema.StringAttribute{
						MarkdownDescription: `Static delegated prefix id.`,
						Computed:            true,
					},
					"updated_at": schema.StringAttribute{
						MarkdownDescription: `Prefix Updated time.`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatics`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"created_at": schema.StringAttribute{
							MarkdownDescription: `Prefix creation time.`,
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: `Identifying description for the prefix.`,
							Computed:            true,
						},
						"origin": schema.SingleNestedAttribute{
							MarkdownDescription: `WAN1/WAN2/Independent prefix.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"interfaces": schema.ListAttribute{
									MarkdownDescription: `Uplink provided or independent`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"type": schema.StringAttribute{
									MarkdownDescription: `Origin type`,
									Computed:            true,
								},
							},
						},
						"prefix": schema.StringAttribute{
							MarkdownDescription: `IPv6 prefix/prefix length.`,
							Computed:            true,
						},
						"static_delegated_prefix_id": schema.StringAttribute{
							MarkdownDescription: `Static delegated prefix id.`,
							Computed:            true,
						},
						"updated_at": schema.StringAttribute{
							MarkdownDescription: `Prefix Updated time.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksAppliancePrefixesDelegatedStaticsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksAppliancePrefixesDelegatedStatics NetworksAppliancePrefixesDelegatedStatics
	diags := req.Config.Get(ctx, &networksAppliancePrefixesDelegatedStatics)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksAppliancePrefixesDelegatedStatics.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksAppliancePrefixesDelegatedStatics.NetworkID.IsNull(), !networksAppliancePrefixesDelegatedStatics.StaticDelegatedPrefixID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkAppliancePrefixesDelegatedStatics")
		vvNetworkID := networksAppliancePrefixesDelegatedStatics.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetNetworkAppliancePrefixesDelegatedStatics(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkAppliancePrefixesDelegatedStatics",
				err.Error(),
			)
			return
		}

		networksAppliancePrefixesDelegatedStatics = ResponseApplianceGetNetworkAppliancePrefixesDelegatedStaticsItemsToBody(networksAppliancePrefixesDelegatedStatics, response1)
		diags = resp.State.Set(ctx, &networksAppliancePrefixesDelegatedStatics)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkAppliancePrefixesDelegatedStatic")
		vvNetworkID := networksAppliancePrefixesDelegatedStatics.NetworkID.ValueString()
		vvStaticDelegatedPrefixID := networksAppliancePrefixesDelegatedStatics.StaticDelegatedPrefixID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Appliance.GetNetworkAppliancePrefixesDelegatedStatic(vvNetworkID, vvStaticDelegatedPrefixID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkAppliancePrefixesDelegatedStatic",
				err.Error(),
			)
			return
		}

		networksAppliancePrefixesDelegatedStatics = ResponseApplianceGetNetworkAppliancePrefixesDelegatedStaticItemToBody(networksAppliancePrefixesDelegatedStatics, response2)
		diags = resp.State.Set(ctx, &networksAppliancePrefixesDelegatedStatics)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksAppliancePrefixesDelegatedStatics struct {
	NetworkID               types.String                                                        `tfsdk:"network_id"`
	StaticDelegatedPrefixID types.String                                                        `tfsdk:"static_delegated_prefix_id"`
	Items                   *[]ResponseItemApplianceGetNetworkAppliancePrefixesDelegatedStatics `tfsdk:"items"`
	Item                    *ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatic        `tfsdk:"item"`
}

type ResponseItemApplianceGetNetworkAppliancePrefixesDelegatedStatics struct {
	CreatedAt               types.String                                                            `tfsdk:"created_at"`
	Description             types.String                                                            `tfsdk:"description"`
	Origin                  *ResponseItemApplianceGetNetworkAppliancePrefixesDelegatedStaticsOrigin `tfsdk:"origin"`
	Prefix                  types.String                                                            `tfsdk:"prefix"`
	StaticDelegatedPrefixID types.String                                                            `tfsdk:"static_delegated_prefix_id"`
	UpdatedAt               types.String                                                            `tfsdk:"updated_at"`
}

type ResponseItemApplianceGetNetworkAppliancePrefixesDelegatedStaticsOrigin struct {
	Interfaces types.List   `tfsdk:"interfaces"`
	Type       types.String `tfsdk:"type"`
}

type ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatic struct {
	CreatedAt               types.String                                                       `tfsdk:"created_at"`
	Description             types.String                                                       `tfsdk:"description"`
	Origin                  *ResponseApplianceGetNetworkAppliancePrefixesDelegatedStaticOrigin `tfsdk:"origin"`
	Prefix                  types.String                                                       `tfsdk:"prefix"`
	StaticDelegatedPrefixID types.String                                                       `tfsdk:"static_delegated_prefix_id"`
	UpdatedAt               types.String                                                       `tfsdk:"updated_at"`
}

type ResponseApplianceGetNetworkAppliancePrefixesDelegatedStaticOrigin struct {
	Interfaces types.List   `tfsdk:"interfaces"`
	Type       types.String `tfsdk:"type"`
}

// ToBody
func ResponseApplianceGetNetworkAppliancePrefixesDelegatedStaticsItemsToBody(state NetworksAppliancePrefixesDelegatedStatics, response *merakigosdk.ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatics) NetworksAppliancePrefixesDelegatedStatics {
	var items []ResponseItemApplianceGetNetworkAppliancePrefixesDelegatedStatics
	for _, item := range *response {
		itemState := ResponseItemApplianceGetNetworkAppliancePrefixesDelegatedStatics{
			CreatedAt:   types.StringValue(item.CreatedAt),
			Description: types.StringValue(item.Description),
			Origin: func() *ResponseItemApplianceGetNetworkAppliancePrefixesDelegatedStaticsOrigin {
				if item.Origin != nil {
					return &ResponseItemApplianceGetNetworkAppliancePrefixesDelegatedStaticsOrigin{
						Interfaces: StringSliceToList(item.Origin.Interfaces),
						Type:       types.StringValue(item.Origin.Type),
					}
				}
				return nil
			}(),
			Prefix:                  types.StringValue(item.Prefix),
			StaticDelegatedPrefixID: types.StringValue(item.StaticDelegatedPrefixID),
			UpdatedAt:               types.StringValue(item.UpdatedAt),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseApplianceGetNetworkAppliancePrefixesDelegatedStaticItemToBody(state NetworksAppliancePrefixesDelegatedStatics, response *merakigosdk.ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatic) NetworksAppliancePrefixesDelegatedStatics {
	itemState := ResponseApplianceGetNetworkAppliancePrefixesDelegatedStatic{
		CreatedAt:   types.StringValue(response.CreatedAt),
		Description: types.StringValue(response.Description),
		Origin: func() *ResponseApplianceGetNetworkAppliancePrefixesDelegatedStaticOrigin {
			if response.Origin != nil {
				return &ResponseApplianceGetNetworkAppliancePrefixesDelegatedStaticOrigin{
					Interfaces: StringSliceToList(response.Origin.Interfaces),
					Type:       types.StringValue(response.Origin.Type),
				}
			}
			return nil
		}(),
		Prefix:                  types.StringValue(response.Prefix),
		StaticDelegatedPrefixID: types.StringValue(response.StaticDelegatedPrefixID),
		UpdatedAt:               types.StringValue(response.UpdatedAt),
	}
	state.Item = &itemState
	return state
}
