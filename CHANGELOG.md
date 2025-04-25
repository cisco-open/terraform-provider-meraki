## Unreleased (- -, -)
## 1.1.2-beta (April 25, 2025)
BUGFIXES:
* Changed the type of the VLAN ID attribute from String to Int64 for better data integrity.

## 1.1.1-beta (April 24, 2025)
BUGFIXES:
* remove `vlan_id` attribute in `resource_meraki_networks_appliance_vlans.go`

## 1.1.0-beta (April 22, 2025)
FEATURES:
* Provider supports v1.57.0 of Meraki Dashboard API.
* **New Data Source** `data_source_meraki_devices_sensor_commands.go`
* **New Data Source** `data_source_meraki_organizations_appliance_dns_local_profiles.go`
* **New Data Source** `data_source_meraki_organizations_appliance_dns_local_profiles_assignments.go`
* **New Data Source** `data_source_meraki_organizations_appliance_dns_local_records.go`
* **New Data Source** `data_source_meraki_organizations_appliance_dns_split_profiles.go`
* **New Data Source** `data_source_meraki_organizations_appliance_dns_split_profiles_assignments.go`
* **New Data Source** `data_source_meraki_organizations_appliance_firewall_multicast_forwarding_by_network.go`
* **New Data Source** `data_source_meraki_organizations_devices_controller_migrations.go`
* **New Data Source** `data_source_meraki_organizations_devices_system_memory_usage_history_by_interval.go`
* **New Data Source** `data_source_meraki_organizations_integrations_xdr_networks.go`
* **New Data Source** `data_source_meraki_organizations_switch_ports_usage_history_by_device_by_interval.go`
* **New Data Source** `data_source_meraki_organizations_wireless_devices_power_mode_history.go`
* **New Data Source** `data_source_meraki_organizations_wireless_devices_system_cpu_load_history.go`
* **New Data Source** `data_source_meraki_organizations_wireless_ssids_firewall_isolation_allowlist_entries.go`
* **New Resource** `resource_meraki_devices_live_tools_throughput_test.go`
* **New Resource** `resource_meraki_devices_sensor_commands_create.go`
* **New Resource** `resource_meraki_networks_appliance_firewall_multicast_forwarding.go`
* **New Resource** `resource_meraki_organizations_appliance_dns_local_profiles.go`
* **New Resource** `resource_meraki_organizations_appliance_dns_local_profiles_assignments_bulk_create.go`
* **New Resource** `resource_meraki_organizations_appliance_dns_local_profiles_assignments_bulk_delete.go`
* **New Resource** `resource_meraki_organizations_appliance_dns_local_records.go`
* **New Resource** `resource_meraki_organizations_appliance_dns_split_profiles.go`
* **New Resource** `resource_meraki_organizations_appliance_dns_split_profiles_assignments_bulk_create.go`
* **New Resource** `resource_meraki_organizations_appliance_dns_split_profiles_assignments_bulk_delete.go`
* **New Resource** `resource_meraki_organizations_devices_controller_migrations_create.go`
* **New Resource** `resource_meraki_organizations_integrations_xdr_networks_disable.go`
* **New Resource** `resource_meraki_organizations_integrations_xdr_networks_enable.go`
* **New Resource** `resource_meraki_organizations_licenses_renew_seats.go`
* **New Resource** `resource_meraki_organizations_wireless_ssids_firewall_isolation_allowlist_entries_create.go`
* **New Resource** `resource_meraki_organizations_wireless_ssids_firewall_isolation_allowlist_entries_delete.go`
* **New Resource** `resource_meraki_organizations_wireless_ssids_firewall_isolation_allowlist_entries_update.go`

## 1.0.7-beta (March 31, 2025)
BUGFIXES:
* meraki_organizations_policy_objects data source cannot unmarshal issue. [Fixed] #232
* meraki_devices_appliance_vmx_authentication_token provider error. [Fixed] #237

## 1.0.6-beta (March 10, 2025)
BUGFIXES:
* meraki_organizations_saml_idps not mapped at the end [Fixed]. #225
* meraki_organizations_camera_roles including ID parameter. #226

## 1.0.5-beta (March 05, 2025)
BUGFIXES:
* Cannot create meraki_organizations_saml_idps resources #225.
* Cannot create meraki_organizations_saml_roles resources #224.
* Cannot import meraki_organizations_camera_roles resources #226.

## 1.0.4-beta (February 27, 2025)
BUGFIXES:
* resource meraki_networks_wireless_ssids - order of list of radius servers not maintained when applied. #218
* resource meraki_networks_switch_access_policies - list of Radius servers not maintaining order. #217

## 1.0.3-beta (February 24, 2025)
BUGFIXES:
* resource meraki_networks_switch_access_policies - guest_vlan_id not being saved #210.
* resource meraki_networks_wireless_ssids - enterprise_admin_access trying to be flagged into state though not configured #211.
* resource "meraki_networks_devices_claim" broken #215.

## 1.0.3-beta (February 13, 2025)
BUGFIXES:
* Update import syntax in Meraki resource documentation and examples to reflect changes in required parameters. #201
* meraki_networks_appliance_vpn_site_to_site_vpn, impossible to manage priority order. Changing to list to preserve order. #203
* Resource 'meraki_networks' possibly missing required argument in documentation. Documentation updated. #202
* Resource 'meraki_networks' map networkIds #181.

## 1.0.1-beta (January 20, 2025)
FEATURES:
* **New Data Source** `data_source_meraki_networks_wireless_air_marshal_rules.go`
* Added support for `-1` value in the `per_page` parameter of the data source. When set to `-1`, the response will include all available data without pagination.

