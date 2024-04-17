
resource "meraki_networks_traffic_analysis" "example" {

  custom_pie_chart_items = [{

    name  = "Item from hostname"
    type  = "host"
    value = "example.com"
  }]
  mode       = "disabled"
  network_id = "string"
}

output "meraki_networks_traffic_analysis_example" {
  value = meraki_networks_traffic_analysis.example
}