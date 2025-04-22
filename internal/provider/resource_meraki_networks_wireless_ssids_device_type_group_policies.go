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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWirelessSSIDsDeviceTypeGroupPoliciesResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessSSIDsDeviceTypeGroupPoliciesResource{}
)

func NewNetworksWirelessSSIDsDeviceTypeGroupPoliciesResource() resource.Resource {
	return &NetworksWirelessSSIDsDeviceTypeGroupPoliciesResource{}
}

type NetworksWirelessSSIDsDeviceTypeGroupPoliciesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessSSIDsDeviceTypeGroupPoliciesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessSSIDsDeviceTypeGroupPoliciesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_device_type_group_policies"
}

func (r *NetworksWirelessSSIDsDeviceTypeGroupPoliciesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"device_type_policies": schema.SetNestedAttribute{
				MarkdownDescription: `List of device type policies.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"device_policy": schema.StringAttribute{
							MarkdownDescription: `The device policy. Can be one of 'Allowed', 'Blocked' or 'Group policy'
                                        Allowed values: [Allowed,Blocked,Group policy]`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"Allowed",
									"Blocked",
									"Group policy",
								),
							},
						},
						"device_type": schema.StringAttribute{
							MarkdownDescription: `The device type. Can be one of 'Android', 'BlackBerry', 'Chrome OS', 'iPad', 'iPhone', 'iPod', 'Mac OS X', 'Windows', 'Windows Phone', 'B&N Nook' or 'Other OS'
                                        Allowed values: [Android,B&N Nook,BlackBerry,Chrome OS,Mac OS X,Other OS,Windows,Windows Phone,iPad,iPhone,iPod]`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"Android",
									"B&N Nook",
									"BlackBerry",
									"Chrome OS",
									"Mac OS X",
									"Other OS",
									"Windows",
									"Windows Phone",
									"iPad",
									"iPhone",
									"iPod",
								),
							},
						},
						"group_policy_id": schema.Int64Attribute{
							MarkdownDescription: `ID of the group policy object.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `If true, the SSID device type group policies are enabled.`,
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
			"number": schema.StringAttribute{
				MarkdownDescription: `number path parameter.`,
				Required:            true,
			},
		},
	}
}

func (r *NetworksWirelessSSIDsDeviceTypeGroupPoliciesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessSSIDsDeviceTypeGroupPoliciesRs

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
		responseVerifyItem, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDDeviceTypeGroupPolicies(vvNetworkID, vvNumber)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksWirelessSsidsDeviceTypeGroupPolicies  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksWirelessSsidsDeviceTypeGroupPolicies only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDDeviceTypeGroupPolicies(vvNetworkID, vvNumber, dataRequest)
	//Update
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDDeviceTypeGroupPolicies",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDDeviceTypeGroupPolicies",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDDeviceTypeGroupPolicies(vvNetworkID, vvNumber)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDDeviceTypeGroupPolicies",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDDeviceTypeGroupPolicies",
			err.Error(),
		)
		return
	}

	data = ResponseWirelessGetNetworkWirelessSSIDDeviceTypeGroupPoliciesItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *NetworksWirelessSSIDsDeviceTypeGroupPoliciesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessSSIDsDeviceTypeGroupPoliciesRs

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
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessSSIDDeviceTypeGroupPolicies(vvNetworkID, vvNumber)
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
				"Failure when executing GetNetworkWirelessSSIDDeviceTypeGroupPolicies",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDDeviceTypeGroupPolicies",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetNetworkWirelessSSIDDeviceTypeGroupPoliciesItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksWirelessSSIDsDeviceTypeGroupPoliciesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

func (r *NetworksWirelessSSIDsDeviceTypeGroupPoliciesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksWirelessSSIDsDeviceTypeGroupPoliciesRs
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
	restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDDeviceTypeGroupPolicies(vvNetworkID, vvNumber, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDDeviceTypeGroupPolicies",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDDeviceTypeGroupPolicies",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessSSIDsDeviceTypeGroupPoliciesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksWirelessSSIDsDeviceTypeGroupPolicies", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessSSIDsDeviceTypeGroupPoliciesRs struct {
	NetworkID          types.String                                                                         `tfsdk:"network_id"`
	Number             types.String                                                                         `tfsdk:"number"`
	DeviceTypePolicies *[]ResponseWirelessGetNetworkWirelessSsidDeviceTypeGroupPoliciesDeviceTypePoliciesRs `tfsdk:"device_type_policies"`
	Enabled            types.Bool                                                                           `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessSsidDeviceTypeGroupPoliciesDeviceTypePoliciesRs struct {
	DevicePolicy  types.String `tfsdk:"device_policy"`
	DeviceType    types.String `tfsdk:"device_type"`
	GroupPolicyID types.Int64  `tfsdk:"group_policy_id"`
}

// FromBody
func (r *NetworksWirelessSSIDsDeviceTypeGroupPoliciesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDDeviceTypeGroupPolicies {
	var requestWirelessUpdateNetworkWirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies

	if r.DeviceTypePolicies != nil {
		for _, rItem1 := range *r.DeviceTypePolicies {
			devicePolicy := rItem1.DevicePolicy.ValueString()
			deviceType := rItem1.DeviceType.ValueString()
			groupPolicyID := func() *int64 {
				if !rItem1.GroupPolicyID.IsUnknown() && !rItem1.GroupPolicyID.IsNull() {
					return rItem1.GroupPolicyID.ValueInt64Pointer()
				}
				return nil
			}()
			requestWirelessUpdateNetworkWirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies = append(requestWirelessUpdateNetworkWirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies{
				DevicePolicy:  devicePolicy,
				DeviceType:    deviceType,
				GroupPolicyID: int64ToIntPointer(groupPolicyID),
			})
			//[debug] Is Array: True
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDDeviceTypeGroupPolicies{
		DeviceTypePolicies: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies {
			if len(requestWirelessUpdateNetworkWirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies) > 0 {
				return &requestWirelessUpdateNetworkWirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies
			}
			return nil
		}(),
		Enabled: enabled,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessSSIDDeviceTypeGroupPoliciesItemToBodyRs(state NetworksWirelessSSIDsDeviceTypeGroupPoliciesRs, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDDeviceTypeGroupPolicies, is_read bool) NetworksWirelessSSIDsDeviceTypeGroupPoliciesRs {
	itemState := NetworksWirelessSSIDsDeviceTypeGroupPoliciesRs{
		DeviceTypePolicies: func() *[]ResponseWirelessGetNetworkWirelessSsidDeviceTypeGroupPoliciesDeviceTypePoliciesRs {
			if response.DeviceTypePolicies != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidDeviceTypeGroupPoliciesDeviceTypePoliciesRs, len(*response.DeviceTypePolicies))
				for i, deviceTypePolicies := range *response.DeviceTypePolicies {
					result[i] = ResponseWirelessGetNetworkWirelessSsidDeviceTypeGroupPoliciesDeviceTypePoliciesRs{
						DevicePolicy: types.StringValue(deviceTypePolicies.DevicePolicy),
						DeviceType:   types.StringValue(deviceTypePolicies.DeviceType),
					}
				}
				return &result
			}
			return nil
		}(),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessSSIDsDeviceTypeGroupPoliciesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessSSIDsDeviceTypeGroupPoliciesRs)
}
