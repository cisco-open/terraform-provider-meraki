terraform {
  required_providers {
    meraki = {
      version = "1.1.5-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug    = "true"
  meraki_base_url = "http://localhost:3002"
}


resource "meraki_organizations_saml_roles" "example" {

  networks = [{

    access = "full"
    id     = "N_24329156"
  }]
  org_access      = "none"
  organization_id = "string"
  role            = "myrole"
  tags = [{

    access = "read-only"
    tag    = "west"
  }]
}

output "meraki_organizations_saml_roles_example" {
  value = meraki_organizations_saml_roles.example
}