
data "meraki_networks_sm_users_softwares" "example" {

  network_id = "string"
  user_id    = "string"
}

output "meraki_networks_sm_users_softwares_example" {
  value = data.meraki_networks_sm_users_softwares.example.items
}
