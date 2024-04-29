
resource "meraki_devices_live_tools_wake_on_lan" "example" {

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
  mac     = "00:11:22:33:44:55"
  serial  = "string"
  vlan_id = 12
}

output "meraki_devices_live_tools_wake_on_lan_example" {
  value = meraki_devices_live_tools_wake_on_lan.example
}