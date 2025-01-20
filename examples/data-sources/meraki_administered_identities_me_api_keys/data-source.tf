
data "meraki_administered_identities_me_api_keys" "example" {

}

output "meraki_administered_identities_me_api_keys_example" {
  value = data.meraki_administered_identities_me_api_keys.example.items
}
