terraform {
  required_providers {
    meraki = {
      version = "0.2.10-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_alerts_settings" "example" {
  network_id = "L_828099381482775374"
  alerts = [{
    type    = "ampMalwareDetected"
    enabled = true
  }]
  default_destinations = {
    all_admins      = true
    # emails          = []
    http_server_ids = ["aHR0cHM6Ly93ZWJob29rLnNpdGU="]
    snmp            = false
  }
}