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
	_ resource.Resource              = &NetworksSmDevicesUninstallAppsResource{}
	_ resource.ResourceWithConfigure = &NetworksSmDevicesUninstallAppsResource{}
)

func NewNetworksSmDevicesUninstallAppsResource() resource.Resource {
	return &NetworksSmDevicesUninstallAppsResource{}
}

type NetworksSmDevicesUninstallAppsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSmDevicesUninstallAppsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSmDevicesUninstallAppsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices_uninstall_apps"
}

// resourceAction
func (r *NetworksSmDevicesUninstallAppsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"device_id": schema.StringAttribute{
				MarkdownDescription: `deviceId path parameter. Device ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"app_ids": schema.ListAttribute{
						MarkdownDescription: `ids of applications to be uninstalled`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}
func (r *NetworksSmDevicesUninstallAppsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSmDevicesUninstallApps

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
	vvDeviceID := data.DeviceID.ValueString()
	dataRequest := data.toSdkApiRequestCreate(ctx)
	restyResp1, err := r.client.Sm.UninstallNetworkSmDeviceApps(vvNetworkID, vvDeviceID, dataRequest)
	if err != nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UninstallNetworkSmDeviceApps",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UninstallNetworkSmDeviceApps",
			err.Error(),
		)
		return
	}
	//Item
	// //entro aqui 2
	// data2 := ResponseSmUninstallNetworkSmDeviceApps(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSmDevicesUninstallAppsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmDevicesUninstallAppsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmDevicesUninstallAppsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSmDevicesUninstallApps struct {
	NetworkID  types.String                             `tfsdk:"network_id"`
	DeviceID   types.String                             `tfsdk:"device_id"`
	Parameters *RequestSmUninstallNetworkSmDeviceAppsRs `tfsdk:"parameters"`
}

type RequestSmUninstallNetworkSmDeviceAppsRs struct {
	AppIDs types.Set `tfsdk:"app_ids"`
}

// FromBody
func (r *NetworksSmDevicesUninstallApps) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSmUninstallNetworkSmDeviceApps {
	re := *r.Parameters
	var appIDs []string = nil
	re.AppIDs.ElementsAs(ctx, &appIDs, false)
	out := merakigosdk.RequestSmUninstallNetworkSmDeviceApps{
		AppIDs: appIDs,
	}
	return &out
}
