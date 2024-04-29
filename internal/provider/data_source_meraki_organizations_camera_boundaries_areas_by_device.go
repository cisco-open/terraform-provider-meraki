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
	_ datasource.DataSource              = &OrganizationsCameraBoundariesAreasByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsCameraBoundariesAreasByDeviceDataSource{}
)

func NewOrganizationsCameraBoundariesAreasByDeviceDataSource() datasource.DataSource {
	return &OrganizationsCameraBoundariesAreasByDeviceDataSource{}
}

type OrganizationsCameraBoundariesAreasByDeviceDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsCameraBoundariesAreasByDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsCameraBoundariesAreasByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_camera_boundaries_areas_by_device"
}

func (d *OrganizationsCameraBoundariesAreasByDeviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				MarkdownDescription: `Array of ResponseCameraGetOrganizationCameraBoundariesAreasByDevice`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"boundaries": schema.SingleNestedAttribute{
							MarkdownDescription: `Configured area boundaries of the camera`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `The area boundary id`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `The area boundary name`,
									Computed:            true,
								},
								"type": schema.StringAttribute{
									MarkdownDescription: `The area boundary type`,
									Computed:            true,
								},
								"vertices": schema.SetNestedAttribute{
									MarkdownDescription: `The area boundary vertices`,
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

func (d *OrganizationsCameraBoundariesAreasByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsCameraBoundariesAreasByDevice OrganizationsCameraBoundariesAreasByDevice
	diags := req.Config.Get(ctx, &organizationsCameraBoundariesAreasByDevice)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationCameraBoundariesAreasByDevice")
		vvOrganizationID := organizationsCameraBoundariesAreasByDevice.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationCameraBoundariesAreasByDeviceQueryParams{}

		queryParams1.Serials = elementsToStrings(ctx, organizationsCameraBoundariesAreasByDevice.Serials)

		response1, restyResp1, err := d.client.Camera.GetOrganizationCameraBoundariesAreasByDevice(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCameraBoundariesAreasByDevice",
				err.Error(),
			)
			return
		}

		organizationsCameraBoundariesAreasByDevice = ResponseCameraGetOrganizationCameraBoundariesAreasByDeviceItemsToBody(organizationsCameraBoundariesAreasByDevice, response1)
		diags = resp.State.Set(ctx, &organizationsCameraBoundariesAreasByDevice)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsCameraBoundariesAreasByDevice struct {
	OrganizationID types.String                                                      `tfsdk:"organization_id"`
	Serials        types.List                                                        `tfsdk:"serials"`
	Items          *[]ResponseItemCameraGetOrganizationCameraBoundariesAreasByDevice `tfsdk:"items"`
}

type ResponseItemCameraGetOrganizationCameraBoundariesAreasByDevice struct {
	Boundaries *ResponseItemCameraGetOrganizationCameraBoundariesAreasByDeviceBoundaries `tfsdk:"boundaries"`
	NetworkID  types.String                                                              `tfsdk:"network_id"`
	Serial     types.String                                                              `tfsdk:"serial"`
}

type ResponseItemCameraGetOrganizationCameraBoundariesAreasByDeviceBoundaries struct {
	ID       types.String                                                                        `tfsdk:"id"`
	Name     types.String                                                                        `tfsdk:"name"`
	Type     types.String                                                                        `tfsdk:"type"`
	Vertices *[]ResponseItemCameraGetOrganizationCameraBoundariesAreasByDeviceBoundariesVertices `tfsdk:"vertices"`
}

type ResponseItemCameraGetOrganizationCameraBoundariesAreasByDeviceBoundariesVertices struct {
	X types.Float64 `tfsdk:"x"`
	Y types.Float64 `tfsdk:"y"`
}

// ToBody
func ResponseCameraGetOrganizationCameraBoundariesAreasByDeviceItemsToBody(state OrganizationsCameraBoundariesAreasByDevice, response *merakigosdk.ResponseCameraGetOrganizationCameraBoundariesAreasByDevice) OrganizationsCameraBoundariesAreasByDevice {
	var items []ResponseItemCameraGetOrganizationCameraBoundariesAreasByDevice
	for _, item := range *response {
		itemState := ResponseItemCameraGetOrganizationCameraBoundariesAreasByDevice{
			Boundaries: func() *ResponseItemCameraGetOrganizationCameraBoundariesAreasByDeviceBoundaries {
				if item.Boundaries != nil {
					return &ResponseItemCameraGetOrganizationCameraBoundariesAreasByDeviceBoundaries{
						ID:   types.StringValue(item.Boundaries.ID),
						Name: types.StringValue(item.Boundaries.Name),
						Type: types.StringValue(item.Boundaries.Type),
						Vertices: func() *[]ResponseItemCameraGetOrganizationCameraBoundariesAreasByDeviceBoundariesVertices {
							if item.Boundaries.Vertices != nil {
								result := make([]ResponseItemCameraGetOrganizationCameraBoundariesAreasByDeviceBoundariesVertices, len(*item.Boundaries.Vertices))
								for i, vertices := range *item.Boundaries.Vertices {
									result[i] = ResponseItemCameraGetOrganizationCameraBoundariesAreasByDeviceBoundariesVertices{
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
							return &[]ResponseItemCameraGetOrganizationCameraBoundariesAreasByDeviceBoundariesVertices{}
						}(),
					}
				}
				return &ResponseItemCameraGetOrganizationCameraBoundariesAreasByDeviceBoundaries{}
			}(),
			NetworkID: types.StringValue(item.NetworkID),
			Serial:    types.StringValue(item.Serial),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
