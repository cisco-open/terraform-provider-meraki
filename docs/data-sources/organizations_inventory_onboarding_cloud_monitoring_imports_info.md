---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_inventory_onboarding_cloud_monitoring_imports_info Data Source - terraform-provider-meraki"
subcategory: "organizations"
description: |-
  
---

# meraki_organizations_inventory_onboarding_cloud_monitoring_imports_info (Data Source)



## Example Usage

```terraform
data "meraki_organizations_inventory_onboarding_cloud_monitoring_imports_info" "example" {

  import_ids      = ["string"]
  organization_id = "string"
}

output "meraki_organizations_inventory_onboarding_cloud_monitoring_imports_info_example" {
  value = data.meraki_organizations_inventory_onboarding_cloud_monitoring_imports_info.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `import_ids` (List of String) importIds query parameter. import ids from an imports
- `organization_id` (String) organizationId path parameter. Organization ID

### Read-Only

- `items` (Attributes List) Array of ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `device` (Attributes) Represents the details of an imported device. (see [below for nested schema](#nestedatt--items--device))
- `import_id` (String) Database ID for the new entity entry.

<a id="nestedatt--items--device"></a>
### Nested Schema for `items.device`

Read-Only:

- `created` (Boolean) Whether or not the device was successfully created in dashboard.
- `status` (String) Represents the current state of importing the device.
- `url` (String) The url to the device details page within dashboard.
