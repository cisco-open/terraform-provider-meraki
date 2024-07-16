terraform {
  required_providers {
    meraki = {
      version = "0.2.7-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

#
# # Fetching a list of networks in the organization
#
data "meraki_networks" "my_networks" {
  provider        = meraki
  organization_id = "828099381482762270"

}

#
# 1. Corporate (802.1x) on vlan 100
#
resource "meraki_networks_wireless_ssids" "my_splash_ssid" {
  for_each = {
    for idx, network in data.meraki_networks.my_networks.items : idx => network 
    if contains(network.product_types, "wireless")
  }

  provider                      = meraki
  network_id                    = each.value.id
  number                        = 12
  auth_mode                     = "open"
  enabled                       = false
  name                          = "Wide open"
  ip_assignment_mode            = "Bridge mode"
  default_vlan_id               = "300"
  splash_page                   = "Click-through splash page"
} 