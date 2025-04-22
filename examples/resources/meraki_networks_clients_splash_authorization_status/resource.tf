
resource "meraki_networks_clients_splash_authorization_status" "example" {

  client_id  = "string"
  network_id = "string"
  ssids = {

    status_0 = {

      is_authorized = true
    }
    status_2 = {

      is_authorized = false
    }
  }
}

output "meraki_networks_clients_splash_authorization_status_example" {
  value = meraki_networks_clients_splash_authorization_status.example
}