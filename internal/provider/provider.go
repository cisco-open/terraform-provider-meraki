// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	CUSTOM_USER_AGENT  = "MerakiTerraform/1.0.3-beta Cisco"
	DEFAULT_USER_AGENT = "Meraki"
)

// terraform-provider-meraki
// Ensure MerakiProvider satisfies various provider interfaces.
var _ provider.Provider = &MerakiProvider{}

// MerakiProvider defines the provider implementation.
type MerakiProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// MerakiProviderModel describes the provider data model.
type MerakiProviderModel struct {
	BaseURL               types.String `tfsdk:"meraki_base_url"`
	MerakiDashboardApiKey types.String `tfsdk:"meraki_dashboard_api_key"`
	Debug                 types.String `tfsdk:"meraki_debug"`
	RequestPerSecond      types.Int64  `tfsdk:"meraki_requests_per_second"`
	UserAgent             types.String `tfsdk:"meraki_user_agent"`
}

type MerakiProviderData struct {
	Client *merakigosdk.Client
}

func (p *MerakiProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "meraki"
	resp.Version = p.version
}

func (p *MerakiProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"meraki_base_url": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Cisco Meraki base URL, FQDN or IP. If not set, it uses the MERAKI_BASE_URL environment variable. Default is (https://api.meraki.com/)",
			},
			"meraki_dashboard_api_key": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				MarkdownDescription: "Cisco Meraki meraki_dashboard_api_key to authenticate. If not set, it uses the MERAKI_DASHBOARD_API_KEY environment variable.",
			},
			"meraki_debug": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Flag for Cisco Meraki to enable debugging. If not set, it uses the MERAKI_DEBUG environment variable defaults to `false`.",
			},
			"meraki_requests_per_second": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Flag requests per second allowed for client. Default is (10)",
			},
			"meraki_user_agent": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: " Define an identifier or User-Agent for API requests to Meraki. Default is (Meraki)",
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^\S*$`),
						"must not contain white spaces",
					),
				},
			},
		},
	}
}

func (p *MerakiProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data MerakiProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.BaseURL.IsUnknown() || data.BaseURL.IsNull() {
		// resp.Diagnostics.AddAttributeError(
		// 	path.Root("base_url"),
		// 	"Unknown Meraki API base_url",
		// 	"The provider cannot create the Meraki API client as there is an unknown configuration value for the Meraki API BaseURL. "+
		// 		"Either target apply the source of the value first, set the value statically in the configuration, or use the MERAKI_BASE_URL environment variable.",
		// )
		data.BaseURL = types.StringValue("https://api.meraki.com/")
		// return
	}

	var requestPerSecond int
	if data.RequestPerSecond.IsUnknown() || data.RequestPerSecond.IsNull() {
		// resp.Diagnostics.AddAttributeError(
		// 	path.Root("base_url"),
		// 	"Unknown Meraki API base_url",
		// 	"The provider cannot create the Meraki API client as there is an unknown configuration value for the Meraki API BaseURL. "+
		// 		"Either target apply the source of the value first, set the value statically in the configuration, or use the MERAKI_BASE_URL environment variable.",
		// )
		requestPerSecond = 10
		// return
	} else {
		requestPerSecondTf := int(data.RequestPerSecond.ValueInt64())
		requestPerSecond = requestPerSecondTf
	}

	if data.MerakiDashboardApiKey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("dashboard_api_key"),
			"Unknown Meraki API dashboard_api_key",
			"The provider cannot create the Meraki API client as there is an unknown configuration value for the Meraki API MerakiDashboardApiKey. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the MERAKI_DASHBOARD_API_KEY environment variable.",
		)
		return
	}
	// if data.Debug.IsUnknown() || data.Debug.IsNull() {
	// 	// resp.Diagnostics.AddAttributeError(
	// 	// 	path.Root("debug"),
	// 	// 	"Unknown Meraki API debug",
	// 	// 	"The provider cannot create the Meraki API client as there is an unknown configuration value for the Meraki API Debug. "+
	// 	// 		"Either target apply the source of the value first, set the value statically in the configuration, or use the MERAKI_DEBUG environment variable.",
	// 	// )
	// 	// return
	// 	data.Debug = types.StringValue("false")
	// }
	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to enviroment variables, but override
	// with Terraform configuration value if set.
	baseURL := os.Getenv("MERAKI_BASE_URL")
	merakiDashboardApiKey := os.Getenv("MERAKI_DASHBOARD_API_KEY")
	debug := os.Getenv("MERAKI_DEBUG")
	userAgent := CUSTOM_USER_AGENT
	if !data.BaseURL.IsNull() && !data.BaseURL.IsUnknown() {
		baseURL = data.BaseURL.ValueString()
	}
	if !data.MerakiDashboardApiKey.IsNull() && !data.MerakiDashboardApiKey.IsUnknown() {
		merakiDashboardApiKey = data.MerakiDashboardApiKey.ValueString()
	}
	if !data.Debug.IsNull() && !data.Debug.IsUnknown() {
		debug = data.Debug.ValueString()
	}

	if !data.UserAgent.IsNull() && !data.UserAgent.IsUnknown() {
		userAgent = fmt.Sprintf("%s %s", CUSTOM_USER_AGENT, data.UserAgent.ValueString())
	} else {
		userAgent = fmt.Sprintf("%s %s", CUSTOM_USER_AGENT, DEFAULT_USER_AGENT)
	}
	log.Printf("[DEBUG] Meraki userAgent: %s", userAgent)
	// if !data.SSLVerify.IsNull() {
	// 	sslVerify = data.SSLVerify.ValueString()
	// }

	// Create a new Meraki client using the configuration values
	client, err := merakigosdk.NewClientWithOptionsAndRequests(baseURL,
		merakiDashboardApiKey, debug, userAgent, requestPerSecond,
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Uneable to Create Meraki API Client",
			"Error: "+err.Error(),
		)
		return
	}
	// client.RestyClient().SetLogger(createLogger())
	// client.SetUserAgent(customUserAgent)
	dataClient := MerakiProviderData{Client: client}

	resp.DataSourceData = dataClient
	resp.ResourceData = dataClient

}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &MerakiProvider{
			version: version,
		}
	}
}
func (p *MerakiProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewOrganizationsResource,
		NewOrganizationsAdminsResource,
		NewDevicesResource,
		NewDevicesApplianceUplinksSettingsResource,
		NewDevicesCameraCustomAnalyticsResource,
		NewDevicesCameraQualityAndRetentionResource,
		NewDevicesCameraSenseResource,
		NewDevicesCameraVideoSettingsResource,
		NewDevicesCameraWirelessProfilesResource,
		NewDevicesCellularSimsResource,
		NewDevicesCellularGatewayLanResource,
		NewDevicesCellularGatewayPortForwardingRulesResource,
		NewDevicesManagementInterfaceResource,
		NewDevicesSensorRelationshipsResource,
		NewDevicesSwitchPortsResource,
		NewDevicesSwitchRoutingInterfacesResource,
		NewDevicesSwitchRoutingInterfacesDhcpResource,
		NewDevicesSwitchRoutingStaticRoutesResource,
		NewDevicesSwitchWarmSpareResource,
		NewDevicesWirelessBluetoothSettingsResource,
		NewDevicesWirelessRadioSettingsResource,
		NewNetworksResource,
		NewNetworksAlertsSettingsResource,
		NewNetworksApplianceConnectivityMonitoringDestinationsResource,
		NewNetworksApplianceContentFilteringResource,
		NewNetworksApplianceStaticRoutesResource,
		NewNetworksApplianceFirewallCellularFirewallRulesResource,
		NewNetworksApplianceFirewallFirewalledServicesResource,
		NewNetworksApplianceFirewallInboundFirewallRulesResource,
		NewNetworksApplianceFirewallL3FirewallRulesResource,
		NewNetworksApplianceFirewallL7FirewallRulesResource,
		NewNetworksApplianceFirewallOneToManyNatRulesResource,
		NewNetworksApplianceFirewallOneToOneNatRulesResource,
		NewNetworksApplianceFirewallPortForwardingRulesResource,
		NewNetworksApplianceFirewallSettingsResource,
		NewNetworksAppliancePortsResource,
		NewNetworksAppliancePrefixesDelegatedStaticsResource,
		NewNetworksApplianceSecurityIntrusionResource,
		NewNetworksApplianceSecurityMalwareResource,
		NewNetworksApplianceSettingsResource,
		NewNetworksApplianceSingleLanResource,
		NewNetworksApplianceSSIDsResource,
		NewNetworksApplianceTrafficShapingResource,
		NewNetworksApplianceTrafficShapingRulesResource,
		NewNetworksApplianceTrafficShapingUplinkBandwidthResource,
		NewNetworksApplianceTrafficShapingUplinkSelectionResource,
		NewNetworksApplianceVLANsSettingsResource,
		NewNetworksApplianceVLANsResource,
		NewNetworksApplianceVpnBgpResource,
		NewNetworksApplianceVpnSiteToSiteVpnResource,
		NewNetworksApplianceWarmSpareResource,
		NewNetworksCameraQualityRetentionProfilesResource,
		NewNetworksCameraWirelessProfilesResource,
		NewNetworksCellularGatewayConnectivityMonitoringDestinationsResource,
		NewNetworksCellularGatewayDhcpResource,
		NewNetworksCellularGatewaySubnetPoolResource,
		NewNetworksCellularGatewayUplinkResource,
		NewNetworksClientsPolicyResource,
		NewNetworksClientsSplashAuthorizationStatusResource,
		NewNetworksFirmwareUpgradesResource,
		NewNetworksFirmwareUpgradesStagedEventsResource,
		NewNetworksFirmwareUpgradesStagedGroupsResource,
		NewNetworksFirmwareUpgradesStagedStagesResource,
		NewNetworksFloorPlansResource,
		NewNetworksGroupPoliciesResource,
		NewNetworksMerakiAuthUsersResource,
		NewNetworksNetflowResource,
		NewNetworksSensorAlertsProfilesResource,
		NewNetworksSensorMqttBrokersResource,
		NewNetworksSettingsResource,
		NewNetworksSmBypassActivationLockAttemptsResource,
		NewNetworksSmTargetGroupsResource,
		NewNetworksSNMPResource,
		NewNetworksSwitchAccessControlListsResource,
		NewNetworksSwitchAccessPoliciesResource,
		NewNetworksSwitchAlternateManagementInterfaceResource,
		NewNetworksSwitchDhcpServerPolicyResource,
		NewNetworksSwitchDhcpServerPolicyArpInspectionTrustedServersResource,
		NewNetworksSwitchDscpToCosMappingsResource,
		NewNetworksSwitchLinkAggregationsResource,
		NewNetworksSwitchMtuResource,
		NewNetworksSwitchPortSchedulesResource,
		NewNetworksSwitchQosRulesOrderResource,
		NewNetworksSwitchRoutingMulticastResource,
		NewNetworksSwitchRoutingMulticastRendezvousPointsResource,
		NewNetworksSwitchRoutingOspfResource,
		NewNetworksSwitchSettingsResource,
		NewNetworksSwitchStacksResource,
		NewNetworksSwitchStacksRoutingInterfacesResource,
		NewNetworksSwitchStacksRoutingInterfacesDhcpResource,
		NewNetworksSwitchStacksRoutingStaticRoutesResource,
		NewNetworksSwitchStormControlResource,
		NewNetworksSwitchStpResource,
		NewNetworksSyslogServersResource,
		NewNetworksTrafficAnalysisResource,
		NewNetworksWebhooksHTTPServersResource,
		NewNetworksWebhooksPayloadTemplatesResource,
		NewNetworksWirelessAlternateManagementInterfaceResource,
		NewNetworksWirelessBillingResource,
		NewNetworksWirelessBluetoothSettingsResource,
		NewNetworksWirelessRfProfilesResource,
		NewNetworksWirelessSettingsResource,
		NewNetworksWirelessSSIDsResource,
		NewNetworksWirelessSSIDsBonjourForwardingResource,
		NewNetworksWirelessSSIDsDeviceTypeGroupPoliciesResource,
		NewNetworksWirelessSSIDsEapOverrideResource,
		NewNetworksWirelessSSIDsFirewallL3FirewallRulesResource,
		NewNetworksWirelessSSIDsFirewallL7FirewallRulesResource,
		NewNetworksWirelessSSIDsHotspot20Resource,
		NewNetworksWirelessSSIDsIDentityPsksResource,
		NewNetworksWirelessSSIDsSchedulesResource,
		NewNetworksWirelessSSIDsSplashSettingsResource,
		NewNetworksWirelessSSIDsTrafficShapingRulesResource,
		NewNetworksWirelessSSIDsVpnResource,
		NewOrganizationsActionBatchesResource,
		NewOrganizationsAdaptivePolicyACLsResource,
		NewOrganizationsAdaptivePolicyGroupsResource,
		NewOrganizationsAdaptivePolicyPoliciesResource,
		NewOrganizationsAdaptivePolicySettingsResource,
		NewOrganizationsAlertsProfilesResource,
		NewOrganizationsApplianceSecurityIntrusionResource,
		NewOrganizationsApplianceVpnThirdPartyVpnpeersResource,
		NewOrganizationsApplianceVpnVpnFirewallRulesResource,
		NewOrganizationsBrandingPoliciesResource,
		NewOrganizationsBrandingPoliciesPrioritiesResource,
		NewOrganizationsCameraCustomAnalyticsArtifactsResource,
		NewOrganizationsConfigTemplatesResource,
		NewOrganizationsConfigTemplatesSwitchProfilesPortsResource,
		NewOrganizationsEarlyAccessFeaturesOptInsResource,
		NewOrganizationsInsightMonitoredMediaServersResource,
		NewOrganizationsLicensesResource,
		NewOrganizationsLoginSecurityResource,
		NewOrganizationsPolicyObjectsGroupsResource,
		NewOrganizationsPolicyObjectsResource,
		NewOrganizationsSamlResource,
		NewOrganizationsSamlIDpsResource,
		NewOrganizationsSamlRolesResource,
		NewOrganizationsSNMPResource,
		NewDevicesApplianceRadioSettingsResource,
		NewDevicesLiveToolsArpTableResource,
		NewDevicesLiveToolsCableResource,
		NewDevicesLiveToolsWakeOnLanResource,
		NewNetworksApplianceRfProfilesResource,
		NewNetworksVLANProfilesResource,
		NewNetworksWirelessEthernetPortsProfilesResource,
		NewOrganizationsCameraRolesResource,
		NewOrganizationsSmAdminsRolesResource,
		NewDevicesWirelessElectronicShelfLabelResource,
		NewNetworksWirelessElectronicShelfLabelResource,
		NewOrganizationsCellularGatewayEsimsServiceProvidersAccountsResource,
		NewOrganizationsSplashThemesResource,
		NewDevicesApplianceVmxAuthenticationTokenResource,
		NewDevicesBlinkLedsResource,
		NewDevicesCameraGenerateSnapshotResource,
		NewDevicesLiveToolsPingResource,
		NewDevicesLiveToolsPingDeviceResource,
		NewDevicesSwitchPortsCycleResource,
		NewNetworksApplianceTrafficShapingCustomPerformanceClassesResource,
		NewNetworksApplianceWarmSpareSwapResource,
		NewNetworksBindResource,
		NewNetworksClientsProvisionResource,
		NewNetworksDevicesClaimResource,
		NewNetworksDevicesClaimVmxResource,
		NewNetworksDevicesRemoveResource,
		NewNetworksFirmwareUpgradesRollbacksResource,
		NewNetworksFirmwareUpgradesStagedEventsDeferResource,
		NewNetworksFirmwareUpgradesStagedEventsRollbacksResource,
		NewNetworksMqttBrokersResource,
		NewNetworksPiiRequestsDeleteResource,
		NewNetworksSmDevicesCheckinResource,
		NewNetworksSmDevicesFieldsResource,
		NewNetworksSmDevicesLockResource,
		NewNetworksSmDevicesModifyTagsResource,
		NewNetworksSmDevicesMoveResource,
		NewNetworksSmDevicesWipeResource,
		NewNetworksSmDevicesRefreshDetailsResource,
		NewNetworksSmDevicesUnenrollResource,
		NewNetworksSmUserAccessDevicesDeleteResource,
		NewNetworksSplitResource,
		NewNetworksSwitchStacksAddResource,
		NewNetworksSwitchStacksRemoveResource,
		NewNetworksUnbindResource,
		NewOrganizationsClaimResource,
		NewOrganizationsCloneResource,
		NewOrganizationsInventoryClaimResource,
		NewOrganizationsInventoryOnboardingCloudMonitoringExportEventsResource,
		NewOrganizationsInventoryOnboardingCloudMonitoringImportsResource,
		NewOrganizationsInventoryOnboardingCloudMonitoringPrepareResource,
		NewOrganizationsInventoryReleaseResource,
		NewOrganizationsLicensesAssignSeatsResource,
		NewOrganizationsLicensesMoveResource,
		NewOrganizationsLicensesMoveSeatsResource,
		NewOrganizationsLicensesRenewSeatsResource,
		NewOrganizationsLicensingCotermLicensesMoveResource,
		NewOrganizationsNetworksCombineResource,
		NewOrganizationsSwitchDevicesCloneResource,
		NewAdministeredLicensingSubscriptionSubscriptionsClaimResource,
		NewAdministeredLicensingSubscriptionSubscriptionsClaimKeyValidateResource,
		NewAdministeredLicensingSubscriptionSubscriptionsBindResource,
		NewDevicesWirelessAlternateManagementInterfaceIPv6Resource,
		NewNetworksApplianceTrafficShapingVpnExclusionsResource,
		NewNetworksSmDevicesRebootResource,
		NewNetworksSmDevicesShutdownResource,
		NewNetworksSmDevicesInstallAppsResource,
		NewNetworksSmDevicesUninstallAppsResource,
		NewNetworksVLANProfilesAssignmentsReassignResource,
		NewNetworksWirelessEthernetPortsProfilesAssignResource,
		NewNetworksWirelessEthernetPortsProfilesSetDefaultResource,
		NewOrganizationsSmSentryPoliciesAssignmentsResource,
		NewAdministeredIDentitiesMeAPIKeysGenerateResource,
		NewAdministeredIDentitiesMeAPIKeysRevokeResource,
		NewDevicesLiveToolsLedsBlinkResource,
		NewNetworksApplianceSdwanInternetPoliciesResource,
		NewNetworksFloorPlansAutoLocateJobsBatchResource,
		NewNetworksCancelResource,
		NewNetworksPublishResource,
		NewNetworksRecalculateResource,
		NewNetworksFloorPlansDevicesBatchUpdateResource,
		NewNetworksWirelessAirMarshalRulesResource,
		NewNetworksWirelessAirMarshalRulesUpdateResource,
		NewNetworksWirelessAirMarshalRulesDeleteResource,
		NewNetworksWirelessAirMarshalSettingsResource,
		NewOrganizationsAssuranceAlertsDismissResource,
		NewOrganizationsAssuranceAlertsRestoreResource,
		NewOrganizationsCellularGatewayEsimsSwapResource,
		NewOrganizationsDevicesDetailsBulkUpdateResource,
		NewOrganizationsAssetsResource,
		NewOrganizationsWirelessRadioAutoRfChannelsRecalculateResource,
	}
}

func (p *MerakiProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewOrganizationsDataSource,
		NewOrganizationsAdminsDataSource,
		NewAdministeredIDentitiesMeDataSource,
		NewDevicesDataSource,
		NewDevicesAppliancePerformanceDataSource,
		NewDevicesApplianceUplinksSettingsDataSource,
		NewDevicesCameraAnalyticsLiveDataSource,
		NewDevicesCameraCustomAnalyticsDataSource,
		NewDevicesCameraQualityAndRetentionDataSource,
		NewDevicesCameraSenseDataSource,
		NewDevicesCameraVideoSettingsDataSource,
		NewDevicesCameraVideoLinkDataSource,
		NewDevicesCameraWirelessProfilesDataSource,
		NewDevicesCellularSimsDataSource,
		NewDevicesCellularGatewayLanDataSource,
		NewDevicesCellularGatewayPortForwardingRulesDataSource,
		NewDevicesLiveToolsPingInfoDataSource,
		NewDevicesLiveToolsPingDeviceInfoDataSource,
		NewDevicesLldpCdpDataSource,
		NewDevicesManagementInterfaceDataSource,
		NewDevicesSensorRelationshipsDataSource,
		NewDevicesSwitchPortsDataSource,
		NewDevicesSwitchPortsStatusesDataSource,
		NewDevicesSwitchRoutingInterfacesDataSource,
		NewDevicesSwitchRoutingInterfacesDhcpDataSource,
		NewDevicesSwitchRoutingStaticRoutesDataSource,
		NewDevicesSwitchWarmSpareDataSource,
		NewDevicesWirelessBluetoothSettingsDataSource,
		NewDevicesWirelessConnectionStatsDataSource,
		NewDevicesWirelessLatencyStatsDataSource,
		NewDevicesWirelessRadioSettingsDataSource,
		NewDevicesWirelessStatusDataSource,
		NewNetworksDataSource,
		NewNetworksAlertsHistoryDataSource,
		NewNetworksAlertsSettingsDataSource,
		NewNetworksApplianceConnectivityMonitoringDestinationsDataSource,
		NewNetworksApplianceContentFilteringDataSource,
		NewNetworksApplianceStaticRoutesDataSource,
		NewNetworksApplianceContentFilteringCategoriesDataSource,
		NewNetworksApplianceFirewallCellularFirewallRulesDataSource,
		NewNetworksApplianceFirewallFirewalledServicesDataSource,
		NewNetworksApplianceFirewallInboundFirewallRulesDataSource,
		NewNetworksApplianceFirewallL3FirewallRulesDataSource,
		NewNetworksApplianceFirewallL7FirewallRulesDataSource,
		NewNetworksApplianceFirewallL7FirewallRulesApplicationCategoriesDataSource,
		NewNetworksApplianceFirewallOneToManyNatRulesDataSource,
		NewNetworksApplianceFirewallOneToOneNatRulesDataSource,
		NewNetworksApplianceFirewallPortForwardingRulesDataSource,
		NewNetworksApplianceFirewallSettingsDataSource,
		NewNetworksAppliancePortsDataSource,
		NewNetworksAppliancePrefixesDelegatedStaticsDataSource,
		NewNetworksApplianceSecurityIntrusionDataSource,
		NewNetworksApplianceSecurityMalwareDataSource,
		NewNetworksApplianceSettingsDataSource,
		NewNetworksApplianceSingleLanDataSource,
		NewNetworksApplianceSSIDsDataSource,
		NewNetworksApplianceTrafficShapingDataSource,
		NewNetworksApplianceTrafficShapingRulesDataSource,
		NewNetworksApplianceTrafficShapingUplinkBandwidthDataSource,
		NewNetworksApplianceTrafficShapingUplinkSelectionDataSource,
		NewNetworksApplianceVLANsSettingsDataSource,
		NewNetworksApplianceVLANsDataSource,
		NewNetworksApplianceVpnBgpDataSource,
		NewNetworksApplianceVpnSiteToSiteVpnDataSource,
		NewNetworksApplianceWarmSpareDataSource,
		NewNetworksBluetoothClientsDataSource,
		NewNetworksCameraQualityRetentionProfilesDataSource,
		NewNetworksCameraWirelessProfilesDataSource,
		NewNetworksCellularGatewayConnectivityMonitoringDestinationsDataSource,
		NewNetworksCellularGatewayDhcpDataSource,
		NewNetworksCellularGatewaySubnetPoolDataSource,
		NewNetworksCellularGatewayUplinkDataSource,
		NewNetworksClientsDataSource,
		NewNetworksClientsOverviewDataSource,
		NewNetworksClientsPolicyDataSource,
		NewNetworksClientsSplashAuthorizationStatusDataSource,
		NewNetworksEventsDataSource,
		NewNetworksEventsEventTypesDataSource,
		NewNetworksFirmwareUpgradesDataSource,
		NewNetworksFirmwareUpgradesStagedEventsDataSource,
		NewNetworksFirmwareUpgradesStagedGroupsDataSource,
		NewNetworksFirmwareUpgradesStagedStagesDataSource,
		NewNetworksFloorPlansDataSource,
		NewNetworksGroupPoliciesDataSource,
		NewNetworksHealthAlertsDataSource,
		NewNetworksInsightApplicationsHealthByTimeDataSource,
		NewNetworksMerakiAuthUsersDataSource,
		NewNetworksNetflowDataSource,
		NewNetworksPiiPiiKeysDataSource,
		NewNetworksPiiRequestsDataSource,
		NewNetworksPiiSmDevicesForKeyDataSource,
		NewNetworksPiiSmOwnersForKeyDataSource,
		NewNetworksPoliciesByClientDataSource,
		NewNetworksSensorAlertsCurrentOverviewByMetricDataSource,
		NewNetworksSensorAlertsOverviewByMetricDataSource,
		NewNetworksSensorAlertsProfilesDataSource,
		NewNetworksSensorMqttBrokersDataSource,
		NewNetworksSensorRelationshipsDataSource,
		NewNetworksSettingsDataSource,
		NewNetworksSmBypassActivationLockAttemptsDataSource,
		NewNetworksSmDevicesDataSource,
		NewNetworksSmDevicesCellularUsageHistoryDataSource,
		NewNetworksSmDevicesCertsDataSource,
		NewNetworksSmDevicesConnectivityDataSource,
		NewNetworksSmDevicesDesktopLogsDataSource,
		NewNetworksSmDevicesDeviceCommandLogsDataSource,
		NewNetworksSmDevicesDeviceProfilesDataSource,
		NewNetworksSmDevicesNetworkAdaptersDataSource,
		NewNetworksSmDevicesPerformanceHistoryDataSource,
		NewNetworksSmDevicesSecurityCentersDataSource,
		NewNetworksSmDevicesWLANListsDataSource,
		NewNetworksSmProfilesDataSource,
		NewNetworksSmTargetGroupsDataSource,
		NewNetworksSmTrustedAccessConfigsDataSource,
		NewNetworksSmUserAccessDevicesDataSource,
		NewNetworksSmUsersDataSource,
		NewNetworksSmUsersDeviceProfilesDataSource,
		NewNetworksSmUsersSoftwaresDataSource,
		NewNetworksSNMPDataSource,
		NewNetworksSwitchAccessControlListsDataSource,
		NewNetworksSwitchAccessPoliciesDataSource,
		NewNetworksSwitchAlternateManagementInterfaceDataSource,
		NewNetworksSwitchDhcpV4ServersSeenDataSource,
		NewNetworksSwitchDhcpServerPolicyDataSource,
		NewNetworksSwitchDhcpServerPolicyArpInspectionTrustedServersDataSource,
		NewNetworksSwitchDhcpServerPolicyArpInspectionWarningsByDeviceDataSource,
		NewNetworksSwitchDscpToCosMappingsDataSource,
		NewNetworksSwitchLinkAggregationsDataSource,
		NewNetworksSwitchMtuDataSource,
		NewNetworksSwitchPortSchedulesDataSource,
		NewNetworksSwitchQosRulesOrderDataSource,
		NewNetworksSwitchRoutingMulticastDataSource,
		NewNetworksSwitchRoutingMulticastRendezvousPointsDataSource,
		NewNetworksSwitchRoutingOspfDataSource,
		NewNetworksSwitchSettingsDataSource,
		NewNetworksSwitchStacksDataSource,
		NewNetworksSwitchStacksRoutingInterfacesDataSource,
		NewNetworksSwitchStacksRoutingInterfacesDhcpDataSource,
		NewNetworksSwitchStacksRoutingStaticRoutesDataSource,
		NewNetworksSwitchStormControlDataSource,
		NewNetworksSwitchStpDataSource,
		NewNetworksSyslogServersDataSource,
		NewNetworksTopologyLinkLayerDataSource,
		NewNetworksTrafficAnalysisDataSource,
		NewNetworksTrafficShapingApplicationCategoriesDataSource,
		NewNetworksTrafficShapingDscpTaggingOptionsDataSource,
		NewNetworksWebhooksHTTPServersDataSource,
		NewNetworksWebhooksPayloadTemplatesDataSource,
		NewNetworksWebhooksWebhookTestsDataSource,
		NewNetworksWirelessAlternateManagementInterfaceDataSource,
		NewNetworksWirelessBillingDataSource,
		NewNetworksWirelessBluetoothSettingsDataSource,
		NewNetworksWirelessChannelUtilizationHistoryDataSource,
		NewNetworksWirelessClientCountHistoryDataSource,
		NewNetworksWirelessClientsConnectionStatsDataSource,
		NewNetworksWirelessClientsLatencyStatsDataSource,
		NewNetworksWirelessConnectionStatsDataSource,
		NewNetworksWirelessDataRateHistoryDataSource,
		NewNetworksWirelessDevicesConnectionStatsDataSource,
		NewNetworksWirelessFailedConnectionsDataSource,
		NewNetworksWirelessLatencyHistoryDataSource,
		NewNetworksWirelessLatencyStatsDataSource,
		NewNetworksWirelessMeshStatusesDataSource,
		NewNetworksWirelessRfProfilesDataSource,
		NewNetworksWirelessSettingsDataSource,
		NewNetworksWirelessSignalQualityHistoryDataSource,
		NewNetworksWirelessSSIDsDataSource,
		NewNetworksWirelessSSIDsBonjourForwardingDataSource,
		NewNetworksWirelessSSIDsDeviceTypeGroupPoliciesDataSource,
		NewNetworksWirelessSSIDsEapOverrideDataSource,
		NewNetworksWirelessSSIDsFirewallL3FirewallRulesDataSource,
		NewNetworksWirelessSSIDsFirewallL7FirewallRulesDataSource,
		NewNetworksWirelessSSIDsHotspot20DataSource,
		NewNetworksWirelessSSIDsIDentityPsksDataSource,
		NewNetworksWirelessSSIDsSchedulesDataSource,
		NewNetworksWirelessSSIDsSplashSettingsDataSource,
		NewNetworksWirelessSSIDsTrafficShapingRulesDataSource,
		NewNetworksWirelessSSIDsVpnDataSource,
		NewNetworksWirelessUsageHistoryDataSource,
		NewOrganizationsActionBatchesDataSource,
		NewOrganizationsAdaptivePolicyACLsDataSource,
		NewOrganizationsAdaptivePolicyGroupsDataSource,
		NewOrganizationsAdaptivePolicyOverviewDataSource,
		NewOrganizationsAdaptivePolicyPoliciesDataSource,
		NewOrganizationsAdaptivePolicySettingsDataSource,
		NewOrganizationsAlertsProfilesDataSource,
		NewOrganizationsAPIRequestsDataSource,
		NewOrganizationsAPIRequestsOverviewDataSource,
		NewOrganizationsAPIRequestsOverviewResponseCodesByIntervalDataSource,
		NewOrganizationsApplianceSecurityIntrusionDataSource,
		NewOrganizationsApplianceVpnThirdPartyVpnpeersDataSource,
		NewOrganizationsApplianceVpnVpnFirewallRulesDataSource,
		NewOrganizationsBrandingPoliciesDataSource,
		NewOrganizationsBrandingPoliciesPrioritiesDataSource,
		NewOrganizationsCameraCustomAnalyticsArtifactsDataSource,
		NewOrganizationsCellularGatewayUplinkStatusesDataSource,
		NewOrganizationsClientsBandwidthUsageHistoryDataSource,
		NewOrganizationsClientsOverviewDataSource,
		NewOrganizationsClientsSearchDataSource,
		NewOrganizationsConfigTemplatesDataSource,
		NewOrganizationsConfigTemplatesSwitchProfilesDataSource,
		NewOrganizationsConfigTemplatesSwitchProfilesPortsDataSource,
		NewOrganizationsDevicesDataSource,
		NewOrganizationsDevicesAvailabilitiesDataSource,
		NewOrganizationsDevicesPowerModulesStatusesByDeviceDataSource,
		NewOrganizationsDevicesProvisioningStatusesDataSource,
		NewOrganizationsDevicesStatusesDataSource,
		NewOrganizationsDevicesStatusesOverviewDataSource,
		NewOrganizationsDevicesUplinksAddressesByDeviceDataSource,
		NewOrganizationsDevicesUplinksLossAndLatencyDataSource,
		NewOrganizationsEarlyAccessFeaturesDataSource,
		NewOrganizationsEarlyAccessFeaturesOptInsDataSource,
		NewOrganizationsFirmwareUpgradesDataSource,
		NewOrganizationsFirmwareUpgradesByDeviceDataSource,
		NewOrganizationsInsightApplicationsDataSource,
		NewOrganizationsInsightMonitoredMediaServersDataSource,
		NewOrganizationsInventoryDevicesDataSource,
		NewOrganizationsInventoryOnboardingCloudMonitoringImportsInfoDataSource,
		NewOrganizationsInventoryOnboardingCloudMonitoringNetworksDataSource,
		NewOrganizationsLicensesDataSource,
		NewOrganizationsLicensesOverviewDataSource,
		NewOrganizationsLicensingCotermLicensesDataSource,
		NewOrganizationsLoginSecurityDataSource,
		NewOrganizationsOpenapiSpecDataSource,
		NewOrganizationsPolicyObjectsGroupsDataSource,
		NewOrganizationsPolicyObjectsDataSource,
		NewOrganizationsSamlDataSource,
		NewOrganizationsSamlIDpsDataSource,
		NewOrganizationsSamlRolesDataSource,
		NewOrganizationsSensorReadingsHistoryDataSource,
		NewOrganizationsSensorReadingsLatestDataSource,
		NewOrganizationsSmApnsCertDataSource,
		NewOrganizationsSmVppAccountsDataSource,
		NewOrganizationsSNMPDataSource,
		NewOrganizationsSummaryTopAppliancesByUtilizationDataSource,
		NewOrganizationsSummaryTopClientsByUsageDataSource,
		NewOrganizationsSummaryTopClientsManufacturersByUsageDataSource,
		NewOrganizationsSummaryTopDevicesByUsageDataSource,
		NewOrganizationsSummaryTopDevicesModelsByUsageDataSource,
		NewOrganizationsSummaryTopSSIDsByUsageDataSource,
		NewOrganizationsSummaryTopSwitchesByEnergyUsageDataSource,
		NewOrganizationsSwitchPortsBySwitchDataSource,
		NewOrganizationsUplinksStatusesDataSource,
		NewOrganizationsWebhooksLogsDataSource,
		NewOrganizationsWirelessDevicesEthernetStatusesDataSource,
		NewAdministeredLicensingSubscriptionEntitlementsDataSource,
		NewAdministeredLicensingSubscriptionSubscriptionsDataSource,
		NewAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesDataSource,
		NewDevicesApplianceRadioSettingsDataSource,
		NewDevicesLiveToolsArpTableDataSource,
		NewDevicesLiveToolsCableDataSource,
		NewDevicesLiveToolsWakeOnLanDataSource,
		NewNetworksApplianceRfProfilesDataSource,
		NewNetworksVLANProfilesDataSource,
		NewNetworksVLANProfilesAssignmentsByDeviceDataSource,
		NewNetworksWirelessEthernetPortsProfilesDataSource,
		NewOrganizationsApplianceTrafficShapingVpnExclusionsByNetworkDataSource,
		NewOrganizationsApplianceUplinksStatusesOverviewDataSource,
		NewOrganizationsApplianceUplinksUsageByNetworkDataSource,
		NewOrganizationsCameraBoundariesAreasByDeviceDataSource,
		NewOrganizationsCameraBoundariesLinesByDeviceDataSource,
		NewOrganizationsCameraDetectionsHistoryByBoundaryByIntervalDataSource,
		NewOrganizationsCameraPermissionsDataSource,
		NewOrganizationsCameraRolesDataSource,
		NewOrganizationsDevicesAvailabilitiesChangeHistoryDataSource,
		NewOrganizationsSmAdminsRolesDataSource,
		NewOrganizationsSmSentryPoliciesAssignmentsByNetworkDataSource,
		NewOrganizationsSummaryTopNetworksByStatusDataSource,
		NewOrganizationsWebhooksCallbacksStatusesDataSource,
		NewOrganizationsWirelessDevicesChannelUtilizationByDeviceDataSource,
		NewOrganizationsWirelessDevicesChannelUtilizationByNetworkDataSource,
		NewOrganizationsWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalDataSource,
		NewOrganizationsWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalDataSource,
		NewOrganizationsWirelessDevicesPacketLossByClientDataSource,
		NewOrganizationsWirelessDevicesPacketLossByDeviceDataSource,
		NewOrganizationsWirelessDevicesPacketLossByNetworkDataSource,
		NewAdministeredIDentitiesMeAPIKeysDataSource,
		NewDevicesWirelessElectronicShelfLabelDataSource,
		NewNetworksWirelessElectronicShelfLabelDataSource,
		NewNetworksWirelessElectronicShelfLabelConfiguredDevicesDataSource,
		NewOrganizationsAssuranceAlertsDataSource,
		NewOrganizationsAssuranceAlertsOverviewDataSource,
		NewOrganizationsAssuranceAlertsOverviewByNetworkDataSource,
		NewOrganizationsAssuranceAlertsOverviewByTypeDataSource,
		NewOrganizationsAssuranceAlertsOverviewHistoricalDataSource,
		NewOrganizationsCellularGatewayEsimsInventoryDataSource,
		NewOrganizationsCellularGatewayEsimsServiceProvidersDataSource,
		NewOrganizationsCellularGatewayEsimsServiceProvidersAccountsDataSource,
		NewOrganizationsCellularGatewayEsimsServiceProvidersAccountsCommunicationPlansDataSource,
		NewOrganizationsCellularGatewayEsimsServiceProvidersAccountsRatePlansDataSource,
		NewOrganizationsDevicesOverviewByModelDataSource,
		NewOrganizationsFloorPlansAutoLocateDevicesDataSource,
		NewOrganizationsFloorPlansAutoLocateStatusesDataSource,
		NewOrganizationsSplashThemesDataSource,
		NewOrganizationsSummaryTopApplicationsByUsageDataSource,
		NewOrganizationsSummaryTopApplicationsCategoriesByUsageDataSource,
		NewOrganizationsSwitchPortsClientsOverviewByDeviceDataSource,
		NewOrganizationsSwitchPortsOverviewDataSource,
		NewOrganizationsSwitchPortsStatusesBySwitchDataSource,
		NewOrganizationsSwitchPortsTopologyDiscoveryByDeviceDataSource,
		NewOrganizationsWirelessAirMarshalRulesDataSource,
		NewOrganizationsWirelessAirMarshalSettingsByNetworkDataSource,
		NewOrganizationsWirelessClientsOverviewByDeviceDataSource,
		NewOrganizationsWirelessDevicesWirelessControllersByDeviceDataSource,
		NewOrganizationsWirelessRfProfilesAssignmentsByDeviceDataSource,
		NewOrganizationsWirelessSSIDsStatusesByDeviceDataSource,
		NewOrganizationsWirelessControllerAvailabilitiesChangeHistoryDataSource,
		NewOrganizationsWirelessControllerClientsOverviewHistoryByDeviceByIntervalDataSource,
		NewOrganizationsWirelessControllerConnectionsDataSource,
		NewOrganizationsWirelessControllerDevicesInterfacesL2ByDeviceDataSource,
		NewOrganizationsWirelessControllerDevicesInterfacesL2StatusesChangeHistoryByDeviceDataSource,
		NewOrganizationsWirelessControllerDevicesInterfacesL2UsageHistoryByIntervalDataSource,
		NewOrganizationsWirelessControllerDevicesInterfacesL3ByDeviceDataSource,
		NewOrganizationsWirelessControllerDevicesInterfacesL3StatusesChangeHistoryByDeviceDataSource,
		NewOrganizationsWirelessControllerDevicesInterfacesL3UsageHistoryByIntervalDataSource,
		NewOrganizationsWirelessControllerDevicesInterfacesPacketsOverviewByDeviceDataSource,
		NewOrganizationsWirelessControllerDevicesInterfacesUsageHistoryByIntervalDataSource,
		NewOrganizationsWirelessControllerDevicesRedundancyFailoverHistoryDataSource,
		NewOrganizationsWirelessControllerDevicesRedundancyStatusesDataSource,
		NewOrganizationsWirelessControllerDevicesSystemUtilizationHistoryByIntervalDataSource,
		NewOrganizationsWirelessControllerOverviewByDeviceDataSource,
	}
}
