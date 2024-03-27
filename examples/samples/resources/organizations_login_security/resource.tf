terraform {
  required_providers {
    meraki = {
      version = "0.1.0-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-open/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_organizations_login_security" "example" {
  provider = meraki


  account_lockout_attempts = 3
  enforce_idle_timeout     = false
  enforce_two_factor_auth  = false
  organization_id          = "828099381482762270"
}

output "meraki_organizations_login_security_example" {
  value = meraki_organizations_login_security.example
}