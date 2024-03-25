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
	_ datasource.DataSource              = &NetworksSwitchPortSchedulesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchPortSchedulesDataSource{}
)

func NewNetworksSwitchPortSchedulesDataSource() datasource.DataSource {
	return &NetworksSwitchPortSchedulesDataSource{}
}

type NetworksSwitchPortSchedulesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchPortSchedulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchPortSchedulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_port_schedules"
}

func (d *NetworksSwitchPortSchedulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSwitchGetNetworkSwitchPortSchedules`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"id": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name for your port schedule. Required`,
							Computed:            true,
						},
						"port_schedule": schema.SingleNestedAttribute{
							MarkdownDescription: `    The schedule for switch port scheduling. Schedules are applied to days of the week.
    When it's empty, default schedule with all days of a week are configured.
    Any unspecified day in the schedule is added as a default schedule configuration of the day.
`,
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"friday": schema.SingleNestedAttribute{
									MarkdownDescription: `The schedule object for Friday.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"active": schema.BoolAttribute{
											MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
											Computed:            true,
										},
										"from": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
										"to": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
									},
								},
								"monday": schema.SingleNestedAttribute{
									MarkdownDescription: `The schedule object for Monday.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"active": schema.BoolAttribute{
											MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
											Computed:            true,
										},
										"from": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
										"to": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
									},
								},
								"saturday": schema.SingleNestedAttribute{
									MarkdownDescription: `The schedule object for Saturday.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"active": schema.BoolAttribute{
											MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
											Computed:            true,
										},
										"from": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
										"to": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
									},
								},
								"sunday": schema.SingleNestedAttribute{
									MarkdownDescription: `The schedule object for Sunday.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"active": schema.BoolAttribute{
											MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
											Computed:            true,
										},
										"from": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
										"to": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
									},
								},
								"thursday": schema.SingleNestedAttribute{
									MarkdownDescription: `The schedule object for Thursday.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"active": schema.BoolAttribute{
											MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
											Computed:            true,
										},
										"from": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
										"to": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
									},
								},
								"tuesday": schema.SingleNestedAttribute{
									MarkdownDescription: `The schedule object for Tuesday.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"active": schema.BoolAttribute{
											MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
											Computed:            true,
										},
										"from": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
										"to": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
									},
								},
								"wednesday": schema.SingleNestedAttribute{
									MarkdownDescription: `The schedule object for Wednesday.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"active": schema.BoolAttribute{
											MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
											Computed:            true,
										},
										"from": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
										"to": schema.StringAttribute{
											MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
											Computed:            true,
										},
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

func (d *NetworksSwitchPortSchedulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchPortSchedules NetworksSwitchPortSchedules
	diags := req.Config.Get(ctx, &networksSwitchPortSchedules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchPortSchedules")
		vvNetworkID := networksSwitchPortSchedules.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchPortSchedules(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchPortSchedules",
				err.Error(),
			)
			return
		}

		networksSwitchPortSchedules = ResponseSwitchGetNetworkSwitchPortSchedulesItemsToBody(networksSwitchPortSchedules, response1)
		diags = resp.State.Set(ctx, &networksSwitchPortSchedules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchPortSchedules struct {
	NetworkID types.String                                       `tfsdk:"network_id"`
	Items     *[]ResponseItemSwitchGetNetworkSwitchPortSchedules `tfsdk:"items"`
}

type ResponseItemSwitchGetNetworkSwitchPortSchedules struct {
	ID           types.String                                                 `tfsdk:"id"`
	Name         types.String                                                 `tfsdk:"name"`
	PortSchedule *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortSchedule `tfsdk:"port_schedule"`
}

type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortSchedule struct {
	Friday    *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleFriday    `tfsdk:"friday"`
	Monday    *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleMonday    `tfsdk:"monday"`
	Saturday  *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSaturday  `tfsdk:"saturday"`
	Sunday    *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSunday    `tfsdk:"sunday"`
	Thursday  *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleThursday  `tfsdk:"thursday"`
	Tuesday   *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleTuesday   `tfsdk:"tuesday"`
	Wednesday *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleWednesday `tfsdk:"wednesday"`
}

type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleFriday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleMonday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSaturday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSunday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleThursday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleTuesday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleWednesday struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchPortSchedulesItemsToBody(state NetworksSwitchPortSchedules, response *merakigosdk.ResponseSwitchGetNetworkSwitchPortSchedules) NetworksSwitchPortSchedules {
	var items []ResponseItemSwitchGetNetworkSwitchPortSchedules
	for _, item := range *response {
		itemState := ResponseItemSwitchGetNetworkSwitchPortSchedules{
			ID:   types.StringValue(item.ID),
			Name: types.StringValue(item.Name),
			PortSchedule: func() *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortSchedule {
				if item.PortSchedule != nil {
					return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortSchedule{
						Friday: func() *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleFriday {
							if item.PortSchedule.Friday != nil {
								return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleFriday{
									Active: func() types.Bool {
										if item.PortSchedule.Friday.Active != nil {
											return types.BoolValue(*item.PortSchedule.Friday.Active)
										}
										return types.Bool{}
									}(),
									From: types.StringValue(item.PortSchedule.Friday.From),
									To:   types.StringValue(item.PortSchedule.Friday.To),
								}
							}
							return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleFriday{}
						}(),
						Monday: func() *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleMonday {
							if item.PortSchedule.Monday != nil {
								return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleMonday{
									Active: func() types.Bool {
										if item.PortSchedule.Monday.Active != nil {
											return types.BoolValue(*item.PortSchedule.Monday.Active)
										}
										return types.Bool{}
									}(),
									From: types.StringValue(item.PortSchedule.Monday.From),
									To:   types.StringValue(item.PortSchedule.Monday.To),
								}
							}
							return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleMonday{}
						}(),
						Saturday: func() *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSaturday {
							if item.PortSchedule.Saturday != nil {
								return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSaturday{
									Active: func() types.Bool {
										if item.PortSchedule.Saturday.Active != nil {
											return types.BoolValue(*item.PortSchedule.Saturday.Active)
										}
										return types.Bool{}
									}(),
									From: types.StringValue(item.PortSchedule.Saturday.From),
									To:   types.StringValue(item.PortSchedule.Saturday.To),
								}
							}
							return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSaturday{}
						}(),
						Sunday: func() *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSunday {
							if item.PortSchedule.Sunday != nil {
								return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSunday{
									Active: func() types.Bool {
										if item.PortSchedule.Sunday.Active != nil {
											return types.BoolValue(*item.PortSchedule.Sunday.Active)
										}
										return types.Bool{}
									}(),
									From: types.StringValue(item.PortSchedule.Sunday.From),
									To:   types.StringValue(item.PortSchedule.Sunday.To),
								}
							}
							return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSunday{}
						}(),
						Thursday: func() *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleThursday {
							if item.PortSchedule.Thursday != nil {
								return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleThursday{
									Active: func() types.Bool {
										if item.PortSchedule.Thursday.Active != nil {
											return types.BoolValue(*item.PortSchedule.Thursday.Active)
										}
										return types.Bool{}
									}(),
									From: types.StringValue(item.PortSchedule.Thursday.From),
									To:   types.StringValue(item.PortSchedule.Thursday.To),
								}
							}
							return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleThursday{}
						}(),
						Tuesday: func() *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleTuesday {
							if item.PortSchedule.Tuesday != nil {
								return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleTuesday{
									Active: func() types.Bool {
										if item.PortSchedule.Tuesday.Active != nil {
											return types.BoolValue(*item.PortSchedule.Tuesday.Active)
										}
										return types.Bool{}
									}(),
									From: types.StringValue(item.PortSchedule.Tuesday.From),
									To:   types.StringValue(item.PortSchedule.Tuesday.To),
								}
							}
							return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleTuesday{}
						}(),
						Wednesday: func() *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleWednesday {
							if item.PortSchedule.Wednesday != nil {
								return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleWednesday{
									Active: func() types.Bool {
										if item.PortSchedule.Wednesday.Active != nil {
											return types.BoolValue(*item.PortSchedule.Wednesday.Active)
										}
										return types.Bool{}
									}(),
									From: types.StringValue(item.PortSchedule.Wednesday.From),
									To:   types.StringValue(item.PortSchedule.Wednesday.To),
								}
							}
							return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleWednesday{}
						}(),
					}
				}
				return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortSchedule{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
