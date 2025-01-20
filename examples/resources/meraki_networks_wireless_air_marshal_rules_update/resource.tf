
resource "meraki_networks_wireless_air_marshal_rules_update" "example" {

  network_id = "string"
  rule_id    = "string"
  parameters = {

    match = {

      string = "00:11:22:33:44:55"
      type   = "bssid"
    }
    type = "allow"
  }
}

output "meraki_networks_wireless_air_marshal_rules_update_example" {
  value = meraki_networks_wireless_air_marshal_rules_update.example
}