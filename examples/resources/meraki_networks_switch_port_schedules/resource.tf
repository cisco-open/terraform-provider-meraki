
resource "meraki_networks_switch_port_schedules" "example" {

  name       = "Weekdays schedule"
  network_id = "string"
  port_schedule = {

    friday = {

      active = true
      from   = "9:00"
      to     = "17:00"
    }
    monday = {

      active = true
      from   = "9:00"
      to     = "17:00"
    }
    saturday = {

      active = false
      from   = "0:00"
      to     = "24:00"
    }
    sunday = {

      active = false
      from   = "0:00"
      to     = "24:00"
    }
    thursday = {

      active = true
      from   = "9:00"
      to     = "17:00"
    }
    tuesday = {

      active = true
      from   = "9:00"
      to     = "17:00"
    }
    wednesday = {

      active = true
      from   = "9:00"
      to     = "17:00"
    }
  }
}

output "meraki_networks_switch_port_schedules_example" {
  value = meraki_networks_switch_port_schedules.example
}