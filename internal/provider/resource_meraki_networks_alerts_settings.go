package provider

// RESOURCE NORMAL
import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksAlertsSettingsResource{}
	_ resource.ResourceWithConfigure = &NetworksAlertsSettingsResource{}
)

func NewNetworksAlertsSettingsResource() resource.Resource {
	return &NetworksAlertsSettingsResource{}
}

type NetworksAlertsSettingsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksAlertsSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksAlertsSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_alerts_settings"
}

func (r *NetworksAlertsSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"alerts_response": schema.SetNestedAttribute{
				MarkdownDescription: `Alert-specific configuration for each type. Only alerts that pertain to the network can be updated.`,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"alert_destinations": schema.SingleNestedAttribute{
							MarkdownDescription: `A hash of destinations for this specific alert`,
							Computed:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"all_admins": schema.BoolAttribute{
									MarkdownDescription: `If true, then all network admins will receive emails for this alert`,
									Computed:            true,
									PlanModifiers: []planmodifier.Bool{
										boolplanmodifier.UseStateForUnknown(),
									},
								},
								"emails": schema.SetAttribute{
									MarkdownDescription: `A list of emails that will receive information about the alert`,
									Computed:            true,
									PlanModifiers: []planmodifier.Set{
										setplanmodifier.UseStateForUnknown(),
									},

									ElementType: types.StringType,
								},
								"http_server_ids": schema.SetAttribute{
									MarkdownDescription: `A list of HTTP server IDs to send a Webhook to for this alert`,
									Computed:            true,
									PlanModifiers: []planmodifier.Set{
										setplanmodifier.UseStateForUnknown(),
									},
									Default:     setdefault.StaticValue(types.SetNull(types.StringType)),
									ElementType: types.StringType,
								},
								"snmp": schema.BoolAttribute{
									MarkdownDescription: `If true, then an SNMP trap will be sent for this alert if there is an SNMP trap server configured for this network`,
									Computed:            true,
									PlanModifiers: []planmodifier.Bool{
										boolplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: `A boolean depicting if the alert is turned on or off`,
							Computed:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
						"filters": schema.SingleNestedAttribute{
							MarkdownDescription: `A hash of specific configuration data for the alert. Only filters specific to the alert will be updated.`,
							Computed:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"timeout": schema.Int64Attribute{
									Computed: true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"period": schema.Int64Attribute{
									Computed: true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"threshold": schema.Int64Attribute{
									Computed: true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `The type of alert`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"alerts": schema.SetNestedAttribute{
				MarkdownDescription: `Alert-specific configuration for each type. Only alerts that pertain to the network can be updated.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"alert_destinations": schema.SingleNestedAttribute{
							MarkdownDescription: `A hash of destinations for this specific alert`,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"all_admins": schema.BoolAttribute{
									MarkdownDescription: `If true, then all network admins will receive emails for this alert`,
									Optional:            true,
									PlanModifiers: []planmodifier.Bool{
										boolplanmodifier.UseStateForUnknown(),
									},
								},
								"emails": schema.SetAttribute{
									MarkdownDescription: `A list of emails that will receive information about the alert`,
									Optional:            true,
									PlanModifiers: []planmodifier.Set{
										setplanmodifier.UseStateForUnknown(),
									},

									ElementType: types.StringType,
								},
								"http_server_ids": schema.SetAttribute{
									MarkdownDescription: `A list of HTTP server IDs to send a Webhook to for this alert`,
									Optional:            true,
									PlanModifiers: []planmodifier.Set{
										setplanmodifier.UseStateForUnknown(),
									},
									ElementType: types.StringType,
								},
								"snmp": schema.BoolAttribute{
									MarkdownDescription: `If true, then an SNMP trap will be sent for this alert if there is an SNMP trap server configured for this network`,
									Optional:            true,
									PlanModifiers: []planmodifier.Bool{
										boolplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: `A boolean depicting if the alert is turned on or off`,
							Optional:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
						"filters": schema.SingleNestedAttribute{
							MarkdownDescription: `A hash of specific configuration data for the alert. Only filters specific to the alert will be updated.`,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"timeout": schema.Int64Attribute{
									Optional: true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"period": schema.Int64Attribute{
									Optional: true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"threshold": schema.Int64Attribute{
									Optional: true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `The type of alert`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"default_destinations": schema.SingleNestedAttribute{
				MarkdownDescription: `The network-wide destinations for all alerts on the network.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"all_admins": schema.BoolAttribute{
						MarkdownDescription: `If true, then all network admins will receive emails.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"emails": schema.SetAttribute{
						MarkdownDescription: `A list of emails that will receive the alert(s).`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						Default:     setdefault.StaticValue(types.SetNull(types.StringType)),
						ElementType: types.StringType,
					},
					"http_server_ids": schema.SetAttribute{
						MarkdownDescription: `A list of HTTP server IDs to send a Webhook to`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						Default:     setdefault.StaticValue(types.SetNull(types.StringType)),
						ElementType: types.StringType,
					},
					"snmp": schema.BoolAttribute{
						MarkdownDescription: `If true, then an SNMP trap will be sent if there is an SNMP trap server configured for this network.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Default: booldefault.StaticBool(false),
					},
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
		},
	}
}

func (r *NetworksAlertsSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksAlertsSettingsRs

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
	//Item
	responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkAlertsSettings(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksAlertsSettings only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksAlertsSettings only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Networks.UpdateNetworkAlertsSettings(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkAlertsSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkAlertsSettings",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Networks.GetNetworkAlertsSettings(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkAlertsSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkAlertsSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkAlertsSettingsItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksAlertsSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksAlertsSettingsRs

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
	responseGet, restyRespGet, err := r.client.Networks.GetNetworkAlertsSettings(vvNetworkID)
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
				"Failure when executing GetNetworkAlertsSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkAlertsSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkAlertsSettingsItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksAlertsSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksAlertsSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksAlertsSettingsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Networks.UpdateNetworkAlertsSettings(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkAlertsSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkAlertsSettings",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksAlertsSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksAlertsSettings", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksAlertsSettingsRs struct {
	NetworkID           types.String                                                   `tfsdk:"network_id"`
	Alerts              *[]ResponseNetworksGetNetworkAlertsSettingsAlertsRs            `tfsdk:"alerts"`
	AlertsResponse      *[]ResponseNetworksGetNetworkAlertsSettingsAlertsRs            `tfsdk:"alerts_response"`
	DefaultDestinations *ResponseNetworksGetNetworkAlertsSettingsDefaultDestinationsRs `tfsdk:"default_destinations"`
	Muting              *RequestNetworksUpdateNetworkAlertsSettingsMutingRs            `tfsdk:"muting"`
}

type ResponseNetworksGetNetworkAlertsSettingsAlertsRs struct {
	AlertDestinations *ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinationsRs `tfsdk:"alert_destinations"`
	Enabled           types.Bool                                                         `tfsdk:"enabled"`
	Filters           *ResponseNetworksGetNetworkAlertsSettingsAlertsFiltersRs           `tfsdk:"filters"`
	Type              types.String                                                       `tfsdk:"type"`
}

type ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinationsRs struct {
	AllAdmins     types.Bool `tfsdk:"all_admins"`
	Emails        types.Set  `tfsdk:"emails"`
	HTTPServerIDs types.Set  `tfsdk:"http_server_ids"`
	SNMP          types.Bool `tfsdk:"snmp"`
}

type ResponseNetworksGetNetworkAlertsSettingsAlertsFiltersRs struct {
	Timeout   types.Int64 `tfsdk:"timeout"`
	Threshold types.Int64 `tfsdk:"threshold"`
	Period    types.Int64 `tfsdk:"period"`
}

type ResponseNetworksGetNetworkAlertsSettingsDefaultDestinationsRs struct {
	AllAdmins     types.Bool `tfsdk:"all_admins"`
	Emails        types.Set  `tfsdk:"emails"`
	HTTPServerIDs types.Set  `tfsdk:"http_server_ids"`
	SNMP          types.Bool `tfsdk:"snmp"`
}

type RequestNetworksUpdateNetworkAlertsSettingsMutingRs struct {
	ByPortSchedules *RequestNetworksUpdateNetworkAlertsSettingsMutingByPortSchedulesRs `tfsdk:"by_port_schedules"`
}

type RequestNetworksUpdateNetworkAlertsSettingsMutingByPortSchedulesRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// FromBody
func (r *NetworksAlertsSettingsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetworkAlertsSettings {
	var requestNetworksUpdateNetworkAlertsSettingsAlerts []merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsAlerts
	if r.Alerts != nil {
		for _, rItem1 := range *r.Alerts {
			var requestNetworksUpdateNetworkAlertsSettingsAlertsAlertDestinations *merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsAlertsAlertDestinations
			if rItem1.AlertDestinations != nil {
				allAdmins := func() *bool {
					if !rItem1.AlertDestinations.AllAdmins.IsUnknown() && !rItem1.AlertDestinations.AllAdmins.IsNull() {
						return rItem1.AlertDestinations.AllAdmins.ValueBoolPointer()
					}
					return nil
				}()
				var emails []string = nil
				//Hoola aqui
				rItem1.AlertDestinations.Emails.ElementsAs(ctx, &emails, false)
				var httpServerIDs []string = nil
				//Hoola aqui
				rItem1.AlertDestinations.HTTPServerIDs.ElementsAs(ctx, &httpServerIDs, false)
				sNMP := func() *bool {
					if !rItem1.AlertDestinations.SNMP.IsUnknown() && !rItem1.AlertDestinations.SNMP.IsNull() {
						return rItem1.AlertDestinations.SNMP.ValueBoolPointer()
					}
					return nil
				}()
				requestNetworksUpdateNetworkAlertsSettingsAlertsAlertDestinations = &merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsAlertsAlertDestinations{
					AllAdmins:     allAdmins,
					Emails:        emails,
					HTTPServerIDs: httpServerIDs,
					SNMP:          sNMP,
				}
			}
			enabled := func() *bool {
				if !rItem1.Enabled.IsUnknown() && !rItem1.Enabled.IsNull() {
					return rItem1.Enabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestNetworksUpdateNetworkAlertsSettingsAlertsFilters *merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsAlertsFilters
			if rItem1.Filters != nil {
				timeout := func() *int64 {
					if !rItem1.Filters.Timeout.IsUnknown() && !rItem1.Filters.Timeout.IsNull() {
						return rItem1.Filters.Timeout.ValueInt64Pointer()
					}
					return nil
				}()
				period := func() *int64 {
					if !rItem1.Filters.Period.IsUnknown() && !rItem1.Filters.Period.IsNull() {
						return rItem1.Filters.Period.ValueInt64Pointer()
					}
					return nil
				}()
				threshold := func() *int64 {
					if !rItem1.Filters.Threshold.IsUnknown() && !rItem1.Filters.Threshold.IsNull() {
						return rItem1.Filters.Threshold.ValueInt64Pointer()
					}
					return nil
				}()
				requestNetworksUpdateNetworkAlertsSettingsAlertsFilters = &merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsAlertsFilters{
					Timeout:   int64ToIntPointer(timeout),
					Period:    int64ToIntPointer(period),
					Threshold: int64ToIntPointer(threshold),
				}
			}
			typeR := rItem1.Type.ValueString()
			requestNetworksUpdateNetworkAlertsSettingsAlerts = append(requestNetworksUpdateNetworkAlertsSettingsAlerts, merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsAlerts{
				AlertDestinations: requestNetworksUpdateNetworkAlertsSettingsAlertsAlertDestinations,
				Enabled:           enabled,
				Filters:           requestNetworksUpdateNetworkAlertsSettingsAlertsFilters,
				Type:              typeR,
			})
		}
	}
	var requestNetworksUpdateNetworkAlertsSettingsDefaultDestinations *merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsDefaultDestinations
	if r.DefaultDestinations != nil {
		allAdmins := func() *bool {
			if !r.DefaultDestinations.AllAdmins.IsUnknown() && !r.DefaultDestinations.AllAdmins.IsNull() {
				return r.DefaultDestinations.AllAdmins.ValueBoolPointer()
			}
			return nil
		}()
		var emails []string = nil
		//Hoola aqui
		r.DefaultDestinations.Emails.ElementsAs(ctx, &emails, false)
		var httpServerIDs []string = nil
		//Hoola aqui
		r.DefaultDestinations.HTTPServerIDs.ElementsAs(ctx, &httpServerIDs, false)
		sNMP := func() *bool {
			if !r.DefaultDestinations.SNMP.IsUnknown() && !r.DefaultDestinations.SNMP.IsNull() {
				return r.DefaultDestinations.SNMP.ValueBoolPointer()
			}
			return nil
		}()
		requestNetworksUpdateNetworkAlertsSettingsDefaultDestinations = &merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsDefaultDestinations{
			AllAdmins:     allAdmins,
			Emails:        emails,
			HTTPServerIDs: httpServerIDs,
			SNMP:          sNMP,
		}
	}
	var requestNetworksUpdateNetworkAlertsSettingsMuting *merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsMuting
	if r.Muting != nil {
		var requestNetworksUpdateNetworkAlertsSettingsMutingByPortSchedules *merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsMutingByPortSchedules
		if r.Muting.ByPortSchedules != nil {
			enabled := func() *bool {
				if !r.Muting.ByPortSchedules.Enabled.IsUnknown() && !r.Muting.ByPortSchedules.Enabled.IsNull() {
					return r.Muting.ByPortSchedules.Enabled.ValueBoolPointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkAlertsSettingsMutingByPortSchedules = &merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsMutingByPortSchedules{
				Enabled: enabled,
			}
		}
		requestNetworksUpdateNetworkAlertsSettingsMuting = &merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsMuting{
			ByPortSchedules: requestNetworksUpdateNetworkAlertsSettingsMutingByPortSchedules,
		}
	}
	out := merakigosdk.RequestNetworksUpdateNetworkAlertsSettings{
		Alerts: func() *[]merakigosdk.RequestNetworksUpdateNetworkAlertsSettingsAlerts {
			if len(requestNetworksUpdateNetworkAlertsSettingsAlerts) > 0 {
				return &requestNetworksUpdateNetworkAlertsSettingsAlerts
			}
			return nil
		}(),
		DefaultDestinations: requestNetworksUpdateNetworkAlertsSettingsDefaultDestinations,
		Muting:              requestNetworksUpdateNetworkAlertsSettingsMuting,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkAlertsSettingsItemToBodyRs(state NetworksAlertsSettingsRs, response *merakigosdk.ResponseNetworksGetNetworkAlertsSettings, is_read bool) NetworksAlertsSettingsRs {
	itemState := NetworksAlertsSettingsRs{
		AlertsResponse: func() *[]ResponseNetworksGetNetworkAlertsSettingsAlertsRs {
			if response.Alerts != nil {
				result := make([]ResponseNetworksGetNetworkAlertsSettingsAlertsRs, len(*response.Alerts))
				for i, alerts := range *response.Alerts {
					result[i] = ResponseNetworksGetNetworkAlertsSettingsAlertsRs{
						AlertDestinations: func() *ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinationsRs {
							if alerts.AlertDestinations != nil {
								return &ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinationsRs{
									AllAdmins: func() types.Bool {
										if alerts.AlertDestinations.AllAdmins != nil {
											return types.BoolValue(*alerts.AlertDestinations.AllAdmins)
										}
										return types.Bool{}
									}(),
									Emails:        StringSliceToSet(alerts.AlertDestinations.Emails),
									HTTPServerIDs: StringSliceToSet(alerts.AlertDestinations.HTTPServerIDs),
									SNMP: func() types.Bool {
										if alerts.AlertDestinations.SNMP != nil {
											return types.BoolValue(*alerts.AlertDestinations.SNMP)
										}
										return types.Bool{}
									}(),
								}
							}
							return &ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinationsRs{}
						}(),
						Enabled: func() types.Bool {
							if alerts.Enabled != nil {
								return types.BoolValue(*alerts.Enabled)
							}
							return types.Bool{}
						}(),
						Filters: func() *ResponseNetworksGetNetworkAlertsSettingsAlertsFiltersRs {
							if alerts.Filters != nil {
								return &ResponseNetworksGetNetworkAlertsSettingsAlertsFiltersRs{
									Timeout: func() types.Int64 {
										if alerts.Filters.Timeout != nil {
											return types.Int64Value(int64(*alerts.Filters.Timeout))
										}
										return types.Int64{}
									}(),
								}
							}
							return &ResponseNetworksGetNetworkAlertsSettingsAlertsFiltersRs{}
						}(),
						Type: types.StringValue(alerts.Type),
					}
				}
				return &result
			}
			return &[]ResponseNetworksGetNetworkAlertsSettingsAlertsRs{}
		}(),
		DefaultDestinations: func() *ResponseNetworksGetNetworkAlertsSettingsDefaultDestinationsRs {
			if response.DefaultDestinations != nil {
				return &ResponseNetworksGetNetworkAlertsSettingsDefaultDestinationsRs{
					AllAdmins: func() types.Bool {
						if response.DefaultDestinations.AllAdmins != nil {
							return types.BoolValue(*response.DefaultDestinations.AllAdmins)
						}
						return types.Bool{}
					}(),
					Emails:        StringSliceToSet(response.DefaultDestinations.Emails),
					HTTPServerIDs: StringSliceToSet(response.DefaultDestinations.HTTPServerIDs),
					SNMP: func() types.Bool {
						if response.DefaultDestinations.SNMP != nil {
							return types.BoolValue(*response.DefaultDestinations.SNMP)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkAlertsSettingsDefaultDestinationsRs{}
		}(),
	}

	// itemState.DefaultDestinations.SNMP = state.DefaultDestinations.SNMP

	itemState.Alerts = state.Alerts
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksAlertsSettingsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksAlertsSettingsRs)
}
