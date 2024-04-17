
data "meraki_networks_sm_devices" "example" {

  ending_before  = "string"
  fields         = ["string"]
  ids            = ["string"]
  network_id     = "string"
  per_page       = 1
  scope          = ["string"]
  serials        = ["string"]
  starting_after = "string"
  system_types   = ["string"]
  uuids          = ["string"]
  wifi_macs      = ["string"]
}

output "meraki_networks_sm_devices_example" {
  value = data.meraki_networks_sm_devices.example.items
}
