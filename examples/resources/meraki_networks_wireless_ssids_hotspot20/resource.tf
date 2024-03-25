
resource "meraki_networks_wireless_ssids_hotspot20" "example" {

  domains = ["meraki.local", "domain2.com"]
  enabled = true
  mcc_mncs = [{

    mcc = "123"
    mnc = "456"
  }]
  nai_realms = [{

    format = "1"
    methods = [{

      authentication_types = {

        eapinner_authentication     = ["EAP-TTLS with MSCHAPv2"]
        non_eapinner_authentication = ["MSCHAP"]
      }
      id = "1"
    }]
  }]
  network_access_type = "Private network"
  network_id          = "string"
  number              = "string"
  operator = {

    name = "Meraki Product Management"
  }
  roam_consort_ois = ["ABC123", "456EFG"]
  venue = {

    name = "SF Branch"
    type = "Unspecified Assembly"
  }
}

output "meraki_networks_wireless_ssids_hotspot20_example" {
  value = meraki_networks_wireless_ssids_hotspot20.example
}