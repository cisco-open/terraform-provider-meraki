terraform {
  required_providers {
    meraki = {
      version = "1.0.3-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_appliance_content_filtering" "example" {

  allowed_url_patterns   = ["http://www.example2.org", "http://help.com.au"]
  blocked_url_categories = []
  blocked_url_patterns   = ["http://www.example.com", "http://www.betting.com"]
  network_id             = "L_828099381482771185"
  url_category_list_size = "fullList"
}

output "meraki_networks_appliance_content_filtering_example" {
  value = meraki_networks_appliance_content_filtering.example
}