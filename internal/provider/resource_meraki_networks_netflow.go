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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksNetflowResource{}
	_ resource.ResourceWithConfigure = &NetworksNetflowResource{}
)

func NewNetworksNetflowResource() resource.Resource {
	return &NetworksNetflowResource{}
}

type NetworksNetflowResource struct {
	client *merakigosdk.Client
}

func (r *NetworksNetflowResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksNetflowResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_netflow"
}

func (r *NetworksNetflowResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"collector_ip": schema.StringAttribute{
				MarkdownDescription: `The IPv4 address of the NetFlow collector.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"collector_port": schema.Int64Attribute{
				MarkdownDescription: `The port that the NetFlow collector will be listening on.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"eta_dst_port": schema.Int64Attribute{
				MarkdownDescription: `The port that the Encrypted Traffic Analytics collector will be listening on.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"eta_enabled": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating whether Encrypted Traffic Analytics is enabled (true) or disabled (false).`,
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
			"reporting_enabled": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating whether NetFlow traffic reporting is enabled (true) or disabled (false).`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *NetworksNetflowResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksNetflowRs

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
	//Item
	responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkNetflow(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksNetflow only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksNetflow only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkNetflow(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkNetflow",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkNetflow",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Networks.GetNetworkNetflow(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkNetflow",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkNetflow",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkNetflowItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksNetflowResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksNetflowRs

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
	responseGet, restyRespGet, err := r.client.Networks.GetNetworkNetflow(vvNetworkID)
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
				"Failure when executing GetNetworkNetflow",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkNetflow",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkNetflowItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksNetflowResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksNetflowResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksNetflowRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkNetflow(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkNetflow",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkNetflow",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksNetflowResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksNetflow", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksNetflowRs struct {
	NetworkID        types.String `tfsdk:"network_id"`
	CollectorIP      types.String `tfsdk:"collector_ip"`
	CollectorPort    types.Int64  `tfsdk:"collector_port"`
	EtaDstPort       types.Int64  `tfsdk:"eta_dst_port"`
	EtaEnabled       types.Bool   `tfsdk:"eta_enabled"`
	ReportingEnabled types.Bool   `tfsdk:"reporting_enabled"`
}

// FromBody
func (r *NetworksNetflowRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetworkNetflow {
	emptyString := ""
	collectorIP := new(string)
	if !r.CollectorIP.IsUnknown() && !r.CollectorIP.IsNull() {
		*collectorIP = r.CollectorIP.ValueString()
	} else {
		collectorIP = &emptyString
	}
	collectorPort := new(int64)
	if !r.CollectorPort.IsUnknown() && !r.CollectorPort.IsNull() {
		*collectorPort = r.CollectorPort.ValueInt64()
	} else {
		collectorPort = nil
	}
	etaDstPort := new(int64)
	if !r.EtaDstPort.IsUnknown() && !r.EtaDstPort.IsNull() {
		*etaDstPort = r.EtaDstPort.ValueInt64()
	} else {
		etaDstPort = nil
	}
	etaEnabled := new(bool)
	if !r.EtaEnabled.IsUnknown() && !r.EtaEnabled.IsNull() {
		*etaEnabled = r.EtaEnabled.ValueBool()
	} else {
		etaEnabled = nil
	}
	reportingEnabled := new(bool)
	if !r.ReportingEnabled.IsUnknown() && !r.ReportingEnabled.IsNull() {
		*reportingEnabled = r.ReportingEnabled.ValueBool()
	} else {
		reportingEnabled = nil
	}
	out := merakigosdk.RequestNetworksUpdateNetworkNetflow{
		CollectorIP:      *collectorIP,
		CollectorPort:    int64ToIntPointer(collectorPort),
		EtaDstPort:       int64ToIntPointer(etaDstPort),
		EtaEnabled:       etaEnabled,
		ReportingEnabled: reportingEnabled,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkNetflowItemToBodyRs(state NetworksNetflowRs, response *merakigosdk.ResponseNetworksGetNetworkNetflow, is_read bool) NetworksNetflowRs {
	itemState := NetworksNetflowRs{
		CollectorIP: types.StringValue(response.CollectorIP),
		CollectorPort: func() types.Int64 {
			if response.CollectorPort != nil {
				return types.Int64Value(int64(*response.CollectorPort))
			}
			return types.Int64{}
		}(),
		EtaDstPort: func() types.Int64 {
			if response.EtaDstPort != nil {
				return types.Int64Value(int64(*response.EtaDstPort))
			}
			return types.Int64{}
		}(),
		EtaEnabled: func() types.Bool {
			if response.EtaEnabled != nil {
				return types.BoolValue(*response.EtaEnabled)
			}
			return types.Bool{}
		}(),
		ReportingEnabled: func() types.Bool {
			if response.ReportingEnabled != nil {
				return types.BoolValue(*response.ReportingEnabled)
			}
			return types.Bool{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksNetflowRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksNetflowRs)
}
