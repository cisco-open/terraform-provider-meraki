---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_early_access_features_opt_ins Data Source - terraform-provider-meraki"
subcategory: "organizations"
description: |-
  
---

# meraki_organizations_early_access_features_opt_ins (Data Source)



## Example Usage

```terraform
data "meraki_organizations_early_access_features_opt_ins" "example" {

  organization_id = "string"
}

output "meraki_organizations_early_access_features_opt_ins_example" {
  value = data.meraki_organizations_early_access_features_opt_ins.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `opt_in_id` (String) optInId path parameter. Opt in ID
- `organization_id` (String) organizationId path parameter. Organization ID

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `created_at` (String) Time when Early Access Feature was created
- `id` (String) ID of Early Access Feature
- `limit_scope_to_networks` (Attributes Set) Networks assigned to the Early Access Feature (see [below for nested schema](#nestedatt--item--limit_scope_to_networks))
- `opt_out_eligibility` (Attributes) Descriptions of the early access feature (see [below for nested schema](#nestedatt--item--opt_out_eligibility))
- `short_name` (String) Name of Early Access Feature

<a id="nestedatt--item--limit_scope_to_networks"></a>
### Nested Schema for `item.limit_scope_to_networks`

Read-Only:

- `id` (String) ID of Network
- `name` (String) Name of Network


<a id="nestedatt--item--opt_out_eligibility"></a>
### Nested Schema for `item.opt_out_eligibility`

Read-Only:

- `eligible` (Boolean) Condition flag to opt out from the feature
- `help` (Attributes) Additional help information (see [below for nested schema](#nestedatt--item--opt_out_eligibility--help))
- `reason` (String) User friendly message regarding opt-out eligibility

<a id="nestedatt--item--opt_out_eligibility--help"></a>
### Nested Schema for `item.opt_out_eligibility.help`

Read-Only:

- `label` (String) Help link label
- `url` (String) Help link url
