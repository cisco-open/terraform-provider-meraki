
data "meraki_networks_events_event_types" "example" {

  network_id = "string"
}

output "meraki_networks_events_event_types_example" {
  value = data.meraki_networks_events_event_types.example.items
}
