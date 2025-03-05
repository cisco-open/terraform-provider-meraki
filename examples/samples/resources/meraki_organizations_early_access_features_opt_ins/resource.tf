terraform {
  required_providers {
    meraki = {
      version = "1.0.5-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_organizations_early_access_features_opt_ins" "example" {

  limit_scope_to_networks = ["N_12345"]
  organization_id         = "828099381482762270"
  short_name              = "has_mx_no_nat_early_access"
  opt_in_id               = "828099381482925915"
}

output "meraki_organizations_early_access_features_opt_ins_example" {
  value = meraki_organizations_early_access_features_opt_ins.example
}