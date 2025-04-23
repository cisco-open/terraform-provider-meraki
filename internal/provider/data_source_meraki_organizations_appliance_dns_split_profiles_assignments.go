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
	_ datasource.DataSource              = &OrganizationsApplianceDNSSplitProfilesAssignmentsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsApplianceDNSSplitProfilesAssignmentsDataSource{}
)

func NewOrganizationsApplianceDNSSplitProfilesAssignmentsDataSource() datasource.DataSource {
	return &OrganizationsApplianceDNSSplitProfilesAssignmentsDataSource{}
}

type OrganizationsApplianceDNSSplitProfilesAssignmentsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsApplianceDNSSplitProfilesAssignmentsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsApplianceDNSSplitProfilesAssignmentsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_appliance_dns_split_profiles_assignments"
}

func (d *OrganizationsApplianceDNSSplitProfilesAssignmentsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter the results by network IDs`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"profile_ids": schema.ListAttribute{
				MarkdownDescription: `profileIds query parameter. Optional parameter to filter the results by profile IDs`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"items": schema.ListNestedAttribute{
						MarkdownDescription: `List of split DNS profile assignment`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"assignment_id": schema.StringAttribute{
									MarkdownDescription: `ID of the assignment`,
									Computed:            true,
								},
								"network": schema.SingleNestedAttribute{
									MarkdownDescription: `The network attached to the profile`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `ID of the network`,
											Computed:            true,
										},
									},
								},
								"profile": schema.SingleNestedAttribute{
									MarkdownDescription: `The profile the network is attached to`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `ID of the profile`,
											Computed:            true,
										},
									},
								},
							},
						},
					},
					"meta": schema.SingleNestedAttribute{
						MarkdownDescription: `Metadata relevant to the paginated dataset`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"counts": schema.SingleNestedAttribute{
								MarkdownDescription: `Counts relating to the paginated dataset`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"items": schema.SingleNestedAttribute{
										MarkdownDescription: `Counts relating to the paginated items`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"remaining": schema.Int64Attribute{
												MarkdownDescription: `The number of items in the dataset that are available on subsequent pages`,
												Computed:            true,
											},
											"total": schema.Int64Attribute{
												MarkdownDescription: `The total number of items in the dataset`,
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
		},
	}
}

func (d *OrganizationsApplianceDNSSplitProfilesAssignmentsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsApplianceDNSSplitProfilesAssignments OrganizationsApplianceDNSSplitProfilesAssignments
	diags := req.Config.Get(ctx, &organizationsApplianceDNSSplitProfilesAssignments)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationApplianceDNSSplitProfilesAssignments")
		vvOrganizationID := organizationsApplianceDNSSplitProfilesAssignments.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationApplianceDNSSplitProfilesAssignmentsQueryParams{}

		queryParams1.ProfileIDs = elementsToStrings(ctx, organizationsApplianceDNSSplitProfilesAssignments.ProfileIDs)
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsApplianceDNSSplitProfilesAssignments.NetworkIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetOrganizationApplianceDNSSplitProfilesAssignments(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationApplianceDNSSplitProfilesAssignments",
				err.Error(),
			)
			return
		}

		organizationsApplianceDNSSplitProfilesAssignments = ResponseApplianceGetOrganizationApplianceDNSSplitProfilesAssignmentsItemToBody(organizationsApplianceDNSSplitProfilesAssignments, response1)
		diags = resp.State.Set(ctx, &organizationsApplianceDNSSplitProfilesAssignments)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsApplianceDNSSplitProfilesAssignments struct {
	OrganizationID types.String                                                          `tfsdk:"organization_id"`
	ProfileIDs     types.List                                                            `tfsdk:"profile_ids"`
	NetworkIDs     types.List                                                            `tfsdk:"network_ids"`
	Item           *ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignments `tfsdk:"item"`
}

type ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignments struct {
	Items *[]ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsItems `tfsdk:"items"`
	Meta  *ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsMeta    `tfsdk:"meta"`
}

type ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsItems struct {
	AssignmentID types.String                                                                      `tfsdk:"assignment_id"`
	Network      *ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsItemsNetwork `tfsdk:"network"`
	Profile      *ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsItemsProfile `tfsdk:"profile"`
}

type ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsItemsNetwork struct {
	ID types.String `tfsdk:"id"`
}

type ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsItemsProfile struct {
	ID types.String `tfsdk:"id"`
}

type ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsMeta struct {
	Counts *ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsMetaCounts `tfsdk:"counts"`
}

type ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsMetaCounts struct {
	Items *ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsMetaCountsItems `tfsdk:"items"`
}

type ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseApplianceGetOrganizationApplianceDNSSplitProfilesAssignmentsItemToBody(state OrganizationsApplianceDNSSplitProfilesAssignments, response *merakigosdk.ResponseApplianceGetOrganizationApplianceDNSSplitProfilesAssignments) OrganizationsApplianceDNSSplitProfilesAssignments {
	itemState := ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignments{
		Items: func() *[]ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsItems {
			if response.Items != nil {
				result := make([]ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsItems{
						AssignmentID: types.StringValue(items.AssignmentID),
						Network: func() *ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsItemsNetwork {
							if items.Network != nil {
								return &ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsItemsNetwork{
									ID: types.StringValue(items.Network.ID),
								}
							}
							return nil
						}(),
						Profile: func() *ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsItemsProfile {
							if items.Profile != nil {
								return &ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsItemsProfile{
									ID: types.StringValue(items.Profile.ID),
								}
							}
							return nil
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Meta: func() *ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsMeta {
			if response.Meta != nil {
				return &ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsMeta{
					Counts: func() *ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsMetaCounts{
								Items: func() *ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseApplianceGetOrganizationApplianceDnsSplitProfilesAssignmentsMetaCountsItems{
											Remaining: func() types.Int64 {
												if response.Meta.Counts.Items.Remaining != nil {
													return types.Int64Value(int64(*response.Meta.Counts.Items.Remaining))
												}
												return types.Int64{}
											}(),
											Total: func() types.Int64 {
												if response.Meta.Counts.Items.Total != nil {
													return types.Int64Value(int64(*response.Meta.Counts.Items.Total))
												}
												return types.Int64{}
											}(),
										}
									}
									return nil
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
