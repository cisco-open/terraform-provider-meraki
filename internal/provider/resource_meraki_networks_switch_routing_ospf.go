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

	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSwitchRoutingOspfResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchRoutingOspfResource{}
)

func NewNetworksSwitchRoutingOspfResource() resource.Resource {
	return &NetworksSwitchRoutingOspfResource{}
}

type NetworksSwitchRoutingOspfResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchRoutingOspfResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchRoutingOspfResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_routing_ospf"
}

func (r *NetworksSwitchRoutingOspfResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"areas": schema.ListNestedAttribute{
				MarkdownDescription: `OSPF areas`,
				Optional:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"area_id": schema.StringAttribute{
							MarkdownDescription: `OSPF area ID`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"area_name": schema.StringAttribute{
							MarkdownDescription: `Name of the OSPF area`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"area_type": schema.StringAttribute{
							MarkdownDescription: `Area types in OSPF. Must be one of: ["normal", "stub", "nssa"]
                                        Allowed values: [normal,nssa,stub]`,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"normal",
									"nssa",
									"stub",
								),
							},
						},
					},
				},
			},
			"dead_timer_in_seconds": schema.Int64Attribute{
				MarkdownDescription: `Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `Boolean value to enable or disable OSPF routing. OSPF routing is disabled by default.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"hello_timer_in_seconds": schema.Int64Attribute{
				MarkdownDescription: `Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"md5_authentication_enabled": schema.BoolAttribute{
				MarkdownDescription: `Boolean value to enable or disable MD5 authentication. MD5 authentication is disabled by default.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"md5_authentication_key": schema.SingleNestedAttribute{
				MarkdownDescription: `MD5 authentication credentials. This param is only relevant if md5AuthenticationEnabled is true`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"id": schema.Int64Attribute{
						MarkdownDescription: `MD5 authentication key index. Key index must be between 1 to 255`,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"passphrase": schema.StringAttribute{
						MarkdownDescription: `MD5 authentication passphrase`,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"v3": schema.SingleNestedAttribute{
				MarkdownDescription: `OSPF v3 configuration`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"areas": schema.ListNestedAttribute{
						MarkdownDescription: `OSPF v3 areas`,
						Optional:            true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"area_id": schema.StringAttribute{
									MarkdownDescription: `OSPF area ID`,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"area_name": schema.StringAttribute{
									MarkdownDescription: `Name of the OSPF area`,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"area_type": schema.StringAttribute{
									MarkdownDescription: `Area types in OSPF. Must be one of: ["normal", "stub", "nssa"]
                                              Allowed values: [normal,nssa,stub]`,
									Optional: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
									Validators: []validator.String{
										stringvalidator.OneOf(
											"normal",
											"nssa",
											"stub",
										),
									},
								},
							},
						},
					},
					"dead_timer_in_seconds": schema.Int64Attribute{
						MarkdownDescription: `Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535`,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Boolean value to enable or disable V3 OSPF routing. OSPF V3 routing is disabled by default.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"hello_timer_in_seconds": schema.Int64Attribute{
						MarkdownDescription: `Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
		},
	}
}

func (r *NetworksSwitchRoutingOspfResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchRoutingOspfRs

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

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchRoutingOspf(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchRoutingOspf",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchRoutingOspf",
			err.Error(),
		)
		return
	}

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *NetworksSwitchRoutingOspfResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchRoutingOspfRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	responseGet, restyRespGet, err := r.client.Switch.GetNetworkSwitchRoutingOspf(vvNetworkID)
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
				"Failure when executing GetNetworkSwitchRoutingOspf",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchRoutingOspf",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSwitchGetNetworkSwitchRoutingOspfItemToBodyRs(data, responseGet, true)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksSwitchRoutingOspfResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksSwitchRoutingOspfResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksSwitchRoutingOspfRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchRoutingOspf(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchRoutingOspf",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchRoutingOspf",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworksSwitchRoutingOspfResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksSwitchRoutingOspf", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSwitchRoutingOspfRs struct {
	NetworkID                types.String                                                     `tfsdk:"network_id"`
	Areas                    *[]ResponseSwitchGetNetworkSwitchRoutingOspfAreasRs              `tfsdk:"areas"`
	DeadTimerInSeconds       types.Int64                                                      `tfsdk:"dead_timer_in_seconds"`
	Enabled                  types.Bool                                                       `tfsdk:"enabled"`
	HelloTimerInSeconds      types.Int64                                                      `tfsdk:"hello_timer_in_seconds"`
	Md5AuthenticationEnabled types.Bool                                                       `tfsdk:"md5_authentication_enabled"`
	Md5AuthenticationKey     *ResponseSwitchGetNetworkSwitchRoutingOspfMd5AuthenticationKeyRs `tfsdk:"md5_authentication_key"`
	V3                       *ResponseSwitchGetNetworkSwitchRoutingOspfV3Rs                   `tfsdk:"v3"`
}

type ResponseSwitchGetNetworkSwitchRoutingOspfAreasRs struct {
	Areaid   types.String `tfsdk:"area_id"`
	AreaName types.String `tfsdk:"area_name"`
	AreaType types.String `tfsdk:"area_type"`
}

type ResponseSwitchGetNetworkSwitchRoutingOspfMd5AuthenticationKeyRs struct {
	ID         types.Int64  `tfsdk:"id"`
	Passphrase types.String `tfsdk:"passphrase"`
}

type ResponseSwitchGetNetworkSwitchRoutingOspfV3Rs struct {
	Areas               *[]ResponseSwitchGetNetworkSwitchRoutingOspfV3AreasRs `tfsdk:"areas"`
	DeadTimerInSeconds  types.Int64                                           `tfsdk:"dead_timer_in_seconds"`
	Enabled             types.Bool                                            `tfsdk:"enabled"`
	HelloTimerInSeconds types.Int64                                           `tfsdk:"hello_timer_in_seconds"`
}

type ResponseSwitchGetNetworkSwitchRoutingOspfV3AreasRs struct {
	Areaid   types.String `tfsdk:"area_id"`
	AreaName types.String `tfsdk:"area_name"`
	AreaType types.String `tfsdk:"area_type"`
}

// FromBody
func (r *NetworksSwitchRoutingOspfRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspf {
	var requestSwitchUpdateNetworkSwitchRoutingOspfAreas []merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfAreas

	if r.Areas != nil {
		for _, rItem1 := range *r.Areas {
			areaID := rItem1.Areaid.ValueString()
			areaName := rItem1.AreaName.ValueString()
			areaType := rItem1.AreaType.ValueString()
			requestSwitchUpdateNetworkSwitchRoutingOspfAreas = append(requestSwitchUpdateNetworkSwitchRoutingOspfAreas, merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfAreas{
				AreaID:   areaID,
				AreaName: areaName,
				AreaType: areaType,
			})
			//[debug] Is Array: True
		}
	}
	deadTimerInSeconds := new(int64)
	if !r.DeadTimerInSeconds.IsUnknown() && !r.DeadTimerInSeconds.IsNull() {
		*deadTimerInSeconds = r.DeadTimerInSeconds.ValueInt64()
	} else {
		deadTimerInSeconds = nil
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	helloTimerInSeconds := new(int64)
	if !r.HelloTimerInSeconds.IsUnknown() && !r.HelloTimerInSeconds.IsNull() {
		*helloTimerInSeconds = r.HelloTimerInSeconds.ValueInt64()
	} else {
		helloTimerInSeconds = nil
	}
	md5AuthenticationEnabled := new(bool)
	if !r.Md5AuthenticationEnabled.IsUnknown() && !r.Md5AuthenticationEnabled.IsNull() {
		*md5AuthenticationEnabled = r.Md5AuthenticationEnabled.ValueBool()
	} else {
		md5AuthenticationEnabled = nil
	}
	var requestSwitchUpdateNetworkSwitchRoutingOspfMd5AuthenticationKey *merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfMd5AuthenticationKey

	if r.Md5AuthenticationKey != nil {
		id := func() *int64 {
			if !r.Md5AuthenticationKey.ID.IsUnknown() && !r.Md5AuthenticationKey.ID.IsNull() {
				return r.Md5AuthenticationKey.ID.ValueInt64Pointer()
			}
			return nil
		}()
		passphrase := r.Md5AuthenticationKey.Passphrase.ValueString()
		requestSwitchUpdateNetworkSwitchRoutingOspfMd5AuthenticationKey = &merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfMd5AuthenticationKey{
			ID:         int64ToIntPointer(id),
			Passphrase: passphrase,
		}
		//[debug] Is Array: False
	}
	var requestSwitchUpdateNetworkSwitchRoutingOspfV3 *merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfV3

	if r.V3 != nil {

		log.Printf("[DEBUG] #TODO []RequestSwitchUpdateNetworkSwitchRoutingOspfV3Areas")
		var requestSwitchUpdateNetworkSwitchRoutingOspfV3Areas []merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfV3Areas

		if r.V3.Areas != nil {
			for _, rItem1 := range *r.V3.Areas {
				areaID := rItem1.Areaid.ValueString()
				areaName := rItem1.AreaName.ValueString()
				areaType := rItem1.AreaType.ValueString()
				requestSwitchUpdateNetworkSwitchRoutingOspfV3Areas = append(requestSwitchUpdateNetworkSwitchRoutingOspfV3Areas, merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfV3Areas{
					AreaID:   areaID,
					AreaName: areaName,
					AreaType: areaType,
				})
				//[debug] Is Array: True
			}
		}
		deadTimerInSeconds := func() *int64 {
			if !r.V3.DeadTimerInSeconds.IsUnknown() && !r.V3.DeadTimerInSeconds.IsNull() {
				return r.V3.DeadTimerInSeconds.ValueInt64Pointer()
			}
			return nil
		}()
		enabled := func() *bool {
			if !r.V3.Enabled.IsUnknown() && !r.V3.Enabled.IsNull() {
				return r.V3.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		helloTimerInSeconds := func() *int64 {
			if !r.V3.HelloTimerInSeconds.IsUnknown() && !r.V3.HelloTimerInSeconds.IsNull() {
				return r.V3.HelloTimerInSeconds.ValueInt64Pointer()
			}
			return nil
		}()
		requestSwitchUpdateNetworkSwitchRoutingOspfV3 = &merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfV3{
			Areas: func() *[]merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfV3Areas {
				if len(requestSwitchUpdateNetworkSwitchRoutingOspfV3Areas) > 0 {
					return &requestSwitchUpdateNetworkSwitchRoutingOspfV3Areas
				}
				return nil
			}(),
			DeadTimerInSeconds:  int64ToIntPointer(deadTimerInSeconds),
			Enabled:             enabled,
			HelloTimerInSeconds: int64ToIntPointer(helloTimerInSeconds),
		}
		//[debug] Is Array: False
	}
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspf{
		Areas: func() *[]merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfAreas {
			if len(requestSwitchUpdateNetworkSwitchRoutingOspfAreas) > 0 {
				return &requestSwitchUpdateNetworkSwitchRoutingOspfAreas
			}
			return nil
		}(),
		DeadTimerInSeconds:       int64ToIntPointer(deadTimerInSeconds),
		Enabled:                  enabled,
		HelloTimerInSeconds:      int64ToIntPointer(helloTimerInSeconds),
		Md5AuthenticationEnabled: md5AuthenticationEnabled,
		Md5AuthenticationKey:     requestSwitchUpdateNetworkSwitchRoutingOspfMd5AuthenticationKey,
		V3:                       requestSwitchUpdateNetworkSwitchRoutingOspfV3,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchRoutingOspfItemToBodyRs(state NetworksSwitchRoutingOspfRs, response *merakigosdk.ResponseSwitchGetNetworkSwitchRoutingOspf, is_read bool) NetworksSwitchRoutingOspfRs {
	itemState := NetworksSwitchRoutingOspfRs{
		Areas: func() *[]ResponseSwitchGetNetworkSwitchRoutingOspfAreasRs {
			if response.Areas != nil {
				result := make([]ResponseSwitchGetNetworkSwitchRoutingOspfAreasRs, len(*response.Areas))
				for i, areas := range *response.Areas {
					result[i] = ResponseSwitchGetNetworkSwitchRoutingOspfAreasRs{
						Areaid: func() types.String {
							if strconv.Itoa(*areas.AreaID) != "" {
								return types.StringValue(strconv.Itoa(*areas.AreaID))
							}
							return types.String{}
						}(),
						AreaName: func() types.String {
							if areas.AreaName != "" {
								return types.StringValue(areas.AreaName)
							}
							return types.String{}
						}(),
						AreaType: func() types.String {
							if areas.AreaType != "" {
								return types.StringValue(areas.AreaType)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		DeadTimerInSeconds: func() types.Int64 {
			if response.DeadTimerInSeconds != nil {
				return types.Int64Value(int64(*response.DeadTimerInSeconds))
			}
			return types.Int64{}
		}(),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		HelloTimerInSeconds: func() types.Int64 {
			if response.HelloTimerInSeconds != nil {
				return types.Int64Value(int64(*response.HelloTimerInSeconds))
			}
			return types.Int64{}
		}(),
		Md5AuthenticationEnabled: func() types.Bool {
			if response.Md5AuthenticationEnabled != nil {
				return types.BoolValue(*response.Md5AuthenticationEnabled)
			}
			return types.Bool{}
		}(),
		Md5AuthenticationKey: func() *ResponseSwitchGetNetworkSwitchRoutingOspfMd5AuthenticationKeyRs {
			if response.Md5AuthenticationKey != nil {
				return &ResponseSwitchGetNetworkSwitchRoutingOspfMd5AuthenticationKeyRs{
					ID: func() types.Int64 {
						if response.Md5AuthenticationKey.ID != nil {
							return types.Int64Value(int64(*response.Md5AuthenticationKey.ID))
						}
						return types.Int64{}
					}(),
					Passphrase: func() types.String {
						if response.Md5AuthenticationKey.Passphrase != "" {
							return types.StringValue(response.Md5AuthenticationKey.Passphrase)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
		V3: func() *ResponseSwitchGetNetworkSwitchRoutingOspfV3Rs {
			if response.V3 != nil {
				return &ResponseSwitchGetNetworkSwitchRoutingOspfV3Rs{
					Areas: func() *[]ResponseSwitchGetNetworkSwitchRoutingOspfV3AreasRs {
						if response.V3.Areas != nil {
							result := make([]ResponseSwitchGetNetworkSwitchRoutingOspfV3AreasRs, len(*response.V3.Areas))
							for i, areas := range *response.V3.Areas {
								result[i] = ResponseSwitchGetNetworkSwitchRoutingOspfV3AreasRs{
									Areaid: func() types.String {
										if strconv.Itoa(*areas.AreaID) != "" {
											return types.StringValue(strconv.Itoa(*areas.AreaID))
										}
										return types.String{}
									}(),
									AreaName: func() types.String {
										if areas.AreaName != "" {
											return types.StringValue(areas.AreaName)
										}
										return types.String{}
									}(),
									AreaType: func() types.String {
										if areas.AreaType != "" {
											return types.StringValue(areas.AreaType)
										}
										return types.String{}
									}(),
								}
							}
							return &result
						}
						return nil
					}(),
					DeadTimerInSeconds: func() types.Int64 {
						if response.V3.DeadTimerInSeconds != nil {
							return types.Int64Value(int64(*response.V3.DeadTimerInSeconds))
						}
						return types.Int64{}
					}(),
					Enabled: func() types.Bool {
						if response.V3.Enabled != nil {
							return types.BoolValue(*response.V3.Enabled)
						}
						return types.Bool{}
					}(),
					HelloTimerInSeconds: func() types.Int64 {
						if response.V3.HelloTimerInSeconds != nil {
							return types.Int64Value(int64(*response.V3.HelloTimerInSeconds))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSwitchRoutingOspfRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSwitchRoutingOspfRs)
}
