
resource "meraki_networks_wireless_ethernet_ports_profiles_assign" "example" {

  network_id = "string"
  parameters = {

    profile_id = "1001"
    serials    = ["Q234-ABCD-0001", "Q234-ABCD-0002", "Q234-ABCD-0003"]
  }
}

output "meraki_networks_wireless_ethernet_ports_profiles_assign_example" {
  value = meraki_networks_wireless_ethernet_ports_profiles_assign.example
}