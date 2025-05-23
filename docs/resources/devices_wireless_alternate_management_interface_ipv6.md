---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_devices_wireless_alternate_management_interface_ipv6 Resource - terraform-provider-meraki"
subcategory: "wireless"
description: |-
  
---

# meraki_devices_wireless_alternate_management_interface_ipv6 (Resource)





~>Warning: This resource does not represent a real-world entity in Meraki Dashboard, therefore changing or deleting this resource on its own has no immediate effect. Instead, it is a task part of a Meraki Dashboard workflow. It is executed in Meraki without any additional verification. It does not check if it was executed before or if a similar configuration or action 
already existed previously.


## Example Usage

```terraform
resource "meraki_devices_wireless_alternate_management_interface_ipv6" "example" {

  serial = "string"
  parameters = {

    addresses = [{

      address         = "2001:db8:3c4d:15::1"
      assignment_mode = "static"
      gateway         = "fe80:db8:c15:c0:d0c::10ca:1d02"
      nameservers = {

        addresses = ["2001:db8:3c4d:15::1", "2001:db8:3c4d:15::1"]
      }
      prefix   = "2001:db8:3c4d:15::/64"
      protocol = "ipv6"
    }]
  }
}

output "meraki_devices_wireless_alternate_management_interface_ipv6_example" {
  value = meraki_devices_wireless_alternate_management_interface_ipv6.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Attributes) (see [below for nested schema](#nestedatt--parameters))
- `serial` (String) serial path parameter.

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `addresses` (Attributes Set) configured alternate management interface addresses (see [below for nested schema](#nestedatt--parameters--addresses))

<a id="nestedatt--parameters--addresses"></a>
### Nested Schema for `parameters.addresses`

Optional:

- `address` (String) The IP address configured for the alternate management interface
- `assignment_mode` (String) The type of address assignment. Either static or dynamic.
                                              Allowed values: [dynamic,static]
- `gateway` (String) The gateway address configured for the alternate managment interface
- `nameservers` (Attributes) The DNS servers settings for this address. (see [below for nested schema](#nestedatt--parameters--addresses--nameservers))
- `prefix` (String) The IPv6 prefix length of the IPv6 interface. Required if IPv6 object is included.
- `protocol` (String) The IP protocol used for the address
                                              Allowed values: [ipv4,ipv6]

<a id="nestedatt--parameters--addresses--nameservers"></a>
### Nested Schema for `parameters.addresses.nameservers`

Optional:

- `addresses` (List of String) Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.




<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `addresses` (Attributes Set) configured alternate management interface addresses (see [below for nested schema](#nestedatt--item--addresses))

<a id="nestedatt--item--addresses"></a>
### Nested Schema for `item.addresses`

Read-Only:

- `address` (String) The IP address configured for the alternate management interface
- `assignment_mode` (String) The type of address assignment. Either static or dynamic.
                                                Allowed values: [dynamic,static]
- `gateway` (String) The gateway address configured for the alternate managment interface
- `nameservers` (Attributes) The DNS servers settings for this address. (see [below for nested schema](#nestedatt--item--addresses--nameservers))
- `prefix` (String) The IPv6 prefix of the interface. Required if IPv6 object is included.
- `protocol` (String) The IP protocol used for the address
                                                Allowed values: [ipv4,ipv6]

<a id="nestedatt--item--addresses--nameservers"></a>
### Nested Schema for `item.addresses.nameservers`

Read-Only:

- `addresses` (Set of String) Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.
