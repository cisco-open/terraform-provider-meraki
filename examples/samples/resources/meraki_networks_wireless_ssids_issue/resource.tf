terraform {
  required_providers {
    meraki = {
      version = "1.1.3-beta"
      source  = "hashicorp.com/edu/meraki"
      # "hashicorp.com/edu/meraki" is the local built source, change to "cisco-en-programmability/meraki" to use downloaded version from registry
    }
  }
}

variable "wifi_disable" {
  description = "Indica si se debe deshabilitar la red Wi-Fi"
  type        = bool
  default     = false # Establecido en falso, indicando que la red Wi-Fi está habilitada por defecto.
}

variable "wpa3_enable" {
  description = "Indica si se debe habilitar WPA3"
  type        = bool
  default     = true # Establecido en verdadero, habilitando WPA3.
}

variable "ise_servers" {
  description = "Lista de servidores ISE con IP y secreto compartido"
  type = list(object({
    server_ip         = string
    ise_shared_secret = string
  }))
  default = [
    {
      server_ip         = "192.168.1.10"
      ise_shared_secret = "secret_key_1"
    },
    {
      server_ip         = "192.168.1.11"
      ise_shared_secret = "secret_key_2"
    }
  ]
}

variable "network_id" {
  description = "ID de la red en Meraki"
  type        = string
  default     = "1234567890" # ID de ejemplo para la red Meraki.
}

variable "SSID_NAME" {
  description = "Nombre de la red SSID"
  type        = string
  default     = "Mi_SSID" # Nombre de ejemplo para el SSID.
}

variable "node_mac" {
  description = "Dirección MAC del nodo"
  type        = string
  default     = "00:11:22:33:44:55" # Dirección MAC de ejemplo.
}

variable "vap_name" {
  description = "Nombre del VAP"
  type        = string
  default     = "VAP_1" # Nombre de ejemplo para el VAP.
}

variable "vap_num" {
  description = "Número del VAP"
  type        = number
  default     = 1 # Número de VAP de ejemplo.
}


provider "meraki" {
  meraki_debug    = true
  meraki_base_url = "http://localhost:3002"
}

resource "meraki_networks_wireless_ssids" "SSID_name" {

  network_id  = "1"
  number      = 0
  name        = "SSID_NAME"
  enabled     = var.wifi_disable == true ? false : true
  splash_page = "None"
  auth_mode   = "8021x-radius"
  dot11w = {
    enabled  = var.wpa3_enable
    required = false
  }
  dot11r = {
    enabled  = true
    adaptive = false
  }
  encryption_mode     = "wpa-eap"
  wpa_encryption_mode = var.wpa3_enable ? "WPA3 Transition Mode" : "WPA2 only"
  radius_servers = [
    for server in var.ise_servers : {
      host   = server.server_ip
      port   = 1812
      secret = server.ise_shared_secret
  }]
  radius_accounting_enabled = true
  radius_accounting_servers = [
    for server in var.ise_servers : {
      host   = server.server_ip
      port   = 1813
      secret = server.ise_shared_secret
  }]
  radius_testing_enabled              = false
  radius_server_timeout               = 1
  radius_server_attempts_limit        = 3
  radius_fallback_enabled             = false
  radius_accounting_interim_interval  = 1200
  radius_proxy_enabled                = false
  radius_coa_enabled                  = false
  radius_called_station_id            = "$NODE_MAC$:$VAP_NAME$"
  radius_authentication_nas_id        = "$NODE_MAC$:$VAP_NUM$"
  radius_attribute_for_group_policies = "Airespace-ACL-Name"
  ip_assignment_mode                  = "Bridge mode"
  use_vlan_tagging                    = false
  radius_override                     = true
  min_bitrate                         = 11
  band_selection                      = "Dual band operation"
  per_client_bandwidth_limit_up       = 0
  per_client_bandwidth_limit_down     = 0
  per_ssid_bandwidth_limit_up         = 0
  per_ssid_bandwidth_limit_down       = 0
  mandatory_dhcp_enabled              = false
  lan_isolation_enabled               = false
  visible                             = true
  available_on_all_aps                = true
}

