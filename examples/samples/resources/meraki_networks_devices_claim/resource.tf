terraform {
  required_providers {
    meraki = {
      version = "1.2.4-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug    = "true"
  meraki_base_url = "http://localhost:3002"
}

resource "meraki_networks_devices_claim" "example" {

  network_id     = "L_828099381482775374"
  add_atomically = false
  parameters = {

    serials = []
  }
}