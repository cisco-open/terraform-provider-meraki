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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsActionBatchesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsActionBatchesDataSource{}
)

func NewOrganizationsActionBatchesDataSource() datasource.DataSource {
	return &OrganizationsActionBatchesDataSource{}
}

type OrganizationsActionBatchesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsActionBatchesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsActionBatchesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_action_batches"
}

func (d *OrganizationsActionBatchesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"action_batch_id": schema.StringAttribute{
				MarkdownDescription: `actionBatchId path parameter. Action batch ID`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: `status query parameter. Filter batches by status. Valid types are pending, completed, and failed.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"actions": schema.SetNestedAttribute{
						MarkdownDescription: `A set of changes made as part of this action (<a href='https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/'>more details</a>)`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"operation": schema.StringAttribute{
									MarkdownDescription: `The operation to be used by this action`,
									Computed:            true,
								},
								"resource": schema.StringAttribute{
									MarkdownDescription: `Unique identifier for the resource to be acted on`,
									Computed:            true,
								},
							},
						},
					},
					"confirmed": schema.BoolAttribute{
						MarkdownDescription: `Flag describing whether the action should be previewed before executing or not`,
						Computed:            true,
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `ID of the action batch. Can be used to check the status of the action batch at /organizations/{organizationId}/actionBatches/{actionBatchId}`,
						Computed:            true,
					},
					"organization_id": schema.StringAttribute{
						MarkdownDescription: `ID of the organization this action batch belongs to`,
						Computed:            true,
					},
					"status": schema.SingleNestedAttribute{
						MarkdownDescription: `Status of action batch`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"completed": schema.BoolAttribute{
								MarkdownDescription: `Flag describing whether all actions in the action batch have completed`,
								Computed:            true,
							},
							"created_resources": schema.SetNestedAttribute{
								MarkdownDescription: `Resources created as a result of this action batch`,
								Computed:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `ID of the created resource`,
											Computed:            true,
										},
										"uri": schema.StringAttribute{
											MarkdownDescription: `URI, not including base, of the created resource`,
											Computed:            true,
										},
									},
								},
							},
							"errors": schema.ListAttribute{
								MarkdownDescription: `List of errors encountered when running actions in the action batch`,
								Computed:            true,
								ElementType:         types.StringType,
							},
							"failed": schema.BoolAttribute{
								MarkdownDescription: `Flag describing whether any actions in the action batch failed`,
								Computed:            true,
							},
						},
					},
					"synchronous": schema.BoolAttribute{
						MarkdownDescription: `Flag describing whether actions should run synchronously or asynchronously`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationActionBatches`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"actions": schema.SetNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"body": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{

											"enabled": schema.BoolAttribute{
												Computed: true,
											},
										},
									},
									"operation": schema.StringAttribute{
										Computed: true,
									},
									"resource": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"confirmed": schema.BoolAttribute{
							Computed: true,
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"organization_id": schema.StringAttribute{
							Computed: true,
						},
						"status": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"completed": schema.BoolAttribute{
									Computed: true,
								},
								"created_resources": schema.SetNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"id": schema.Int64Attribute{
												Computed: true,
											},
											"uri": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
								"errors": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
								"failed": schema.BoolAttribute{
									Computed: true,
								},
							},
						},
						"synchronous": schema.BoolAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsActionBatchesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsActionBatches OrganizationsActionBatches
	diags := req.Config.Get(ctx, &organizationsActionBatches)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsActionBatches.OrganizationID.IsNull(), !organizationsActionBatches.Status.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsActionBatches.OrganizationID.IsNull(), !organizationsActionBatches.ActionBatchID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationActionBatches")
		vvOrganizationID := organizationsActionBatches.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationActionBatchesQueryParams{}

		queryParams1.Status = organizationsActionBatches.Status.ValueString()

		response1, restyResp1, err := d.client.Organizations.GetOrganizationActionBatches(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationActionBatches",
				err.Error(),
			)
			return
		}

		organizationsActionBatches = ResponseOrganizationsGetOrganizationActionBatchesItemsToBody(organizationsActionBatches, response1)
		diags = resp.State.Set(ctx, &organizationsActionBatches)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationActionBatch")
		vvOrganizationID := organizationsActionBatches.OrganizationID.ValueString()
		vvActionBatchID := organizationsActionBatches.ActionBatchID.ValueString()

		response2, restyResp2, err := d.client.Organizations.GetOrganizationActionBatch(vvOrganizationID, vvActionBatchID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationActionBatch",
				err.Error(),
			)
			return
		}

		organizationsActionBatches = ResponseOrganizationsGetOrganizationActionBatchItemToBody(organizationsActionBatches, response2)
		diags = resp.State.Set(ctx, &organizationsActionBatches)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsActionBatches struct {
	OrganizationID types.String                                             `tfsdk:"organization_id"`
	Status         types.String                                             `tfsdk:"status"`
	ActionBatchID  types.String                                             `tfsdk:"action_batch_id"`
	Items          *[]ResponseItemOrganizationsGetOrganizationActionBatches `tfsdk:"items"`
	Item           *ResponseOrganizationsGetOrganizationActionBatch         `tfsdk:"item"`
}

