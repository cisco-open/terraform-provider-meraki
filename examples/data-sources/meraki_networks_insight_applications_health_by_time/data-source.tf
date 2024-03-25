
data "meraki_networks_insight_applications_health_by_time" "example" {

  application_id = "string"
  network_id     = "string"
  resolution     = 1
  t0             = "string"
  t1             = "string"
  timespan       = 1.0
}

output "meraki_networks_insight_applications_health_by_time_example" {
  value = data.meraki_networks_insight_applications_health_by_time.example.items
}
