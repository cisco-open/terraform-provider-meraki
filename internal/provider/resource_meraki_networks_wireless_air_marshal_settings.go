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
	_ resource.Resource              = &NetworksWirelessAirMarshalSettingsResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessAirMarshalSettingsResource{}
)

func NewNetworksWirelessAirMarshalSettingsResource() resource.Resource {
	return &NetworksWirelessAirMarshalSettingsResource{}
}

type NetworksWirelessAirMarshalSettingsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessAirMarshalSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessAirMarshalSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_air_marshal_settings"
}

// resourceAction
func (r *NetworksWirelessAirMarshalSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

					"default_policy": schema.StringAttribute{
						MarkdownDescription: `Indicates whether or not clients are allowed to       connect to rogue SSIDs. (blocked by default)`,
						Computed:            true,
					},
					"network_id": schema.StringAttribute{
						MarkdownDescription: `The network ID`,
						Computed:            true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"default_policy": schema.StringAttribute{
						MarkdownDescription: `Allows clients to access rogue networks. Blocked by default.
                                        Allowed values: [allow,block]`,
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
				},
			},
		},
	}
}
func (r *NetworksWirelessAirMarshalSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessAirMarshalSettings

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
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp1, err := r.client.Wireless.UpdateNetworkWirelessAirMarshalSettings(vvNetworkID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessAirMarshalSettings",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessAirMarshalSettings",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseWirelessUpdateNetworkWirelessAirMarshalSettingsItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessAirMarshalSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksWirelessAirMarshalSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksWirelessAirMarshalSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessAirMarshalSettings struct {
	NetworkID  types.String                                              `tfsdk:"network_id"`
	Item       *ResponseWirelessUpdateNetworkWirelessAirMarshalSettings  `tfsdk:"item"`
	Parameters *RequestWirelessUpdateNetworkWirelessAirMarshalSettingsRs `tfsdk:"parameters"`
}

type ResponseWirelessUpdateNetworkWirelessAirMarshalSettings struct {
	DefaultPolicy types.String `tfsdk:"default_policy"`
	NetworkID     types.String `tfsdk:"network_id"`
}

type RequestWirelessUpdateNetworkWirelessAirMarshalSettingsRs struct {
	DefaultPolicy types.String `tfsdk:"default_policy"`
}

// FromBody
func (r *NetworksWirelessAirMarshalSettings) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessAirMarshalSettings {
	emptyString := ""
	re := *r.Parameters
	defaultPolicy := new(string)
	if !re.DefaultPolicy.IsUnknown() && !re.DefaultPolicy.IsNull() {
		*defaultPolicy = re.DefaultPolicy.ValueString()
	} else {
		defaultPolicy = &emptyString
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessAirMarshalSettings{
		DefaultPolicy: *defaultPolicy,
	}
	return &out
}

// ToBody
func ResponseWirelessUpdateNetworkWirelessAirMarshalSettingsItemToBody(state NetworksWirelessAirMarshalSettings, response *merakigosdk.ResponseWirelessUpdateNetworkWirelessAirMarshalSettings) NetworksWirelessAirMarshalSettings {
	itemState := ResponseWirelessUpdateNetworkWirelessAirMarshalSettings{
		DefaultPolicy: types.StringValue(response.DefaultPolicy),
		NetworkID:     types.StringValue(response.NetworkID),
	}
	state.Item = &itemState
	return state
}
