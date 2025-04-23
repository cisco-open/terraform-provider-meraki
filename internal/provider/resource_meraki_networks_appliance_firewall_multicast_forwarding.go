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

// RESOURCE ACTION

import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceFirewallMulticastForwardingResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceFirewallMulticastForwardingResource{}
)

func NewNetworksApplianceFirewallMulticastForwardingResource() resource.Resource {
	return &NetworksApplianceFirewallMulticastForwardingResource{}
}

type NetworksApplianceFirewallMulticastForwardingResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceFirewallMulticastForwardingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceFirewallMulticastForwardingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_multicast_forwarding"
}

// resourceAction
func (r *NetworksApplianceFirewallMulticastForwardingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"network": schema.SingleNestedAttribute{
						MarkdownDescription: `Network details`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"id": schema.StringAttribute{
								MarkdownDescription: `ID of the network whose multicast forwarding settings are returned.`,
								Computed:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Name of the network whose multicast forwarding settings are returned.`,
								Computed:            true,
							},
						},
					},
					"rules": schema.SetNestedAttribute{
						MarkdownDescription: `Static multicast forwarding rules.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"address": schema.StringAttribute{
									MarkdownDescription: `IP address`,
									Computed:            true,
								},
								"description": schema.StringAttribute{
									MarkdownDescription: `Forwarding rule description.`,
									Computed:            true,
								},
								"vlan_ids": schema.ListAttribute{
									MarkdownDescription: `List of VLAN IDs`,
									Computed:            true,
									ElementType:         types.StringType,
								},
							},
						},
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"rules": schema.SetNestedAttribute{
						MarkdownDescription: `Static multicast forwarding rules. Pass an empty array to clear all rules.`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"address": schema.StringAttribute{
									MarkdownDescription: `IP address`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"description": schema.StringAttribute{
									MarkdownDescription: `Forwarding rule description.`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"vlan_ids": schema.ListAttribute{
									MarkdownDescription: `List of VLAN IDs`,
									Optional:            true,
									Computed:            true,
									ElementType:         types.StringType,
								},
							},
						},
					},
				},
			},
		},
	}
}
func (r *NetworksApplianceFirewallMulticastForwardingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceFirewallMulticastForwarding

	var item types.Object
	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp1, err := r.client.Appliance.UpdateNetworkApplianceFirewallMulticastForwarding(vvNetworkID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceFirewallMulticastForwarding",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceFirewallMulticastForwarding",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwardingItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceFirewallMulticastForwardingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksApplianceFirewallMulticastForwardingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksApplianceFirewallMulticastForwardingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceFirewallMulticastForwarding struct {
	NetworkID  types.String                                                         `tfsdk:"network_id"`
	Item       *ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwarding  `tfsdk:"item"`
	Parameters *RequestApplianceUpdateNetworkApplianceFirewallMulticastForwardingRs `tfsdk:"parameters"`
}

type ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwarding struct {
	Network *ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwardingNetwork `tfsdk:"network"`
	Rules   *[]ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwardingRules `tfsdk:"rules"`
}

type ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwardingNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwardingRules struct {
	Address     types.String `tfsdk:"address"`
	Description types.String `tfsdk:"description"`
	VLANIDs     types.List   `tfsdk:"vlan_ids"`
}

type RequestApplianceUpdateNetworkApplianceFirewallMulticastForwardingRs struct {
	Rules *[]RequestApplianceUpdateNetworkApplianceFirewallMulticastForwardingRulesRs `tfsdk:"rules"`
}

type RequestApplianceUpdateNetworkApplianceFirewallMulticastForwardingRulesRs struct {
	Address     types.String `tfsdk:"address"`
	Description types.String `tfsdk:"description"`
	VLANIDs     types.Set    `tfsdk:"vlan_ids"`
}

// FromBody
func (r *NetworksApplianceFirewallMulticastForwarding) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallMulticastForwarding {
	re := *r.Parameters
	var requestApplianceUpdateNetworkApplianceFirewallMulticastForwardingRules []merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallMulticastForwardingRules

	if re.Rules != nil {
		for _, rItem1 := range *re.Rules {
			address := rItem1.Address.ValueString()
			description := rItem1.Description.ValueString()

			var vlanIDs []string = nil
			rItem1.VLANIDs.ElementsAs(ctx, &vlanIDs, false)
			requestApplianceUpdateNetworkApplianceFirewallMulticastForwardingRules = append(requestApplianceUpdateNetworkApplianceFirewallMulticastForwardingRules, merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallMulticastForwardingRules{
				Address:     address,
				Description: description,
				VLANIDs:     vlanIDs,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallMulticastForwarding{
		Rules: &requestApplianceUpdateNetworkApplianceFirewallMulticastForwardingRules,
	}
	return &out
}

// ToBody
func ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwardingItemToBody(state NetworksApplianceFirewallMulticastForwarding, response *merakigosdk.ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwarding) NetworksApplianceFirewallMulticastForwarding {
	itemState := ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwarding{
		Network: func() *ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwardingNetwork {
			if response.Network != nil {
				return &ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwardingNetwork{
					ID:   types.StringValue(response.Network.ID),
					Name: types.StringValue(response.Network.Name),
				}
			}
			return nil
		}(),
		Rules: func() *[]ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwardingRules {
			if response.Rules != nil {
				result := make([]ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwardingRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseApplianceUpdateNetworkApplianceFirewallMulticastForwardingRules{
						Address:     types.StringValue(rules.Address),
						Description: types.StringValue(rules.Description),
						VLANIDs:     StringSliceToList(rules.VLANIDs),
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
