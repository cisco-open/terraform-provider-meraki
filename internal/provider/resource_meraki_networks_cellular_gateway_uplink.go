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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

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
	_ resource.Resource              = &NetworksCellularGatewayUplinkResource{}
	_ resource.ResourceWithConfigure = &NetworksCellularGatewayUplinkResource{}
)

func NewNetworksCellularGatewayUplinkResource() resource.Resource {
	return &NetworksCellularGatewayUplinkResource{}
}

type NetworksCellularGatewayUplinkResource struct {
	client *merakigosdk.Client
}

func (r *NetworksCellularGatewayUplinkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksCellularGatewayUplinkResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_cellular_gateway_uplink"
}

func (r *NetworksCellularGatewayUplinkResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bandwidth_limits": schema.SingleNestedAttribute{
				MarkdownDescription: `The bandwidth settings for the 'cellular' uplink`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"limit_down": schema.Int64Attribute{
						MarkdownDescription: `The maximum download limit (integer, in Kbps). 'null' indicates no limit.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"limit_up": schema.Int64Attribute{
						MarkdownDescription: `The maximum upload limit (integer, in Kbps). 'null' indicates no limit.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
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

func (r *NetworksCellularGatewayUplinkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksCellularGatewayUplinkRs

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
	vvNetworkID := data.NetworkID.ValueString()
	//Item
	responseVerifyItem, restyResp1, err := r.client.CellularGateway.GetNetworkCellularGatewayUplink(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksCellularGatewayUplink only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksCellularGatewayUplink only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	_, restyResp2, err := r.client.CellularGateway.UpdateNetworkCellularGatewayUplink(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkCellularGatewayUplink",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkCellularGatewayUplink",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.CellularGateway.GetNetworkCellularGatewayUplink(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkCellularGatewayUplink",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkCellularGatewayUplink",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseCellularGatewayGetNetworkCellularGatewayUplinkItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksCellularGatewayUplinkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksCellularGatewayUplinkRs

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
	responseGet, restyRespGet, err := r.client.CellularGateway.GetNetworkCellularGatewayUplink(vvNetworkID)
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
				"Failure when executing GetNetworkCellularGatewayUplink",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkCellularGatewayUplink",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseCellularGatewayGetNetworkCellularGatewayUplinkItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksCellularGatewayUplinkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksCellularGatewayUplinkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksCellularGatewayUplinkRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	_, restyResp2, err := r.client.CellularGateway.UpdateNetworkCellularGatewayUplink(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkCellularGatewayUplink",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkCellularGatewayUplink",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksCellularGatewayUplinkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksCellularGatewayUplink", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksCellularGatewayUplinkRs struct {
	NetworkID       types.String                                                             `tfsdk:"network_id"`
	BandwidthLimits *ResponseCellularGatewayGetNetworkCellularGatewayUplinkBandwidthLimitsRs `tfsdk:"bandwidth_limits"`
}

type ResponseCellularGatewayGetNetworkCellularGatewayUplinkBandwidthLimitsRs struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

// FromBody
func (r *NetworksCellularGatewayUplinkRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestCellularGatewayUpdateNetworkCellularGatewayUplink {
	var requestCellularGatewayUpdateNetworkCellularGatewayUplinkBandwidthLimits *merakigosdk.RequestCellularGatewayUpdateNetworkCellularGatewayUplinkBandwidthLimits
	if r.BandwidthLimits != nil {
		limitDown := func() *int64 {
			if !r.BandwidthLimits.LimitDown.IsUnknown() && !r.BandwidthLimits.LimitDown.IsNull() {
				return r.BandwidthLimits.LimitDown.ValueInt64Pointer()
			}
			return nil
		}()
		limitUp := func() *int64 {
			if !r.BandwidthLimits.LimitUp.IsUnknown() && !r.BandwidthLimits.LimitUp.IsNull() {
				return r.BandwidthLimits.LimitUp.ValueInt64Pointer()
			}
			return nil
		}()
		requestCellularGatewayUpdateNetworkCellularGatewayUplinkBandwidthLimits = &merakigosdk.RequestCellularGatewayUpdateNetworkCellularGatewayUplinkBandwidthLimits{
			LimitDown: int64ToIntPointer(limitDown),
			LimitUp:   int64ToIntPointer(limitUp),
		}
	}
	out := merakigosdk.RequestCellularGatewayUpdateNetworkCellularGatewayUplink{
		BandwidthLimits: requestCellularGatewayUpdateNetworkCellularGatewayUplinkBandwidthLimits,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseCellularGatewayGetNetworkCellularGatewayUplinkItemToBodyRs(state NetworksCellularGatewayUplinkRs, response *merakigosdk.ResponseCellularGatewayGetNetworkCellularGatewayUplink, is_read bool) NetworksCellularGatewayUplinkRs {
	itemState := NetworksCellularGatewayUplinkRs{
		BandwidthLimits: func() *ResponseCellularGatewayGetNetworkCellularGatewayUplinkBandwidthLimitsRs {
			if response.BandwidthLimits != nil {
				return &ResponseCellularGatewayGetNetworkCellularGatewayUplinkBandwidthLimitsRs{
					LimitDown: func() types.Int64 {
						if response.BandwidthLimits.LimitDown != nil {
							return types.Int64Value(int64(*response.BandwidthLimits.LimitDown))
						}
						return types.Int64{}
					}(),
					LimitUp: func() types.Int64 {
						if response.BandwidthLimits.LimitUp != nil {
							return types.Int64Value(int64(*response.BandwidthLimits.LimitUp))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksCellularGatewayUplinkRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksCellularGatewayUplinkRs)
}
