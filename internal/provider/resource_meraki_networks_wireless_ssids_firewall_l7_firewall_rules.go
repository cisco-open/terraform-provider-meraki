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
	"fmt"
	"strconv"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWirelessSSIDsFirewallL7FirewallRulesResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessSSIDsFirewallL7FirewallRulesResource{}
)

func NewNetworksWirelessSSIDsFirewallL7FirewallRulesResource() resource.Resource {
	return &NetworksWirelessSSIDsFirewallL7FirewallRulesResource{}
}

type NetworksWirelessSSIDsFirewallL7FirewallRulesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessSSIDsFirewallL7FirewallRulesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessSSIDsFirewallL7FirewallRulesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_firewall_l7_firewall_rules"
}

func (r *NetworksWirelessSSIDsFirewallL7FirewallRulesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
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

						"policy": schema.StringAttribute{
							MarkdownDescription: `'Deny' traffic specified by this rule
                                        Allowed values: [deny]`,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"deny",
								),
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `Type of the L7 firewall rule. One of: 'application', 'applicationCategory', 'host', 'port', 'ipRange'
                                        Allowed values: [application,applicationCategory,host,ipRange,port]`,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"application",
									"applicationCategory",
									"host",
									"ipRange",
									"port",
								),
							},
						},
						"value": schema.StringAttribute{
							MarkdownDescription: `The value of what needs to get blocked. Format of the value varies depending on type of the firewall rule selected.`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"value_list": schema.SetAttribute{
							MarkdownDescription: `The list of values of what needs to get blocked. Format of the value varies depending on type of the firewall rule selected.`,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},
							ElementType: types.StringType,
						},
						"value_obj": schema.SingleNestedAttribute{
							MarkdownDescription: `The object of what needs to get blocked. Format of the value varies depending on type of the firewall rule selected.`,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Optional: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"name": schema.StringAttribute{
									Optional: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *NetworksWirelessSSIDsFirewallL7FirewallRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessSSIDsFirewallL7FirewallRulesRs

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
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDFirewallL7FirewallRules(vvNetworkID, vvNumber, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDFirewallL7FirewallRules",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDFirewallL7FirewallRules",
			err.Error(),
		)
		return
	}

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *NetworksWirelessSSIDsFirewallL7FirewallRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessSSIDsFirewallL7FirewallRulesRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	vvNumber := data.Number.ValueString()
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessSSIDFirewallL7FirewallRules(vvNetworkID, vvNumber)
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
				"Failure when executing GetNetworkWirelessSSIDFirewallL7FirewallRules",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDFirewallL7FirewallRules",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRulesItemToBodyRs(data, responseGet, true)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksWirelessSSIDsFirewallL7FirewallRulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

func (r *NetworksWirelessSSIDsFirewallL7FirewallRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksWirelessSSIDsFirewallL7FirewallRulesRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	vvNumber := plan.Number.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDFirewallL7FirewallRules(vvNetworkID, vvNumber, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDFirewallL7FirewallRules",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDFirewallL7FirewallRules",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworksWirelessSSIDsFirewallL7FirewallRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksWirelessSSIDsFirewallL7FirewallRules", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessSSIDsFirewallL7FirewallRulesRs struct {
	NetworkID types.String                                                            `tfsdk:"network_id"`
	Number    types.String                                                            `tfsdk:"number"`
	Rules     *[]ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRulesRs `tfsdk:"rules"`
}

type ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRulesRs struct {
	Policy    types.String                                                                `tfsdk:"policy"`
	Type      types.String                                                                `tfsdk:"type"`
	Value     types.String                                                                `tfsdk:"value"`
	ValueList types.Set                                                                   `tfsdk:"value_list"`
	ValueObj  *ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRulesValueObj `tfsdk:"value_obj"`
}

// FromBody
func (r *NetworksWirelessSSIDsFirewallL7FirewallRulesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRules {
	var requestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules

	if r.Rules != nil {
		for _, rItem1 := range *r.Rules {
			var valueR interface{}
			policy := rItem1.Policy.ValueString()
			typeR := rItem1.Type.ValueString()
			value := rItem1.Value.ValueString()
			var valueList []string
			rItem1.ValueList.ElementsAs(ctx, &valueList, false)
			var requestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRulesValue *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRulesValue
			if rItem1.ValueObj != nil {
				name := rItem1.ValueObj.Name.ValueString()
				id := rItem1.ValueObj.ID.ValueString()
				requestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRulesValue = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRulesValue{
					ID:   id,
					Name: name,
				}
			}
			if !rItem1.Value.IsNull() && !rItem1.Value.IsUnknown() && rItem1.Type.ValueString() != "blockedCountries" && rItem1.Type.ValueString() != "applicationCategory" {
				valueR = value
			} else {
				if !rItem1.ValueList.IsNull() && !rItem1.ValueList.IsUnknown() && rItem1.Type.ValueString() == "blockedCountries" {
					valueR = valueList
				} else {
					valueR = requestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRulesValue
				}
			}
			requestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules = append(requestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules{
				Policy: policy,
				Type:   typeR,
				Value:  valueR,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRules{
		Rules: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules {
			if len(requestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules) > 0 || r.Rules != nil {
				return &requestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRulesItemToBodyRs(state NetworksWirelessSSIDsFirewallL7FirewallRulesRs, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRules, is_read bool) NetworksWirelessSSIDsFirewallL7FirewallRulesRs {
	itemState := NetworksWirelessSSIDsFirewallL7FirewallRulesRs{
		Rules: func() *[]ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRulesRs {
			if response.Rules != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRulesRs, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRulesRs{
						Policy: func() types.String {
							if rules.Policy != "" {
								return types.StringValue(rules.Policy)
							}
							return types.String{}
						}(),
						Type: func() types.String {
							if rules.Type != "" {
								return types.StringValue(rules.Type)
							}
							return types.String{}
						}(),
						Value: func() types.String {
							if rules.Value != nil {
								return types.StringValue(*rules.Value)
							}
							return types.String{}
						}(),
						ValueList: func() types.Set {
							if rules.ValueList == nil {
								return types.SetNull(types.StringType)
							}
							return StringSliceToSet(*rules.ValueList)
						}(),
						ValueObj: func() *ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRulesValueObj {
							if rules.ValueObj == nil {
								return nil
							}
							return &ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRulesValueObj{
								ID: func() types.String {
									if rules.ValueObj.ID != "" {
										return types.StringValue(rules.ValueObj.ID)
									}
									return types.String{}
								}(),
								Name: func() types.String {
									if rules.ValueObj.Name != "" {
										return types.StringValue(rules.ValueObj.Name)
									}
									return types.String{}
								}(),
							}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessSSIDsFirewallL7FirewallRulesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessSSIDsFirewallL7FirewallRulesRs)
}
