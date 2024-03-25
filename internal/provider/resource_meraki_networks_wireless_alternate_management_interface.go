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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

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
	_ resource.Resource              = &NetworksWirelessAlternateManagementInterfaceResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessAlternateManagementInterfaceResource{}
)

func NewNetworksWirelessAlternateManagementInterfaceResource() resource.Resource {
	return &NetworksWirelessAlternateManagementInterfaceResource{}
}

type NetworksWirelessAlternateManagementInterfaceResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessAlternateManagementInterfaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessAlternateManagementInterfaceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_alternate_management_interface"
}

func (r *NetworksWirelessAlternateManagementInterfaceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"access_points": schema.SetNestedAttribute{
				MarkdownDescription: `Array of access point serial number and IP assignment. Note: accessPoints IP assignment is not applicable for template networks, in other words, do not put 'accessPoints' in the body when updating template networks. Also, an empty 'accessPoints' array will remove all previous static IP assignments`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"alternate_management_ip": schema.StringAttribute{
							MarkdownDescription: `Wireless alternate management interface device IP. Provide an empty string to remove alternate management IP assignment`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"dns1": schema.StringAttribute{
							MarkdownDescription: `Primary DNS must be in IP format`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"dns2": schema.StringAttribute{
							MarkdownDescription: `Optional secondary DNS must be in IP format`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"gateway": schema.StringAttribute{
							MarkdownDescription: `Gateway must be in IP format`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `Serial number of access point to be configured with alternate management IP`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"subnet_mask": schema.StringAttribute{
							MarkdownDescription: `Subnet mask must be in IP format`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `Boolean value to enable or disable alternate management interface`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"protocols": schema.SetAttribute{
				MarkdownDescription: `Can be one or more of the following values: 'radius', 'snmp', 'syslog' or 'ldap'`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"vlan_id": schema.Int64Attribute{
				MarkdownDescription: `Alternate management interface VLAN, must be between 1 and 4094`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *NetworksWirelessAlternateManagementInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessAlternateManagementInterfaceRs

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
	//Item
	responseVerifyItem, restyResp1, err := r.client.Wireless.GetNetworkWirelessAlternateManagementInterface(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksWirelessAlternateManagementInterface only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksWirelessAlternateManagementInterface only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Wireless.UpdateNetworkWirelessAlternateManagementInterface(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessAlternateManagementInterface",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessAlternateManagementInterface",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Wireless.GetNetworkWirelessAlternateManagementInterface(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessAlternateManagementInterface",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessAlternateManagementInterface",
			err.Error(),
		)
		return
	}

	data = ResponseWirelessGetNetworkWirelessAlternateManagementInterfaceItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessAlternateManagementInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessAlternateManagementInterfaceRs

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
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessAlternateManagementInterface(vvNetworkID)
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
				"Failure when executing GetNetworkWirelessAlternateManagementInterface",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessAlternateManagementInterface",
			err.Error(),
		)
		return
	}

	data = ResponseWirelessGetNetworkWirelessAlternateManagementInterfaceItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksWirelessAlternateManagementInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksWirelessAlternateManagementInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksWirelessAlternateManagementInterfaceRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Wireless.UpdateNetworkWirelessAlternateManagementInterface(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessAlternateManagementInterface",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessAlternateManagementInterface",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessAlternateManagementInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessAlternateManagementInterfaceRs struct {
	NetworkID    types.String                                                                    `tfsdk:"network_id"`
	AccessPoints *[]ResponseWirelessGetNetworkWirelessAlternateManagementInterfaceAccessPointsRs `tfsdk:"access_points"`
	Enabled      types.Bool                                                                      `tfsdk:"enabled"`
	Protocols    types.Set                                                                       `tfsdk:"protocols"`
	VLANID       types.Int64                                                                     `tfsdk:"vlan_id"`
}

type ResponseWirelessGetNetworkWirelessAlternateManagementInterfaceAccessPointsRs struct {
	AlternateManagementIP types.String `tfsdk:"alternate_management_ip"`
	DNS1                  types.String `tfsdk:"dns1"`
	DNS2                  types.String `tfsdk:"dns2"`
	Gateway               types.String `tfsdk:"gateway"`
	Serial                types.String `tfsdk:"serial"`
	SubnetMask            types.String `tfsdk:"subnet_mask"`
}

// FromBody
func (r *NetworksWirelessAlternateManagementInterfaceRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessAlternateManagementInterface {
	var requestWirelessUpdateNetworkWirelessAlternateManagementInterfaceAccessPoints []merakigosdk.RequestWirelessUpdateNetworkWirelessAlternateManagementInterfaceAccessPoints
	if r.AccessPoints != nil {
		for _, rItem1 := range *r.AccessPoints {
			alternateManagementIP := rItem1.AlternateManagementIP.ValueString()
			dNS1 := rItem1.DNS1.ValueString()
			dNS2 := rItem1.DNS2.ValueString()
			gateway := rItem1.Gateway.ValueString()
			serial := rItem1.Serial.ValueString()
			subnetMask := rItem1.SubnetMask.ValueString()
			requestWirelessUpdateNetworkWirelessAlternateManagementInterfaceAccessPoints = append(requestWirelessUpdateNetworkWirelessAlternateManagementInterfaceAccessPoints, merakigosdk.RequestWirelessUpdateNetworkWirelessAlternateManagementInterfaceAccessPoints{
				AlternateManagementIP: alternateManagementIP,
				DNS1:                  dNS1,
				DNS2:                  dNS2,
				Gateway:               gateway,
				Serial:                serial,
				SubnetMask:            subnetMask,
			})
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var protocols []string = nil
	r.Protocols.ElementsAs(ctx, &protocols, false)
	vLANID := new(int64)
	if !r.VLANID.IsUnknown() && !r.VLANID.IsNull() {
		*vLANID = r.VLANID.ValueInt64()
	} else {
		vLANID = nil
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessAlternateManagementInterface{
		AccessPoints: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessAlternateManagementInterfaceAccessPoints {
			if len(requestWirelessUpdateNetworkWirelessAlternateManagementInterfaceAccessPoints) > 0 {
				return &requestWirelessUpdateNetworkWirelessAlternateManagementInterfaceAccessPoints
			}
			return nil
		}(),
		Enabled:   enabled,
		Protocols: protocols,
		VLANID:    int64ToIntPointer(vLANID),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessAlternateManagementInterfaceItemToBodyRs(state NetworksWirelessAlternateManagementInterfaceRs, response *merakigosdk.ResponseWirelessGetNetworkWirelessAlternateManagementInterface, is_read bool) NetworksWirelessAlternateManagementInterfaceRs {
	itemState := NetworksWirelessAlternateManagementInterfaceRs{
		AccessPoints: func() *[]ResponseWirelessGetNetworkWirelessAlternateManagementInterfaceAccessPointsRs {
			if response.AccessPoints != nil {
				result := make([]ResponseWirelessGetNetworkWirelessAlternateManagementInterfaceAccessPointsRs, len(*response.AccessPoints))
				for i, accessPoints := range *response.AccessPoints {
					result[i] = ResponseWirelessGetNetworkWirelessAlternateManagementInterfaceAccessPointsRs{
						AlternateManagementIP: types.StringValue(accessPoints.AlternateManagementIP),
						DNS1:                  types.StringValue(accessPoints.DNS1),
						DNS2:                  types.StringValue(accessPoints.DNS2),
						Gateway:               types.StringValue(accessPoints.Gateway),
						Serial:                types.StringValue(accessPoints.Serial),
						SubnetMask:            types.StringValue(accessPoints.SubnetMask),
					}
				}
				return &result
			}
			return &[]ResponseWirelessGetNetworkWirelessAlternateManagementInterfaceAccessPointsRs{}
		}(),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		Protocols: StringSliceToSet(response.Protocols),
		VLANID: func() types.Int64 {
			if response.VLANID != nil {
				return types.Int64Value(int64(*response.VLANID))
			}
			return types.Int64{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessAlternateManagementInterfaceRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessAlternateManagementInterfaceRs)
}
