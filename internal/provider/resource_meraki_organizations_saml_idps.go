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
	"fmt"
	"strings"

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsSamlIDpsResource{}
	_ resource.ResourceWithConfigure = &OrganizationsSamlIDpsResource{}
)

func NewOrganizationsSamlIDpsResource() resource.Resource {
	return &OrganizationsSamlIDpsResource{}
}

type OrganizationsSamlIDpsResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsSamlIDpsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsSamlIDpsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_saml_idps"
}

func (r *OrganizationsSamlIDpsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"consumer_url": schema.StringAttribute{
				MarkdownDescription: `URL that is consuming SAML Identity Provider (IdP)`,
				Computed:            true,
			},
			"idp_id": schema.StringAttribute{
				MarkdownDescription: `ID associated with the SAML Identity Provider (IdP)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"slo_logout_url": schema.StringAttribute{
				MarkdownDescription: `Dashboard will redirect users to this URL when they sign out.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"x509cert_sha1_fingerprint": schema.StringAttribute{
				MarkdownDescription: `Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

//path params to set ['idpId']

func (r *OrganizationsSamlIDpsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsSamlIDpsRs

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
	//Reviw This  Has Item and item
	//HAS CREATE

	vvIDpID := data.IDpID.ValueString()
	if vvIDpID != "" {
		responseVerifyItem, restyRespGet, err := r.client.Organizations.GetOrganizationSamlIDp(vvOrganizationID, vvIDpID)
		if err != nil || responseVerifyItem == nil {
			if restyRespGet != nil {
				if restyRespGet.StatusCode() != 404 {

					resp.Diagnostics.AddError(
						"Failure when executing GetOrganizationSamlIDp",
						err.Error(),
					)
					return
				}
			}
		}

		if responseVerifyItem != nil {
			data := ResponseOrganizationsGetOrganizationSamlIDpItemToBodyRs(data, responseVerifyItem, false)
			diags := resp.State.Set(ctx, &data)
			resp.Diagnostics.Append(diags...)
			return
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Organizations.CreateOrganizationSamlIDp(vvOrganizationID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationSamlIDp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationSamlIDp",
			err.Error(),
		)
		return
	}
	//Items
	res := *response
	vvIDpID = res.IDpID
	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationSamlIDp(vvOrganizationID, vvIDpID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSamlIDps",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationSamlIDps",
			err.Error(),
		)
		return
	}
	data = ResponseOrganizationsGetOrganizationSamlIDpItemToBodyRs(data, responseGet, false)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
	return

}

func (r *OrganizationsSamlIDpsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsSamlIDpsRs

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
	vvIDpID := data.IDpID.ValueString()
	responseGet, restyRespGet, err := r.client.Organizations.GetOrganizationSamlIDp(vvOrganizationID, vvIDpID)
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
				"Failure when executing GetOrganizationSamlIDp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationSamlIDp",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseOrganizationsGetOrganizationSamlIDpItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsSamlIDpsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("idp_id"), idParts[1])...)
}

func (r *OrganizationsSamlIDpsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsSamlIDpsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	vvIDpID := data.IDpID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	_, restyResp2, err := r.client.Organizations.UpdateOrganizationSamlIDp(vvOrganizationID, vvIDpID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationSamlIDp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationSamlIDp",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsSamlIDpsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsSamlIDpsRs
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &state, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}

	vvOrganizationID := state.OrganizationID.ValueString()
	vvIDpID := state.IDpID.ValueString()
	_, err := r.client.Organizations.DeleteOrganizationSamlIDp(vvOrganizationID, vvIDpID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationSamlIDp", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsSamlIDpsRs struct {
	OrganizationID          types.String `tfsdk:"organization_id"`
	IDpID                   types.String `tfsdk:"idp_id"`
	ConsumerURL             types.String `tfsdk:"consumer_url"`
	SloLogoutURL            types.String `tfsdk:"slo_logout_url"`
	X509CertSha1Fingerprint types.String `tfsdk:"x509cert_sha1_fingerprint"`
}

// FromBody
func (r *OrganizationsSamlIDpsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationSamlIDp {
	emptyString := ""
	sloLogoutURL := new(string)
	if !r.SloLogoutURL.IsUnknown() && !r.SloLogoutURL.IsNull() {
		*sloLogoutURL = r.SloLogoutURL.ValueString()
	} else {
		sloLogoutURL = &emptyString
	}
	x509CertSha1Fingerprint := new(string)
	if !r.X509CertSha1Fingerprint.IsUnknown() && !r.X509CertSha1Fingerprint.IsNull() {
		*x509CertSha1Fingerprint = r.X509CertSha1Fingerprint.ValueString()
	} else {
		x509CertSha1Fingerprint = &emptyString
	}
	out := merakigosdk.RequestOrganizationsCreateOrganizationSamlIDp{
		SloLogoutURL:            *sloLogoutURL,
		X509CertSha1Fingerprint: *x509CertSha1Fingerprint,
	}
	return &out
}
func (r *OrganizationsSamlIDpsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestOrganizationsUpdateOrganizationSamlIDp {
	emptyString := ""
	sloLogoutURL := new(string)
	if !r.SloLogoutURL.IsUnknown() && !r.SloLogoutURL.IsNull() {
		*sloLogoutURL = r.SloLogoutURL.ValueString()
	} else {
		sloLogoutURL = &emptyString
	}
	x509CertSha1Fingerprint := new(string)
	if !r.X509CertSha1Fingerprint.IsUnknown() && !r.X509CertSha1Fingerprint.IsNull() {
		*x509CertSha1Fingerprint = r.X509CertSha1Fingerprint.ValueString()
	} else {
		x509CertSha1Fingerprint = &emptyString
	}
	out := merakigosdk.RequestOrganizationsUpdateOrganizationSamlIDp{
		SloLogoutURL:            *sloLogoutURL,
		X509CertSha1Fingerprint: *x509CertSha1Fingerprint,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseOrganizationsGetOrganizationSamlIDpItemToBodyRs(state OrganizationsSamlIDpsRs, response *merakigosdk.ResponseOrganizationsGetOrganizationSamlIDp, is_read bool) OrganizationsSamlIDpsRs {
	itemState := OrganizationsSamlIDpsRs{
		ConsumerURL:             types.StringValue(response.ConsumerURL),
		IDpID:                   types.StringValue(response.IDpID),
		SloLogoutURL:            types.StringValue(response.SloLogoutURL),
		X509CertSha1Fingerprint: types.StringValue(response.X509CertSha1Fingerprint),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsSamlIDpsRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsSamlIDpsRs)
}
