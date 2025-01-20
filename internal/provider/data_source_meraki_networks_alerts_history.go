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
	_ datasource.DataSource              = &NetworksAlertsHistoryDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksAlertsHistoryDataSource{}
)

func NewNetworksAlertsHistoryDataSource() datasource.DataSource {
	return &NetworksAlertsHistoryDataSource{}
}

type NetworksAlertsHistoryDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksAlertsHistoryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksAlertsHistoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_alerts_history"
}

func (d *NetworksAlertsHistoryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				MarkdownDescription: `Array of ResponseNetworksGetNetworkAlertsHistory`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"alert_data": schema.StringAttribute{
							//Entro en string ds
							//TODO interface
							MarkdownDescription: `relevant data about the event that caused the alert`,
							Computed:            true,
						},
						"alert_type": schema.StringAttribute{
							MarkdownDescription: `user friendly alert type`,
							Computed:            true,
						},
						"alert_type_id": schema.StringAttribute{
							MarkdownDescription: `type of alert`,
							Computed:            true,
						},
						"destinations": schema.SingleNestedAttribute{
							MarkdownDescription: `the destinations this alert is configured to be delivered to`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"email": schema.SingleNestedAttribute{
									MarkdownDescription: `email destinations for this alert`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"sent_at": schema.StringAttribute{
											MarkdownDescription: `time when the alert was sent to the user(s) for this channel`,
											Computed:            true,
										},
									},
								},
								"push": schema.SingleNestedAttribute{
									MarkdownDescription: `push destinations for this alert`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"sent_at": schema.StringAttribute{
											MarkdownDescription: `time when the alert was sent to the user(s) for this channel`,
											Computed:            true,
										},
									},
								},
								"sms": schema.SingleNestedAttribute{
									MarkdownDescription: `sms destinations for this alert`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"sent_at": schema.StringAttribute{
											MarkdownDescription: `time when the alert was sent to the user(s) for this channel`,
											Computed:            true,
										},
									},
								},
								"webhook": schema.SingleNestedAttribute{
									MarkdownDescription: `webhook destinations for this alert`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"sent_at": schema.StringAttribute{
											MarkdownDescription: `time when the alert was sent to the user(s) for this channel`,
											Computed:            true,
										},
									},
								},
							},
						},
						"device": schema.SingleNestedAttribute{
							MarkdownDescription: `info related to the device that caused the alert`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"serial": schema.StringAttribute{
									MarkdownDescription: `device serial`,
									Computed:            true,
								},
							},
						},
						"occurred_at": schema.StringAttribute{
							MarkdownDescription: `time when the event occurred`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksAlertsHistoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksAlertsHistory NetworksAlertsHistory
	diags := req.Config.Get(ctx, &networksAlertsHistory)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkAlertsHistory")
		vvNetworkID := networksAlertsHistory.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkAlertsHistoryQueryParams{}

		queryParams1.PerPage = int(networksAlertsHistory.PerPage.ValueInt64())
		queryParams1.StartingAfter = networksAlertsHistory.StartingAfter.ValueString()
		queryParams1.EndingBefore = networksAlertsHistory.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkAlertsHistory(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkAlertsHistory",
				err.Error(),
			)
			return
		}

		networksAlertsHistory = ResponseNetworksGetNetworkAlertsHistoryItemsToBody(networksAlertsHistory, response1)
		diags = resp.State.Set(ctx, &networksAlertsHistory)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksAlertsHistory struct {
	NetworkID     types.String                                   `tfsdk:"network_id"`
	PerPage       types.Int64                                    `tfsdk:"per_page"`
	StartingAfter types.String                                   `tfsdk:"starting_after"`
	EndingBefore  types.String                                   `tfsdk:"ending_before"`
	Items         *[]ResponseItemNetworksGetNetworkAlertsHistory `tfsdk:"items"`
}

type ResponseItemNetworksGetNetworkAlertsHistory struct {
	AlertData    *ResponseItemNetworksGetNetworkAlertsHistoryAlertData    `tfsdk:"alert_data"`
	AlertType    types.String                                             `tfsdk:"alert_type"`
	AlertTypeID  types.String                                             `tfsdk:"alert_type_id"`
	Destinations *ResponseItemNetworksGetNetworkAlertsHistoryDestinations `tfsdk:"destinations"`
	Device       *ResponseItemNetworksGetNetworkAlertsHistoryDevice       `tfsdk:"device"`
	OccurredAt   types.String                                             `tfsdk:"occurred_at"`
}

