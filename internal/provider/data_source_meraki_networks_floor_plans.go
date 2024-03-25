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
	_ datasource.DataSource              = &NetworksFloorPlansDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksFloorPlansDataSource{}
)

func NewNetworksFloorPlansDataSource() datasource.DataSource {
	return &NetworksFloorPlansDataSource{}
}

type NetworksFloorPlansDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksFloorPlansDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksFloorPlansDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_floor_plans"
}

func (d *NetworksFloorPlansDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"floor_plan_id": schema.StringAttribute{
				MarkdownDescription: `floorPlanId path parameter. Floor plan ID`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"bottom_left_corner": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"lat": schema.Float64Attribute{
								Computed: true,
							},
							"lng": schema.Float64Attribute{
								Computed: true,
							},
						},
					},
					"bottom_right_corner": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"lat": schema.Float64Attribute{
								Computed: true,
							},
							"lng": schema.Float64Attribute{
								Computed: true,
							},
						},
					},
					"center": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"lat": schema.Float64Attribute{
								Computed: true,
							},
							"lng": schema.Float64Attribute{
								Computed: true,
							},
						},
					},
					"devices": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"address": schema.StringAttribute{
									Computed: true,
								},
								"beacon_id_params": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"major": schema.Int64Attribute{
											Computed: true,
										},
										"minor": schema.Int64Attribute{
											Computed: true,
										},
										"uuid": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"firmware": schema.StringAttribute{
									Computed: true,
								},
								"floor_plan_id": schema.StringAttribute{
									Computed: true,
								},
								"lan_ip": schema.StringAttribute{
									Computed: true,
								},
								"lat": schema.Float64Attribute{
									Computed: true,
								},
								"lng": schema.Float64Attribute{
									Computed: true,
								},
								"mac": schema.StringAttribute{
									Computed: true,
								},
								"model": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"network_id": schema.StringAttribute{
									Computed: true,
								},
								"notes": schema.StringAttribute{
									Computed: true,
								},
								"serial": schema.StringAttribute{
									Computed: true,
								},
								"tags": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
							},
						},
					},
					"floor_plan_id": schema.StringAttribute{
						Computed: true,
					},
					"height": schema.Float64Attribute{
						Computed: true,
					},
					"image_extension": schema.StringAttribute{
						Computed: true,
					},
					"image_md5": schema.StringAttribute{
						Computed: true,
					},
					"image_url": schema.StringAttribute{
						Computed: true,
					},
					"image_url_expires_at": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"top_left_corner": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"lat": schema.Float64Attribute{
								Computed: true,
							},
							"lng": schema.Float64Attribute{
								Computed: true,
							},
						},
					},
					"top_right_corner": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"lat": schema.Float64Attribute{
								Computed: true,
							},
							"lng": schema.Float64Attribute{
								Computed: true,
							},
						},
					},
					"width": schema.Int64Attribute{
						Computed: true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseNetworksGetNetworkFloorPlans`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"bottom_left_corner": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"lat": schema.Float64Attribute{
									Computed: true,
								},
								"lng": schema.Float64Attribute{
									Computed: true,
								},
							},
						},
						"bottom_right_corner": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"lat": schema.Float64Attribute{
									Computed: true,
								},
								"lng": schema.Float64Attribute{
									Computed: true,
								},
							},
						},
						"center": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"lat": schema.Float64Attribute{
									Computed: true,
								},
								"lng": schema.Float64Attribute{
									Computed: true,
								},
							},
						},
						"devices": schema.SetNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"address": schema.StringAttribute{
										Computed: true,
									},
									"beacon_id_params": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{

											"major": schema.Int64Attribute{
												Computed: true,
											},
											"minor": schema.Int64Attribute{
												Computed: true,
											},
											"uuid": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"firmware": schema.StringAttribute{
										Computed: true,
									},
									"floor_plan_id": schema.StringAttribute{
										Computed: true,
									},
									"lan_ip": schema.StringAttribute{
										Computed: true,
									},
									"lat": schema.Float64Attribute{
										Computed: true,
									},
									"lng": schema.Float64Attribute{
										Computed: true,
									},
									"mac": schema.StringAttribute{
										Computed: true,
									},
									"model": schema.StringAttribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
									"network_id": schema.StringAttribute{
										Computed: true,
									},
									"notes": schema.StringAttribute{
										Computed: true,
									},
									"serial": schema.StringAttribute{
										Computed: true,
									},
									"tags": schema.ListAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
								},
							},
						},
						"floor_plan_id": schema.StringAttribute{
							Computed: true,
						},
						"height": schema.Float64Attribute{
							Computed: true,
						},
						"image_extension": schema.StringAttribute{
							Computed: true,
						},
						"image_md5": schema.StringAttribute{
							Computed: true,
						},
						"image_url": schema.StringAttribute{
							Computed: true,
						},
						"image_url_expires_at": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"top_left_corner": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"lat": schema.Float64Attribute{
									Computed: true,
								},
								"lng": schema.Float64Attribute{
									Computed: true,
								},
							},
						},
						"top_right_corner": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"lat": schema.Float64Attribute{
									Computed: true,
								},
								"lng": schema.Float64Attribute{
									Computed: true,
								},
							},
						},
						"width": schema.Int64Attribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksFloorPlansDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksFloorPlans NetworksFloorPlans
	diags := req.Config.Get(ctx, &networksFloorPlans)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksFloorPlans.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksFloorPlans.NetworkID.IsNull(), !networksFloorPlans.FloorPlanID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkFloorPlans")
		vvNetworkID := networksFloorPlans.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Networks.GetNetworkFloorPlans(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFloorPlans",
				err.Error(),
			)
			return
		}

		networksFloorPlans = ResponseNetworksGetNetworkFloorPlansItemsToBody(networksFloorPlans, response1)
		diags = resp.State.Set(ctx, &networksFloorPlans)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkFloorPlan")
		vvNetworkID := networksFloorPlans.NetworkID.ValueString()
		vvFloorPlanID := networksFloorPlans.FloorPlanID.ValueString()

		response2, restyResp2, err := d.client.Networks.GetNetworkFloorPlan(vvNetworkID, vvFloorPlanID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFloorPlan",
				err.Error(),
			)
			return
		}

		networksFloorPlans = ResponseNetworksGetNetworkFloorPlanItemToBody(networksFloorPlans, response2)
		diags = resp.State.Set(ctx, &networksFloorPlans)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksFloorPlans struct {
	NetworkID   types.String                                `tfsdk:"network_id"`
	FloorPlanID types.String                                `tfsdk:"floor_plan_id"`
	Items       *[]ResponseItemNetworksGetNetworkFloorPlans `tfsdk:"items"`
	Item        *ResponseNetworksGetNetworkFloorPlan        `tfsdk:"item"`
}

type ResponseItemNetworksGetNetworkFloorPlans struct {
	BottomLeftCorner  *ResponseItemNetworksGetNetworkFloorPlansBottomLeftCorner  `tfsdk:"bottom_left_corner"`
	BottomRightCorner *ResponseItemNetworksGetNetworkFloorPlansBottomRightCorner `tfsdk:"bottom_right_corner"`
	Center            *ResponseItemNetworksGetNetworkFloorPlansCenter            `tfsdk:"center"`
	Devices           *[]ResponseItemNetworksGetNetworkFloorPlansDevices         `tfsdk:"devices"`
	FloorPlanID       types.String                                               `tfsdk:"floor_plan_id"`
	Height            types.Float64                                              `tfsdk:"height"`
	ImageExtension    types.String                                               `tfsdk:"image_extension"`
	ImageMd5          types.String                                               `tfsdk:"image_md5"`
	ImageURL          types.String                                               `tfsdk:"image_url"`
	ImageURLExpiresAt types.String                                               `tfsdk:"image_url_expires_at"`
	Name              types.String                                               `tfsdk:"name"`
	TopLeftCorner     *ResponseItemNetworksGetNetworkFloorPlansTopLeftCorner     `tfsdk:"top_left_corner"`
	TopRightCorner    *ResponseItemNetworksGetNetworkFloorPlansTopRightCorner    `tfsdk:"top_right_corner"`
	Width             types.Int64                                                `tfsdk:"width"`
}

type ResponseItemNetworksGetNetworkFloorPlansBottomLeftCorner struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseItemNetworksGetNetworkFloorPlansBottomRightCorner struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseItemNetworksGetNetworkFloorPlansCenter struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseItemNetworksGetNetworkFloorPlansDevices struct {
	Address        types.String                                                   `tfsdk:"address"`
	BeaconIDParams *ResponseItemNetworksGetNetworkFloorPlansDevicesBeaconIdParams `tfsdk:"beacon_id_params"`
	Firmware       types.String                                                   `tfsdk:"firmware"`
	FloorPlanID    types.String                                                   `tfsdk:"floor_plan_id"`
	LanIP          types.String                                                   `tfsdk:"lan_ip"`
	Lat            types.Float64                                                  `tfsdk:"lat"`
	Lng            types.Float64                                                  `tfsdk:"lng"`
	Mac            types.String                                                   `tfsdk:"mac"`
	Model          types.String                                                   `tfsdk:"model"`
	Name           types.String                                                   `tfsdk:"name"`
	NetworkID      types.String                                                   `tfsdk:"network_id"`
	Notes          types.String                                                   `tfsdk:"notes"`
	Serial         types.String                                                   `tfsdk:"serial"`
	Tags           types.List                                                     `tfsdk:"tags"`
}

type ResponseItemNetworksGetNetworkFloorPlansDevicesBeaconIdParams struct {
	Major types.Int64  `tfsdk:"major"`
	Minor types.Int64  `tfsdk:"minor"`
	UUID  types.String `tfsdk:"uuid"`
}

type ResponseItemNetworksGetNetworkFloorPlansTopLeftCorner struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseItemNetworksGetNetworkFloorPlansTopRightCorner struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseNetworksGetNetworkFloorPlan struct {
	BottomLeftCorner  *ResponseNetworksGetNetworkFloorPlanBottomLeftCorner  `tfsdk:"bottom_left_corner"`
	BottomRightCorner *ResponseNetworksGetNetworkFloorPlanBottomRightCorner `tfsdk:"bottom_right_corner"`
	Center            *ResponseNetworksGetNetworkFloorPlanCenter            `tfsdk:"center"`
	Devices           *[]ResponseNetworksGetNetworkFloorPlanDevices         `tfsdk:"devices"`
	FloorPlanID       types.String                                          `tfsdk:"floor_plan_id"`
	Height            types.Float64                                         `tfsdk:"height"`
	ImageExtension    types.String                                          `tfsdk:"image_extension"`
	ImageMd5          types.String                                          `tfsdk:"image_md5"`
	ImageURL          types.String                                          `tfsdk:"image_url"`
	ImageURLExpiresAt types.String                                          `tfsdk:"image_url_expires_at"`
	Name              types.String                                          `tfsdk:"name"`
	TopLeftCorner     *ResponseNetworksGetNetworkFloorPlanTopLeftCorner     `tfsdk:"top_left_corner"`
	TopRightCorner    *ResponseNetworksGetNetworkFloorPlanTopRightCorner    `tfsdk:"top_right_corner"`
	Width             types.Int64                                           `tfsdk:"width"`
}

type ResponseNetworksGetNetworkFloorPlanBottomLeftCorner struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseNetworksGetNetworkFloorPlanBottomRightCorner struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseNetworksGetNetworkFloorPlanCenter struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseNetworksGetNetworkFloorPlanDevices struct {
	Address        types.String                                              `tfsdk:"address"`
	BeaconIDParams *ResponseNetworksGetNetworkFloorPlanDevicesBeaconIdParams `tfsdk:"beacon_id_params"`
	Firmware       types.String                                              `tfsdk:"firmware"`
	FloorPlanID    types.String                                              `tfsdk:"floor_plan_id"`
	LanIP          types.String                                              `tfsdk:"lan_ip"`
	Lat            types.Float64                                             `tfsdk:"lat"`
	Lng            types.Float64                                             `tfsdk:"lng"`
	Mac            types.String                                              `tfsdk:"mac"`
	Model          types.String                                              `tfsdk:"model"`
	Name           types.String                                              `tfsdk:"name"`
	NetworkID      types.String                                              `tfsdk:"network_id"`
	Notes          types.String                                              `tfsdk:"notes"`
	Serial         types.String                                              `tfsdk:"serial"`
	Tags           types.List                                                `tfsdk:"tags"`
}

type ResponseNetworksGetNetworkFloorPlanDevicesBeaconIdParams struct {
	Major types.Int64  `tfsdk:"major"`
	Minor types.Int64  `tfsdk:"minor"`
	UUID  types.String `tfsdk:"uuid"`
}

type ResponseNetworksGetNetworkFloorPlanTopLeftCorner struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseNetworksGetNetworkFloorPlanTopRightCorner struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

// ToBody
func ResponseNetworksGetNetworkFloorPlansItemsToBody(state NetworksFloorPlans, response *merakigosdk.ResponseNetworksGetNetworkFloorPlans) NetworksFloorPlans {
	var items []ResponseItemNetworksGetNetworkFloorPlans
	for _, item := range *response {
		itemState := ResponseItemNetworksGetNetworkFloorPlans{
			BottomLeftCorner: func() *ResponseItemNetworksGetNetworkFloorPlansBottomLeftCorner {
				if item.BottomLeftCorner != nil {
					return &ResponseItemNetworksGetNetworkFloorPlansBottomLeftCorner{
						Lat: func() types.Float64 {
							if item.BottomLeftCorner.Lat != nil {
								return types.Float64Value(float64(*item.BottomLeftCorner.Lat))
							}
							return types.Float64{}
						}(),
						Lng: func() types.Float64 {
							if item.BottomLeftCorner.Lng != nil {
								return types.Float64Value(float64(*item.BottomLeftCorner.Lng))
							}
							return types.Float64{}
						}(),
					}
				}
				return &ResponseItemNetworksGetNetworkFloorPlansBottomLeftCorner{}
			}(),
			BottomRightCorner: func() *ResponseItemNetworksGetNetworkFloorPlansBottomRightCorner {
				if item.BottomRightCorner != nil {
					return &ResponseItemNetworksGetNetworkFloorPlansBottomRightCorner{
						Lat: func() types.Float64 {
							if item.BottomRightCorner.Lat != nil {
								return types.Float64Value(float64(*item.BottomRightCorner.Lat))
							}
							return types.Float64{}
						}(),
						Lng: func() types.Float64 {
							if item.BottomRightCorner.Lng != nil {
								return types.Float64Value(float64(*item.BottomRightCorner.Lng))
							}
							return types.Float64{}
						}(),
					}
				}
				return &ResponseItemNetworksGetNetworkFloorPlansBottomRightCorner{}
			}(),
			Center: func() *ResponseItemNetworksGetNetworkFloorPlansCenter {
				if item.Center != nil {
					return &ResponseItemNetworksGetNetworkFloorPlansCenter{
						Lat: func() types.Float64 {
							if item.Center.Lat != nil {
								return types.Float64Value(float64(*item.Center.Lat))
							}
							return types.Float64{}
						}(),
						Lng: func() types.Float64 {
							if item.Center.Lng != nil {
								return types.Float64Value(float64(*item.Center.Lng))
							}
							return types.Float64{}
						}(),
					}
				}
				return &ResponseItemNetworksGetNetworkFloorPlansCenter{}
			}(),
			Devices: func() *[]ResponseItemNetworksGetNetworkFloorPlansDevices {
				if item.Devices != nil {
					result := make([]ResponseItemNetworksGetNetworkFloorPlansDevices, len(*item.Devices))
					for i, devices := range *item.Devices {
						result[i] = ResponseItemNetworksGetNetworkFloorPlansDevices{
							Address: types.StringValue(devices.Address),
							BeaconIDParams: func() *ResponseItemNetworksGetNetworkFloorPlansDevicesBeaconIdParams {
								if devices.BeaconIDParams != nil {
									return &ResponseItemNetworksGetNetworkFloorPlansDevicesBeaconIdParams{
										Major: func() types.Int64 {
											if devices.BeaconIDParams.Major != nil {
												return types.Int64Value(int64(*devices.BeaconIDParams.Major))
											}
											return types.Int64{}
										}(),
										Minor: func() types.Int64 {
											if devices.BeaconIDParams.Minor != nil {
												return types.Int64Value(int64(*devices.BeaconIDParams.Minor))
											}
											return types.Int64{}
										}(),
										UUID: types.StringValue(devices.BeaconIDParams.UUID),
									}
								}
								return &ResponseItemNetworksGetNetworkFloorPlansDevicesBeaconIdParams{}
							}(),
							Firmware:    types.StringValue(devices.Firmware),
							FloorPlanID: types.StringValue(devices.FloorPlanID),
							LanIP:       types.StringValue(devices.LanIP),
							Lat: func() types.Float64 {
								if devices.Lat != nil {
									return types.Float64Value(float64(*devices.Lat))
								}
								return types.Float64{}
							}(),
							Lng: func() types.Float64 {
								if devices.Lng != nil {
									return types.Float64Value(float64(*devices.Lng))
								}
								return types.Float64{}
							}(),
							Mac:       types.StringValue(devices.Mac),
							Model:     types.StringValue(devices.Model),
							Name:      types.StringValue(devices.Name),
							NetworkID: types.StringValue(devices.NetworkID),
							Notes:     types.StringValue(devices.Notes),
							Serial:    types.StringValue(devices.Serial),
							Tags:      StringSliceToList(devices.Tags),
						}
					}
					return &result
				}
				return &[]ResponseItemNetworksGetNetworkFloorPlansDevices{}
			}(),
			FloorPlanID: types.StringValue(item.FloorPlanID),
			Height: func() types.Float64 {
				if item.Height != nil {
					return types.Float64Value(float64(*item.Height))
				}
				return types.Float64{}
			}(),
			ImageExtension:    types.StringValue(item.ImageExtension),
			ImageMd5:          types.StringValue(item.ImageMd5),
			ImageURL:          types.StringValue(item.ImageURL),
			ImageURLExpiresAt: types.StringValue(item.ImageURLExpiresAt),
			Name:              types.StringValue(item.Name),
			TopLeftCorner: func() *ResponseItemNetworksGetNetworkFloorPlansTopLeftCorner {
				if item.TopLeftCorner != nil {
					return &ResponseItemNetworksGetNetworkFloorPlansTopLeftCorner{
						Lat: func() types.Float64 {
							if item.TopLeftCorner.Lat != nil {
								return types.Float64Value(float64(*item.TopLeftCorner.Lat))
							}
							return types.Float64{}
						}(),
						Lng: func() types.Float64 {
							if item.TopLeftCorner.Lng != nil {
								return types.Float64Value(float64(*item.TopLeftCorner.Lng))
							}
							return types.Float64{}
						}(),
					}
				}
				return &ResponseItemNetworksGetNetworkFloorPlansTopLeftCorner{}
			}(),
			TopRightCorner: func() *ResponseItemNetworksGetNetworkFloorPlansTopRightCorner {
				if item.TopRightCorner != nil {
					return &ResponseItemNetworksGetNetworkFloorPlansTopRightCorner{
						Lat: func() types.Float64 {
							if item.TopRightCorner.Lat != nil {
								return types.Float64Value(float64(*item.TopRightCorner.Lat))
							}
							return types.Float64{}
						}(),
						Lng: func() types.Float64 {
							if item.TopRightCorner.Lng != nil {
								return types.Float64Value(float64(*item.TopRightCorner.Lng))
							}
							return types.Float64{}
						}(),
					}
				}
				return &ResponseItemNetworksGetNetworkFloorPlansTopRightCorner{}
			}(),
			Width: func() types.Int64 {
				if item.Width != nil {
					return types.Int64Value(int64(*item.Width))
				}
				return types.Int64{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseNetworksGetNetworkFloorPlanItemToBody(state NetworksFloorPlans, response *merakigosdk.ResponseNetworksGetNetworkFloorPlan) NetworksFloorPlans {
	itemState := ResponseNetworksGetNetworkFloorPlan{
		BottomLeftCorner: func() *ResponseNetworksGetNetworkFloorPlanBottomLeftCorner {
			if response.BottomLeftCorner != nil {
				return &ResponseNetworksGetNetworkFloorPlanBottomLeftCorner{
					Lat: func() types.Float64 {
						if response.BottomLeftCorner.Lat != nil {
							return types.Float64Value(float64(*response.BottomLeftCorner.Lat))
						}
						return types.Float64{}
					}(),
					Lng: func() types.Float64 {
						if response.BottomLeftCorner.Lng != nil {
							return types.Float64Value(float64(*response.BottomLeftCorner.Lng))
						}
						return types.Float64{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkFloorPlanBottomLeftCorner{}
		}(),
		BottomRightCorner: func() *ResponseNetworksGetNetworkFloorPlanBottomRightCorner {
			if response.BottomRightCorner != nil {
				return &ResponseNetworksGetNetworkFloorPlanBottomRightCorner{
					Lat: func() types.Float64 {
						if response.BottomRightCorner.Lat != nil {
							return types.Float64Value(float64(*response.BottomRightCorner.Lat))
						}
						return types.Float64{}
					}(),
					Lng: func() types.Float64 {
						if response.BottomRightCorner.Lng != nil {
							return types.Float64Value(float64(*response.BottomRightCorner.Lng))
						}
						return types.Float64{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkFloorPlanBottomRightCorner{}
		}(),
		Center: func() *ResponseNetworksGetNetworkFloorPlanCenter {
			if response.Center != nil {
				return &ResponseNetworksGetNetworkFloorPlanCenter{
					Lat: func() types.Float64 {
						if response.Center.Lat != nil {
							return types.Float64Value(float64(*response.Center.Lat))
						}
						return types.Float64{}
					}(),
					Lng: func() types.Float64 {
						if response.Center.Lng != nil {
							return types.Float64Value(float64(*response.Center.Lng))
						}
						return types.Float64{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkFloorPlanCenter{}
		}(),
		Devices: func() *[]ResponseNetworksGetNetworkFloorPlanDevices {
			if response.Devices != nil {
				result := make([]ResponseNetworksGetNetworkFloorPlanDevices, len(*response.Devices))
				for i, devices := range *response.Devices {
					result[i] = ResponseNetworksGetNetworkFloorPlanDevices{
						Address: types.StringValue(devices.Address),
						BeaconIDParams: func() *ResponseNetworksGetNetworkFloorPlanDevicesBeaconIdParams {
							if devices.BeaconIDParams != nil {
								return &ResponseNetworksGetNetworkFloorPlanDevicesBeaconIdParams{
									Major: func() types.Int64 {
										if devices.BeaconIDParams.Major != nil {
											return types.Int64Value(int64(*devices.BeaconIDParams.Major))
										}
										return types.Int64{}
									}(),
									Minor: func() types.Int64 {
										if devices.BeaconIDParams.Minor != nil {
											return types.Int64Value(int64(*devices.BeaconIDParams.Minor))
										}
										return types.Int64{}
									}(),
									UUID: types.StringValue(devices.BeaconIDParams.UUID),
								}
							}
							return &ResponseNetworksGetNetworkFloorPlanDevicesBeaconIdParams{}
						}(),
						Firmware:    types.StringValue(devices.Firmware),
						FloorPlanID: types.StringValue(devices.FloorPlanID),
						LanIP:       types.StringValue(devices.LanIP),
						Lat: func() types.Float64 {
							if devices.Lat != nil {
								return types.Float64Value(float64(*devices.Lat))
							}
							return types.Float64{}
						}(),
						Lng: func() types.Float64 {
							if devices.Lng != nil {
								return types.Float64Value(float64(*devices.Lng))
							}
							return types.Float64{}
						}(),
						Mac:       types.StringValue(devices.Mac),
						Model:     types.StringValue(devices.Model),
						Name:      types.StringValue(devices.Name),
						NetworkID: types.StringValue(devices.NetworkID),
						Notes:     types.StringValue(devices.Notes),
						Serial:    types.StringValue(devices.Serial),
						Tags:      StringSliceToList(devices.Tags),
					}
				}
				return &result
			}
			return &[]ResponseNetworksGetNetworkFloorPlanDevices{}
		}(),
		FloorPlanID: types.StringValue(response.FloorPlanID),
		Height: func() types.Float64 {
			if response.Height != nil {
				return types.Float64Value(float64(*response.Height))
			}
			return types.Float64{}
		}(),
		ImageExtension:    types.StringValue(response.ImageExtension),
		ImageMd5:          types.StringValue(response.ImageMd5),
		ImageURL:          types.StringValue(response.ImageURL),
		ImageURLExpiresAt: types.StringValue(response.ImageURLExpiresAt),
		Name:              types.StringValue(response.Name),
		TopLeftCorner: func() *ResponseNetworksGetNetworkFloorPlanTopLeftCorner {
			if response.TopLeftCorner != nil {
				return &ResponseNetworksGetNetworkFloorPlanTopLeftCorner{
					Lat: func() types.Float64 {
						if response.TopLeftCorner.Lat != nil {
							return types.Float64Value(float64(*response.TopLeftCorner.Lat))
						}
						return types.Float64{}
					}(),
					Lng: func() types.Float64 {
						if response.TopLeftCorner.Lng != nil {
							return types.Float64Value(float64(*response.TopLeftCorner.Lng))
						}
						return types.Float64{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkFloorPlanTopLeftCorner{}
		}(),
		TopRightCorner: func() *ResponseNetworksGetNetworkFloorPlanTopRightCorner {
			if response.TopRightCorner != nil {
				return &ResponseNetworksGetNetworkFloorPlanTopRightCorner{
					Lat: func() types.Float64 {
						if response.TopRightCorner.Lat != nil {
							return types.Float64Value(float64(*response.TopRightCorner.Lat))
						}
						return types.Float64{}
					}(),
					Lng: func() types.Float64 {
						if response.TopRightCorner.Lng != nil {
							return types.Float64Value(float64(*response.TopRightCorner.Lng))
						}
						return types.Float64{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkFloorPlanTopRightCorner{}
		}(),
		Width: func() types.Int64 {
			if response.Width != nil {
				return types.Int64Value(int64(*response.Width))
			}
			return types.Int64{}
		}(),
	}
	state.Item = &itemState
	return state
}
