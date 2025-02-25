---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_appliance_firewall_port_forwarding_rules Data Source - terraform-provider-meraki"
subcategory: "appliance"
description: |-
  
---

# meraki_networks_appliance_firewall_port_forwarding_rules (Data Source)



## Example Usage

```terraform
data "meraki_networks_appliance_firewall_port_forwarding_rules" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_firewall_port_forwarding_rules_example" {
  value = data.meraki_networks_appliance_firewall_port_forwarding_rules.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `rules` (Attributes Set) An array of port forwarding rules (see [below for nested schema](#nestedatt--item--rules))

<a id="nestedatt--item--rules"></a>
### Nested Schema for `item.rules`

Read-Only:

- `allowed_ips` (List of String) An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges (or any)
- `lan_ip` (String) IP address of the device subject to port forwarding
- `local_port` (String) The port or port range that receives forwarded traffic from the WAN
- `name` (String) Name of the rule
- `protocol` (String) Protocol the rule applies to
- `public_port` (String) The port or port range forwarded to the host on the LAN
- `uplink` (String) The physical WAN interface on which the traffic arrives; allowed values vary by appliance model and configuration
