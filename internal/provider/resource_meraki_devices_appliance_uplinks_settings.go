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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesApplianceUplinksSettingsResource{}
	_ resource.ResourceWithConfigure = &DevicesApplianceUplinksSettingsResource{}
)

func NewDevicesApplianceUplinksSettingsResource() resource.Resource {
	return &DevicesApplianceUplinksSettingsResource{}
}

type DevicesApplianceUplinksSettingsResource struct {
	client *merakigosdk.Client
}

func (r *DevicesApplianceUplinksSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesApplianceUplinksSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_appliance_uplinks_settings"
}

func (r *DevicesApplianceUplinksSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"interfaces": schema.SingleNestedAttribute{
				MarkdownDescription: `Interface settings.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"wan1": schema.SingleNestedAttribute{
						MarkdownDescription: `WAN 1 settings.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Enable or disable the interface.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"pppoe": schema.SingleNestedAttribute{
								MarkdownDescription: `Configuration options for PPPoE.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"authentication": schema.SingleNestedAttribute{
										MarkdownDescription: `Settings for PPPoE Authentication.`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"enabled": schema.BoolAttribute{
												MarkdownDescription: `Whether PPPoE authentication is enabled.`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.Bool{
													boolplanmodifier.UseStateForUnknown(),
												},
											},
											"password": schema.StringAttribute{
												MarkdownDescription: `Password for PPPoE authentication. This parameter is not returned.`,
												Sensitive:           true,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
											"username": schema.StringAttribute{
												MarkdownDescription: `Username for PPPoE authentication.`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
										},
									},
									"enabled": schema.BoolAttribute{
										MarkdownDescription: `Whether PPPoE is enabled.`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Bool{
											boolplanmodifier.UseStateForUnknown(),
										},
									},
								},
							},
							"svis": schema.SingleNestedAttribute{
								MarkdownDescription: `SVI settings by protocol.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"ipv4": schema.SingleNestedAttribute{
										MarkdownDescription: `IPv4 settings for static/dynamic mode.`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"address": schema.StringAttribute{
												MarkdownDescription: `IP address and subnet mask when in static mode.`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
											"assignment_mode": schema.StringAttribute{
												MarkdownDescription: `The assignment mode for this SVI. Applies only when PPPoE is disabled.
                                                          Allowed values: [dynamic,static]`,
												Computed: true,
												Optional: true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
												Validators: []validator.String{
													stringvalidator.OneOf(
														"dynamic",
														"static",
													),
												},
											},
											"gateway": schema.StringAttribute{
												MarkdownDescription: `Gateway IP address when in static mode.`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
											"nameservers": schema.SingleNestedAttribute{
												MarkdownDescription: `The nameserver settings for this SVI.`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.Object{
													objectplanmodifier.UseStateForUnknown(),
												},
												Attributes: map[string]schema.Attribute{

													"addresses": schema.SetAttribute{
														MarkdownDescription: `Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.`,
														Computed:            true,
														Optional:            true,
														Default:             setdefault.StaticValue(types.SetValueMust(types.StringType, make([]attr.Value, 0))),
														PlanModifiers: []planmodifier.Set{
															setplanmodifier.UseStateForUnknown(),
														},

														ElementType: types.StringType,
													},
												},
											},
										},
									},
									"ipv6": schema.SingleNestedAttribute{
										MarkdownDescription: `IPv6 settings for static/dynamic mode.`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"address": schema.StringAttribute{
												MarkdownDescription: `Static address that will override the one(s) received by SLAAC.`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
											"assignment_mode": schema.StringAttribute{
												MarkdownDescription: `The assignment mode for this SVI. Applies only when PPPoE is disabled.
                                                          Allowed values: [dynamic,static]`,
												Computed: true,
												Optional: true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
												Validators: []validator.String{
													stringvalidator.OneOf(
														"dynamic",
														"static",
													),
												},
											},
											"gateway": schema.StringAttribute{
												MarkdownDescription: `Static gateway that will override the one received by autoconf.`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
											"nameservers": schema.SingleNestedAttribute{
												MarkdownDescription: `The nameserver settings for this SVI.`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.Object{
													objectplanmodifier.UseStateForUnknown(),
												},
												Attributes: map[string]schema.Attribute{

													"addresses": schema.SetAttribute{
														MarkdownDescription: `Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.`,
														Computed:            true,
														Default:             setdefault.StaticValue(types.SetNull(types.StringType)),
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
								},
							},
							"vlan_tagging": schema.SingleNestedAttribute{
								MarkdownDescription: `VLAN tagging settings.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"enabled": schema.BoolAttribute{
										MarkdownDescription: `Whether VLAN tagging is enabled.`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Bool{
											boolplanmodifier.UseStateForUnknown(),
										},
									},
									"vlan_id": schema.Int64Attribute{
										MarkdownDescription: `The ID of the VLAN to use for VLAN tagging.`,
										// Computed:            true,
										Optional: true,
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.UseStateForUnknown(),
										},
									},
								},
							},
						},
					},
					"wan2": schema.SingleNestedAttribute{
						MarkdownDescription: `WAN 2 settings.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Enable or disable the interface.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"pppoe": schema.SingleNestedAttribute{
								MarkdownDescription: `Configuration options for PPPoE.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"authentication": schema.SingleNestedAttribute{
										MarkdownDescription: `Settings for PPPoE Authentication.`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"enabled": schema.BoolAttribute{
												MarkdownDescription: `Whether PPPoE authentication is enabled.`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.Bool{
													boolplanmodifier.UseStateForUnknown(),
												},
											},
											"password": schema.StringAttribute{
												MarkdownDescription: `Password for PPPoE authentication. This parameter is not returned.`,
												Sensitive:           true,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
											"username": schema.StringAttribute{
												MarkdownDescription: `Username for PPPoE authentication.`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
										},
									},
									"enabled": schema.BoolAttribute{
										MarkdownDescription: `Whether PPPoE is enabled.`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Bool{
											boolplanmodifier.UseStateForUnknown(),
										},
									},
								},
							},
							"svis": schema.SingleNestedAttribute{
								MarkdownDescription: `SVI settings by protocol.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"ipv4": schema.SingleNestedAttribute{
										MarkdownDescription: `IPv4 settings for static/dynamic mode.`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"address": schema.StringAttribute{
												MarkdownDescription: `IP address and subnet mask when in static mode.`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
											"assignment_mode": schema.StringAttribute{
												MarkdownDescription: `The assignment mode for this SVI. Applies only when PPPoE is disabled.
                                                          Allowed values: [dynamic,static]`,
												Computed: true,
												Optional: true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
												Validators: []validator.String{
													stringvalidator.OneOf(
														"dynamic",
														"static",
													),
												},
											},
											"gateway": schema.StringAttribute{
												MarkdownDescription: `Gateway IP address when in static mode.`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
											"nameservers": schema.SingleNestedAttribute{
												MarkdownDescription: `The nameserver settings for this SVI.`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.Object{
													objectplanmodifier.UseStateForUnknown(),
												},
												Attributes: map[string]schema.Attribute{

													"addresses": schema.SetAttribute{
														MarkdownDescription: `Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.`,
														Computed:            true,
														Default:             setdefault.StaticValue(types.SetNull(types.StringType)),
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
									"ipv6": schema.SingleNestedAttribute{
										MarkdownDescription: `IPv6 settings for static/dynamic mode.`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"address": schema.StringAttribute{
												MarkdownDescription: `Static address that will override the one(s) received by SLAAC.`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
											"assignment_mode": schema.StringAttribute{
												MarkdownDescription: `The assignment mode for this SVI. Applies only when PPPoE is disabled.
                                                          Allowed values: [dynamic,static]`,
												Computed: true,
												Optional: true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
												Validators: []validator.String{
													stringvalidator.OneOf(
														"dynamic",
														"static",
													),
												},
											},
											"gateway": schema.StringAttribute{
												MarkdownDescription: `Static gateway that will override the one received by autoconf.`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
											"nameservers": schema.SingleNestedAttribute{
												MarkdownDescription: `The nameserver settings for this SVI.`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.Object{
													objectplanmodifier.UseStateForUnknown(),
												},
												Attributes: map[string]schema.Attribute{

													"addresses": schema.SetAttribute{
														MarkdownDescription: `Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.`,
														Computed:            true,
														Default:             setdefault.StaticValue(types.SetNull(types.StringType)),
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
								},
							},
							"vlan_tagging": schema.SingleNestedAttribute{
								MarkdownDescription: `VLAN tagging settings.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"enabled": schema.BoolAttribute{
										MarkdownDescription: `Whether VLAN tagging is enabled.`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Bool{
											boolplanmodifier.UseStateForUnknown(),
										},
									},
									"vlan_id": schema.Int64Attribute{
										MarkdownDescription: `The ID of the VLAN to use for VLAN tagging.`,
										// Computed:            true,
										Optional: true,
										PlanModifiers: []planmodifier.Int64{
											int64planmodifier.UseStateForUnknown(),
										},
									},
								},
							},
						},
					},
				},
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
		},
	}
}

func (r *DevicesApplianceUplinksSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesApplianceUplinksSettingsRs

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
	vvSerial := data.Serial.ValueString()
	//Has Item and not has items

	if vvSerial != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Appliance.GetDeviceApplianceUplinksSettings(vvSerial)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource DevicesApplianceUplinksSettings  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource DevicesApplianceUplinksSettings only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateDeviceApplianceUplinksSettings(vvSerial, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceApplianceUplinksSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceApplianceUplinksSettings",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Appliance.GetDeviceApplianceUplinksSettings(vvSerial)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceApplianceUplinksSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceApplianceUplinksSettings",
			err.Error(),
		)
		return
	}

	data = ResponseApplianceGetDeviceApplianceUplinksSettingsItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *DevicesApplianceUplinksSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesApplianceUplinksSettingsRs

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

	vvSerial := data.Serial.ValueString()
	responseGet, restyRespGet, err := r.client.Appliance.GetDeviceApplianceUplinksSettings(vvSerial)
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
				"Failure when executing GetDeviceApplianceUplinksSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceApplianceUplinksSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetDeviceApplianceUplinksSettingsItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *DevicesApplianceUplinksSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), req.ID)...)
}

func (r *DevicesApplianceUplinksSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesApplianceUplinksSettingsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvSerial := data.Serial.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateDeviceApplianceUplinksSettings(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceApplianceUplinksSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceApplianceUplinksSettings",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesApplianceUplinksSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting DevicesApplianceUplinksSettings", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesApplianceUplinksSettingsRs struct {
	Serial     types.String                                                    `tfsdk:"serial"`
	Interfaces *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesRs `tfsdk:"interfaces"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesRs struct {
	Wan1 *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1Rs `tfsdk:"wan1"`
	Wan2 *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2Rs `tfsdk:"wan2"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1Rs struct {
	Enabled     types.Bool                                                                     `tfsdk:"enabled"`
	Pppoe       *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1PppoeRs       `tfsdk:"pppoe"`
	Svis        *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisRs        `tfsdk:"svis"`
	VLANTagging *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1VlanTaggingRs `tfsdk:"vlan_tagging"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1PppoeRs struct {
	Authentication *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthenticationRs `tfsdk:"authentication"`
	Enabled        types.Bool                                                                             `tfsdk:"enabled"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthenticationRs struct {
	Enabled  types.Bool   `tfsdk:"enabled"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisRs struct {
	IPv4 *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv4Rs `tfsdk:"ipv4"`
	IPv6 *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv6Rs `tfsdk:"ipv6"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv4Rs struct {
	Address        types.String                                                                           `tfsdk:"address"`
	AssignmentMode types.String                                                                           `tfsdk:"assignment_mode"`
	Gateway        types.String                                                                           `tfsdk:"gateway"`
	Nameservers    *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv4NameserversRs `tfsdk:"nameservers"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv4NameserversRs struct {
	Addresses types.Set `tfsdk:"addresses"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv6Rs struct {
	Address        types.String                                                                           `tfsdk:"address"`
	AssignmentMode types.String                                                                           `tfsdk:"assignment_mode"`
	Gateway        types.String                                                                           `tfsdk:"gateway"`
	Nameservers    *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv6NameserversRs `tfsdk:"nameservers"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv6NameserversRs struct {
	Addresses types.Set `tfsdk:"addresses"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1VlanTaggingRs struct {
	Enabled types.Bool  `tfsdk:"enabled"`
	VLANID  types.Int64 `tfsdk:"vlan_id"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2Rs struct {
	Enabled     types.Bool                                                                     `tfsdk:"enabled"`
	Pppoe       *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2PppoeRs       `tfsdk:"pppoe"`
	Svis        *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisRs        `tfsdk:"svis"`
	VLANTagging *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2VlanTaggingRs `tfsdk:"vlan_tagging"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2PppoeRs struct {
	Authentication *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthenticationRs `tfsdk:"authentication"`
	Enabled        types.Bool                                                                             `tfsdk:"enabled"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthenticationRs struct {
	Enabled  types.Bool   `tfsdk:"enabled"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisRs struct {
	IPv4 *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv4Rs `tfsdk:"ipv4"`
	IPv6 *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv6Rs `tfsdk:"ipv6"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv4Rs struct {
	Address        types.String                                                                           `tfsdk:"address"`
	AssignmentMode types.String                                                                           `tfsdk:"assignment_mode"`
	Gateway        types.String                                                                           `tfsdk:"gateway"`
	Nameservers    *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv4NameserversRs `tfsdk:"nameservers"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv4NameserversRs struct {
	Addresses types.Set `tfsdk:"addresses"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv6Rs struct {
	Address        types.String                                                                           `tfsdk:"address"`
	AssignmentMode types.String                                                                           `tfsdk:"assignment_mode"`
	Gateway        types.String                                                                           `tfsdk:"gateway"`
	Nameservers    *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv6NameserversRs `tfsdk:"nameservers"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv6NameserversRs struct {
	Addresses types.Set `tfsdk:"addresses"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2VlanTaggingRs struct {
	Enabled types.Bool  `tfsdk:"enabled"`
	VLANID  types.Int64 `tfsdk:"vlan_id"`
}

// FromBody
func (r *DevicesApplianceUplinksSettingsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettings {
	var requestApplianceUpdateDeviceApplianceUplinksSettingsInterfaces *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfaces

	if r.Interfaces != nil {
		var requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1 *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1

		if r.Interfaces.Wan1 != nil {
			enabled := func() *bool {
				if !r.Interfaces.Wan1.Enabled.IsUnknown() && !r.Interfaces.Wan1.Enabled.IsNull() {
					return r.Interfaces.Wan1.Enabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1Pppoe *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1Pppoe

			if r.Interfaces.Wan1.Pppoe != nil {
				var requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthentication *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthentication

				if r.Interfaces.Wan1.Pppoe.Authentication != nil {
					enabled := func() *bool {
						if !r.Interfaces.Wan1.Pppoe.Authentication.Enabled.IsUnknown() && !r.Interfaces.Wan1.Pppoe.Authentication.Enabled.IsNull() {
							return r.Interfaces.Wan1.Pppoe.Authentication.Enabled.ValueBoolPointer()
						}
						return nil
					}()
					password := r.Interfaces.Wan1.Pppoe.Authentication.Password.ValueString()
					username := r.Interfaces.Wan1.Pppoe.Authentication.Username.ValueString()
					requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthentication = &merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthentication{
						Enabled:  enabled,
						Password: password,
						Username: username,
					}
					//[debug] Is Array: False
				}
				enabled := func() *bool {
					if !r.Interfaces.Wan1.Pppoe.Enabled.IsUnknown() && !r.Interfaces.Wan1.Pppoe.Enabled.IsNull() {
						return r.Interfaces.Wan1.Pppoe.Enabled.ValueBoolPointer()
					}
					return nil
				}()
				requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1Pppoe = &merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1Pppoe{
					Authentication: requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthentication,
					Enabled:        enabled,
				}
				//[debug] Is Array: False
			}
			var requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1Svis *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1Svis

			if r.Interfaces.Wan1.Svis != nil {
				var requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4 *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4

				if r.Interfaces.Wan1.Svis.IPv4 != nil {
					address := r.Interfaces.Wan1.Svis.IPv4.Address.ValueString()
					assignmentMode := r.Interfaces.Wan1.Svis.IPv4.AssignmentMode.ValueString()
					gateway := r.Interfaces.Wan1.Svis.IPv4.Gateway.ValueString()
					var requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4Nameservers *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4Nameservers

					if r.Interfaces.Wan1.Svis.IPv4.Nameservers != nil {

						var addresses []string = nil
						r.Interfaces.Wan1.Svis.IPv4.Nameservers.Addresses.ElementsAs(ctx, &addresses, false)
						requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4Nameservers = &merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4Nameservers{
							Addresses: addresses,
						}
						//[debug] Is Array: False
					}
					requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4 = &merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4{
						Address:        address,
						AssignmentMode: assignmentMode,
						Gateway:        gateway,
						Nameservers:    requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4Nameservers,
					}
					//[debug] Is Array: False
				}
				var requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6 *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6

				if r.Interfaces.Wan1.Svis.IPv6 != nil {
					address := r.Interfaces.Wan1.Svis.IPv6.Address.ValueString()
					assignmentMode := r.Interfaces.Wan1.Svis.IPv6.AssignmentMode.ValueString()
					gateway := r.Interfaces.Wan1.Svis.IPv6.Gateway.ValueString()
					var requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6Nameservers *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6Nameservers

					if r.Interfaces.Wan1.Svis.IPv6.Nameservers != nil {

						var addresses []string = nil
						r.Interfaces.Wan1.Svis.IPv6.Nameservers.Addresses.ElementsAs(ctx, &addresses, false)
						requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6Nameservers = &merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6Nameservers{
							Addresses: addresses,
						}
						//[debug] Is Array: False
					}
					requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6 = &merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6{
						Address:        address,
						AssignmentMode: assignmentMode,
						Gateway:        gateway,
						Nameservers:    requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6Nameservers,
					}
					//[debug] Is Array: False
				}
				requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1Svis = &merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1Svis{
					IPv4: requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv4,
					IPv6: requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1SvisIPv6,
				}
				//[debug] Is Array: False
			}
			var requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1VLANTagging *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1VLANTagging

			if r.Interfaces.Wan1.VLANTagging != nil {
				enabled := func() *bool {
					if !r.Interfaces.Wan1.VLANTagging.Enabled.IsUnknown() && !r.Interfaces.Wan1.VLANTagging.Enabled.IsNull() {
						return r.Interfaces.Wan1.VLANTagging.Enabled.ValueBoolPointer()
					}
					return nil
				}()
				vlanID := func() *int64 {
					if !r.Interfaces.Wan1.VLANTagging.VLANID.IsUnknown() && !r.Interfaces.Wan1.VLANTagging.VLANID.IsNull() {
						return r.Interfaces.Wan1.VLANTagging.VLANID.ValueInt64Pointer()
					}
					return nil
				}()
				requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1VLANTagging = &merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1VLANTagging{
					Enabled: enabled,
					VLANID:  int64ToIntPointer(vlanID),
				}
				//[debug] Is Array: False
			}
			requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1 = &merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1{
				Enabled:     enabled,
				Pppoe:       requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1Pppoe,
				Svis:        requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1Svis,
				VLANTagging: requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1VLANTagging,
			}
			//[debug] Is Array: False
		}
		var requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2 *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2

		if r.Interfaces.Wan2 != nil {
			enabled := func() *bool {
				if !r.Interfaces.Wan2.Enabled.IsUnknown() && !r.Interfaces.Wan2.Enabled.IsNull() {
					return r.Interfaces.Wan2.Enabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2Pppoe *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2Pppoe

			if r.Interfaces.Wan2.Pppoe != nil {
				var requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthentication *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthentication

				if r.Interfaces.Wan2.Pppoe.Authentication != nil {
					enabled := func() *bool {
						if !r.Interfaces.Wan2.Pppoe.Authentication.Enabled.IsUnknown() && !r.Interfaces.Wan2.Pppoe.Authentication.Enabled.IsNull() {
							return r.Interfaces.Wan2.Pppoe.Authentication.Enabled.ValueBoolPointer()
						}
						return nil
					}()
					password := r.Interfaces.Wan2.Pppoe.Authentication.Password.ValueString()
					username := r.Interfaces.Wan2.Pppoe.Authentication.Username.ValueString()
					requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthentication = &merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthentication{
						Enabled:  enabled,
						Password: password,
						Username: username,
					}
					//[debug] Is Array: False
				}
				enabled := func() *bool {
					if !r.Interfaces.Wan2.Pppoe.Enabled.IsUnknown() && !r.Interfaces.Wan2.Pppoe.Enabled.IsNull() {
						return r.Interfaces.Wan2.Pppoe.Enabled.ValueBoolPointer()
					}
					return nil
				}()
				requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2Pppoe = &merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2Pppoe{
					Authentication: requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthentication,
					Enabled:        enabled,
				}
				//[debug] Is Array: False
			}
			var requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2Svis *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2Svis

			if r.Interfaces.Wan2.Svis != nil {
				var requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4 *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4

				if r.Interfaces.Wan2.Svis.IPv4 != nil {
					address := r.Interfaces.Wan2.Svis.IPv4.Address.ValueString()
					assignmentMode := r.Interfaces.Wan2.Svis.IPv4.AssignmentMode.ValueString()
					gateway := r.Interfaces.Wan2.Svis.IPv4.Gateway.ValueString()
					var requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4Nameservers *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4Nameservers

					if r.Interfaces.Wan2.Svis.IPv4.Nameservers != nil {

						var addresses []string = nil
						r.Interfaces.Wan2.Svis.IPv4.Nameservers.Addresses.ElementsAs(ctx, &addresses, false)
						requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4Nameservers = &merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4Nameservers{
							Addresses: addresses,
						}
						//[debug] Is Array: False
					}
					requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4 = &merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4{
						Address:        address,
						AssignmentMode: assignmentMode,
						Gateway:        gateway,
						Nameservers:    requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4Nameservers,
					}
					//[debug] Is Array: False
				}
				var requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6 *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6

				if r.Interfaces.Wan2.Svis.IPv6 != nil {
					address := r.Interfaces.Wan2.Svis.IPv6.Address.ValueString()
					assignmentMode := r.Interfaces.Wan2.Svis.IPv6.AssignmentMode.ValueString()
					gateway := r.Interfaces.Wan2.Svis.IPv6.Gateway.ValueString()
					var requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6Nameservers *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6Nameservers

					if r.Interfaces.Wan2.Svis.IPv6.Nameservers != nil {

						var addresses []string = nil
						r.Interfaces.Wan2.Svis.IPv6.Nameservers.Addresses.ElementsAs(ctx, &addresses, false)
						requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6Nameservers = &merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6Nameservers{
							Addresses: addresses,
						}
						//[debug] Is Array: False
					}
					requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6 = &merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6{
						Address:        address,
						AssignmentMode: assignmentMode,
						Gateway:        gateway,
						Nameservers:    requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6Nameservers,
					}
					//[debug] Is Array: False
				}
				requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2Svis = &merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2Svis{
					IPv4: requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv4,
					IPv6: requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2SvisIPv6,
				}
				//[debug] Is Array: False
			}
			var requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2VLANTagging *merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2VLANTagging

			if r.Interfaces.Wan2.VLANTagging != nil {
				enabled := func() *bool {
					if !r.Interfaces.Wan2.VLANTagging.Enabled.IsUnknown() && !r.Interfaces.Wan2.VLANTagging.Enabled.IsNull() {
						return r.Interfaces.Wan2.VLANTagging.Enabled.ValueBoolPointer()
					}
					return nil
				}()
				vlanID := func() *int64 {
					if !r.Interfaces.Wan2.VLANTagging.VLANID.IsUnknown() && !r.Interfaces.Wan2.VLANTagging.VLANID.IsNull() {
						return r.Interfaces.Wan2.VLANTagging.VLANID.ValueInt64Pointer()
					}
					return nil
				}()
				requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2VLANTagging = &merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2VLANTagging{
					Enabled: enabled,
					VLANID:  int64ToIntPointer(vlanID),
				}
				//[debug] Is Array: False
			}
			requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2 = &merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2{
				Enabled:     enabled,
				Pppoe:       requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2Pppoe,
				Svis:        requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2Svis,
				VLANTagging: requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2VLANTagging,
			}
			//[debug] Is Array: False
		}
		requestApplianceUpdateDeviceApplianceUplinksSettingsInterfaces = &merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettingsInterfaces{
			Wan1: requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan1,
			Wan2: requestApplianceUpdateDeviceApplianceUplinksSettingsInterfacesWan2,
		}
		//[debug] Is Array: False
	}
	out := merakigosdk.RequestApplianceUpdateDeviceApplianceUplinksSettings{
		Interfaces: requestApplianceUpdateDeviceApplianceUplinksSettingsInterfaces,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetDeviceApplianceUplinksSettingsItemToBodyRs(state DevicesApplianceUplinksSettingsRs, response *merakigosdk.ResponseApplianceGetDeviceApplianceUplinksSettings, is_read bool) DevicesApplianceUplinksSettingsRs {
	itemState := DevicesApplianceUplinksSettingsRs{
		Interfaces: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesRs {
			if response.Interfaces != nil {
				return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesRs{
					Wan1: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1Rs {
						if response.Interfaces.Wan1 != nil {
							return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1Rs{
								Enabled: func() types.Bool {
									if response.Interfaces.Wan1.Enabled != nil {
										return types.BoolValue(*response.Interfaces.Wan1.Enabled)
									}
									return types.Bool{}
								}(),
								Pppoe: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1PppoeRs {
									if response.Interfaces.Wan1.Pppoe != nil {
										return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1PppoeRs{
											Authentication: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthenticationRs {
												if response.Interfaces.Wan1.Pppoe.Authentication != nil {
													return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthenticationRs{
														Enabled: func() types.Bool {
															if response.Interfaces.Wan1.Pppoe.Authentication.Enabled != nil {
																return types.BoolValue(*response.Interfaces.Wan1.Pppoe.Authentication.Enabled)
															}
															return types.Bool{}
														}(),
														Username: types.StringValue(response.Interfaces.Wan1.Pppoe.Authentication.Username),
													}
												}
												return nil
											}(),
											Enabled: func() types.Bool {
												if response.Interfaces.Wan1.Pppoe.Enabled != nil {
													return types.BoolValue(*response.Interfaces.Wan1.Pppoe.Enabled)
												}
												return types.Bool{}
											}(),
										}
									}
									return nil
								}(),
								Svis: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisRs {
									if response.Interfaces.Wan1.Svis != nil {
										return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisRs{
											IPv4: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv4Rs {
												if response.Interfaces.Wan1.Svis.IPv4 != nil {
													return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv4Rs{
														Address:        types.StringValue(response.Interfaces.Wan1.Svis.IPv4.Address),
														AssignmentMode: types.StringValue(response.Interfaces.Wan1.Svis.IPv4.AssignmentMode),
														Gateway:        types.StringValue(response.Interfaces.Wan1.Svis.IPv4.Gateway),
														Nameservers: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv4NameserversRs {
															if response.Interfaces.Wan1.Svis.IPv4.Nameservers != nil {
																return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv4NameserversRs{
																	Addresses: StringSliceToSet(response.Interfaces.Wan1.Svis.IPv4.Nameservers.Addresses),
																}
															}
															return nil
														}(),
													}
												}
												return nil
											}(),
											IPv6: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv6Rs {
												if response.Interfaces.Wan1.Svis.IPv6 != nil {
													return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv6Rs{
														Address:        types.StringValue(response.Interfaces.Wan1.Svis.IPv6.Address),
														AssignmentMode: types.StringValue(response.Interfaces.Wan1.Svis.IPv6.AssignmentMode),
														Gateway:        types.StringValue(response.Interfaces.Wan1.Svis.IPv6.Gateway),
														Nameservers: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv6NameserversRs {
															if response.Interfaces.Wan1.Svis.IPv6.Nameservers != nil {
																return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv6NameserversRs{
																	Addresses: StringSliceToSet(response.Interfaces.Wan1.Svis.IPv6.Nameservers.Addresses),
																}
															}
															return nil
														}(),
													}
												}
												return nil
											}(),
										}
									}
									return nil
								}(),
								VLANTagging: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1VlanTaggingRs {
									if response.Interfaces.Wan1.VLANTagging != nil {
										return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1VlanTaggingRs{
											Enabled: func() types.Bool {
												if response.Interfaces.Wan1.VLANTagging.Enabled != nil {
													return types.BoolValue(*response.Interfaces.Wan1.VLANTagging.Enabled)
												}
												return types.Bool{}
											}(),
											VLANID: func() types.Int64 {
												if response.Interfaces.Wan1.VLANTagging.VLANID != nil {
													return types.Int64Value(int64(*response.Interfaces.Wan1.VLANTagging.VLANID))
												}
												return types.Int64{}
											}(),
										}
									}
									return nil
								}(),
							}
						}
						return nil
					}(),
					Wan2: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2Rs {
						if response.Interfaces.Wan2 != nil {
							return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2Rs{
								Enabled: func() types.Bool {
									if response.Interfaces.Wan2.Enabled != nil {
										return types.BoolValue(*response.Interfaces.Wan2.Enabled)
									}
									return types.Bool{}
								}(),
								Pppoe: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2PppoeRs {
									if response.Interfaces.Wan2.Pppoe != nil {
										return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2PppoeRs{
											Authentication: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthenticationRs {
												if response.Interfaces.Wan2.Pppoe.Authentication != nil {
													return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthenticationRs{
														Enabled: func() types.Bool {
															if response.Interfaces.Wan2.Pppoe.Authentication.Enabled != nil {
																return types.BoolValue(*response.Interfaces.Wan2.Pppoe.Authentication.Enabled)
															}
															return types.Bool{}
														}(),
														Username: types.StringValue(response.Interfaces.Wan2.Pppoe.Authentication.Username),
													}
												}
												return nil
											}(),
											Enabled: func() types.Bool {
												if response.Interfaces.Wan2.Pppoe.Enabled != nil {
													return types.BoolValue(*response.Interfaces.Wan2.Pppoe.Enabled)
												}
												return types.Bool{}
											}(),
										}
									}
									return nil
								}(),
								Svis: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisRs {
									if response.Interfaces.Wan2.Svis != nil {
										return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisRs{
											IPv4: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv4Rs {
												if response.Interfaces.Wan2.Svis.IPv4 != nil {
													return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv4Rs{
														Address:        types.StringValue(response.Interfaces.Wan2.Svis.IPv4.Address),
														AssignmentMode: types.StringValue(response.Interfaces.Wan2.Svis.IPv4.AssignmentMode),
														Gateway:        types.StringValue(response.Interfaces.Wan2.Svis.IPv4.Gateway),
														Nameservers: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv4NameserversRs {
															if response.Interfaces.Wan2.Svis.IPv4.Nameservers != nil {
																return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv4NameserversRs{
																	Addresses: StringSliceToSet(response.Interfaces.Wan2.Svis.IPv4.Nameservers.Addresses),
																}
															}
															return nil
														}(),
													}
												}
												return nil
											}(),
											IPv6: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv6Rs {
												if response.Interfaces.Wan2.Svis.IPv6 != nil {
													return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv6Rs{
														Address:        types.StringValue(response.Interfaces.Wan2.Svis.IPv6.Address),
														AssignmentMode: types.StringValue(response.Interfaces.Wan2.Svis.IPv6.AssignmentMode),
														Gateway:        types.StringValue(response.Interfaces.Wan2.Svis.IPv6.Gateway),
														Nameservers: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv6NameserversRs {
															if response.Interfaces.Wan2.Svis.IPv6.Nameservers != nil {
																return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv6NameserversRs{
																	Addresses: StringSliceToSet(response.Interfaces.Wan2.Svis.IPv6.Nameservers.Addresses),
																}
															}
															return nil
														}(),
													}
												}
												return nil
											}(),
										}
									}
									return nil
								}(),
								VLANTagging: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2VlanTaggingRs {
									if response.Interfaces.Wan2.VLANTagging != nil {
										return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2VlanTaggingRs{
											Enabled: func() types.Bool {
												if response.Interfaces.Wan2.VLANTagging.Enabled != nil {
													return types.BoolValue(*response.Interfaces.Wan2.VLANTagging.Enabled)
												}
												return types.Bool{}
											}(),
											VLANID: func() types.Int64 {
												if response.Interfaces.Wan2.VLANTagging.VLANID != nil {
													return types.Int64Value(int64(*response.Interfaces.Wan2.VLANTagging.VLANID))
												}
												return types.Int64{}
											}(),
										}
									}
									return nil
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesApplianceUplinksSettingsRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesApplianceUplinksSettingsRs)
}
