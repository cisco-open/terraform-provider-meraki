---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_appliance_settings Resource - terraform-provider-meraki"
subcategory: "appliance"
description: |-
  
---

# meraki_networks_appliance_settings (Resource)



## Example Usage

```terraform
resource "meraki_networks_appliance_settings" "example" {

  client_tracking_method = "MAC address"
  deployment_mode        = "routed"
  dynamic_dns = {

    enabled = true
    prefix  = "test"
  }
  network_id = "string"
}

output "meraki_networks_appliance_settings_example" {
  value = meraki_networks_appliance_settings.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Optional

- `client_tracking_method` (String) Client tracking method of a network
                                  Allowed values: [IP address,MAC address,Unique client identifier]
- `deployment_mode` (String) Deployment mode of a network
                                  Allowed values: [passthrough,routed]
- `dynamic_dns` (Attributes) Dynamic DNS settings for a network (see [below for nested schema](#nestedatt--dynamic_dns))

<a id="nestedatt--dynamic_dns"></a>
### Nested Schema for `dynamic_dns`

Optional:

- `enabled` (Boolean) Dynamic DNS enabled
- `prefix` (String) Dynamic DNS url prefix. DDNS must be enabled to update

Read-Only:

- `url` (String) Dynamic DNS url. DDNS must be enabled to update

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_networks_appliance_settings.example "network_id"
```
