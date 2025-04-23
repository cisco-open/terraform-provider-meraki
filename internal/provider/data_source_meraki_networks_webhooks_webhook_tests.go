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
	_ datasource.DataSource              = &NetworksWebhooksWebhookTestsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWebhooksWebhookTestsDataSource{}
)

func NewNetworksWebhooksWebhookTestsDataSource() datasource.DataSource {
	return &NetworksWebhooksWebhookTestsDataSource{}
}

type NetworksWebhooksWebhookTestsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWebhooksWebhookTestsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWebhooksWebhookTestsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_webhooks_webhook_tests"
}

func (d *NetworksWebhooksWebhookTestsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"webhook_test_id": schema.StringAttribute{
				MarkdownDescription: `webhookTestId path parameter. Webhook test ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"id": schema.StringAttribute{
						MarkdownDescription: `Webhook delivery identifier`,
						Computed:            true,
					},
					"status": schema.StringAttribute{
						MarkdownDescription: `Current status of the webhook delivery`,
						Computed:            true,
					},
					"url": schema.StringAttribute{
						MarkdownDescription: `URL where the webhook was delivered`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *NetworksWebhooksWebhookTestsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWebhooksWebhookTests NetworksWebhooksWebhookTests
	diags := req.Config.Get(ctx, &networksWebhooksWebhookTests)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWebhooksWebhookTest")
		vvNetworkID := networksWebhooksWebhookTests.NetworkID.ValueString()
		vvWebhookTestID := networksWebhooksWebhookTests.WebhookTestID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkWebhooksWebhookTest(vvNetworkID, vvWebhookTestID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWebhooksWebhookTest",
				err.Error(),
			)
			return
		}

		networksWebhooksWebhookTests = ResponseNetworksGetNetworkWebhooksWebhookTestItemToBody(networksWebhooksWebhookTests, response1)
		diags = resp.State.Set(ctx, &networksWebhooksWebhookTests)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWebhooksWebhookTests struct {
	NetworkID     types.String                                   `tfsdk:"network_id"`
	WebhookTestID types.String                                   `tfsdk:"webhook_test_id"`
	Item          *ResponseNetworksGetNetworkWebhooksWebhookTest `tfsdk:"item"`
}

type ResponseNetworksGetNetworkWebhooksWebhookTest struct {
	ID     types.String `tfsdk:"id"`
	Status types.String `tfsdk:"status"`
	URL    types.String `tfsdk:"url"`
}

// ToBody
func ResponseNetworksGetNetworkWebhooksWebhookTestItemToBody(state NetworksWebhooksWebhookTests, response *merakigosdk.ResponseNetworksGetNetworkWebhooksWebhookTest) NetworksWebhooksWebhookTests {
	itemState := ResponseNetworksGetNetworkWebhooksWebhookTest{
		ID:     types.StringValue(response.ID),
		Status: types.StringValue(response.Status),
		URL:    types.StringValue(response.URL),
	}
	state.Item = &itemState
	return state
}
