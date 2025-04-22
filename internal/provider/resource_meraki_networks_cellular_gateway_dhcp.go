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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksCellularGatewayDhcpResource{}
	_ resource.ResourceWithConfigure = &NetworksCellularGatewayDhcpResource{}
)

func NewNetworksCellularGatewayDhcpResource() resource.Resource {
	return &NetworksCellularGatewayDhcpResource{}
}

type NetworksCellularGatewayDhcpResource struct {
	client *merakigosdk.Client
}

func (r *NetworksCellularGatewayDhcpResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksCellularGatewayDhcpResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_cellular_gateway_dhcp"
}

func (r *NetworksCellularGatewayDhcpResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"dhcp_lease_time": schema.StringAttribute{
				MarkdownDescription: `DHCP Lease time for all MG in the network.
                                  Allowed values: [1 day,1 hour,1 week,12 hours,30 minutes,4 hours]`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"1 day",
						"1 hour",
						"1 week",
						"12 hours",
						"30 minutes",
						"4 hours",
					),
				},
			},
			"dns_custom_nameservers": schema.SetAttribute{
				MarkdownDescription: `List of fixed IPs representing the the DNS Name servers when the mode is 'custom'.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"dns_nameservers": schema.StringAttribute{
				MarkdownDescription: `DNS name servers mode for all MG in the network.
                                  Allowed values: [custom,google_dns,opendns,upstream_dns]`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"custom",
						"google_dns",
						"opendns",
						"upstream_dns",
					),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
		},
	}
}

func (r *NetworksCellularGatewayDhcpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksCellularGatewayDhcpRs

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
		responseVerifyItem, restyResp1, err := r.client.CellularGateway.GetNetworkCellularGatewayDhcp(vvNetworkID)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksCellularGatewayDhcp  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksCellularGatewayDhcp only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.CellularGateway.UpdateNetworkCellularGatewayDhcp(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkCellularGatewayDhcp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkCellularGatewayDhcp",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.CellularGateway.GetNetworkCellularGatewayDhcp(vvNetworkID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkCellularGatewayDhcp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkCellularGatewayDhcp",
			err.Error(),
		)
		return
	}

	data = ResponseCellularGatewayGetNetworkCellularGatewayDhcpItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *NetworksCellularGatewayDhcpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksCellularGatewayDhcpRs

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
	responseGet, restyRespGet, err := r.client.CellularGateway.GetNetworkCellularGatewayDhcp(vvNetworkID)
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
				"Failure when executing GetNetworkCellularGatewayDhcp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkCellularGatewayDhcp",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseCellularGatewayGetNetworkCellularGatewayDhcpItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksCellularGatewayDhcpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksCellularGatewayDhcpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksCellularGatewayDhcpRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.CellularGateway.UpdateNetworkCellularGatewayDhcp(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkCellularGatewayDhcp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkCellularGatewayDhcp",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksCellularGatewayDhcpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksCellularGatewayDhcp", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksCellularGatewayDhcpRs struct {
	NetworkID            types.String `tfsdk:"network_id"`
	DhcpLeaseTime        types.String `tfsdk:"dhcp_lease_time"`
	DNSCustomNameservers types.Set    `tfsdk:"dns_custom_nameservers"`
	DNSNameservers       types.String `tfsdk:"dns_nameservers"`
}

// FromBody
func (r *NetworksCellularGatewayDhcpRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestCellularGatewayUpdateNetworkCellularGatewayDhcp {
	emptyString := ""
	dhcpLeaseTime := new(string)
	if !r.DhcpLeaseTime.IsUnknown() && !r.DhcpLeaseTime.IsNull() {
		*dhcpLeaseTime = r.DhcpLeaseTime.ValueString()
	} else {
		dhcpLeaseTime = &emptyString
	}
	var dNSCustomNameservers []string = nil
	r.DNSCustomNameservers.ElementsAs(ctx, &dNSCustomNameservers, false)
	dNSNameservers := new(string)
	if !r.DNSNameservers.IsUnknown() && !r.DNSNameservers.IsNull() {
		*dNSNameservers = r.DNSNameservers.ValueString()
	} else {
		dNSNameservers = &emptyString
	}
	out := merakigosdk.RequestCellularGatewayUpdateNetworkCellularGatewayDhcp{
		DhcpLeaseTime:        *dhcpLeaseTime,
		DNSCustomNameservers: dNSCustomNameservers,
		DNSNameservers:       *dNSNameservers,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseCellularGatewayGetNetworkCellularGatewayDhcpItemToBodyRs(state NetworksCellularGatewayDhcpRs, response *merakigosdk.ResponseCellularGatewayGetNetworkCellularGatewayDhcp, is_read bool) NetworksCellularGatewayDhcpRs {
	itemState := NetworksCellularGatewayDhcpRs{
		DhcpLeaseTime:        types.StringValue(response.DhcpLeaseTime),
		DNSCustomNameservers: StringSliceToSet(response.DNSCustomNameservers),
		DNSNameservers:       types.StringValue(response.DNSNameservers),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksCellularGatewayDhcpRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksCellularGatewayDhcpRs)
}
