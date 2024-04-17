terraform {
  required_providers {
    meraki = {
      version = "0.2.0-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

resource "meraki_networks_unbind" "example" {

  network_id = "L_828099381482771185"
  parameters = {
    retain_configs = false
  }
}