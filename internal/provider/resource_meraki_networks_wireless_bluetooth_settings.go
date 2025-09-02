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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWirelessBluetoothSettingsResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessBluetoothSettingsResource{}
)

func NewNetworksWirelessBluetoothSettingsResource() resource.Resource {
	return &NetworksWirelessBluetoothSettingsResource{}
}

type NetworksWirelessBluetoothSettingsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessBluetoothSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessBluetoothSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_bluetooth_settings"
}

func (r *NetworksWirelessBluetoothSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"advertising_enabled": schema.BoolAttribute{
				MarkdownDescription: `Whether APs will advertise beacons.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"esl_enabled": schema.BoolAttribute{
				MarkdownDescription: `Whether ESL is enabled on this network.`,
				Computed:            true,
			},
			"major": schema.Int64Attribute{
				MarkdownDescription: `The major number to be used in the beacon identifier. Only valid in 'Non-unique' mode.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"major_minor_assignment_mode": schema.StringAttribute{
				MarkdownDescription: `The way major and minor number should be assigned to nodes in the network. ('Unique', 'Non-unique')
                                  Allowed values: [Non-unique,Unique]`,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Non-unique",
						"Unique",
					),
				},
			},
			"minor": schema.Int64Attribute{
				MarkdownDescription: `The minor number to be used in the beacon identifier. Only valid in 'Non-unique' mode.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"scanning_enabled": schema.BoolAttribute{
				MarkdownDescription: `Whether APs will scan for Bluetooth enabled clients.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"uuid": schema.StringAttribute{
				MarkdownDescription: `The UUID to be used in the beacon identifier.`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *NetworksWirelessBluetoothSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessBluetoothSettingsRs

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
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessBluetoothSettings(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessBluetoothSettings",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessBluetoothSettings",
			err.Error(),
		)
		return
	}

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *NetworksWirelessBluetoothSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessBluetoothSettingsRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessBluetoothSettings(vvNetworkID)
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
				"Failure when executing GetNetworkWirelessBluetoothSettings",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessBluetoothSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetNetworkWirelessBluetoothSettingsItemToBodyRs(data, responseGet, true)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksWirelessBluetoothSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksWirelessBluetoothSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksWirelessBluetoothSettingsRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessBluetoothSettings(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessBluetoothSettings",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessBluetoothSettings",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworksWirelessBluetoothSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksWirelessBluetoothSettings", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessBluetoothSettingsRs struct {
	NetworkID                types.String `tfsdk:"network_id"`
	AdvertisingEnabled       types.Bool   `tfsdk:"advertising_enabled"`
	EslEnabled               types.Bool   `tfsdk:"esl_enabled"`
	Major                    types.Int64  `tfsdk:"major"`
	MajorMinorAssignmentMode types.String `tfsdk:"major_minor_assignment_mode"`
	Minor                    types.Int64  `tfsdk:"minor"`
	ScanningEnabled          types.Bool   `tfsdk:"scanning_enabled"`
	UUID                     types.String `tfsdk:"uuid"`
}

// FromBody
func (r *NetworksWirelessBluetoothSettingsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessBluetoothSettings {
	emptyString := ""
	advertisingEnabled := new(bool)
	if !r.AdvertisingEnabled.IsUnknown() && !r.AdvertisingEnabled.IsNull() {
		*advertisingEnabled = r.AdvertisingEnabled.ValueBool()
	} else {
		advertisingEnabled = nil
	}
	major := new(int64)
	if !r.Major.IsUnknown() && !r.Major.IsNull() {
		*major = r.Major.ValueInt64()
	} else {
		major = nil
	}
	majorMinorAssignmentMode := new(string)
	if !r.MajorMinorAssignmentMode.IsUnknown() && !r.MajorMinorAssignmentMode.IsNull() {
		*majorMinorAssignmentMode = r.MajorMinorAssignmentMode.ValueString()
	} else {
		majorMinorAssignmentMode = &emptyString
	}
	minor := new(int64)
	if !r.Minor.IsUnknown() && !r.Minor.IsNull() {
		*minor = r.Minor.ValueInt64()
	} else {
		minor = nil
	}
	scanningEnabled := new(bool)
	if !r.ScanningEnabled.IsUnknown() && !r.ScanningEnabled.IsNull() {
		*scanningEnabled = r.ScanningEnabled.ValueBool()
	} else {
		scanningEnabled = nil
	}
	uUID := new(string)
	if !r.UUID.IsUnknown() && !r.UUID.IsNull() {
		*uUID = r.UUID.ValueString()
	} else {
		uUID = &emptyString
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessBluetoothSettings{
		AdvertisingEnabled:       advertisingEnabled,
		Major:                    int64ToIntPointer(major),
		MajorMinorAssignmentMode: *majorMinorAssignmentMode,
		Minor:                    int64ToIntPointer(minor),
		ScanningEnabled:          scanningEnabled,
		UUID:                     *uUID,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessBluetoothSettingsItemToBodyRs(state NetworksWirelessBluetoothSettingsRs, response *merakigosdk.ResponseWirelessGetNetworkWirelessBluetoothSettings, is_read bool) NetworksWirelessBluetoothSettingsRs {
	itemState := NetworksWirelessBluetoothSettingsRs{
		AdvertisingEnabled: func() types.Bool {
			if response.AdvertisingEnabled != nil {
				return types.BoolValue(*response.AdvertisingEnabled)
			}
			return types.Bool{}
		}(),
		EslEnabled: func() types.Bool {
			if response.EslEnabled != nil {
				return types.BoolValue(*response.EslEnabled)
			}
			return types.Bool{}
		}(),
		Major: func() types.Int64 {
			if response.Major != nil {
				return types.Int64Value(int64(*response.Major))
			}
			return types.Int64{}
		}(),
		MajorMinorAssignmentMode: func() types.String {
			if response.MajorMinorAssignmentMode != "" {
				return types.StringValue(response.MajorMinorAssignmentMode)
			}
			return types.String{}
		}(),
		Minor: func() types.Int64 {
			if response.Minor != nil {
				return types.Int64Value(int64(*response.Minor))
			}
			return types.Int64{}
		}(),
		ScanningEnabled: func() types.Bool {
			if response.ScanningEnabled != nil {
				return types.BoolValue(*response.ScanningEnabled)
			}
			return types.Bool{}
		}(),
		UUID: func() types.String {
			if response.UUID != "" {
				return types.StringValue(response.UUID)
			}
			return types.String{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessBluetoothSettingsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessBluetoothSettingsRs)
}
