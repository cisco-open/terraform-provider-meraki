
resource "meraki_networks_switch_routing_ospf" "example" {

  areas = [{

    area_id   = "1284392014819"
    area_name = "Backbone"
    area_type = "normal"
  }]
  dead_timer_in_seconds      = 40
  enabled                    = true
  hello_timer_in_seconds     = 10
  md5_authentication_enabled = true
  md5_authentication_key = {

    id         = 1234
    passphrase = "abc1234"
  }
  network_id = "string"
  v3 = {

    areas = [{

      area_id   = "1284392014819"
      area_name = "V3 Backbone"
      area_type = "normal"
    }]
    dead_timer_in_seconds  = 40
    enabled                = true
    hello_timer_in_seconds = 10
  }
}

output "meraki_networks_switch_routing_ospf_example" {
  value = meraki_networks_switch_routing_ospf.example
}