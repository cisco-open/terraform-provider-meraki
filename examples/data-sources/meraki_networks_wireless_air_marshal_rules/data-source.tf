
data "meraki_networks_wireless_air_marshal_rules" "example" {

  network_id = "string"
  t0         = "string"
  timespan   = 1.0
}

output "meraki_networks_wireless_air_marshal_rules_example" {
  value = data.meraki_networks_wireless_air_marshal_rules.example.items
}
