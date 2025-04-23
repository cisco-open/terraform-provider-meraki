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
	_ datasource.DataSource              = &NetworksSmTargetGroupsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSmTargetGroupsDataSource{}
)

func NewNetworksSmTargetGroupsDataSource() datasource.DataSource {
	return &NetworksSmTargetGroupsDataSource{}
}

type NetworksSmTargetGroupsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSmTargetGroupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSmTargetGroupsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_target_groups"
}

func (d *NetworksSmTargetGroupsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"target_group_id": schema.StringAttribute{
				MarkdownDescription: `targetGroupId path parameter. Target group ID`,
				Optional:            true,
			},
			"with_details": schema.BoolAttribute{
				MarkdownDescription: `withDetails query parameter. Boolean indicating if the the ids of the devices or users scoped by the target group should be included in the response`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"id": schema.StringAttribute{
						MarkdownDescription: `The ID of this target group.`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of this target group.`,
						Computed:            true,
					},
					"scope": schema.StringAttribute{
						MarkdownDescription: `The scope of the target group.`,
						Computed:            true,
					},
					"tags": schema.ListAttribute{
						MarkdownDescription: `The tags of the target group.`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSmGetNetworkSmTargetGroups`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"id": schema.StringAttribute{
							MarkdownDescription: `The ID of this target group.`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of this target group.`,
							Computed:            true,
						},
						"scope": schema.StringAttribute{
							MarkdownDescription: `The scope of the target group.`,
							Computed:            true,
						},
						"tags": schema.ListAttribute{
							MarkdownDescription: `The tags of the target group.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSmTargetGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSmTargetGroups NetworksSmTargetGroups
	diags := req.Config.Get(ctx, &networksSmTargetGroups)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksSmTargetGroups.NetworkID.IsNull(), !networksSmTargetGroups.WithDetails.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksSmTargetGroups.NetworkID.IsNull(), !networksSmTargetGroups.TargetGroupID.IsNull(), !networksSmTargetGroups.WithDetails.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmTargetGroups")
		vvNetworkID := networksSmTargetGroups.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkSmTargetGroupsQueryParams{}

		queryParams1.WithDetails = networksSmTargetGroups.WithDetails.ValueBool()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sm.GetNetworkSmTargetGroups(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmTargetGroups",
				err.Error(),
			)
			return
		}

		networksSmTargetGroups = ResponseSmGetNetworkSmTargetGroupsItemsToBody(networksSmTargetGroups, response1)
		diags = resp.State.Set(ctx, &networksSmTargetGroups)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmTargetGroup")
		vvNetworkID := networksSmTargetGroups.NetworkID.ValueString()
		vvTargetGroupID := networksSmTargetGroups.TargetGroupID.ValueString()
		queryParams2 := merakigosdk.GetNetworkSmTargetGroupQueryParams{}

		queryParams2.WithDetails = networksSmTargetGroups.WithDetails.ValueBool()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Sm.GetNetworkSmTargetGroup(vvNetworkID, vvTargetGroupID, &queryParams2)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmTargetGroup",
				err.Error(),
			)
			return
		}

		networksSmTargetGroups = ResponseSmGetNetworkSmTargetGroupItemToBody(networksSmTargetGroups, response2)
		diags = resp.State.Set(ctx, &networksSmTargetGroups)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSmTargetGroups struct {
	NetworkID     types.String                              `tfsdk:"network_id"`
	WithDetails   types.Bool                                `tfsdk:"with_details"`
	TargetGroupID types.String                              `tfsdk:"target_group_id"`
	Items         *[]ResponseItemSmGetNetworkSmTargetGroups `tfsdk:"items"`
	Item          *ResponseSmGetNetworkSmTargetGroup        `tfsdk:"item"`
}

type ResponseItemSmGetNetworkSmTargetGroups struct {
	ID    types.String `tfsdk:"id"`
	Name  types.String `tfsdk:"name"`
	Scope types.String `tfsdk:"scope"`
	Tags  types.List   `tfsdk:"tags"`
}

type ResponseSmGetNetworkSmTargetGroup struct {
	ID    types.String `tfsdk:"id"`
	Name  types.String `tfsdk:"name"`
	Scope types.String `tfsdk:"scope"`
	Tags  types.List   `tfsdk:"tags"`
}

// ToBody
func ResponseSmGetNetworkSmTargetGroupsItemsToBody(state NetworksSmTargetGroups, response *merakigosdk.ResponseSmGetNetworkSmTargetGroups) NetworksSmTargetGroups {
	var items []ResponseItemSmGetNetworkSmTargetGroups
	for _, item := range *response {
		itemState := ResponseItemSmGetNetworkSmTargetGroups{
			ID:    types.StringValue(item.ID),
			Name:  types.StringValue(item.Name),
			Scope: types.StringValue(item.Scope),
			Tags:  StringSliceToList(item.Tags),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseSmGetNetworkSmTargetGroupItemToBody(state NetworksSmTargetGroups, response *merakigosdk.ResponseSmGetNetworkSmTargetGroup) NetworksSmTargetGroups {
	itemState := ResponseSmGetNetworkSmTargetGroup{
		ID:    types.StringValue(response.ID),
		Name:  types.StringValue(response.Name),
		Scope: types.StringValue(response.Scope),
		Tags:  StringSliceToList(response.Tags),
	}
	state.Item = &itemState
	return state
}
