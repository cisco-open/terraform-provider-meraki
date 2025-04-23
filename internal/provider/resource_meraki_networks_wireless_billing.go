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

// RESOURCE NORMAL
import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWirelessBillingResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessBillingResource{}
)

func NewNetworksWirelessBillingResource() resource.Resource {
	return &NetworksWirelessBillingResource{}
}

type NetworksWirelessBillingResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessBillingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessBillingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_billing"
}

func (r *NetworksWirelessBillingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"currency": schema.StringAttribute{
				MarkdownDescription: `The currency code of this node group's billing plans`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"plans": schema.SetNestedAttribute{
				MarkdownDescription: `Array of billing plans in the node group. (Can configure a maximum of 5)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"bandwidth_limits": schema.SingleNestedAttribute{
							MarkdownDescription: `The uplink bandwidth settings for the pricing plan.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"limit_down": schema.Int64Attribute{
									MarkdownDescription: `The maximum download limit (integer, in Kbps).`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"limit_up": schema.Int64Attribute{
									MarkdownDescription: `The maximum upload limit (integer, in Kbps).`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
							},
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `The id of the pricing plan to update.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"price": schema.Float64Attribute{
							MarkdownDescription: `The price of the billing plan.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Float64{
								float64planmodifier.UseStateForUnknown(),
							},
						},
						"time_limit": schema.StringAttribute{
							MarkdownDescription: `The time limit of the pricing plan in minutes.
                                        Allowed values: [1 day,1 hour,1 week,30 days]`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"1 day",
									"1 hour",
									"1 week",
									"30 days",
								),
							},
						},
					},
				},
			},
		},
	}
}

