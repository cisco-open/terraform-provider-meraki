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
	_ datasource.DataSource              = &NetworksClientsPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksClientsPolicyDataSource{}
)

func NewNetworksClientsPolicyDataSource() datasource.DataSource {
	return &NetworksClientsPolicyDataSource{}
}

type NetworksClientsPolicyDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksClientsPolicyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksClientsPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_clients_policy"
}

func (d *NetworksClientsPolicyDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"client_id": schema.StringAttribute{
				MarkdownDescription: `clientId path parameter. Client ID`,
				Required:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"device_policy": schema.StringAttribute{
						MarkdownDescription: `The name of the client's policy`,
						Computed:            true,
					},
					"group_policy_id": schema.StringAttribute{
						MarkdownDescription: `The group policy identifier of the client`,
						Computed:            true,
					},
					"mac": schema.StringAttribute{
						MarkdownDescription: `The MAC address of the client`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *NetworksClientsPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksClientsPolicy NetworksClientsPolicy
	diags := req.Config.Get(ctx, &networksClientsPolicy)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkClientPolicy")
		vvNetworkID := networksClientsPolicy.NetworkID.ValueString()
		vvClientID := networksClientsPolicy.ClientID.ValueString()

		response1, restyResp1, err := d.client.Networks.GetNetworkClientPolicy(vvNetworkID, vvClientID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkClientPolicy",
				err.Error(),
			)
			return
		}

		networksClientsPolicy = ResponseNetworksGetNetworkClientPolicyItemToBody(networksClientsPolicy, response1)
		diags = resp.State.Set(ctx, &networksClientsPolicy)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksClientsPolicy struct {
	NetworkID types.String                            `tfsdk:"network_id"`
	ClientID  types.String                            `tfsdk:"client_id"`
	Item      *ResponseNetworksGetNetworkClientPolicy `tfsdk:"item"`
}

type ResponseNetworksGetNetworkClientPolicy struct {
	DevicePolicy  types.String `tfsdk:"device_policy"`
	GroupPolicyID types.String `tfsdk:"group_policy_id"`
	Mac           types.String `tfsdk:"mac"`
}

// ToBody
func ResponseNetworksGetNetworkClientPolicyItemToBody(state NetworksClientsPolicy, response *merakigosdk.ResponseNetworksGetNetworkClientPolicy) NetworksClientsPolicy {
	itemState := ResponseNetworksGetNetworkClientPolicy{
		DevicePolicy:  types.StringValue(response.DevicePolicy),
		GroupPolicyID: types.StringValue(response.GroupPolicyID),
		Mac:           types.StringValue(response.Mac),
	}
	state.Item = &itemState
	return state
}
