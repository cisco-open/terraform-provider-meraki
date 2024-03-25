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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsLicensesOverviewDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsLicensesOverviewDataSource{}
)

func NewOrganizationsLicensesOverviewDataSource() datasource.DataSource {
	return &OrganizationsLicensesOverviewDataSource{}
}

type OrganizationsLicensesOverviewDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsLicensesOverviewDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsLicensesOverviewDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_licenses_overview"
}

func (d *OrganizationsLicensesOverviewDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"expiration_date": schema.StringAttribute{
						Computed: true,
					},
					"licensed_device_counts": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"ms": schema.Int64Attribute{
								Computed: true,
							},
						},
					},
					"status": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *OrganizationsLicensesOverviewDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsLicensesOverview OrganizationsLicensesOverview
	diags := req.Config.Get(ctx, &organizationsLicensesOverview)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationLicensesOverview")
		vvOrganizationID := organizationsLicensesOverview.OrganizationID.ValueString()

		response1, restyResp1, err := d.client.Organizations.GetOrganizationLicensesOverview(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationLicensesOverview",
				err.Error(),
			)
			return
		}

		organizationsLicensesOverview = ResponseOrganizationsGetOrganizationLicensesOverviewItemToBody(organizationsLicensesOverview, response1)
		diags = resp.State.Set(ctx, &organizationsLicensesOverview)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsLicensesOverview struct {
	OrganizationID types.String                                          `tfsdk:"organization_id"`
	Item           *ResponseOrganizationsGetOrganizationLicensesOverview `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationLicensesOverview struct {
	ExpirationDate       types.String                                                              `tfsdk:"expiration_date"`
	LicensedDeviceCounts *ResponseOrganizationsGetOrganizationLicensesOverviewLicensedDeviceCounts `tfsdk:"licensed_device_counts"`
	Status               types.String                                                              `tfsdk:"status"`
}

type ResponseOrganizationsGetOrganizationLicensesOverviewLicensedDeviceCounts struct {
	MS types.Int64 `tfsdk:"ms"`
}

// ToBody
func ResponseOrganizationsGetOrganizationLicensesOverviewItemToBody(state OrganizationsLicensesOverview, response *merakigosdk.ResponseOrganizationsGetOrganizationLicensesOverview) OrganizationsLicensesOverview {
	itemState := ResponseOrganizationsGetOrganizationLicensesOverview{
		ExpirationDate: types.StringValue(response.ExpirationDate),
		LicensedDeviceCounts: func() *ResponseOrganizationsGetOrganizationLicensesOverviewLicensedDeviceCounts {
			if response.LicensedDeviceCounts != nil {
				return &ResponseOrganizationsGetOrganizationLicensesOverviewLicensedDeviceCounts{
					MS: func() types.Int64 {
						if response.LicensedDeviceCounts.MS != nil {
							return types.Int64Value(int64(*response.LicensedDeviceCounts.MS))
						}
						return types.Int64{}
					}(),
				}
			}
			return &ResponseOrganizationsGetOrganizationLicensesOverviewLicensedDeviceCounts{}
		}(),
		Status: types.StringValue(response.Status),
	}
	state.Item = &itemState
	return state
}
