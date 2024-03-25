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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

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

					"afc_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"ap_band_settings": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								Computed: true,
							},
							"band_steering_enabled": schema.BoolAttribute{
								Computed: true,
							},
						},
					},
					"band_selection_type": schema.StringAttribute{
						Computed: true,
					},
					"client_balancing_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"five_ghz_settings": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"channel_width": schema.StringAttribute{
								Computed: true,
							},
							"max_power": schema.Int64Attribute{
								Computed: true,
							},
							"min_bitrate": schema.Int64Attribute{
								Computed: true,
							},
							"min_power": schema.Int64Attribute{
								Computed: true,
							},
							"valid_auto_channels": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
					"id": schema.StringAttribute{
						Computed: true,
					},
					"min_bitrate_type": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"network_id": schema.StringAttribute{
						Computed: true,
					},
					"per_ssid_settings": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"status_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										Computed: true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"min_bitrate": schema.Int64Attribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"status_1": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										Computed: true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"min_bitrate": schema.Int64Attribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"status_10": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										Computed: true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"min_bitrate": schema.Int64Attribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"status_11": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										Computed: true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"min_bitrate": schema.Int64Attribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"status_12": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										Computed: true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"min_bitrate": schema.Int64Attribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"status_13": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										Computed: true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"min_bitrate": schema.Int64Attribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"status_14": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										Computed: true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"min_bitrate": schema.Int64Attribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"status_2": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										Computed: true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"min_bitrate": schema.Int64Attribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"status_3": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										Computed: true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"min_bitrate": schema.Int64Attribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"status_4": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										Computed: true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"min_bitrate": schema.Int64Attribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"status_5": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										Computed: true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"min_bitrate": schema.Int64Attribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"status_6": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										Computed: true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"min_bitrate": schema.Int64Attribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"status_7": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										Computed: true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"min_bitrate": schema.Int64Attribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"status_8": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										Computed: true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"min_bitrate": schema.Int64Attribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"status_9": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"band_operation_mode": schema.StringAttribute{
										Computed: true,
									},
									"band_steering_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"min_bitrate": schema.Int64Attribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
					},
					"six_ghz_settings": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"afc_enabled": schema.BoolAttribute{
								Computed: true,
							},
							"channel_width": schema.StringAttribute{
								Computed: true,
							},
							"max_power": schema.Int64Attribute{
								Computed: true,
							},
							"min_bitrate": schema.Int64Attribute{
								Computed: true,
							},
							"min_power": schema.Int64Attribute{
								Computed: true,
							},
							"valid_auto_channels": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
					"transmission": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								Computed: true,
							},
						},
					},
					"two_four_ghz_settings": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"ax_enabled": schema.BoolAttribute{
								Computed: true,
							},
							"max_power": schema.Int64Attribute{
								Computed: true,
							},
							"min_bitrate": schema.Int64Attribute{
								Computed: true,
							},
							"min_power": schema.Int64Attribute{
								Computed: true,
							},
							"valid_auto_channels": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseWirelessGetNetworkWirelessRfProfiles`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"afc_enabled": schema.BoolAttribute{
							Computed: true,
						},
						"ap_band_settings": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"band_operation_mode": schema.StringAttribute{
									Computed: true,
								},
								"band_steering_enabled": schema.BoolAttribute{
									Computed: true,
								},
							},
						},
						"band_selection_type": schema.StringAttribute{
							Computed: true,
						},
						"client_balancing_enabled": schema.BoolAttribute{
							Computed: true,
						},
						"five_ghz_settings": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"channel_width": schema.StringAttribute{
									Computed: true,
								},
								"max_power": schema.Int64Attribute{
									Computed: true,
								},
								"min_bitrate": schema.Int64Attribute{
									Computed: true,
								},
								"min_power": schema.Int64Attribute{
									Computed: true,
								},
								"valid_auto_channels": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
							},
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"min_bitrate_type": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"network_id": schema.StringAttribute{
							Computed: true,
						},
						"per_ssid_settings": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"status_0": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"band_operation_mode": schema.StringAttribute{
											Computed: true,
										},
										"band_steering_enabled": schema.BoolAttribute{
											Computed: true,
										},
										"min_bitrate": schema.Int64Attribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"status_1": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"band_operation_mode": schema.StringAttribute{
											Computed: true,
										},
										"band_steering_enabled": schema.BoolAttribute{
											Computed: true,
										},
										"min_bitrate": schema.Int64Attribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"status_10": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"band_operation_mode": schema.StringAttribute{
											Computed: true,
										},
										"band_steering_enabled": schema.BoolAttribute{
											Computed: true,
										},
										"min_bitrate": schema.Int64Attribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"status_11": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"band_operation_mode": schema.StringAttribute{
											Computed: true,
										},
										"band_steering_enabled": schema.BoolAttribute{
											Computed: true,
										},
										"min_bitrate": schema.Int64Attribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"status_12": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"band_operation_mode": schema.StringAttribute{
											Computed: true,
										},
										"band_steering_enabled": schema.BoolAttribute{
											Computed: true,
										},
										"min_bitrate": schema.Int64Attribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"status_13": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"band_operation_mode": schema.StringAttribute{
											Computed: true,
										},
										"band_steering_enabled": schema.BoolAttribute{
											Computed: true,
										},
										"min_bitrate": schema.Int64Attribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"status_14": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"band_operation_mode": schema.StringAttribute{
											Computed: true,
										},
										"band_steering_enabled": schema.BoolAttribute{
											Computed: true,
										},
										"min_bitrate": schema.Int64Attribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"status_2": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"band_operation_mode": schema.StringAttribute{
											Computed: true,
										},
										"band_steering_enabled": schema.BoolAttribute{
											Computed: true,
										},
										"min_bitrate": schema.Int64Attribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"status_3": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"band_operation_mode": schema.StringAttribute{
											Computed: true,
										},
										"band_steering_enabled": schema.BoolAttribute{
											Computed: true,
										},
										"min_bitrate": schema.Int64Attribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"status_4": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"band_operation_mode": schema.StringAttribute{
											Computed: true,
										},
										"band_steering_enabled": schema.BoolAttribute{
											Computed: true,
										},
										"min_bitrate": schema.Int64Attribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"status_5": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"band_operation_mode": schema.StringAttribute{
											Computed: true,
										},
										"band_steering_enabled": schema.BoolAttribute{
											Computed: true,
										},
										"min_bitrate": schema.Int64Attribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"status_6": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"band_operation_mode": schema.StringAttribute{
											Computed: true,
										},
										"band_steering_enabled": schema.BoolAttribute{
											Computed: true,
										},
										"min_bitrate": schema.Int64Attribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"status_7": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"band_operation_mode": schema.StringAttribute{
											Computed: true,
										},
										"band_steering_enabled": schema.BoolAttribute{
											Computed: true,
										},
										"min_bitrate": schema.Int64Attribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"status_8": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"band_operation_mode": schema.StringAttribute{
											Computed: true,
										},
										"band_steering_enabled": schema.BoolAttribute{
											Computed: true,
										},
										"min_bitrate": schema.Int64Attribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"status_9": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"band_operation_mode": schema.StringAttribute{
											Computed: true,
										},
										"band_steering_enabled": schema.BoolAttribute{
											Computed: true,
										},
										"min_bitrate": schema.Int64Attribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
						},
						"six_ghz_settings": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"afc_enabled": schema.BoolAttribute{
									Computed: true,
								},
								"channel_width": schema.StringAttribute{
									Computed: true,
								},
								"max_power": schema.Int64Attribute{
									Computed: true,
								},
								"min_bitrate": schema.Int64Attribute{
									Computed: true,
								},
								"min_power": schema.Int64Attribute{
									Computed: true,
								},
								"valid_auto_channels": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
							},
						},
						"transmission": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"enabled": schema.BoolAttribute{
									Computed: true,
								},
							},
						},
						"two_four_ghz_settings": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"ax_enabled": schema.BoolAttribute{
									Computed: true,
								},
								"max_power": schema.Int64Attribute{
									Computed: true,
								},
								"min_bitrate": schema.Int64Attribute{
									Computed: true,
								},
								"min_power": schema.Int64Attribute{
									Computed: true,
								},
								"valid_auto_channels": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
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

		networksWirelessRfProfiles = ResponseWirelessGetNetworkWirelessRfProfilesItemsToBody(networksWirelessRfProfiles, response1)
		diags = resp.State.Set(ctx, &networksWirelessRfProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessRfProfile")
		vvNetworkID := networksWirelessRfProfiles.NetworkID.ValueString()
		vvRfProfileID := networksWirelessRfProfiles.RfProfileID.ValueString()

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
	NetworkID               types.String                                        `tfsdk:"network_id"`
	IncludeTemplateProfiles types.Bool                                          `tfsdk:"include_template_profiles"`
	RfProfileID             types.String                                        `tfsdk:"rf_profile_id"`
	Items                   *[]ResponseItemWirelessGetNetworkWirelessRfProfiles `tfsdk:"items"`
	Item                    *ResponseWirelessGetNetworkWirelessRfProfile        `tfsdk:"item"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfiles struct {
	AfcEnabled             types.Bool                                                          `tfsdk:"afc_enabled"`
	ApBandSettings         *ResponseItemWirelessGetNetworkWirelessRfProfilesApBandSettings     `tfsdk:"ap_band_settings"`
	BandSelectionType      types.String                                                        `tfsdk:"band_selection_type"`
	ClientBalancingEnabled types.Bool                                                          `tfsdk:"client_balancing_enabled"`
	FiveGhzSettings        *ResponseItemWirelessGetNetworkWirelessRfProfilesFiveGhzSettings    `tfsdk:"five_ghz_settings"`
	ID                     types.String                                                        `tfsdk:"id"`
	MinBitrateType         types.String                                                        `tfsdk:"min_bitrate_type"`
	Name                   types.String                                                        `tfsdk:"name"`
	NetworkID              types.String                                                        `tfsdk:"network_id"`
	PerSSIDSettings        *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings    `tfsdk:"per_ssid_settings"`
	SixGhzSettings         *ResponseItemWirelessGetNetworkWirelessRfProfilesSixGhzSettings     `tfsdk:"six_ghz_settings"`
	Transmission           *ResponseItemWirelessGetNetworkWirelessRfProfilesTransmission       `tfsdk:"transmission"`
	TwoFourGhzSettings     *ResponseItemWirelessGetNetworkWirelessRfProfilesTwoFourGhzSettings `tfsdk:"two_four_ghz_settings"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesApBandSettings struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesFiveGhzSettings struct {
	ChannelWidth      types.String `tfsdk:"channel_width"`
	MaxPower          types.Int64  `tfsdk:"max_power"`
	MinBitrate        types.Int64  `tfsdk:"min_bitrate"`
	MinPower          types.Int64  `tfsdk:"min_power"`
	ValidAutoChannels types.List   `tfsdk:"valid_auto_channels"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings struct {
	Status0  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings0  `tfsdk:"status_0"`
	Status1  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings1  `tfsdk:"status_1"`
	Status10 *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings10 `tfsdk:"status_10"`
	Status11 *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings11 `tfsdk:"status_11"`
	Status12 *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings12 `tfsdk:"status_12"`
	Status13 *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings13 `tfsdk:"status_13"`
	Status14 *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings14 `tfsdk:"status_14"`
	Status2  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings2  `tfsdk:"status_2"`
	Status3  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings3  `tfsdk:"status_3"`
	Status4  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings4  `tfsdk:"status_4"`
	Status5  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings5  `tfsdk:"status_5"`
	Status6  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings6  `tfsdk:"status_6"`
	Status7  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings7  `tfsdk:"status_7"`
	Status8  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings8  `tfsdk:"status_8"`
	Status9  *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings9  `tfsdk:"status_9"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings0 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings1 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings10 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings11 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings12 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings13 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings14 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings2 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings3 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings4 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings5 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings6 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings7 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings8 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings9 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesSixGhzSettings struct {
	AfcEnabled        types.Bool   `tfsdk:"afc_enabled"`
	ChannelWidth      types.String `tfsdk:"channel_width"`
	MaxPower          types.Int64  `tfsdk:"max_power"`
	MinBitrate        types.Int64  `tfsdk:"min_bitrate"`
	MinPower          types.Int64  `tfsdk:"min_power"`
	ValidAutoChannels types.List   `tfsdk:"valid_auto_channels"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesTransmission struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseItemWirelessGetNetworkWirelessRfProfilesTwoFourGhzSettings struct {
	AxEnabled         types.Bool  `tfsdk:"ax_enabled"`
	MaxPower          types.Int64 `tfsdk:"max_power"`
	MinBitrate        types.Int64 `tfsdk:"min_bitrate"`
	MinPower          types.Int64 `tfsdk:"min_power"`
	ValidAutoChannels types.List  `tfsdk:"valid_auto_channels"`
}

type ResponseWirelessGetNetworkWirelessRfProfile struct {
	AfcEnabled             types.Bool                                                     `tfsdk:"afc_enabled"`
	ApBandSettings         *ResponseWirelessGetNetworkWirelessRfProfileApBandSettings     `tfsdk:"ap_band_settings"`
	BandSelectionType      types.String                                                   `tfsdk:"band_selection_type"`
	ClientBalancingEnabled types.Bool                                                     `tfsdk:"client_balancing_enabled"`
	FiveGhzSettings        *ResponseWirelessGetNetworkWirelessRfProfileFiveGhzSettings    `tfsdk:"five_ghz_settings"`
	ID                     types.String                                                   `tfsdk:"id"`
	MinBitrateType         types.String                                                   `tfsdk:"min_bitrate_type"`
	Name                   types.String                                                   `tfsdk:"name"`
	NetworkID              types.String                                                   `tfsdk:"network_id"`
	PerSSIDSettings        *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings    `tfsdk:"per_ssid_settings"`
	SixGhzSettings         *ResponseWirelessGetNetworkWirelessRfProfileSixGhzSettings     `tfsdk:"six_ghz_settings"`
	Transmission           *ResponseWirelessGetNetworkWirelessRfProfileTransmission       `tfsdk:"transmission"`
	TwoFourGhzSettings     *ResponseWirelessGetNetworkWirelessRfProfileTwoFourGhzSettings `tfsdk:"two_four_ghz_settings"`
}

type ResponseWirelessGetNetworkWirelessRfProfileApBandSettings struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfileFiveGhzSettings struct {
	ChannelWidth      types.String `tfsdk:"channel_width"`
	MaxPower          types.Int64  `tfsdk:"max_power"`
	MinBitrate        types.Int64  `tfsdk:"min_bitrate"`
	MinPower          types.Int64  `tfsdk:"min_power"`
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
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9 struct {
	BandOperationMode   types.String `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool   `tfsdk:"band_steering_enabled"`
	MinBitrate          types.Int64  `tfsdk:"min_bitrate"`
	Name                types.String `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfileSixGhzSettings struct {
	AfcEnabled        types.Bool   `tfsdk:"afc_enabled"`
	ChannelWidth      types.String `tfsdk:"channel_width"`
	MaxPower          types.Int64  `tfsdk:"max_power"`
	MinBitrate        types.Int64  `tfsdk:"min_bitrate"`
	MinPower          types.Int64  `tfsdk:"min_power"`
	ValidAutoChannels types.List   `tfsdk:"valid_auto_channels"`
}

type ResponseWirelessGetNetworkWirelessRfProfileTransmission struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfileTwoFourGhzSettings struct {
	AxEnabled         types.Bool  `tfsdk:"ax_enabled"`
	MaxPower          types.Int64 `tfsdk:"max_power"`
	MinBitrate        types.Int64 `tfsdk:"min_bitrate"`
	MinPower          types.Int64 `tfsdk:"min_power"`
	ValidAutoChannels types.List  `tfsdk:"valid_auto_channels"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessRfProfilesItemsToBody(state NetworksWirelessRfProfiles, response *merakigosdk.ResponseWirelessGetNetworkWirelessRfProfiles) NetworksWirelessRfProfiles {
	var items []ResponseItemWirelessGetNetworkWirelessRfProfiles
	for _, item := range *response {
		itemState := ResponseItemWirelessGetNetworkWirelessRfProfiles{
			AfcEnabled: func() types.Bool {
				if item.AfcEnabled != nil {
					return types.BoolValue(*item.AfcEnabled)
				}
				return types.Bool{}
			}(),
			ApBandSettings: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesApBandSettings {
				if item.ApBandSettings != nil {
					return &ResponseItemWirelessGetNetworkWirelessRfProfilesApBandSettings{
						BandOperationMode: types.StringValue(item.ApBandSettings.BandOperationMode),
						BandSteeringEnabled: func() types.Bool {
							if item.ApBandSettings.BandSteeringEnabled != nil {
								return types.BoolValue(*item.ApBandSettings.BandSteeringEnabled)
							}
							return types.Bool{}
						}(),
					}
				}
				return &ResponseItemWirelessGetNetworkWirelessRfProfilesApBandSettings{}
			}(),
			BandSelectionType: types.StringValue(item.BandSelectionType),
			ClientBalancingEnabled: func() types.Bool {
				if item.ClientBalancingEnabled != nil {
					return types.BoolValue(*item.ClientBalancingEnabled)
				}
				return types.Bool{}
			}(),
			FiveGhzSettings: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesFiveGhzSettings {
				if item.FiveGhzSettings != nil {
					return &ResponseItemWirelessGetNetworkWirelessRfProfilesFiveGhzSettings{
						ChannelWidth: types.StringValue(item.FiveGhzSettings.ChannelWidth),
						MaxPower: func() types.Int64 {
							if item.FiveGhzSettings.MaxPower != nil {
								return types.Int64Value(int64(*item.FiveGhzSettings.MaxPower))
							}
							return types.Int64{}
						}(),
						MinBitrate: func() types.Int64 {
							if item.FiveGhzSettings.MinBitrate != nil {
								return types.Int64Value(int64(*item.FiveGhzSettings.MinBitrate))
							}
							return types.Int64{}
						}(),
						MinPower: func() types.Int64 {
							if item.FiveGhzSettings.MinPower != nil {
								return types.Int64Value(int64(*item.FiveGhzSettings.MinPower))
							}
							return types.Int64{}
						}(),
						ValidAutoChannels: StringSliceToList(item.FiveGhzSettings.ValidAutoChannels),
					}
				}
				return &ResponseItemWirelessGetNetworkWirelessRfProfilesFiveGhzSettings{}
			}(),
			ID:             types.StringValue(item.ID),
			MinBitrateType: types.StringValue(item.MinBitrateType),
			Name:           types.StringValue(item.Name),
			NetworkID:      types.StringValue(item.NetworkID),
			PerSSIDSettings: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings {
				if item.PerSSIDSettings != nil {
					return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings{
						Status0: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings0 {
							if item.PerSSIDSettings.Status0 != nil {
								return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings0{
									BandOperationMode: types.StringValue(item.PerSSIDSettings.Status0.BandOperationMode),
									BandSteeringEnabled: func() types.Bool {
										if item.PerSSIDSettings.Status0.BandSteeringEnabled != nil {
											return types.BoolValue(*item.PerSSIDSettings.Status0.BandSteeringEnabled)
										}
										return types.Bool{}
									}(),
									MinBitrate: func() types.Int64 {
										if item.PerSSIDSettings.Status0.MinBitrate != nil {
											return types.Int64Value(int64(*item.PerSSIDSettings.Status0.MinBitrate))
										}
										return types.Int64{}
									}(),
									Name: types.StringValue(item.PerSSIDSettings.Status0.Name),
								}
							}
							return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings0{}
						}(),
						Status1: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings1 {
							if item.PerSSIDSettings.Status1 != nil {
								return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings1{
									BandOperationMode: types.StringValue(item.PerSSIDSettings.Status1.BandOperationMode),
									BandSteeringEnabled: func() types.Bool {
										if item.PerSSIDSettings.Status1.BandSteeringEnabled != nil {
											return types.BoolValue(*item.PerSSIDSettings.Status1.BandSteeringEnabled)
										}
										return types.Bool{}
									}(),
									MinBitrate: func() types.Int64 {
										if item.PerSSIDSettings.Status1.MinBitrate != nil {
											return types.Int64Value(int64(*item.PerSSIDSettings.Status1.MinBitrate))
										}
										return types.Int64{}
									}(),
									Name: types.StringValue(item.PerSSIDSettings.Status1.Name),
								}
							}
							return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings1{}
						}(),
						Status10: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings10 {
							if item.PerSSIDSettings.Status10 != nil {
								return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings10{
									BandOperationMode: types.StringValue(item.PerSSIDSettings.Status10.BandOperationMode),
									BandSteeringEnabled: func() types.Bool {
										if item.PerSSIDSettings.Status10.BandSteeringEnabled != nil {
											return types.BoolValue(*item.PerSSIDSettings.Status10.BandSteeringEnabled)
										}
										return types.Bool{}
									}(),
									MinBitrate: func() types.Int64 {
										if item.PerSSIDSettings.Status10.MinBitrate != nil {
											return types.Int64Value(int64(*item.PerSSIDSettings.Status10.MinBitrate))
										}
										return types.Int64{}
									}(),
									Name: types.StringValue(item.PerSSIDSettings.Status10.Name),
								}
							}
							return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings10{}
						}(),
						Status11: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings11 {
							if item.PerSSIDSettings.Status11 != nil {
								return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings11{
									BandOperationMode: types.StringValue(item.PerSSIDSettings.Status11.BandOperationMode),
									BandSteeringEnabled: func() types.Bool {
										if item.PerSSIDSettings.Status11.BandSteeringEnabled != nil {
											return types.BoolValue(*item.PerSSIDSettings.Status11.BandSteeringEnabled)
										}
										return types.Bool{}
									}(),
									MinBitrate: func() types.Int64 {
										if item.PerSSIDSettings.Status11.MinBitrate != nil {
											return types.Int64Value(int64(*item.PerSSIDSettings.Status11.MinBitrate))
										}
										return types.Int64{}
									}(),
									Name: types.StringValue(item.PerSSIDSettings.Status11.Name),
								}
							}
							return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings11{}
						}(),
						Status12: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings12 {
							if item.PerSSIDSettings.Status12 != nil {
								return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings12{
									BandOperationMode: types.StringValue(item.PerSSIDSettings.Status12.BandOperationMode),
									BandSteeringEnabled: func() types.Bool {
										if item.PerSSIDSettings.Status12.BandSteeringEnabled != nil {
											return types.BoolValue(*item.PerSSIDSettings.Status12.BandSteeringEnabled)
										}
										return types.Bool{}
									}(),
									MinBitrate: func() types.Int64 {
										if item.PerSSIDSettings.Status12.MinBitrate != nil {
											return types.Int64Value(int64(*item.PerSSIDSettings.Status12.MinBitrate))
										}
										return types.Int64{}
									}(),
									Name: types.StringValue(item.PerSSIDSettings.Status12.Name),
								}
							}
							return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings12{}
						}(),
						Status13: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings13 {
							if item.PerSSIDSettings.Status13 != nil {
								return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings13{
									BandOperationMode: types.StringValue(item.PerSSIDSettings.Status13.BandOperationMode),
									BandSteeringEnabled: func() types.Bool {
										if item.PerSSIDSettings.Status13.BandSteeringEnabled != nil {
											return types.BoolValue(*item.PerSSIDSettings.Status13.BandSteeringEnabled)
										}
										return types.Bool{}
									}(),
									MinBitrate: func() types.Int64 {
										if item.PerSSIDSettings.Status13.MinBitrate != nil {
											return types.Int64Value(int64(*item.PerSSIDSettings.Status13.MinBitrate))
										}
										return types.Int64{}
									}(),
									Name: types.StringValue(item.PerSSIDSettings.Status13.Name),
								}
							}
							return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings13{}
						}(),
						Status14: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings14 {
							if item.PerSSIDSettings.Status14 != nil {
								return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings14{
									BandOperationMode: types.StringValue(item.PerSSIDSettings.Status14.BandOperationMode),
									BandSteeringEnabled: func() types.Bool {
										if item.PerSSIDSettings.Status14.BandSteeringEnabled != nil {
											return types.BoolValue(*item.PerSSIDSettings.Status14.BandSteeringEnabled)
										}
										return types.Bool{}
									}(),
									MinBitrate: func() types.Int64 {
										if item.PerSSIDSettings.Status14.MinBitrate != nil {
											return types.Int64Value(int64(*item.PerSSIDSettings.Status14.MinBitrate))
										}
										return types.Int64{}
									}(),
									Name: types.StringValue(item.PerSSIDSettings.Status14.Name),
								}
							}
							return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings14{}
						}(),
						Status2: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings2 {
							if item.PerSSIDSettings.Status2 != nil {
								return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings2{
									BandOperationMode: types.StringValue(item.PerSSIDSettings.Status2.BandOperationMode),
									BandSteeringEnabled: func() types.Bool {
										if item.PerSSIDSettings.Status2.BandSteeringEnabled != nil {
											return types.BoolValue(*item.PerSSIDSettings.Status2.BandSteeringEnabled)
										}
										return types.Bool{}
									}(),
									MinBitrate: func() types.Int64 {
										if item.PerSSIDSettings.Status2.MinBitrate != nil {
											return types.Int64Value(int64(*item.PerSSIDSettings.Status2.MinBitrate))
										}
										return types.Int64{}
									}(),
									Name: types.StringValue(item.PerSSIDSettings.Status2.Name),
								}
							}
							return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings2{}
						}(),
						Status3: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings3 {
							if item.PerSSIDSettings.Status3 != nil {
								return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings3{
									BandOperationMode: types.StringValue(item.PerSSIDSettings.Status3.BandOperationMode),
									BandSteeringEnabled: func() types.Bool {
										if item.PerSSIDSettings.Status3.BandSteeringEnabled != nil {
											return types.BoolValue(*item.PerSSIDSettings.Status3.BandSteeringEnabled)
										}
										return types.Bool{}
									}(),
									MinBitrate: func() types.Int64 {
										if item.PerSSIDSettings.Status3.MinBitrate != nil {
											return types.Int64Value(int64(*item.PerSSIDSettings.Status3.MinBitrate))
										}
										return types.Int64{}
									}(),
									Name: types.StringValue(item.PerSSIDSettings.Status3.Name),
								}
							}
							return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings3{}
						}(),
						Status4: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings4 {
							if item.PerSSIDSettings.Status4 != nil {
								return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings4{
									BandOperationMode: types.StringValue(item.PerSSIDSettings.Status4.BandOperationMode),
									BandSteeringEnabled: func() types.Bool {
										if item.PerSSIDSettings.Status4.BandSteeringEnabled != nil {
											return types.BoolValue(*item.PerSSIDSettings.Status4.BandSteeringEnabled)
										}
										return types.Bool{}
									}(),
									MinBitrate: func() types.Int64 {
										if item.PerSSIDSettings.Status4.MinBitrate != nil {
											return types.Int64Value(int64(*item.PerSSIDSettings.Status4.MinBitrate))
										}
										return types.Int64{}
									}(),
									Name: types.StringValue(item.PerSSIDSettings.Status4.Name),
								}
							}
							return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings4{}
						}(),
						Status5: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings5 {
							if item.PerSSIDSettings.Status5 != nil {
								return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings5{
									BandOperationMode: types.StringValue(item.PerSSIDSettings.Status5.BandOperationMode),
									BandSteeringEnabled: func() types.Bool {
										if item.PerSSIDSettings.Status5.BandSteeringEnabled != nil {
											return types.BoolValue(*item.PerSSIDSettings.Status5.BandSteeringEnabled)
										}
										return types.Bool{}
									}(),
									MinBitrate: func() types.Int64 {
										if item.PerSSIDSettings.Status5.MinBitrate != nil {
											return types.Int64Value(int64(*item.PerSSIDSettings.Status5.MinBitrate))
										}
										return types.Int64{}
									}(),
									Name: types.StringValue(item.PerSSIDSettings.Status5.Name),
								}
							}
							return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings5{}
						}(),
						Status6: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings6 {
							if item.PerSSIDSettings.Status6 != nil {
								return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings6{
									BandOperationMode: types.StringValue(item.PerSSIDSettings.Status6.BandOperationMode),
									BandSteeringEnabled: func() types.Bool {
										if item.PerSSIDSettings.Status6.BandSteeringEnabled != nil {
											return types.BoolValue(*item.PerSSIDSettings.Status6.BandSteeringEnabled)
										}
										return types.Bool{}
									}(),
									MinBitrate: func() types.Int64 {
										if item.PerSSIDSettings.Status6.MinBitrate != nil {
											return types.Int64Value(int64(*item.PerSSIDSettings.Status6.MinBitrate))
										}
										return types.Int64{}
									}(),
									Name: types.StringValue(item.PerSSIDSettings.Status6.Name),
								}
							}
							return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings6{}
						}(),
						Status7: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings7 {
							if item.PerSSIDSettings.Status7 != nil {
								return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings7{
									BandOperationMode: types.StringValue(item.PerSSIDSettings.Status7.BandOperationMode),
									BandSteeringEnabled: func() types.Bool {
										if item.PerSSIDSettings.Status7.BandSteeringEnabled != nil {
											return types.BoolValue(*item.PerSSIDSettings.Status7.BandSteeringEnabled)
										}
										return types.Bool{}
									}(),
									MinBitrate: func() types.Int64 {
										if item.PerSSIDSettings.Status7.MinBitrate != nil {
											return types.Int64Value(int64(*item.PerSSIDSettings.Status7.MinBitrate))
										}
										return types.Int64{}
									}(),
									Name: types.StringValue(item.PerSSIDSettings.Status7.Name),
								}
							}
							return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings7{}
						}(),
						Status8: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings8 {
							if item.PerSSIDSettings.Status8 != nil {
								return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings8{
									BandOperationMode: types.StringValue(item.PerSSIDSettings.Status8.BandOperationMode),
									BandSteeringEnabled: func() types.Bool {
										if item.PerSSIDSettings.Status8.BandSteeringEnabled != nil {
											return types.BoolValue(*item.PerSSIDSettings.Status8.BandSteeringEnabled)
										}
										return types.Bool{}
									}(),
									MinBitrate: func() types.Int64 {
										if item.PerSSIDSettings.Status8.MinBitrate != nil {
											return types.Int64Value(int64(*item.PerSSIDSettings.Status8.MinBitrate))
										}
										return types.Int64{}
									}(),
									Name: types.StringValue(item.PerSSIDSettings.Status8.Name),
								}
							}
							return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings8{}
						}(),
						Status9: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings9 {
							if item.PerSSIDSettings.Status9 != nil {
								return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings9{
									BandOperationMode: types.StringValue(item.PerSSIDSettings.Status9.BandOperationMode),
									BandSteeringEnabled: func() types.Bool {
										if item.PerSSIDSettings.Status9.BandSteeringEnabled != nil {
											return types.BoolValue(*item.PerSSIDSettings.Status9.BandSteeringEnabled)
										}
										return types.Bool{}
									}(),
									MinBitrate: func() types.Int64 {
										if item.PerSSIDSettings.Status9.MinBitrate != nil {
											return types.Int64Value(int64(*item.PerSSIDSettings.Status9.MinBitrate))
										}
										return types.Int64{}
									}(),
									Name: types.StringValue(item.PerSSIDSettings.Status9.Name),
								}
							}
							return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings9{}
						}(),
					}
				}
				return &ResponseItemWirelessGetNetworkWirelessRfProfilesPerSsidSettings{}
			}(),
			SixGhzSettings: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesSixGhzSettings {
				if item.SixGhzSettings != nil {
					return &ResponseItemWirelessGetNetworkWirelessRfProfilesSixGhzSettings{
						AfcEnabled: func() types.Bool {
							if item.SixGhzSettings.AfcEnabled != nil {
								return types.BoolValue(*item.SixGhzSettings.AfcEnabled)
							}
							return types.Bool{}
						}(),
						ChannelWidth: types.StringValue(item.SixGhzSettings.ChannelWidth),
						MaxPower: func() types.Int64 {
							if item.SixGhzSettings.MaxPower != nil {
								return types.Int64Value(int64(*item.SixGhzSettings.MaxPower))
							}
							return types.Int64{}
						}(),
						MinBitrate: func() types.Int64 {
							if item.SixGhzSettings.MinBitrate != nil {
								return types.Int64Value(int64(*item.SixGhzSettings.MinBitrate))
							}
							return types.Int64{}
						}(),
						MinPower: func() types.Int64 {
							if item.SixGhzSettings.MinPower != nil {
								return types.Int64Value(int64(*item.SixGhzSettings.MinPower))
							}
							return types.Int64{}
						}(),
						ValidAutoChannels: StringSliceToList(item.SixGhzSettings.ValidAutoChannels),
					}
				}
				return &ResponseItemWirelessGetNetworkWirelessRfProfilesSixGhzSettings{}
			}(),
			Transmission: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesTransmission {
				if item.Transmission != nil {
					return &ResponseItemWirelessGetNetworkWirelessRfProfilesTransmission{
						Enabled: func() types.Bool {
							if item.Transmission.Enabled != nil {
								return types.BoolValue(*item.Transmission.Enabled)
							}
							return types.Bool{}
						}(),
					}
				}
				return &ResponseItemWirelessGetNetworkWirelessRfProfilesTransmission{}
			}(),
			TwoFourGhzSettings: func() *ResponseItemWirelessGetNetworkWirelessRfProfilesTwoFourGhzSettings {
				if item.TwoFourGhzSettings != nil {
					return &ResponseItemWirelessGetNetworkWirelessRfProfilesTwoFourGhzSettings{
						AxEnabled: func() types.Bool {
							if item.TwoFourGhzSettings.AxEnabled != nil {
								return types.BoolValue(*item.TwoFourGhzSettings.AxEnabled)
							}
							return types.Bool{}
						}(),
						MaxPower: func() types.Int64 {
							if item.TwoFourGhzSettings.MaxPower != nil {
								return types.Int64Value(int64(*item.TwoFourGhzSettings.MaxPower))
							}
							return types.Int64{}
						}(),
						MinBitrate: func() types.Int64 {
							if item.TwoFourGhzSettings.MinBitrate != nil {
								return types.Int64Value(int64(*item.TwoFourGhzSettings.MinBitrate))
							}
							return types.Int64{}
						}(),
						MinPower: func() types.Int64 {
							if item.TwoFourGhzSettings.MinPower != nil {
								return types.Int64Value(int64(*item.TwoFourGhzSettings.MinPower))
							}
							return types.Int64{}
						}(),
						ValidAutoChannels: StringSliceToList(item.TwoFourGhzSettings.ValidAutoChannels),
					}
				}
				return &ResponseItemWirelessGetNetworkWirelessRfProfilesTwoFourGhzSettings{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseWirelessGetNetworkWirelessRfProfileItemToBody(state NetworksWirelessRfProfiles, response *merakigosdk.ResponseWirelessGetNetworkWirelessRfProfile) NetworksWirelessRfProfiles {
	itemState := ResponseWirelessGetNetworkWirelessRfProfile{
		AfcEnabled: func() types.Bool {
			if response.AfcEnabled != nil {
				return types.BoolValue(*response.AfcEnabled)
			}
			return types.Bool{}
		}(),
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
				}
			}
			return &ResponseWirelessGetNetworkWirelessRfProfileApBandSettings{}
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
					ValidAutoChannels: StringSliceToList(response.FiveGhzSettings.ValidAutoChannels),
				}
			}
			return &ResponseWirelessGetNetworkWirelessRfProfileFiveGhzSettings{}
		}(),
		ID:             types.StringValue(response.ID),
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
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status0.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status0.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status0.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings0{}
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
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status1.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status1.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status1.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1{}
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
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status10.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status10.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status10.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10{}
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
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status11.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status11.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status11.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11{}
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
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status12.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status12.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status12.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12{}
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
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status13.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status13.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status13.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13{}
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
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status14.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status14.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status14.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14{}
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
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status2.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status2.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status2.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2{}
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
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status3.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status3.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status3.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3{}
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
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status4.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status4.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status4.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4{}
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
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status5.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status5.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status5.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5{}
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
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status6.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status6.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status6.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6{}
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
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status7.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status7.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status7.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7{}
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
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status8.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status8.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status8.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8{}
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
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status9.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status9.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status9.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9{}
					}(),
				}
			}
			return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings{}
		}(),
		SixGhzSettings: func() *ResponseWirelessGetNetworkWirelessRfProfileSixGhzSettings {
			if response.SixGhzSettings != nil {
				return &ResponseWirelessGetNetworkWirelessRfProfileSixGhzSettings{
					AfcEnabled: func() types.Bool {
						if response.SixGhzSettings.AfcEnabled != nil {
							return types.BoolValue(*response.SixGhzSettings.AfcEnabled)
						}
						return types.Bool{}
					}(),
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
					ValidAutoChannels: StringSliceToList(response.SixGhzSettings.ValidAutoChannels),
				}
			}
			return &ResponseWirelessGetNetworkWirelessRfProfileSixGhzSettings{}
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
			return &ResponseWirelessGetNetworkWirelessRfProfileTransmission{}
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
					MinBitrate: func() types.Int64 {
						if response.TwoFourGhzSettings.MinBitrate != nil {
							return types.Int64Value(int64(*response.TwoFourGhzSettings.MinBitrate))
						}
						return types.Int64{}
					}(),
					MinPower: func() types.Int64 {
						if response.TwoFourGhzSettings.MinPower != nil {
							return types.Int64Value(int64(*response.TwoFourGhzSettings.MinPower))
						}
						return types.Int64{}
					}(),
					ValidAutoChannels: StringSliceToList(response.TwoFourGhzSettings.ValidAutoChannels),
				}
			}
			return &ResponseWirelessGetNetworkWirelessRfProfileTwoFourGhzSettings{}
		}(),
	}
	state.Item = &itemState
	return state
}
