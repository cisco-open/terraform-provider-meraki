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
	_ datasource.DataSource              = &OrganizationsDevicesAvailabilitiesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsDevicesAvailabilitiesDataSource{}
)

func NewOrganizationsDevicesAvailabilitiesDataSource() datasource.DataSource {
	return &OrganizationsDevicesAvailabilitiesDataSource{}
}

type OrganizationsDevicesAvailabilitiesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsDevicesAvailabilitiesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsDevicesAvailabilitiesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_devices_availabilities"
}

func (d *OrganizationsDevicesAvailabilitiesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				MarkdownDescription: `productTypes query parameter. Optional parameter to filter device availabilities by device product types. This filter uses multiple exact matches. Valid types are wireless, appliance, switch, camera, cellularGateway, sensor, wirelessController, and campusGateway`,
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
			"statuses": schema.ListAttribute{
				MarkdownDescription: `statuses query parameter. Optional parameter to filter device availabilities by device status. This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
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
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationDevicesAvailabilities`,
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
									MarkdownDescription: `ID for the network containing the device.`,
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
						"status": schema.StringAttribute{
							MarkdownDescription: `Status of the device. Possible values are: online, alerting, offline, dormant.`,
							Computed:            true,
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

func (d *OrganizationsDevicesAvailabilitiesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsDevicesAvailabilities OrganizationsDevicesAvailabilities
	diags := req.Config.Get(ctx, &organizationsDevicesAvailabilities)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationDevicesAvailabilities")
		vvOrganizationID := organizationsDevicesAvailabilities.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationDevicesAvailabilitiesQueryParams{}

		queryParams1.PerPage = int(organizationsDevicesAvailabilities.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsDevicesAvailabilities.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsDevicesAvailabilities.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsDevicesAvailabilities.NetworkIDs)
		queryParams1.ProductTypes = elementsToStrings(ctx, organizationsDevicesAvailabilities.ProductTypes)
		queryParams1.Serials = elementsToStrings(ctx, organizationsDevicesAvailabilities.Serials)
		queryParams1.Tags = elementsToStrings(ctx, organizationsDevicesAvailabilities.Tags)
		queryParams1.TagsFilterType = organizationsDevicesAvailabilities.TagsFilterType.ValueString()
		queryParams1.Statuses = elementsToStrings(ctx, organizationsDevicesAvailabilities.Statuses)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationDevicesAvailabilities(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationDevicesAvailabilities",
				err.Error(),
			)
			return
		}

		organizationsDevicesAvailabilities = ResponseOrganizationsGetOrganizationDevicesAvailabilitiesItemsToBody(organizationsDevicesAvailabilities, response1)
		diags = resp.State.Set(ctx, &organizationsDevicesAvailabilities)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsDevicesAvailabilities struct {
	OrganizationID types.String                                                     `tfsdk:"organization_id"`
	PerPage        types.Int64                                                      `tfsdk:"per_page"`
	StartingAfter  types.String                                                     `tfsdk:"starting_after"`
	EndingBefore   types.String                                                     `tfsdk:"ending_before"`
	NetworkIDs     types.List                                                       `tfsdk:"network_ids"`
	ProductTypes   types.List                                                       `tfsdk:"product_types"`
	Serials        types.List                                                       `tfsdk:"serials"`
	Tags           types.List                                                       `tfsdk:"tags"`
	TagsFilterType types.String                                                     `tfsdk:"tags_filter_type"`
	Statuses       types.List                                                       `tfsdk:"statuses"`
	Items          *[]ResponseItemOrganizationsGetOrganizationDevicesAvailabilities `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationDevicesAvailabilities struct {
	Mac         types.String                                                          `tfsdk:"mac"`
	Name        types.String                                                          `tfsdk:"name"`
	Network     *ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesNetwork `tfsdk:"network"`
	ProductType types.String                                                          `tfsdk:"product_type"`
	Serial      types.String                                                          `tfsdk:"serial"`
	Status      types.String                                                          `tfsdk:"status"`
	Tags        types.List                                                            `tfsdk:"tags"`
}

type ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesNetwork struct {
	ID types.String `tfsdk:"id"`
}

// ToBody
func ResponseOrganizationsGetOrganizationDevicesAvailabilitiesItemsToBody(state OrganizationsDevicesAvailabilities, response *merakigosdk.ResponseOrganizationsGetOrganizationDevicesAvailabilities) OrganizationsDevicesAvailabilities {
	var items []ResponseItemOrganizationsGetOrganizationDevicesAvailabilities
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationDevicesAvailabilities{
			Mac: func() types.String {
				if item.Mac != "" {
					return types.StringValue(item.Mac)
				}
				return types.String{}
			}(),
			Name: func() types.String {
				if item.Name != "" {
					return types.StringValue(item.Name)
				}
				return types.String{}
			}(),
			Network: func() *ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesNetwork {
				if item.Network != nil {
					return &ResponseItemOrganizationsGetOrganizationDevicesAvailabilitiesNetwork{
						ID: func() types.String {
							if item.Network.ID != "" {
								return types.StringValue(item.Network.ID)
							}
							return types.String{}
						}(),
					}
				}
				return nil
			}(),
			ProductType: func() types.String {
				if item.ProductType != "" {
					return types.StringValue(item.ProductType)
				}
				return types.String{}
			}(),
			Serial: func() types.String {
				if item.Serial != "" {
					return types.StringValue(item.Serial)
				}
				return types.String{}
			}(),
			Status: func() types.String {
				if item.Status != "" {
					return types.StringValue(item.Status)
				}
				return types.String{}
			}(),
			Tags: StringSliceToList(item.Tags),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
