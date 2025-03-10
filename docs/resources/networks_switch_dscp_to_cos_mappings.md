---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_switch_dscp_to_cos_mappings Resource - terraform-provider-meraki"
subcategory: "switch"
description: |-
  
---

# meraki_networks_switch_dscp_to_cos_mappings (Resource)



## Example Usage

```terraform
resource "meraki_networks_switch_dscp_to_cos_mappings" "example" {

  mappings = [{

    cos   = 1
    dscp  = 1
    title = "Video"
  }]
  network_id = "string"
}

output "meraki_networks_switch_dscp_to_cos_mappings_example" {
  value = meraki_networks_switch_dscp_to_cos_mappings.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Optional

- `mappings` (Attributes Set) An array of DSCP to CoS mappings. An empty array will reset the mappings to default. (see [below for nested schema](#nestedatt--mappings))

<a id="nestedatt--mappings"></a>
### Nested Schema for `mappings`

Optional:

- `cos` (Number) The actual layer-2 CoS queue the DSCP value is mapped to. These are not bits set on outgoing frames. Value can be in the range of 0 to 5 inclusive.
- `dscp` (Number) The Differentiated Services Code Point (DSCP) tag in the IP header that will be mapped to a particular Class-of-Service (CoS) queue. Value can be in the range of 0 to 63 inclusive.
- `title` (String) Label for the mapping (optional).

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_networks_switch_dscp_to_cos_mappings.example "network_id"
```
