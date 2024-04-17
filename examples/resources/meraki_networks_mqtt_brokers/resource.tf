
resource "meraki_networks_mqtt_brokers" "example" {

  network_id = "string"
  parameters = {

    authentication = {

      password = "*****"
      username = "milesmeraki"
    }
    host = "1.2.3.4"
    name = "MQTT_Broker_1"
    port = 443
    security = {

      mode = "tls"
      tls = {

        ca_certificate   = "*****"
        verify_hostnames = true
      }
    }
  }
}

output "meraki_networks_mqtt_brokers_example" {
  value = meraki_networks_mqtt_brokers.example
}