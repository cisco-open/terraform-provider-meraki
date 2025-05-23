---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_appliance_content_filtering Data Source - terraform-provider-meraki"
subcategory: "appliance"
description: |-
  
---

# meraki_networks_appliance_content_filtering (Data Source)



## Example Usage

```terraform
data "meraki_networks_appliance_content_filtering" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_content_filtering_example" {
  value = data.meraki_networks_appliance_content_filtering.example.item
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

- `allowed_url_patterns` (List of String)
- `blocked_url_categories` (Attributes Set) (see [below for nested schema](#nestedatt--item--blocked_url_categories))
- `blocked_url_patterns` (List of String)
- `url_category_list_size` (String)

<a id="nestedatt--item--blocked_url_categories"></a>
### Nested Schema for `item.blocked_url_categories`

Read-Only:

- `id` (String)
- `name` (String)
