terraform {
  required_providers {
    meraki = {
      source  = "hashicorp.com/edu/meraki"
      version = "1.2.4-beta"

    }
  }
}


provider "meraki" {
  # Configuration options
  meraki_debug = "true"
}

data "meraki_networks" "my_networks" {
  provider        = meraki
  organization_id = "828099381482762270"
}

output "networks" {
  value = data.meraki_networks.my_networks.items
}
#  resource "meraki_networks_group_policies" "foobar" {
#    network_id = "L_828099381482771185"
#    name = "foobar"
# }

resource "meraki_networks_group_policies" "foobar" {
  for_each = {
    for idx, network in data.meraki_networks.my_networks.items : idx => network
    if contains(network.product_types, "appliance") || contains(network.product_types, "wireless")
  }
  network_id = each.value.id
  name = "foobar"
}
