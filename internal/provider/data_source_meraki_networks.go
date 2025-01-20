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
	_ datasource.DataSource              = &NetworksDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksDataSource{}
)

func NewNetworksDataSource() datasource.DataSource {
	return &NetworksDataSource{}
}

type NetworksDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks"
}

func (d *NetworksDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"config_template_id": schema.StringAttribute{
				MarkdownDescription: `configTemplateId query parameter. An optional parameter that is the ID of a config template. Will return all networks bound to that template.`,
				Optional:            true,
			},
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"is_bound_to_config_template": schema.BoolAttribute{
				MarkdownDescription: `isBoundToConfigTemplate query parameter. An optional parameter to filter config template bound networks. If configTemplateId is set, this cannot be false.`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 100000. Default is 1000.`,
				Optional:            true,
			},
			"product_types": schema.ListAttribute{
				MarkdownDescription: `productTypes query parameter. An optional parameter to filter networks by product type. Results will have at least one of the included product types.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"tags": schema.ListAttribute{
				MarkdownDescription: `tags query parameter. An optional parameter to filter networks by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below).`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"tags_filter_type": schema.StringAttribute{
				MarkdownDescription: `tagsFilterType query parameter. An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"enrollment_string": schema.StringAttribute{
						MarkdownDescription: `Enrollment string for the network`,
						Computed:            true,
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `Network ID`,
						Computed:            true,
					},
					"is_bound_to_config_template": schema.BoolAttribute{
						MarkdownDescription: `If the network is bound to a config template`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Network name`,
						Computed:            true,
					},
					"notes": schema.StringAttribute{
						MarkdownDescription: `Notes for the network`,
						Computed:            true,
					},
					"organization_id": schema.StringAttribute{
						MarkdownDescription: `Organization ID`,
						Computed:            true,
					},
					"product_types": schema.ListAttribute{
						MarkdownDescription: `List of the product types that the network supports`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"tags": schema.ListAttribute{
						MarkdownDescription: `Network tags`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"time_zone": schema.StringAttribute{
						MarkdownDescription: `Timezone of the network`,
						Computed:            true,
					},
					"url": schema.StringAttribute{
						MarkdownDescription: `URL to the network Dashboard UI`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationNetworks`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"enrollment_string": schema.StringAttribute{
							MarkdownDescription: `Enrollment string for the network`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `Network ID`,
							Computed:            true,
						},
						"is_bound_to_config_template": schema.BoolAttribute{
							MarkdownDescription: `If the network is bound to a config template`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Network name`,
							Computed:            true,
						},
						"notes": schema.StringAttribute{
							MarkdownDescription: `Notes for the network`,
							Computed:            true,
						},
						"organization_id": schema.StringAttribute{
							MarkdownDescription: `Organization ID`,
							Computed:            true,
						},
						"product_types": schema.ListAttribute{
							MarkdownDescription: `List of the product types that the network supports`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"tags": schema.ListAttribute{
							MarkdownDescription: `Network tags`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"time_zone": schema.StringAttribute{
							MarkdownDescription: `Timezone of the network`,
							Computed:            true,
						},
						"url": schema.StringAttribute{
							MarkdownDescription: `URL to the network Dashboard UI`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networks Networks
	diags := req.Config.Get(ctx, &networks)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networks.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networks.OrganizationID.IsNull(), !networks.ConfigTemplateID.IsNull(), !networks.IsBoundToConfigTemplate.IsNull(), !networks.Tags.IsNull(), !networks.TagsFilterType.IsNull(), !networks.ProductTypes.IsNull(), !networks.PerPage.IsNull(), !networks.StartingAfter.IsNull(), !networks.EndingBefore.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetwork")
		vvNetworkID := networks.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetwork(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetwork",
				err.Error(),
			)
			return
		}

		networks = ResponseNetworksGetNetworkItemToBody(networks, response1)
		diags = resp.State.Set(ctx, &networks)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationNetworks")
		vvOrganizationID := networks.OrganizationID.ValueString()
		queryParams2 := merakigosdk.GetOrganizationNetworksQueryParams{}

		queryParams2.ConfigTemplateID = networks.ConfigTemplateID.ValueString()
		queryParams2.IsBoundToConfigTemplate = networks.IsBoundToConfigTemplate.ValueBool()
		queryParams2.Tags = elementsToStrings(ctx, networks.Tags)
		queryParams2.TagsFilterType = networks.TagsFilterType.ValueString()
		queryParams2.ProductTypes = elementsToStrings(ctx, networks.ProductTypes)
		queryParams2.PerPage = int(networks.PerPage.ValueInt64())
		queryParams2.StartingAfter = networks.StartingAfter.ValueString()
		queryParams2.EndingBefore = networks.EndingBefore.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Organizations.GetOrganizationNetworks(vvOrganizationID, &queryParams2)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationNetworks",
				err.Error(),
			)
			return
		}

		networks = ResponseNetworksGetOrganizationNetworksItemsToBody(networks, response2)
		diags = resp.State.Set(ctx, &networks)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type Networks struct {
	NetworkID               types.String                                        `tfsdk:"network_id"`
	OrganizationID          types.String                                        `tfsdk:"organization_id"`
	ConfigTemplateID        types.String                                        `tfsdk:"config_template_id"`
	IsBoundToConfigTemplate types.Bool                                          `tfsdk:"is_bound_to_config_template"`
	Tags                    types.List                                          `tfsdk:"tags"`
	TagsFilterType          types.String                                        `tfsdk:"tags_filter_type"`
	ProductTypes            types.List                                          `tfsdk:"product_types"`
	PerPage                 types.Int64                                         `tfsdk:"per_page"`
	StartingAfter           types.String                                        `tfsdk:"starting_after"`
	EndingBefore            types.String                                        `tfsdk:"ending_before"`
	Items                   *[]ResponseItemOrganizationsGetOrganizationNetworks `tfsdk:"items"`
	Item                    *ResponseNetworksGetNetwork                         `tfsdk:"item"`
}

type ResponseItemOrganizationsGetOrganizationNetworks struct {
	EnrollmentString        types.String `tfsdk:"enrollment_string"`
	ID                      types.String `tfsdk:"id"`
	IsBoundToConfigTemplate types.Bool   `tfsdk:"is_bound_to_config_template"`
	Name                    types.String `tfsdk:"name"`
	Notes                   types.String `tfsdk:"notes"`
	OrganizationID          types.String `tfsdk:"organization_id"`
	ProductTypes            types.List   `tfsdk:"product_types"`
	Tags                    types.List   `tfsdk:"tags"`
	TimeZone                types.String `tfsdk:"time_zone"`
	URL                     types.String `tfsdk:"url"`
}

type ResponseNetworksGetNetwork struct {
	EnrollmentString        types.String `tfsdk:"enrollment_string"`
	ID                      types.String `tfsdk:"id"`
	IsBoundToConfigTemplate types.Bool   `tfsdk:"is_bound_to_config_template"`
	Name                    types.String `tfsdk:"name"`
	Notes                   types.String `tfsdk:"notes"`
	OrganizationID          types.String `tfsdk:"organization_id"`
	ProductTypes            types.List   `tfsdk:"product_types"`
	Tags                    types.List   `tfsdk:"tags"`
	TimeZone                types.String `tfsdk:"time_zone"`
	URL                     types.String `tfsdk:"url"`
}

// ToBody
func ResponseNetworksGetOrganizationNetworksItemsToBody(state Networks, response *merakigosdk.ResponseOrganizationsGetOrganizationNetworks) Networks {
	var items []ResponseItemOrganizationsGetOrganizationNetworks
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationNetworks{
			EnrollmentString: types.StringValue(item.EnrollmentString),
			ID:               types.StringValue(item.ID),
			IsBoundToConfigTemplate: func() types.Bool {
				if item.IsBoundToConfigTemplate != nil {
					return types.BoolValue(*item.IsBoundToConfigTemplate)
				}
				return types.Bool{}
			}(),
			Name:           types.StringValue(item.Name),
			Notes:          types.StringValue(item.Notes),
			OrganizationID: types.StringValue(item.OrganizationID),
			ProductTypes:   StringSliceToList(item.ProductTypes),
			Tags:           StringSliceToList(item.Tags),
			TimeZone:       types.StringValue(item.TimeZone),
			URL:            types.StringValue(item.URL),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseNetworksGetNetworkItemToBody(state Networks, response *merakigosdk.ResponseNetworksGetNetwork) Networks {
	itemState := ResponseNetworksGetNetwork{
		EnrollmentString: types.StringValue(response.EnrollmentString),
		ID:               types.StringValue(response.ID),
		IsBoundToConfigTemplate: func() types.Bool {
			if response.IsBoundToConfigTemplate != nil {
				return types.BoolValue(*response.IsBoundToConfigTemplate)
			}
			return types.Bool{}
		}(),
		Name:           types.StringValue(response.Name),
		Notes:          types.StringValue(response.Notes),
		OrganizationID: types.StringValue(response.OrganizationID),
		ProductTypes:   StringSliceToList(response.ProductTypes),
		Tags:           StringSliceToList(response.Tags),
		TimeZone:       types.StringValue(response.TimeZone),
		URL:            types.StringValue(response.URL),
	}
	state.Item = &itemState
	return state
}
