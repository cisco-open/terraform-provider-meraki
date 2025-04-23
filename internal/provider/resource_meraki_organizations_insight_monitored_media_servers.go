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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsInsightMonitoredMediaServersResource{}
	_ resource.ResourceWithConfigure = &OrganizationsInsightMonitoredMediaServersResource{}
)

func NewOrganizationsInsightMonitoredMediaServersResource() resource.Resource {
	return &OrganizationsInsightMonitoredMediaServersResource{}
}

type OrganizationsInsightMonitoredMediaServersResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsInsightMonitoredMediaServersResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsInsightMonitoredMediaServersResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_insight_monitored_media_servers"
}

func (r *OrganizationsInsightMonitoredMediaServersResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"address": schema.StringAttribute{
				MarkdownDescription: `The IP address (IPv4 only) or hostname of the media server to monitor`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"best_effort_monitoring_enabled": schema.BoolAttribute{
				MarkdownDescription: `Indicates that if the media server doesn't respond to ICMP pings, the nearest hop will be used in its stead`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: `Monitored media server id`,
				Computed:            true,
			},
			"monitored_media_server_id": schema.StringAttribute{
				MarkdownDescription: `monitoredMediaServerId path parameter. Monitored media server ID`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name of the VoIP provider`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
		},
	}
}

//path params to set ['monitoredMediaServerId']

func (r *OrganizationsInsightMonitoredMediaServersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsInsightMonitoredMediaServersRs

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
	//Has Item and has items and post

	vvName := data.Name.ValueString()

	responseVerifyItem, restyResp1, err := r.client.Insight.GetOrganizationInsightMonitoredMediaServers(vvOrganizationID)
	//Has Post
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationInsightMonitoredMediaServers",
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
			vvMonitoredMediaServerID, ok := result2["ID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter MonitoredMediaServerID",
					"Fail Parsing MonitoredMediaServerID",
				)
				return
			}
			r.client.Insight.UpdateOrganizationInsightMonitoredMediaServer(vvOrganizationID, vvMonitoredMediaServerID, data.toSdkApiRequestUpdate(ctx))

			responseVerifyItem2, _, _ := r.client.Insight.GetOrganizationInsightMonitoredMediaServer(vvOrganizationID, vvMonitoredMediaServerID)
			if responseVerifyItem2 != nil {
				data = ResponseInsightGetOrganizationInsightMonitoredMediaServerItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Insight.CreateOrganizationInsightMonitoredMediaServer(vvOrganizationID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationInsightMonitoredMediaServer",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationInsightMonitoredMediaServer",
			err.Error(),
		)
		return
	}

	responseGet, restyResp1, err := r.client.Insight.GetOrganizationInsightMonitoredMediaServers(vvOrganizationID)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationInsightMonitoredMediaServers",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationInsightMonitoredMediaServers",
			err.Error(),
		)
		return
	}

	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvMonitoredMediaServerID, ok := result2["ID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter MonitoredMediaServerID",
				"Fail Parsing MonitoredMediaServerID",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Insight.GetOrganizationInsightMonitoredMediaServer(vvOrganizationID, vvMonitoredMediaServerID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseInsightGetOrganizationInsightMonitoredMediaServerItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationInsightMonitoredMediaServer",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationInsightMonitoredMediaServer",
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

func (r *OrganizationsInsightMonitoredMediaServersResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsInsightMonitoredMediaServersRs

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
	vvMonitoredMediaServerID := data.MonitoredMediaServerID.ValueString()
	responseGet, restyRespGet, err := r.client.Insight.GetOrganizationInsightMonitoredMediaServer(vvOrganizationID, vvMonitoredMediaServerID)
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
				"Failure when executing GetOrganizationInsightMonitoredMediaServer",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationInsightMonitoredMediaServer",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseInsightGetOrganizationInsightMonitoredMediaServerItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *OrganizationsInsightMonitoredMediaServersResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("monitered_media_server_id"), idParts[1])...)
}

func (r *OrganizationsInsightMonitoredMediaServersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsInsightMonitoredMediaServersRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	vvMonitoredMediaServerID := data.MonitoredMediaServerID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Insight.UpdateOrganizationInsightMonitoredMediaServer(vvOrganizationID, vvMonitoredMediaServerID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationInsightMonitoredMediaServer",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationInsightMonitoredMediaServer",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsInsightMonitoredMediaServersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsInsightMonitoredMediaServersRs
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
	vvMonitoredMediaServerID := state.MonitoredMediaServerID.ValueString()
	_, err := r.client.Insight.DeleteOrganizationInsightMonitoredMediaServer(vvOrganizationID, vvMonitoredMediaServerID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationInsightMonitoredMediaServer", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsInsightMonitoredMediaServersRs struct {
	OrganizationID              types.String `tfsdk:"organization_id"`
	MonitoredMediaServerID      types.String `tfsdk:"monitored_media_server_id"`
	Address                     types.String `tfsdk:"address"`
	BestEffortMonitoringEnabled types.Bool   `tfsdk:"best_effort_monitoring_enabled"`
	ID                          types.String `tfsdk:"id"`
	Name                        types.String `tfsdk:"name"`
}

// FromBody
func (r *OrganizationsInsightMonitoredMediaServersRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestInsightCreateOrganizationInsightMonitoredMediaServer {
	emptyString := ""
	address := new(string)
	if !r.Address.IsUnknown() && !r.Address.IsNull() {
		*address = r.Address.ValueString()
	} else {
		address = &emptyString
	}
	bestEffortMonitoringEnabled := new(bool)
	if !r.BestEffortMonitoringEnabled.IsUnknown() && !r.BestEffortMonitoringEnabled.IsNull() {
		*bestEffortMonitoringEnabled = r.BestEffortMonitoringEnabled.ValueBool()
	} else {
		bestEffortMonitoringEnabled = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	out := merakigosdk.RequestInsightCreateOrganizationInsightMonitoredMediaServer{
		Address:                     *address,
		BestEffortMonitoringEnabled: bestEffortMonitoringEnabled,
		Name:                        *name,
	}
	return &out
}
func (r *OrganizationsInsightMonitoredMediaServersRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestInsightUpdateOrganizationInsightMonitoredMediaServer {
	emptyString := ""
	address := new(string)
	if !r.Address.IsUnknown() && !r.Address.IsNull() {
		*address = r.Address.ValueString()
	} else {
		address = &emptyString
	}
	bestEffortMonitoringEnabled := new(bool)
	if !r.BestEffortMonitoringEnabled.IsUnknown() && !r.BestEffortMonitoringEnabled.IsNull() {
		*bestEffortMonitoringEnabled = r.BestEffortMonitoringEnabled.ValueBool()
	} else {
		bestEffortMonitoringEnabled = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	out := merakigosdk.RequestInsightUpdateOrganizationInsightMonitoredMediaServer{
		Address:                     *address,
		BestEffortMonitoringEnabled: bestEffortMonitoringEnabled,
		Name:                        *name,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseInsightGetOrganizationInsightMonitoredMediaServerItemToBodyRs(state OrganizationsInsightMonitoredMediaServersRs, response *merakigosdk.ResponseInsightGetOrganizationInsightMonitoredMediaServer, is_read bool) OrganizationsInsightMonitoredMediaServersRs {
	itemState := OrganizationsInsightMonitoredMediaServersRs{
		Address: types.StringValue(response.Address),
		BestEffortMonitoringEnabled: func() types.Bool {
			if response.BestEffortMonitoringEnabled != nil {
				return types.BoolValue(*response.BestEffortMonitoringEnabled)
			}
			return types.Bool{}
		}(),
		ID:   types.StringValue(response.ID),
		Name: types.StringValue(response.Name),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsInsightMonitoredMediaServersRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsInsightMonitoredMediaServersRs)
}
