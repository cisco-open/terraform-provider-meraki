---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_administered_licensing_subscription_subscriptions Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_administered_licensing_subscription_subscriptions (Data Source)



## Example Usage

```terraform
data "meraki_administered_licensing_subscription_subscriptions" "example" {

  end_date         = "string"
  ending_before    = "string"
  organization_ids = ["string"]
  per_page         = 1
  product_types    = ["string"]
  start_date       = "string"
  starting_after   = "string"
  statuses         = ["string"]
  subscription_ids = ["string"]
}

output "meraki_administered_licensing_subscription_subscriptions_example" {
  value = data.meraki_administered_licensing_subscription_subscriptions.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `end_date` (String) endDate query parameter. Filter subscriptions by end date, ISO 8601 format. To filter with a range of dates, use 'endDate[
]=?' in the request. Accepted options include lt, gt, lte, gte.
- `ending_before` (String) endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `organization_ids` (List of String) organizationIds query parameter. Organizations to get associated subscriptions for
- `per_page` (Number) perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.
- `product_types` (List of String) productTypes query parameter. List of product types that returned subscriptions need to have entitlements for.
- `start_date` (String) startDate query parameter. Filter subscriptions by start date, ISO 8601 format. To filter with a range of dates, use 'startDate[
]=?' in the request. Accepted options include lt, gt, lte, gte.
- `starting_after` (String) startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `statuses` (List of String) statuses query parameter. List of statuses that returned subscriptions can have
- `subscription_ids` (List of String) subscriptionIds query parameter. List of subscription ids to fetch

### Read-Only

- `items` (Attributes List) Array of ResponseLicensingGetAdministeredLicensingSubscriptionSubscriptions (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `counts` (Attributes) Numeric breakdown of network and entitlement counts (see [below for nested schema](#nestedatt--items--counts))
- `description` (String) Subscription description
- `end_date` (String) Subscription expiration date
- `entitlements` (Attributes Set) Entitlement info (see [below for nested schema](#nestedatt--items--entitlements))
- `name` (String) Subscription name
- `product_types` (List of String) Products the subscription has entitlements for
- `start_date` (String) Subscription start date
- `status` (String) Subscription status
- `subscription_id` (String) Subscription's ID
- `web_order_id` (String) Web order id

<a id="nestedatt--items--counts"></a>
### Nested Schema for `items.counts`

Read-Only:

- `networks` (Number) Number of networks bound to this subscription
- `seats` (Attributes) Seat distribution (see [below for nested schema](#nestedatt--items--counts--seats))

<a id="nestedatt--items--counts--seats"></a>
### Nested Schema for `items.counts.seats`

Read-Only:

- `assigned` (Number) Number of seats in use
- `available` (Number) Number of seats available for use
- `limit` (Number) Total number of seats provided by this subscription



<a id="nestedatt--items--entitlements"></a>
### Nested Schema for `items.entitlements`

Read-Only:

- `seats` (Attributes) Seat distribution (see [below for nested schema](#nestedatt--items--entitlements--seats))
- `sku` (String) SKU of the required product

<a id="nestedatt--items--entitlements--seats"></a>
### Nested Schema for `items.entitlements.seats`

Read-Only:

- `assigned` (Number) Number of seats in use
- `available` (Number) Number of seats available for use
- `limit` (Number) Total number of seats provided by this subscription for this sku