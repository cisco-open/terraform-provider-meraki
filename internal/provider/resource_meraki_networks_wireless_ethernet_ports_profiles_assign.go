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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWirelessEthernetPortsProfilesAssignResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessEthernetPortsProfilesAssignResource{}
)

func NewNetworksWirelessEthernetPortsProfilesAssignResource() resource.Resource {
	return &NetworksWirelessEthernetPortsProfilesAssignResource{}
}

type NetworksWirelessEthernetPortsProfilesAssignResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessEthernetPortsProfilesAssignResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessEthernetPortsProfilesAssignResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ethernet_ports_profiles_assign"
}

// resourceAction
func (r *NetworksWirelessEthernetPortsProfilesAssignResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

					"profile_id": schema.StringAttribute{
						MarkdownDescription: `AP profile ID`,
						Computed:            true,
					},
					"serials": schema.SetAttribute{
						MarkdownDescription: `List of updated AP serials`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"profile_id": schema.StringAttribute{
						MarkdownDescription: `AP profile ID`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"serials": schema.SetAttribute{
						MarkdownDescription: `List of AP serials`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}
func (r *NetworksWirelessEthernetPortsProfilesAssignResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessEthernetPortsProfilesAssign

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
	response, restyResp1, err := r.client.Wireless.AssignNetworkWirelessEthernetPortsProfiles(vvNetworkID, dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing AssignNetworkWirelessEthernetPortsProfiles",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing AssignNetworkWirelessEthernetPortsProfiles",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseWirelessAssignNetworkWirelessEthernetPortsProfilesItemToBody(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessEthernetPortsProfilesAssignResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksWirelessEthernetPortsProfilesAssignResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksWirelessEthernetPortsProfilesAssignResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessEthernetPortsProfilesAssign struct {
	NetworkID  types.String                                                 `tfsdk:"network_id"`
	Item       *ResponseWirelessAssignNetworkWirelessEthernetPortsProfiles  `tfsdk:"item"`
	Parameters *RequestWirelessAssignNetworkWirelessEthernetPortsProfilesRs `tfsdk:"parameters"`
}

type ResponseWirelessAssignNetworkWirelessEthernetPortsProfiles struct {
	ProfileID types.String `tfsdk:"profile_id"`
	Serials   types.Set    `tfsdk:"serials"`
}

type RequestWirelessAssignNetworkWirelessEthernetPortsProfilesRs struct {
	ProfileID types.String `tfsdk:"profile_id"`
	Serials   types.Set    `tfsdk:"serials"`
}

// FromBody
func (r *NetworksWirelessEthernetPortsProfilesAssign) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestWirelessAssignNetworkWirelessEthernetPortsProfiles {
	emptyString := ""
	re := *r.Parameters
	profileID := new(string)
	if !re.ProfileID.IsUnknown() && !re.ProfileID.IsNull() {
		*profileID = re.ProfileID.ValueString()
	} else {
		profileID = &emptyString
	}
	var serials []string = nil
	re.Serials.ElementsAs(ctx, &serials, false)
	out := merakigosdk.RequestWirelessAssignNetworkWirelessEthernetPortsProfiles{
		ProfileID: *profileID,
		Serials:   serials,
	}
	return &out
}

// ToBody
func ResponseWirelessAssignNetworkWirelessEthernetPortsProfilesItemToBody(state NetworksWirelessEthernetPortsProfilesAssign, response *merakigosdk.ResponseWirelessAssignNetworkWirelessEthernetPortsProfiles) NetworksWirelessEthernetPortsProfilesAssign {
	itemState := ResponseWirelessAssignNetworkWirelessEthernetPortsProfiles{
		ProfileID: types.StringValue(response.ProfileID),
		Serials:   StringSliceToSet(response.Serials),
	}
	state.Item = &itemState
	return state
}
