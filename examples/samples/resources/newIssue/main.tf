terraform {
  required_providers {
    meraki = {
      version = "0.2.8-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}


resource "meraki_networks_switch_stacks" "this_ms_stack" {
  network_id = "L_828099381482775374"
  name       = "750-comm-meraki-lab-sw-stack"
  serials    = ["QBSB-AX45-LY9A", "QBSB-BNH2-KDXJ"]
}