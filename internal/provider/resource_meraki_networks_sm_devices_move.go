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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSmDevicesMoveResource{}
	_ resource.ResourceWithConfigure = &NetworksSmDevicesMoveResource{}
)

func NewNetworksSmDevicesMoveResource() resource.Resource {
	return &NetworksSmDevicesMoveResource{}
}

type NetworksSmDevicesMoveResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSmDevicesMoveResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSmDevicesMoveResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices_move"
}

// resourceAction
func (r *NetworksSmDevicesMoveResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

					"ids": schema.ListAttribute{
						MarkdownDescription: `The Meraki Ids of the set of devices.`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"new_network": schema.StringAttribute{
						MarkdownDescription: `The network to which the devices was moved.`,
						Computed:            true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"ids": schema.ListAttribute{
						MarkdownDescription: `The ids of the devices to be moved.`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"new_network": schema.StringAttribute{
						MarkdownDescription: `The new network to which the devices will be moved.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"scope": schema.ListAttribute{
						MarkdownDescription: `The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be moved.`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"serials": schema.ListAttribute{
						MarkdownDescription: `The serials of the devices to be moved.`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"wifi_macs": schema.ListAttribute{
						MarkdownDescription: `The wifiMacs of the devices to be moved.`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}
func (r *NetworksSmDevicesMoveResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSmDevicesMove

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
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Sm.MoveNetworkSmDevices(vvNetworkID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing MoveNetworkSmDevices",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing MoveNetworkSmDevices",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseSmMoveNetworkSmDevicesItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSmDevicesMoveResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmDevicesMoveResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmDevicesMoveResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSmDevicesMove struct {
	NetworkID  types.String                     `tfsdk:"network_id"`
	Item       *ResponseSmMoveNetworkSmDevices  `tfsdk:"item"`
	Parameters *RequestSmMoveNetworkSmDevicesRs `tfsdk:"parameters"`
}

type ResponseSmMoveNetworkSmDevices struct {
	IDs        types.List   `tfsdk:"ids"`
	NewNetwork types.String `tfsdk:"new_network"`
}

type RequestSmMoveNetworkSmDevicesRs struct {
	IDs        types.Set    `tfsdk:"ids"`
	NewNetwork types.String `tfsdk:"new_network"`
	Scope      types.Set    `tfsdk:"scope"`
	Serials    types.Set    `tfsdk:"serials"`
	WifiMacs   types.Set    `tfsdk:"wifi_macs"`
}

// FromBody
func (r *NetworksSmDevicesMove) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSmMoveNetworkSmDevices {
	emptyString := ""
	re := *r.Parameters
	var iDs []string = nil
	re.IDs.ElementsAs(ctx, &iDs, false)
	newNetwork := new(string)
	if !re.NewNetwork.IsUnknown() && !re.NewNetwork.IsNull() {
		*newNetwork = re.NewNetwork.ValueString()
	} else {
		newNetwork = &emptyString
	}
	var scope []string = nil
	re.Scope.ElementsAs(ctx, &scope, false)
	var serials []string = nil
	re.Serials.ElementsAs(ctx, &serials, false)
	var wifiMacs []string = nil
	re.WifiMacs.ElementsAs(ctx, &wifiMacs, false)
	out := merakigosdk.RequestSmMoveNetworkSmDevices{
		IDs:        iDs,
		NewNetwork: *newNetwork,
		Scope:      scope,
		Serials:    serials,
		WifiMacs:   wifiMacs,
	}
	return &out
}

// ToBody
func ResponseSmMoveNetworkSmDevicesItemToBody(state NetworksSmDevicesMove, response *merakigosdk.ResponseSmMoveNetworkSmDevices) NetworksSmDevicesMove {
	itemState := ResponseSmMoveNetworkSmDevices{
		IDs:        StringSliceToList(response.IDs),
		NewNetwork: types.StringValue(response.NewNetwork),
	}
	state.Item = &itemState
	return state
}
