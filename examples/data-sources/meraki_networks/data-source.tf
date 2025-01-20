
data "meraki_networks" "example" {

  config_template_id          = "string"
  ending_before               = "string"
  is_bound_to_config_template = false
  organization_id             = "string"
  per_page                    = 1
  product_types               = ["string"]
  starting_after              = "string"
  tags                        = ["string"]
  tags_filter_type            = "string"
}

output "meraki_networks_example" {
  value = data.meraki_networks.example.items
}

data "meraki_networks" "example" {

  network_id = "string"
}

output "meraki_networks_example" {
  value = data.meraki_networks.example.item
}
