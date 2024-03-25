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
	_ datasource.DataSource              = &OrganizationsCellularGatewayUplinkStatusesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsCellularGatewayUplinkStatusesDataSource{}
)

func NewOrganizationsCellularGatewayUplinkStatusesDataSource() datasource.DataSource {
	return &OrganizationsCellularGatewayUplinkStatusesDataSource{}
}

type OrganizationsCellularGatewayUplinkStatusesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsCellularGatewayUplinkStatusesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsCellularGatewayUplinkStatusesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_cellular_gateway_uplink_statuses"
}

func (d *OrganizationsCellularGatewayUplinkStatusesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"iccids": schema.ListAttribute{
				MarkdownDescription: `iccids query parameter. A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. A list of network IDs. The returned devices will be filtered to only include these networks.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.`,
				Optional:            true,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. A list of serial numbers. The returned devices will be filtered to only include these serials.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseCellularGatewayGetOrganizationCellularGatewayUplinkStatuses`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"last_reported_at": schema.StringAttribute{
							MarkdownDescription: `Last reported time for the device`,
							Computed:            true,
						},
						"model": schema.StringAttribute{
							MarkdownDescription: `Device model`,
							Computed:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: `Network Id`,
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `Serial number of the device`,
							Computed:            true,
						},
						"uplinks": schema.SetNestedAttribute{
							MarkdownDescription: `Uplinks info`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"apn": schema.StringAttribute{
										MarkdownDescription: `Access Point Name`,
										Computed:            true,
									},
									"connection_type": schema.StringAttribute{
										MarkdownDescription: `Connection Type`,
										Computed:            true,
									},
									"dns1": schema.StringAttribute{
										MarkdownDescription: `Primary DNS IP`,
										Computed:            true,
									},
									"dns2": schema.StringAttribute{
										MarkdownDescription: `Secondary DNS IP`,
										Computed:            true,
									},
									"gateway": schema.StringAttribute{
										MarkdownDescription: `Gateway IP`,
										Computed:            true,
									},
									"iccid": schema.StringAttribute{
										MarkdownDescription: `Integrated Circuit Card Identification Number`,
										Computed:            true,
									},
									"interface": schema.StringAttribute{
										MarkdownDescription: `Uplink interface`,
										Computed:            true,
									},
									"ip": schema.StringAttribute{
										MarkdownDescription: `Uplink IP`,
										Computed:            true,
									},
									"model": schema.StringAttribute{
										MarkdownDescription: `Uplink model`,
										Computed:            true,
									},
									"provider": schema.StringAttribute{
										MarkdownDescription: `Network Provider`,
										Computed:            true,
									},
									"public_ip": schema.StringAttribute{
										MarkdownDescription: `Public IP`,
										Computed:            true,
									},
									"signal_stat": schema.SingleNestedAttribute{
										MarkdownDescription: `Tower Signal Status`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"rsrp": schema.StringAttribute{
												MarkdownDescription: `Reference Signal Received Power`,
												Computed:            true,
											},
											"rsrq": schema.StringAttribute{
												MarkdownDescription: `Reference Signal Received Quality`,
												Computed:            true,
											},
										},
									},
									"signal_type": schema.StringAttribute{
										MarkdownDescription: `Signal Type`,
										Computed:            true,
									},
									"status": schema.StringAttribute{
										MarkdownDescription: `Uplink status`,
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

func (d *OrganizationsCellularGatewayUplinkStatusesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsCellularGatewayUplinkStatuses OrganizationsCellularGatewayUplinkStatuses
	diags := req.Config.Get(ctx, &organizationsCellularGatewayUplinkStatuses)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationCellularGatewayUplinkStatuses")
		vvOrganizationID := organizationsCellularGatewayUplinkStatuses.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationCellularGatewayUplinkStatusesQueryParams{}

		queryParams1.PerPage = int(organizationsCellularGatewayUplinkStatuses.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsCellularGatewayUplinkStatuses.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsCellularGatewayUplinkStatuses.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsCellularGatewayUplinkStatuses.NetworkIDs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsCellularGatewayUplinkStatuses.Serials)
		queryParams1.Iccids = elementsToStrings(ctx, organizationsCellularGatewayUplinkStatuses.Iccids)

		response1, restyResp1, err := d.client.CellularGateway.GetOrganizationCellularGatewayUplinkStatuses(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCellularGatewayUplinkStatuses",
				err.Error(),
			)
			return
		}

		organizationsCellularGatewayUplinkStatuses = ResponseCellularGatewayGetOrganizationCellularGatewayUplinkStatusesItemsToBody(organizationsCellularGatewayUplinkStatuses, response1)
		diags = resp.State.Set(ctx, &organizationsCellularGatewayUplinkStatuses)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsCellularGatewayUplinkStatuses struct {
	OrganizationID types.String                                                               `tfsdk:"organization_id"`
	PerPage        types.Int64                                                                `tfsdk:"per_page"`
	StartingAfter  types.String                                                               `tfsdk:"starting_after"`
	EndingBefore   types.String                                                               `tfsdk:"ending_before"`
	NetworkIDs     types.List                                                                 `tfsdk:"network_ids"`
	Serials        types.List                                                                 `tfsdk:"serials"`
	Iccids         types.List                                                                 `tfsdk:"iccids"`
	Items          *[]ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatuses `tfsdk:"items"`
}

type ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatuses struct {
	LastReportedAt types.String                                                                      `tfsdk:"last_reported_at"`
	Model          types.String                                                                      `tfsdk:"model"`
	NetworkID      types.String                                                                      `tfsdk:"network_id"`
	Serial         types.String                                                                      `tfsdk:"serial"`
	Uplinks        *[]ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinks `tfsdk:"uplinks"`
}

type ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinks struct {
	Apn            types.String                                                                              `tfsdk:"apn"`
	ConnectionType types.String                                                                              `tfsdk:"connection_type"`
	DNS1           types.String                                                                              `tfsdk:"dns1"`
	DNS2           types.String                                                                              `tfsdk:"dns2"`
	Gateway        types.String                                                                              `tfsdk:"gateway"`
	Iccid          types.String                                                                              `tfsdk:"iccid"`
	Interface      types.String                                                                              `tfsdk:"interface"`
	IP             types.String                                                                              `tfsdk:"ip"`
	Model          types.String                                                                              `tfsdk:"model"`
	Provider       types.String                                                                              `tfsdk:"provider"`
	PublicIP       types.String                                                                              `tfsdk:"public_ip"`
	SignalStat     *ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinksSignalStat `tfsdk:"signal_stat"`
	SignalType     types.String                                                                              `tfsdk:"signal_type"`
	Status         types.String                                                                              `tfsdk:"status"`
}

type ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinksSignalStat struct {
	Rsrp types.String `tfsdk:"rsrp"`
	Rsrq types.String `tfsdk:"rsrq"`
}

// ToBody
func ResponseCellularGatewayGetOrganizationCellularGatewayUplinkStatusesItemsToBody(state OrganizationsCellularGatewayUplinkStatuses, response *merakigosdk.ResponseCellularGatewayGetOrganizationCellularGatewayUplinkStatuses) OrganizationsCellularGatewayUplinkStatuses {
	var items []ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatuses
	for _, item := range *response {
		itemState := ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatuses{
			LastReportedAt: types.StringValue(item.LastReportedAt),
			Model:          types.StringValue(item.Model),
			NetworkID:      types.StringValue(item.NetworkID),
			Serial:         types.StringValue(item.Serial),
			Uplinks: func() *[]ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinks {
				if item.Uplinks != nil {
					result := make([]ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinks, len(*item.Uplinks))
					for i, uplinks := range *item.Uplinks {
						result[i] = ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinks{
							Apn:            types.StringValue(uplinks.Apn),
							ConnectionType: types.StringValue(uplinks.ConnectionType),
							DNS1:           types.StringValue(uplinks.DNS1),
							DNS2:           types.StringValue(uplinks.DNS2),
							Gateway:        types.StringValue(uplinks.Gateway),
							Iccid:          types.StringValue(uplinks.Iccid),
							Interface:      types.StringValue(uplinks.Interface),
							IP:             types.StringValue(uplinks.IP),
							Model:          types.StringValue(uplinks.Model),
							Provider:       types.StringValue(uplinks.Provider),
							PublicIP:       types.StringValue(uplinks.PublicIP),
							SignalStat: func() *ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinksSignalStat {
								if uplinks.SignalStat != nil {
									return &ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinksSignalStat{
										Rsrp: types.StringValue(uplinks.SignalStat.Rsrp),
										Rsrq: types.StringValue(uplinks.SignalStat.Rsrq),
									}
								}
								return &ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinksSignalStat{}
							}(),
							SignalType: types.StringValue(uplinks.SignalType),
							Status:     types.StringValue(uplinks.Status),
						}
					}
					return &result
				}
				return &[]ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinks{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
