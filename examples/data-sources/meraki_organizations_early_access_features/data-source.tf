
data "meraki_organizations_early_access_features" "example" {

  organization_id = "string"
}

output "meraki_organizations_early_access_features_example" {
  value = data.meraki_organizations_early_access_features.example.items
}
