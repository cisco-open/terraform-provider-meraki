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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksTrafficShapingDscpTaggingOptionsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksTrafficShapingDscpTaggingOptionsDataSource{}
)

func NewNetworksTrafficShapingDscpTaggingOptionsDataSource() datasource.DataSource {
	return &NetworksTrafficShapingDscpTaggingOptionsDataSource{}
}

type NetworksTrafficShapingDscpTaggingOptionsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksTrafficShapingDscpTaggingOptionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksTrafficShapingDscpTaggingOptionsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_traffic_shaping_dscp_tagging_options"
}

func (d *NetworksTrafficShapingDscpTaggingOptionsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseNetworksGetNetworkTrafficShapingDscpTaggingOptions`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"description": schema.StringAttribute{
							Computed: true,
						},
						"dscp_tag_value": schema.Int64Attribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksTrafficShapingDscpTaggingOptionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksTrafficShapingDscpTaggingOptions NetworksTrafficShapingDscpTaggingOptions
	diags := req.Config.Get(ctx, &networksTrafficShapingDscpTaggingOptions)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkTrafficShapingDscpTaggingOptions")
		vvNetworkID := networksTrafficShapingDscpTaggingOptions.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkTrafficShapingDscpTaggingOptions(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkTrafficShapingDscpTaggingOptions",
				err.Error(),
			)
			return
		}

		networksTrafficShapingDscpTaggingOptions = ResponseNetworksGetNetworkTrafficShapingDscpTaggingOptionsItemsToBody(networksTrafficShapingDscpTaggingOptions, response1)
		diags = resp.State.Set(ctx, &networksTrafficShapingDscpTaggingOptions)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksTrafficShapingDscpTaggingOptions struct {
	NetworkID types.String                                                      `tfsdk:"network_id"`
	Items     *[]ResponseItemNetworksGetNetworkTrafficShapingDscpTaggingOptions `tfsdk:"items"`
}

type ResponseItemNetworksGetNetworkTrafficShapingDscpTaggingOptions struct {
	Description  types.String `tfsdk:"description"`
	DscpTagValue types.Int64  `tfsdk:"dscp_tag_value"`
}

// ToBody
func ResponseNetworksGetNetworkTrafficShapingDscpTaggingOptionsItemsToBody(state NetworksTrafficShapingDscpTaggingOptions, response *merakigosdk.ResponseNetworksGetNetworkTrafficShapingDscpTaggingOptions) NetworksTrafficShapingDscpTaggingOptions {
	var items []ResponseItemNetworksGetNetworkTrafficShapingDscpTaggingOptions
	for _, item := range *response {
		itemState := ResponseItemNetworksGetNetworkTrafficShapingDscpTaggingOptions{
			Description: types.StringValue(item.Description),
			DscpTagValue: func() types.Int64 {
				if item.DscpTagValue != nil {
					return types.Int64Value(int64(*item.DscpTagValue))
				}
				return types.Int64{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
