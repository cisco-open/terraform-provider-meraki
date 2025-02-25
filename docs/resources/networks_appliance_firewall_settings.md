---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_appliance_firewall_settings Resource - terraform-provider-meraki"
subcategory: "appliance"
description: |-
  
---

# meraki_networks_appliance_firewall_settings (Resource)



## Example Usage

```terraform
resource "meraki_networks_appliance_firewall_settings" "example" {

  network_id = "string"
  spoofing_protection = {

    ip_source_guard = {

      mode = "block"
    }
  }
}

output "meraki_networks_appliance_firewall_settings_example" {
  value = meraki_networks_appliance_firewall_settings.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Optional

- `spoofing_protection` (Attributes) Spoofing protection settings (see [below for nested schema](#nestedatt--spoofing_protection))

<a id="nestedatt--spoofing_protection"></a>
### Nested Schema for `spoofing_protection`

Optional:

- `ip_source_guard` (Attributes) IP source address spoofing settings (see [below for nested schema](#nestedatt--spoofing_protection--ip_source_guard))

<a id="nestedatt--spoofing_protection--ip_source_guard"></a>
### Nested Schema for `spoofing_protection.ip_source_guard`

Optional:

- `mode` (String) Mode of protection
                                              Allowed values: [block,log]

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_networks_appliance_firewall_settings.example "network_id"
```
