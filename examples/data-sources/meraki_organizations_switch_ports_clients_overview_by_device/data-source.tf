
data "meraki_organizations_switch_ports_clients_overview_by_device" "example" {

  configuration_updated_after = "string"
  ending_before               = "string"
  mac                         = "string"
  macs                        = ["string"]
  name                        = "string"
  network_ids                 = ["string"]
  organization_id             = "string"
  per_page                    = 1
  port_profile_ids            = ["string"]
  serial                      = "string"
  serials                     = ["string"]
  starting_after              = "string"
  t0                          = "string"
  timespan                    = 1.0
}

output "meraki_organizations_switch_ports_clients_overview_by_device_example" {
  value = data.meraki_organizations_switch_ports_clients_overview_by_device.example.item
}
