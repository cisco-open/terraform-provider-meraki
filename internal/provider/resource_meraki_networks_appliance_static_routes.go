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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceStaticRoutesResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceStaticRoutesResource{}
)

func NewNetworksApplianceStaticRoutesResource() resource.Resource {
	return &NetworksApplianceStaticRoutesResource{}
}

type NetworksApplianceStaticRoutesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceStaticRoutesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceStaticRoutesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_static_routes"
}

func (r *NetworksApplianceStaticRoutesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `Whether the route is enabled or not`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"fixed_ip_assignments": schema.SingleNestedAttribute{
				MarkdownDescription: `The DHCP fixed IP assignments on the static route. This should be an object that contains mappings from MAC addresses to objects that themselves each contain "ip" and "name" string fields. See the sample request/response for more details.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"attribute_22_33_44_55_66_77": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"ip": schema.StringAttribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
						},
					},
				},
			},
			"gateway_ip": schema.StringAttribute{
				MarkdownDescription: `Gateway IP address (next hop)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"gateway_vlan_id": schema.Int64Attribute{
				MarkdownDescription: `Gateway VLAN ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				//            Differents_types: `   parameter: schema.TypeString, item: schema.TypeInt`,
			},
			"gateway_vlan_id_rs": schema.StringAttribute{
				MarkdownDescription: `Gateway VLAN ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				//            Differents_types: `   parameter: schema.TypeString, item: schema.TypeInt`,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: `Route ID`,
				Computed:            true,
			},
			"ip_version": schema.Int64Attribute{
				MarkdownDescription: `IP protocol version`,
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Name of the route`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `Network ID`,
				Required:            true,
			},
			"reserved_ip_ranges": schema.SetNestedAttribute{
				MarkdownDescription: `DHCP reserved IP ranges`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"comment": schema.StringAttribute{
							MarkdownDescription: `Description of the range`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"end": schema.StringAttribute{
							MarkdownDescription: `Last address in the reserved range`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"start": schema.StringAttribute{
							MarkdownDescription: `First address in the reserved range`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"static_route_id": schema.StringAttribute{
				MarkdownDescription: `staticRouteId path parameter. Static route ID`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"subnet": schema.StringAttribute{
				MarkdownDescription: `Subnet of the route`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

//path params to set ['staticRouteId']

func (r *NetworksApplianceStaticRoutesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceStaticRoutesRs

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
	//Has Item and has items and post

	vvName := data.Name.ValueString()

	responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceStaticRoutes(vvNetworkID)
	//Has Post
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkApplianceStaticRoutes",
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
			vvStaticRouteID, ok := result2["ID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter StaticRouteID",
					"Fail Parsing StaticRouteID",
				)
				return
			}
			r.client.Appliance.UpdateNetworkApplianceStaticRoute(vvNetworkID, vvStaticRouteID, data.toSdkApiRequestUpdate(ctx))

			responseVerifyItem2, _, _ := r.client.Appliance.GetNetworkApplianceStaticRoute(vvNetworkID, vvStaticRouteID)
			if responseVerifyItem2 != nil {
				data = ResponseApplianceGetNetworkApplianceStaticRouteItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Appliance.CreateNetworkApplianceStaticRoute(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkApplianceStaticRoute",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkApplianceStaticRoute",
			err.Error(),
		)
		return
	}

	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceStaticRoutes(vvNetworkID)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceStaticRoutes",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceStaticRoutes",
			err.Error(),
		)
		return
	}

	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvStaticRouteID, ok := result2["ID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter StaticRouteID",
				"Fail Parsing StaticRouteID",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Appliance.GetNetworkApplianceStaticRoute(vvNetworkID, vvStaticRouteID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseApplianceGetNetworkApplianceStaticRouteItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkApplianceStaticRoute",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceStaticRoute",
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

func (r *NetworksApplianceStaticRoutesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceStaticRoutesRs

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
	vvStaticRouteID := data.StaticRouteID.ValueString()
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceStaticRoute(vvNetworkID, vvStaticRouteID)
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
				"Failure when executing GetNetworkApplianceStaticRoute",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceStaticRoute",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceStaticRouteItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceStaticRoutesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("static_route_id"), idParts[1])...)
}

func (r *NetworksApplianceStaticRoutesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceStaticRoutesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvStaticRouteID := data.StaticRouteID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceStaticRoute(vvNetworkID, vvStaticRouteID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceStaticRoute",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceStaticRoute",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceStaticRoutesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksApplianceStaticRoutesRs
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
	vvStaticRouteID := state.StaticRouteID.ValueString()
	_, err := r.client.Appliance.DeleteNetworkApplianceStaticRoute(vvNetworkID, vvStaticRouteID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkApplianceStaticRoute", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksApplianceStaticRoutesRs struct {
	NetworkID          types.String                                                         `tfsdk:"network_id"`
	StaticRouteID      types.String                                                         `tfsdk:"static_route_id"`
	Enabled            types.Bool                                                           `tfsdk:"enabled"`
	FixedIPAssignments *ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignmentsRs `tfsdk:"fixed_ip_assignments"`
	GatewayIP          types.String                                                         `tfsdk:"gateway_ip"`
	GatewayVLANID      types.Int64                                                          `tfsdk:"gateway_vlan_id"`
	GatewayVLANIDRs    types.String                                                         `tfsdk:"gateway_vlan_id_rs"`
	ID                 types.String                                                         `tfsdk:"id"`
	IPVersion          types.Int64                                                          `tfsdk:"ip_version"`
	Name               types.String                                                         `tfsdk:"name"`
	ReservedIPRanges   *[]ResponseApplianceGetNetworkApplianceStaticRouteReservedIpRangesRs `tfsdk:"reserved_ip_ranges"`
	Subnet             types.String                                                         `tfsdk:"subnet"`
}

type ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignmentsRs struct {
	Status223344556677 *ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignments223344556677Rs `tfsdk:"attribute_22_33_44_55_66_77"`
}

type ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignments223344556677Rs struct {
	IP   types.String `tfsdk:"ip"`
	Name types.String `tfsdk:"name"`
}

type ResponseApplianceGetNetworkApplianceStaticRouteReservedIpRangesRs struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

// FromBody
func (r *NetworksApplianceStaticRoutesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestApplianceCreateNetworkApplianceStaticRoute {
	emptyString := ""
	gatewayIP := new(string)
	if !r.GatewayIP.IsUnknown() && !r.GatewayIP.IsNull() {
		*gatewayIP = r.GatewayIP.ValueString()
	} else {
		gatewayIP = &emptyString
	}
	gatewayVLANID := new(string)
	if !r.GatewayVLANID.IsUnknown() && !r.GatewayVLANID.IsNull() {
		*gatewayVLANID = r.GatewayVLANIDRs.ValueString()
	} else {
		gatewayVLANID = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	subnet := new(string)
	if !r.Subnet.IsUnknown() && !r.Subnet.IsNull() {
		*subnet = r.Subnet.ValueString()
	} else {
		subnet = &emptyString
	}
	out := merakigosdk.RequestApplianceCreateNetworkApplianceStaticRoute{
		GatewayIP:     *gatewayIP,
		GatewayVLANID: *gatewayVLANID,
		Name:          *name,
		Subnet:        *subnet,
	}
	return &out
}
func (r *NetworksApplianceStaticRoutesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceStaticRoute {
	emptyString := ""
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	// var requestApplianceUpdateNetworkApplianceStaticRouteFixedIPAssignments *merakigosdk.RequestApplianceUpdateNetworkApplianceStaticRouteFixedIPAssignments

	// if r.FixedIPAssignments != nil {
	// 	requestApplianceUpdateNetworkApplianceStaticRouteFixedIPAssignments = &merakigosdk.RequestApplianceUpdateNetworkApplianceStaticRouteFixedIPAssignments{}
	// 	//[debug] Is Array: False
	// }
	gatewayIP := new(string)
	if !r.GatewayIP.IsUnknown() && !r.GatewayIP.IsNull() {
		*gatewayIP = r.GatewayIP.ValueString()
	} else {
		gatewayIP = &emptyString
	}
	gatewayVLANID := new(string)
	if !r.GatewayVLANIDRs.IsUnknown() && !r.GatewayVLANIDRs.IsNull() {
		*gatewayVLANID = r.GatewayVLANIDRs.ValueString()
	} else {
		gatewayVLANID = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestApplianceUpdateNetworkApplianceStaticRouteReservedIPRanges []merakigosdk.RequestApplianceUpdateNetworkApplianceStaticRouteReservedIPRanges

	if r.ReservedIPRanges != nil {
		for _, rItem1 := range *r.ReservedIPRanges {
			comment := rItem1.Comment.ValueString()
			end := rItem1.End.ValueString()
			start := rItem1.Start.ValueString()
			requestApplianceUpdateNetworkApplianceStaticRouteReservedIPRanges = append(requestApplianceUpdateNetworkApplianceStaticRouteReservedIPRanges, merakigosdk.RequestApplianceUpdateNetworkApplianceStaticRouteReservedIPRanges{
				Comment: comment,
				End:     end,
				Start:   start,
			})
			//[debug] Is Array: True
		}
	}
	subnet := new(string)
	if !r.Subnet.IsUnknown() && !r.Subnet.IsNull() {
		*subnet = r.Subnet.ValueString()
	} else {
		subnet = &emptyString
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceStaticRoute{
		Enabled: enabled,
		// FixedIPAssignments: requestApplianceUpdateNetworkApplianceStaticRouteFixedIPAssignments,
		GatewayIP:     *gatewayIP,
		GatewayVLANID: *gatewayVLANID,
		Name:          *name,
		ReservedIPRanges: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceStaticRouteReservedIPRanges {
			if len(requestApplianceUpdateNetworkApplianceStaticRouteReservedIPRanges) > 0 {
				return &requestApplianceUpdateNetworkApplianceStaticRouteReservedIPRanges
			}
			return nil
		}(),
		Subnet: *subnet,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceStaticRouteItemToBodyRs(state NetworksApplianceStaticRoutesRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceStaticRoute, is_read bool) NetworksApplianceStaticRoutesRs {
	itemState := NetworksApplianceStaticRoutesRs{
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		FixedIPAssignments: func() *ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignmentsRs {
			if response.FixedIPAssignments != nil {
				return &ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignmentsRs{
					Status223344556677: func() *ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignments223344556677Rs {
						if response.FixedIPAssignments.Status223344556677 != nil {
							return &ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignments223344556677Rs{
								IP:   types.StringValue(response.FixedIPAssignments.Status223344556677.IP),
								Name: types.StringValue(response.FixedIPAssignments.Status223344556677.Name),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		GatewayIP: types.StringValue(response.GatewayIP),
		GatewayVLANID: func() types.Int64 {
			if response.GatewayVLANID != nil {
				return types.Int64Value(int64(*response.GatewayVLANID))
			}
			return types.Int64{}
		}(),
		ID: types.StringValue(response.ID),
		IPVersion: func() types.Int64 {
			if response.IPVersion != nil {
				return types.Int64Value(int64(*response.IPVersion))
			}
			return types.Int64{}
		}(),
		Name:      types.StringValue(response.Name),
		NetworkID: types.StringValue(response.NetworkID),
		ReservedIPRanges: func() *[]ResponseApplianceGetNetworkApplianceStaticRouteReservedIpRangesRs {
			if response.ReservedIPRanges != nil {
				result := make([]ResponseApplianceGetNetworkApplianceStaticRouteReservedIpRangesRs, len(*response.ReservedIPRanges))
				for i, reservedIPRanges := range *response.ReservedIPRanges {
					result[i] = ResponseApplianceGetNetworkApplianceStaticRouteReservedIpRangesRs{
						Comment: types.StringValue(reservedIPRanges.Comment),
						End:     types.StringValue(reservedIPRanges.End),
						Start:   types.StringValue(reservedIPRanges.Start),
					}
				}
				return &result
			}
			return nil
		}(),
		Subnet:        types.StringValue(response.Subnet),
		StaticRouteID: types.StringValue(response.ID),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceStaticRoutesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceStaticRoutesRs)
}
