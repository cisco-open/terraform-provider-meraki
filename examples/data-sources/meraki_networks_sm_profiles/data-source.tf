
data "meraki_networks_sm_profiles" "example" {

  network_id    = "string"
  payload_types = ["string"]
}

output "meraki_networks_sm_profiles_example" {
  value = data.meraki_networks_sm_profiles.example.items
}
