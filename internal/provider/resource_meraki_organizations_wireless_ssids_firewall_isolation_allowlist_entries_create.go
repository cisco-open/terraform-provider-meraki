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

// RESOURCE ACTION

import (
	"context"

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesCreateResource{}
	_ resource.ResourceWithConfigure = &WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesCreateResource{}
)

func NewWirelessWirelessSSIDsFirewallIsolationAllowlistEntriesCreateResource() resource.Resource {
	return &WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesCreateResource{}
}

type WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesCreateResource struct {
	client *merakigosdk.Client
}

func (r *WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesCreateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesCreateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_Wireless_wireless_ssids_firewall_isolation_allowlist_entries_create"
}

// resourceAction
func (r *WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesCreateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"client": schema.SingleNestedAttribute{
						MarkdownDescription: `The client of allowlist`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"mac": schema.StringAttribute{
								MarkdownDescription: `L2 Isolation mac address`,
								Computed:            true,
							},
						},
					},
					"created_at": schema.StringAttribute{
						MarkdownDescription: `Created at timestamp for the adaptive policy group`,
						Computed:            true,
					},
					"description": schema.StringAttribute{
						MarkdownDescription: `The description of mac address`,
						Computed:            true,
					},
					"entry_id": schema.StringAttribute{
						MarkdownDescription: `The id of entry`,
						Computed:            true,
					},
					"last_updated_at": schema.StringAttribute{
						MarkdownDescription: `Updated at timestamp for the adaptive policy group`,
						Computed:            true,
					},
					"network": schema.SingleNestedAttribute{
						MarkdownDescription: `The network that allowlist SSID belongs to`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"id": schema.StringAttribute{
								MarkdownDescription: `The index of network`,
								Computed:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `The name of network`,
								Computed:            true,
							},
						},
					},
					"ssid": schema.SingleNestedAttribute{
						MarkdownDescription: `The SSID that allowlist belongs to`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"id": schema.StringAttribute{
								MarkdownDescription: `The index of SSID`,
								Computed:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `The name of SSID`,
								Computed:            true,
							},
							"number": schema.Int64Attribute{
								MarkdownDescription: `The number of SSID`,
								Computed:            true,
							},
						},
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"client": schema.SingleNestedAttribute{
						MarkdownDescription: `The client of allowlist`,
						Optional:            true,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"mac": schema.StringAttribute{
								MarkdownDescription: `L2 Isolation mac address`,
								Optional:            true,
								Computed:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
							},
						},
					},
					"description": schema.StringAttribute{
						MarkdownDescription: `The description of mac address`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"network": schema.SingleNestedAttribute{
						MarkdownDescription: `The Network that allowlist belongs to`,
						Optional:            true,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"id": schema.StringAttribute{
								MarkdownDescription: `The ID of network`,
								Optional:            true,
								Computed:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
							},
						},
					},
					"ssid": schema.SingleNestedAttribute{
						MarkdownDescription: `The SSID that allowlist belongs to`,
						Optional:            true,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"number": schema.Int64Attribute{
								MarkdownDescription: `The number of SSID`,
								Optional:            true,
								Computed:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.RequiresReplace(),
								},
							},
						},
					},
				},
			},
		},
	}
}
func (r *WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesCreateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesCreate

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
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Wireless.CreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntry(vvOrganizationID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntry",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntry",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntryItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesCreateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesCreateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesCreateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesCreate struct {
	OrganizationID types.String                                                                     `tfsdk:"organization_id"`
	Item           *ResponseWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry  `tfsdk:"item"`
	Parameters     *RequestWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryRs `tfsdk:"parameters"`
}

type ResponseWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry struct {
	Client        *ResponseWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryClient  `tfsdk:"client"`
	CreatedAt     types.String                                                                           `tfsdk:"created_at"`
	Description   types.String                                                                           `tfsdk:"description"`
	EntryID       types.String                                                                           `tfsdk:"entry_id"`
	LastUpdatedAt types.String                                                                           `tfsdk:"last_updated_at"`
	Network       *ResponseWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryNetwork `tfsdk:"network"`
	SSID          *ResponseWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntrySsid    `tfsdk:"ssid"`
}

type ResponseWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryClient struct {
	Mac types.String `tfsdk:"mac"`
}

type ResponseWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntrySsid struct {
	ID     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Number types.Int64  `tfsdk:"number"`
}

type RequestWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryRs struct {
	Client      *RequestWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryClientRs  `tfsdk:"client"`
	Description types.String                                                                            `tfsdk:"description"`
	Network     *RequestWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryNetworkRs `tfsdk:"network"`
	SSID        *RequestWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntrySsidRs    `tfsdk:"ssid"`
}

type RequestWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryClientRs struct {
	Mac types.String `tfsdk:"mac"`
}

type RequestWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryNetworkRs struct {
	ID types.String `tfsdk:"id"`
}

type RequestWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntrySsidRs struct {
	Number types.Int64 `tfsdk:"number"`
}

// FromBody
func (r *WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesCreate) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntry {
	emptyString := ""
	re := *r.Parameters
	var requestWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntryClient *merakigosdk.RequestWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntryClient

	if re.Client != nil {
		mac := re.Client.Mac.ValueString()
		requestWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntryClient = &merakigosdk.RequestWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntryClient{
			Mac: mac,
		}
		//[debug] Is Array: False
	}
	description := new(string)
	if !re.Description.IsUnknown() && !re.Description.IsNull() {
		*description = re.Description.ValueString()
	} else {
		description = &emptyString
	}
	var requestWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntryNetwork *merakigosdk.RequestWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntryNetwork

	if re.Network != nil {
		id := re.Network.ID.ValueString()
		requestWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntryNetwork = &merakigosdk.RequestWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntryNetwork{
			ID: id,
		}
		//[debug] Is Array: False
	}
	var requestWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntrySSID *merakigosdk.RequestWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntrySSID

	if re.SSID != nil {
		number := func() *int64 {
			if !re.SSID.Number.IsUnknown() && !re.SSID.Number.IsNull() {
				return re.SSID.Number.ValueInt64Pointer()
			}
			return nil
		}()
		requestWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntrySSID = &merakigosdk.RequestWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntrySSID{
			Number: int64ToIntPointer(number),
		}
		//[debug] Is Array: False
	}
	out := merakigosdk.RequestWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntry{
		Client:      requestWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntryClient,
		Description: *description,
		Network:     requestWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntryNetwork,
		SSID:        requestWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntrySSID,
	}
	return &out
}

// ToBody
func ResponseWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntryItemToBody(state WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesCreate, response *merakigosdk.ResponseWirelessCreateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntry) WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesCreate {
	itemState := ResponseWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry{
		Client: func() *ResponseWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryClient {
			if response.Client != nil {
				return &ResponseWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryClient{
					Mac: types.StringValue(response.Client.Mac),
				}
			}
			return nil
		}(),
		CreatedAt:     types.StringValue(response.CreatedAt),
		Description:   types.StringValue(response.Description),
		EntryID:       types.StringValue(response.EntryID),
		LastUpdatedAt: types.StringValue(response.LastUpdatedAt),
		Network: func() *ResponseWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryNetwork {
			if response.Network != nil {
				return &ResponseWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryNetwork{
					ID:   types.StringValue(response.Network.ID),
					Name: types.StringValue(response.Network.Name),
				}
			}
			return nil
		}(),
		SSID: func() *ResponseWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntrySsid {
			if response.SSID != nil {
				return &ResponseWirelessCreateOrganizationWirelessSsidsFirewallIsolationAllowlistEntrySsid{
					ID:   types.StringValue(response.SSID.ID),
					Name: types.StringValue(response.SSID.Name),
					Number: func() types.Int64 {
						if response.SSID.Number != nil {
							return types.Int64Value(int64(*response.SSID.Number))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
