
data "meraki_networks_cellular_gateway_dhcp" "example" {

  network_id = "string"
}

output "meraki_networks_cellular_gateway_dhcp_example" {
  value = data.meraki_networks_cellular_gateway_dhcp.example.item
}
