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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksCameraWirelessProfilesResource{}
	_ resource.ResourceWithConfigure = &NetworksCameraWirelessProfilesResource{}
)

func NewNetworksCameraWirelessProfilesResource() resource.Resource {
	return &NetworksCameraWirelessProfilesResource{}
}

type NetworksCameraWirelessProfilesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksCameraWirelessProfilesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksCameraWirelessProfilesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_camera_wireless_profiles"
}

func (r *NetworksCameraWirelessProfilesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"applied_device_count": schema.Int64Attribute{
				MarkdownDescription: `The count of the applied devices.`,
				Computed:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: `The ID of the camera wireless profile.`,
				Computed:            true,
			},
			"identity": schema.SingleNestedAttribute{
				MarkdownDescription: `The identity of the wireless profile. Required for creating wireless profiles in 8021x-radius auth mode.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"password": schema.StringAttribute{
						MarkdownDescription: `The password of the identity.`,
						Sensitive:           true,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"username": schema.StringAttribute{
						MarkdownDescription: `The username of the identity.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name of the camera wireless profile.`,
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
			"ssid": schema.SingleNestedAttribute{
				MarkdownDescription: `The details of the SSID config.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"auth_mode": schema.StringAttribute{
						MarkdownDescription: `The auth mode of the SSID. It can be set to ('psk', '8021x-radius').
                                        Allowed values: [8021x-radius,psk]`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"8021x-radius",
								"psk",
							),
						},
					},
					"encryption_mode": schema.StringAttribute{
						MarkdownDescription: `The encryption mode of the SSID.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the SSID.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"psk": schema.StringAttribute{
						MarkdownDescription: `The pre-shared key of the SSID, if mode is PSK`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"wireless_profile_id": schema.StringAttribute{
				MarkdownDescription: `wirelessProfileId path parameter. Wireless profile ID`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

//path params to set ['wirelessProfileId']

func (r *NetworksCameraWirelessProfilesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksCameraWirelessProfilesRs

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
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Camera.GetNetworkCameraWirelessProfiles(vvNetworkID)
	//Have Create
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkCameraWirelessProfiles",
					err.Error(),
				)
				return
			}
		}
	}
	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvWirelessProfileID, ok := result2["ID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter WirelessProfileID",
					"Fail Parsing WirelessProfileID",
				)
				return
			}
			r.client.Camera.UpdateNetworkCameraWirelessProfile(vvNetworkID, vvWirelessProfileID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Camera.GetNetworkCameraWirelessProfile(vvNetworkID, vvWirelessProfileID)
			if responseVerifyItem2 != nil {
				data = ResponseCameraGetNetworkCameraWirelessProfileItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Camera.CreateNetworkCameraWirelessProfile(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkCameraWirelessProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkCameraWirelessProfile",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Camera.GetNetworkCameraWirelessProfiles(vvNetworkID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkCameraWirelessProfiles",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkCameraWirelessProfiles",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvWirelessProfileID, ok := result2["ID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter WirelessProfileID",
				"Error",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Camera.GetNetworkCameraWirelessProfile(vvNetworkID, vvWirelessProfileID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseCameraGetNetworkCameraWirelessProfileItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkCameraWirelessProfile",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkCameraWirelessProfile",
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

func (r *NetworksCameraWirelessProfilesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksCameraWirelessProfilesRs

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
	vvWirelessProfileID := data.WirelessProfileID.ValueString()
	responseGet, restyRespGet, err := r.client.Camera.GetNetworkCameraWirelessProfile(vvNetworkID, vvWirelessProfileID)
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
				"Failure when executing GetNetworkCameraWirelessProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkCameraWirelessProfile",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseCameraGetNetworkCameraWirelessProfileItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksCameraWirelessProfilesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("wireless_profile_id"), idParts[1])...)
}

func (r *NetworksCameraWirelessProfilesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksCameraWirelessProfilesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvWirelessProfileID := data.WirelessProfileID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Camera.UpdateNetworkCameraWirelessProfile(vvNetworkID, vvWirelessProfileID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkCameraWirelessProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkCameraWirelessProfile",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksCameraWirelessProfilesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksCameraWirelessProfilesRs
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
	vvWirelessProfileID := state.WirelessProfileID.ValueString()
	_, err := r.client.Camera.DeleteNetworkCameraWirelessProfile(vvNetworkID, vvWirelessProfileID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkCameraWirelessProfile", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksCameraWirelessProfilesRs struct {
	NetworkID          types.String                                             `tfsdk:"network_id"`
	WirelessProfileID  types.String                                             `tfsdk:"wireless_profile_id"`
	AppliedDeviceCount types.Int64                                              `tfsdk:"applied_device_count"`
	ID                 types.String                                             `tfsdk:"id"`
	IDentity           *ResponseCameraGetNetworkCameraWirelessProfileIdentityRs `tfsdk:"identity"`
	Name               types.String                                             `tfsdk:"name"`
	SSID               *ResponseCameraGetNetworkCameraWirelessProfileSsidRs     `tfsdk:"ssid"`
}

type ResponseCameraGetNetworkCameraWirelessProfileIdentityRs struct {
	Password types.String `tfsdk:"password"`
	Username types.String `tfsdk:"username"`
}

type ResponseCameraGetNetworkCameraWirelessProfileSsidRs struct {
	AuthMode       types.String `tfsdk:"auth_mode"`
	EncryptionMode types.String `tfsdk:"encryption_mode"`
	Name           types.String `tfsdk:"name"`
	Psk            types.String `tfsdk:"psk"`
}

// FromBody
func (r *NetworksCameraWirelessProfilesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestCameraCreateNetworkCameraWirelessProfile {
	emptyString := ""
	var requestCameraCreateNetworkCameraWirelessProfileIDentity *merakigosdk.RequestCameraCreateNetworkCameraWirelessProfileIDentity
	if r.IDentity != nil {
		password := r.IDentity.Password.ValueString()
		username := r.IDentity.Username.ValueString()
		requestCameraCreateNetworkCameraWirelessProfileIDentity = &merakigosdk.RequestCameraCreateNetworkCameraWirelessProfileIDentity{
			Password: password,
			Username: username,
		}
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestCameraCreateNetworkCameraWirelessProfileSSID *merakigosdk.RequestCameraCreateNetworkCameraWirelessProfileSSID
	if r.SSID != nil {
		authMode := r.SSID.AuthMode.ValueString()
		encryptionMode := r.SSID.EncryptionMode.ValueString()
		name := r.SSID.Name.ValueString()
		psk := r.SSID.Psk.ValueString()
		requestCameraCreateNetworkCameraWirelessProfileSSID = &merakigosdk.RequestCameraCreateNetworkCameraWirelessProfileSSID{
			AuthMode:       authMode,
			EncryptionMode: encryptionMode,
			Name:           name,
			Psk:            psk,
		}
	}
	out := merakigosdk.RequestCameraCreateNetworkCameraWirelessProfile{
		IDentity: requestCameraCreateNetworkCameraWirelessProfileIDentity,
		Name:     *name,
		SSID:     requestCameraCreateNetworkCameraWirelessProfileSSID,
	}
	return &out
}
func (r *NetworksCameraWirelessProfilesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestCameraUpdateNetworkCameraWirelessProfile {
	emptyString := ""
	var requestCameraUpdateNetworkCameraWirelessProfileIDentity *merakigosdk.RequestCameraUpdateNetworkCameraWirelessProfileIDentity
	if r.IDentity != nil {
		password := r.IDentity.Password.ValueString()
		username := r.IDentity.Username.ValueString()
		requestCameraUpdateNetworkCameraWirelessProfileIDentity = &merakigosdk.RequestCameraUpdateNetworkCameraWirelessProfileIDentity{
			Password: password,
			Username: username,
		}
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestCameraUpdateNetworkCameraWirelessProfileSSID *merakigosdk.RequestCameraUpdateNetworkCameraWirelessProfileSSID
	if r.SSID != nil {
		authMode := r.SSID.AuthMode.ValueString()
		encryptionMode := r.SSID.EncryptionMode.ValueString()
		name := r.SSID.Name.ValueString()
		psk := r.SSID.Psk.ValueString()
		requestCameraUpdateNetworkCameraWirelessProfileSSID = &merakigosdk.RequestCameraUpdateNetworkCameraWirelessProfileSSID{
			AuthMode:       authMode,
			EncryptionMode: encryptionMode,
			Name:           name,
			Psk:            psk,
		}
	}
	out := merakigosdk.RequestCameraUpdateNetworkCameraWirelessProfile{
		IDentity: requestCameraUpdateNetworkCameraWirelessProfileIDentity,
		Name:     *name,
		SSID:     requestCameraUpdateNetworkCameraWirelessProfileSSID,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseCameraGetNetworkCameraWirelessProfileItemToBodyRs(state NetworksCameraWirelessProfilesRs, response *merakigosdk.ResponseCameraGetNetworkCameraWirelessProfile, is_read bool) NetworksCameraWirelessProfilesRs {
	itemState := NetworksCameraWirelessProfilesRs{
		AppliedDeviceCount: func() types.Int64 {
			if response.AppliedDeviceCount != nil {
				return types.Int64Value(int64(*response.AppliedDeviceCount))
			}
			return types.Int64{}
		}(),
		ID: types.StringValue(response.ID),
		IDentity: func() *ResponseCameraGetNetworkCameraWirelessProfileIdentityRs {
			if response.IDentity != nil {
				return &ResponseCameraGetNetworkCameraWirelessProfileIdentityRs{
					Password: types.StringValue(response.IDentity.Password),
					Username: types.StringValue(response.IDentity.Username),
				}
			}
			return nil
		}(),
		Name: types.StringValue(response.Name),
		SSID: func() *ResponseCameraGetNetworkCameraWirelessProfileSsidRs {
			if response.SSID != nil {
				return &ResponseCameraGetNetworkCameraWirelessProfileSsidRs{
					AuthMode:       types.StringValue(response.SSID.AuthMode),
					EncryptionMode: types.StringValue(response.SSID.EncryptionMode),
					Name:           types.StringValue(response.SSID.Name),
					Psk:            types.StringValue(response.SSID.Psk),
				}
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksCameraWirelessProfilesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksCameraWirelessProfilesRs)
}
