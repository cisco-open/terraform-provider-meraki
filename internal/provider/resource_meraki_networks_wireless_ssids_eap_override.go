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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWirelessSSIDsEapOverrideResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessSSIDsEapOverrideResource{}
)

func NewNetworksWirelessSSIDsEapOverrideResource() resource.Resource {
	return &NetworksWirelessSSIDsEapOverrideResource{}
}

type NetworksWirelessSSIDsEapOverrideResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessSSIDsEapOverrideResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessSSIDsEapOverrideResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_eap_override"
}

func (r *NetworksWirelessSSIDsEapOverrideResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"eapol_key": schema.SingleNestedAttribute{
				MarkdownDescription: `EAPOL Key settings.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"retries": schema.Int64Attribute{
						MarkdownDescription: `Maximum number of EAPOL key retries.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"timeout_in_ms": schema.Int64Attribute{
						MarkdownDescription: `EAPOL Key timeout in milliseconds.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"identity": schema.SingleNestedAttribute{
				MarkdownDescription: `EAP settings for identity requests.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"retries": schema.Int64Attribute{
						MarkdownDescription: `Maximum number of EAP retries.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"timeout": schema.Int64Attribute{
						MarkdownDescription: `EAP timeout in seconds.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"max_retries": schema.Int64Attribute{
				MarkdownDescription: `Maximum number of general EAP retries.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
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
			"timeout": schema.Int64Attribute{
				MarkdownDescription: `General EAP timeout in seconds.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *NetworksWirelessSSIDsEapOverrideResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessSSIDsEapOverrideRs

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
	// Has Paths
	vvNetworkID := data.NetworkID.ValueString()
	vvNumber := data.Number.ValueString()
	//Has Item and not has items

	if vvNetworkID != "" && vvNumber != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDEapOverride(vvNetworkID, vvNumber)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksWirelessSsidsEapOverride  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksWirelessSsidsEapOverride only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDEapOverride(vvNetworkID, vvNumber, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDEapOverride",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDEapOverride",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDEapOverride(vvNetworkID, vvNumber)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDEapOverride",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDEapOverride",
			err.Error(),
		)
		return
	}

	data = ResponseWirelessGetNetworkWirelessSSIDEapOverrideItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *NetworksWirelessSSIDsEapOverrideResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessSSIDsEapOverrideRs

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
	vvNumber := data.Number.ValueString()
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessSSIDEapOverride(vvNetworkID, vvNumber)
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
				"Failure when executing GetNetworkWirelessSSIDEapOverride",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDEapOverride",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetNetworkWirelessSSIDEapOverrideItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksWirelessSSIDsEapOverrideResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("number"), idParts[1])...)
}

func (r *NetworksWirelessSSIDsEapOverrideResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksWirelessSSIDsEapOverrideRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvNumber := data.Number.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDEapOverride(vvNetworkID, vvNumber, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDEapOverride",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDEapOverride",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessSSIDsEapOverrideResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksWirelessSSIDsEapOverride", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessSSIDsEapOverrideRs struct {
	NetworkID  types.String                                                 `tfsdk:"network_id"`
	Number     types.String                                                 `tfsdk:"number"`
	EapolKey   *ResponseWirelessGetNetworkWirelessSsidEapOverrideEapolKeyRs `tfsdk:"eapol_key"`
	IDentity   *ResponseWirelessGetNetworkWirelessSsidEapOverrideIdentityRs `tfsdk:"identity"`
	MaxRetries types.Int64                                                  `tfsdk:"max_retries"`
	Timeout    types.Int64                                                  `tfsdk:"timeout"`
}

type ResponseWirelessGetNetworkWirelessSsidEapOverrideEapolKeyRs struct {
	Retries     types.Int64 `tfsdk:"retries"`
	TimeoutInMs types.Int64 `tfsdk:"timeout_in_ms"`
}

type ResponseWirelessGetNetworkWirelessSsidEapOverrideIdentityRs struct {
	Retries types.Int64 `tfsdk:"retries"`
	Timeout types.Int64 `tfsdk:"timeout"`
}

// FromBody
func (r *NetworksWirelessSSIDsEapOverrideRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDEapOverride {
	var requestWirelessUpdateNetworkWirelessSSIDEapOverrideEapolKey *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDEapOverrideEapolKey

	if r.EapolKey != nil {
		retries := func() *int64 {
			if !r.EapolKey.Retries.IsUnknown() && !r.EapolKey.Retries.IsNull() {
				return r.EapolKey.Retries.ValueInt64Pointer()
			}
			return nil
		}()
		timeoutInMs := func() *int64 {
			if !r.EapolKey.TimeoutInMs.IsUnknown() && !r.EapolKey.TimeoutInMs.IsNull() {
				return r.EapolKey.TimeoutInMs.ValueInt64Pointer()
			}
			return nil
		}()
		requestWirelessUpdateNetworkWirelessSSIDEapOverrideEapolKey = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDEapOverrideEapolKey{
			Retries:     int64ToIntPointer(retries),
			TimeoutInMs: int64ToIntPointer(timeoutInMs),
		}
		//[debug] Is Array: False
	}
	var requestWirelessUpdateNetworkWirelessSSIDEapOverrideIDentity *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDEapOverrideIDentity

	if r.IDentity != nil {
		retries := func() *int64 {
			if !r.IDentity.Retries.IsUnknown() && !r.IDentity.Retries.IsNull() {
				return r.IDentity.Retries.ValueInt64Pointer()
			}
			return nil
		}()
		timeout := func() *int64 {
			if !r.IDentity.Timeout.IsUnknown() && !r.IDentity.Timeout.IsNull() {
				return r.IDentity.Timeout.ValueInt64Pointer()
			}
			return nil
		}()
		requestWirelessUpdateNetworkWirelessSSIDEapOverrideIDentity = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDEapOverrideIDentity{
			Retries: int64ToIntPointer(retries),
			Timeout: int64ToIntPointer(timeout),
		}
		//[debug] Is Array: False
	}
	maxRetries := new(int64)
	if !r.MaxRetries.IsUnknown() && !r.MaxRetries.IsNull() {
		*maxRetries = r.MaxRetries.ValueInt64()
	} else {
		maxRetries = nil
	}
	timeout := new(int64)
	if !r.Timeout.IsUnknown() && !r.Timeout.IsNull() {
		*timeout = r.Timeout.ValueInt64()
	} else {
		timeout = nil
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDEapOverride{
		EapolKey:   requestWirelessUpdateNetworkWirelessSSIDEapOverrideEapolKey,
		IDentity:   requestWirelessUpdateNetworkWirelessSSIDEapOverrideIDentity,
		MaxRetries: int64ToIntPointer(maxRetries),
		Timeout:    int64ToIntPointer(timeout),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessSSIDEapOverrideItemToBodyRs(state NetworksWirelessSSIDsEapOverrideRs, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDEapOverride, is_read bool) NetworksWirelessSSIDsEapOverrideRs {
	itemState := NetworksWirelessSSIDsEapOverrideRs{
		EapolKey: func() *ResponseWirelessGetNetworkWirelessSsidEapOverrideEapolKeyRs {
			if response.EapolKey != nil {
				return &ResponseWirelessGetNetworkWirelessSsidEapOverrideEapolKeyRs{
					Retries: func() types.Int64 {
						if response.EapolKey.Retries != nil {
							return types.Int64Value(int64(*response.EapolKey.Retries))
						}
						return types.Int64{}
					}(),
					TimeoutInMs: func() types.Int64 {
						if response.EapolKey.TimeoutInMs != nil {
							return types.Int64Value(int64(*response.EapolKey.TimeoutInMs))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
		IDentity: func() *ResponseWirelessGetNetworkWirelessSsidEapOverrideIdentityRs {
			if response.IDentity != nil {
				return &ResponseWirelessGetNetworkWirelessSsidEapOverrideIdentityRs{
					Retries: func() types.Int64 {
						if response.IDentity.Retries != nil {
							return types.Int64Value(int64(*response.IDentity.Retries))
						}
						return types.Int64{}
					}(),
					Timeout: func() types.Int64 {
						if response.IDentity.Timeout != nil {
							return types.Int64Value(int64(*response.IDentity.Timeout))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
		MaxRetries: func() types.Int64 {
			if response.MaxRetries != nil {
				return types.Int64Value(int64(*response.MaxRetries))
			}
			return types.Int64{}
		}(),
		Timeout: func() types.Int64 {
			if response.Timeout != nil {
				return types.Int64Value(int64(*response.Timeout))
			}
			return types.Int64{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessSSIDsEapOverrideRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessSSIDsEapOverrideRs)
}
