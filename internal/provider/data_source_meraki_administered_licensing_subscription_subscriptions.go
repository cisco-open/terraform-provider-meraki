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
	_ datasource.DataSource              = &AdministeredLicensingSubscriptionSubscriptionsDataSource{}
	_ datasource.DataSourceWithConfigure = &AdministeredLicensingSubscriptionSubscriptionsDataSource{}
)

func NewAdministeredLicensingSubscriptionSubscriptionsDataSource() datasource.DataSource {
	return &AdministeredLicensingSubscriptionSubscriptionsDataSource{}
}

type AdministeredLicensingSubscriptionSubscriptionsDataSource struct {
	client *merakigosdk.Client
}

func (d *AdministeredLicensingSubscriptionSubscriptionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *AdministeredLicensingSubscriptionSubscriptionsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_administered_licensing_subscription_subscriptions"
}

func (d *AdministeredLicensingSubscriptionSubscriptionsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"end_date": schema.StringAttribute{
				MarkdownDescription: `endDate query parameter. Filter subscriptions by end date, ISO 8601 format. To filter with a range of dates, use 'endDate[
]=?' in the request. Accepted options include lt, gt, lte, gte.`,
				Optional: true,
			},
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"organization_ids": schema.ListAttribute{
				MarkdownDescription: `organizationIds query parameter. Organizations to get associated subscriptions for`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.`,
				Optional:            true,
			},
			"product_types": schema.ListAttribute{
				MarkdownDescription: `productTypes query parameter. List of product types that returned subscriptions need to have entitlements for.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"start_date": schema.StringAttribute{
				MarkdownDescription: `startDate query parameter. Filter subscriptions by start date, ISO 8601 format. To filter with a range of dates, use 'startDate[
]=?' in the request. Accepted options include lt, gt, lte, gte.`,
				Optional: true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"statuses": schema.ListAttribute{
				MarkdownDescription: `statuses query parameter. List of statuses that returned subscriptions can have`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"subscription_ids": schema.ListAttribute{
				MarkdownDescription: `subscriptionIds query parameter. List of subscription ids to fetch`,
				Optional:            true,
				ElementType:         types.StringType,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseLicensingGetAdministeredLicensingSubscriptionSubscriptions`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"counts": schema.SingleNestedAttribute{
							MarkdownDescription: `Numeric breakdown of network and entitlement counts`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"networks": schema.Int64Attribute{
									MarkdownDescription: `Number of networks bound to this subscription`,
									Computed:            true,
								},
								"seats": schema.SingleNestedAttribute{
									MarkdownDescription: `Seat distribution`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"assigned": schema.Int64Attribute{
											MarkdownDescription: `Number of seats in use`,
											Computed:            true,
										},
										"available": schema.Int64Attribute{
											MarkdownDescription: `Number of seats available for use`,
											Computed:            true,
										},
										"limit": schema.Int64Attribute{
											MarkdownDescription: `Total number of seats provided by this subscription`,
											Computed:            true,
										},
									},
								},
							},
						},
						"description": schema.StringAttribute{
							MarkdownDescription: `Subscription description`,
							Computed:            true,
						},
						"end_date": schema.StringAttribute{
							MarkdownDescription: `Subscription expiration date`,
							Computed:            true,
						},
						"entitlements": schema.SetNestedAttribute{
							MarkdownDescription: `Entitlement info`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"seats": schema.SingleNestedAttribute{
										MarkdownDescription: `Seat distribution`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"assigned": schema.Int64Attribute{
												MarkdownDescription: `Number of seats in use`,
												Computed:            true,
											},
											"available": schema.Int64Attribute{
												MarkdownDescription: `Number of seats available for use`,
												Computed:            true,
											},
											"limit": schema.Int64Attribute{
												MarkdownDescription: `Total number of seats provided by this subscription for this sku`,
												Computed:            true,
											},
										},
									},
									"sku": schema.StringAttribute{
										MarkdownDescription: `SKU of the required product`,
										Computed:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Subscription name`,
							Computed:            true,
						},
						"product_types": schema.ListAttribute{
							MarkdownDescription: `Products the subscription has entitlements for`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"start_date": schema.StringAttribute{
							MarkdownDescription: `Subscription start date`,
							Computed:            true,
						},
						"status": schema.StringAttribute{
							MarkdownDescription: `Subscription status`,
							Computed:            true,
						},
						"subscription_id": schema.StringAttribute{
							MarkdownDescription: `Subscription's ID`,
							Computed:            true,
						},
						"web_order_id": schema.StringAttribute{
							MarkdownDescription: `Web order id`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *AdministeredLicensingSubscriptionSubscriptionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var administeredLicensingSubscriptionSubscriptions AdministeredLicensingSubscriptionSubscriptions
	diags := req.Config.Get(ctx, &administeredLicensingSubscriptionSubscriptions)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAdministeredLicensingSubscriptionSubscriptions")
		queryParams1 := merakigosdk.GetAdministeredLicensingSubscriptionSubscriptionsQueryParams{}

		queryParams1.PerPage = int(administeredLicensingSubscriptionSubscriptions.PerPage.ValueInt64())
		queryParams1.StartingAfter = administeredLicensingSubscriptionSubscriptions.StartingAfter.ValueString()
		queryParams1.EndingBefore = administeredLicensingSubscriptionSubscriptions.EndingBefore.ValueString()
		queryParams1.SubscriptionIDs = elementsToStrings(ctx, administeredLicensingSubscriptionSubscriptions.SubscriptionIDs)
		queryParams1.OrganizationIDs = elementsToStrings(ctx, administeredLicensingSubscriptionSubscriptions.OrganizationIDs)
		queryParams1.Statuses = elementsToStrings(ctx, administeredLicensingSubscriptionSubscriptions.Statuses)
		queryParams1.ProductTypes = elementsToStrings(ctx, administeredLicensingSubscriptionSubscriptions.ProductTypes)
		queryParams1.StartDate = administeredLicensingSubscriptionSubscriptions.StartDate.ValueString()
		queryParams1.EndDate = administeredLicensingSubscriptionSubscriptions.EndDate.ValueString()

		response1, restyResp1, err := d.client.Licensing.GetAdministeredLicensingSubscriptionSubscriptions(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetAdministeredLicensingSubscriptionSubscriptions",
				err.Error(),
			)
			return
		}

		administeredLicensingSubscriptionSubscriptions = ResponseLicensingGetAdministeredLicensingSubscriptionSubscriptionsItemsToBody(administeredLicensingSubscriptionSubscriptions, response1)
		diags = resp.State.Set(ctx, &administeredLicensingSubscriptionSubscriptions)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type AdministeredLicensingSubscriptionSubscriptions struct {
	PerPage         types.Int64                                                               `tfsdk:"per_page"`
	StartingAfter   types.String                                                              `tfsdk:"starting_after"`
	EndingBefore    types.String                                                              `tfsdk:"ending_before"`
	SubscriptionIDs types.List                                                                `tfsdk:"subscription_ids"`
	OrganizationIDs types.List                                                                `tfsdk:"organization_ids"`
	Statuses        types.List                                                                `tfsdk:"statuses"`
	ProductTypes    types.List                                                                `tfsdk:"product_types"`
	StartDate       types.String                                                              `tfsdk:"start_date"`
	EndDate         types.String                                                              `tfsdk:"end_date"`
	Items           *[]ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptions `tfsdk:"items"`
}

type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptions struct {
	Counts         *ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsCounts         `tfsdk:"counts"`
	Description    types.String                                                                          `tfsdk:"description"`
	EndDate        types.String                                                                          `tfsdk:"end_date"`
	Entitlements   *[]ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsEntitlements `tfsdk:"entitlements"`
	Name           types.String                                                                          `tfsdk:"name"`
	ProductTypes   types.List                                                                            `tfsdk:"product_types"`
	StartDate      types.String                                                                          `tfsdk:"start_date"`
	Status         types.String                                                                          `tfsdk:"status"`
	SubscriptionID types.String                                                                          `tfsdk:"subscription_id"`
	WebOrderID     types.String                                                                          `tfsdk:"web_order_id"`
}

type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsCounts struct {
	Networks types.Int64                                                                        `tfsdk:"networks"`
	Seats    *ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsCountsSeats `tfsdk:"seats"`
}

type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsCountsSeats struct {
	Assigned  types.Int64 `tfsdk:"assigned"`
	Available types.Int64 `tfsdk:"available"`
	Limit     types.Int64 `tfsdk:"limit"`
}

type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsEntitlements struct {
	Seats *ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsEntitlementsSeats `tfsdk:"seats"`
	Sku   types.String                                                                             `tfsdk:"sku"`
}

type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsEntitlementsSeats struct {
	Assigned  types.Int64 `tfsdk:"assigned"`
	Available types.Int64 `tfsdk:"available"`
	Limit     types.Int64 `tfsdk:"limit"`
}

// ToBody
func ResponseLicensingGetAdministeredLicensingSubscriptionSubscriptionsItemsToBody(state AdministeredLicensingSubscriptionSubscriptions, response *merakigosdk.ResponseLicensingGetAdministeredLicensingSubscriptionSubscriptions) AdministeredLicensingSubscriptionSubscriptions {
	var items []ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptions
	for _, item := range *response {
		itemState := ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptions{
			Counts: func() *ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsCounts {
				if item.Counts != nil {
					return &ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsCounts{
						Networks: func() types.Int64 {
							if item.Counts.Networks != nil {
								return types.Int64Value(int64(*item.Counts.Networks))
							}
							return types.Int64{}
						}(),
						Seats: func() *ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsCountsSeats {
							if item.Counts.Seats != nil {
								return &ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsCountsSeats{
									Assigned: func() types.Int64 {
										if item.Counts.Seats.Assigned != nil {
											return types.Int64Value(int64(*item.Counts.Seats.Assigned))
										}
										return types.Int64{}
									}(),
									Available: func() types.Int64 {
										if item.Counts.Seats.Available != nil {
											return types.Int64Value(int64(*item.Counts.Seats.Available))
										}
										return types.Int64{}
									}(),
									Limit: func() types.Int64 {
										if item.Counts.Seats.Limit != nil {
											return types.Int64Value(int64(*item.Counts.Seats.Limit))
										}
										return types.Int64{}
									}(),
								}
							}
							return &ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsCountsSeats{}
						}(),
					}
				}
				return &ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsCounts{}
			}(),
			Description: types.StringValue(item.Description),
			EndDate:     types.StringValue(item.EndDate),
			Entitlements: func() *[]ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsEntitlements {
				if item.Entitlements != nil {
					result := make([]ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsEntitlements, len(*item.Entitlements))
					for i, entitlements := range *item.Entitlements {
						result[i] = ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsEntitlements{
							Seats: func() *ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsEntitlementsSeats {
								if entitlements.Seats != nil {
									return &ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsEntitlementsSeats{
										Assigned: func() types.Int64 {
											if entitlements.Seats.Assigned != nil {
												return types.Int64Value(int64(*entitlements.Seats.Assigned))
											}
											return types.Int64{}
										}(),
										Available: func() types.Int64 {
											if entitlements.Seats.Available != nil {
												return types.Int64Value(int64(*entitlements.Seats.Available))
											}
											return types.Int64{}
										}(),
										Limit: func() types.Int64 {
											if entitlements.Seats.Limit != nil {
												return types.Int64Value(int64(*entitlements.Seats.Limit))
											}
											return types.Int64{}
										}(),
									}
								}
								return &ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsEntitlementsSeats{}
							}(),
							Sku: types.StringValue(entitlements.Sku),
						}
					}
					return &result
				}
				return &[]ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsEntitlements{}
			}(),
			Name:           types.StringValue(item.Name),
			ProductTypes:   StringSliceToList(item.ProductTypes),
			StartDate:      types.StringValue(item.StartDate),
			Status:         types.StringValue(item.Status),
			SubscriptionID: types.StringValue(item.SubscriptionID),
			WebOrderID:     types.StringValue(item.WebOrderID),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
