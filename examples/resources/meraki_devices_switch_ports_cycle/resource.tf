
resource "meraki_devices_switch_ports_cycle" "example" {

  serial = "string"
  parameters = {

    ports = ["1", "2-5", "1_MA-MOD-8X10G_1", "1_MA-MOD-8X10G_2-1_MA-MOD-8X10G_8"]
  }
}

output "meraki_devices_switch_ports_cycle_example" {
  value = meraki_devices_switch_ports_cycle.example
}