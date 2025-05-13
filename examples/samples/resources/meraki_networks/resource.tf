terraform {
  required_providers {
    meraki = {
      version = "1.1.3-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks" "example" {
  name            = "My Test Network"
  organization_id = "828099381482762270"
  notes           = "This is a test network created by my team by terraform"
  product_types   = ["appliance", "switch", "wireless", "cellularGateway"]
  time_zone       = "America/Los_Angeles"
}
output "meraki_networks_example" {
  value = resource.meraki_networks.example
}
