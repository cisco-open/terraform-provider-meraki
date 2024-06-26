---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_administered_licensing_subscription_subscriptions_claim Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_administered_licensing_subscription_subscriptions_claim (Resource)

~>Warning: This resource does not represent a real-world entity in Meraki Dashboard, therefore changing or deleting this resource on its own has no immediate effect. Instead, it is a task part of a Meraki Dashboard workflow. It is executed in Meraki without any additional verification. It does not check if it was executed before or if a similar configuration or action 
already existed previously.

## Example Usage

```terraform
resource "meraki_administered_licensing_subscription_subscriptions_claim" "example" {

  validate = false
  parameters = {

    claim_key       = "S2345-6789A-BCDEF-GHJKM"
    description     = "Subscription for all main offices"
    name            = "Corporate subscription"
    organization_id = "12345678910"
  }
}

output "meraki_administered_licensing_subscription_subscriptions_claim_example" {
  value = meraki_administered_licensing_subscription_subscriptions_claim.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Attributes) (see [below for nested schema](#nestedatt--parameters))

### Optional

- `validate` (Boolean) validate query parameter. Check if the provided claim key is valid and can be claimed into the organization.

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `claim_key` (String) The subscription's claim key
- `description` (String) Extra details or notes about the subscription
- `name` (String) Friendly name to identify the subscription
- `organization_id` (String) The id of the organization claiming the subscription


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `counts` (Attributes) Numeric breakdown of network and entitlement counts (see [below for nested schema](#nestedatt--item--counts))
- `description` (String) Subscription description
- `end_date` (String) Subscription expiration date
- `entitlements` (Attributes Set) Entitlement info (see [below for nested schema](#nestedatt--item--entitlements))
- `name` (String) Subscription name
- `product_types` (List of String) Products the subscription has entitlements for
- `start_date` (String) Subscription start date
- `status` (String) Subscription status
- `subscription_id` (String) Subscription's ID
- `web_order_id` (String) Web order id

<a id="nestedatt--item--counts"></a>
### Nested Schema for `item.counts`

Read-Only:

- `networks` (Number) Number of networks bound to this subscription
- `seats` (Attributes) Seat distribution (see [below for nested schema](#nestedatt--item--counts--seats))

<a id="nestedatt--item--counts--seats"></a>
### Nested Schema for `item.counts.seats`

Read-Only:

- `assigned` (Number) Number of seats in use
- `available` (Number) Number of seats available for use
- `limit` (Number) Total number of seats provided by this subscription



<a id="nestedatt--item--entitlements"></a>
### Nested Schema for `item.entitlements`

Read-Only:

- `seats` (Attributes) Seat distribution (see [below for nested schema](#nestedatt--item--entitlements--seats))
- `sku` (String) SKU of the required product

<a id="nestedatt--item--entitlements--seats"></a>
### Nested Schema for `item.entitlements.seats`

Read-Only:

- `assigned` (Number) Number of seats in use
- `available` (Number) Number of seats available for use
- `limit` (Number) Total number of seats provided by this subscription for this sku
