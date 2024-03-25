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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksCellularGatewaySubnetPoolResource{}
	_ resource.ResourceWithConfigure = &NetworksCellularGatewaySubnetPoolResource{}
)

func NewNetworksCellularGatewaySubnetPoolResource() resource.Resource {
	return &NetworksCellularGatewaySubnetPoolResource{}
}

type NetworksCellularGatewaySubnetPoolResource struct {
	client *merakigosdk.Client
}

func (r *NetworksCellularGatewaySubnetPoolResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksCellularGatewaySubnetPoolResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_cellular_gateway_subnet_pool"
}

func (r *NetworksCellularGatewaySubnetPoolResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"cidr": schema.StringAttribute{
				MarkdownDescription: `CIDR of the pool of subnets. Each MG in this network will automatically pick a subnet from this pool.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"deployment_mode": schema.StringAttribute{
				Computed: true,
			},
			"mask": schema.Int64Attribute{
				MarkdownDescription: `Mask used for the subnet of all MGs in  this network.`,
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
			"subnets": schema.SetNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"appliance_ip": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"serial": schema.StringAttribute{
							Computed: true,
						},
						"subnet": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (r *NetworksCellularGatewaySubnetPoolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksCellularGatewaySubnetPoolRs

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
	responseVerifyItem, restyResp1, err := r.client.CellularGateway.GetNetworkCellularGatewaySubnetPool(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksCellularGatewaySubnetPool only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksCellularGatewaySubnetPool only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.CellularGateway.UpdateNetworkCellularGatewaySubnetPool(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkCellularGatewaySubnetPool",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkCellularGatewaySubnetPool",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.CellularGateway.GetNetworkCellularGatewaySubnetPool(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkCellularGatewaySubnetPool",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkCellularGatewaySubnetPool",
			err.Error(),
		)
		return
	}

	data = ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksCellularGatewaySubnetPoolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksCellularGatewaySubnetPoolRs

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
	responseGet, restyRespGet, err := r.client.CellularGateway.GetNetworkCellularGatewaySubnetPool(vvNetworkID)
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
				"Failure when executing GetNetworkCellularGatewaySubnetPool",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkCellularGatewaySubnetPool",
			err.Error(),
		)
		return
	}

	data = ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksCellularGatewaySubnetPoolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksCellularGatewaySubnetPoolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksCellularGatewaySubnetPoolRs
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
	restyResp2, err := r.client.CellularGateway.UpdateNetworkCellularGatewaySubnetPool(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkCellularGatewaySubnetPool",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkCellularGatewaySubnetPool",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksCellularGatewaySubnetPoolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksCellularGatewaySubnetPoolRs struct {
	NetworkID      types.String                                                           `tfsdk:"network_id"`
	Cidr           types.String                                                           `tfsdk:"cidr"`
	DeploymentMode types.String                                                           `tfsdk:"deployment_mode"`
	Mask           types.Int64                                                            `tfsdk:"mask"`
	Subnets        *[]ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnetsRs `tfsdk:"subnets"`
}

type ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnetsRs struct {
	ApplianceIP types.String `tfsdk:"appliance_ip"`
	Name        types.String `tfsdk:"name"`
	Serial      types.String `tfsdk:"serial"`
	Subnet      types.String `tfsdk:"subnet"`
}

// FromBody
func (r *NetworksCellularGatewaySubnetPoolRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestCellularGatewayUpdateNetworkCellularGatewaySubnetPool {
	emptyString := ""
	cidr := new(string)
	if !r.Cidr.IsUnknown() && !r.Cidr.IsNull() {
		*cidr = r.Cidr.ValueString()
	} else {
		cidr = &emptyString
	}
	mask := new(int64)
	if !r.Mask.IsUnknown() && !r.Mask.IsNull() {
		*mask = r.Mask.ValueInt64()
	} else {
		mask = nil
	}
	out := merakigosdk.RequestCellularGatewayUpdateNetworkCellularGatewaySubnetPool{
		Cidr: *cidr,
		Mask: int64ToIntPointer(mask),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolItemToBodyRs(state NetworksCellularGatewaySubnetPoolRs, response *merakigosdk.ResponseCellularGatewayGetNetworkCellularGatewaySubnetPool, is_read bool) NetworksCellularGatewaySubnetPoolRs {
	itemState := NetworksCellularGatewaySubnetPoolRs{
		Cidr:           types.StringValue(response.Cidr),
		DeploymentMode: types.StringValue(response.DeploymentMode),
		Mask: func() types.Int64 {
			if response.Mask != nil {
				return types.Int64Value(int64(*response.Mask))
			}
			return types.Int64{}
		}(),
		Subnets: func() *[]ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnetsRs {
			if response.Subnets != nil {
				result := make([]ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnetsRs, len(*response.Subnets))
				for i, subnets := range *response.Subnets {
					result[i] = ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnetsRs{
						ApplianceIP: types.StringValue(subnets.ApplianceIP),
						Name:        types.StringValue(subnets.Name),
						Serial:      types.StringValue(subnets.Serial),
						Subnet:      types.StringValue(subnets.Subnet),
					}
				}
				return &result
			}
			return &[]ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnetsRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksCellularGatewaySubnetPoolRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksCellularGatewaySubnetPoolRs)
}
