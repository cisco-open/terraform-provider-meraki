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
	_ datasource.DataSource              = &OrganizationsWirelessDevicesWirelessControllersByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessDevicesWirelessControllersByDeviceDataSource{}
)

func NewOrganizationsWirelessDevicesWirelessControllersByDeviceDataSource() datasource.DataSource {
	return &OrganizationsWirelessDevicesWirelessControllersByDeviceDataSource{}
}

type OrganizationsWirelessDevicesWirelessControllersByDeviceDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessDevicesWirelessControllersByDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessDevicesWirelessControllersByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_devices_wireless_controllers_by_device"
}

func (d *OrganizationsWirelessDevicesWirelessControllersByDeviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"controller_serials": schema.ListAttribute{
				MarkdownDescription: `controllerSerials query parameter. Optional parameter to filter access points by its wireless LAN controller cloud ID. This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter access points by network ID. This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 100.`,
				Optional:            true,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Optional parameter to filter access points by its cloud ID. This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"items": schema.ListNestedAttribute{
						MarkdownDescription: `List of Catalyst access points information`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"controller": schema.SingleNestedAttribute{
									MarkdownDescription: `Associated wireless controller`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"serial": schema.StringAttribute{
											MarkdownDescription: `Associated wireless controller cloud ID`,
											Computed:            true,
										},
									},
								},
								"country_code": schema.StringAttribute{
									MarkdownDescription: `Country code (2 characters)`,
									Computed:            true,
								},
								"details": schema.SetNestedAttribute{
									MarkdownDescription: `Catalyst access point details`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"name": schema.StringAttribute{
												MarkdownDescription: `Item name`,
												Computed:            true,
											},
											"value": schema.StringAttribute{
												MarkdownDescription: `Item value`,
												Computed:            true,
											},
										},
									},
								},
								"joined_at": schema.StringAttribute{
									MarkdownDescription: `The time when AP joins wireless controller`,
									Computed:            true,
								},
								"mode": schema.StringAttribute{
									MarkdownDescription: `AP mode (local, flex, etc.)`,
									Computed:            true,
								},
								"model": schema.StringAttribute{
									MarkdownDescription: `AP model`,
									Computed:            true,
								},
								"network": schema.SingleNestedAttribute{
									MarkdownDescription: `Catalyst access point network`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `Catalyst access point network ID`,
											Computed:            true,
										},
									},
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `AP cloud ID`,
									Computed:            true,
								},
								"tags": schema.SetNestedAttribute{
									MarkdownDescription: `The tags of the catalyst access point`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"policy": schema.StringAttribute{
												MarkdownDescription: `Policy tag`,
												Computed:            true,
											},
											"rf": schema.StringAttribute{
												MarkdownDescription: `RF tag`,
												Computed:            true,
											},
											"site": schema.StringAttribute{
												MarkdownDescription: `Site tag`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
					},
					"meta": schema.SingleNestedAttribute{
						MarkdownDescription: `Metadata relevant to the paginated dataset`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"counts": schema.SingleNestedAttribute{
								MarkdownDescription: `Counts relating to the paginated dataset`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"items": schema.SingleNestedAttribute{
										MarkdownDescription: `Counts relating to the paginated items`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"remaining": schema.Int64Attribute{
												MarkdownDescription: `The number of items in the dataset that are available on subsequent pages`,
												Computed:            true,
											},
											"total": schema.Int64Attribute{
												MarkdownDescription: `The total number of items in the dataset`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsWirelessDevicesWirelessControllersByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessDevicesWirelessControllersByDevice OrganizationsWirelessDevicesWirelessControllersByDevice
	diags := req.Config.Get(ctx, &organizationsWirelessDevicesWirelessControllersByDevice)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessDevicesWirelessControllersByDevice")
		vvOrganizationID := organizationsWirelessDevicesWirelessControllersByDevice.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessDevicesWirelessControllersByDeviceQueryParams{}

		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsWirelessDevicesWirelessControllersByDevice.NetworkIDs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessDevicesWirelessControllersByDevice.Serials)
		queryParams1.ControllerSerials = elementsToStrings(ctx, organizationsWirelessDevicesWirelessControllersByDevice.ControllerSerials)
		queryParams1.PerPage = int(organizationsWirelessDevicesWirelessControllersByDevice.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessDevicesWirelessControllersByDevice.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessDevicesWirelessControllersByDevice.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetOrganizationWirelessDevicesWirelessControllersByDevice(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessDevicesWirelessControllersByDevice",
				err.Error(),
			)
			return
		}

		organizationsWirelessDevicesWirelessControllersByDevice = ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemToBody(organizationsWirelessDevicesWirelessControllersByDevice, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessDevicesWirelessControllersByDevice)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessDevicesWirelessControllersByDevice struct {
	OrganizationID    types.String                                                               `tfsdk:"organization_id"`
	NetworkIDs        types.List                                                                 `tfsdk:"network_ids"`
	Serials           types.List                                                                 `tfsdk:"serials"`
	ControllerSerials types.List                                                                 `tfsdk:"controller_serials"`
	PerPage           types.Int64                                                                `tfsdk:"per_page"`
	StartingAfter     types.String                                                               `tfsdk:"starting_after"`
	EndingBefore      types.String                                                               `tfsdk:"ending_before"`
	Item              *ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDevice `tfsdk:"item"`
}

type ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDevice struct {
	Items *[]ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItems `tfsdk:"items"`
	Meta  *ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceMeta    `tfsdk:"meta"`
}

type ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItems struct {
	Controller  *ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsController `tfsdk:"controller"`
	CountryCode types.String                                                                              `tfsdk:"country_code"`
	Details     *[]ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsDetails  `tfsdk:"details"`
	JoinedAt    types.String                                                                              `tfsdk:"joined_at"`
	Mode        types.String                                                                              `tfsdk:"mode"`
	Model       types.String                                                                              `tfsdk:"model"`
	Network     *ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsNetwork    `tfsdk:"network"`
	Serial      types.String                                                                              `tfsdk:"serial"`
	Tags        *[]ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsTags     `tfsdk:"tags"`
}

type ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsController struct {
	Serial types.String `tfsdk:"serial"`
}

type ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsDetails struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsNetwork struct {
	ID types.String `tfsdk:"id"`
}

type ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsTags struct {
	Policy types.String `tfsdk:"policy"`
	Rf     types.String `tfsdk:"rf"`
	Site   types.String `tfsdk:"site"`
}

type ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceMeta struct {
	Counts *ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceMetaCounts `tfsdk:"counts"`
}

type ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceMetaCounts struct {
	Items *ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceMetaCountsItems `tfsdk:"items"`
}

type ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemToBody(state OrganizationsWirelessDevicesWirelessControllersByDevice, response *merakigosdk.ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDevice) OrganizationsWirelessDevicesWirelessControllersByDevice {
	itemState := ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDevice{
		Items: func() *[]ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItems {
			if response.Items != nil {
				result := make([]ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItems{
						Controller: func() *ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsController {
							if items.Controller != nil {
								return &ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsController{
									Serial: types.StringValue(items.Controller.Serial),
								}
							}
							return nil
						}(),
						CountryCode: types.StringValue(items.CountryCode),
						Details: func() *[]ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsDetails {
							if items.Details != nil {
								result := make([]ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsDetails, len(*items.Details))
								for i, details := range *items.Details {
									result[i] = ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsDetails{
										Name:  types.StringValue(details.Name),
										Value: types.StringValue(details.Value),
									}
								}
								return &result
							}
							return nil
						}(),
						JoinedAt: types.StringValue(items.JoinedAt),
						Mode:     types.StringValue(items.Mode),
						Model:    types.StringValue(items.Model),
						Network: func() *ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsNetwork {
							if items.Network != nil {
								return &ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsNetwork{
									ID: types.StringValue(items.Network.ID),
								}
							}
							return nil
						}(),
						Serial: types.StringValue(items.Serial),
						Tags: func() *[]ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsTags {
							if items.Tags != nil {
								result := make([]ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsTags, len(*items.Tags))
								for i, tags := range *items.Tags {
									result[i] = ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceItemsTags{
										Policy: types.StringValue(tags.Policy),
										Rf:     types.StringValue(tags.Rf),
										Site:   types.StringValue(tags.Site),
									}
								}
								return &result
							}
							return nil
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Meta: func() *ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceMeta {
			if response.Meta != nil {
				return &ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceMeta{
					Counts: func() *ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceMetaCounts{
								Items: func() *ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseWirelessGetOrganizationWirelessDevicesWirelessControllersByDeviceMetaCountsItems{
											Remaining: func() types.Int64 {
												if response.Meta.Counts.Items.Remaining != nil {
													return types.Int64Value(int64(*response.Meta.Counts.Items.Remaining))
												}
												return types.Int64{}
											}(),
											Total: func() types.Int64 {
												if response.Meta.Counts.Items.Total != nil {
													return types.Int64Value(int64(*response.Meta.Counts.Items.Total))
												}
												return types.Int64{}
											}(),
										}
									}
									return nil
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
