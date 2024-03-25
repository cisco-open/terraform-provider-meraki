
data "meraki_devices_sensor_relationships" "example" {

  serial = "string"
}

output "meraki_devices_sensor_relationships_example" {
  value = data.meraki_devices_sensor_relationships.example.items
}
