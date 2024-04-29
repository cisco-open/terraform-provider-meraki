package provider

// DATA SOURCE NORMAL
import (
	"context"
	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksAlertsSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksAlertsSettingsDataSource{}
)

func NewNetworksAlertsSettingsDataSource() datasource.DataSource {
	return &NetworksAlertsSettingsDataSource{}
}

type NetworksAlertsSettingsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksAlertsSettingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksAlertsSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_alerts_settings"
}

func (d *NetworksAlertsSettingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"alerts": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"alert_destinations": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"all_admins": schema.BoolAttribute{
											Computed: true,
										},
										"emails": schema.ListAttribute{
											Computed:    true,
											ElementType: types.StringType,
										},
										"http_server_ids": schema.ListAttribute{
											Computed:    true,
											ElementType: types.StringType,
										},
										"snmp": schema.BoolAttribute{
											Computed: true,
										},
									},
								},
								"enabled": schema.BoolAttribute{
									Computed: true,
								},
								"filters": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"timeout": schema.Int64Attribute{
											Computed: true,
										},
									},
								},
								"type": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"default_destinations": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"all_admins": schema.BoolAttribute{
								Computed: true,
							},
							"emails": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"http_server_ids": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"snmp": schema.BoolAttribute{
								Computed: true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksAlertsSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksAlertsSettings NetworksAlertsSettings
	diags := req.Config.Get(ctx, &networksAlertsSettings)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkAlertsSettings")
		vvNetworkID := networksAlertsSettings.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Networks.GetNetworkAlertsSettings(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkAlertsSettings",
				err.Error(),
			)
			return
		}

		networksAlertsSettings = ResponseNetworksGetNetworkAlertsSettingsItemToBody(networksAlertsSettings, response1)
		diags = resp.State.Set(ctx, &networksAlertsSettings)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksAlertsSettings struct {
	NetworkID types.String                              `tfsdk:"network_id"`
	Item      *ResponseNetworksGetNetworkAlertsSettings `tfsdk:"item"`
}

type ResponseNetworksGetNetworkAlertsSettings struct {
	Alerts              *[]ResponseNetworksGetNetworkAlertsSettingsAlerts            `tfsdk:"alerts"`
	DefaultDestinations *ResponseNetworksGetNetworkAlertsSettingsDefaultDestinations `tfsdk:"default_destinations"`
}

type ResponseNetworksGetNetworkAlertsSettingsAlerts struct {
	AlertDestinations *ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinations `tfsdk:"alert_destinations"`
	Enabled           types.Bool                                                       `tfsdk:"enabled"`
	Filters           *ResponseNetworksGetNetworkAlertsSettingsAlertsFilters           `tfsdk:"filters"`
	Type              types.String                                                     `tfsdk:"type"`
}

type ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinations struct {
	AllAdmins     types.Bool `tfsdk:"all_admins"`
	Emails        types.List `tfsdk:"emails"`
	HTTPServerIDs types.List `tfsdk:"http_server_ids"`
	SNMP          types.Bool `tfsdk:"snmp"`
}

type ResponseNetworksGetNetworkAlertsSettingsAlertsFilters struct {
	Timeout types.Int64 `tfsdk:"timeout"`
}

type ResponseNetworksGetNetworkAlertsSettingsDefaultDestinations struct {
	AllAdmins     types.Bool `tfsdk:"all_admins"`
	Emails        types.List `tfsdk:"emails"`
	HTTPServerIDs types.List `tfsdk:"http_server_ids"`
	SNMP          types.Bool `tfsdk:"snmp"`
}

// ToBody
func ResponseNetworksGetNetworkAlertsSettingsItemToBody(state NetworksAlertsSettings, response *merakigosdk.ResponseNetworksGetNetworkAlertsSettings) NetworksAlertsSettings {
	itemState := ResponseNetworksGetNetworkAlertsSettings{
		Alerts: func() *[]ResponseNetworksGetNetworkAlertsSettingsAlerts {
			if response.Alerts != nil {
				result := make([]ResponseNetworksGetNetworkAlertsSettingsAlerts, len(*response.Alerts))
				for i, alerts := range *response.Alerts {
					result[i] = ResponseNetworksGetNetworkAlertsSettingsAlerts{
						AlertDestinations: func() *ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinations {
							if alerts.AlertDestinations != nil {
								return &ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinations{
									AllAdmins: func() types.Bool {
										if alerts.AlertDestinations.AllAdmins != nil {
											return types.BoolValue(*alerts.AlertDestinations.AllAdmins)
										}
										return types.Bool{}
									}(),
									Emails:        StringSliceToList(alerts.AlertDestinations.Emails),
									HTTPServerIDs: StringSliceToList(alerts.AlertDestinations.HTTPServerIDs),
									SNMP: func() types.Bool {
										if alerts.AlertDestinations.SNMP != nil {
											return types.BoolValue(*alerts.AlertDestinations.SNMP)
										}
										return types.Bool{}
									}(),
								}
							}
							return &ResponseNetworksGetNetworkAlertsSettingsAlertsAlertDestinations{}
						}(),
						Enabled: func() types.Bool {
							if alerts.Enabled != nil {
								return types.BoolValue(*alerts.Enabled)
							}
							return types.Bool{}
						}(),
						Filters: func() *ResponseNetworksGetNetworkAlertsSettingsAlertsFilters {
							if alerts.Filters != nil {
								return &ResponseNetworksGetNetworkAlertsSettingsAlertsFilters{
									Timeout: func() types.Int64 {
										if alerts.Filters.Timeout != nil {
											return types.Int64Value(int64(*alerts.Filters.Timeout))
										}
										return types.Int64{}
									}(),
								}
							}
							return &ResponseNetworksGetNetworkAlertsSettingsAlertsFilters{}
						}(),
						Type: types.StringValue(alerts.Type),
					}
				}
				return &result
			}
			return &[]ResponseNetworksGetNetworkAlertsSettingsAlerts{}
		}(),
		DefaultDestinations: func() *ResponseNetworksGetNetworkAlertsSettingsDefaultDestinations {
			if response.DefaultDestinations != nil {
				return &ResponseNetworksGetNetworkAlertsSettingsDefaultDestinations{
					AllAdmins: func() types.Bool {
						if response.DefaultDestinations.AllAdmins != nil {
							return types.BoolValue(*response.DefaultDestinations.AllAdmins)
						}
						return types.Bool{}
					}(),
					Emails:        StringSliceToList(response.DefaultDestinations.Emails),
					HTTPServerIDs: StringSliceToList(response.DefaultDestinations.HTTPServerIDs),
					SNMP: func() types.Bool {
						if response.DefaultDestinations.SNMP != nil {
							return types.BoolValue(*response.DefaultDestinations.SNMP)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkAlertsSettingsDefaultDestinations{}
		}(),
	}
	state.Item = &itemState
	return state
}