## 1.0.0-beta (January 20, 2025)
BREAKING CHANGES:
* Resource `resource_meraki_organizations_users.go` has been removed.
FEATURES:
* **New Data Source** `data_source_meraki_administered_identities_me_api_keys.go`
* **New Data Source** `data_source_meraki_devices_wireless_electronic_shelf_label.go`
* **New Data Source** `data_source_meraki_networks_wireless_electronic_shelf_label.go`
* **New Data Source** `data_source_meraki_networks_wireless_electronic_shelf_label_configured_devices.go`
* **New Data Source** `data_source_meraki_organizations_assurance_alerts.go`
* **New Data Source** `data_source_meraki_organizations_assurance_alerts_overview.go`
* **New Data Source** `data_source_meraki_organizations_assurance_alerts_overview_by_network.go`
* **New Data Source** `data_source_meraki_organizations_assurance_alerts_overview_by_type.go`
* **New Data Source** `data_source_meraki_organizations_assurance_alerts_overview_historical.go`
* **New Data Source** `data_source_meraki_organizations_cellular_gateway_esims_inventory.go`
* **New Data Source** `data_source_meraki_organizations_cellular_gateway_esims_service_providers.go`
* **New Data Source** `data_source_meraki_organizations_cellular_gateway_esims_service_providers_accounts.go`
* **New Data Source** `data_source_meraki_organizations_cellular_gateway_esims_service_providers_accounts_communication_plans.go`
* **New Data Source** `data_source_meraki_organizations_cellular_gateway_esims_service_providers_accounts_rate_plans.go`
* **New Data Source** `data_source_meraki_organizations_devices_overview_by_model.go`
* **New Data Source** `data_source_meraki_organizations_floor_plans_auto_locate_devices.go`
* **New Data Source** `data_source_meraki_organizations_floor_plans_auto_locate_statuses.go`
* **New Data Source** `data_source_meraki_organizations_splash_themes.go`
* **New Data Source** `data_source_meraki_organizations_summary_top_applications_by_usage.go`
* **New Data Source** `data_source_meraki_organizations_summary_top_applications_categories_by_usage.go`
* **New Data Source** `data_source_meraki_organizations_switch_ports_clients_overview_by_device.go`
* **New Data Source** `data_source_meraki_organizations_switch_ports_overview.go`
* **New Data Source** `data_source_meraki_organizations_switch_ports_statuses_by_switch.go`
* **New Data Source** `data_source_meraki_organizations_switch_ports_topology_discovery_by_device.go`
* **New Data Source** `data_source_meraki_organizations_wireless_air_marshal_rules.go`
* **New Data Source** `data_source_meraki_organizations_wireless_air_marshal_settings_by_network.go`
* **New Data Source** `data_source_meraki_organizations_wireless_clients_overview_by_device.go`
* **New Data Source** `data_source_meraki_organizations_wireless_controller_availabilities_change_history.go`
* **New Data Source** `data_source_meraki_organizations_wireless_controller_clients_overview_history_by_device_by_interval.go`
* **New Data Source** `data_source_meraki_organizations_wireless_controller_connections.go`
* **New Data Source** `data_source_meraki_organizations_wireless_controller_devices_interfaces_l2_by_device.go`
* **New Data Source** `data_source_meraki_organizations_wireless_controller_devices_interfaces_l2_statuses_change_history_by_device.go`
* **New Data Source** `data_source_meraki_organizations_wireless_controller_devices_interfaces_l2_usage_history_by_interval.go`
* **New Data Source** `data_source_meraki_organizations_wireless_controller_devices_interfaces_l3_by_device.go`
* **New Data Source** `data_source_meraki_organizations_wireless_controller_devices_interfaces_l3_statuses_change_history_by_device.go`
* **New Data Source** `data_source_meraki_organizations_wireless_controller_devices_interfaces_l3_usage_history_by_interval.go`
* **New Data Source** `data_source_meraki_organizations_wireless_controller_devices_interfaces_packets_overview_by_device.go`
* **New Data Source** `data_source_meraki_organizations_wireless_controller_devices_interfaces_usage_history_by_interval.`go
* **New Data Source** `data_source_meraki_organizations_wireless_controller_devices_redundancy_failover_history.go`
* **New Data Source** `data_source_meraki_organizations_wireless_controller_devices_redundancy_statuses.go`
* **New Data Source** `data_source_meraki_organizations_wireless_controller_devices_system_utilization_history_by_interval.go`
* **New Data Source** `data_source_meraki_organizations_wireless_controller_overview_by_device.go`
* **New Data Source** `data_source_meraki_organizations_wireless_devices_wireless_controllers_by_device.go`
* **New Data Source** `data_source_meraki_organizations_wireless_rf_profiles_assignments_by_device.go`
* **New Data Source** `data_source_meraki_organizations_wireless_ssids_statuses_by_device.go`
* **New Resource** `resource_meraki_administered_identities_me_api_keys_generate.go`
* **New Resource** `resource_meraki_administered_identities_me_api_keys_revoke.go`
* **New Resource** `resource_meraki_devices_live_tools_leds_blink.go`
* **New Resource** `resource_meraki_devices_wireless_electronic_shelf_label.go`
* **New Resource** `resource_meraki_networks_appliance_sdwan_internet_policies.go`
* **New Resource** `resource_meraki_networks_cancel.go`
* **New Resource** `resource_meraki_networks_floor_plans_auto_locate_jobs_batch.go`
* **New Resource** `resource_meraki_networks_floor_plans_devices_batch_update.go`
* **New Resource** `resource_meraki_networks_publish.go`
* **New Resource** `resource_meraki_networks_recalculate.go`
* **New Resource** `resource_meraki_networks_wireless_air_marshal_rules.go`
* **New Resource** `resource_meraki_networks_wireless_air_marshal_rules_delete.go`
* **New Resource** `resource_meraki_networks_wireless_air_marshal_rules_update.go`
* **New Resource** `resource_meraki_networks_wireless_air_marshal_settings.go`
* **New Resource** `resource_meraki_networks_wireless_electronic_shelf_label.go`
* **New Resource** `resource_meraki_organizations_assets.go`
* **New Resource** `resource_meraki_organizations_assurance_alerts_dismiss.go`
* **New Resource** `resource_meraki_organizations_assurance_alerts_restore.go`
* **New Resource** `resource_meraki_organizations_cellular_gateway_esims_service_providers_accounts.go`
* **New Resource** `resource_meraki_organizations_cellular_gateway_esims_swap.go`
* **New Resource** `resource_meraki_organizations_devices_details_bulk_update.go`
* **New Resource** `resource_meraki_organizations_licenses_renew_seats.go`
* **New Resource** `resource_meraki_organizations_splash_themes.go`
* **New Resource** `resource_meraki_organizations_wireless_radio_auto_rf_channels_recalculate.go`
* **New Resource** `resource_meraki_organizations_licenses_renew_seats.go`
IMPROVEMENTS:
* Provider supports v1.53.0 of Meraki Dashboard API.

## 0.2.13-alpha (November 27, 2024)
BUGFIXES:
* meraki_debug no longer works #179 [fixed].
* appliance_traffic_shaping_rules does not store order #173, Changing to list to preserve the order.
* Terraform registry documentation updates #180.
* meraki_organizations_policy_objects_groups broken, both the data and resource #178, documentation issue, changing types to umarshal struct.
FEATURES:
* **New Resource** `resource_meraki_networks_appliance_static_routes`
* **New Data Source** `data_source_meraki_networks_appliance_static_routes`

## 0.2.12-alpha (October 01, 2024)
BUGFIXES:
* Updating logs to only english.
* meraki_networks_wireless_ssids fix in request, ignore empty fields.
* meraki_networks_group_policies fix in request, ignore empty fields.
* meraki_organizations_snmp fixing only read fields.
* meraki_networks_appliance_traffic_shaping_rules input variable issue #134. Doc updated.
* meraki_networks_switch_dhcp_server_policy_arp_inspection_trusted_servers read method fixed.
* meraki_networks_group_policies Fail in toSdkFunction.
* meraki_networks_switch_link_aggregations, ID not mapped in delete method.
* meraki_networks_switch_link_aggregations, Incorrect error handling.
* meraki_organizations_admins fixing import.
* meraki_organizations_appliance_vpn_third_party_vpnpeers, fixig array converion fail.

