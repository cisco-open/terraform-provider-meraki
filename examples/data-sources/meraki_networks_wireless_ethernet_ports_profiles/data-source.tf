
data "meraki_networks_wireless_ethernet_ports_profiles" "example" {

  network_id = "string"
  profile_id = "string"
}

output "meraki_networks_wireless_ethernet_ports_profiles_example" {
  value = data.meraki_networks_wireless_ethernet_ports_profiles.example.item
}
