
data "meraki_devices_live_tools_wake_on_lan" "example" {

  serial         = "string"
  wake_on_lan_id = "string"
}

output "meraki_devices_live_tools_wake_on_lan_example" {
  value = data.meraki_devices_live_tools_wake_on_lan.example.item
}
