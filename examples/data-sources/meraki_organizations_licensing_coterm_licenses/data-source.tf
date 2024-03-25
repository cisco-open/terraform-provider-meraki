
data "meraki_organizations_licensing_coterm_licenses" "example" {

  ending_before   = "string"
  expired         = false
  invalidated     = false
  organization_id = "string"
  per_page        = 1
  starting_after  = "string"
}

output "meraki_organizations_licensing_coterm_licenses_example" {
  value = data.meraki_organizations_licensing_coterm_licenses.example.items
}
