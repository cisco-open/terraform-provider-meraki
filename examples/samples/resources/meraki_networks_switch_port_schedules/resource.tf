terraform {
  required_providers {
    meraki = {
      version = "1.1.3-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_switch_port_schedules" "example" {

  name       = "Weekdays schedule"
  network_id = "L_828099381482771185"
  port_schedule = {

    friday = {

      active = true
      from   = "09:00"
      to     = "24:00"
    }
    monday = {

      active = true
      from   = "09:00"
      to     = "17:00"
    }
    saturday = {

      active = false
      from   = "00:00"
      to     = "24:00"
    }
    sunday = {

      active = false
      from   = "00:00"
      to     = "24:00"
    }
    thursday = {

      active = true
      from   = "09:00"
      to     = "17:00"
    }
    tuesday = {

      active = true
      from   = "09:00"
      to     = "17:00"
    }
    wednesday = {

      active = true
      from   = "09:00"
      to     = "17:00"
    }
  }
}

output "meraki_networks_switch_port_schedules_example" {
  value = meraki_networks_switch_port_schedules.example
}