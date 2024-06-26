---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_adaptive_policy_policies Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_organizations_adaptive_policy_policies (Data Source)



## Example Usage

```terraform
data "meraki_organizations_adaptive_policy_policies" "example" {

  organization_id = "string"
}

output "meraki_organizations_adaptive_policy_policies_example" {
  value = data.meraki_organizations_adaptive_policy_policies.example.items
}

data "meraki_organizations_adaptive_policy_policies" "example" {

  organization_id = "string"
}

output "meraki_organizations_adaptive_policy_policies_example" {
  value = data.meraki_organizations_adaptive_policy_policies.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) id path parameter.
- `organization_id` (String) organizationId path parameter. Organization ID

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))
- `items` (Attributes List) Array of ResponseOrganizationsGetOrganizationAdaptivePolicyPolicies (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `acls` (Attributes Set) (see [below for nested schema](#nestedatt--item--acls))
- `adaptive_policy_id` (String)
- `created_at` (String)
- `destination_group` (Attributes) (see [below for nested schema](#nestedatt--item--destination_group))
- `last_entry_rule` (String)
- `source_group` (Attributes) (see [below for nested schema](#nestedatt--item--source_group))
- `updated_at` (String)

<a id="nestedatt--item--acls"></a>
### Nested Schema for `item.acls`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedatt--item--destination_group"></a>
### Nested Schema for `item.destination_group`

Read-Only:

- `id` (String)
- `name` (String)
- `sgt` (Number)


<a id="nestedatt--item--source_group"></a>
### Nested Schema for `item.source_group`

Read-Only:

- `id` (String)
- `name` (String)
- `sgt` (Number)



<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `acls` (Attributes Set) (see [below for nested schema](#nestedatt--items--acls))
- `adaptive_policy_id` (String)
- `created_at` (String)
- `destination_group` (Attributes) (see [below for nested schema](#nestedatt--items--destination_group))
- `last_entry_rule` (String)
- `source_group` (Attributes) (see [below for nested schema](#nestedatt--items--source_group))
- `updated_at` (String)

<a id="nestedatt--items--acls"></a>
### Nested Schema for `items.acls`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedatt--items--destination_group"></a>
### Nested Schema for `items.destination_group`

Read-Only:

- `id` (String)
- `name` (String)
- `sgt` (Number)


<a id="nestedatt--items--source_group"></a>
### Nested Schema for `items.source_group`

Read-Only:

- `id` (String)
- `name` (String)
- `sgt` (Number)
