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
	"net/url"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersResource{}
)

func NewNetworksSwitchDhcpServerPolicyArpInspectionTrustedServersResource() resource.Resource {
	return &NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersResource{}
}

type NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_dhcp_server_policy_arp_inspection_trusted_servers"
}

func (r *NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ipv4": schema.SingleNestedAttribute{
				MarkdownDescription: `IPv4 attributes of the trusted server.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"address": schema.StringAttribute{
						MarkdownDescription: `IPv4 address of the trusted server.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"mac": schema.StringAttribute{
				MarkdownDescription: `Mac address of the trusted server.`,
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
			"trusted_server_id": schema.StringAttribute{
				MarkdownDescription: `ID of the trusted server.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"vlan": schema.Int64Attribute{
				MarkdownDescription: `Vlan ID of the trusted server.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

//path params to set ['trustedServerId']

func (r *NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersRs

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
	vvMac := data.Mac.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := getAllItemsNetworksSwitchDhcpServerPolicyArpInspectionTrustedServers(*r.client, vvNetworkID)
	//Have Create
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers",
					err.Error(),
				)
				return
			}
		}
	}
	//TODO HAS ONLY ITEMS
	// Create

	responseStruct := structToMap(responseVerifyItem)
	result := getDictResult(responseStruct, "Mac", vvMac, simpleCmp)
	var responseVerifyItem2 merakigosdk.ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers
	if result != nil {
		err := mapToStruct(result.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		data = ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersItemToBodyRs(data, &responseVerifyItem2, false)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Switch.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := getAllItemsNetworksSwitchDhcpServerPolicyArpInspectionTrustedServers(*r.client, vvNetworkID)
	// Has only items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers",
			err.Error(),
		)
		return
	}
	responseStruct2 := structToMap(responseGet)
	result2 := getDictResult(responseStruct2, "Mac", vvMac, simpleCmp)
	if result2 != nil {
		err := mapToStruct(result2.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		data = ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersItemToBodyRs(data, &responseVerifyItem2, false)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	} else {
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchLinkAggregations Result",
			"Not Found",
		)
		return
	}
}

func (r *NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersRs

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
	// Not has Item

	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvMac := data.Mac.ValueString()
	// mac

	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers(vvNetworkID, nil)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers",
			err.Error(),
		)
		return
	}
	responseStruct2 := structToMap(responseGet)
	result2 := getDictResult(responseStruct2, "Mac", vvMac, simpleCmp)
	var responseVerifyItem2 merakigosdk.ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers
	if result2 != nil {
		err := mapToStruct(result2.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		//entro aqui
		data = ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersItemToBodyRs(data, &responseVerifyItem2, true)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	} else {
		resp.Diagnostics.AddWarning(
			"Resource not found",
			"Deleting resource",
		)
		resp.State.RemoveResource(ctx)
		return
	}
}

func (r *NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvTrustedServerID := data.TrustedServerID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(vvNetworkID, vvTrustedServerID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersRs
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
	vvTrustedServerID := state.TrustedServerID.ValueString()
	_, err := r.client.Switch.DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(vvNetworkID, vvTrustedServerID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersRs struct {
	NetworkID       types.String `tfsdk:"network_id"`
	TrustedServerID types.String `tfsdk:"trusted_server_id"`
	//TIENE ITEMS
	IPv4 *ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4Rs `tfsdk:"ipv4"`
	Mac  types.String                                                                         `tfsdk:"mac"`
	VLAN types.Int64                                                                          `tfsdk:"vlan"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4Rs struct {
	Address types.String `tfsdk:"address"`
}

// FromBody
func (r *NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer {
	emptyString := ""
	var requestSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerIPv4 *merakigosdk.RequestSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerIPv4
	if r.IPv4 != nil {
		address := r.IPv4.Address.ValueString()
		requestSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerIPv4 = &merakigosdk.RequestSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerIPv4{
			Address: address,
		}
	}
	mac := new(string)
	if !r.Mac.IsUnknown() && !r.Mac.IsNull() {
		*mac = r.Mac.ValueString()
	} else {
		mac = &emptyString
	}
	vLAN := new(int64)
	if !r.VLAN.IsUnknown() && !r.VLAN.IsNull() {
		*vLAN = r.VLAN.ValueInt64()
	} else {
		vLAN = nil
	}
	out := merakigosdk.RequestSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer{
		IPv4: requestSwitchCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerIPv4,
		Mac:  *mac,
		VLAN: int64ToIntPointer(vLAN),
	}
	return &out
}
func (r *NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer {
	emptyString := ""
	var requestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerIPv4 *merakigosdk.RequestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerIPv4
	if r.IPv4 != nil {
		address := r.IPv4.Address.ValueString()
		requestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerIPv4 = &merakigosdk.RequestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerIPv4{
			Address: address,
		}
	}
	mac := new(string)
	if !r.Mac.IsUnknown() && !r.Mac.IsNull() {
		*mac = r.Mac.ValueString()
	} else {
		mac = &emptyString
	}
	vLAN := new(int64)
	if !r.VLAN.IsUnknown() && !r.VLAN.IsNull() {
		*vLAN = r.VLAN.ValueInt64()
	} else {
		vLAN = nil
	}
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer{
		IPv4: requestSwitchUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerIPv4,
		Mac:  *mac,
		VLAN: int64ToIntPointer(vLAN),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersItemToBodyRs(state NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersRs, response *merakigosdk.ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers, is_read bool) NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersRs {
	itemState := NetworksSwitchDhcpServerPolicyArpInspectionTrustedServersRs{
		IPv4: func() *ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4Rs {
			if response.IPv4 != nil {
				return &ResponseItemSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersIpv4Rs{
					Address: types.StringValue(response.IPv4.Address),
				}
			}
			return nil
		}(),
		Mac:             types.StringValue(response.Mac),
		TrustedServerID: types.StringValue(response.TrustedServerID),
		VLAN: func() types.Int64 {
			if response.VLAN != nil {
				return types.Int64Value(int64(*response.VLAN))
			}
			return types.Int64{}
		}(),
		NetworkID: state.NetworkID,
	}
	state = itemState
	return state
}

func getAllItemsNetworksSwitchDhcpServerPolicyArpInspectionTrustedServers(client merakigosdk.Client, networkId string) (merakigosdk.ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers, *resty.Response, error) {
	var all_response merakigosdk.ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers
	response, r2, er := client.Switch.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers(networkId, &merakigosdk.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersQueryParams{
		PerPage: 1000,
	})
	count := 0
	all_response = append(all_response, *response...)
	for len(*response) >= 1000 {
		count += 1
		fmt.Println(count)
		links := strings.Split(r2.Header().Get("Link"), ",")
		var link string
		if count > 1 {
			link = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.Split(links[2], ";")[0], ">", ""), "<", ""), client.RestyClient().BaseURL, "")
		} else {
			link = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.Split(links[1], ";")[0], ">", ""), "<", ""), client.RestyClient().BaseURL, "")
		}
		myUrl, _ := url.Parse(link)
		params, _ := url.ParseQuery(myUrl.RawQuery)
		if params["endingBefore"] != nil {
			response, r2, er = client.Switch.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers(networkId, &merakigosdk.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersQueryParams{
				PerPage:      1000,
				EndingBefore: params["endingBefore"][0],
			})
			all_response = append(all_response, *response...)
		} else {
			break
		}
	}

	return all_response, r2, er
}
