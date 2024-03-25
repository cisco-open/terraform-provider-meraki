
data "meraki_networks_appliance_connectivity_monitoring_destinations" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_connectivity_monitoring_destinations_example" {
  value = data.meraki_networks_appliance_connectivity_monitoring_destinations.example.item
}
