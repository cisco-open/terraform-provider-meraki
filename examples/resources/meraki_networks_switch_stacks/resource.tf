
resource "meraki_networks_switch_stacks" "example" {

  name       = "A cool stack"
  network_id = "string"
  serials    = ["QBZY-XWVU-TSRQ", "QBAB-CDEF-GHIJ"]
}

output "meraki_networks_switch_stacks_example" {
  value = meraki_networks_switch_stacks.example
}