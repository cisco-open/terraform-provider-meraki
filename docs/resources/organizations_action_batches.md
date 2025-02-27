---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_action_batches Resource - terraform-provider-meraki"
subcategory: "organizations"
description: |-
  
---

# meraki_organizations_action_batches (Resource)



## Example Usage

```terraform
resource "meraki_organizations_action_batches" "example" {

  actions = [{

    operation = "create"
    resource  = "/devices/QXXX-XXXX-XXXX/switch/ports/3"
  }]
  callback = {

    http_server = {

      id = "aHR0cHM6Ly93d3cuZXhhbXBsZS5jb20vd2ViaG9va3M="
    }
    payload_template = {

      id = "wpt_2100"
    }
    shared_secret = "secret"
    url           = "https://webhook.site/28efa24e-f830-4d9f-a12b-fbb9e5035031"
  }
  confirmed       = true
  organization_id = "string"
  synchronous     = true
}

output "meraki_organizations_action_batches_example" {
  value = meraki_organizations_action_batches.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) ID of the organization this action batch belongs to

### Optional

- `action_batch_id` (String) actionBatchId path parameter. Action batch ID
- `actions` (Attributes Set) A set of changes made as part of this action (<a href='https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/'>more details</a>) (see [below for nested schema](#nestedatt--actions))
- `callback` (Attributes) Information for callback used to send back results (see [below for nested schema](#nestedatt--callback))
- `confirmed` (Boolean) Flag describing whether the action should be previewed before executing or not
- `synchronous` (Boolean) Flag describing whether actions should run synchronously or asynchronously

### Read-Only

- `id` (String) ID of the action batch. Can be used to check the status of the action batch at /organizations/{organizationId}/actionBatches/{actionBatchId}
- `status` (Attributes) Status of action batch (see [below for nested schema](#nestedatt--status))

<a id="nestedatt--actions"></a>
### Nested Schema for `actions`

Optional:

- `body` (String) Data provided in the body of the Action. Contents depend on the Action type
- `operation` (String) The operation to be used by this action
- `resource` (String) Unique identifier for the resource to be acted on


<a id="nestedatt--callback"></a>
### Nested Schema for `callback`

Optional:

- `http_server` (Attributes) The webhook receiver used for the callback webhook. (see [below for nested schema](#nestedatt--callback--http_server))
- `payload_template` (Attributes) The payload template of the webhook used for the callback (see [below for nested schema](#nestedatt--callback--payload_template))
- `shared_secret` (String) A shared secret that will be included in the requests sent to the callback URL. It can be used to verify that the request was sent by Meraki. If using this field, please also specify an url.
- `url` (String) The callback URL for the webhook target. This was either provided in the original request or comes from a configured webhook receiver

Read-Only:

- `id` (String) The ID of the callback. To check the status of the callback, use this ID in a request to /webhooks/callbacks/statuses/{id}
- `status` (String) The status of the callback

<a id="nestedatt--callback--http_server"></a>
### Nested Schema for `callback.http_server`

Optional:

- `id` (String) The webhook receiver ID that will receive information. If specifying this, please leave the url and sharedSecret fields blank.


<a id="nestedatt--callback--payload_template"></a>
### Nested Schema for `callback.payload_template`

Optional:

- `id` (String) The ID of the payload template. Defaults to 'wpt_00005' for the Callback (included) template.



<a id="nestedatt--status"></a>
### Nested Schema for `status`

Read-Only:

- `completed` (Boolean) Flag describing whether all actions in the action batch have completed
- `created_resources` (Attributes Set) Resources created as a result of this action batch (see [below for nested schema](#nestedatt--status--created_resources))
- `errors` (Set of String) List of errors encountered when running actions in the action batch
- `failed` (Boolean) Flag describing whether any actions in the action batch failed

<a id="nestedatt--status--created_resources"></a>
### Nested Schema for `status.created_resources`

Read-Only:

- `id` (String) ID of the created resource
- `uri` (String) URI, not including base, of the created resource

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_organizations_action_batches.example "action_batch_id,organization_id"
```
