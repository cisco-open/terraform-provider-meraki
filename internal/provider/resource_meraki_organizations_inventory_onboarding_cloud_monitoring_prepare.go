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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

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
	_ resource.Resource              = &OrganizationsInventoryOnboardingCloudMonitoringPrepareResource{}
	_ resource.ResourceWithConfigure = &OrganizationsInventoryOnboardingCloudMonitoringPrepareResource{}
)

func NewOrganizationsInventoryOnboardingCloudMonitoringPrepareResource() resource.Resource {
	return &OrganizationsInventoryOnboardingCloudMonitoringPrepareResource{}
}

type OrganizationsInventoryOnboardingCloudMonitoringPrepareResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsInventoryOnboardingCloudMonitoringPrepareResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsInventoryOnboardingCloudMonitoringPrepareResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_inventory_onboarding_cloud_monitoring_prepare"
}

// resourceAction
func (r *OrganizationsInventoryOnboardingCloudMonitoringPrepareResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"devices": schema.SetNestedAttribute{
						MarkdownDescription: `A set of devices to import (or update)`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"sudi": schema.StringAttribute{
									MarkdownDescription: `Device SUDI certificate`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"tunnel": schema.SingleNestedAttribute{
									MarkdownDescription: `TLS Related Parameters`,
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"certificate_name": schema.StringAttribute{
											MarkdownDescription: `Name of the configured TLS certificate`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.RequiresReplace(),
											},
										},
										"local_interface": schema.Int64Attribute{
											MarkdownDescription: `Number of the vlan expected to be used to connect to the cloud`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.RequiresReplace(),
											},
										},
										"loopback_number": schema.Int64Attribute{
											MarkdownDescription: `Number of the configured Loopback Interface used for TLS overlay`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.RequiresReplace(),
											},
										},
										"name": schema.StringAttribute{
											MarkdownDescription: `Name of the configured TLS tunnel`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.RequiresReplace(),
											},
										},
									},
								},
								"user": schema.SingleNestedAttribute{
									MarkdownDescription: `User parameters`,
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"username": schema.StringAttribute{
											MarkdownDescription: `The name of the device user for Meraki monitoring`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.RequiresReplace(),
											},
										},
									},
								},
								"vty": schema.SingleNestedAttribute{
									MarkdownDescription: `VTY Related Parameters`,
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"access_list": schema.SingleNestedAttribute{
											MarkdownDescription: `AccessList details`,
											Optional:            true,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"vty_in": schema.SingleNestedAttribute{
													MarkdownDescription: `VTY in ACL`,
													Optional:            true,
													Computed:            true,
													Attributes: map[string]schema.Attribute{

														"name": schema.StringAttribute{
															MarkdownDescription: `Name`,
															Optional:            true,
															Computed:            true,
															PlanModifiers: []planmodifier.String{
																stringplanmodifier.RequiresReplace(),
															},
														},
													},
												},
												"vty_out": schema.SingleNestedAttribute{
													MarkdownDescription: `VTY out ACL`,
													Optional:            true,
													Computed:            true,
													Attributes: map[string]schema.Attribute{

														"name": schema.StringAttribute{
															MarkdownDescription: `Name`,
															Optional:            true,
															Computed:            true,
															PlanModifiers: []planmodifier.String{
																stringplanmodifier.RequiresReplace(),
															},
														},
													},
												},
											},
										},
										"authentication": schema.SingleNestedAttribute{
											MarkdownDescription: `VTY AAA authentication`,
											Optional:            true,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"group": schema.SingleNestedAttribute{
													MarkdownDescription: `Group Details`,
													Optional:            true,
													Computed:            true,
													Attributes: map[string]schema.Attribute{

														"name": schema.StringAttribute{
															MarkdownDescription: `Group Name`,
															Optional:            true,
															Computed:            true,
															PlanModifiers: []planmodifier.String{
																stringplanmodifier.RequiresReplace(),
															},
														},
													},
												},
											},
										},
										"authorization": schema.SingleNestedAttribute{
											MarkdownDescription: `VTY AAA authorization`,
											Optional:            true,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"group": schema.SingleNestedAttribute{
													MarkdownDescription: `Group Details`,
													Optional:            true,
													Computed:            true,
													Attributes: map[string]schema.Attribute{

														"name": schema.StringAttribute{
															MarkdownDescription: `Group Name`,
															Optional:            true,
															Computed:            true,
															PlanModifiers: []planmodifier.String{
																stringplanmodifier.RequiresReplace(),
															},
														},
													},
												},
											},
										},
										"end_line_number": schema.Int64Attribute{
											MarkdownDescription: `Ending line VTY number`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.RequiresReplace(),
											},
										},
										"rotary_number": schema.Int64Attribute{
											MarkdownDescription: `SSH rotary number`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.RequiresReplace(),
											},
										},
										"start_line_number": schema.Int64Attribute{
											MarkdownDescription: `Starting line VTY number`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.RequiresReplace(),
											},
										},
									},
								},
							},
						},
					},
					"items": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"items": schema.ListNestedAttribute{
									MarkdownDescription: `Array of ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepare`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"config_params": schema.SingleNestedAttribute{
												MarkdownDescription: `Params used in order to connect to the device`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"cloud_static_ip": schema.StringAttribute{
														MarkdownDescription: `Static IP Address used to connect to the device`,
														Computed:            true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.RequiresReplace(),
														},
													},
													"tunnel": schema.SingleNestedAttribute{
														MarkdownDescription: `Configuration options used to connect to the device`,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"host": schema.StringAttribute{
																MarkdownDescription: `SSH tunnel URL used to connect to the device`,
																Computed:            true,
																PlanModifiers: []planmodifier.String{
																	stringplanmodifier.RequiresReplace(),
																},
															},
															"mode": schema.StringAttribute{
																Computed: true,
																PlanModifiers: []planmodifier.String{
																	stringplanmodifier.RequiresReplace(),
																},
															},
															"name": schema.StringAttribute{
																MarkdownDescription: `The name of the tunnel we are attempting to connect to`,
																Computed:            true,
																PlanModifiers: []planmodifier.String{
																	stringplanmodifier.RequiresReplace(),
																},
															},
															"port": schema.StringAttribute{
																MarkdownDescription: `The port used for the ssh tunnel.`,
																Computed:            true,
																PlanModifiers: []planmodifier.String{
																	stringplanmodifier.RequiresReplace(),
																},
															},
															"root_certificate": schema.SingleNestedAttribute{
																MarkdownDescription: `Root certificate information`,
																Computed:            true,
																Attributes: map[string]schema.Attribute{

																	"content": schema.StringAttribute{
																		MarkdownDescription: `Public certificate value`,
																		Computed:            true,
																		PlanModifiers: []planmodifier.String{
																			stringplanmodifier.RequiresReplace(),
																		},
																	},
																	"name": schema.StringAttribute{
																		MarkdownDescription: `The name of the server protected by the certificate`,
																		Computed:            true,
																		PlanModifiers: []planmodifier.String{
																			stringplanmodifier.RequiresReplace(),
																		},
																	},
																},
															},
														},
													},
													"user": schema.SingleNestedAttribute{
														MarkdownDescription: `User credentials used to connect to the device`,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"public_key": schema.StringAttribute{
																MarkdownDescription: `The public key for the registered user`,
																Computed:            true,
																PlanModifiers: []planmodifier.String{
																	stringplanmodifier.RequiresReplace(),
																},
															},
															"secret": schema.SingleNestedAttribute{
																MarkdownDescription: `Stores the user secret hash`,
																Computed:            true,
																Attributes: map[string]schema.Attribute{

																	"hash": schema.StringAttribute{
																		MarkdownDescription: `The hashed secret`,
																		Computed:            true,
																		PlanModifiers: []planmodifier.String{
																			stringplanmodifier.RequiresReplace(),
																		},
																	},
																},
															},
															"username": schema.StringAttribute{
																MarkdownDescription: `The username added to Catalyst device`,
																Computed:            true,
																PlanModifiers: []planmodifier.String{
																	stringplanmodifier.RequiresReplace(),
																},
															},
														},
													},
												},
											},
											"device_id": schema.StringAttribute{
												MarkdownDescription: `Import ID from the Import operation`,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
											"message": schema.StringAttribute{
												MarkdownDescription: `Message related to whether or not the device was found and can be imported.`,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
											"status": schema.StringAttribute{
												MarkdownDescription: `The import status of the device`,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
											"udi": schema.StringAttribute{
												MarkdownDescription: `Device UDI certificate`,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
										},
									},
								},
							},
						},
					},
					"options": schema.SingleNestedAttribute{
						MarkdownDescription: `Additional options for the import`,
						Optional:            true,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"skip_commit": schema.BoolAttribute{
								MarkdownDescription: `Flag to skip adding the device to RDM`,
								Optional:            true,
								Computed:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.RequiresReplace(),
								},
							},
						},
					},
				},
			},
		},
	}
}
func (r *OrganizationsInventoryOnboardingCloudMonitoringPrepareResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsInventoryOnboardingCloudMonitoringPrepare

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
	vvOrganizationID := data.OrganizationID.ValueString()
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Organizations.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(vvOrganizationID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationInventoryOnboardingCloudMonitoringPrepare",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationInventoryOnboardingCloudMonitoringPrepare",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareItemsToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsInventoryOnboardingCloudMonitoringPrepareResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsInventoryOnboardingCloudMonitoringPrepareResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsInventoryOnboardingCloudMonitoringPrepareResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsInventoryOnboardingCloudMonitoringPrepare struct {
	OrganizationID types.String                                                                            `tfsdk:"organization_id"`
	Items          *[]ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepare `tfsdk:"items"`
	Parameters     *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRs      `tfsdk:"parameters"`
}

type ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepare struct {
	ConfigParams *ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParams `tfsdk:"config_params"`
	DeviceID     types.String                                                                                      `tfsdk:"device_id"`
	Message      types.String                                                                                      `tfsdk:"message"`
	Status       types.String                                                                                      `tfsdk:"status"`
	Udi          types.String                                                                                      `tfsdk:"udi"`
}

type ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParams struct {
	CloudStaticIP types.String                                                                                            `tfsdk:"cloud_static_ip"`
	Tunnel        *ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsTunnel `tfsdk:"tunnel"`
	User          *ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsUser   `tfsdk:"user"`
}

type ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsTunnel struct {
	Host            types.String                                                                                                           `tfsdk:"host"`
	Mode            types.String                                                                                                           `tfsdk:"mode"`
	Name            types.String                                                                                                           `tfsdk:"name"`
	Port            types.String                                                                                                           `tfsdk:"port"`
	RootCertificate *ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsTunnelRootCertificate `tfsdk:"root_certificate"`
}

type ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsTunnelRootCertificate struct {
	Content types.String `tfsdk:"content"`
	Name    types.String `tfsdk:"name"`
}

type ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsUser struct {
	PublicKey types.String                                                                                                `tfsdk:"public_key"`
	Secret    *ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret `tfsdk:"secret"`
	Username  types.String                                                                                                `tfsdk:"username"`
}

type ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret struct {
	Hash types.String `tfsdk:"hash"`
}

type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRs struct {
	Devices *[]RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesRs `tfsdk:"devices"`
	Options *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareOptionsRs   `tfsdk:"options"`
}

type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesRs struct {
	Sudi   types.String                                                                                    `tfsdk:"sudi"`
	Tunnel *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesTunnelRs `tfsdk:"tunnel"`
	User   *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesUserRs   `tfsdk:"user"`
	Vty    *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyRs    `tfsdk:"vty"`
}

type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesTunnelRs struct {
	CertificateName types.String `tfsdk:"certificate_name"`
	LocalInterface  types.Int64  `tfsdk:"local_interface"`
	LoopbackNumber  types.Int64  `tfsdk:"loopback_number"`
	Name            types.String `tfsdk:"name"`
}

type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesUserRs struct {
	Username types.String `tfsdk:"username"`
}

type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyRs struct {
	AccessList      *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListRs     `tfsdk:"access_list"`
	Authentication  *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthenticationRs `tfsdk:"authentication"`
	Authorization   *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthorizationRs  `tfsdk:"authorization"`
	EndLineNumber   types.Int64                                                                                                `tfsdk:"end_line_number"`
	RotaryNumber    types.Int64                                                                                                `tfsdk:"rotary_number"`
	StartLineNumber types.Int64                                                                                                `tfsdk:"start_line_number"`
}

type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListRs struct {
	VtyIn  *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListVtyInRs  `tfsdk:"vty_in"`
	VtyOut *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListVtyOutRs `tfsdk:"vty_out"`
}

type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListVtyInRs struct {
	Name types.String `tfsdk:"name"`
}

type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListVtyOutRs struct {
	Name types.String `tfsdk:"name"`
}

type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthenticationRs struct {
	Group *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthenticationGroupRs `tfsdk:"group"`
}

type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthenticationGroupRs struct {
	Name types.String `tfsdk:"name"`
}

type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthorizationRs struct {
	Group *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthorizationGroupRs `tfsdk:"group"`
}

type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthorizationGroupRs struct {
	Name types.String `tfsdk:"name"`
}

type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareOptionsRs struct {
	SkipCommit types.Bool `tfsdk:"skip_commit"`
}

// FromBody
func (r *OrganizationsInventoryOnboardingCloudMonitoringPrepare) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepare {
	re := *r.Parameters
	var requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevices []merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevices

	if re.Devices != nil {
		for _, rItem1 := range *re.Devices {
			sudi := rItem1.Sudi.ValueString()
			var requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesTunnel *merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesTunnel

			if rItem1.Tunnel != nil {
				certificateName := rItem1.Tunnel.CertificateName.ValueString()
				localInterface := func() *int64 {
					if !rItem1.Tunnel.LocalInterface.IsUnknown() && !rItem1.Tunnel.LocalInterface.IsNull() {
						return rItem1.Tunnel.LocalInterface.ValueInt64Pointer()
					}
					return nil
				}()
				loopbackNumber := func() *int64 {
					if !rItem1.Tunnel.LoopbackNumber.IsUnknown() && !rItem1.Tunnel.LoopbackNumber.IsNull() {
						return rItem1.Tunnel.LoopbackNumber.ValueInt64Pointer()
					}
					return nil
				}()
				name := rItem1.Tunnel.Name.ValueString()
				requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesTunnel = &merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesTunnel{
					CertificateName: certificateName,
					LocalInterface:  int64ToIntPointer(localInterface),
					LoopbackNumber:  int64ToIntPointer(loopbackNumber),
					Name:            name,
				}
				//[debug] Is Array: False
			}
			var requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesUser *merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesUser

			if rItem1.User != nil {
				username := rItem1.User.Username.ValueString()
				requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesUser = &merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesUser{
					Username: username,
				}
				//[debug] Is Array: False
			}
			var requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVty *merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVty

			if rItem1.Vty != nil {
				var requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessList *merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessList

				if rItem1.Vty.AccessList != nil {
					var requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListVtyIn *merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListVtyIn

					if rItem1.Vty.AccessList.VtyIn != nil {
						name := rItem1.Vty.AccessList.VtyIn.Name.ValueString()
						requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListVtyIn = &merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListVtyIn{
							Name: name,
						}
						//[debug] Is Array: False
					}
					var requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListVtyOut *merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListVtyOut

					if rItem1.Vty.AccessList.VtyOut != nil {
						name := rItem1.Vty.AccessList.VtyOut.Name.ValueString()
						requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListVtyOut = &merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListVtyOut{
							Name: name,
						}
						//[debug] Is Array: False
					}
					requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessList = &merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessList{
						VtyIn:  requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListVtyIn,
						VtyOut: requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessListVtyOut,
					}
					//[debug] Is Array: False
				}
				var requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthentication *merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthentication

				if rItem1.Vty.Authentication != nil {
					var requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthenticationGroup *merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthenticationGroup

					if rItem1.Vty.Authentication.Group != nil {
						name := rItem1.Vty.Authentication.Group.Name.ValueString()
						requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthenticationGroup = &merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthenticationGroup{
							Name: name,
						}
						//[debug] Is Array: False
					}
					requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthentication = &merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthentication{
						Group: requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthenticationGroup,
					}
					//[debug] Is Array: False
				}
				var requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthorization *merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthorization

				if rItem1.Vty.Authorization != nil {
					var requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthorizationGroup *merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthorizationGroup

					if rItem1.Vty.Authorization.Group != nil {
						name := rItem1.Vty.Authorization.Group.Name.ValueString()
						requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthorizationGroup = &merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthorizationGroup{
							Name: name,
						}
						//[debug] Is Array: False
					}
					requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthorization = &merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthorization{
						Group: requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthorizationGroup,
					}
					//[debug] Is Array: False
				}
				endLineNumber := func() *int64 {
					if !rItem1.Vty.EndLineNumber.IsUnknown() && !rItem1.Vty.EndLineNumber.IsNull() {
						return rItem1.Vty.EndLineNumber.ValueInt64Pointer()
					}
					return nil
				}()
				rotaryNumber := func() *int64 {
					if !rItem1.Vty.RotaryNumber.IsUnknown() && !rItem1.Vty.RotaryNumber.IsNull() {
						return rItem1.Vty.RotaryNumber.ValueInt64Pointer()
					}
					return nil
				}()
				startLineNumber := func() *int64 {
					if !rItem1.Vty.StartLineNumber.IsUnknown() && !rItem1.Vty.StartLineNumber.IsNull() {
						return rItem1.Vty.StartLineNumber.ValueInt64Pointer()
					}
					return nil
				}()
				requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVty = &merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVty{
					AccessList:      requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAccessList,
					Authentication:  requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthentication,
					Authorization:   requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVtyAuthorization,
					EndLineNumber:   int64ToIntPointer(endLineNumber),
					RotaryNumber:    int64ToIntPointer(rotaryNumber),
					StartLineNumber: int64ToIntPointer(startLineNumber),
				}
				//[debug] Is Array: False
			}
			requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevices = append(requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevices, merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevices{
				Sudi:   sudi,
				Tunnel: requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesTunnel,
				User:   requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesUser,
				Vty:    requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevicesVty,
			})
			//[debug] Is Array: True
		}
	}
	var requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareOptions *merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareOptions

	if re.Options != nil {
		skipCommit := func() *bool {
			if !re.Options.SkipCommit.IsUnknown() && !re.Options.SkipCommit.IsNull() {
				return re.Options.SkipCommit.ValueBoolPointer()
			}
			return nil
		}()
		requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareOptions = &merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareOptions{
			SkipCommit: skipCommit,
		}
		//[debug] Is Array: False
	}
	out := merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepare{
		Devices: func() *[]merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevices {
			if len(requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevices) > 0 {
				return &requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareDevices
			}
			return nil
		}(),
		Options: requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareOptions,
	}
	return &out
}

// ToBody
func ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareItemsToBody(state OrganizationsInventoryOnboardingCloudMonitoringPrepare, response *merakigosdk.ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepare) OrganizationsInventoryOnboardingCloudMonitoringPrepare {
	var items []ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepare
	for _, item := range *response {
		itemState := ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepare{
			ConfigParams: func() *ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParams {
				if item.ConfigParams != nil {
					return &ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParams{
						CloudStaticIP: types.StringValue(item.ConfigParams.CloudStaticIP),
						Tunnel: func() *ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsTunnel {
							if item.ConfigParams.Tunnel != nil {
								return &ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsTunnel{
									Host: types.StringValue(item.ConfigParams.Tunnel.Host),
									Mode: types.StringValue(item.ConfigParams.Tunnel.Mode),
									Name: types.StringValue(item.ConfigParams.Tunnel.Name),
									Port: types.StringValue(item.ConfigParams.Tunnel.Port),
									RootCertificate: func() *ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsTunnelRootCertificate {
										if item.ConfigParams.Tunnel.RootCertificate != nil {
											return &ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsTunnelRootCertificate{
												Content: types.StringValue(item.ConfigParams.Tunnel.RootCertificate.Content),
												Name:    types.StringValue(item.ConfigParams.Tunnel.RootCertificate.Name),
											}
										}
										return nil
									}(),
								}
							}
							return nil
						}(),
						User: func() *ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsUser {
							if item.ConfigParams.User != nil {
								return &ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsUser{
									PublicKey: types.StringValue(item.ConfigParams.User.PublicKey),
									Secret: func() *ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret {
										if item.ConfigParams.User.Secret != nil {
											return &ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringPrepareConfigParamsUserSecret{
												Hash: types.StringValue(item.ConfigParams.User.Secret.Hash),
											}
										}
										return nil
									}(),
									Username: types.StringValue(item.ConfigParams.User.Username),
								}
							}
							return nil
						}(),
					}
				}
				return nil
			}(),
			DeviceID: types.StringValue(item.DeviceID),
			Message:  types.StringValue(item.Message),
			Status:   types.StringValue(item.Status),
			Udi:      types.StringValue(item.Udi),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
