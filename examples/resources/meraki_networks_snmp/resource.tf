
resource "meraki_networks_snmp" "example" {

  access     = "users"
  network_id = "string"
  users = [{

    passphrase = "hunter2"
    username   = "AzureDiamond"
  }]
}

output "meraki_networks_snmp_example" {
  value = meraki_networks_snmp.example
}