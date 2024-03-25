
data "meraki_organizations_early_access_features_opt_ins" "example" {

  organization_id = "string"
}

output "meraki_organizations_early_access_features_opt_ins_example" {
  value = data.meraki_organizations_early_access_features_opt_ins.example.items
}

data "meraki_organizations_early_access_features_opt_ins" "example" {

  organization_id = "string"
}

output "meraki_organizations_early_access_features_opt_ins_example" {
  value = data.meraki_organizations_early_access_features_opt_ins.example.item
}
