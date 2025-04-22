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
	_ datasource.DataSource              = &OrganizationsApplianceDNSLocalRecordsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsApplianceDNSLocalRecordsDataSource{}
)

func NewOrganizationsApplianceDNSLocalRecordsDataSource() datasource.DataSource {
	return &OrganizationsApplianceDNSLocalRecordsDataSource{}
}

type OrganizationsApplianceDNSLocalRecordsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsApplianceDNSLocalRecordsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsApplianceDNSLocalRecordsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_appliance_dns_local_records"
}

func (d *OrganizationsApplianceDNSLocalRecordsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"profile_ids": schema.ListAttribute{
				MarkdownDescription: `profileIds query parameter. Optional parameter to filter the results by profile IDs`,
				Optional:            true,
				ElementType:         types.StringType,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseApplianceGetOrganizationApplianceDnsLocalRecords`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"address": schema.StringAttribute{
							MarkdownDescription: `IP for the DNS record`,
							Computed:            true,
						},
						"hostname": schema.StringAttribute{
							MarkdownDescription: `Hostname for the DNS record`,
							Computed:            true,
						},
						"profile": schema.SingleNestedAttribute{
							MarkdownDescription: `The profile the DNS record is associated with`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `Profile ID`,
									Computed:            true,
								},
							},
						},
						"record_id": schema.StringAttribute{
							MarkdownDescription: `Record ID`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsApplianceDNSLocalRecordsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsApplianceDNSLocalRecords OrganizationsApplianceDNSLocalRecords
	diags := req.Config.Get(ctx, &organizationsApplianceDNSLocalRecords)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationApplianceDNSLocalRecords")
		vvOrganizationID := organizationsApplianceDNSLocalRecords.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationApplianceDNSLocalRecordsQueryParams{}

		queryParams1.ProfileIDs = elementsToStrings(ctx, organizationsApplianceDNSLocalRecords.ProfileIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetOrganizationApplianceDNSLocalRecords(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationApplianceDNSLocalRecords",
				err.Error(),
			)
			return
		}

		organizationsApplianceDNSLocalRecords = ResponseApplianceGetOrganizationApplianceDNSLocalRecordsItemsToBody(organizationsApplianceDNSLocalRecords, response1)
		diags = resp.State.Set(ctx, &organizationsApplianceDNSLocalRecords)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsApplianceDNSLocalRecords struct {
	OrganizationID types.String                                                    `tfsdk:"organization_id"`
	ProfileIDs     types.List                                                      `tfsdk:"profile_ids"`
	Items          *[]ResponseItemApplianceGetOrganizationApplianceDnsLocalRecords `tfsdk:"items"`
}

type ResponseItemApplianceGetOrganizationApplianceDnsLocalRecords struct {
	Address  types.String                                                         `tfsdk:"address"`
	Hostname types.String                                                         `tfsdk:"hostname"`
	Profile  *ResponseItemApplianceGetOrganizationApplianceDnsLocalRecordsProfile `tfsdk:"profile"`
	RecordID types.String                                                         `tfsdk:"record_id"`
}

type ResponseItemApplianceGetOrganizationApplianceDnsLocalRecordsProfile struct {
	ID types.String `tfsdk:"id"`
}

// ToBody
func ResponseApplianceGetOrganizationApplianceDNSLocalRecordsItemsToBody(state OrganizationsApplianceDNSLocalRecords, response *merakigosdk.ResponseApplianceGetOrganizationApplianceDNSLocalRecords) OrganizationsApplianceDNSLocalRecords {
	var items []ResponseItemApplianceGetOrganizationApplianceDnsLocalRecords
	for _, item := range *response {
		itemState := ResponseItemApplianceGetOrganizationApplianceDnsLocalRecords{
			Address:  types.StringValue(item.Address),
			Hostname: types.StringValue(item.Hostname),
			Profile: func() *ResponseItemApplianceGetOrganizationApplianceDnsLocalRecordsProfile {
				if item.Profile != nil {
					return &ResponseItemApplianceGetOrganizationApplianceDnsLocalRecordsProfile{
						ID: types.StringValue(item.Profile.ID),
					}
				}
				return nil
			}(),
			RecordID: types.StringValue(item.RecordID),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
