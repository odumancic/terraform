# Terraform VTM

The VTM Terraform provider is used to interact with the VTM load balancer based on
```atlassian/go-vtm``` and ```atlassian/terraform-provider-vtm```

* Added support for managing Traffic Managers .

## Supported Resources

See the `resource_*.go` files for available resources and the
supported arguments for each resource.

Support for resources is being added as needed. **Bold** resources are
fully supported.

- [x] **Action Program**
- [x] **Alerting Action**
- [ ] Aptimizer Application Scope
- [ ] Aptimizer Profile
- [ ] Bandwidth Class
- [ ] Cloud Credentials
- [ ] Custom configuration set
- [x] **Event Type**
- [x] **Extra File**
- [ ] GLB Service
- [ ] Global Settings
- [x] **License**
- [ ] Location
- [x] **Monitor**
- [x] **Monitor Program**
- [ ] NAT Configuration
- [x] Pool
- [x] **Protection Class**
- [x] **Rate Shaping Class**
- [x] **Rule**
- [x] **SLM Class**
- [ ] SSL Client Key Pair
- [x] **SSL Key Pair**
- [x] **SSL Trusted Certificate**
- [ ] Security Settings
- [ ] Session Persistence Class
- [x] **Traffic IP Group**
- [x] **Traffic Manager**
- [ ] TrafficScript Authenticator
- [ ] User Authenticator
- [ ] User Group
- [x] Virtual Server

## Example usage
~~~
provider "vtm" {
    url = "https://example:9070"
    username = "username"
    password = "password"
    verify_ssl = "false"
}

resource "vtm_traffic_manager" "LON3"{
        provider = "vtm.LON3"
        name = "a198d9291760"
        bind_ips = ["*"]
        ntpservers = ["0.vyatta.pool.ntp.org","1.vyatta.pool.ntp.org"]
}

resource "vtm_traffic_ip_group" "IP_GROUP" {
  name        = "LON3"
  enabled     = "true"
  ipaddresses = ["172.17.0.3"]
  machines    = ["570f964cf9bb"]
  note        = "This is test IP traffic group ${var.provider}
}

resource "vtm_virtual_server" "LON3" {
  name                  = "virtual_server"
  enabled               = "true"
  listen_on_traffic_ips = ["LON3"]
  port                  = "80"
  protocol              = "http"
  pool                  = "test"
  note                  = "This is test Virtual Server ${var.provider}"
}

resource "vtm_pool" "Pool" {
  name = "test"
  node = {
    node  = "172.22.212.26:8444"
    state = "active"
  }
  note = "This is test Pool ${var.provider}"
}

resource "vtm_license_key" "DEV" {
  name    = "dev_license"
  content = "${var.license}"
}
~~~

## Argument Reference
### vtm

* ```url``` - The protocol, host name, and port for the VTM REST API
* ```username``` - The username for authenticating against the API
* ```password``` - The password for authenticating against the API
* ```valid_networks``` - A comma separated list of valid traffic IP networks (in CIDR notation)
* ```verify_ssl``` - Perform SSL verification, default is true

The provider can also be configured through the environmental variables VTM_URL, VTM_USERNAME, VTM_PASSWORD, VTM_VALID_NETWORKS, and VTM_VERIFY_SSL.

### vtm_traffic_manager

* ```name``` - Name of traffic manager __(required)__
* ```adminMasterXMLIP``` - Default: "0.0.0.0"
* ```adminSlaveXMLIP``` - Default: "0.0.0.0"
* ```allow``` -
* ```allow_update``` - Default: true
* ```api_port``` - Default: 9070
* ```authenticationServerIP``` Default: "0.0.0.0"
* ```bgp_router_id```
* ```bind_ip``` - Default: "*"
* ```bind_ips``` - Default: "*" __(required)__
* ```cloud_platform```
* ```community``` - Default: "public"
* ```config_enabled``` - Default: true
* ```external_ip```
* ```fwmark``` - Default: 320
* ```gateway_ipv4```
* ```gateway_ipv6```
* ```hash_algorithm``` - Default: "md5"
* ```hostname```
* ```hosts``` - Default: 0
* ```if``` - Default: 0
* ```ip``` - Default: 0
* ```iptables_enabled``` - Default: true
* ```ipv4_forwarding``` - Default: false
* ```ipv6_forwarding``` - Default: false
* ```java_port``` - Default: 9060
* ```licence_agreed``` - Default: false
* ```location```
* ```manageazureroutes``` - Default: true
* ```managedpa``` - ~~Default: true~~ Not implemented
* ```manageec2conf``` - Default: true
* ```manageiptrans``` - Default: true
* ```managereturnpath``` - Default: true
* ```managevpcconf``` - Default: true
* ```name_servers```
* ```nameip```
* ```ntpservers```
* ```num_aptimizer_threads``` - Default: 0
* ```num_children``` - Default: 0
* ```numberOfCPUs``` - Default: 0
* ```ospfv2_ip```
* ```port``` - Default: 9080
* ```restServerPort``` - Default: 11003
* ```routes``` - Default: none
* ```routing_table``` - Default: 320
* ```search_domains``` - Default: none
* ```security_level``` - Default: "noauthnopriv"
* ```shim_client_id```
* ```shim_client_key```
* ```shim_enabled``` - Default: false
* ```shim_ips```
* ```shim_load_balance``` - Default: "round_robin"
* ```shim_log_level``` - Default: "notice"
* ```shim_mode``` - Default: "portal"
* ```shim_portal_url```
* ```shim_proxy_host```
* ```snmp_bind_ip``` - Default: "*"
* ```snmp_enabled``` - Default: "false"
* ```snmp_port``` - Default: "default"
* ```ssh_enabled``` - Default: "true"
* ```ssh_password_allowed``` - Default: "true"
* ```ssh_port``` - Default: 22
* ```timezone``` - Default: "US/Pacific"
* ```trafficip```
* ```updaterIP``` - Default: "0.0.0.0"
* ```username```

