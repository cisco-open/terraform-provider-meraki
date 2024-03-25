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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksPiiRequestsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksPiiRequestsDataSource{}
)

func NewNetworksPiiRequestsDataSource() datasource.DataSource {
	return &NetworksPiiRequestsDataSource{}
}

type NetworksPiiRequestsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksPiiRequestsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksPiiRequestsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_pii_requests"
}

func (d *NetworksPiiRequestsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"request_id": schema.StringAttribute{
				MarkdownDescription: `requestId path parameter. Request ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"completed_at": schema.Int64Attribute{
						Computed: true,
					},
					"created_at": schema.Int64Attribute{
						Computed: true,
					},
					"datasets": schema.StringAttribute{
						Computed: true,
					},
					"id": schema.StringAttribute{
						Computed: true,
					},
					"mac": schema.StringAttribute{
						Computed: true,
					},
					"network_id": schema.StringAttribute{
						Computed: true,
					},
					"organization_wide": schema.BoolAttribute{
						Computed: true,
					},
					"status": schema.StringAttribute{
						Computed: true,
					},
					"type": schema.StringAttribute{
						Computed: true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseNetworksGetNetworkPiiRequests`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"completed_at": schema.Int64Attribute{
							Computed: true,
						},
						"created_at": schema.Int64Attribute{
							Computed: true,
						},
						"datasets": schema.StringAttribute{
							Computed: true,
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"mac": schema.StringAttribute{
							Computed: true,
						},
						"network_id": schema.StringAttribute{
							Computed: true,
						},
						"organization_wide": schema.BoolAttribute{
							Computed: true,
						},
						"status": schema.StringAttribute{
							Computed: true,
						},
						"type": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksPiiRequestsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksPiiRequests NetworksPiiRequests
	diags := req.Config.Get(ctx, &networksPiiRequests)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksPiiRequests.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksPiiRequests.NetworkID.IsNull(), !networksPiiRequests.RequestID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkPiiRequests")
		vvNetworkID := networksPiiRequests.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Networks.GetNetworkPiiRequests(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkPiiRequests",
				err.Error(),
			)
			return
		}

		networksPiiRequests = ResponseNetworksGetNetworkPiiRequestsItemsToBody(networksPiiRequests, response1)
		diags = resp.State.Set(ctx, &networksPiiRequests)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkPiiRequest")
		vvNetworkID := networksPiiRequests.NetworkID.ValueString()
		vvRequestID := networksPiiRequests.RequestID.ValueString()

		response2, restyResp2, err := d.client.Networks.GetNetworkPiiRequest(vvNetworkID, vvRequestID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkPiiRequest",
				err.Error(),
			)
			return
		}

		networksPiiRequests = ResponseNetworksGetNetworkPiiRequestItemToBody(networksPiiRequests, response2)
		diags = resp.State.Set(ctx, &networksPiiRequests)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksPiiRequests struct {
	NetworkID types.String                                 `tfsdk:"network_id"`
	RequestID types.String                                 `tfsdk:"request_id"`
	Items     *[]ResponseItemNetworksGetNetworkPiiRequests `tfsdk:"items"`
	Item      *ResponseNetworksGetNetworkPiiRequest        `tfsdk:"item"`
}

type ResponseItemNetworksGetNetworkPiiRequests struct {
	CompletedAt      types.Int64  `tfsdk:"completed_at"`
	CreatedAt        types.Int64  `tfsdk:"created_at"`
	Datasets         types.String `tfsdk:"datasets"`
	ID               types.String `tfsdk:"id"`
	Mac              types.String `tfsdk:"mac"`
	NetworkID        types.String `tfsdk:"network_id"`
	OrganizationWide types.Bool   `tfsdk:"organization_wide"`
	Status           types.String `tfsdk:"status"`
	Type             types.String `tfsdk:"type"`
}

type ResponseNetworksGetNetworkPiiRequest struct {
	CompletedAt      types.Int64  `tfsdk:"completed_at"`
	CreatedAt        types.Int64  `tfsdk:"created_at"`
	Datasets         types.String `tfsdk:"datasets"`
	ID               types.String `tfsdk:"id"`
	Mac              types.String `tfsdk:"mac"`
	NetworkID        types.String `tfsdk:"network_id"`
	OrganizationWide types.Bool   `tfsdk:"organization_wide"`
	Status           types.String `tfsdk:"status"`
	Type             types.String `tfsdk:"type"`
}

// ToBody
func ResponseNetworksGetNetworkPiiRequestsItemsToBody(state NetworksPiiRequests, response *merakigosdk.ResponseNetworksGetNetworkPiiRequests) NetworksPiiRequests {
	var items []ResponseItemNetworksGetNetworkPiiRequests
	for _, item := range *response {
		itemState := ResponseItemNetworksGetNetworkPiiRequests{
			CompletedAt: func() types.Int64 {
				if item.CompletedAt != nil {
					return types.Int64Value(int64(*item.CompletedAt))
				}
				return types.Int64{}
			}(),
			CreatedAt: func() types.Int64 {
				if item.CreatedAt != nil {
					return types.Int64Value(int64(*item.CreatedAt))
				}
				return types.Int64{}
			}(),
			Datasets:  types.StringValue(item.Datasets),
			ID:        types.StringValue(item.ID),
			Mac:       types.StringValue(item.Mac),
			NetworkID: types.StringValue(item.NetworkID),
			OrganizationWide: func() types.Bool {
				if item.OrganizationWide != nil {
					return types.BoolValue(*item.OrganizationWide)
				}
				return types.Bool{}
			}(),
			Status: types.StringValue(item.Status),
			Type:   types.StringValue(item.Type),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseNetworksGetNetworkPiiRequestItemToBody(state NetworksPiiRequests, response *merakigosdk.ResponseNetworksGetNetworkPiiRequest) NetworksPiiRequests {
	itemState := ResponseNetworksGetNetworkPiiRequest{
		CompletedAt: func() types.Int64 {
			if response.CompletedAt != nil {
				return types.Int64Value(int64(*response.CompletedAt))
			}
			return types.Int64{}
		}(),
		CreatedAt: func() types.Int64 {
			if response.CreatedAt != nil {
				return types.Int64Value(int64(*response.CreatedAt))
			}
			return types.Int64{}
		}(),
		Datasets:  types.StringValue(response.Datasets),
		ID:        types.StringValue(response.ID),
		Mac:       types.StringValue(response.Mac),
		NetworkID: types.StringValue(response.NetworkID),
		OrganizationWide: func() types.Bool {
			if response.OrganizationWide != nil {
				return types.BoolValue(*response.OrganizationWide)
			}
			return types.Bool{}
		}(),
		Status: types.StringValue(response.Status),
		Type:   types.StringValue(response.Type),
	}
	state.Item = &itemState
	return state
}
