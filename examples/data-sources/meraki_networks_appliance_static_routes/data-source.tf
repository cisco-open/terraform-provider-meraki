
data "meraki_networks_appliance_static_routes" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_static_routes_example" {
  value = data.meraki_networks_appliance_static_routes.example.items
}

data "meraki_networks_appliance_static_routes" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_static_routes_example" {
  value = data.meraki_networks_appliance_static_routes.example.item
}
