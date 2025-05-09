---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_devices_cellular_gateway_port_forwarding_rules Resource - terraform-provider-meraki"
subcategory: "cellularGateway"
description: |-
  
---

# meraki_devices_cellular_gateway_port_forwarding_rules (Resource)



## Example Usage

```terraform
resource "meraki_devices_cellular_gateway_port_forwarding_rules" "example" {

  rules = [{

    access      = "any"
    allowed_ips = ["10.10.10.10", "10.10.10.11"]
    lan_ip      = "172.31.128.5"
    local_port  = "4"
    name        = "test"
    protocol    = "tcp"
    public_port = "11-12"
  }]
  serial = "string"
}

output "meraki_devices_cellular_gateway_port_forwarding_rules_example" {
  value = meraki_devices_cellular_gateway_port_forwarding_rules.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `serial` (String) serial path parameter.

### Optional

- `rules` (Attributes Set) An array of port forwarding params (see [below for nested schema](#nestedatt--rules))

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Optional:

- `access` (String) **any** or **restricted**. Specify the right to make inbound connections on the specified ports or port ranges. If **restricted**, a list of allowed IPs is mandatory.
- `allowed_ips` (Set of String) An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges.
- `lan_ip` (String) The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN
- `local_port` (String) A port or port ranges that will receive the forwarded traffic from the WAN
- `name` (String) A descriptive name for the rule
- `protocol` (String) TCP or UDP
- `public_port` (String) A port or port ranges that will be forwarded to the host on the LAN

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_devices_cellular_gateway_port_forwarding_rules.example "serial"
```
