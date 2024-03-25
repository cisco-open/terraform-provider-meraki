
data "meraki_networks_bluetooth_clients" "example" {

  bluetooth_client_id           = "string"
  connectivity_history_timespan = 1
  include_connectivity_history  = false
  network_id                    = "string"
}

output "meraki_networks_bluetooth_clients_example" {
  value = data.meraki_networks_bluetooth_clients.example.item
}
