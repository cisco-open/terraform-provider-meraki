terraform {
  required_providers {
    meraki = {
      version = "0.1.0-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks" "example" {
  provider        = meraki
  organization_id = "828099381482762270"
  name            = "TEST_TERRAFORM2"
  product_types   = ["appliance", "switch", "wireless"]
  # , "switch", "wireless"
  time_zone = "America/Los_Angeles"
  notes     = "Additional description of the network2"
}
output "meraki_networks_example" {
  value = resource.meraki_networks.example
}
