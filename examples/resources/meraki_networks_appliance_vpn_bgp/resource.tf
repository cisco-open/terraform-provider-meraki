
resource "meraki_networks_appliance_vpn_bgp" "example" {

  as_number       = 64515
  enabled         = true
  ibgp_hold_timer = 120
  neighbors = [{

    allow_transit = true
    authentication = {

      password = "abc123"
    }
    ebgp_hold_timer = 180
    ebgp_multihop   = 2
    ip              = "10.10.10.22"
    ipv6 = {

      address = "2002::1234:abcd:ffff:c0a8:101"
    }
    next_hop_ip      = "1.2.3.4"
    receive_limit    = 120
    remote_as_number = 64343
    source_interface = "wan1"
    ttl_security = {

      enabled = false
    }
  }]
  network_id = "string"
}

output "meraki_networks_appliance_vpn_bgp_example" {
  value = meraki_networks_appliance_vpn_bgp.example
}