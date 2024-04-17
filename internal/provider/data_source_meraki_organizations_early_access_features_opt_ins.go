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

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsEarlyAccessFeaturesOptInsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsEarlyAccessFeaturesOptInsDataSource{}
)

func NewOrganizationsEarlyAccessFeaturesOptInsDataSource() datasource.DataSource {
	return &OrganizationsEarlyAccessFeaturesOptInsDataSource{}
}

type OrganizationsEarlyAccessFeaturesOptInsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsEarlyAccessFeaturesOptInsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsEarlyAccessFeaturesOptInsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_early_access_features_opt_ins"
}

func (d *OrganizationsEarlyAccessFeaturesOptInsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"opt_in_id": schema.StringAttribute{
				MarkdownDescription: `optInId path parameter. Opt in ID`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"created_at": schema.StringAttribute{
						MarkdownDescription: `Time when Early Access Feature was created`,
						Computed:            true,
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `ID of Early Access Feature`,
						Computed:            true,
					},
					"limit_scope_to_networks": schema.SetNestedAttribute{
						MarkdownDescription: `Networks assigned to the Early Access Feature`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `ID of Network`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name of Network`,
									Computed:            true,
								},
							},
						},
					},
					"short_name": schema.StringAttribute{
						MarkdownDescription: `Name of Early Access Feature`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *OrganizationsEarlyAccessFeaturesOptInsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsEarlyAccessFeaturesOptIns OrganizationsEarlyAccessFeaturesOptIns
	diags := req.Config.Get(ctx, &organizationsEarlyAccessFeaturesOptIns)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsEarlyAccessFeaturesOptIns.OrganizationID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsEarlyAccessFeaturesOptIns.OrganizationID.IsNull(), !organizationsEarlyAccessFeaturesOptIns.OptInID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationEarlyAccessFeaturesOptIns")
		vvOrganizationID := organizationsEarlyAccessFeaturesOptIns.OrganizationID.ValueString()

		response1, restyResp1, err := d.client.Organizations.GetOrganizationEarlyAccessFeaturesOptIns(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationEarlyAccessFeaturesOptIns",
				err.Error(),
			)
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationEarlyAccessFeaturesOptIn")
		vvOrganizationID := organizationsEarlyAccessFeaturesOptIns.OrganizationID.ValueString()
		vvOptInID := organizationsEarlyAccessFeaturesOptIns.OptInID.ValueString()

		response2, restyResp2, err := d.client.Organizations.GetOrganizationEarlyAccessFeaturesOptIn(vvOrganizationID, vvOptInID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationEarlyAccessFeaturesOptIn",
				err.Error(),
			)
			return
		}

		organizationsEarlyAccessFeaturesOptIns = ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInItemToBody(organizationsEarlyAccessFeaturesOptIns, response2)
		diags = resp.State.Set(ctx, &organizationsEarlyAccessFeaturesOptIns)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsEarlyAccessFeaturesOptIns struct {
	OrganizationID types.String                                                  `tfsdk:"organization_id"`
	OptInID        types.String                                                  `tfsdk:"opt_in_id"`
	Item           *ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptIn `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptIn struct {
	CreatedAt            types.String                                                                        `tfsdk:"created_at"`
	ID                   types.String                                                                        `tfsdk:"id"`
	LimitScopeToNetworks *[]ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInLimitScopeToNetworks `tfsdk:"limit_scope_to_networks"`
	ShortName            types.String                                                                        `tfsdk:"short_name"`
}

type ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInLimitScopeToNetworks struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// ToBody
func ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInItemToBody(state OrganizationsEarlyAccessFeaturesOptIns, response *merakigosdk.ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptIn) OrganizationsEarlyAccessFeaturesOptIns {
	itemState := ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptIn{
		CreatedAt: types.StringValue(response.CreatedAt),
		ID:        types.StringValue(response.ID),
		LimitScopeToNetworks: func() *[]ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInLimitScopeToNetworks {
			if response.LimitScopeToNetworks != nil {
				result := make([]ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInLimitScopeToNetworks, len(*response.LimitScopeToNetworks))
				for i, limitScopeToNetworks := range *response.LimitScopeToNetworks {
					result[i] = ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInLimitScopeToNetworks{
						ID:   types.StringValue(limitScopeToNetworks.ID),
						Name: types.StringValue(limitScopeToNetworks.Name),
					}
				}
				return &result
			}
			return &[]ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInLimitScopeToNetworks{}
		}(),
		ShortName: types.StringValue(response.ShortName),
	}
	state.Item = &itemState
	return state
}
