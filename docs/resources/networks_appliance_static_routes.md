---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_appliance_static_routes Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_appliance_static_routes (Resource)



## Example Usage

```terraform
resource "meraki_networks_appliance_static_routes" "example" {
  provider   = meraki
  gateway_ip = "1.2.3.5"
  name       = "My route"
  network_id = "string"
  subnet     = "192.168.1.0/24"
}

output "meraki_networks_appliance_static_routes_example" {
  value = meraki_networks_appliance_static_routes.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Optional

- `enabled` (Boolean) The enabled state of the static route
- `fixed_ip_assignments` (Attributes) The DHCP fixed IP assignments on the static route. This should be an object that contains mappings from MAC addresses to objects that themselves each contain "ip" and "name" string fields. See the sample request/response for more details. (see [below for nested schema](#nestedatt--fixed_ip_assignments))
- `gateway_ip` (String) The gateway IP (next hop) of the static route
- `gateway_vlan_id` (Number) The gateway IP (next hop) VLAN ID of the static route
- `gateway_vlan_id_rs` (String) The gateway IP (next hop) VLAN ID of the static route
- `name` (String) The name of the new static route
- `reserved_ip_ranges` (Attributes Set) The DHCP reserved IP ranges on the static route (see [below for nested schema](#nestedatt--reserved_ip_ranges))
- `static_route_id` (String) staticRouteId path parameter. Static route ID
- `subnet` (String) The subnet of the static route

### Read-Only

- `id` (String) The ID of this resource.
- `ip_version` (Number)

<a id="nestedatt--fixed_ip_assignments"></a>
### Nested Schema for `fixed_ip_assignments`

Read-Only:

- `attribute_22_33_44_55_66_77` (Attributes) (see [below for nested schema](#nestedatt--fixed_ip_assignments--attribute_22_33_44_55_66_77))

<a id="nestedatt--fixed_ip_assignments--attribute_22_33_44_55_66_77"></a>
### Nested Schema for `fixed_ip_assignments.attribute_22_33_44_55_66_77`

Read-Only:

- `ip` (String)
- `name` (String)



<a id="nestedatt--reserved_ip_ranges"></a>
### Nested Schema for `reserved_ip_ranges`

Optional:

- `comment` (String) A text comment for the reserved range
- `end` (String) The last IP in the reserved range
- `start` (String) The first IP in the reserved range

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_networks_appliance_static_routes.example "network_id,static_route_id"
```