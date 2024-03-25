
resource "meraki_networks_switch_stacks_add" "example" {

  network_id      = "string"
  switch_stack_id = "string"
  parameters = {

    serial = "QBZY-XWVU-TSRQ"
  }
}

output "meraki_networks_switch_stacks_add_example" {
  value = meraki_networks_switch_stacks_add.example
}