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

// RESOURCE ACTION

import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &AdministeredLicensingSubscriptionSubscriptionsClaimKeyValidateResource{}
	_ resource.ResourceWithConfigure = &AdministeredLicensingSubscriptionSubscriptionsClaimKeyValidateResource{}
)

func NewAdministeredLicensingSubscriptionSubscriptionsClaimKeyValidateResource() resource.Resource {
	return &AdministeredLicensingSubscriptionSubscriptionsClaimKeyValidateResource{}
}

type AdministeredLicensingSubscriptionSubscriptionsClaimKeyValidateResource struct {
	client *merakigosdk.Client
}

func (r *AdministeredLicensingSubscriptionSubscriptionsClaimKeyValidateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *AdministeredLicensingSubscriptionSubscriptionsClaimKeyValidateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_administered_licensing_subscription_subscriptions_claim_key_validate"
}

// resourceAction
func (r *AdministeredLicensingSubscriptionSubscriptionsClaimKeyValidateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"item": schema.SingleNestedAttribute{
				Computed: true,
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
					"product_types": schema.SetAttribute{
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
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"claim_key": schema.StringAttribute{
						MarkdownDescription: `The subscription's claim key`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
				},
			},
		},
	}
}
func (r *AdministeredLicensingSubscriptionSubscriptionsClaimKeyValidateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data AdministeredLicensingSubscriptionSubscriptionsClaimKeyValidate

	var item types.Object
	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Licensing.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey(dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyItemToBody(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *AdministeredLicensingSubscriptionSubscriptionsClaimKeyValidateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *AdministeredLicensingSubscriptionSubscriptionsClaimKeyValidateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *AdministeredLicensingSubscriptionSubscriptionsClaimKeyValidateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type AdministeredLicensingSubscriptionSubscriptionsClaimKeyValidate struct {
	Item       *ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey  `tfsdk:"item"`
	Parameters *RequestLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRs `tfsdk:"parameters"`
}

type ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey struct {
	Counts         *ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyCounts         `tfsdk:"counts"`
	Description    types.String                                                                                   `tfsdk:"description"`
	EndDate        types.String                                                                                   `tfsdk:"end_date"`
	Entitlements   *[]ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyEntitlements `tfsdk:"entitlements"`
	Name           types.String                                                                                   `tfsdk:"name"`
	ProductTypes   types.Set                                                                                      `tfsdk:"product_types"`
	StartDate      types.String                                                                                   `tfsdk:"start_date"`
	Status         types.String                                                                                   `tfsdk:"status"`
	SubscriptionID types.String                                                                                   `tfsdk:"subscription_id"`
	WebOrderID     types.String                                                                                   `tfsdk:"web_order_id"`
}

type ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyCounts struct {
	Networks types.Int64                                                                                 `tfsdk:"networks"`
	Seats    *ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyCountsSeats `tfsdk:"seats"`
}

type ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyCountsSeats struct {
	Assigned  types.Int64 `tfsdk:"assigned"`
	Available types.Int64 `tfsdk:"available"`
	Limit     types.Int64 `tfsdk:"limit"`
}

type ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyEntitlements struct {
	Seats *ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyEntitlementsSeats `tfsdk:"seats"`
	Sku   types.String                                                                                      `tfsdk:"sku"`
}

type ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyEntitlementsSeats struct {
	Assigned  types.Int64 `tfsdk:"assigned"`
	Available types.Int64 `tfsdk:"available"`
	Limit     types.Int64 `tfsdk:"limit"`
}

type RequestLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRs struct {
	ClaimKey types.String `tfsdk:"claim_key"`
}

// FromBody
func (r *AdministeredLicensingSubscriptionSubscriptionsClaimKeyValidate) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey {
	emptyString := ""
	re := *r.Parameters
	claimKey := new(string)
	if !re.ClaimKey.IsUnknown() && !re.ClaimKey.IsNull() {
		*claimKey = re.ClaimKey.ValueString()
	} else {
		claimKey = &emptyString
	}
	out := merakigosdk.RequestLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey{
		ClaimKey: *claimKey,
	}
	return &out
}

// ToBody
func ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyItemToBody(state AdministeredLicensingSubscriptionSubscriptionsClaimKeyValidate, response *merakigosdk.ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey) AdministeredLicensingSubscriptionSubscriptionsClaimKeyValidate {
	itemState := ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey{
		Counts: func() *ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyCounts {
			if response.Counts != nil {
				return &ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyCounts{
					Networks: func() types.Int64 {
						if response.Counts.Networks != nil {
							return types.Int64Value(int64(*response.Counts.Networks))
						}
						return types.Int64{}
					}(),
					Seats: func() *ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyCountsSeats {
						if response.Counts.Seats != nil {
							return &ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyCountsSeats{
								Assigned: func() types.Int64 {
									if response.Counts.Seats.Assigned != nil {
										return types.Int64Value(int64(*response.Counts.Seats.Assigned))
									}
									return types.Int64{}
								}(),
								Available: func() types.Int64 {
									if response.Counts.Seats.Available != nil {
										return types.Int64Value(int64(*response.Counts.Seats.Available))
									}
									return types.Int64{}
								}(),
								Limit: func() types.Int64 {
									if response.Counts.Seats.Limit != nil {
										return types.Int64Value(int64(*response.Counts.Seats.Limit))
									}
									return types.Int64{}
								}(),
							}
						}
						return &ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyCountsSeats{}
					}(),
				}
			}
			return &ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyCounts{}
		}(),
		Description: types.StringValue(response.Description),
		EndDate:     types.StringValue(response.EndDate),
		Entitlements: func() *[]ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyEntitlements {
			if response.Entitlements != nil {
				result := make([]ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyEntitlements, len(*response.Entitlements))
				for i, entitlements := range *response.Entitlements {
					result[i] = ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyEntitlements{
						Seats: func() *ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyEntitlementsSeats {
							if entitlements.Seats != nil {
								return &ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyEntitlementsSeats{
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
							return &ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyEntitlementsSeats{}
						}(),
						Sku: types.StringValue(entitlements.Sku),
					}
				}
				return &result
			}
			return &[]ResponseLicensingValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyEntitlements{}
		}(),
		Name:           types.StringValue(response.Name),
		ProductTypes:   StringSliceToSet(response.ProductTypes),
		StartDate:      types.StringValue(response.StartDate),
		Status:         types.StringValue(response.Status),
		SubscriptionID: types.StringValue(response.SubscriptionID),
		WebOrderID:     types.StringValue(response.WebOrderID),
	}
	state.Item = &itemState
	return state
}
