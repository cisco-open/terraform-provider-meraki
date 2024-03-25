
resource "meraki_networks_appliance_traffic_shaping_custom_performance_classes" "example" {

  network_id = "string"
  parameters = {

    max_jitter          = 100
    max_latency         = 100
    max_loss_percentage = 5
    name                = "myCustomPerformanceClass"
  }
}

output "meraki_networks_appliance_traffic_shaping_custom_performance_classes_example" {
  value = meraki_networks_appliance_traffic_shaping_custom_performance_classes.example
}