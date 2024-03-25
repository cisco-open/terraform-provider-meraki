
data "meraki_networks_meraki_auth_users" "example" {

  network_id = "string"
}

output "meraki_networks_meraki_auth_users_example" {
  value = data.meraki_networks_meraki_auth_users.example.items
}

data "meraki_networks_meraki_auth_users" "example" {

  network_id = "string"
}

output "meraki_networks_meraki_auth_users_example" {
  value = data.meraki_networks_meraki_auth_users.example.item
}
