
data "meraki_networks_policies_by_client" "example" {

  ending_before  = "string"
  network_id     = "string"
  per_page       = 1
  starting_after = "string"
  t0             = "string"
  timespan       = 1.0
}

output "meraki_networks_policies_by_client_example" {
  value = data.meraki_networks_policies_by_client.example.items
}
