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

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsActionBatchesResource{}
	_ resource.ResourceWithConfigure = &OrganizationsActionBatchesResource{}
)

func NewOrganizationsActionBatchesResource() resource.Resource {
	return &OrganizationsActionBatchesResource{}
}

type OrganizationsActionBatchesResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsActionBatchesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsActionBatchesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_action_batches"
}

func (r *OrganizationsActionBatchesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"action_batch_id": schema.StringAttribute{
				MarkdownDescription: `actionBatchId path parameter. Action batch ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"actions": schema.SetNestedAttribute{
				MarkdownDescription: `A set of changes made as part of this action (<a href='https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/'>more details</a>)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
					SuppressDiffSet(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"body": schema.StringAttribute{
							//Todo interface
							MarkdownDescription: `Data provided in the body of the Action. Contents depend on the Action type`,
							Computed:            true,
							Optional:            true,
						},
						"operation": schema.StringAttribute{
							MarkdownDescription: `The operation to be used by this action`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"resource": schema.StringAttribute{
							MarkdownDescription: `Unique identifier for the resource to be acted on`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"callback": schema.SingleNestedAttribute{
				MarkdownDescription: `Information for callback used to send back results`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"http_server": schema.SingleNestedAttribute{
						MarkdownDescription: `The webhook receiver used for the callback webhook.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"id": schema.StringAttribute{
								MarkdownDescription: `The webhook receiver ID that will receive information. If specifying this, please leave the url and sharedSecret fields blank.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
									SuppressDiffString(),
								},
							},
						},
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `The ID of the callback. To check the status of the callback, use this ID in a request to /webhooks/callbacks/statuses/{id}`,
						Computed:            true,
					},
					"payload_template": schema.SingleNestedAttribute{
						MarkdownDescription: `The payload template of the webhook used for the callback`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"id": schema.StringAttribute{
								MarkdownDescription: `The ID of the payload template. Defaults to 'wpt_00005' for the Callback (included) template.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
									SuppressDiffString(),
								},
							},
						},
					},
					"shared_secret": schema.StringAttribute{
						MarkdownDescription: `A shared secret that will be included in the requests sent to the callback URL. It can be used to verify that the request was sent by Meraki. If using this field, please also specify an url.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
							SuppressDiffString(),
						},
					},
					"status": schema.StringAttribute{
						MarkdownDescription: `The status of the callback`,
						Computed:            true,
					},
					"url": schema.StringAttribute{
						MarkdownDescription: `The callback URL for the webhook target. This was either provided in the original request or comes from a configured webhook receiver`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
							SuppressDiffString(),
						},
					},
				},
			},
			"confirmed": schema.BoolAttribute{
				MarkdownDescription: `Flag describing whether the action should be previewed before executing or not`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: `ID of the action batch. Can be used to check the status of the action batch at /organizations/{organizationId}/actionBatches/{actionBatchId}`,
				Computed:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `ID of the organization this action batch belongs to`,
				Required:            true,
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
					"errors": schema.SetAttribute{
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
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

//path params to set ['actionBatchId']
//path params to assign NOT EDITABLE ['actions', 'callback']

func (r *OrganizationsActionBatchesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsActionBatchesRs

	var item types.Object
	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	vvOrganizationID := data.OrganizationID.ValueString()
	// organization_id
	//Reviw This  Has Item and item
	//HAS CREATE

	vvActionBatchID := data.ActionBatchID.ValueString()
	if vvActionBatchID != "" {
		responseVerifyItem, restyRespGet, err := r.client.Organizations.GetOrganizationActionBatch(vvOrganizationID, vvActionBatchID)
		if err != nil || responseVerifyItem == nil {
			if restyRespGet != nil {
				if restyRespGet.StatusCode() != 404 {

					resp.Diagnostics.AddError(
						"Failure when executing GetOrganizationActionBatch",
						err.Error(),
					)
					return
				}
			}
		}

		if responseVerifyItem != nil {
			data = ResponseOrganizationsGetOrganizationActionBatchItemToBodyRs(data, responseVerifyItem, false)
			diags := resp.State.Set(ctx, &data)
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	response, restyResp2, err := r.client.Organizations.CreateOrganizationActionBatch(vvOrganizationID, data.toSdkApiRequestCreate(ctx))

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing ",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing ",
			err.Error(),
		)
		return
	}
	//Items
	vvActionBatchID = response.ID
	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationActionBatch(vvOrganizationID, vvActionBatchID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationActionBatches",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationActionBatches",
			err.Error(),
		)
		return
	} else {
		data = ResponseOrganizationsGetOrganizationActionBatchItemToBodyRs(data, responseGet, false)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
	}

}

func (r *OrganizationsActionBatchesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsActionBatchesRs

	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	// Has Item2

	vvOrganizationID := data.OrganizationID.ValueString()
	vvActionBatchID := data.ActionBatchID.ValueString()
	responseGet, restyRespGet, err := r.client.Organizations.GetOrganizationActionBatch(vvOrganizationID, vvActionBatchID)
	if err != nil || restyRespGet == nil {
		if restyRespGet != nil {
			if restyRespGet.StatusCode() == 404 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationActionBatch",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationActionBatch",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseOrganizationsGetOrganizationActionBatchItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *OrganizationsActionBatchesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("action_batch_id"), idParts[0])...)
}

func (r *OrganizationsActionBatchesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsActionBatchesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	vvActionBatchID := data.ActionBatchID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	_, restyResp2, err := r.client.Organizations.UpdateOrganizationActionBatch(vvOrganizationID, vvActionBatchID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationActionBatch",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationActionBatch",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsActionBatchesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsActionBatchesRs
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &state, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}

	vvOrganizationID := state.OrganizationID.ValueString()
	vvActionBatchID := state.ActionBatchID.ValueString()
	_, err := r.client.Organizations.DeleteOrganizationActionBatch(vvOrganizationID, vvActionBatchID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationActionBatch", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsActionBatchesRs struct {
	OrganizationID types.String                                                `tfsdk:"organization_id"`
	ActionBatchID  types.String                                                `tfsdk:"action_batch_id"`
	Actions        *[]ResponseOrganizationsGetOrganizationActionBatchActionsRs `tfsdk:"actions"`
	Callback       *ResponseOrganizationsGetOrganizationActionBatchCallbackRs  `tfsdk:"callback"`
	Confirmed      types.Bool                                                  `tfsdk:"confirmed"`
	ID             types.String                                                `tfsdk:"id"`
	Status         *ResponseOrganizationsGetOrganizationActionBatchStatusRs    `tfsdk:"status"`
	Synchronous    types.Bool                                                  `tfsdk:"synchronous"`
}

type ResponseOrganizationsGetOrganizationActionBatchActionsRs struct {
	Body      *ResponseOrganizationsGetOrganizationActionBatchActionsBodyRs `tfsdk:"body"`
	Operation types.String                                                  `tfsdk:"operation"`
	Resource  types.String                                                  `tfsdk:"resource"`
}

type ResponseOrganizationsGetOrganizationActionBatchCallbackHTTPServerRs struct {
	ID types.String `tfsdk:"id"`
}

type ResponseOrganizationsGetOrganizationActionBatchActionsBodyRs interface{}

type ResponseOrganizationsGetOrganizationActionBatchCallbackRs struct {
	ID              types.String                                                              `tfsdk:"id"`
	HTTPServer      *ResponseOrganizationsGetOrganizationActionBatchCallbackHTTPServerRs      `tfsdk:"http_server"`
	PayloadTemplate *ResponseOrganizationsGetOrganizationActionBatchCallbackPayloadTemplateRs `tfsdk:"payload_template"`
	SharedSecret    types.String                                                              `tfsdk:"shared_secret"`
	Status          types.String                                                              `tfsdk:"status"`
	URL             types.String                                                              `tfsdk:"url"`
}

type ResponseOrganizationsGetOrganizationActionBatchStatusRs struct {
	Completed        types.Bool                                                                 `tfsdk:"completed"`
	CreatedResources *[]ResponseOrganizationsGetOrganizationActionBatchStatusCreatedResourcesRs `tfsdk:"created_resources"`
	Errors           types.Set                                                                  `tfsdk:"errors"`
	Failed           types.Bool                                                                 `tfsdk:"failed"`
}

type ResponseOrganizationsGetOrganizationActionBatchStatusCreatedResourcesRs struct {
	ID  types.String `tfsdk:"id"`
	URI types.String `tfsdk:"uri"`
}

type ResponseOrganizationsGetOrganizationActionBatchCallbackPayloadTemplateRs struct {
	ID types.String `tfsdk:"id"`
}

// FromBody
func (r *OrganizationsActionBatchesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationActionBatch {
	var requestOrganizationsCreateOrganizationActionBatchActions []merakigosdk.RequestOrganizationsCreateOrganizationActionBatchActions
	if r.Actions != nil {
		for _, rItem1 := range *r.Actions {
			// var requestOrganizationsCreateOrganizationActionBatchActionsBody *merakigosdk.RequestOrganizationsCreateOrganizationActionBatchActionsBody
			// if rItem1.Body != nil {
			// 	requestOrganizationsCreateOrganizationActionBatchActionsBody = &merakigosdk.RequestOrganizationsCreateOrganizationActionBatchActionsBody{}
			// }
			operation := rItem1.Operation.ValueString()
			resource := rItem1.Resource.ValueString()
			requestOrganizationsCreateOrganizationActionBatchActions = append(requestOrganizationsCreateOrganizationActionBatchActions, merakigosdk.RequestOrganizationsCreateOrganizationActionBatchActions{
				// Body:      requestOrganizationsCreateOrganizationActionBatchActionsBody,
				Operation: operation,
				Resource:  resource,
			})
		}
	}
	var requestOrganizationsCreateOrganizationActionBatchCallback *merakigosdk.RequestOrganizationsCreateOrganizationActionBatchCallback
	if r.Callback != nil {
		var requestOrganizationsCreateOrganizationActionBatchCallbackHTTPServer *merakigosdk.RequestOrganizationsCreateOrganizationActionBatchCallbackHTTPServer
		if r.Callback != nil {
			iD := r.Callback.HTTPServer.ID.ValueString()
			requestOrganizationsCreateOrganizationActionBatchCallbackHTTPServer = &merakigosdk.RequestOrganizationsCreateOrganizationActionBatchCallbackHTTPServer{
				ID: iD,
			}
		}
		var requestOrganizationsCreateOrganizationActionBatchCallbackPayloadTemplate *merakigosdk.RequestOrganizationsCreateOrganizationActionBatchCallbackPayloadTemplate
		if r.Callback.PayloadTemplate != nil {
			iD := r.Callback.PayloadTemplate.ID.ValueString()
			requestOrganizationsCreateOrganizationActionBatchCallbackPayloadTemplate = &merakigosdk.RequestOrganizationsCreateOrganizationActionBatchCallbackPayloadTemplate{
				ID: iD,
			}
		}
		sharedSecret := r.Callback.SharedSecret.ValueString()
		uRL := r.Callback.URL.ValueString()
		requestOrganizationsCreateOrganizationActionBatchCallback = &merakigosdk.RequestOrganizationsCreateOrganizationActionBatchCallback{
			HTTPServer:      requestOrganizationsCreateOrganizationActionBatchCallbackHTTPServer,
			PayloadTemplate: requestOrganizationsCreateOrganizationActionBatchCallbackPayloadTemplate,
			SharedSecret:    sharedSecret,
			URL:             uRL,
		}
	}
	confirmed := new(bool)
	if !r.Confirmed.IsUnknown() && !r.Confirmed.IsNull() {
		*confirmed = r.Confirmed.ValueBool()
	} else {
		confirmed = nil
	}
	synchronous := new(bool)
	if !r.Synchronous.IsUnknown() && !r.Synchronous.IsNull() {
		*synchronous = r.Synchronous.ValueBool()
	} else {
		synchronous = nil
	}
	out := merakigosdk.RequestOrganizationsCreateOrganizationActionBatch{
		Actions: func() *[]merakigosdk.RequestOrganizationsCreateOrganizationActionBatchActions {
			if len(requestOrganizationsCreateOrganizationActionBatchActions) > 0 {
				return &requestOrganizationsCreateOrganizationActionBatchActions
			}
			return nil
		}(),
		Callback:    requestOrganizationsCreateOrganizationActionBatchCallback,
		Confirmed:   confirmed,
		Synchronous: synchronous,
	}
	return &out
}
func (r *OrganizationsActionBatchesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestOrganizationsUpdateOrganizationActionBatch {
	confirmed := new(bool)
	if !r.Confirmed.IsUnknown() && !r.Confirmed.IsNull() {
		*confirmed = r.Confirmed.ValueBool()
	} else {
		confirmed = nil
	}
	synchronous := new(bool)
	if !r.Synchronous.IsUnknown() && !r.Synchronous.IsNull() {
		*synchronous = r.Synchronous.ValueBool()
	} else {
		synchronous = nil
	}
	out := merakigosdk.RequestOrganizationsUpdateOrganizationActionBatch{
		Confirmed:   confirmed,
		Synchronous: synchronous,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseOrganizationsGetOrganizationActionBatchItemToBodyRs(state OrganizationsActionBatchesRs, response *merakigosdk.ResponseOrganizationsGetOrganizationActionBatch, is_read bool) OrganizationsActionBatchesRs {
	itemState := OrganizationsActionBatchesRs{
		Actions: func() *[]ResponseOrganizationsGetOrganizationActionBatchActionsRs {
			if response.Actions != nil {
				result := make([]ResponseOrganizationsGetOrganizationActionBatchActionsRs, len(*response.Actions))
				for i, actions := range *response.Actions {
					result[i] = ResponseOrganizationsGetOrganizationActionBatchActionsRs{
						// Body:      types.StringValue(actions.Body), //TODO POSIBLE interface
						Operation: types.StringValue(actions.Operation),
						Resource:  types.StringValue(actions.Resource),
					}
				}
				return &result
			}
			return nil
		}(),
		Callback: func() *ResponseOrganizationsGetOrganizationActionBatchCallbackRs {
			if response.Callback != nil {
				return &ResponseOrganizationsGetOrganizationActionBatchCallbackRs{
					ID:     types.StringValue(response.Callback.ID),
					Status: types.StringValue(response.Callback.Status),
					URL:    types.StringValue(response.Callback.URL),
				}
			}
			return nil
		}(),
		Confirmed: func() types.Bool {
			if response.Confirmed != nil {
				return types.BoolValue(*response.Confirmed)
			}
			return types.Bool{}
		}(),
		ID:             types.StringValue(response.ID),
		OrganizationID: types.StringValue(response.OrganizationID),
		Status: func() *ResponseOrganizationsGetOrganizationActionBatchStatusRs {
			if response.Status != nil {
				return &ResponseOrganizationsGetOrganizationActionBatchStatusRs{
					Completed: func() types.Bool {
						if response.Status.Completed != nil {
							return types.BoolValue(*response.Status.Completed)
						}
						return types.Bool{}
					}(),
					CreatedResources: func() *[]ResponseOrganizationsGetOrganizationActionBatchStatusCreatedResourcesRs {
						if response.Status.CreatedResources != nil {
							result := make([]ResponseOrganizationsGetOrganizationActionBatchStatusCreatedResourcesRs, len(*response.Status.CreatedResources))
							for i, createdResources := range *response.Status.CreatedResources {
								result[i] = ResponseOrganizationsGetOrganizationActionBatchStatusCreatedResourcesRs{
									ID:  types.StringValue(createdResources.ID),
									URI: types.StringValue(createdResources.URI),
								}
							}
							return &result
						}
						return nil
					}(),
					Errors: StringSliceToSet(response.Status.Errors),
					Failed: func() types.Bool {
						if response.Status.Failed != nil {
							return types.BoolValue(*response.Status.Failed)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
		Synchronous: func() types.Bool {
			if response.Synchronous != nil {
				return types.BoolValue(*response.Synchronous)
			}
			return types.Bool{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsActionBatchesRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsActionBatchesRs)
}
