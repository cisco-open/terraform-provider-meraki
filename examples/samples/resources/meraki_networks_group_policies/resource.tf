terraform {
  required_providers {
    meraki = {
      version = "1.0.2-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}
provider "meraki" {
  meraki_debug    = "true"
  meraki_base_url = "http://localhost:3000"
}


resource "meraki_networks_group_policies" "example" {

  bandwidth = {

    bandwidth_limits = {

      limit_down = 1000000
      limit_up   = 1000000
    }
    settings = "custom"
  }
  bonjour_forwarding = {

    rules = [{

      description = "A simple bonjour rule"
      services    = ["All Services"]
      vlan_id     = "1"
    }]
    settings = "custom"
  }
  # content_filtering = {

  #   allowed_url_patterns = {

  #     settings = "network default"
  #   }
  #   blocked_url_categories = {

  #     # categories = ["meraki:contentFiltering/category/1", "meraki:contentFiltering/category/7"]
  #     settings   = "override"
  #   }
  #   blocked_url_patterns = {

  #     patterns = ["http://www.example.com", "http://www.betting.com"]
  #     settings = "append"
  #   }
  # }
  firewall_and_traffic_shaping = {

    l3_firewall_rules = [{

      comment   = "Allow TCP traffic to subnet with HTTP servers."
      dest_cidr = "192.168.1.0/24"
      dest_port = "443"
      policy    = "allow"
      protocol  = "tcp"
    }]
    l7_firewall_rules = [{

      policy = "deny"
      type   = "host"
      value  = "google.com"
    }]
    settings = "custom"
    traffic_shaping_rules = [{

      definitions = [{

        type  = "host"
        value = "google.com"
      }]
      # dscp_tag_value = 1
      pcp_tag_value = 1
      per_client_bandwidth_limits = {

        bandwidth_limits = {

          limit_down = 1000000
          limit_up   = 1000000
        }
        settings = "custom"
      }
      priority = "normal"
    }]
  }
  name       = "IOT"
  network_id = "L_828099381482775375"
  scheduling = {

    enabled = true
    friday = {

      active = true
      from   = "9:00"
      to     = "17:00"
    }
    monday = {

      active = true
      from   = "9:00"
      to     = "17:00"
    }
    saturday = {

      active = true
      from   = "9:00"
      to     = "17:00"
    }
    sunday = {

      active = true
      from   = "9:00"
      to     = "17:00"
    }
    thursday = {

      active = true
      from   = "9:00"
      to     = "17:00"
    }
    tuesday = {

      active = true
      from   = "9:00"
      to     = "17:00"
    }
    wednesday = {

      active = true
      from   = "9:00"
      to     = "17:00"
    }
  }
  splash_auth_settings = "bypass"
  vlan_tagging = {

    settings = "custom"
    vlan_id  = "1"
  }
}