### vtm_traffic_ip_group example


* ```name``` - Name of traffic ip group __(required)__
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

* ```content``` - Licence key content __(required)__
* ```name``` - Licence key name __(required)__

## Variable file example

This is example of variables file.

~~~
variable "provider" {
  default = "LON1"
}
variable "license" {
  default = "6430da6559e3812eb048205460d5250c06024a84"
}
variable "url" {
  default = {
    LON1 = "https://192.168.0.15:9070"
    LON2 = "https://192.168.0.15:9071"
  }
}
variable "username" {
  default = {
    LON1 = "admin"
    LON2 = "admin"
  }
}
variable "password" {
  default = {
    LON1 = "admin"
    LON2 = "password"
  }
}
~~~

## vtm.tf example using variables and multiple providers
~~~
provider "vtm" {
  url        = "${var.url["${var.provider}"]}"
  username   = "${var.username["${var.provider}"]}"
  password   = "${var.password["${var.provider}"]}"
  verify_ssl = "false"
}

provider "vtm" {
  alias      = "LON1"
  url        = "${var.url["LON1"]}"
  username   = "${var.username["LON1"]}"
  password   = "${var.password["LON1"]}"
  verify_ssl = "false"
}

provider "vtm" {
  alias      = "LON2"
  url        = "${var.url["LON2"]}"
  username   = "${var.username["LON2"]}"
  password   = "${var.password["LON2"]}"
  verify_ssl = "false"
}

## License

resource "vtm_license_key" "DEV" {
  provider= "vtm.LON1"
  name    = "dev_license"
  content = "${var.license}"
}

## LON1

resource "vtm_traffic_manager" "LON1"{
  provider 		= "vtm.LON1"
  name			= "a198d9291760"
  bind_ips 		= ["*"]
  ntpservers 	= ["0.vyatta.pool.ntp.org","1.vyatta.pool.ntp.org", "2.vyatta.pool.ntp.org", "3.vyatta.pool.ntp.org"]
}

resource "vtm_traffic_ip_group" "LON1" {
  provider    = "vtm.LON1"
  name        = "LON1"
  enabled     = "true"
  ipaddresses = ["172.17.0.4"]
  machines    = ["${vtm_traffic_manager.LON1.id}"]
  note        = "This is test IP traffic group ${vtm_traffic_manager.LON1.id}"
}

resource "vtm_virtual_server" "LON1" {
  provider 				= "vtm.LON1"
  name                  = "virtual_server"
  enabled               = "true"
  listen_on_traffic_ips = ["${vtm_traffic_ip_group.LON1.name}"]
  port                  = "80"
  protocol              = "http"
  pool                  = "Bar"
  note                  = "This is test Virtual Server ${vtm_traffic_ip_group.LON1.name}"
}

resource "vtm_pool" "LON1" {
  provider 	= "vtm.LON1"
  name 		= "Bar"
  node 		= {
    node  	= "172.22.212.26:8444"
    state 	= "active"
  }
  note 		= "This is test Pool ${vtm_virtual_server.LON1.name}"
}

# LON2

resource "vtm_traffic_manager" "LON2"{
  provider              = "vtm.LON2"
  name                  = "1f14b390e423"
  bind_ips              = ["*"]
  ntpservers    = ["0.vyatta.pool.ntp.org","1.vyatta.pool.ntp.org", "2.vyatta.pool.ntp.org", "3.vyatta.pool.ntp.org"]
}

resource "vtm_traffic_ip_group" "LON2" {
  provider    = "vtm.LON2"
  name        = "LON2"
  enabled     = "true"
  ipaddresses = ["172.17.0.2"]
  machines    = ["${vtm_traffic_manager.LON2.id}"]
  note        = "This is test IP traffic group ${vtm_traffic_manager.LON2.id}"
}

resource "vtm_virtual_server" "LON2" {
  provider              = "vtm.LON2"
  name                  = "virtual_server"
  enabled               = "true"
  listen_on_traffic_ips = ["${vtm_traffic_ip_group.LON2.name}"]
  port                  = "80"
  protocol              = "http"
  pool                  = "Bar"
  note                  = "This is test Virtual Server ${vtm_traffic_ip_group.LON2.name}"
}

resource "vtm_pool" "LON2" {
  provider      = "vtm.LON2"
  name          = "Bar"
  node          = {
    node        = "172.22.212.26:8444"
    state       = "active"
  }
  note          = "This is test Pool ${vtm_virtual_server.LON2.name}"
}


output "LON1" {
  value = "${vtm_traffic_manager.LON1.id}"
}

output "LON2" {
  value = "${vtm_traffic_manager.LON2.id}"
}
~~~


## Execute plan

~~~
terraform plan -var 'provider=LON5'

terraform apply -var 'provider=LON5'
~~~
