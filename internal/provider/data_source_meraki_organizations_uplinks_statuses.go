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
	_ datasource.DataSource              = &OrganizationsUplinksStatusesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsUplinksStatusesDataSource{}
)

func NewOrganizationsUplinksStatusesDataSource() datasource.DataSource {
	return &OrganizationsUplinksStatusesDataSource{}
}

type OrganizationsUplinksStatusesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsUplinksStatusesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsUplinksStatusesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_uplinks_statuses"
}

func (d *OrganizationsUplinksStatusesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"iccids": schema.ListAttribute{
				MarkdownDescription: `iccids query parameter. A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. A list of network IDs. The returned devices will be filtered to only include these networks.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.`,
				Optional:            true,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. A list of serial numbers. The returned devices will be filtered to only include these serials.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationUplinksStatuses`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"high_availability": schema.SingleNestedAttribute{
							MarkdownDescription: `Device High Availability Capabilities`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"enabled": schema.BoolAttribute{
									MarkdownDescription: `Indicates whether High Availability is enabled for the device. For devices that do not support HA, this will be 'false'`,
									Computed:            true,
								},
								"role": schema.StringAttribute{
									MarkdownDescription: `The HA role of the device on the network. For devices that do not support HA, this will be 'primary'`,
									Computed:            true,
								},
							},
						},
						"last_reported_at": schema.StringAttribute{
							MarkdownDescription: `Last reported time for the device`,
							Computed:            true,
						},
						"model": schema.StringAttribute{
							MarkdownDescription: `The uplink model`,
							Computed:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: `Network identifier`,
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `The uplink serial`,
							Computed:            true,
						},
						"uplinks": schema.SetNestedAttribute{
							MarkdownDescription: `Uplinks`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"apn": schema.StringAttribute{
										MarkdownDescription: `Access Point Name`,
										Computed:            true,
									},
									"connection_type": schema.StringAttribute{
										MarkdownDescription: `Connection Type`,
										Computed:            true,
									},
									"dns1": schema.StringAttribute{
										MarkdownDescription: `Primary DNS IP`,
										Computed:            true,
									},
									"dns2": schema.StringAttribute{
										MarkdownDescription: `Secondary DNS IP`,
										Computed:            true,
									},
									"gateway": schema.StringAttribute{
										MarkdownDescription: `Gateway IP`,
										Computed:            true,
									},
									"iccid": schema.StringAttribute{
										MarkdownDescription: `Integrated Circuit Card Identification Number`,
										Computed:            true,
									},
									"interface": schema.StringAttribute{
										MarkdownDescription: `Uplink interface`,
										Computed:            true,
									},
									"ip": schema.StringAttribute{
										MarkdownDescription: `Uplink IP`,
										Computed:            true,
									},
									"ip_assigned_by": schema.StringAttribute{
										MarkdownDescription: `The way in which the IP is assigned`,
										Computed:            true,
									},
									"primary_dns": schema.StringAttribute{
										MarkdownDescription: `Primary DNS IP`,
										Computed:            true,
									},
									"provider": schema.StringAttribute{
										MarkdownDescription: `Network Provider`,
										Computed:            true,
									},
									"public_ip": schema.StringAttribute{
										MarkdownDescription: `Public IP`,
										Computed:            true,
									},
									"secondary_dns": schema.StringAttribute{
										MarkdownDescription: `Secondary DNS IP`,
										Computed:            true,
									},
									"signal_stat": schema.SingleNestedAttribute{
										MarkdownDescription: `Tower Signal Status`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"rsrp": schema.StringAttribute{
												MarkdownDescription: `Reference Signal Received Power`,
												Computed:            true,
											},
											"rsrq": schema.StringAttribute{
												MarkdownDescription: `Reference Signal Received Quality`,
												Computed:            true,
											},
										},
									},
									"signal_type": schema.StringAttribute{
										MarkdownDescription: `Signal Type`,
										Computed:            true,
									},
									"status": schema.StringAttribute{
										MarkdownDescription: `Uplink status`,
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsUplinksStatusesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsUplinksStatuses OrganizationsUplinksStatuses
	diags := req.Config.Get(ctx, &organizationsUplinksStatuses)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationUplinksStatuses")
		vvOrganizationID := organizationsUplinksStatuses.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationUplinksStatusesQueryParams{}

		queryParams1.PerPage = int(organizationsUplinksStatuses.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsUplinksStatuses.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsUplinksStatuses.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsUplinksStatuses.NetworkIDs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsUplinksStatuses.Serials)
		queryParams1.Iccids = elementsToStrings(ctx, organizationsUplinksStatuses.Iccids)

		response1, restyResp1, err := d.client.Organizations.GetOrganizationUplinksStatuses(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationUplinksStatuses",
				err.Error(),
			)
			return
		}

		organizationsUplinksStatuses = ResponseOrganizationsGetOrganizationUplinksStatusesItemsToBody(organizationsUplinksStatuses, response1)
		diags = resp.State.Set(ctx, &organizationsUplinksStatuses)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsUplinksStatuses struct {
	OrganizationID types.String                                               `tfsdk:"organization_id"`
	PerPage        types.Int64                                                `tfsdk:"per_page"`
	StartingAfter  types.String                                               `tfsdk:"starting_after"`
	EndingBefore   types.String                                               `tfsdk:"ending_before"`
	NetworkIDs     types.List                                                 `tfsdk:"network_ids"`
	Serials        types.List                                                 `tfsdk:"serials"`
	Iccids         types.List                                                 `tfsdk:"iccids"`
	Items          *[]ResponseItemOrganizationsGetOrganizationUplinksStatuses `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationUplinksStatuses struct {
	HighAvailability *ResponseItemOrganizationsGetOrganizationUplinksStatusesHighAvailability `tfsdk:"high_availability"`
	LastReportedAt   types.String                                                             `tfsdk:"last_reported_at"`
	Model            types.String                                                             `tfsdk:"model"`
	NetworkID        types.String                                                             `tfsdk:"network_id"`
	Serial           types.String                                                             `tfsdk:"serial"`
	Uplinks          *[]ResponseItemOrganizationsGetOrganizationUplinksStatusesUplinks        `tfsdk:"uplinks"`
}

type ResponseItemOrganizationsGetOrganizationUplinksStatusesHighAvailability struct {
	Enabled types.Bool   `tfsdk:"enabled"`
	Role    types.String `tfsdk:"role"`
}

type ResponseItemOrganizationsGetOrganizationUplinksStatusesUplinks struct {
	Apn            types.String                                                              `tfsdk:"apn"`
	ConnectionType types.String                                                              `tfsdk:"connection_type"`
	DNS1           types.String                                                              `tfsdk:"dns1"`
	DNS2           types.String                                                              `tfsdk:"dns2"`
	Gateway        types.String                                                              `tfsdk:"gateway"`
	Iccid          types.String                                                              `tfsdk:"iccid"`
	Interface      types.String                                                              `tfsdk:"interface"`
	IP             types.String                                                              `tfsdk:"ip"`
	IPAssignedBy   types.String                                                              `tfsdk:"ip_assigned_by"`
	PrimaryDNS     types.String                                                              `tfsdk:"primary_dns"`
	Provider       types.String                                                              `tfsdk:"provider"`
	PublicIP       types.String                                                              `tfsdk:"public_ip"`
	SecondaryDNS   types.String                                                              `tfsdk:"secondary_dns"`
	SignalStat     *ResponseItemOrganizationsGetOrganizationUplinksStatusesUplinksSignalStat `tfsdk:"signal_stat"`
	SignalType     types.String                                                              `tfsdk:"signal_type"`
	Status         types.String                                                              `tfsdk:"status"`
}

type ResponseItemOrganizationsGetOrganizationUplinksStatusesUplinksSignalStat struct {
	Rsrp types.String `tfsdk:"rsrp"`
	Rsrq types.String `tfsdk:"rsrq"`
}

// ToBody
func ResponseOrganizationsGetOrganizationUplinksStatusesItemsToBody(state OrganizationsUplinksStatuses, response *merakigosdk.ResponseOrganizationsGetOrganizationUplinksStatuses) OrganizationsUplinksStatuses {
	var items []ResponseItemOrganizationsGetOrganizationUplinksStatuses
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationUplinksStatuses{
			HighAvailability: func() *ResponseItemOrganizationsGetOrganizationUplinksStatusesHighAvailability {
				if item.HighAvailability != nil {
					return &ResponseItemOrganizationsGetOrganizationUplinksStatusesHighAvailability{
						Enabled: func() types.Bool {
							if item.HighAvailability.Enabled != nil {
								return types.BoolValue(*item.HighAvailability.Enabled)
							}
							return types.Bool{}
						}(),
						Role: types.StringValue(item.HighAvailability.Role),
					}
				}
				return &ResponseItemOrganizationsGetOrganizationUplinksStatusesHighAvailability{}
			}(),
			LastReportedAt: types.StringValue(item.LastReportedAt),
			Model:          types.StringValue(item.Model),
			NetworkID:      types.StringValue(item.NetworkID),
			Serial:         types.StringValue(item.Serial),
			Uplinks: func() *[]ResponseItemOrganizationsGetOrganizationUplinksStatusesUplinks {
				if item.Uplinks != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationUplinksStatusesUplinks, len(*item.Uplinks))
					for i, uplinks := range *item.Uplinks {
						result[i] = ResponseItemOrganizationsGetOrganizationUplinksStatusesUplinks{
							Apn:            types.StringValue(uplinks.Apn),
							ConnectionType: types.StringValue(uplinks.ConnectionType),
							DNS1:           types.StringValue(uplinks.DNS1),
							DNS2:           types.StringValue(uplinks.DNS2),
							Gateway:        types.StringValue(uplinks.Gateway),
							Iccid:          types.StringValue(uplinks.Iccid),
							Interface:      types.StringValue(uplinks.Interface),
							IP:             types.StringValue(uplinks.IP),
							IPAssignedBy:   types.StringValue(uplinks.IPAssignedBy),
							PrimaryDNS:     types.StringValue(uplinks.PrimaryDNS),
							Provider:       types.StringValue(uplinks.Provider),
							PublicIP:       types.StringValue(uplinks.PublicIP),
							SecondaryDNS:   types.StringValue(uplinks.SecondaryDNS),
							SignalStat: func() *ResponseItemOrganizationsGetOrganizationUplinksStatusesUplinksSignalStat {
								if uplinks.SignalStat != nil {
									return &ResponseItemOrganizationsGetOrganizationUplinksStatusesUplinksSignalStat{
										Rsrp: types.StringValue(uplinks.SignalStat.Rsrp),
										Rsrq: types.StringValue(uplinks.SignalStat.Rsrq),
									}
								}
								return &ResponseItemOrganizationsGetOrganizationUplinksStatusesUplinksSignalStat{}
							}(),
							SignalType: types.StringValue(uplinks.SignalType),
							Status:     types.StringValue(uplinks.Status),
						}
					}
					return &result
				}
				return &[]ResponseItemOrganizationsGetOrganizationUplinksStatusesUplinks{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
