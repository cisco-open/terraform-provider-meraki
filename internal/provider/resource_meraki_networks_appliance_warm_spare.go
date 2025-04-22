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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceWarmSpareResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceWarmSpareResource{}
)

func NewNetworksApplianceWarmSpareResource() resource.Resource {
	return &NetworksApplianceWarmSpareResource{}
}

type NetworksApplianceWarmSpareResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceWarmSpareResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceWarmSpareResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_warm_spare"
}

func (r *NetworksApplianceWarmSpareResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `Is the warm spare enabled`,
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
			"primary_serial": schema.StringAttribute{
				MarkdownDescription: `Serial number of the primary appliance`,
				Computed:            true,
			},
			"spare_serial": schema.StringAttribute{
				MarkdownDescription: `Serial number of the warm spare appliance`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"uplink_mode": schema.StringAttribute{
				MarkdownDescription: `Uplink mode, either virtual or public`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"virtual_ip1": schema.StringAttribute{
				MarkdownDescription: `The WAN 1 shared IP`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"virtual_ip2": schema.StringAttribute{
				MarkdownDescription: `The WAN 2 shared IP`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"wan1": schema.SingleNestedAttribute{
				MarkdownDescription: `WAN 1 IP and subnet`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{

					"ip": schema.StringAttribute{
						MarkdownDescription: `IP address used for WAN 1`,
						Computed:            true,
					},
					"subnet": schema.StringAttribute{
						MarkdownDescription: `Subnet used for WAN 1`,
						Computed:            true,
					},
				},
			},
			"wan2": schema.SingleNestedAttribute{
				MarkdownDescription: `WAN 2 IP and subnet`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{

					"ip": schema.StringAttribute{
						MarkdownDescription: `IP address used for WAN 2`,
						Computed:            true,
					},
					"subnet": schema.StringAttribute{
						MarkdownDescription: `Subnet used for WAN 2`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (r *NetworksApplianceWarmSpareResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceWarmSpareRs

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
		responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceWarmSpare(vvNetworkID)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksApplianceWarmSpare  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksApplianceWarmSpare only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceWarmSpare(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceWarmSpare",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceWarmSpare",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceWarmSpare(vvNetworkID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceWarmSpare",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceWarmSpare",
			err.Error(),
		)
		return
	}

	data = ResponseApplianceGetNetworkApplianceWarmSpareItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *NetworksApplianceWarmSpareResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceWarmSpareRs

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
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceWarmSpare(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceWarmSpare",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceWarmSpare",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceWarmSpareItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksApplianceWarmSpareResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceWarmSpareResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceWarmSpareRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceWarmSpare(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceWarmSpare",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceWarmSpare",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceWarmSpareResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceWarmSpare", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceWarmSpareRs struct {
	NetworkID     types.String                                         `tfsdk:"network_id"`
	Enabled       types.Bool                                           `tfsdk:"enabled"`
	PrimarySerial types.String                                         `tfsdk:"primary_serial"`
	SpareSerial   types.String                                         `tfsdk:"spare_serial"`
	UplinkMode    types.String                                         `tfsdk:"uplink_mode"`
	Wan1          *ResponseApplianceGetNetworkApplianceWarmSpareWan1Rs `tfsdk:"wan1"`
	Wan2          *ResponseApplianceGetNetworkApplianceWarmSpareWan2Rs `tfsdk:"wan2"`
	VirtualIP1    types.String                                         `tfsdk:"virtual_ip1"`
	VirtualIP2    types.String                                         `tfsdk:"virtual_ip2"`
}

type ResponseApplianceGetNetworkApplianceWarmSpareWan1Rs struct {
	IP     types.String `tfsdk:"ip"`
	Subnet types.String `tfsdk:"subnet"`
}

type ResponseApplianceGetNetworkApplianceWarmSpareWan2Rs struct {
	IP     types.String `tfsdk:"ip"`
	Subnet types.String `tfsdk:"subnet"`
}

// FromBody
func (r *NetworksApplianceWarmSpareRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceWarmSpare {
	emptyString := ""
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	spareSerial := new(string)
	if !r.SpareSerial.IsUnknown() && !r.SpareSerial.IsNull() {
		*spareSerial = r.SpareSerial.ValueString()
	} else {
		spareSerial = &emptyString
	}
	uplinkMode := new(string)
	if !r.UplinkMode.IsUnknown() && !r.UplinkMode.IsNull() {
		*uplinkMode = r.UplinkMode.ValueString()
	} else {
		uplinkMode = &emptyString
	}
	virtualIP1 := new(string)
	if !r.VirtualIP1.IsUnknown() && !r.VirtualIP1.IsNull() {
		*virtualIP1 = r.VirtualIP1.ValueString()
	} else {
		virtualIP1 = &emptyString
	}
	virtualIP2 := new(string)
	if !r.VirtualIP2.IsUnknown() && !r.VirtualIP2.IsNull() {
		*virtualIP2 = r.VirtualIP2.ValueString()
	} else {
		virtualIP2 = &emptyString
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceWarmSpare{
		Enabled:     enabled,
		SpareSerial: *spareSerial,
		UplinkMode:  *uplinkMode,
		VirtualIP1:  *virtualIP1,
		VirtualIP2:  *virtualIP2,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceWarmSpareItemToBodyRs(state NetworksApplianceWarmSpareRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceWarmSpare, is_read bool) NetworksApplianceWarmSpareRs {
	itemState := NetworksApplianceWarmSpareRs{
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		PrimarySerial: types.StringValue(response.PrimarySerial),
		SpareSerial:   types.StringValue(response.SpareSerial),
		UplinkMode:    types.StringValue(response.UplinkMode),
		Wan1: func() *ResponseApplianceGetNetworkApplianceWarmSpareWan1Rs {
			if response.Wan1 != nil {
				return &ResponseApplianceGetNetworkApplianceWarmSpareWan1Rs{
					IP:     types.StringValue(response.Wan1.IP),
					Subnet: types.StringValue(response.Wan1.Subnet),
				}
			}
			return nil
		}(),
		Wan2: func() *ResponseApplianceGetNetworkApplianceWarmSpareWan2Rs {
			if response.Wan2 != nil {
				return &ResponseApplianceGetNetworkApplianceWarmSpareWan2Rs{
					IP:     types.StringValue(response.Wan2.IP),
					Subnet: types.StringValue(response.Wan2.Subnet),
				}
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceWarmSpareRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceWarmSpareRs)
}
