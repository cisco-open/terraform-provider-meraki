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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSmDevicesLockResource{}
	_ resource.ResourceWithConfigure = &NetworksSmDevicesLockResource{}
)

func NewNetworksSmDevicesLockResource() resource.Resource {
	return &NetworksSmDevicesLockResource{}
}

type NetworksSmDevicesLockResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSmDevicesLockResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSmDevicesLockResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices_lock"
}

// resourceAction
func (r *NetworksSmDevicesLockResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"ids": schema.ListAttribute{
						MarkdownDescription: `The ids of the devices to be locked.`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"pin": schema.Int64Attribute{
						MarkdownDescription: `The pin number for locking macOS devices (a six digit number). Required only for macOS devices.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
					},
					"scope": schema.ListAttribute{
						MarkdownDescription: `The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be locked.`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"serials": schema.ListAttribute{
						MarkdownDescription: `The serials of the devices to be locked.`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"wifi_macs": schema.ListAttribute{
						MarkdownDescription: `The wifiMacs of the devices to be locked.`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}
func (r *NetworksSmDevicesLockResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSmDevicesLock

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
	response, restyResp1, err := r.client.Sm.LockNetworkSmDevices(vvNetworkID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing LockNetworkSmDevices",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing LockNetworkSmDevices",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseSmLockNetworkSmDevicesItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSmDevicesLockResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmDevicesLockResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmDevicesLockResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSmDevicesLock struct {
	NetworkID  types.String                     `tfsdk:"network_id"`
	Item       *ResponseSmLockNetworkSmDevices  `tfsdk:"item"`
	Parameters *RequestSmLockNetworkSmDevicesRs `tfsdk:"parameters"`
}

type ResponseSmLockNetworkSmDevices struct {
	IDs types.List `tfsdk:"ids"`
}

type RequestSmLockNetworkSmDevicesRs struct {
	IDs      types.Set   `tfsdk:"ids"`
	Pin      types.Int64 `tfsdk:"pin"`
	Scope    types.Set   `tfsdk:"scope"`
	Serials  types.Set   `tfsdk:"serials"`
	WifiMacs types.Set   `tfsdk:"wifi_macs"`
}

// FromBody
func (r *NetworksSmDevicesLock) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSmLockNetworkSmDevices {
	re := *r.Parameters
	var iDs []string = nil
	re.IDs.ElementsAs(ctx, &iDs, false)
	pin := new(int64)
	if !re.Pin.IsUnknown() && !re.Pin.IsNull() {
		*pin = re.Pin.ValueInt64()
	} else {
		pin = nil
	}
	var scope []string = nil
	re.Scope.ElementsAs(ctx, &scope, false)
	var serials []string = nil
	re.Serials.ElementsAs(ctx, &serials, false)
	var wifiMacs []string = nil
	re.WifiMacs.ElementsAs(ctx, &wifiMacs, false)
	out := merakigosdk.RequestSmLockNetworkSmDevices{
		IDs:      iDs,
		Pin:      int64ToIntPointer(pin),
		Scope:    scope,
		Serials:  serials,
		WifiMacs: wifiMacs,
	}
	return &out
}

// ToBody
func ResponseSmLockNetworkSmDevicesItemToBody(state NetworksSmDevicesLock, response *merakigosdk.ResponseSmLockNetworkSmDevices) NetworksSmDevicesLock {
	itemState := ResponseSmLockNetworkSmDevices{
		IDs: StringSliceToList(response.IDs),
	}
	state.Item = &itemState
	return state
}
