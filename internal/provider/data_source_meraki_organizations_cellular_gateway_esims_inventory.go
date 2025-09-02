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
	_ datasource.DataSource              = &OrganizationsCellularGatewayEsimsInventoryDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsCellularGatewayEsimsInventoryDataSource{}
)

func NewOrganizationsCellularGatewayEsimsInventoryDataSource() datasource.DataSource {
	return &OrganizationsCellularGatewayEsimsInventoryDataSource{}
}

type OrganizationsCellularGatewayEsimsInventoryDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsCellularGatewayEsimsInventoryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsCellularGatewayEsimsInventoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_cellular_gateway_esims_inventory"
}

func (d *OrganizationsCellularGatewayEsimsInventoryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"eids": schema.ListAttribute{
				MarkdownDescription: `eids query parameter. Optional parameter to filter the results by EID.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"items": schema.ListNestedAttribute{
						MarkdownDescription: `List of eSIM Devices`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"active": schema.BoolAttribute{
									MarkdownDescription: `Whether eSIM is currently active SIM on Device`,
									Computed:            true,
								},
								"device": schema.SingleNestedAttribute{
									MarkdownDescription: `Meraki Device properties`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"model": schema.StringAttribute{
											MarkdownDescription: `Device model`,
											Computed:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: `Device name`,
											Computed:            true,
										},
										"serial": schema.StringAttribute{
											MarkdownDescription: `Device serial number`,
											Computed:            true,
										},
										"status": schema.StringAttribute{
											MarkdownDescription: `Device status`,
											Computed:            true,
										},
										"url": schema.StringAttribute{
											MarkdownDescription: `Device URL`,
											Computed:            true,
										},
									},
								},
								"eid": schema.StringAttribute{
									MarkdownDescription: `eSIM EID`,
									Computed:            true,
								},
								"last_updated_at": schema.StringAttribute{
									MarkdownDescription: `Last update of eSIM`,
									Computed:            true,
								},
								"network": schema.SingleNestedAttribute{
									MarkdownDescription: `Meraki Network properties`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `Network ID for this eSIM`,
											Computed:            true,
										},
									},
								},
								"profiles": schema.SetNestedAttribute{
									MarkdownDescription: `eSIM Profile Information`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"custom_apns": schema.ListAttribute{
												MarkdownDescription: `Available custom APNs for the profile`,
												Computed:            true,
												ElementType:         types.StringType,
											},
											"iccid": schema.StringAttribute{
												MarkdownDescription: `eSIM profile ID`,
												Computed:            true,
											},
											"service_provider": schema.SingleNestedAttribute{
												MarkdownDescription: `Service Provider information`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"name": schema.StringAttribute{
														MarkdownDescription: `Service Provider name`,
														Computed:            true,
													},
													"plans": schema.SetNestedAttribute{
														MarkdownDescription: `Plans currently active on the eSIM`,
														Computed:            true,
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{

																"name": schema.StringAttribute{
																	MarkdownDescription: `Plan name`,
																	Computed:            true,
																},
																"type": schema.StringAttribute{
																	MarkdownDescription: `Plan type (communication, rate)`,
																	Computed:            true,
																},
															},
														},
													},
												},
											},
											"status": schema.StringAttribute{
												MarkdownDescription: `eSIM profile status`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
					},
					"meta": schema.SingleNestedAttribute{
						MarkdownDescription: `Meta details about the result`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"counts": schema.SingleNestedAttribute{
								MarkdownDescription: `Counts of involved entities`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"items": schema.SingleNestedAttribute{
										MarkdownDescription: `Count of eSIM Devices available`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"remaining": schema.Int64Attribute{
												MarkdownDescription: `Remaining number of eSIM Devices`,
												Computed:            true,
											},
											"total": schema.Int64Attribute{
												MarkdownDescription: `Total number of eSIM Devices`,
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

func (d *OrganizationsCellularGatewayEsimsInventoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsCellularGatewayEsimsInventory OrganizationsCellularGatewayEsimsInventory
	diags := req.Config.Get(ctx, &organizationsCellularGatewayEsimsInventory)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationCellularGatewayEsimsInventory")
		vvOrganizationID := organizationsCellularGatewayEsimsInventory.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationCellularGatewayEsimsInventoryQueryParams{}

		queryParams1.Eids = elementsToStrings(ctx, organizationsCellularGatewayEsimsInventory.Eids)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.CellularGateway.GetOrganizationCellularGatewayEsimsInventory(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCellularGatewayEsimsInventory",
				err.Error(),
			)
			return
		}

		organizationsCellularGatewayEsimsInventory = ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemToBody(organizationsCellularGatewayEsimsInventory, response1)
		diags = resp.State.Set(ctx, &organizationsCellularGatewayEsimsInventory)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsCellularGatewayEsimsInventory struct {
	OrganizationID types.String                                                         `tfsdk:"organization_id"`
	Eids           types.List                                                           `tfsdk:"eids"`
	Item           *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventory `tfsdk:"item"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventory struct {
	Items *[]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItems `tfsdk:"items"`
	Meta  *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryMeta    `tfsdk:"meta"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItems struct {
	Active        types.Bool                                                                          `tfsdk:"active"`
	Device        *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsDevice     `tfsdk:"device"`
	Eid           types.String                                                                        `tfsdk:"eid"`
	LastUpdatedAt types.String                                                                        `tfsdk:"last_updated_at"`
	Network       *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsNetwork    `tfsdk:"network"`
	Profiles      *[]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfiles `tfsdk:"profiles"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsDevice struct {
	Model  types.String `tfsdk:"model"`
	Name   types.String `tfsdk:"name"`
	Serial types.String `tfsdk:"serial"`
	Status types.String `tfsdk:"status"`
	URL    types.String `tfsdk:"url"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsNetwork struct {
	ID types.String `tfsdk:"id"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfiles struct {
	CustomApns      types.List                                                                                       `tfsdk:"custom_apns"`
	Iccid           types.String                                                                                     `tfsdk:"iccid"`
	ServiceProvider *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfilesServiceProvider `tfsdk:"service_provider"`
	Status          types.String                                                                                     `tfsdk:"status"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfilesServiceProvider struct {
	Name  types.String                                                                                            `tfsdk:"name"`
	Plans *[]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfilesServiceProviderPlans `tfsdk:"plans"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfilesServiceProviderPlans struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryMeta struct {
	Counts *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryMetaCounts `tfsdk:"counts"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryMetaCounts struct {
	Items *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryMetaCountsItems `tfsdk:"items"`
}

type ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemToBody(state OrganizationsCellularGatewayEsimsInventory, response *merakigosdk.ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventory) OrganizationsCellularGatewayEsimsInventory {
	itemState := ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventory{
		Items: func() *[]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItems {
			if response.Items != nil {
				result := make([]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItems{
						Active: func() types.Bool {
							if items.Active != nil {
								return types.BoolValue(*items.Active)
							}
							return types.Bool{}
						}(),
						Device: func() *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsDevice {
							if items.Device != nil {
								return &ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsDevice{
									Model: func() types.String {
										if items.Device.Model != "" {
											return types.StringValue(items.Device.Model)
										}
										return types.String{}
									}(),
									Name: func() types.String {
										if items.Device.Name != "" {
											return types.StringValue(items.Device.Name)
										}
										return types.String{}
									}(),
									Serial: func() types.String {
										if items.Device.Serial != "" {
											return types.StringValue(items.Device.Serial)
										}
										return types.String{}
									}(),
									Status: func() types.String {
										if items.Device.Status != "" {
											return types.StringValue(items.Device.Status)
										}
										return types.String{}
									}(),
									URL: func() types.String {
										if items.Device.URL != "" {
											return types.StringValue(items.Device.URL)
										}
										return types.String{}
									}(),
								}
							}
							return nil
						}(),
						Eid: func() types.String {
							if items.Eid != "" {
								return types.StringValue(items.Eid)
							}
							return types.String{}
						}(),
						LastUpdatedAt: func() types.String {
							if items.LastUpdatedAt != "" {
								return types.StringValue(items.LastUpdatedAt)
							}
							return types.String{}
						}(),
						Network: func() *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsNetwork {
							if items.Network != nil {
								return &ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsNetwork{
									ID: func() types.String {
										if items.Network.ID != "" {
											return types.StringValue(items.Network.ID)
										}
										return types.String{}
									}(),
								}
							}
							return nil
						}(),
						Profiles: func() *[]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfiles {
							if items.Profiles != nil {
								result := make([]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfiles, len(*items.Profiles))
								for i, profiles := range *items.Profiles {
									result[i] = ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfiles{
										CustomApns: StringSliceToList(profiles.CustomApns),
										Iccid: func() types.String {
											if profiles.Iccid != "" {
												return types.StringValue(profiles.Iccid)
											}
											return types.String{}
										}(),
										ServiceProvider: func() *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfilesServiceProvider {
											if profiles.ServiceProvider != nil {
												return &ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfilesServiceProvider{
													Name: func() types.String {
														if profiles.ServiceProvider.Name != "" {
															return types.StringValue(profiles.ServiceProvider.Name)
														}
														return types.String{}
													}(),
													Plans: func() *[]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfilesServiceProviderPlans {
														if profiles.ServiceProvider.Plans != nil {
															result := make([]ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfilesServiceProviderPlans, len(*profiles.ServiceProvider.Plans))
															for i, plans := range *profiles.ServiceProvider.Plans {
																result[i] = ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryItemsProfilesServiceProviderPlans{
																	Name: func() types.String {
																		if plans.Name != "" {
																			return types.StringValue(plans.Name)
																		}
																		return types.String{}
																	}(),
																	Type: func() types.String {
																		if plans.Type != "" {
																			return types.StringValue(plans.Type)
																		}
																		return types.String{}
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
										Status: func() types.String {
											if profiles.Status != "" {
												return types.StringValue(profiles.Status)
											}
											return types.String{}
										}(),
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
		Meta: func() *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryMeta {
			if response.Meta != nil {
				return &ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryMeta{
					Counts: func() *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryMetaCounts{
								Items: func() *ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseCellularGatewayGetOrganizationCellularGatewayEsimsInventoryMetaCountsItems{
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
