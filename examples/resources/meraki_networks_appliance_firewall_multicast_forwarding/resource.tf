
resource "meraki_networks_appliance_firewall_multicast_forwarding" "example" {

  network_id = "string"
  parameters = {

    rules = [{

      address     = "224.0.0.1"
      description = "test"
      vlan_ids    = ["1"]
    }]
  }
}

output "meraki_networks_appliance_firewall_multicast_forwarding_example" {
  value = meraki_networks_appliance_firewall_multicast_forwarding.example
}