
data "meraki_networks_events" "example" {

  client_ip            = "string"
  client_mac           = "string"
  client_name          = "string"
  device_mac           = "string"
  device_name          = "string"
  device_serial        = "string"
  ending_before        = "string"
  excluded_event_types = ["string"]
  included_event_types = ["string"]
  network_id           = "string"
  per_page             = 1
  product_type         = "string"
  sm_device_mac        = "string"
  sm_device_name       = "string"
  starting_after       = "string"
}

output "meraki_networks_events_example" {
  value = data.meraki_networks_events.example.item
}
