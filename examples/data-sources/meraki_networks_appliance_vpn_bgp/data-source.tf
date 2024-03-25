
data "meraki_networks_appliance_vpn_bgp" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_vpn_bgp_example" {
  value = data.meraki_networks_appliance_vpn_bgp.example.item
}
