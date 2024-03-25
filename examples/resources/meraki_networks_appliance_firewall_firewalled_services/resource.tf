
resource "meraki_networks_appliance_firewall_firewalled_services" "example" {

  access      = "restricted"
  allowed_ips = ["123.123.123.1"]
  network_id  = "string"
  service     = "string"
}

output "meraki_networks_appliance_firewall_firewalled_services_example" {
  value = meraki_networks_appliance_firewall_firewalled_services.example
}