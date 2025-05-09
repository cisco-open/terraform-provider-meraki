---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_appliance_vpn_bgp Resource - terraform-provider-meraki"
subcategory: "appliance"
description: |-
  
---

# meraki_networks_appliance_vpn_bgp (Resource)



## Example Usage

```terraform
resource "meraki_networks_appliance_vpn_bgp" "example" {

  as_number       = 64515
  enabled         = true
  ibgp_hold_timer = 120
  neighbors = [{

    allow_transit = true
    authentication = {

      password = "abc123"
    }
    ebgp_hold_timer = 180
    ebgp_multihop   = 2
    ip              = "10.10.10.22"
    ipv6 = {

      address = "2002::1234:abcd:ffff:c0a8:101"
    }
    next_hop_ip      = "1.2.3.4"
    receive_limit    = 120
    remote_as_number = 64343
    source_interface = "wan1"
    ttl_security = {

      enabled = false
    }
  }]
  network_id = "string"
}

output "meraki_networks_appliance_vpn_bgp_example" {
  value = meraki_networks_appliance_vpn_bgp.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Optional

- `as_number` (Number) The number of the Autonomous System to which the appliance belongs
- `enabled` (Boolean) Whether BGP is enabled on the appliance
- `ibgp_hold_timer` (Number) The iBGP hold time in seconds
- `neighbors` (Attributes Set) List of eBGP neighbor configurations (see [below for nested schema](#nestedatt--neighbors))

<a id="nestedatt--neighbors"></a>
### Nested Schema for `neighbors`

Optional:

- `allow_transit` (Boolean) Whether the appliance will advertise routes learned from other Autonomous Systems
- `authentication` (Attributes) Authentication settings between BGP peers (see [below for nested schema](#nestedatt--neighbors--authentication))
- `ebgp_hold_timer` (Number) The eBGP hold time in seconds for the neighbor
- `ebgp_multihop` (Number) The number of hops the appliance must traverse to establish a peering relationship with the neighbor
- `ip` (String) The IPv4 address of the neighbor
- `ipv6` (Attributes) Information regarding IPv6 address of the neighbor (see [below for nested schema](#nestedatt--neighbors--ipv6))
- `next_hop_ip` (String) The IPv4 address of the neighbor that will establish a TCP session with the appliance
- `receive_limit` (Number) The maximum number of routes that the appliance can receive from the neighbor
- `remote_as_number` (Number) Remote AS number of the neighbor
- `source_interface` (String) The output interface the appliance uses to establish a peering relationship with the neighbor
- `ttl_security` (Attributes) Settings for BGP TTL security to protect BGP peering sessions from forged IP attacks (see [below for nested schema](#nestedatt--neighbors--ttl_security))

<a id="nestedatt--neighbors--authentication"></a>
### Nested Schema for `neighbors.authentication`

Optional:

- `password` (String, Sensitive) Password to configure MD5 authentication between BGP peers


<a id="nestedatt--neighbors--ipv6"></a>
### Nested Schema for `neighbors.ipv6`

Optional:

- `address` (String) The IPv6 address of the neighbor


<a id="nestedatt--neighbors--ttl_security"></a>
### Nested Schema for `neighbors.ttl_security`

Optional:

- `enabled` (Boolean) Whether BGP TTL security is enabled

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_networks_appliance_vpn_bgp.example "network_id"
```
