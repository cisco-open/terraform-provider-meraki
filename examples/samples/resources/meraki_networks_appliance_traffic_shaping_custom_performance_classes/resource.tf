terraform {
  required_providers {
    meraki = {
      source  = "hashicorp.com/edu/meraki"
      version = "0.2.10-alpha"
    }
  }

  required_version = ">= 1.2.0"
}

provider "meraki" {
  meraki_debug = "true"
}


resource "meraki_networks_appliance_traffic_shaping_custom_performance_classes" "example" {

  network_id = "L_828099381482775486"
  parameters = {

    max_jitter          = 100
    max_latency         = 100
    max_loss_percentage = 5
    name                = "myCustomPerformanceClass"
  }
}

output "meraki_networks_appliance_traffic_shaping_custom_performance_classes_example" {
  value = meraki_networks_appliance_traffic_shaping_custom_performance_classes.example
}