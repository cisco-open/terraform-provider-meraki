
resource "meraki_networks_vlan_profiles" "example" {

  iname      = "string"
  name       = "My VLAN profile name"
  network_id = "string"
  vlan_groups = [{

    name     = "named-group-1"
    vlan_ids = "2,5-7"
  }]
  vlan_names = [{

    adaptive_policy_group = {

      id = "791"
    }
    name    = "named-1"
    vlan_id = "1"
  }]
}

output "meraki_networks_vlan_profiles_example" {
  value = meraki_networks_vlan_profiles.example
}