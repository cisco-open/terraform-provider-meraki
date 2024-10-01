terraform {
  required_providers {
    meraki = {
      version = "0.2.12-alpha"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_organizations_adaptive_policy_groups" "repro" {
  organization_id = "828099381482762270"
  description     = "A repro"
  name            = "pulumi testing 2"
  sgt             = 1006
}