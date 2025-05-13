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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesUpdateResource{}
	_ resource.ResourceWithConfigure = &WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesUpdateResource{}
)

func NewWirelessWirelessSSIDsFirewallIsolationAllowlistEntriesUpdateResource() resource.Resource {
	return &WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesUpdateResource{}
}

type WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesUpdateResource struct {
	client *merakigosdk.Client
}

func (r *WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesUpdateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesUpdateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_Wireless_wireless_ssids_firewall_isolation_allowlist_entries_update"
}

// resourceAction
func (r *WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesUpdateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"entry_id": schema.StringAttribute{
				MarkdownDescription: `entryId path parameter. Entry ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
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
				},
			},
		},
	}
}
func (r *WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesUpdateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesUpdate

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
	vvEntryID := data.EntryID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp1, err := r.client.Wireless.UpdateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntry(vvOrganizationID, vvEntryID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntry",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntry",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseWirelessUpdateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntryItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesUpdateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesUpdateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesUpdateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesUpdate struct {
	OrganizationID types.String                                                                     `tfsdk:"organization_id"`
	EntryID        types.String                                                                     `tfsdk:"entry_id"`
	Item           *ResponseWirelessUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry  `tfsdk:"item"`
	Parameters     *RequestWirelessUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryRs `tfsdk:"parameters"`
}

type ResponseWirelessUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry struct {
	Client        *ResponseWirelessUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryClient  `tfsdk:"client"`
	CreatedAt     types.String                                                                           `tfsdk:"created_at"`
	Description   types.String                                                                           `tfsdk:"description"`
	EntryID       types.String                                                                           `tfsdk:"entry_id"`
	LastUpdatedAt types.String                                                                           `tfsdk:"last_updated_at"`
	Network       *ResponseWirelessUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryNetwork `tfsdk:"network"`
	SSID          *ResponseWirelessUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntrySsid    `tfsdk:"ssid"`
}

type ResponseWirelessUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryClient struct {
	Mac types.String `tfsdk:"mac"`
}

type ResponseWirelessUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseWirelessUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntrySsid struct {
	ID     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Number types.Int64  `tfsdk:"number"`
}

type RequestWirelessUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryRs struct {
	Client      *RequestWirelessUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryClientRs `tfsdk:"client"`
	Description types.String                                                                           `tfsdk:"description"`
}

type RequestWirelessUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryClientRs struct {
	Mac types.String `tfsdk:"mac"`
}

// FromBody
func (r *WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesUpdate) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntry {
	emptyString := ""
	re := *r.Parameters
	var requestWirelessUpdateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntryClient *merakigosdk.RequestWirelessUpdateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntryClient

	if re.Client != nil {
		mac := re.Client.Mac.ValueString()
		requestWirelessUpdateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntryClient = &merakigosdk.RequestWirelessUpdateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntryClient{
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
	out := merakigosdk.RequestWirelessUpdateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntry{
		Client:      requestWirelessUpdateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntryClient,
		Description: *description,
	}
	return &out
}

// ToBody
func ResponseWirelessUpdateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntryItemToBody(state WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesUpdate, response *merakigosdk.ResponseWirelessUpdateOrganizationWirelessSSIDsFirewallIsolationAllowlistEntry) WirelessWirelessSSIDsFirewallIsolationAllowlistEntriesUpdate {
	itemState := ResponseWirelessUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntry{
		Client: func() *ResponseWirelessUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryClient {
			if response.Client != nil {
				return &ResponseWirelessUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryClient{
					Mac: types.StringValue(response.Client.Mac),
				}
			}
			return nil
		}(),
		CreatedAt:     types.StringValue(response.CreatedAt),
		Description:   types.StringValue(response.Description),
		EntryID:       types.StringValue(response.EntryID),
		LastUpdatedAt: types.StringValue(response.LastUpdatedAt),
		Network: func() *ResponseWirelessUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryNetwork {
			if response.Network != nil {
				return &ResponseWirelessUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntryNetwork{
					ID:   types.StringValue(response.Network.ID),
					Name: types.StringValue(response.Network.Name),
				}
			}
			return nil
		}(),
		SSID: func() *ResponseWirelessUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntrySsid {
			if response.SSID != nil {
				return &ResponseWirelessUpdateOrganizationWirelessSsidsFirewallIsolationAllowlistEntrySsid{
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
