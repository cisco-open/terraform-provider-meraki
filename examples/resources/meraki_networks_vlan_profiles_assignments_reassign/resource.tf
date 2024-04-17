
resource "meraki_networks_vlan_profiles_assignments_reassign" "example" {

  network_id = "string"
  parameters = {

    serials   = ["Q234-ABCD-5678"]
    stack_ids = ["1234"]
    vlan_profile = {

      iname = "Profile1"
    }
  }
}

output "meraki_networks_vlan_profiles_assignments_reassign_example" {
  value = meraki_networks_vlan_profiles_assignments_reassign.example
}