## 0.2.11-alpha (August 19, 2024)
BUGFIXES:
* problematic field is ipv6: {} in the body.
* peer_ips field detecting unexpected changes #123.
* provider produced inconsistent result after apply, meraki_organizations_early_access_features_opt_ins.
* meraki_networks_appliance_traffic_shaping_rules input variable issue #134. Doc updated.

## 0.2.10-alpha (August 05, 2024)
BUGFIXES:
* Issue #117 Fixed.
* Issue #118 Fixed.
* Issue #120 Fixed.
* Issue #121 Fixed.
* Issue #122 Fixed.
* Issue #123 Fixed.
* Issue #124 Fixed.
* Issue #125 Fixed.
* Issue #126 Fixed.
* Issue #127 Fixed.
* Issue #128 Fixed.
* Issue #129 Fixed.
* Issue #130 Fixed.
* Issue #131 Fixed.

## 0.2.9-alpha (July 23, 2024)
BUGFIXES:
* meraki_networks_switch_stp - plugin crash when reading #114.
* Failure when executing GetNetworkSwitchRoutingOspf #115.
* meraki_networks_vlan_profiles silently fails when attempting to create a resource on the default profile #94.
* Resource meraki_networks_sensor_alerts_profiles error with operation CreateNetworkSensorAlertsProfile #91.
* Resource meraki_networks_wireless_rf_profiles error when trying to create new profile.
#88.

## 0.2.8-alpha (July 16, 2024)
BUGFIXES:
* meraki_devices_management_interface error on import state #109.
* Encountering panic in latest version when attempting to create meraki_networks (Resource) #111. Solved in utils file.

## 0.2.6-alpha (July 15, 2024)
BUGFIXES:
* Issue #88 fixed.
* Issue #89 fixed.
* Issue #92 fixed.
* Issue #93 fixed.
* Issue #94 fixed.
* Issue #96 fixed.
* Issue #97 fixed.
* Issue #98 fixed.
* Issue #99 fixed.
* Issue #100 fixed.
* Issue #101 fixed.
* Issue #107 fixed.
* Issue #108 fixed.
* Issue #109 fixed.

## 0.2.5-alpha (June 28, 2024)
BUGFIXES:
* resource meraki_networks_appliance_firewall_l3_firewall_rules this_rule - rules issues for empty ruleset and syslog_default_rule #76.
* resource meraki_networks_appliance_firewall_l7_firewall_rules this_rule - fails to apply with empty ruleset #79.
* error when attempting to import vlan into resource meraki_networks_appliance_vlans this_vlan #81.
* resource meraki_organizations_appliance_vpn_third_party_vpnpeers fails to create #80.
* error on meraki_networks_appliance_content_filtering │ Path: blocked_url_categories_rs #82.
* resource meraki_networks_vlan_profiles causes error when attempting to create #73.

## 0.2.4-alpha (June 13, 2024)
BUGFIXES:
* Resource meraki_networks_sensor_alerts_profiles error with operation CreateNetworkSensorAlertsProfile #31.
* resource meraki_networks_vlan_profiles causes error when attempting to create #73.
* resource meraki_networks_webhooks_http_servers marked as tainted and results in an error after creating #72.
* resource meraki_networks_alerts_settings causes a panic on create when muting is set #71.
* resource meraki_networks_alerts_settings causes a panic when attempting to import #70.
* resource meraki_networks_wireless_ssids this_ssid - 8021x-radius auth tfprotov6 plug crashing on apply #64.
* resource meraki_networks_appliance_firewall_one_to_many_nat_rules this - failure on apply with empty ruleset to ensure no rules applied. #63
* resource meraki_networks_appliance_firewall_inbound_firewall_rules this - Error out due to default rule not meshing with configured rule block. #62
* Resource meraki_networks_sensor_alerts_profiles error with operation CreateNetworkSensorAlertsProfile #31.
* meraki_networks_wireless_ssids_firewall_l3_firewall_rules rule attribute #28.

## 0.2.3-alpha (May 31, 2024)
BUGFIXES:
* Removing unnecessary attributes from `meraki_devices_appliance_vmx_authentication`.
* Attribute `group_policy_id` turns to an only `Computed` parameter.
* Attribute `switch_stack_id` is now a `Computed` attribute too, to avoid mismatches between state and response.

## 0.2.2-alpha (May 29, 2024)
BUGFIXES:
* Manage `null` values for array into objects in `meraki_devices_appliance_uplinks_settings_resource`.
* Manage `null` values for array into objects in `meraki_networks_appliance_security_intrusion_resource`.
* Include exception to expects `400` code instead of `404` in `meraki_networks_group_policies_resource`.
* Attribute `password` is now just am `Optional` parameter, not `Computed` in `meraki_networks_settings_resource`.
* Avoid unexpected or non-existent changes in `meraki_networks_switch_access_policies_resource`.
* Mapping ID post creation in `meraki_networks_switch_stacks_resource`.
* Attribute `default_vlan_id` is now just am `Optional` parameter, not `Computed` in `meraki_networks_wireless_ssids_resource`.
* Attribute `updatestrategy` included in `meraki_networks_wireless_settings_resource`.
* Managing `null` interfaces, to avoid panics in `utils` at `changeUnknowns` function.

## 0.2.1-alpha (April 16, 2024)
BUGFIXES:
* NetworksWirelessSSIDsResource attributte marked as `Computed` removed cause were only `Optional` params.

