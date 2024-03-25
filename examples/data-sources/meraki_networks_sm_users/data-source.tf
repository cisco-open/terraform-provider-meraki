
data "meraki_networks_sm_users" "example" {

  emails     = ["string"]
  ids        = ["string"]
  network_id = "string"
  scope      = ["string"]
  usernames  = ["string"]
}

output "meraki_networks_sm_users_example" {
  value = data.meraki_networks_sm_users.example.items
}
