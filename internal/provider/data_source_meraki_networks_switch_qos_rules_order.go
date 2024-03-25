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
	_ datasource.DataSource              = &NetworksSwitchQosRulesOrderDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchQosRulesOrderDataSource{}
)

func NewNetworksSwitchQosRulesOrderDataSource() datasource.DataSource {
	return &NetworksSwitchQosRulesOrderDataSource{}
}

type NetworksSwitchQosRulesOrderDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchQosRulesOrderDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchQosRulesOrderDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_qos_rules_order"
}

func (d *NetworksSwitchQosRulesOrderDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"qos_rule_id": schema.StringAttribute{
				MarkdownDescription: `qosRuleId path parameter. Qos rule ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"dscp": schema.Int64Attribute{
						Computed: true,
					},
					"dst_port": schema.Int64Attribute{
						Computed: true,
					},
					"dst_port_range": schema.StringAttribute{
						Computed: true,
					},
					"id": schema.StringAttribute{
						Computed: true,
					},
					"protocol": schema.StringAttribute{
						Computed: true,
					},
					"src_port": schema.Int64Attribute{
						Computed: true,
					},
					"src_port_range": schema.StringAttribute{
						Computed: true,
					},
					"vlan": schema.Int64Attribute{
						Computed: true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSwitchGetNetworkSwitchQosRules`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"dscp": schema.Int64Attribute{
							Computed: true,
						},
						"dst_port": schema.Int64Attribute{
							Computed: true,
						},
						"dst_port_range": schema.StringAttribute{
							Computed: true,
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"protocol": schema.StringAttribute{
							Computed: true,
						},
						"src_port": schema.Int64Attribute{
							Computed: true,
						},
						"src_port_range": schema.StringAttribute{
							Computed: true,
						},
						"vlan": schema.Int64Attribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchQosRulesOrderDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchQosRulesOrder NetworksSwitchQosRulesOrder
	diags := req.Config.Get(ctx, &networksSwitchQosRulesOrder)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksSwitchQosRulesOrder.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksSwitchQosRulesOrder.NetworkID.IsNull(), !networksSwitchQosRulesOrder.QosRuleID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchQosRules")
		vvNetworkID := networksSwitchQosRulesOrder.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchQosRules(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchQosRules",
				err.Error(),
			)
			return
		}

		networksSwitchQosRulesOrder = ResponseSwitchGetNetworkSwitchQosRulesItemsToBody(networksSwitchQosRulesOrder, response1)
		diags = resp.State.Set(ctx, &networksSwitchQosRulesOrder)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchQosRule")
		vvNetworkID := networksSwitchQosRulesOrder.NetworkID.ValueString()
		vvQosRuleID := networksSwitchQosRulesOrder.QosRuleID.ValueString()

		response2, restyResp2, err := d.client.Switch.GetNetworkSwitchQosRule(vvNetworkID, vvQosRuleID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchQosRule",
				err.Error(),
			)
			return
		}

		networksSwitchQosRulesOrder = ResponseSwitchGetNetworkSwitchQosRuleItemToBody(networksSwitchQosRulesOrder, response2)
		diags = resp.State.Set(ctx, &networksSwitchQosRulesOrder)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchQosRulesOrder struct {
	NetworkID types.String                                  `tfsdk:"network_id"`
	QosRuleID types.String                                  `tfsdk:"qos_rule_id"`
	Items     *[]ResponseItemSwitchGetNetworkSwitchQosRules `tfsdk:"items"`
	Item      *ResponseSwitchGetNetworkSwitchQosRule        `tfsdk:"item"`
}

type ResponseItemSwitchGetNetworkSwitchQosRules struct {
	Dscp         types.Int64  `tfsdk:"dscp"`
	DstPort      types.Int64  `tfsdk:"dst_port"`
	DstPortRange types.String `tfsdk:"dst_port_range"`
	ID           types.String `tfsdk:"id"`
	Protocol     types.String `tfsdk:"protocol"`
	SrcPort      types.Int64  `tfsdk:"src_port"`
	SrcPortRange types.String `tfsdk:"src_port_range"`
	VLAN         types.Int64  `tfsdk:"vlan"`
}

type ResponseSwitchGetNetworkSwitchQosRule struct {
	Dscp         types.Int64  `tfsdk:"dscp"`
	DstPort      types.Int64  `tfsdk:"dst_port"`
	DstPortRange types.String `tfsdk:"dst_port_range"`
	ID           types.String `tfsdk:"id"`
	Protocol     types.String `tfsdk:"protocol"`
	SrcPort      types.Int64  `tfsdk:"src_port"`
	SrcPortRange types.String `tfsdk:"src_port_range"`
	VLAN         types.Int64  `tfsdk:"vlan"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchQosRulesItemsToBody(state NetworksSwitchQosRulesOrder, response *merakigosdk.ResponseSwitchGetNetworkSwitchQosRules) NetworksSwitchQosRulesOrder {
	var items []ResponseItemSwitchGetNetworkSwitchQosRules
	for _, item := range *response {
		itemState := ResponseItemSwitchGetNetworkSwitchQosRules{
			Dscp: func() types.Int64 {
				if item.Dscp != nil {
					return types.Int64Value(int64(*item.Dscp))
				}
				return types.Int64{}
			}(),
			DstPort: func() types.Int64 {
				if item.DstPort != nil {
					return types.Int64Value(int64(*item.DstPort))
				}
				return types.Int64{}
			}(),
			DstPortRange: types.StringValue(item.DstPortRange),
			ID:           types.StringValue(item.ID),
			Protocol:     types.StringValue(item.Protocol),
			SrcPort: func() types.Int64 {
				if item.SrcPort != nil {
					return types.Int64Value(int64(*item.SrcPort))
				}
				return types.Int64{}
			}(),
			SrcPortRange: types.StringValue(item.SrcPortRange),
			VLAN: func() types.Int64 {
				if item.VLAN != nil {
					return types.Int64Value(int64(*item.VLAN))
				}
				return types.Int64{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseSwitchGetNetworkSwitchQosRuleItemToBody(state NetworksSwitchQosRulesOrder, response *merakigosdk.ResponseSwitchGetNetworkSwitchQosRule) NetworksSwitchQosRulesOrder {
	itemState := ResponseSwitchGetNetworkSwitchQosRule{
		Dscp: func() types.Int64 {
			if response.Dscp != nil {
				return types.Int64Value(int64(*response.Dscp))
			}
			return types.Int64{}
		}(),
		DstPort: func() types.Int64 {
			if response.DstPort != nil {
				return types.Int64Value(int64(*response.DstPort))
			}
			return types.Int64{}
		}(),
		DstPortRange: types.StringValue(response.DstPortRange),
		ID:           types.StringValue(response.ID),
		Protocol:     types.StringValue(response.Protocol),
		SrcPort: func() types.Int64 {
			if response.SrcPort != nil {
				return types.Int64Value(int64(*response.SrcPort))
			}
			return types.Int64{}
		}(),
		SrcPortRange: types.StringValue(response.SrcPortRange),
		VLAN: func() types.Int64 {
			if response.VLAN != nil {
				return types.Int64Value(int64(*response.VLAN))
			}
			return types.Int64{}
		}(),
	}
	state.Item = &itemState
	return state
}
