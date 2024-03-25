
resource "meraki_networks_appliance_vpn_bgp" "example" {

  as_number       = 64515
  enabled         = true
  ibgp_hold_timer = 120
  neighbors = [{

    allow_transit    = true
    ebgp_hold_timer  = 180
    ebgp_multihop    = 2
    ip               = "10.10.10.22"
    receive_limit    = 120
    remote_as_number = 64343
  }]
  network_id = "string"
}

output "meraki_networks_appliance_vpn_bgp_example" {
  value = meraki_networks_appliance_vpn_bgp.example
}