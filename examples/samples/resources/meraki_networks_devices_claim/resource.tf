terraform {
  required_providers {
    meraki = {
      version = "0.2.13-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_devices_claim" "example" {

  network_id = "L_828099381482775374"
  parameters = {

    serials = []
  }
}