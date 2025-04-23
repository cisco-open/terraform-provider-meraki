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
	_ datasource.DataSource              = &OrganizationsAlertsProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsAlertsProfilesDataSource{}
)

func NewOrganizationsAlertsProfilesDataSource() datasource.DataSource {
	return &OrganizationsAlertsProfilesDataSource{}
}

type OrganizationsAlertsProfilesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsAlertsProfilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsAlertsProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_alerts_profiles"
}

func (d *OrganizationsAlertsProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationAlertsProfiles`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"alert_condition": schema.SingleNestedAttribute{
							MarkdownDescription: `The conditions that determine if the alert triggers`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"bit_rate_bps": schema.Int64Attribute{
									MarkdownDescription: `The threshold the metric must cross to be valid for alerting. Used only for WAN Utilization alerts.`,
									Computed:            true,
								},
								"duration": schema.Int64Attribute{
									MarkdownDescription: `The total duration in seconds that the threshold should be crossed before alerting`,
									Computed:            true,
								},
								"interface": schema.StringAttribute{
									MarkdownDescription: `The uplink observed for the alert`,
									Computed:            true,
								},
								"window": schema.Int64Attribute{
									MarkdownDescription: `The look back period in seconds for sensing the alert`,
									Computed:            true,
								},
							},
						},
						"description": schema.StringAttribute{
							MarkdownDescription: `User supplied description of the alert`,
							Computed:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: `Is the alert config enabled`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `The alert config ID`,
							Computed:            true,
						},
						"network_tags": schema.ListAttribute{
							MarkdownDescription: `Networks with these tags will be monitored for the alert`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"recipients": schema.SingleNestedAttribute{
							MarkdownDescription: `List of recipients that will recieve the alert.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"emails": schema.ListAttribute{
									MarkdownDescription: `A list of emails that will receive information about the alert`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"http_server_ids": schema.ListAttribute{
									MarkdownDescription: `A list base64 encoded urls of webhook endpoints that will receive information about the alert`,
									Computed:            true,
									ElementType:         types.StringType,
								},
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `The alert type`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsAlertsProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsAlertsProfiles OrganizationsAlertsProfiles
	diags := req.Config.Get(ctx, &organizationsAlertsProfiles)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAlertsProfiles")
		vvOrganizationID := organizationsAlertsProfiles.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationAlertsProfiles(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAlertsProfiles",
				err.Error(),
			)
			return
		}

		organizationsAlertsProfiles = ResponseOrganizationsGetOrganizationAlertsProfilesItemsToBody(organizationsAlertsProfiles, response1)
		diags = resp.State.Set(ctx, &organizationsAlertsProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsAlertsProfiles struct {
	OrganizationID types.String                                              `tfsdk:"organization_id"`
	Items          *[]ResponseItemOrganizationsGetOrganizationAlertsProfiles `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationAlertsProfiles struct {
	AlertCondition *ResponseItemOrganizationsGetOrganizationAlertsProfilesAlertCondition `tfsdk:"alert_condition"`
	Description    types.String                                                          `tfsdk:"description"`
	Enabled        types.Bool                                                            `tfsdk:"enabled"`
	ID             types.String                                                          `tfsdk:"id"`
	NetworkTags    types.List                                                            `tfsdk:"network_tags"`
	Recipients     *ResponseItemOrganizationsGetOrganizationAlertsProfilesRecipients     `tfsdk:"recipients"`
	Type           types.String                                                          `tfsdk:"type"`
}

type ResponseItemOrganizationsGetOrganizationAlertsProfilesAlertCondition struct {
	BitRateBps types.Int64  `tfsdk:"bit_rate_bps"`
	Duration   types.Int64  `tfsdk:"duration"`
	Interface  types.String `tfsdk:"interface"`
	Window     types.Int64  `tfsdk:"window"`
}

type ResponseItemOrganizationsGetOrganizationAlertsProfilesRecipients struct {
	Emails        types.List `tfsdk:"emails"`
	HTTPServerIDs types.List `tfsdk:"http_server_ids"`
}

// ToBody
func ResponseOrganizationsGetOrganizationAlertsProfilesItemsToBody(state OrganizationsAlertsProfiles, response *merakigosdk.ResponseOrganizationsGetOrganizationAlertsProfiles) OrganizationsAlertsProfiles {
	var items []ResponseItemOrganizationsGetOrganizationAlertsProfiles
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationAlertsProfiles{
			AlertCondition: func() *ResponseItemOrganizationsGetOrganizationAlertsProfilesAlertCondition {
				if item.AlertCondition != nil {
					return &ResponseItemOrganizationsGetOrganizationAlertsProfilesAlertCondition{
						BitRateBps: func() types.Int64 {
							if item.AlertCondition.BitRateBps != nil {
								return types.Int64Value(int64(*item.AlertCondition.BitRateBps))
							}
							return types.Int64{}
						}(),
						Duration: func() types.Int64 {
							if item.AlertCondition.Duration != nil {
								return types.Int64Value(int64(*item.AlertCondition.Duration))
							}
							return types.Int64{}
						}(),
						Interface: types.StringValue(item.AlertCondition.Interface),
						Window: func() types.Int64 {
							if item.AlertCondition.Window != nil {
								return types.Int64Value(int64(*item.AlertCondition.Window))
							}
							return types.Int64{}
						}(),
					}
				}
				return nil
			}(),
			Description: types.StringValue(item.Description),
			Enabled: func() types.Bool {
				if item.Enabled != nil {
					return types.BoolValue(*item.Enabled)
				}
				return types.Bool{}
			}(),
			ID:          types.StringValue(item.ID),
			NetworkTags: StringSliceToList(item.NetworkTags),
			Recipients: func() *ResponseItemOrganizationsGetOrganizationAlertsProfilesRecipients {
				if item.Recipients != nil {
					return &ResponseItemOrganizationsGetOrganizationAlertsProfilesRecipients{
						Emails:        StringSliceToList(item.Recipients.Emails),
						HTTPServerIDs: StringSliceToList(item.Recipients.HTTPServerIDs),
					}
				}
				return nil
			}(),
			Type: types.StringValue(item.Type),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
