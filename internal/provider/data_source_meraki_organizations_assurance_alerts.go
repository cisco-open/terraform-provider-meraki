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
	_ datasource.DataSource              = &OrganizationsAssuranceAlertsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsAssuranceAlertsDataSource{}
)

func NewOrganizationsAssuranceAlertsDataSource() datasource.DataSource {
	return &OrganizationsAssuranceAlertsDataSource{}
}

type OrganizationsAssuranceAlertsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsAssuranceAlertsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsAssuranceAlertsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_assurance_alerts"
}

func (d *OrganizationsAssuranceAlertsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"active": schema.BoolAttribute{
				MarkdownDescription: `active query parameter. Optional parameter to filter by active alerts defaults to true`,
				Optional:            true,
			},
			"category": schema.StringAttribute{
				MarkdownDescription: `category query parameter. Optional parameter to filter by category.`,
				Optional:            true,
			},
			"device_tags": schema.ListAttribute{
				MarkdownDescription: `deviceTags query parameter. Optional parameter to filter by device tags`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"device_types": schema.ListAttribute{
				MarkdownDescription: `deviceTypes query parameter. Optional parameter to filter by device types`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"dismissed": schema.BoolAttribute{
				MarkdownDescription: `dismissed query parameter. Optional parameter to filter by dismissed alerts defaults to false`,
				Optional:            true,
			},
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId query parameter. Optional parameter to filter alerts by network ids.`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 4 300. Default is 30.`,
				Optional:            true,
			},
			"resolved": schema.BoolAttribute{
				MarkdownDescription: `resolved query parameter. Optional parameter to filter by resolved alerts defaults to false`,
				Optional:            true,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Optional parameter to filter by primary device serial`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"severity": schema.StringAttribute{
				MarkdownDescription: `severity query parameter. Optional parameter to filter by severity type.`,
				Optional:            true,
			},
			"sort_by": schema.StringAttribute{
				MarkdownDescription: `sortBy query parameter. Optional parameter to set column to sort by.`,
				Optional:            true,
			},
			"sort_order": schema.StringAttribute{
				MarkdownDescription: `sortOrder query parameter. Sorted order of entries. Order options are 'ascending' and 'descending'. Default is 'ascending'.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"suppress_alerts_for_offline_nodes": schema.BoolAttribute{
				MarkdownDescription: `suppressAlertsForOfflineNodes query parameter. When set to true the api will only return connectivity alerts for a given device if that device is in an offline state. This only applies to devices. This is ignored when resolved is true. Example:  If a Switch has a VLan Mismatch and is Unreachable. only the Unreachable alert will be returned. Defaults to false.`,
				Optional:            true,
			},
			"ts_end": schema.StringAttribute{
				MarkdownDescription: `tsEnd query parameter. Optional parameter to filter by end timestamp`,
				Optional:            true,
			},
			"ts_start": schema.StringAttribute{
				MarkdownDescription: `tsStart query parameter. Optional parameter to filter by starting timestamp`,
				Optional:            true,
			},
			"types": schema.ListAttribute{
				MarkdownDescription: `types query parameter. Optional parameter to filter by alert type.`,
				Optional:            true,
				ElementType:         types.StringType,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationAssuranceAlerts`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"category_type": schema.StringAttribute{
							MarkdownDescription: `Category type that the health alert belongs to`,
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: `Description of the alert`,
							Computed:            true,
						},
						"device_type": schema.StringAttribute{
							MarkdownDescription: `Device Type that the alert occurred on`,
							Computed:            true,
						},
						"dismissed_at": schema.StringAttribute{
							MarkdownDescription: `Time when the alert was dismissed`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `ID of the health alert`,
							Computed:            true,
						},
						"network": schema.SingleNestedAttribute{
							MarkdownDescription: `Network details`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `ID of the network where alert appears`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name of the network where alert appears`,
									Computed:            true,
								},
							},
						},
						"resolved_at": schema.StringAttribute{
							MarkdownDescription: `Time when the alert was resolved`,
							Computed:            true,
						},
						"scope": schema.SingleNestedAttribute{
							MarkdownDescription: `Scope of the alert (which devices and networks are affected)`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"applications": schema.ListAttribute{
									MarkdownDescription: `Applications affected by the alert`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"devices": schema.SetNestedAttribute{
									MarkdownDescription: `Description of affected devices`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"imei": schema.StringAttribute{
												MarkdownDescription: `IMEI of affected device`,
												Computed:            true,
											},
											"lldp": schema.SingleNestedAttribute{
												MarkdownDescription: `Port of affected device`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"port": schema.StringAttribute{
														MarkdownDescription: `Port of affect device`,
														Computed:            true,
													},
												},
											},
											"mac": schema.StringAttribute{
												MarkdownDescription: `MAC address of affected device`,
												Computed:            true,
											},
											"name": schema.StringAttribute{
												MarkdownDescription: `Name of affected device`,
												Computed:            true,
											},
											"order": schema.Int64Attribute{
												MarkdownDescription: `Order of affected device in array`,
												Computed:            true,
											},
											"product_type": schema.StringAttribute{
												MarkdownDescription: `Type of affected device`,
												Computed:            true,
											},
											"serial": schema.StringAttribute{
												MarkdownDescription: `Serial of affected device`,
												Computed:            true,
											},
											"url": schema.StringAttribute{
												MarkdownDescription: `URL of affected device`,
												Computed:            true,
											},
										},
									},
								},
								"peers": schema.SetNestedAttribute{
									MarkdownDescription: `Peers related to the alert`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"network": schema.SingleNestedAttribute{
												MarkdownDescription: `Network of the peer`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"id": schema.StringAttribute{
														MarkdownDescription: `Id of the network`,
														Computed:            true,
													},
													"name": schema.StringAttribute{
														MarkdownDescription: `Name of the network`,
														Computed:            true,
													},
												},
											},
											"url": schema.StringAttribute{
												MarkdownDescription: `URL to the peer`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
						"severity": schema.StringAttribute{
							MarkdownDescription: `Alert severity`,
							Computed:            true,
						},
						"started_at": schema.StringAttribute{
							MarkdownDescription: `Time when the alert started`,
							Computed:            true,
						},
						"title": schema.StringAttribute{
							MarkdownDescription: `Human Readable Title for Alert type`,
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `Alert Type`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsAssuranceAlertsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsAssuranceAlerts OrganizationsAssuranceAlerts
	diags := req.Config.Get(ctx, &organizationsAssuranceAlerts)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAssuranceAlerts")
		vvOrganizationID := organizationsAssuranceAlerts.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationAssuranceAlertsQueryParams{}

		queryParams1.PerPage = int(organizationsAssuranceAlerts.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsAssuranceAlerts.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsAssuranceAlerts.EndingBefore.ValueString()
		queryParams1.SortOrder = organizationsAssuranceAlerts.SortOrder.ValueString()
		queryParams1.NetworkID = organizationsAssuranceAlerts.NetworkID.ValueString()
		queryParams1.Severity = organizationsAssuranceAlerts.Severity.ValueString()
		queryParams1.Types = elementsToStrings(ctx, organizationsAssuranceAlerts.Types)
		queryParams1.TsStart = organizationsAssuranceAlerts.TsStart.ValueString()
		queryParams1.TsEnd = organizationsAssuranceAlerts.TsEnd.ValueString()
		queryParams1.Category = organizationsAssuranceAlerts.Category.ValueString()
		queryParams1.SortBy = organizationsAssuranceAlerts.SortBy.ValueString()
		queryParams1.Serials = elementsToStrings(ctx, organizationsAssuranceAlerts.Serials)
		queryParams1.DeviceTypes = elementsToStrings(ctx, organizationsAssuranceAlerts.DeviceTypes)
		queryParams1.DeviceTags = elementsToStrings(ctx, organizationsAssuranceAlerts.DeviceTags)
		queryParams1.Active = organizationsAssuranceAlerts.Active.ValueBool()
		queryParams1.Dismissed = organizationsAssuranceAlerts.Dismissed.ValueBool()
		queryParams1.Resolved = organizationsAssuranceAlerts.Resolved.ValueBool()
		queryParams1.SuppressAlertsForOfflineNodes = organizationsAssuranceAlerts.SuppressAlertsForOfflineNodes.ValueBool()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationAssuranceAlerts(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAssuranceAlerts",
				err.Error(),
			)
			return
		}

		organizationsAssuranceAlerts = ResponseOrganizationsGetOrganizationAssuranceAlertsItemsToBody(organizationsAssuranceAlerts, response1)
		diags = resp.State.Set(ctx, &organizationsAssuranceAlerts)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsAssuranceAlerts struct {
	OrganizationID                types.String                                               `tfsdk:"organization_id"`
	PerPage                       types.Int64                                                `tfsdk:"per_page"`
	StartingAfter                 types.String                                               `tfsdk:"starting_after"`
	EndingBefore                  types.String                                               `tfsdk:"ending_before"`
	SortOrder                     types.String                                               `tfsdk:"sort_order"`
	NetworkID                     types.String                                               `tfsdk:"network_id"`
	Severity                      types.String                                               `tfsdk:"severity"`
	Types                         types.List                                                 `tfsdk:"types"`
	TsStart                       types.String                                               `tfsdk:"ts_start"`
	TsEnd                         types.String                                               `tfsdk:"ts_end"`
	Category                      types.String                                               `tfsdk:"category"`
	SortBy                        types.String                                               `tfsdk:"sort_by"`
	Serials                       types.List                                                 `tfsdk:"serials"`
	DeviceTypes                   types.List                                                 `tfsdk:"device_types"`
	DeviceTags                    types.List                                                 `tfsdk:"device_tags"`
	Active                        types.Bool                                                 `tfsdk:"active"`
	Dismissed                     types.Bool                                                 `tfsdk:"dismissed"`
	Resolved                      types.Bool                                                 `tfsdk:"resolved"`
	SuppressAlertsForOfflineNodes types.Bool                                                 `tfsdk:"suppress_alerts_for_offline_nodes"`
	Items                         *[]ResponseItemOrganizationsGetOrganizationAssuranceAlerts `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationAssuranceAlerts struct {
	CategoryType types.String                                                    `tfsdk:"category_type"`
	Description  types.String                                                    `tfsdk:"description"`
	DeviceType   types.String                                                    `tfsdk:"device_type"`
	DismissedAt  types.String                                                    `tfsdk:"dismissed_at"`
	ID           types.String                                                    `tfsdk:"id"`
	Network      *ResponseItemOrganizationsGetOrganizationAssuranceAlertsNetwork `tfsdk:"network"`
	ResolvedAt   types.String                                                    `tfsdk:"resolved_at"`
	Scope        *ResponseItemOrganizationsGetOrganizationAssuranceAlertsScope   `tfsdk:"scope"`
	Severity     types.String                                                    `tfsdk:"severity"`
	StartedAt    types.String                                                    `tfsdk:"started_at"`
	Title        types.String                                                    `tfsdk:"title"`
	Type         types.String                                                    `tfsdk:"type"`
}

type ResponseItemOrganizationsGetOrganizationAssuranceAlertsNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseItemOrganizationsGetOrganizationAssuranceAlertsScope struct {
	Applications *[]ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopeApplications `tfsdk:"applications"`
	Devices      *[]ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopeDevices      `tfsdk:"devices"`
	Peers        *[]ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopePeers        `tfsdk:"peers"`
}

type ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopeApplications interface{}

type ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopeDevices struct {
	Imei        types.String                                                             `tfsdk:"imei"`
	Lldp        *ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopeDevicesLldp `tfsdk:"lldp"`
	Mac         types.String                                                             `tfsdk:"mac"`
	Name        types.String                                                             `tfsdk:"name"`
	Order       types.Int64                                                              `tfsdk:"order"`
	ProductType types.String                                                             `tfsdk:"product_type"`
	Serial      types.String                                                             `tfsdk:"serial"`
	URL         types.String                                                             `tfsdk:"url"`
}

type ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopeDevicesLldp struct {
	Port types.String `tfsdk:"port"`
}

type ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopePeers struct {
	Network *ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopePeersNetwork `tfsdk:"network"`
	URL     types.String                                                              `tfsdk:"url"`
}

type ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopePeersNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// ToBody
func ResponseOrganizationsGetOrganizationAssuranceAlertsItemsToBody(state OrganizationsAssuranceAlerts, response *merakigosdk.ResponseOrganizationsGetOrganizationAssuranceAlerts) OrganizationsAssuranceAlerts {
	var items []ResponseItemOrganizationsGetOrganizationAssuranceAlerts
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationAssuranceAlerts{
			CategoryType: types.StringValue(item.CategoryType),
			Description:  types.StringValue(item.Description),
			DeviceType:   types.StringValue(item.DeviceType),
			DismissedAt:  types.StringValue(item.DismissedAt),
			ID:           types.StringValue(item.ID),
			Network: func() *ResponseItemOrganizationsGetOrganizationAssuranceAlertsNetwork {
				if item.Network != nil {
					return &ResponseItemOrganizationsGetOrganizationAssuranceAlertsNetwork{
						ID:   types.StringValue(item.Network.ID),
						Name: types.StringValue(item.Network.Name),
					}
				}
				return nil
			}(),
			ResolvedAt: types.StringValue(item.ResolvedAt),
			Scope: func() *ResponseItemOrganizationsGetOrganizationAssuranceAlertsScope {
				if item.Scope != nil {
					return &ResponseItemOrganizationsGetOrganizationAssuranceAlertsScope{
						// Applications: StringSliceToList(item.Scope.Applications), //Interface{}
						Devices: func() *[]ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopeDevices {
							if item.Scope.Devices != nil {
								result := make([]ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopeDevices, len(*item.Scope.Devices))
								for i, devices := range *item.Scope.Devices {
									result[i] = ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopeDevices{
										Imei: types.StringValue(devices.Imei),
										Lldp: func() *ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopeDevicesLldp {
											if devices.Lldp != nil {
												return &ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopeDevicesLldp{
													Port: types.StringValue(devices.Lldp.Port),
												}
											}
											return nil
										}(),
										Mac:  types.StringValue(devices.Mac),
										Name: types.StringValue(devices.Name),
										Order: func() types.Int64 {
											if devices.Order != nil {
												return types.Int64Value(int64(*devices.Order))
											}
											return types.Int64{}
										}(),
										ProductType: types.StringValue(devices.ProductType),
										Serial:      types.StringValue(devices.Serial),
										URL:         types.StringValue(devices.URL),
									}
								}
								return &result
							}
							return nil
						}(),
						Peers: func() *[]ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopePeers {
							if item.Scope.Peers != nil {
								result := make([]ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopePeers, len(*item.Scope.Peers))
								for i, peers := range *item.Scope.Peers {
									result[i] = ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopePeers{
										Network: func() *ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopePeersNetwork {
											if peers.Network != nil {
												return &ResponseItemOrganizationsGetOrganizationAssuranceAlertsScopePeersNetwork{
													ID:   types.StringValue(peers.Network.ID),
													Name: types.StringValue(peers.Network.Name),
												}
											}
											return nil
										}(),
										URL: types.StringValue(peers.URL),
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
			Severity:  types.StringValue(item.Severity),
			StartedAt: types.StringValue(item.StartedAt),
			Title:     types.StringValue(item.Title),
			Type:      types.StringValue(item.Type),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
