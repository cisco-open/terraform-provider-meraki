---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_summary_top_appliances_by_utilization Data Source - terraform-provider-meraki"
subcategory: ""
description: |-
  
---

# meraki_organizations_summary_top_appliances_by_utilization (Data Source)



## Example Usage

```terraform
data "meraki_organizations_summary_top_appliances_by_utilization" "example" {

  organization_id = "string"
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
}

output "meraki_organizations_summary_top_appliances_by_utilization_example" {
  value = data.meraki_organizations_summary_top_appliances_by_utilization.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `t0` (String) t0 query parameter. The beginning of the timespan for the data.
- `t1` (String) t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
- `timespan` (Number) timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 25 minutes and be less than or equal to 31 days. The default is 1 day.

### Read-Only

- `items` (Attributes List) Array of ResponseOrganizationsGetOrganizationSummaryTopAppliancesByUtilization (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `mac` (String) Mac address of the appliance
- `model` (String) Model of the appliance
- `name` (String) Name of the appliance
- `network` (Attributes) Network info (see [below for nested schema](#nestedatt--items--network))
- `serial` (String) Serial number of the appliance
- `utilization` (Attributes) Utilization of the appliance (see [below for nested schema](#nestedatt--items--utilization))

<a id="nestedatt--items--network"></a>
### Nested Schema for `items.network`

Read-Only:

- `id` (String) Network id
- `name` (String) Network name


<a id="nestedatt--items--utilization"></a>
### Nested Schema for `items.utilization`

Read-Only:

- `average` (Attributes) Average utilization of the appliance (see [below for nested schema](#nestedatt--items--utilization--average))

<a id="nestedatt--items--utilization--average"></a>
### Nested Schema for `items.utilization.average`

Read-Only:

- `percentage` (Number) Average percentage utilization of the appliance
