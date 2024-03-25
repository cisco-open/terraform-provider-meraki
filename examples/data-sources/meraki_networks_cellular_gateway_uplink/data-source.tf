
data "meraki_networks_cellular_gateway_uplink" "example" {

  network_id = "string"
}

output "meraki_networks_cellular_gateway_uplink_example" {
  value = data.meraki_networks_cellular_gateway_uplink.example.item
}
