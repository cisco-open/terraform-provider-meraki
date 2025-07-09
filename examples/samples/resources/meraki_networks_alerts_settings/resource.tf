terraform {
  required_providers {
    meraki = {
      version = "1.1.6-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

################################################################

resource "meraki_networks_alerts_settings" "example" {

  alerts = [{

    alert_destinations = {

      all_admins = false
      emails     = ["miles2@meraki.com"]
    }
    filters = {
      timeout = 10
    }
    enabled = true
    type    = "gatewayDown"
  }]
  default_destinations = {

    all_admins = true
    emails     = ["miles@meraki.com"]
  }
  network_id = "L_828099381482775375"
}

output "meraki_networks_alerts_settings_example" {
  value = meraki_networks_alerts_settings.example
}