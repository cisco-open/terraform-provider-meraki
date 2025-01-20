
resource "meraki_networks_switch_mtu" "example" {

  default_mtu_size = 9578
  network_id       = "string"
  overrides = [{

    mtu_size        = 1500
    switch_profiles = ["1284392014819", "2983092129865"]
    switches        = ["Q234-ABCD-0001", "Q234-ABCD-0002", "Q234-ABCD-0003"]
  }]
}

output "meraki_networks_switch_mtu_example" {
  value = meraki_networks_switch_mtu.example
}