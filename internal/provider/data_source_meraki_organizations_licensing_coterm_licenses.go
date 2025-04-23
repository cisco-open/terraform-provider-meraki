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
	_ datasource.DataSource              = &OrganizationsLicensingCotermLicensesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsLicensingCotermLicensesDataSource{}
)

func NewOrganizationsLicensingCotermLicensesDataSource() datasource.DataSource {
	return &OrganizationsLicensingCotermLicensesDataSource{}
}

type OrganizationsLicensingCotermLicensesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsLicensingCotermLicensesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsLicensingCotermLicensesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_licensing_coterm_licenses"
}

func (d *OrganizationsLicensingCotermLicensesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"expired": schema.BoolAttribute{
				MarkdownDescription: `expired query parameter. Filter for licenses that are expired`,
				Optional:            true,
			},
			"invalidated": schema.BoolAttribute{
				MarkdownDescription: `invalidated query parameter. Filter for licenses that are invalidated`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseLicensingGetOrganizationLicensingCotermLicenses`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"claimed_at": schema.StringAttribute{
							MarkdownDescription: `When the license was claimed into the organization`,
							Computed:            true,
						},
						"counts": schema.SetNestedAttribute{
							MarkdownDescription: `The counts of the license by model type`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"count": schema.Int64Attribute{
										MarkdownDescription: `The number of counts the license contains of this model`,
										Computed:            true,
									},
									"model": schema.StringAttribute{
										MarkdownDescription: `The license model type`,
										Computed:            true,
									},
								},
							},
						},
						"duration": schema.Int64Attribute{
							MarkdownDescription: `The duration (term length) of the license, measured in days`,
							Computed:            true,
						},
						"editions": schema.SetNestedAttribute{
							MarkdownDescription: `The editions of the license for each relevant product type`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"edition": schema.StringAttribute{
										MarkdownDescription: `The name of the license edition`,
										Computed:            true,
									},
									"product_type": schema.StringAttribute{
										MarkdownDescription: `The product type of the license edition`,
										Computed:            true,
									},
								},
							},
						},
						"expired": schema.BoolAttribute{
							MarkdownDescription: `Flag to indicate if the license is expired`,
							Computed:            true,
						},
						"invalidated": schema.BoolAttribute{
							MarkdownDescription: `Flag to indicated that the license is invalidated`,
							Computed:            true,
						},
						"invalidated_at": schema.StringAttribute{
							MarkdownDescription: `When the license was invalidated. Will be null for active licenses`,
							Computed:            true,
						},
						"key": schema.StringAttribute{
							MarkdownDescription: `The key of the license`,
							Computed:            true,
						},
						"mode": schema.StringAttribute{
							MarkdownDescription: `The operation mode of the license when it was claimed`,
							Computed:            true,
						},
						"organization_id": schema.StringAttribute{
							MarkdownDescription: `The ID of the organization that the license is claimed in`,
							Computed:            true,
						},
						"started_at": schema.StringAttribute{
							MarkdownDescription: `When the license's term began (approximately the date when the license was created)`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsLicensingCotermLicensesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsLicensingCotermLicenses OrganizationsLicensingCotermLicenses
	diags := req.Config.Get(ctx, &organizationsLicensingCotermLicenses)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationLicensingCotermLicenses")
		vvOrganizationID := organizationsLicensingCotermLicenses.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationLicensingCotermLicensesQueryParams{}

		queryParams1.PerPage = int(organizationsLicensingCotermLicenses.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsLicensingCotermLicenses.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsLicensingCotermLicenses.EndingBefore.ValueString()
		queryParams1.Invalidated = organizationsLicensingCotermLicenses.Invalidated.ValueBool()
		queryParams1.Expired = organizationsLicensingCotermLicenses.Expired.ValueBool()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Licensing.GetOrganizationLicensingCotermLicenses(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationLicensingCotermLicenses",
				err.Error(),
			)
			return
		}

		organizationsLicensingCotermLicenses = ResponseLicensingGetOrganizationLicensingCotermLicensesItemsToBody(organizationsLicensingCotermLicenses, response1)
		diags = resp.State.Set(ctx, &organizationsLicensingCotermLicenses)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsLicensingCotermLicenses struct {
	OrganizationID types.String                                                   `tfsdk:"organization_id"`
	PerPage        types.Int64                                                    `tfsdk:"per_page"`
	StartingAfter  types.String                                                   `tfsdk:"starting_after"`
	EndingBefore   types.String                                                   `tfsdk:"ending_before"`
	Invalidated    types.Bool                                                     `tfsdk:"invalidated"`
	Expired        types.Bool                                                     `tfsdk:"expired"`
	Items          *[]ResponseItemLicensingGetOrganizationLicensingCotermLicenses `tfsdk:"items"`
}

type ResponseItemLicensingGetOrganizationLicensingCotermLicenses struct {
	ClaimedAt      types.String                                                           `tfsdk:"claimed_at"`
	Counts         *[]ResponseItemLicensingGetOrganizationLicensingCotermLicensesCounts   `tfsdk:"counts"`
	Duration       types.Int64                                                            `tfsdk:"duration"`
	Editions       *[]ResponseItemLicensingGetOrganizationLicensingCotermLicensesEditions `tfsdk:"editions"`
	Expired        types.Bool                                                             `tfsdk:"expired"`
	Invalidated    types.Bool                                                             `tfsdk:"invalidated"`
	InvalidatedAt  types.String                                                           `tfsdk:"invalidated_at"`
	Key            types.String                                                           `tfsdk:"key"`
	Mode           types.String                                                           `tfsdk:"mode"`
	OrganizationID types.String                                                           `tfsdk:"organization_id"`
	StartedAt      types.String                                                           `tfsdk:"started_at"`
}

type ResponseItemLicensingGetOrganizationLicensingCotermLicensesCounts struct {
	Count types.Int64  `tfsdk:"count"`
	Model types.String `tfsdk:"model"`
}

type ResponseItemLicensingGetOrganizationLicensingCotermLicensesEditions struct {
	Edition     types.String `tfsdk:"edition"`
	ProductType types.String `tfsdk:"product_type"`
}

// ToBody
func ResponseLicensingGetOrganizationLicensingCotermLicensesItemsToBody(state OrganizationsLicensingCotermLicenses, response *merakigosdk.ResponseLicensingGetOrganizationLicensingCotermLicenses) OrganizationsLicensingCotermLicenses {
	var items []ResponseItemLicensingGetOrganizationLicensingCotermLicenses
	for _, item := range *response {
		itemState := ResponseItemLicensingGetOrganizationLicensingCotermLicenses{
			ClaimedAt: types.StringValue(item.ClaimedAt),
			Counts: func() *[]ResponseItemLicensingGetOrganizationLicensingCotermLicensesCounts {
				if item.Counts != nil {
					result := make([]ResponseItemLicensingGetOrganizationLicensingCotermLicensesCounts, len(*item.Counts))
					for i, counts := range *item.Counts {
						result[i] = ResponseItemLicensingGetOrganizationLicensingCotermLicensesCounts{
							Count: func() types.Int64 {
								if counts.Count != nil {
									return types.Int64Value(int64(*counts.Count))
								}
								return types.Int64{}
							}(),
							Model: types.StringValue(counts.Model),
						}
					}
					return &result
				}
				return nil
			}(),
			Duration: func() types.Int64 {
				if item.Duration != nil {
					return types.Int64Value(int64(*item.Duration))
				}
				return types.Int64{}
			}(),
			Editions: func() *[]ResponseItemLicensingGetOrganizationLicensingCotermLicensesEditions {
				if item.Editions != nil {
					result := make([]ResponseItemLicensingGetOrganizationLicensingCotermLicensesEditions, len(*item.Editions))
					for i, editions := range *item.Editions {
						result[i] = ResponseItemLicensingGetOrganizationLicensingCotermLicensesEditions{
							Edition:     types.StringValue(editions.Edition),
							ProductType: types.StringValue(editions.ProductType),
						}
					}
					return &result
				}
				return nil
			}(),
			Expired: func() types.Bool {
				if item.Expired != nil {
					return types.BoolValue(*item.Expired)
				}
				return types.Bool{}
			}(),
			Invalidated: func() types.Bool {
				if item.Invalidated != nil {
					return types.BoolValue(*item.Invalidated)
				}
				return types.Bool{}
			}(),
			InvalidatedAt:  types.StringValue(item.InvalidatedAt),
			Key:            types.StringValue(item.Key),
			Mode:           types.StringValue(item.Mode),
			OrganizationID: types.StringValue(item.OrganizationID),
			StartedAt:      types.StringValue(item.StartedAt),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
