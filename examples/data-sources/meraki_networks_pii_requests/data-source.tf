
data "meraki_networks_pii_requests" "example" {

  network_id = "string"
}

output "meraki_networks_pii_requests_example" {
  value = data.meraki_networks_pii_requests.example.items
}

data "meraki_networks_pii_requests" "example" {

  network_id = "string"
}

output "meraki_networks_pii_requests_example" {
  value = data.meraki_networks_pii_requests.example.item
}