func (r *NetworksWirelessBillingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessBillingRs

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
	// Has Paths
	vvNetworkID := data.NetworkID.ValueString()
	//Has Item and not has items

	if vvNetworkID != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Wireless.GetNetworkWirelessBilling(vvNetworkID)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksWirelessBilling  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksWirelessBilling only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessBilling(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessBilling",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessBilling",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Wireless.GetNetworkWirelessBilling(vvNetworkID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessBilling",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessBilling",
			err.Error(),
		)
		return
	}

	data = ResponseWirelessGetNetworkWirelessBillingItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *NetworksWirelessBillingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessBillingRs

	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessBilling(vvNetworkID)
	if err != nil || restyRespGet == nil {
		if restyRespGet != nil {
			if restyRespGet.StatusCode() == 404 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessBilling",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessBilling",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetNetworkWirelessBillingItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksWirelessBillingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksWirelessBillingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksWirelessBillingRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessBilling(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessBilling",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessBilling",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessBillingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksWirelessBilling", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessBillingRs struct {
	NetworkID types.String                                        `tfsdk:"network_id"`
	Currency  types.String                                        `tfsdk:"currency"`
	Plans     *[]ResponseWirelessGetNetworkWirelessBillingPlansRs `tfsdk:"plans"`
}

type ResponseWirelessGetNetworkWirelessBillingPlansRs struct {
	BandwidthLimits *ResponseWirelessGetNetworkWirelessBillingPlansBandwidthLimitsRs `tfsdk:"bandwidth_limits"`
	ID              types.String                                                     `tfsdk:"id"`
	Price           types.Float64                                                    `tfsdk:"price"`
	TimeLimit       types.String                                                     `tfsdk:"time_limit"`
}

type ResponseWirelessGetNetworkWirelessBillingPlansBandwidthLimitsRs struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

// FromBody
func (r *NetworksWirelessBillingRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessBilling {
	emptyString := ""
	currency := new(string)
	if !r.Currency.IsUnknown() && !r.Currency.IsNull() {
		*currency = r.Currency.ValueString()
	} else {
		currency = &emptyString
	}
	var requestWirelessUpdateNetworkWirelessBillingPlans []merakigosdk.RequestWirelessUpdateNetworkWirelessBillingPlans

	if r.Plans != nil {
		for _, rItem1 := range *r.Plans {
			var requestWirelessUpdateNetworkWirelessBillingPlansBandwidthLimits *merakigosdk.RequestWirelessUpdateNetworkWirelessBillingPlansBandwidthLimits

			if rItem1.BandwidthLimits != nil {
				limitDown := func() *int64 {
					if !rItem1.BandwidthLimits.LimitDown.IsUnknown() && !rItem1.BandwidthLimits.LimitDown.IsNull() {
						return rItem1.BandwidthLimits.LimitDown.ValueInt64Pointer()
					}
					return nil
				}()
				limitUp := func() *int64 {
					if !rItem1.BandwidthLimits.LimitUp.IsUnknown() && !rItem1.BandwidthLimits.LimitUp.IsNull() {
						return rItem1.BandwidthLimits.LimitUp.ValueInt64Pointer()
					}
					return nil
				}()
				requestWirelessUpdateNetworkWirelessBillingPlansBandwidthLimits = &merakigosdk.RequestWirelessUpdateNetworkWirelessBillingPlansBandwidthLimits{
					LimitDown: int64ToIntPointer(limitDown),
					LimitUp:   int64ToIntPointer(limitUp),
				}
				//[debug] Is Array: False
			}
			id := rItem1.ID.ValueString()
			price := func() *float64 {
				if !rItem1.Price.IsUnknown() && !rItem1.Price.IsNull() {
					return rItem1.Price.ValueFloat64Pointer()
				}
				return nil
			}()
			timeLimit := rItem1.TimeLimit.ValueString()
			requestWirelessUpdateNetworkWirelessBillingPlans = append(requestWirelessUpdateNetworkWirelessBillingPlans, merakigosdk.RequestWirelessUpdateNetworkWirelessBillingPlans{
				BandwidthLimits: requestWirelessUpdateNetworkWirelessBillingPlansBandwidthLimits,
				ID:              id,
				Price:           price,
				TimeLimit:       timeLimit,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessBilling{
		Currency: *currency,
		Plans: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessBillingPlans {
			if len(requestWirelessUpdateNetworkWirelessBillingPlans) > 0 {
				return &requestWirelessUpdateNetworkWirelessBillingPlans
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessBillingItemToBodyRs(state NetworksWirelessBillingRs, response *merakigosdk.ResponseWirelessGetNetworkWirelessBilling, is_read bool) NetworksWirelessBillingRs {
	itemState := NetworksWirelessBillingRs{
		Currency: types.StringValue(response.Currency),
		Plans: func() *[]ResponseWirelessGetNetworkWirelessBillingPlansRs {
			if response.Plans != nil {
				result := make([]ResponseWirelessGetNetworkWirelessBillingPlansRs, len(*response.Plans))
				for i, plans := range *response.Plans {
					result[i] = ResponseWirelessGetNetworkWirelessBillingPlansRs{
						BandwidthLimits: func() *ResponseWirelessGetNetworkWirelessBillingPlansBandwidthLimitsRs {
							if plans.BandwidthLimits != nil {
								return &ResponseWirelessGetNetworkWirelessBillingPlansBandwidthLimitsRs{
									LimitDown: func() types.Int64 {
										if plans.BandwidthLimits.LimitDown != nil {
											return types.Int64Value(int64(*plans.BandwidthLimits.LimitDown))
										}
										return types.Int64{}
									}(),
									LimitUp: func() types.Int64 {
										if plans.BandwidthLimits.LimitUp != nil {
											return types.Int64Value(int64(*plans.BandwidthLimits.LimitUp))
										}
										return types.Int64{}
									}(),
								}
							}
							return nil
						}(),
						ID: types.StringValue(plans.ID),
						Price: func() types.Float64 {
							if plans.Price != nil {
								return types.Float64Value(float64(*plans.Price))
							}
							return types.Float64{}
						}(),
						TimeLimit: types.StringValue(plans.TimeLimit),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessBillingRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessBillingRs)
}
