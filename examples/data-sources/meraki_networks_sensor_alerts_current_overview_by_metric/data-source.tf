
data "meraki_networks_sensor_alerts_current_overview_by_metric" "example" {

  network_id = "string"
}

output "meraki_networks_sensor_alerts_current_overview_by_metric_example" {
  value = data.meraki_networks_sensor_alerts_current_overview_by_metric.example.item
}
