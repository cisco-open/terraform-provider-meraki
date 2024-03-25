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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksClientsSplashAuthorizationStatusResource{}
	_ resource.ResourceWithConfigure = &NetworksClientsSplashAuthorizationStatusResource{}
)

func NewNetworksClientsSplashAuthorizationStatusResource() resource.Resource {
	return &NetworksClientsSplashAuthorizationStatusResource{}
}

type NetworksClientsSplashAuthorizationStatusResource struct {
	client *merakigosdk.Client
}

func (r *NetworksClientsSplashAuthorizationStatusResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksClientsSplashAuthorizationStatusResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_clients_splash_authorization_status"
}

func (r *NetworksClientsSplashAuthorizationStatusResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"client_id": schema.StringAttribute{
				MarkdownDescription: `clientId path parameter. Client ID`,
				Required:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"ssids": schema.SingleNestedAttribute{
				MarkdownDescription: `The target SSIDs. Each SSID must be enabled and must have Click-through splash enabled. For each SSID where isAuthorized is true, the expiration time will automatically be set according to the SSID's splash frequency. Not all networks support configuring all SSIDs`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"status_0": schema.SingleNestedAttribute{
						MarkdownDescription: `Splash authorization for SSID 0`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"authorized_at": schema.StringAttribute{
								Computed: true,
							},
							"expires_at": schema.StringAttribute{
								Computed: true,
							},
							"is_authorized": schema.BoolAttribute{
								MarkdownDescription: `New authorization status for the SSID (true, false).`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"status_1": schema.SingleNestedAttribute{
						MarkdownDescription: `Splash authorization for SSID 1`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"is_authorized": schema.BoolAttribute{
								MarkdownDescription: `New authorization status for the SSID (true, false).`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"status_10": schema.SingleNestedAttribute{
						MarkdownDescription: `Splash authorization for SSID 10`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"is_authorized": schema.BoolAttribute{
								MarkdownDescription: `New authorization status for the SSID (true, false).`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"status_11": schema.SingleNestedAttribute{
						MarkdownDescription: `Splash authorization for SSID 11`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"is_authorized": schema.BoolAttribute{
								MarkdownDescription: `New authorization status for the SSID (true, false).`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"status_12": schema.SingleNestedAttribute{
						MarkdownDescription: `Splash authorization for SSID 12`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"is_authorized": schema.BoolAttribute{
								MarkdownDescription: `New authorization status for the SSID (true, false).`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"status_13": schema.SingleNestedAttribute{
						MarkdownDescription: `Splash authorization for SSID 13`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"is_authorized": schema.BoolAttribute{
								MarkdownDescription: `New authorization status for the SSID (true, false).`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"status_14": schema.SingleNestedAttribute{
						MarkdownDescription: `Splash authorization for SSID 14`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"is_authorized": schema.BoolAttribute{
								MarkdownDescription: `New authorization status for the SSID (true, false).`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"status_2": schema.SingleNestedAttribute{
						MarkdownDescription: `Splash authorization for SSID 2`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"is_authorized": schema.BoolAttribute{
								MarkdownDescription: `New authorization status for the SSID (true, false).`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"status_3": schema.SingleNestedAttribute{
						MarkdownDescription: `Splash authorization for SSID 3`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"is_authorized": schema.BoolAttribute{
								MarkdownDescription: `New authorization status for the SSID (true, false).`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"status_4": schema.SingleNestedAttribute{
						MarkdownDescription: `Splash authorization for SSID 4`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"is_authorized": schema.BoolAttribute{
								MarkdownDescription: `New authorization status for the SSID (true, false).`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"status_5": schema.SingleNestedAttribute{
						MarkdownDescription: `Splash authorization for SSID 5`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"is_authorized": schema.BoolAttribute{
								MarkdownDescription: `New authorization status for the SSID (true, false).`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"status_6": schema.SingleNestedAttribute{
						MarkdownDescription: `Splash authorization for SSID 6`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"is_authorized": schema.BoolAttribute{
								MarkdownDescription: `New authorization status for the SSID (true, false).`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"status_7": schema.SingleNestedAttribute{
						MarkdownDescription: `Splash authorization for SSID 7`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"is_authorized": schema.BoolAttribute{
								MarkdownDescription: `New authorization status for the SSID (true, false).`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"status_8": schema.SingleNestedAttribute{
						MarkdownDescription: `Splash authorization for SSID 8`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"is_authorized": schema.BoolAttribute{
								MarkdownDescription: `New authorization status for the SSID (true, false).`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"status_9": schema.SingleNestedAttribute{
						MarkdownDescription: `Splash authorization for SSID 9`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"is_authorized": schema.BoolAttribute{
								MarkdownDescription: `New authorization status for the SSID (true, false).`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *NetworksClientsSplashAuthorizationStatusResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksClientsSplashAuthorizationStatusRs

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
	vvClientID := data.ClientID.ValueString()
	//Item
	responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkClientSplashAuthorizationStatus(vvNetworkID, vvClientID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksClientsSplashAuthorizationStatus only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksClientsSplashAuthorizationStatus only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Networks.UpdateNetworkClientSplashAuthorizationStatus(vvNetworkID, vvClientID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkClientSplashAuthorizationStatus",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkClientSplashAuthorizationStatus",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Networks.GetNetworkClientSplashAuthorizationStatus(vvNetworkID, vvClientID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkClientSplashAuthorizationStatus",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkClientSplashAuthorizationStatus",
			err.Error(),
		)
		return
	}

	data = ResponseNetworksGetNetworkClientSplashAuthorizationStatusItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksClientsSplashAuthorizationStatusResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksClientsSplashAuthorizationStatusRs

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
	vvClientID := data.ClientID.ValueString()
	// client_id
	responseGet, restyRespGet, err := r.client.Networks.GetNetworkClientSplashAuthorizationStatus(vvNetworkID, vvClientID)
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
				"Failure when executing GetNetworkClientSplashAuthorizationStatus",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkClientSplashAuthorizationStatus",
			err.Error(),
		)
		return
	}

	data = ResponseNetworksGetNetworkClientSplashAuthorizationStatusItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksClientsSplashAuthorizationStatusResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("client_id"), idParts[1])...)
}

func (r *NetworksClientsSplashAuthorizationStatusResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksClientsSplashAuthorizationStatusRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvClientID := data.ClientID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Networks.UpdateNetworkClientSplashAuthorizationStatus(vvNetworkID, vvClientID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkClientSplashAuthorizationStatus",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkClientSplashAuthorizationStatus",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksClientsSplashAuthorizationStatusResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksClientsSplashAuthorizationStatusRs struct {
	NetworkID types.String                                                      `tfsdk:"network_id"`
	ClientID  types.String                                                      `tfsdk:"client_id"`
	SSIDs     *ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsidsRs `tfsdk:"ssids"`
}

type ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsidsRs struct {
	Status0  *ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsids0Rs    `tfsdk:"status_0"`
	Status2  *ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsids2Rs    `tfsdk:"status_2"`
	Status1  *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids1Rs  `tfsdk:"status_1"`
	Status10 *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids10Rs `tfsdk:"status_10"`
	Status11 *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids11Rs `tfsdk:"status_11"`
	Status12 *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids12Rs `tfsdk:"status_12"`
	Status13 *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids13Rs `tfsdk:"status_13"`
	Status14 *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids14Rs `tfsdk:"status_14"`
	Status3  *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids3Rs  `tfsdk:"status_3"`
	Status4  *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids4Rs  `tfsdk:"status_4"`
	Status5  *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids5Rs  `tfsdk:"status_5"`
	Status6  *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids6Rs  `tfsdk:"status_6"`
	Status7  *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids7Rs  `tfsdk:"status_7"`
	Status8  *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids8Rs  `tfsdk:"status_8"`
	Status9  *RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids9Rs  `tfsdk:"status_9"`
}

type ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsids0Rs struct {
	AuthorizedAt types.String `tfsdk:"authorized_at"`
	ExpiresAt    types.String `tfsdk:"expires_at"`
	IsAuthorized types.Bool   `tfsdk:"is_authorized"`
}

type ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsids2Rs struct {
	IsAuthorized types.Bool `tfsdk:"is_authorized"`
}

type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids1Rs struct {
	IsAuthorized types.Bool `tfsdk:"is_authorized"`
}

type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids10Rs struct {
	IsAuthorized types.Bool `tfsdk:"is_authorized"`
}

type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids11Rs struct {
	IsAuthorized types.Bool `tfsdk:"is_authorized"`
}

type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids12Rs struct {
	IsAuthorized types.Bool `tfsdk:"is_authorized"`
}

type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids13Rs struct {
	IsAuthorized types.Bool `tfsdk:"is_authorized"`
}

type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids14Rs struct {
	IsAuthorized types.Bool `tfsdk:"is_authorized"`
}

type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids3Rs struct {
	IsAuthorized types.Bool `tfsdk:"is_authorized"`
}

type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids4Rs struct {
	IsAuthorized types.Bool `tfsdk:"is_authorized"`
}

type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids5Rs struct {
	IsAuthorized types.Bool `tfsdk:"is_authorized"`
}

type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids6Rs struct {
	IsAuthorized types.Bool `tfsdk:"is_authorized"`
}

type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids7Rs struct {
	IsAuthorized types.Bool `tfsdk:"is_authorized"`
}

type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids8Rs struct {
	IsAuthorized types.Bool `tfsdk:"is_authorized"`
}

type RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSsids9Rs struct {
	IsAuthorized types.Bool `tfsdk:"is_authorized"`
}

// FromBody
func (r *NetworksClientsSplashAuthorizationStatusRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatus {
	var requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs *merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs
	if r.SSIDs != nil {
		var requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs0 *merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs0
		if r.SSIDs.Status0 != nil {
			isAuthorized := func() *bool {
				if !r.SSIDs.Status0.IsAuthorized.IsUnknown() && !r.SSIDs.Status0.IsAuthorized.IsNull() {
					return r.SSIDs.Status0.IsAuthorized.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs0 = &merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs0{
				IsAuthorized: isAuthorized,
			}
		}
		var requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs1 *merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs1
		if r.SSIDs.Status1 != nil {
			isAuthorized := func() *bool {
				if !r.SSIDs.Status1.IsAuthorized.IsUnknown() && !r.SSIDs.Status1.IsAuthorized.IsNull() {
					return r.SSIDs.Status1.IsAuthorized.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs1 = &merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs1{
				IsAuthorized: isAuthorized,
			}
		}
		var requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs10 *merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs10
		if r.SSIDs.Status10 != nil {
			isAuthorized := func() *bool {
				if !r.SSIDs.Status10.IsAuthorized.IsUnknown() && !r.SSIDs.Status10.IsAuthorized.IsNull() {
					return r.SSIDs.Status10.IsAuthorized.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs10 = &merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs10{
				IsAuthorized: isAuthorized,
			}
		}
		var requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs11 *merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs11
		if r.SSIDs.Status11 != nil {
			isAuthorized := func() *bool {
				if !r.SSIDs.Status11.IsAuthorized.IsUnknown() && !r.SSIDs.Status11.IsAuthorized.IsNull() {
					return r.SSIDs.Status11.IsAuthorized.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs11 = &merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs11{
				IsAuthorized: isAuthorized,
			}
		}
		var requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs12 *merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs12
		if r.SSIDs.Status12 != nil {
			isAuthorized := func() *bool {
				if !r.SSIDs.Status12.IsAuthorized.IsUnknown() && !r.SSIDs.Status12.IsAuthorized.IsNull() {
					return r.SSIDs.Status12.IsAuthorized.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs12 = &merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs12{
				IsAuthorized: isAuthorized,
			}
		}
		var requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs13 *merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs13
		if r.SSIDs.Status13 != nil {
			isAuthorized := func() *bool {
				if !r.SSIDs.Status13.IsAuthorized.IsUnknown() && !r.SSIDs.Status13.IsAuthorized.IsNull() {
					return r.SSIDs.Status13.IsAuthorized.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs13 = &merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs13{
				IsAuthorized: isAuthorized,
			}
		}
		var requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs14 *merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs14
		if r.SSIDs.Status14 != nil {
			isAuthorized := func() *bool {
				if !r.SSIDs.Status14.IsAuthorized.IsUnknown() && !r.SSIDs.Status14.IsAuthorized.IsNull() {
					return r.SSIDs.Status14.IsAuthorized.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs14 = &merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs14{
				IsAuthorized: isAuthorized,
			}
		}
		var requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs2 *merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs2
		if r.SSIDs.Status2 != nil {
			isAuthorized := func() *bool {
				if !r.SSIDs.Status2.IsAuthorized.IsUnknown() && !r.SSIDs.Status2.IsAuthorized.IsNull() {
					return r.SSIDs.Status2.IsAuthorized.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs2 = &merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs2{
				IsAuthorized: isAuthorized,
			}
		}
		var requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs3 *merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs3
		if r.SSIDs.Status3 != nil {
			isAuthorized := func() *bool {
				if !r.SSIDs.Status3.IsAuthorized.IsUnknown() && !r.SSIDs.Status3.IsAuthorized.IsNull() {
					return r.SSIDs.Status3.IsAuthorized.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs3 = &merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs3{
				IsAuthorized: isAuthorized,
			}
		}
		var requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs4 *merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs4
		if r.SSIDs.Status4 != nil {
			isAuthorized := func() *bool {
				if !r.SSIDs.Status4.IsAuthorized.IsUnknown() && !r.SSIDs.Status4.IsAuthorized.IsNull() {
					return r.SSIDs.Status4.IsAuthorized.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs4 = &merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs4{
				IsAuthorized: isAuthorized,
			}
		}
		var requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs5 *merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs5
		if r.SSIDs.Status5 != nil {
			isAuthorized := func() *bool {
				if !r.SSIDs.Status5.IsAuthorized.IsUnknown() && !r.SSIDs.Status5.IsAuthorized.IsNull() {
					return r.SSIDs.Status5.IsAuthorized.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs5 = &merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs5{
				IsAuthorized: isAuthorized,
			}
		}
		var requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs6 *merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs6
		if r.SSIDs.Status6 != nil {
			isAuthorized := func() *bool {
				if !r.SSIDs.Status6.IsAuthorized.IsUnknown() && !r.SSIDs.Status6.IsAuthorized.IsNull() {
					return r.SSIDs.Status6.IsAuthorized.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs6 = &merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs6{
				IsAuthorized: isAuthorized,
			}
		}
		var requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs7 *merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs7
		if r.SSIDs.Status7 != nil {
			isAuthorized := func() *bool {
				if !r.SSIDs.Status7.IsAuthorized.IsUnknown() && !r.SSIDs.Status7.IsAuthorized.IsNull() {
					return r.SSIDs.Status7.IsAuthorized.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs7 = &merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs7{
				IsAuthorized: isAuthorized,
			}
		}
		var requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs8 *merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs8
		if r.SSIDs.Status8 != nil {
			isAuthorized := func() *bool {
				if !r.SSIDs.Status8.IsAuthorized.IsUnknown() && !r.SSIDs.Status8.IsAuthorized.IsNull() {
					return r.SSIDs.Status8.IsAuthorized.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs8 = &merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs8{
				IsAuthorized: isAuthorized,
			}
		}
		var requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs9 *merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs9
		if r.SSIDs.Status9 != nil {
			isAuthorized := func() *bool {
				if !r.SSIDs.Status9.IsAuthorized.IsUnknown() && !r.SSIDs.Status9.IsAuthorized.IsNull() {
					return r.SSIDs.Status9.IsAuthorized.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs9 = &merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs9{
				IsAuthorized: isAuthorized,
			}
		}
		requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs = &merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs{
			Status0:  requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs0,
			Status1:  requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs1,
			Status10: requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs10,
			Status11: requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs11,
			Status12: requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs12,
			Status13: requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs13,
			Status14: requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs14,
			Status2:  requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs2,
			Status3:  requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs3,
			Status4:  requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs4,
			Status5:  requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs5,
			Status6:  requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs6,
			Status7:  requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs7,
			Status8:  requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs8,
			Status9:  requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs9,
		}
	}
	out := merakigosdk.RequestNetworksUpdateNetworkClientSplashAuthorizationStatus{
		SSIDs: requestNetworksUpdateNetworkClientSplashAuthorizationStatusSSIDs,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkClientSplashAuthorizationStatusItemToBodyRs(state NetworksClientsSplashAuthorizationStatusRs, response *merakigosdk.ResponseNetworksGetNetworkClientSplashAuthorizationStatus, is_read bool) NetworksClientsSplashAuthorizationStatusRs {
	itemState := NetworksClientsSplashAuthorizationStatusRs{
		SSIDs: func() *ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsidsRs {
			if response.SSIDs != nil {
				return &ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsidsRs{
					Status0: func() *ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsids0Rs {
						if response.SSIDs.Status0 != nil {
							return &ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsids0Rs{
								AuthorizedAt: types.StringValue(response.SSIDs.Status0.AuthorizedAt),
								ExpiresAt:    types.StringValue(response.SSIDs.Status0.ExpiresAt),
								IsAuthorized: func() types.Bool {
									if response.SSIDs.Status0.IsAuthorized != nil {
										return types.BoolValue(*response.SSIDs.Status0.IsAuthorized)
									}
									return types.Bool{}
								}(),
							}
						}
						return &ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsids0Rs{}
					}(),
					Status2: func() *ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsids2Rs {
						if response.SSIDs.Status2 != nil {
							return &ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsids2Rs{
								IsAuthorized: func() types.Bool {
									if response.SSIDs.Status2.IsAuthorized != nil {
										return types.BoolValue(*response.SSIDs.Status2.IsAuthorized)
									}
									return types.Bool{}
								}(),
							}
						}
						return &ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsids2Rs{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsidsRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksClientsSplashAuthorizationStatusRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksClientsSplashAuthorizationStatusRs)
}
