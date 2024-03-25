
data "meraki_networks_clients_policy" "example" {

  client_id  = "string"
  network_id = "string"
}

output "meraki_networks_clients_policy_example" {
  value = data.meraki_networks_clients_policy.example.item
}
