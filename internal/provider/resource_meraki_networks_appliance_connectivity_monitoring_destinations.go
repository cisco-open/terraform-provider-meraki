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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceConnectivityMonitoringDestinationsResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceConnectivityMonitoringDestinationsResource{}
)

func NewNetworksApplianceConnectivityMonitoringDestinationsResource() resource.Resource {
	return &NetworksApplianceConnectivityMonitoringDestinationsResource{}
}

type NetworksApplianceConnectivityMonitoringDestinationsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceConnectivityMonitoringDestinationsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceConnectivityMonitoringDestinationsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_connectivity_monitoring_destinations"
}

func (r *NetworksApplianceConnectivityMonitoringDestinationsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"destinations": schema.SetNestedAttribute{
				MarkdownDescription: `The list of connectivity monitoring destinations`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"default": schema.BoolAttribute{
							MarkdownDescription: `Boolean indicating whether this is the default testing destination (true) or not (false). Defaults to false. Only one default is allowed`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
						"description": schema.StringAttribute{
							MarkdownDescription: `Description of the testing destination. Optional, defaults to an empty string`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"ip": schema.StringAttribute{
							MarkdownDescription: `The IP address to test connectivity with`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
		},
	}
}

func (r *NetworksApplianceConnectivityMonitoringDestinationsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceConnectivityMonitoringDestinationsRs

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
		responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceConnectivityMonitoringDestinations(vvNetworkID)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksApplianceConnectivityMonitoringDestinations  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksApplianceConnectivityMonitoringDestinations only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceConnectivityMonitoringDestinations(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceConnectivityMonitoringDestinations",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceConnectivityMonitoringDestinations",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceConnectivityMonitoringDestinations(vvNetworkID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceConnectivityMonitoringDestinations",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceConnectivityMonitoringDestinations",
			err.Error(),
		)
		return
	}

	data = ResponseApplianceGetNetworkApplianceConnectivityMonitoringDestinationsItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *NetworksApplianceConnectivityMonitoringDestinationsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceConnectivityMonitoringDestinationsRs

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
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceConnectivityMonitoringDestinations(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceConnectivityMonitoringDestinations",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceConnectivityMonitoringDestinations",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceConnectivityMonitoringDestinationsItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksApplianceConnectivityMonitoringDestinationsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceConnectivityMonitoringDestinationsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceConnectivityMonitoringDestinationsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceConnectivityMonitoringDestinations(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceConnectivityMonitoringDestinations",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceConnectivityMonitoringDestinations",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceConnectivityMonitoringDestinationsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceConnectivityMonitoringDestinations", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceConnectivityMonitoringDestinationsRs struct {
	NetworkID    types.String                                                                            `tfsdk:"network_id"`
	Destinations *[]ResponseApplianceGetNetworkApplianceConnectivityMonitoringDestinationsDestinationsRs `tfsdk:"destinations"`
}

type ResponseApplianceGetNetworkApplianceConnectivityMonitoringDestinationsDestinationsRs struct {
	Default     types.Bool   `tfsdk:"default"`
	Description types.String `tfsdk:"description"`
	IP          types.String `tfsdk:"ip"`
}

// FromBody
func (r *NetworksApplianceConnectivityMonitoringDestinationsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceConnectivityMonitoringDestinations {
	var requestApplianceUpdateNetworkApplianceConnectivityMonitoringDestinationsDestinations []merakigosdk.RequestApplianceUpdateNetworkApplianceConnectivityMonitoringDestinationsDestinations

	if r.Destinations != nil {
		for _, rItem1 := range *r.Destinations {
			defaultR := func() *bool {
				if !rItem1.Default.IsUnknown() && !rItem1.Default.IsNull() {
					return rItem1.Default.ValueBoolPointer()
				}
				return nil
			}()
			description := rItem1.Description.ValueString()
			ip := rItem1.IP.ValueString()
			requestApplianceUpdateNetworkApplianceConnectivityMonitoringDestinationsDestinations = append(requestApplianceUpdateNetworkApplianceConnectivityMonitoringDestinationsDestinations, merakigosdk.RequestApplianceUpdateNetworkApplianceConnectivityMonitoringDestinationsDestinations{
				Default:     defaultR,
				Description: description,
				IP:          ip,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceConnectivityMonitoringDestinations{
		Destinations: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceConnectivityMonitoringDestinationsDestinations {
			if len(requestApplianceUpdateNetworkApplianceConnectivityMonitoringDestinationsDestinations) > 0 {
				return &requestApplianceUpdateNetworkApplianceConnectivityMonitoringDestinationsDestinations
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceConnectivityMonitoringDestinationsItemToBodyRs(state NetworksApplianceConnectivityMonitoringDestinationsRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceConnectivityMonitoringDestinations, is_read bool) NetworksApplianceConnectivityMonitoringDestinationsRs {
	itemState := NetworksApplianceConnectivityMonitoringDestinationsRs{
		Destinations: func() *[]ResponseApplianceGetNetworkApplianceConnectivityMonitoringDestinationsDestinationsRs {
			if response.Destinations != nil {
				result := make([]ResponseApplianceGetNetworkApplianceConnectivityMonitoringDestinationsDestinationsRs, len(*response.Destinations))
				for i, destinations := range *response.Destinations {
					result[i] = ResponseApplianceGetNetworkApplianceConnectivityMonitoringDestinationsDestinationsRs{
						Default: func() types.Bool {
							if destinations.Default != nil {
								return types.BoolValue(*destinations.Default)
							}
							return types.Bool{}
						}(),
						Description: types.StringValue(destinations.Description),
						IP:          types.StringValue(destinations.IP),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceConnectivityMonitoringDestinationsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceConnectivityMonitoringDestinationsRs)
}
