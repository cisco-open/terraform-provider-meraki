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
	_ datasource.DataSource              = &NetworksSwitchStacksDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchStacksDataSource{}
)

func NewNetworksSwitchStacksDataSource() datasource.DataSource {
	return &NetworksSwitchStacksDataSource{}
}

type NetworksSwitchStacksDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchStacksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchStacksDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_stacks"
}

func (d *NetworksSwitchStacksDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"switch_stack_id": schema.StringAttribute{
				MarkdownDescription: `switchStackId path parameter. Switch stack ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"id": schema.StringAttribute{
						MarkdownDescription: `ID of the Switch stack`,
						Computed:            true,
					},
					"is_monitor_only": schema.BoolAttribute{
						MarkdownDescription: `Tells if stack is Monitored Stack.`,
						Computed:            true,
					},
					"members": schema.SetNestedAttribute{
						MarkdownDescription: `Members of the Stack`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"mac": schema.StringAttribute{
									MarkdownDescription: `MAC address of the device`,
									Computed:            true,
								},
								"model": schema.StringAttribute{
									MarkdownDescription: `Model of the device`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name of the device`,
									Computed:            true,
								},
								"role": schema.StringAttribute{
									MarkdownDescription: `Role of the device`,
									Computed:            true,
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `Serial number of the device`,
									Computed:            true,
								},
							},
						},
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of the Switch stack`,
						Computed:            true,
					},
					"serials": schema.ListAttribute{
						MarkdownDescription: `Serials of the switches in the switch stack`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSwitchGetNetworkSwitchStacks`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"id": schema.StringAttribute{
							MarkdownDescription: `ID of the Switch stack`,
							Computed:            true,
						},
						"is_monitor_only": schema.BoolAttribute{
							MarkdownDescription: `Tells if stack is Monitored Stack.`,
							Computed:            true,
						},
						"members": schema.SetNestedAttribute{
							MarkdownDescription: `Members of the Stack`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"mac": schema.StringAttribute{
										MarkdownDescription: `MAC address of the device`,
										Computed:            true,
									},
									"model": schema.StringAttribute{
										MarkdownDescription: `Model of the device`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Name of the device`,
										Computed:            true,
									},
									"role": schema.StringAttribute{
										MarkdownDescription: `Role of the device`,
										Computed:            true,
									},
									"serial": schema.StringAttribute{
										MarkdownDescription: `Serial number of the device`,
										Computed:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the Switch stack`,
							Computed:            true,
						},
						"serials": schema.ListAttribute{
							MarkdownDescription: `Serials of the switches in the switch stack`,
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchStacksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchStacks NetworksSwitchStacks
	diags := req.Config.Get(ctx, &networksSwitchStacks)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksSwitchStacks.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksSwitchStacks.NetworkID.IsNull(), !networksSwitchStacks.SwitchStackID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchStacks")
		vvNetworkID := networksSwitchStacks.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchStacks(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStacks",
				err.Error(),
			)
			return
		}

		networksSwitchStacks = ResponseSwitchGetNetworkSwitchStacksItemsToBody(networksSwitchStacks, response1)
		diags = resp.State.Set(ctx, &networksSwitchStacks)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchStack")
		vvNetworkID := networksSwitchStacks.NetworkID.ValueString()
		vvSwitchStackID := networksSwitchStacks.SwitchStackID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Switch.GetNetworkSwitchStack(vvNetworkID, vvSwitchStackID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStack",
				err.Error(),
			)
			return
		}

		networksSwitchStacks = ResponseSwitchGetNetworkSwitchStackItemToBody(networksSwitchStacks, response2)
		diags = resp.State.Set(ctx, &networksSwitchStacks)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchStacks struct {
	NetworkID     types.String                                `tfsdk:"network_id"`
	SwitchStackID types.String                                `tfsdk:"switch_stack_id"`
	Items         *[]ResponseItemSwitchGetNetworkSwitchStacks `tfsdk:"items"`
	Item          *ResponseSwitchGetNetworkSwitchStack        `tfsdk:"item"`
}

