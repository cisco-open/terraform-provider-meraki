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

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesApplianceRadioSettingsResource{}
	_ resource.ResourceWithConfigure = &DevicesApplianceRadioSettingsResource{}
)

func NewDevicesApplianceRadioSettingsResource() resource.Resource {
	return &DevicesApplianceRadioSettingsResource{}
}

type DevicesApplianceRadioSettingsResource struct {
	client *merakigosdk.Client
}

func (r *DevicesApplianceRadioSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesApplianceRadioSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_appliance_radio_settings"
}

func (r *DevicesApplianceRadioSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"five_ghz_settings": schema.SingleNestedAttribute{
				MarkdownDescription: `Manual radio settings for 5 GHz`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"channel": schema.Int64Attribute{
						MarkdownDescription: `Manual channel for 5 GHz
                                        Allowed values: [36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161,165,169,173,177]`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"channel_width": schema.Int64Attribute{
						MarkdownDescription: `Manual channel width for 5 GHz
                                        Allowed values: [0,20,40,80,160]`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"target_power": schema.Int64Attribute{
						MarkdownDescription: `Manual target power for 5 GHz`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"rf_profile_id": schema.StringAttribute{
				MarkdownDescription: `RF Profile ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `The device serial`,
				Required:            true,
			},
			"two_four_ghz_settings": schema.SingleNestedAttribute{
				MarkdownDescription: `Manual radio settings for 2.4 GHz`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"channel": schema.Int64Attribute{
						MarkdownDescription: `Manual channel for 2.4 GHz
                                        Allowed values: [1,2,3,4,5,6,7,8,9,10,11,12,13,14]`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"target_power": schema.Int64Attribute{
						MarkdownDescription: `Manual target power for 2.4 GHz`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
		},
	}
}

func (r *DevicesApplianceRadioSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesApplianceRadioSettingsRs

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
	//Item
	responseVerifyItem, restyResp1, err := r.client.Appliance.GetDeviceApplianceRadioSettings(vvSerial)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource DevicesApplianceRadioSettings only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource DevicesApplianceRadioSettings only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateDeviceApplianceRadioSettings(vvSerial, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceApplianceRadioSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceApplianceRadioSettings",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Appliance.GetDeviceApplianceRadioSettings(vvSerial)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceApplianceRadioSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceApplianceRadioSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetDeviceApplianceRadioSettingsItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesApplianceRadioSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesApplianceRadioSettingsRs

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
	responseGet, restyRespGet, err := r.client.Appliance.GetDeviceApplianceRadioSettings(vvSerial)
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
				"Failure when executing GetDeviceApplianceRadioSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceApplianceRadioSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetDeviceApplianceRadioSettingsItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *DevicesApplianceRadioSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), req.ID)...)
}

func (r *DevicesApplianceRadioSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesApplianceRadioSettingsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvSerial := data.Serial.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateDeviceApplianceRadioSettings(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceApplianceRadioSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceApplianceRadioSettings",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesApplianceRadioSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting DevicesApplianceRadioSettings", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesApplianceRadioSettingsRs struct {
	Serial             types.String                                                          `tfsdk:"serial"`
	FiveGhzSettings    *ResponseApplianceGetDeviceApplianceRadioSettingsFiveGhzSettingsRs    `tfsdk:"five_ghz_settings"`
	RfProfileID        types.String                                                          `tfsdk:"rf_profile_id"`
	TwoFourGhzSettings *ResponseApplianceGetDeviceApplianceRadioSettingsTwoFourGhzSettingsRs `tfsdk:"two_four_ghz_settings"`
}

type ResponseApplianceGetDeviceApplianceRadioSettingsFiveGhzSettingsRs struct {
	Channel      types.Int64 `tfsdk:"channel"`
	ChannelWidth types.Int64 `tfsdk:"channel_width"`
	TargetPower  types.Int64 `tfsdk:"target_power"`
}

type ResponseApplianceGetDeviceApplianceRadioSettingsTwoFourGhzSettingsRs struct {
	Channel     types.Int64 `tfsdk:"channel"`
	TargetPower types.Int64 `tfsdk:"target_power"`
}

// FromBody
func (r *DevicesApplianceRadioSettingsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateDeviceApplianceRadioSettings {
	emptyString := ""
	var requestApplianceUpdateDeviceApplianceRadioSettingsFiveGhzSettings *merakigosdk.RequestApplianceUpdateDeviceApplianceRadioSettingsFiveGhzSettings
	if r.FiveGhzSettings != nil {
		channel := func() *int64 {
			if !r.FiveGhzSettings.Channel.IsUnknown() && !r.FiveGhzSettings.Channel.IsNull() {
				return r.FiveGhzSettings.Channel.ValueInt64Pointer()
			}
			return nil
		}()
		channelWidth := func() *int64 {
			if !r.FiveGhzSettings.ChannelWidth.IsUnknown() && !r.FiveGhzSettings.ChannelWidth.IsNull() {
				return r.FiveGhzSettings.ChannelWidth.ValueInt64Pointer()
			}
			return nil
		}()
		targetPower := func() *int64 {
			if !r.FiveGhzSettings.TargetPower.IsUnknown() && !r.FiveGhzSettings.TargetPower.IsNull() {
				return r.FiveGhzSettings.TargetPower.ValueInt64Pointer()
			}
			return nil
		}()
		requestApplianceUpdateDeviceApplianceRadioSettingsFiveGhzSettings = &merakigosdk.RequestApplianceUpdateDeviceApplianceRadioSettingsFiveGhzSettings{
			Channel:      int64ToIntPointer(channel),
			ChannelWidth: int64ToIntPointer(channelWidth),
			TargetPower:  int64ToIntPointer(targetPower),
		}
	}
	rfProfileID := new(string)
	if !r.RfProfileID.IsUnknown() && !r.RfProfileID.IsNull() {
		*rfProfileID = r.RfProfileID.ValueString()
	} else {
		rfProfileID = &emptyString
	}
	var requestApplianceUpdateDeviceApplianceRadioSettingsTwoFourGhzSettings *merakigosdk.RequestApplianceUpdateDeviceApplianceRadioSettingsTwoFourGhzSettings
	if r.TwoFourGhzSettings != nil {
		channel := func() *int64 {
			if !r.TwoFourGhzSettings.Channel.IsUnknown() && !r.TwoFourGhzSettings.Channel.IsNull() {
				return r.TwoFourGhzSettings.Channel.ValueInt64Pointer()
			}
			return nil
		}()
		targetPower := func() *int64 {
			if !r.TwoFourGhzSettings.TargetPower.IsUnknown() && !r.TwoFourGhzSettings.TargetPower.IsNull() {
				return r.TwoFourGhzSettings.TargetPower.ValueInt64Pointer()
			}
			return nil
		}()
		requestApplianceUpdateDeviceApplianceRadioSettingsTwoFourGhzSettings = &merakigosdk.RequestApplianceUpdateDeviceApplianceRadioSettingsTwoFourGhzSettings{
			Channel:     int64ToIntPointer(channel),
			TargetPower: int64ToIntPointer(targetPower),
		}
	}
	out := merakigosdk.RequestApplianceUpdateDeviceApplianceRadioSettings{
		FiveGhzSettings:    requestApplianceUpdateDeviceApplianceRadioSettingsFiveGhzSettings,
		RfProfileID:        *rfProfileID,
		TwoFourGhzSettings: requestApplianceUpdateDeviceApplianceRadioSettingsTwoFourGhzSettings,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetDeviceApplianceRadioSettingsItemToBodyRs(state DevicesApplianceRadioSettingsRs, response *merakigosdk.ResponseApplianceGetDeviceApplianceRadioSettings, is_read bool) DevicesApplianceRadioSettingsRs {
	itemState := DevicesApplianceRadioSettingsRs{
		FiveGhzSettings: func() *ResponseApplianceGetDeviceApplianceRadioSettingsFiveGhzSettingsRs {
			if response.FiveGhzSettings != nil {
				return &ResponseApplianceGetDeviceApplianceRadioSettingsFiveGhzSettingsRs{
					Channel: func() types.Int64 {
						if response.FiveGhzSettings.Channel != nil {
							return types.Int64Value(int64(*response.FiveGhzSettings.Channel))
						}
						return types.Int64{}
					}(),
					ChannelWidth: func() types.Int64 {
						if response.FiveGhzSettings.ChannelWidth != nil {
							return types.Int64Value(int64(*response.FiveGhzSettings.ChannelWidth))
						}
						return types.Int64{}
					}(),
					TargetPower: func() types.Int64 {
						if response.FiveGhzSettings.TargetPower != nil {
							return types.Int64Value(int64(*response.FiveGhzSettings.TargetPower))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
		RfProfileID: types.StringValue(response.RfProfileID),
		Serial:      types.StringValue(response.Serial),
		TwoFourGhzSettings: func() *ResponseApplianceGetDeviceApplianceRadioSettingsTwoFourGhzSettingsRs {
			if response.TwoFourGhzSettings != nil {
				return &ResponseApplianceGetDeviceApplianceRadioSettingsTwoFourGhzSettingsRs{
					Channel: func() types.Int64 {
						if response.TwoFourGhzSettings.Channel != nil {
							return types.Int64Value(int64(*response.TwoFourGhzSettings.Channel))
						}
						return types.Int64{}
					}(),
					TargetPower: func() types.Int64 {
						if response.TwoFourGhzSettings.TargetPower != nil {
							return types.Int64Value(int64(*response.TwoFourGhzSettings.TargetPower))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesApplianceRadioSettingsRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesApplianceRadioSettingsRs)
}
