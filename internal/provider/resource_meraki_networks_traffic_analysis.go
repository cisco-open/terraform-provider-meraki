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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksTrafficAnalysisResource{}
	_ resource.ResourceWithConfigure = &NetworksTrafficAnalysisResource{}
)

func NewNetworksTrafficAnalysisResource() resource.Resource {
	return &NetworksTrafficAnalysisResource{}
}

type NetworksTrafficAnalysisResource struct {
	client *merakigosdk.Client
}

func (r *NetworksTrafficAnalysisResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksTrafficAnalysisResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_traffic_analysis"
}

func (r *NetworksTrafficAnalysisResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"custom_pie_chart_items": schema.SetNestedAttribute{
				MarkdownDescription: `The list of items that make up the custom pie chart for traffic reporting.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the custom pie chart item.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `    The signature type for the custom pie chart item. Can be one of 'host', 'port' or 'ipRange'.
`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"value": schema.StringAttribute{
							MarkdownDescription: `    The value of the custom pie chart item. Valid syntax depends on the signature type of the chart item
    (see sample request/response for more details).
`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: `    The traffic analysis mode for the network. Can be one of 'disabled' (do not collect traffic types),
    'basic' (collect generic traffic categories), or 'detailed' (collect destination hostnames).
`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
		},
	}
}

func (r *NetworksTrafficAnalysisResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksTrafficAnalysisRs

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
	responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkTrafficAnalysis(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksTrafficAnalysis only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksTrafficAnalysis only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Networks.UpdateNetworkTrafficAnalysis(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkTrafficAnalysis",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkTrafficAnalysis",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Networks.GetNetworkTrafficAnalysis(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkTrafficAnalysis",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkTrafficAnalysis",
			err.Error(),
		)
		return
	}

	data = ResponseNetworksGetNetworkTrafficAnalysisItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksTrafficAnalysisResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksTrafficAnalysisRs

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
	responseGet, restyRespGet, err := r.client.Networks.GetNetworkTrafficAnalysis(vvNetworkID)
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
				"Failure when executing GetNetworkTrafficAnalysis",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkTrafficAnalysis",
			err.Error(),
		)
		return
	}

	data = ResponseNetworksGetNetworkTrafficAnalysisItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksTrafficAnalysisResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksTrafficAnalysisResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksTrafficAnalysisRs
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
	restyResp2, err := r.client.Networks.UpdateNetworkTrafficAnalysis(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkTrafficAnalysis",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkTrafficAnalysis",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksTrafficAnalysisResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksTrafficAnalysisRs struct {
	NetworkID           types.String                                                      `tfsdk:"network_id"`
	CustomPieChartItems *[]ResponseNetworksGetNetworkTrafficAnalysisCustomPieChartItemsRs `tfsdk:"custom_pie_chart_items"`
	Mode                types.String                                                      `tfsdk:"mode"`
}

type ResponseNetworksGetNetworkTrafficAnalysisCustomPieChartItemsRs struct {
	Name  types.String `tfsdk:"name"`
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

// FromBody
func (r *NetworksTrafficAnalysisRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetworkTrafficAnalysis {
	emptyString := ""
	var requestNetworksUpdateNetworkTrafficAnalysisCustomPieChartItems []merakigosdk.RequestNetworksUpdateNetworkTrafficAnalysisCustomPieChartItems
	if r.CustomPieChartItems != nil {
		for _, rItem1 := range *r.CustomPieChartItems {
			name := rItem1.Name.ValueString()
			typeR := rItem1.Type.ValueString()
			value := rItem1.Value.ValueString()
			requestNetworksUpdateNetworkTrafficAnalysisCustomPieChartItems = append(requestNetworksUpdateNetworkTrafficAnalysisCustomPieChartItems, merakigosdk.RequestNetworksUpdateNetworkTrafficAnalysisCustomPieChartItems{
				Name:  name,
				Type:  typeR,
				Value: value,
			})
		}
	}
	mode := new(string)
	if !r.Mode.IsUnknown() && !r.Mode.IsNull() {
		*mode = r.Mode.ValueString()
	} else {
		mode = &emptyString
	}
	out := merakigosdk.RequestNetworksUpdateNetworkTrafficAnalysis{
		CustomPieChartItems: func() *[]merakigosdk.RequestNetworksUpdateNetworkTrafficAnalysisCustomPieChartItems {
			if len(requestNetworksUpdateNetworkTrafficAnalysisCustomPieChartItems) > 0 {
				return &requestNetworksUpdateNetworkTrafficAnalysisCustomPieChartItems
			}
			return nil
		}(),
		Mode: *mode,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkTrafficAnalysisItemToBodyRs(state NetworksTrafficAnalysisRs, response *merakigosdk.ResponseNetworksGetNetworkTrafficAnalysis, is_read bool) NetworksTrafficAnalysisRs {
	itemState := NetworksTrafficAnalysisRs{
		CustomPieChartItems: func() *[]ResponseNetworksGetNetworkTrafficAnalysisCustomPieChartItemsRs {
			if response.CustomPieChartItems != nil {
				result := make([]ResponseNetworksGetNetworkTrafficAnalysisCustomPieChartItemsRs, len(*response.CustomPieChartItems))
				for i, customPieChartItems := range *response.CustomPieChartItems {
					result[i] = ResponseNetworksGetNetworkTrafficAnalysisCustomPieChartItemsRs{
						Name:  types.StringValue(customPieChartItems.Name),
						Type:  types.StringValue(customPieChartItems.Type),
						Value: types.StringValue(customPieChartItems.Value),
					}
				}
				return &result
			}
			return &[]ResponseNetworksGetNetworkTrafficAnalysisCustomPieChartItemsRs{}
		}(),
		Mode: types.StringValue(response.Mode),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksTrafficAnalysisRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksTrafficAnalysisRs)
}
