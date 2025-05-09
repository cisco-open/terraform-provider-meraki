
resource "meraki_organizations_early_access_features_opt_ins" "example" {

  limit_scope_to_networks = ["N_12345"]
  organization_id         = "string"
  short_name              = "has_beta_api"
}

output "meraki_organizations_early_access_features_opt_ins_example" {
  value = meraki_organizations_early_access_features_opt_ins.example
}