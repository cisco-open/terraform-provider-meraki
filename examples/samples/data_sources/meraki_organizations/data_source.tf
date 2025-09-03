terraform {
  required_providers {
    meraki = {
      version = "1.2.1-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}
data "meraki_organizations" "example" {
  organization_id = "828099381482762270"

}
output "meraki_organizations_example" {
  value = data.meraki_organizations.example.item
}