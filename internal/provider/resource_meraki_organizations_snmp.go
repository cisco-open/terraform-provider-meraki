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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsSNMPResource{}
	_ resource.ResourceWithConfigure = &OrganizationsSNMPResource{}
)

func NewOrganizationsSNMPResource() resource.Resource {
	return &OrganizationsSNMPResource{}
}

type OrganizationsSNMPResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsSNMPResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsSNMPResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_snmp"
}

func (r *OrganizationsSNMPResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"hostname": schema.StringAttribute{
				MarkdownDescription: `The hostname of the SNMP server.`,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"peer_ips": schema.SetAttribute{
				MarkdownDescription: `The list of IPv4 addresses that are allowed to access the SNMP server.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				Default:     setdefault.StaticValue(types.SetNull(types.StringType)),
				ElementType: types.StringType,
			},
			"port": schema.Int64Attribute{
				MarkdownDescription: `The port of the SNMP server.`,
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"v2_community_string": schema.StringAttribute{
				MarkdownDescription: `The community string for SNMP version 2c, if enabled.`,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"v2c_enabled": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating whether SNMP version 2c is enabled for the organization.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"v3_auth_mode": schema.StringAttribute{
				MarkdownDescription: `The SNMP version 3 authentication mode. Can be either 'MD5' or 'SHA'.
                                  Allowed values: [MD5,SHA]`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"MD5",
						"SHA",
					),
				},
			},
			"v3_auth_pass": schema.StringAttribute{
				MarkdownDescription: `The SNMP version 3 authentication password. Must be at least 8 characters if specified.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"v3_enabled": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating whether SNMP version 3 is enabled for the organization.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"v3_priv_mode": schema.StringAttribute{
				MarkdownDescription: `The SNMP version 3 privacy mode. Can be either 'DES' or 'AES128'.
                                  Allowed values: [AES128,DES]`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"AES128",
						"DES",
					),
				},
			},
			"v3_priv_pass": schema.StringAttribute{
				MarkdownDescription: `The SNMP version 3 privacy password. Must be at least 8 characters if specified.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"v3_user": schema.StringAttribute{
				MarkdownDescription: `The user for SNMP version 3, if enabled.`,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *OrganizationsSNMPResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsSNMPRs

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
	//Item
	responseVerifyItem, restyResp1, err := r.client.Organizations.GetOrganizationSNMP(vvOrganizationID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource OrganizationsSNMP only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource OrganizationsSNMP only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Organizations.UpdateOrganizationSNMP(vvOrganizationID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationSNMP",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationSNMP",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationSNMP(vvOrganizationID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSNMP",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationSNMP",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseOrganizationsGetOrganizationSNMPItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsSNMPResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsSNMPRs

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
	responseGet, restyRespGet, err := r.client.Organizations.GetOrganizationSNMP(vvOrganizationID)
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
				"Failure when executing GetOrganizationSNMP",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationSNMP",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseOrganizationsGetOrganizationSNMPItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *OrganizationsSNMPResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), req.ID)...)
}

func (r *OrganizationsSNMPResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsSNMPRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Organizations.UpdateOrganizationSNMP(vvOrganizationID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationSNMP",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationSNMP",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsSNMPResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting OrganizationsSNMP", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsSNMPRs struct {
	OrganizationID    types.String `tfsdk:"organization_id"`
	Hostname          types.String `tfsdk:"hostname"`
	PeerIPs           types.Set    `tfsdk:"peer_ips"`
	Port              types.Int64  `tfsdk:"port"`
	V2CommunityString types.String `tfsdk:"v2_community_string"`
	V2CEnabled        types.Bool   `tfsdk:"v2c_enabled"`
	V3AuthMode        types.String `tfsdk:"v3_auth_mode"`
	V3Enabled         types.Bool   `tfsdk:"v3_enabled"`
	V3PrivMode        types.String `tfsdk:"v3_priv_mode"`
	V3User            types.String `tfsdk:"v3_user"`
	V3AuthPass        types.String `tfsdk:"v3_auth_pass"`
	V3PrivPass        types.String `tfsdk:"v3_priv_pass"`
}

// FromBody
func (r *OrganizationsSNMPRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestOrganizationsUpdateOrganizationSNMP {
	emptyString := ""
	var peerIPs []string = nil
	r.PeerIPs.ElementsAs(ctx, &peerIPs, false)
	v2CEnabled := new(bool)
	if !r.V2CEnabled.IsUnknown() && !r.V2CEnabled.IsNull() {
		*v2CEnabled = r.V2CEnabled.ValueBool()
	} else {
		v2CEnabled = nil
	}
	v3AuthMode := new(string)
	if !r.V3AuthMode.IsUnknown() && !r.V3AuthMode.IsNull() {
		*v3AuthMode = r.V3AuthMode.ValueString()
	} else {
		v3AuthMode = &emptyString
	}
	v3AuthPass := new(string)
	if !r.V3AuthPass.IsUnknown() && !r.V3AuthPass.IsNull() {
		*v3AuthPass = r.V3AuthPass.ValueString()
	} else {
		v3AuthPass = &emptyString
	}
	v3Enabled := new(bool)
	if !r.V3Enabled.IsUnknown() && !r.V3Enabled.IsNull() {
		*v3Enabled = r.V3Enabled.ValueBool()
	} else {
		v3Enabled = nil
	}
	v3PrivMode := new(string)
	if !r.V3PrivMode.IsUnknown() && !r.V3PrivMode.IsNull() {
		*v3PrivMode = r.V3PrivMode.ValueString()
	} else {
		v3PrivMode = &emptyString
	}
	v3PrivPass := new(string)
	if !r.V3PrivPass.IsUnknown() && !r.V3PrivPass.IsNull() {
		*v3PrivPass = r.V3PrivPass.ValueString()
	} else {
		v3PrivPass = &emptyString
	}
	out := merakigosdk.RequestOrganizationsUpdateOrganizationSNMP{
		PeerIPs:    peerIPs,
		V2CEnabled: v2CEnabled,
		V3AuthMode: *v3AuthMode,
		V3AuthPass: *v3AuthPass,
		V3Enabled:  v3Enabled,
		V3PrivMode: *v3PrivMode,
		V3PrivPass: *v3PrivPass,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseOrganizationsGetOrganizationSNMPItemToBodyRs(state OrganizationsSNMPRs, response *merakigosdk.ResponseOrganizationsGetOrganizationSNMP, is_read bool) OrganizationsSNMPRs {
	itemState := OrganizationsSNMPRs{
		Hostname: types.StringValue(response.Hostname),
		PeerIPs:  state.PeerIPs,
		Port: func() types.Int64 {
			if response.Port != nil {
				return types.Int64Value(int64(*response.Port))
			}
			return types.Int64{}
		}(),
		V2CommunityString: types.StringValue(response.V2CommunityString),
		V2CEnabled: func() types.Bool {
			if response.V2CEnabled != nil {
				return types.BoolValue(*response.V2CEnabled)
			}
			return types.Bool{}
		}(),
		V3AuthMode: types.StringValue(response.V3AuthMode),
		V3Enabled: func() types.Bool {
			if response.V3Enabled != nil {
				return types.BoolValue(*response.V3Enabled)
			}
			return types.Bool{}
		}(),
		V3PrivMode: types.StringValue(response.V3PrivMode),
		V3User:     types.StringValue(response.V3User),
		V3AuthPass: state.V3AuthPass,
		V3PrivPass: state.V3PrivPass,
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsSNMPRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsSNMPRs)
}
