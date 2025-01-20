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
	_ datasource.DataSource              = &NetworksSmUserAccessDevicesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSmUserAccessDevicesDataSource{}
)

func NewNetworksSmUserAccessDevicesDataSource() datasource.DataSource {
	return &NetworksSmUserAccessDevicesDataSource{}
}

type NetworksSmUserAccessDevicesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSmUserAccessDevicesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSmUserAccessDevicesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_user_access_devices"
}

func (d *NetworksSmUserAccessDevicesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 100.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSmGetNetworkSmUserAccessDevices`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"email": schema.StringAttribute{
							MarkdownDescription: `user email`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `device ID`,
							Computed:            true,
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: `mac address`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `device name`,
							Computed:            true,
						},
						"system_type": schema.StringAttribute{
							MarkdownDescription: `system type`,
							Computed:            true,
						},
						"tags": schema.ListAttribute{
							MarkdownDescription: `device tags`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"trusted_access_connections": schema.SetNestedAttribute{
							MarkdownDescription: `Array of trusted access configs`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"downloaded_at": schema.StringAttribute{
										MarkdownDescription: `time that config was downloaded`,
										Computed:            true,
									},
									"last_connected_at": schema.StringAttribute{
										MarkdownDescription: `time of last connection`,
										Computed:            true,
									},
									"scep_completed_at": schema.StringAttribute{
										MarkdownDescription: `time that SCEP completed`,
										Computed:            true,
									},
									"trusted_access_config_id": schema.StringAttribute{
										MarkdownDescription: `config id`,
										Computed:            true,
									},
								},
							},
						},
						"username": schema.StringAttribute{
							MarkdownDescription: `username`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSmUserAccessDevicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSmUserAccessDevices NetworksSmUserAccessDevices
	diags := req.Config.Get(ctx, &networksSmUserAccessDevices)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmUserAccessDevices")
		vvNetworkID := networksSmUserAccessDevices.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkSmUserAccessDevicesQueryParams{}

		queryParams1.PerPage = int(networksSmUserAccessDevices.PerPage.ValueInt64())
		queryParams1.StartingAfter = networksSmUserAccessDevices.StartingAfter.ValueString()
		queryParams1.EndingBefore = networksSmUserAccessDevices.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sm.GetNetworkSmUserAccessDevices(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmUserAccessDevices",
				err.Error(),
			)
			return
		}

		networksSmUserAccessDevices = ResponseSmGetNetworkSmUserAccessDevicesItemsToBody(networksSmUserAccessDevices, response1)
		diags = resp.State.Set(ctx, &networksSmUserAccessDevices)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSmUserAccessDevices struct {
	NetworkID     types.String                                   `tfsdk:"network_id"`
	PerPage       types.Int64                                    `tfsdk:"per_page"`
	StartingAfter types.String                                   `tfsdk:"starting_after"`
	EndingBefore  types.String                                   `tfsdk:"ending_before"`
	Items         *[]ResponseItemSmGetNetworkSmUserAccessDevices `tfsdk:"items"`
}

type ResponseItemSmGetNetworkSmUserAccessDevices struct {
	Email                    types.String                                                           `tfsdk:"email"`
	ID                       types.String                                                           `tfsdk:"id"`
	Mac                      types.String                                                           `tfsdk:"mac"`
	Name                     types.String                                                           `tfsdk:"name"`
	SystemType               types.String                                                           `tfsdk:"system_type"`
	Tags                     types.List                                                             `tfsdk:"tags"`
	TrustedAccessConnections *[]ResponseItemSmGetNetworkSmUserAccessDevicesTrustedAccessConnections `tfsdk:"trusted_access_connections"`
	Username                 types.String                                                           `tfsdk:"username"`
}

type ResponseItemSmGetNetworkSmUserAccessDevicesTrustedAccessConnections struct {
	DownloadedAt          types.String `tfsdk:"downloaded_at"`
	LastConnectedAt       types.String `tfsdk:"last_connected_at"`
	ScepCompletedAt       types.String `tfsdk:"scep_completed_at"`
	TrustedAccessConfigID types.String `tfsdk:"trusted_access_config_id"`
}

// ToBody
func ResponseSmGetNetworkSmUserAccessDevicesItemsToBody(state NetworksSmUserAccessDevices, response *merakigosdk.ResponseSmGetNetworkSmUserAccessDevices) NetworksSmUserAccessDevices {
	var items []ResponseItemSmGetNetworkSmUserAccessDevices
	for _, item := range *response {
		itemState := ResponseItemSmGetNetworkSmUserAccessDevices{
			Email:      types.StringValue(item.Email),
			ID:         types.StringValue(item.ID),
			Mac:        types.StringValue(item.Mac),
			Name:       types.StringValue(item.Name),
			SystemType: types.StringValue(item.SystemType),
			Tags:       StringSliceToList(item.Tags),
			TrustedAccessConnections: func() *[]ResponseItemSmGetNetworkSmUserAccessDevicesTrustedAccessConnections {
				if item.TrustedAccessConnections != nil {
					result := make([]ResponseItemSmGetNetworkSmUserAccessDevicesTrustedAccessConnections, len(*item.TrustedAccessConnections))
					for i, trustedAccessConnections := range *item.TrustedAccessConnections {
						result[i] = ResponseItemSmGetNetworkSmUserAccessDevicesTrustedAccessConnections{
							DownloadedAt:          types.StringValue(trustedAccessConnections.DownloadedAt),
							LastConnectedAt:       types.StringValue(trustedAccessConnections.LastConnectedAt),
							ScepCompletedAt:       types.StringValue(trustedAccessConnections.ScepCompletedAt),
							TrustedAccessConfigID: types.StringValue(trustedAccessConnections.TrustedAccessConfigID),
						}
					}
					return &result
				}
				return nil
			}(),
			Username: types.StringValue(item.Username),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
