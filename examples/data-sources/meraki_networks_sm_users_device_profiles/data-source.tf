
data "meraki_networks_sm_users_device_profiles" "example" {

  network_id = "string"
  user_id    = "string"
}

output "meraki_networks_sm_users_device_profiles_example" {
  value = data.meraki_networks_sm_users_device_profiles.example.items
}
