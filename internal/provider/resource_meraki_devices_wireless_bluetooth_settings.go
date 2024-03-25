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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesWirelessBluetoothSettingsResource{}
	_ resource.ResourceWithConfigure = &DevicesWirelessBluetoothSettingsResource{}
)

func NewDevicesWirelessBluetoothSettingsResource() resource.Resource {
	return &DevicesWirelessBluetoothSettingsResource{}
}

type DevicesWirelessBluetoothSettingsResource struct {
	client *merakigosdk.Client
}

func (r *DevicesWirelessBluetoothSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesWirelessBluetoothSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_wireless_bluetooth_settings"
}

func (r *DevicesWirelessBluetoothSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"major": schema.Int64Attribute{
				MarkdownDescription: `Desired major value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"minor": schema.Int64Attribute{
				MarkdownDescription: `Desired minor value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"uuid": schema.StringAttribute{
				MarkdownDescription: `Desired UUID of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *DevicesWirelessBluetoothSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesWirelessBluetoothSettingsRs

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
	vvSerial := data.Serial.ValueString()
	// serial
	//Item
	responseVerifyItem, restyResp1, err := r.client.Wireless.GetDeviceWirelessBluetoothSettings(vvSerial)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource DevicesWirelessBluetoothSettings only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource DevicesWirelessBluetoothSettings only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateDeviceWirelessBluetoothSettings(vvSerial, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceWirelessBluetoothSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceWirelessBluetoothSettings",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Wireless.GetDeviceWirelessBluetoothSettings(vvSerial)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceWirelessBluetoothSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceWirelessBluetoothSettings",
			err.Error(),
		)
		return
	}

	data = ResponseWirelessGetDeviceWirelessBluetoothSettingsItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesWirelessBluetoothSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesWirelessBluetoothSettingsRs

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
	// serial
	responseGet, restyRespGet, err := r.client.Wireless.GetDeviceWirelessBluetoothSettings(vvSerial)
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
				"Failure when executing GetDeviceWirelessBluetoothSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceWirelessBluetoothSettings",
			err.Error(),
		)
		return
	}

	data = ResponseWirelessGetDeviceWirelessBluetoothSettingsItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *DevicesWirelessBluetoothSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), req.ID)...)
}

func (r *DevicesWirelessBluetoothSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesWirelessBluetoothSettingsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvSerial := data.Serial.ValueString()
	// serial
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateDeviceWirelessBluetoothSettings(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceWirelessBluetoothSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceWirelessBluetoothSettings",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesWirelessBluetoothSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesWirelessBluetoothSettingsRs struct {
	Serial types.String `tfsdk:"serial"`
	Major  types.Int64  `tfsdk:"major"`
	Minor  types.Int64  `tfsdk:"minor"`
	UUID   types.String `tfsdk:"uuid"`
}

// FromBody
func (r *DevicesWirelessBluetoothSettingsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateDeviceWirelessBluetoothSettings {
	emptyString := ""
	major := new(int64)
	if !r.Major.IsUnknown() && !r.Major.IsNull() {
		*major = r.Major.ValueInt64()
	} else {
		major = nil
	}
	minor := new(int64)
	if !r.Minor.IsUnknown() && !r.Minor.IsNull() {
		*minor = r.Minor.ValueInt64()
	} else {
		minor = nil
	}
	uUID := new(string)
	if !r.UUID.IsUnknown() && !r.UUID.IsNull() {
		*uUID = r.UUID.ValueString()
	} else {
		uUID = &emptyString
	}
	out := merakigosdk.RequestWirelessUpdateDeviceWirelessBluetoothSettings{
		Major: int64ToIntPointer(major),
		Minor: int64ToIntPointer(minor),
		UUID:  *uUID,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetDeviceWirelessBluetoothSettingsItemToBodyRs(state DevicesWirelessBluetoothSettingsRs, response *merakigosdk.ResponseWirelessGetDeviceWirelessBluetoothSettings, is_read bool) DevicesWirelessBluetoothSettingsRs {
	itemState := DevicesWirelessBluetoothSettingsRs{
		Major: func() types.Int64 {
			if response.Major != nil {
				return types.Int64Value(int64(*response.Major))
			}
			return types.Int64{}
		}(),
		Minor: func() types.Int64 {
			if response.Minor != nil {
				return types.Int64Value(int64(*response.Minor))
			}
			return types.Int64{}
		}(),
		UUID: types.StringValue(response.UUID),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesWirelessBluetoothSettingsRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesWirelessBluetoothSettingsRs)
}
