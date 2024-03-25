
resource "meraki_networks_cellular_gateway_dhcp" "example" {

  dhcp_lease_time        = "1 hour"
  dns_custom_nameservers = ["172.16.2.111", "172.16.2.30"]
  dns_nameservers        = "custom"
  network_id             = "string"
}

output "meraki_networks_cellular_gateway_dhcp_example" {
  value = meraki_networks_cellular_gateway_dhcp.example
}