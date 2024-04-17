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
	_ datasource.DataSource              = &NetworksSmTrustedAccessConfigsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSmTrustedAccessConfigsDataSource{}
)

func NewNetworksSmTrustedAccessConfigsDataSource() datasource.DataSource {
	return &NetworksSmTrustedAccessConfigsDataSource{}
}

type NetworksSmTrustedAccessConfigsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSmTrustedAccessConfigsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSmTrustedAccessConfigsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_trusted_access_configs"
}

func (d *NetworksSmTrustedAccessConfigsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 100.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSmGetNetworkSmTrustedAccessConfigs`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"access_end_at": schema.StringAttribute{
							MarkdownDescription: `time that access ends`,
							Computed:            true,
						},
						"access_start_at": schema.StringAttribute{
							MarkdownDescription: `time that access starts`,
							Computed:            true,
						},
						"additional_email_text": schema.StringAttribute{
							MarkdownDescription: `Optional email text`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `device ID`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `device name`,
							Computed:            true,
						},
						"notify_time_before_access_ends": schema.Int64Attribute{
							MarkdownDescription: `Time before access expiration reminder email sends`,
							Computed:            true,
						},
						"scope": schema.StringAttribute{
							MarkdownDescription: `scope`,
							Computed:            true,
						},
						"send_expiration_emails": schema.BoolAttribute{
							MarkdownDescription: `Send Email Notifications`,
							Computed:            true,
						},
						"ssid_name": schema.StringAttribute{
							MarkdownDescription: `SSID name`,
							Computed:            true,
						},
						"tags": schema.ListAttribute{
							MarkdownDescription: `device tags`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"timebound_type": schema.StringAttribute{
							MarkdownDescription: `type of access period, either a static range or a dynamic period`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSmTrustedAccessConfigsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSmTrustedAccessConfigs NetworksSmTrustedAccessConfigs
	diags := req.Config.Get(ctx, &networksSmTrustedAccessConfigs)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmTrustedAccessConfigs")
		vvNetworkID := networksSmTrustedAccessConfigs.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkSmTrustedAccessConfigsQueryParams{}

		queryParams1.PerPage = int(networksSmTrustedAccessConfigs.PerPage.ValueInt64())
		queryParams1.StartingAfter = networksSmTrustedAccessConfigs.StartingAfter.ValueString()
		queryParams1.EndingBefore = networksSmTrustedAccessConfigs.EndingBefore.ValueString()

		response1, restyResp1, err := d.client.Sm.GetNetworkSmTrustedAccessConfigs(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmTrustedAccessConfigs",
				err.Error(),
			)
			return
		}

		networksSmTrustedAccessConfigs = ResponseSmGetNetworkSmTrustedAccessConfigsItemsToBody(networksSmTrustedAccessConfigs, response1)
		diags = resp.State.Set(ctx, &networksSmTrustedAccessConfigs)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSmTrustedAccessConfigs struct {
	NetworkID     types.String                                      `tfsdk:"network_id"`
	PerPage       types.Int64                                       `tfsdk:"per_page"`
	StartingAfter types.String                                      `tfsdk:"starting_after"`
	EndingBefore  types.String                                      `tfsdk:"ending_before"`
	Items         *[]ResponseItemSmGetNetworkSmTrustedAccessConfigs `tfsdk:"items"`
}

type ResponseItemSmGetNetworkSmTrustedAccessConfigs struct {
	AccessEndAt                types.String `tfsdk:"access_end_at"`
	AccessStartAt              types.String `tfsdk:"access_start_at"`
	AdditionalEmailText        types.String `tfsdk:"additional_email_text"`
	ID                         types.String `tfsdk:"id"`
	Name                       types.String `tfsdk:"name"`
	NotifyTimeBeforeAccessEnds types.Int64  `tfsdk:"notify_time_before_access_ends"`
	Scope                      types.String `tfsdk:"scope"`
	SendExpirationEmails       types.Bool   `tfsdk:"send_expiration_emails"`
	SSIDName                   types.String `tfsdk:"ssid_name"`
	Tags                       types.List   `tfsdk:"tags"`
	TimeboundType              types.String `tfsdk:"timebound_type"`
}

// ToBody
func ResponseSmGetNetworkSmTrustedAccessConfigsItemsToBody(state NetworksSmTrustedAccessConfigs, response *merakigosdk.ResponseSmGetNetworkSmTrustedAccessConfigs) NetworksSmTrustedAccessConfigs {
	var items []ResponseItemSmGetNetworkSmTrustedAccessConfigs
	for _, item := range *response {
		itemState := ResponseItemSmGetNetworkSmTrustedAccessConfigs{
			AccessEndAt:         types.StringValue(item.AccessEndAt),
			AccessStartAt:       types.StringValue(item.AccessStartAt),
			AdditionalEmailText: types.StringValue(item.AdditionalEmailText),
			ID:                  types.StringValue(item.ID),
			Name:                types.StringValue(item.Name),
			NotifyTimeBeforeAccessEnds: func() types.Int64 {
				if item.NotifyTimeBeforeAccessEnds != nil {
					return types.Int64Value(int64(*item.NotifyTimeBeforeAccessEnds))
				}
				return types.Int64{}
			}(),
			Scope: types.StringValue(item.Scope),
			SendExpirationEmails: func() types.Bool {
				if item.SendExpirationEmails != nil {
					return types.BoolValue(*item.SendExpirationEmails)
				}
				return types.Bool{}
			}(),
			SSIDName:      types.StringValue(item.SSIDName),
			Tags:          StringSliceToList(item.Tags),
			TimeboundType: types.StringValue(item.TimeboundType),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
