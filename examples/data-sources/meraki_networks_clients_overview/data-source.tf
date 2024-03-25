
data "meraki_networks_clients_overview" "example" {

  network_id = "string"
  resolution = 1
  t0         = "string"
  t1         = "string"
  timespan   = 1.0
}

output "meraki_networks_clients_overview_example" {
  value = data.meraki_networks_clients_overview.example.item
}
