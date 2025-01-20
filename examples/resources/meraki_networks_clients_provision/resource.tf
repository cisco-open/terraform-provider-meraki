
resource "meraki_networks_clients_provision" "example" {

  network_id = "string"
  parameters = {

    clients = [{

      mac  = "00:11:22:33:44:55"
      name = "Miles's phone"
    }]
    device_policy   = "Group policy"
    group_policy_id = "101"
    policies_by_security_appliance = {

      device_policy = "Normal"
    }
    policies_by_ssid = {

      0 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      1 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      10 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      11 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      12 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      13 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      14 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      2 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      3 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      4 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      5 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      6 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      7 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      8 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      9 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
    }
  }
}

output "meraki_networks_clients_provision_example" {
  value = meraki_networks_clients_provision.example
}