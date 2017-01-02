# Terraform VTM

The VTM Terraform provider is used to interact with the VTM load balancer.

## Example usage
~~~
provider "vtm" {
    url = "https://example:9070"
    username = "username"
    password = "password"
    verify_ssl = "false"
}
~~~

## Argument Reference

* ```url``` - The protocol, host name, and port for the VTM REST API
* ```username``` - The username for authenticating against the API
* ```password``` - The password for authenticating against the API
* ```valid_networks``` - A comma separated list of valid traffic IP networks (in CIDR notation)
* ```verify_ssl``` - Perform SSL verification, default is true

The provider can also be configured through the environmental variables VTM_URL, VTM_USERNAME, VTM_PASSWORD, VTM_VALID_NETWORKS, and VTM_VERIFY_SSL.

## Supported Resources
This is short overview of the common resource configuration
### vtm_traffic_ip_group example

~~~
resource "vtm_traffic_ip_group" "IP_GROUP" {
  name        = "${var.provider}"
  enabled     = "true"
  ipaddresses = ["172.17.0.3"]
  machines    = ["570f964cf9bb"]
  note        = "This is test IP traffic group ${var.provider}
}
~~~

* ```name``` - Name of traffic group __(required)__
* ```enabled``` - Traffic group is disabled or enabled, Default:  true
* ```ipaddresses``` - IP address, separated by comma
* ```machines``` - Traffic manager ID
* ```note``` - Default: ""
* ```hash_source_port``` - Default:  false
* ```keeptogether``` - Default:  false
* ```location``` - Default:  0
* ```mode``` - Default:  "singlehosted"
* ```multicast``` - Default:  ""
* ```rhi_ospfv2_metric_base``` - Default:  10
* ```rhi_ospfv2_passive_metric_offset``` - Default:  10
* ```slaves```

### vtm_virtual_server example

~~~
resource "vtm_virtual_server" "LON3" {
  name                  = "virtual_server1"
  enabled               = "true"
  listen_on_traffic_ips = ["LON5"]
  port                  = "80"
  protocol              = "http"
  pool                  = "test1"
  note                  = "This is test Virtual Server ${var.provider}"
}
~~~

* ```name``` - Name of virtual Server __(required)__
* ```enabled``` - Virtual server is disabled or enabled, Default:  true
* ```listen_on_any``` - Listen on any IP address, Default:  false
* ```listen_on_traffic_ips``` - Listen on Traffic Group
* ```port``` - Port __(required)__
* ```pool``` - Name of the virtual pool __(required)__
* ```protocol``` - Internal protocol
* ```add_x_forwarded_for``` - Default:  false
* ```add_x_forwarded_proto``` - Default:  false
* ```connection_errors_error_file``` - Default:  "Default"
* ```connection_keepalive_timeout``` - Default:  10
* ```connection_timeout``` - Default:  300
* ```connect_timeout``` - Default:  10
* ```gzip_enabled``` - Default:  false
* ```gzip_compress_level``` - Default:  1
* ```gzip_max_size``` - Default:  100000
* ```gzip_min_size``` -  Default:  1000
* ```gzip_include_mime```
* ```http_location_rewrite``` - Default:  "if_host_matches"
* ```log_enabled``` - Default:  false
* ```log_filename``` - Default:  "%zeushome%/zxtm/log/%v.log"
* ```log_format``` - Default "%h %l %u %t \"%r\" %s %b \"%{Referer}i\" \"%{User-agent}i\""
* ```log_server_connection_failures``` - Default:  false
* ```note``` - Default: ""
* ```recent_connections_save_all``` - Default:  false
* ```protection_class```
* ```request_rules```
* ```response_rules```
* ```completion_rules```
* ```ssl_add_http_headers``` - Default:  false
* ```ssl_trust_magic``` - Default: false
* ```ssl_decrypt``` - Default: false
* ```ssl_server_cert_default``` - Default: ""
* ```ssl_server_cert_host_mapping```
* ```syslog_enabled``` - Default:  false
* ```syslog_format``` - Default:  "%h %l %u %t \"%r\" %s %b \"%{Referer}i\" \"%{User-agent}i\""
* ```syslog_ip_end_point```
* ```syslog_msg_len_limit``` - Default:  1024
* ```web_cache_enabled``` - Default:  false
* ```web_cache_max_time``` - Default:  600,
* ```web_cache_refresh_time``` - Default:  2
* ```web_cache_control_out```
* ```web_cache_error_page_time``` - Default:  30

### vtm_pool example
~~~
resource "vtm_pool" "Pool1" {
  name = "test1"

  node = {
    node  = "172.22.212.26:8444"
    state = "active"
  }
  note = "This is test Pool ${var.provider}"
}
~~~

* ```name``` - Name of Virtual pool __(required)__
* ```bandwidth_class``` - Default:  ""
* ```connection_max_connect_time``` - Default:  4
* ```connection_max_connections_per_node``` - Default:  0
* ```connection_max_queue_size``` - Default:  0
* ```connection_max_reply_time``` - Default:  30
* ```connection_queue_timeout``` - Default:  10
* ```dns_autoscale_enabled``` - Default:  false
* ```dns_autoscale_hostnames```
* ```dns_autoscale_port``` - Default:  80
* ```failure_pool``` - Default:  ""
* ```load_balancing_algorithm``` - Default:  "round_robin"
* ```load_balancing_priority_enabled``` - Default:  false
* ```load_balancing_priority_nodes``` - Default:  1
* ```max_connection_attempts``` - Default:  0
* ```max_idle_connections_pernode``` - Default:  50
* ```max_timed_out_connection_attempts``` - Default:  2
* ```monitors```
* ```node_close_with_rst``` - Default:  false
* ```node_connection_attempts``` - Default:  3
* ```node``` -
  * ```node``` - IP address of the node __(required)__
  * ```weight``` - Default:  1
  * ```state``` - Default:  "active"
  * ```priority``` - Default:  1
* ```note``` - Default: ""
* ```passive_monitoring``` - Default:  true
* ```persistence_class``` - Default:  ""
* ```tcp_nagle``` - Default:  true
* ```transparent``` - Default:  false
* ```udp_accept_from``` - Default:  "dest_only"
* ```udp_accept_from_mask``` - Default:  ""

### vtm_license_key example
~~~
resource "vtm_license_key" "DEV" {
  name    = "dev_license"
  content = "${var.license}"
}
~~~

* ```content``` - Licence key content __(required)__
* ```name``` - Licence key name __(required)__
