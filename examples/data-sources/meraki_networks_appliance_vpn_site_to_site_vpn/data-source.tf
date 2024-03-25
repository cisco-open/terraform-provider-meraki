
data "meraki_networks_appliance_vpn_site_to_site_vpn" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_vpn_site_to_site_vpn_example" {
  value = data.meraki_networks_appliance_vpn_site_to_site_vpn.example.item
}
