
data "meraki_networks_wireless_mesh_statuses" "example" {

  ending_before  = "string"
  network_id     = "string"
  per_page       = 1
  starting_after = "string"
}

output "meraki_networks_wireless_mesh_statuses_example" {
  value = data.meraki_networks_wireless_mesh_statuses.example.item
}
