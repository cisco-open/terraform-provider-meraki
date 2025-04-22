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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksResource{}
	_ resource.ResourceWithConfigure = &NetworksResource{}
)

func NewNetworksResource() resource.Resource {
	return &NetworksResource{}
}

type NetworksResource struct {
	client *merakigosdk.Client
}

func (r *NetworksResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks"
}

func (r *NetworksResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// "copy_from_network_id": schema.StringAttribute{
			// 	MarkdownDescription: `The ID of the network to copy configuration from. Other provided parameters will override the copied configuration, except type which must match this network's type exactly.`,
			// 	Computed:            true,
			// 	Optional:            true,
			// 	PlanModifiers: []planmodifier.String{
			// 		stringplanmodifier.UseStateForUnknown(),
			// 		SuppressDiffString(),
			// 	},
			// },
			"enrollment_string": schema.StringAttribute{
				MarkdownDescription: `Enrollment string for the network`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"notes": schema.StringAttribute{
				MarkdownDescription: `Notes for the network`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `Organization ID`,
				Required:            true,
			},
			"product_types": schema.SetAttribute{
				MarkdownDescription: `List of the product types that the network supports`,
				Required:            true,
				PlanModifiers: []planmodifier.Set{
					SuppressDiffSet(),
				},

				ElementType: types.StringType,
			},
			"tags": schema.SetAttribute{
				MarkdownDescription: `Network tags`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"time_zone": schema.StringAttribute{
				MarkdownDescription: `Timezone of the network`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"url": schema.StringAttribute{
				MarkdownDescription: `URL to the network Dashboard UI`,
				Computed:            true,
			},
		},
	}
}

//path params to set ['networkId']
//path params to assign NOT EDITABLE ['copyFromNetworkId', 'productTypes']

func (r *NetworksResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksRs

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
	vvName := data.Name.ValueString()

	responseVerifyItem, restyResp1, err := r.client.Organizations.GetOrganizationNetworks(vvOrganizationID, &merakigosdk.GetOrganizationNetworksQueryParams{
		PerPage: -1,
	})
	//Has Post
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationNetworks",
					err.Error(),
				)
				return
			}
		}
	}

	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvNetworkID, ok := result2["ID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter NetworkID",
					"Fail Parsing NetworkID",
				)
				return
			}
			r.client.Networks.UpdateNetwork(vvNetworkID, data.toSdkApiRequestUpdate(ctx))

			responseVerifyItem2, _, _ := r.client.Networks.GetNetwork(vvNetworkID)
			if responseVerifyItem2 != nil {
				data = ResponseNetworksGetNetworkItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Organizations.CreateOrganizationNetwork(vvOrganizationID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationNetwork",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationNetwork",
			err.Error(),
		)
		return
	}

	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationNetworks(vvOrganizationID, &merakigosdk.GetOrganizationNetworksQueryParams{
		PerPage: -1,
	})

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationNetworks",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationNetworks",
			err.Error(),
		)
		return
	}

	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvNetworkID, ok := result2["ID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter NetworkID",
				"Fail Parsing NetworkID",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Networks.GetNetwork(vvNetworkID)
		if responseVerifyItem2 != nil && err == nil {
			data.NetworkID = types.StringValue(responseVerifyItem2.ID)
			data = ResponseNetworksGetNetworkItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetwork",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetwork",
				err.Error(),
			)
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Error in result.",
			"Error in result.",
		)
		return
	}

}

func (r *NetworksResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksRs

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

	vvNetworkID := data.ID.ValueString()
	responseGet, restyRespGet, err := r.client.Networks.GetNetwork(vvNetworkID)
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
				"Failure when executing GetNetwork",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetwork",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}

func (r *NetworksResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.ID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetwork(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetwork",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetwork",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksRs
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

	vvNetworkID := state.ID.ValueString()
	_, err := r.client.Networks.DeleteNetwork(vvNetworkID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetwork", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksRs struct {
	NetworkID               types.String `tfsdk:"network_id"`
	OrganizationID          types.String `tfsdk:"organization_id"`
	EnrollmentString        types.String `tfsdk:"enrollment_string"`
	ID                      types.String `tfsdk:"id"`
	IsBoundToConfigTemplate types.Bool   `tfsdk:"is_bound_to_config_template"`
	Name                    types.String `tfsdk:"name"`
	Notes                   types.String `tfsdk:"notes"`
	ProductTypes            types.Set    `tfsdk:"product_types"`
	Tags                    types.Set    `tfsdk:"tags"`
	TimeZone                types.String `tfsdk:"time_zone"`
	URL                     types.String `tfsdk:"url"`
}

// FromBody
func (r *NetworksRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationNetwork {
	emptyString := ""
	// copyFromNetworkID := new(string)
	// if !r.CopyFromNetworkID.IsUnknown() && !r.CopyFromNetworkID.IsNull() {
	// 	*copyFromNetworkID = r.CopyFromNetworkID.ValueString()
	// } else {
	// 	copyFromNetworkID = &emptyString
	// }
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	notes := new(string)
	if !r.Notes.IsUnknown() && !r.Notes.IsNull() {
		*notes = r.Notes.ValueString()
	} else {
		notes = &emptyString
	}
	var productTypes []string = nil
	r.ProductTypes.ElementsAs(ctx, &productTypes, false)
	var tags []string = nil
	r.Tags.ElementsAs(ctx, &tags, false)
	timeZone := new(string)
	if !r.TimeZone.IsUnknown() && !r.TimeZone.IsNull() {
		*timeZone = r.TimeZone.ValueString()
	} else {
		timeZone = &emptyString
	}
	out := merakigosdk.RequestOrganizationsCreateOrganizationNetwork{
		// CopyFromNetworkID: *copyFromNetworkID,
		Name:         *name,
		Notes:        *notes,
		ProductTypes: productTypes,
		Tags:         tags,
		TimeZone:     *timeZone,
	}
	return &out
}
func (r *NetworksRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetwork {
	emptyString := ""
	enrollmentString := new(string)
	if !r.EnrollmentString.IsUnknown() && !r.EnrollmentString.IsNull() {
		*enrollmentString = r.EnrollmentString.ValueString()
	} else {
		enrollmentString = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	notes := new(string)
	if !r.Notes.IsUnknown() && !r.Notes.IsNull() {
		*notes = r.Notes.ValueString()
	} else {
		notes = &emptyString
	}
	var tags []string = nil
	r.Tags.ElementsAs(ctx, &tags, false)
	timeZone := new(string)
	if !r.TimeZone.IsUnknown() && !r.TimeZone.IsNull() {
		*timeZone = r.TimeZone.ValueString()
	} else {
		timeZone = &emptyString
	}
	out := merakigosdk.RequestNetworksUpdateNetwork{
		EnrollmentString: *enrollmentString,
		Name:             *name,
		Notes:            *notes,
		Tags:             tags,
		TimeZone:         *timeZone,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkItemToBodyRs(state NetworksRs, response *merakigosdk.ResponseNetworksGetNetwork, is_read bool) NetworksRs {
	itemState := NetworksRs{
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
		ProductTypes:   StringSliceToSet(response.ProductTypes),
		Tags:           StringSliceToSet(response.Tags),
		TimeZone:       types.StringValue(response.TimeZone),
		URL:            types.StringValue(response.URL),
	}
	itemState.NetworkID = types.StringValue(response.ID)
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksRs)
}
