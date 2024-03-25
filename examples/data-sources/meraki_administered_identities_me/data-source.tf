
data "meraki_administered_identities_me" "example" {

}

output "meraki_administered_identities_me_example" {
  value = data.meraki_administered_identities_me.example.item
}
