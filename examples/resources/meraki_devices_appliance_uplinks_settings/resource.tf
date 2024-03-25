
resource "meraki_devices_appliance_uplinks_settings" "example" {

  interfaces = {

    wan1 = {

      enabled = true
      pppoe = {

        authentication = {

          enabled  = true
          password = "password"
          username = "username"
        }
        enabled = true
      }
      svis = {

        ipv4 = {

          address         = "9.10.11.10/16"
          assignment_mode = "static"
          gateway         = "13.14.15.16"
          nameservers = {

            addresses = ["1.2.3.4"]
          }
        }
        ipv6 = {

          address         = "1:2:3::4"
          assignment_mode = "static"
          gateway         = "1:2:3::5"
          nameservers = {

            addresses = ["1001:4860:4860::8888", "1001:4860:4860::8844"]
          }
        }
      }
      vlan_tagging = {

        enabled = true
        vlan_id = 1
      }
    }
    wan2 = {

      enabled = true
      pppoe = {

        authentication = {

          enabled  = true
          password = "password"
          username = "username"
        }
        enabled = true
      }
      svis = {

        ipv4 = {

          address         = "9.10.11.10/16"
          assignment_mode = "static"
          gateway         = "13.14.15.16"
          nameservers = {

            addresses = ["1.2.3.4"]
          }
        }
        ipv6 = {

          address         = "1:2:3::4"
          assignment_mode = "static"
          gateway         = "1:2:3::5"
          nameservers = {

            addresses = ["1001:4860:4860::8888", "1001:4860:4860::8844"]
          }
        }
      }
      vlan_tagging = {

        enabled = true
        vlan_id = 1
      }
    }
  }
  serial = "string"
}

output "meraki_devices_appliance_uplinks_settings_example" {
  value = meraki_devices_appliance_uplinks_settings.example
}