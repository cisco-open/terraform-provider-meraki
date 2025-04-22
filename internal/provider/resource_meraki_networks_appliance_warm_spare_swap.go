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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceWarmSpareSwapResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceWarmSpareSwapResource{}
)

func NewNetworksApplianceWarmSpareSwapResource() resource.Resource {
	return &NetworksApplianceWarmSpareSwapResource{}
}

type NetworksApplianceWarmSpareSwapResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceWarmSpareSwapResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceWarmSpareSwapResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_warm_spare_swap"
}

// resourceAction
func (r *NetworksApplianceWarmSpareSwapResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Is the warm spare enabled`,
						Computed:            true,
					},
					"primary_serial": schema.StringAttribute{
						MarkdownDescription: `Serial number of the primary appliance`,
						Computed:            true,
					},
					"spare_serial": schema.StringAttribute{
						MarkdownDescription: `Serial number of the warm spare appliance`,
						Computed:            true,
					},
					"uplink_mode": schema.StringAttribute{
						MarkdownDescription: `Uplink mode, either virtual or public`,
						Computed:            true,
					},
					"wan1": schema.SingleNestedAttribute{
						MarkdownDescription: `WAN 1 IP and subnet`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"ip": schema.StringAttribute{
								MarkdownDescription: `IP address used for WAN 1`,
								Computed:            true,
							},
							"subnet": schema.StringAttribute{
								MarkdownDescription: `Subnet used for WAN 1`,
								Computed:            true,
							},
						},
					},
					"wan2": schema.SingleNestedAttribute{
						MarkdownDescription: `WAN 2 IP and subnet`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"ip": schema.StringAttribute{
								MarkdownDescription: `IP address used for WAN 2`,
								Computed:            true,
							},
							"subnet": schema.StringAttribute{
								MarkdownDescription: `Subnet used for WAN 2`,
								Computed:            true,
							},
						},
					},
				},
			},
		},
	}
}
func (r *NetworksApplianceWarmSpareSwapResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceWarmSpareSwap

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
	response, restyResp1, err := r.client.Appliance.SwapNetworkApplianceWarmSpare(vvNetworkID)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing SwapNetworkApplianceWarmSpare",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing SwapNetworkApplianceWarmSpare",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseApplianceSwapNetworkApplianceWarmSpareItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceWarmSpareSwapResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksApplianceWarmSpareSwapResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksApplianceWarmSpareSwapResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceWarmSpareSwap struct {
	NetworkID types.String                                    `tfsdk:"network_id"`
	Item      *ResponseApplianceSwapNetworkApplianceWarmSpare `tfsdk:"item"`
}

type ResponseApplianceSwapNetworkApplianceWarmSpare struct {
	Enabled       types.Bool                                          `tfsdk:"enabled"`
	PrimarySerial types.String                                        `tfsdk:"primary_serial"`
	SpareSerial   types.String                                        `tfsdk:"spare_serial"`
	UplinkMode    types.String                                        `tfsdk:"uplink_mode"`
	Wan1          *ResponseApplianceSwapNetworkApplianceWarmSpareWan1 `tfsdk:"wan1"`
	Wan2          *ResponseApplianceSwapNetworkApplianceWarmSpareWan2 `tfsdk:"wan2"`
}

type ResponseApplianceSwapNetworkApplianceWarmSpareWan1 struct {
	IP     types.String `tfsdk:"ip"`
	Subnet types.String `tfsdk:"subnet"`
}

type ResponseApplianceSwapNetworkApplianceWarmSpareWan2 struct {
	IP     types.String `tfsdk:"ip"`
	Subnet types.String `tfsdk:"subnet"`
}

type RequestApplianceSwapNetworkApplianceWarmSpareRs interface{}

// FromBody
// ToBody
func ResponseApplianceSwapNetworkApplianceWarmSpareItemToBody(state NetworksApplianceWarmSpareSwap, response *merakigosdk.ResponseApplianceSwapNetworkApplianceWarmSpare) NetworksApplianceWarmSpareSwap {
	itemState := ResponseApplianceSwapNetworkApplianceWarmSpare{
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		PrimarySerial: types.StringValue(response.PrimarySerial),
		SpareSerial:   types.StringValue(response.SpareSerial),
		UplinkMode:    types.StringValue(response.UplinkMode),
		Wan1: func() *ResponseApplianceSwapNetworkApplianceWarmSpareWan1 {
			if response.Wan1 != nil {
				return &ResponseApplianceSwapNetworkApplianceWarmSpareWan1{
					IP:     types.StringValue(response.Wan1.IP),
					Subnet: types.StringValue(response.Wan1.Subnet),
				}
			}
			return nil
		}(),
		Wan2: func() *ResponseApplianceSwapNetworkApplianceWarmSpareWan2 {
			if response.Wan2 != nil {
				return &ResponseApplianceSwapNetworkApplianceWarmSpareWan2{
					IP:     types.StringValue(response.Wan2.IP),
					Subnet: types.StringValue(response.Wan2.Subnet),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
