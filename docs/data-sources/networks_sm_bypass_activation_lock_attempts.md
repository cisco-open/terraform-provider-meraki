---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_sm_bypass_activation_lock_attempts Data Source - terraform-provider-meraki"
subcategory: "sm"
description: |-
  
---

# meraki_networks_sm_bypass_activation_lock_attempts (Data Source)



## Example Usage

```terraform
data "meraki_networks_sm_bypass_activation_lock_attempts" "example" {

  attempt_id = "string"
  network_id = "string"
}

output "meraki_networks_sm_bypass_activation_lock_attempts_example" {
  value = data.meraki_networks_sm_bypass_activation_lock_attempts.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `attempt_id` (String) attemptId path parameter. Attempt ID
- `network_id` (String) networkId path parameter. Network ID

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `data` (Attributes) (see [below for nested schema](#nestedatt--item--data))
- `id` (String)
- `status` (String)

<a id="nestedatt--item--data"></a>
### Nested Schema for `item.data`

Read-Only:

- `status_2090938209` (Attributes) (see [below for nested schema](#nestedatt--item--data--status_2090938209))
- `status_38290139892` (Attributes) (see [below for nested schema](#nestedatt--item--data--status_38290139892))

<a id="nestedatt--item--data--status_2090938209"></a>
### Nested Schema for `item.data.status_2090938209`

Read-Only:

- `errors` (List of String)
- `success` (Boolean)


<a id="nestedatt--item--data--status_38290139892"></a>
### Nested Schema for `item.data.status_38290139892`

Read-Only:

- `success` (Boolean)
