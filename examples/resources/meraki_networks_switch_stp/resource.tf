
resource "meraki_networks_switch_stp" "example" {

  network_id   = "string"
  rstp_enabled = true
  stp_bridge_priority = [{

    stp_priority = 4096
    switches     = ["Q234-ABCD-0001", "Q234-ABCD-0002", "Q234-ABCD-0003"]
  }]
}

output "meraki_networks_switch_stp_example" {
  value = meraki_networks_switch_stp.example
}