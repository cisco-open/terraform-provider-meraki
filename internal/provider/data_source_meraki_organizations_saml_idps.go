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

// DATA SOURCE NORMAL
import (
	"context"
	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsSamlIDpsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSamlIDpsDataSource{}
)

func NewOrganizationsSamlIDpsDataSource() datasource.DataSource {
	return &OrganizationsSamlIDpsDataSource{}
}

type OrganizationsSamlIDpsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSamlIDpsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSamlIDpsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_saml_idps"
}

func (d *OrganizationsSamlIDpsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"idp_id": schema.StringAttribute{
				MarkdownDescription: `idpId path parameter. Idp ID`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"consumer_url": schema.StringAttribute{
						MarkdownDescription: `URL that is consuming SAML Identity Provider (IdP)`,
						Computed:            true,
					},
					"idp_id": schema.StringAttribute{
						MarkdownDescription: `ID associated with the SAML Identity Provider (IdP)`,
						Computed:            true,
					},
					"slo_logout_url": schema.StringAttribute{
						MarkdownDescription: `Dashboard will redirect users to this URL when they sign out.`,
						Computed:            true,
					},
					"x509cert_sha1_fingerprint": schema.StringAttribute{
						MarkdownDescription: `Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationSamlIdps`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"consumer_url": schema.StringAttribute{
							MarkdownDescription: `URL that is consuming SAML Identity Provider (IdP)`,
							Computed:            true,
						},
						"idp_id": schema.StringAttribute{
							MarkdownDescription: `ID associated with the SAML Identity Provider (IdP)`,
							Computed:            true,
						},
						"slo_logout_url": schema.StringAttribute{
							MarkdownDescription: `Dashboard will redirect users to this URL when they sign out.`,
							Computed:            true,
						},
						"x509cert_sha1_fingerprint": schema.StringAttribute{
							MarkdownDescription: `Fingerprint (SHA1) of the SAML certificate provided by your Identity Provider (IdP). This will be used for encryption / validation.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsSamlIDpsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSamlIDps OrganizationsSamlIDps
	diags := req.Config.Get(ctx, &organizationsSamlIDps)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsSamlIDps.OrganizationID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsSamlIDps.OrganizationID.IsNull(), !organizationsSamlIDps.IDpID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSamlIDps")
		vvOrganizationID := organizationsSamlIDps.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationSamlIDps(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSamlIDps",
				err.Error(),
			)
			return
		}

		organizationsSamlIDps = ResponseOrganizationsGetOrganizationSamlIDpsItemsToBody(organizationsSamlIDps, response1)
		diags = resp.State.Set(ctx, &organizationsSamlIDps)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSamlIDp")
		vvOrganizationID := organizationsSamlIDps.OrganizationID.ValueString()
		vvIDpID := organizationsSamlIDps.IDpID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Organizations.GetOrganizationSamlIDp(vvOrganizationID, vvIDpID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSamlIDp",
				err.Error(),
			)
			return
		}

		organizationsSamlIDps = ResponseOrganizationsGetOrganizationSamlIDpItemToBody(organizationsSamlIDps, response2)
		diags = resp.State.Set(ctx, &organizationsSamlIDps)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSamlIDps struct {
	OrganizationID types.String                                        `tfsdk:"organization_id"`
	IDpID          types.String                                        `tfsdk:"idp_id"`
	Items          *[]ResponseItemOrganizationsGetOrganizationSamlIdps `tfsdk:"items"`
	Item           *ResponseOrganizationsGetOrganizationSamlIdp        `tfsdk:"item"`
}

type ResponseItemOrganizationsGetOrganizationSamlIdps struct {
	ConsumerURL             types.String `tfsdk:"consumer_url"`
	IDpID                   types.String `tfsdk:"idp_id"`
	SloLogoutURL            types.String `tfsdk:"slo_logout_url"`
	X509CertSha1Fingerprint types.String `tfsdk:"x509cert_sha1_fingerprint"`
}

type ResponseOrganizationsGetOrganizationSamlIdp struct {
	ConsumerURL             types.String `tfsdk:"consumer_url"`
	IDpID                   types.String `tfsdk:"idp_id"`
	SloLogoutURL            types.String `tfsdk:"slo_logout_url"`
	X509CertSha1Fingerprint types.String `tfsdk:"x509cert_sha1_fingerprint"`
}

// ToBody
func ResponseOrganizationsGetOrganizationSamlIDpsItemsToBody(state OrganizationsSamlIDps, response *merakigosdk.ResponseOrganizationsGetOrganizationSamlIDps) OrganizationsSamlIDps {
	var items []ResponseItemOrganizationsGetOrganizationSamlIdps
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationSamlIdps{
			ConsumerURL:             types.StringValue(item.ConsumerURL),
			IDpID:                   types.StringValue(item.IDpID),
			SloLogoutURL:            types.StringValue(item.SloLogoutURL),
			X509CertSha1Fingerprint: types.StringValue(item.X509CertSha1Fingerprint),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseOrganizationsGetOrganizationSamlIDpItemToBody(state OrganizationsSamlIDps, response *merakigosdk.ResponseOrganizationsGetOrganizationSamlIDp) OrganizationsSamlIDps {
	itemState := ResponseOrganizationsGetOrganizationSamlIdp{
		ConsumerURL:             types.StringValue(response.ConsumerURL),
		IDpID:                   types.StringValue(response.IDpID),
		SloLogoutURL:            types.StringValue(response.SloLogoutURL),
		X509CertSha1Fingerprint: types.StringValue(response.X509CertSha1Fingerprint),
	}
	state.Item = &itemState
	return state
}
