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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsCameraBoundariesLinesByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsCameraBoundariesLinesByDeviceDataSource{}
)

func NewOrganizationsCameraBoundariesLinesByDeviceDataSource() datasource.DataSource {
	return &OrganizationsCameraBoundariesLinesByDeviceDataSource{}
}

type OrganizationsCameraBoundariesLinesByDeviceDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsCameraBoundariesLinesByDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsCameraBoundariesLinesByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_camera_boundaries_lines_by_device"
}

func (d *OrganizationsCameraBoundariesLinesByDeviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. A list of serial numbers. The returned cameras will be filtered to only include these serials.`,
				Optional:            true,
				ElementType:         types.StringType,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseCameraGetOrganizationCameraBoundariesLinesByDevice`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"boundaries": schema.SingleNestedAttribute{
							MarkdownDescription: `Configured line boundaries of the camera`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"direction_vertex": schema.SingleNestedAttribute{
									MarkdownDescription: `The line boundary crossing direction vertex`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"x": schema.Float64Attribute{
											MarkdownDescription: `The vertex x coordinate`,
											Computed:            true,
										},
										"y": schema.Float64Attribute{
											MarkdownDescription: `The vertex y coordinate`,
											Computed:            true,
										},
									},
								},
								"id": schema.StringAttribute{
									MarkdownDescription: `The line boundary id`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `The line boundary name`,
									Computed:            true,
								},
								"type": schema.StringAttribute{
									MarkdownDescription: `The line boundary type`,
									Computed:            true,
								},
								"vertices": schema.SetNestedAttribute{
									MarkdownDescription: `The line boundary vertices`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"x": schema.Float64Attribute{
												MarkdownDescription: `The vertex x coordinate`,
												Computed:            true,
											},
											"y": schema.Float64Attribute{
												MarkdownDescription: `The vertex y coordinate`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: `The network id of the camera`,
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `The serial number of the camera`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsCameraBoundariesLinesByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsCameraBoundariesLinesByDevice OrganizationsCameraBoundariesLinesByDevice
	diags := req.Config.Get(ctx, &organizationsCameraBoundariesLinesByDevice)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationCameraBoundariesLinesByDevice")
		vvOrganizationID := organizationsCameraBoundariesLinesByDevice.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationCameraBoundariesLinesByDeviceQueryParams{}

		queryParams1.Serials = elementsToStrings(ctx, organizationsCameraBoundariesLinesByDevice.Serials)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Camera.GetOrganizationCameraBoundariesLinesByDevice(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCameraBoundariesLinesByDevice",
				err.Error(),
			)
			return
		}

		organizationsCameraBoundariesLinesByDevice = ResponseCameraGetOrganizationCameraBoundariesLinesByDeviceItemsToBody(organizationsCameraBoundariesLinesByDevice, response1)
		diags = resp.State.Set(ctx, &organizationsCameraBoundariesLinesByDevice)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsCameraBoundariesLinesByDevice struct {
	OrganizationID types.String                                                      `tfsdk:"organization_id"`
	Serials        types.List                                                        `tfsdk:"serials"`
	Items          *[]ResponseItemCameraGetOrganizationCameraBoundariesLinesByDevice `tfsdk:"items"`
}

type ResponseItemCameraGetOrganizationCameraBoundariesLinesByDevice struct {
	Boundaries *ResponseItemCameraGetOrganizationCameraBoundariesLinesByDeviceBoundaries `tfsdk:"boundaries"`
	NetworkID  types.String                                                              `tfsdk:"network_id"`
	Serial     types.String                                                              `tfsdk:"serial"`
}

type ResponseItemCameraGetOrganizationCameraBoundariesLinesByDeviceBoundaries struct {
	DirectionVertex *ResponseItemCameraGetOrganizationCameraBoundariesLinesByDeviceBoundariesDirectionVertex `tfsdk:"direction_vertex"`
	ID              types.String                                                                             `tfsdk:"id"`
	Name            types.String                                                                             `tfsdk:"name"`
	Type            types.String                                                                             `tfsdk:"type"`
	Vertices        *[]ResponseItemCameraGetOrganizationCameraBoundariesLinesByDeviceBoundariesVertices      `tfsdk:"vertices"`
}

type ResponseItemCameraGetOrganizationCameraBoundariesLinesByDeviceBoundariesDirectionVertex struct {
	X types.Float64 `tfsdk:"x"`
	Y types.Float64 `tfsdk:"y"`
}

type ResponseItemCameraGetOrganizationCameraBoundariesLinesByDeviceBoundariesVertices struct {
	X types.Float64 `tfsdk:"x"`
	Y types.Float64 `tfsdk:"y"`
}

// ToBody
func ResponseCameraGetOrganizationCameraBoundariesLinesByDeviceItemsToBody(state OrganizationsCameraBoundariesLinesByDevice, response *merakigosdk.ResponseCameraGetOrganizationCameraBoundariesLinesByDevice) OrganizationsCameraBoundariesLinesByDevice {
	var items []ResponseItemCameraGetOrganizationCameraBoundariesLinesByDevice
	for _, item := range *response {
		itemState := ResponseItemCameraGetOrganizationCameraBoundariesLinesByDevice{
			Boundaries: func() *ResponseItemCameraGetOrganizationCameraBoundariesLinesByDeviceBoundaries {
				if item.Boundaries != nil {
					return &ResponseItemCameraGetOrganizationCameraBoundariesLinesByDeviceBoundaries{
						DirectionVertex: func() *ResponseItemCameraGetOrganizationCameraBoundariesLinesByDeviceBoundariesDirectionVertex {
							if item.Boundaries.DirectionVertex != nil {
								return &ResponseItemCameraGetOrganizationCameraBoundariesLinesByDeviceBoundariesDirectionVertex{
									X: func() types.Float64 {
										if item.Boundaries.DirectionVertex.X != nil {
											return types.Float64Value(float64(*item.Boundaries.DirectionVertex.X))
										}
										return types.Float64{}
									}(),
									Y: func() types.Float64 {
										if item.Boundaries.DirectionVertex.Y != nil {
											return types.Float64Value(float64(*item.Boundaries.DirectionVertex.Y))
										}
										return types.Float64{}
									}(),
								}
							}
							return nil
						}(),
						ID: func() types.String {
							if item.Boundaries.ID != "" {
								return types.StringValue(item.Boundaries.ID)
							}
							return types.String{}
						}(),
						Name: func() types.String {
							if item.Boundaries.Name != "" {
								return types.StringValue(item.Boundaries.Name)
							}
							return types.String{}
						}(),
						Type: func() types.String {
							if item.Boundaries.Type != "" {
								return types.StringValue(item.Boundaries.Type)
							}
							return types.String{}
						}(),
						Vertices: func() *[]ResponseItemCameraGetOrganizationCameraBoundariesLinesByDeviceBoundariesVertices {
							if item.Boundaries.Vertices != nil {
								result := make([]ResponseItemCameraGetOrganizationCameraBoundariesLinesByDeviceBoundariesVertices, len(*item.Boundaries.Vertices))
								for i, vertices := range *item.Boundaries.Vertices {
									result[i] = ResponseItemCameraGetOrganizationCameraBoundariesLinesByDeviceBoundariesVertices{
										X: func() types.Float64 {
											if vertices.X != nil {
												return types.Float64Value(float64(*vertices.X))
											}
											return types.Float64{}
										}(),
										Y: func() types.Float64 {
											if vertices.Y != nil {
												return types.Float64Value(float64(*vertices.Y))
											}
											return types.Float64{}
										}(),
									}
								}
								return &result
							}
							return nil
						}(),
					}
				}
				return nil
			}(),
			NetworkID: func() types.String {
				if item.NetworkID != "" {
					return types.StringValue(item.NetworkID)
				}
				return types.String{}
			}(),
			Serial: func() types.String {
				if item.Serial != "" {
					return types.StringValue(item.Serial)
				}
				return types.String{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
