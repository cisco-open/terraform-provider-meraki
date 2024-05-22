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
	_ resource.Resource              = &NetworksSmDevicesShutdownResource{}
	_ resource.ResourceWithConfigure = &NetworksSmDevicesShutdownResource{}
)

func NewNetworksSmDevicesShutdownResource() resource.Resource {
	return &NetworksSmDevicesShutdownResource{}
}

type NetworksSmDevicesShutdownResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSmDevicesShutdownResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSmDevicesShutdownResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices_shutdown"
}

// resourceAction
func (r *NetworksSmDevicesShutdownResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

					"ids": schema.SetAttribute{
						MarkdownDescription: `The Meraki Ids of the set of endpoints.`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"ids": schema.SetAttribute{
						MarkdownDescription: `The ids of the endpoints to be shutdown.`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"scope": schema.SetAttribute{
						MarkdownDescription: `The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the endpoints to be shutdown.`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"serials": schema.SetAttribute{
						MarkdownDescription: `The serials of the endpoints to be shutdown.`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"wifi_macs": schema.SetAttribute{
						MarkdownDescription: `The wifiMacs of the endpoints to be shutdown.`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}
func (r *NetworksSmDevicesShutdownResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSmDevicesShutdown

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
	response, restyResp1, err := r.client.Sm.ShutdownNetworkSmDevices(vvNetworkID, dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing ShutdownNetworkSmDevices",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing ShutdownNetworkSmDevices",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseSmShutdownNetworkSmDevicesItemToBody(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSmDevicesShutdownResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmDevicesShutdownResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmDevicesShutdownResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSmDevicesShutdown struct {
	NetworkID  types.String                         `tfsdk:"network_id"`
	Item       *ResponseSmShutdownNetworkSmDevices  `tfsdk:"item"`
	Parameters *RequestSmShutdownNetworkSmDevicesRs `tfsdk:"parameters"`
}

type ResponseSmShutdownNetworkSmDevices struct {
	IDs types.Set `tfsdk:"ids"`
}

type RequestSmShutdownNetworkSmDevicesRs struct {
	IDs      types.Set `tfsdk:"ids"`
	Scope    types.Set `tfsdk:"scope"`
	Serials  types.Set `tfsdk:"serials"`
	WifiMacs types.Set `tfsdk:"wifi_macs"`
}

// FromBody
func (r *NetworksSmDevicesShutdown) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSmShutdownNetworkSmDevices {
	re := *r.Parameters
	var iDs []string = nil
	re.IDs.ElementsAs(ctx, &iDs, false)
	var scope []string = nil
	re.Scope.ElementsAs(ctx, &scope, false)
	var serials []string = nil
	re.Serials.ElementsAs(ctx, &serials, false)
	var wifiMacs []string = nil
	re.WifiMacs.ElementsAs(ctx, &wifiMacs, false)
	out := merakigosdk.RequestSmShutdownNetworkSmDevices{
		IDs:      iDs,
		Scope:    scope,
		Serials:  serials,
		WifiMacs: wifiMacs,
	}
	return &out
}

// ToBody
func ResponseSmShutdownNetworkSmDevicesItemToBody(state NetworksSmDevicesShutdown, response *merakigosdk.ResponseSmShutdownNetworkSmDevices) NetworksSmDevicesShutdown {
	itemState := ResponseSmShutdownNetworkSmDevices{
		IDs: StringSliceToSet(response.IDs),
	}
	state.Item = &itemState
	return state
}
