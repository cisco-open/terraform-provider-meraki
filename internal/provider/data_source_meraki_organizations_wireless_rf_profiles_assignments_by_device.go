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
	_ datasource.DataSource              = &OrganizationsWirelessRfProfilesAssignmentsByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsWirelessRfProfilesAssignmentsByDeviceDataSource{}
)

func NewOrganizationsWirelessRfProfilesAssignmentsByDeviceDataSource() datasource.DataSource {
	return &OrganizationsWirelessRfProfilesAssignmentsByDeviceDataSource{}
}

type OrganizationsWirelessRfProfilesAssignmentsByDeviceDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsWirelessRfProfilesAssignmentsByDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsWirelessRfProfilesAssignmentsByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_rf_profiles_assignments_by_device"
}

func (d *OrganizationsWirelessRfProfilesAssignmentsByDeviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"mac": schema.StringAttribute{
				MarkdownDescription: `mac query parameter. Optional parameter to filter RF profiles by device MAC address. All returned devices will have a MAC address that contains the search term or is an exact match.`,
				Optional:            true,
			},
			"macs": schema.ListAttribute{
				MarkdownDescription: `macs query parameter. Optional parameter to filter RF profiles by one or more device MAC addresses. All returned devices will have a MAC address that is an exact match.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"model": schema.StringAttribute{
				MarkdownDescription: `model query parameter. Optional parameter to filter RF profiles by device model. All returned devices will have a model that contains the search term or is an exact match.`,
				Optional:            true,
			},
			"models": schema.ListAttribute{
				MarkdownDescription: `models query parameter. Optional parameter to filter RF profiles by one or more device models. All returned devices will have a model that is an exact match.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `name query parameter. Optional parameter to filter RF profiles by device name. All returned devices will have a name that contains the search term or is an exact match.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter devices by network.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.`,
				Optional:            true,
			},
			"product_types": schema.ListAttribute{
				MarkdownDescription: `productTypes query parameter. Optional parameter to filter devices by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, and secureConnect.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial query parameter. Optional parameter to filter RF profiles by device serial number. All returned devices will have a serial number that contains the search term or is an exact match.`,
				Optional:            true,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Optional parameter to filter RF profiles by one or more device serial numbers. All returned devices will have a serial number that is an exact match.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseWirelessGetOrganizationWirelessRfProfilesAssignmentsByDevice`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"items": schema.ListNestedAttribute{
							MarkdownDescription: `The top-level propery containing all status data.`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"model": schema.StringAttribute{
										MarkdownDescription: `Model number of the device.`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Name of the device.`,
										Computed:            true,
									},
									"network": schema.SingleNestedAttribute{
										MarkdownDescription: `Information regarding the network the device belongs to.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"id": schema.StringAttribute{
												MarkdownDescription: `The network ID.`,
												Computed:            true,
											},
										},
									},
									"rf_profile": schema.SingleNestedAttribute{
										MarkdownDescription: `Information regarding the RF Profile of the device.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"id": schema.StringAttribute{
												MarkdownDescription: `The ID of the RF Profile the device belongs to.`,
												Computed:            true,
											},
											"is_indoor_default": schema.BoolAttribute{
												MarkdownDescription: `Status to show if this profile is default indoor profile.`,
												Computed:            true,
											},
											"is_outdoor_default": schema.BoolAttribute{
												MarkdownDescription: `Status to show if this profile is default outdoor profile.`,
												Computed:            true,
											},
											"name": schema.StringAttribute{
												MarkdownDescription: `The name of the RF Profile the device belongs to.`,
												Computed:            true,
											},
										},
									},
									"serial": schema.StringAttribute{
										MarkdownDescription: `Unique serial number for device.`,
										Computed:            true,
									},
								},
							},
						},
						"meta": schema.SingleNestedAttribute{
							MarkdownDescription: `Other metadata related to this result set.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"counts": schema.SingleNestedAttribute{
									MarkdownDescription: `Count metadata related to this result set.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"items": schema.SingleNestedAttribute{
											MarkdownDescription: `The count metadata.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"remaining": schema.Int64Attribute{
													MarkdownDescription: `The number of serials remaining based on current pagination location within the dataset.`,
													Computed:            true,
												},
												"total": schema.Int64Attribute{
													MarkdownDescription: `The total number of serials.`,
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
		},
	}
}

func (d *OrganizationsWirelessRfProfilesAssignmentsByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsWirelessRfProfilesAssignmentsByDevice OrganizationsWirelessRfProfilesAssignmentsByDevice
	diags := req.Config.Get(ctx, &organizationsWirelessRfProfilesAssignmentsByDevice)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationWirelessRfProfilesAssignmentsByDevice")
		vvOrganizationID := organizationsWirelessRfProfilesAssignmentsByDevice.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationWirelessRfProfilesAssignmentsByDeviceQueryParams{}

		queryParams1.PerPage = int(organizationsWirelessRfProfilesAssignmentsByDevice.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsWirelessRfProfilesAssignmentsByDevice.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsWirelessRfProfilesAssignmentsByDevice.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsWirelessRfProfilesAssignmentsByDevice.NetworkIDs)
		queryParams1.ProductTypes = elementsToStrings(ctx, organizationsWirelessRfProfilesAssignmentsByDevice.ProductTypes)
		queryParams1.Name = organizationsWirelessRfProfilesAssignmentsByDevice.Name.ValueString()
		queryParams1.Mac = organizationsWirelessRfProfilesAssignmentsByDevice.Mac.ValueString()
		queryParams1.Serial = organizationsWirelessRfProfilesAssignmentsByDevice.Serial.ValueString()
		queryParams1.Model = organizationsWirelessRfProfilesAssignmentsByDevice.Model.ValueString()
		queryParams1.Macs = elementsToStrings(ctx, organizationsWirelessRfProfilesAssignmentsByDevice.Macs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsWirelessRfProfilesAssignmentsByDevice.Serials)
		queryParams1.Models = elementsToStrings(ctx, organizationsWirelessRfProfilesAssignmentsByDevice.Models)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetOrganizationWirelessRfProfilesAssignmentsByDevice(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationWirelessRfProfilesAssignmentsByDevice",
				err.Error(),
			)
			return
		}

		organizationsWirelessRfProfilesAssignmentsByDevice = ResponseWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceItemsToBody(organizationsWirelessRfProfilesAssignmentsByDevice, response1)
		diags = resp.State.Set(ctx, &organizationsWirelessRfProfilesAssignmentsByDevice)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsWirelessRfProfilesAssignmentsByDevice struct {
	OrganizationID types.String                                                                `tfsdk:"organization_id"`
	PerPage        types.Int64                                                                 `tfsdk:"per_page"`
	StartingAfter  types.String                                                                `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                `tfsdk:"ending_before"`
	NetworkIDs     types.List                                                                  `tfsdk:"network_ids"`
	ProductTypes   types.List                                                                  `tfsdk:"product_types"`
	Name           types.String                                                                `tfsdk:"name"`
	Mac            types.String                                                                `tfsdk:"mac"`
	Serial         types.String                                                                `tfsdk:"serial"`
	Model          types.String                                                                `tfsdk:"model"`
	Macs           types.List                                                                  `tfsdk:"macs"`
	Serials        types.List                                                                  `tfsdk:"serials"`
	Models         types.List                                                                  `tfsdk:"models"`
	Items          *[]ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDevice `tfsdk:"items"`
}

type ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDevice struct {
	Items *[]ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceItems `tfsdk:"items"`
	Meta  *ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceMeta    `tfsdk:"meta"`
}

type ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceItems struct {
	Model     types.String                                                                            `tfsdk:"model"`
	Name      types.String                                                                            `tfsdk:"name"`
	Network   *ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceItemsNetwork   `tfsdk:"network"`
	RfProfile *ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceItemsRfProfile `tfsdk:"rf_profile"`
	Serial    types.String                                                                            `tfsdk:"serial"`
}

type ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceItemsNetwork struct {
	ID types.String `tfsdk:"id"`
}

type ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceItemsRfProfile struct {
	ID               types.String `tfsdk:"id"`
	IsIndoorDefault  types.Bool   `tfsdk:"is_indoor_default"`
	IsOutdoorDefault types.Bool   `tfsdk:"is_outdoor_default"`
	Name             types.String `tfsdk:"name"`
}

type ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceMeta struct {
	Counts *ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceMetaCounts `tfsdk:"counts"`
}

type ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceMetaCounts struct {
	Items *ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceMetaCountsItems `tfsdk:"items"`
}

type ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceItemsToBody(state OrganizationsWirelessRfProfilesAssignmentsByDevice, response *merakigosdk.ResponseWirelessGetOrganizationWirelessRfProfilesAssignmentsByDevice) OrganizationsWirelessRfProfilesAssignmentsByDevice {
	var items []ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDevice
	for _, item := range *response {
		itemState := ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDevice{
			Items: func() *[]ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceItems {
				if item.Items != nil {
					result := make([]ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceItems, len(*item.Items))
					for i, items := range *item.Items {
						result[i] = ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceItems{
							Model: types.StringValue(items.Model),
							Name:  types.StringValue(items.Name),
							Network: func() *ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceItemsNetwork {
								if items.Network != nil {
									return &ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceItemsNetwork{
										ID: types.StringValue(items.Network.ID),
									}
								}
								return nil
							}(),
							RfProfile: func() *ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceItemsRfProfile {
								if items.RfProfile != nil {
									return &ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceItemsRfProfile{
										ID: types.StringValue(items.RfProfile.ID),
										IsIndoorDefault: func() types.Bool {
											if items.RfProfile.IsIndoorDefault != nil {
												return types.BoolValue(*items.RfProfile.IsIndoorDefault)
											}
											return types.Bool{}
										}(),
										IsOutdoorDefault: func() types.Bool {
											if items.RfProfile.IsOutdoorDefault != nil {
												return types.BoolValue(*items.RfProfile.IsOutdoorDefault)
											}
											return types.Bool{}
										}(),
										Name: types.StringValue(items.RfProfile.Name),
									}
								}
								return nil
							}(),
							Serial: types.StringValue(items.Serial),
						}
					}
					return &result
				}
				return nil
			}(),
			Meta: func() *ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceMeta {
				if item.Meta != nil {
					return &ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceMeta{
						Counts: func() *ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceMetaCounts {
							if item.Meta.Counts != nil {
								return &ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceMetaCounts{
									Items: func() *ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceMetaCountsItems {
										if item.Meta.Counts.Items != nil {
											return &ResponseItemWirelessGetOrganizationWirelessRfProfilesAssignmentsByDeviceMetaCountsItems{
												Remaining: func() types.Int64 {
													if item.Meta.Counts.Items.Remaining != nil {
														return types.Int64Value(int64(*item.Meta.Counts.Items.Remaining))
													}
													return types.Int64{}
												}(),
												Total: func() types.Int64 {
													if item.Meta.Counts.Items.Total != nil {
														return types.Int64Value(int64(*item.Meta.Counts.Items.Total))
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
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
