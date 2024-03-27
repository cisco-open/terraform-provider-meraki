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
data "meraki_administered_identities_me" "example" {
  count = 1000
  provider = meraki
}
# output "meraki_administered_identities_me_example" {
#   value = data.meraki_administered_identities_me.example.item
# }