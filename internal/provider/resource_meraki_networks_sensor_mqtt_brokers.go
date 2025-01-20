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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSensorMqttBrokersResource{}
	_ resource.ResourceWithConfigure = &NetworksSensorMqttBrokersResource{}
)

func NewNetworksSensorMqttBrokersResource() resource.Resource {
	return &NetworksSensorMqttBrokersResource{}
}

type NetworksSensorMqttBrokersResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSensorMqttBrokersResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSensorMqttBrokersResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sensor_mqtt_brokers"
}

func (r *NetworksSensorMqttBrokersResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `Specifies whether the broker is enabled for sensor data. Currently, only a single broker may be enabled for sensor data.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"mqtt_broker_id": schema.StringAttribute{
				MarkdownDescription: `ID of the MQTT Broker.`,
				Required:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
		},
	}
}

//path params to set ['mqttBrokerId']

func (r *NetworksSensorMqttBrokersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSensorMqttBrokersRs

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
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvMqttBrokerID := data.MqttBrokerID.ValueString()
	//Reviw This  Has Item and item
	//Revisar

	// vvMqttBrokerID := data.MqttBrokerID.ValueString()
	if vvMqttBrokerID != "" {
		responseVerifyItem, _, err := r.client.Sensor.GetNetworkSensorMqttBroker(vvNetworkID, vvMqttBrokerID)
		if err != nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksSensorMqttBrokers only have update context, not create.",
				err.Error(),
			)
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Resource NetworksSensorMqttBrokers only have update context, not create.",
			"Path parameter MqttBrokerID expected",
		)
		return
	}

	response, restyResp1, err := r.client.Sensor.UpdateNetworkSensorMqttBroker(vvNetworkID, vvMqttBrokerID, data.toSdkApiRequestUpdate(ctx))

	if err != nil || restyResp1 == nil || response == nil {
		if restyResp1 != nil {
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
	//Item
	responseGet, restyResp1, err := r.client.Sensor.GetNetworkSensorMqttBroker(vvNetworkID, vvMqttBrokerID)
	// Has only items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSensorMqttBrokers",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSensorMqttBrokers",
			err.Error(),
		)
		return
	}
	data = ResponseSensorGetNetworkSensorMqttBrokerItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSensorMqttBrokersResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSensorMqttBrokersRs

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

	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvMqttBrokerID := data.MqttBrokerID.ValueString()
	// mqtt_broker_id
	responseGet, restyRespGet, err := r.client.Sensor.GetNetworkSensorMqttBroker(vvNetworkID, vvMqttBrokerID)
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
				"Failure when executing GetNetworkSensorMqttBroker",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSensorMqttBroker",
			err.Error(),
		)
		return
	}

	data = ResponseSensorGetNetworkSensorMqttBrokerItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksSensorMqttBrokersResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("mqtt_broker_id"), idParts[1])...)
}

func (r *NetworksSensorMqttBrokersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSensorMqttBrokersRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvMqttBrokerID := data.MqttBrokerID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Sensor.UpdateNetworkSensorMqttBroker(vvNetworkID, vvMqttBrokerID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSensorMqttBroker",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSensorMqttBroker",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSensorMqttBrokersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSensorMqttBrokersRs struct {
	NetworkID    types.String `tfsdk:"network_id"`
	MqttBrokerID types.String `tfsdk:"mqtt_broker_id"`
	Enabled      types.Bool   `tfsdk:"enabled"`
}

// FromBody
func (r *NetworksSensorMqttBrokersRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSensorUpdateNetworkSensorMqttBroker {
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	out := merakigosdk.RequestSensorUpdateNetworkSensorMqttBroker{
		Enabled: enabled,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSensorGetNetworkSensorMqttBrokerItemToBodyRs(state NetworksSensorMqttBrokersRs, response *merakigosdk.ResponseSensorGetNetworkSensorMqttBroker, is_read bool) NetworksSensorMqttBrokersRs {
	itemState := NetworksSensorMqttBrokersRs{
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		MqttBrokerID: types.StringValue(response.MqttBrokerID),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSensorMqttBrokersRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSensorMqttBrokersRs)
}
