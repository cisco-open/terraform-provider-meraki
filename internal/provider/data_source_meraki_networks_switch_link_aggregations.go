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
	_ datasource.DataSource              = &NetworksSwitchLinkAggregationsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchLinkAggregationsDataSource{}
)

func NewNetworksSwitchLinkAggregationsDataSource() datasource.DataSource {
	return &NetworksSwitchLinkAggregationsDataSource{}
}

type NetworksSwitchLinkAggregationsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchLinkAggregationsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchLinkAggregationsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_link_aggregations"
}

func (d *NetworksSwitchLinkAggregationsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSwitchGetNetworkSwitchLinkAggregations`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"id": schema.StringAttribute{
							Computed: true,
						},
						"switch_ports": schema.SetNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"port_id": schema.StringAttribute{
										Computed: true,
									},
									"serial": schema.StringAttribute{
										Computed: true,
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

func (d *NetworksSwitchLinkAggregationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchLinkAggregations NetworksSwitchLinkAggregations
	diags := req.Config.Get(ctx, &networksSwitchLinkAggregations)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchLinkAggregations")
		vvNetworkID := networksSwitchLinkAggregations.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchLinkAggregations(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchLinkAggregations",
				err.Error(),
			)
			return
		}

		networksSwitchLinkAggregations = ResponseSwitchGetNetworkSwitchLinkAggregationsItemsToBody(networksSwitchLinkAggregations, response1)
		diags = resp.State.Set(ctx, &networksSwitchLinkAggregations)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchLinkAggregations struct {
	NetworkID types.String                                          `tfsdk:"network_id"`
	Items     *[]ResponseItemSwitchGetNetworkSwitchLinkAggregations `tfsdk:"items"`
}

type ResponseItemSwitchGetNetworkSwitchLinkAggregations struct {
	ID          types.String                                                     `tfsdk:"id"`
	SwitchPorts *[]ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPorts `tfsdk:"switch_ports"`
}

type ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPorts struct {
	PortID types.String `tfsdk:"port_id"`
	Serial types.String `tfsdk:"serial"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchLinkAggregationsItemsToBody(state NetworksSwitchLinkAggregations, response *merakigosdk.ResponseSwitchGetNetworkSwitchLinkAggregations) NetworksSwitchLinkAggregations {
	var items []ResponseItemSwitchGetNetworkSwitchLinkAggregations
	for _, item := range *response {
		itemState := ResponseItemSwitchGetNetworkSwitchLinkAggregations{
			ID: types.StringValue(item.ID),
			SwitchPorts: func() *[]ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPorts {
				if item.SwitchPorts != nil {
					result := make([]ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPorts, len(*item.SwitchPorts))
					for i, switchPorts := range *item.SwitchPorts {
						result[i] = ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPorts{
							PortID: types.StringValue(switchPorts.PortID),
							Serial: types.StringValue(switchPorts.Serial),
						}
					}
					return &result
				}
				return &[]ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPorts{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
