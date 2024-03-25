
data "meraki_networks_sensor_alerts_overview_by_metric" "example" {

  interval   = 1
  network_id = "string"
  t0         = "string"
  t1         = "string"
  timespan   = 1.0
}

output "meraki_networks_sensor_alerts_overview_by_metric_example" {
  value = data.meraki_networks_sensor_alerts_overview_by_metric.example.items
}
