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
	_ datasource.DataSource              = &AdministeredLicensingSubscriptionEntitlementsDataSource{}
	_ datasource.DataSourceWithConfigure = &AdministeredLicensingSubscriptionEntitlementsDataSource{}
)

func NewAdministeredLicensingSubscriptionEntitlementsDataSource() datasource.DataSource {
	return &AdministeredLicensingSubscriptionEntitlementsDataSource{}
}

type AdministeredLicensingSubscriptionEntitlementsDataSource struct {
	client *merakigosdk.Client
}

func (d *AdministeredLicensingSubscriptionEntitlementsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *AdministeredLicensingSubscriptionEntitlementsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_administered_licensing_subscription_entitlements"
}

func (d *AdministeredLicensingSubscriptionEntitlementsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"skus": schema.ListAttribute{
				MarkdownDescription: `skus query parameter. Filter to entitlements with the specified SKUs`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"feature_tier": schema.StringAttribute{
						MarkdownDescription: `The feature tier associated with the entitlement (null for add-ons)`,
						Computed:            true,
					},
					"is_add_on": schema.BoolAttribute{
						MarkdownDescription: `Whether or not the entitlement is an add-on`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The user-facing name of the entitlement`,
						Computed:            true,
					},
					"product_class": schema.StringAttribute{
						MarkdownDescription: `The product class associated with the entitlement`,
						Computed:            true,
					},
					"product_type": schema.StringAttribute{
						MarkdownDescription: `The product type of the entitlement`,
						Computed:            true,
					},
					"sku": schema.StringAttribute{
						MarkdownDescription: `The SKU identifier of the entitlement`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *AdministeredLicensingSubscriptionEntitlementsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var administeredLicensingSubscriptionEntitlements AdministeredLicensingSubscriptionEntitlements
	diags := req.Config.Get(ctx, &administeredLicensingSubscriptionEntitlements)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAdministeredLicensingSubscriptionEntitlements")
		queryParams1 := merakigosdk.GetAdministeredLicensingSubscriptionEntitlementsQueryParams{}

		queryParams1.Skus = elementsToStrings(ctx, administeredLicensingSubscriptionEntitlements.Skus)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Licensing.GetAdministeredLicensingSubscriptionEntitlements(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetAdministeredLicensingSubscriptionEntitlements",
				err.Error(),
			)
			return
		}

		administeredLicensingSubscriptionEntitlements = ResponseLicensingGetAdministeredLicensingSubscriptionEntitlementsItemToBody(administeredLicensingSubscriptionEntitlements, response1)
		diags = resp.State.Set(ctx, &administeredLicensingSubscriptionEntitlements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type AdministeredLicensingSubscriptionEntitlements struct {
	Skus types.List                                                         `tfsdk:"skus"`
	Item *ResponseLicensingGetAdministeredLicensingSubscriptionEntitlements `tfsdk:"item"`
}

type ResponseLicensingGetAdministeredLicensingSubscriptionEntitlements struct {
	FeatureTier  types.String `tfsdk:"feature_tier"`
	IsAddOn      types.Bool   `tfsdk:"is_add_on"`
	Name         types.String `tfsdk:"name"`
	ProductClass types.String `tfsdk:"product_class"`
	ProductType  types.String `tfsdk:"product_type"`
	Sku          types.String `tfsdk:"sku"`
}

// ToBody
func ResponseLicensingGetAdministeredLicensingSubscriptionEntitlementsItemToBody(state AdministeredLicensingSubscriptionEntitlements, response *merakigosdk.ResponseLicensingGetAdministeredLicensingSubscriptionEntitlements) AdministeredLicensingSubscriptionEntitlements {
	itemState := ResponseLicensingGetAdministeredLicensingSubscriptionEntitlements{
		FeatureTier: types.StringValue(response.FeatureTier),
		IsAddOn: func() types.Bool {
			if response.IsAddOn != nil {
				return types.BoolValue(*response.IsAddOn)
			}
			return types.Bool{}
		}(),
		Name:         types.StringValue(response.Name),
		ProductClass: types.StringValue(response.ProductClass),
		ProductType:  types.StringValue(response.ProductType),
		Sku:          types.StringValue(response.Sku),
	}
	state.Item = &itemState
	return state
}
