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

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWirelessEthernetPortsProfilesResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessEthernetPortsProfilesResource{}
)

func NewNetworksWirelessEthernetPortsProfilesResource() resource.Resource {
	return &NetworksWirelessEthernetPortsProfilesResource{}
}

type NetworksWirelessEthernetPortsProfilesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessEthernetPortsProfilesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessEthernetPortsProfilesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ethernet_ports_profiles"
}

func (r *NetworksWirelessEthernetPortsProfilesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"is_default": schema.BoolAttribute{
				MarkdownDescription: `Is default profile`,
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `AP port profile name`,
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
			"ports": schema.SetNestedAttribute{
				MarkdownDescription: `Ports config`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"enabled": schema.BoolAttribute{
							MarkdownDescription: `Enabled`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"number": schema.Int64Attribute{
							MarkdownDescription: `Number`,
							Computed:            true,
						},
						"psk_group_id": schema.StringAttribute{
							MarkdownDescription: `PSK Group number`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"ssid": schema.Int64Attribute{
							MarkdownDescription: `Ssid number`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"profile_id": schema.StringAttribute{
				MarkdownDescription: `AP port profile ID`,
				Required:            true,
			},
			"usb_ports": schema.SetNestedAttribute{
				MarkdownDescription: `Usb ports config`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"enabled": schema.BoolAttribute{
							MarkdownDescription: `Enabled`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"ssid": schema.Int64Attribute{
							MarkdownDescription: `Ssid number`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
		},
	}
}

func (r *NetworksWirelessEthernetPortsProfilesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessEthernetPortsProfilesRs

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
	vvProfileID := data.ProfileID.ValueString()
	//Item
	responseVerifyItem, restyResp1, err := r.client.Wireless.GetNetworkWirelessEthernetPortsProfile(vvNetworkID, vvProfileID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksWirelessEthernetPortsProfiles only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksWirelessEthernetPortsProfiles only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessEthernetPortsProfile(vvNetworkID, vvProfileID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessEthernetPortsProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessEthernetPortsProfile",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Wireless.GetNetworkWirelessEthernetPortsProfile(vvNetworkID, vvProfileID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessEthernetPortsProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessEthernetPortsProfile",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetNetworkWirelessEthernetPortsProfileItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessEthernetPortsProfilesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessEthernetPortsProfilesRs

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
	vvProfileID := data.ProfileID.ValueString()
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessEthernetPortsProfile(vvNetworkID, vvProfileID)
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
				"Failure when executing GetNetworkWirelessEthernetPortsProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessEthernetPortsProfile",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetNetworkWirelessEthernetPortsProfileItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessEthernetPortsProfilesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("profile_id"), idParts[1])...)
}

func (r *NetworksWirelessEthernetPortsProfilesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksWirelessEthernetPortsProfilesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvProfileID := data.ProfileID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessEthernetPortsProfile(vvNetworkID, vvProfileID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessEthernetPortsProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessEthernetPortsProfile",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessEthernetPortsProfilesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksWirelessEthernetPortsProfilesRs
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
	vvProfileID := state.ProfileID.ValueString()
	_, err := r.client.Wireless.DeleteNetworkWirelessEthernetPortsProfile(vvNetworkID, vvProfileID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkWirelessEthernetPortsProfile", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksWirelessEthernetPortsProfilesRs struct {
	NetworkID types.String                                                        `tfsdk:"network_id"`
	ProfileID types.String                                                        `tfsdk:"profile_id"`
	IsDefault types.Bool                                                          `tfsdk:"is_default"`
	Name      types.String                                                        `tfsdk:"name"`
	Ports     *[]ResponseWirelessGetNetworkWirelessEthernetPortsProfilePortsRs    `tfsdk:"ports"`
	UsbPorts  *[]ResponseWirelessGetNetworkWirelessEthernetPortsProfileUsbPortsRs `tfsdk:"usb_ports"`
}

type ResponseWirelessGetNetworkWirelessEthernetPortsProfilePortsRs struct {
	Enabled    types.Bool   `tfsdk:"enabled"`
	Name       types.String `tfsdk:"name"`
	Number     types.Int64  `tfsdk:"number"`
	PskGroupID types.String `tfsdk:"psk_group_id"`
	SSID       types.Int64  `tfsdk:"ssid"`
}

type ResponseWirelessGetNetworkWirelessEthernetPortsProfileUsbPortsRs struct {
	Enabled types.Bool   `tfsdk:"enabled"`
	Name    types.String `tfsdk:"name"`
	SSID    types.Int64  `tfsdk:"ssid"`
}

// FromBody
func (r *NetworksWirelessEthernetPortsProfilesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessEthernetPortsProfile {
	emptyString := ""
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestWirelessUpdateNetworkWirelessEthernetPortsProfilePorts []merakigosdk.RequestWirelessUpdateNetworkWirelessEthernetPortsProfilePorts
	if r.Ports != nil {
		for _, rItem1 := range *r.Ports {
			enabled := func() *bool {
				if !rItem1.Enabled.IsUnknown() && !rItem1.Enabled.IsNull() {
					return rItem1.Enabled.ValueBoolPointer()
				}
				return nil
			}()
			name := rItem1.Name.ValueString()
			pskGroupID := rItem1.PskGroupID.ValueString()
			sSID := func() *int64 {
				if !rItem1.SSID.IsUnknown() && !rItem1.SSID.IsNull() {
					return rItem1.SSID.ValueInt64Pointer()
				}
				return nil
			}()
			requestWirelessUpdateNetworkWirelessEthernetPortsProfilePorts = append(requestWirelessUpdateNetworkWirelessEthernetPortsProfilePorts, merakigosdk.RequestWirelessUpdateNetworkWirelessEthernetPortsProfilePorts{
				Enabled:    enabled,
				Name:       name,
				PskGroupID: pskGroupID,
				SSID:       int64ToIntPointer(sSID),
			})
		}
	}
	var requestWirelessUpdateNetworkWirelessEthernetPortsProfileUsbPorts []merakigosdk.RequestWirelessUpdateNetworkWirelessEthernetPortsProfileUsbPorts
	if r.UsbPorts != nil {
		for _, rItem1 := range *r.UsbPorts {
			enabled := func() *bool {
				if !rItem1.Enabled.IsUnknown() && !rItem1.Enabled.IsNull() {
					return rItem1.Enabled.ValueBoolPointer()
				}
				return nil
			}()
			name := rItem1.Name.ValueString()
			sSID := func() *int64 {
				if !rItem1.SSID.IsUnknown() && !rItem1.SSID.IsNull() {
					return rItem1.SSID.ValueInt64Pointer()
				}
				return nil
			}()
			requestWirelessUpdateNetworkWirelessEthernetPortsProfileUsbPorts = append(requestWirelessUpdateNetworkWirelessEthernetPortsProfileUsbPorts, merakigosdk.RequestWirelessUpdateNetworkWirelessEthernetPortsProfileUsbPorts{
				Enabled: enabled,
				Name:    name,
				SSID:    int64ToIntPointer(sSID),
			})
		}
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessEthernetPortsProfile{
		Name: *name,
		Ports: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessEthernetPortsProfilePorts {
			if len(requestWirelessUpdateNetworkWirelessEthernetPortsProfilePorts) > 0 {
				return &requestWirelessUpdateNetworkWirelessEthernetPortsProfilePorts
			}
			return nil
		}(),
		UsbPorts: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessEthernetPortsProfileUsbPorts {
			if len(requestWirelessUpdateNetworkWirelessEthernetPortsProfileUsbPorts) > 0 {
				return &requestWirelessUpdateNetworkWirelessEthernetPortsProfileUsbPorts
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessEthernetPortsProfileItemToBodyRs(state NetworksWirelessEthernetPortsProfilesRs, response *merakigosdk.ResponseWirelessGetNetworkWirelessEthernetPortsProfile, is_read bool) NetworksWirelessEthernetPortsProfilesRs {
	itemState := NetworksWirelessEthernetPortsProfilesRs{
		IsDefault: func() types.Bool {
			if response.IsDefault != nil {
				return types.BoolValue(*response.IsDefault)
			}
			return types.Bool{}
		}(),
		Name: types.StringValue(response.Name),
		Ports: func() *[]ResponseWirelessGetNetworkWirelessEthernetPortsProfilePortsRs {
			if response.Ports != nil {
				result := make([]ResponseWirelessGetNetworkWirelessEthernetPortsProfilePortsRs, len(*response.Ports))
				for i, ports := range *response.Ports {
					result[i] = ResponseWirelessGetNetworkWirelessEthernetPortsProfilePortsRs{
						Enabled: func() types.Bool {
							if ports.Enabled != nil {
								return types.BoolValue(*ports.Enabled)
							}
							return types.Bool{}
						}(),
						Name: types.StringValue(ports.Name),
						Number: func() types.Int64 {
							if ports.Number != nil {
								return types.Int64Value(int64(*ports.Number))
							}
							return types.Int64{}
						}(),
						PskGroupID: types.StringValue(ports.PskGroupID),
						SSID: func() types.Int64 {
							if ports.SSID != nil {
								return types.Int64Value(int64(*ports.SSID))
							}
							return types.Int64{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		ProfileID: types.StringValue(response.ProfileID),
		UsbPorts: func() *[]ResponseWirelessGetNetworkWirelessEthernetPortsProfileUsbPortsRs {
			if response.UsbPorts != nil {
				result := make([]ResponseWirelessGetNetworkWirelessEthernetPortsProfileUsbPortsRs, len(*response.UsbPorts))
				for i, usbPorts := range *response.UsbPorts {
					result[i] = ResponseWirelessGetNetworkWirelessEthernetPortsProfileUsbPortsRs{
						Enabled: func() types.Bool {
							if usbPorts.Enabled != nil {
								return types.BoolValue(*usbPorts.Enabled)
							}
							return types.Bool{}
						}(),
						Name: types.StringValue(usbPorts.Name),
						SSID: func() types.Int64 {
							if usbPorts.SSID != nil {
								return types.Int64Value(int64(*usbPorts.SSID))
							}
							return types.Int64{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessEthernetPortsProfilesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessEthernetPortsProfilesRs)
}
