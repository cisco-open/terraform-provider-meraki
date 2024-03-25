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
	_ datasource.DataSource              = &NetworksEventsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksEventsDataSource{}
)

func NewNetworksEventsDataSource() datasource.DataSource {
	return &NetworksEventsDataSource{}
}

type NetworksEventsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksEventsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksEventsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_events"
}

func (d *NetworksEventsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"client_ip": schema.StringAttribute{
				MarkdownDescription: `clientIp query parameter. The IP of the client which the list of events will be filtered with. Only supported for track-by-IP networks.`,
				Optional:            true,
			},
			"client_mac": schema.StringAttribute{
				MarkdownDescription: `clientMac query parameter. The MAC address of the client which the list of events will be filtered with. Only supported for track-by-MAC networks.`,
				Optional:            true,
			},
			"client_name": schema.StringAttribute{
				MarkdownDescription: `clientName query parameter. The name, or partial name, of the client which the list of events will be filtered with`,
				Optional:            true,
			},
			"device_mac": schema.StringAttribute{
				MarkdownDescription: `deviceMac query parameter. The MAC address of the Meraki device which the list of events will be filtered with`,
				Optional:            true,
			},
			"device_name": schema.StringAttribute{
				MarkdownDescription: `deviceName query parameter. The name of the Meraki device which the list of events will be filtered with`,
				Optional:            true,
			},
			"device_serial": schema.StringAttribute{
				MarkdownDescription: `deviceSerial query parameter. The serial of the Meraki device which the list of events will be filtered with`,
				Optional:            true,
			},
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"excluded_event_types": schema.ListAttribute{
				MarkdownDescription: `excludedEventTypes query parameter. A list of event types. The returned events will be filtered to exclude events with these types.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"included_event_types": schema.ListAttribute{
				MarkdownDescription: `includedEventTypes query parameter. A list of event types. The returned events will be filtered to only include events with these types.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 10.`,
				Optional:            true,
			},
			"product_type": schema.StringAttribute{
				MarkdownDescription: `productType query parameter. The product type to fetch events for. This parameter is required for networks with multiple device types. Valid types are wireless, appliance, switch, systemsManager, camera, and cellularGateway`,
				Optional:            true,
			},
			"sm_device_mac": schema.StringAttribute{
				MarkdownDescription: `smDeviceMac query parameter. The MAC address of the Systems Manager device which the list of events will be filtered with`,
				Optional:            true,
			},
			"sm_device_name": schema.StringAttribute{
				MarkdownDescription: `smDeviceName query parameter. The name of the Systems Manager device which the list of events will be filtered with`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"events": schema.SetNestedAttribute{
						MarkdownDescription: `An array of events that took place in the network.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"category": schema.StringAttribute{
									MarkdownDescription: `The category that the event type belongs to`,
									Computed:            true,
								},
								"client_description": schema.StringAttribute{
									MarkdownDescription: `A description of the client. This is usually the client's device name.`,
									Computed:            true,
								},
								"client_id": schema.StringAttribute{
									MarkdownDescription: `A string identifying the client. This could be a client's MAC or IP address`,
									Computed:            true,
								},
								"client_mac": schema.StringAttribute{
									MarkdownDescription: `The client's MAC address.`,
									Computed:            true,
								},
								"description": schema.StringAttribute{
									MarkdownDescription: `A description of the event the happened.`,
									Computed:            true,
								},
								"device_name": schema.StringAttribute{
									MarkdownDescription: `The name of the device. Only shown if the device is an access point.`,
									Computed:            true,
								},
								"device_serial": schema.StringAttribute{
									MarkdownDescription: `The serial number of the device. Only shown if the device is an access point.`,
									Computed:            true,
								},
								"event_data": schema.SingleNestedAttribute{
									MarkdownDescription: `An object containing more data related to the event.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"aid": schema.StringAttribute{
											MarkdownDescription: `The association ID of the client.`,
											Computed:            true,
										},
										"channel": schema.StringAttribute{
											MarkdownDescription: `The radio channel the client is connecting to.`,
											Computed:            true,
										},
										"client_ip": schema.StringAttribute{
											MarkdownDescription: `The client's IP address`,
											Computed:            true,
										},
										"client_mac": schema.StringAttribute{
											MarkdownDescription: `The client's MAC address`,
											Computed:            true,
										},
										"radio": schema.StringAttribute{
											MarkdownDescription: `The radio band number the client is trying to connect to.`,
											Computed:            true,
										},
										"rssi": schema.StringAttribute{
											MarkdownDescription: `The current received signal strength indication (RSSI) of the client connected to an AP.`,
											Computed:            true,
										},
										"vap": schema.StringAttribute{
											MarkdownDescription: `The virtual access point (VAP) number the client is connecting to.`,
											Computed:            true,
										},
									},
								},
								"network_id": schema.StringAttribute{
									MarkdownDescription: `The ID of the network.`,
									Computed:            true,
								},
								"occurred_at": schema.StringAttribute{
									MarkdownDescription: `An UTC ISO8601 string of the time the event occurred at.`,
									Computed:            true,
								},
								"ssid_number": schema.Int64Attribute{
									MarkdownDescription: `The SSID number of the device.`,
									Computed:            true,
								},
								"type": schema.StringAttribute{
									MarkdownDescription: `The type of event being listed.`,
									Computed:            true,
								},
							},
						},
					},
					"message": schema.StringAttribute{
						MarkdownDescription: `A message regarding the events sent. Usually 'null' unless there are no events`,
						Computed:            true,
					},
					"page_end_at": schema.StringAttribute{
						MarkdownDescription: `An UTC ISO8601 string of the latest occured at time of the listed events of the page.`,
						Computed:            true,
					},
					"page_start_at": schema.StringAttribute{
						MarkdownDescription: `An UTC ISO8601 string of the earliest occured at time of the listed events of the page.`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *NetworksEventsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksEvents NetworksEvents
	diags := req.Config.Get(ctx, &networksEvents)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkEvents")
		vvNetworkID := networksEvents.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkEventsQueryParams{}

		queryParams1.ProductType = networksEvents.ProductType.ValueString()
		queryParams1.IncludedEventTypes = elementsToStrings(ctx, networksEvents.IncludedEventTypes)
		queryParams1.ExcludedEventTypes = elementsToStrings(ctx, networksEvents.ExcludedEventTypes)
		queryParams1.DeviceMac = networksEvents.DeviceMac.ValueString()
		queryParams1.DeviceSerial = networksEvents.DeviceSerial.ValueString()
		queryParams1.DeviceName = networksEvents.DeviceName.ValueString()
		queryParams1.ClientIP = networksEvents.ClientIP.ValueString()
		queryParams1.ClientMac = networksEvents.ClientMac.ValueString()
		queryParams1.ClientName = networksEvents.ClientName.ValueString()
		queryParams1.SmDeviceMac = networksEvents.SmDeviceMac.ValueString()
		queryParams1.SmDeviceName = networksEvents.SmDeviceName.ValueString()
		queryParams1.PerPage = int(networksEvents.PerPage.ValueInt64())
		queryParams1.StartingAfter = networksEvents.StartingAfter.ValueString()
		queryParams1.EndingBefore = networksEvents.EndingBefore.ValueString()

		response1, restyResp1, err := d.client.Networks.GetNetworkEvents(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkEvents",
				err.Error(),
			)
			return
		}

		networksEvents = ResponseNetworksGetNetworkEventsItemToBody(networksEvents, response1)
		diags = resp.State.Set(ctx, &networksEvents)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksEvents struct {
	NetworkID          types.String                      `tfsdk:"network_id"`
	ProductType        types.String                      `tfsdk:"product_type"`
	IncludedEventTypes types.List                        `tfsdk:"included_event_types"`
	ExcludedEventTypes types.List                        `tfsdk:"excluded_event_types"`
	DeviceMac          types.String                      `tfsdk:"device_mac"`
	DeviceSerial       types.String                      `tfsdk:"device_serial"`
	DeviceName         types.String                      `tfsdk:"device_name"`
	ClientIP           types.String                      `tfsdk:"client_ip"`
	ClientMac          types.String                      `tfsdk:"client_mac"`
	ClientName         types.String                      `tfsdk:"client_name"`
	SmDeviceMac        types.String                      `tfsdk:"sm_device_mac"`
	SmDeviceName       types.String                      `tfsdk:"sm_device_name"`
	PerPage            types.Int64                       `tfsdk:"per_page"`
	StartingAfter      types.String                      `tfsdk:"starting_after"`
	EndingBefore       types.String                      `tfsdk:"ending_before"`
	Item               *ResponseNetworksGetNetworkEvents `tfsdk:"item"`
}

