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
	_ datasource.DataSource              = &OrganizationsApplianceUplinksStatusesOverviewDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsApplianceUplinksStatusesOverviewDataSource{}
)

func NewOrganizationsApplianceUplinksStatusesOverviewDataSource() datasource.DataSource {
	return &OrganizationsApplianceUplinksStatusesOverviewDataSource{}
}

type OrganizationsApplianceUplinksStatusesOverviewDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsApplianceUplinksStatusesOverviewDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsApplianceUplinksStatusesOverviewDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_appliance_uplinks_statuses_overview"
}

func (d *OrganizationsApplianceUplinksStatusesOverviewDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"counts": schema.SingleNestedAttribute{
						MarkdownDescription: `counts`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"by_status": schema.SingleNestedAttribute{
								MarkdownDescription: `byStatus`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"active": schema.Int64Attribute{
										MarkdownDescription: `number of uplinks that are active and working`,
										Computed:            true,
									},
									"connecting": schema.Int64Attribute{
										MarkdownDescription: `number of uplinks currently connecting`,
										Computed:            true,
									},
									"failed": schema.Int64Attribute{
										MarkdownDescription: `number of uplinks that were working but have failed`,
										Computed:            true,
									},
									"not_connected": schema.Int64Attribute{
										MarkdownDescription: `number of uplinks currently where nothing is plugged in`,
										Computed:            true,
									},
									"ready": schema.Int64Attribute{
										MarkdownDescription: `number of uplinks that are working but on standby`,
										Computed:            true,
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

func (d *OrganizationsApplianceUplinksStatusesOverviewDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsApplianceUplinksStatusesOverview OrganizationsApplianceUplinksStatusesOverview
	diags := req.Config.Get(ctx, &organizationsApplianceUplinksStatusesOverview)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationApplianceUplinksStatusesOverview")
		vvOrganizationID := organizationsApplianceUplinksStatusesOverview.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetOrganizationApplianceUplinksStatusesOverview(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationApplianceUplinksStatusesOverview",
				err.Error(),
			)
			return
		}

		organizationsApplianceUplinksStatusesOverview = ResponseApplianceGetOrganizationApplianceUplinksStatusesOverviewItemToBody(organizationsApplianceUplinksStatusesOverview, response1)
		diags = resp.State.Set(ctx, &organizationsApplianceUplinksStatusesOverview)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsApplianceUplinksStatusesOverview struct {
	OrganizationID types.String                                                      `tfsdk:"organization_id"`
	Item           *ResponseApplianceGetOrganizationApplianceUplinksStatusesOverview `tfsdk:"item"`
}

type ResponseApplianceGetOrganizationApplianceUplinksStatusesOverview struct {
	Counts *ResponseApplianceGetOrganizationApplianceUplinksStatusesOverviewCounts `tfsdk:"counts"`
}

type ResponseApplianceGetOrganizationApplianceUplinksStatusesOverviewCounts struct {
	ByStatus *ResponseApplianceGetOrganizationApplianceUplinksStatusesOverviewCountsByStatus `tfsdk:"by_status"`
}

type ResponseApplianceGetOrganizationApplianceUplinksStatusesOverviewCountsByStatus struct {
	Active       types.Int64 `tfsdk:"active"`
	Connecting   types.Int64 `tfsdk:"connecting"`
	Failed       types.Int64 `tfsdk:"failed"`
	NotConnected types.Int64 `tfsdk:"not_connected"`
	Ready        types.Int64 `tfsdk:"ready"`
}

// ToBody
func ResponseApplianceGetOrganizationApplianceUplinksStatusesOverviewItemToBody(state OrganizationsApplianceUplinksStatusesOverview, response *merakigosdk.ResponseApplianceGetOrganizationApplianceUplinksStatusesOverview) OrganizationsApplianceUplinksStatusesOverview {
	itemState := ResponseApplianceGetOrganizationApplianceUplinksStatusesOverview{
		Counts: func() *ResponseApplianceGetOrganizationApplianceUplinksStatusesOverviewCounts {
			if response.Counts != nil {
				return &ResponseApplianceGetOrganizationApplianceUplinksStatusesOverviewCounts{
					ByStatus: func() *ResponseApplianceGetOrganizationApplianceUplinksStatusesOverviewCountsByStatus {
						if response.Counts.ByStatus != nil {
							return &ResponseApplianceGetOrganizationApplianceUplinksStatusesOverviewCountsByStatus{
								Active: func() types.Int64 {
									if response.Counts.ByStatus.Active != nil {
										return types.Int64Value(int64(*response.Counts.ByStatus.Active))
									}
									return types.Int64{}
								}(),
								Connecting: func() types.Int64 {
									if response.Counts.ByStatus.Connecting != nil {
										return types.Int64Value(int64(*response.Counts.ByStatus.Connecting))
									}
									return types.Int64{}
								}(),
								Failed: func() types.Int64 {
									if response.Counts.ByStatus.Failed != nil {
										return types.Int64Value(int64(*response.Counts.ByStatus.Failed))
									}
									return types.Int64{}
								}(),
								NotConnected: func() types.Int64 {
									if response.Counts.ByStatus.NotConnected != nil {
										return types.Int64Value(int64(*response.Counts.ByStatus.NotConnected))
									}
									return types.Int64{}
								}(),
								Ready: func() types.Int64 {
									if response.Counts.ByStatus.Ready != nil {
										return types.Int64Value(int64(*response.Counts.ByStatus.Ready))
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
	state.Item = &itemState
	return state
}
