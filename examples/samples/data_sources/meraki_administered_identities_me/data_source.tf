terraform {
  required_providers {
    meraki = {
      version = "1.1.6-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_user_agent = "Testing_aaa"
  meraki_base_url = "http://localhost:3006/api/v1"
  # meraki_requests_per_second = 10
}
data "meraki_administered_identities_me" "example" {
  # count = 10

}
output "meraki_administered_identities_me_example" {
  value = data.meraki_administered_identities_me.example.item
}