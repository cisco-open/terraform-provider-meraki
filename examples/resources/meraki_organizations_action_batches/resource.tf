
resource "meraki_organizations_action_batches" "example" {

  actions = [{

    operation = "create"
    resource  = "/devices/QXXX-XXXX-XXXX/switch/ports/3"
  }]
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
  confirmed       = true
  organization_id = "string"
  synchronous     = true
}

output "meraki_organizations_action_batches_example" {
  value = meraki_organizations_action_batches.example
}