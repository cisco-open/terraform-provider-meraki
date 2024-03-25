
data "meraki_networks_alerts_history" "example" {

  ending_before  = "string"
  network_id     = "string"
  per_page       = 1
  starting_after = "string"
}

output "meraki_networks_alerts_history_example" {
  value = data.meraki_networks_alerts_history.example.items
}
