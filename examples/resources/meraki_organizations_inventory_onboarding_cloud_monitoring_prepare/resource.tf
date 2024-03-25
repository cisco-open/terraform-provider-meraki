
resource "meraki_organizations_inventory_onboarding_cloud_monitoring_prepare" "example" {

  organization_id = "string"
  parameters = {

    devices = [{

      sudi = <<EOT
          "-----BEGIN CERTIFICATE-----
        MIIDyTCCArGgAwIBAgIKBBNXOVCGU1YztjANBgkqhkiG9w0BAQsFADAnMQ4wDAYD
        VQQKEwVDaXNjbzEVMBMGA1UEAxMMQUNUMiBTVURJIENBMB4XDTIxMDUzMTEzNTUx
        NVoXDTI5MDUxNDIwMjU0MVowbTEpMCcGA1UEBRMgUElEOkM5MjAwTC0yNFAtNEcg
        U046SkFFMjUyMjBSMksxDjAMBgNVBAoTBUNpc2NvMRgwFgYDVQQLEw9BQ1QtMiBM
        aXRlIFNVREkxFjAUBgNVBAMTDUM5MjAwTC0yNFAtNEcwggEiMA0GCSqGSIb3DQEB
        AQUAA4IBDwAwggEKAoIBAQDaUPxW76gT5MdoEAt+UrDFiYA9RYh2iHicDViBEyow
        TR1TuP36bHh13X3vtGiDsCD88Ci2TZIqd/EDkkc7v9ipUUYVVH+YDrPt2Aukb1PH
        D6K0R+KhgEzRo5x54TlU6oWvjUpwNZUwwdhMWIQaUVkMyZBYNy0jGPLO8jwZhyBg
        1Fneybr9pwedGbLrAaz+gdEikB8B4a/fvPjVfL5Ngb4QRjFqWuE+X3nLc0kHedep
        6nfgpUNXMlStVm5nIXKP6OjmzfCHPYh9L2Ehs1TrSk1ser9Ofx0ZMVL/jBZR2EIj
        OZ8tH6KlX2/B2pbSPIO6kD5c4UA8Cf1SbDJCwJ/kI9ihAgMBAAGjgbAwga0wDgYD
        VR0PAQH/BAQDAgXgMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgwFoAUSNjx8cJw1Vu7
        fHMJk6+4uDAD+H8wTQYDVR0RBEYwRKBCBgkrBgEEAQkVAgOgNRMzQ2hpcElEPVVV
        VUNNaElGcUVFMklFUUVBQWNBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUE9MB0GA1Ud
        DgQWBBRdhMkFD/z5hokaQeLbaRsp4hkvbzANBgkqhkiG9w0BAQsFAAOCAQEAMtuh
        YpBz4xEZ7YdJsLpw67Q0TTJGnTBRpzAeY1urYDoDz8TSx556XG7z3IRzuED5KVSp
        OwmH/iZ+tDfYQ3W3ElWTW93871DkuW4WQIfbnoHg/F7bF0DKYVkD3rpZjyz3NhzH
        d7cjTdJXQ85bTAOXDuxKH3qewrXxxOGXgh3I6NUq0UwMTWh84lND7Jl+ZAQkYNS2
        iHanTZFQBk3ML0NUb7fKDYGRTZRqwQ/upIO4S6LV1cxH/6V0qbMy3sCSHZoMLrW3
        0m3M6yKpe5+VZzHZwmWdUf3Ot+zKjhveK5/YNsMIASdvtvymxUizq2Hr1hvR/kPc
        p1vuyWxipU8JfzOh/A==
        -----END CERTIFICATE-----
        "
          EOT
      tunnel = {

        certificate_name = "DeviceSUDI"
        local_interface  = 1
        loopback_number  = 1000
        name             = "MERAKI"
      }
      user = {

        username = "Meraki"
      }
      vty = {

        access_list = {

          vty_in = {

            name = "MERAKI_IN"
          }
          vty_out = {

            name = "MERAKI_OUT"
          }
        }
        authentication = {

          group = {

            name = ""
          }
        }
        authorization = {

          group = {

            name = "MERAKI"
          }
        }
        end_line_number   = 17
        rotary_number     = 50
        start_line_number = 16
      }
    }]
  }
}

output "meraki_organizations_inventory_onboarding_cloud_monitoring_prepare_example" {
  value = meraki_organizations_inventory_onboarding_cloud_monitoring_prepare.example
}