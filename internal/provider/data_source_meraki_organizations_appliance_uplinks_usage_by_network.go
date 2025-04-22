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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsApplianceUplinksUsageByNetworkDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsApplianceUplinksUsageByNetworkDataSource{}
)

func NewOrganizationsApplianceUplinksUsageByNetworkDataSource() datasource.DataSource {
	return &OrganizationsApplianceUplinksUsageByNetworkDataSource{}
}

type OrganizationsApplianceUplinksUsageByNetworkDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsApplianceUplinksUsageByNetworkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsApplianceUplinksUsageByNetworkDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_appliance_uplinks_usage_by_network"
}

func (d *OrganizationsApplianceUplinksUsageByNetworkDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 365 days from today.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 14 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 days. The default is 1 day.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseApplianceGetOrganizationApplianceUplinksUsageByNetwork`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"by_uplink": schema.SetNestedAttribute{
							MarkdownDescription: `Uplink usage`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"interface": schema.StringAttribute{
										MarkdownDescription: `Uplink name`,
										Computed:            true,
									},
									"received": schema.Int64Attribute{
										MarkdownDescription: `Bytes received`,
										Computed:            true,
									},
									"sent": schema.Int64Attribute{
										MarkdownDescription: `Bytes sent`,
										Computed:            true,
									},
									"serial": schema.StringAttribute{
										MarkdownDescription: `Uplink serial`,
										Computed:            true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Network name`,
							Computed:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: `Network identifier`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsApplianceUplinksUsageByNetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsApplianceUplinksUsageByNetwork OrganizationsApplianceUplinksUsageByNetwork
	diags := req.Config.Get(ctx, &organizationsApplianceUplinksUsageByNetwork)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationApplianceUplinksUsageByNetwork")
		vvOrganizationID := organizationsApplianceUplinksUsageByNetwork.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationApplianceUplinksUsageByNetworkQueryParams{}

		queryParams1.T0 = organizationsApplianceUplinksUsageByNetwork.T0.ValueString()
		queryParams1.T1 = organizationsApplianceUplinksUsageByNetwork.T1.ValueString()
		queryParams1.Timespan = organizationsApplianceUplinksUsageByNetwork.Timespan.ValueFloat64()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetOrganizationApplianceUplinksUsageByNetwork(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationApplianceUplinksUsageByNetwork",
				err.Error(),
			)
			return
		}

		organizationsApplianceUplinksUsageByNetwork = ResponseApplianceGetOrganizationApplianceUplinksUsageByNetworkItemsToBody(organizationsApplianceUplinksUsageByNetwork, response1)
		diags = resp.State.Set(ctx, &organizationsApplianceUplinksUsageByNetwork)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsApplianceUplinksUsageByNetwork struct {
	OrganizationID types.String                                                          `tfsdk:"organization_id"`
	T0             types.String                                                          `tfsdk:"t0"`
	T1             types.String                                                          `tfsdk:"t1"`
	Timespan       types.Float64                                                         `tfsdk:"timespan"`
	Items          *[]ResponseItemApplianceGetOrganizationApplianceUplinksUsageByNetwork `tfsdk:"items"`
}

type ResponseItemApplianceGetOrganizationApplianceUplinksUsageByNetwork struct {
	ByUplink  *[]ResponseItemApplianceGetOrganizationApplianceUplinksUsageByNetworkByUplink `tfsdk:"by_uplink"`
	Name      types.String                                                                  `tfsdk:"name"`
	NetworkID types.String                                                                  `tfsdk:"network_id"`
}

type ResponseItemApplianceGetOrganizationApplianceUplinksUsageByNetworkByUplink struct {
	Interface types.String `tfsdk:"interface"`
	Received  types.Int64  `tfsdk:"received"`
	Sent      types.Int64  `tfsdk:"sent"`
	Serial    types.String `tfsdk:"serial"`
}

// ToBody
func ResponseApplianceGetOrganizationApplianceUplinksUsageByNetworkItemsToBody(state OrganizationsApplianceUplinksUsageByNetwork, response *merakigosdk.ResponseApplianceGetOrganizationApplianceUplinksUsageByNetwork) OrganizationsApplianceUplinksUsageByNetwork {
	var items []ResponseItemApplianceGetOrganizationApplianceUplinksUsageByNetwork
	for _, item := range *response {
		itemState := ResponseItemApplianceGetOrganizationApplianceUplinksUsageByNetwork{
			ByUplink: func() *[]ResponseItemApplianceGetOrganizationApplianceUplinksUsageByNetworkByUplink {
				if item.ByUplink != nil {
					result := make([]ResponseItemApplianceGetOrganizationApplianceUplinksUsageByNetworkByUplink, len(*item.ByUplink))
					for i, byUplink := range *item.ByUplink {
						result[i] = ResponseItemApplianceGetOrganizationApplianceUplinksUsageByNetworkByUplink{
							Interface: types.StringValue(byUplink.Interface),
							Received: func() types.Int64 {
								if byUplink.Received != nil {
									return types.Int64Value(int64(*byUplink.Received))
								}
								return types.Int64{}
							}(),
							Sent: func() types.Int64 {
								if byUplink.Sent != nil {
									return types.Int64Value(int64(*byUplink.Sent))
								}
								return types.Int64{}
							}(),
							Serial: types.StringValue(byUplink.Serial),
						}
					}
					return &result
				}
				return nil
			}(),
			Name:      types.StringValue(item.Name),
			NetworkID: types.StringValue(item.NetworkID),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
