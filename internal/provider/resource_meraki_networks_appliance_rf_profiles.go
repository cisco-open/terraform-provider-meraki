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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceRfProfilesResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceRfProfilesResource{}
)

func NewNetworksApplianceRfProfilesResource() resource.Resource {
	return &NetworksApplianceRfProfilesResource{}
}

type NetworksApplianceRfProfilesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceRfProfilesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceRfProfilesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_rf_profiles"
}

func (r *NetworksApplianceRfProfilesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"five_ghz_settings": schema.SingleNestedAttribute{
				MarkdownDescription: `Settings related to 5Ghz band.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"ax_enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether ax radio on 5Ghz band is on or off.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"min_bitrate": schema.Int64Attribute{
						MarkdownDescription: `Min bitrate (Mbps) of 2.4Ghz band.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: `ID of the RF Profile.`,
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name of the profile.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `ID of network this RF Profile belongs in.`,
				Required:            true,
			},
			"per_ssid_settings": schema.SingleNestedAttribute{
				MarkdownDescription: `Per-SSID radio settings by number.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"status_1": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for SSID 1.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Band mode of this SSID
                                              Allowed values: [2.4ghz,5ghz,6ghz,dual,multi]`,
								Computed: true,
								Optional: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"2.4ghz",
										"5ghz",
										"6ghz",
										"dual",
										"multi",
									),
								},
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"status_2": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for SSID 2.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Band mode of this SSID
                                              Allowed values: [2.4ghz,5ghz,6ghz,dual,multi]`,
								Computed: true,
								Optional: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"2.4ghz",
										"5ghz",
										"6ghz",
										"dual",
										"multi",
									),
								},
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"status_3": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for SSID 3.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Band mode of this SSID
                                              Allowed values: [2.4ghz,5ghz,6ghz,dual,multi]`,
								Computed: true,
								Optional: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"2.4ghz",
										"5ghz",
										"6ghz",
										"dual",
										"multi",
									),
								},
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"status_4": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for SSID 4.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Band mode of this SSID
                                              Allowed values: [2.4ghz,5ghz,6ghz,dual,multi]`,
								Computed: true,
								Optional: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"2.4ghz",
										"5ghz",
										"6ghz",
										"dual",
										"multi",
									),
								},
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
				},
			},
			"rf_profile_id": schema.StringAttribute{
				MarkdownDescription: `rfProfileId path parameter. Rf profile ID`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"two_four_ghz_settings": schema.SingleNestedAttribute{
				MarkdownDescription: `Settings related to 2.4Ghz band.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"ax_enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether ax radio on 2.4Ghz band is on or off.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"min_bitrate": schema.Float64Attribute{
						MarkdownDescription: `Min bitrate (Mbps) of 2.4Ghz band.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
		},
	}
}

func (r *NetworksApplianceRfProfilesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceRfProfilesRs

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
	vvName := data.Name.ValueString()
	//Item
	responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceRfProfiles(vvNetworkID)
	//Have Create
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkApplianceRfProfiles",
					err.Error(),
				)
				return
			}
		}
	}

	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem.Assigned)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvRfProfileID, ok := result2["ID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter vvRfProfileID",
					"Error",
				)
				return
			}
			r.client.Appliance.UpdateNetworkApplianceRfProfile(vvNetworkID, vvRfProfileID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Appliance.GetNetworkApplianceRfProfile(vvNetworkID, vvRfProfileID)
			if responseVerifyItem2 != nil {
				data = ResponseApplianceGetNetworkApplianceRfProfileItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Appliance.CreateNetworkApplianceRfProfile(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkApplianceRfProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkApplianceRfProfile",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceRfProfiles(vvNetworkID)
	// Has item and not has items
	responseStruct := structToMap(responseGet.Assigned)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvRfProfileID, ok := result2["ID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter vvRfProfileID",
				"Error",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Appliance.GetNetworkApplianceRfProfile(vvNetworkID, vvRfProfileID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseApplianceGetNetworkApplianceRfProfileItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkApplianceRfProfile",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceRfProfile",
				err.Error(),
			)
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Error in result.",
			"Error in result.",
		)
		return
	}
}

func (r *NetworksApplianceRfProfilesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceRfProfilesRs

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
	vvID := data.ID.ValueString()
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceRfProfile(vvNetworkID, vvID)
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
				"Failure when executing GetNetworkApplianceRfProfiles",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceRfProfiles",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceRfProfileItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksApplianceRfProfilesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
}

func (r *NetworksApplianceRfProfilesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceRfProfilesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvRfProfileID := data.ID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceRfProfile(vvNetworkID, vvRfProfileID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceRfProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceRfProfile",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceRfProfilesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksApplianceRfProfilesRs
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
	vvRfProfileID := state.RfProfileID.ValueString()
	_, err := r.client.Appliance.DeleteNetworkApplianceRfProfile(vvNetworkID, vvRfProfileID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkApplianceRfProfile", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksApplianceRfProfilesRs struct {
	NetworkID          types.String                                                       `tfsdk:"network_id"`
	RfProfileID        types.String                                                       `tfsdk:"rf_profile_id"`
	FiveGhzSettings    *ResponseApplianceGetNetworkApplianceRfProfileFiveGhzSettingsRs    `tfsdk:"five_ghz_settings"`
	ID                 types.String                                                       `tfsdk:"id"`
	Name               types.String                                                       `tfsdk:"name"`
	PerSSIDSettings    *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettingsRs    `tfsdk:"per_ssid_settings"`
	TwoFourGhzSettings *ResponseApplianceGetNetworkApplianceRfProfileTwoFourGhzSettingsRs `tfsdk:"two_four_ghz_settings"`
}

type ResponseApplianceGetNetworkApplianceRfProfileFiveGhzSettingsRs struct {
	AxEnabled  types.Bool  `tfsdk:"ax_enabled"`
	MinBitrate types.Int64 `tfsdk:"min_bitrate"`
}

type ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettingsRs struct {
	Status1 *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings1Rs `tfsdk:"status_1"`
	Status2 *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings2Rs `tfsdk:"status_2"`
	Status3 *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings3Rs `tfsdk:"status_3"`
	Status4 *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings4Rs `tfsdk:"status_4"`
}

type ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings1Rs struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
}

type ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings2Rs struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
}

type ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings3Rs struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
}

type ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings4Rs struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
}

type ResponseApplianceGetNetworkApplianceRfProfileTwoFourGhzSettingsRs struct {
	AxEnabled  types.Bool    `tfsdk:"ax_enabled"`
	MinBitrate types.Float64 `tfsdk:"min_bitrate"`
}

// FromBody
func (r *NetworksApplianceRfProfilesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestApplianceCreateNetworkApplianceRfProfile {
	emptyString := ""
	var requestApplianceCreateNetworkApplianceRfProfileFiveGhzSettings *merakigosdk.RequestApplianceCreateNetworkApplianceRfProfileFiveGhzSettings
	if r.FiveGhzSettings != nil {
		axEnabled := func() *bool {
			if !r.FiveGhzSettings.AxEnabled.IsUnknown() && !r.FiveGhzSettings.AxEnabled.IsNull() {
				return r.FiveGhzSettings.AxEnabled.ValueBoolPointer()
			}
			return nil
		}()
		minBitrate := func() *int64 {
			if !r.FiveGhzSettings.MinBitrate.IsUnknown() && !r.FiveGhzSettings.MinBitrate.IsNull() {
				return r.FiveGhzSettings.MinBitrate.ValueInt64Pointer()
			}
			return nil
		}()
		requestApplianceCreateNetworkApplianceRfProfileFiveGhzSettings = &merakigosdk.RequestApplianceCreateNetworkApplianceRfProfileFiveGhzSettings{
			AxEnabled:  axEnabled,
			MinBitrate: int64ToIntPointer(minBitrate),
		}
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings *merakigosdk.RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings
	if r.PerSSIDSettings != nil {
		var requestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings1 *merakigosdk.RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings1
		if r.PerSSIDSettings.Status1 != nil {
			bandOperationMode := r.PerSSIDSettings.Status1.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status1.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status1.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status1.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			requestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings1 = &merakigosdk.RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings1{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
			}
		}
		var requestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings2 *merakigosdk.RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings2
		if r.PerSSIDSettings.Status2 != nil {
			bandOperationMode := r.PerSSIDSettings.Status2.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status2.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status2.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status2.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			requestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings2 = &merakigosdk.RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings2{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
			}
		}
		var requestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings3 *merakigosdk.RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings3
		if r.PerSSIDSettings.Status3 != nil {
			bandOperationMode := r.PerSSIDSettings.Status3.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status3.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status3.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status3.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			requestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings3 = &merakigosdk.RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings3{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
			}
		}
		var requestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings4 *merakigosdk.RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings4
		if r.PerSSIDSettings.Status4 != nil {
			bandOperationMode := r.PerSSIDSettings.Status4.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status4.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status4.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status4.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			requestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings4 = &merakigosdk.RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings4{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
			}
		}
		requestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings = &merakigosdk.RequestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings{
			Status1: requestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings1,
			Status2: requestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings2,
			Status3: requestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings3,
			Status4: requestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings4,
		}
	}
	var requestApplianceCreateNetworkApplianceRfProfileTwoFourGhzSettings *merakigosdk.RequestApplianceCreateNetworkApplianceRfProfileTwoFourGhzSettings
	if r.TwoFourGhzSettings != nil {
		axEnabled := func() *bool {
			if !r.TwoFourGhzSettings.AxEnabled.IsUnknown() && !r.TwoFourGhzSettings.AxEnabled.IsNull() {
				return r.TwoFourGhzSettings.AxEnabled.ValueBoolPointer()
			}
			return nil
		}()
		minBitrate := func() *float64 {
			if !r.TwoFourGhzSettings.MinBitrate.IsUnknown() && !r.TwoFourGhzSettings.MinBitrate.IsNull() {
				return r.TwoFourGhzSettings.MinBitrate.ValueFloat64Pointer()
			}
			return nil
		}()
		requestApplianceCreateNetworkApplianceRfProfileTwoFourGhzSettings = &merakigosdk.RequestApplianceCreateNetworkApplianceRfProfileTwoFourGhzSettings{
			AxEnabled:  axEnabled,
			MinBitrate: minBitrate,
		}
	}
	out := merakigosdk.RequestApplianceCreateNetworkApplianceRfProfile{
		FiveGhzSettings:    requestApplianceCreateNetworkApplianceRfProfileFiveGhzSettings,
		Name:               *name,
		PerSSIDSettings:    requestApplianceCreateNetworkApplianceRfProfilePerSSIDSettings,
		TwoFourGhzSettings: requestApplianceCreateNetworkApplianceRfProfileTwoFourGhzSettings,
	}
	return &out
}
func (r *NetworksApplianceRfProfilesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceRfProfile {
	emptyString := ""
	var requestApplianceUpdateNetworkApplianceRfProfileFiveGhzSettings *merakigosdk.RequestApplianceUpdateNetworkApplianceRfProfileFiveGhzSettings
	if r.FiveGhzSettings != nil {
		axEnabled := func() *bool {
			if !r.FiveGhzSettings.AxEnabled.IsUnknown() && !r.FiveGhzSettings.AxEnabled.IsNull() {
				return r.FiveGhzSettings.AxEnabled.ValueBoolPointer()
			}
			return nil
		}()
		minBitrate := func() *int64 {
			if !r.FiveGhzSettings.MinBitrate.IsUnknown() && !r.FiveGhzSettings.MinBitrate.IsNull() {
				return r.FiveGhzSettings.MinBitrate.ValueInt64Pointer()
			}
			return nil
		}()
		requestApplianceUpdateNetworkApplianceRfProfileFiveGhzSettings = &merakigosdk.RequestApplianceUpdateNetworkApplianceRfProfileFiveGhzSettings{
			AxEnabled:  axEnabled,
			MinBitrate: int64ToIntPointer(minBitrate),
		}
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings *merakigosdk.RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings
	if r.PerSSIDSettings != nil {
		var requestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings1 *merakigosdk.RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings1
		if r.PerSSIDSettings.Status1 != nil {
			bandOperationMode := r.PerSSIDSettings.Status1.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status1.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status1.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status1.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			requestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings1 = &merakigosdk.RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings1{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
			}
		}
		var requestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings2 *merakigosdk.RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings2
		if r.PerSSIDSettings.Status2 != nil {
			bandOperationMode := r.PerSSIDSettings.Status2.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status2.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status2.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status2.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			requestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings2 = &merakigosdk.RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings2{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
			}
		}
		var requestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings3 *merakigosdk.RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings3
		if r.PerSSIDSettings.Status3 != nil {
			bandOperationMode := r.PerSSIDSettings.Status3.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status3.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status3.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status3.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			requestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings3 = &merakigosdk.RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings3{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
			}
		}
		var requestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings4 *merakigosdk.RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings4
		if r.PerSSIDSettings.Status4 != nil {
			bandOperationMode := r.PerSSIDSettings.Status4.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status4.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status4.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status4.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			requestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings4 = &merakigosdk.RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings4{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
			}
		}
		requestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings = &merakigosdk.RequestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings{
			Status1: requestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings1,
			Status2: requestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings2,
			Status3: requestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings3,
			Status4: requestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings4,
		}
	}
	var requestApplianceUpdateNetworkApplianceRfProfileTwoFourGhzSettings *merakigosdk.RequestApplianceUpdateNetworkApplianceRfProfileTwoFourGhzSettings
	if r.TwoFourGhzSettings != nil {
		axEnabled := func() *bool {
			if !r.TwoFourGhzSettings.AxEnabled.IsUnknown() && !r.TwoFourGhzSettings.AxEnabled.IsNull() {
				return r.TwoFourGhzSettings.AxEnabled.ValueBoolPointer()
			}
			return nil
		}()
		minBitrate := func() *float64 {
			if !r.TwoFourGhzSettings.MinBitrate.IsUnknown() && !r.TwoFourGhzSettings.MinBitrate.IsNull() {
				return r.TwoFourGhzSettings.MinBitrate.ValueFloat64Pointer()
			}
			return nil
		}()
		requestApplianceUpdateNetworkApplianceRfProfileTwoFourGhzSettings = &merakigosdk.RequestApplianceUpdateNetworkApplianceRfProfileTwoFourGhzSettings{
			AxEnabled:  axEnabled,
			MinBitrate: minBitrate,
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceRfProfile{
		FiveGhzSettings:    requestApplianceUpdateNetworkApplianceRfProfileFiveGhzSettings,
		Name:               *name,
		PerSSIDSettings:    requestApplianceUpdateNetworkApplianceRfProfilePerSSIDSettings,
		TwoFourGhzSettings: requestApplianceUpdateNetworkApplianceRfProfileTwoFourGhzSettings,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceRfProfileItemToBodyRs(state NetworksApplianceRfProfilesRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceRfProfile, is_read bool) NetworksApplianceRfProfilesRs {
	itemState := NetworksApplianceRfProfilesRs{
		FiveGhzSettings: func() *ResponseApplianceGetNetworkApplianceRfProfileFiveGhzSettingsRs {
			if response.FiveGhzSettings != nil {
				return &ResponseApplianceGetNetworkApplianceRfProfileFiveGhzSettingsRs{
					AxEnabled: func() types.Bool {
						if response.FiveGhzSettings.AxEnabled != nil {
							return types.BoolValue(*response.FiveGhzSettings.AxEnabled)
						}
						return types.Bool{}
					}(),
					MinBitrate: func() types.Int64 {
						if response.FiveGhzSettings.MinBitrate != nil {
							return types.Int64Value(int64(*response.FiveGhzSettings.MinBitrate))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
		ID:        types.StringValue(response.ID),
		Name:      types.StringValue(response.Name),
		NetworkID: types.StringValue(response.NetworkID),
		PerSSIDSettings: func() *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettingsRs {
			if response.PerSSIDSettings != nil {
				return &ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettingsRs{
					Status1: func() *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings1Rs {
						if response.PerSSIDSettings.Status1 != nil {
							return &ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings1Rs{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status1.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status1.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status1.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
					Status2: func() *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings2Rs {
						if response.PerSSIDSettings.Status2 != nil {
							return &ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings2Rs{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status2.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status2.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status2.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
					Status3: func() *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings3Rs {
						if response.PerSSIDSettings.Status3 != nil {
							return &ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings3Rs{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status3.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status3.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status3.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
					Status4: func() *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings4Rs {
						if response.PerSSIDSettings.Status4 != nil {
							return &ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings4Rs{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status4.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status4.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status4.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		TwoFourGhzSettings: func() *ResponseApplianceGetNetworkApplianceRfProfileTwoFourGhzSettingsRs {
			if response.TwoFourGhzSettings != nil {
				return &ResponseApplianceGetNetworkApplianceRfProfileTwoFourGhzSettingsRs{
					AxEnabled: func() types.Bool {
						if response.TwoFourGhzSettings.AxEnabled != nil {
							return types.BoolValue(*response.TwoFourGhzSettings.AxEnabled)
						}
						return types.Bool{}
					}(),
					MinBitrate: func() types.Float64 {
						if response.TwoFourGhzSettings.MinBitrate != nil {
							return types.Float64Value(float64(*response.TwoFourGhzSettings.MinBitrate))
						}
						return types.Float64{}
					}(),
				}
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceRfProfilesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceRfProfilesRs)
}
