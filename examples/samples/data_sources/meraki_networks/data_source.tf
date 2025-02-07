terraform {
  required_providers {
    meraki = {
      version = "1.0.1-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}
data "meraki_networks" "example" {

  organization_id = "828099381482762270"
  # organization_id = "828099381482762270"
  # network_id = "N_828099381482850169"

}
output "meraki_networks_example" {
  value = data.meraki_networks.example.items
}