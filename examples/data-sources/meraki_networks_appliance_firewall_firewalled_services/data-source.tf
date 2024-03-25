
data "meraki_networks_appliance_firewall_firewalled_services" "example" {

  network_id = "string"
  service    = "string"
}

output "meraki_networks_appliance_firewall_firewalled_services_example" {
  value = data.meraki_networks_appliance_firewall_firewalled_services.example.item
}
