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
	"fmt"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWirelessSSIDsIDentityPsksResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessSSIDsIDentityPsksResource{}
)

func NewNetworksWirelessSSIDsIDentityPsksResource() resource.Resource {
	return &NetworksWirelessSSIDsIDentityPsksResource{}
}

type NetworksWirelessSSIDsIDentityPsksResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessSSIDsIDentityPsksResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessSSIDsIDentityPsksResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_identity_psks"
}

func (r *NetworksWirelessSSIDsIDentityPsksResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"email": schema.StringAttribute{
				MarkdownDescription: `The email associated with the System's Manager User`,
				Computed:            true,
			},
			"expires_at": schema.StringAttribute{
				MarkdownDescription: `Timestamp for when the Identity PSK expires, or 'null' to never expire`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"group_policy_id": schema.StringAttribute{
				MarkdownDescription: `The group policy to be applied to clients`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: `The unique identifier of the Identity PSK`,
				Computed:            true,
			},
			"identity_psk_id": schema.StringAttribute{
				MarkdownDescription: `identityPskId path parameter. Identity psk ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name of the Identity PSK`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: `number path parameter.`,
				Required:            true,
			},
			"passphrase": schema.StringAttribute{
				MarkdownDescription: `The passphrase for client authentication`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"wifi_personal_network_id": schema.StringAttribute{
				MarkdownDescription: `The WiFi Personal Network unique identifier`,
				Computed:            true,
			},
		},
	}
}

//path params to set ['identityPskId']

func (r *NetworksWirelessSSIDsIDentityPsksResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessSSIDsIDentityPsksRs

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
	// network_id
	vvNumber := data.Number.ValueString()
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDIDentityPsks(vvNetworkID, vvNumber)
	//Have Create
	if err != nil || restyResp1 == nil {
		if restyResp1.StatusCode() != 404 {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDIDentityPsks",
				err.Error(),
			)
			return
		}
	}
	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvIDentityPskID, ok := result2["ID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter IDentityPskID",
					"Error",
				)
				return
			}
			data.IDentityPskID = types.StringValue(vvIDentityPskID)
			r.client.Wireless.UpdateNetworkWirelessSSIDIDentityPsk(vvNetworkID, vvNumber, vvIDentityPskID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Wireless.GetNetworkWirelessSSIDIDentityPsk(vvNetworkID, vvNumber, vvIDentityPskID)
			if responseVerifyItem2 != nil {
				data = ResponseWirelessGetNetworkWirelessSSIDIDentityPskItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	restyResp2, err := r.client.Wireless.CreateNetworkWirelessSSIDIDentityPsk(vvNetworkID, vvNumber, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkWirelessSSIDIDentityPsk",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkWirelessSSIDIDentityPsk",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDIDentityPsks(vvNetworkID, vvNumber)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDIDentityPsks",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDIDentityPsks",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvIDentityPskID, ok := result2["ID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter IDentityPskID",
				"Error",
			)
			return
		}
		data.IDentityPskID = types.StringValue(vvIDentityPskID)
		responseVerifyItem2, restyRespGet, err := r.client.Wireless.GetNetworkWirelessSSIDIDentityPsk(vvNetworkID, vvNumber, vvIDentityPskID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseWirelessGetNetworkWirelessSSIDIDentityPskItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkWirelessSSIDIDentityPsk",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDIDentityPsk",
				err.Error(),
			)
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Error in result.",
			"Error in result.",
		)
		return
	}
}

func (r *NetworksWirelessSSIDsIDentityPsksResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessSSIDsIDentityPsksRs

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
	// network_id
	vvNumber := data.Number.ValueString()
	// number
	vvIDentityPskID := data.ID.ValueString()
	// identity_psk_id
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessSSIDIDentityPsk(vvNetworkID, vvNumber, vvIDentityPskID)
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
				"Failure when executing GetNetworkWirelessSSIDIDentityPsk",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDIDentityPsk",
			err.Error(),
		)
		return
	}

	data = ResponseWirelessGetNetworkWirelessSSIDIDentityPskItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessSSIDsIDentityPsksResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 3 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("number"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("identity_psk_id"), idParts[2])...)
}

