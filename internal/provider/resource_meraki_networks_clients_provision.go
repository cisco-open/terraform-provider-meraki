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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksClientsProvisionResource{}
	_ resource.ResourceWithConfigure = &NetworksClientsProvisionResource{}
)

func NewNetworksClientsProvisionResource() resource.Resource {
	return &NetworksClientsProvisionResource{}
}

type NetworksClientsProvisionResource struct {
	client *merakigosdk.Client
}

func (r *NetworksClientsProvisionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksClientsProvisionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_clients_provision"
}

// resourceAction
func (r *NetworksClientsProvisionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"clients": schema.SetNestedAttribute{
						MarkdownDescription: `The list of clients to provision`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"client_id": schema.StringAttribute{
									MarkdownDescription: `The identifier of the client`,
									Computed:            true,
								},
								"mac": schema.StringAttribute{
									MarkdownDescription: `The MAC address of the client`,
									Computed:            true,
								},
								"message": schema.StringAttribute{
									MarkdownDescription: `The client's display message if its group policy is 'Blocked'`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `The name of the client`,
									Computed:            true,
								},
							},
						},
					},
					"device_policy": schema.StringAttribute{
						MarkdownDescription: `The name of the client's policy`,
						Computed:            true,
					},
					"group_policy_id": schema.StringAttribute{
						MarkdownDescription: `The group policy identifier of the client`,
						Computed:            true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"clients": schema.SetNestedAttribute{
						MarkdownDescription: `The array of clients to provision`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"mac": schema.StringAttribute{
									MarkdownDescription: `The MAC address of the client. Required.`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `The display name for the client. Optional. Limited to 255 bytes.`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
							},
						},
					},
					"device_policy": schema.StringAttribute{
						MarkdownDescription: `The policy to apply to the specified client. Can be 'Group policy', 'Allowed', 'Blocked', 'Per connection' or 'Normal'. Required.
                                        Allowed values: [Allowed,Blocked,Group policy,Normal,Per connection]`,
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"group_policy_id": schema.StringAttribute{
						MarkdownDescription: `The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"policies_by_security_appliance": schema.SingleNestedAttribute{
						MarkdownDescription: `An object, describing what the policy-connection association is for the security appliance. (Only relevant if the security appliance is actually within the network)`,
						Optional:            true,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"device_policy": schema.StringAttribute{
								MarkdownDescription: `The policy to apply to the specified client. Can be 'Allowed', 'Blocked' or 'Normal'. Required.
                                              Allowed values: [Allowed,Blocked,Normal]`,
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
							},
						},
					},
					"policies_by_ssid": schema.SingleNestedAttribute{
						MarkdownDescription: `An object, describing the policy-connection associations for each active SSID within the network. Keys should be the number of enabled SSIDs, mapping to an object describing the client's policy`,
						Optional:            true,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"status_0": schema.SingleNestedAttribute{
								MarkdownDescription: `The number for the SSID`,
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"device_policy": schema.StringAttribute{
										MarkdownDescription: `The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
                                                    Allowed values: [Allowed,Blocked,Group policy,Normal]`,
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
									"group_policy_id": schema.StringAttribute{
										MarkdownDescription: `The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.`,
										Optional:            true,
										Computed:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
								},
							},
							"status_1": schema.SingleNestedAttribute{
								MarkdownDescription: `The number for the SSID`,
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"device_policy": schema.StringAttribute{
										MarkdownDescription: `The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
                                                    Allowed values: [Allowed,Blocked,Group policy,Normal]`,
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
									"group_policy_id": schema.StringAttribute{
										MarkdownDescription: `The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.`,
										Optional:            true,
										Computed:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
								},
							},
							"status_10": schema.SingleNestedAttribute{
								MarkdownDescription: `The number for the SSID`,
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"device_policy": schema.StringAttribute{
										MarkdownDescription: `The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
                                                    Allowed values: [Allowed,Blocked,Group policy,Normal]`,
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
									"group_policy_id": schema.StringAttribute{
										MarkdownDescription: `The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.`,
										Optional:            true,
										Computed:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
								},
							},
							"status_11": schema.SingleNestedAttribute{
								MarkdownDescription: `The number for the SSID`,
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"device_policy": schema.StringAttribute{
										MarkdownDescription: `The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
                                                    Allowed values: [Allowed,Blocked,Group policy,Normal]`,
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
									"group_policy_id": schema.StringAttribute{
										MarkdownDescription: `The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.`,
										Optional:            true,
										Computed:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
								},
							},
							"status_12": schema.SingleNestedAttribute{
								MarkdownDescription: `The number for the SSID`,
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"device_policy": schema.StringAttribute{
										MarkdownDescription: `The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
                                                    Allowed values: [Allowed,Blocked,Group policy,Normal]`,
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
									"group_policy_id": schema.StringAttribute{
										MarkdownDescription: `The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.`,
										Optional:            true,
										Computed:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
								},
							},
							"status_13": schema.SingleNestedAttribute{
								MarkdownDescription: `The number for the SSID`,
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"device_policy": schema.StringAttribute{
										MarkdownDescription: `The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
                                                    Allowed values: [Allowed,Blocked,Group policy,Normal]`,
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
									"group_policy_id": schema.StringAttribute{
										MarkdownDescription: `The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.`,
										Optional:            true,
										Computed:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
								},
							},
							"status_14": schema.SingleNestedAttribute{
								MarkdownDescription: `The number for the SSID`,
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"device_policy": schema.StringAttribute{
										MarkdownDescription: `The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
                                                    Allowed values: [Allowed,Blocked,Group policy,Normal]`,
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
									"group_policy_id": schema.StringAttribute{
										MarkdownDescription: `The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.`,
										Optional:            true,
										Computed:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
								},
							},
							"status_2": schema.SingleNestedAttribute{
								MarkdownDescription: `The number for the SSID`,
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"device_policy": schema.StringAttribute{
										MarkdownDescription: `The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
                                                    Allowed values: [Allowed,Blocked,Group policy,Normal]`,
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
									"group_policy_id": schema.StringAttribute{
										MarkdownDescription: `The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.`,
										Optional:            true,
										Computed:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
								},
							},
							"status_3": schema.SingleNestedAttribute{
								MarkdownDescription: `The number for the SSID`,
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"device_policy": schema.StringAttribute{
										MarkdownDescription: `The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
                                                    Allowed values: [Allowed,Blocked,Group policy,Normal]`,
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
									"group_policy_id": schema.StringAttribute{
										MarkdownDescription: `The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.`,
										Optional:            true,
										Computed:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
								},
							},
							"status_4": schema.SingleNestedAttribute{
								MarkdownDescription: `The number for the SSID`,
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"device_policy": schema.StringAttribute{
										MarkdownDescription: `The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
                                                    Allowed values: [Allowed,Blocked,Group policy,Normal]`,
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
									"group_policy_id": schema.StringAttribute{
										MarkdownDescription: `The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.`,
										Optional:            true,
										Computed:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
								},
							},
							"status_5": schema.SingleNestedAttribute{
								MarkdownDescription: `The number for the SSID`,
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"device_policy": schema.StringAttribute{
										MarkdownDescription: `The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
                                                    Allowed values: [Allowed,Blocked,Group policy,Normal]`,
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
									"group_policy_id": schema.StringAttribute{
										MarkdownDescription: `The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.`,
										Optional:            true,
										Computed:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
								},
							},
							"status_6": schema.SingleNestedAttribute{
								MarkdownDescription: `The number for the SSID`,
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"device_policy": schema.StringAttribute{
										MarkdownDescription: `The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
                                                    Allowed values: [Allowed,Blocked,Group policy,Normal]`,
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
									"group_policy_id": schema.StringAttribute{
										MarkdownDescription: `The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.`,
										Optional:            true,
										Computed:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
								},
							},
							"status_7": schema.SingleNestedAttribute{
								MarkdownDescription: `The number for the SSID`,
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"device_policy": schema.StringAttribute{
										MarkdownDescription: `The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
                                                    Allowed values: [Allowed,Blocked,Group policy,Normal]`,
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
									"group_policy_id": schema.StringAttribute{
										MarkdownDescription: `The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.`,
										Optional:            true,
										Computed:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
								},
							},
							"status_8": schema.SingleNestedAttribute{
								MarkdownDescription: `The number for the SSID`,
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"device_policy": schema.StringAttribute{
										MarkdownDescription: `The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
                                                    Allowed values: [Allowed,Blocked,Group policy,Normal]`,
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
									"group_policy_id": schema.StringAttribute{
										MarkdownDescription: `The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.`,
										Optional:            true,
										Computed:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
								},
							},
							"status_9": schema.SingleNestedAttribute{
								MarkdownDescription: `The number for the SSID`,
								Optional:            true,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"device_policy": schema.StringAttribute{
										MarkdownDescription: `The policy to apply to the specified client. Can be 'Allowed', 'Blocked', 'Normal' or 'Group policy'. Required.
                                                    Allowed values: [Allowed,Blocked,Group policy,Normal]`,
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
									"group_policy_id": schema.StringAttribute{
										MarkdownDescription: `The ID of the desired group policy to apply to the client. Required if 'devicePolicy' is set to "Group policy". Otherwise this is ignored.`,
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
				},
			},
		},
	}
}
func (r *NetworksClientsProvisionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksClientsProvision

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
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Networks.ProvisionNetworkClients(vvNetworkID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing ProvisionNetworkClients",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing ProvisionNetworkClients",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseNetworksProvisionNetworkClientsItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksClientsProvisionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksClientsProvisionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksClientsProvisionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksClientsProvision struct {
	NetworkID  types.String                              `tfsdk:"network_id"`
	Item       *ResponseNetworksProvisionNetworkClients  `tfsdk:"item"`
	Parameters *RequestNetworksProvisionNetworkClientsRs `tfsdk:"parameters"`
}

type ResponseNetworksProvisionNetworkClients struct {
	Clients       *[]ResponseNetworksProvisionNetworkClientsClients `tfsdk:"clients"`
	DevicePolicy  types.String                                      `tfsdk:"device_policy"`
	GroupPolicyID types.String                                      `tfsdk:"group_policy_id"`
}

type ResponseNetworksProvisionNetworkClientsClients struct {
	ClientID types.String `tfsdk:"client_id"`
	Mac      types.String `tfsdk:"mac"`
	Message  types.String `tfsdk:"message"`
	Name     types.String `tfsdk:"name"`
}

type RequestNetworksProvisionNetworkClientsRs struct {
	Clients                     *[]RequestNetworksProvisionNetworkClientsClientsRs                   `tfsdk:"clients"`
	DevicePolicy                types.String                                                         `tfsdk:"device_policy"`
	GroupPolicyID               types.String                                                         `tfsdk:"group_policy_id"`
	PoliciesBySecurityAppliance *RequestNetworksProvisionNetworkClientsPoliciesBySecurityApplianceRs `tfsdk:"policies_by_security_appliance"`
	PoliciesBySSID              *RequestNetworksProvisionNetworkClientsPoliciesBySsidRs              `tfsdk:"policies_by_ssid"`
}

type RequestNetworksProvisionNetworkClientsClientsRs struct {
	Mac  types.String `tfsdk:"mac"`
	Name types.String `tfsdk:"name"`
}

type RequestNetworksProvisionNetworkClientsPoliciesBySecurityApplianceRs struct {
	DevicePolicy types.String `tfsdk:"device_policy"`
}

type RequestNetworksProvisionNetworkClientsPoliciesBySsidRs struct {
	Status0  *RequestNetworksProvisionNetworkClientsPoliciesBySsid0Rs  `tfsdk:"status_0"`
	Status1  *RequestNetworksProvisionNetworkClientsPoliciesBySsid1Rs  `tfsdk:"status_1"`
	Status10 *RequestNetworksProvisionNetworkClientsPoliciesBySsid10Rs `tfsdk:"status_10"`
	Status11 *RequestNetworksProvisionNetworkClientsPoliciesBySsid11Rs `tfsdk:"status_11"`
	Status12 *RequestNetworksProvisionNetworkClientsPoliciesBySsid12Rs `tfsdk:"status_12"`
	Status13 *RequestNetworksProvisionNetworkClientsPoliciesBySsid13Rs `tfsdk:"status_13"`
	Status14 *RequestNetworksProvisionNetworkClientsPoliciesBySsid14Rs `tfsdk:"status_14"`
	Status2  *RequestNetworksProvisionNetworkClientsPoliciesBySsid2Rs  `tfsdk:"status_2"`
	Status3  *RequestNetworksProvisionNetworkClientsPoliciesBySsid3Rs  `tfsdk:"status_3"`
	Status4  *RequestNetworksProvisionNetworkClientsPoliciesBySsid4Rs  `tfsdk:"status_4"`
	Status5  *RequestNetworksProvisionNetworkClientsPoliciesBySsid5Rs  `tfsdk:"status_5"`
	Status6  *RequestNetworksProvisionNetworkClientsPoliciesBySsid6Rs  `tfsdk:"status_6"`
	Status7  *RequestNetworksProvisionNetworkClientsPoliciesBySsid7Rs  `tfsdk:"status_7"`
	Status8  *RequestNetworksProvisionNetworkClientsPoliciesBySsid8Rs  `tfsdk:"status_8"`
	Status9  *RequestNetworksProvisionNetworkClientsPoliciesBySsid9Rs  `tfsdk:"status_9"`
}

type RequestNetworksProvisionNetworkClientsPoliciesBySsid0Rs struct {
	DevicePolicy  types.String `tfsdk:"device_policy"`
	GroupPolicyID types.String `tfsdk:"group_policy_id"`
}

type RequestNetworksProvisionNetworkClientsPoliciesBySsid1Rs struct {
	DevicePolicy  types.String `tfsdk:"device_policy"`
	GroupPolicyID types.String `tfsdk:"group_policy_id"`
}

type RequestNetworksProvisionNetworkClientsPoliciesBySsid10Rs struct {
	DevicePolicy  types.String `tfsdk:"device_policy"`
	GroupPolicyID types.String `tfsdk:"group_policy_id"`
}

type RequestNetworksProvisionNetworkClientsPoliciesBySsid11Rs struct {
	DevicePolicy  types.String `tfsdk:"device_policy"`
	GroupPolicyID types.String `tfsdk:"group_policy_id"`
}

type RequestNetworksProvisionNetworkClientsPoliciesBySsid12Rs struct {
	DevicePolicy  types.String `tfsdk:"device_policy"`
	GroupPolicyID types.String `tfsdk:"group_policy_id"`
}

type RequestNetworksProvisionNetworkClientsPoliciesBySsid13Rs struct {
	DevicePolicy  types.String `tfsdk:"device_policy"`
	GroupPolicyID types.String `tfsdk:"group_policy_id"`
}

type RequestNetworksProvisionNetworkClientsPoliciesBySsid14Rs struct {
	DevicePolicy  types.String `tfsdk:"device_policy"`
	GroupPolicyID types.String `tfsdk:"group_policy_id"`
}

type RequestNetworksProvisionNetworkClientsPoliciesBySsid2Rs struct {
	DevicePolicy  types.String `tfsdk:"device_policy"`
	GroupPolicyID types.String `tfsdk:"group_policy_id"`
}

type RequestNetworksProvisionNetworkClientsPoliciesBySsid3Rs struct {
	DevicePolicy  types.String `tfsdk:"device_policy"`
	GroupPolicyID types.String `tfsdk:"group_policy_id"`
}

type RequestNetworksProvisionNetworkClientsPoliciesBySsid4Rs struct {
	DevicePolicy  types.String `tfsdk:"device_policy"`
	GroupPolicyID types.String `tfsdk:"group_policy_id"`
}

type RequestNetworksProvisionNetworkClientsPoliciesBySsid5Rs struct {
	DevicePolicy  types.String `tfsdk:"device_policy"`
	GroupPolicyID types.String `tfsdk:"group_policy_id"`
}

type RequestNetworksProvisionNetworkClientsPoliciesBySsid6Rs struct {
	DevicePolicy  types.String `tfsdk:"device_policy"`
	GroupPolicyID types.String `tfsdk:"group_policy_id"`
}

type RequestNetworksProvisionNetworkClientsPoliciesBySsid7Rs struct {
	DevicePolicy  types.String `tfsdk:"device_policy"`
	GroupPolicyID types.String `tfsdk:"group_policy_id"`
}

type RequestNetworksProvisionNetworkClientsPoliciesBySsid8Rs struct {
	DevicePolicy  types.String `tfsdk:"device_policy"`
	GroupPolicyID types.String `tfsdk:"group_policy_id"`
}

type RequestNetworksProvisionNetworkClientsPoliciesBySsid9Rs struct {
	DevicePolicy  types.String `tfsdk:"device_policy"`
	GroupPolicyID types.String `tfsdk:"group_policy_id"`
}

// FromBody
func (r *NetworksClientsProvision) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksProvisionNetworkClients {
	emptyString := ""
	re := *r.Parameters
	var requestNetworksProvisionNetworkClientsClients []merakigosdk.RequestNetworksProvisionNetworkClientsClients

	if re.Clients != nil {
		for _, rItem1 := range *re.Clients {
			mac := rItem1.Mac.ValueString()
			name := rItem1.Name.ValueString()
			requestNetworksProvisionNetworkClientsClients = append(requestNetworksProvisionNetworkClientsClients, merakigosdk.RequestNetworksProvisionNetworkClientsClients{
				Mac:  mac,
				Name: name,
			})
			//[debug] Is Array: True
		}
	}
	devicePolicy := new(string)
	if !re.DevicePolicy.IsUnknown() && !re.DevicePolicy.IsNull() {
		*devicePolicy = re.DevicePolicy.ValueString()
	} else {
		devicePolicy = &emptyString
	}
	groupPolicyID := new(string)
	if !re.GroupPolicyID.IsUnknown() && !re.GroupPolicyID.IsNull() {
		*groupPolicyID = re.GroupPolicyID.ValueString()
	} else {
		groupPolicyID = &emptyString
	}
	var requestNetworksProvisionNetworkClientsPoliciesBySecurityAppliance *merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySecurityAppliance

	if re.PoliciesBySecurityAppliance != nil {
		devicePolicy := re.PoliciesBySecurityAppliance.DevicePolicy.ValueString()
		requestNetworksProvisionNetworkClientsPoliciesBySecurityAppliance = &merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySecurityAppliance{
			DevicePolicy: devicePolicy,
		}
		//[debug] Is Array: False
	}
	var requestNetworksProvisionNetworkClientsPoliciesBySSID *merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID

	if re.PoliciesBySSID != nil {
		var requestNetworksProvisionNetworkClientsPoliciesBySSID0 *merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID0

		if re.PoliciesBySSID.Status0 != nil {
			devicePolicy := re.PoliciesBySSID.Status0.DevicePolicy.ValueString()
			groupPolicyID := re.PoliciesBySSID.Status0.GroupPolicyID.ValueString()
			requestNetworksProvisionNetworkClientsPoliciesBySSID0 = &merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID0{
				DevicePolicy:  devicePolicy,
				GroupPolicyID: groupPolicyID,
			}
			//[debug] Is Array: False
		}
		var requestNetworksProvisionNetworkClientsPoliciesBySSID1 *merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID1

		if re.PoliciesBySSID.Status1 != nil {
			devicePolicy := re.PoliciesBySSID.Status1.DevicePolicy.ValueString()
			groupPolicyID := re.PoliciesBySSID.Status1.GroupPolicyID.ValueString()
			requestNetworksProvisionNetworkClientsPoliciesBySSID1 = &merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID1{
				DevicePolicy:  devicePolicy,
				GroupPolicyID: groupPolicyID,
			}
			//[debug] Is Array: False
		}
		var requestNetworksProvisionNetworkClientsPoliciesBySSID10 *merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID10

		if re.PoliciesBySSID.Status10 != nil {
			devicePolicy := re.PoliciesBySSID.Status10.DevicePolicy.ValueString()
			groupPolicyID := re.PoliciesBySSID.Status10.GroupPolicyID.ValueString()
			requestNetworksProvisionNetworkClientsPoliciesBySSID10 = &merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID10{
				DevicePolicy:  devicePolicy,
				GroupPolicyID: groupPolicyID,
			}
			//[debug] Is Array: False
		}
		var requestNetworksProvisionNetworkClientsPoliciesBySSID11 *merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID11

		if re.PoliciesBySSID.Status11 != nil {
			devicePolicy := re.PoliciesBySSID.Status11.DevicePolicy.ValueString()
			groupPolicyID := re.PoliciesBySSID.Status11.GroupPolicyID.ValueString()
			requestNetworksProvisionNetworkClientsPoliciesBySSID11 = &merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID11{
				DevicePolicy:  devicePolicy,
				GroupPolicyID: groupPolicyID,
			}
			//[debug] Is Array: False
		}
		var requestNetworksProvisionNetworkClientsPoliciesBySSID12 *merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID12

		if re.PoliciesBySSID.Status12 != nil {
			devicePolicy := re.PoliciesBySSID.Status12.DevicePolicy.ValueString()
			groupPolicyID := re.PoliciesBySSID.Status12.GroupPolicyID.ValueString()
			requestNetworksProvisionNetworkClientsPoliciesBySSID12 = &merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID12{
				DevicePolicy:  devicePolicy,
				GroupPolicyID: groupPolicyID,
			}
			//[debug] Is Array: False
		}
		var requestNetworksProvisionNetworkClientsPoliciesBySSID13 *merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID13

		if re.PoliciesBySSID.Status13 != nil {
			devicePolicy := re.PoliciesBySSID.Status13.DevicePolicy.ValueString()
			groupPolicyID := re.PoliciesBySSID.Status13.GroupPolicyID.ValueString()
			requestNetworksProvisionNetworkClientsPoliciesBySSID13 = &merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID13{
				DevicePolicy:  devicePolicy,
				GroupPolicyID: groupPolicyID,
			}
			//[debug] Is Array: False
		}
		var requestNetworksProvisionNetworkClientsPoliciesBySSID14 *merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID14

		if re.PoliciesBySSID.Status14 != nil {
			devicePolicy := re.PoliciesBySSID.Status14.DevicePolicy.ValueString()
			groupPolicyID := re.PoliciesBySSID.Status14.GroupPolicyID.ValueString()
			requestNetworksProvisionNetworkClientsPoliciesBySSID14 = &merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID14{
				DevicePolicy:  devicePolicy,
				GroupPolicyID: groupPolicyID,
			}
			//[debug] Is Array: False
		}
		var requestNetworksProvisionNetworkClientsPoliciesBySSID2 *merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID2

		if re.PoliciesBySSID.Status2 != nil {
			devicePolicy := re.PoliciesBySSID.Status2.DevicePolicy.ValueString()
			groupPolicyID := re.PoliciesBySSID.Status2.GroupPolicyID.ValueString()
			requestNetworksProvisionNetworkClientsPoliciesBySSID2 = &merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID2{
				DevicePolicy:  devicePolicy,
				GroupPolicyID: groupPolicyID,
			}
			//[debug] Is Array: False
		}
		var requestNetworksProvisionNetworkClientsPoliciesBySSID3 *merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID3

		if re.PoliciesBySSID.Status3 != nil {
			devicePolicy := re.PoliciesBySSID.Status3.DevicePolicy.ValueString()
			groupPolicyID := re.PoliciesBySSID.Status3.GroupPolicyID.ValueString()
			requestNetworksProvisionNetworkClientsPoliciesBySSID3 = &merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID3{
				DevicePolicy:  devicePolicy,
				GroupPolicyID: groupPolicyID,
			}
			//[debug] Is Array: False
		}
		var requestNetworksProvisionNetworkClientsPoliciesBySSID4 *merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID4

		if re.PoliciesBySSID.Status4 != nil {
			devicePolicy := re.PoliciesBySSID.Status4.DevicePolicy.ValueString()
			groupPolicyID := re.PoliciesBySSID.Status4.GroupPolicyID.ValueString()
			requestNetworksProvisionNetworkClientsPoliciesBySSID4 = &merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID4{
				DevicePolicy:  devicePolicy,
				GroupPolicyID: groupPolicyID,
			}
			//[debug] Is Array: False
		}
		var requestNetworksProvisionNetworkClientsPoliciesBySSID5 *merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID5

		if re.PoliciesBySSID.Status5 != nil {
			devicePolicy := re.PoliciesBySSID.Status5.DevicePolicy.ValueString()
			groupPolicyID := re.PoliciesBySSID.Status5.GroupPolicyID.ValueString()
			requestNetworksProvisionNetworkClientsPoliciesBySSID5 = &merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID5{
				DevicePolicy:  devicePolicy,
				GroupPolicyID: groupPolicyID,
			}
			//[debug] Is Array: False
		}
		var requestNetworksProvisionNetworkClientsPoliciesBySSID6 *merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID6

		if re.PoliciesBySSID.Status6 != nil {
			devicePolicy := re.PoliciesBySSID.Status6.DevicePolicy.ValueString()
			groupPolicyID := re.PoliciesBySSID.Status6.GroupPolicyID.ValueString()
			requestNetworksProvisionNetworkClientsPoliciesBySSID6 = &merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID6{
				DevicePolicy:  devicePolicy,
				GroupPolicyID: groupPolicyID,
			}
			//[debug] Is Array: False
		}
		var requestNetworksProvisionNetworkClientsPoliciesBySSID7 *merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID7

		if re.PoliciesBySSID.Status7 != nil {
			devicePolicy := re.PoliciesBySSID.Status7.DevicePolicy.ValueString()
			groupPolicyID := re.PoliciesBySSID.Status7.GroupPolicyID.ValueString()
			requestNetworksProvisionNetworkClientsPoliciesBySSID7 = &merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID7{
				DevicePolicy:  devicePolicy,
				GroupPolicyID: groupPolicyID,
			}
			//[debug] Is Array: False
		}
		var requestNetworksProvisionNetworkClientsPoliciesBySSID8 *merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID8

		if re.PoliciesBySSID.Status8 != nil {
			devicePolicy := re.PoliciesBySSID.Status8.DevicePolicy.ValueString()
			groupPolicyID := re.PoliciesBySSID.Status8.GroupPolicyID.ValueString()
			requestNetworksProvisionNetworkClientsPoliciesBySSID8 = &merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID8{
				DevicePolicy:  devicePolicy,
				GroupPolicyID: groupPolicyID,
			}
			//[debug] Is Array: False
		}
		var requestNetworksProvisionNetworkClientsPoliciesBySSID9 *merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID9

		if re.PoliciesBySSID.Status9 != nil {
			devicePolicy := re.PoliciesBySSID.Status9.DevicePolicy.ValueString()
			groupPolicyID := re.PoliciesBySSID.Status9.GroupPolicyID.ValueString()
			requestNetworksProvisionNetworkClientsPoliciesBySSID9 = &merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID9{
				DevicePolicy:  devicePolicy,
				GroupPolicyID: groupPolicyID,
			}
			//[debug] Is Array: False
		}
		requestNetworksProvisionNetworkClientsPoliciesBySSID = &merakigosdk.RequestNetworksProvisionNetworkClientsPoliciesBySSID{
			Status0:  requestNetworksProvisionNetworkClientsPoliciesBySSID0,
			Status1:  requestNetworksProvisionNetworkClientsPoliciesBySSID1,
			Status10: requestNetworksProvisionNetworkClientsPoliciesBySSID10,
			Status11: requestNetworksProvisionNetworkClientsPoliciesBySSID11,
			Status12: requestNetworksProvisionNetworkClientsPoliciesBySSID12,
			Status13: requestNetworksProvisionNetworkClientsPoliciesBySSID13,
			Status14: requestNetworksProvisionNetworkClientsPoliciesBySSID14,
			Status2:  requestNetworksProvisionNetworkClientsPoliciesBySSID2,
			Status3:  requestNetworksProvisionNetworkClientsPoliciesBySSID3,
			Status4:  requestNetworksProvisionNetworkClientsPoliciesBySSID4,
			Status5:  requestNetworksProvisionNetworkClientsPoliciesBySSID5,
			Status6:  requestNetworksProvisionNetworkClientsPoliciesBySSID6,
			Status7:  requestNetworksProvisionNetworkClientsPoliciesBySSID7,
			Status8:  requestNetworksProvisionNetworkClientsPoliciesBySSID8,
			Status9:  requestNetworksProvisionNetworkClientsPoliciesBySSID9,
		}
		//[debug] Is Array: False
	}
	out := merakigosdk.RequestNetworksProvisionNetworkClients{
		Clients:                     &requestNetworksProvisionNetworkClientsClients,
		DevicePolicy:                *devicePolicy,
		GroupPolicyID:               *groupPolicyID,
		PoliciesBySecurityAppliance: requestNetworksProvisionNetworkClientsPoliciesBySecurityAppliance,
		PoliciesBySSID:              requestNetworksProvisionNetworkClientsPoliciesBySSID,
	}
	return &out
}

// ToBody
func ResponseNetworksProvisionNetworkClientsItemToBody(state NetworksClientsProvision, response *merakigosdk.ResponseNetworksProvisionNetworkClients) NetworksClientsProvision {
	itemState := ResponseNetworksProvisionNetworkClients{
		Clients: func() *[]ResponseNetworksProvisionNetworkClientsClients {
			if response.Clients != nil {
				result := make([]ResponseNetworksProvisionNetworkClientsClients, len(*response.Clients))
				for i, clients := range *response.Clients {
					result[i] = ResponseNetworksProvisionNetworkClientsClients{
						ClientID: func() types.String {
							if clients.ClientID != "" {
								return types.StringValue(clients.ClientID)
							}
							return types.String{}
						}(),
						Mac: func() types.String {
							if clients.Mac != "" {
								return types.StringValue(clients.Mac)
							}
							return types.String{}
						}(),
						Message: func() types.String {
							if clients.Message != "" {
								return types.StringValue(clients.Message)
							}
							return types.String{}
						}(),
						Name: func() types.String {
							if clients.Name != "" {
								return types.StringValue(clients.Name)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		DevicePolicy: func() types.String {
			if response.DevicePolicy != "" {
				return types.StringValue(response.DevicePolicy)
			}
			return types.String{}
		}(),
		GroupPolicyID: func() types.String {
			if response.GroupPolicyID != "" {
				return types.StringValue(response.GroupPolicyID)
			}
			return types.String{}
		}(),
	}
	state.Item = &itemState
	return state
}
