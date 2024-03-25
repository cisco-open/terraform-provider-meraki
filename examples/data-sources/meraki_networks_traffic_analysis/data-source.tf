
data "meraki_networks_traffic_analysis" "example" {

  network_id = "string"
}

output "meraki_networks_traffic_analysis_example" {
  value = data.meraki_networks_traffic_analysis.example.item
}