type ResponseItemOrganizationsGetOrganizationActionBatches struct {
	Actions        *[]ResponseItemOrganizationsGetOrganizationActionBatchesActions `tfsdk:"actions"`
	Confirmed      types.Bool                                                      `tfsdk:"confirmed"`
	ID             types.String                                                    `tfsdk:"id"`
	OrganizationID types.String                                                    `tfsdk:"organization_id"`
	Status         *ResponseItemOrganizationsGetOrganizationActionBatchesStatus    `tfsdk:"status"`
	Synchronous    types.Bool                                                      `tfsdk:"synchronous"`
}

type ResponseItemOrganizationsGetOrganizationActionBatchesActions struct {
	Body      *ResponseItemOrganizationsGetOrganizationActionBatchesActionsBody `tfsdk:"body"`
	Operation types.String                                                      `tfsdk:"operation"`
	Resource  types.String                                                      `tfsdk:"resource"`
}

type ResponseItemOrganizationsGetOrganizationActionBatchesActionsBody struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseItemOrganizationsGetOrganizationActionBatchesStatus struct {
	Completed        types.Bool                                                                     `tfsdk:"completed"`
	CreatedResources *[]ResponseItemOrganizationsGetOrganizationActionBatchesStatusCreatedResources `tfsdk:"created_resources"`
	Errors           types.List                                                                     `tfsdk:"errors"`
	Failed           types.Bool                                                                     `tfsdk:"failed"`
}

type ResponseItemOrganizationsGetOrganizationActionBatchesStatusCreatedResources struct {
	ID  types.Int64  `tfsdk:"id"`
	URI types.String `tfsdk:"uri"`
}

type ResponseOrganizationsGetOrganizationActionBatch struct {
	Actions        *[]ResponseOrganizationsGetOrganizationActionBatchActions `tfsdk:"actions"`
	Confirmed      types.Bool                                                `tfsdk:"confirmed"`
	ID             types.String                                              `tfsdk:"id"`
	OrganizationID types.String                                              `tfsdk:"organization_id"`
	Status         *ResponseOrganizationsGetOrganizationActionBatchStatus    `tfsdk:"status"`
	Synchronous    types.Bool                                                `tfsdk:"synchronous"`
}

type ResponseOrganizationsGetOrganizationActionBatchActions struct {
	Operation types.String `tfsdk:"operation"`
	Resource  types.String `tfsdk:"resource"`
}

type ResponseOrganizationsGetOrganizationActionBatchStatus struct {
	Completed        types.Bool                                                               `tfsdk:"completed"`
	CreatedResources *[]ResponseOrganizationsGetOrganizationActionBatchStatusCreatedResources `tfsdk:"created_resources"`
	Errors           types.List                                                               `tfsdk:"errors"`
	Failed           types.Bool                                                               `tfsdk:"failed"`
}

type ResponseOrganizationsGetOrganizationActionBatchStatusCreatedResources struct {
	ID  types.String `tfsdk:"id"`
	URI types.String `tfsdk:"uri"`
}

