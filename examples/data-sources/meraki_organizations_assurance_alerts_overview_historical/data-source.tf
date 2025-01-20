
data "meraki_organizations_assurance_alerts_overview_historical" "example" {

  category         = "string"
  device_types     = ["string"]
  network_id       = "string"
  organization_id  = "string"
  segment_duration = 1
  serials          = ["string"]
  severity         = "string"
  ts_end           = "string"
  ts_start         = "string"
  types            = ["string"]
}

output "meraki_organizations_assurance_alerts_overview_historical_example" {
  value = data.meraki_organizations_assurance_alerts_overview_historical.example.item
}
