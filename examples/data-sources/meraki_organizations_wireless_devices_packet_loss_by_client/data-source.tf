
data "meraki_organizations_wireless_devices_packet_loss_by_client" "example" {

  bands           = ["string"]
  ending_before   = "string"
  macs            = ["string"]
  network_ids     = ["string"]
  organization_id = "string"
  per_page        = 1
  ssids           = ["string"]
  starting_after  = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
}

output "meraki_organizations_wireless_devices_packet_loss_by_client_example" {
  value = data.meraki_organizations_wireless_devices_packet_loss_by_client.example.items
}
