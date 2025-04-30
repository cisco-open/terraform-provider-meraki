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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSwitchAccessControlListsResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchAccessControlListsResource{}
)

func NewNetworksSwitchAccessControlListsResource() resource.Resource {
	return &NetworksSwitchAccessControlListsResource{}
}

type NetworksSwitchAccessControlListsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchAccessControlListsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchAccessControlListsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_access_control_lists"
}

func (r *NetworksSwitchAccessControlListsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"rules_response": schema.SetNestedAttribute{
				MarkdownDescription: `An ordered array of the access control list rules`,
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
						"dst_cidr": schema.StringAttribute{
							MarkdownDescription: `Destination IP address (in IP or CIDR notation)`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"dst_port": schema.StringAttribute{
							MarkdownDescription: `Destination port`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"ip_version": schema.StringAttribute{
							MarkdownDescription: `IP address version`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"policy": schema.StringAttribute{
							MarkdownDescription: `'allow' or 'deny' traffic specified by this rule`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: `The type of protocol`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"src_cidr": schema.StringAttribute{
							MarkdownDescription: `Source IP address (in IP or CIDR notation)`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"src_port": schema.StringAttribute{
							MarkdownDescription: `Source port`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"vlan": schema.StringAttribute{
							MarkdownDescription: `ncoming traffic VLAN`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"rules": schema.SetNestedAttribute{
				MarkdownDescription: `An ordered array of the access control list rules`,
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
						"dst_cidr": schema.StringAttribute{
							MarkdownDescription: `Destination IP address (in IP or CIDR notation)`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"dst_port": schema.StringAttribute{
							MarkdownDescription: `Destination port`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"ip_version": schema.StringAttribute{
							MarkdownDescription: `IP address version
                                        Allowed values: [any,ipv4,ipv6]`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"any",
									"ipv4",
									"ipv6",
								),
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
							MarkdownDescription: `The type of protocol
                                        Allowed values: [any,tcp,udp]`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"any",
									"tcp",
									"udp",
								),
							},
						},
						"src_cidr": schema.StringAttribute{
							MarkdownDescription: `Source IP address (in IP or CIDR notation)`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"src_port": schema.StringAttribute{
							MarkdownDescription: `Source port`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"vlan": schema.StringAttribute{
							MarkdownDescription: `ncoming traffic VLAN`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
		},
	}
}

func (r *NetworksSwitchAccessControlListsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchAccessControlListsRs

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

	if vvNetworkID != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Switch.GetNetworkSwitchAccessControlLists(vvNetworkID)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksSwitchAccessControlLists  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksSwitchAccessControlLists only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchAccessControlLists(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchAccessControlLists",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchAccessControlLists",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchAccessControlLists(vvNetworkID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchAccessControlLists",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchAccessControlLists",
			err.Error(),
		)
		return
	}

	data = ResponseSwitchGetNetworkSwitchAccessControlListsItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *NetworksSwitchAccessControlListsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchAccessControlListsRs

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
	responseGet, restyRespGet, err := r.client.Switch.GetNetworkSwitchAccessControlLists(vvNetworkID)
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
				"Failure when executing GetNetworkSwitchAccessControlLists",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchAccessControlLists",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSwitchGetNetworkSwitchAccessControlListsItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksSwitchAccessControlListsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksSwitchAccessControlListsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSwitchAccessControlListsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchAccessControlLists(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchAccessControlLists",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchAccessControlLists",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchAccessControlListsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksSwitchAccessControlLists", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSwitchAccessControlListsRs struct {
	NetworkID     types.String                                               `tfsdk:"network_id"`
	Rules         *[]ResponseSwitchGetNetworkSwitchAccessControlListsRulesRs `tfsdk:"rules"`
	RulesResponse *[]ResponseSwitchGetNetworkSwitchAccessControlListsRulesRs `tfsdk:"rules_response"`
}

type ResponseSwitchGetNetworkSwitchAccessControlListsRulesRs struct {
	Comment   types.String `tfsdk:"comment"`
	DstCidr   types.String `tfsdk:"dst_cidr"`
	DstPort   types.String `tfsdk:"dst_port"`
	IPVersion types.String `tfsdk:"ip_version"`
	Policy    types.String `tfsdk:"policy"`
	Protocol  types.String `tfsdk:"protocol"`
	SrcCidr   types.String `tfsdk:"src_cidr"`
	SrcPort   types.String `tfsdk:"src_port"`
	VLAN      types.String `tfsdk:"vlan"`
}

// FromBody
func (r *NetworksSwitchAccessControlListsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchAccessControlLists {
	var requestSwitchUpdateNetworkSwitchAccessControlListsRules []merakigosdk.RequestSwitchUpdateNetworkSwitchAccessControlListsRules

	if r.Rules != nil {
		for _, rItem1 := range *r.Rules {
			comment := rItem1.Comment.ValueString()
			dstCidr := rItem1.DstCidr.ValueString()
			dstPort := rItem1.DstPort.ValueString()
			ipVersion := rItem1.IPVersion.ValueString()
			policy := rItem1.Policy.ValueString()
			protocol := rItem1.Protocol.ValueString()
			srcCidr := rItem1.SrcCidr.ValueString()
			srcPort := rItem1.SrcPort.ValueString()
			vlan := rItem1.VLAN.ValueString()
			requestSwitchUpdateNetworkSwitchAccessControlListsRules = append(requestSwitchUpdateNetworkSwitchAccessControlListsRules, merakigosdk.RequestSwitchUpdateNetworkSwitchAccessControlListsRules{
				Comment:   comment,
				DstCidr:   dstCidr,
				DstPort:   dstPort,
				IPVersion: ipVersion,
				Policy:    policy,
				Protocol:  protocol,
				SrcCidr:   srcCidr,
				SrcPort:   srcPort,
				VLAN:      vlan,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchAccessControlLists{
		Rules: func() *[]merakigosdk.RequestSwitchUpdateNetworkSwitchAccessControlListsRules {
			if len(requestSwitchUpdateNetworkSwitchAccessControlListsRules) > 0 {
				return &requestSwitchUpdateNetworkSwitchAccessControlListsRules
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchAccessControlListsItemToBodyRs(state NetworksSwitchAccessControlListsRs, response *merakigosdk.ResponseSwitchGetNetworkSwitchAccessControlLists, is_read bool) NetworksSwitchAccessControlListsRs {
	itemState := NetworksSwitchAccessControlListsRs{
		RulesResponse: func() *[]ResponseSwitchGetNetworkSwitchAccessControlListsRulesRs {
			if response.Rules != nil {
				result := make([]ResponseSwitchGetNetworkSwitchAccessControlListsRulesRs, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseSwitchGetNetworkSwitchAccessControlListsRulesRs{
						Comment:   types.StringValue(rules.Comment),
						DstCidr:   types.StringValue(rules.DstCidr),
						DstPort:   types.StringValue(rules.DstPort),
						IPVersion: types.StringValue(rules.IPVersion),
						Policy:    types.StringValue(rules.Policy),
						Protocol:  types.StringValue(rules.Protocol),
						SrcCidr:   types.StringValue(rules.SrcCidr),
						SrcPort:   types.StringValue(rules.SrcPort),
						VLAN:      types.StringValue(rules.VLAN),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	itemState.Rules = state.Rules
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSwitchAccessControlListsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSwitchAccessControlListsRs)
}
