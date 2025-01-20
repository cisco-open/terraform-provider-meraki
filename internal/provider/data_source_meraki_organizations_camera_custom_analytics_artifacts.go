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
	_ datasource.DataSource              = &OrganizationsCameraCustomAnalyticsArtifactsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsCameraCustomAnalyticsArtifactsDataSource{}
)

func NewOrganizationsCameraCustomAnalyticsArtifactsDataSource() datasource.DataSource {
	return &OrganizationsCameraCustomAnalyticsArtifactsDataSource{}
}

type OrganizationsCameraCustomAnalyticsArtifactsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsCameraCustomAnalyticsArtifactsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsCameraCustomAnalyticsArtifactsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_camera_custom_analytics_artifacts"
}

func (d *OrganizationsCameraCustomAnalyticsArtifactsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"artifact_id": schema.StringAttribute{
				MarkdownDescription: `artifactId path parameter. Artifact ID`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"artifact_id": schema.StringAttribute{
						MarkdownDescription: `Custom analytics artifact ID`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Custom analytics artifact name`,
						Computed:            true,
					},
					"organization_id": schema.StringAttribute{
						MarkdownDescription: `Organization ID`,
						Computed:            true,
					},
					"status": schema.SingleNestedAttribute{
						MarkdownDescription: `Custom analytics artifact status`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"message": schema.StringAttribute{
								MarkdownDescription: `Status message`,
								Computed:            true,
							},
							"type": schema.StringAttribute{
								MarkdownDescription: `Status type`,
								Computed:            true,
							},
						},
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseCameraGetOrganizationCameraCustomAnalyticsArtifacts`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"artifact_id": schema.StringAttribute{
							MarkdownDescription: `Custom analytics artifact ID`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Custom analytics artifact name`,
							Computed:            true,
						},
						"organization_id": schema.StringAttribute{
							MarkdownDescription: `Organization ID`,
							Computed:            true,
						},
						"status": schema.SingleNestedAttribute{
							MarkdownDescription: `Custom analytics artifact status`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"message": schema.StringAttribute{
									MarkdownDescription: `Status message`,
									Computed:            true,
								},
								"type": schema.StringAttribute{
									MarkdownDescription: `Status type`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsCameraCustomAnalyticsArtifactsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsCameraCustomAnalyticsArtifacts OrganizationsCameraCustomAnalyticsArtifacts
	diags := req.Config.Get(ctx, &organizationsCameraCustomAnalyticsArtifacts)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsCameraCustomAnalyticsArtifacts.OrganizationID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsCameraCustomAnalyticsArtifacts.OrganizationID.IsNull(), !organizationsCameraCustomAnalyticsArtifacts.ArtifactID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationCameraCustomAnalyticsArtifacts")
		vvOrganizationID := organizationsCameraCustomAnalyticsArtifacts.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Camera.GetOrganizationCameraCustomAnalyticsArtifacts(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCameraCustomAnalyticsArtifacts",
				err.Error(),
			)
			return
		}

		organizationsCameraCustomAnalyticsArtifacts = ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactsItemsToBody(organizationsCameraCustomAnalyticsArtifacts, response1)
		diags = resp.State.Set(ctx, &organizationsCameraCustomAnalyticsArtifacts)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationCameraCustomAnalyticsArtifact")
		vvOrganizationID := organizationsCameraCustomAnalyticsArtifacts.OrganizationID.ValueString()
		vvArtifactID := organizationsCameraCustomAnalyticsArtifacts.ArtifactID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Camera.GetOrganizationCameraCustomAnalyticsArtifact(vvOrganizationID, vvArtifactID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCameraCustomAnalyticsArtifact",
				err.Error(),
			)
			return
		}

		organizationsCameraCustomAnalyticsArtifacts = ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactItemToBody(organizationsCameraCustomAnalyticsArtifacts, response2)
		diags = resp.State.Set(ctx, &organizationsCameraCustomAnalyticsArtifacts)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsCameraCustomAnalyticsArtifacts struct {
	OrganizationID types.String                                                       `tfsdk:"organization_id"`
	ArtifactID     types.String                                                       `tfsdk:"artifact_id"`
	Items          *[]ResponseItemCameraGetOrganizationCameraCustomAnalyticsArtifacts `tfsdk:"items"`
	Item           *ResponseCameraGetOrganizationCameraCustomAnalyticsArtifact        `tfsdk:"item"`
}

type ResponseItemCameraGetOrganizationCameraCustomAnalyticsArtifacts struct {
	ArtifactID     types.String                                                           `tfsdk:"artifact_id"`
	Name           types.String                                                           `tfsdk:"name"`
	OrganizationID types.String                                                           `tfsdk:"organization_id"`
	Status         *ResponseItemCameraGetOrganizationCameraCustomAnalyticsArtifactsStatus `tfsdk:"status"`
}

type ResponseItemCameraGetOrganizationCameraCustomAnalyticsArtifactsStatus struct {
	Message types.String `tfsdk:"message"`
	Type    types.String `tfsdk:"type"`
}

type ResponseCameraGetOrganizationCameraCustomAnalyticsArtifact struct {
	ArtifactID     types.String                                                      `tfsdk:"artifact_id"`
	Name           types.String                                                      `tfsdk:"name"`
	OrganizationID types.String                                                      `tfsdk:"organization_id"`
	Status         *ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactStatus `tfsdk:"status"`
}

type ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactStatus struct {
	Message types.String `tfsdk:"message"`
	Type    types.String `tfsdk:"type"`
}

// ToBody
func ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactsItemsToBody(state OrganizationsCameraCustomAnalyticsArtifacts, response *merakigosdk.ResponseCameraGetOrganizationCameraCustomAnalyticsArtifacts) OrganizationsCameraCustomAnalyticsArtifacts {
	var items []ResponseItemCameraGetOrganizationCameraCustomAnalyticsArtifacts
	for _, item := range *response {
		itemState := ResponseItemCameraGetOrganizationCameraCustomAnalyticsArtifacts{
			ArtifactID:     types.StringValue(item.ArtifactID),
			Name:           types.StringValue(item.Name),
			OrganizationID: types.StringValue(item.OrganizationID),
			Status: func() *ResponseItemCameraGetOrganizationCameraCustomAnalyticsArtifactsStatus {
				if item.Status != nil {
					return &ResponseItemCameraGetOrganizationCameraCustomAnalyticsArtifactsStatus{
						Message: types.StringValue(item.Status.Message),
						Type:    types.StringValue(item.Status.Type),
					}
				}
				return nil
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactItemToBody(state OrganizationsCameraCustomAnalyticsArtifacts, response *merakigosdk.ResponseCameraGetOrganizationCameraCustomAnalyticsArtifact) OrganizationsCameraCustomAnalyticsArtifacts {
	itemState := ResponseCameraGetOrganizationCameraCustomAnalyticsArtifact{
		ArtifactID:     types.StringValue(response.ArtifactID),
		Name:           types.StringValue(response.Name),
		OrganizationID: types.StringValue(response.OrganizationID),
		Status: func() *ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactStatus {
			if response.Status != nil {
				return &ResponseCameraGetOrganizationCameraCustomAnalyticsArtifactStatus{
					Message: types.StringValue(response.Status.Message),
					Type:    types.StringValue(response.Status.Type),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
