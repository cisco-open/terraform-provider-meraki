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
	_ datasource.DataSource              = &AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesDataSource{}
	_ datasource.DataSourceWithConfigure = &AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesDataSource{}
)

func NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesDataSource() datasource.DataSource {
	return &AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesDataSource{}
}

type AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesDataSource struct {
	client *merakigosdk.Client
}

func (d *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_administered_licensing_subscription_subscriptions_compliance_statuses"
}

func (d *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_ids": schema.ListAttribute{
				MarkdownDescription: `organizationIds query parameter. Organizations to get subscription compliance information for`,
				Required:            true,
				ElementType:         types.StringType,
			},
			"subscription_ids": schema.ListAttribute{
				MarkdownDescription: `subscriptionIds query parameter. Subscription ids`,
				Optional:            true,
				ElementType:         types.StringType,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"subscription": schema.SingleNestedAttribute{
							MarkdownDescription: `Subscription details`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `Subscription's ID`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Friendly name to identify the subscription`,
									Computed:            true,
								},
								"status": schema.StringAttribute{
									MarkdownDescription: `One of the following: "inactive" | "active" | "out_of_compliance" | "expired" | "canceled"`,
									Computed:            true,
								},
							},
						},
						"violations": schema.SingleNestedAttribute{
							MarkdownDescription: `Violations`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"by_product_class": schema.SetNestedAttribute{
									MarkdownDescription: `List of violations by product class that are not compliance`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"grace_period_ends_at": schema.StringAttribute{
												MarkdownDescription: `End date of the grace period in ISO 8601 format`,
												Computed:            true,
											},
											"missing": schema.SingleNestedAttribute{
												MarkdownDescription: `Missing entitlements details`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"entitlements": schema.SetNestedAttribute{
														MarkdownDescription: `List of missing entitlements`,
														Computed:            true,
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{

																"quantity": schema.Int64Attribute{
																	MarkdownDescription: `Number required`,
																	Computed:            true,
																},
																"sku": schema.StringAttribute{
																	MarkdownDescription: `SKU of the required product`,
																	Computed:            true,
																},
															},
														},
													},
												},
											},
											"product_class": schema.StringAttribute{
												MarkdownDescription: `Name of the product class`,
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

func (d *AdministeredLicensingSubscriptionSubscriptionsComplianceStatusesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var administeredLicensingSubscriptionSubscriptionsComplianceStatuses AdministeredLicensingSubscriptionSubscriptionsComplianceStatuses
	diags := req.Config.Get(ctx, &administeredLicensingSubscriptionSubscriptionsComplianceStatuses)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses")
		queryParams1 := merakigosdk.GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesQueryParams{}

		queryParams1.OrganizationIDs = elementsToStrings(ctx, administeredLicensingSubscriptionSubscriptionsComplianceStatuses.OrganizationIDs)

		queryParams1.SubscriptionIDs = elementsToStrings(ctx, administeredLicensingSubscriptionSubscriptionsComplianceStatuses.SubscriptionIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Licensing.GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses",
				err.Error(),
			)
			return
		}

		administeredLicensingSubscriptionSubscriptionsComplianceStatuses = ResponseLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesItemsToBody(administeredLicensingSubscriptionSubscriptionsComplianceStatuses, response1)
		diags = resp.State.Set(ctx, &administeredLicensingSubscriptionSubscriptionsComplianceStatuses)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type AdministeredLicensingSubscriptionSubscriptionsComplianceStatuses struct {
	OrganizationIDs types.List                                                                                  `tfsdk:"organization_ids"`
	SubscriptionIDs types.List                                                                                  `tfsdk:"subscription_ids"`
	Items           *[]ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses `tfsdk:"items"`
}

type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses struct {
	Subscription *ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription `tfsdk:"subscription"`
	Violations   *ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations   `tfsdk:"violations"`
}

type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription struct {
	ID     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Status types.String `tfsdk:"status"`
}

type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations struct {
	ByProductClass *[]ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass `tfsdk:"by_product_class"`
}

type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass struct {
	GracePeriodEndsAt types.String                                                                                                             `tfsdk:"grace_period_ends_at"`
	Missing           *ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClassMissing `tfsdk:"missing"`
	ProductClass      types.String                                                                                                             `tfsdk:"product_class"`
}

type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClassMissing struct {
	Entitlements *[]ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClassMissingEntitlements `tfsdk:"entitlements"`
}

type ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClassMissingEntitlements struct {
	Quantity types.Int64  `tfsdk:"quantity"`
	Sku      types.String `tfsdk:"sku"`
}

// ToBody
func ResponseLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesItemsToBody(state AdministeredLicensingSubscriptionSubscriptionsComplianceStatuses, response *merakigosdk.ResponseLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses) AdministeredLicensingSubscriptionSubscriptionsComplianceStatuses {
	var items []ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses
	for _, item := range *response {
		itemState := ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses{
			Subscription: func() *ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription {
				if item.Subscription != nil {
					return &ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesSubscription{
						ID:     types.StringValue(item.Subscription.ID),
						Name:   types.StringValue(item.Subscription.Name),
						Status: types.StringValue(item.Subscription.Status),
					}
				}
				return nil
			}(),
			Violations: func() *ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations {
				if item.Violations != nil {
					return &ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolations{
						ByProductClass: func() *[]ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass {
							if item.Violations.ByProductClass != nil {
								result := make([]ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass, len(*item.Violations.ByProductClass))
								for i, byProductClass := range *item.Violations.ByProductClass {
									result[i] = ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClass{
										GracePeriodEndsAt: types.StringValue(byProductClass.GracePeriodEndsAt),
										Missing: func() *ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClassMissing {
											if byProductClass.Missing != nil {
												return &ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClassMissing{
													Entitlements: func() *[]ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClassMissingEntitlements {
														if byProductClass.Missing.Entitlements != nil {
															result := make([]ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClassMissingEntitlements, len(*byProductClass.Missing.Entitlements))
															for i, entitlements := range *byProductClass.Missing.Entitlements {
																result[i] = ResponseItemLicensingGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesViolationsByProductClassMissingEntitlements{
																	Quantity: func() types.Int64 {
																		if entitlements.Quantity != nil {
																			return types.Int64Value(int64(*entitlements.Quantity))
																		}
																		return types.Int64{}
																	}(),
																	Sku: types.StringValue(entitlements.Sku),
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
										ProductClass: types.StringValue(byProductClass.ProductClass),
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
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
