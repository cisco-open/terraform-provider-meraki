
data "meraki_networks_netflow" "example" {

  network_id = "string"
}

output "meraki_networks_netflow_example" {
  value = data.meraki_networks_netflow.example.item
}
