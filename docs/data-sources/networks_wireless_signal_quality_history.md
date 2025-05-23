---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_wireless_signal_quality_history Data Source - terraform-provider-meraki"
subcategory: "wireless"
description: |-
  
---

# meraki_networks_wireless_signal_quality_history (Data Source)



## Example Usage

```terraform
data "meraki_networks_wireless_signal_quality_history" "example" {

  ap_tag          = "string"
  auto_resolution = false
  band            = "string"
  client_id       = "string"
  device_serial   = "string"
  network_id      = "string"
  resolution      = 1
  ssid            = 1
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
}

output "meraki_networks_wireless_signal_quality_history_example" {
  value = data.meraki_networks_wireless_signal_quality_history.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Optional

- `ap_tag` (String) apTag query parameter. Filter results by AP tag; either :clientId or :deviceSerial must be jointly specified.
- `auto_resolution` (Boolean) autoResolution query parameter. Automatically select a data resolution based on the given timespan; this overrides the value specified by the 'resolution' parameter. The default setting is false.
- `band` (String) band query parameter. Filter results by band (either '2.4', '5' or '6').
- `client_id` (String) clientId query parameter. Filter results by network client.
- `device_serial` (String) deviceSerial query parameter. Filter results by device.
- `resolution` (Number) resolution query parameter. The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400.
- `ssid` (Number) ssid query parameter. Filter results by SSID number.
- `t0` (String) t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
- `t1` (String) t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
- `timespan` (Number) timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.

### Read-Only

- `items` (Attributes List) Array of ResponseWirelessGetNetworkWirelessSignalQualityHistory (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `end_ts` (String) The end time of the query range
- `rssi` (Number) Received signal strength indicator
- `snr` (Number) Signal to noise ratio
- `start_ts` (String) The start time of the query range