type ResponseItemNetworksGetNetworkAlertsHistoryAlertData interface{}

type ResponseItemNetworksGetNetworkAlertsHistoryDestinations struct {
	Email   *ResponseItemNetworksGetNetworkAlertsHistoryDestinationsEmail   `tfsdk:"email"`
	Push    *ResponseItemNetworksGetNetworkAlertsHistoryDestinationsPush    `tfsdk:"push"`
	Sms     *ResponseItemNetworksGetNetworkAlertsHistoryDestinationsSms     `tfsdk:"sms"`
	Webhook *ResponseItemNetworksGetNetworkAlertsHistoryDestinationsWebhook `tfsdk:"webhook"`
}

type ResponseItemNetworksGetNetworkAlertsHistoryDestinationsEmail struct {
	SentAt types.String `tfsdk:"sent_at"`
}

type ResponseItemNetworksGetNetworkAlertsHistoryDestinationsPush struct {
	SentAt types.String `tfsdk:"sent_at"`
}

type ResponseItemNetworksGetNetworkAlertsHistoryDestinationsSms struct {
	SentAt types.String `tfsdk:"sent_at"`
}

type ResponseItemNetworksGetNetworkAlertsHistoryDestinationsWebhook struct {
	SentAt types.String `tfsdk:"sent_at"`
}

type ResponseItemNetworksGetNetworkAlertsHistoryDevice struct {
	Serial types.String `tfsdk:"serial"`
}

// ToBody
func ResponseNetworksGetNetworkAlertsHistoryItemsToBody(state NetworksAlertsHistory, response *merakigosdk.ResponseNetworksGetNetworkAlertsHistory) NetworksAlertsHistory {
	var items []ResponseItemNetworksGetNetworkAlertsHistory
	for _, item := range *response {
		itemState := ResponseItemNetworksGetNetworkAlertsHistory{
			// AlertData:   types.StringValue(item.AlertData), //TODO POSIBLE interface
			AlertType:   types.StringValue(item.AlertType),
			AlertTypeID: types.StringValue(item.AlertTypeID),
			Destinations: func() *ResponseItemNetworksGetNetworkAlertsHistoryDestinations {
				if item.Destinations != nil {
					return &ResponseItemNetworksGetNetworkAlertsHistoryDestinations{
						Email: func() *ResponseItemNetworksGetNetworkAlertsHistoryDestinationsEmail {
							if item.Destinations.Email != nil {
								return &ResponseItemNetworksGetNetworkAlertsHistoryDestinationsEmail{
									SentAt: types.StringValue(item.Destinations.Email.SentAt),
								}
							}
							return nil
						}(),
						Push: func() *ResponseItemNetworksGetNetworkAlertsHistoryDestinationsPush {
							if item.Destinations.Push != nil {
								return &ResponseItemNetworksGetNetworkAlertsHistoryDestinationsPush{
									SentAt: types.StringValue(item.Destinations.Push.SentAt),
								}
							}
							return nil
						}(),
						Sms: func() *ResponseItemNetworksGetNetworkAlertsHistoryDestinationsSms {
							if item.Destinations.Sms != nil {
								return &ResponseItemNetworksGetNetworkAlertsHistoryDestinationsSms{
									SentAt: types.StringValue(item.Destinations.Sms.SentAt),
								}
							}
							return nil
						}(),
						Webhook: func() *ResponseItemNetworksGetNetworkAlertsHistoryDestinationsWebhook {
							if item.Destinations.Webhook != nil {
								return &ResponseItemNetworksGetNetworkAlertsHistoryDestinationsWebhook{
									SentAt: types.StringValue(item.Destinations.Webhook.SentAt),
								}
							}
							return nil
						}(),
					}
				}
				return nil
			}(),
			Device: func() *ResponseItemNetworksGetNetworkAlertsHistoryDevice {
				if item.Device != nil {
					return &ResponseItemNetworksGetNetworkAlertsHistoryDevice{
						Serial: types.StringValue(item.Device.Serial),
					}
				}
				return nil
			}(),
			OccurredAt: types.StringValue(item.OccurredAt),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
