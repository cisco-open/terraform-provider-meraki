terraform {
  required_providers {
    meraki = {
      version = "1.2.4-beta"
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