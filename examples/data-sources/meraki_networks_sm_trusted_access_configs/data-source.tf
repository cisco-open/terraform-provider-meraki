
data "meraki_networks_sm_trusted_access_configs" "example" {

  ending_before  = "string"
  network_id     = "string"
  per_page       = 1
  starting_after = "string"
}

output "meraki_networks_sm_trusted_access_configs_example" {
  value = data.meraki_networks_sm_trusted_access_configs.example.items
}
