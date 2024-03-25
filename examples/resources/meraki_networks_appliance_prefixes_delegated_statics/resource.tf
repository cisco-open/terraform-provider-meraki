
resource "meraki_networks_appliance_prefixes_delegated_statics" "example" {

  description = "Prefix on WAN 1 of Long Island Office network"
  network_id  = "string"
  origin = {

    interfaces = ["wan1"]
    type       = "internet"
  }
  prefix = "2001:db8:3c4d:15::/64"
}

output "meraki_networks_appliance_prefixes_delegated_statics_example" {
  value = meraki_networks_appliance_prefixes_delegated_statics.example
}