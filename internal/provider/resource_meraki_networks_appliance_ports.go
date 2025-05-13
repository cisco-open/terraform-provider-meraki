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
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksAppliancePortsResource{}
	_ resource.ResourceWithConfigure = &NetworksAppliancePortsResource{}
)

func NewNetworksAppliancePortsResource() resource.Resource {
	return &NetworksAppliancePortsResource{}
}

type NetworksAppliancePortsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksAppliancePortsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksAppliancePortsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_ports"
}

func (r *NetworksAppliancePortsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"access_policy": schema.StringAttribute{
				MarkdownDescription: `The name of the policy. Only applicable to Access ports.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"allowed_vlans": schema.StringAttribute{
				MarkdownDescription: `Comma-delimited list of the VLAN ID's allowed on the port, or 'all' to permit all VLAN's on the port.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"drop_untagged_traffic": schema.BoolAttribute{
				MarkdownDescription: `Whether the trunk port can drop all untagged traffic.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `The status of the port`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"number": schema.Int64Attribute{
				MarkdownDescription: `Number of the port`,
				Computed:            true,
			},
			"port_id": schema.StringAttribute{
				MarkdownDescription: `portId path parameter. Port ID`,
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: `The type of the port: 'access' or 'trunk'.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"vlan": schema.Int64Attribute{
				MarkdownDescription: `Native VLAN when the port is in Trunk mode. Access VLAN when the port is in Access mode.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

//path params to set ['portId']

func (r *NetworksAppliancePortsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksAppliancePortsRs

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
	vvPortID := data.PortID.ValueString()
	//Has Item and has items and not post

	if vvNetworkID != "" && vvPortID != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkAppliancePort(vvNetworkID, vvPortID)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksAppliancePorts  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksAppliancePorts only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkAppliancePort(vvNetworkID, vvPortID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkAppliancePort",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkAppliancePort",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Appliance.GetNetworkAppliancePort(vvNetworkID, vvPortID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkAppliancePort",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkAppliancePort",
			err.Error(),
		)
		return
	}

	data = ResponseApplianceGetNetworkAppliancePortItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *NetworksAppliancePortsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksAppliancePortsRs

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
	vvPortID := data.PortID.ValueString()
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkAppliancePort(vvNetworkID, vvPortID)
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
				"Failure when executing GetNetworkAppliancePort",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkAppliancePort",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkAppliancePortItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksAppliancePortsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("port_id"), idParts[1])...)
}

func (r *NetworksAppliancePortsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksAppliancePortsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvPortID := data.PortID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkAppliancePort(vvNetworkID, vvPortID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkAppliancePort",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkAppliancePort",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksAppliancePortsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksAppliancePorts", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksAppliancePortsRs struct {
	NetworkID           types.String `tfsdk:"network_id"`
	PortID              types.String `tfsdk:"port_id"`
	AccessPolicy        types.String `tfsdk:"access_policy"`
	AllowedVLANs        types.String `tfsdk:"allowed_vlans"`
	DropUntaggedTraffic types.Bool   `tfsdk:"drop_untagged_traffic"`
	Enabled             types.Bool   `tfsdk:"enabled"`
	Number              types.Int64  `tfsdk:"number"`
	Type                types.String `tfsdk:"type"`
	VLAN                types.Int64  `tfsdk:"vlan"`
}

// FromBody
func (r *NetworksAppliancePortsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkAppliancePort {
	emptyString := ""
	accessPolicy := new(string)
	if !r.AccessPolicy.IsUnknown() && !r.AccessPolicy.IsNull() {
		*accessPolicy = r.AccessPolicy.ValueString()
	} else {
		accessPolicy = &emptyString
	}
	allowedVLANs := new(string)
	if !r.AllowedVLANs.IsUnknown() && !r.AllowedVLANs.IsNull() {
		*allowedVLANs = r.AllowedVLANs.ValueString()
	} else {
		allowedVLANs = &emptyString
	}
	dropUntaggedTraffic := new(bool)
	if !r.DropUntaggedTraffic.IsUnknown() && !r.DropUntaggedTraffic.IsNull() {
		*dropUntaggedTraffic = r.DropUntaggedTraffic.ValueBool()
	} else {
		dropUntaggedTraffic = nil
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	typeR := new(string)
	if !r.Type.IsUnknown() && !r.Type.IsNull() {
		*typeR = r.Type.ValueString()
	} else {
		typeR = &emptyString
	}
	vLAN := new(int64)
	if !r.VLAN.IsUnknown() && !r.VLAN.IsNull() {
		*vLAN = r.VLAN.ValueInt64()
	} else {
		vLAN = nil
	}
	out := merakigosdk.RequestApplianceUpdateNetworkAppliancePort{
		AccessPolicy:        *accessPolicy,
		AllowedVLANs:        *allowedVLANs,
		DropUntaggedTraffic: dropUntaggedTraffic,
		Enabled:             enabled,
		Type:                *typeR,
		VLAN:                int64ToIntPointer(vLAN),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkAppliancePortItemToBodyRs(state NetworksAppliancePortsRs, response *merakigosdk.ResponseApplianceGetNetworkAppliancePort, is_read bool) NetworksAppliancePortsRs {
	itemState := NetworksAppliancePortsRs{
		AccessPolicy: types.StringValue(response.AccessPolicy),
		AllowedVLANs: types.StringValue(response.AllowedVLANs),
		DropUntaggedTraffic: func() types.Bool {
			if response.DropUntaggedTraffic != nil {
				return types.BoolValue(*response.DropUntaggedTraffic)
			}
			return types.Bool{}
		}(),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		Number: func() types.Int64 {
			if response.Number != nil {
				return types.Int64Value(int64(*response.Number))
			}
			return types.Int64{}
		}(),
		Type: types.StringValue(response.Type),
		VLAN: func() types.Int64 {
			if response.VLAN != nil {
				return types.Int64Value(int64(*response.VLAN))
			}
			return types.Int64{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksAppliancePortsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksAppliancePortsRs)
}
