
data "meraki_organizations_appliance_traffic_shaping_vpn_exclusions_by_network" "example" {

  ending_before   = "string"
  network_ids     = ["string"]
  organization_id = "string"
  per_page        = 1
  starting_after  = "string"
}

output "meraki_organizations_appliance_traffic_shaping_vpn_exclusions_by_network_example" {
  value = data.meraki_organizations_appliance_traffic_shaping_vpn_exclusions_by_network.example.item
}
