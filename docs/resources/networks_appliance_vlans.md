---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_appliance_vlans Resource - terraform-provider-meraki"
subcategory: "appliance"
description: |-
  
---

# meraki_networks_appliance_vlans (Resource)



## Example Usage

```terraform
resource "meraki_networks_appliance_vlans" "example" {

  appliance_ip              = "192.168.1.2"
  cidr                      = "192.168.1.0/24"
  dhcp_boot_options_enabled = true
  dhcp_handling             = "Run a DHCP server"
  dhcp_lease_time           = "30 minutes"
  dhcp_options = [{

    code  = "3"
    type  = "text"
    value = "five"
  }]
  group_policy_id = "101"
  id              = "1234"
  ipv6 = {

    enabled = true
    prefix_assignments = [{

      autonomous = false
      origin = {

        interfaces = ["wan0"]
        type       = "internet"
      }
      static_appliance_ip6 = "2001:db8:3c4d:15::1"
      static_prefix        = "2001:db8:3c4d:15::/64"
    }]
  }
  mandatory_dhcp = {

    enabled = true
  }
  mask               = 28
  name               = "My VLAN"
  network_id         = "string"
  subnet             = "192.168.1.0/24"
  template_vlan_type = "same"
}

output "meraki_networks_appliance_vlans_example" {
  value = meraki_networks_appliance_vlans.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Optional

- `appliance_ip` (String) The local IP of the appliance on the VLAN
- `cidr` (String) CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.
- `dhcp_boot_filename` (String) DHCP boot option for boot filename
- `dhcp_boot_next_server` (String) DHCP boot option to direct boot clients to the server to load the boot file from
- `dhcp_boot_options_enabled` (Boolean) Use DHCP boot options specified in other properties
- `dhcp_handling` (String) The appliance's handling of DHCP requests on this VLAN. One of: 'Run a DHCP server', 'Relay DHCP to another server' or 'Do not respond to DHCP requests'
                                  Allowed values: [Do not respond to DHCP requests,Relay DHCP to another server,Run a DHCP server]
- `dhcp_lease_time` (String) The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'
                                  Allowed values: [1 day,1 hour,1 week,12 hours,30 minutes,4 hours]
- `dhcp_options` (Attributes Set) The list of DHCP options that will be included in DHCP responses. Each object in the list should have "code", "type", and "value" properties. (see [below for nested schema](#nestedatt--dhcp_options))
- `dhcp_relay_server_ips` (Set of String) The IPs of the DHCP servers that DHCP requests should be relayed to
- `dns_nameservers` (String) The DNS nameservers used for DHCP responses, either "upstream_dns", "google_dns", "opendns", or a newline seperated string of IP addresses or domain names
- `group_policy_id` (String) The id of the desired group policy to apply to the VLAN
- `id` (String) The VLAN ID of the VLAN
- `ipv6` (Attributes) IPv6 configuration on the VLAN (see [below for nested schema](#nestedatt--ipv6))
- `mandatory_dhcp` (Attributes) Mandatory DHCP will enforce that clients connecting to this VLAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above (see [below for nested schema](#nestedatt--mandatory_dhcp))
- `mask` (Number) Mask used for the subnet of all bound to the template networks. Applicable only for template network.
- `name` (String) The name of the VLAN
- `reserved_ip_ranges` (Attributes Set) The DHCP reserved IP ranges on the VLAN (see [below for nested schema](#nestedatt--reserved_ip_ranges))
- `subnet` (String) The subnet of the VLAN
- `template_vlan_type` (String) Type of subnetting of the VLAN. Applicable only for template network.
                                  Allowed values: [same,unique]
- `vlan_id` (String) vlanId path parameter. Vlan ID
- `vpn_nat_subnet` (String) The translated VPN subnet if VPN and VPN subnet translation are enabled on the VLAN

### Read-Only

- `interface_id` (String) The interface ID of the VLAN

<a id="nestedatt--dhcp_options"></a>
### Nested Schema for `dhcp_options`

Optional:

- `code` (String) The code for the DHCP option. This should be an integer between 2 and 254.
- `type` (String) The type for the DHCP option. One of: 'text', 'ip', 'hex' or 'integer'
                                        Allowed values: [hex,integer,ip,text]
- `value` (String) The value for the DHCP option


<a id="nestedatt--ipv6"></a>
### Nested Schema for `ipv6`

Optional:

- `enabled` (Boolean) Enable IPv6 on VLAN
- `prefix_assignments` (Attributes Set) Prefix assignments on the VLAN (see [below for nested schema](#nestedatt--ipv6--prefix_assignments))

<a id="nestedatt--ipv6--prefix_assignments"></a>
### Nested Schema for `ipv6.prefix_assignments`

Optional:

- `autonomous` (Boolean) Auto assign a /64 prefix from the origin to the VLAN
- `origin` (Attributes) The origin of the prefix (see [below for nested schema](#nestedatt--ipv6--prefix_assignments--origin))
- `static_appliance_ip6` (String) Manual configuration of the IPv6 Appliance IP
- `static_prefix` (String) Manual configuration of a /64 prefix on the VLAN

<a id="nestedatt--ipv6--prefix_assignments--origin"></a>
### Nested Schema for `ipv6.prefix_assignments.origin`

Optional:

- `interfaces` (Set of String) Interfaces associated with the prefix
- `type` (String) Type of the origin
                                                    Allowed values: [independent,internet]




<a id="nestedatt--mandatory_dhcp"></a>
### Nested Schema for `mandatory_dhcp`

Optional:

- `enabled` (Boolean) Enable Mandatory DHCP on VLAN.


<a id="nestedatt--reserved_ip_ranges"></a>
### Nested Schema for `reserved_ip_ranges`

Optional:

- `comment` (String) A text comment for the reserved range
- `end` (String) The last IP in the reserved range
- `start` (String) The first IP in the reserved range

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_networks_appliance_vlans.example "network_id,vlan_id"
```
