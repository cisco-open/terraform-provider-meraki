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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsApplianceSecurityIntrusionDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsApplianceSecurityIntrusionDataSource{}
)

func NewOrganizationsApplianceSecurityIntrusionDataSource() datasource.DataSource {
	return &OrganizationsApplianceSecurityIntrusionDataSource{}
}

type OrganizationsApplianceSecurityIntrusionDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsApplianceSecurityIntrusionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsApplianceSecurityIntrusionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_appliance_security_intrusion"
}

func (d *OrganizationsApplianceSecurityIntrusionDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"allowed_rules": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"message": schema.StringAttribute{
									Computed: true,
								},
								"rule_id": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsApplianceSecurityIntrusionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsApplianceSecurityIntrusion OrganizationsApplianceSecurityIntrusion
	diags := req.Config.Get(ctx, &organizationsApplianceSecurityIntrusion)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationApplianceSecurityIntrusion")
		vvOrganizationID := organizationsApplianceSecurityIntrusion.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetOrganizationApplianceSecurityIntrusion(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationApplianceSecurityIntrusion",
				err.Error(),
			)
			return
		}

		organizationsApplianceSecurityIntrusion = ResponseApplianceGetOrganizationApplianceSecurityIntrusionItemToBody(organizationsApplianceSecurityIntrusion, response1)
		diags = resp.State.Set(ctx, &organizationsApplianceSecurityIntrusion)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsApplianceSecurityIntrusion struct {
	OrganizationID types.String                                                `tfsdk:"organization_id"`
	Item           *ResponseApplianceGetOrganizationApplianceSecurityIntrusion `tfsdk:"item"`
}

type ResponseApplianceGetOrganizationApplianceSecurityIntrusion struct {
	AllowedRules *[]ResponseApplianceGetOrganizationApplianceSecurityIntrusionAllowedRules `tfsdk:"allowed_rules"`
}

type ResponseApplianceGetOrganizationApplianceSecurityIntrusionAllowedRules struct {
	Message types.String `tfsdk:"message"`
	RuleID  types.String `tfsdk:"rule_id"`
}

// ToBody
func ResponseApplianceGetOrganizationApplianceSecurityIntrusionItemToBody(state OrganizationsApplianceSecurityIntrusion, response *merakigosdk.ResponseApplianceGetOrganizationApplianceSecurityIntrusion) OrganizationsApplianceSecurityIntrusion {
	itemState := ResponseApplianceGetOrganizationApplianceSecurityIntrusion{
		AllowedRules: func() *[]ResponseApplianceGetOrganizationApplianceSecurityIntrusionAllowedRules {
			if response.AllowedRules != nil {
				result := make([]ResponseApplianceGetOrganizationApplianceSecurityIntrusionAllowedRules, len(*response.AllowedRules))
				for i, allowedRules := range *response.AllowedRules {
					result[i] = ResponseApplianceGetOrganizationApplianceSecurityIntrusionAllowedRules{
						Message: func() types.String {
							if allowedRules.Message != "" {
								return types.StringValue(allowedRules.Message)
							}
							return types.String{}
						}(),
						RuleID: func() types.String {
							if allowedRules.RuleID != "" {
								return types.StringValue(allowedRules.RuleID)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
