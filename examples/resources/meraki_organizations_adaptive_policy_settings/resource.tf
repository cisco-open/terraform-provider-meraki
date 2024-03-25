
resource "meraki_organizations_adaptive_policy_settings" "example" {

  enabled_networks = ["L_11111111", "L_22222222", "N_33333333", "L_44444444"]
  organization_id  = "string"
}

output "meraki_organizations_adaptive_policy_settings_example" {
  value = meraki_organizations_adaptive_policy_settings.example
}