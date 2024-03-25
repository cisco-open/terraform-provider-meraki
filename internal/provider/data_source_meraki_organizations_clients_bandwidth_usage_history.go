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
	_ datasource.DataSource              = &OrganizationsClientsBandwidthUsageHistoryDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsClientsBandwidthUsageHistoryDataSource{}
)

func NewOrganizationsClientsBandwidthUsageHistoryDataSource() datasource.DataSource {
	return &OrganizationsClientsBandwidthUsageHistoryDataSource{}
}

type OrganizationsClientsBandwidthUsageHistoryDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsClientsBandwidthUsageHistoryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsClientsBandwidthUsageHistoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_clients_bandwidth_usage_history"
}

func (d *OrganizationsClientsBandwidthUsageHistoryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 31 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationClientsBandwidthUsageHistory`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"downstream": schema.Int64Attribute{
							MarkdownDescription: `Downloaded data, in mbps.`,
							Computed:            true,
						},
						"total": schema.Int64Attribute{
							MarkdownDescription: `Total bandwidth usage, in mbps.`,
							Computed:            true,
						},
						"ts": schema.StringAttribute{
							MarkdownDescription: `Timestamp for the bandwidth usage snapshot.`,
							Computed:            true,
						},
						"upstream": schema.Int64Attribute{
							MarkdownDescription: `Uploaded data, in mbps.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsClientsBandwidthUsageHistoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsClientsBandwidthUsageHistory OrganizationsClientsBandwidthUsageHistory
	diags := req.Config.Get(ctx, &organizationsClientsBandwidthUsageHistory)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationClientsBandwidthUsageHistory")
		vvOrganizationID := organizationsClientsBandwidthUsageHistory.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationClientsBandwidthUsageHistoryQueryParams{}

		queryParams1.T0 = organizationsClientsBandwidthUsageHistory.T0.ValueString()
		queryParams1.T1 = organizationsClientsBandwidthUsageHistory.T1.ValueString()
		queryParams1.Timespan = organizationsClientsBandwidthUsageHistory.Timespan.ValueFloat64()

		response1, restyResp1, err := d.client.Organizations.GetOrganizationClientsBandwidthUsageHistory(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationClientsBandwidthUsageHistory",
				err.Error(),
			)
			return
		}

		organizationsClientsBandwidthUsageHistory = ResponseOrganizationsGetOrganizationClientsBandwidthUsageHistoryItemsToBody(organizationsClientsBandwidthUsageHistory, response1)
		diags = resp.State.Set(ctx, &organizationsClientsBandwidthUsageHistory)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsClientsBandwidthUsageHistory struct {
	OrganizationID types.String                                                            `tfsdk:"organization_id"`
	T0             types.String                                                            `tfsdk:"t0"`
	T1             types.String                                                            `tfsdk:"t1"`
	Timespan       types.Float64                                                           `tfsdk:"timespan"`
	Items          *[]ResponseItemOrganizationsGetOrganizationClientsBandwidthUsageHistory `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationClientsBandwidthUsageHistory struct {
	Downstream types.Int64  `tfsdk:"downstream"`
	Total      types.Int64  `tfsdk:"total"`
	Ts         types.String `tfsdk:"ts"`
	Upstream   types.Int64  `tfsdk:"upstream"`
}

// ToBody
func ResponseOrganizationsGetOrganizationClientsBandwidthUsageHistoryItemsToBody(state OrganizationsClientsBandwidthUsageHistory, response *merakigosdk.ResponseOrganizationsGetOrganizationClientsBandwidthUsageHistory) OrganizationsClientsBandwidthUsageHistory {
	var items []ResponseItemOrganizationsGetOrganizationClientsBandwidthUsageHistory
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationClientsBandwidthUsageHistory{
			Downstream: func() types.Int64 {
				if item.Downstream != nil {
					return types.Int64Value(int64(*item.Downstream))
				}
				return types.Int64{}
			}(),
			Total: func() types.Int64 {
				if item.Total != nil {
					return types.Int64Value(int64(*item.Total))
				}
				return types.Int64{}
			}(),
			Ts: types.StringValue(item.Ts),
			Upstream: func() types.Int64 {
				if item.Upstream != nil {
					return types.Int64Value(int64(*item.Upstream))
				}
				return types.Int64{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
