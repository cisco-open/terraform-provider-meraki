
resource "meraki_networks_floor_plans" "example" {

  bottom_left_corner = {

    lat = 37.770040510499996
    lng = -122.38714009525
  }
  bottom_right_corner = {

    lat = 37.770040510499996
    lng = -122.38714009525
  }
  center = {

    lat = 37.770040510499996
    lng = -122.38714009525
  }
  floor_number   = 5.0
  image_contents = "2a9edd3f4ffd80130c647d13eacb59f3"
  name           = "HQ Floor Plan"
  network_id     = "string"
  top_left_corner = {

    lat = 37.770040510499996
    lng = -122.38714009525
  }
  top_right_corner = {

    lat = 37.770040510499996
    lng = -122.38714009525
  }
}

output "meraki_networks_floor_plans_example" {
  value = meraki_networks_floor_plans.example
}