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

// RESOURCE ACTION

import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksDevicesClaimVmxResource{}
	_ resource.ResourceWithConfigure = &NetworksDevicesClaimVmxResource{}
)

func NewNetworksDevicesClaimVmxResource() resource.Resource {
	return &NetworksDevicesClaimVmxResource{}
}

type NetworksDevicesClaimVmxResource struct {
	client *merakigosdk.Client
}

func (r *NetworksDevicesClaimVmxResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksDevicesClaimVmxResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_devices_claim_vmx"
}

// resourceAction
func (r *NetworksDevicesClaimVmxResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{

			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"size": schema.StringAttribute{
						MarkdownDescription: `The size of the vMX you claim. It can be one of: small, medium, large, 100`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
				},
			},
		},
	}
}
func (r *NetworksDevicesClaimVmxResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksDevicesClaimVmx

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
	dataRequest := data.toSdkApiRequestCreate(ctx)
	restyResp1, err := r.client.Networks.VmxNetworkDevicesClaim(vvNetworkID, dataRequest)

	if err != nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing VmxNetworkDevicesClaim",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing VmxNetworkDevicesClaim",
			err.Error(),
		)
		return
	}
	//Item

	// data2 := ResponseNetworksVmxNetworkDevicesClaim(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksDevicesClaimVmxResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksDevicesClaimVmxResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksDevicesClaimVmxResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksDevicesClaimVmx struct {
	NetworkID  types.String                             `tfsdk:"network_id"`
	Parameters *RequestNetworksVmxNetworkDevicesClaimRs `tfsdk:"parameters"`
}

type RequestNetworksVmxNetworkDevicesClaimRs struct {
	Size types.String `tfsdk:"size"`
}

// FromBody
func (r *NetworksDevicesClaimVmx) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksVmxNetworkDevicesClaim {
	emptyString := ""
	re := *r.Parameters
	size := new(string)
	if !re.Size.IsUnknown() && !re.Size.IsNull() {
		*size = re.Size.ValueString()
	} else {
		size = &emptyString
	}
	out := merakigosdk.RequestNetworksVmxNetworkDevicesClaim{
		Size: *size,
	}
	return &out
}

//ToBody
