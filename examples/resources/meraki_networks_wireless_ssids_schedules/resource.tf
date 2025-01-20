
resource "meraki_networks_wireless_ssids_schedules" "example" {

  enabled    = true
  network_id = "string"
  number     = "string"
  ranges = [{

    end_day    = "Tuesday"
    end_time   = "05:00"
    start_day  = "Tuesday"
    start_time = "01:00"
  }]
  ranges_in_seconds = [{

    end   = 1
    start = 604800
  }]
}

output "meraki_networks_wireless_ssids_schedules_example" {
  value = meraki_networks_wireless_ssids_schedules.example
}