
resource "meraki_networks_publish" "example" {

  job_id     = "string"
  network_id = "string"
  parameters = {

    devices = [{

      auto_locate = {

        is_anchor = true
      }
      lat    = 37.4180951010362
      lng    = -122.098531723022
      serial = "Q234-ABCD-5678"
    }]
  }
}

output "meraki_networks_publish_example" {
  value = meraki_networks_publish.example
}