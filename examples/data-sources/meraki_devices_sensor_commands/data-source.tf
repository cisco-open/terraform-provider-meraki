
data "meraki_devices_sensor_commands" "example" {

  ending_before  = "string"
  operations     = ["string"]
  per_page       = 1
  serial         = "string"
  sort_order     = "string"
  starting_after = "string"
  t0             = "string"
  t1             = "string"
  timespan       = 1.0
}

output "meraki_devices_sensor_commands_example" {
  value = data.meraki_devices_sensor_commands.example.items
}

data "meraki_devices_sensor_commands" "example" {

  ending_before  = "string"
  operations     = ["string"]
  per_page       = 1
  serial         = "string"
  sort_order     = "string"
  starting_after = "string"
  t0             = "string"
  t1             = "string"
  timespan       = 1.0
}

output "meraki_devices_sensor_commands_example" {
  value = data.meraki_devices_sensor_commands.example.item
}
