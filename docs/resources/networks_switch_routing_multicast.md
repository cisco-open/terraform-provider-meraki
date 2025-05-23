---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_switch_routing_multicast Resource - terraform-provider-meraki"
subcategory: "switch"
description: |-
  
---

# meraki_networks_switch_routing_multicast (Resource)



## Example Usage

```terraform
resource "meraki_networks_switch_routing_multicast" "example" {

  default_settings = {

    flood_unknown_multicast_traffic_enabled = true
    igmp_snooping_enabled                   = true
  }
  network_id = "string"
  overrides = [{

    flood_unknown_multicast_traffic_enabled = true
    igmp_snooping_enabled                   = true
    stacks                                  = ["789102", "123456", "129102"]
    switch_profiles                         = ["1234", "4567"]
    switches                                = ["Q234-ABCD-0001", "Q234-ABCD-0002", "Q234-ABCD-0003"]
  }]
}

output "meraki_networks_switch_routing_multicast_example" {
  value = meraki_networks_switch_routing_multicast.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Optional

- `default_settings` (Attributes) Default multicast setting for entire network. IGMP snooping and Flood unknown
      multicast traffic settings are enabled by default. (see [below for nested schema](#nestedatt--default_settings))
- `overrides` (Attributes Set) Array of paired switches/stacks/profiles and corresponding multicast settings.
      An empty array will clear the multicast settings. (see [below for nested schema](#nestedatt--overrides))

<a id="nestedatt--default_settings"></a>
### Nested Schema for `default_settings`

Optional:

- `flood_unknown_multicast_traffic_enabled` (Boolean) Flood unknown multicast traffic enabled for the entire network
- `igmp_snooping_enabled` (Boolean) IGMP snooping enabled for the entire network


<a id="nestedatt--overrides"></a>
### Nested Schema for `overrides`

Optional:

- `flood_unknown_multicast_traffic_enabled` (Boolean) Flood unknown multicast traffic enabled for switches, switch stacks or switch templates
- `igmp_snooping_enabled` (Boolean) IGMP snooping enabled for switches, switch stacks or switch templates
- `stacks` (Set of String) (optional) List of switch stack ids for non-template network
- `switch_profiles` (Set of String) (optional) List of switch templates ids for template network
- `switches` (Set of String) (optional) List of switch serials for non-template network

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_networks_switch_routing_multicast.example "network_id"
```
