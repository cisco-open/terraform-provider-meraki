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