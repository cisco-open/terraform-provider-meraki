
data "meraki_networks_clients_splash_authorization_status" "example" {

  client_id  = "string"
  network_id = "string"
}

output "meraki_networks_clients_splash_authorization_status_example" {
  value = data.meraki_networks_clients_splash_authorization_status.example.item
}
