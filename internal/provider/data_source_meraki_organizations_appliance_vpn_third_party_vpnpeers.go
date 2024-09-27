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
	_ datasource.DataSource              = &OrganizationsApplianceVpnThirdPartyVpnpeersDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsApplianceVpnThirdPartyVpnpeersDataSource{}
)

func NewOrganizationsApplianceVpnThirdPartyVpnpeersDataSource() datasource.DataSource {
	return &OrganizationsApplianceVpnThirdPartyVpnpeersDataSource{}
}

type OrganizationsApplianceVpnThirdPartyVpnpeersDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsApplianceVpnThirdPartyVpnpeersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsApplianceVpnThirdPartyVpnpeersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_appliance_vpn_third_party_vpnpeers"
}

func (d *OrganizationsApplianceVpnThirdPartyVpnpeersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"peers": schema.SetNestedAttribute{
						MarkdownDescription: `The list of VPN peers`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"ike_version": schema.StringAttribute{
									MarkdownDescription: `[optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to '1' when omitted.`,
									Computed:            true,
								},
								"ipsec_policies": schema.SingleNestedAttribute{
									MarkdownDescription: `Custom IPSec policies for the VPN peer. If not included and a preset has not been chosen, the default preset for IPSec policies will be used.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"child_auth_algo": schema.ListAttribute{
											MarkdownDescription: `This is the authentication algorithms to be used in Phase 2. The value should be an array with one of the following algorithms: 'sha256', 'sha1', 'md5'`,
											Computed:            true,
											ElementType:         types.StringType,
										},
										"child_cipher_algo": schema.ListAttribute{
											MarkdownDescription: `This is the cipher algorithms to be used in Phase 2. The value should be an array with one or more of the following algorithms: 'aes256', 'aes192', 'aes128', 'tripledes', 'des', 'null'`,
											Computed:            true,
											ElementType:         types.StringType,
										},
										"child_lifetime": schema.Int64Attribute{
											MarkdownDescription: `The lifetime of the Phase 2 SA in seconds.`,
											Computed:            true,
										},
										"child_pfs_group": schema.ListAttribute{
											MarkdownDescription: `This is the Diffie-Hellman group to be used for Perfect Forward Secrecy in Phase 2. The value should be an array with one of the following values: 'disabled','group14', 'group5', 'group2', 'group1'`,
											Computed:            true,
											ElementType:         types.StringType,
										},
										"ike_auth_algo": schema.ListAttribute{
											MarkdownDescription: `This is the authentication algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: 'sha256', 'sha1', 'md5'`,
											Computed:            true,
											ElementType:         types.StringType,
										},
										"ike_cipher_algo": schema.ListAttribute{
											MarkdownDescription: `This is the cipher algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: 'aes256', 'aes192', 'aes128', 'tripledes', 'des'`,
											Computed:            true,
											ElementType:         types.StringType,
										},
										"ike_diffie_hellman_group": schema.ListAttribute{
											MarkdownDescription: `This is the Diffie-Hellman group to be used in Phase 1. The value should be an array with one of the following algorithms: 'group14', 'group5', 'group2', 'group1'`,
											Computed:            true,
											ElementType:         types.StringType,
										},
										"ike_lifetime": schema.Int64Attribute{
											MarkdownDescription: `The lifetime of the Phase 1 SA in seconds.`,
											Computed:            true,
										},
										"ike_prf_algo": schema.ListAttribute{
											MarkdownDescription: `[optional] This is the pseudo-random function to be used in IKE_SA. The value should be an array with one of the following algorithms: 'prfsha256', 'prfsha1', 'prfmd5', 'default'. The 'default' option can be used to default to the Authentication algorithm.`,
											Computed:            true,
											ElementType:         types.StringType,
										},
									},
								},
								"ipsec_policies_preset": schema.StringAttribute{
									MarkdownDescription: `One of the following available presets: 'default', 'aws', 'azure'. If this is provided, the 'ipsecPolicies' parameter is ignored.`,
									Computed:            true,
								},
								"local_id": schema.StringAttribute{
									MarkdownDescription: `[optional] The local ID is used to identify the MX to the peer. This will apply to all MXs this peer applies to.`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `The name of the VPN peer`,
									Computed:            true,
								},
								"network_tags": schema.ListAttribute{
									MarkdownDescription: `A list of network tags that will connect with this peer. Use ['all'] for all networks. Use ['none'] for no networks. If not included, the default is ['all'].`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"private_subnets": schema.ListAttribute{
									MarkdownDescription: `The list of the private subnets of the VPN peer`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"public_ip": schema.StringAttribute{
									MarkdownDescription: `[optional] The public IP of the VPN peer`,
									Computed:            true,
								},
								"remote_id": schema.StringAttribute{
									MarkdownDescription: `[optional] The remote ID is used to identify the connecting VPN peer. This can either be a valid IPv4 Address, FQDN or User FQDN.`,
									Computed:            true,
								},
								"secret": schema.StringAttribute{
									MarkdownDescription: `The shared secret with the VPN peer`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsApplianceVpnThirdPartyVpnpeersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsApplianceVpnThirdPartyVpnpeers OrganizationsApplianceVpnThirdPartyVpnpeers
	diags := req.Config.Get(ctx, &organizationsApplianceVpnThirdPartyVpnpeers)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationApplianceVpnThirdPartyVpnpeers")
		vvOrganizationID := organizationsApplianceVpnThirdPartyVpnpeers.OrganizationID.ValueString()

		response1, restyResp1, err := d.client.Appliance.GetOrganizationApplianceVpnThirdPartyVpnpeers(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationApplianceVpnThirdPartyVpnpeers",
				err.Error(),
			)
			return
		}

		organizationsApplianceVpnThirdPartyVpnpeers = ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersItemToBody(organizationsApplianceVpnThirdPartyVpnpeers, response1)
		diags = resp.State.Set(ctx, &organizationsApplianceVpnThirdPartyVpnpeers)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsApplianceVpnThirdPartyVpnpeers struct {
	OrganizationID types.String                                                    `tfsdk:"organization_id"`
	Item           *ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeers `tfsdk:"item"`
}

type ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeers struct {
	Peers *[]ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeers `tfsdk:"peers"`
}

type ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeers struct {
	IkeVersion          types.String                                                                      `tfsdk:"ike_version"`
	IPsecPolicies       *ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeersIpsecPolicies `tfsdk:"ipsec_policies"`
	IPsecPoliciesPreset types.String                                                                      `tfsdk:"ipsec_policies_preset"`
	LocalID             types.String                                                                      `tfsdk:"local_id"`
	Name                types.String                                                                      `tfsdk:"name"`
	NetworkTags         types.List                                                                        `tfsdk:"network_tags"`
	PrivateSubnets      types.List                                                                        `tfsdk:"private_subnets"`
	PublicIP            types.String                                                                      `tfsdk:"public_ip"`
	RemoteID            types.String                                                                      `tfsdk:"remote_id"`
	Secret              types.String                                                                      `tfsdk:"secret"`
}

type ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeersIpsecPolicies struct {
	ChildAuthAlgo         types.List  `tfsdk:"child_auth_algo"`
	ChildCipherAlgo       types.List  `tfsdk:"child_cipher_algo"`
	ChildLifetime         types.Int64 `tfsdk:"child_lifetime"`
	ChildPfsGroup         types.List  `tfsdk:"child_pfs_group"`
	IkeAuthAlgo           types.List  `tfsdk:"ike_auth_algo"`
	IkeCipherAlgo         types.List  `tfsdk:"ike_cipher_algo"`
	IkeDiffieHellmanGroup types.List  `tfsdk:"ike_diffie_hellman_group"`
	IkeLifetime           types.Int64 `tfsdk:"ike_lifetime"`
	IkePrfAlgo            types.List  `tfsdk:"ike_prf_algo"`
}

// ToBody
func ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersItemToBody(state OrganizationsApplianceVpnThirdPartyVpnpeers, response *merakigosdk.ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeers) OrganizationsApplianceVpnThirdPartyVpnpeers {
	itemState := ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeers{
		Peers: func() *[]ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeers {
			if response.Peers != nil {
				result := make([]ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeers, len(*response.Peers))
				for i, peers := range *response.Peers {
					result[i] = ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeers{
						IkeVersion: types.StringValue(peers.IkeVersion),
						IPsecPolicies: func() *ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeersIpsecPolicies {
							if peers.IPsecPolicies != nil {
								return &ResponseApplianceGetOrganizationApplianceVpnThirdPartyVpnpeersPeersIpsecPolicies{
									ChildAuthAlgo:   StringSliceToList(peers.IPsecPolicies.ChildAuthAlgo),
									ChildCipherAlgo: StringSliceToList(peers.IPsecPolicies.ChildCipherAlgo),
									ChildLifetime: func() types.Int64 {
										if peers.IPsecPolicies.ChildLifetime != nil {
											return types.Int64Value(int64(*peers.IPsecPolicies.ChildLifetime))
										}
										return types.Int64{}
									}(),
									ChildPfsGroup:         StringSliceToList(peers.IPsecPolicies.ChildPfsGroup),
									IkeAuthAlgo:           StringSliceToList(peers.IPsecPolicies.IkeAuthAlgo),
									IkeCipherAlgo:         StringSliceToList(peers.IPsecPolicies.IkeCipherAlgo),
									IkeDiffieHellmanGroup: StringSliceToList(peers.IPsecPolicies.IkeDiffieHellmanGroup),
									IkeLifetime: func() types.Int64 {
										if peers.IPsecPolicies.IkeLifetime != nil {
											return types.Int64Value(int64(*peers.IPsecPolicies.IkeLifetime))
										}
										return types.Int64{}
									}(),
									IkePrfAlgo: StringSliceToList(peers.IPsecPolicies.IkePrfAlgo),
								}
							}
							return nil
						}(),
						IPsecPoliciesPreset: types.StringValue(peers.IPsecPoliciesPreset),
						LocalID:             types.StringValue(peers.LocalID),
						Name:                types.StringValue(peers.Name),
						NetworkTags:         StringSliceToList(peers.NetworkTags),
						PrivateSubnets:      StringSliceToList(peers.PrivateSubnets),
						PublicIP:            types.StringValue(peers.PublicIP),
						RemoteID:            types.StringValue(peers.RemoteID),
						Secret:              types.StringValue(peers.Secret),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
