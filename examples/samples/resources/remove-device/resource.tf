terraform {
  required_providers {
    meraki = {
      version = "0.1.0-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-open/meraki" to use downloaded version from registry
    }
  }
}

resource "meraki_networks_devices_remove" "example" {
  provider   = meraki
  network_id = "L_828099381482771185"
  parameters = {
    serial = "QBSA-D8CD-5LR6"
  }
}