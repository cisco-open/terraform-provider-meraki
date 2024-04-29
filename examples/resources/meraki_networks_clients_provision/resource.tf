
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

      status_0 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      status_1 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      status_10 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      status_11 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      status_12 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      status_13 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      status_14 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      status_2 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      status_3 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      status_4 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      status_5 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      status_6 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      status_7 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      status_8 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
      status_9 = {

        device_policy   = "Group policy"
        group_policy_id = "101"
      }
    }
  }
}

output "meraki_networks_clients_provision_example" {
  value = meraki_networks_clients_provision.example
}