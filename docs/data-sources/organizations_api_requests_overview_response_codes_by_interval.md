---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_organizations_api_requests_overview_response_codes_by_interval Data Source - terraform-provider-meraki"
subcategory: "organizations"
description: |-
  
---

# meraki_organizations_api_requests_overview_response_codes_by_interval (Data Source)



## Example Usage

```terraform
data "meraki_organizations_api_requests_overview_response_codes_by_interval" "example" {

  admin_ids       = ["string"]
  interval        = 1
  operation_ids   = ["string"]
  organization_id = "string"
  source_ips      = ["string"]
  t0              = "string"
  t1              = "string"
  timespan        = 1.0
  user_agent      = "string"
  version         = 1
}

output "meraki_organizations_api_requests_overview_response_codes_by_interval_example" {
  value = data.meraki_organizations_api_requests_overview_response_codes_by_interval.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `organization_id` (String) organizationId path parameter. Organization ID

### Optional

- `admin_ids` (List of String) adminIds query parameter. Filter by admin ID of user that made the API request
- `interval` (Number) interval query parameter. The time interval in seconds for returned data. The valid intervals are: 120, 3600, 14400, 21600. The default is 21600. Interval is calculated if time params are provided.
- `operation_ids` (List of String) operationIds query parameter. Filter by operation ID of the endpoint
- `source_ips` (List of String) sourceIps query parameter. Filter by source IP that made the API request
- `t0` (String) t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
- `t1` (String) t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
- `timespan` (Number) timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. If interval is provided, the timespan will be autocalculated.
- `user_agent` (String) userAgent query parameter. Filter by user agent string for API request. This will filter by a complete or partial match.
- `version` (Number) version query parameter. Filter by API version of the endpoint. Allowable values are: [0, 1]

### Read-Only

- `items` (Attributes List) Array of ResponseOrganizationsGetOrganizationApiRequestsOverviewResponseCodesByInterval (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `counts` (Attributes Set) list of response codes and a count of how many requests had that code in the given time period (see [below for nested schema](#nestedatt--items--counts))
- `end_ts` (String) The end time of the access period
- `start_ts` (String) The start time of the access period

<a id="nestedatt--items--counts"></a>
### Nested Schema for `items.counts`

Read-Only:

- `code` (Number) Response status code of the API response
- `count` (Number) Number of records that match the status code
