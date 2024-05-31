terraform {
  required_providers {
    meraki = {
      version = "0.2.3-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
  # meraki_requests_per_second = 10
}
data "meraki_administered_identities_me" "example" {
  count = 10

}
# output "meraki_administered_identities_me_example" {
#   value = data.meraki_administered_identities_me.example.item
# }