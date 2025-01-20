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
	_ datasource.DataSource              = &OrganizationsDevicesPowerModulesStatusesByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsDevicesPowerModulesStatusesByDeviceDataSource{}
)

func NewOrganizationsDevicesPowerModulesStatusesByDeviceDataSource() datasource.DataSource {
	return &OrganizationsDevicesPowerModulesStatusesByDeviceDataSource{}
}

type OrganizationsDevicesPowerModulesStatusesByDeviceDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsDevicesPowerModulesStatusesByDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsDevicesPowerModulesStatusesByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_devices_power_modules_statuses_by_device"
}

func (d *OrganizationsDevicesPowerModulesStatusesByDeviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter device availabilities by network ID. This filter uses multiple exact matches.`,
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
				MarkdownDescription: `productTypes query parameter. Optional parameter to filter device availabilities by device product types. This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"tags": schema.ListAttribute{
				MarkdownDescription: `tags query parameter. An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"tags_filter_type": schema.StringAttribute{
				MarkdownDescription: `tagsFilterType query parameter. An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationDevicesPowerModulesStatusesByDevice`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"mac": schema.StringAttribute{
							MarkdownDescription: `The device MAC address.`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The device name.`,
							Computed:            true,
						},
						"network": schema.SingleNestedAttribute{
							MarkdownDescription: `Network info.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `ID for the network that the device is associated with.`,
									Computed:            true,
								},
							},
						},
						"product_type": schema.StringAttribute{
							MarkdownDescription: `Device product type.`,
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `The device serial number.`,
							Computed:            true,
						},
						"slots": schema.SetNestedAttribute{
							MarkdownDescription: `Information for the device's AC power supplies.`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"model": schema.StringAttribute{
										MarkdownDescription: `The power supply unit model.`,
										Computed:            true,
									},
									"number": schema.Int64Attribute{
										MarkdownDescription: `Which slot the AC power supply occupies. Possible values are: 0, 1, 2.`,
										Computed:            true,
									},
									"serial": schema.StringAttribute{
										MarkdownDescription: `The power supply unit serial number.`,
										Computed:            true,
									},
									"status": schema.StringAttribute{
										MarkdownDescription: `Status of the power supply unit. Possible values are: connected, not connected, powering.`,
										Computed:            true,
									},
								},
							},
						},
						"tags": schema.ListAttribute{
							MarkdownDescription: `List of custom tags for the device.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsDevicesPowerModulesStatusesByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsDevicesPowerModulesStatusesByDevice OrganizationsDevicesPowerModulesStatusesByDevice
	diags := req.Config.Get(ctx, &organizationsDevicesPowerModulesStatusesByDevice)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationDevicesPowerModulesStatusesByDevice")
		vvOrganizationID := organizationsDevicesPowerModulesStatusesByDevice.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationDevicesPowerModulesStatusesByDeviceQueryParams{}

		queryParams1.PerPage = int(organizationsDevicesPowerModulesStatusesByDevice.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsDevicesPowerModulesStatusesByDevice.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsDevicesPowerModulesStatusesByDevice.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsDevicesPowerModulesStatusesByDevice.NetworkIDs)
		queryParams1.ProductTypes = elementsToStrings(ctx, organizationsDevicesPowerModulesStatusesByDevice.ProductTypes)
		queryParams1.Serials = elementsToStrings(ctx, organizationsDevicesPowerModulesStatusesByDevice.Serials)
		queryParams1.Tags = elementsToStrings(ctx, organizationsDevicesPowerModulesStatusesByDevice.Tags)
		queryParams1.TagsFilterType = organizationsDevicesPowerModulesStatusesByDevice.TagsFilterType.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationDevicesPowerModulesStatusesByDevice(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationDevicesPowerModulesStatusesByDevice",
				err.Error(),
			)
			return
		}

		organizationsDevicesPowerModulesStatusesByDevice = ResponseOrganizationsGetOrganizationDevicesPowerModulesStatusesByDeviceItemsToBody(organizationsDevicesPowerModulesStatusesByDevice, response1)
		diags = resp.State.Set(ctx, &organizationsDevicesPowerModulesStatusesByDevice)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsDevicesPowerModulesStatusesByDevice struct {
	OrganizationID types.String                                                                   `tfsdk:"organization_id"`
	PerPage        types.Int64                                                                    `tfsdk:"per_page"`
	StartingAfter  types.String                                                                   `tfsdk:"starting_after"`
	EndingBefore   types.String                                                                   `tfsdk:"ending_before"`
	NetworkIDs     types.List                                                                     `tfsdk:"network_ids"`
	ProductTypes   types.List                                                                     `tfsdk:"product_types"`
	Serials        types.List                                                                     `tfsdk:"serials"`
	Tags           types.List                                                                     `tfsdk:"tags"`
	TagsFilterType types.String                                                                   `tfsdk:"tags_filter_type"`
	Items          *[]ResponseItemOrganizationsGetOrganizationDevicesPowerModulesStatusesByDevice `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationDevicesPowerModulesStatusesByDevice struct {
	Mac         types.String                                                                        `tfsdk:"mac"`
	Name        types.String                                                                        `tfsdk:"name"`
	Network     *ResponseItemOrganizationsGetOrganizationDevicesPowerModulesStatusesByDeviceNetwork `tfsdk:"network"`
	ProductType types.String                                                                        `tfsdk:"product_type"`
	Serial      types.String                                                                        `tfsdk:"serial"`
	Slots       *[]ResponseItemOrganizationsGetOrganizationDevicesPowerModulesStatusesByDeviceSlots `tfsdk:"slots"`
	Tags        types.List                                                                          `tfsdk:"tags"`
}

type ResponseItemOrganizationsGetOrganizationDevicesPowerModulesStatusesByDeviceNetwork struct {
	ID types.String `tfsdk:"id"`
}

type ResponseItemOrganizationsGetOrganizationDevicesPowerModulesStatusesByDeviceSlots struct {
	Model  types.String `tfsdk:"model"`
	Number types.Int64  `tfsdk:"number"`
	Serial types.String `tfsdk:"serial"`
	Status types.String `tfsdk:"status"`
}

// ToBody
func ResponseOrganizationsGetOrganizationDevicesPowerModulesStatusesByDeviceItemsToBody(state OrganizationsDevicesPowerModulesStatusesByDevice, response *merakigosdk.ResponseOrganizationsGetOrganizationDevicesPowerModulesStatusesByDevice) OrganizationsDevicesPowerModulesStatusesByDevice {
	var items []ResponseItemOrganizationsGetOrganizationDevicesPowerModulesStatusesByDevice
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationDevicesPowerModulesStatusesByDevice{
			Mac:  types.StringValue(item.Mac),
			Name: types.StringValue(item.Name),
			Network: func() *ResponseItemOrganizationsGetOrganizationDevicesPowerModulesStatusesByDeviceNetwork {
				if item.Network != nil {
					return &ResponseItemOrganizationsGetOrganizationDevicesPowerModulesStatusesByDeviceNetwork{
						ID: types.StringValue(item.Network.ID),
					}
				}
				return nil
			}(),
			ProductType: types.StringValue(item.ProductType),
			Serial:      types.StringValue(item.Serial),
			Slots: func() *[]ResponseItemOrganizationsGetOrganizationDevicesPowerModulesStatusesByDeviceSlots {
				if item.Slots != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationDevicesPowerModulesStatusesByDeviceSlots, len(*item.Slots))
					for i, slots := range *item.Slots {
						result[i] = ResponseItemOrganizationsGetOrganizationDevicesPowerModulesStatusesByDeviceSlots{
							Model: types.StringValue(slots.Model),
							Number: func() types.Int64 {
								if slots.Number != nil {
									return types.Int64Value(int64(*slots.Number))
								}
								return types.Int64{}
							}(),
							Serial: types.StringValue(slots.Serial),
							Status: types.StringValue(slots.Status),
						}
					}
					return &result
				}
				return nil
			}(),
			Tags: StringSliceToList(item.Tags),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
