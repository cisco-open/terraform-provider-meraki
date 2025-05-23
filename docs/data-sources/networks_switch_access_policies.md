---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_switch_access_policies Data Source - terraform-provider-meraki"
subcategory: "switch"
description: |-
  
---

# meraki_networks_switch_access_policies (Data Source)



## Example Usage

```terraform
data "meraki_networks_switch_access_policies" "example" {

  network_id = "string"
}

output "meraki_networks_switch_access_policies_example" {
  value = data.meraki_networks_switch_access_policies.example.items
}

data "meraki_networks_switch_access_policies" "example" {

  network_id = "string"
}

output "meraki_networks_switch_access_policies_example" {
  value = data.meraki_networks_switch_access_policies.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `access_policy_number` (String) accessPolicyNumber path parameter. Access policy number
- `network_id` (String) networkId path parameter. Network ID

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))
- `items` (Attributes List) Array of ResponseSwitchGetNetworkSwitchAccessPolicies (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `access_policy_type` (String) Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
- `counts` (Attributes) Counts associated with the access policy (see [below for nested schema](#nestedatt--item--counts))
- `dot1x` (Attributes) 802.1x Settings (see [below for nested schema](#nestedatt--item--dot1x))
- `guest_port_bouncing` (Boolean) If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
- `guest_vlan_id` (Number) ID for the guest VLAN allow unauthorized devices access to limited network resources
- `host_mode` (String) Choose the Host Mode for the access policy.
- `increase_access_speed` (Boolean) Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is 'Hybrid Authentication.
- `name` (String) Name of the access policy
- `radius` (Attributes) Object for RADIUS Settings (see [below for nested schema](#nestedatt--item--radius))
- `radius_accounting_enabled` (Boolean) Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients
- `radius_accounting_servers` (Attributes Set) List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access (see [below for nested schema](#nestedatt--item--radius_accounting_servers))
- `radius_coa_support_enabled` (Boolean) Change of authentication for RADIUS re-authentication and disconnection
- `radius_group_attribute` (String) Acceptable values are **""** for None, or **"11"** for Group Policies ACL
- `radius_servers` (Attributes Set) List of RADIUS servers to require connecting devices to authenticate against before granting network access (see [below for nested schema](#nestedatt--item--radius_servers))
- `radius_testing_enabled` (Boolean) If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
- `url_redirect_walled_garden_enabled` (Boolean) Enable to restrict access for clients to a response_objectific set of IP addresses or hostnames prior to authentication
- `url_redirect_walled_garden_ranges` (List of String) IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
- `voice_vlan_clients` (Boolean) CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.

<a id="nestedatt--item--counts"></a>
### Nested Schema for `item.counts`

Read-Only:

- `ports` (Attributes) Counts associated with ports (see [below for nested schema](#nestedatt--item--counts--ports))

<a id="nestedatt--item--counts--ports"></a>
### Nested Schema for `item.counts.ports`

Read-Only:

- `with_this_policy` (Number) Number of ports in the network with this policy. For template networks, this is the number of template ports (not child ports) with this policy.



<a id="nestedatt--item--dot1x"></a>
### Nested Schema for `item.dot1x`

Read-Only:

- `control_direction` (String) Supports either 'both' or 'inbound'. Set to 'inbound' to allow unauthorized egress on the switchport. Set to 'both' to control both traffic directions with authorization. Defaults to 'both'


<a id="nestedatt--item--radius"></a>
### Nested Schema for `item.radius`

Read-Only:

- `cache` (Attributes) Object for RADIUS Cache Settings (see [below for nested schema](#nestedatt--item--radius--cache))
- `critical_auth` (Attributes) Critical auth settings for when authentication is rejected by the RADIUS server (see [below for nested schema](#nestedatt--item--radius--critical_auth))
- `failed_auth_vlan_id` (Number) VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
- `re_authentication_interval` (Number) Re-authentication period in seconds. Will be null if hostMode is Multi-Auth

<a id="nestedatt--item--radius--cache"></a>
### Nested Schema for `item.radius.cache`

Read-Only:

- `enabled` (Boolean) Enable to cache authorization and authentication responses on the RADIUS server
- `timeout` (Number) If RADIUS caching is enabled, this value dictates how long the cache will remain in the RADIUS server, in hours, to allow network access without authentication


<a id="nestedatt--item--radius--critical_auth"></a>
### Nested Schema for `item.radius.critical_auth`

Read-Only:

- `data_vlan_id` (Number) VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
- `suspend_port_bounce` (Boolean) Enable to suspend port bounce when RADIUS servers are unreachable
- `voice_vlan_id` (Number) VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth



<a id="nestedatt--item--radius_accounting_servers"></a>
### Nested Schema for `item.radius_accounting_servers`

Read-Only:

- `host` (String) Public IP address of the RADIUS accounting server
- `organization_radius_server_id` (String) Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server
- `port` (Number) UDP port that the RADIUS Accounting server listens on for access requests
- `server_id` (String) Unique ID of the RADIUS accounting server


<a id="nestedatt--item--radius_servers"></a>
### Nested Schema for `item.radius_servers`

Read-Only:

- `host` (String) Public IP address of the RADIUS server
- `organization_radius_server_id` (String) Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server
- `port` (Number) UDP port that the RADIUS server listens on for access requests
- `server_id` (String) Unique ID of the RADIUS server



<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `access_policy_type` (String) Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
- `counts` (Attributes) Counts associated with the access policy (see [below for nested schema](#nestedatt--items--counts))
- `dot1x` (Attributes) 802.1x Settings (see [below for nested schema](#nestedatt--items--dot1x))
- `guest_port_bouncing` (Boolean) If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
- `guest_vlan_id` (Number) ID for the guest VLAN allow unauthorized devices access to limited network resources
- `host_mode` (String) Choose the Host Mode for the access policy.
- `increase_access_speed` (Boolean) Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is 'Hybrid Authentication.
- `name` (String) Name of the access policy
- `radius` (Attributes) Object for RADIUS Settings (see [below for nested schema](#nestedatt--items--radius))
- `radius_accounting_enabled` (Boolean) Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients
- `radius_accounting_servers` (Attributes Set) List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access (see [below for nested schema](#nestedatt--items--radius_accounting_servers))
- `radius_coa_support_enabled` (Boolean) Change of authentication for RADIUS re-authentication and disconnection
- `radius_group_attribute` (String) Acceptable values are **""** for None, or **"11"** for Group Policies ACL
- `radius_servers` (Attributes Set) List of RADIUS servers to require connecting devices to authenticate against before granting network access (see [below for nested schema](#nestedatt--items--radius_servers))
- `radius_testing_enabled` (Boolean) If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
- `url_redirect_walled_garden_enabled` (Boolean) Enable to restrict access for clients to a response_objectific set of IP addresses or hostnames prior to authentication
- `url_redirect_walled_garden_ranges` (List of String) IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
- `voice_vlan_clients` (Boolean) CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.

<a id="nestedatt--items--counts"></a>
### Nested Schema for `items.counts`

Read-Only:

- `ports` (Attributes) Counts associated with ports (see [below for nested schema](#nestedatt--items--counts--ports))

<a id="nestedatt--items--counts--ports"></a>
### Nested Schema for `items.counts.ports`

Read-Only:

- `with_this_policy` (Number) Number of ports in the network with this policy. For template networks, this is the number of template ports (not child ports) with this policy.



<a id="nestedatt--items--dot1x"></a>
### Nested Schema for `items.dot1x`

Read-Only:

- `control_direction` (String) Supports either 'both' or 'inbound'. Set to 'inbound' to allow unauthorized egress on the switchport. Set to 'both' to control both traffic directions with authorization. Defaults to 'both'


<a id="nestedatt--items--radius"></a>
### Nested Schema for `items.radius`

Read-Only:

- `cache` (Attributes) Object for RADIUS Cache Settings (see [below for nested schema](#nestedatt--items--radius--cache))
- `critical_auth` (Attributes) Critical auth settings for when authentication is rejected by the RADIUS server (see [below for nested schema](#nestedatt--items--radius--critical_auth))
- `failed_auth_vlan_id` (Number) VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
- `re_authentication_interval` (Number) Re-authentication period in seconds. Will be null if hostMode is Multi-Auth

<a id="nestedatt--items--radius--cache"></a>
### Nested Schema for `items.radius.cache`

Read-Only:

- `enabled` (Boolean) Enable to cache authorization and authentication responses on the RADIUS server
- `timeout` (Number) If RADIUS caching is enabled, this value dictates how long the cache will remain in the RADIUS server, in hours, to allow network access without authentication


<a id="nestedatt--items--radius--critical_auth"></a>
### Nested Schema for `items.radius.critical_auth`

Read-Only:

- `data_vlan_id` (Number) VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
- `suspend_port_bounce` (Boolean) Enable to suspend port bounce when RADIUS servers are unreachable
- `voice_vlan_id` (Number) VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth



<a id="nestedatt--items--radius_accounting_servers"></a>
### Nested Schema for `items.radius_accounting_servers`

Read-Only:

- `host` (String) Public IP address of the RADIUS accounting server
- `organization_radius_server_id` (String) Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server
- `port` (Number) UDP port that the RADIUS Accounting server listens on for access requests
- `server_id` (String) Unique ID of the RADIUS accounting server


<a id="nestedatt--items--radius_servers"></a>
### Nested Schema for `items.radius_servers`

Read-Only:

- `host` (String) Public IP address of the RADIUS server
- `organization_radius_server_id` (String) Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server
- `port` (Number) UDP port that the RADIUS server listens on for access requests
- `server_id` (String) Unique ID of the RADIUS server
