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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSmDevicesInstallAppsResource{}
	_ resource.ResourceWithConfigure = &NetworksSmDevicesInstallAppsResource{}
)

func NewNetworksSmDevicesInstallAppsResource() resource.Resource {
	return &NetworksSmDevicesInstallAppsResource{}
}

type NetworksSmDevicesInstallAppsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSmDevicesInstallAppsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSmDevicesInstallAppsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices_install_apps"
}

// resourceAction
func (r *NetworksSmDevicesInstallAppsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
						MarkdownDescription: `ids of applications to be installed`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"force": schema.BoolAttribute{
						MarkdownDescription: `By default, installation of an app which is believed to already be present on the device will be skipped. If you'd like to force the installation of the app, set this parameter to true.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.RequiresReplace(),
						},
					},
				},
			},
		},
	}
}
func (r *NetworksSmDevicesInstallAppsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSmDevicesInstallApps

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
	restyResp1, err := r.client.Sm.InstallNetworkSmDeviceApps(vvNetworkID, vvDeviceID, dataRequest)
	if err != nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing InstallNetworkSmDeviceApps",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing InstallNetworkSmDeviceApps",
			err.Error(),
		)
		return
	}
	//Item
	// //entro aqui 2
	// data2 := ResponseSmInstallNetworkSmDeviceApps(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSmDevicesInstallAppsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmDevicesInstallAppsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmDevicesInstallAppsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSmDevicesInstallApps struct {
	NetworkID  types.String                           `tfsdk:"network_id"`
	DeviceID   types.String                           `tfsdk:"device_id"`
	Parameters *RequestSmInstallNetworkSmDeviceAppsRs `tfsdk:"parameters"`
}

type RequestSmInstallNetworkSmDeviceAppsRs struct {
	AppIDs types.Set  `tfsdk:"app_ids"`
	Force  types.Bool `tfsdk:"force"`
}

// FromBody
func (r *NetworksSmDevicesInstallApps) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSmInstallNetworkSmDeviceApps {
	re := *r.Parameters
	var appIDs []string = nil
	re.AppIDs.ElementsAs(ctx, &appIDs, false)
	force := new(bool)
	if !re.Force.IsUnknown() && !re.Force.IsNull() {
		*force = re.Force.ValueBool()
	} else {
		force = nil
	}
	out := merakigosdk.RequestSmInstallNetworkSmDeviceApps{
		AppIDs: appIDs,
		Force:  force,
	}
	return &out
}
