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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksFirmwareUpgradesResource{}
	_ resource.ResourceWithConfigure = &NetworksFirmwareUpgradesResource{}
)

func NewNetworksFirmwareUpgradesResource() resource.Resource {
	return &NetworksFirmwareUpgradesResource{}
}

type NetworksFirmwareUpgradesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksFirmwareUpgradesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksFirmwareUpgradesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_firmware_upgrades"
}

func (r *NetworksFirmwareUpgradesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"products": schema.SingleNestedAttribute{
				MarkdownDescription: `The network devices to be updated`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"appliance": schema.SingleNestedAttribute{
						MarkdownDescription: `The network device to be updated`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
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
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"time": schema.StringAttribute{
										MarkdownDescription: `Timestamp of the next scheduled firmware upgrade`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"to_version": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the version the device will upgrade to if it exists`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"firmware": schema.StringAttribute{
												MarkdownDescription: `Name of the firmware version`,
												Computed:            true,
											},
											"id": schema.StringAttribute{
												MarkdownDescription: `Firmware version identifier`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
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
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"camera": schema.SingleNestedAttribute{
						MarkdownDescription: `The network device to be updated`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
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
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"time": schema.StringAttribute{
										MarkdownDescription: `Timestamp of the next scheduled firmware upgrade`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"to_version": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the version the device will upgrade to if it exists`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"firmware": schema.StringAttribute{
												MarkdownDescription: `Name of the firmware version`,
												Computed:            true,
											},
											"id": schema.StringAttribute{
												MarkdownDescription: `Firmware version identifier`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
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
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"cellular_gateway": schema.SingleNestedAttribute{
						MarkdownDescription: `The network device to be updated`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
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
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"time": schema.StringAttribute{
										MarkdownDescription: `Timestamp of the next scheduled firmware upgrade`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"to_version": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the version the device will upgrade to if it exists`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"firmware": schema.StringAttribute{
												MarkdownDescription: `Name of the firmware version`,
												Computed:            true,
											},
											"id": schema.StringAttribute{
												MarkdownDescription: `Firmware version identifier`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
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
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"secure_connect": schema.SingleNestedAttribute{
						MarkdownDescription: `The network device to be updated`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
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
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"time": schema.StringAttribute{
										MarkdownDescription: `Timestamp of the next scheduled firmware upgrade`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"to_version": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the version the device will upgrade to if it exists`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"firmware": schema.StringAttribute{
												MarkdownDescription: `Name of the firmware version`,
												Computed:            true,
											},
											"id": schema.StringAttribute{
												MarkdownDescription: `Firmware version identifier`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
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
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"sensor": schema.SingleNestedAttribute{
						MarkdownDescription: `The network device to be updated`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
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
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"time": schema.StringAttribute{
										MarkdownDescription: `Timestamp of the next scheduled firmware upgrade`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"to_version": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the version the device will upgrade to if it exists`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"firmware": schema.StringAttribute{
												MarkdownDescription: `Name of the firmware version`,
												Computed:            true,
											},
											"id": schema.StringAttribute{
												MarkdownDescription: `Firmware version identifier`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
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
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"switch": schema.SingleNestedAttribute{
						MarkdownDescription: `The network device to be updated`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
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
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"time": schema.StringAttribute{
										MarkdownDescription: `Timestamp of the next scheduled firmware upgrade`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"to_version": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the version the device will upgrade to if it exists`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"firmware": schema.StringAttribute{
												MarkdownDescription: `Name of the firmware version`,
												Computed:            true,
											},
											"id": schema.StringAttribute{
												MarkdownDescription: `Firmware version identifier`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
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
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"switch_catalyst": schema.SingleNestedAttribute{
						MarkdownDescription: `The network device to be updated`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"next_upgrade": schema.SingleNestedAttribute{
								MarkdownDescription: `The pending firmware upgrade if it exists`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"time": schema.StringAttribute{
										MarkdownDescription: `The time of the last successful upgrade`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"to_version": schema.SingleNestedAttribute{
										MarkdownDescription: `The version to be updated to`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"id": schema.StringAttribute{
												MarkdownDescription: `The version ID`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
										},
									},
								},
							},
							"participate_in_next_beta_release": schema.BoolAttribute{
								MarkdownDescription: `Whether or not the network wants beta firmware`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"wireless": schema.SingleNestedAttribute{
						MarkdownDescription: `The network device to be updated`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
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
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"time": schema.StringAttribute{
										MarkdownDescription: `Timestamp of the next scheduled firmware upgrade`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"to_version": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the version the device will upgrade to if it exists`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"firmware": schema.StringAttribute{
												MarkdownDescription: `Name of the firmware version`,
												Computed:            true,
											},
											"id": schema.StringAttribute{
												MarkdownDescription: `Firmware version identifier`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
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
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"wireless_controller": schema.SingleNestedAttribute{
						MarkdownDescription: `The network device to be updated`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
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
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"time": schema.StringAttribute{
										MarkdownDescription: `Timestamp of the next scheduled firmware upgrade`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"to_version": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the version the device will upgrade to if it exists`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"firmware": schema.StringAttribute{
												MarkdownDescription: `Name of the firmware version`,
												Computed:            true,
											},
											"id": schema.StringAttribute{
												MarkdownDescription: `Firmware version identifier`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
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
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
				},
			},
			"timezone": schema.StringAttribute{
				MarkdownDescription: `The timezone for the network`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"upgrade_window": schema.SingleNestedAttribute{
				MarkdownDescription: `Upgrade window for devices in network`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"day_of_week": schema.StringAttribute{
						MarkdownDescription: `Day of the week
                                        Allowed values: [fri,friday,mon,monday,sat,saturday,sun,sunday,thu,thursday,tue,tuesday,wed,wednesday]`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"fri",
								"friday",
								"mon",
								"monday",
								"sat",
								"saturday",
								"sun",
								"sunday",
								"thu",
								"thursday",
								"tue",
								"tuesday",
								"wed",
								"wednesday",
							),
						},
					},
					"hour_of_day": schema.StringAttribute{
						MarkdownDescription: `Hour of the day
                                        Allowed values: [0:00,10:00,11:00,12:00,13:00,14:00,15:00,16:00,17:00,18:00,19:00,1:00,20:00,21:00,22:00,23:00,2:00,3:00,4:00,5:00,6:00,7:00,8:00,9:00]`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"0:00",
								"10:00",
								"11:00",
								"12:00",
								"13:00",
								"14:00",
								"15:00",
								"16:00",
								"17:00",
								"18:00",
								"19:00",
								"1:00",
								"20:00",
								"21:00",
								"22:00",
								"23:00",
								"2:00",
								"3:00",
								"4:00",
								"5:00",
								"6:00",
								"7:00",
								"8:00",
								"9:00",
							),
						},
					},
				},
			},
		},
	}
}

func (r *NetworksFirmwareUpgradesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksFirmwareUpgradesRs

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
		responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkFirmwareUpgrades(vvNetworkID)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksFirmwareUpgrades  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksFirmwareUpgrades only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkFirmwareUpgrades(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkFirmwareUpgrades",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkFirmwareUpgrades",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Networks.GetNetworkFirmwareUpgrades(vvNetworkID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFirmwareUpgrades",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkFirmwareUpgrades",
			err.Error(),
		)
		return
	}

	data = ResponseNetworksGetNetworkFirmwareUpgradesItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *NetworksFirmwareUpgradesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksFirmwareUpgradesRs

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
	responseGet, restyRespGet, err := r.client.Networks.GetNetworkFirmwareUpgrades(vvNetworkID)
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
				"Failure when executing GetNetworkFirmwareUpgrades",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkFirmwareUpgrades",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkFirmwareUpgradesItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksFirmwareUpgradesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksFirmwareUpgradesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksFirmwareUpgradesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkFirmwareUpgrades(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkFirmwareUpgrades",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkFirmwareUpgrades",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksFirmwareUpgradesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksFirmwareUpgrades", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksFirmwareUpgradesRs struct {
	NetworkID     types.String                                               `tfsdk:"network_id"`
	Products      *ResponseNetworksGetNetworkFirmwareUpgradesProductsRs      `tfsdk:"products"`
	Timezone      types.String                                               `tfsdk:"timezone"`
	UpgradeWindow *ResponseNetworksGetNetworkFirmwareUpgradesUpgradeWindowRs `tfsdk:"upgrade_window"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsRs struct {
	Appliance          *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceRs          `tfsdk:"appliance"`
	Camera             *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraRs             `tfsdk:"camera"`
	CellularGateway    *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayRs    `tfsdk:"cellular_gateway"`
	SecureConnect      *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectRs      `tfsdk:"secure_connect"`
	Sensor             *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorRs             `tfsdk:"sensor"`
	Switch             *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchRs             `tfsdk:"switch"`
	Wireless           *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessRs           `tfsdk:"wireless"`
	WirelessController *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerRs `tfsdk:"wireless_controller"`
	SwitchCatalyst     *RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalystRs   `tfsdk:"switch_catalyst"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceRs struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceAvailableVersionsRs `tfsdk:"available_versions"`
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceCurrentVersionRs      `tfsdk:"current_version"`
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeRs         `tfsdk:"last_upgrade"`
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgradeRs         `tfsdk:"next_upgrade"`
	ParticipateInNextBetaRelease types.Bool                                                                        `tfsdk:"participate_in_next_beta_release"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceAvailableVersionsRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceCurrentVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeRs struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeFromVersionRs `tfsdk:"from_version"`
	Time        types.String                                                                         `tfsdk:"time"`
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeToVersionRs   `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeFromVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeToVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgradeRs struct {
	Time      types.String                                                                       `tfsdk:"time"`
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgradeToVersionRs `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgradeToVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraRs struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraAvailableVersionsRs `tfsdk:"available_versions"`
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraCurrentVersionRs      `tfsdk:"current_version"`
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeRs         `tfsdk:"last_upgrade"`
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgradeRs         `tfsdk:"next_upgrade"`
	ParticipateInNextBetaRelease types.Bool                                                                     `tfsdk:"participate_in_next_beta_release"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraAvailableVersionsRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraCurrentVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeRs struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeFromVersionRs `tfsdk:"from_version"`
	Time        types.String                                                                      `tfsdk:"time"`
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeToVersionRs   `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeFromVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeToVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgradeRs struct {
	Time      types.String                                                                    `tfsdk:"time"`
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgradeToVersionRs `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgradeToVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayRs struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayAvailableVersionsRs `tfsdk:"available_versions"`
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayCurrentVersionRs      `tfsdk:"current_version"`
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeRs         `tfsdk:"last_upgrade"`
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeRs         `tfsdk:"next_upgrade"`
	ParticipateInNextBetaRelease types.Bool                                                                              `tfsdk:"participate_in_next_beta_release"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayAvailableVersionsRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayCurrentVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeRs struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeFromVersionRs `tfsdk:"from_version"`
	Time        types.String                                                                               `tfsdk:"time"`
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeToVersionRs   `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeFromVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeToVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeRs struct {
	Time      types.String                                                                             `tfsdk:"time"`
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeToVersionRs `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeToVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectRs struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectAvailableVersionsRs `tfsdk:"available_versions"`
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectCurrentVersionRs      `tfsdk:"current_version"`
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeRs         `tfsdk:"last_upgrade"`
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectNextUpgradeRs         `tfsdk:"next_upgrade"`
	ParticipateInNextBetaRelease types.Bool                                                                            `tfsdk:"participate_in_next_beta_release"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectAvailableVersionsRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectCurrentVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeRs struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeFromVersionRs `tfsdk:"from_version"`
	Time        types.String                                                                             `tfsdk:"time"`
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeToVersionRs   `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeFromVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeToVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectNextUpgradeRs struct {
	Time      types.String                                                                           `tfsdk:"time"`
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectNextUpgradeToVersionRs `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectNextUpgradeToVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorRs struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorAvailableVersionsRs `tfsdk:"available_versions"`
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorCurrentVersionRs      `tfsdk:"current_version"`
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeRs         `tfsdk:"last_upgrade"`
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgradeRs         `tfsdk:"next_upgrade"`
	ParticipateInNextBetaRelease types.Bool                                                                     `tfsdk:"participate_in_next_beta_release"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorAvailableVersionsRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorCurrentVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeRs struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeFromVersionRs `tfsdk:"from_version"`
	Time        types.String                                                                      `tfsdk:"time"`
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeToVersionRs   `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeFromVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeToVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgradeRs struct {
	Time      types.String                                                                    `tfsdk:"time"`
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgradeToVersionRs `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgradeToVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchRs struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchAvailableVersionsRs `tfsdk:"available_versions"`
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchCurrentVersionRs      `tfsdk:"current_version"`
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeRs         `tfsdk:"last_upgrade"`
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgradeRs         `tfsdk:"next_upgrade"`
	ParticipateInNextBetaRelease types.Bool                                                                     `tfsdk:"participate_in_next_beta_release"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchAvailableVersionsRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchCurrentVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeRs struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeFromVersionRs `tfsdk:"from_version"`
	Time        types.String                                                                      `tfsdk:"time"`
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeToVersionRs   `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeFromVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeToVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgradeRs struct {
	Time      types.String                                                                    `tfsdk:"time"`
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgradeToVersionRs `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgradeToVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessRs struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessAvailableVersionsRs `tfsdk:"available_versions"`
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessCurrentVersionRs      `tfsdk:"current_version"`
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeRs         `tfsdk:"last_upgrade"`
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgradeRs         `tfsdk:"next_upgrade"`
	ParticipateInNextBetaRelease types.Bool                                                                       `tfsdk:"participate_in_next_beta_release"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessAvailableVersionsRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessCurrentVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeRs struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeFromVersionRs `tfsdk:"from_version"`
	Time        types.String                                                                        `tfsdk:"time"`
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeToVersionRs   `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeFromVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeToVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgradeRs struct {
	Time      types.String                                                                      `tfsdk:"time"`
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgradeToVersionRs `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgradeToVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerRs struct {
	AvailableVersions            *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerAvailableVersionsRs `tfsdk:"available_versions"`
	CurrentVersion               *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerCurrentVersionRs      `tfsdk:"current_version"`
	LastUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeRs         `tfsdk:"last_upgrade"`
	NextUpgrade                  *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerNextUpgradeRs         `tfsdk:"next_upgrade"`
	ParticipateInNextBetaRelease types.Bool                                                                                 `tfsdk:"participate_in_next_beta_release"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerAvailableVersionsRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerCurrentVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeRs struct {
	FromVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeFromVersionRs `tfsdk:"from_version"`
	Time        types.String                                                                                  `tfsdk:"time"`
	ToVersion   *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeToVersionRs   `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeFromVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeToVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerNextUpgradeRs struct {
	Time      types.String                                                                                `tfsdk:"time"`
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerNextUpgradeToVersionRs `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerNextUpgradeToVersionRs struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesUpgradeWindowRs struct {
	DayOfWeek types.String `tfsdk:"day_of_week"`
	HourOfDay types.String `tfsdk:"hour_of_day"`
}

type RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalystRs struct {
	NextUpgrade                  *RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalystNextUpgradeRs `tfsdk:"next_upgrade"`
	ParticipateInNextBetaRelease types.Bool                                                                       `tfsdk:"participate_in_next_beta_release"`
}

type RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalystNextUpgradeRs struct {
	Time      types.String                                                                              `tfsdk:"time"`
	ToVersion *RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalystNextUpgradeToVersionRs `tfsdk:"to_version"`
}

type RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalystNextUpgradeToVersionRs struct {
	ID types.String `tfsdk:"id"`
}

// FromBody
func (r *NetworksFirmwareUpgradesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgrades {
	emptyString := ""
	var requestNetworksUpdateNetworkFirmwareUpgradesProducts *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProducts

	if r.Products != nil {
		var requestNetworksUpdateNetworkFirmwareUpgradesProductsAppliance *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsAppliance

		if r.Products.Appliance != nil {
			var requestNetworksUpdateNetworkFirmwareUpgradesProductsApplianceNextUpgrade *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsApplianceNextUpgrade

			if r.Products.Appliance.NextUpgrade != nil {
				time := r.Products.Appliance.NextUpgrade.Time.ValueString()
				var requestNetworksUpdateNetworkFirmwareUpgradesProductsApplianceNextUpgradeToVersion *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsApplianceNextUpgradeToVersion

				if r.Products.Appliance.NextUpgrade.ToVersion != nil {
					id := r.Products.Appliance.NextUpgrade.ToVersion.ID.ValueString()
					requestNetworksUpdateNetworkFirmwareUpgradesProductsApplianceNextUpgradeToVersion = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsApplianceNextUpgradeToVersion{
						ID: id,
					}
					//[debug] Is Array: False
				}
				requestNetworksUpdateNetworkFirmwareUpgradesProductsApplianceNextUpgrade = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsApplianceNextUpgrade{
					Time:      time,
					ToVersion: requestNetworksUpdateNetworkFirmwareUpgradesProductsApplianceNextUpgradeToVersion,
				}
				//[debug] Is Array: False
			}
			participateInNextBetaRelease := func() *bool {
				if !r.Products.Appliance.ParticipateInNextBetaRelease.IsUnknown() && !r.Products.Appliance.ParticipateInNextBetaRelease.IsNull() {
					return r.Products.Appliance.ParticipateInNextBetaRelease.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkFirmwareUpgradesProductsAppliance = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsAppliance{
				NextUpgrade:                  requestNetworksUpdateNetworkFirmwareUpgradesProductsApplianceNextUpgrade,
				ParticipateInNextBetaRelease: participateInNextBetaRelease,
			}
			//[debug] Is Array: False
		}
		var requestNetworksUpdateNetworkFirmwareUpgradesProductsCamera *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsCamera

		if r.Products.Camera != nil {
			var requestNetworksUpdateNetworkFirmwareUpgradesProductsCameraNextUpgrade *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsCameraNextUpgrade

			if r.Products.Camera.NextUpgrade != nil {
				time := r.Products.Camera.NextUpgrade.Time.ValueString()
				var requestNetworksUpdateNetworkFirmwareUpgradesProductsCameraNextUpgradeToVersion *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsCameraNextUpgradeToVersion

				if r.Products.Camera.NextUpgrade.ToVersion != nil {
					id := r.Products.Camera.NextUpgrade.ToVersion.ID.ValueString()
					requestNetworksUpdateNetworkFirmwareUpgradesProductsCameraNextUpgradeToVersion = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsCameraNextUpgradeToVersion{
						ID: id,
					}
					//[debug] Is Array: False
				}
				requestNetworksUpdateNetworkFirmwareUpgradesProductsCameraNextUpgrade = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsCameraNextUpgrade{
					Time:      time,
					ToVersion: requestNetworksUpdateNetworkFirmwareUpgradesProductsCameraNextUpgradeToVersion,
				}
				//[debug] Is Array: False
			}
			participateInNextBetaRelease := func() *bool {
				if !r.Products.Camera.ParticipateInNextBetaRelease.IsUnknown() && !r.Products.Camera.ParticipateInNextBetaRelease.IsNull() {
					return r.Products.Camera.ParticipateInNextBetaRelease.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkFirmwareUpgradesProductsCamera = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsCamera{
				NextUpgrade:                  requestNetworksUpdateNetworkFirmwareUpgradesProductsCameraNextUpgrade,
				ParticipateInNextBetaRelease: participateInNextBetaRelease,
			}
			//[debug] Is Array: False
		}
		var requestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGateway *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGateway

		if r.Products.CellularGateway != nil {
			var requestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayNextUpgrade *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayNextUpgrade

			if r.Products.CellularGateway.NextUpgrade != nil {
				time := r.Products.CellularGateway.NextUpgrade.Time.ValueString()
				var requestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeToVersion *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeToVersion

				if r.Products.CellularGateway.NextUpgrade.ToVersion != nil {
					id := r.Products.CellularGateway.NextUpgrade.ToVersion.ID.ValueString()
					requestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeToVersion = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeToVersion{
						ID: id,
					}
					//[debug] Is Array: False
				}
				requestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayNextUpgrade = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayNextUpgrade{
					Time:      time,
					ToVersion: requestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeToVersion,
				}
				//[debug] Is Array: False
			}
			participateInNextBetaRelease := func() *bool {
				if !r.Products.CellularGateway.ParticipateInNextBetaRelease.IsUnknown() && !r.Products.CellularGateway.ParticipateInNextBetaRelease.IsNull() {
					return r.Products.CellularGateway.ParticipateInNextBetaRelease.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGateway = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGateway{
				NextUpgrade:                  requestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGatewayNextUpgrade,
				ParticipateInNextBetaRelease: participateInNextBetaRelease,
			}
			//[debug] Is Array: False
		}
		var requestNetworksUpdateNetworkFirmwareUpgradesProductsSecureConnect *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSecureConnect

		if r.Products.SecureConnect != nil {
			var requestNetworksUpdateNetworkFirmwareUpgradesProductsSecureConnectNextUpgrade *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSecureConnectNextUpgrade

			if r.Products.SecureConnect.NextUpgrade != nil {
				time := r.Products.SecureConnect.NextUpgrade.Time.ValueString()
				var requestNetworksUpdateNetworkFirmwareUpgradesProductsSecureConnectNextUpgradeToVersion *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSecureConnectNextUpgradeToVersion

				if r.Products.SecureConnect.NextUpgrade.ToVersion != nil {
					id := r.Products.SecureConnect.NextUpgrade.ToVersion.ID.ValueString()
					requestNetworksUpdateNetworkFirmwareUpgradesProductsSecureConnectNextUpgradeToVersion = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSecureConnectNextUpgradeToVersion{
						ID: id,
					}
					//[debug] Is Array: False
				}
				requestNetworksUpdateNetworkFirmwareUpgradesProductsSecureConnectNextUpgrade = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSecureConnectNextUpgrade{
					Time:      time,
					ToVersion: requestNetworksUpdateNetworkFirmwareUpgradesProductsSecureConnectNextUpgradeToVersion,
				}
				//[debug] Is Array: False
			}
			participateInNextBetaRelease := func() *bool {
				if !r.Products.SecureConnect.ParticipateInNextBetaRelease.IsUnknown() && !r.Products.SecureConnect.ParticipateInNextBetaRelease.IsNull() {
					return r.Products.SecureConnect.ParticipateInNextBetaRelease.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkFirmwareUpgradesProductsSecureConnect = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSecureConnect{
				NextUpgrade:                  requestNetworksUpdateNetworkFirmwareUpgradesProductsSecureConnectNextUpgrade,
				ParticipateInNextBetaRelease: participateInNextBetaRelease,
			}
			//[debug] Is Array: False
		}
		var requestNetworksUpdateNetworkFirmwareUpgradesProductsSensor *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSensor

		if r.Products.Sensor != nil {
			var requestNetworksUpdateNetworkFirmwareUpgradesProductsSensorNextUpgrade *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSensorNextUpgrade

			if r.Products.Sensor.NextUpgrade != nil {
				time := r.Products.Sensor.NextUpgrade.Time.ValueString()
				var requestNetworksUpdateNetworkFirmwareUpgradesProductsSensorNextUpgradeToVersion *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSensorNextUpgradeToVersion

				if r.Products.Sensor.NextUpgrade.ToVersion != nil {
					id := r.Products.Sensor.NextUpgrade.ToVersion.ID.ValueString()
					requestNetworksUpdateNetworkFirmwareUpgradesProductsSensorNextUpgradeToVersion = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSensorNextUpgradeToVersion{
						ID: id,
					}
					//[debug] Is Array: False
				}
				requestNetworksUpdateNetworkFirmwareUpgradesProductsSensorNextUpgrade = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSensorNextUpgrade{
					Time:      time,
					ToVersion: requestNetworksUpdateNetworkFirmwareUpgradesProductsSensorNextUpgradeToVersion,
				}
				//[debug] Is Array: False
			}
			participateInNextBetaRelease := func() *bool {
				if !r.Products.Sensor.ParticipateInNextBetaRelease.IsUnknown() && !r.Products.Sensor.ParticipateInNextBetaRelease.IsNull() {
					return r.Products.Sensor.ParticipateInNextBetaRelease.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkFirmwareUpgradesProductsSensor = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSensor{
				NextUpgrade:                  requestNetworksUpdateNetworkFirmwareUpgradesProductsSensorNextUpgrade,
				ParticipateInNextBetaRelease: participateInNextBetaRelease,
			}
			//[debug] Is Array: False
		}
		var requestNetworksUpdateNetworkFirmwareUpgradesProductsSwitch *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitch

		if r.Products.Switch != nil {
			var requestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchNextUpgrade *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchNextUpgrade

			if r.Products.Switch.NextUpgrade != nil {
				time := r.Products.Switch.NextUpgrade.Time.ValueString()
				var requestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchNextUpgradeToVersion *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchNextUpgradeToVersion

				if r.Products.Switch.NextUpgrade.ToVersion != nil {
					id := r.Products.Switch.NextUpgrade.ToVersion.ID.ValueString()
					requestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchNextUpgradeToVersion = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchNextUpgradeToVersion{
						ID: id,
					}
					//[debug] Is Array: False
				}
				requestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchNextUpgrade = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchNextUpgrade{
					Time:      time,
					ToVersion: requestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchNextUpgradeToVersion,
				}
				//[debug] Is Array: False
			}
			participateInNextBetaRelease := func() *bool {
				if !r.Products.Switch.ParticipateInNextBetaRelease.IsUnknown() && !r.Products.Switch.ParticipateInNextBetaRelease.IsNull() {
					return r.Products.Switch.ParticipateInNextBetaRelease.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkFirmwareUpgradesProductsSwitch = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitch{
				NextUpgrade:                  requestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchNextUpgrade,
				ParticipateInNextBetaRelease: participateInNextBetaRelease,
			}
			//[debug] Is Array: False
		}
		var requestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalyst *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalyst

		if r.Products.SwitchCatalyst != nil {
			var requestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalystNextUpgrade *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalystNextUpgrade

			if r.Products.SwitchCatalyst.NextUpgrade != nil {
				time := r.Products.SwitchCatalyst.NextUpgrade.Time.ValueString()
				var requestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalystNextUpgradeToVersion *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalystNextUpgradeToVersion

				if r.Products.SwitchCatalyst.NextUpgrade.ToVersion != nil {
					id := r.Products.SwitchCatalyst.NextUpgrade.ToVersion.ID.ValueString()
					requestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalystNextUpgradeToVersion = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalystNextUpgradeToVersion{
						ID: id,
					}
					//[debug] Is Array: False
				}
				requestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalystNextUpgrade = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalystNextUpgrade{
					Time:      time,
					ToVersion: requestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalystNextUpgradeToVersion,
				}
				//[debug] Is Array: False
			}
			participateInNextBetaRelease := func() *bool {
				if !r.Products.SwitchCatalyst.ParticipateInNextBetaRelease.IsUnknown() && !r.Products.SwitchCatalyst.ParticipateInNextBetaRelease.IsNull() {
					return r.Products.SwitchCatalyst.ParticipateInNextBetaRelease.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalyst = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalyst{
				NextUpgrade:                  requestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalystNextUpgrade,
				ParticipateInNextBetaRelease: participateInNextBetaRelease,
			}
			//[debug] Is Array: False
		}
		var requestNetworksUpdateNetworkFirmwareUpgradesProductsWireless *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsWireless

		if r.Products.Wireless != nil {
			var requestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessNextUpgrade *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessNextUpgrade

			if r.Products.Wireless.NextUpgrade != nil {
				time := r.Products.Wireless.NextUpgrade.Time.ValueString()
				var requestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessNextUpgradeToVersion *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessNextUpgradeToVersion

				if r.Products.Wireless.NextUpgrade.ToVersion != nil {
					id := r.Products.Wireless.NextUpgrade.ToVersion.ID.ValueString()
					requestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessNextUpgradeToVersion = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessNextUpgradeToVersion{
						ID: id,
					}
					//[debug] Is Array: False
				}
				requestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessNextUpgrade = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessNextUpgrade{
					Time:      time,
					ToVersion: requestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessNextUpgradeToVersion,
				}
				//[debug] Is Array: False
			}
			participateInNextBetaRelease := func() *bool {
				if !r.Products.Wireless.ParticipateInNextBetaRelease.IsUnknown() && !r.Products.Wireless.ParticipateInNextBetaRelease.IsNull() {
					return r.Products.Wireless.ParticipateInNextBetaRelease.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkFirmwareUpgradesProductsWireless = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsWireless{
				NextUpgrade:                  requestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessNextUpgrade,
				ParticipateInNextBetaRelease: participateInNextBetaRelease,
			}
			//[debug] Is Array: False
		}
		var requestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessController *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessController

		if r.Products.WirelessController != nil {
			var requestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessControllerNextUpgrade *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessControllerNextUpgrade

			if r.Products.WirelessController.NextUpgrade != nil {
				time := r.Products.WirelessController.NextUpgrade.Time.ValueString()
				var requestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessControllerNextUpgradeToVersion *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessControllerNextUpgradeToVersion

				if r.Products.WirelessController.NextUpgrade.ToVersion != nil {
					id := r.Products.WirelessController.NextUpgrade.ToVersion.ID.ValueString()
					requestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessControllerNextUpgradeToVersion = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessControllerNextUpgradeToVersion{
						ID: id,
					}
					//[debug] Is Array: False
				}
				requestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessControllerNextUpgrade = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessControllerNextUpgrade{
					Time:      time,
					ToVersion: requestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessControllerNextUpgradeToVersion,
				}
				//[debug] Is Array: False
			}
			participateInNextBetaRelease := func() *bool {
				if !r.Products.WirelessController.ParticipateInNextBetaRelease.IsUnknown() && !r.Products.WirelessController.ParticipateInNextBetaRelease.IsNull() {
					return r.Products.WirelessController.ParticipateInNextBetaRelease.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessController = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessController{
				NextUpgrade:                  requestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessControllerNextUpgrade,
				ParticipateInNextBetaRelease: participateInNextBetaRelease,
			}
			//[debug] Is Array: False
		}
		requestNetworksUpdateNetworkFirmwareUpgradesProducts = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesProducts{
			Appliance:          requestNetworksUpdateNetworkFirmwareUpgradesProductsAppliance,
			Camera:             requestNetworksUpdateNetworkFirmwareUpgradesProductsCamera,
			CellularGateway:    requestNetworksUpdateNetworkFirmwareUpgradesProductsCellularGateway,
			SecureConnect:      requestNetworksUpdateNetworkFirmwareUpgradesProductsSecureConnect,
			Sensor:             requestNetworksUpdateNetworkFirmwareUpgradesProductsSensor,
			Switch:             requestNetworksUpdateNetworkFirmwareUpgradesProductsSwitch,
			SwitchCatalyst:     requestNetworksUpdateNetworkFirmwareUpgradesProductsSwitchCatalyst,
			Wireless:           requestNetworksUpdateNetworkFirmwareUpgradesProductsWireless,
			WirelessController: requestNetworksUpdateNetworkFirmwareUpgradesProductsWirelessController,
		}
		//[debug] Is Array: False
	}
	timezone := new(string)
	if !r.Timezone.IsUnknown() && !r.Timezone.IsNull() {
		*timezone = r.Timezone.ValueString()
	} else {
		timezone = &emptyString
	}
	var requestNetworksUpdateNetworkFirmwareUpgradesUpgradeWindow *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesUpgradeWindow

	if r.UpgradeWindow != nil {
		dayOfWeek := r.UpgradeWindow.DayOfWeek.ValueString()
		hourOfDay := r.UpgradeWindow.HourOfDay.ValueString()
		requestNetworksUpdateNetworkFirmwareUpgradesUpgradeWindow = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesUpgradeWindow{
			DayOfWeek: dayOfWeek,
			HourOfDay: hourOfDay,
		}
		//[debug] Is Array: False
	}
	out := merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgrades{
		Products:      requestNetworksUpdateNetworkFirmwareUpgradesProducts,
		Timezone:      *timezone,
		UpgradeWindow: requestNetworksUpdateNetworkFirmwareUpgradesUpgradeWindow,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkFirmwareUpgradesItemToBodyRs(state NetworksFirmwareUpgradesRs, response *merakigosdk.ResponseNetworksGetNetworkFirmwareUpgrades, is_read bool) NetworksFirmwareUpgradesRs {
	itemState := NetworksFirmwareUpgradesRs{
		Products: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsRs {
			if response.Products != nil {
				return &ResponseNetworksGetNetworkFirmwareUpgradesProductsRs{
					Appliance: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceRs {
						if response.Products.Appliance != nil {
							return &ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceRs{
								AvailableVersions: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceAvailableVersionsRs {
									if response.Products.Appliance.AvailableVersions != nil {
										result := make([]ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceAvailableVersionsRs, len(*response.Products.Appliance.AvailableVersions))
										for i, availableVersions := range *response.Products.Appliance.AvailableVersions {
											result[i] = ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceAvailableVersionsRs{
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
								CurrentVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceCurrentVersionRs {
									if response.Products.Appliance.CurrentVersion != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceCurrentVersionRs{
											Firmware:    types.StringValue(response.Products.Appliance.CurrentVersion.Firmware),
											ID:          types.StringValue(response.Products.Appliance.CurrentVersion.ID),
											ReleaseDate: types.StringValue(response.Products.Appliance.CurrentVersion.ReleaseDate),
											ReleaseType: types.StringValue(response.Products.Appliance.CurrentVersion.ReleaseType),
											ShortName:   types.StringValue(response.Products.Appliance.CurrentVersion.ShortName),
										}
									}
									return nil
								}(),
								LastUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeRs {
									if response.Products.Appliance.LastUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeRs{
											FromVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeFromVersionRs {
												if response.Products.Appliance.LastUpgrade.FromVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeFromVersionRs{
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
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeToVersionRs {
												if response.Products.Appliance.LastUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceLastUpgradeToVersionRs{
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
								NextUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgradeRs {
									if response.Products.Appliance.NextUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgradeRs{
											Time: types.StringValue(response.Products.Appliance.NextUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgradeToVersionRs {
												if response.Products.Appliance.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsApplianceNextUpgradeToVersionRs{
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
					Camera: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraRs {
						if response.Products.Camera != nil {
							return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraRs{
								AvailableVersions: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraAvailableVersionsRs {
									if response.Products.Camera.AvailableVersions != nil {
										result := make([]ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraAvailableVersionsRs, len(*response.Products.Camera.AvailableVersions))
										for i, availableVersions := range *response.Products.Camera.AvailableVersions {
											result[i] = ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraAvailableVersionsRs{
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
								CurrentVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraCurrentVersionRs {
									if response.Products.Camera.CurrentVersion != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraCurrentVersionRs{
											Firmware:    types.StringValue(response.Products.Camera.CurrentVersion.Firmware),
											ID:          types.StringValue(response.Products.Camera.CurrentVersion.ID),
											ReleaseDate: types.StringValue(response.Products.Camera.CurrentVersion.ReleaseDate),
											ReleaseType: types.StringValue(response.Products.Camera.CurrentVersion.ReleaseType),
											ShortName:   types.StringValue(response.Products.Camera.CurrentVersion.ShortName),
										}
									}
									return nil
								}(),
								LastUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeRs {
									if response.Products.Camera.LastUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeRs{
											FromVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeFromVersionRs {
												if response.Products.Camera.LastUpgrade.FromVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeFromVersionRs{
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
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeToVersionRs {
												if response.Products.Camera.LastUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraLastUpgradeToVersionRs{
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
								NextUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgradeRs {
									if response.Products.Camera.NextUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgradeRs{
											Time: types.StringValue(response.Products.Camera.NextUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgradeToVersionRs {
												if response.Products.Camera.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCameraNextUpgradeToVersionRs{
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
					CellularGateway: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayRs {
						if response.Products.CellularGateway != nil {
							return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayRs{
								AvailableVersions: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayAvailableVersionsRs {
									if response.Products.CellularGateway.AvailableVersions != nil {
										result := make([]ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayAvailableVersionsRs, len(*response.Products.CellularGateway.AvailableVersions))
										for i, availableVersions := range *response.Products.CellularGateway.AvailableVersions {
											result[i] = ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayAvailableVersionsRs{
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
								CurrentVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayCurrentVersionRs {
									if response.Products.CellularGateway.CurrentVersion != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayCurrentVersionRs{
											Firmware:    types.StringValue(response.Products.CellularGateway.CurrentVersion.Firmware),
											ID:          types.StringValue(response.Products.CellularGateway.CurrentVersion.ID),
											ReleaseDate: types.StringValue(response.Products.CellularGateway.CurrentVersion.ReleaseDate),
											ReleaseType: types.StringValue(response.Products.CellularGateway.CurrentVersion.ReleaseType),
											ShortName:   types.StringValue(response.Products.CellularGateway.CurrentVersion.ShortName),
										}
									}
									return nil
								}(),
								LastUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeRs {
									if response.Products.CellularGateway.LastUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeRs{
											FromVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeFromVersionRs {
												if response.Products.CellularGateway.LastUpgrade.FromVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeFromVersionRs{
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
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeToVersionRs {
												if response.Products.CellularGateway.LastUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayLastUpgradeToVersionRs{
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
								NextUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeRs {
									if response.Products.CellularGateway.NextUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeRs{
											Time: types.StringValue(response.Products.CellularGateway.NextUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeToVersionRs {
												if response.Products.CellularGateway.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsCellularGatewayNextUpgradeToVersionRs{
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
					SecureConnect: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectRs {
						if response.Products.SecureConnect != nil {
							return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectRs{
								AvailableVersions: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectAvailableVersionsRs {
									if response.Products.SecureConnect.AvailableVersions != nil {
										result := make([]ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectAvailableVersionsRs, len(*response.Products.SecureConnect.AvailableVersions))
										for i, availableVersions := range *response.Products.SecureConnect.AvailableVersions {
											result[i] = ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectAvailableVersionsRs{
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
								CurrentVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectCurrentVersionRs {
									if response.Products.SecureConnect.CurrentVersion != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectCurrentVersionRs{
											Firmware:    types.StringValue(response.Products.SecureConnect.CurrentVersion.Firmware),
											ID:          types.StringValue(response.Products.SecureConnect.CurrentVersion.ID),
											ReleaseDate: types.StringValue(response.Products.SecureConnect.CurrentVersion.ReleaseDate),
											ReleaseType: types.StringValue(response.Products.SecureConnect.CurrentVersion.ReleaseType),
											ShortName:   types.StringValue(response.Products.SecureConnect.CurrentVersion.ShortName),
										}
									}
									return nil
								}(),
								LastUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeRs {
									if response.Products.SecureConnect.LastUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeRs{
											FromVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeFromVersionRs {
												if response.Products.SecureConnect.LastUpgrade.FromVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeFromVersionRs{
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
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeToVersionRs {
												if response.Products.SecureConnect.LastUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectLastUpgradeToVersionRs{
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
								NextUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectNextUpgradeRs {
									if response.Products.SecureConnect.NextUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectNextUpgradeRs{
											Time: types.StringValue(response.Products.SecureConnect.NextUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectNextUpgradeToVersionRs {
												if response.Products.SecureConnect.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSecureConnectNextUpgradeToVersionRs{
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
					Sensor: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorRs {
						if response.Products.Sensor != nil {
							return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorRs{
								AvailableVersions: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorAvailableVersionsRs {
									if response.Products.Sensor.AvailableVersions != nil {
										result := make([]ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorAvailableVersionsRs, len(*response.Products.Sensor.AvailableVersions))
										for i, availableVersions := range *response.Products.Sensor.AvailableVersions {
											result[i] = ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorAvailableVersionsRs{
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
								CurrentVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorCurrentVersionRs {
									if response.Products.Sensor.CurrentVersion != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorCurrentVersionRs{
											Firmware:    types.StringValue(response.Products.Sensor.CurrentVersion.Firmware),
											ID:          types.StringValue(response.Products.Sensor.CurrentVersion.ID),
											ReleaseDate: types.StringValue(response.Products.Sensor.CurrentVersion.ReleaseDate),
											ReleaseType: types.StringValue(response.Products.Sensor.CurrentVersion.ReleaseType),
											ShortName:   types.StringValue(response.Products.Sensor.CurrentVersion.ShortName),
										}
									}
									return nil
								}(),
								LastUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeRs {
									if response.Products.Sensor.LastUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeRs{
											FromVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeFromVersionRs {
												if response.Products.Sensor.LastUpgrade.FromVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeFromVersionRs{
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
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeToVersionRs {
												if response.Products.Sensor.LastUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorLastUpgradeToVersionRs{
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
								NextUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgradeRs {
									if response.Products.Sensor.NextUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgradeRs{
											Time: types.StringValue(response.Products.Sensor.NextUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgradeToVersionRs {
												if response.Products.Sensor.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSensorNextUpgradeToVersionRs{
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
					Switch: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchRs {
						if response.Products.Switch != nil {
							return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchRs{
								AvailableVersions: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchAvailableVersionsRs {
									if response.Products.Switch.AvailableVersions != nil {
										result := make([]ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchAvailableVersionsRs, len(*response.Products.Switch.AvailableVersions))
										for i, availableVersions := range *response.Products.Switch.AvailableVersions {
											result[i] = ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchAvailableVersionsRs{
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
								CurrentVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchCurrentVersionRs {
									if response.Products.Switch.CurrentVersion != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchCurrentVersionRs{
											Firmware:    types.StringValue(response.Products.Switch.CurrentVersion.Firmware),
											ID:          types.StringValue(response.Products.Switch.CurrentVersion.ID),
											ReleaseDate: types.StringValue(response.Products.Switch.CurrentVersion.ReleaseDate),
											ReleaseType: types.StringValue(response.Products.Switch.CurrentVersion.ReleaseType),
											ShortName:   types.StringValue(response.Products.Switch.CurrentVersion.ShortName),
										}
									}
									return nil
								}(),
								LastUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeRs {
									if response.Products.Switch.LastUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeRs{
											FromVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeFromVersionRs {
												if response.Products.Switch.LastUpgrade.FromVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeFromVersionRs{
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
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeToVersionRs {
												if response.Products.Switch.LastUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchLastUpgradeToVersionRs{
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
								NextUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgradeRs {
									if response.Products.Switch.NextUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgradeRs{
											Time: types.StringValue(response.Products.Switch.NextUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgradeToVersionRs {
												if response.Products.Switch.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsSwitchNextUpgradeToVersionRs{
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
					Wireless: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessRs {
						if response.Products.Wireless != nil {
							return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessRs{
								AvailableVersions: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessAvailableVersionsRs {
									if response.Products.Wireless.AvailableVersions != nil {
										result := make([]ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessAvailableVersionsRs, len(*response.Products.Wireless.AvailableVersions))
										for i, availableVersions := range *response.Products.Wireless.AvailableVersions {
											result[i] = ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessAvailableVersionsRs{
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
								CurrentVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessCurrentVersionRs {
									if response.Products.Wireless.CurrentVersion != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessCurrentVersionRs{
											Firmware:    types.StringValue(response.Products.Wireless.CurrentVersion.Firmware),
											ID:          types.StringValue(response.Products.Wireless.CurrentVersion.ID),
											ReleaseDate: types.StringValue(response.Products.Wireless.CurrentVersion.ReleaseDate),
											ReleaseType: types.StringValue(response.Products.Wireless.CurrentVersion.ReleaseType),
											ShortName:   types.StringValue(response.Products.Wireless.CurrentVersion.ShortName),
										}
									}
									return nil
								}(),
								LastUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeRs {
									if response.Products.Wireless.LastUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeRs{
											FromVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeFromVersionRs {
												if response.Products.Wireless.LastUpgrade.FromVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeFromVersionRs{
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
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeToVersionRs {
												if response.Products.Wireless.LastUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessLastUpgradeToVersionRs{
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
								NextUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgradeRs {
									if response.Products.Wireless.NextUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgradeRs{
											Time: types.StringValue(response.Products.Wireless.NextUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgradeToVersionRs {
												if response.Products.Wireless.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessNextUpgradeToVersionRs{
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
					WirelessController: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerRs {
						if response.Products.WirelessController != nil {
							return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerRs{
								AvailableVersions: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerAvailableVersionsRs {
									if response.Products.WirelessController.AvailableVersions != nil {
										result := make([]ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerAvailableVersionsRs, len(*response.Products.WirelessController.AvailableVersions))
										for i, availableVersions := range *response.Products.WirelessController.AvailableVersions {
											result[i] = ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerAvailableVersionsRs{
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
								CurrentVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerCurrentVersionRs {
									if response.Products.WirelessController.CurrentVersion != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerCurrentVersionRs{
											Firmware:    types.StringValue(response.Products.WirelessController.CurrentVersion.Firmware),
											ID:          types.StringValue(response.Products.WirelessController.CurrentVersion.ID),
											ReleaseDate: types.StringValue(response.Products.WirelessController.CurrentVersion.ReleaseDate),
											ReleaseType: types.StringValue(response.Products.WirelessController.CurrentVersion.ReleaseType),
											ShortName:   types.StringValue(response.Products.WirelessController.CurrentVersion.ShortName),
										}
									}
									return nil
								}(),
								LastUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeRs {
									if response.Products.WirelessController.LastUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeRs{
											FromVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeFromVersionRs {
												if response.Products.WirelessController.LastUpgrade.FromVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeFromVersionRs{
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
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeToVersionRs {
												if response.Products.WirelessController.LastUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerLastUpgradeToVersionRs{
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
								NextUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerNextUpgradeRs {
									if response.Products.WirelessController.NextUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerNextUpgradeRs{
											Time: types.StringValue(response.Products.WirelessController.NextUpgrade.Time),
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerNextUpgradeToVersionRs {
												if response.Products.WirelessController.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesProductsWirelessControllerNextUpgradeToVersionRs{
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
		UpgradeWindow: func() *ResponseNetworksGetNetworkFirmwareUpgradesUpgradeWindowRs {
			if response.UpgradeWindow != nil {
				return &ResponseNetworksGetNetworkFirmwareUpgradesUpgradeWindowRs{
					DayOfWeek: types.StringValue(response.UpgradeWindow.DayOfWeek),
					HourOfDay: types.StringValue(response.UpgradeWindow.HourOfDay),
				}
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksFirmwareUpgradesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksFirmwareUpgradesRs)
}
