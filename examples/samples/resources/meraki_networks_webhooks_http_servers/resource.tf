terraform {
  required_providers {
    meraki = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

provider "meraki" {
  meraki_debug = "true"
}

resource "meraki_networks_webhooks_http_servers" "webhook_http_server" {
  name          = "Pulumi Test Webhook 2"
  url           = "https://example.com/test"
  network_id    = "L_828099381482775375"
  shared_secret = "redactedSecret"
  payload_template = {
    name                = "Meraki (included)"
    payload_template_id = "wpt_00001"
  }
}

