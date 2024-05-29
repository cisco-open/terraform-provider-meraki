terraform {
  required_providers {
    meraki = {
      source  = "hashicorp.com/edu/meraki"
      version = "0.2.2-alpha"
    }
  }

  required_version = ">= 1.2.0"
}

provider meraki {
    meraki_debug = "true"
}

data "meraki_organizations" "example" {}

data "meraki_networks" "example" {
  organization_id = data.meraki_organizations.example.items[0].id
}

resource "meraki_networks_appliance_security_intrusion" "example" {

  ids_rulesets = "balanced"
  mode         = "prevention"
  network_id   = data.meraki_networks.example.items[0].id
}

output "meraki_networks_appliance_security_intrusion_example" {
  value = meraki_networks_appliance_security_intrusion.example
}