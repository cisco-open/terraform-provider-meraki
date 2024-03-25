
resource "meraki_networks_wireless_ssids_vpn" "example" {

  concentrator = {

    network_id = "N_123"
    vlan_id    = 44
  }
  failover = {

    heartbeat_interval = 10
    idle_timeout       = 30
    request_ip         = "1.1.1.1"
  }
  network_id = "string"
  number     = "string"
  split_tunnel = {

    enabled = true
    rules = [{

      comment   = "split tunnel rule 1"
      dest_cidr = "1.1.1.1/32"
      dest_port = "any"
      policy    = "allow"
      protocol  = "Any"
    }]
  }
}

output "meraki_networks_wireless_ssids_vpn_example" {
  value = meraki_networks_wireless_ssids_vpn.example
}