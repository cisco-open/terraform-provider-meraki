
data "meraki_organizations_assurance_alerts_overview" "example" {

  active                            = false
  category                          = "string"
  device_tags                       = ["string"]
  device_types                      = ["string"]
  dismissed                         = false
  network_id                        = "string"
  organization_id                   = "string"
  resolved                          = false
  serials                           = ["string"]
  severity                          = "string"
  suppress_alerts_for_offline_nodes = false
  ts_end                            = "string"
  ts_start                          = "string"
  types                             = ["string"]
}

output "meraki_organizations_assurance_alerts_overview_example" {
  value = data.meraki_organizations_assurance_alerts_overview.example.item
}
