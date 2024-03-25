
data "meraki_networks_appliance_security_intrusion" "example" {

  network_id = "string"
}

output "meraki_networks_appliance_security_intrusion_example" {
  value = data.meraki_networks_appliance_security_intrusion.example.item
}
