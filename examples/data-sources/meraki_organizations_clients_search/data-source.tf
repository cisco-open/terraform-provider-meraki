
data "meraki_organizations_clients_search" "example" {

  ending_before   = "string"
  mac             = "string"
  organization_id = "string"
  per_page        = 1
  starting_after  = "string"
}

output "meraki_organizations_clients_search_example" {
  value = data.meraki_organizations_clients_search.example.item
}
