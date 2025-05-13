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
	"strconv"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSyslogServersResource{}
	_ resource.ResourceWithConfigure = &NetworksSyslogServersResource{}
)

func NewNetworksSyslogServersResource() resource.Resource {
	return &NetworksSyslogServersResource{}
}

type NetworksSyslogServersResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSyslogServersResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSyslogServersResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_syslog_servers"
}

func (r *NetworksSyslogServersResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"servers": schema.SetNestedAttribute{
				MarkdownDescription: `List of the syslog servers for this network`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"host": schema.StringAttribute{
							MarkdownDescription: `The IP address or FQDN of the syslog server`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: `The port of the syslog server`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"roles": schema.SetAttribute{
							MarkdownDescription: `A list of roles for the syslog server. Options (case-insensitive): 'Wireless event log', 'Appliance event log', 'Switch event log', 'Air Marshal events', 'Flows', 'URLs', 'IDS alerts', 'Security events'`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
						},
					},
				},
			},
		},
	}
}

func (r *NetworksSyslogServersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSyslogServersRs

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
	//Has Item and not has items

	if vvNetworkID != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkSyslogServers(vvNetworkID)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksSyslogServers  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksSyslogServers only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkSyslogServers(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSyslogServers",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSyslogServers",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Networks.GetNetworkSyslogServers(vvNetworkID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSyslogServers",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSyslogServers",
			err.Error(),
		)
		return
	}

	data = ResponseNetworksGetNetworkSyslogServersItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *NetworksSyslogServersResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSyslogServersRs

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
	responseGet, restyRespGet, err := r.client.Networks.GetNetworkSyslogServers(vvNetworkID)
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
				"Failure when executing GetNetworkSyslogServers",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSyslogServers",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkSyslogServersItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksSyslogServersResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksSyslogServersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSyslogServersRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkSyslogServers(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSyslogServers",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSyslogServers",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSyslogServersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksSyslogServers", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSyslogServersRs struct {
	NetworkID types.String                                        `tfsdk:"network_id"`
	Servers   *[]ResponseNetworksGetNetworkSyslogServersServersRs `tfsdk:"servers"`
}

type ResponseNetworksGetNetworkSyslogServersServersRs struct {
	Host  types.String `tfsdk:"host"`
	Port  types.Int64  `tfsdk:"port"`
	Roles types.Set    `tfsdk:"roles"`
}

// FromBody
func (r *NetworksSyslogServersRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetworkSyslogServers {
	var requestNetworksUpdateNetworkSyslogServersServers []merakigosdk.RequestNetworksUpdateNetworkSyslogServersServers

	if r.Servers != nil {
		for _, rItem1 := range *r.Servers {
			host := rItem1.Host.ValueString()
			port := func() *int64 {
				if !rItem1.Port.IsUnknown() && !rItem1.Port.IsNull() {
					return rItem1.Port.ValueInt64Pointer()
				}
				return nil
			}()

			var roles []string = nil
			rItem1.Roles.ElementsAs(ctx, &roles, false)
			requestNetworksUpdateNetworkSyslogServersServers = append(requestNetworksUpdateNetworkSyslogServersServers, merakigosdk.RequestNetworksUpdateNetworkSyslogServersServers{
				Host:  host,
				Port:  int64ToIntPointer(port),
				Roles: roles,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestNetworksUpdateNetworkSyslogServers{
		Servers: func() *[]merakigosdk.RequestNetworksUpdateNetworkSyslogServersServers {
			if len(requestNetworksUpdateNetworkSyslogServersServers) > 0 {
				return &requestNetworksUpdateNetworkSyslogServersServers
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkSyslogServersItemToBodyRs(state NetworksSyslogServersRs, response *merakigosdk.ResponseNetworksGetNetworkSyslogServers, is_read bool) NetworksSyslogServersRs {
	itemState := NetworksSyslogServersRs{
		Servers: func() *[]ResponseNetworksGetNetworkSyslogServersServersRs {
			if response.Servers != nil {
				result := make([]ResponseNetworksGetNetworkSyslogServersServersRs, len(*response.Servers))
				for i, servers := range *response.Servers {
					result[i] = ResponseNetworksGetNetworkSyslogServersServersRs{
						Host: types.StringValue(servers.Host),
						Port: func() types.Int64 {
							if servers.Port != "" {
								serverPort, _ := strconv.Atoi(servers.Port)
								return types.Int64Value(int64(serverPort))
							}
							return types.Int64{}
						}(),
						Roles: StringSliceToSet(servers.Roles),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSyslogServersRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSyslogServersRs)
}
