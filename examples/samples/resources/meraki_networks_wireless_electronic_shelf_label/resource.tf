terraform {
  required_providers {
    meraki = {
      version = "1.0.3-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

resource "meraki_networks_wireless_electronic_shelf_label" "example" {

  enabled    = false
  hostname   = "N_24329156"
  network_id = "L_828099381482771185"
}

output "meraki_networks_wireless_electronic_shelf_label_example" {
  value = meraki_networks_wireless_electronic_shelf_label.example
}