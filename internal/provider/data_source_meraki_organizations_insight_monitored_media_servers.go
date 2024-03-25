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
	_ datasource.DataSource              = &OrganizationsInsightMonitoredMediaServersDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsInsightMonitoredMediaServersDataSource{}
)

func NewOrganizationsInsightMonitoredMediaServersDataSource() datasource.DataSource {
	return &OrganizationsInsightMonitoredMediaServersDataSource{}
}

type OrganizationsInsightMonitoredMediaServersDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsInsightMonitoredMediaServersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsInsightMonitoredMediaServersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_insight_monitored_media_servers"
}

func (d *OrganizationsInsightMonitoredMediaServersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"monitored_media_server_id": schema.StringAttribute{
				MarkdownDescription: `monitoredMediaServerId path parameter. Monitored media server ID`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"address": schema.StringAttribute{
						Computed: true,
					},
					"best_effort_monitoring_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"id": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseInsightGetOrganizationInsightMonitoredMediaServers`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"address": schema.StringAttribute{
							MarkdownDescription: `The IP address (IPv4 only) or hostname of the media server to monitor`,
							Computed:            true,
						},
						"best_effort_monitoring_enabled": schema.BoolAttribute{
							MarkdownDescription: `Indicates that if the media server doesn't respond to ICMP pings, the nearest hop will be used in its stead`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `Monitored media server id`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the VoIP provider`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsInsightMonitoredMediaServersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsInsightMonitoredMediaServers OrganizationsInsightMonitoredMediaServers
	diags := req.Config.Get(ctx, &organizationsInsightMonitoredMediaServers)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsInsightMonitoredMediaServers.OrganizationID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsInsightMonitoredMediaServers.OrganizationID.IsNull(), !organizationsInsightMonitoredMediaServers.MonitoredMediaServerID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationInsightMonitoredMediaServers")
		vvOrganizationID := organizationsInsightMonitoredMediaServers.OrganizationID.ValueString()

		response1, restyResp1, err := d.client.Insight.GetOrganizationInsightMonitoredMediaServers(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationInsightMonitoredMediaServers",
				err.Error(),
			)
			return
		}

		organizationsInsightMonitoredMediaServers = ResponseInsightGetOrganizationInsightMonitoredMediaServersItemsToBody(organizationsInsightMonitoredMediaServers, response1)
		diags = resp.State.Set(ctx, &organizationsInsightMonitoredMediaServers)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationInsightMonitoredMediaServer")
		vvOrganizationID := organizationsInsightMonitoredMediaServers.OrganizationID.ValueString()
		vvMonitoredMediaServerID := organizationsInsightMonitoredMediaServers.MonitoredMediaServerID.ValueString()

		response2, restyResp2, err := d.client.Insight.GetOrganizationInsightMonitoredMediaServer(vvOrganizationID, vvMonitoredMediaServerID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationInsightMonitoredMediaServer",
				err.Error(),
			)
			return
		}

		organizationsInsightMonitoredMediaServers = ResponseInsightGetOrganizationInsightMonitoredMediaServerItemToBody(organizationsInsightMonitoredMediaServers, response2)
		diags = resp.State.Set(ctx, &organizationsInsightMonitoredMediaServers)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsInsightMonitoredMediaServers struct {
	OrganizationID         types.String                                                      `tfsdk:"organization_id"`
	MonitoredMediaServerID types.String                                                      `tfsdk:"monitored_media_server_id"`
	Items                  *[]ResponseItemInsightGetOrganizationInsightMonitoredMediaServers `tfsdk:"items"`
	Item                   *ResponseInsightGetOrganizationInsightMonitoredMediaServer        `tfsdk:"item"`
}

type ResponseItemInsightGetOrganizationInsightMonitoredMediaServers struct {
	Address                     types.String `tfsdk:"address"`
	BestEffortMonitoringEnabled types.Bool   `tfsdk:"best_effort_monitoring_enabled"`
	ID                          types.String `tfsdk:"id"`
	Name                        types.String `tfsdk:"name"`
}

type ResponseInsightGetOrganizationInsightMonitoredMediaServer struct {
	Address                     types.String `tfsdk:"address"`
	BestEffortMonitoringEnabled types.Bool   `tfsdk:"best_effort_monitoring_enabled"`
	ID                          types.String `tfsdk:"id"`
	Name                        types.String `tfsdk:"name"`
}

// ToBody
func ResponseInsightGetOrganizationInsightMonitoredMediaServersItemsToBody(state OrganizationsInsightMonitoredMediaServers, response *merakigosdk.ResponseInsightGetOrganizationInsightMonitoredMediaServers) OrganizationsInsightMonitoredMediaServers {
	var items []ResponseItemInsightGetOrganizationInsightMonitoredMediaServers
	for _, item := range *response {
		itemState := ResponseItemInsightGetOrganizationInsightMonitoredMediaServers{
			Address: types.StringValue(item.Address),
			BestEffortMonitoringEnabled: func() types.Bool {
				if item.BestEffortMonitoringEnabled != nil {
					return types.BoolValue(*item.BestEffortMonitoringEnabled)
				}
				return types.Bool{}
			}(),
			ID:   types.StringValue(item.ID),
			Name: types.StringValue(item.Name),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseInsightGetOrganizationInsightMonitoredMediaServerItemToBody(state OrganizationsInsightMonitoredMediaServers, response *merakigosdk.ResponseInsightGetOrganizationInsightMonitoredMediaServer) OrganizationsInsightMonitoredMediaServers {
	itemState := ResponseInsightGetOrganizationInsightMonitoredMediaServer{
		Address: types.StringValue(response.Address),
		BestEffortMonitoringEnabled: func() types.Bool {
			if response.BestEffortMonitoringEnabled != nil {
				return types.BoolValue(*response.BestEffortMonitoringEnabled)
			}
			return types.Bool{}
		}(),
		ID:   types.StringValue(response.ID),
		Name: types.StringValue(response.Name),
	}
	state.Item = &itemState
	return state
}
