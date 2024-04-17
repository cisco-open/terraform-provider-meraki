package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWirelessSSIDsHotspot20Resource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessSSIDsHotspot20Resource{}
)

func NewNetworksWirelessSSIDsHotspot20Resource() resource.Resource {
	return &NetworksWirelessSSIDsHotspot20Resource{}
}

type NetworksWirelessSSIDsHotspot20Resource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessSSIDsHotspot20Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessSSIDsHotspot20Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_hotspot20"
}

func (r *NetworksWirelessSSIDsHotspot20Resource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"domains": schema.SetAttribute{
				MarkdownDescription: `An array of domain names`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `Whether or not Hotspot 2.0 for this SSID is enabled`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"mcc_mncs": schema.SetNestedAttribute{
				MarkdownDescription: `An array of MCC/MNC pairs`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"mcc": schema.StringAttribute{
							MarkdownDescription: `MCC value`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"mnc": schema.StringAttribute{
							MarkdownDescription: `MNC value`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"nai_realms": schema.SetNestedAttribute{
				MarkdownDescription: `An array of NAI realms`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"format": schema.StringAttribute{
							MarkdownDescription: `The format for the realm ('1' or '0')`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"0",
									"1",
								),
							},
						},
						"methods": schema.SetNestedAttribute{
							MarkdownDescription: `An array of EAP methods for the realm.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"authentication_types": schema.SingleNestedAttribute{
										MarkdownDescription: `The authentication types for the method. These should be formatted as an object with the EAP method category in camelcase as the key and the list of types as the value (nonEapInnerAuthentication: Reserved, PAP, CHAP, MSCHAP, MSCHAPV2; eapInnerAuthentication: EAP-TLS, EAP-SIM, EAP-AKA, EAP-TTLS with MSCHAPv2; credentials: SIM, USIM, NFC Secure Element, Hardware Token, Softoken, Certificate, username/password, none, Reserved, Vendor Specific; tunneledEapMethodCredentials: SIM, USIM, NFC Secure Element, Hardware Token, Softoken, Certificate, username/password, Reserved, Anonymous, Vendor Specific)`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"credentials": schema.SetAttribute{
												Computed:    true,
												ElementType: types.StringType,
											},
											"eapinner_authentication": schema.SetAttribute{
												Computed:    true,
												ElementType: types.StringType,
											},
											"non_eapinner_authentication": schema.SetAttribute{
												Computed:    true,
												ElementType: types.StringType,
											},
											"tunneled_eap_method_credentials": schema.SetAttribute{
												Computed:    true,
												ElementType: types.StringType,
											},
										},
									},
									"id": schema.StringAttribute{
										MarkdownDescription: `ID of method`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
								},
							},
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"realm": schema.StringAttribute{
							MarkdownDescription: `The name of the realm`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"network_access_type": schema.StringAttribute{
				MarkdownDescription: `The network type of this SSID ('Private network', 'Private network with guest access', 'Chargeable public network', 'Free public network', 'Personal device network', 'Emergency services only network', 'Test or experimental', 'Wildcard')`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Chargeable public network",
						"Emergency services only network",
						"Free public network",
						"Personal device network",
						"Private network",
						"Private network with guest access",
						"Test or experimental",
						"Wildcard",
					),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: `number path parameter.`,
				Required:            true,
			},
			"operator": schema.SingleNestedAttribute{
				MarkdownDescription: `Operator settings for this SSID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"name": schema.StringAttribute{
						MarkdownDescription: `Operator name`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"roam_consort_ois": schema.SetAttribute{
				MarkdownDescription: `An array of roaming consortium OIs (hexadecimal number 3-5 octets in length)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"venue": schema.SingleNestedAttribute{
				MarkdownDescription: `Venue settings for this SSID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"name": schema.StringAttribute{
						MarkdownDescription: `Venue name`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"type": schema.StringAttribute{
						MarkdownDescription: `Venue type ('Unspecified', 'Unspecified Assembly', 'Arena', 'Stadium', 'Passenger Terminal', 'Amphitheater', 'Amusement Park', 'Place of Worship', 'Convention Center', 'Library', 'Museum', 'Restaurant', 'Theater', 'Bar', 'Coffee Shop', 'Zoo or Aquarium', 'Emergency Coordination Center', 'Unspecified Business', 'Doctor or Dentist office', 'Bank', 'Fire Station', 'Police Station', 'Post Office', 'Professional Office', 'Research and Development Facility', 'Attorney Office', 'Unspecified Educational', 'School, Primary', 'School, Secondary', 'University or College', 'Unspecified Factory and Industrial', 'Factory', 'Unspecified Institutional', 'Hospital', 'Long-Term Care Facility', 'Alcohol and Drug Rehabilitation Center', 'Group Home', 'Prison or Jail', 'Unspecified Mercantile', 'Retail Store', 'Grocery Market', 'Automotive Service Station', 'Shopping Mall', 'Gas Station', 'Unspecified Residential', 'Private Residence', 'Hotel or Motel', 'Dormitory', 'Boarding House', 'Unspecified Storage', 'Unspecified Utility and Miscellaneous', 'Unspecified Vehicular', 'Automobile or Truck', 'Airplane', 'Bus', 'Ferry', 'Ship or Boat', 'Train', 'Motor Bike', 'Unspecified Outdoor', 'Muni-mesh Network', 'City Park', 'Rest Area', 'Traffic Control', 'Bus Stop', 'Kiosk')`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"Airplane",
								"Alcohol and Drug Rehabilitation Center",
								"Amphitheater",
								"Amusement Park",
								"Arena",
								"Attorney Office",
								"Automobile or Truck",
								"Automotive Service Station",
								"Bank",
								"Bar",
								"Boarding House",
								"Bus",
								"Bus Stop",
								"City Park",
								"Coffee Shop",
								"Convention Center",
								"Doctor or Dentist office",
								"Dormitory",
								"Emergency Coordination Center",
								"Factory",
								"Ferry",
								"Fire Station",
								"Gas Station",
								"Grocery Market",
								"Group Home",
								"Hospital",
								"Hotel or Motel",
								"Kiosk",
								"Library",
								"Long-Term Care Facility",
								"Motor Bike",
								"Muni-mesh Network",
								"Museum",
								"Passenger Terminal",
								"Place of Worship",
								"Police Station",
								"Post Office",
								"Prison or Jail",
								"Private Residence",
								"Professional Office",
								"Research and Development Facility",
								"Rest Area",
								"Restaurant",
								"Retail Store",
								"School, Primary",
								"School, Secondary",
								"Ship or Boat",
								"Shopping Mall",
								"Stadium",
								"Theater",
								"Traffic Control",
								"Train",
								"University or College",
								"Unspecified",
								"Unspecified Assembly",
								"Unspecified Business",
								"Unspecified Educational",
								"Unspecified Factory and Industrial",
								"Unspecified Institutional",
								"Unspecified Mercantile",
								"Unspecified Outdoor",
								"Unspecified Residential",
								"Unspecified Storage",
								"Unspecified Utility and Miscellaneous",
								"Unspecified Vehicular",
								"Zoo or Aquarium",
							),
						},
					},
				},
			},
		},
	}
}
func (r *NetworksWirelessSSIDsHotspot20Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessSSIDsHotspot20Rs

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
	vvNumber := data.Number.ValueString()
	//Item
	responseVerifyItem, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDHotspot20(vvNetworkID, vvNumber)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksWirelessSSIDsHotspot20 only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksWirelessSSIDsHotspot20 only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDHotspot20(vvNetworkID, vvNumber, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDHotspot20",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDHotspot20",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDHotspot20(vvNetworkID, vvNumber)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDHotspot20",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDHotspot20",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetNetworkWirelessSSIDHotspot20ItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessSSIDsHotspot20Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessSSIDsHotspot20Rs

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
	vvNumber := data.Number.ValueString()
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessSSIDHotspot20(vvNetworkID, vvNumber)
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
				"Failure when executing GetNetworkWirelessSSIDHotspot20",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDHotspot20",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetNetworkWirelessSSIDHotspot20ItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksWirelessSSIDsHotspot20Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("number"), idParts[1])...)
}

func (r *NetworksWirelessSSIDsHotspot20Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksWirelessSSIDsHotspot20Rs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvNumber := data.Number.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDHotspot20(vvNetworkID, vvNumber, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDHotspot20",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDHotspot20",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessSSIDsHotspot20Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksWirelessSSIDsHotspot20", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessSSIDsHotspot20Rs struct {
	NetworkID         types.String                                                  `tfsdk:"network_id"`
	Number            types.String                                                  `tfsdk:"number"`
	Domains           types.Set                                                     `tfsdk:"domains"`
	Enabled           types.Bool                                                    `tfsdk:"enabled"`
	MccMncs           *[]ResponseWirelessGetNetworkWirelessSsidHotspot20MccMncsRs   `tfsdk:"mcc_mncs"`
	NaiRealms         *[]ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsRs `tfsdk:"nai_realms"`
	NetworkAccessType types.String                                                  `tfsdk:"network_access_type"`
	Operator          *ResponseWirelessGetNetworkWirelessSsidHotspot20OperatorRs    `tfsdk:"operator"`
	RoamConsortOis    types.Set                                                     `tfsdk:"roam_consort_ois"`
	Venue             *ResponseWirelessGetNetworkWirelessSsidHotspot20VenueRs       `tfsdk:"venue"`
}

type ResponseWirelessGetNetworkWirelessSsidHotspot20MccMncsRs struct {
	Mcc types.String `tfsdk:"mcc"`
	Mnc types.String `tfsdk:"mnc"`
}

type ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsRs struct {
	Format  types.String                                                         `tfsdk:"format"`
	Methods *[]ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethodsRs `tfsdk:"methods"`
	Name    types.String                                                         `tfsdk:"name"`
	Realm   types.String                                                         `tfsdk:"realm"`
}

type ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethodsRs struct {
	AuthenticationTypes *ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethodsAuthenticationTypesRs `tfsdk:"authentication_types"`
	ID                  types.String                                                                          `tfsdk:"id"`
}

type ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethodsAuthenticationTypesRs struct {
	Credentials                  types.Set `tfsdk:"credentials"`
	EapinnerAuthentication       types.Set `tfsdk:"eap_inner_authentication"`
	NonEapinnerAuthentication    types.Set `tfsdk:"non_eap_inner_authentication"`
	TunneledEapMethodCredentials types.Set `tfsdk:"tunneled_eap_method_credentials"`
}

type ResponseWirelessGetNetworkWirelessSsidHotspot20OperatorRs struct {
	Name types.String `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessSsidHotspot20VenueRs struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

// FromBody
func (r *NetworksWirelessSSIDsHotspot20Rs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDHotspot20 {
	emptyString := ""
	var domains []string = nil
	r.Domains.ElementsAs(ctx, &domains, false)
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var requestWirelessUpdateNetworkWirelessSSIDHotspot20MccMncs []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDHotspot20MccMncs
	if r.MccMncs != nil {
		for _, rItem1 := range *r.MccMncs {
			mcc := rItem1.Mcc.ValueString()
			mnc := rItem1.Mnc.ValueString()
			requestWirelessUpdateNetworkWirelessSSIDHotspot20MccMncs = append(requestWirelessUpdateNetworkWirelessSSIDHotspot20MccMncs, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDHotspot20MccMncs{
				Mcc: mcc,
				Mnc: mnc,
			})
		}
	}
	var requestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealms []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealms
	if r.NaiRealms != nil {
		for _, rItem1 := range *r.NaiRealms {
			format := rItem1.Format.ValueString()
			var requestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealmsMethods []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealmsMethods
			if rItem1.Methods != nil {
				for _, rItem2 := range *rItem1.Methods { //Methods// name: methods
					// var requestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealmsMethodsAuthenticationTypes *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealmsMethodsAuthenticationTypes
					// if rItem2.AuthenticationTypes != nil {
					// 	requestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealmsMethodsAuthenticationTypes = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealmsMethodsAuthenticationTypes{}
					// }
					iD := rItem2.ID.ValueString()
					requestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealmsMethods = append(requestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealmsMethods, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealmsMethods{
						// AuthenticationTypes: requestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealmsMethodsAuthenticationTypes,
						ID: iD,
					})
				}
			}
			realm := rItem1.Realm.ValueString()
			requestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealms = append(requestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealms, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealms{
				Format: format,
				Methods: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealmsMethods {
					if len(requestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealmsMethods) > 0 {
						return &requestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealmsMethods
					}
					return nil
				}(),
				Realm: realm,
			})
		}
	}
	networkAccessType := new(string)
	if !r.NetworkAccessType.IsUnknown() && !r.NetworkAccessType.IsNull() {
		*networkAccessType = r.NetworkAccessType.ValueString()
	} else {
		networkAccessType = &emptyString
	}
	var requestWirelessUpdateNetworkWirelessSSIDHotspot20Operator *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDHotspot20Operator
	if r.Operator != nil {
		name := r.Operator.Name.ValueString()
		requestWirelessUpdateNetworkWirelessSSIDHotspot20Operator = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDHotspot20Operator{
			Name: name,
		}
	}
	var roamConsortOis []string = nil
	r.RoamConsortOis.ElementsAs(ctx, &roamConsortOis, false)
	var requestWirelessUpdateNetworkWirelessSSIDHotspot20Venue *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDHotspot20Venue
	if r.Venue != nil {
		name := r.Venue.Name.ValueString()
		typeR := r.Venue.Type.ValueString()
		requestWirelessUpdateNetworkWirelessSSIDHotspot20Venue = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDHotspot20Venue{
			Name: name,
			Type: typeR,
		}
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDHotspot20{
		Domains: domains,
		Enabled: enabled,
		MccMncs: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDHotspot20MccMncs {
			if len(requestWirelessUpdateNetworkWirelessSSIDHotspot20MccMncs) > 0 {
				return &requestWirelessUpdateNetworkWirelessSSIDHotspot20MccMncs
			}
			return nil
		}(),
		NaiRealms: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealms {
			if len(requestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealms) > 0 {
				return &requestWirelessUpdateNetworkWirelessSSIDHotspot20NaiRealms
			}
			return nil
		}(),
		NetworkAccessType: *networkAccessType,
		Operator:          requestWirelessUpdateNetworkWirelessSSIDHotspot20Operator,
		RoamConsortOis:    roamConsortOis,
		Venue:             requestWirelessUpdateNetworkWirelessSSIDHotspot20Venue,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessSSIDHotspot20ItemToBodyRs(state NetworksWirelessSSIDsHotspot20Rs, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDHotspot20, is_read bool) NetworksWirelessSSIDsHotspot20Rs {
	itemState := NetworksWirelessSSIDsHotspot20Rs{
		Domains: StringSliceToSet(response.Domains),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		MccMncs: func() *[]ResponseWirelessGetNetworkWirelessSsidHotspot20MccMncsRs {
			if response.MccMncs != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidHotspot20MccMncsRs, len(*response.MccMncs))
				for i, mccMncs := range *response.MccMncs {
					result[i] = ResponseWirelessGetNetworkWirelessSsidHotspot20MccMncsRs{
						Mcc: types.StringValue(mccMncs.Mcc),
						Mnc: types.StringValue(mccMncs.Mnc),
					}
				}
				return &result
			}
			return &[]ResponseWirelessGetNetworkWirelessSsidHotspot20MccMncsRs{}
		}(),
		NaiRealms: func() *[]ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsRs {
			if response.NaiRealms != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsRs, len(*response.NaiRealms))
				for i, naiRealms := range *response.NaiRealms {
					result[i] = ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsRs{
						Format: types.StringValue(naiRealms.Format),
						Methods: func() *[]ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethodsRs {
							if naiRealms.Methods != nil {
								result := make([]ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethodsRs, len(*naiRealms.Methods))
								for i, methods := range *naiRealms.Methods {
									result[i] = ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethodsRs{
										AuthenticationTypes: func() *ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethodsAuthenticationTypesRs {
											if methods.AuthenticationTypes != nil {
												return &ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethodsAuthenticationTypesRs{
													Credentials:                  StringSliceToSet(methods.AuthenticationTypes.Credentials),
													EapinnerAuthentication:       StringSliceToSet(methods.AuthenticationTypes.EapinnerAuthentication),
													NonEapinnerAuthentication:    StringSliceToSet(methods.AuthenticationTypes.NonEapinnerAuthentication),
													TunneledEapMethodCredentials: StringSliceToSet(methods.AuthenticationTypes.TunneledEapMethodCredentials),
												}
											}
											return &ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethodsAuthenticationTypesRs{}
										}(),
										ID: types.StringValue(methods.ID),
									}
								}
								return &result
							}
							return &[]ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsMethodsRs{}
						}(),
						Name: types.StringValue(naiRealms.Name),
					}
				}
				return &result
			}
			return &[]ResponseWirelessGetNetworkWirelessSsidHotspot20NaiRealmsRs{}
		}(),
		NetworkAccessType: types.StringValue(response.NetworkAccessType),
		Operator: func() *ResponseWirelessGetNetworkWirelessSsidHotspot20OperatorRs {
			if response.Operator != nil {
				return &ResponseWirelessGetNetworkWirelessSsidHotspot20OperatorRs{
					Name: types.StringValue(response.Operator.Name),
				}
			}
			return &ResponseWirelessGetNetworkWirelessSsidHotspot20OperatorRs{}
		}(),
		RoamConsortOis: StringSliceToSet(response.RoamConsortOis),
		Venue: func() *ResponseWirelessGetNetworkWirelessSsidHotspot20VenueRs {
			if response.Venue != nil {
				return &ResponseWirelessGetNetworkWirelessSsidHotspot20VenueRs{
					Name: types.StringValue(response.Venue.Name),
					Type: types.StringValue(response.Venue.Type),
				}
			}
			return &ResponseWirelessGetNetworkWirelessSsidHotspot20VenueRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessSSIDsHotspot20Rs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessSSIDsHotspot20Rs)
}
