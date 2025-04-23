# This plan will configure 3 wireless SSIDs across all networks in the Meraki organization:
# 1. Corporate (802.1x) on vlan 100
# 2. Mobile (WPA-PSK) on vlan 200
# 3. Guests (Portal) on vlan 300

terraform {
  required_providers {
    meraki = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/meraki"
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


  network_id         = each.value.id
  number             = 10
  auth_mode          = "8021x-radius"
  name               = "Can't be hacked :-)"
  ip_assignment_mode = "Bridge mode"
  default_vlan_id    = "100"
  radius_servers = [{
    host   = "1.2.3.4"
    port   = 1812
    secret = "SuperSecretPassword"
  }]

}

#
# 2. Mobile (WPA-PSK) on vlan 200
#
resource "meraki_networks_wireless_ssids" "my_wpa_psk_ssid" {
  for_each = {
    for idx, network in data.meraki_networks.my_networks.items : idx => network
    if contains(network.product_types, "wireless")
  }


  network_id          = each.value.id
  number              = 11
  auth_mode           = "psk"
  encryption_mode     = "wpa"
  wpa_encryption_mode = "WPA2 only"
  psk                 = "BadPassword"
  name                = "Somewhat secure"
  ip_assignment_mode  = "Bridge mode"
  default_vlan_id     = "200"
}

# #
# # 3. Guests (Portal) on vlan 300
# #
resource "meraki_networks_wireless_ssids" "my_splash_ssid" {
  for_each = {
    for idx, network in data.meraki_networks.my_networks.items : idx => network
    if contains(network.product_types, "wireless")
  }


  network_id         = each.value.id
  number             = 12
  auth_mode          = "open"
  name               = "Unsecure guests"
  ip_assignment_mode = "Bridge mode"
  default_vlan_id    = "300"
  splash_page        = "Click-through splash page"
}