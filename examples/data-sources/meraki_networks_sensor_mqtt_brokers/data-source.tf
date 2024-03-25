
data "meraki_networks_sensor_mqtt_brokers" "example" {

  network_id = "string"
}

output "meraki_networks_sensor_mqtt_brokers_example" {
  value = data.meraki_networks_sensor_mqtt_brokers.example.items
}

data "meraki_networks_sensor_mqtt_brokers" "example" {

  network_id = "string"
}

output "meraki_networks_sensor_mqtt_brokers_example" {
  value = data.meraki_networks_sensor_mqtt_brokers.example.item
}
