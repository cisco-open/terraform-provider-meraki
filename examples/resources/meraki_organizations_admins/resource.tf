
resource "meraki_organizations_admins" "example" {

  authentication_method = "Email"
  email                 = "miles@meraki.com"
  name                  = "Miles Meraki"
  networks = [{

    access = "full"
    id     = "N_24329156"
  }]
  org_access      = "none"
  organization_id = "string"
  tags = [{

    access = "read-only"
    tag    = "west"
  }]
}

output "meraki_organizations_admins_example" {
  value = meraki_organizations_admins.example
}