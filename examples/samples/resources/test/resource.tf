terraform {
  required_providers {
    meraki = {
      version = "1.2.4-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

data "meraki_networks" "my_networks" {
  organization_id = "828099381482762270"
}


resource "meraki_networks_wireless_ssids" "my_8021x_ssid" {
  for_each = {
    for idx, network in data.meraki_networks.my_networks.items : idx => network
    if contains(network.product_types, "wireless")
  }


  network_id         = each.value.id
  band_selection     = "Dual band operation"
  number             = 10
  visible            = true
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