
data "meraki_networks_sensor_relationships" "example" {

  network_id = "string"
}

output "meraki_networks_sensor_relationships_example" {
  value = data.meraki_networks_sensor_relationships.example.items
}
