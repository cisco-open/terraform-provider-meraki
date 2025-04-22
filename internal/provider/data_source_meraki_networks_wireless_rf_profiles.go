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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksWirelessRfProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessRfProfilesDataSource{}
)

func NewNetworksWirelessRfProfilesDataSource() datasource.DataSource {
	return &NetworksWirelessRfProfilesDataSource{}
}

type NetworksWirelessRfProfilesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessRfProfilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessRfProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_rf_profiles"
}

func (d *NetworksWirelessRfProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"include_template_profiles": schema.BoolAttribute{
				MarkdownDescription: `includeTemplateProfiles query parameter. If the network is bound to a template, this parameter controls whether or not the non-basic RF profiles defined on the template should be included in the response alongside the non-basic profiles defined on the bound network. Defaults to false.`,
				Optional:            true,
			},
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

					"ap_band_settings": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings that will be enabled if selectionType is set to 'ap'.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'. Defaults to dual.`,
								Computed:            true,
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Steers client to most open band. Can be either true or false. Defaults to true.`,
								Computed:            true,
							},
							"bands": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings related to all bands`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"enabled": schema.ListAttribute{
										MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"]`,
										Computed:            true,
										ElementType:         types.StringType,
									},
								},
							},
						},
					},
					"band_selection_type": schema.StringAttribute{
						MarkdownDescription: `Band selection can be set to either 'ssid' or 'ap'. This param is required on creation.`,
						Computed:            true,
					},
					"client_balancing_enabled": schema.BoolAttribute{
						MarkdownDescription: `Steers client to best available access point. Can be either true or false. Defaults to true.`,
						Computed:            true,
					},
					"five_ghz_settings": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings related to 5Ghz band`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"channel_width": schema.StringAttribute{
								MarkdownDescription: `Sets channel width (MHz) for 5Ghz band. Can be one of 'auto', '20', '40' or '80'. Defaults to auto.`,
								Computed:            true,
							},
							"max_power": schema.Int64Attribute{
								MarkdownDescription: `Sets max power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 30.`,
								Computed:            true,
							},
							"min_bitrate": schema.Int64Attribute{
								MarkdownDescription: `Sets min bitrate (Mbps) of 5Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'. Defaults to 12.`,
								Computed:            true,
							},
							"min_power": schema.Int64Attribute{
								MarkdownDescription: `Sets min power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 8.`,
								Computed:            true,
							},
							"rxsop": schema.Int64Attribute{
								MarkdownDescription: `The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.`,
								Computed:            true,
							},
						},
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `The name of the new profile. Must be unique.`,
						Computed:            true,
					},
					"is_indoor_default": schema.BoolAttribute{
						MarkdownDescription: `Set this profile as the default indoor rf profile. If the profile ID is one of 'indoor' or 'outdoor',   then a new profile will be created from the respective ID and set as the default`,
						Computed:            true,
					},
					"is_outdoor_default": schema.BoolAttribute{
						MarkdownDescription: `Set this profile as the default outdoor rf profile. If the profile ID is one of 'indoor' or 'outdoor',   then a new profile will be created from the respective ID and set as the default`,
						Computed:            true,
					},
					"min_bitrate_type": schema.StringAttribute{
						MarkdownDescription: `Minimum bitrate can be set to either 'band' or 'ssid'. Defaults to band.`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the new profile. Must be unique. This param is required on creation.`,
						Computed:            true,
					},
					"network_id": schema.StringAttribute{
						MarkdownDescription: `The network ID of the RF Profile`,
						Computed:            true,
					},
					"per_ssid_settings": schema.SingleNestedAttribute{
						MarkdownDescription: `Per-SSID radio settings by number.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"status_0": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for SSID 0`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
										Computed:            true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
										Computed:            true,
									},
									"bands": schema.SingleNestedAttribute{
										MarkdownDescription: `Settings related to all bands`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"enabled": schema.ListAttribute{
												MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"]`,
												Computed:            true,
												ElementType:         types.StringType,
											},
										},
									},
									"min_bitrate": schema.Int64Attribute{
										MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Name of SSID`,
										Computed:            true,
									},
								},
							},
							"status_1": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for SSID 1`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
										Computed:            true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
										Computed:            true,
									},
									"bands": schema.SingleNestedAttribute{
										MarkdownDescription: `Settings related to all bands`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"enabled": schema.ListAttribute{
												MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"]`,
												Computed:            true,
												ElementType:         types.StringType,
											},
										},
									},
									"min_bitrate": schema.Int64Attribute{
										MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Name of SSID`,
										Computed:            true,
									},
								},
							},
							"status_10": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for SSID 10`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
										Computed:            true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
										Computed:            true,
									},
									"bands": schema.SingleNestedAttribute{
										MarkdownDescription: `Settings related to all bands`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"enabled": schema.ListAttribute{
												MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"]`,
												Computed:            true,
												ElementType:         types.StringType,
											},
										},
									},
									"min_bitrate": schema.Int64Attribute{
										MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Name of SSID`,
										Computed:            true,
									},
								},
							},
							"status_11": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for SSID 11`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
										Computed:            true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
										Computed:            true,
									},
									"bands": schema.SingleNestedAttribute{
										MarkdownDescription: `Settings related to all bands`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"enabled": schema.ListAttribute{
												MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"]`,
												Computed:            true,
												ElementType:         types.StringType,
											},
										},
									},
									"min_bitrate": schema.Int64Attribute{
										MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Name of SSID`,
										Computed:            true,
									},
								},
							},
							"status_12": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for SSID 12`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
										Computed:            true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
										Computed:            true,
									},
									"bands": schema.SingleNestedAttribute{
										MarkdownDescription: `Settings related to all bands`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"enabled": schema.ListAttribute{
												MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"]`,
												Computed:            true,
												ElementType:         types.StringType,
											},
										},
									},
									"min_bitrate": schema.Int64Attribute{
										MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Name of SSID`,
										Computed:            true,
									},
								},
							},
							"status_13": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for SSID 13`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
										Computed:            true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
										Computed:            true,
									},
									"bands": schema.SingleNestedAttribute{
										MarkdownDescription: `Settings related to all bands`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"enabled": schema.ListAttribute{
												MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"]`,
												Computed:            true,
												ElementType:         types.StringType,
											},
										},
									},
									"min_bitrate": schema.Int64Attribute{
										MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Name of SSID`,
										Computed:            true,
									},
								},
							},
							"status_14": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for SSID 14`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
										Computed:            true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
										Computed:            true,
									},
									"bands": schema.SingleNestedAttribute{
										MarkdownDescription: `Settings related to all bands`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"enabled": schema.ListAttribute{
												MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"]`,
												Computed:            true,
												ElementType:         types.StringType,
											},
										},
									},
									"min_bitrate": schema.Int64Attribute{
										MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Name of SSID`,
										Computed:            true,
									},
								},
							},
							"status_2": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for SSID 2`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
										Computed:            true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
										Computed:            true,
									},
									"bands": schema.SingleNestedAttribute{
										MarkdownDescription: `Settings related to all bands`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"enabled": schema.ListAttribute{
												MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"]`,
												Computed:            true,
												ElementType:         types.StringType,
											},
										},
									},
									"min_bitrate": schema.Int64Attribute{
										MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Name of SSID`,
										Computed:            true,
									},
								},
							},
							"status_3": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for SSID 3`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
										Computed:            true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
										Computed:            true,
									},
									"bands": schema.SingleNestedAttribute{
										MarkdownDescription: `Settings related to all bands`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"enabled": schema.ListAttribute{
												MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"]`,
												Computed:            true,
												ElementType:         types.StringType,
											},
										},
									},
									"min_bitrate": schema.Int64Attribute{
										MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Name of SSID`,
										Computed:            true,
									},
								},
							},
							"status_4": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for SSID 4`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
										Computed:            true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
										Computed:            true,
									},
									"bands": schema.SingleNestedAttribute{
										MarkdownDescription: `Settings related to all bands`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"enabled": schema.ListAttribute{
												MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"]`,
												Computed:            true,
												ElementType:         types.StringType,
											},
										},
									},
									"min_bitrate": schema.Int64Attribute{
										MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Name of SSID`,
										Computed:            true,
									},
								},
							},
							"status_5": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for SSID 5`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
										Computed:            true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
										Computed:            true,
									},
									"bands": schema.SingleNestedAttribute{
										MarkdownDescription: `Settings related to all bands`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"enabled": schema.ListAttribute{
												MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"]`,
												Computed:            true,
												ElementType:         types.StringType,
											},
										},
									},
									"min_bitrate": schema.Int64Attribute{
										MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Name of SSID`,
										Computed:            true,
									},
								},
							},
							"status_6": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for SSID 6`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
										Computed:            true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
										Computed:            true,
									},
									"bands": schema.SingleNestedAttribute{
										MarkdownDescription: `Settings related to all bands`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"enabled": schema.ListAttribute{
												MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"]`,
												Computed:            true,
												ElementType:         types.StringType,
											},
										},
									},
									"min_bitrate": schema.Int64Attribute{
										MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Name of SSID`,
										Computed:            true,
									},
								},
							},
							"status_7": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for SSID 7`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
										Computed:            true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
										Computed:            true,
									},
									"bands": schema.SingleNestedAttribute{
										MarkdownDescription: `Settings related to all bands`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"enabled": schema.ListAttribute{
												MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"]`,
												Computed:            true,
												ElementType:         types.StringType,
											},
										},
									},
									"min_bitrate": schema.Int64Attribute{
										MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Name of SSID`,
										Computed:            true,
									},
								},
							},
							"status_8": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for SSID 8`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
										Computed:            true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
										Computed:            true,
									},
									"bands": schema.SingleNestedAttribute{
										MarkdownDescription: `Settings related to all bands`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"enabled": schema.ListAttribute{
												MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"]`,
												Computed:            true,
												ElementType:         types.StringType,
											},
										},
									},
									"min_bitrate": schema.Int64Attribute{
										MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Name of SSID`,
										Computed:            true,
									},
								},
							},
							"status_9": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings for SSID 9`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
										Computed:            true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
										Computed:            true,
									},
									"bands": schema.SingleNestedAttribute{
										MarkdownDescription: `Settings related to all bands`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"enabled": schema.ListAttribute{
												MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"]`,
												Computed:            true,
												ElementType:         types.StringType,
											},
										},
									},
									"min_bitrate": schema.Int64Attribute{
										MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Name of SSID`,
										Computed:            true,
									},
								},
							},
						},
					},
					"six_ghz_settings": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings related to 6Ghz band. Only applicable to networks with 6Ghz capable APs`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"channel_width": schema.StringAttribute{
								MarkdownDescription: `Sets channel width (MHz) for 6Ghz band. Can be one of '0', '20', '40', '80' or '160'. Defaults to auto.`,
								Computed:            true,
							},
							"max_power": schema.Int64Attribute{
								MarkdownDescription: `Sets max power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 30.`,
								Computed:            true,
							},
							"min_bitrate": schema.Int64Attribute{
								MarkdownDescription: `Sets min bitrate (Mbps) of 6Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'. Defaults to 12.`,
								Computed:            true,
							},
							"min_power": schema.Int64Attribute{
								MarkdownDescription: `Sets min power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 8.`,
								Computed:            true,
							},
							"rxsop": schema.Int64Attribute{
								MarkdownDescription: `The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.`,
								Computed:            true,
							},
							"valid_auto_channels": schema.ListAttribute{
								MarkdownDescription: `Sets valid auto channels for 6Ghz band. Can be one of '1', '5', '9', '13', '17', '21', '25', '29', '33', '37', '41', '45', '49', '53', '57', '61', '65', '69', '73', '77', '81', '85', '89', '93', '97', '101', '105', '109', '113', '117', '121', '125', '129', '133', '137', '141', '145', '149', '153', '157', '161', '165', '169', '173', '177', '181', '185', '189', '193', '197', '201', '205', '209', '213', '217', '221', '225', '229' or '233'. Defaults to auto.`,
								Computed:            true,
								ElementType:         types.StringType, //TODO FINAL ELSE param_schema.Elem.Type para revisar
								// {'Type': 'schema.TypeInt'}
							},
						},
					},
					"transmission": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings related to radio transmission.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Toggle for radio transmission. When false, radios will not transmit at all.`,
								Computed:            true,
							},
						},
					},
					"two_four_ghz_settings": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings related to 2.4Ghz band`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"ax_enabled": schema.BoolAttribute{
								MarkdownDescription: `Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.`,
								Computed:            true,
							},
							"max_power": schema.Int64Attribute{
								MarkdownDescription: `Sets max power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 30.`,
								Computed:            true,
							},
							"min_bitrate": schema.Float64Attribute{
								MarkdownDescription: `Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'. Defaults to 11.`,
								Computed:            true,
							},
							"min_power": schema.Int64Attribute{
								MarkdownDescription: `Sets min power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 5.`,
								Computed:            true,
							},
							"rxsop": schema.Int64Attribute{
								MarkdownDescription: `The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.`,
								Computed:            true,
							},
							"valid_auto_channels": schema.ListAttribute{
								MarkdownDescription: `Sets valid auto channels for 2.4Ghz band. Can be one of '1', '6' or '11'. Defaults to [1, 6, 11].`,
								Computed:            true,
								ElementType:         types.StringType, //TODO FINAL ELSE param_schema.Elem.Type para revisar

							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessRfProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessRfProfiles NetworksWirelessRfProfiles
	diags := req.Config.Get(ctx, &networksWirelessRfProfiles)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksWirelessRfProfiles.NetworkID.IsNull(), !networksWirelessRfProfiles.IncludeTemplateProfiles.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksWirelessRfProfiles.NetworkID.IsNull(), !networksWirelessRfProfiles.RfProfileID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessRfProfiles")
		vvNetworkID := networksWirelessRfProfiles.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkWirelessRfProfilesQueryParams{}

		queryParams1.IncludeTemplateProfiles = networksWirelessRfProfiles.IncludeTemplateProfiles.ValueBool()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessRfProfiles(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessRfProfiles",
				err.Error(),
			)
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessRfProfile")
		vvNetworkID := networksWirelessRfProfiles.NetworkID.ValueString()
		vvRfProfileID := networksWirelessRfProfiles.RfProfileID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Wireless.GetNetworkWirelessRfProfile(vvNetworkID, vvRfProfileID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessRfProfile",
				err.Error(),
			)
			return
		}

		networksWirelessRfProfiles = ResponseWirelessGetNetworkWirelessRfProfileItemToBody(networksWirelessRfProfiles, response2)
		diags = resp.State.Set(ctx, &networksWirelessRfProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessRfProfiles struct {
	NetworkID               types.String                                 `tfsdk:"network_id"`
	IncludeTemplateProfiles types.Bool                                   `tfsdk:"include_template_profiles"`
	RfProfileID             types.String                                 `tfsdk:"rf_profile_id"`
	Item                    *ResponseWirelessGetNetworkWirelessRfProfile `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessRfProfile struct {
	ApBandSettings         *ResponseWirelessGetNetworkWirelessRfProfileApBandSettings     `tfsdk:"ap_band_settings"`
	BandSelectionType      types.String                                                   `tfsdk:"band_selection_type"`
	ClientBalancingEnabled types.Bool                                                     `tfsdk:"client_balancing_enabled"`
	FiveGhzSettings        *ResponseWirelessGetNetworkWirelessRfProfileFiveGhzSettings    `tfsdk:"five_ghz_settings"`
	ID                     types.String                                                   `tfsdk:"id"`
	IsIndoorDefault        types.Bool                                                     `tfsdk:"is_indoor_default"`
	IsOutdoorDefault       types.Bool                                                     `tfsdk:"is_outdoor_default"`
	MinBitrateType         types.String                                                   `tfsdk:"min_bitrate_type"`
	Name                   types.String                                                   `tfsdk:"name"`
	NetworkID              types.String                                                   `tfsdk:"network_id"`
	PerSSIDSettings        *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings    `tfsdk:"per_ssid_settings"`
	SixGhzSettings         *ResponseWirelessGetNetworkWirelessRfProfileSixGhzSettings     `tfsdk:"six_ghz_settings"`
	Transmission           *ResponseWirelessGetNetworkWirelessRfProfileTransmission       `tfsdk:"transmission"`
	TwoFourGhzSettings     *ResponseWirelessGetNetworkWirelessRfProfileTwoFourGhzSettings `tfsdk:"two_four_ghz_settings"`
}

type ResponseWirelessGetNetworkWirelessRfProfileApBandSettings struct {
	BandOperationMode   types.String                                                    `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                      `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfileApBandSettingsBands `tfsdk:"bands"`
}

type ResponseWirelessGetNetworkWirelessRfProfileApBandSettingsBands struct {
	Enabled types.List `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfileFiveGhzSettings struct {
	ChannelWidth      types.String `tfsdk:"channel_width"`
	MaxPower          types.Int64  `tfsdk:"max_power"`
	MinBitrate        types.Int64  `tfsdk:"min_bitrate"`
	MinPower          types.Int64  `tfsdk:"min_power"`
	Rxsop             types.Int64  `tfsdk:"rxsop"`
	ValidAutoChannels types.List   `tfsdk:"valid_auto_channels"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings struct {
	Status0  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings0  `tfsdk:"status_0"`
	Status1  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1  `tfsdk:"status_1"`
	Status10 *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10 `tfsdk:"status_10"`
	Status11 *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11 `tfsdk:"status_11"`
	Status12 *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12 `tfsdk:"status_12"`
	Status13 *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13 `tfsdk:"status_13"`
	Status14 *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14 `tfsdk:"status_14"`
	Status2  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2  `tfsdk:"status_2"`
	Status3  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3  `tfsdk:"status_3"`
	Status4  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4  `tfsdk:"status_4"`
	Status5  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5  `tfsdk:"status_5"`
	Status6  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6  `tfsdk:"status_6"`
	Status7  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7  `tfsdk:"status_7"`
	Status8  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8  `tfsdk:"status_8"`
	Status9  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9  `tfsdk:"status_9"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings0 struct {
	BandOperationMode   types.String                                                      `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                        `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings0Bands `tfsdk:"bands"`
	MinBitrate          types.Int64                                                       `tfsdk:"min_bitrate"`
	Name                types.String                                                      `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings0Bands struct {
	Enabled types.List `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1 struct {
	BandOperationMode   types.String                                                      `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                        `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1Bands `tfsdk:"bands"`
	MinBitrate          types.Int64                                                       `tfsdk:"min_bitrate"`
	Name                types.String                                                      `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1Bands struct {
	Enabled types.List `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10 struct {
	BandOperationMode   types.String                                                       `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                         `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10Bands `tfsdk:"bands"`
	MinBitrate          types.Int64                                                        `tfsdk:"min_bitrate"`
	Name                types.String                                                       `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10Bands struct {
	Enabled types.List `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11 struct {
	BandOperationMode   types.String                                                       `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                         `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11Bands `tfsdk:"bands"`
	MinBitrate          types.Int64                                                        `tfsdk:"min_bitrate"`
	Name                types.String                                                       `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11Bands struct {
	Enabled types.List `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12 struct {
	BandOperationMode   types.String                                                       `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                         `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12Bands `tfsdk:"bands"`
	MinBitrate          types.Int64                                                        `tfsdk:"min_bitrate"`
	Name                types.String                                                       `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12Bands struct {
	Enabled types.List `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13 struct {
	BandOperationMode   types.String                                                       `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                         `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13Bands `tfsdk:"bands"`
	MinBitrate          types.Int64                                                        `tfsdk:"min_bitrate"`
	Name                types.String                                                       `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13Bands struct {
	Enabled types.List `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14 struct {
	BandOperationMode   types.String                                                       `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                         `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14Bands `tfsdk:"bands"`
	MinBitrate          types.Int64                                                        `tfsdk:"min_bitrate"`
	Name                types.String                                                       `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14Bands struct {
	Enabled types.List `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2 struct {
	BandOperationMode   types.String                                                      `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                        `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2Bands `tfsdk:"bands"`
	MinBitrate          types.Int64                                                       `tfsdk:"min_bitrate"`
	Name                types.String                                                      `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2Bands struct {
	Enabled types.List `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3 struct {
	BandOperationMode   types.String                                                      `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                        `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3Bands `tfsdk:"bands"`
	MinBitrate          types.Int64                                                       `tfsdk:"min_bitrate"`
	Name                types.String                                                      `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3Bands struct {
	Enabled types.List `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4 struct {
	BandOperationMode   types.String                                                      `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                        `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4Bands `tfsdk:"bands"`
	MinBitrate          types.Int64                                                       `tfsdk:"min_bitrate"`
	Name                types.String                                                      `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4Bands struct {
	Enabled types.List `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5 struct {
	BandOperationMode   types.String                                                      `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                        `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5Bands `tfsdk:"bands"`
	MinBitrate          types.Int64                                                       `tfsdk:"min_bitrate"`
	Name                types.String                                                      `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5Bands struct {
	Enabled types.List `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6 struct {
	BandOperationMode   types.String                                                      `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                        `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6Bands `tfsdk:"bands"`
	MinBitrate          types.Int64                                                       `tfsdk:"min_bitrate"`
	Name                types.String                                                      `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6Bands struct {
	Enabled types.List `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7 struct {
	BandOperationMode   types.String                                                      `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                        `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7Bands `tfsdk:"bands"`
	MinBitrate          types.Int64                                                       `tfsdk:"min_bitrate"`
	Name                types.String                                                      `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7Bands struct {
	Enabled types.List `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8 struct {
	BandOperationMode   types.String                                                      `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                        `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8Bands `tfsdk:"bands"`
	MinBitrate          types.Int64                                                       `tfsdk:"min_bitrate"`
	Name                types.String                                                      `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8Bands struct {
	Enabled types.List `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9 struct {
	BandOperationMode   types.String                                                      `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                        `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9Bands `tfsdk:"bands"`
	MinBitrate          types.Int64                                                       `tfsdk:"min_bitrate"`
	Name                types.String                                                      `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9Bands struct {
	Enabled types.List `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfileSixGhzSettings struct {
	ChannelWidth      types.String `tfsdk:"channel_width"`
	MaxPower          types.Int64  `tfsdk:"max_power"`
	MinBitrate        types.Int64  `tfsdk:"min_bitrate"`
	MinPower          types.Int64  `tfsdk:"min_power"`
	Rxsop             types.Int64  `tfsdk:"rxsop"`
	ValidAutoChannels types.List   `tfsdk:"valid_auto_channels"`
}

type ResponseWirelessGetNetworkWirelessRfProfileTransmission struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfileTwoFourGhzSettings struct {
	AxEnabled         types.Bool    `tfsdk:"ax_enabled"`
	MaxPower          types.Int64   `tfsdk:"max_power"`
	MinBitrate        types.Float64 `tfsdk:"min_bitrate"`
	MinPower          types.Int64   `tfsdk:"min_power"`
	Rxsop             types.Int64   `tfsdk:"rxsop"`
	ValidAutoChannels types.List    `tfsdk:"valid_auto_channels"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessRfProfileItemToBody(state NetworksWirelessRfProfiles, response *merakigosdk.ResponseWirelessGetNetworkWirelessRfProfile) NetworksWirelessRfProfiles {
	itemState := ResponseWirelessGetNetworkWirelessRfProfile{
		ApBandSettings: func() *ResponseWirelessGetNetworkWirelessRfProfileApBandSettings {
			if response.ApBandSettings != nil {
				return &ResponseWirelessGetNetworkWirelessRfProfileApBandSettings{
					BandOperationMode: types.StringValue(response.ApBandSettings.BandOperationMode),
					BandSteeringEnabled: func() types.Bool {
						if response.ApBandSettings.BandSteeringEnabled != nil {
							return types.BoolValue(*response.ApBandSettings.BandSteeringEnabled)
						}
						return types.Bool{}
					}(),
					Bands: func() *ResponseWirelessGetNetworkWirelessRfProfileApBandSettingsBands {
						if response.ApBandSettings.Bands != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfileApBandSettingsBands{
								Enabled: StringSliceToList(response.ApBandSettings.Bands.Enabled),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		BandSelectionType: types.StringValue(response.BandSelectionType),
		ClientBalancingEnabled: func() types.Bool {
			if response.ClientBalancingEnabled != nil {
				return types.BoolValue(*response.ClientBalancingEnabled)
			}
			return types.Bool{}
		}(),
		FiveGhzSettings: func() *ResponseWirelessGetNetworkWirelessRfProfileFiveGhzSettings {
			if response.FiveGhzSettings != nil {
				return &ResponseWirelessGetNetworkWirelessRfProfileFiveGhzSettings{
					ChannelWidth: types.StringValue(response.FiveGhzSettings.ChannelWidth),
					MaxPower: func() types.Int64 {
						if response.FiveGhzSettings.MaxPower != nil {
							return types.Int64Value(int64(*response.FiveGhzSettings.MaxPower))
						}
						return types.Int64{}
					}(),
					MinBitrate: func() types.Int64 {
						if response.FiveGhzSettings.MinBitrate != nil {
							return types.Int64Value(int64(*response.FiveGhzSettings.MinBitrate))
						}
						return types.Int64{}
					}(),
					MinPower: func() types.Int64 {
						if response.FiveGhzSettings.MinPower != nil {
							return types.Int64Value(int64(*response.FiveGhzSettings.MinPower))
						}
						return types.Int64{}
					}(),
					Rxsop: func() types.Int64 {
						if response.FiveGhzSettings.Rxsop != nil {
							return types.Int64Value(int64(*response.FiveGhzSettings.Rxsop))
						}
						return types.Int64{}
					}(),
					ValidAutoChannels: StringSliceToListInt(response.FiveGhzSettings.ValidAutoChannels),
				}
			}
			return nil
		}(),
		ID: types.StringValue(response.ID),
		IsIndoorDefault: func() types.Bool {
			if response.IsIndoorDefault != nil {
				return types.BoolValue(*response.IsIndoorDefault)
			}
			return types.Bool{}
		}(),
		IsOutdoorDefault: func() types.Bool {
			if response.IsOutdoorDefault != nil {
				return types.BoolValue(*response.IsOutdoorDefault)
			}
			return types.Bool{}
		}(),
		MinBitrateType: types.StringValue(response.MinBitrateType),
		Name:           types.StringValue(response.Name),
		NetworkID:      types.StringValue(response.NetworkID),
		PerSSIDSettings: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings {
			if response.PerSSIDSettings != nil {
				return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings{
					Status0: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings0 {
						if response.PerSSIDSettings.Status0 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings0{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status0.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status0.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status0.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings0Bands {
									if response.PerSSIDSettings.Status0.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings0Bands{
											Enabled: StringSliceToList(response.PerSSIDSettings.Status0.Bands.Enabled),
										}
									}
									return nil
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status0.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status0.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status0.Name),
							}
						}
						return nil
					}(),
					Status1: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1 {
						if response.PerSSIDSettings.Status1 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status1.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status1.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status1.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1Bands {
									if response.PerSSIDSettings.Status1.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1Bands{
											Enabled: StringSliceToList(response.PerSSIDSettings.Status1.Bands.Enabled),
										}
									}
									return nil
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status1.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status1.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status1.Name),
							}
						}
						return nil
					}(),
					Status10: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10 {
						if response.PerSSIDSettings.Status10 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status10.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status10.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status10.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10Bands {
									if response.PerSSIDSettings.Status10.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10Bands{
											Enabled: StringSliceToList(response.PerSSIDSettings.Status10.Bands.Enabled),
										}
									}
									return nil
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status10.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status10.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status10.Name),
							}
						}
						return nil
					}(),
					Status11: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11 {
						if response.PerSSIDSettings.Status11 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status11.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status11.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status11.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11Bands {
									if response.PerSSIDSettings.Status11.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11Bands{
											Enabled: StringSliceToList(response.PerSSIDSettings.Status11.Bands.Enabled),
										}
									}
									return nil
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status11.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status11.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status11.Name),
							}
						}
						return nil
					}(),
					Status12: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12 {
						if response.PerSSIDSettings.Status12 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status12.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status12.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status12.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12Bands {
									if response.PerSSIDSettings.Status12.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12Bands{
											Enabled: StringSliceToList(response.PerSSIDSettings.Status12.Bands.Enabled),
										}
									}
									return nil
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status12.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status12.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status12.Name),
							}
						}
						return nil
					}(),
					Status13: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13 {
						if response.PerSSIDSettings.Status13 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status13.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status13.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status13.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13Bands {
									if response.PerSSIDSettings.Status13.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13Bands{
											Enabled: StringSliceToList(response.PerSSIDSettings.Status13.Bands.Enabled),
										}
									}
									return nil
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status13.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status13.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status13.Name),
							}
						}
						return nil
					}(),
					Status14: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14 {
						if response.PerSSIDSettings.Status14 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status14.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status14.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status14.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14Bands {
									if response.PerSSIDSettings.Status14.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14Bands{
											Enabled: StringSliceToList(response.PerSSIDSettings.Status14.Bands.Enabled),
										}
									}
									return nil
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status14.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status14.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status14.Name),
							}
						}
						return nil
					}(),
					Status2: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2 {
						if response.PerSSIDSettings.Status2 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status2.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status2.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status2.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2Bands {
									if response.PerSSIDSettings.Status2.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2Bands{
											Enabled: StringSliceToList(response.PerSSIDSettings.Status2.Bands.Enabled),
										}
									}
									return nil
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status2.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status2.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status2.Name),
							}
						}
						return nil
					}(),
					Status3: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3 {
						if response.PerSSIDSettings.Status3 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status3.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status3.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status3.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3Bands {
									if response.PerSSIDSettings.Status3.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3Bands{
											Enabled: StringSliceToList(response.PerSSIDSettings.Status3.Bands.Enabled),
										}
									}
									return nil
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status3.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status3.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status3.Name),
							}
						}
						return nil
					}(),
					Status4: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4 {
						if response.PerSSIDSettings.Status4 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status4.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status4.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status4.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4Bands {
									if response.PerSSIDSettings.Status4.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4Bands{
											Enabled: StringSliceToList(response.PerSSIDSettings.Status4.Bands.Enabled),
										}
									}
									return nil
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status4.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status4.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status4.Name),
							}
						}
						return nil
					}(),
					Status5: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5 {
						if response.PerSSIDSettings.Status5 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status5.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status5.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status5.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5Bands {
									if response.PerSSIDSettings.Status5.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5Bands{
											Enabled: StringSliceToList(response.PerSSIDSettings.Status5.Bands.Enabled),
										}
									}
									return nil
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status5.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status5.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status5.Name),
							}
						}
						return nil
					}(),
					Status6: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6 {
						if response.PerSSIDSettings.Status6 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status6.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status6.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status6.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6Bands {
									if response.PerSSIDSettings.Status6.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6Bands{
											Enabled: StringSliceToList(response.PerSSIDSettings.Status6.Bands.Enabled),
										}
									}
									return nil
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status6.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status6.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status6.Name),
							}
						}
						return nil
					}(),
					Status7: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7 {
						if response.PerSSIDSettings.Status7 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status7.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status7.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status7.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7Bands {
									if response.PerSSIDSettings.Status7.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7Bands{
											Enabled: StringSliceToList(response.PerSSIDSettings.Status7.Bands.Enabled),
										}
									}
									return nil
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status7.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status7.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status7.Name),
							}
						}
						return nil
					}(),
					Status8: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8 {
						if response.PerSSIDSettings.Status8 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status8.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status8.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status8.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8Bands {
									if response.PerSSIDSettings.Status8.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8Bands{
											Enabled: StringSliceToList(response.PerSSIDSettings.Status8.Bands.Enabled),
										}
									}
									return nil
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status8.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status8.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status8.Name),
							}
						}
						return nil
					}(),
					Status9: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9 {
						if response.PerSSIDSettings.Status9 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status9.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status9.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status9.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9Bands {
									if response.PerSSIDSettings.Status9.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9Bands{
											Enabled: StringSliceToList(response.PerSSIDSettings.Status9.Bands.Enabled),
										}
									}
									return nil
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status9.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status9.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status9.Name),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		SixGhzSettings: func() *ResponseWirelessGetNetworkWirelessRfProfileSixGhzSettings {
			if response.SixGhzSettings != nil {
				return &ResponseWirelessGetNetworkWirelessRfProfileSixGhzSettings{
					ChannelWidth: types.StringValue(response.SixGhzSettings.ChannelWidth),
					MaxPower: func() types.Int64 {
						if response.SixGhzSettings.MaxPower != nil {
							return types.Int64Value(int64(*response.SixGhzSettings.MaxPower))
						}
						return types.Int64{}
					}(),
					MinBitrate: func() types.Int64 {
						if response.SixGhzSettings.MinBitrate != nil {
							return types.Int64Value(int64(*response.SixGhzSettings.MinBitrate))
						}
						return types.Int64{}
					}(),
					MinPower: func() types.Int64 {
						if response.SixGhzSettings.MinPower != nil {
							return types.Int64Value(int64(*response.SixGhzSettings.MinPower))
						}
						return types.Int64{}
					}(),
					Rxsop: func() types.Int64 {
						if response.SixGhzSettings.Rxsop != nil {
							return types.Int64Value(int64(*response.SixGhzSettings.Rxsop))
						}
						return types.Int64{}
					}(),
					ValidAutoChannels: StringSliceToListInt(response.SixGhzSettings.ValidAutoChannels),
				}
			}
			return nil
		}(),
		Transmission: func() *ResponseWirelessGetNetworkWirelessRfProfileTransmission {
			if response.Transmission != nil {
				return &ResponseWirelessGetNetworkWirelessRfProfileTransmission{
					Enabled: func() types.Bool {
						if response.Transmission.Enabled != nil {
							return types.BoolValue(*response.Transmission.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
		TwoFourGhzSettings: func() *ResponseWirelessGetNetworkWirelessRfProfileTwoFourGhzSettings {
			if response.TwoFourGhzSettings != nil {
				return &ResponseWirelessGetNetworkWirelessRfProfileTwoFourGhzSettings{
					AxEnabled: func() types.Bool {
						if response.TwoFourGhzSettings.AxEnabled != nil {
							return types.BoolValue(*response.TwoFourGhzSettings.AxEnabled)
						}
						return types.Bool{}
					}(),
					MaxPower: func() types.Int64 {
						if response.TwoFourGhzSettings.MaxPower != nil {
							return types.Int64Value(int64(*response.TwoFourGhzSettings.MaxPower))
						}
						return types.Int64{}
					}(),
					MinBitrate: func() types.Float64 {
						if response.TwoFourGhzSettings.MinBitrate != nil {
							return types.Float64Value(float64(*response.TwoFourGhzSettings.MinBitrate))
						}
						return types.Float64{}
					}(),
					MinPower: func() types.Int64 {
						if response.TwoFourGhzSettings.MinPower != nil {
							return types.Int64Value(int64(*response.TwoFourGhzSettings.MinPower))
						}
						return types.Int64{}
					}(),
					Rxsop: func() types.Int64 {
						if response.TwoFourGhzSettings.Rxsop != nil {
							return types.Int64Value(int64(*response.TwoFourGhzSettings.Rxsop))
						}
						return types.Int64{}
					}(),
					ValidAutoChannels: StringSliceToListInt(response.TwoFourGhzSettings.ValidAutoChannels),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
