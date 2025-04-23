
resource "meraki_devices_sensor_commands_create" "example" {

  serial = "string"
  parameters = {

    operation = "disableDownstreamPower"
  }
}

output "meraki_devices_sensor_commands_create_example" {
  value = meraki_devices_sensor_commands_create.example
}