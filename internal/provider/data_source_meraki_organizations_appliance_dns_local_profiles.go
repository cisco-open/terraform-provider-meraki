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
	_ datasource.DataSource              = &OrganizationsApplianceDNSLocalProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsApplianceDNSLocalProfilesDataSource{}
)

func NewOrganizationsApplianceDNSLocalProfilesDataSource() datasource.DataSource {
	return &OrganizationsApplianceDNSLocalProfilesDataSource{}
}

type OrganizationsApplianceDNSLocalProfilesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsApplianceDNSLocalProfilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsApplianceDNSLocalProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_appliance_dns_local_profiles"
}

func (d *OrganizationsApplianceDNSLocalProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				MarkdownDescription: `Array of ResponseApplianceGetOrganizationApplianceDnsLocalProfiles`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"name": schema.StringAttribute{
							MarkdownDescription: `Name of profile`,
							Computed:            true,
						},
						"profile_id": schema.StringAttribute{
							MarkdownDescription: `Profile ID`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsApplianceDNSLocalProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsApplianceDNSLocalProfiles OrganizationsApplianceDNSLocalProfiles
	diags := req.Config.Get(ctx, &organizationsApplianceDNSLocalProfiles)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationApplianceDNSLocalProfiles")
		vvOrganizationID := organizationsApplianceDNSLocalProfiles.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationApplianceDNSLocalProfilesQueryParams{}

		queryParams1.ProfileIDs = elementsToStrings(ctx, organizationsApplianceDNSLocalProfiles.ProfileIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetOrganizationApplianceDNSLocalProfiles(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationApplianceDNSLocalProfiles",
				err.Error(),
			)
			return
		}

		organizationsApplianceDNSLocalProfiles = ResponseApplianceGetOrganizationApplianceDNSLocalProfilesItemsToBody(organizationsApplianceDNSLocalProfiles, response1)
		diags = resp.State.Set(ctx, &organizationsApplianceDNSLocalProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsApplianceDNSLocalProfiles struct {
	OrganizationID types.String                                                     `tfsdk:"organization_id"`
	ProfileIDs     types.List                                                       `tfsdk:"profile_ids"`
	Items          *[]ResponseItemApplianceGetOrganizationApplianceDnsLocalProfiles `tfsdk:"items"`
}

type ResponseItemApplianceGetOrganizationApplianceDnsLocalProfiles struct {
	Name      types.String `tfsdk:"name"`
	ProfileID types.String `tfsdk:"profile_id"`
}

// ToBody
func ResponseApplianceGetOrganizationApplianceDNSLocalProfilesItemsToBody(state OrganizationsApplianceDNSLocalProfiles, response *merakigosdk.ResponseApplianceGetOrganizationApplianceDNSLocalProfiles) OrganizationsApplianceDNSLocalProfiles {
	var items []ResponseItemApplianceGetOrganizationApplianceDnsLocalProfiles
	for _, item := range *response {
		itemState := ResponseItemApplianceGetOrganizationApplianceDnsLocalProfiles{
			Name:      types.StringValue(item.Name),
			ProfileID: types.StringValue(item.ProfileID),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