func (r *NetworksWirelessSSIDsIDentityPsksResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksWirelessSSIDsIDentityPsksRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvNumber := data.Number.ValueString()
	vvIDentityPskID := data.ID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDIDentityPsk(vvNetworkID, vvNumber, vvIDentityPskID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDIDentityPsk",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDIDentityPsk",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessSSIDsIDentityPsksResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksWirelessSSIDsIDentityPsksRs
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &state, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}

	vvNetworkID := state.NetworkID.ValueString()
	vvNumber := state.Number.ValueString()
	vvIDentityPskID := state.ID.ValueString()
	_, err := r.client.Wireless.DeleteNetworkWirelessSSIDIDentityPsk(vvNetworkID, vvNumber, vvIDentityPskID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkWirelessSSIDIDentityPsk", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksWirelessSSIDsIDentityPsksRs struct {
	NetworkID             types.String `tfsdk:"network_id"`
	Number                types.String `tfsdk:"number"`
	IDentityPskID         types.String `tfsdk:"identity_psk_id"`
	Email                 types.String `tfsdk:"email"`
	ExpiresAt             types.String `tfsdk:"expires_at"`
	GroupPolicyID         types.String `tfsdk:"group_policy_id"`
	ID                    types.String `tfsdk:"id"`
	Name                  types.String `tfsdk:"name"`
	Passphrase            types.String `tfsdk:"passphrase"`
	WifiPersonalNetworkID types.String `tfsdk:"wifi_personal_network_id"`
}

// FromBody
func (r *NetworksWirelessSSIDsIDentityPsksRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestWirelessCreateNetworkWirelessSSIDIDentityPsk {
	emptyString := ""
	expiresAt := new(string)
	if !r.ExpiresAt.IsUnknown() && !r.ExpiresAt.IsNull() {
		*expiresAt = r.ExpiresAt.ValueString()
	} else {
		expiresAt = &emptyString
	}
	groupPolicyID := new(string)
	if !r.GroupPolicyID.IsUnknown() && !r.GroupPolicyID.IsNull() {
		*groupPolicyID = r.GroupPolicyID.ValueString()
	} else {
		groupPolicyID = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	passphrase := new(string)
	if !r.Passphrase.IsUnknown() && !r.Passphrase.IsNull() {
		*passphrase = r.Passphrase.ValueString()
	} else {
		passphrase = &emptyString
	}
	out := merakigosdk.RequestWirelessCreateNetworkWirelessSSIDIDentityPsk{
		ExpiresAt:     *expiresAt,
		GroupPolicyID: *groupPolicyID,
		Name:          *name,
		Passphrase:    *passphrase,
	}
	return &out
}
func (r *NetworksWirelessSSIDsIDentityPsksRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDIDentityPsk {
	emptyString := ""
	expiresAt := new(string)
	if !r.ExpiresAt.IsUnknown() && !r.ExpiresAt.IsNull() {
		*expiresAt = r.ExpiresAt.ValueString()
	} else {
		expiresAt = &emptyString
	}
	groupPolicyID := new(string)
	if !r.GroupPolicyID.IsUnknown() && !r.GroupPolicyID.IsNull() {
		*groupPolicyID = r.GroupPolicyID.ValueString()
	} else {
		groupPolicyID = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	passphrase := new(string)
	if !r.Passphrase.IsUnknown() && !r.Passphrase.IsNull() {
		*passphrase = r.Passphrase.ValueString()
	} else {
		passphrase = &emptyString
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDIDentityPsk{
		ExpiresAt:     *expiresAt,
		GroupPolicyID: *groupPolicyID,
		Name:          *name,
		Passphrase:    *passphrase,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessSSIDIDentityPskItemToBodyRs(state NetworksWirelessSSIDsIDentityPsksRs, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDIDentityPsk, is_read bool) NetworksWirelessSSIDsIDentityPsksRs {
	itemState := NetworksWirelessSSIDsIDentityPsksRs{
		Email:                 types.StringValue(response.Email),
		ExpiresAt:             types.StringValue(response.ExpiresAt),
		GroupPolicyID:         types.StringValue(response.GroupPolicyID),
		ID:                    types.StringValue(response.ID),
		Name:                  types.StringValue(response.Name),
		Passphrase:            types.StringValue(response.Passphrase),
		WifiPersonalNetworkID: types.StringValue(response.WifiPersonalNetworkID),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessSSIDsIDentityPsksRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessSSIDsIDentityPsksRs)
}
