
data "meraki_networks_snmp" "example" {

  network_id = "string"
}

output "meraki_networks_snmp_example" {
  value = data.meraki_networks_snmp.example.item
}
