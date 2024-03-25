
resource "meraki_devices_sensor_relationships" "example" {

  livestream = {

    related_devices = [{

      serial = "string"
    }]
  }
  serial = "string"
}

output "meraki_devices_sensor_relationships_example" {
  value = meraki_devices_sensor_relationships.example
}