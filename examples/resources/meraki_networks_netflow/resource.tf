
resource "meraki_networks_netflow" "example" {

  collector_ip      = "string"
  collector_port    = 1
  eta_dst_port      = 1
  eta_enabled       = false
  network_id        = "string"
  reporting_enabled = false
}

output "meraki_networks_netflow_example" {
  value = meraki_networks_netflow.example
}