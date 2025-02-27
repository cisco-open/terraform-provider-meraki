
resource "meraki_organizations_appliance_vpn_third_party_vpnpeers" "example" {

  organization_id = "string"
  peers = [{

    ike_version = "2"
    ipsec_policies = {

      child_auth_algo          = ["sha1"]
      child_cipher_algo        = ["aes128"]
      child_lifetime           = 28800
      child_pfs_group          = ["disabled"]
      ike_auth_algo            = ["sha1"]
      ike_cipher_algo          = ["tripledes"]
      ike_diffie_hellman_group = ["group2"]
      ike_lifetime             = 28800
      ike_prf_algo             = ["prfsha1"]
    }
    ipsec_policies_preset = "default"
    local_id              = "myMXId@meraki.com"
    name                  = "Peer Name"
    network_tags          = ["none"]
    private_subnets       = ["192.168.1.0/24", "192.168.128.0/24"]
    public_hostname       = "example.com"
    public_ip             = "123.123.123.1"
    remote_id             = "miles@meraki.com"
    secret                = "Sample Password"
  }]
}

output "meraki_organizations_appliance_vpn_third_party_vpnpeers_example" {
  value = meraki_organizations_appliance_vpn_third_party_vpnpeers.example
}