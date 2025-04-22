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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceTrafficShapingUplinkBandwidthResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceTrafficShapingUplinkBandwidthResource{}
)

func NewNetworksApplianceTrafficShapingUplinkBandwidthResource() resource.Resource {
	return &NetworksApplianceTrafficShapingUplinkBandwidthResource{}
}

type NetworksApplianceTrafficShapingUplinkBandwidthResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceTrafficShapingUplinkBandwidthResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceTrafficShapingUplinkBandwidthResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_traffic_shaping_uplink_bandwidth"
}

func (r *NetworksApplianceTrafficShapingUplinkBandwidthResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bandwidth_limits": schema.SingleNestedAttribute{
				MarkdownDescription: `A hash uplink keys and their configured settings for the Appliance`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"cellular": schema.SingleNestedAttribute{
						MarkdownDescription: `uplink cellular configured limits [optional]`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"limit_down": schema.Int64Attribute{
								MarkdownDescription: `configured DOWN limit for the uplink (in Kbps).  Null indicated unlimited`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
							},
							"limit_up": schema.Int64Attribute{
								MarkdownDescription: `configured UP limit for the uplink (in Kbps).  Null indicated unlimited`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"wan1": schema.SingleNestedAttribute{
						MarkdownDescription: `uplink wan1 configured limits [optional]`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"limit_down": schema.Int64Attribute{
								MarkdownDescription: `configured DOWN limit for the uplink (in Kbps).  Null indicated unlimited`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
							},
							"limit_up": schema.Int64Attribute{
								MarkdownDescription: `configured UP limit for the uplink (in Kbps).  Null indicated unlimited`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"wan2": schema.SingleNestedAttribute{
						MarkdownDescription: `uplink wan2 configured limits [optional]`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"limit_down": schema.Int64Attribute{
								MarkdownDescription: `configured DOWN limit for the uplink (in Kbps).  Null indicated unlimited`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
							},
							"limit_up": schema.Int64Attribute{
								MarkdownDescription: `configured UP limit for the uplink (in Kbps).  Null indicated unlimited`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
							},
						},
					},
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
		},
	}
}

func (r *NetworksApplianceTrafficShapingUplinkBandwidthResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceTrafficShapingUplinkBandwidthRs

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
		responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceTrafficShapingUplinkBandwidth(vvNetworkID)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksApplianceTrafficShapingUplinkBandwidth  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksApplianceTrafficShapingUplinkBandwidth only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceTrafficShapingUplinkBandwidth(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceTrafficShapingUplinkBandwidth",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceTrafficShapingUplinkBandwidth",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceTrafficShapingUplinkBandwidth(vvNetworkID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceTrafficShapingUplinkBandwidth",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceTrafficShapingUplinkBandwidth",
			err.Error(),
		)
		return
	}

	data = ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *NetworksApplianceTrafficShapingUplinkBandwidthResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceTrafficShapingUplinkBandwidthRs

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
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceTrafficShapingUplinkBandwidth(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceTrafficShapingUplinkBandwidth",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceTrafficShapingUplinkBandwidth",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksApplianceTrafficShapingUplinkBandwidthResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceTrafficShapingUplinkBandwidthResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceTrafficShapingUplinkBandwidthRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceTrafficShapingUplinkBandwidth(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceTrafficShapingUplinkBandwidth",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceTrafficShapingUplinkBandwidth",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceTrafficShapingUplinkBandwidthResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceTrafficShapingUplinkBandwidth", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceTrafficShapingUplinkBandwidthRs struct {
	NetworkID       types.String                                                                        `tfsdk:"network_id"`
	BandwidthLimits *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsRs `tfsdk:"bandwidth_limits"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsRs struct {
	Cellular *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellularRs `tfsdk:"cellular"`
	Wan1     *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1Rs     `tfsdk:"wan1"`
	Wan2     *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2Rs     `tfsdk:"wan2"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellularRs struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1Rs struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2Rs struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

// FromBody
func (r *NetworksApplianceTrafficShapingUplinkBandwidthRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidth {
	var requestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimits *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimits

	if r.BandwidthLimits != nil {
		var requestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular

		if r.BandwidthLimits.Cellular != nil {
			limitDown := func() *int64 {
				if !r.BandwidthLimits.Cellular.LimitDown.IsUnknown() && !r.BandwidthLimits.Cellular.LimitDown.IsNull() {
					return r.BandwidthLimits.Cellular.LimitDown.ValueInt64Pointer()
				}
				return nil
			}()
			limitUp := func() *int64 {
				if !r.BandwidthLimits.Cellular.LimitUp.IsUnknown() && !r.BandwidthLimits.Cellular.LimitUp.IsNull() {
					return r.BandwidthLimits.Cellular.LimitUp.ValueInt64Pointer()
				}
				return nil
			}()
			requestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular = &merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular{
				LimitDown: int64ToIntPointer(limitDown),
				LimitUp:   int64ToIntPointer(limitUp),
			}
			//[debug] Is Array: False
		}
		var requestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1 *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1

		if r.BandwidthLimits.Wan1 != nil {
			limitDown := func() *int64 {
				if !r.BandwidthLimits.Wan1.LimitDown.IsUnknown() && !r.BandwidthLimits.Wan1.LimitDown.IsNull() {
					return r.BandwidthLimits.Wan1.LimitDown.ValueInt64Pointer()
				}
				return nil
			}()
			limitUp := func() *int64 {
				if !r.BandwidthLimits.Wan1.LimitUp.IsUnknown() && !r.BandwidthLimits.Wan1.LimitUp.IsNull() {
					return r.BandwidthLimits.Wan1.LimitUp.ValueInt64Pointer()
				}
				return nil
			}()
			requestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1 = &merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1{
				LimitDown: int64ToIntPointer(limitDown),
				LimitUp:   int64ToIntPointer(limitUp),
			}
			//[debug] Is Array: False
		}
		var requestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2 *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2

		if r.BandwidthLimits.Wan2 != nil {
			limitDown := func() *int64 {
				if !r.BandwidthLimits.Wan2.LimitDown.IsUnknown() && !r.BandwidthLimits.Wan2.LimitDown.IsNull() {
					return r.BandwidthLimits.Wan2.LimitDown.ValueInt64Pointer()
				}
				return nil
			}()
			limitUp := func() *int64 {
				if !r.BandwidthLimits.Wan2.LimitUp.IsUnknown() && !r.BandwidthLimits.Wan2.LimitUp.IsNull() {
					return r.BandwidthLimits.Wan2.LimitUp.ValueInt64Pointer()
				}
				return nil
			}()
			requestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2 = &merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2{
				LimitDown: int64ToIntPointer(limitDown),
				LimitUp:   int64ToIntPointer(limitUp),
			}
			//[debug] Is Array: False
		}
		requestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimits = &merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimits{
			Cellular: requestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellular,
			Wan1:     requestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1,
			Wan2:     requestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2,
		}
		//[debug] Is Array: False
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidth{
		BandwidthLimits: requestApplianceUpdateNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimits,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthItemToBodyRs(state NetworksApplianceTrafficShapingUplinkBandwidthRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidth, is_read bool) NetworksApplianceTrafficShapingUplinkBandwidthRs {
	itemState := NetworksApplianceTrafficShapingUplinkBandwidthRs{
		BandwidthLimits: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsRs {
			if response.BandwidthLimits != nil {
				return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsRs{
					Cellular: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellularRs {
						if response.BandwidthLimits.Cellular != nil {
							return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsCellularRs{
								LimitDown: func() types.Int64 {
									if response.BandwidthLimits.Cellular.LimitDown != nil {
										return types.Int64Value(int64(*response.BandwidthLimits.Cellular.LimitDown))
									}
									return types.Int64{}
								}(),
								LimitUp: func() types.Int64 {
									if response.BandwidthLimits.Cellular.LimitUp != nil {
										return types.Int64Value(int64(*response.BandwidthLimits.Cellular.LimitUp))
									}
									return types.Int64{}
								}(),
							}
						}
						return nil
					}(),
					Wan1: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1Rs {
						if response.BandwidthLimits.Wan1 != nil {
							return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan1Rs{
								LimitDown: func() types.Int64 {
									if response.BandwidthLimits.Wan1.LimitDown != nil {
										return types.Int64Value(int64(*response.BandwidthLimits.Wan1.LimitDown))
									}
									return types.Int64{}
								}(),
								LimitUp: func() types.Int64 {
									if response.BandwidthLimits.Wan1.LimitUp != nil {
										return types.Int64Value(int64(*response.BandwidthLimits.Wan1.LimitUp))
									}
									return types.Int64{}
								}(),
							}
						}
						return nil
					}(),
					Wan2: func() *ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2Rs {
						if response.BandwidthLimits.Wan2 != nil {
							return &ResponseApplianceGetNetworkApplianceTrafficShapingUplinkBandwidthBandwidthLimitsWan2Rs{
								LimitDown: func() types.Int64 {
									if response.BandwidthLimits.Wan2.LimitDown != nil {
										return types.Int64Value(int64(*response.BandwidthLimits.Wan2.LimitDown))
									}
									return types.Int64{}
								}(),
								LimitUp: func() types.Int64 {
									if response.BandwidthLimits.Wan2.LimitUp != nil {
										return types.Int64Value(int64(*response.BandwidthLimits.Wan2.LimitUp))
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
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceTrafficShapingUplinkBandwidthRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceTrafficShapingUplinkBandwidthRs)
}
