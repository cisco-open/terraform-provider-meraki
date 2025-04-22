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
	_ resource.Resource              = &OrganizationsApplianceVpnThirdPartyVpnpeersResource{}
	_ resource.ResourceWithConfigure = &OrganizationsApplianceVpnThirdPartyVpnpeersResource{}
)

func NewOrganizationsApplianceVpnThirdPartyVpnpeersResource() resource.Resource {
	return &OrganizationsApplianceVpnThirdPartyVpnpeersResource{}
}

type OrganizationsApplianceVpnThirdPartyVpnpeersResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsApplianceVpnThirdPartyVpnpeersResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsApplianceVpnThirdPartyVpnpeersResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_appliance_vpn_third_party_vpnpeers"
}

func (r *OrganizationsApplianceVpnThirdPartyVpnpeersResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"peers": schema.SetNestedAttribute{
				MarkdownDescription: `The list of VPN peers`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"ike_version": schema.StringAttribute{
							MarkdownDescription: `[optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to '1' when omitted.
                                        Allowed values: [1,2]`,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"1",
									"2",
								),
							},
						},
						"ipsec_policies": schema.SingleNestedAttribute{
							MarkdownDescription: `Custom IPSec policies for the VPN peer. If not included and a preset has not been chosen, the default preset for IPSec policies will be used.`,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"child_auth_algo": schema.SetAttribute{
									MarkdownDescription: `This is the authentication algorithms to be used in Phase 2. The value should be an array with one of the following algorithms: 'sha256', 'sha1', 'md5'`,
									Optional:            true,
									Default:             setdefault.StaticValue(types.SetNull(types.StringType)),
									ElementType:         types.StringType,
									Computed:            true,
								},
								"child_cipher_algo": schema.SetAttribute{
									MarkdownDescription: `This is the cipher algorithms to be used in Phase 2. The value should be an array with one or more of the following algorithms: 'aes256', 'aes192', 'aes128', 'tripledes', 'des', 'null'`,
									Optional:            true,
									Default:             setdefault.StaticValue(types.SetNull(types.StringType)),
									ElementType:         types.StringType,
									Computed:            true,
								},
								"child_lifetime": schema.Int64Attribute{
									MarkdownDescription: `The lifetime of the Phase 2 SA in seconds.`,
									Optional:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"child_pfs_group": schema.SetAttribute{
									MarkdownDescription: `This is the Diffie-Hellman group to be used for Perfect Forward Secrecy in Phase 2. The value should be an array with one of the following values: 'disabled','group14', 'group5', 'group2', 'group1'`,
									Optional:            true,
									Default:             setdefault.StaticValue(types.SetNull(types.StringType)),
									ElementType:         types.StringType,
									Computed:            true,
								},
								"ike_auth_algo": schema.SetAttribute{
									MarkdownDescription: `This is the authentication algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: 'sha256', 'sha1', 'md5'`,
									Optional:            true,
									Default:             setdefault.StaticValue(types.SetNull(types.StringType)),
									ElementType:         types.StringType,
									Computed:            true,
								},
								"ike_cipher_algo": schema.SetAttribute{
									MarkdownDescription: `This is the cipher algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: 'aes256', 'aes192', 'aes128', 'tripledes', 'des'`,
									Optional:            true,
									Default:             setdefault.StaticValue(types.SetNull(types.StringType)),
									ElementType:         types.StringType,
									Computed:            true,
								},
								"ike_diffie_hellman_group": schema.SetAttribute{
									MarkdownDescription: `This is the Diffie-Hellman group to be used in Phase 1. The value should be an array with one of the following algorithms: 'group14', 'group5', 'group2', 'group1'`,
									Optional:            true,
									Default:             setdefault.StaticValue(types.SetNull(types.StringType)),
									ElementType:         types.StringType,
									Computed:            true,
								},
								"ike_lifetime": schema.Int64Attribute{
									MarkdownDescription: `The lifetime of the Phase 1 SA in seconds.`,
									Optional:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"ike_prf_algo": schema.SetAttribute{
									MarkdownDescription: `[optional] This is the pseudo-random function to be used in IKE_SA. The value should be an array with one of the following algorithms: 'prfsha256', 'prfsha1', 'prfmd5', 'default'. The 'default' option can be used to default to the Authentication algorithm.`,
									Optional:            true,
									Default:             setdefault.StaticValue(types.SetNull(types.StringType)),
									ElementType:         types.StringType,
									Computed:            true,
								},
							},
						},
						"ipsec_policies_preset": schema.StringAttribute{
							MarkdownDescription: `One of the following available presets: 'default', 'aws', 'azure', 'umbrella', 'zscaler'. If this is provided, the 'ipsecPolicies' parameter is ignored.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"local_id": schema.StringAttribute{
							MarkdownDescription: `[optional] The local ID is used to identify the MX to the peer. This will apply to all MXs this peer applies to.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the VPN peer`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"network_tags": schema.SetAttribute{
							MarkdownDescription: `A list of network tags that will connect with this peer. Use ['all'] for all networks. Use ['none'] for no networks. If not included, the default is ['all'].`,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
						},
						"private_subnets": schema.SetAttribute{
							MarkdownDescription: `The list of the private subnets of the VPN peer`,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
						},
						"public_hostname": schema.StringAttribute{
							MarkdownDescription: `[optional] The public hostname of the VPN peer`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"public_ip": schema.StringAttribute{
							MarkdownDescription: `[optional] The public IP of the VPN peer`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"remote_id": schema.StringAttribute{
							MarkdownDescription: `[optional] The remote ID is used to identify the connecting VPN peer. This can either be a valid IPv4 Address, FQDN or User FQDN.`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: `The shared secret with the VPN peer`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"peers_response": schema.SetNestedAttribute{
				MarkdownDescription: `The list of VPN peers`,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"ike_version": schema.StringAttribute{
							MarkdownDescription: `[optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to '1' when omitted.
                                        Allowed values: [1,2]`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"1",
									"2",
								),
							},
						},
						"ipsec_policies": schema.SingleNestedAttribute{
							MarkdownDescription: `Custom IPSec policies for the VPN peer. If not included and a preset has not been chosen, the default preset for IPSec policies will be used.`,
							Computed:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"child_auth_algo": schema.SetAttribute{
									MarkdownDescription: `This is the authentication algorithms to be used in Phase 2. The value should be an array with one of the following algorithms: 'sha256', 'sha1', 'md5'`,
									Computed:            true,
									Default:             setdefault.StaticValue(types.SetNull(types.StringType)),

									ElementType: types.StringType,
								},
								"child_cipher_algo": schema.SetAttribute{
									MarkdownDescription: `This is the cipher algorithms to be used in Phase 2. The value should be an array with one or more of the following algorithms: 'aes256', 'aes192', 'aes128', 'tripledes', 'des', 'null'`,
									Computed:            true,
									Default:             setdefault.StaticValue(types.SetNull(types.StringType)),

									ElementType: types.StringType,
								},
								"child_lifetime": schema.Int64Attribute{
									MarkdownDescription: `The lifetime of the Phase 2 SA in seconds.`,
									Computed:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"child_pfs_group": schema.SetAttribute{
									MarkdownDescription: `This is the Diffie-Hellman group to be used for Perfect Forward Secrecy in Phase 2. The value should be an array with one of the following values: 'disabled','group14', 'group5', 'group2', 'group1'`,
									Computed:            true,
									Default:             setdefault.StaticValue(types.SetNull(types.StringType)),

									ElementType: types.StringType,
								},
								"ike_auth_algo": schema.SetAttribute{
									MarkdownDescription: `This is the authentication algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: 'sha256', 'sha1', 'md5'`,
									Computed:            true,
									Default:             setdefault.StaticValue(types.SetNull(types.StringType)),

									ElementType: types.StringType,
								},
								"ike_cipher_algo": schema.SetAttribute{
									MarkdownDescription: `This is the cipher algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: 'aes256', 'aes192', 'aes128', 'tripledes', 'des'`,
									Computed:            true,
									Default:             setdefault.StaticValue(types.SetNull(types.StringType)),

									ElementType: types.StringType,
								},
								"ike_diffie_hellman_group": schema.SetAttribute{
									MarkdownDescription: `This is the Diffie-Hellman group to be used in Phase 1. The value should be an array with one of the following algorithms: 'group14', 'group5', 'group2', 'group1'`,
									Computed:            true,
									Default:             setdefault.StaticValue(types.SetNull(types.StringType)),

									ElementType: types.StringType,
								},
								"ike_lifetime": schema.Int64Attribute{
									MarkdownDescription: `The lifetime of the Phase 1 SA in seconds.`,
									Computed:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"ike_prf_algo": schema.SetAttribute{
									MarkdownDescription: `[optional] This is the pseudo-random function to be used in IKE_SA. The value should be an array with one of the following algorithms: 'prfsha256', 'prfsha1', 'prfmd5', 'default'. The 'default' option can be used to default to the Authentication algorithm.`,
									Computed:            true,
									Default:             setdefault.StaticValue(types.SetNull(types.StringType)),

									ElementType: types.StringType,
								},
							},
						},
						"ipsec_policies_preset": schema.StringAttribute{
							MarkdownDescription: `One of the following available presets: 'default', 'aws', 'azure', 'umbrella', 'zscaler'. If this is provided, the 'ipsecPolicies' parameter is ignored.`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"local_id": schema.StringAttribute{
							MarkdownDescription: `[optional] The local ID is used to identify the MX to the peer. This will apply to all MXs this peer applies to.`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the VPN peer`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"network_tags": schema.SetAttribute{
							MarkdownDescription: `A list of network tags that will connect with this peer. Use ['all'] for all networks. Use ['none'] for no networks. If not included, the default is ['all'].`,
							Computed:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
						},
						"private_subnets": schema.SetAttribute{
							MarkdownDescription: `The list of the private subnets of the VPN peer`,
							Computed:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
						},
						"public_hostname": schema.StringAttribute{
							MarkdownDescription: `[optional] The public hostname of the VPN peer`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"public_ip": schema.StringAttribute{
							MarkdownDescription: `[optional] The public IP of the VPN peer`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"remote_id": schema.StringAttribute{
							MarkdownDescription: `[optional] The remote ID is used to identify the connecting VPN peer. This can either be a valid IPv4 Address, FQDN or User FQDN.`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: `The shared secret with the VPN peer`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
		},
	}
}

func (r *OrganizationsApplianceVpnThirdPartyVpnpeersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsApplianceVpnThirdPartyVpnpeersRs

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
	vvOrganizationID := data.OrganizationID.ValueString()
	//Has Item and not has items

	if vvOrganizationID != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Appliance.GetOrganizationApplianceVpnThirdPartyVpnpeers(vvOrganizationID)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource OrganizationsApplianceVpnThirdPartyVpnpeers  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource OrganizationsApplianceVpnThirdPartyVpnpeers only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateOrganizationApplianceVpnThirdPartyVpnpeers(vvOrganizationID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationApplianceVpnThirdPartyVpnpeers",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationApplianceVpnThirdPartyVpnpeers",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Appliance.GetOrganizationApplianceVpnThirdPartyVpnpeers(vvOrganizationID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationApplianceVpnThirdPartyVpnpeers",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationApplianceVpnThirdPartyVpnpeers",
			err.Error(),
		)
		return
	}

	data = ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *OrganizationsApplianceVpnThirdPartyVpnpeersResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsApplianceVpnThirdPartyVpnpeersRs

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

	vvOrganizationID := data.OrganizationID.ValueString()
	responseGet, restyRespGet, err := r.client.Appliance.GetOrganizationApplianceVpnThirdPartyVpnpeers(vvOrganizationID)
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
				"Failure when executing GetOrganizationApplianceVpnThirdPartyVpnpeers",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationApplianceVpnThirdPartyVpnpeers",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *OrganizationsApplianceVpnThirdPartyVpnpeersResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), req.ID)...)
}

func (r *OrganizationsApplianceVpnThirdPartyVpnpeersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsApplianceVpnThirdPartyVpnpeersRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateOrganizationApplianceVpnThirdPartyVpnpeers(vvOrganizationID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationApplianceVpnThirdPartyVpnpeers",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationApplianceVpnThirdPartyVpnpeers",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsApplianceVpnThirdPartyVpnpeersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting OrganizationsApplianceVpnThirdPartyVpnpeers", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsApplianceVpnThirdPartyVpnpeersRs struct {
	OrganizationID types.String                                                             `tfsdk:"organization_id"`
	Peers          *[]ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeersRs `tfsdk:"peers"`
	PeersResponse  *[]ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeersRs `tfsdk:"peers_response"`
}

type ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeersRs struct {
	IkeVersion          types.String                                                                        `tfsdk:"ike_version"`
	IPsecPolicies       *ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeersIpsecPoliciesRs `tfsdk:"ipsec_policies"`
	IPsecPoliciesPreset types.String                                                                        `tfsdk:"ipsec_policies_preset"`
	LocalID             types.String                                                                        `tfsdk:"local_id"`
	Name                types.String                                                                        `tfsdk:"name"`
	NetworkTags         types.Set                                                                           `tfsdk:"network_tags"`
	PrivateSubnets      types.Set                                                                           `tfsdk:"private_subnets"`
	PublicIP            types.String                                                                        `tfsdk:"public_ip"`
	RemoteID            types.String                                                                        `tfsdk:"remote_id"`
	Secret              types.String                                                                        `tfsdk:"secret"`
	PublicHostname      types.String                                                                        `tfsdk:"public_hostname"`
}

type ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeersIpsecPoliciesRs struct {
	ChildAuthAlgo         types.Set   `tfsdk:"child_auth_algo"`
	ChildCipherAlgo       types.Set   `tfsdk:"child_cipher_algo"`
	ChildLifetime         types.Int64 `tfsdk:"child_lifetime"`
	ChildPfsGroup         types.Set   `tfsdk:"child_pfs_group"`
	IkeAuthAlgo           types.Set   `tfsdk:"ike_auth_algo"`
	IkeCipherAlgo         types.Set   `tfsdk:"ike_cipher_algo"`
	IkeDiffieHellmanGroup types.Set   `tfsdk:"ike_diffie_hellman_group"`
	IkeLifetime           types.Int64 `tfsdk:"ike_lifetime"`
	IkePrfAlgo            types.Set   `tfsdk:"ike_prf_algo"`
}

// FromBody
func (r *OrganizationsApplianceVpnThirdPartyVpnpeersRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeers {
	var requestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeers []merakigosdk.RequestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeers

	if r.Peers != nil {
		for _, rItem1 := range *r.Peers {
			ikeVersion := rItem1.IkeVersion.ValueString()
			var requestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeersIPsecPolicies *merakigosdk.RequestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeersIPsecPolicies

			if rItem1.IPsecPolicies != nil {

				var childAuthAlgo []string = nil
				rItem1.IPsecPolicies.ChildAuthAlgo.ElementsAs(ctx, &childAuthAlgo, false)

				var childCipherAlgo []string = nil
				rItem1.IPsecPolicies.ChildCipherAlgo.ElementsAs(ctx, &childCipherAlgo, false)
				childLifetime := func() *int64 {
					if !rItem1.IPsecPolicies.ChildLifetime.IsUnknown() && !rItem1.IPsecPolicies.ChildLifetime.IsNull() {
						return rItem1.IPsecPolicies.ChildLifetime.ValueInt64Pointer()
					}
					return nil
				}()

				var childPfsGroup []string = nil
				rItem1.IPsecPolicies.ChildPfsGroup.ElementsAs(ctx, &childPfsGroup, false)

				var ikeAuthAlgo []string = nil
				rItem1.IPsecPolicies.IkeAuthAlgo.ElementsAs(ctx, &ikeAuthAlgo, false)

				var ikeCipherAlgo []string = nil
				rItem1.IPsecPolicies.IkeCipherAlgo.ElementsAs(ctx, &ikeCipherAlgo, false)

				var ikeDiffieHellmanGroup []string = nil
				rItem1.IPsecPolicies.IkeDiffieHellmanGroup.ElementsAs(ctx, &ikeDiffieHellmanGroup, false)
				ikeLifetime := func() *int64 {
					if !rItem1.IPsecPolicies.IkeLifetime.IsUnknown() && !rItem1.IPsecPolicies.IkeLifetime.IsNull() {
						return rItem1.IPsecPolicies.IkeLifetime.ValueInt64Pointer()
					}
					return nil
				}()

				var ikePrfAlgo []string = nil
				rItem1.IPsecPolicies.IkePrfAlgo.ElementsAs(ctx, &ikePrfAlgo, false)
				requestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeersIPsecPolicies = &merakigosdk.RequestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeersIPsecPolicies{
					ChildAuthAlgo:         childAuthAlgo,
					ChildCipherAlgo:       childCipherAlgo,
					ChildLifetime:         int64ToIntPointer(childLifetime),
					ChildPfsGroup:         childPfsGroup,
					IkeAuthAlgo:           ikeAuthAlgo,
					IkeCipherAlgo:         ikeCipherAlgo,
					IkeDiffieHellmanGroup: ikeDiffieHellmanGroup,
					IkeLifetime:           int64ToIntPointer(ikeLifetime),
					IkePrfAlgo:            ikePrfAlgo,
				}
				//[debug] Is Array: False
			}
			ipsecPoliciesPreset := rItem1.IPsecPoliciesPreset.ValueString()
			localID := rItem1.LocalID.ValueString()
			name := rItem1.Name.ValueString()

			var networkTags []string = nil
			rItem1.NetworkTags.ElementsAs(ctx, &networkTags, false)

			var privateSubnets []string = nil
			rItem1.PrivateSubnets.ElementsAs(ctx, &privateSubnets, false)
			publicHostname := rItem1.PublicHostname.ValueString()
			publicIP := rItem1.PublicIP.ValueString()
			remoteID := rItem1.RemoteID.ValueString()
			secret := rItem1.Secret.ValueString()
			requestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeers = append(requestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeers, merakigosdk.RequestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeers{
				IkeVersion:          ikeVersion,
				IPsecPolicies:       requestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeersIPsecPolicies,
				IPsecPoliciesPreset: ipsecPoliciesPreset,
				LocalID:             localID,
				Name:                name,
				NetworkTags:         networkTags,
				PrivateSubnets:      privateSubnets,
				PublicHostname:      publicHostname,
				PublicIP:            publicIP,
				RemoteID:            remoteID,
				Secret:              secret,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeers{
		Peers: func() *[]merakigosdk.RequestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeers {
			if len(requestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeers) > 0 {
				return &requestApplianceUpdateOrganizationApplianceVpnThirdPartyVpnpeersPeers
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersItemToBodyRs(state OrganizationsApplianceVpnThirdPartyVpnpeersRs, response *merakigosdk.ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeers, is_read bool) OrganizationsApplianceVpnThirdPartyVpnpeersRs {
	itemState := OrganizationsApplianceVpnThirdPartyVpnpeersRs{
		PeersResponse: func() *[]ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeersRs {
			if response.Peers != nil {
				result := make([]ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeersRs, len(*response.Peers))
				for i, peers := range *response.Peers {
					result[i] = ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeersRs{
						IkeVersion: types.StringValue(peers.IkeVersion),
						IPsecPolicies: func() *ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeersIpsecPoliciesRs {
							if peers.IPsecPolicies != nil {
								return &ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeersIpsecPoliciesRs{
									ChildAuthAlgo:   StringSliceToSet(peers.IPsecPolicies.ChildAuthAlgo),
									ChildCipherAlgo: StringSliceToSet(peers.IPsecPolicies.ChildCipherAlgo),
									ChildLifetime: func() types.Int64 {
										if peers.IPsecPolicies.ChildLifetime != nil {
											return types.Int64Value(int64(*peers.IPsecPolicies.ChildLifetime))
										}
										return types.Int64{}
									}(),
									ChildPfsGroup:         StringSliceToSet(peers.IPsecPolicies.ChildPfsGroup),
									IkeAuthAlgo:           StringSliceToSet(peers.IPsecPolicies.IkeAuthAlgo),
									IkeCipherAlgo:         StringSliceToSet(peers.IPsecPolicies.IkeCipherAlgo),
									IkeDiffieHellmanGroup: StringSliceToSet(peers.IPsecPolicies.IkeDiffieHellmanGroup),
									IkeLifetime: func() types.Int64 {
										if peers.IPsecPolicies.IkeLifetime != nil {
											return types.Int64Value(int64(*peers.IPsecPolicies.IkeLifetime))
										}
										return types.Int64{}
									}(),
									IkePrfAlgo: StringSliceToSet(peers.IPsecPolicies.IkePrfAlgo),
								}
							}
							return nil
						}(),
						IPsecPoliciesPreset: types.StringValue(peers.IPsecPoliciesPreset),
						LocalID:             types.StringValue(peers.LocalID),
						Name:                types.StringValue(peers.Name),
						NetworkTags:         StringSliceToSet(peers.NetworkTags),
						PrivateSubnets:      StringSliceToSet(peers.PrivateSubnets),
						PublicIP:            types.StringValue(peers.PublicIP),
						RemoteID:            types.StringValue(peers.RemoteID),
						Secret:              types.StringValue(peers.Secret),
					}
				}
				return &result
			}
			return nil
		}(),
		Peers: state.Peers,
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsApplianceVpnThirdPartyVpnpeersRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsApplianceVpnThirdPartyVpnpeersRs)
}
