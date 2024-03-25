
resource "meraki_networks_sensor_mqtt_brokers" "example" {

  enabled        = true
  mqtt_broker_id = "string"
  network_id     = "string"
}

output "meraki_networks_sensor_mqtt_brokers_example" {
  value = meraki_networks_sensor_mqtt_brokers.example
}