
resource "meraki_organizations_early_access_features_opt_ins" "example" {

  limit_scope_to_networks = ["N_12345"]
  organization_id         = "string"
  short_name              = "has_magnetic_beta"
  opt_in_id               = "828099381482925914"
}

output "meraki_organizations_early_access_features_opt_ins_example" {
  value = meraki_organizations_early_access_features_opt_ins.example
}