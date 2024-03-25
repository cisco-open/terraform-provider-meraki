
resource "meraki_organizations_login_security" "example" {

  account_lockout_attempts = 3
  api_authentication = {

    ip_restrictions_for_keys = {

      enabled = true
      ranges  = ["192.195.83.1", "192.168.33.33"]
    }
  }
  enforce_account_lockout     = true
  enforce_different_passwords = true
  enforce_idle_timeout        = true
  enforce_login_ip_ranges     = true
  enforce_password_expiration = true
  enforce_strong_passwords    = true
  enforce_two_factor_auth     = true
  idle_timeout_minutes        = 30
  login_ip_ranges             = ["192.195.83.1", "192.195.83.255"]
  num_different_passwords     = 3
  organization_id             = "string"
  password_expiration_days    = 90
}

output "meraki_organizations_login_security_example" {
  value = meraki_organizations_login_security.example
}