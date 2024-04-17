
data "meraki_devices_live_tools_arp_table" "example" {

  arp_table_id = "string"
  serial       = "string"
}

output "meraki_devices_live_tools_arp_table_example" {
  value = data.meraki_devices_live_tools_arp_table.example.item
}
