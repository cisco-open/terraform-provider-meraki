
data "meraki_organizations_switch_ports_overview" "example" {

  organization_id = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
}

output "meraki_organizations_switch_ports_overview_example" {
  value = data.meraki_organizations_switch_ports_overview.example.item
}
