
resource "meraki_networks_cellular_gateway_connectivity_monitoring_destinations" "example" {

  destinations = [{

    default     = false
    description = "Google"
    ip          = "8.8.8.8"
  }]
  network_id = "string"
}

output "meraki_networks_cellular_gateway_connectivity_monitoring_destinations_example" {
  value = meraki_networks_cellular_gateway_connectivity_monitoring_destinations.example
}