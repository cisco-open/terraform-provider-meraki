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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksSwitchDscpToCosMappingsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchDscpToCosMappingsDataSource{}
)

func NewNetworksSwitchDscpToCosMappingsDataSource() datasource.DataSource {
	return &NetworksSwitchDscpToCosMappingsDataSource{}
}

type NetworksSwitchDscpToCosMappingsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchDscpToCosMappingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchDscpToCosMappingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_dscp_to_cos_mappings"
}

func (d *NetworksSwitchDscpToCosMappingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"mappings": schema.SetNestedAttribute{
						MarkdownDescription: `An array of DSCP to CoS mappings. An empty array will reset the mappings to default.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"cos": schema.Int64Attribute{
									MarkdownDescription: `The actual layer-2 CoS queue the DSCP value is mapped to. These are not bits set on outgoing frames. Value can be in the range of 0 to 5 inclusive.`,
									Computed:            true,
								},
								"dscp": schema.Int64Attribute{
									MarkdownDescription: `The Differentiated Services Code Point (DSCP) tag in the IP header that will be mapped to a particular Class-of-Service (CoS) queue. Value can be in the range of 0 to 63 inclusive.`,
									Computed:            true,
								},
								"title": schema.StringAttribute{
									MarkdownDescription: `Label for the mapping (optional).`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchDscpToCosMappingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchDscpToCosMappings NetworksSwitchDscpToCosMappings
	diags := req.Config.Get(ctx, &networksSwitchDscpToCosMappings)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchDscpToCosMappings")
		vvNetworkID := networksSwitchDscpToCosMappings.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchDscpToCosMappings(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchDscpToCosMappings",
				err.Error(),
			)
			return
		}

		networksSwitchDscpToCosMappings = ResponseSwitchGetNetworkSwitchDscpToCosMappingsItemToBody(networksSwitchDscpToCosMappings, response1)
		diags = resp.State.Set(ctx, &networksSwitchDscpToCosMappings)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchDscpToCosMappings struct {
	NetworkID types.String                                     `tfsdk:"network_id"`
	Item      *ResponseSwitchGetNetworkSwitchDscpToCosMappings `tfsdk:"item"`
}

type ResponseSwitchGetNetworkSwitchDscpToCosMappings struct {
	Mappings *[]ResponseSwitchGetNetworkSwitchDscpToCosMappingsMappings `tfsdk:"mappings"`
}

type ResponseSwitchGetNetworkSwitchDscpToCosMappingsMappings struct {
	Cos   types.Int64  `tfsdk:"cos"`
	Dscp  types.Int64  `tfsdk:"dscp"`
	Title types.String `tfsdk:"title"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchDscpToCosMappingsItemToBody(state NetworksSwitchDscpToCosMappings, response *merakigosdk.ResponseSwitchGetNetworkSwitchDscpToCosMappings) NetworksSwitchDscpToCosMappings {
	itemState := ResponseSwitchGetNetworkSwitchDscpToCosMappings{
		Mappings: func() *[]ResponseSwitchGetNetworkSwitchDscpToCosMappingsMappings {
			if response.Mappings != nil {
				result := make([]ResponseSwitchGetNetworkSwitchDscpToCosMappingsMappings, len(*response.Mappings))
				for i, mappings := range *response.Mappings {
					result[i] = ResponseSwitchGetNetworkSwitchDscpToCosMappingsMappings{
						Cos: func() types.Int64 {
							if mappings.Cos != nil {
								return types.Int64Value(int64(*mappings.Cos))
							}
							return types.Int64{}
						}(),
						Dscp: func() types.Int64 {
							if mappings.Dscp != nil {
								return types.Int64Value(int64(*mappings.Dscp))
							}
							return types.Int64{}
						}(),
						Title: types.StringValue(mappings.Title),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
