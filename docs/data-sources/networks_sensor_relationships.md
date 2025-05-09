---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_networks_sensor_relationships Data Source - terraform-provider-meraki"
subcategory: "sensor"
description: |-
  
---

# meraki_networks_sensor_relationships (Data Source)



## Example Usage

```terraform
data "meraki_networks_sensor_relationships" "example" {

  network_id = "string"
}

output "meraki_networks_sensor_relationships_example" {
  value = data.meraki_networks_sensor_relationships.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) networkId path parameter. Network ID

### Read-Only

- `items` (Attributes List) Array of ResponseSensorGetNetworkSensorRelationships (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `device` (Attributes) A sensor or gateway device in the network (see [below for nested schema](#nestedatt--items--device))
- `relationships` (Attributes) An object describing the relationships defined between the device and other devices (see [below for nested schema](#nestedatt--items--relationships))

<a id="nestedatt--items--device"></a>
### Nested Schema for `items.device`

Read-Only:

- `name` (String) The name of the device
- `product_type` (String) The product type of the device
- `serial` (String) The serial of the device


<a id="nestedatt--items--relationships"></a>
### Nested Schema for `items.relationships`

Read-Only:

- `livestream` (Attributes) A role defined between an MT sensor and an MV camera that adds the camera's livestream to the sensor's details page. Snapshots from the camera will also appear in alert notifications that the sensor triggers. (see [below for nested schema](#nestedatt--items--relationships--livestream))

<a id="nestedatt--items--relationships--livestream"></a>
### Nested Schema for `items.relationships.livestream`

Read-Only:

- `related_devices` (Attributes Set) An array of the related devices for the role (see [below for nested schema](#nestedatt--items--relationships--livestream--related_devices))

<a id="nestedatt--items--relationships--livestream--related_devices"></a>
### Nested Schema for `items.relationships.livestream.related_devices`

Read-Only:

- `product_type` (String) The product type of the related device
- `serial` (String) The serial of the related device