## 0.2.0-alpha (April 16, 2024)
FEATURES:
* Provider supports v1.44.1 of Meraki Dashboard API.
* **New Resource** `meraki_administered_licensing_subscription_subscriptions_bind`
* **New Resource** `meraki_administered_licensing_subscription_subscriptions_claim`
* **New Resource** `meraki_administered_licensing_subscription_subscriptions_claim_key_validate`
* **New Resource** `meraki_devices_appliance_radio_settings`
* **New Resource** `meraki_devices_live_tools_arp_table`
* **New Resource** `meraki_devices_live_tools_cable_test`
* **New Resource** `meraki_devices_live_tools_ping`
* **New Resource** `meraki_devices_live_tools_throughput_test`
* **New Resource** `meraki_devices_live_tools_ping_device`
* **New Resource** `meraki_devices_live_tools_wake_on_lan`
* **New Resource** `meraki_devices_wireless_alternate_management_interface_ipv6`
* **New Resource** `meraki_networks_appliance_rf_profiles`
* **New Resource** `meraki_networks_appliance_traffic_shaping_vpn_exclusions`
* **New Resource** `meraki_networks_sm_devices_install_apps`
* **New Resource** `meraki_networks_sm_devices_reboot`
* **New Resource** `meraki_networks_sm_devices_shutdown`
* **New Resource** `meraki_networks_sm_devices_uninstall_apps`
* **New Resource** `meraki_networks_vlan_profiles`
* **New Resource** `meraki_networks_vlan_profiles_assignments_reassign`
* **New Resource** `meraki_networks_wireless_ethernet_ports_profiles`
* **New Resource** `meraki_networks_wireless_ethernet_ports_profiles_assign`
* **New Resource** `meraki_networks_wireless_ethernet_ports_profiles_set_default`
* **New Resource** `meraki_organizations_camera_roles`
* **New Resource** `meraki_organizations_sm_admins_roles`
* **New Resource** `meraki_organizations_sm_sentry_policies_assignments`
* **New Data Source** `meraki_administered_licensing_subscription_entitlements`
* **New Data Source** `meraki_administered_licensing_subscription_subscriptions`
* **New Data Source** `meraki_administered_licensing_subscription_subscriptions_compliance_statuses`
* **New Data Source** `meraki_devices_appliance_radio_settings`
* **New Data Source** `meraki_devices_live_tools_arp_table`
* **New Data Source** `meraki_devices_live_tools_cable_test`
* **New Data Source** `meraki_devices_live_tools_ping`
* **New Data Source** `meraki_devices_live_tools_ping_device`
* **New Data Source** `meraki_devices_live_tools_wake_on_lan`
* **New Data Source** `meraki_networks_appliance_rf_profiles`
* **New Data Source** `meraki_networks_vlan_profiles`
* **New Data Source** `meraki_networks_vlan_profiles_assignments_by_device`
* **New Data Source** `meraki_networks_wireless_ethernet_ports_profiles`
* **New Data Source** `meraki_organizations_appliance_traffic_shaping_vpn_exclusions_by_network`
* **New Data Source** `meraki_organizations_appliance_uplinks_statuses_overview`
* **New Data Source** `meraki_organizations_appliance_uplinks_usage_by_network`
* **New Data Source** `meraki_organizations_camera_boundaries_areas_by_device`
* **New Data Source** `meraki_organizations_camera_boundaries_lines_by_device`
* **New Data Source** `meraki_organizations_camera_detections_history_by_boundary_by_interval`
* **New Data Source** `meraki_organizations_camera_permissions`
* **New Data Source** `meraki_organizations_camera_roles`
* **New Data Source** `meraki_organizations_devices_availabilities_change_history`
* **New Data Source** `meraki_organizations_devices_boots_history`
* **New Data Source** `meraki_organizations_inventory_onboarding_cloud_monitoring_imports_info`
* **New Data Source** `meraki_organizations_sm_admins_roles`
* **New Data Source** `meraki_organizations_sm_sentry_policies_assignments_by_network`
* **New Data Source** `meraki_organizations_summary_top_networks_by_status`
* **New Data Source** `meraki_organizations_webhooks_callbacks_statuses`
* **New Data Source** `meraki_organizations_wireless_devices_channel_utilization_by_device`
* **New Data Source** `meraki_organizations_wireless_devices_channel_utilization_by_network`
* **New Data Source** `meraki_organizations_wireless_devices_channel_utilization_history_by_device_by_interval`
* **New Data Source** `meraki_organizations_wireless_devices_channel_utilization_history_by_network_by_interval`
* **New Data Source** `meraki_organizations_wireless_devices_packet_loss_by_client`
* **New Data Source** `meraki_organizations_ wireless_devices_packet_loss_by_device`
* **New Data Source** `meraki_organizations_wireless_devices_packet_loss_by_network`

BUGFIXES:
* `meraki_networks_appliance_vlans`, `id` parameter now is required for create context.
* `meraki_devices_management_interface` resource seems to be broken #12. [FIXED]

