---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_devices_live_tools_throughput_test Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_devices_live_tools_throughput_test (Data Source)



## Example Usage

```terraform
data "meraki_devices_live_tools_throughput_test" "example" {

  serial             = "string"
  throughput_test_id = "string"
}

output "meraki_devices_live_tools_throughput_test_example" {
  value = data.meraki_devices_live_tools_throughput_test.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `serial` (String) serial path parameter.
- `throughput_test_id` (String) throughputTestId path parameter. Throughput test ID

### Read-Only

- `item` (Attributes) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `error` (String) Description of the error.
- `request` (Attributes) The parameters of the throughput test request (see [below for nested schema](#nestedatt--item--request))
- `result` (Attributes) Result of the throughput test request (see [below for nested schema](#nestedatt--item--result))
- `status` (String) Status of the throughput test request
- `throughput_test_id` (String) ID of throughput test job
- `url` (String) GET this url to check the status of your throughput test request

<a id="nestedatt--item--request"></a>
### Nested Schema for `item.request`

Read-Only:

- `serial` (String) Device serial number


<a id="nestedatt--item--result"></a>
### Nested Schema for `item.result`

Read-Only:

- `speeds` (Attributes) Shows the speeds (Mbps) (see [below for nested schema](#nestedatt--item--result--speeds))

<a id="nestedatt--item--result--speeds"></a>
### Nested Schema for `item.result.speeds`

Read-Only:

- `downstream` (Number) Shows the download speed from shard (Mbps)