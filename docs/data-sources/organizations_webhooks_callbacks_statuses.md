---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_webhooks_callbacks_statuses Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_organizations_webhooks_callbacks_statuses (Data Source)



## Example Usage

```terraform
data "meraki_organizations_webhooks_callbacks_statuses" "example" {

  callback_id     = "string"
  organization_id = "string"
}

output "meraki_organizations_webhooks_callbacks_statuses_example" {
  value = data.meraki_organizations_webhooks_callbacks_statuses.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `callback_id` (String) callbackId path parameter. Callback ID
- `organization_id` (String) organizationId path parameter. Organization ID

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `callback_id` (String) The ID of the callback
- `created_by` (Attributes) Information around who initiated the callback (see [below for nested schema](#nestedatt--item--created_by))
- `errors` (List of String) The errors returned by the callback
- `status` (String) The status of the callback
- `webhook` (Attributes) The webhook receiver used by the callback to send results (see [below for nested schema](#nestedatt--item--webhook))

<a id="nestedatt--item--created_by"></a>
### Nested Schema for `item.created_by`

Read-Only:

- `admin_id` (String) The ID of the user who initiated the callback


<a id="nestedatt--item--webhook"></a>
### Nested Schema for `item.webhook`

Read-Only:

- `http_server` (Attributes) The webhook receiver used for the callback webhook (see [below for nested schema](#nestedatt--item--webhook--http_server))
- `payload_template` (Attributes) The payload template of the webhook used for the callback (see [below for nested schema](#nestedatt--item--webhook--payload_template))
- `sent_at` (String) The timestamp the callback was sent to the webhook receiver
- `url` (String) The webhook receiver URL where the callback will be sent

<a id="nestedatt--item--webhook--http_server"></a>
### Nested Schema for `item.webhook.http_server`

Read-Only:

- `id` (String) The webhook receiver ID that will receive information


<a id="nestedatt--item--webhook--payload_template"></a>
### Nested Schema for `item.webhook.payload_template`

Read-Only:

- `id` (String) The ID of the payload template