type ResponseNetworksGetNetworkEvents struct {
	Events      *[]ResponseNetworksGetNetworkEventsEvents `tfsdk:"events"`
	Message     types.String                              `tfsdk:"message"`
	PageEndAt   types.String                              `tfsdk:"page_end_at"`
	PageStartAt types.String                              `tfsdk:"page_start_at"`
}

type ResponseNetworksGetNetworkEventsEvents struct {
	Category          types.String                                     `tfsdk:"category"`
	ClientDescription types.String                                     `tfsdk:"client_description"`
	ClientID          types.String                                     `tfsdk:"client_id"`
	ClientMac         types.String                                     `tfsdk:"client_mac"`
	Description       types.String                                     `tfsdk:"description"`
	DeviceName        types.String                                     `tfsdk:"device_name"`
	DeviceSerial      types.String                                     `tfsdk:"device_serial"`
	EventData         *ResponseNetworksGetNetworkEventsEventsEventData `tfsdk:"event_data"`
	NetworkID         types.String                                     `tfsdk:"network_id"`
	OccurredAt        types.String                                     `tfsdk:"occurred_at"`
	SSIDNumber        types.Int64                                      `tfsdk:"ssid_number"`
	Type              types.String                                     `tfsdk:"type"`
}

type ResponseNetworksGetNetworkEventsEventsEventData struct {
	Aid       types.String `tfsdk:"aid"`
	Channel   types.String `tfsdk:"channel"`
	ClientIP  types.String `tfsdk:"client_ip"`
	ClientMac types.String `tfsdk:"client_mac"`
	Radio     types.String `tfsdk:"radio"`
	Rssi      types.String `tfsdk:"rssi"`
	Vap       types.String `tfsdk:"vap"`
}

