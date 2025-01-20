
data "meraki_organizations_assurance_alerts" "example" {

  active                            = false
  category                          = "string"
  device_tags                       = ["string"]
  device_types                      = ["string"]
  dismissed                         = false
  ending_before                     = "string"
  network_id                        = "string"
  organization_id                   = "string"
  per_page                          = 1
  resolved                          = false
  serials                           = ["string"]
  severity                          = "string"
  sort_by                           = "string"
  sort_order                        = "string"
  starting_after                    = "string"
  suppress_alerts_for_offline_nodes = false
  ts_end                            = "string"
  ts_start                          = "string"
  types                             = ["string"]
}

output "meraki_organizations_assurance_alerts_example" {
  value = data.meraki_organizations_assurance_alerts.example.items
}
