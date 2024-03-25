
resource "meraki_networks_appliance_single_lan" "example" {

  appliance_ip = "string"
  ipv6 = {

    enabled = false
    prefix_assignments = [{

      autonomous = false
      origin = {

        interfaces = ["string"]
        type       = "string"
      }
      static_appliance_ip6 = "string"
      static_prefix        = "string"
    }]
  }
  mandatory_dhcp = {

    enabled = false
  }
  network_id = "string"
  subnet     = "string"
}

output "meraki_networks_appliance_single_lan_example" {
  value = meraki_networks_appliance_single_lan.example
}