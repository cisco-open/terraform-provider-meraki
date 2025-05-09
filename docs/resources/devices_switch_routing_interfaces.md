---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_devices_switch_routing_interfaces Resource - terraform-provider-meraki"
subcategory: "switch"
description: |-
  
---

# meraki_devices_switch_routing_interfaces (Resource)



## Example Usage

```terraform
resource "meraki_devices_switch_routing_interfaces" "example" {

  default_gateway = "192.168.1.1"
  interface_ip    = "192.168.1.2"
  ipv6 = {

    address         = "2001:db8::1"
    assignment_mode = "static"
    gateway         = "2001:db8::2"
    prefix          = "2001:db8::/32"
  }
  multicast_routing = "disabled"
  name              = "L3 interface"
  ospf_settings = {

    area               = "0"
    cost               = 1
    is_passive_enabled = true
  }
  serial  = "string"
  subnet  = "192.168.1.0/24"
  vlan_id = 100
}

output "meraki_devices_switch_routing_interfaces_example" {
  value = meraki_devices_switch_routing_interfaces.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `serial` (String) serial path parameter.

### Optional

- `default_gateway` (String) IPv4 default gateway
- `interface_id` (String) The ID
- `interface_ip` (String) IPv4 address
- `ipv6` (Attributes) IPv6 addressing (see [below for nested schema](#nestedatt--ipv6))
- `multicast_routing` (String) Multicast routing status
                                  Allowed values: [IGMP snooping querier,disabled,enabled]
- `name` (String) The name
- `ospf_settings` (Attributes) IPv4 OSPF Settings (see [below for nested schema](#nestedatt--ospf_settings))
- `ospf_v3` (Attributes) IPv6 OSPF Settings (see [below for nested schema](#nestedatt--ospf_v3))
- `subnet` (String) IPv4 subnet
- `vlan_id` (Number) VLAN ID

### Read-Only

- `uplink_v4` (Boolean) When true, this interface is used as static IPv4 uplink
- `uplink_v6` (Boolean) When true, this interface is used as static IPv6 uplink

<a id="nestedatt--ipv6"></a>
### Nested Schema for `ipv6`

Optional:

- `address` (String) IPv6 address
- `assignment_mode` (String) Assignment mode
- `gateway` (String) IPv6 gateway
- `prefix` (String) IPv6 subnet


<a id="nestedatt--ospf_settings"></a>
### Nested Schema for `ospf_settings`

Optional:

- `area` (String) Area ID
- `cost` (Number) OSPF Cost
- `is_passive_enabled` (Boolean) Disable sending Hello packets on this interface's IPv4 area


<a id="nestedatt--ospf_v3"></a>
### Nested Schema for `ospf_v3`

Optional:

- `area` (String) Area ID
- `cost` (Number) OSPF Cost
- `is_passive_enabled` (Boolean) Disable sending Hello packets on this interface's IPv6 area

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_devices_switch_routing_interfaces.example "interface_id,serial"
```
