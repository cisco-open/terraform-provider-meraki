
data "meraki_networks_clients" "example" {

  client_id  = "string"
  network_id = "string"
}

output "meraki_networks_clients_example" {
  value = data.meraki_networks_clients.example.item
}
