
resource "meraki_organizations_branding_policies" "example" {

  admin_settings = {

    applies_to = "All admins of networks..."
    values     = ["N_1234", "L_5678"]
  }
  custom_logo = {

    enabled = true
    image = {

      contents = "Hyperg26C8F4h8CvcoUqpA=="
      format   = "jpg"
    }
  }
  enabled = true
  help_settings = {

    api_docs_subtab                        = "default or inherit"
    cases_subtab                           = "hide"
    cisco_meraki_product_documentation     = "show"
    community_subtab                       = "show"
    data_protection_requests_subtab        = "default or inherit"
    firewall_info_subtab                   = "hide"
    get_help_subtab                        = "default or inherit"
    get_help_subtab_knowledge_base_search  = "<h1>Some custom HTML content</h1>"
    hardware_replacements_subtab           = "hide"
    help_tab                               = "show"
    help_widget                            = "hide"
    new_features_subtab                    = "show"
    sm_forums                              = "hide"
    support_contact_info                   = "show"
    universal_search_knowledge_base_search = "hide"
  }
  name            = "My Branding Policy"
  organization_id = "string"
}

output "meraki_organizations_branding_policies_example" {
  value = meraki_organizations_branding_policies.example
}