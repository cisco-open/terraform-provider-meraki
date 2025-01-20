
data "meraki_networks_wireless_electronic_shelf_label" "example" {

  network_id = "string"
}

output "meraki_networks_wireless_electronic_shelf_label_example" {
  value = data.meraki_networks_wireless_electronic_shelf_label.example.item
}
