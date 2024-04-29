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
	_ datasource.DataSource              = &OrganizationsClientsSearchDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsClientsSearchDataSource{}
)

func NewOrganizationsClientsSearchDataSource() datasource.DataSource {
	return &OrganizationsClientsSearchDataSource{}
}

type OrganizationsClientsSearchDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsClientsSearchDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsClientsSearchDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_clients_search"
}

func (d *OrganizationsClientsSearchDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"mac": schema.StringAttribute{
				MarkdownDescription: `mac query parameter. The MAC address of the client. Required.`,
				Required:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 5. Default is 5.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"client_id": schema.StringAttribute{
						Computed: true,
					},
					"mac": schema.StringAttribute{
						Computed: true,
					},
					"manufacturer": schema.StringAttribute{
						Computed: true,
					},
					"records": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"cdp": schema.StringAttribute{
									Computed: true,
								},
								"client_vpn_connections": schema.SetNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"connected_at": schema.Int64Attribute{
												Computed: true,
											},
											"disconnected_at": schema.Int64Attribute{
												Computed: true,
											},
											"remote_ip": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
								"description": schema.StringAttribute{
									Computed: true,
								},
								"first_seen": schema.Int64Attribute{
									Computed: true,
								},
								"ip": schema.StringAttribute{
									Computed: true,
								},
								"ip6": schema.StringAttribute{
									Computed: true,
								},
								"last_seen": schema.Int64Attribute{
									Computed: true,
								},
								"lldp": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
								"network": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"enrollment_string": schema.StringAttribute{
											Computed: true,
										},
										"id": schema.StringAttribute{
											Computed: true,
										},
										"is_bound_to_config_template": schema.BoolAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
										"notes": schema.StringAttribute{
											Computed: true,
										},
										"organization_id": schema.StringAttribute{
											Computed: true,
										},
										"product_types": schema.ListAttribute{
											Computed:    true,
											ElementType: types.StringType,
										},
										"tags": schema.ListAttribute{
											Computed:    true,
											ElementType: types.StringType,
										},
										"time_zone": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"os": schema.StringAttribute{
									Computed: true,
								},
								"sm_installed": schema.BoolAttribute{
									Computed: true,
								},
								"ssid": schema.StringAttribute{
									Computed: true,
								},
								"status": schema.StringAttribute{
									Computed: true,
								},
								"switchport": schema.StringAttribute{
									Computed: true,
								},
								"user": schema.StringAttribute{
									Computed: true,
								},
								"vlan": schema.StringAttribute{
									Computed: true,
								},
								"wireless_capabilities": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsClientsSearchDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsClientsSearch OrganizationsClientsSearch
	diags := req.Config.Get(ctx, &organizationsClientsSearch)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationClientsSearch")
		vvOrganizationID := organizationsClientsSearch.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationClientsSearchQueryParams{}

		queryParams1.Mac = organizationsClientsSearch.Mac.ValueString()

		queryParams1.PerPage = int(organizationsClientsSearch.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsClientsSearch.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsClientsSearch.EndingBefore.ValueString()

		response1, restyResp1, err := d.client.Organizations.GetOrganizationClientsSearch(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationClientsSearch",
				err.Error(),
			)
			return
		}

		organizationsClientsSearch = ResponseOrganizationsGetOrganizationClientsSearchItemToBody(organizationsClientsSearch, response1)
		diags = resp.State.Set(ctx, &organizationsClientsSearch)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsClientsSearch struct {
	OrganizationID types.String                                       `tfsdk:"organization_id"`
	Mac            types.String                                       `tfsdk:"mac"`
	PerPage        types.Int64                                        `tfsdk:"per_page"`
	StartingAfter  types.String                                       `tfsdk:"starting_after"`
	EndingBefore   types.String                                       `tfsdk:"ending_before"`
	Item           *ResponseOrganizationsGetOrganizationClientsSearch `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationClientsSearch struct {
	ClientID     types.String                                                `tfsdk:"client_id"`
	Mac          types.String                                                `tfsdk:"mac"`
	Manufacturer types.String                                                `tfsdk:"manufacturer"`
	Records      *[]ResponseOrganizationsGetOrganizationClientsSearchRecords `tfsdk:"records"`
}

type ResponseOrganizationsGetOrganizationClientsSearchRecords struct {
	Cdp                  types.String                                                                    `tfsdk:"cdp"`
	ClientVpnConnections *[]ResponseOrganizationsGetOrganizationClientsSearchRecordsClientVpnConnections `tfsdk:"client_vpn_connections"`
	Description          types.String                                                                    `tfsdk:"description"`
	FirstSeen            types.Int64                                                                     `tfsdk:"first_seen"`
	IP                   types.String                                                                    `tfsdk:"ip"`
	IP6                  types.String                                                                    `tfsdk:"ip6"`
	LastSeen             types.Int64                                                                     `tfsdk:"last_seen"`
	Lldp                 types.List                                                                      `tfsdk:"lldp"`
	Network              *ResponseOrganizationsGetOrganizationClientsSearchRecordsNetwork                `tfsdk:"network"`
	Os                   types.String                                                                    `tfsdk:"os"`
	SmInstalled          types.Bool                                                                      `tfsdk:"sm_installed"`
	SSID                 types.String                                                                    `tfsdk:"ssid"`
	Status               types.String                                                                    `tfsdk:"status"`
	Switchport           types.String                                                                    `tfsdk:"switchport"`
	User                 types.String                                                                    `tfsdk:"user"`
	VLAN                 types.String                                                                    `tfsdk:"vlan"`
	WirelessCapabilities types.String                                                                    `tfsdk:"wireless_capabilities"`
}

type ResponseOrganizationsGetOrganizationClientsSearchRecordsClientVpnConnections struct {
	ConnectedAt    types.Int64  `tfsdk:"connected_at"`
	DisconnectedAt types.Int64  `tfsdk:"disconnected_at"`
	RemoteIP       types.String `tfsdk:"remote_ip"`
}

type ResponseOrganizationsGetOrganizationClientsSearchRecordsNetwork struct {
	EnrollmentString        types.String `tfsdk:"enrollment_string"`
	ID                      types.String `tfsdk:"id"`
	IsBoundToConfigTemplate types.Bool   `tfsdk:"is_bound_to_config_template"`
	Name                    types.String `tfsdk:"name"`
	Notes                   types.String `tfsdk:"notes"`
	OrganizationID          types.String `tfsdk:"organization_id"`
	ProductTypes            types.List   `tfsdk:"product_types"`
	Tags                    types.List   `tfsdk:"tags"`
	TimeZone                types.String `tfsdk:"time_zone"`
}

// ToBody
func ResponseOrganizationsGetOrganizationClientsSearchItemToBody(state OrganizationsClientsSearch, response *merakigosdk.ResponseOrganizationsGetOrganizationClientsSearch) OrganizationsClientsSearch {
	itemState := ResponseOrganizationsGetOrganizationClientsSearch{
		ClientID:     types.StringValue(response.ClientID),
		Mac:          types.StringValue(response.Mac),
		Manufacturer: types.StringValue(response.Manufacturer),
		Records: func() *[]ResponseOrganizationsGetOrganizationClientsSearchRecords {
			if response.Records != nil {
				result := make([]ResponseOrganizationsGetOrganizationClientsSearchRecords, len(*response.Records))
				for i, records := range *response.Records {
					result[i] = ResponseOrganizationsGetOrganizationClientsSearchRecords{
						Cdp: types.StringValue(records.Cdp),
						ClientVpnConnections: func() *[]ResponseOrganizationsGetOrganizationClientsSearchRecordsClientVpnConnections {
							if records.ClientVpnConnections != nil {
								result := make([]ResponseOrganizationsGetOrganizationClientsSearchRecordsClientVpnConnections, len(*records.ClientVpnConnections))
								for i, clientVpnConnections := range *records.ClientVpnConnections {
									result[i] = ResponseOrganizationsGetOrganizationClientsSearchRecordsClientVpnConnections{
										ConnectedAt: func() types.Int64 {
											if clientVpnConnections.ConnectedAt != nil {
												return types.Int64Value(int64(*clientVpnConnections.ConnectedAt))
											}
											return types.Int64{}
										}(),
										DisconnectedAt: func() types.Int64 {
											if clientVpnConnections.DisconnectedAt != nil {
												return types.Int64Value(int64(*clientVpnConnections.DisconnectedAt))
											}
											return types.Int64{}
										}(),
										RemoteIP: types.StringValue(clientVpnConnections.RemoteIP),
									}
								}
								return &result
							}
							return &[]ResponseOrganizationsGetOrganizationClientsSearchRecordsClientVpnConnections{}
						}(),
						Description: types.StringValue(records.Description),
						FirstSeen: func() types.Int64 {
							if records.FirstSeen != nil {
								return types.Int64Value(int64(*records.FirstSeen))
							}
							return types.Int64{}
						}(),
						IP:  types.StringValue(records.IP),
						IP6: types.StringValue(records.IP6),
						LastSeen: func() types.Int64 {
							if records.LastSeen != nil {
								return types.Int64Value(int64(*records.LastSeen))
							}
							return types.Int64{}
						}(),
						Lldp: StringSliceToList(records.Lldp),
						Network: func() *ResponseOrganizationsGetOrganizationClientsSearchRecordsNetwork {
							if records.Network != nil {
								return &ResponseOrganizationsGetOrganizationClientsSearchRecordsNetwork{
									EnrollmentString: types.StringValue(records.Network.EnrollmentString),
									ID:               types.StringValue(records.Network.ID),
									IsBoundToConfigTemplate: func() types.Bool {
										if records.Network.IsBoundToConfigTemplate != nil {
											return types.BoolValue(*records.Network.IsBoundToConfigTemplate)
										}
										return types.Bool{}
									}(),
									Name:           types.StringValue(records.Network.Name),
									Notes:          types.StringValue(records.Network.Notes),
									OrganizationID: types.StringValue(records.Network.OrganizationID),
									ProductTypes:   StringSliceToList(records.Network.ProductTypes),
									Tags:           StringSliceToList(records.Network.Tags),
									TimeZone:       types.StringValue(records.Network.TimeZone),
								}
							}
							return &ResponseOrganizationsGetOrganizationClientsSearchRecordsNetwork{}
						}(),
						Os: types.StringValue(records.Os),
						SmInstalled: func() types.Bool {
							if records.SmInstalled != nil {
								return types.BoolValue(*records.SmInstalled)
							}
							return types.Bool{}
						}(),
						SSID:                 types.StringValue(records.SSID),
						Status:               types.StringValue(records.Status),
						Switchport:           types.StringValue(records.Switchport),
						User:                 types.StringValue(records.User),
						VLAN:                 types.StringValue(records.VLAN),
						WirelessCapabilities: types.StringValue(records.WirelessCapabilities),
					}
				}
				return &result
			}
			return &[]ResponseOrganizationsGetOrganizationClientsSearchRecords{}
		}(),
	}
	state.Item = &itemState
	return state
}
