terraform {
  required_providers {
    meraki = {
      version = "0.2.3-alpha"
      source = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

variable "my_network_id" {
  type                      = string
  default                   = "*******" # Branch-1234
}
resource "meraki_networks_appliance_vlans" "my_mx" {
  appliance_ip              = "192.168.14.2"
  id                        = "14"
  name                      = "My VLAN"
  network_id                = var.my_network_id 
  subnet                    = "192.168.14.0/24"
}