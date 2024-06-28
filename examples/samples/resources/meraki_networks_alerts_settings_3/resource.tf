terraform {
  required_providers {
    meraki = {
      version = "0.2.5-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_alerts_settings" "example" {
  network_id = "L_828099381482771185"
  alerts = [{
    alert_destinations = {
      emails = ["test_email@meraki.com"]
      snmp = false
      all_admins = false
    }
    type = "ampMalwareDetected"
    enabled = true
  }]
  default_destinations = {
    all_admins = true
    emails = ["tf@domain.com"]
    snmp = false
  }
  # muting = {
  #   by_port_schedules = {
  #     enabled = true
  #   }
  # }
}

