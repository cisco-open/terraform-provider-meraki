---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_integrations_xdr_networks Data Source - terraform-provider-meraki"
subcategory: "organizations"
description: |-
  
---

# meraki_organizations_integrations_xdr_networks (Data Source)



## Example Usage

```terraform
data "meraki_organizations_integrations_xdr_networks" "example" {

  ending_before   = "string"
  network_ids     = ["string"]
  organization_id = "string"
  per_page        = 1
  starting_after  = "string"
}

output "meraki_organizations_integrations_xdr_networks_example" {
  value = data.meraki_organizations_integrations_xdr_networks.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `ending_before` (String) endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `network_ids` (List of String) networkIds query parameter. Optional parameter to filter the results by network IDs
- `per_page` (Number) perPage query parameter. The number of entries per page returned. Acceptable range is 3 100. Default is 20.
- `starting_after` (String) startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `items` (Attributes List) List of networks with XDR enabled (see [below for nested schema](#nestedatt--item--items))
- `meta` (Attributes) Metadata relevant to the paginated dataset (see [below for nested schema](#nestedatt--item--meta))

<a id="nestedatt--item--items"></a>
### Nested Schema for `item.items`

Read-Only:

- `enabled` (Boolean) Represents whether XDR is enabled for the network
- `is_eligible` (Boolean) Represents whether the network is eligible for XDR
- `name` (String) The name of the network
- `network_id` (String) Network ID
- `product_types` (List of String) List of products that have XDR enabled


<a id="nestedatt--item--meta"></a>
### Nested Schema for `item.meta`

Read-Only:

- `counts` (Attributes) Counts relating to the paginated dataset (see [below for nested schema](#nestedatt--item--meta--counts))

<a id="nestedatt--item--meta--counts"></a>
### Nested Schema for `item.meta.counts`

Read-Only:

- `items` (Attributes) Counts relating to the paginated networks (see [below for nested schema](#nestedatt--item--meta--counts--items))

<a id="nestedatt--item--meta--counts--items"></a>
### Nested Schema for `item.meta.counts.items`

Read-Only:

- `remaining` (Number) The number of networks in the dataset that are available on subsequent pages
- `total` (Number) The total number of networks in the dataset
