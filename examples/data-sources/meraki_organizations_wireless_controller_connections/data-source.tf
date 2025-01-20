
data "meraki_organizations_wireless_controller_connections" "example" {

  controller_serials = ["string"]
  ending_before      = "string"
  network_ids        = ["string"]
  organization_id    = "string"
  per_page           = 1
  starting_after     = "string"
}

output "meraki_organizations_wireless_controller_connections_example" {
  value = data.meraki_organizations_wireless_controller_connections.example.item
}
