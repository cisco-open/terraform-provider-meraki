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
	_ resource.Resource              = &DevicesWirelessElectronicShelfLabelResource{}
	_ resource.ResourceWithConfigure = &DevicesWirelessElectronicShelfLabelResource{}
)

func NewDevicesWirelessElectronicShelfLabelResource() resource.Resource {
	return &DevicesWirelessElectronicShelfLabelResource{}
}

type DevicesWirelessElectronicShelfLabelResource struct {
	client *merakigosdk.Client
}

func (r *DevicesWirelessElectronicShelfLabelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesWirelessElectronicShelfLabelResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_wireless_electronic_shelf_label"
}

func (r *DevicesWirelessElectronicShelfLabelResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ap_esl_id": schema.Int64Attribute{
				MarkdownDescription: `An identifier for the device used by the ESL system`,
				Computed:            true,
			},
			"channel": schema.StringAttribute{
				MarkdownDescription: `Desired ESL channel for the device, or 'Auto' (case insensitive) to use the recommended channel`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `Turn ESL features on and off for this device`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"hostname": schema.StringAttribute{
				MarkdownDescription: `Hostname of the ESL management service`,
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `The identifier for the device's network`,
				Computed:            true,
			},
			"provider_r": schema.StringAttribute{
				MarkdownDescription: `The service providing ESL functionality`,
				Computed:            true,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `The serial number of the device`,
				Required:            true,
			},
		},
	}
}

func (r *DevicesWirelessElectronicShelfLabelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesWirelessElectronicShelfLabelRs

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
	vvSerial := data.Serial.ValueString()
	//Has Item and not has items

	if vvSerial != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Wireless.GetDeviceWirelessElectronicShelfLabel(vvSerial)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource DevicesWirelessElectronicShelfLabel  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource DevicesWirelessElectronicShelfLabel only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateDeviceWirelessElectronicShelfLabel(vvSerial, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceWirelessElectronicShelfLabel",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceWirelessElectronicShelfLabel",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Wireless.GetDeviceWirelessElectronicShelfLabel(vvSerial)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceWirelessElectronicShelfLabel",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceWirelessElectronicShelfLabel",
			err.Error(),
		)
		return
	}

	data = ResponseWirelessGetDeviceWirelessElectronicShelfLabelItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *DevicesWirelessElectronicShelfLabelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesWirelessElectronicShelfLabelRs

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

	vvSerial := data.Serial.ValueString()
	responseGet, restyRespGet, err := r.client.Wireless.GetDeviceWirelessElectronicShelfLabel(vvSerial)
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
				"Failure when executing GetDeviceWirelessElectronicShelfLabel",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceWirelessElectronicShelfLabel",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetDeviceWirelessElectronicShelfLabelItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *DevicesWirelessElectronicShelfLabelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), req.ID)...)
}

func (r *DevicesWirelessElectronicShelfLabelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesWirelessElectronicShelfLabelRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvSerial := data.Serial.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateDeviceWirelessElectronicShelfLabel(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceWirelessElectronicShelfLabel",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceWirelessElectronicShelfLabel",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesWirelessElectronicShelfLabelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting DevicesWirelessElectronicShelfLabel", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesWirelessElectronicShelfLabelRs struct {
	Serial    types.String `tfsdk:"serial"`
	ApEslID   types.Int64  `tfsdk:"ap_esl_id"`
	Channel   types.String `tfsdk:"channel"`
	Enabled   types.Bool   `tfsdk:"enabled"`
	Hostname  types.String `tfsdk:"hostname"`
	NetworkID types.String `tfsdk:"network_id"`
	Provider  types.String `tfsdk:"provider_r"`
}

// FromBody
func (r *DevicesWirelessElectronicShelfLabelRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateDeviceWirelessElectronicShelfLabel {
	emptyString := ""
	channel := new(string)
	if !r.Channel.IsUnknown() && !r.Channel.IsNull() {
		*channel = r.Channel.ValueString()
	} else {
		channel = &emptyString
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	out := merakigosdk.RequestWirelessUpdateDeviceWirelessElectronicShelfLabel{
		Channel: *channel,
		Enabled: enabled,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetDeviceWirelessElectronicShelfLabelItemToBodyRs(state DevicesWirelessElectronicShelfLabelRs, response *merakigosdk.ResponseWirelessGetDeviceWirelessElectronicShelfLabel, is_read bool) DevicesWirelessElectronicShelfLabelRs {
	itemState := DevicesWirelessElectronicShelfLabelRs{
		ApEslID: func() types.Int64 {
			if response.ApEslID != nil {
				return types.Int64Value(int64(*response.ApEslID))
			}
			return types.Int64{}
		}(),
		Channel: types.StringValue(response.Channel),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		Hostname:  types.StringValue(response.Hostname),
		NetworkID: types.StringValue(response.NetworkID),
		Provider:  types.StringValue(response.Provider),
		Serial:    types.StringValue(response.Serial),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesWirelessElectronicShelfLabelRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesWirelessElectronicShelfLabelRs)
}
