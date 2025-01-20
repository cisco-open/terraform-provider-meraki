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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksFirmwareUpgradesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksFirmwareUpgradesDataSource{}
)

func NewNetworksFirmwareUpgradesDataSource() datasource.DataSource {
	return &NetworksFirmwareUpgradesDataSource{}
}

type NetworksFirmwareUpgradesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksFirmwareUpgradesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksFirmwareUpgradesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_firmware_upgrades"
}

func (d *NetworksFirmwareUpgradesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"products": schema.SingleNestedAttribute{
						MarkdownDescription: `The network devices to be updated`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"appliance": schema.SingleNestedAttribute{
								MarkdownDescription: `The network device to be updated`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"available_versions": schema.SetNestedAttribute{
										MarkdownDescription: `Firmware versions available for upgrade`,
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{

												"firmware": schema.StringAttribute{
													MarkdownDescription: `Name of the firmware version`,
													Computed:            true,
												},
												"id": schema.StringAttribute{
													MarkdownDescription: `Firmware version identifier`,
													Computed:            true,
												},
												"release_date": schema.StringAttribute{
													MarkdownDescription: `Release date of the firmware version`,
													Computed:            true,
												},
												"release_type": schema.StringAttribute{
													MarkdownDescription: `Release type of the firmware version`,
													Computed:            true,
												},
												"short_name": schema.StringAttribute{
													MarkdownDescription: `Firmware version short name`,
													Computed:            true,
												},
											},
										},
									},
									"current_version": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the current version on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"firmware": schema.StringAttribute{
												MarkdownDescription: `Name of the firmware version`,
												Computed:            true,
											},
											"id": schema.StringAttribute{
												MarkdownDescription: `Firmware version identifier`,
												Computed:            true,
											},
											"release_date": schema.StringAttribute{
												MarkdownDescription: `Release date of the firmware version`,
												Computed:            true,
											},
											"release_type": schema.StringAttribute{
												MarkdownDescription: `Release type of the firmware version`,
												Computed:            true,
											},
											"short_name": schema.StringAttribute{
												MarkdownDescription: `Firmware version short name`,
												Computed:            true,
											},
										},
									},
									"last_upgrade": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the last firmware upgrade on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"from_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device upgraded from`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
											"time": schema.StringAttribute{
												MarkdownDescription: `Timestamp of the last successful firmware upgrade`,
												Computed:            true,
											},
											"to_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device upgraded to`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
										},
									},
									"next_upgrade": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the next firmware upgrade on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"time": schema.StringAttribute{
												MarkdownDescription: `Timestamp of the next scheduled firmware upgrade`,
												Computed:            true,
											},
											"to_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device will upgrade to if it exists`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
										},
									},
									"participate_in_next_beta_release": schema.BoolAttribute{
										MarkdownDescription: `Whether or not the network wants beta firmware`,
										Computed:            true,
									},
								},
							},
							"camera": schema.SingleNestedAttribute{
								MarkdownDescription: `The network device to be updated`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"available_versions": schema.SetNestedAttribute{
										MarkdownDescription: `Firmware versions available for upgrade`,
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{

												"firmware": schema.StringAttribute{
													MarkdownDescription: `Name of the firmware version`,
													Computed:            true,
												},
												"id": schema.StringAttribute{
													MarkdownDescription: `Firmware version identifier`,
													Computed:            true,
												},
												"release_date": schema.StringAttribute{
													MarkdownDescription: `Release date of the firmware version`,
													Computed:            true,
												},
												"release_type": schema.StringAttribute{
													MarkdownDescription: `Release type of the firmware version`,
													Computed:            true,
												},
												"short_name": schema.StringAttribute{
													MarkdownDescription: `Firmware version short name`,
													Computed:            true,
												},
											},
										},
									},
									"current_version": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the current version on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"firmware": schema.StringAttribute{
												MarkdownDescription: `Name of the firmware version`,
												Computed:            true,
											},
											"id": schema.StringAttribute{
												MarkdownDescription: `Firmware version identifier`,
												Computed:            true,
											},
											"release_date": schema.StringAttribute{
												MarkdownDescription: `Release date of the firmware version`,
												Computed:            true,
											},
											"release_type": schema.StringAttribute{
												MarkdownDescription: `Release type of the firmware version`,
												Computed:            true,
											},
											"short_name": schema.StringAttribute{
												MarkdownDescription: `Firmware version short name`,
												Computed:            true,
											},
										},
									},
									"last_upgrade": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the last firmware upgrade on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"from_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device upgraded from`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
											"time": schema.StringAttribute{
												MarkdownDescription: `Timestamp of the last successful firmware upgrade`,
												Computed:            true,
											},
											"to_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device upgraded to`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
										},
									},
									"next_upgrade": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the next firmware upgrade on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"time": schema.StringAttribute{
												MarkdownDescription: `Timestamp of the next scheduled firmware upgrade`,
												Computed:            true,
											},
											"to_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device will upgrade to if it exists`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
										},
									},
									"participate_in_next_beta_release": schema.BoolAttribute{
										MarkdownDescription: `Whether or not the network wants beta firmware`,
										Computed:            true,
									},
								},
							},
							"cellular_gateway": schema.SingleNestedAttribute{
								MarkdownDescription: `The network device to be updated`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"available_versions": schema.SetNestedAttribute{
										MarkdownDescription: `Firmware versions available for upgrade`,
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{

												"firmware": schema.StringAttribute{
													MarkdownDescription: `Name of the firmware version`,
													Computed:            true,
												},
												"id": schema.StringAttribute{
													MarkdownDescription: `Firmware version identifier`,
													Computed:            true,
												},
												"release_date": schema.StringAttribute{
													MarkdownDescription: `Release date of the firmware version`,
													Computed:            true,
												},
												"release_type": schema.StringAttribute{
													MarkdownDescription: `Release type of the firmware version`,
													Computed:            true,
												},
												"short_name": schema.StringAttribute{
													MarkdownDescription: `Firmware version short name`,
													Computed:            true,
												},
											},
										},
									},
									"current_version": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the current version on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"firmware": schema.StringAttribute{
												MarkdownDescription: `Name of the firmware version`,
												Computed:            true,
											},
											"id": schema.StringAttribute{
												MarkdownDescription: `Firmware version identifier`,
												Computed:            true,
											},
											"release_date": schema.StringAttribute{
												MarkdownDescription: `Release date of the firmware version`,
												Computed:            true,
											},
											"release_type": schema.StringAttribute{
												MarkdownDescription: `Release type of the firmware version`,
												Computed:            true,
											},
											"short_name": schema.StringAttribute{
												MarkdownDescription: `Firmware version short name`,
												Computed:            true,
											},
										},
									},
									"last_upgrade": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the last firmware upgrade on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"from_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device upgraded from`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
											"time": schema.StringAttribute{
												MarkdownDescription: `Timestamp of the last successful firmware upgrade`,
												Computed:            true,
											},
											"to_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device upgraded to`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
										},
									},
									"next_upgrade": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the next firmware upgrade on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"time": schema.StringAttribute{
												MarkdownDescription: `Timestamp of the next scheduled firmware upgrade`,
												Computed:            true,
											},
											"to_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device will upgrade to if it exists`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
										},
									},
									"participate_in_next_beta_release": schema.BoolAttribute{
										MarkdownDescription: `Whether or not the network wants beta firmware`,
										Computed:            true,
									},
								},
							},
							"secure_connect": schema.SingleNestedAttribute{
								MarkdownDescription: `The network device to be updated`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"available_versions": schema.SetNestedAttribute{
										MarkdownDescription: `Firmware versions available for upgrade`,
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{

												"firmware": schema.StringAttribute{
													MarkdownDescription: `Name of the firmware version`,
													Computed:            true,
												},
												"id": schema.StringAttribute{
													MarkdownDescription: `Firmware version identifier`,
													Computed:            true,
												},
												"release_date": schema.StringAttribute{
													MarkdownDescription: `Release date of the firmware version`,
													Computed:            true,
												},
												"release_type": schema.StringAttribute{
													MarkdownDescription: `Release type of the firmware version`,
													Computed:            true,
												},
												"short_name": schema.StringAttribute{
													MarkdownDescription: `Firmware version short name`,
													Computed:            true,
												},
											},
										},
									},
									"current_version": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the current version on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"firmware": schema.StringAttribute{
												MarkdownDescription: `Name of the firmware version`,
												Computed:            true,
											},
											"id": schema.StringAttribute{
												MarkdownDescription: `Firmware version identifier`,
												Computed:            true,
											},
											"release_date": schema.StringAttribute{
												MarkdownDescription: `Release date of the firmware version`,
												Computed:            true,
											},
											"release_type": schema.StringAttribute{
												MarkdownDescription: `Release type of the firmware version`,
												Computed:            true,
											},
											"short_name": schema.StringAttribute{
												MarkdownDescription: `Firmware version short name`,
												Computed:            true,
											},
										},
									},
									"last_upgrade": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the last firmware upgrade on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"from_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device upgraded from`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
											"time": schema.StringAttribute{
												MarkdownDescription: `Timestamp of the last successful firmware upgrade`,
												Computed:            true,
											},
											"to_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device upgraded to`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
										},
									},
									"next_upgrade": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the next firmware upgrade on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"time": schema.StringAttribute{
												MarkdownDescription: `Timestamp of the next scheduled firmware upgrade`,
												Computed:            true,
											},
											"to_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device will upgrade to if it exists`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
										},
									},
									"participate_in_next_beta_release": schema.BoolAttribute{
										MarkdownDescription: `Whether or not the network wants beta firmware`,
										Computed:            true,
									},
								},
							},
							"sensor": schema.SingleNestedAttribute{
								MarkdownDescription: `The network device to be updated`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"available_versions": schema.SetNestedAttribute{
										MarkdownDescription: `Firmware versions available for upgrade`,
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{

												"firmware": schema.StringAttribute{
													MarkdownDescription: `Name of the firmware version`,
													Computed:            true,
												},
												"id": schema.StringAttribute{
													MarkdownDescription: `Firmware version identifier`,
													Computed:            true,
												},
												"release_date": schema.StringAttribute{
													MarkdownDescription: `Release date of the firmware version`,
													Computed:            true,
												},
												"release_type": schema.StringAttribute{
													MarkdownDescription: `Release type of the firmware version`,
													Computed:            true,
												},
												"short_name": schema.StringAttribute{
													MarkdownDescription: `Firmware version short name`,
													Computed:            true,
												},
											},
										},
									},
									"current_version": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the current version on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"firmware": schema.StringAttribute{
												MarkdownDescription: `Name of the firmware version`,
												Computed:            true,
											},
											"id": schema.StringAttribute{
												MarkdownDescription: `Firmware version identifier`,
												Computed:            true,
											},
											"release_date": schema.StringAttribute{
												MarkdownDescription: `Release date of the firmware version`,
												Computed:            true,
											},
											"release_type": schema.StringAttribute{
												MarkdownDescription: `Release type of the firmware version`,
												Computed:            true,
											},
											"short_name": schema.StringAttribute{
												MarkdownDescription: `Firmware version short name`,
												Computed:            true,
											},
										},
									},
									"last_upgrade": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the last firmware upgrade on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"from_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device upgraded from`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
											"time": schema.StringAttribute{
												MarkdownDescription: `Timestamp of the last successful firmware upgrade`,
												Computed:            true,
											},
											"to_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device upgraded to`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
										},
									},
									"next_upgrade": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the next firmware upgrade on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"time": schema.StringAttribute{
												MarkdownDescription: `Timestamp of the next scheduled firmware upgrade`,
												Computed:            true,
											},
											"to_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device will upgrade to if it exists`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
										},
									},
									"participate_in_next_beta_release": schema.BoolAttribute{
										MarkdownDescription: `Whether or not the network wants beta firmware`,
										Computed:            true,
									},
								},
							},
							"switch": schema.SingleNestedAttribute{
								MarkdownDescription: `The network device to be updated`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"available_versions": schema.SetNestedAttribute{
										MarkdownDescription: `Firmware versions available for upgrade`,
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{

												"firmware": schema.StringAttribute{
													MarkdownDescription: `Name of the firmware version`,
													Computed:            true,
												},
												"id": schema.StringAttribute{
													MarkdownDescription: `Firmware version identifier`,
													Computed:            true,
												},
												"release_date": schema.StringAttribute{
													MarkdownDescription: `Release date of the firmware version`,
													Computed:            true,
												},
												"release_type": schema.StringAttribute{
													MarkdownDescription: `Release type of the firmware version`,
													Computed:            true,
												},
												"short_name": schema.StringAttribute{
													MarkdownDescription: `Firmware version short name`,
													Computed:            true,
												},
											},
										},
									},
									"current_version": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the current version on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"firmware": schema.StringAttribute{
												MarkdownDescription: `Name of the firmware version`,
												Computed:            true,
											},
											"id": schema.StringAttribute{
												MarkdownDescription: `Firmware version identifier`,
												Computed:            true,
											},
											"release_date": schema.StringAttribute{
												MarkdownDescription: `Release date of the firmware version`,
												Computed:            true,
											},
											"release_type": schema.StringAttribute{
												MarkdownDescription: `Release type of the firmware version`,
												Computed:            true,
											},
											"short_name": schema.StringAttribute{
												MarkdownDescription: `Firmware version short name`,
												Computed:            true,
											},
										},
									},
									"last_upgrade": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the last firmware upgrade on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"from_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device upgraded from`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
											"time": schema.StringAttribute{
												MarkdownDescription: `Timestamp of the last successful firmware upgrade`,
												Computed:            true,
											},
											"to_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device upgraded to`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
										},
									},
									"next_upgrade": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the next firmware upgrade on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"time": schema.StringAttribute{
												MarkdownDescription: `Timestamp of the next scheduled firmware upgrade`,
												Computed:            true,
											},
											"to_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device will upgrade to if it exists`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
										},
									},
									"participate_in_next_beta_release": schema.BoolAttribute{
										MarkdownDescription: `Whether or not the network wants beta firmware`,
										Computed:            true,
									},
								},
							},
							"wireless": schema.SingleNestedAttribute{
								MarkdownDescription: `The network device to be updated`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"available_versions": schema.SetNestedAttribute{
										MarkdownDescription: `Firmware versions available for upgrade`,
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{

												"firmware": schema.StringAttribute{
													MarkdownDescription: `Name of the firmware version`,
													Computed:            true,
												},
												"id": schema.StringAttribute{
													MarkdownDescription: `Firmware version identifier`,
													Computed:            true,
												},
												"release_date": schema.StringAttribute{
													MarkdownDescription: `Release date of the firmware version`,
													Computed:            true,
												},
												"release_type": schema.StringAttribute{
													MarkdownDescription: `Release type of the firmware version`,
													Computed:            true,
												},
												"short_name": schema.StringAttribute{
													MarkdownDescription: `Firmware version short name`,
													Computed:            true,
												},
											},
										},
									},
									"current_version": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the current version on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"firmware": schema.StringAttribute{
												MarkdownDescription: `Name of the firmware version`,
												Computed:            true,
											},
											"id": schema.StringAttribute{
												MarkdownDescription: `Firmware version identifier`,
												Computed:            true,
											},
											"release_date": schema.StringAttribute{
												MarkdownDescription: `Release date of the firmware version`,
												Computed:            true,
											},
											"release_type": schema.StringAttribute{
												MarkdownDescription: `Release type of the firmware version`,
												Computed:            true,
											},
											"short_name": schema.StringAttribute{
												MarkdownDescription: `Firmware version short name`,
												Computed:            true,
											},
										},
									},
									"last_upgrade": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the last firmware upgrade on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"from_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device upgraded from`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
											"time": schema.StringAttribute{
												MarkdownDescription: `Timestamp of the last successful firmware upgrade`,
												Computed:            true,
											},
											"to_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device upgraded to`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
										},
									},
									"next_upgrade": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the next firmware upgrade on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"time": schema.StringAttribute{
												MarkdownDescription: `Timestamp of the next scheduled firmware upgrade`,
												Computed:            true,
											},
											"to_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device will upgrade to if it exists`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
										},
									},
									"participate_in_next_beta_release": schema.BoolAttribute{
										MarkdownDescription: `Whether or not the network wants beta firmware`,
										Computed:            true,
									},
								},
							},
							"wireless_controller": schema.SingleNestedAttribute{
								MarkdownDescription: `The network device to be updated`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"available_versions": schema.SetNestedAttribute{
										MarkdownDescription: `Firmware versions available for upgrade`,
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{

												"firmware": schema.StringAttribute{
													MarkdownDescription: `Name of the firmware version`,
													Computed:            true,
												},
												"id": schema.StringAttribute{
													MarkdownDescription: `Firmware version identifier`,
													Computed:            true,
												},
												"release_date": schema.StringAttribute{
													MarkdownDescription: `Release date of the firmware version`,
													Computed:            true,
												},
												"release_type": schema.StringAttribute{
													MarkdownDescription: `Release type of the firmware version`,
													Computed:            true,
												},
												"short_name": schema.StringAttribute{
													MarkdownDescription: `Firmware version short name`,
													Computed:            true,
												},
											},
										},
									},
									"current_version": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the current version on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"firmware": schema.StringAttribute{
												MarkdownDescription: `Name of the firmware version`,
												Computed:            true,
											},
											"id": schema.StringAttribute{
												MarkdownDescription: `Firmware version identifier`,
												Computed:            true,
											},
											"release_date": schema.StringAttribute{
												MarkdownDescription: `Release date of the firmware version`,
												Computed:            true,
											},
											"release_type": schema.StringAttribute{
												MarkdownDescription: `Release type of the firmware version`,
												Computed:            true,
											},
											"short_name": schema.StringAttribute{
												MarkdownDescription: `Firmware version short name`,
												Computed:            true,
											},
										},
									},
									"last_upgrade": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the last firmware upgrade on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"from_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device upgraded from`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
											"time": schema.StringAttribute{
												MarkdownDescription: `Timestamp of the last successful firmware upgrade`,
												Computed:            true,
											},
											"to_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device upgraded to`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
										},
									},
									"next_upgrade": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the next firmware upgrade on the device`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"time": schema.StringAttribute{
												MarkdownDescription: `Timestamp of the next scheduled firmware upgrade`,
												Computed:            true,
											},
											"to_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device will upgrade to if it exists`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"firmware": schema.StringAttribute{
														MarkdownDescription: `Name of the firmware version`,
														Computed:            true,
													},
													"id": schema.StringAttribute{
														MarkdownDescription: `Firmware version identifier`,
														Computed:            true,
													},
													"release_date": schema.StringAttribute{
														MarkdownDescription: `Release date of the firmware version`,
														Computed:            true,
													},
													"release_type": schema.StringAttribute{
														MarkdownDescription: `Release type of the firmware version`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
														Computed:            true,
													},
												},
											},
										},
									},
									"participate_in_next_beta_release": schema.BoolAttribute{
										MarkdownDescription: `Whether or not the network wants beta firmware`,
										Computed:            true,
									},
								},
							},
						},
					},
					"timezone": schema.StringAttribute{
						MarkdownDescription: `The timezone for the network`,
						Computed:            true,
					},
					"upgrade_window": schema.SingleNestedAttribute{
						MarkdownDescription: `Upgrade window for devices in network`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"day_of_week": schema.StringAttribute{
								MarkdownDescription: `Day of the week`,
								Computed:            true,
							},
							"hour_of_day": schema.StringAttribute{
								MarkdownDescription: `Hour of the day`,
								Computed:            true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksFirmwareUpgradesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksFirmwareUpgrades NetworksFirmwareUpgrades
	diags := req.Config.Get(ctx, &networksFirmwareUpgrades)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkFirmwareUpgrades")
		vvNetworkID := networksFirmwareUpgrades.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkFirmwareUpgrades(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFirmwareUpgrades",
				err.Error(),
			)
			return
		}

		networksFirmwareUpgrades = ResponseNetworksGetNetworkFirmwareUpgradesItemToBody(networksFirmwareUpgrades, response1)
		diags = resp.State.Set(ctx, &networksFirmwareUpgrades)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksFirmwareUpgrades struct {
	NetworkID types.String                                `tfsdk:"network_id"`
	Item      *ResponseNetworksGetNetworkFirmwareUpgrades `tfsdk:"item"`
}

type ResponseNetworksGetNetworkFirmwareUpgrades struct {
	Products      *ResponseNetworksGetNetworkFirmwareUpgradesProducts      `tfsdk:"products"`
	Timezone      types.String                                             `tfsdk:"timezone"`
	UpgradeWindow *ResponseNetworksGetNetworkFirmwareUpgradesUpgradeWindow `tfsdk:"upgrade_window"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProducts struct {
	Appliance          *ResponseNetworksGetNetworkFirmwareUpgradesProductsAppliance          `tfsdk:"appliance"`
	Camera             *ResponseNetworksGetNetworkFirmwareUpgradesProductsCamera             `tfsdk:"camera"`
	CellularGateway    *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGateway    `tfsdk:"cellular_gateway"`
	SecureConnect      *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnect      `tfsdk:"secure_connect"`
	Sensor             *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensor             `tfsdk:"sensor"`
	Switch             *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitch             `tfsdk:"switch"`
	Wireless           *ResponseNetworksGetNetworkFirmwareUpgradesProductsWireless           `tfsdk:"wireless"`
	WirelessController *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessController `tfsdk:"wireless_controller"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsAppliance struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceAvailableVersions `tfsdk:"available_versions"`
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceCurrentVersion      `tfsdk:"current_version"`
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgrade         `tfsdk:"last_upgrade"`
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgrade         `tfsdk:"next_upgrade"`
	ParticipateInNextBetaRelease types.Bool                                                                      `tfsdk:"participate_in_next_beta_release"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceCurrentVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgrade struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeFromVersion `tfsdk:"from_version"`
	Time        types.String                                                                       `tfsdk:"time"`
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeToVersion   `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeFromVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeToVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgrade struct {
	Time      types.String                                                                     `tfsdk:"time"`
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgradeToVersion `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgradeToVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCamera struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraAvailableVersions `tfsdk:"available_versions"`
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraCurrentVersion      `tfsdk:"current_version"`
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgrade         `tfsdk:"last_upgrade"`
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgrade         `tfsdk:"next_upgrade"`
	ParticipateInNextBetaRelease types.Bool                                                                   `tfsdk:"participate_in_next_beta_release"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraCurrentVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgrade struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeFromVersion `tfsdk:"from_version"`
	Time        types.String                                                                    `tfsdk:"time"`
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeToVersion   `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeFromVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeToVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgrade struct {
	Time      types.String                                                                  `tfsdk:"time"`
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgradeToVersion `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgradeToVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGateway struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayAvailableVersions `tfsdk:"available_versions"`
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayCurrentVersion      `tfsdk:"current_version"`
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgrade         `tfsdk:"last_upgrade"`
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgrade         `tfsdk:"next_upgrade"`
	ParticipateInNextBetaRelease types.Bool                                                                            `tfsdk:"participate_in_next_beta_release"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayCurrentVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgrade struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeFromVersion `tfsdk:"from_version"`
	Time        types.String                                                                             `tfsdk:"time"`
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeToVersion   `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeFromVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeToVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgrade struct {
	Time      types.String                                                                           `tfsdk:"time"`
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeToVersion `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeToVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnect struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectAvailableVersions `tfsdk:"available_versions"`
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectCurrentVersion      `tfsdk:"current_version"`
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgrade         `tfsdk:"last_upgrade"`
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectNextUpgrade         `tfsdk:"next_upgrade"`
	ParticipateInNextBetaRelease types.Bool                                                                          `tfsdk:"participate_in_next_beta_release"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectCurrentVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgrade struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeFromVersion `tfsdk:"from_version"`
	Time        types.String                                                                           `tfsdk:"time"`
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeToVersion   `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeFromVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeToVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectNextUpgrade struct {
	Time      types.String                                                                         `tfsdk:"time"`
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectNextUpgradeToVersion `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectNextUpgradeToVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensor struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorAvailableVersions `tfsdk:"available_versions"`
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorCurrentVersion      `tfsdk:"current_version"`
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgrade         `tfsdk:"last_upgrade"`
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgrade         `tfsdk:"next_upgrade"`
	ParticipateInNextBetaRelease types.Bool                                                                   `tfsdk:"participate_in_next_beta_release"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorCurrentVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgrade struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeFromVersion `tfsdk:"from_version"`
	Time        types.String                                                                    `tfsdk:"time"`
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeToVersion   `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeFromVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeToVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgrade struct {
	Time      types.String                                                                  `tfsdk:"time"`
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgradeToVersion `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgradeToVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitch struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchAvailableVersions `tfsdk:"available_versions"`
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchCurrentVersion      `tfsdk:"current_version"`
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgrade         `tfsdk:"last_upgrade"`
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgrade         `tfsdk:"next_upgrade"`
	ParticipateInNextBetaRelease types.Bool                                                                   `tfsdk:"participate_in_next_beta_release"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchCurrentVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgrade struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeFromVersion `tfsdk:"from_version"`
	Time        types.String                                                                    `tfsdk:"time"`
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeToVersion   `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeFromVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeToVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgrade struct {
	Time      types.String                                                                  `tfsdk:"time"`
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgradeToVersion `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgradeToVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWireless struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessAvailableVersions `tfsdk:"available_versions"`
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessCurrentVersion      `tfsdk:"current_version"`
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgrade         `tfsdk:"last_upgrade"`
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgrade         `tfsdk:"next_upgrade"`
	ParticipateInNextBetaRelease types.Bool                                                                     `tfsdk:"participate_in_next_beta_release"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessCurrentVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgrade struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeFromVersion `tfsdk:"from_version"`
	Time        types.String                                                                      `tfsdk:"time"`
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeToVersion   `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeFromVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeToVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgrade struct {
	Time      types.String                                                                    `tfsdk:"time"`
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgradeToVersion `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgradeToVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessController struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerAvailableVersions `tfsdk:"available_versions"`
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerCurrentVersion      `tfsdk:"current_version"`
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgrade         `tfsdk:"last_upgrade"`
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerNextUpgrade         `tfsdk:"next_upgrade"`
	ParticipateInNextBetaRelease types.Bool                                                                               `tfsdk:"participate_in_next_beta_release"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerCurrentVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgrade struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeFromVersion `tfsdk:"from_version"`
	Time        types.String                                                                                `tfsdk:"time"`
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeToVersion   `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeFromVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeToVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerNextUpgrade struct {
	Time      types.String                                                                              `tfsdk:"time"`
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerNextUpgradeToVersion `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerNextUpgradeToVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesUpgradeWindow struct {
	DayOfWeek types.String `tfsdk:"day_of_week"`
	HourOfDay types.String `tfsdk:"hour_of_day"`
}

// ToBody
func ResponseNetworksGetNetworkFirmwareUpgradesItemToBody(state NetworksFirmwareUpgrades, response *merakigosdk.ResponseNetworksGetNetworkFirmwareUpgrades) NetworksFirmwareUpgrades {
	itemState := ResponseNetworksGetNetworkFirmwareUpgrades{
		Products: func() *ResponseNetworksGetNetworkFirmwareUpgradesProducts {
			if response.Products != nil {
				return &ResponseNetworksGetNetworkFirmwareUpgradesProducts{
					Appliance: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsAppliance {
						if response.Products.Appliance != nil {
							return &ResponseNetworksGetNetworkFirmwareUpgradesProductsAppliance{
								AvailableVersions: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceAvailableVersions {
									if response.Products.Appliance.AvailableVersions != nil {
										result := make([]ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceAvailableVersions, len(*response.Products.Appliance.AvailableVersions))
										for i, availableVersions := range *response.Products.Appliance.AvailableVersions {
											result[i] = ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceAvailableVersions{
												Firmware:    types.StringValue(availableVersions.Firmware),
												ID:          types.StringValue(availableVersions.ID),
												ReleaseDate: types.StringValue(availableVersions.ReleaseDate),
												ReleaseType: types.StringValue(availableVersions.ReleaseType),
												ShortName:   types.StringValue(availableVersions.ShortName),
											}
										}
										return &result
									}
									return nil
								}(),
								CurrentVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceCurrentVersion {
									if response.Products.Appliance.CurrentVersion != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceCurrentVersion{
											Firmware:    types.StringValue(response.Products.Appliance.CurrentVersion.Firmware),
											ID:          types.StringValue(response.Products.Appliance.CurrentVersion.ID),
											ReleaseDate: types.StringValue(response.Products.Appliance.CurrentVersion.ReleaseDate),
											ReleaseType: types.StringValue(response.Products.Appliance.CurrentVersion.ReleaseType),
											ShortName:   types.StringValue(response.Products.Appliance.CurrentVersion.ShortName),
										}
									}
									return nil
								}(),
								LastUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgrade {
									if response.Products.Appliance.LastUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgrade{
											FromVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeFromVersion {
												if response.Products.Appliance.LastUpgrade.FromVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeFromVersion{
														Firmware:    types.StringValue(response.Products.Appliance.LastUpgrade.FromVersion.Firmware),
														ID:          types.StringValue(response.Products.Appliance.LastUpgrade.FromVersion.ID),
														ReleaseDate: types.StringValue(response.Products.Appliance.LastUpgrade.FromVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.Appliance.LastUpgrade.FromVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.Appliance.LastUpgrade.FromVersion.ShortName),
													}
												}
												return nil
											}(),
											Time: types.StringValue(response.Products.Appliance.LastUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeToVersion {
												if response.Products.Appliance.LastUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeToVersion{
														Firmware:    types.StringValue(response.Products.Appliance.LastUpgrade.ToVersion.Firmware),
														ID:          types.StringValue(response.Products.Appliance.LastUpgrade.ToVersion.ID),
														ReleaseDate: types.StringValue(response.Products.Appliance.LastUpgrade.ToVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.Appliance.LastUpgrade.ToVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.Appliance.LastUpgrade.ToVersion.ShortName),
													}
												}
												return nil
											}(),
										}
									}
									return nil
								}(),
								NextUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgrade {
									if response.Products.Appliance.NextUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgrade{
											Time: types.StringValue(response.Products.Appliance.NextUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgradeToVersion {
												if response.Products.Appliance.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgradeToVersion{
														Firmware:    types.StringValue(response.Products.Appliance.NextUpgrade.ToVersion.Firmware),
														ID:          types.StringValue(response.Products.Appliance.NextUpgrade.ToVersion.ID),
														ReleaseDate: types.StringValue(response.Products.Appliance.NextUpgrade.ToVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.Appliance.NextUpgrade.ToVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.Appliance.NextUpgrade.ToVersion.ShortName),
													}
												}
												return nil
											}(),
										}
									}
									return nil
								}(),
								ParticipateInNextBetaRelease: func() types.Bool {
									if response.Products.Appliance.ParticipateInNextBetaRelease != nil {
										return types.BoolValue(*response.Products.Appliance.ParticipateInNextBetaRelease)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
					Camera: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCamera {
						if response.Products.Camera != nil {
							return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCamera{
								AvailableVersions: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraAvailableVersions {
									if response.Products.Camera.AvailableVersions != nil {
										result := make([]ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraAvailableVersions, len(*response.Products.Camera.AvailableVersions))
										for i, availableVersions := range *response.Products.Camera.AvailableVersions {
											result[i] = ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraAvailableVersions{
												Firmware:    types.StringValue(availableVersions.Firmware),
												ID:          types.StringValue(availableVersions.ID),
												ReleaseDate: types.StringValue(availableVersions.ReleaseDate),
												ReleaseType: types.StringValue(availableVersions.ReleaseType),
												ShortName:   types.StringValue(availableVersions.ShortName),
											}
										}
										return &result
									}
									return nil
								}(),
								CurrentVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraCurrentVersion {
									if response.Products.Camera.CurrentVersion != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraCurrentVersion{
											Firmware:    types.StringValue(response.Products.Camera.CurrentVersion.Firmware),
											ID:          types.StringValue(response.Products.Camera.CurrentVersion.ID),
											ReleaseDate: types.StringValue(response.Products.Camera.CurrentVersion.ReleaseDate),
											ReleaseType: types.StringValue(response.Products.Camera.CurrentVersion.ReleaseType),
											ShortName:   types.StringValue(response.Products.Camera.CurrentVersion.ShortName),
										}
									}
									return nil
								}(),
								LastUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgrade {
									if response.Products.Camera.LastUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgrade{
											FromVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeFromVersion {
												if response.Products.Camera.LastUpgrade.FromVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeFromVersion{
														Firmware:    types.StringValue(response.Products.Camera.LastUpgrade.FromVersion.Firmware),
														ID:          types.StringValue(response.Products.Camera.LastUpgrade.FromVersion.ID),
														ReleaseDate: types.StringValue(response.Products.Camera.LastUpgrade.FromVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.Camera.LastUpgrade.FromVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.Camera.LastUpgrade.FromVersion.ShortName),
													}
												}
												return nil
											}(),
											Time: types.StringValue(response.Products.Camera.LastUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeToVersion {
												if response.Products.Camera.LastUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeToVersion{
														Firmware:    types.StringValue(response.Products.Camera.LastUpgrade.ToVersion.Firmware),
														ID:          types.StringValue(response.Products.Camera.LastUpgrade.ToVersion.ID),
														ReleaseDate: types.StringValue(response.Products.Camera.LastUpgrade.ToVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.Camera.LastUpgrade.ToVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.Camera.LastUpgrade.ToVersion.ShortName),
													}
												}
												return nil
											}(),
										}
									}
									return nil
								}(),
								NextUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgrade {
									if response.Products.Camera.NextUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgrade{
											Time: types.StringValue(response.Products.Camera.NextUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgradeToVersion {
												if response.Products.Camera.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgradeToVersion{
														Firmware:    types.StringValue(response.Products.Camera.NextUpgrade.ToVersion.Firmware),
														ID:          types.StringValue(response.Products.Camera.NextUpgrade.ToVersion.ID),
														ReleaseDate: types.StringValue(response.Products.Camera.NextUpgrade.ToVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.Camera.NextUpgrade.ToVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.Camera.NextUpgrade.ToVersion.ShortName),
													}
												}
												return nil
											}(),
										}
									}
									return nil
								}(),
								ParticipateInNextBetaRelease: func() types.Bool {
									if response.Products.Camera.ParticipateInNextBetaRelease != nil {
										return types.BoolValue(*response.Products.Camera.ParticipateInNextBetaRelease)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
					CellularGateway: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGateway {
						if response.Products.CellularGateway != nil {
							return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGateway{
								AvailableVersions: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayAvailableVersions {
									if response.Products.CellularGateway.AvailableVersions != nil {
										result := make([]ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayAvailableVersions, len(*response.Products.CellularGateway.AvailableVersions))
										for i, availableVersions := range *response.Products.CellularGateway.AvailableVersions {
											result[i] = ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayAvailableVersions{
												Firmware:    types.StringValue(availableVersions.Firmware),
												ID:          types.StringValue(availableVersions.ID),
												ReleaseDate: types.StringValue(availableVersions.ReleaseDate),
												ReleaseType: types.StringValue(availableVersions.ReleaseType),
												ShortName:   types.StringValue(availableVersions.ShortName),
											}
										}
										return &result
									}
									return nil
								}(),
								CurrentVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayCurrentVersion {
									if response.Products.CellularGateway.CurrentVersion != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayCurrentVersion{
											Firmware:    types.StringValue(response.Products.CellularGateway.CurrentVersion.Firmware),
											ID:          types.StringValue(response.Products.CellularGateway.CurrentVersion.ID),
											ReleaseDate: types.StringValue(response.Products.CellularGateway.CurrentVersion.ReleaseDate),
											ReleaseType: types.StringValue(response.Products.CellularGateway.CurrentVersion.ReleaseType),
											ShortName:   types.StringValue(response.Products.CellularGateway.CurrentVersion.ShortName),
										}
									}
									return nil
								}(),
								LastUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgrade {
									if response.Products.CellularGateway.LastUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgrade{
											FromVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeFromVersion {
												if response.Products.CellularGateway.LastUpgrade.FromVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeFromVersion{
														Firmware:    types.StringValue(response.Products.CellularGateway.LastUpgrade.FromVersion.Firmware),
														ID:          types.StringValue(response.Products.CellularGateway.LastUpgrade.FromVersion.ID),
														ReleaseDate: types.StringValue(response.Products.CellularGateway.LastUpgrade.FromVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.CellularGateway.LastUpgrade.FromVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.CellularGateway.LastUpgrade.FromVersion.ShortName),
													}
												}
												return nil
											}(),
											Time: types.StringValue(response.Products.CellularGateway.LastUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeToVersion {
												if response.Products.CellularGateway.LastUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeToVersion{
														Firmware:    types.StringValue(response.Products.CellularGateway.LastUpgrade.ToVersion.Firmware),
														ID:          types.StringValue(response.Products.CellularGateway.LastUpgrade.ToVersion.ID),
														ReleaseDate: types.StringValue(response.Products.CellularGateway.LastUpgrade.ToVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.CellularGateway.LastUpgrade.ToVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.CellularGateway.LastUpgrade.ToVersion.ShortName),
													}
												}
												return nil
											}(),
										}
									}
									return nil
								}(),
								NextUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgrade {
									if response.Products.CellularGateway.NextUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgrade{
											Time: types.StringValue(response.Products.CellularGateway.NextUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeToVersion {
												if response.Products.CellularGateway.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeToVersion{
														Firmware:    types.StringValue(response.Products.CellularGateway.NextUpgrade.ToVersion.Firmware),
														ID:          types.StringValue(response.Products.CellularGateway.NextUpgrade.ToVersion.ID),
														ReleaseDate: types.StringValue(response.Products.CellularGateway.NextUpgrade.ToVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.CellularGateway.NextUpgrade.ToVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.CellularGateway.NextUpgrade.ToVersion.ShortName),
													}
												}
												return nil
											}(),
										}
									}
									return nil
								}(),
								ParticipateInNextBetaRelease: func() types.Bool {
									if response.Products.CellularGateway.ParticipateInNextBetaRelease != nil {
										return types.BoolValue(*response.Products.CellularGateway.ParticipateInNextBetaRelease)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
					SecureConnect: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnect {
						if response.Products.SecureConnect != nil {
							return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnect{
								AvailableVersions: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectAvailableVersions {
									if response.Products.SecureConnect.AvailableVersions != nil {
										result := make([]ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectAvailableVersions, len(*response.Products.SecureConnect.AvailableVersions))
										for i, availableVersions := range *response.Products.SecureConnect.AvailableVersions {
											result[i] = ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectAvailableVersions{
												Firmware:    types.StringValue(availableVersions.Firmware),
												ID:          types.StringValue(availableVersions.ID),
												ReleaseDate: types.StringValue(availableVersions.ReleaseDate),
												ReleaseType: types.StringValue(availableVersions.ReleaseType),
												ShortName:   types.StringValue(availableVersions.ShortName),
											}
										}
										return &result
									}
									return nil
								}(),
								CurrentVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectCurrentVersion {
									if response.Products.SecureConnect.CurrentVersion != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectCurrentVersion{
											Firmware:    types.StringValue(response.Products.SecureConnect.CurrentVersion.Firmware),
											ID:          types.StringValue(response.Products.SecureConnect.CurrentVersion.ID),
											ReleaseDate: types.StringValue(response.Products.SecureConnect.CurrentVersion.ReleaseDate),
											ReleaseType: types.StringValue(response.Products.SecureConnect.CurrentVersion.ReleaseType),
											ShortName:   types.StringValue(response.Products.SecureConnect.CurrentVersion.ShortName),
										}
									}
									return nil
								}(),
								LastUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgrade {
									if response.Products.SecureConnect.LastUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgrade{
											FromVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeFromVersion {
												if response.Products.SecureConnect.LastUpgrade.FromVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeFromVersion{
														Firmware:    types.StringValue(response.Products.SecureConnect.LastUpgrade.FromVersion.Firmware),
														ID:          types.StringValue(response.Products.SecureConnect.LastUpgrade.FromVersion.ID),
														ReleaseDate: types.StringValue(response.Products.SecureConnect.LastUpgrade.FromVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.SecureConnect.LastUpgrade.FromVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.SecureConnect.LastUpgrade.FromVersion.ShortName),
													}
												}
												return nil
											}(),
											Time: types.StringValue(response.Products.SecureConnect.LastUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeToVersion {
												if response.Products.SecureConnect.LastUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeToVersion{
														Firmware:    types.StringValue(response.Products.SecureConnect.LastUpgrade.ToVersion.Firmware),
														ID:          types.StringValue(response.Products.SecureConnect.LastUpgrade.ToVersion.ID),
														ReleaseDate: types.StringValue(response.Products.SecureConnect.LastUpgrade.ToVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.SecureConnect.LastUpgrade.ToVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.SecureConnect.LastUpgrade.ToVersion.ShortName),
													}
												}
												return nil
											}(),
										}
									}
									return nil
								}(),
								NextUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectNextUpgrade {
									if response.Products.SecureConnect.NextUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectNextUpgrade{
											Time: types.StringValue(response.Products.SecureConnect.NextUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectNextUpgradeToVersion {
												if response.Products.SecureConnect.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectNextUpgradeToVersion{
														Firmware:    types.StringValue(response.Products.SecureConnect.NextUpgrade.ToVersion.Firmware),
														ID:          types.StringValue(response.Products.SecureConnect.NextUpgrade.ToVersion.ID),
														ReleaseDate: types.StringValue(response.Products.SecureConnect.NextUpgrade.ToVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.SecureConnect.NextUpgrade.ToVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.SecureConnect.NextUpgrade.ToVersion.ShortName),
													}
												}
												return nil
											}(),
										}
									}
									return nil
								}(),
								ParticipateInNextBetaRelease: func() types.Bool {
									if response.Products.SecureConnect.ParticipateInNextBetaRelease != nil {
										return types.BoolValue(*response.Products.SecureConnect.ParticipateInNextBetaRelease)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
					Sensor: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensor {
						if response.Products.Sensor != nil {
							return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSensor{
								AvailableVersions: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorAvailableVersions {
									if response.Products.Sensor.AvailableVersions != nil {
										result := make([]ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorAvailableVersions, len(*response.Products.Sensor.AvailableVersions))
										for i, availableVersions := range *response.Products.Sensor.AvailableVersions {
											result[i] = ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorAvailableVersions{
												Firmware:    types.StringValue(availableVersions.Firmware),
												ID:          types.StringValue(availableVersions.ID),
												ReleaseDate: types.StringValue(availableVersions.ReleaseDate),
												ReleaseType: types.StringValue(availableVersions.ReleaseType),
												ShortName:   types.StringValue(availableVersions.ShortName),
											}
										}
										return &result
									}
									return nil
								}(),
								CurrentVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorCurrentVersion {
									if response.Products.Sensor.CurrentVersion != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorCurrentVersion{
											Firmware:    types.StringValue(response.Products.Sensor.CurrentVersion.Firmware),
											ID:          types.StringValue(response.Products.Sensor.CurrentVersion.ID),
											ReleaseDate: types.StringValue(response.Products.Sensor.CurrentVersion.ReleaseDate),
											ReleaseType: types.StringValue(response.Products.Sensor.CurrentVersion.ReleaseType),
											ShortName:   types.StringValue(response.Products.Sensor.CurrentVersion.ShortName),
										}
									}
									return nil
								}(),
								LastUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgrade {
									if response.Products.Sensor.LastUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgrade{
											FromVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeFromVersion {
												if response.Products.Sensor.LastUpgrade.FromVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeFromVersion{
														Firmware:    types.StringValue(response.Products.Sensor.LastUpgrade.FromVersion.Firmware),
														ID:          types.StringValue(response.Products.Sensor.LastUpgrade.FromVersion.ID),
														ReleaseDate: types.StringValue(response.Products.Sensor.LastUpgrade.FromVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.Sensor.LastUpgrade.FromVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.Sensor.LastUpgrade.FromVersion.ShortName),
													}
												}
												return nil
											}(),
											Time: types.StringValue(response.Products.Sensor.LastUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeToVersion {
												if response.Products.Sensor.LastUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeToVersion{
														Firmware:    types.StringValue(response.Products.Sensor.LastUpgrade.ToVersion.Firmware),
														ID:          types.StringValue(response.Products.Sensor.LastUpgrade.ToVersion.ID),
														ReleaseDate: types.StringValue(response.Products.Sensor.LastUpgrade.ToVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.Sensor.LastUpgrade.ToVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.Sensor.LastUpgrade.ToVersion.ShortName),
													}
												}
												return nil
											}(),
										}
									}
									return nil
								}(),
								NextUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgrade {
									if response.Products.Sensor.NextUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgrade{
											Time: types.StringValue(response.Products.Sensor.NextUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgradeToVersion {
												if response.Products.Sensor.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgradeToVersion{
														Firmware:    types.StringValue(response.Products.Sensor.NextUpgrade.ToVersion.Firmware),
														ID:          types.StringValue(response.Products.Sensor.NextUpgrade.ToVersion.ID),
														ReleaseDate: types.StringValue(response.Products.Sensor.NextUpgrade.ToVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.Sensor.NextUpgrade.ToVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.Sensor.NextUpgrade.ToVersion.ShortName),
													}
												}
												return nil
											}(),
										}
									}
									return nil
								}(),
								ParticipateInNextBetaRelease: func() types.Bool {
									if response.Products.Sensor.ParticipateInNextBetaRelease != nil {
										return types.BoolValue(*response.Products.Sensor.ParticipateInNextBetaRelease)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
					Switch: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitch {
						if response.Products.Switch != nil {
							return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitch{
								AvailableVersions: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchAvailableVersions {
									if response.Products.Switch.AvailableVersions != nil {
										result := make([]ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchAvailableVersions, len(*response.Products.Switch.AvailableVersions))
										for i, availableVersions := range *response.Products.Switch.AvailableVersions {
											result[i] = ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchAvailableVersions{
												Firmware:    types.StringValue(availableVersions.Firmware),
												ID:          types.StringValue(availableVersions.ID),
												ReleaseDate: types.StringValue(availableVersions.ReleaseDate),
												ReleaseType: types.StringValue(availableVersions.ReleaseType),
												ShortName:   types.StringValue(availableVersions.ShortName),
											}
										}
										return &result
									}
									return nil
								}(),
								CurrentVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchCurrentVersion {
									if response.Products.Switch.CurrentVersion != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchCurrentVersion{
											Firmware:    types.StringValue(response.Products.Switch.CurrentVersion.Firmware),
											ID:          types.StringValue(response.Products.Switch.CurrentVersion.ID),
											ReleaseDate: types.StringValue(response.Products.Switch.CurrentVersion.ReleaseDate),
											ReleaseType: types.StringValue(response.Products.Switch.CurrentVersion.ReleaseType),
											ShortName:   types.StringValue(response.Products.Switch.CurrentVersion.ShortName),
										}
									}
									return nil
								}(),
								LastUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgrade {
									if response.Products.Switch.LastUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgrade{
											FromVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeFromVersion {
												if response.Products.Switch.LastUpgrade.FromVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeFromVersion{
														Firmware:    types.StringValue(response.Products.Switch.LastUpgrade.FromVersion.Firmware),
														ID:          types.StringValue(response.Products.Switch.LastUpgrade.FromVersion.ID),
														ReleaseDate: types.StringValue(response.Products.Switch.LastUpgrade.FromVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.Switch.LastUpgrade.FromVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.Switch.LastUpgrade.FromVersion.ShortName),
													}
												}
												return nil
											}(),
											Time: types.StringValue(response.Products.Switch.LastUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeToVersion {
												if response.Products.Switch.LastUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeToVersion{
														Firmware:    types.StringValue(response.Products.Switch.LastUpgrade.ToVersion.Firmware),
														ID:          types.StringValue(response.Products.Switch.LastUpgrade.ToVersion.ID),
														ReleaseDate: types.StringValue(response.Products.Switch.LastUpgrade.ToVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.Switch.LastUpgrade.ToVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.Switch.LastUpgrade.ToVersion.ShortName),
													}
												}
												return nil
											}(),
										}
									}
									return nil
								}(),
								NextUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgrade {
									if response.Products.Switch.NextUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgrade{
											Time: types.StringValue(response.Products.Switch.NextUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgradeToVersion {
												if response.Products.Switch.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgradeToVersion{
														Firmware:    types.StringValue(response.Products.Switch.NextUpgrade.ToVersion.Firmware),
														ID:          types.StringValue(response.Products.Switch.NextUpgrade.ToVersion.ID),
														ReleaseDate: types.StringValue(response.Products.Switch.NextUpgrade.ToVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.Switch.NextUpgrade.ToVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.Switch.NextUpgrade.ToVersion.ShortName),
													}
												}
												return nil
											}(),
										}
									}
									return nil
								}(),
								ParticipateInNextBetaRelease: func() types.Bool {
									if response.Products.Switch.ParticipateInNextBetaRelease != nil {
										return types.BoolValue(*response.Products.Switch.ParticipateInNextBetaRelease)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
					Wireless: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWireless {
						if response.Products.Wireless != nil {
							return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWireless{
								AvailableVersions: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessAvailableVersions {
									if response.Products.Wireless.AvailableVersions != nil {
										result := make([]ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessAvailableVersions, len(*response.Products.Wireless.AvailableVersions))
										for i, availableVersions := range *response.Products.Wireless.AvailableVersions {
											result[i] = ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessAvailableVersions{
												Firmware:    types.StringValue(availableVersions.Firmware),
												ID:          types.StringValue(availableVersions.ID),
												ReleaseDate: types.StringValue(availableVersions.ReleaseDate),
												ReleaseType: types.StringValue(availableVersions.ReleaseType),
												ShortName:   types.StringValue(availableVersions.ShortName),
											}
										}
										return &result
									}
									return nil
								}(),
								CurrentVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessCurrentVersion {
									if response.Products.Wireless.CurrentVersion != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessCurrentVersion{
											Firmware:    types.StringValue(response.Products.Wireless.CurrentVersion.Firmware),
											ID:          types.StringValue(response.Products.Wireless.CurrentVersion.ID),
											ReleaseDate: types.StringValue(response.Products.Wireless.CurrentVersion.ReleaseDate),
											ReleaseType: types.StringValue(response.Products.Wireless.CurrentVersion.ReleaseType),
											ShortName:   types.StringValue(response.Products.Wireless.CurrentVersion.ShortName),
										}
									}
									return nil
								}(),
								LastUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgrade {
									if response.Products.Wireless.LastUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgrade{
											FromVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeFromVersion {
												if response.Products.Wireless.LastUpgrade.FromVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeFromVersion{
														Firmware:    types.StringValue(response.Products.Wireless.LastUpgrade.FromVersion.Firmware),
														ID:          types.StringValue(response.Products.Wireless.LastUpgrade.FromVersion.ID),
														ReleaseDate: types.StringValue(response.Products.Wireless.LastUpgrade.FromVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.Wireless.LastUpgrade.FromVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.Wireless.LastUpgrade.FromVersion.ShortName),
													}
												}
												return nil
											}(),
											Time: types.StringValue(response.Products.Wireless.LastUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeToVersion {
												if response.Products.Wireless.LastUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeToVersion{
														Firmware:    types.StringValue(response.Products.Wireless.LastUpgrade.ToVersion.Firmware),
														ID:          types.StringValue(response.Products.Wireless.LastUpgrade.ToVersion.ID),
														ReleaseDate: types.StringValue(response.Products.Wireless.LastUpgrade.ToVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.Wireless.LastUpgrade.ToVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.Wireless.LastUpgrade.ToVersion.ShortName),
													}
												}
												return nil
											}(),
										}
									}
									return nil
								}(),
								NextUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgrade {
									if response.Products.Wireless.NextUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgrade{
											Time: types.StringValue(response.Products.Wireless.NextUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgradeToVersion {
												if response.Products.Wireless.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgradeToVersion{
														Firmware:    types.StringValue(response.Products.Wireless.NextUpgrade.ToVersion.Firmware),
														ID:          types.StringValue(response.Products.Wireless.NextUpgrade.ToVersion.ID),
														ReleaseDate: types.StringValue(response.Products.Wireless.NextUpgrade.ToVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.Wireless.NextUpgrade.ToVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.Wireless.NextUpgrade.ToVersion.ShortName),
													}
												}
												return nil
											}(),
										}
									}
									return nil
								}(),
								ParticipateInNextBetaRelease: func() types.Bool {
									if response.Products.Wireless.ParticipateInNextBetaRelease != nil {
										return types.BoolValue(*response.Products.Wireless.ParticipateInNextBetaRelease)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
					WirelessController: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessController {
						if response.Products.WirelessController != nil {
							return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessController{
								AvailableVersions: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerAvailableVersions {
									if response.Products.WirelessController.AvailableVersions != nil {
										result := make([]ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerAvailableVersions, len(*response.Products.WirelessController.AvailableVersions))
										for i, availableVersions := range *response.Products.WirelessController.AvailableVersions {
											result[i] = ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerAvailableVersions{
												Firmware:    types.StringValue(availableVersions.Firmware),
												ID:          types.StringValue(availableVersions.ID),
												ReleaseDate: types.StringValue(availableVersions.ReleaseDate),
												ReleaseType: types.StringValue(availableVersions.ReleaseType),
												ShortName:   types.StringValue(availableVersions.ShortName),
											}
										}
										return &result
									}
									return nil
								}(),
								CurrentVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerCurrentVersion {
									if response.Products.WirelessController.CurrentVersion != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerCurrentVersion{
											Firmware:    types.StringValue(response.Products.WirelessController.CurrentVersion.Firmware),
											ID:          types.StringValue(response.Products.WirelessController.CurrentVersion.ID),
											ReleaseDate: types.StringValue(response.Products.WirelessController.CurrentVersion.ReleaseDate),
											ReleaseType: types.StringValue(response.Products.WirelessController.CurrentVersion.ReleaseType),
											ShortName:   types.StringValue(response.Products.WirelessController.CurrentVersion.ShortName),
										}
									}
									return nil
								}(),
								LastUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgrade {
									if response.Products.WirelessController.LastUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgrade{
											FromVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeFromVersion {
												if response.Products.WirelessController.LastUpgrade.FromVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeFromVersion{
														Firmware:    types.StringValue(response.Products.WirelessController.LastUpgrade.FromVersion.Firmware),
														ID:          types.StringValue(response.Products.WirelessController.LastUpgrade.FromVersion.ID),
														ReleaseDate: types.StringValue(response.Products.WirelessController.LastUpgrade.FromVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.WirelessController.LastUpgrade.FromVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.WirelessController.LastUpgrade.FromVersion.ShortName),
													}
												}
												return nil
											}(),
											Time: types.StringValue(response.Products.WirelessController.LastUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeToVersion {
												if response.Products.WirelessController.LastUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeToVersion{
														Firmware:    types.StringValue(response.Products.WirelessController.LastUpgrade.ToVersion.Firmware),
														ID:          types.StringValue(response.Products.WirelessController.LastUpgrade.ToVersion.ID),
														ReleaseDate: types.StringValue(response.Products.WirelessController.LastUpgrade.ToVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.WirelessController.LastUpgrade.ToVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.WirelessController.LastUpgrade.ToVersion.ShortName),
													}
												}
												return nil
											}(),
										}
									}
									return nil
								}(),
								NextUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerNextUpgrade {
									if response.Products.WirelessController.NextUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerNextUpgrade{
											Time: types.StringValue(response.Products.WirelessController.NextUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerNextUpgradeToVersion {
												if response.Products.WirelessController.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerNextUpgradeToVersion{
														Firmware:    types.StringValue(response.Products.WirelessController.NextUpgrade.ToVersion.Firmware),
														ID:          types.StringValue(response.Products.WirelessController.NextUpgrade.ToVersion.ID),
														ReleaseDate: types.StringValue(response.Products.WirelessController.NextUpgrade.ToVersion.ReleaseDate),
														ReleaseType: types.StringValue(response.Products.WirelessController.NextUpgrade.ToVersion.ReleaseType),
														ShortName:   types.StringValue(response.Products.WirelessController.NextUpgrade.ToVersion.ShortName),
													}
												}
												return nil
											}(),
										}
									}
									return nil
								}(),
								ParticipateInNextBetaRelease: func() types.Bool {
									if response.Products.WirelessController.ParticipateInNextBetaRelease != nil {
										return types.BoolValue(*response.Products.WirelessController.ParticipateInNextBetaRelease)
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
		Timezone: types.StringValue(response.Timezone),
		UpgradeWindow: func() *ResponseNetworksGetNetworkFirmwareUpgradesUpgradeWindow {
			if response.UpgradeWindow != nil {
				return &ResponseNetworksGetNetworkFirmwareUpgradesUpgradeWindow{
					DayOfWeek: types.StringValue(response.UpgradeWindow.DayOfWeek),
					HourOfDay: types.StringValue(response.UpgradeWindow.HourOfDay),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
