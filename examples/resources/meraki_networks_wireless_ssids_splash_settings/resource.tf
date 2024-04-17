
resource "meraki_networks_wireless_ssids_splash_settings" "example" {

  allow_simultaneous_logins = false
  billing = {

    free_access = {

      duration_in_minutes = 120
      enabled             = true
    }
    prepaid_access_fast_login_enabled = true
    reply_to_email_address            = "user@email.com"
  }
  block_all_traffic_before_sign_on  = false
  controller_disconnection_behavior = "default"
  guest_sponsorship = {

    duration_in_minutes         = 30
    guest_can_request_timeframe = false
  }
  network_id   = "string"
  number       = "string"
  redirect_url = "https://example.com"
  sentry_enrollment = {

    enforced_systems = ["iOS"]
    strength         = "focused"
    systems_manager_network = {

      id = "N_1234"
    }
  }
  splash_image = {

    extension = "jpg"
    image = {

      contents = "Q2lzY28gTWVyYWtp"
      format   = "jpg"
    }
    md5 = "542cccac8d7dedee0f185311d154d194"
  }
  splash_logo = {

    extension = "jpg"
    image = {

      contents = "Q2lzY28gTWVyYWtp"
      format   = "jpg"
    }
    md5 = "abcd1234"
  }
  splash_prepaid_front = {

    extension = "jpg"
    image = {

      contents = "Q2lzY28gTWVyYWtp"
      format   = "jpg"
    }
    md5 = "542cccac8d7dedee0f185311d154d194"
  }
  splash_timeout   = 1440
  splash_url       = "https://www.custom_splash_url.com"
  theme_id         = "c3ddcb4f16785ee747ab5ffc10867d6c8ea704be"
  use_redirect_url = true
  use_splash_url   = true
  welcome_message  = "Welcome!"
}

output "meraki_networks_wireless_ssids_splash_settings_example" {
  value = meraki_networks_wireless_ssids_splash_settings.example
}