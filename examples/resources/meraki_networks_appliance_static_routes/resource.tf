
resource "meraki_networks_appliance_static_routes" "example" {

  gateway_ip      = "1.2.3.5"
  gateway_vlan_id = "100"
  name            = "My route"
  network_id      = "string"
  subnet          = "192.168.1.0/24"
}

output "meraki_networks_appliance_static_routes_example" {
  value = meraki_networks_appliance_static_routes.example
}