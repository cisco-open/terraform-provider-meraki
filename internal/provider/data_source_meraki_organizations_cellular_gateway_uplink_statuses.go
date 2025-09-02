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
									"imsi": schema.StringAttribute{
										MarkdownDescription: `International Mobile Subscriber Identity`,
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
									"mcc": schema.StringAttribute{
										MarkdownDescription: `Mobile Country Code`,
										Computed:            true,
									},
									"mnc": schema.StringAttribute{
										MarkdownDescription: `Mobile Network Code`,
										Computed:            true,
									},
									"model": schema.StringAttribute{
										MarkdownDescription: `Uplink model`,
										Computed:            true,
									},
									"msisdn": schema.StringAttribute{
										MarkdownDescription: `Mobile Station Integrated Services Digital Network`,
										Computed:            true,
									},
									"mtu": schema.Int64Attribute{
										MarkdownDescription: `Maximum Transmission Unit`,
										Computed:            true,
									},
									"provider_r": schema.StringAttribute{
										MarkdownDescription: `Network Provider`,
										Computed:            true,
									},
									"public_ip": schema.StringAttribute{
										MarkdownDescription: `Public IP`,
										Computed:            true,
									},
									"roaming": schema.SingleNestedAttribute{
										MarkdownDescription: `Roaming Status`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"status": schema.StringAttribute{
												MarkdownDescription: `Roaming Status`,
												Computed:            true,
											},
										},
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

		// has_unknown_response: None

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
	Imsi           types.String                                                                              `tfsdk:"imsi"`
	Interface      types.String                                                                              `tfsdk:"interface"`
	IP             types.String                                                                              `tfsdk:"ip"`
	Mcc            types.String                                                                              `tfsdk:"mcc"`
	Mnc            types.String                                                                              `tfsdk:"mnc"`
	Model          types.String                                                                              `tfsdk:"model"`
	Msisdn         types.String                                                                              `tfsdk:"msisdn"`
	Mtu            types.Int64                                                                               `tfsdk:"mtu"`
	Provider       types.String                                                                              `tfsdk:"provider_r"`
	PublicIP       types.String                                                                              `tfsdk:"public_ip"`
	Roaming        *ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinksRoaming    `tfsdk:"roaming"`
	SignalStat     *ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinksSignalStat `tfsdk:"signal_stat"`
	SignalType     types.String                                                                              `tfsdk:"signal_type"`
	Status         types.String                                                                              `tfsdk:"status"`
}

type ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinksRoaming struct {
	Status types.String `tfsdk:"status"`
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
			LastReportedAt: func() types.String {
				if item.LastReportedAt != "" {
					return types.StringValue(item.LastReportedAt)
				}
				return types.String{}
			}(),
			Model: func() types.String {
				if item.Model != "" {
					return types.StringValue(item.Model)
				}
				return types.String{}
			}(),
			NetworkID: func() types.String {
				if item.NetworkID != "" {
					return types.StringValue(item.NetworkID)
				}
				return types.String{}
			}(),
			Serial: func() types.String {
				if item.Serial != "" {
					return types.StringValue(item.Serial)
				}
				return types.String{}
			}(),
			Uplinks: func() *[]ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinks {
				if item.Uplinks != nil {
					result := make([]ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinks, len(*item.Uplinks))
					for i, uplinks := range *item.Uplinks {
						result[i] = ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinks{
							Apn: func() types.String {
								if uplinks.Apn != "" {
									return types.StringValue(uplinks.Apn)
								}
								return types.String{}
							}(),
							ConnectionType: func() types.String {
								if uplinks.ConnectionType != "" {
									return types.StringValue(uplinks.ConnectionType)
								}
								return types.String{}
							}(),
							DNS1: func() types.String {
								if uplinks.DNS1 != "" {
									return types.StringValue(uplinks.DNS1)
								}
								return types.String{}
							}(),
							DNS2: func() types.String {
								if uplinks.DNS2 != "" {
									return types.StringValue(uplinks.DNS2)
								}
								return types.String{}
							}(),
							Gateway: func() types.String {
								if uplinks.Gateway != "" {
									return types.StringValue(uplinks.Gateway)
								}
								return types.String{}
							}(),
							Iccid: func() types.String {
								if uplinks.Iccid != "" {
									return types.StringValue(uplinks.Iccid)
								}
								return types.String{}
							}(),
							Imsi: func() types.String {
								if uplinks.Imsi != "" {
									return types.StringValue(uplinks.Imsi)
								}
								return types.String{}
							}(),
							Interface: func() types.String {
								if uplinks.Interface != "" {
									return types.StringValue(uplinks.Interface)
								}
								return types.String{}
							}(),
							IP: func() types.String {
								if uplinks.IP != "" {
									return types.StringValue(uplinks.IP)
								}
								return types.String{}
							}(),
							Mcc: func() types.String {
								if uplinks.Mcc != "" {
									return types.StringValue(uplinks.Mcc)
								}
								return types.String{}
							}(),
							Mnc: func() types.String {
								if uplinks.Mnc != "" {
									return types.StringValue(uplinks.Mnc)
								}
								return types.String{}
							}(),
							Model: func() types.String {
								if uplinks.Model != "" {
									return types.StringValue(uplinks.Model)
								}
								return types.String{}
							}(),
							Msisdn: func() types.String {
								if uplinks.Msisdn != "" {
									return types.StringValue(uplinks.Msisdn)
								}
								return types.String{}
							}(),
							Mtu: func() types.Int64 {
								if uplinks.Mtu != nil {
									return types.Int64Value(int64(*uplinks.Mtu))
								}
								return types.Int64{}
							}(),
							Provider: func() types.String {
								if uplinks.Provider != "" {
									return types.StringValue(uplinks.Provider)
								}
								return types.String{}
							}(),
							PublicIP: func() types.String {
								if uplinks.PublicIP != "" {
									return types.StringValue(uplinks.PublicIP)
								}
								return types.String{}
							}(),
							Roaming: func() *ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinksRoaming {
								if uplinks.Roaming != nil {
									return &ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinksRoaming{
										Status: func() types.String {
											if uplinks.Roaming.Status != "" {
												return types.StringValue(uplinks.Roaming.Status)
											}
											return types.String{}
										}(),
									}
								}
								return nil
							}(),
							SignalStat: func() *ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinksSignalStat {
								if uplinks.SignalStat != nil {
									return &ResponseItemCellularGatewayGetOrganizationCellularGatewayUplinkStatusesUplinksSignalStat{
										Rsrp: func() types.String {
											if uplinks.SignalStat.Rsrp != "" {
												return types.StringValue(uplinks.SignalStat.Rsrp)
											}
											return types.String{}
										}(),
										Rsrq: func() types.String {
											if uplinks.SignalStat.Rsrq != "" {
												return types.StringValue(uplinks.SignalStat.Rsrq)
											}
											return types.String{}
										}(),
									}
								}
								return nil
							}(),
							SignalType: func() types.String {
								if uplinks.SignalType != "" {
									return types.StringValue(uplinks.SignalType)
								}
								return types.String{}
							}(),
							Status: func() types.String {
								if uplinks.Status != "" {
									return types.StringValue(uplinks.Status)
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
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
