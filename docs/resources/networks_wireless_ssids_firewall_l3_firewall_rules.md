---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_wireless_ssids_firewall_l3_firewall_rules Resource - terraform-provider-meraki"
subcategory: "wireless"
description: |-
  
---

# meraki_networks_wireless_ssids_firewall_l3_firewall_rules (Resource)



## Example Usage

```terraform
resource "meraki_networks_wireless_ssids_firewall_l3_firewall_rules" "example" {

  allow_lan_access = true
  network_id       = "string"
  number           = "string"
  rules = [{

    comment   = "Allow TCP traffic to subnet with HTTP servers."
    dest_cidr = "192.168.1.0/24"
    dest_port = "443"
    policy    = "allow"
    protocol  = "tcp"
  }]
}

output "meraki_networks_wireless_ssids_firewall_l3_firewall_rules_example" {
  value = meraki_networks_wireless_ssids_firewall_l3_firewall_rules.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID
- `number` (String) number path parameter.

### Optional

- `allow_lan_access` (Boolean) Allows wireless client access to local LAN (boolean value - true allows access and false denies access)
- `rules` (Attributes Set) An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule). (see [below for nested schema](#nestedatt--rules))
- `rules_response` (Attributes Set) An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule). (see [below for nested schema](#nestedatt--rules_response))

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Optional:

- `comment` (String) Description of the rule (optional)
- `dest_cidr` (String) Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
- `dest_port` (String) Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
- `ip_ver` (String) Ip Ver
- `policy` (String) 'allow' or 'deny' traffic specified by this rule
                                        Allowed values: [allow,deny]
- `protocol` (String) The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
                                        Allowed values: [any,icmp,icmp6,tcp,udp]


<a id="nestedatt--rules_response"></a>
### Nested Schema for `rules_response`

Read-Only:

- `comment` (String) Description of the rule (optional)
- `dest_cidr` (String) Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
- `dest_port` (String) Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
- `ip_ver` (String) Ip Version
- `policy` (String) 'allow' or 'deny' traffic specified by this rule
- `protocol` (String) The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_networks_wireless_ssids_firewall_l3_firewall_rules.example "network_id,number"
```