// ToBody
func ResponseOrganizationsGetOrganizationActionBatchesItemsToBody(state OrganizationsActionBatches, response *merakigosdk.ResponseOrganizationsGetOrganizationActionBatches) OrganizationsActionBatches {
	var items []ResponseItemOrganizationsGetOrganizationActionBatches
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationActionBatches{
			Actions: func() *[]ResponseItemOrganizationsGetOrganizationActionBatchesActions {
				if item.Actions != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationActionBatchesActions, len(*item.Actions))
					for i, actions := range *item.Actions {
						result[i] = ResponseItemOrganizationsGetOrganizationActionBatchesActions{
							Body: func() *ResponseItemOrganizationsGetOrganizationActionBatchesActionsBody {
								if actions.Body != nil {
									return &ResponseItemOrganizationsGetOrganizationActionBatchesActionsBody{
										Enabled: func() types.Bool {
											if actions.Body.Enabled != nil {
												return types.BoolValue(*actions.Body.Enabled)
											}
											return types.Bool{}
										}(),
									}
								}
								return &ResponseItemOrganizationsGetOrganizationActionBatchesActionsBody{}
							}(),
							Operation: types.StringValue(actions.Operation),
							Resource:  types.StringValue(actions.Resource),
						}
					}
					return &result
				}
				return &[]ResponseItemOrganizationsGetOrganizationActionBatchesActions{}
			}(),
			Confirmed: func() types.Bool {
				if item.Confirmed != nil {
					return types.BoolValue(*item.Confirmed)
				}
				return types.Bool{}
			}(),
			ID:             types.StringValue(item.ID),
			OrganizationID: types.StringValue(item.OrganizationID),
			Status: func() *ResponseItemOrganizationsGetOrganizationActionBatchesStatus {
				if item.Status != nil {
					return &ResponseItemOrganizationsGetOrganizationActionBatchesStatus{
						Completed: func() types.Bool {
							if item.Status.Completed != nil {
								return types.BoolValue(*item.Status.Completed)
							}
							return types.Bool{}
						}(),
						CreatedResources: func() *[]ResponseItemOrganizationsGetOrganizationActionBatchesStatusCreatedResources {
							if item.Status.CreatedResources != nil {
								result := make([]ResponseItemOrganizationsGetOrganizationActionBatchesStatusCreatedResources, len(*item.Status.CreatedResources))
								for i, createdResources := range *item.Status.CreatedResources {
									result[i] = ResponseItemOrganizationsGetOrganizationActionBatchesStatusCreatedResources{
										ID: func() types.Int64 {
											if createdResources.ID != nil {
												return types.Int64Value(int64(*createdResources.ID))
											}
											return types.Int64{}
										}(),
										URI: types.StringValue(createdResources.URI),
									}
								}
								return &result
							}
							return &[]ResponseItemOrganizationsGetOrganizationActionBatchesStatusCreatedResources{}
						}(),
						Errors: StringSliceToList(item.Status.Errors),
						Failed: func() types.Bool {
							if item.Status.Failed != nil {
								return types.BoolValue(*item.Status.Failed)
							}
							return types.Bool{}
						}(),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationActionBatchesStatus{}
			}(),
			Synchronous: func() types.Bool {
				if item.Synchronous != nil {
					return types.BoolValue(*item.Synchronous)
				}
				return types.Bool{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseOrganizationsGetOrganizationActionBatchItemToBody(state OrganizationsActionBatches, response *merakigosdk.ResponseOrganizationsGetOrganizationActionBatch) OrganizationsActionBatches {
	itemState := ResponseOrganizationsGetOrganizationActionBatch{
		Actions: func() *[]ResponseOrganizationsGetOrganizationActionBatchActions {
			if response.Actions != nil {
				result := make([]ResponseOrganizationsGetOrganizationActionBatchActions, len(*response.Actions))
				for i, actions := range *response.Actions {
					result[i] = ResponseOrganizationsGetOrganizationActionBatchActions{
						Operation: types.StringValue(actions.Operation),
						Resource:  types.StringValue(actions.Resource),
					}
				}
				return &result
			}
			return &[]ResponseOrganizationsGetOrganizationActionBatchActions{}
		}(),
		Confirmed: func() types.Bool {
			if response.Confirmed != nil {
				return types.BoolValue(*response.Confirmed)
			}
			return types.Bool{}
		}(),
		ID:             types.StringValue(response.ID),
		OrganizationID: types.StringValue(response.OrganizationID),
		Status: func() *ResponseOrganizationsGetOrganizationActionBatchStatus {
			if response.Status != nil {
				return &ResponseOrganizationsGetOrganizationActionBatchStatus{
					Completed: func() types.Bool {
						if response.Status.Completed != nil {
							return types.BoolValue(*response.Status.Completed)
						}
						return types.Bool{}
					}(),
					CreatedResources: func() *[]ResponseOrganizationsGetOrganizationActionBatchStatusCreatedResources {
						if response.Status.CreatedResources != nil {
							result := make([]ResponseOrganizationsGetOrganizationActionBatchStatusCreatedResources, len(*response.Status.CreatedResources))
							for i, createdResources := range *response.Status.CreatedResources {
								result[i] = ResponseOrganizationsGetOrganizationActionBatchStatusCreatedResources{
									ID:  types.StringValue(createdResources.ID),
									URI: types.StringValue(createdResources.URI),
								}
							}
							return &result
						}
						return &[]ResponseOrganizationsGetOrganizationActionBatchStatusCreatedResources{}
					}(),
					Errors: StringSliceToList(response.Status.Errors),
					Failed: func() types.Bool {
						if response.Status.Failed != nil {
							return types.BoolValue(*response.Status.Failed)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseOrganizationsGetOrganizationActionBatchStatus{}
		}(),
		Synchronous: func() types.Bool {
			if response.Synchronous != nil {
				return types.BoolValue(*response.Synchronous)
			}
			return types.Bool{}
		}(),
	}
	state.Item = &itemState
	return state
}
