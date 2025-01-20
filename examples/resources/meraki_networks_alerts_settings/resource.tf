
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