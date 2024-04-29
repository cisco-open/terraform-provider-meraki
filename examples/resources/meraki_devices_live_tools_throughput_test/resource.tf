
resource "meraki_devices_live_tools_throughput_test" "example" {

  callback = {

    http_server = {

      id = "aHR0cHM6Ly93d3cuZXhhbXBsZS5jb20vd2ViaG9va3M="
    }
    payload_template = {

      id = "wpt_2100"
    }
    shared_secret = "secret"
    url           = "https://webhook.site/28efa24e-f830-4d9f-a12b-fbb9e5035031"
  }
  serial = "string"
}

output "meraki_devices_live_tools_throughput_test_example" {
  value = meraki_devices_live_tools_throughput_test.example
}