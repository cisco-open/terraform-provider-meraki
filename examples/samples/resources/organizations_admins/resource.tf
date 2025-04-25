terraform {
  required_providers {
    meraki = {
      version = "1.1.2-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_organizations_admins" "example" {

  authentication_method = "Email"
  email                 = "devnetmerakiadmin23@yopmail.com"
  name                  = "DevNet Admin 345 test"
  org_access            = "full"
  organization_id       = "828099381482762270"
}

output "meraki_organizations_admins_example" {
  value = resource.meraki_organizations_admins.example
}