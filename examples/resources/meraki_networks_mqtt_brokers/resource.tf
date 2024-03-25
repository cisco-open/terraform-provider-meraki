
resource "meraki_networks_mqtt_brokers" "example" {

  network_id = "string"
  parameters = {

    authentication = {

      username = "Username"
    }
    host = "1.1.1.1"
    name = "MQTT_Broker_1"
    port = 1234
    security = {

      mode = "tls"
    }
  }
}

output "meraki_networks_mqtt_brokers_example" {
  value = meraki_networks_mqtt_brokers.example
}