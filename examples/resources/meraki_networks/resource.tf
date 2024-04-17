
resource "meraki_networks" "example" {

  name            = "Main Office"
  notes           = "Additional description of the network"
  organization_id = "string"
  product_types   = ["appliance", "switch", "wireless"]
  tags            = ["tag1", "tag2"]
  time_zone       = "America/Los_Angeles"
}

output "meraki_networks_example" {
  value = meraki_networks.example
}