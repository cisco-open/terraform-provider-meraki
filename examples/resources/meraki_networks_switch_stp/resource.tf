
resource "meraki_networks_switch_stp" "example" {

  network_id   = "string"
  rstp_enabled = true
  stp_bridge_priority = [{

    stacks          = ["789102", "123456", "129102"]
    stp_priority    = 4096
    switch_profiles = ["1098", "1099", "1100"]
    switches        = ["Q234-ABCD-0001", "Q234-ABCD-0002", "Q234-ABCD-0003"]
  }]
}

output "meraki_networks_switch_stp_example" {
  value = meraki_networks_switch_stp.example
}