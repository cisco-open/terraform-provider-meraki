terraform {
  required_providers {
    meraki = {
      version = "1.2.2-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug    = "true"

}

resource "meraki_networks_group_policies" "foobar" {
  network_id = "L_828099381482777237"
  name = "foobar"
}
