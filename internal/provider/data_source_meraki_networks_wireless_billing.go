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
	_ datasource.DataSource              = &NetworksWirelessBillingDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessBillingDataSource{}
)

func NewNetworksWirelessBillingDataSource() datasource.DataSource {
	return &NetworksWirelessBillingDataSource{}
}

type NetworksWirelessBillingDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessBillingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessBillingDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_billing"
}

func (d *NetworksWirelessBillingDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"currency": schema.StringAttribute{
						Computed: true,
					},
					"plans": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"bandwidth_limits": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"limit_down": schema.Int64Attribute{
											Computed: true,
										},
										"limit_up": schema.Int64Attribute{
											Computed: true,
										},
									},
								},
								"id": schema.StringAttribute{
									Computed: true,
								},
								"price": schema.Int64Attribute{
									Computed: true,
								},
								"time_limit": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessBillingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessBilling NetworksWirelessBilling
	diags := req.Config.Get(ctx, &networksWirelessBilling)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessBilling")
		vvNetworkID := networksWirelessBilling.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessBilling(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessBilling",
				err.Error(),
			)
			return
		}

		networksWirelessBilling = ResponseWirelessGetNetworkWirelessBillingItemToBody(networksWirelessBilling, response1)
		diags = resp.State.Set(ctx, &networksWirelessBilling)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessBilling struct {
	NetworkID types.String                               `tfsdk:"network_id"`
	Item      *ResponseWirelessGetNetworkWirelessBilling `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessBilling struct {
	Currency types.String                                      `tfsdk:"currency"`
	Plans    *[]ResponseWirelessGetNetworkWirelessBillingPlans `tfsdk:"plans"`
}

type ResponseWirelessGetNetworkWirelessBillingPlans struct {
	BandwidthLimits *ResponseWirelessGetNetworkWirelessBillingPlansBandwidthLimits `tfsdk:"bandwidth_limits"`
	ID              types.String                                                   `tfsdk:"id"`
	Price           types.Int64                                                    `tfsdk:"price"`
	TimeLimit       types.String                                                   `tfsdk:"time_limit"`
}

type ResponseWirelessGetNetworkWirelessBillingPlansBandwidthLimits struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessBillingItemToBody(state NetworksWirelessBilling, response *merakigosdk.ResponseWirelessGetNetworkWirelessBilling) NetworksWirelessBilling {
	itemState := ResponseWirelessGetNetworkWirelessBilling{
		Currency: types.StringValue(response.Currency),
		Plans: func() *[]ResponseWirelessGetNetworkWirelessBillingPlans {
			if response.Plans != nil {
				result := make([]ResponseWirelessGetNetworkWirelessBillingPlans, len(*response.Plans))
				for i, plans := range *response.Plans {
					result[i] = ResponseWirelessGetNetworkWirelessBillingPlans{
						BandwidthLimits: func() *ResponseWirelessGetNetworkWirelessBillingPlansBandwidthLimits {
							if plans.BandwidthLimits != nil {
								return &ResponseWirelessGetNetworkWirelessBillingPlansBandwidthLimits{
									LimitDown: func() types.Int64 {
										if plans.BandwidthLimits.LimitDown != nil {
											return types.Int64Value(int64(*plans.BandwidthLimits.LimitDown))
										}
										return types.Int64{}
									}(),
									LimitUp: func() types.Int64 {
										if plans.BandwidthLimits.LimitUp != nil {
											return types.Int64Value(int64(*plans.BandwidthLimits.LimitUp))
										}
										return types.Int64{}
									}(),
								}
							}
							return &ResponseWirelessGetNetworkWirelessBillingPlansBandwidthLimits{}
						}(),
						ID: types.StringValue(plans.ID),
						Price: func() types.Int64 {
							if plans.Price != nil {
								return types.Int64Value(int64(*plans.Price))
							}
							return types.Int64{}
						}(),
						TimeLimit: types.StringValue(plans.TimeLimit),
					}
				}
				return &result
			}
			return &[]ResponseWirelessGetNetworkWirelessBillingPlans{}
		}(),
	}
	state.Item = &itemState
	return state
}
