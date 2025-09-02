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

// RESOURCE NORMAL
import (
	"context"
	"strconv"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceFirewallPortForwardingRulesResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceFirewallPortForwardingRulesResource{}
)

func NewNetworksApplianceFirewallPortForwardingRulesResource() resource.Resource {
	return &NetworksApplianceFirewallPortForwardingRulesResource{}
}

type NetworksApplianceFirewallPortForwardingRulesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceFirewallPortForwardingRulesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceFirewallPortForwardingRulesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_port_forwarding_rules"
}

func (r *NetworksApplianceFirewallPortForwardingRulesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: `An array of port forwarding rules`,
				Optional:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"allowed_ips": schema.ListAttribute{
							MarkdownDescription: `An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges (or any)`,
							Optional:            true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
						},
						"lan_ip": schema.StringAttribute{
							MarkdownDescription: `IP address of the device subject to port forwarding`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"local_port": schema.StringAttribute{
							MarkdownDescription: `The port or port range that receives forwarded traffic from the WAN`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the rule`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: `Protocol the rule applies to
                                        Allowed values: [tcp,udp]`,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"tcp",
									"udp",
								),
							},
						},
						"public_port": schema.StringAttribute{
							MarkdownDescription: `The port or port range forwarded to the host on the LAN`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"uplink": schema.StringAttribute{
							MarkdownDescription: `The physical WAN interface on which the traffic arrives; allowed values vary by appliance model and configuration
                                        Allowed values: [both,internet1,internet2,internet3]`,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"both",
									"internet1",
									"internet2",
									"internet3",
								),
							},
						},
					},
				},
			},
		},
	}
}

func (r *NetworksApplianceFirewallPortForwardingRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceFirewallPortForwardingRulesRs

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
	// Has Paths
	vvNetworkID := data.NetworkID.ValueString()
	//Has Item and not has items

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceFirewallPortForwardingRules(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceFirewallPortForwardingRules",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceFirewallPortForwardingRules",
			err.Error(),
		)
		return
	}

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *NetworksApplianceFirewallPortForwardingRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceFirewallPortForwardingRulesRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceFirewallPortForwardingRules(vvNetworkID)
	if err != nil || restyRespGet == nil {
		if restyRespGet != nil {
			if restyRespGet.StatusCode() == 404 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceFirewallPortForwardingRules",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceFirewallPortForwardingRules",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesItemToBodyRs(data, responseGet, true)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksApplianceFirewallPortForwardingRulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceFirewallPortForwardingRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksApplianceFirewallPortForwardingRulesRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceFirewallPortForwardingRules(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceFirewallPortForwardingRules",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceFirewallPortForwardingRules",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworksApplianceFirewallPortForwardingRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceFirewallPortForwardingRules", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceFirewallPortForwardingRulesRs struct {
	NetworkID types.String                                                              `tfsdk:"network_id"`
	Rules     *[]ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRulesRs `tfsdk:"rules"`
}

type ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRulesRs struct {
	AllowedIPs types.List   `tfsdk:"allowed_ips"`
	LanIP      types.String `tfsdk:"lan_ip"`
	LocalPort  types.String `tfsdk:"local_port"`
	Name       types.String `tfsdk:"name"`
	Protocol   types.String `tfsdk:"protocol"`
	PublicPort types.String `tfsdk:"public_port"`
	Uplink     types.String `tfsdk:"uplink"`
}

// FromBody
func (r *NetworksApplianceFirewallPortForwardingRulesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallPortForwardingRules {
	var requestApplianceUpdateNetworkApplianceFirewallPortForwardingRulesRules []merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallPortForwardingRulesRules

	if r.Rules != nil {
		for _, rItem1 := range *r.Rules {

			var allowedIPs []string = nil
			rItem1.AllowedIPs.ElementsAs(ctx, &allowedIPs, false)
			lanIP := rItem1.LanIP.ValueString()
			localPort := rItem1.LocalPort.ValueString()
			name := rItem1.Name.ValueString()
			protocol := rItem1.Protocol.ValueString()
			publicPort := rItem1.PublicPort.ValueString()
			uplink := rItem1.Uplink.ValueString()
			requestApplianceUpdateNetworkApplianceFirewallPortForwardingRulesRules = append(requestApplianceUpdateNetworkApplianceFirewallPortForwardingRulesRules, merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallPortForwardingRulesRules{
				AllowedIPs: allowedIPs,
				LanIP:      lanIP,
				LocalPort:  localPort,
				Name:       name,
				Protocol:   protocol,
				PublicPort: publicPort,
				Uplink:     uplink,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallPortForwardingRules{
		Rules: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallPortForwardingRulesRules {
			if len(requestApplianceUpdateNetworkApplianceFirewallPortForwardingRulesRules) > 0 || r.Rules != nil {
				return &requestApplianceUpdateNetworkApplianceFirewallPortForwardingRulesRules
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesItemToBodyRs(state NetworksApplianceFirewallPortForwardingRulesRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceFirewallPortForwardingRules, is_read bool) NetworksApplianceFirewallPortForwardingRulesRs {
	itemState := NetworksApplianceFirewallPortForwardingRulesRs{
		Rules: func() *[]ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRulesRs {
			if response.Rules != nil {
				result := make([]ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRulesRs, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRulesRs{
						AllowedIPs: StringSliceToList(rules.AllowedIPs),
						LanIP: func() types.String {
							if rules.LanIP != "" {
								return types.StringValue(rules.LanIP)
							}
							return types.String{}
						}(),
						LocalPort: func() types.String {
							if rules.LocalPort != "" {
								return types.StringValue(rules.LocalPort)
							}
							return types.String{}
						}(),
						Name: func() types.String {
							if rules.Name != "" {
								return types.StringValue(rules.Name)
							}
							return types.String{}
						}(),
						Protocol: func() types.String {
							if rules.Protocol != "" {
								return types.StringValue(rules.Protocol)
							}
							return types.String{}
						}(),
						PublicPort: func() types.String {
							if rules.PublicPort != "" {
								return types.StringValue(rules.PublicPort)
							}
							return types.String{}
						}(),
						Uplink: func() types.String {
							if rules.Uplink != "" {
								return types.StringValue(rules.Uplink)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceFirewallPortForwardingRulesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceFirewallPortForwardingRulesRs)
}
