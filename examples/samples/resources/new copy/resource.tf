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
resource "meraki_networks_wireless_ssids" "my_8021x_ssid" {
  for_each = {
    for idx, network in data.meraki_networks.my_networks.items : idx => network 
    if contains(network.product_types, "wireless")
  }

  provider                      = meraki
  network_id                    = each.value.id
  number                        = 10
  auth_mode                     = "8021x-radius"
  encryption_mode               = "wpa-eap"
  wpa_encryption_mode           = "WPA2 only"
  name                          = "Try me"
  ip_assignment_mode            = "Bridge mode"
  enabled                       = true
  default_vlan_id               = "100"
  radius_servers                = [{
    host                        = "1.2.3.4"
    port                        = 1812
    secret                      = "SuperSecretPassword"
  }]
  
} 