// ToBody
func ResponseNetworksGetNetworkEventsItemToBody(state NetworksEvents, response *merakigosdk.ResponseNetworksGetNetworkEvents) NetworksEvents {
	itemState := ResponseNetworksGetNetworkEvents{
		Events: func() *[]ResponseNetworksGetNetworkEventsEvents {
			if response.Events != nil {
				result := make([]ResponseNetworksGetNetworkEventsEvents, len(*response.Events))
				for i, events := range *response.Events {
					result[i] = ResponseNetworksGetNetworkEventsEvents{
						Category:          types.StringValue(events.Category),
						ClientDescription: types.StringValue(events.ClientDescription),
						ClientID:          types.StringValue(events.ClientID),
						ClientMac:         types.StringValue(events.ClientMac),
						Description:       types.StringValue(events.Description),
						DeviceName:        types.StringValue(events.DeviceName),
						DeviceSerial:      types.StringValue(events.DeviceSerial),
						EventData: func() *ResponseNetworksGetNetworkEventsEventsEventData {
							if events.EventData != nil {
								return &ResponseNetworksGetNetworkEventsEventsEventData{
									Aid:       types.StringValue(events.EventData.Aid),
									Channel:   types.StringValue(events.EventData.Channel),
									ClientIP:  types.StringValue(events.EventData.ClientIP),
									ClientMac: types.StringValue(events.EventData.ClientMac),
									Radio:     types.StringValue(events.EventData.Radio),
									Rssi:      types.StringValue(events.EventData.Rssi),
									Vap:       types.StringValue(events.EventData.Vap),
								}
							}
							return &ResponseNetworksGetNetworkEventsEventsEventData{}
						}(),
						NetworkID:  types.StringValue(events.NetworkID),
						OccurredAt: types.StringValue(events.OccurredAt),
						SSIDNumber: func() types.Int64 {
							if events.SSIDNumber != nil {
								return types.Int64Value(int64(*events.SSIDNumber))
							}
							return types.Int64{}
						}(),
						Type: types.StringValue(events.Type),
					}
				}
				return &result
			}
			return &[]ResponseNetworksGetNetworkEventsEvents{}
		}(),
		Message:     types.StringValue(response.Message),
		PageEndAt:   types.StringValue(response.PageEndAt),
		PageStartAt: types.StringValue(response.PageStartAt),
	}
	state.Item = &itemState
	return state
}
