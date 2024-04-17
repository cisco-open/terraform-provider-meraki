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

// DATA SOURCE NORMAL
import (
	"context"
	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksApplianceRfProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceRfProfilesDataSource{}
)

func NewNetworksApplianceRfProfilesDataSource() datasource.DataSource {
	return &NetworksApplianceRfProfilesDataSource{}
}

type NetworksApplianceRfProfilesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceRfProfilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceRfProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_rf_profiles"
}

func (d *NetworksApplianceRfProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"rf_profile_id": schema.StringAttribute{
				MarkdownDescription: `rfProfileId path parameter. Rf profile ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"five_ghz_settings": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings related to 5Ghz band.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"ax_enabled": schema.BoolAttribute{
								MarkdownDescription: `Whether ax radio on 5Ghz band is on or off.`,
								Computed:            true,
							},
							"min_bitrate": schema.Int64Attribute{
								MarkdownDescription: `Min bitrate (Mbps) of 2.4Ghz band.`,
								Computed:            true,
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
					},
					"network_id": schema.StringAttribute{
						MarkdownDescription: `ID of network this RF Profile belongs in.`,
						Computed:            true,
					},
					"per_ssid_settings": schema.SingleNestedAttribute{
						MarkdownDescription: `Per-SSID radio settings by number.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"status_1": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for SSID 1.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										MarkdownDescription: `Band mode of this SSID`,
										Computed:            true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										MarkdownDescription: `Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.`,
										Computed:            true,
									},
								},
							},
							"status_2": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for SSID 2.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										MarkdownDescription: `Band mode of this SSID`,
										Computed:            true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										MarkdownDescription: `Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.`,
										Computed:            true,
									},
								},
							},
							"status_3": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for SSID 3.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										MarkdownDescription: `Band mode of this SSID`,
										Computed:            true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										MarkdownDescription: `Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.`,
										Computed:            true,
									},
								},
							},
							"status_4": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for SSID 4.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										MarkdownDescription: `Band mode of this SSID`,
										Computed:            true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										MarkdownDescription: `Whether this SSID steers clients to the most open band between 2.4 GHz and 5 GHz.`,
										Computed:            true,
									},
								},
							},
						},
					},
					"two_four_ghz_settings": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings related to 2.4Ghz band.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"ax_enabled": schema.BoolAttribute{
								MarkdownDescription: `Whether ax radio on 2.4Ghz band is on or off.`,
								Computed:            true,
							},
							"min_bitrate": schema.Float64Attribute{
								MarkdownDescription: `Min bitrate (Mbps) of 2.4Ghz band.`,
								Computed:            true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksApplianceRfProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceRfProfiles NetworksApplianceRfProfiles
	diags := req.Config.Get(ctx, &networksApplianceRfProfiles)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksApplianceRfProfiles.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksApplianceRfProfiles.NetworkID.IsNull(), !networksApplianceRfProfiles.RfProfileID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceRfProfiles")
		vvNetworkID := networksApplianceRfProfiles.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceRfProfiles(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceRfProfiles",
				err.Error(),
			)
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceRfProfile")
		vvNetworkID := networksApplianceRfProfiles.NetworkID.ValueString()
		vvRfProfileID := networksApplianceRfProfiles.RfProfileID.ValueString()

		response2, restyResp2, err := d.client.Appliance.GetNetworkApplianceRfProfile(vvNetworkID, vvRfProfileID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceRfProfile",
				err.Error(),
			)
			return
		}

		networksApplianceRfProfiles = ResponseApplianceGetNetworkApplianceRfProfileItemToBody(networksApplianceRfProfiles, response2)
		diags = resp.State.Set(ctx, &networksApplianceRfProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceRfProfiles struct {
	NetworkID   types.String                                   `tfsdk:"network_id"`
	RfProfileID types.String                                   `tfsdk:"rf_profile_id"`
	Item        *ResponseApplianceGetNetworkApplianceRfProfile `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceRfProfile struct {
	FiveGhzSettings    *ResponseApplianceGetNetworkApplianceRfProfileFiveGhzSettings    `tfsdk:"five_ghz_settings"`
	ID                 types.String                                                     `tfsdk:"id"`
	Name               types.String                                                     `tfsdk:"name"`
	NetworkID          types.String                                                     `tfsdk:"network_id"`
	PerSSIDSettings    *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings    `tfsdk:"per_ssid_settings"`
	TwoFourGhzSettings *ResponseApplianceGetNetworkApplianceRfProfileTwoFourGhzSettings `tfsdk:"two_four_ghz_settings"`
}

type ResponseApplianceGetNetworkApplianceRfProfileFiveGhzSettings struct {
	AxEnabled  types.Bool  `tfsdk:"ax_enabled"`
	MinBitrate types.Int64 `tfsdk:"min_bitrate"`
}

type ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings struct {
	Status1 *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings1 `tfsdk:"1"`
	Status2 *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings2 `tfsdk:"2"`
	Status3 *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings3 `tfsdk:"3"`
	Status4 *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings4 `tfsdk:"4"`
}

type ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings1 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
}

type ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings2 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
}

type ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings3 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
}

type ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings4 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
}

type ResponseApplianceGetNetworkApplianceRfProfileTwoFourGhzSettings struct {
	AxEnabled  types.Bool    `tfsdk:"ax_enabled"`
	MinBitrate types.Float64 `tfsdk:"min_bitrate"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceRfProfileItemToBody(state NetworksApplianceRfProfiles, response *merakigosdk.ResponseApplianceGetNetworkApplianceRfProfile) NetworksApplianceRfProfiles {
	itemState := ResponseApplianceGetNetworkApplianceRfProfile{
		FiveGhzSettings: func() *ResponseApplianceGetNetworkApplianceRfProfileFiveGhzSettings {
			if response.FiveGhzSettings != nil {
				return &ResponseApplianceGetNetworkApplianceRfProfileFiveGhzSettings{
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
			return &ResponseApplianceGetNetworkApplianceRfProfileFiveGhzSettings{}
		}(),
		ID:        types.StringValue(response.ID),
		Name:      types.StringValue(response.Name),
		NetworkID: types.StringValue(response.NetworkID),
		PerSSIDSettings: func() *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings {
			if response.PerSSIDSettings != nil {
				return &ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings{
					Status1: func() *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings1 {
						if response.PerSSIDSettings.Status1 != nil {
							return &ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings1{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status1.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status1.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status1.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
							}
						}
						return &ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings1{}
					}(),
					Status2: func() *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings2 {
						if response.PerSSIDSettings.Status2 != nil {
							return &ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings2{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status2.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status2.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status2.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
							}
						}
						return &ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings2{}
					}(),
					Status3: func() *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings3 {
						if response.PerSSIDSettings.Status3 != nil {
							return &ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings3{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status3.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status3.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status3.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
							}
						}
						return &ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings3{}
					}(),
					Status4: func() *ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings4 {
						if response.PerSSIDSettings.Status4 != nil {
							return &ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings4{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status4.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status4.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status4.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
							}
						}
						return &ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings4{}
					}(),
				}
			}
			return &ResponseApplianceGetNetworkApplianceRfProfilePerSsidSettings{}
		}(),
		TwoFourGhzSettings: func() *ResponseApplianceGetNetworkApplianceRfProfileTwoFourGhzSettings {
			if response.TwoFourGhzSettings != nil {
				return &ResponseApplianceGetNetworkApplianceRfProfileTwoFourGhzSettings{
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
			return &ResponseApplianceGetNetworkApplianceRfProfileTwoFourGhzSettings{}
		}(),
	}
	state.Item = &itemState
	return state
}
