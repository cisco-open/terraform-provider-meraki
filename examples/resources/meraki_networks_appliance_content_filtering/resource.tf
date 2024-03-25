
resource "meraki_networks_appliance_content_filtering" "example" {

  allowed_url_patterns   = ["http://www.example.org", "http://help.com.au"]
  blocked_url_categories = ["meraki:contentFiltering/category/1", "meraki:contentFiltering/category/7"]
  blocked_url_patterns   = ["http://www.example.com", "http://www.betting.com"]
  network_id             = "string"
  url_category_list_size = "topSites"
}

output "meraki_networks_appliance_content_filtering_example" {
  value = meraki_networks_appliance_content_filtering.example
}