---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_appliance_warm_spare Data Source - terraform-provider-meraki"
subcategory: "appliance"
description: |-
  
---

# meraki_networks_appliance_warm_spare (Data Source)



## Example Usage

```terraform
data "meraki_networks_appliance_warm_spare" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_warm_spare_example" {
  value = data.meraki_networks_appliance_warm_spare.example.item
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

- `enabled` (Boolean) Is the warm spare enabled
- `primary_serial` (String) Serial number of the primary appliance
- `spare_serial` (String) Serial number of the warm spare appliance
- `uplink_mode` (String) Uplink mode, either virtual or public
- `wan1` (Attributes) WAN 1 IP and subnet (see [below for nested schema](#nestedatt--item--wan1))
- `wan2` (Attributes) WAN 2 IP and subnet (see [below for nested schema](#nestedatt--item--wan2))

<a id="nestedatt--item--wan1"></a>
### Nested Schema for `item.wan1`

Read-Only:

- `ip` (String) IP address used for WAN 1
- `subnet` (String) Subnet used for WAN 1


<a id="nestedatt--item--wan2"></a>
### Nested Schema for `item.wan2`

Read-Only:

- `ip` (String) IP address used for WAN 2
- `subnet` (String) Subnet used for WAN 2
