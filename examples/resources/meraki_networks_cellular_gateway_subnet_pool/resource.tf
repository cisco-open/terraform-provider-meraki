
resource "meraki_networks_cellular_gateway_subnet_pool" "example" {

  cidr       = "192.168.0.0/16"
  mask       = 24
  network_id = "string"
}

output "meraki_networks_cellular_gateway_subnet_pool_example" {
  value = meraki_networks_cellular_gateway_subnet_pool.example
}