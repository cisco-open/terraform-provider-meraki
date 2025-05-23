---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_devices_sensor_commands Data Source - terraform-provider-meraki"
subcategory: "sensor"
description: |-
  
---

# meraki_devices_sensor_commands (Data Source)



## Example Usage

```terraform
data "meraki_devices_sensor_commands" "example" {

  ending_before  = "string"
  operations     = ["string"]
  per_page       = 1
  serial         = "string"
  sort_order     = "string"
  starting_after = "string"
  t0             = "string"
  t1             = "string"
  timespan       = 1.0
}

output "meraki_devices_sensor_commands_example" {
  value = data.meraki_devices_sensor_commands.example.items
}

data "meraki_devices_sensor_commands" "example" {

  ending_before  = "string"
  operations     = ["string"]
  per_page       = 1
  serial         = "string"
  sort_order     = "string"
  starting_after = "string"
  t0             = "string"
  t1             = "string"
  timespan       = 1.0
}

output "meraki_devices_sensor_commands_example" {
  value = data.meraki_devices_sensor_commands.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `command_id` (String) commandId path parameter. Command ID
- `ending_before` (String) endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `operations` (List of String) operations query parameter. Optional parameter to filter commands by operation. Allowed values are disableDownstreamPower, enableDownstreamPower, cycleDownstreamPower, and refreshData.
- `per_page` (Number) perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 10.
- `serial` (String) serial path parameter.
- `sort_order` (String) sortOrder query parameter. Sorted order of entries. Order options are 'ascending' and 'descending'. Default is 'descending'.
- `starting_after` (String) startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
- `t0` (String) t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 30 days from today.
- `t1` (String) t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 30 days after t0.
- `timespan` (Number) timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 30 days. The default is 30 days.

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))
- `items` (Attributes List) Array of ResponseSensorGetDeviceSensorCommands (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `command_id` (String) ID to check the status of the command request
- `completed_at` (String) Time when the command was completed
- `created_at` (String) Time when the command was triggered
- `created_by` (Attributes) Information about the admin who triggered the command (see [below for nested schema](#nestedatt--item--created_by))
- `errors` (List of String) Array of errors if failed
- `operation` (String) Operation run on the sensor
- `status` (String) Status of the command request

<a id="nestedatt--item--created_by"></a>
### Nested Schema for `item.created_by`

Read-Only:

- `admin_id` (String) ID of the admin
- `email` (String) Email of the admin
- `name` (String) Name of the admin



<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `command_id` (String) ID to check the status of the command request
- `completed_at` (String) Time when the command was completed
- `created_at` (String) Time when the command was triggered
- `created_by` (Attributes) Information about the admin who triggered the command (see [below for nested schema](#nestedatt--items--created_by))
- `errors` (List of String) Array of errors if failed
- `operation` (String) Operation run on the sensor
- `status` (String) Status of the command request

<a id="nestedatt--items--created_by"></a>
### Nested Schema for `items.created_by`

Read-Only:

- `admin_id` (String) ID of the admin
- `email` (String) Email of the admin
- `name` (String) Name of the admin
