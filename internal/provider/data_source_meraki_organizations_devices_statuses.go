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
	_ datasource.DataSource              = &OrganizationsDevicesStatusesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsDevicesStatusesDataSource{}
)

func NewOrganizationsDevicesStatusesDataSource() datasource.DataSource {
	return &OrganizationsDevicesStatusesDataSource{}
}

type OrganizationsDevicesStatusesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsDevicesStatusesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsDevicesStatusesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_devices_statuses"
}

func (d *OrganizationsDevicesStatusesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"models": schema.ListAttribute{
				MarkdownDescription: `models query parameter. Optional parameter to filter devices by models.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter devices by network ids.`,
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
				MarkdownDescription: `productTypes query parameter. An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, and secureConnect.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Optional parameter to filter devices by serials.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"statuses": schema.ListAttribute{
				MarkdownDescription: `statuses query parameter. Optional parameter to filter devices by statuses. Valid statuses are ["online", "alerting", "offline", "dormant"].`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"tags": schema.ListAttribute{
				MarkdownDescription: `tags query parameter. An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below).`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"tags_filter_type": schema.StringAttribute{
				MarkdownDescription: `tagsFilterType query parameter. An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationDevicesStatuses`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"components": schema.SingleNestedAttribute{
							MarkdownDescription: `Components`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"power_supplies": schema.SetNestedAttribute{
									MarkdownDescription: `Power Supplies`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"model": schema.StringAttribute{
												MarkdownDescription: `Model of the power supply`,
												Computed:            true,
											},
											"poe": schema.SingleNestedAttribute{
												MarkdownDescription: `PoE info of the power supply`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"maximum": schema.Int64Attribute{
														MarkdownDescription: `Maximum PoE this power supply can provide when connected to the current switch model`,
														Computed:            true,
													},
													"unit": schema.StringAttribute{
														MarkdownDescription: `Unit of the PoE maximum`,
														Computed:            true,
													},
												},
											},
											"serial": schema.StringAttribute{
												MarkdownDescription: `Serial of the power supply`,
												Computed:            true,
											},
											"slot": schema.Int64Attribute{
												MarkdownDescription: `Slot the power supply is in`,
												Computed:            true,
											},
											"status": schema.StringAttribute{
												MarkdownDescription: `Status of the power supply`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
						"gateway": schema.StringAttribute{
							MarkdownDescription: `IP Gateway`,
							Computed:            true,
						},
						"ip_type": schema.StringAttribute{
							MarkdownDescription: `IP Type`,
							Computed:            true,
						},
						"lan_ip": schema.StringAttribute{
							MarkdownDescription: `LAN IP Address`,
							Computed:            true,
						},
						"last_reported_at": schema.StringAttribute{
							MarkdownDescription: `Device Last Reported Location`,
							Computed:            true,
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: `MAC Address`,
							Computed:            true,
						},
						"model": schema.StringAttribute{
							MarkdownDescription: `Model`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Device Name`,
							Computed:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: `Network ID`,
							Computed:            true,
						},
						"primary_dns": schema.StringAttribute{
							MarkdownDescription: `Primary DNS`,
							Computed:            true,
						},
						"product_type": schema.StringAttribute{
							MarkdownDescription: `Product Type`,
							Computed:            true,
						},
						"public_ip": schema.StringAttribute{
							MarkdownDescription: `Public IP Address`,
							Computed:            true,
						},
						"secondary_dns": schema.StringAttribute{
							MarkdownDescription: `Secondary DNS`,
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `Device Serial Number`,
							Computed:            true,
						},
						"status": schema.StringAttribute{
							MarkdownDescription: `Device Status`,
							Computed:            true,
						},
						"tags": schema.ListAttribute{
							MarkdownDescription: `Tags`,
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsDevicesStatusesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsDevicesStatuses OrganizationsDevicesStatuses
	diags := req.Config.Get(ctx, &organizationsDevicesStatuses)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationDevicesStatuses")
		vvOrganizationID := organizationsDevicesStatuses.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationDevicesStatusesQueryParams{}

		queryParams1.PerPage = int(organizationsDevicesStatuses.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsDevicesStatuses.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsDevicesStatuses.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsDevicesStatuses.NetworkIDs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsDevicesStatuses.Serials)
		queryParams1.Statuses = elementsToStrings(ctx, organizationsDevicesStatuses.Statuses)
		queryParams1.ProductTypes = elementsToStrings(ctx, organizationsDevicesStatuses.ProductTypes)
		queryParams1.Models = elementsToStrings(ctx, organizationsDevicesStatuses.Models)
		queryParams1.Tags = elementsToStrings(ctx, organizationsDevicesStatuses.Tags)
		queryParams1.TagsFilterType = organizationsDevicesStatuses.TagsFilterType.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationDevicesStatuses(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationDevicesStatuses",
				err.Error(),
			)
			return
		}

		organizationsDevicesStatuses = ResponseOrganizationsGetOrganizationDevicesStatusesItemsToBody(organizationsDevicesStatuses, response1)
		diags = resp.State.Set(ctx, &organizationsDevicesStatuses)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsDevicesStatuses struct {
	OrganizationID types.String                                               `tfsdk:"organization_id"`
	PerPage        types.Int64                                                `tfsdk:"per_page"`
	StartingAfter  types.String                                               `tfsdk:"starting_after"`
	EndingBefore   types.String                                               `tfsdk:"ending_before"`
	NetworkIDs     types.List                                                 `tfsdk:"network_ids"`
	Serials        types.List                                                 `tfsdk:"serials"`
	Statuses       types.List                                                 `tfsdk:"statuses"`
	ProductTypes   types.List                                                 `tfsdk:"product_types"`
	Models         types.List                                                 `tfsdk:"models"`
	Tags           types.List                                                 `tfsdk:"tags"`
	TagsFilterType types.String                                               `tfsdk:"tags_filter_type"`
	Items          *[]ResponseItemOrganizationsGetOrganizationDevicesStatuses `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationDevicesStatuses struct {
	Components     *ResponseItemOrganizationsGetOrganizationDevicesStatusesComponents `tfsdk:"components"`
	Gateway        types.String                                                       `tfsdk:"gateway"`
	IPType         types.String                                                       `tfsdk:"ip_type"`
	LanIP          types.String                                                       `tfsdk:"lan_ip"`
	LastReportedAt types.String                                                       `tfsdk:"last_reported_at"`
	Mac            types.String                                                       `tfsdk:"mac"`
	Model          types.String                                                       `tfsdk:"model"`
	Name           types.String                                                       `tfsdk:"name"`
	NetworkID      types.String                                                       `tfsdk:"network_id"`
	PrimaryDNS     types.String                                                       `tfsdk:"primary_dns"`
	ProductType    types.String                                                       `tfsdk:"product_type"`
	PublicIP       types.String                                                       `tfsdk:"public_ip"`
	SecondaryDNS   types.String                                                       `tfsdk:"secondary_dns"`
	Serial         types.String                                                       `tfsdk:"serial"`
	Status         types.String                                                       `tfsdk:"status"`
	Tags           types.List                                                         `tfsdk:"tags"`
}

type ResponseItemOrganizationsGetOrganizationDevicesStatusesComponents struct {
	PowerSupplies *[]ResponseItemOrganizationsGetOrganizationDevicesStatusesComponentsPowerSupplies `tfsdk:"power_supplies"`
}

type ResponseItemOrganizationsGetOrganizationDevicesStatusesComponentsPowerSupplies struct {
	Model  types.String                                                                       `tfsdk:"model"`
	Poe    *ResponseItemOrganizationsGetOrganizationDevicesStatusesComponentsPowerSuppliesPoe `tfsdk:"poe"`
	Serial types.String                                                                       `tfsdk:"serial"`
	Slot   types.Int64                                                                        `tfsdk:"slot"`
	Status types.String                                                                       `tfsdk:"status"`
}

type ResponseItemOrganizationsGetOrganizationDevicesStatusesComponentsPowerSuppliesPoe struct {
	Maximum types.Int64  `tfsdk:"maximum"`
	Unit    types.String `tfsdk:"unit"`
}

// ToBody
func ResponseOrganizationsGetOrganizationDevicesStatusesItemsToBody(state OrganizationsDevicesStatuses, response *merakigosdk.ResponseOrganizationsGetOrganizationDevicesStatuses) OrganizationsDevicesStatuses {
	var items []ResponseItemOrganizationsGetOrganizationDevicesStatuses
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationDevicesStatuses{
			Components: func() *ResponseItemOrganizationsGetOrganizationDevicesStatusesComponents {
				if item.Components != nil {
					return &ResponseItemOrganizationsGetOrganizationDevicesStatusesComponents{
						PowerSupplies: func() *[]ResponseItemOrganizationsGetOrganizationDevicesStatusesComponentsPowerSupplies {
							if item.Components.PowerSupplies != nil {
								result := make([]ResponseItemOrganizationsGetOrganizationDevicesStatusesComponentsPowerSupplies, len(*item.Components.PowerSupplies))
								for i, powerSupplies := range *item.Components.PowerSupplies {
									result[i] = ResponseItemOrganizationsGetOrganizationDevicesStatusesComponentsPowerSupplies{
										Model: types.StringValue(powerSupplies.Model),
										Poe: func() *ResponseItemOrganizationsGetOrganizationDevicesStatusesComponentsPowerSuppliesPoe {
											if powerSupplies.Poe != nil {
												return &ResponseItemOrganizationsGetOrganizationDevicesStatusesComponentsPowerSuppliesPoe{
													Maximum: func() types.Int64 {
														if powerSupplies.Poe.Maximum != nil {
															return types.Int64Value(int64(*powerSupplies.Poe.Maximum))
														}
														return types.Int64{}
													}(),
													Unit: types.StringValue(powerSupplies.Poe.Unit),
												}
											}
											return nil
										}(),
										Serial: types.StringValue(powerSupplies.Serial),
										Slot: func() types.Int64 {
											if powerSupplies.Slot != nil {
												return types.Int64Value(int64(*powerSupplies.Slot))
											}
											return types.Int64{}
										}(),
										Status: types.StringValue(powerSupplies.Status),
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
			Gateway:        types.StringValue(item.Gateway),
			IPType:         types.StringValue(item.IPType),
			LanIP:          types.StringValue(item.LanIP),
			LastReportedAt: types.StringValue(item.LastReportedAt),
			Mac:            types.StringValue(item.Mac),
			Model:          types.StringValue(item.Model),
			Name:           types.StringValue(item.Name),
			NetworkID:      types.StringValue(item.NetworkID),
			PrimaryDNS:     types.StringValue(item.PrimaryDNS),
			ProductType:    types.StringValue(item.ProductType),
			PublicIP:       types.StringValue(item.PublicIP),
			SecondaryDNS:   types.StringValue(item.SecondaryDNS),
			Serial:         types.StringValue(item.Serial),
			Status:         types.StringValue(item.Status),
			Tags:           StringSliceToList(item.Tags),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
