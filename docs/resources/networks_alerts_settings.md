---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_alerts_settings Resource - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_networks_alerts_settings (Resource)



## Example Usage

```terraform
resource "meraki_networks_alerts_settings" "example" {

  alerts = [{

    alert_destinations = {

      all_admins      = false
      emails          = ["miles@meraki.com"]
      http_server_ids = ["aHR0cHM6Ly93d3cuZXhhbXBsZS5jb20vd2ViaG9va3M="]
      sms_numbers     = ["+15555555555"]
      snmp            = false
    }
    enabled = true
    filters = {

      conditions = [{

        direction = "+"
        duration  = 1
        threshold = 72.5
        type      = "temperature"
        unit      = "celsius"
      }]
      failure_type    = "802.1X auth fail"
      lookback_window = 360
      min_duration    = 60
      name            = "Filter"
      period          = 1800
      priority        = ""
      regex           = "[a-z]"
      selector        = "{'smartSensitivity':'medium','smartEnabled':false,'eventReminderPeriodSecs':10800}"
      serials         = ["Q234-ABCD-0001", "Q234-ABCD-0002", "Q234-ABCD-0003"]
      ssid_num        = 1
      tag             = "tag1"
      threshold       = 30
      timeout         = 60
    }
    type = "gatewayDown"
  }]
  default_destinations = {

    all_admins      = true
    emails          = ["miles@meraki.com"]
    http_server_ids = ["aHR0cHM6Ly93d3cuZXhhbXBsZS5jb20vd2ViaG9va3M="]
    snmp            = true
  }
  muting = {

    by_port_schedules = {

      enabled = true
    }
  }
  network_id = "string"
}

output "meraki_networks_alerts_settings_example" {
  value = meraki_networks_alerts_settings.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Optional

- `alerts` (Attributes Set) Alert-specific configuration for each type. Only alerts that pertain to the network can be updated. (see [below for nested schema](#nestedatt--alerts))
- `default_destinations` (Attributes) The network-wide destinations for all alerts on the network. (see [below for nested schema](#nestedatt--default_destinations))
- `muting` (Attributes) Mute alerts under certain conditions (see [below for nested schema](#nestedatt--muting))

### Read-Only

- `alerts_response` (Attributes Set) Alert-specific configuration for each type. Only alerts that pertain to the network can be updated. (see [below for nested schema](#nestedatt--alerts_response))

<a id="nestedatt--alerts"></a>
### Nested Schema for `alerts`

Optional:

- `alert_destinations` (Attributes) A hash of destinations for this specific alert (see [below for nested schema](#nestedatt--alerts--alert_destinations))
- `enabled` (Boolean) A boolean depicting if the alert is turned on or off
- `filters` (Attributes) A hash of specific configuration data for the alert. Only filters specific to the alert will be updated. (see [below for nested schema](#nestedatt--alerts--filters))
- `type` (String) The type of alert

<a id="nestedatt--alerts--alert_destinations"></a>
### Nested Schema for `alerts.alert_destinations`

Optional:

- `all_admins` (Boolean) If true, then all network admins will receive emails for this alert
- `emails` (Set of String) A list of emails that will receive information about the alert
- `http_server_ids` (Set of String) A list of HTTP server IDs to send a Webhook to for this alert
- `sms_numbers` (Set of String) A list of phone numbers that will receive text messages about the alert. Only available for sensors status alerts.
- `snmp` (Boolean) If true, then an SNMP trap will be sent for this alert if there is an SNMP trap server configured for this network


<a id="nestedatt--alerts--filters"></a>
### Nested Schema for `alerts.filters`

Optional:

- `conditions` (Attributes Set) Conditions (see [below for nested schema](#nestedatt--alerts--filters--conditions))
- `failure_type` (String) Failure Type
- `lookback_window` (Number) Loopback Window (in sec)
- `min_duration` (Number) Min Duration
- `name` (String) Name
- `period` (Number) Period
- `priority` (String) Priority
- `regex` (String) Regex
- `selector` (String) Selector
- `serials` (Set of String) Serials
- `ssid_num` (Number) SSID Number
- `tag` (String) Tag
- `threshold` (Number) Threshold
- `timeout` (Number) Timeout

<a id="nestedatt--alerts--filters--conditions"></a>
### Nested Schema for `alerts.filters.conditions`

Optional:

- `direction` (String) Direction
                                                    Allowed values: [+,-]
- `duration` (Number) Duration
- `threshold` (Number) Threshold
- `type` (String) Type of condition
- `unit` (String) Unit




<a id="nestedatt--default_destinations"></a>
### Nested Schema for `default_destinations`

Optional:

- `all_admins` (Boolean) If true, then all network admins will receive emails.
- `emails` (Set of String) A list of emails that will receive the alert(s).
- `http_server_ids` (Set of String) A list of HTTP server IDs to send a Webhook to
- `snmp` (Boolean) If true, then an SNMP trap will be sent if there is an SNMP trap server configured for this network.


<a id="nestedatt--muting"></a>
### Nested Schema for `muting`

Optional:

- `by_port_schedules` (Attributes) Mute wireless unreachable alerts based on switch port schedules (see [below for nested schema](#nestedatt--muting--by_port_schedules))

<a id="nestedatt--muting--by_port_schedules"></a>
### Nested Schema for `muting.by_port_schedules`

Optional:

- `enabled` (Boolean) If true, then wireless unreachable alerts will be muted when caused by a port schedule



<a id="nestedatt--alerts_response"></a>
### Nested Schema for `alerts_response`

Read-Only:

- `alert_destinations` (Attributes) A hash of destinations for this specific alert (see [below for nested schema](#nestedatt--alerts_response--alert_destinations))
- `enabled` (Boolean) A boolean depicting if the alert is turned on or off
- `filters` (Attributes) A hash of specific configuration data for the alert. Only filters specific to the alert will be updated. (see [below for nested schema](#nestedatt--alerts_response--filters))
- `type` (String) The type of alert

<a id="nestedatt--alerts_response--alert_destinations"></a>
### Nested Schema for `alerts_response.alert_destinations`

Read-Only:

- `all_admins` (Boolean) If true, then all network admins will receive emails for this alert
- `emails` (Set of String) A list of emails that will receive information about the alert
- `http_server_ids` (Set of String) A list of HTTP server IDs to send a Webhook to for this alert
- `sms_numbers` (Set of String) A list of phone numbers that will receive text messages about the alert. Only available for sensors status alerts.
- `snmp` (Boolean) If true, then an SNMP trap will be sent for this alert if there is an SNMP trap server configured for this network


<a id="nestedatt--alerts_response--filters"></a>
### Nested Schema for `alerts_response.filters`

Read-Only:

- `conditions` (Attributes Set) Conditions (see [below for nested schema](#nestedatt--alerts_response--filters--conditions))
- `failure_type` (String) Failure Type
- `lookback_window` (Number) Loopback Window (in sec)
- `min_duration` (Number) Min Duration
- `name` (String) Name
- `period` (Number) Period
- `priority` (String) Priority
- `regex` (String) Regex
- `selector` (String) Selector
- `serials` (Set of String) Serials
- `ssid_num` (Number) SSID Number
- `tag` (String) Tag
- `threshold` (Number) Threshold
- `timeout` (Number) Timeout

<a id="nestedatt--alerts_response--filters--conditions"></a>
### Nested Schema for `alerts_response.filters.conditions`

Optional:

- `direction` (String) Direction
                                                    Allowed values: [+,-]

Read-Only:

- `duration` (Number) Duration
- `threshold` (Number) Threshold
- `type` (String) Type of condition
- `unit` (String) Unit

## Import

Import is supported using the following syntax:

```shell
terraform import meraki_networks_alerts_settings.example "network_id"
```
