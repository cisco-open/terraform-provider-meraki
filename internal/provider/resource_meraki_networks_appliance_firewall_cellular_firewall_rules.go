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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceFirewallCellularFirewallRulesResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceFirewallCellularFirewallRulesResource{}
)

func NewNetworksApplianceFirewallCellularFirewallRulesResource() resource.Resource {
	return &NetworksApplianceFirewallCellularFirewallRulesResource{}
}

type NetworksApplianceFirewallCellularFirewallRulesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceFirewallCellularFirewallRulesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceFirewallCellularFirewallRulesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_cellular_firewall_rules"
}

func (r *NetworksApplianceFirewallCellularFirewallRulesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"rules": schema.SetNestedAttribute{
				MarkdownDescription: `An ordered array of the firewall rules (not including the default rule)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"comment": schema.StringAttribute{
							MarkdownDescription: `Description of the rule (optional)`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"dest_cidr": schema.StringAttribute{
							MarkdownDescription: `Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"dest_port": schema.StringAttribute{
							MarkdownDescription: `Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"policy": schema.StringAttribute{
							MarkdownDescription: `'allow' or 'deny' traffic specified by this rule
                                        Allowed values: [allow,deny]`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"allow",
									"deny",
								),
							},
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: `The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
                                        Allowed values: [any,icmp,icmp6,tcp,udp]`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"any",
									"icmp",
									"icmp6",
									"tcp",
									"udp",
								),
							},
						},
						"src_cidr": schema.StringAttribute{
							MarkdownDescription: `Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (note: FQDN not supported for source addresses)`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"src_port": schema.StringAttribute{
							MarkdownDescription: `Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"syslog_enabled": schema.BoolAttribute{
							MarkdownDescription: `Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
		},
	}
}

func (r *NetworksApplianceFirewallCellularFirewallRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceFirewallCellularFirewallRulesRs

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
	//Item
	responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceFirewallCellularFirewallRules(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceFirewallCellularFirewallRules only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceFirewallCellularFirewallRules only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceFirewallCellularFirewallRules(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceFirewallCellularFirewallRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceFirewallCellularFirewallRules",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceFirewallCellularFirewallRules(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceFirewallCellularFirewallRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceFirewallCellularFirewallRules",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRulesItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceFirewallCellularFirewallRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceFirewallCellularFirewallRulesRs

	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceFirewallCellularFirewallRules(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceFirewallCellularFirewallRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceFirewallCellularFirewallRules",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRulesItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksApplianceFirewallCellularFirewallRulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceFirewallCellularFirewallRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceFirewallCellularFirewallRulesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceFirewallCellularFirewallRules(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceFirewallCellularFirewallRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceFirewallCellularFirewallRules",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceFirewallCellularFirewallRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceFirewallCellularFirewallRules", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceFirewallCellularFirewallRulesRs struct {
	NetworkID types.String                                                                `tfsdk:"network_id"`
	Rules     *[]ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRulesRulesRs `tfsdk:"rules"`
}

type ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRulesRulesRs struct {
	Comment       types.String `tfsdk:"comment"`
	DestCidr      types.String `tfsdk:"dest_cidr"`
	DestPort      types.String `tfsdk:"dest_port"`
	Policy        types.String `tfsdk:"policy"`
	Protocol      types.String `tfsdk:"protocol"`
	SrcCidr       types.String `tfsdk:"src_cidr"`
	SrcPort       types.String `tfsdk:"src_port"`
	SyslogEnabled types.Bool   `tfsdk:"syslog_enabled"`
}

// FromBody
func (r *NetworksApplianceFirewallCellularFirewallRulesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallCellularFirewallRules {
	var requestApplianceUpdateNetworkApplianceFirewallCellularFirewallRulesRules []merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallCellularFirewallRulesRules
	if r.Rules != nil {
		for _, rItem1 := range *r.Rules {
			comment := rItem1.Comment.ValueString()
			destCidr := rItem1.DestCidr.ValueString()
			destPort := rItem1.DestPort.ValueString()
			policy := rItem1.Policy.ValueString()
			protocol := rItem1.Protocol.ValueString()
			srcCidr := rItem1.SrcCidr.ValueString()
			srcPort := rItem1.SrcPort.ValueString()
			syslogEnabled := func() *bool {
				if !rItem1.SyslogEnabled.IsUnknown() && !rItem1.SyslogEnabled.IsNull() {
					return rItem1.SyslogEnabled.ValueBoolPointer()
				}
				return nil
			}()
			requestApplianceUpdateNetworkApplianceFirewallCellularFirewallRulesRules = append(requestApplianceUpdateNetworkApplianceFirewallCellularFirewallRulesRules, merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallCellularFirewallRulesRules{
				Comment:       comment,
				DestCidr:      destCidr,
				DestPort:      destPort,
				Policy:        policy,
				Protocol:      protocol,
				SrcCidr:       srcCidr,
				SrcPort:       srcPort,
				SyslogEnabled: syslogEnabled,
			})
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallCellularFirewallRules{
		Rules: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallCellularFirewallRulesRules {
			if len(requestApplianceUpdateNetworkApplianceFirewallCellularFirewallRulesRules) > 0 || r.Rules != nil {
				return &requestApplianceUpdateNetworkApplianceFirewallCellularFirewallRulesRules
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRulesItemToBodyRs(state NetworksApplianceFirewallCellularFirewallRulesRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRules, is_read bool) NetworksApplianceFirewallCellularFirewallRulesRs {
	itemState := NetworksApplianceFirewallCellularFirewallRulesRs{
		Rules: func() *[]ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRulesRulesRs {
			if response.Rules != nil {
				result := make([]ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRulesRulesRs, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseApplianceGetNetworkApplianceFirewallCellularFirewallRulesRulesRs{
						Comment:  types.StringValue(rules.Comment),
						DestCidr: types.StringValue(rules.DestCidr),
						DestPort: types.StringValue(rules.DestPort),
						Policy:   types.StringValue(rules.Policy),
						Protocol: types.StringValue(rules.Protocol),
						SrcCidr:  types.StringValue(rules.SrcCidr),
						SrcPort:  types.StringValue(rules.SrcPort),
						SyslogEnabled: func() types.Bool {
							if rules.SyslogEnabled != nil {
								return types.BoolValue(*rules.SyslogEnabled)
							}
							return types.Bool{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceFirewallCellularFirewallRulesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceFirewallCellularFirewallRulesRs)
}
