
resource "meraki_networks_meraki_auth_users" "example" {

  account_type = "802.1X"
  authorizations = [{

    expires_at  = "2018-03-13T00:00:00.090210Z"
    ssid_number = 1
  }]
  email                  = "miles@meraki.com"
  email_password_to_user = false
  is_admin               = false
  name                   = "Miles Meraki"
  network_id             = "string"
  password               = "secret"
}

output "meraki_networks_meraki_auth_users_example" {
  value = meraki_networks_meraki_auth_users.example
}