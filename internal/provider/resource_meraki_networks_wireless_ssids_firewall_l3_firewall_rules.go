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
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWirelessSSIDsFirewallL3FirewallRulesResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessSSIDsFirewallL3FirewallRulesResource{}
)

func NewNetworksWirelessSSIDsFirewallL3FirewallRulesResource() resource.Resource {
	return &NetworksWirelessSSIDsFirewallL3FirewallRulesResource{}
}

type NetworksWirelessSSIDsFirewallL3FirewallRulesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessSSIDsFirewallL3FirewallRulesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessSSIDsFirewallL3FirewallRulesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_firewall_l3_firewall_rules"
}

func (r *NetworksWirelessSSIDsFirewallL3FirewallRulesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"allow_lan_access": schema.BoolAttribute{
				MarkdownDescription: `Allows wireless client access to local LAN (boolean value - true allows access and false denies access)`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: `number path parameter.`,
				Required:            true,
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: `An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule).`,
				Optional:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"comment": schema.StringAttribute{
							MarkdownDescription: `Description of the rule (optional)`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"dest_cidr": schema.StringAttribute{
							MarkdownDescription: `Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'`,
							Optional:            true,

							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								&CaseInsensitiveStringPlanModifier{},
							},
						},
						"dest_port": schema.StringAttribute{
							MarkdownDescription: `Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								&CaseInsensitiveStringPlanModifier{},
							},
						},
						"policy": schema.StringAttribute{
							MarkdownDescription: `'allow' or 'deny' traffic specified by this rule
                                        Allowed values: [allow,deny]`,
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
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								&CaseInsensitiveStringPlanModifier{},
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
					},
				},
			},
		},
	}
}

func (r *NetworksWirelessSSIDsFirewallL3FirewallRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessSSIDsFirewallL3FirewallRulesRs

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
	vvNumber := data.Number.ValueString()
	//Has Item and not has items

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDFirewallL3FirewallRules(vvNetworkID, vvNumber, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDFirewallL3FirewallRules",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDFirewallL3FirewallRules",
			err.Error(),
		)
		return
	}

	// Update response to read Struct
	var responseRead *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRules
	err = json.Unmarshal(restyResp2.Body(), &responseRead)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when unmarshalling response",
			err.Error(),
		)
	}
	data = ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRulesItemToBodyRs(data, responseRead, true)

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *NetworksWirelessSSIDsFirewallL3FirewallRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessSSIDsFirewallL3FirewallRulesRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	vvNumber := data.Number.ValueString()
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessSSIDFirewallL3FirewallRules(vvNetworkID, vvNumber)
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
				"Failure when executing GetNetworkWirelessSSIDFirewallL3FirewallRules",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDFirewallL3FirewallRules",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRulesItemToBodyRs(data, responseGet, true)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksWirelessSSIDsFirewallL3FirewallRulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: networkId,number. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("number"), idParts[1])...)
}

func (r *NetworksWirelessSSIDsFirewallL3FirewallRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksWirelessSSIDsFirewallL3FirewallRulesRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	vvNumber := plan.Number.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDFirewallL3FirewallRules(vvNetworkID, vvNumber, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDFirewallL3FirewallRules",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDFirewallL3FirewallRules",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworksWirelessSSIDsFirewallL3FirewallRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksWirelessSSIDsFirewallL3FirewallRules", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessSSIDsFirewallL3FirewallRulesRs struct {
	NetworkID      types.String                                                            `tfsdk:"network_id"`
	Number         types.String                                                            `tfsdk:"number"`
	AllowLanAccess types.Bool                                                              `tfsdk:"allow_lan_access"`
	Rules          *[]ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRulesRulesRs `tfsdk:"rules"`
}

type ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRulesRulesRs struct {
	Comment  types.String `tfsdk:"comment"`
	DestCidr types.String `tfsdk:"dest_cidr"`
	DestPort types.String `tfsdk:"dest_port"`
	Policy   types.String `tfsdk:"policy"`
	Protocol types.String `tfsdk:"protocol"`
}

// FromBody
func (r *NetworksWirelessSSIDsFirewallL3FirewallRulesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRules {
	allowLanAccess := new(bool)
	if !r.AllowLanAccess.IsUnknown() && !r.AllowLanAccess.IsNull() {
		*allowLanAccess = r.AllowLanAccess.ValueBool()
	} else {
		allowLanAccess = nil
	}
	var requestWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRulesRules []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRulesRules

	if r.Rules != nil {
		for _, rItem1 := range *r.Rules {
			comment := rItem1.Comment.ValueString()
			destCidr := rItem1.DestCidr.ValueString()
			destPort := rItem1.DestPort.ValueString()
			policy := rItem1.Policy.ValueString()
			protocol := rItem1.Protocol.ValueString()
			requestWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRulesRules = append(requestWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRulesRules, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRulesRules{
				Comment:  comment,
				DestCidr: destCidr,
				DestPort: destPort,
				Policy:   policy,
				Protocol: protocol,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRules{
		AllowLanAccess: allowLanAccess,
		Rules: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRulesRules {
			if len(requestWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRulesRules) > 0 {
				return &requestWirelessUpdateNetworkWirelessSSIDFirewallL3FirewallRulesRules
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRulesItemToBodyRs(state NetworksWirelessSSIDsFirewallL3FirewallRulesRs, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRules, is_read bool) NetworksWirelessSSIDsFirewallL3FirewallRulesRs {
	if response.Rules != nil {
		var filteredRules []merakigosdk.ResponseWirelessGetNetworkWirelessSSIDFirewallL3FirewallRulesRules

		for _, rule := range *response.Rules {
			// Skip the default rule since it's managed by the system
			if rule.Comment != "Default rule" {
				if rule.Comment == "Wireless clients accessing LAN" {
					continue
				}
				filteredRules = append(filteredRules, rule)
			}
		}
		// Update response with filtered rules, excluding default rule
		response.Rules = &filteredRules
	}
	itemState := NetworksWirelessSSIDsFirewallL3FirewallRulesRs{
		AllowLanAccess: func() types.Bool {
			if response.AllowLanAccess != nil {
				return types.BoolValue(*response.AllowLanAccess)
			}
			return types.Bool{}
		}(),
		Rules: func() *[]ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRulesRulesRs {
			if response.Rules != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRulesRulesRs, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseWirelessGetNetworkWirelessSsidFirewallL3FirewallRulesRulesRs{
						Comment: func() types.String {
							if rules.Comment != "" {
								return types.StringValue(rules.Comment)
							}
							return types.StringNull()
						}(),
						DestCidr: func() types.String {
							if rules.DestCidr != "" {
								return types.StringValue(strings.ToLower(rules.DestCidr))
							}
							return types.String{}
						}(),
						DestPort: func() types.String {
							if strings.ToLower(rules.DestPort) != "" {
								return types.StringValue(strings.ToLower(rules.DestPort))
							}
							return types.String{}
						}(),
						Policy: func() types.String {
							if rules.Policy != "" {
								return types.StringValue(rules.Policy)
							}
							return types.String{}
						}(),
						Protocol: func() types.String {
							if strings.ToLower(rules.Protocol) != "" {
								return types.StringValue(strings.ToLower(rules.Protocol))
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
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessSSIDsFirewallL3FirewallRulesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessSSIDsFirewallL3FirewallRulesRs)
}

// CaseInsensitiveStringPlanModifier is a plan modifier that normalizes string values to lowercase
// for case-insensitive comparison during plan operations
type CaseInsensitiveStringPlanModifier struct{}

func (m *CaseInsensitiveStringPlanModifier) Description(ctx context.Context) string {
	return "Normalizes string values to lowercase for case-insensitive comparison"
}

func (m *CaseInsensitiveStringPlanModifier) MarkdownDescription(ctx context.Context) string {
	return "Normalizes string values to lowercase for case-insensitive comparison"
}

func (m *CaseInsensitiveStringPlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// If the plan value is unknown, don't modify it
	if req.PlanValue.IsUnknown() {
		return
	}

	// If the state value is unknown, don't modify it
	if req.StateValue.IsUnknown() {
		return
	}

	// If both values are the same (case-insensitive), use the state value to avoid unnecessary changes
	if strings.EqualFold(req.PlanValue.ValueString(), req.StateValue.ValueString()) {
		resp.PlanValue = req.StateValue
		return
	}

	// Otherwise, keep the plan value as is
	resp.PlanValue = req.PlanValue
}
