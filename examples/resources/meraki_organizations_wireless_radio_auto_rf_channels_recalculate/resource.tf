
resource "meraki_organizations_wireless_radio_auto_rf_channels_recalculate" "example" {

  organization_id = "string"
  parameters = {

    network_ids = ["N_678910", "L_12345"]
  }
}

output "meraki_organizations_wireless_radio_auto_rf_channels_recalculate_example" {
  value = meraki_organizations_wireless_radio_auto_rf_channels_recalculate.example
}