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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &AdministeredLicensingSubscriptionSubscriptionsBindResource{}
	_ resource.ResourceWithConfigure = &AdministeredLicensingSubscriptionSubscriptionsBindResource{}
)

func NewAdministeredLicensingSubscriptionSubscriptionsBindResource() resource.Resource {
	return &AdministeredLicensingSubscriptionSubscriptionsBindResource{}
}

type AdministeredLicensingSubscriptionSubscriptionsBindResource struct {
	client *merakigosdk.Client
}

func (r *AdministeredLicensingSubscriptionSubscriptionsBindResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *AdministeredLicensingSubscriptionSubscriptionsBindResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_administered_licensing_subscription_subscriptions_bind"
}

// resourceAction
func (r *AdministeredLicensingSubscriptionSubscriptionsBindResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"subscription_id": schema.StringAttribute{
				MarkdownDescription: `subscriptionId path parameter. Subscription ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"validate": schema.BoolAttribute{
				MarkdownDescription: `validate query parameter. Check if the provided networks can be bound to the subscription. Returns any licensing problems and does not commit the results.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"errors": schema.ListAttribute{
						MarkdownDescription: `Array of errors if failed`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"insufficient_entitlements": schema.SetNestedAttribute{
						MarkdownDescription: `A list of entitlements required to successfully bind the networks to the subscription`,
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
					"networks": schema.SetNestedAttribute{
						MarkdownDescription: `Unbound networks`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `Network ID`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Network name`,
									Computed:            true,
								},
							},
						},
					},
					"subscription_id": schema.StringAttribute{
						MarkdownDescription: `Subscription ID`,
						Computed:            true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"network_ids": schema.ListAttribute{
						MarkdownDescription: `List of network ids to bind to the subscription`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}
func (r *AdministeredLicensingSubscriptionSubscriptionsBindResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data AdministeredLicensingSubscriptionSubscriptionsBind

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
	vvSubscriptionID := data.SubscriptionID.ValueString()
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Licensing.BindAdministeredLicensingSubscriptionSubscription(vvSubscriptionID, dataRequest, nil)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing BindAdministeredLicensingSubscriptionSubscription",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing BindAdministeredLicensingSubscriptionSubscription",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseLicensingBindAdministeredLicensingSubscriptionSubscriptionItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *AdministeredLicensingSubscriptionSubscriptionsBindResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *AdministeredLicensingSubscriptionSubscriptionsBindResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *AdministeredLicensingSubscriptionSubscriptionsBindResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type AdministeredLicensingSubscriptionSubscriptionsBind struct {
	SubscriptionID types.String                                                         `tfsdk:"subscription_id"`
	Item           *ResponseLicensingBindAdministeredLicensingSubscriptionSubscription  `tfsdk:"item"`
	Parameters     *RequestLicensingBindAdministeredLicensingSubscriptionSubscriptionRs `tfsdk:"parameters"`
}

type ResponseLicensingBindAdministeredLicensingSubscriptionSubscription struct {
	Errors                   types.List                                                                                    `tfsdk:"errors"`
	InsufficientEntitlements *[]ResponseLicensingBindAdministeredLicensingSubscriptionSubscriptionInsufficientEntitlements `tfsdk:"insufficient_entitlements"`
	Networks                 *[]ResponseLicensingBindAdministeredLicensingSubscriptionSubscriptionNetworks                 `tfsdk:"networks"`
	SubscriptionID           types.String                                                                                  `tfsdk:"subscription_id"`
}

type ResponseLicensingBindAdministeredLicensingSubscriptionSubscriptionInsufficientEntitlements struct {
	Quantity types.Int64  `tfsdk:"quantity"`
	Sku      types.String `tfsdk:"sku"`
}

type ResponseLicensingBindAdministeredLicensingSubscriptionSubscriptionNetworks struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type RequestLicensingBindAdministeredLicensingSubscriptionSubscriptionRs struct {
	NetworkIDs types.Set `tfsdk:"network_ids"`
}

// FromBody
func (r *AdministeredLicensingSubscriptionSubscriptionsBind) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestLicensingBindAdministeredLicensingSubscriptionSubscription {
	re := *r.Parameters
	var networkIDs []string = nil
	re.NetworkIDs.ElementsAs(ctx, &networkIDs, false)
	out := merakigosdk.RequestLicensingBindAdministeredLicensingSubscriptionSubscription{
		NetworkIDs: networkIDs,
	}
	return &out
}

// ToBody
func ResponseLicensingBindAdministeredLicensingSubscriptionSubscriptionItemToBody(state AdministeredLicensingSubscriptionSubscriptionsBind, response *merakigosdk.ResponseLicensingBindAdministeredLicensingSubscriptionSubscription) AdministeredLicensingSubscriptionSubscriptionsBind {
	itemState := ResponseLicensingBindAdministeredLicensingSubscriptionSubscription{
		Errors: StringSliceToList(response.Errors),
		InsufficientEntitlements: func() *[]ResponseLicensingBindAdministeredLicensingSubscriptionSubscriptionInsufficientEntitlements {
			if response.InsufficientEntitlements != nil {
				result := make([]ResponseLicensingBindAdministeredLicensingSubscriptionSubscriptionInsufficientEntitlements, len(*response.InsufficientEntitlements))
				for i, insufficientEntitlements := range *response.InsufficientEntitlements {
					result[i] = ResponseLicensingBindAdministeredLicensingSubscriptionSubscriptionInsufficientEntitlements{
						Quantity: func() types.Int64 {
							if insufficientEntitlements.Quantity != nil {
								return types.Int64Value(int64(*insufficientEntitlements.Quantity))
							}
							return types.Int64{}
						}(),
						Sku: types.StringValue(insufficientEntitlements.Sku),
					}
				}
				return &result
			}
			return nil
		}(),
		Networks: func() *[]ResponseLicensingBindAdministeredLicensingSubscriptionSubscriptionNetworks {
			if response.Networks != nil {
				result := make([]ResponseLicensingBindAdministeredLicensingSubscriptionSubscriptionNetworks, len(*response.Networks))
				for i, networks := range *response.Networks {
					result[i] = ResponseLicensingBindAdministeredLicensingSubscriptionSubscriptionNetworks{
						ID:   types.StringValue(networks.ID),
						Name: types.StringValue(networks.Name),
					}
				}
				return &result
			}
			return nil
		}(),
		SubscriptionID: types.StringValue(response.SubscriptionID),
	}
	state.Item = &itemState
	return state
}
