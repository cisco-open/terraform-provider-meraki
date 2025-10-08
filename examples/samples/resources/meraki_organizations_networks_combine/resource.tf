terraform {
  required_providers {
    meraki = {
      version = "1.2.4-beta"
      source  = "hashicorp.com/edu/meraki"
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_organizations_networks_combine" "example" {

  organization_id = "828099381482762270"
  parameters = {

    # enrollment_string = "hello-world"
    name              = "Long Island world"
    network_ids       = ["N_828099381482850162", "L_828099381482777316"]
  }
}

output "meraki_organizations_networks_combine_example" {
  value = meraki_organizations_networks_combine.example
}