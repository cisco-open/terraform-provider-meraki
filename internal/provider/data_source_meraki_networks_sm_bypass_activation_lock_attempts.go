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
	_ datasource.DataSource              = &NetworksSmBypassActivationLockAttemptsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSmBypassActivationLockAttemptsDataSource{}
)

func NewNetworksSmBypassActivationLockAttemptsDataSource() datasource.DataSource {
	return &NetworksSmBypassActivationLockAttemptsDataSource{}
}

type NetworksSmBypassActivationLockAttemptsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSmBypassActivationLockAttemptsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSmBypassActivationLockAttemptsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_bypass_activation_lock_attempts"
}

func (d *NetworksSmBypassActivationLockAttemptsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"attempt_id": schema.StringAttribute{
				MarkdownDescription: `attemptId path parameter. Attempt ID`,
				Required:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"data": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"status_2090938209": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"errors": schema.ListAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
									"success": schema.BoolAttribute{
										Computed: true,
									},
								},
							},
							"status_38290139892": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"success": schema.BoolAttribute{
										Computed: true,
									},
								},
							},
						},
					},
					"id": schema.StringAttribute{
						Computed: true,
					},
					"status": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *NetworksSmBypassActivationLockAttemptsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSmBypassActivationLockAttempts NetworksSmBypassActivationLockAttempts
	diags := req.Config.Get(ctx, &networksSmBypassActivationLockAttempts)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmBypassActivationLockAttempt")
		vvNetworkID := networksSmBypassActivationLockAttempts.NetworkID.ValueString()
		vvAttemptID := networksSmBypassActivationLockAttempts.AttemptID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sm.GetNetworkSmBypassActivationLockAttempt(vvNetworkID, vvAttemptID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmBypassActivationLockAttempt",
				err.Error(),
			)
			return
		}

		networksSmBypassActivationLockAttempts = ResponseSmGetNetworkSmBypassActivationLockAttemptItemToBody(networksSmBypassActivationLockAttempts, response1)
		diags = resp.State.Set(ctx, &networksSmBypassActivationLockAttempts)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSmBypassActivationLockAttempts struct {
	NetworkID types.String                                       `tfsdk:"network_id"`
	AttemptID types.String                                       `tfsdk:"attempt_id"`
	Item      *ResponseSmGetNetworkSmBypassActivationLockAttempt `tfsdk:"item"`
}

type ResponseSmGetNetworkSmBypassActivationLockAttempt struct {
	Data   *ResponseSmGetNetworkSmBypassActivationLockAttemptData `tfsdk:"data"`
	ID     types.String                                           `tfsdk:"id"`
	Status types.String                                           `tfsdk:"status"`
}

type ResponseSmGetNetworkSmBypassActivationLockAttemptData struct {
	Status2090938209  *ResponseSmGetNetworkSmBypassActivationLockAttemptData2090938209  `tfsdk:"status_2090938209"`
	Status38290139892 *ResponseSmGetNetworkSmBypassActivationLockAttemptData38290139892 `tfsdk:"status_38290139892"`
}

type ResponseSmGetNetworkSmBypassActivationLockAttemptData2090938209 struct {
	Errors  types.List `tfsdk:"errors"`
	Success types.Bool `tfsdk:"success"`
}

type ResponseSmGetNetworkSmBypassActivationLockAttemptData38290139892 struct {
	Success types.Bool `tfsdk:"success"`
}

// ToBody
func ResponseSmGetNetworkSmBypassActivationLockAttemptItemToBody(state NetworksSmBypassActivationLockAttempts, response *merakigosdk.ResponseSmGetNetworkSmBypassActivationLockAttempt) NetworksSmBypassActivationLockAttempts {
	itemState := ResponseSmGetNetworkSmBypassActivationLockAttempt{
		Data: func() *ResponseSmGetNetworkSmBypassActivationLockAttemptData {
			if response.Data != nil {
				return &ResponseSmGetNetworkSmBypassActivationLockAttemptData{
					Status2090938209: func() *ResponseSmGetNetworkSmBypassActivationLockAttemptData2090938209 {
						if response.Data.Status2090938209 != nil {
							return &ResponseSmGetNetworkSmBypassActivationLockAttemptData2090938209{
								Errors: StringSliceToList(response.Data.Status2090938209.Errors),
								Success: func() types.Bool {
									if response.Data.Status2090938209.Success != nil {
										return types.BoolValue(*response.Data.Status2090938209.Success)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
					Status38290139892: func() *ResponseSmGetNetworkSmBypassActivationLockAttemptData38290139892 {
						if response.Data.Status38290139892 != nil {
							return &ResponseSmGetNetworkSmBypassActivationLockAttemptData38290139892{
								Success: func() types.Bool {
									if response.Data.Status38290139892.Success != nil {
										return types.BoolValue(*response.Data.Status38290139892.Success)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		ID:     types.StringValue(response.ID),
		Status: types.StringValue(response.Status),
	}
	state.Item = &itemState
	return state
}
