
resource "meraki_networks_netflow" "example" {

  collector_ip      = "1.2.3.4"
  collector_port    = 443
  eta_dst_port      = 443
  eta_enabled       = true
  network_id        = "string"
  reporting_enabled = true
}

output "meraki_networks_netflow_example" {
  value = meraki_networks_netflow.example
}