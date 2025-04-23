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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsApplianceDNSSplitProfilesResource{}
	_ resource.ResourceWithConfigure = &OrganizationsApplianceDNSSplitProfilesResource{}
)

func NewOrganizationsApplianceDNSSplitProfilesResource() resource.Resource {
	return &OrganizationsApplianceDNSSplitProfilesResource{}
}

type OrganizationsApplianceDNSSplitProfilesResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsApplianceDNSSplitProfilesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsApplianceDNSSplitProfilesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_appliance_dns_split_profiles"
}

func (r *OrganizationsApplianceDNSSplitProfilesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"hostnames": schema.SetAttribute{
				MarkdownDescription: `The hostname patterns to match for redirection. For more information on Split DNS hostname pattern formatting, please consult the Split DNS KB.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Name of profile`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"nameservers": schema.SingleNestedAttribute{
				MarkdownDescription: `Contains the nameserver information for redirection.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"addresses": schema.SetAttribute{
						MarkdownDescription: `The nameserver address(es) to use for redirection. A maximum of one address is supported.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},

						ElementType: types.StringType,
					},
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"profile_id": schema.StringAttribute{
				MarkdownDescription: `Profile ID`,
				Computed:            true,
			},
		},
	}
}

func (r *OrganizationsApplianceDNSSplitProfilesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsApplianceDNSSplitProfilesRs

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

	vvName := data.Name.ValueString()

	responseVerifyItem, restyResp1, err := r.client.Appliance.GetOrganizationApplianceDNSSplitProfiles(vvOrganizationID, nil)
	//Has Post
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationApplianceDNSSplitProfiles",
					err.Error(),
				)
				return
			}
		}
	}

	var responseVerifyItem2 merakigosdk.ResponseItemApplianceGetOrganizationApplianceDNSSplitProfiles
	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			err := mapToStruct(result.(map[string]interface{}), &responseVerifyItem2)
			if err != nil {
				resp.Diagnostics.AddError(
					"Failure when executing mapToStruct in resource",
					err.Error(),
				)
				return
			}
			data = ResponseApplianceGetOrganizationApplianceDNSSplitProfilesItemToBodyRs(data, &responseVerifyItem2, false)
			// Path params update assigned
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return

		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Appliance.CreateOrganizationApplianceDNSSplitProfile(vvOrganizationID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationApplianceDNSSplitProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationApplianceDNSSplitProfile",
			err.Error(),
		)
		return
	}

	responseGet, restyResp1, err := r.client.Appliance.GetOrganizationApplianceDNSSplitProfiles(vvOrganizationID, nil)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationApplianceDNSSplitProfiles",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationApplianceDNSSplitProfiles",
			err.Error(),
		)
		return
	}

	responseStruct := structToMap(responseGet)
	result2 := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result2 != nil {
		err := mapToStruct(result2.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		data = ResponseApplianceGetOrganizationApplianceDNSSplitProfilesItemToBodyRs(data, &responseVerifyItem2, false)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	} else {
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationApplianceDNSSplitProfiles Result",
			"Not Found",
		)
		return
	}

}

func (r *OrganizationsApplianceDNSSplitProfilesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsApplianceDNSSplitProfilesRs

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
	vvName := data.Name.ValueString()

	responseGet, restyResp1, err := r.client.Appliance.GetOrganizationApplianceDNSSplitProfiles(vvOrganizationID, nil)

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
				"Failure when executing GetOrganizationApplianceDNSSplitProfiles",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationApplianceDNSSplitProfiles",
			err.Error(),
		)
		return
	}
	responseStruct2 := structToMap(responseGet)
	result2 := getDictResult(responseStruct2, "Name", vvName, simpleCmp)
	var responseVerifyItem2 merakigosdk.ResponseItemApplianceGetOrganizationApplianceDNSSplitProfiles
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
		data = ResponseApplianceGetOrganizationApplianceDNSSplitProfilesItemToBodyRs(data, &responseVerifyItem2, true)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	} else {
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationApplianceDNSSplitProfiles Result",
			err.Error(),
		)
		return
	}
}

func (r *OrganizationsApplianceDNSSplitProfilesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), idParts[1])...)
}

func (r *OrganizationsApplianceDNSSplitProfilesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsApplianceDNSSplitProfilesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update
	// No update
	resp.Diagnostics.AddError(
		"Update operation not supported in OrganizationsApplianceDNSSplitProfiles",
		"Update operation not supported in OrganizationsApplianceDNSSplitProfiles",
	)
	return
}

func (r *OrganizationsApplianceDNSSplitProfilesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting OrganizationsApplianceDNSSplitProfiles", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsApplianceDNSSplitProfilesRs struct {
	OrganizationID types.String `tfsdk:"organization_id"`
	//TIENE ITEMS
	Hostnames   types.Set                                                                   `tfsdk:"hostnames"`
	Name        types.String                                                                `tfsdk:"name"`
	Nameservers *ResponseItemApplianceGetOrganizationApplianceDnsSplitProfilesNameserversRs `tfsdk:"nameservers"`
	ProfileID   types.String                                                                `tfsdk:"profile_id"`
}

type ResponseItemApplianceGetOrganizationApplianceDnsSplitProfilesNameserversRs struct {
	Addresses types.Set `tfsdk:"addresses"`
}

// FromBody
func (r *OrganizationsApplianceDNSSplitProfilesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestApplianceCreateOrganizationApplianceDNSSplitProfile {
	emptyString := ""
	var hostnames []string = nil
	r.Hostnames.ElementsAs(ctx, &hostnames, false)
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestApplianceCreateOrganizationApplianceDNSSplitProfileNameservers *merakigosdk.RequestApplianceCreateOrganizationApplianceDNSSplitProfileNameservers

	if r.Nameservers != nil {

		var addresses []string = nil
		r.Nameservers.Addresses.ElementsAs(ctx, &addresses, false)
		requestApplianceCreateOrganizationApplianceDNSSplitProfileNameservers = &merakigosdk.RequestApplianceCreateOrganizationApplianceDNSSplitProfileNameservers{
			Addresses: addresses,
		}
		//[debug] Is Array: False
	}
	out := merakigosdk.RequestApplianceCreateOrganizationApplianceDNSSplitProfile{
		Hostnames:   hostnames,
		Name:        *name,
		Nameservers: requestApplianceCreateOrganizationApplianceDNSSplitProfileNameservers,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetOrganizationApplianceDNSSplitProfilesItemToBodyRs(state OrganizationsApplianceDNSSplitProfilesRs, response *merakigosdk.ResponseItemApplianceGetOrganizationApplianceDNSSplitProfiles, is_read bool) OrganizationsApplianceDNSSplitProfilesRs {
	itemState := OrganizationsApplianceDNSSplitProfilesRs{
		Hostnames: StringSliceToSet(response.Hostnames),
		Name:      types.StringValue(response.Name),
		Nameservers: func() *ResponseItemApplianceGetOrganizationApplianceDnsSplitProfilesNameserversRs {
			if response.Nameservers != nil {
				return &ResponseItemApplianceGetOrganizationApplianceDnsSplitProfilesNameserversRs{
					Addresses: StringSliceToSet(response.Nameservers.Addresses),
				}
			}
			return nil
		}(),
		ProfileID: types.StringValue(response.ProfileID),
	}
	state = itemState
	return state
}
