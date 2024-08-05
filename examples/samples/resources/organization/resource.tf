terraform {
  required_providers {
    meraki = {
      version = "0.2.10-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_organizations" "example" {



  api = {

    enabled = true
  }
  management = {

    details = [{

      name  = "MSP ID"
      value = "12345678"
    }]
  }
  name = "Test Terraform2"
  #   organization_id = "575334852396583071"

}

output "meraki_organizations_example" {
  value = meraki_organizations.example
}