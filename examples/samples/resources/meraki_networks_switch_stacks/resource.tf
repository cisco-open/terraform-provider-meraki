terraform {
  required_providers {
    meraki = {
      version = "1.1.5-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}


resource "meraki_networks_switch_stacks" "example" {

  name       = "A cool stack"
  network_id = "N_828099381482850157"
  serials    = ["QBSB-AX45-LY9A", "QBSB-BNH2-KDXJ"]
}

output "meraki_networks_switch_stacks_example" {
  value = meraki_networks_switch_stacks.example
}