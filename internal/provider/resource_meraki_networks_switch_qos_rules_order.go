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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSwitchQosRulesOrderResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchQosRulesOrderResource{}
)

func NewNetworksSwitchQosRulesOrderResource() resource.Resource {
	return &NetworksSwitchQosRulesOrderResource{}
}

type NetworksSwitchQosRulesOrderResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchQosRulesOrderResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchQosRulesOrderResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_qos_rules_order"
}

func (r *NetworksSwitchQosRulesOrderResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"dscp": schema.Int64Attribute{
				MarkdownDescription: `DSCP tag for the incoming packet. Set this to -1 to trust incoming DSCP. Default value is 0`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"dst_port": schema.Int64Attribute{
				MarkdownDescription: `The destination port of the incoming packet. Applicable only if protocol is TCP or UDP.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"dst_port_range": schema.StringAttribute{
				MarkdownDescription: `The destination port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: `Qos Rule id`,
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: `The protocol of the incoming packet. Can be one of "ANY", "TCP" or "UDP". Default value is "ANY"
                                  Allowed values: [ANY,TCP,UDP]`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"ANY",
						"TCP",
						"UDP",
					),
				},
			},
			"qos_rule_id": schema.StringAttribute{
				MarkdownDescription: `qosRuleId path parameter. Qos rule ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"src_port": schema.Int64Attribute{
				MarkdownDescription: `The source port of the incoming packet. Applicable only if protocol is TCP or UDP.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"src_port_range": schema.StringAttribute{
				MarkdownDescription: `The source port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"vlan": schema.Int64Attribute{
				MarkdownDescription: `The VLAN of the incoming packet. A null value will match any VLAN.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

//path params to set ['qosRuleId']

func (r *NetworksSwitchQosRulesOrderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchQosRulesOrderRs

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
	// network_id
	//Reviw This  Has Item and item
	//HAS CREATE

	vvQosRuleID := data.QosRuleID.ValueString()
	if vvQosRuleID != "" {
		responseVerifyItem, restyRespGet, err := r.client.Switch.GetNetworkSwitchQosRule(vvNetworkID, vvQosRuleID)
		if err != nil || responseVerifyItem == nil {
			if restyRespGet != nil {
				if restyRespGet.StatusCode() != 404 {

					resp.Diagnostics.AddError(
						"Failure when executing GetNetworkSwitchQosRule",
						err.Error(),
					)
					return
				}
			}
		}

		if responseVerifyItem != nil {
			data = ResponseSwitchGetNetworkSwitchQosRuleItemToBodyRs(data, responseVerifyItem, false)
			diags := resp.State.Set(ctx, &data)
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	response, restyResp2, err := r.client.Switch.CreateNetworkSwitchQosRule(vvNetworkID, data.toSdkApiRequestCreate(ctx))

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing ",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing ",
			err.Error(),
		)
		return
	}
	vvQosRuleID = response.ID
	//Items
	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchQosRule(vvNetworkID, vvQosRuleID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchQosRules",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchQosRules",
			err.Error(),
		)
		return
	} else {
		data = ResponseSwitchGetNetworkSwitchQosRuleItemToBodyRs(data, responseGet, false)
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}
}

func (r *NetworksSwitchQosRulesOrderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchQosRulesOrderRs

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
	vvQosRuleID := data.QosRuleID.ValueString()
	responseGet, restyRespGet, err := r.client.Switch.GetNetworkSwitchQosRule(vvNetworkID, vvQosRuleID)
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
				"Failure when executing GetNetworkSwitchQosRule",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchQosRule",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSwitchGetNetworkSwitchQosRuleItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchQosRulesOrderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("qos_rule_id"), idParts[1])...)
}

func (r *NetworksSwitchQosRulesOrderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSwitchQosRulesOrderRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvQosRuleID := data.QosRuleID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchQosRule(vvNetworkID, vvQosRuleID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchQosRule",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchQosRule",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchQosRulesOrderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksSwitchQosRulesOrderRs
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &state, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}

	vvNetworkID := state.NetworkID.ValueString()
	vvQosRuleID := state.QosRuleID.ValueString()
	_, err := r.client.Switch.DeleteNetworkSwitchQosRule(vvNetworkID, vvQosRuleID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkSwitchQosRule", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksSwitchQosRulesOrderRs struct {
	NetworkID    types.String `tfsdk:"network_id"`
	QosRuleID    types.String `tfsdk:"qos_rule_id"`
	Dscp         types.Int64  `tfsdk:"dscp"`
	DstPort      types.Int64  `tfsdk:"dst_port"`
	DstPortRange types.String `tfsdk:"dst_port_range"`
	ID           types.String `tfsdk:"id"`
	Protocol     types.String `tfsdk:"protocol"`
	SrcPort      types.Int64  `tfsdk:"src_port"`
	SrcPortRange types.String `tfsdk:"src_port_range"`
	VLAN         types.Int64  `tfsdk:"vlan"`
}

// FromBody
func (r *NetworksSwitchQosRulesOrderRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSwitchCreateNetworkSwitchQosRule {
	emptyString := ""
	dscp := new(int64)
	if !r.Dscp.IsUnknown() && !r.Dscp.IsNull() {
		*dscp = r.Dscp.ValueInt64()
	} else {
		dscp = nil
	}
	dstPort := new(int64)
	if !r.DstPort.IsUnknown() && !r.DstPort.IsNull() {
		*dstPort = r.DstPort.ValueInt64()
	} else {
		dstPort = nil
	}
	dstPortRange := new(string)
	if !r.DstPortRange.IsUnknown() && !r.DstPortRange.IsNull() {
		*dstPortRange = r.DstPortRange.ValueString()
	} else {
		dstPortRange = &emptyString
	}
	protocol := new(string)
	if !r.Protocol.IsUnknown() && !r.Protocol.IsNull() {
		*protocol = r.Protocol.ValueString()
	} else {
		protocol = &emptyString
	}
	srcPort := new(int64)
	if !r.SrcPort.IsUnknown() && !r.SrcPort.IsNull() {
		*srcPort = r.SrcPort.ValueInt64()
	} else {
		srcPort = nil
	}
	srcPortRange := new(string)
	if !r.SrcPortRange.IsUnknown() && !r.SrcPortRange.IsNull() {
		*srcPortRange = r.SrcPortRange.ValueString()
	} else {
		srcPortRange = &emptyString
	}
	vLAN := new(int64)
	if !r.VLAN.IsUnknown() && !r.VLAN.IsNull() {
		*vLAN = r.VLAN.ValueInt64()
	} else {
		vLAN = nil
	}
	out := merakigosdk.RequestSwitchCreateNetworkSwitchQosRule{
		Dscp:         int64ToIntPointer(dscp),
		DstPort:      int64ToIntPointer(dstPort),
		DstPortRange: *dstPortRange,
		Protocol:     *protocol,
		SrcPort:      int64ToIntPointer(srcPort),
		SrcPortRange: *srcPortRange,
		VLAN:         int64ToIntPointer(vLAN),
	}
	return &out
}
func (r *NetworksSwitchQosRulesOrderRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchQosRule {
	emptyString := ""
	dscp := new(int64)
	if !r.Dscp.IsUnknown() && !r.Dscp.IsNull() {
		*dscp = r.Dscp.ValueInt64()
	} else {
		dscp = nil
	}
	dstPort := new(int64)
	if !r.DstPort.IsUnknown() && !r.DstPort.IsNull() {
		*dstPort = r.DstPort.ValueInt64()
	} else {
		dstPort = nil
	}
	dstPortRange := new(string)
	if !r.DstPortRange.IsUnknown() && !r.DstPortRange.IsNull() {
		*dstPortRange = r.DstPortRange.ValueString()
	} else {
		dstPortRange = &emptyString
	}
	protocol := new(string)
	if !r.Protocol.IsUnknown() && !r.Protocol.IsNull() {
		*protocol = r.Protocol.ValueString()
	} else {
		protocol = &emptyString
	}
	srcPort := new(int64)
	if !r.SrcPort.IsUnknown() && !r.SrcPort.IsNull() {
		*srcPort = r.SrcPort.ValueInt64()
	} else {
		srcPort = nil
	}
	srcPortRange := new(string)
	if !r.SrcPortRange.IsUnknown() && !r.SrcPortRange.IsNull() {
		*srcPortRange = r.SrcPortRange.ValueString()
	} else {
		srcPortRange = &emptyString
	}
	vLAN := new(int64)
	if !r.VLAN.IsUnknown() && !r.VLAN.IsNull() {
		*vLAN = r.VLAN.ValueInt64()
	} else {
		vLAN = nil
	}
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchQosRule{
		Dscp:         int64ToIntPointer(dscp),
		DstPort:      int64ToIntPointer(dstPort),
		DstPortRange: *dstPortRange,
		Protocol:     *protocol,
		SrcPort:      int64ToIntPointer(srcPort),
		SrcPortRange: *srcPortRange,
		VLAN:         int64ToIntPointer(vLAN),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchQosRuleItemToBodyRs(state NetworksSwitchQosRulesOrderRs, response *merakigosdk.ResponseSwitchGetNetworkSwitchQosRule, is_read bool) NetworksSwitchQosRulesOrderRs {
	itemState := NetworksSwitchQosRulesOrderRs{
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
		DstPortRange: func() types.String {
			if response.DstPortRange != "" {
				return types.StringValue(response.DstPortRange)
			}
			return types.String{}
		}(),
		ID: func() types.String {
			if response.ID != "" {
				return types.StringValue(response.ID)
			}
			return types.String{}
		}(),
		Protocol: func() types.String {
			if response.Protocol != "" {
				return types.StringValue(response.Protocol)
			}
			return types.String{}
		}(),
		SrcPort: func() types.Int64 {
			if response.SrcPort != nil {
				return types.Int64Value(int64(*response.SrcPort))
			}
			return types.Int64{}
		}(),
		SrcPortRange: func() types.String {
			if response.SrcPortRange != "" {
				return types.StringValue(response.SrcPortRange)
			}
			return types.String{}
		}(),
		VLAN: func() types.Int64 {
			if response.VLAN != nil {
				return types.Int64Value(int64(*response.VLAN))
			}
			return types.Int64{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSwitchQosRulesOrderRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSwitchQosRulesOrderRs)
}
