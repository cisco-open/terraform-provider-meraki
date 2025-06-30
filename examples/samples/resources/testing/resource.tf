terraform {
  required_providers {
    meraki = {
      version = "1.1.5-beta"
      source  = "hashicorp.com/edu/meraki"
    }
  }
}



provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_group_policies" "group_policy_byod" {
  name       = "BYOD"
  network_id = "L_828099381482771185"
  vlan_tagging = {
    settings = "custom"
    vlan_id  = "20"
  }
}

resource "meraki_networks_group_policies" "group_policy_guest" {
  depends_on = [meraki_networks_group_policies.group_policy_byod]
  name       = "GUEST"
  network_id = "L_828099381482771185"
  vlan_tagging = {
    settings = "custom"
    vlan_id  = "30"
  }
}

# resource "meraki_networks_group_policies" "group_policy_pan" {
#   depends_on = [ meraki_networks_group_policies.group_policy_guest ]
#   name       = "PAN"
#   network_id = "L_828099381482771185"
#   vlan_tagging = {
#     settings = "custom"
#     vlan_id  = "15"
#   }
# }

# resource "meraki_networks_group_policies" "group_policy_internal" {
#   depends_on = [ meraki_networks_group_policies.group_policy_pan ]
#   name       = "INTERNAL 2"
#   network_id = "L_828099381482771185"
#   vlan_tagging = {
#     settings = "custom"
#     vlan_id  = "5"
#   }
# }