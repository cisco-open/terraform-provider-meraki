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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsCellularGatewayEsimsServiceProvidersAccountsResource{}
	_ resource.ResourceWithConfigure = &OrganizationsCellularGatewayEsimsServiceProvidersAccountsResource{}
)

func NewOrganizationsCellularGatewayEsimsServiceProvidersAccountsResource() resource.Resource {
	return &OrganizationsCellularGatewayEsimsServiceProvidersAccountsResource{}
}

type OrganizationsCellularGatewayEsimsServiceProvidersAccountsResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsCellularGatewayEsimsServiceProvidersAccountsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsCellularGatewayEsimsServiceProvidersAccountsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_cellular_gateway_esims_service_providers_accounts"
}

func (r *OrganizationsCellularGatewayEsimsServiceProvidersAccountsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				MarkdownDescription: `Service provider account ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SuppressDiffString(),
				},
			},
			"api_key": schema.StringAttribute{
				MarkdownDescription: `Service provider account API key`,
				Computed:            true,
				Optional:            true,
				Sensitive:           true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"service_provider": schema.SingleNestedAttribute{
				MarkdownDescription: `Service Provider information`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"name": schema.StringAttribute{
						MarkdownDescription: `Service provider name`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"logo": schema.StringAttribute{
						MarkdownDescription: `Service provider logo`,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"title": schema.StringAttribute{
				MarkdownDescription: `Service provider account name`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"username": schema.StringAttribute{
				MarkdownDescription: `Service provider account username`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SuppressDiffString(),
				},
			},
		},
	}
}

//path params to set ['accountId']
//path params to assign NOT EDITABLE ['accountId', 'serviceProvider', 'username']

func (r *OrganizationsCellularGatewayEsimsServiceProvidersAccountsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsCellularGatewayEsimsServiceProvidersAccountsRs

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
	// Has Paths
	vvOrganizationID := data.OrganizationID.ValueString()
	//Only Items

	vvUsername := data.Username.ValueString()

	responseVerifyItem, restyResp1, err := r.client.CellularGateway.GetOrganizationCellularGatewayEsimsServiceProvidersAccounts(vvOrganizationID, nil)
	//Has Post
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationCellularGatewayEsimsServiceProvidersAccounts",
					err.Error(),
				)
				return
			}
		}
	}

	var responseVerifyItem2 merakigosdk.ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItems
	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem.Items)
		result := getDictResult(responseStruct, "Username", vvUsername, simpleCmp)
		if result != nil {
			err := mapToStruct(result.(map[string]interface{}), &responseVerifyItem2)
			if err != nil {
				resp.Diagnostics.AddError(
					"Failure when executing mapToStruct in resource",
					err.Error(),
				)
				return
			}
			r.client.CellularGateway.UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount(vvOrganizationID, responseVerifyItem2.AccountID, data.toSdkApiRequestUpdate(ctx))
			data = ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemToBodyRs(data, &responseVerifyItem2, false)
			// Path params update assigned
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return

		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.CellularGateway.CreateOrganizationCellularGatewayEsimsServiceProvidersAccount(vvOrganizationID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationCellularGatewayEsimsServiceProvidersAccount",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationCellularGatewayEsimsServiceProvidersAccount",
			err.Error(),
		)
		return
	}

	responseGet, restyResp1, err := r.client.CellularGateway.GetOrganizationCellularGatewayEsimsServiceProvidersAccounts(vvOrganizationID, nil)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCellularGatewayEsimsServiceProvidersAccounts",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationCellularGatewayEsimsServiceProvidersAccounts",
			err.Error(),
		)
		return
	}

	responseStruct := structToMap(responseGet.Items)
	result2 := getDictResult(responseStruct, "Username", vvUsername, simpleCmp)
	if result2 != nil {
		err := mapToStruct(result2.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		data = ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemToBodyRs(data, &responseVerifyItem2, false)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	} else {
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationCellularGatewayEsimsServiceProvidersAccounts Result",
			"Not Found",
		)
		return
	}

}

func (r *OrganizationsCellularGatewayEsimsServiceProvidersAccountsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsCellularGatewayEsimsServiceProvidersAccountsRs

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
	// Not has Item

	vvOrganizationID := data.OrganizationID.ValueString()
	vvUsername := data.Username.ValueString()

	responseGet, restyResp1, err := r.client.CellularGateway.GetOrganizationCellularGatewayEsimsServiceProvidersAccounts(vvOrganizationID, nil)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() == 404 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCellularGatewayEsimsServiceProvidersAccounts",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationCellularGatewayEsimsServiceProvidersAccounts",
			err.Error(),
		)
		return
	}
	responseStruct2 := structToMap(responseGet.Items)
	result2 := getDictResult(responseStruct2, "Username", vvUsername, simpleCmp)
	var responseVerifyItem2 merakigosdk.ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItems
	if result2 != nil {
		err := mapToStruct(result2.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		//entro aqui
		data = ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemToBodyRs(data, &responseVerifyItem2, true)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	} else {
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationCellularGatewayEsimsServiceProvidersAccounts Result",
			err.Error(),
		)
		return
	}
}
func (r *OrganizationsCellularGatewayEsimsServiceProvidersAccountsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("username"), idParts[1])...)
}

func (r *OrganizationsCellularGatewayEsimsServiceProvidersAccountsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsCellularGatewayEsimsServiceProvidersAccountsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	vvAccountID := data.AccountID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.CellularGateway.UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount(vvOrganizationID, vvAccountID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationCellularGatewayEsimsServiceProvidersAccount",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsCellularGatewayEsimsServiceProvidersAccountsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsCellularGatewayEsimsServiceProvidersAccountsRs
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
	vvAccountID := state.AccountID.ValueString()
	_, err := r.client.CellularGateway.DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount(vvOrganizationID, vvAccountID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationCellularGatewayEsimsServiceProvidersAccount", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsCellularGatewayEsimsServiceProvidersAccountsRs struct {
	OrganizationID  types.String                                                                                                  `tfsdk:"organization_id"`
	AccountID       types.String                                                                                                  `tfsdk:"account_id"`
	LastUpdatedAt   types.String                                                                                                  `tfsdk:"last_updated_at"`
	ServiceProvider *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProviderRs `tfsdk:"service_provider"`
	Title           types.String                                                                                                  `tfsdk:"title"`
	Username        types.String                                                                                                  `tfsdk:"username"`
	APIKey          types.String                                                                                                  `tfsdk:"api_key"`
}

type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProviderRs struct {
	Logo *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProviderLogoRs `tfsdk:"logo"`
	Name types.String                                                                                                      `tfsdk:"name"`
}

type ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProviderLogoRs struct {
	URL types.String `tfsdk:"url"`
}

// FromBody
func (r *OrganizationsCellularGatewayEsimsServiceProvidersAccountsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccount {
	emptyString := ""
	accountID := new(string)
	if !r.AccountID.IsUnknown() && !r.AccountID.IsNull() {
		*accountID = r.AccountID.ValueString()
	} else {
		accountID = &emptyString
	}
	aPIKey := new(string)
	if !r.APIKey.IsUnknown() && !r.APIKey.IsNull() {
		*aPIKey = r.APIKey.ValueString()
	} else {
		aPIKey = &emptyString
	}
	var requestCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccountServiceProvider *merakigosdk.RequestCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccountServiceProvider

	if r.ServiceProvider != nil {
		name := r.ServiceProvider.Name.ValueString()
		requestCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccountServiceProvider = &merakigosdk.RequestCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccountServiceProvider{
			Name: name,
		}
		//[debug] Is Array: False
	}
	title := new(string)
	if !r.Title.IsUnknown() && !r.Title.IsNull() {
		*title = r.Title.ValueString()
	} else {
		title = &emptyString
	}
	username := new(string)
	if !r.Username.IsUnknown() && !r.Username.IsNull() {
		*username = r.Username.ValueString()
	} else {
		username = &emptyString
	}
	out := merakigosdk.RequestCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccount{
		AccountID:       *accountID,
		APIKey:          *aPIKey,
		ServiceProvider: requestCellularGatewayCreateOrganizationCellularGatewayEsimsServiceProvidersAccountServiceProvider,
		Title:           *title,
		Username:        *username,
	}
	return &out
}
func (r *OrganizationsCellularGatewayEsimsServiceProvidersAccountsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestCellularGatewayUpdateOrganizationCellularGatewayEsimsServiceProvidersAccount {
	emptyString := ""
	aPIKey := new(string)
	if !r.APIKey.IsUnknown() && !r.APIKey.IsNull() {
		*aPIKey = r.APIKey.ValueString()
	} else {
		aPIKey = &emptyString
	}
	title := new(string)
	if !r.Title.IsUnknown() && !r.Title.IsNull() {
		*title = r.Title.ValueString()
	} else {
		title = &emptyString
	}
	out := merakigosdk.RequestCellularGatewayUpdateOrganizationCellularGatewayEsimsServiceProvidersAccount{
		APIKey: *aPIKey,
		Title:  *title,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemToBodyRs(state OrganizationsCellularGatewayEsimsServiceProvidersAccountsRs, response *merakigosdk.ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItems, is_read bool) OrganizationsCellularGatewayEsimsServiceProvidersAccountsRs {
	itemState := OrganizationsCellularGatewayEsimsServiceProvidersAccountsRs{
		AccountID:     types.StringValue(response.AccountID),
		LastUpdatedAt: types.StringValue(response.LastUpdatedAt),
		ServiceProvider: func() *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProviderRs {
			if response.ServiceProvider != nil {
				return &ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProviderRs{
					Logo: func() *ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProviderLogoRs {
						if response.ServiceProvider.Logo != nil {
							return &ResponseItemCellularGatewayGetOrganizationCellularGatewayEsimsServiceProvidersAccountsItemsServiceProviderLogoRs{
								URL: types.StringValue(response.ServiceProvider.Logo.URL),
							}
						}
						return nil
					}(),
					Name: types.StringValue(response.ServiceProvider.Name),
				}
			}
			return nil
		}(),
		Title:    types.StringValue(response.Title),
		Username: types.StringValue(response.Username),
	}
	state = itemState
	return state
}
