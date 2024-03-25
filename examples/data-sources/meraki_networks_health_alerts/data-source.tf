
data "meraki_networks_health_alerts" "example" {

  network_id = "string"
}

output "meraki_networks_health_alerts_example" {
  value = data.meraki_networks_health_alerts.example.items
}
