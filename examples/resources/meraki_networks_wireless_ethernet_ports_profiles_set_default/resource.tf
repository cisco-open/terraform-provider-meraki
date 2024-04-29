
resource "meraki_networks_wireless_ethernet_ports_profiles_set_default" "example" {

  network_id = "string"
  parameters = {

    profile_id = "1001"
  }
}

output "meraki_networks_wireless_ethernet_ports_profiles_set_default_example" {
  value = meraki_networks_wireless_ethernet_ports_profiles_set_default.example
}