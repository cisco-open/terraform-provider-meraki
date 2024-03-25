
resource "meraki_networks_clients_provision" "example" {

  network_id = "string"
  parameters = {

    clients = [{

      mac  = "00:11:22:33:44:55"
      name = "Miles's phone"
    }]
    device_policy   = "Group policy"
    group_policy_id = "101"
  }
}

output "meraki_networks_clients_provision_example" {
  value = meraki_networks_clients_provision.example
}