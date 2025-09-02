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

	"log"

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
	_ resource.Resource              = &NetworksApplianceFirewallOneToManyNatRulesResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceFirewallOneToManyNatRulesResource{}
)

func NewNetworksApplianceFirewallOneToManyNatRulesResource() resource.Resource {
	return &NetworksApplianceFirewallOneToManyNatRulesResource{}
}

type NetworksApplianceFirewallOneToManyNatRulesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceFirewallOneToManyNatRulesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceFirewallOneToManyNatRulesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_one_to_many_nat_rules"
}

func (r *NetworksApplianceFirewallOneToManyNatRulesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: `An array of 1:Many nat rules`,
				Optional:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"port_rules": schema.ListNestedAttribute{
							MarkdownDescription: `An array of associated forwarding rules`,
							Optional:            true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.UseStateForUnknown(),
							},
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"allowed_ips": schema.ListAttribute{
										MarkdownDescription: `Remote IP addresses or ranges that are permitted to access the internal resource via this port forwarding rule, or 'any'`,
										Optional:            true,
										PlanModifiers: []planmodifier.List{
											listplanmodifier.UseStateForUnknown(),
										},

										ElementType: types.StringType,
									},
									"local_ip": schema.StringAttribute{
										MarkdownDescription: `Local IP address to which traffic will be forwarded`,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"local_port": schema.StringAttribute{
										MarkdownDescription: `Destination port of the forwarded traffic that will be sent from the MX to the specified host on the LAN. If you simply wish to forward the traffic without translating the port, this should be the same as the Public port`,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `A description of the rule`,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: `'tcp' or 'udp'
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
										MarkdownDescription: `Destination port of the traffic that is arriving on the WAN`,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
								},
							},
						},
						"public_ip": schema.StringAttribute{
							MarkdownDescription: `The IP address that will be used to access the internal resource from the WAN`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"uplink": schema.StringAttribute{
							MarkdownDescription: `The physical WAN interface on which the traffic will arrive, formatted as 'internetN' where N is an integer representing a valid uplink for the network's appliance`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"internet1",
									"internet2",
								),
							},
						},
					},
				},
			},
		},
	}
}

func (r *NetworksApplianceFirewallOneToManyNatRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceFirewallOneToManyNatRulesRs

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
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceFirewallOneToManyNatRules(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceFirewallOneToManyNatRules",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceFirewallOneToManyNatRules",
			err.Error(),
		)
		return
	}

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *NetworksApplianceFirewallOneToManyNatRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceFirewallOneToManyNatRulesRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceFirewallOneToManyNatRules(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceFirewallOneToManyNatRules",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceFirewallOneToManyNatRules",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesItemToBodyRs(data, responseGet, true)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksApplianceFirewallOneToManyNatRulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceFirewallOneToManyNatRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksApplianceFirewallOneToManyNatRulesRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceFirewallOneToManyNatRules(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceFirewallOneToManyNatRules",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceFirewallOneToManyNatRules",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworksApplianceFirewallOneToManyNatRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceFirewallOneToManyNatRules", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceFirewallOneToManyNatRulesRs struct {
	NetworkID types.String                                                            `tfsdk:"network_id"`
	Rules     *[]ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRulesRs `tfsdk:"rules"`
}

type ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRulesRs struct {
	PortRules *[]ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRulesPortRulesRs `tfsdk:"port_rules"`
	PublicIP  types.String                                                                     `tfsdk:"public_ip"`
	Uplink    types.String                                                                     `tfsdk:"uplink"`
}

type ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRulesPortRulesRs struct {
	AllowedIPs types.List   `tfsdk:"allowed_ips"`
	LocalIP    types.String `tfsdk:"local_ip"`
	LocalPort  types.String `tfsdk:"local_port"`
	Name       types.String `tfsdk:"name"`
	Protocol   types.String `tfsdk:"protocol"`
	PublicPort types.String `tfsdk:"public_port"`
}

// FromBody
func (r *NetworksApplianceFirewallOneToManyNatRulesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRules {
	var requestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRules []merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRules

	if r.Rules != nil {
		for _, rItem1 := range *r.Rules {

			log.Printf("[DEBUG] #TODO []RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRulesPortRules")
			var requestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRulesPortRules []merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRulesPortRules

			if rItem1.PortRules != nil {
				for _, rItem2 := range *rItem1.PortRules {

					var allowedIPs []string = nil
					rItem2.AllowedIPs.ElementsAs(ctx, &allowedIPs, false)
					localIP := rItem2.LocalIP.ValueString()
					localPort := rItem2.LocalPort.ValueString()
					name := rItem2.Name.ValueString()
					protocol := rItem2.Protocol.ValueString()
					publicPort := rItem2.PublicPort.ValueString()
					requestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRulesPortRules = append(requestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRulesPortRules, merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRulesPortRules{
						AllowedIPs: allowedIPs,
						LocalIP:    localIP,
						LocalPort:  localPort,
						Name:       name,
						Protocol:   protocol,
						PublicPort: publicPort,
					})
					//[debug] Is Array: True
				}
			}
			publicIP := rItem1.PublicIP.ValueString()
			uplink := rItem1.Uplink.ValueString()
			requestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRules = append(requestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRules, merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRules{
				PortRules: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRulesPortRules {
					if len(requestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRulesPortRules) > 0 {
						return &requestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRulesPortRules
					}
					return nil
				}(),
				PublicIP: publicIP,
				Uplink:   uplink,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRules{
		Rules: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRules {
			// Always return the rules array, even if empty, to avoid sending null
			return &requestApplianceUpdateNetworkApplianceFirewallOneToManyNatRulesRules
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesItemToBodyRs(state NetworksApplianceFirewallOneToManyNatRulesRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRules, is_read bool) NetworksApplianceFirewallOneToManyNatRulesRs {
	itemState := NetworksApplianceFirewallOneToManyNatRulesRs{
		Rules: func() *[]ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRulesRs {
			if response.Rules != nil {
				result := make([]ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRulesRs, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRulesRs{
						PortRules: func() *[]ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRulesPortRulesRs {
							if rules.PortRules != nil {
								result := make([]ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRulesPortRulesRs, len(*rules.PortRules))
								for i, portRules := range *rules.PortRules {
									result[i] = ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRulesPortRulesRs{
										AllowedIPs: StringSliceToList(portRules.AllowedIPs),
										LocalIP: func() types.String {
											if portRules.LocalIP != "" {
												return types.StringValue(portRules.LocalIP)
											}
											return types.String{}
										}(),
										LocalPort: func() types.String {
											if portRules.LocalPort != "" {
												return types.StringValue(portRules.LocalPort)
											}
											return types.String{}
										}(),
										Name: func() types.String {
											if portRules.Name != "" {
												return types.StringValue(portRules.Name)
											}
											return types.String{}
										}(),
										Protocol: func() types.String {
											if portRules.Protocol != "" {
												return types.StringValue(portRules.Protocol)
											}
											return types.String{}
										}(),
										PublicPort: func() types.String {
											if portRules.PublicPort != "" {
												return types.StringValue(portRules.PublicPort)
											}
											return types.String{}
										}(),
									}
								}
								return &result
							}
							return nil
						}(),
						PublicIP: func() types.String {
							if rules.PublicIP != "" {
								return types.StringValue(rules.PublicIP)
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
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceFirewallOneToManyNatRulesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceFirewallOneToManyNatRulesRs)
}