type ResponseItemSwitchGetNetworkSwitchStacks struct {
	ID            types.String                                       `tfsdk:"id"`
	IsMonitorOnly types.Bool                                         `tfsdk:"is_monitor_only"`
	Members       *[]ResponseItemSwitchGetNetworkSwitchStacksMembers `tfsdk:"members"`
	Name          types.String                                       `tfsdk:"name"`
	Serials       types.List                                         `tfsdk:"serials"`
}

type ResponseItemSwitchGetNetworkSwitchStacksMembers struct {
	Mac    types.String `tfsdk:"mac"`
	Model  types.String `tfsdk:"model"`
	Name   types.String `tfsdk:"name"`
	Role   types.String `tfsdk:"role"`
	Serial types.String `tfsdk:"serial"`
}

type ResponseSwitchGetNetworkSwitchStack struct {
	ID            types.String                                  `tfsdk:"id"`
	IsMonitorOnly types.Bool                                    `tfsdk:"is_monitor_only"`
	Members       *[]ResponseSwitchGetNetworkSwitchStackMembers `tfsdk:"members"`
	Name          types.String                                  `tfsdk:"name"`
	Serials       types.List                                    `tfsdk:"serials"`
}

type ResponseSwitchGetNetworkSwitchStackMembers struct {
	Mac    types.String `tfsdk:"mac"`
	Model  types.String `tfsdk:"model"`
	Name   types.String `tfsdk:"name"`
	Role   types.String `tfsdk:"role"`
	Serial types.String `tfsdk:"serial"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchStacksItemsToBody(state NetworksSwitchStacks, response *merakigosdk.ResponseSwitchGetNetworkSwitchStacks) NetworksSwitchStacks {
	var items []ResponseItemSwitchGetNetworkSwitchStacks
	for _, item := range *response {
		itemState := ResponseItemSwitchGetNetworkSwitchStacks{
			ID: types.StringValue(item.ID),
			IsMonitorOnly: func() types.Bool {
				if item.IsMonitorOnly != nil {
					return types.BoolValue(*item.IsMonitorOnly)
				}
				return types.Bool{}
			}(),
			Members: func() *[]ResponseItemSwitchGetNetworkSwitchStacksMembers {
				if item.Members != nil {
					result := make([]ResponseItemSwitchGetNetworkSwitchStacksMembers, len(*item.Members))
					for i, members := range *item.Members {
						result[i] = ResponseItemSwitchGetNetworkSwitchStacksMembers{
							Mac:    types.StringValue(members.Mac),
							Model:  types.StringValue(members.Model),
							Name:   types.StringValue(members.Name),
							Role:   types.StringValue(members.Role),
							Serial: types.StringValue(members.Serial),
						}
					}
					return &result
				}
				return nil
			}(),
			Name:    types.StringValue(item.Name),
			Serials: StringSliceToList(item.Serials),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseSwitchGetNetworkSwitchStackItemToBody(state NetworksSwitchStacks, response *merakigosdk.ResponseSwitchGetNetworkSwitchStack) NetworksSwitchStacks {
	itemState := ResponseSwitchGetNetworkSwitchStack{
		ID: types.StringValue(response.ID),
		IsMonitorOnly: func() types.Bool {
			if response.IsMonitorOnly != nil {
				return types.BoolValue(*response.IsMonitorOnly)
			}
			return types.Bool{}
		}(),
		Members: func() *[]ResponseSwitchGetNetworkSwitchStackMembers {
			if response.Members != nil {
				result := make([]ResponseSwitchGetNetworkSwitchStackMembers, len(*response.Members))
				for i, members := range *response.Members {
					result[i] = ResponseSwitchGetNetworkSwitchStackMembers{
						Mac:    types.StringValue(members.Mac),
						Model:  types.StringValue(members.Model),
						Name:   types.StringValue(members.Name),
						Role:   types.StringValue(members.Role),
						Serial: types.StringValue(members.Serial),
					}
				}
				return &result
			}
			return nil
		}(),
		Name:    types.StringValue(response.Name),
		Serials: StringSliceToList(response.Serials),
	}
	state.Item = &itemState
	return state
}
