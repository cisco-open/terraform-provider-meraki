
data "meraki_networks_cellular_gateway_subnet_pool" "example" {

  network_id = "string"
}

output "meraki_networks_cellular_gateway_subnet_pool_example" {
  value = data.meraki_networks_cellular_gateway_subnet_pool.example.item
}