## 0.1.0-alpha (March 26, 2024)
FEATURES:
* **New Data Source:** `meraki_administered_identities_me`
* **New Data Source:** `meraki_devices_appliance_performance`
* **New Data Source:** `meraki_devices_appliance_uplinks_settings`
* **New Data Source:** `meraki_devices_camera_analytics_live`
* **New Data Source:** `meraki_devices_camera_custom_analytics`
* **New Data Source:** `meraki_devices_camera_quality_and_retention`
* **New Data Source:** `meraki_devices_camera_sense`
* **New Data Source:** `meraki_devices_camera_video_link`
* **New Data Source:** `meraki_devices_camera_video_settings`
* **New Data Source:** `meraki_devices_camera_wireless_profiles`
* **New Data Source:** `meraki_devices_cellular_gateway_lan`
* **New Data Source:** `meraki_devices_cellular_gateway_port_forwarding_rules`
* **New Data Source:** `meraki_devices_cellular_sims`
* **New Data Source:** `meraki_devices_live_tools_ping_device`
* **New Data Source:** `meraki_devices_live_tools_ping`
* **New Data Source:** `meraki_devices_lldp_cdp`
* **New Data Source:** `meraki_devices_management_interface`
* **New Data Source:** `meraki_devices_sensor_relationships`
* **New Data Source:** `meraki_devices_switch_ports_statuses`
* **New Data Source:** `meraki_devices_switch_ports`
* **New Data Source:** `meraki_devices_switch_routing_interfaces_dhcp`
* **New Data Source:** `meraki_devices_switch_routing_interfaces`
* **New Data Source:** `meraki_devices_switch_routing_static_routes`
* **New Data Source:** `meraki_devices_switch_warm_spare`
* **New Data Source:** `meraki_devices_wireless_bluetooth_settings`
* **New Data Source:** `meraki_devices_wireless_connection_stats`
* **New Data Source:** `meraki_devices_wireless_latency_stats`
* **New Data Source:** `meraki_devices_wireless_radio_settings`
* **New Data Source:** `meraki_devices_wireless_status`
* **New Data Source:** `meraki_devices`
* **New Data Source:** `meraki_networks_alerts_history`
* **New Data Source:** `meraki_networks_alerts_settings`
* **New Data Source:** `meraki_networks_appliance_connectivity_monitoring_destinations`
* **New Data Source:** `meraki_networks_appliance_content_filtering_categories`
* **New Data Source:** `meraki_networks_appliance_content_filtering`
* **New Data Source:** `meraki_networks_appliance_firewall_cellular_firewall_rules`
* **New Data Source:** `meraki_networks_appliance_firewall_firewalled_services`
* **New Data Source:** `meraki_networks_appliance_firewall_inbound_firewall_rules`
* **New Data Source:** `meraki_networks_appliance_firewall_l3_firewall_rules`
* **New Data Source:** `meraki_networks_appliance_firewall_l7_firewall_rules_application_categories`
* **New Data Source:** `meraki_networks_appliance_firewall_l7_firewall_rules`
* **New Data Source:** `meraki_networks_appliance_firewall_one_to_many_nat_rules`
* **New Data Source:** `meraki_networks_appliance_firewall_one_to_one_nat_rules`
* **New Data Source:** `meraki_networks_appliance_firewall_port_forwarding_rules`
* **New Data Source:** `meraki_networks_appliance_firewall_settings`
* **New Data Source:** `meraki_networks_appliance_ports`
* **New Data Source:** `meraki_networks_appliance_prefixes_delegated_statics`
* **New Data Source:** `meraki_networks_appliance_security_intrusion`
* **New Data Source:** `meraki_networks_appliance_security_malware`
* **New Data Source:** `meraki_networks_appliance_settings`
* **New Data Source:** `meraki_networks_appliance_single_lan`
* **New Data Source:** `meraki_networks_appliance_ssids`
* **New Data Source:** `meraki_networks_appliance_traffic_shaping_rules`
* **New Data Source:** `meraki_networks_appliance_traffic_shaping_uplink_bandwidth`
* **New Data Source:** `meraki_networks_appliance_traffic_shaping_uplink_selection`
* **New Data Source:** `meraki_networks_appliance_traffic_shaping`
* **New Data Source:** `meraki_networks_appliance_vlans_settings`
* **New Data Source:** `meraki_networks_appliance_vlans`
* **New Data Source:** `meraki_networks_appliance_vpn_bgp`
* **New Data Source:** `meraki_networks_appliance_vpn_site_to_site_vpn`
* **New Data Source:** `meraki_networks_appliance_warm_spare`
* **New Data Source:** `meraki_networks_bluetooth_clients`
* **New Data Source:** `meraki_networks_camera_quality_retention_profiles`
* **New Data Source:** `meraki_networks_camera_wireless_profiles`
* **New Data Source:** `meraki_networks_cellular_gateway_connectivity_monitoring_destinations`
* **New Data Source:** `meraki_networks_cellular_gateway_dhcp`
* **New Data Source:** `meraki_networks_cellular_gateway_subnet_pool`
* **New Data Source:** `meraki_networks_cellular_gateway_uplink`
* **New Data Source:** `meraki_networks_clients_overview`
* **New Data Source:** `meraki_networks_clients_policy`
* **New Data Source:** `meraki_networks_clients_splash_authorization_status`
* **New Data Source:** `meraki_networks_clients`
* **New Data Source:** `meraki_networks_events_event_types`
* **New Data Source:** `meraki_networks_events`
* **New Data Source:** `meraki_networks_firmware_upgrades_staged_events`
* **New Data Source:** `meraki_networks_firmware_upgrades_staged_groups`
* **New Data Source:** `meraki_networks_firmware_upgrades_staged_stages`
* **New Data Source:** `meraki_networks_firmware_upgrades`
* **New Data Source:** `meraki_networks_floor_plans`
* **New Data Source:** `meraki_networks_group_policies`
* **New Data Source:** `meraki_networks_health_alerts`
* **New Data Source:** `meraki_networks_insight_applications_health_by_time`
* **New Data Source:** `meraki_networks_meraki_auth_users`
* **New Data Source:** `meraki_networks_netflow`
* **New Data Source:** `meraki_networks_pii_pii_keys`
* **New Data Source:** `meraki_networks_pii_requests`
* **New Data Source:** `meraki_networks_pii_sm_devices_for_key`
* **New Data Source:** `meraki_networks_pii_sm_owners_for_key`
* **New Data Source:** `meraki_networks_policies_by_client`
* **New Data Source:** `meraki_networks_sensor_alerts_current_overview_by_metric`
* **New Data Source:** `meraki_networks_sensor_alerts_overview_by_metric`
* **New Data Source:** `meraki_networks_sensor_alerts_profiles`
* **New Data Source:** `meraki_networks_sensor_mqtt_brokers`
* **New Data Source:** `meraki_networks_sensor_relationships`
* **New Data Source:** `meraki_networks_settings`
* **New Data Source:** `meraki_networks_sm_bypass_activation_lock_attempts`
* **New Data Source:** `meraki_networks_sm_devices_cellular_usage_history`
* **New Data Source:** `meraki_networks_sm_devices_certs`
* **New Data Source:** `meraki_networks_sm_devices_connectivity`
* **New Data Source:** `meraki_networks_sm_devices_desktop_logs`
* **New Data Source:** `meraki_networks_sm_devices_device_command_logs`
* **New Data Source:** `meraki_networks_sm_devices_device_profiles`
* **New Data Source:** `meraki_networks_sm_devices_network_adapters`
* **New Data Source:** `meraki_networks_sm_devices_performance_history`
* **New Data Source:** `meraki_networks_sm_devices_security_centers`
* **New Data Source:** `meraki_networks_sm_devices_wlan_lists`
* **New Data Source:** `meraki_networks_sm_devices`
* **New Data Source:** `meraki_networks_sm_profiles`
* **New Data Source:** `meraki_networks_sm_target_groups`
* **New Data Source:** `meraki_networks_sm_trusted_access_configs`
* **New Data Source:** `meraki_networks_sm_user_access_devices`
* **New Data Source:** `meraki_networks_sm_users_device_profiles`
* **New Data Source:** `meraki_networks_sm_users_softwares`
* **New Data Source:** `meraki_networks_sm_users`
* **New Data Source:** `meraki_networks_snmp`
* **New Data Source:** `meraki_networks_switch_access_control_lists`
* **New Data Source:** `meraki_networks_switch_access_policies`
* **New Data Source:** `meraki_networks_switch_alternate_management_interface`
* **New Data Source:** `meraki_networks_switch_dhcp_server_policy_arp_inspection_trusted_servers`
* **New Data Source:** `meraki_networks_switch_dhcp_server_policy_arp_inspection_warnings_by_device`
* **New Data Source:** `meraki_networks_switch_dhcp_server_policy`
* **New Data Source:** `meraki_networks_switch_dhcp_v4_servers_seen`
* **New Data Source:** `meraki_networks_switch_dscp_to_cos_mappings`
* **New Data Source:** `meraki_networks_switch_link_aggregations`
* **New Data Source:** `meraki_networks_switch_mtu`
* **New Data Source:** `meraki_networks_switch_port_schedules`
* **New Data Source:** `meraki_networks_switch_qos_rules_order`
* **New Data Source:** `meraki_networks_switch_routing_multicast`
* **New Data Source:** `meraki_networks_switch_routing_ospf`
* **New Data Source:** `meraki_networks_switch_settings`
* **New Data Source:** `meraki_networks_switch_stacks_routing_interfaces_dhcp`
* **New Data Source:** `meraki_networks_switch_stacks_routing_interfaces`
* **New Data Source:** `meraki_networks_switch_stacks_routing_static_routes`
* **New Data Source:** `meraki_networks_switch_stacks`
* **New Data Source:** `meraki_networks_switch_storm_control`
* **New Data Source:** `meraki_networks_switch_stp`
* **New Data Source:** `meraki_networks_syslog_servers`
* **New Data Source:** `meraki_networks_topology_link_layer`
* **New Data Source:** `meraki_networks_traffic_analysis`
* **New Data Source:** `meraki_networks_traffic_shaping_application_categories`
* **New Data Source:** `meraki_networks_traffic_shaping_dscp_tagging_options`
* **New Data Source:** `meraki_networks_webhooks_http_servers`
* **New Data Source:** `meraki_networks_webhooks_payload_templates`
* **New Data Source:** `meraki_networks_webhooks_webhook_tests`
* **New Data Source:** `meraki_networks_wireless_alternate_management_interface`
* **New Data Source:** `meraki_networks_wireless_billing`
* **New Data Source:** `meraki_networks_wireless_bluetooth_settings`
* **New Data Source:** `meraki_networks_wireless_channel_utilization_history`
* **New Data Source:** `meraki_networks_wireless_client_count_history`
* **New Data Source:** `meraki_networks_wireless_clients_connection_stats`
* **New Data Source:** `meraki_networks_wireless_clients_latency_stats`
* **New Data Source:** `meraki_networks_wireless_connection_stats`
* **New Data Source:** `meraki_networks_wireless_data_rate_history`
* **New Data Source:** `meraki_networks_wireless_devices_connection_stats`
* **New Data Source:** `meraki_networks_wireless_failed_connections`
* **New Data Source:** `meraki_networks_wireless_latency_history`
* **New Data Source:** `meraki_networks_wireless_latency_stats`
* **New Data Source:** `meraki_networks_wireless_mesh_statuses`
* **New Data Source:** `meraki_networks_wireless_rf_profiles`
* **New Data Source:** `meraki_networks_wireless_settings`
* **New Data Source:** `meraki_networks_wireless_signal_quality_history`
* **New Data Source:** `meraki_networks_wireless_ssids_bonjour_forwarding`
* **New Data Source:** `meraki_networks_wireless_ssids_device_type_group_policies`
* **New Data Source:** `meraki_networks_wireless_ssids_eap_override`
* **New Data Source:** `meraki_networks_wireless_ssids_firewall_l3_firewall_rules`
* **New Data Source:** `meraki_networks_wireless_ssids_firewall_l7_firewall_rules`
* **New Data Source:** `meraki_networks_wireless_ssids_hotspot20`
* **New Data Source:** `meraki_networks_wireless_ssids_identity_psks`
* **New Data Source:** `meraki_networks_wireless_ssids_schedules`
* **New Data Source:** `meraki_networks_wireless_ssids_splash_settings`
* **New Data Source:** `meraki_networks_wireless_ssids_traffic_shaping_rules`
* **New Data Source:** `meraki_networks_wireless_ssids_vpn`
* **New Data Source:** `meraki_networks_wireless_ssids`
* **New Data Source:** `meraki_networks_wireless_usage_history`
* **New Data Source:** `meraki_networks`
* **New Data Source:** `meraki_organizations_action_batches`
* **New Data Source:** `meraki_organizations_adaptive_policy_acls`
* **New Data Source:** `meraki_organizations_adaptive_policy_groups`
* **New Data Source:** `meraki_organizations_adaptive_policy_overview`
* **New Data Source:** `meraki_organizations_adaptive_policy_policies`
* **New Data Source:** `meraki_organizations_adaptive_policy_settings`
* **New Data Source:** `meraki_organizations_admins`
* **New Data Source:** `meraki_organizations_alerts_profiles`
* **New Data Source:** `meraki_organizations_api_requests_overview_response_codes_by_interval`
* **New Data Source:** `meraki_organizations_api_requests_overview`
* **New Data Source:** `meraki_organizations_api_requests`
* **New Data Source:** `meraki_organizations_appliance_security_intrusion`
* **New Data Source:** `meraki_organizations_appliance_vpn_third_party_vpnpeers`
* **New Data Source:** `meraki_organizations_appliance_vpn_vpn_firewall_rules`
* **New Data Source:** `meraki_organizations_branding_policies_priorities`
* **New Data Source:** `meraki_organizations_branding_policies`
* **New Data Source:** `meraki_organizations_camera_custom_analytics_artifacts`
* **New Data Source:** `meraki_organizations_cellular_gateway_uplink_statuses`
* **New Data Source:** `meraki_organizations_clients_bandwidth_usage_history`
* **New Data Source:** `meraki_organizations_clients_overview`
* **New Data Source:** `meraki_organizations_clients_search`
* **New Data Source:** `meraki_organizations_config_templates_switch_profiles_ports`
* **New Data Source:** `meraki_organizations_config_templates_switch_profiles`
* **New Data Source:** `meraki_organizations_config_templates`
* **New Data Source:** `meraki_organizations_devices_availabilities`
* **New Data Source:** `meraki_organizations_devices_power_modules_statuses_by_device`
* **New Data Source:** `meraki_organizations_devices_provisioning_statuses`
* **New Data Source:** `meraki_organizations_devices_statuses_overview`
* **New Data Source:** `meraki_organizations_devices_statuses`
* **New Data Source:** `meraki_organizations_devices_uplinks_addresses_by_device`
* **New Data Source:** `meraki_organizations_devices_uplinks_loss_and_latency`
* **New Data Source:** `meraki_organizations_devices`
* **New Data Source:** `meraki_organizations_early_access_features_opt_ins`
* **New Data Source:** `meraki_organizations_early_access_features`
* **New Data Source:** `meraki_organizations_firmware_upgrades_by_device`
* **New Data Source:** `meraki_organizations_firmware_upgrades`
* **New Data Source:** `meraki_organizations_insight_applications`
* **New Data Source:** `meraki_organizations_insight_monitored_media_servers`
* **New Data Source:** `meraki_organizations_inventory_devices`
* **New Data Source:** `meraki_organizations_inventory_onboarding_cloud_monitoring_imports`
* **New Data Source:** `meraki_organizations_inventory_onboarding_cloud_monitoring_networks`
* **New Data Source:** `meraki_organizations_licenses_overview`
* **New Data Source:** `meraki_organizations_licenses`
* **New Data Source:** `meraki_organizations_licensing_coterm_licenses`
* **New Data Source:** `meraki_organizations_login_security`
* **New Data Source:** `meraki_organizations_openapi_spec`
* **New Data Source:** `meraki_organizations_policy_objects_groups`
* **New Data Source:** `meraki_organizations_policy_objects`
* **New Data Source:** `meraki_organizations_saml_idps`
* **New Data Source:** `meraki_organizations_saml_roles`
* **New Data Source:** `meraki_organizations_saml`
* **New Data Source:** `meraki_organizations_sensor_readings_history`
* **New Data Source:** `meraki_organizations_sensor_readings_latest`
* **New Data Source:** `meraki_organizations_sm_apns_cert`
* **New Data Source:** `meraki_organizations_sm_vpp_accounts`
* **New Data Source:** `meraki_organizations_snmp`
* **New Data Source:** `meraki_organizations_summary_top_appliances_by_utilization`
* **New Data Source:** `meraki_organizations_summary_top_clients_by_usage`
* **New Data Source:** `meraki_organizations_summary_top_clients_manufacturers_by_usage`
* **New Data Source:** `meraki_organizations_summary_top_devices_by_usage`
* **New Data Source:** `meraki_organizations_summary_top_devices_models_by_usage`
* **New Data Source:** `meraki_organizations_summary_top_ssids_by_usage`
* **New Data Source:** `meraki_organizations_summary_top_switches_by_energy_usage`
* **New Data Source:** `meraki_organizations_switch_ports_by_switch`
* **New Data Source:** `meraki_organizations_uplinks_statuses`
* **New Data Source:** `meraki_organizations_webhooks_logs`
* **New Data Source:** `meraki_organizations`
* **New Resource:** `meraki_devices_appliance_uplinks_settings`
* **New Resource:** `meraki_devices_appliance_vmx_authentication_token`
* **New Resource:** `meraki_devices_blink_leds`
* **New Resource:** `meraki_devices_camera_custom_analytics`
* **New Resource:** `meraki_devices_camera_generate_snapshot`
* **New Resource:** `meraki_devices_camera_quality_and_retention`
* **New Resource:** `meraki_devices_camera_sense`
* **New Resource:** `meraki_devices_camera_video_settings`
* **New Resource:** `meraki_devices_camera_wireless_profiles`
* **New Resource:** `meraki_devices_cellular_gateway_lan`
* **New Resource:** `meraki_devices_cellular_gateway_port_forwarding_rules`
* **New Resource:** `meraki_devices_cellular_sims`
* **New Resource:** `meraki_devices_live_tools_ping_device`
* **New Resource:** `meraki_devices_live_tools_ping`
* **New Resource:** `meraki_devices_management_interface`
* **New Resource:** `meraki_devices_sensor_relationships`
* **New Resource:** `meraki_devices_switch_ports_cycle`
* **New Resource:** `meraki_devices_switch_ports`
* **New Resource:** `meraki_devices_switch_routing_interfaces_dhcp`
* **New Resource:** `meraki_devices_switch_routing_interfaces`
* **New Resource:** `meraki_devices_switch_routing_static_routes`
* **New Resource:** `meraki_devices_switch_warm_spare`
* **New Resource:** `meraki_devices_wireless_bluetooth_settings`
* **New Resource:** `meraki_devices_wireless_radio_settings`
* **New Resource:** `meraki_devices`
* **New Resource:** `meraki_networks_alerts_settings`
* **New Resource:** `meraki_networks_appliance_connectivity_monitoring_destinations`
* **New Resource:** `meraki_networks_appliance_content_filtering`
* **New Resource:** `meraki_networks_appliance_firewall_cellular_firewall_rules`
* **New Resource:** `meraki_networks_appliance_firewall_firewalled_services`
* **New Resource:** `meraki_networks_appliance_firewall_inbound_firewall_rules`
* **New Resource:** `meraki_networks_appliance_firewall_l3_firewall_rules`
* **New Resource:** `meraki_networks_appliance_firewall_l7_firewall_rules`
* **New Resource:** `meraki_networks_appliance_firewall_one_to_many_nat_rules`
* **New Resource:** `meraki_networks_appliance_firewall_one_to_one_nat_rules`
* **New Resource:** `meraki_networks_appliance_firewall_port_forwarding_rules`
* **New Resource:** `meraki_networks_appliance_firewall_settings`
* **New Resource:** `meraki_networks_appliance_ports`
* **New Resource:** `meraki_networks_appliance_prefixes_delegated_statics`
* **New Resource:** `meraki_networks_appliance_security_intrusion`
* **New Resource:** `meraki_networks_appliance_security_malware`
* **New Resource:** `meraki_networks_appliance_settings`
* **New Resource:** `meraki_networks_appliance_single_lan`
* **New Resource:** `meraki_networks_appliance_ssids`
* **New Resource:** `meraki_networks_appliance_traffic_shaping_custom_performance_classes`
* **New Resource:** `meraki_networks_appliance_traffic_shaping_rules`
* **New Resource:** `meraki_networks_appliance_traffic_shaping_uplink_bandwidth`
* **New Resource:** `meraki_networks_appliance_traffic_shaping_uplink_selection`
* **New Resource:** `meraki_networks_appliance_traffic_shaping`
* **New Resource:** `meraki_networks_appliance_vlans_settings`
* **New Resource:** `meraki_networks_appliance_vlans`
* **New Resource:** `meraki_networks_appliance_vpn_bgp`
* **New Resource:** `meraki_networks_appliance_vpn_site_to_site_vpn`
* **New Resource:** `meraki_networks_appliance_warm_spare_swap`
* **New Resource:** `meraki_networks_appliance_warm_spare`
* **New Resource:** `meraki_networks_bind`
* **New Resource:** `meraki_networks_camera_quality_retention_profiles`
* **New Resource:** `meraki_networks_camera_wireless_profiles`
* **New Resource:** `meraki_networks_cellular_gateway_connectivity_monitoring_destinations`
* **New Resource:** `meraki_networks_cellular_gateway_dhcp`
* **New Resource:** `meraki_networks_cellular_gateway_subnet_pool`
* **New Resource:** `meraki_networks_cellular_gateway_uplink`
* **New Resource:** `meraki_networks_clients_policy`
* **New Resource:** `meraki_networks_clients_provision`
* **New Resource:** `meraki_networks_clients_splash_authorization_status`
* **New Resource:** `meraki_networks_devices_claim_vmx`
* **New Resource:** `meraki_networks_devices_claim`
* **New Resource:** `meraki_networks_devices_remove`
* **New Resource:** `meraki_networks_firmware_upgrades_rollbacks`
* **New Resource:** `meraki_networks_firmware_upgrades_staged_events_defer`
* **New Resource:** `meraki_networks_firmware_upgrades_staged_events_rollbacks`
* **New Resource:** `meraki_networks_firmware_upgrades_staged_events`
* **New Resource:** `meraki_networks_firmware_upgrades_staged_groups`
* **New Resource:** `meraki_networks_firmware_upgrades_staged_stages`
* **New Resource:** `meraki_networks_firmware_upgrades`
* **New Resource:** `meraki_networks_floor_plans`
* **New Resource:** `meraki_networks_group_policies`
* **New Resource:** `meraki_networks_meraki_auth_users`
* **New Resource:** `meraki_networks_mqtt_brokers`
* **New Resource:** `meraki_networks_netflow`
* **New Resource:** `meraki_networks_pii_requests_delete`
* **New Resource:** `meraki_networks_sensor_alerts_profiles`
* **New Resource:** `meraki_networks_sensor_mqtt_brokers`
* **New Resource:** `meraki_networks_settings`
* **New Resource:** `meraki_networks_sm_devices_checkin`
* **New Resource:** `meraki_networks_sm_devices_fields`
* **New Resource:** `meraki_networks_sm_devices_lock`
* **New Resource:** `meraki_networks_sm_devices_modify_tags`
* **New Resource:** `meraki_networks_sm_devices_move`
* **New Resource:** `meraki_networks_sm_devices_refresh_details`
* **New Resource:** `meraki_networks_sm_devices_unenroll`
* **New Resource:** `meraki_networks_sm_devices_wipe`
* **New Resource:** `meraki_networks_sm_target_groups`
* **New Resource:** `meraki_networks_sm_user_access_devices_delete`
* **New Resource:** `meraki_networks_snmp`
* **New Resource:** `meraki_networks_split`
* **New Resource:** `meraki_networks_switch_access_control_lists`
* **New Resource:** `meraki_networks_switch_access_policies`
* **New Resource:** `meraki_networks_switch_alternate_management_interface`
* **New Resource:** `meraki_networks_switch_dhcp_server_policy_arp_inspection_trusted_servers`
* **New Resource:** `meraki_networks_switch_dhcp_server_policy`
* **New Resource:** `meraki_networks_switch_dscp_to_cos_mappings`
* **New Resource:** `meraki_networks_switch_link_aggregations`
* **New Resource:** `meraki_networks_switch_mtu`
* **New Resource:** `meraki_networks_switch_port_schedules`
* **New Resource:** `meraki_networks_switch_qos_rules_order`
* **New Resource:** `meraki_networks_switch_routing_multicast_rendezvous_points`
* **New Resource:** `meraki_networks_switch_routing_multicast`
* **New Resource:** `meraki_networks_switch_routing_ospf`
* **New Resource:** `meraki_networks_switch_settings`
* **New Resource:** `meraki_networks_switch_stacks_add`
* **New Resource:** `meraki_networks_switch_stacks_remove`
* **New Resource:** `meraki_networks_switch_stacks_routing_interfaces_dhcp`
* **New Resource:** `meraki_networks_switch_stacks_routing_interfaces`
* **New Resource:** `meraki_networks_switch_stacks_routing_static_routes`
* **New Resource:** `meraki_networks_switch_stacks`
* **New Resource:** `meraki_networks_switch_storm_control`
* **New Resource:** `meraki_networks_switch_stp`
* **New Resource:** `meraki_networks_syslog_servers`
* **New Resource:** `meraki_networks_traffic_analysis`
* **New Resource:** `meraki_networks_unbind`
* **New Resource:** `meraki_networks_webhooks_http_servers`
* **New Resource:** `meraki_networks_webhooks_payload_templates`
* **New Resource:** `meraki_networks_wireless_alternate_management_interface`
* **New Resource:** `meraki_networks_wireless_billing`
* **New Resource:** `meraki_networks_wireless_bluetooth_settings`
* **New Resource:** `meraki_networks_wireless_rf_profiles`
* **New Resource:** `meraki_networks_wireless_settings`
* **New Resource:** `meraki_networks_wireless_ssids_bonjour_forwarding`
* **New Resource:** `meraki_networks_wireless_ssids_device_type_group_policies`
* **New Resource:** `meraki_networks_wireless_ssids_eap_override`
* **New Resource:** `meraki_networks_wireless_ssids_firewall_l3_firewall_rules`
* **New Resource:** `meraki_networks_wireless_ssids_firewall_l7_firewall_rules`
* **New Resource:** `meraki_networks_wireless_ssids_hotspot20`
* **New Resource:** `meraki_networks_wireless_ssids_identity_psks`
* **New Resource:** `meraki_networks_wireless_ssids_schedules`
* **New Resource:** `meraki_networks_wireless_ssids_splash_settings`
* **New Resource:** `meraki_networks_wireless_ssids_traffic_shaping_rules`
* **New Resource:** `meraki_networks_wireless_ssids_vpn`
* **New Resource:** `meraki_networks_wireless_ssids`
* **New Resource:** `meraki_networks`
* **New Resource:** `meraki_organizations_action_batches`
* **New Resource:** `meraki_organizations_adaptive_policy_acls`
* **New Resource:** `meraki_organizations_adaptive_policy_groups`
* **New Resource:** `meraki_organizations_adaptive_policy_policies`
* **New Resource:** `meraki_organizations_adaptive_policy_settings`
* **New Resource:** `meraki_organizations_admins`
* **New Resource:** `meraki_organizations_alerts_profiles`
* **New Resource:** `meraki_organizations_appliance_security_intrusion`
* **New Resource:** `meraki_organizations_appliance_vpn_third_party_vpnpeers`
* **New Resource:** `meraki_organizations_appliance_vpn_vpn_firewall_rules`
* **New Resource:** `meraki_organizations_branding_policies_priorities`
* **New Resource:** `meraki_organizations_branding_policies`
* **New Resource:** `meraki_organizations_camera_custom_analytics_artifacts`
* **New Resource:** `meraki_organizations_claim`
* **New Resource:** `meraki_organizations_clone`
* **New Resource:** `meraki_organizations_config_templates_switch_profiles_ports`
* **New Resource:** `meraki_organizations_config_templates`
* **New Resource:** `meraki_organizations_early_access_features_opt_ins`
* **New Resource:** `meraki_organizations_insight_monitored_media_servers`
* **New Resource:** `meraki_organizations_inventory_claim`
* **New Resource:** `meraki_organizations_inventory_onboarding_cloud_monitoring_export_events`
* **New Resource:** `meraki_organizations_inventory_onboarding_cloud_monitoring_prepare`
* **New Resource:** `meraki_organizations_inventory_release`
* **New Resource:** `meraki_organizations_licenses_assign_seats`
* **New Resource:** `meraki_organizations_licenses_move_seats`
* **New Resource:** `meraki_organizations_licenses_move`
* **New Resource:** `meraki_organizations_licenses_renew_seats`
* **New Resource:** `meraki_organizations_licenses`
* **New Resource:** `meraki_organizations_licensing_coterm_licenses_move`
* **New Resource:** `meraki_organizations_login_security`
* **New Resource:** `meraki_organizations_networks_combine`
* **New Resource:** `meraki_organizations_policy_objects_groups`
* **New Resource:** `meraki_organizations_policy_objects`
* **New Resource:** `meraki_organizations_saml_roles`
* **New Resource:** `meraki_organizations_saml`
* **New Resource:** `meraki_organizations_snmp`
* **New Resource:** `meraki_organizations_switch_devices_clone`
* **New Resource:** `meraki_organizations_users`
* **New Resource:** `meraki_organizations`