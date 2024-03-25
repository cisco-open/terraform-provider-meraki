
data "meraki_networks_switch_stp" "example" {

  network_id = "string"
}

output "meraki_networks_switch_stp_example" {
  value = data.meraki_networks_switch_stp.example.item
}
