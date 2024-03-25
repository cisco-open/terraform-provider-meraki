
resource "meraki_networks_clients_policy" "example" {

  client_id       = "string"
  device_policy   = "Group policy"
  group_policy_id = "101"
  network_id      = "string"
}

output "meraki_networks_clients_policy_example" {
  value = meraki_networks_clients_policy.example
}