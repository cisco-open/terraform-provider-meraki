terraform {
  required_providers {
    meraki = {
      version = "1.0.6-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_organizations_camera_custom_analytics_artifacts" "example" {

  name            = "Test Terraform4"
  organization_id = "828099381482762270"
}

output "meraki_organizations_camera_custom_analytics_artifacts_example" {
  value = resource.meraki_organizations_camera_custom_analytics_artifacts.example
}