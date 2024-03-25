
resource "meraki_networks_sm_devices_modify_tags" "example" {

  network_id = "string"
  parameters = {

    ids           = ["1284392014819", "2983092129865"]
    scope         = ["withAny, old_tag"]
    serials       = ["Q234-ABCD-0001", "Q234-ABCD-0002", "Q234-ABCD-0003"]
    tags          = ["tag1", "tag2"]
    update_action = "add"
    wifi_macs     = ["00:11:22:33:44:55"]
  }
}

output "meraki_networks_sm_devices_modify_tags_example" {
  value = meraki_networks_sm_devices_modify_tags.example
}