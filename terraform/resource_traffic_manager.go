package main

import (
	"fmt"

	"github.com/atlassian/go-vtm"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceTrafficManager() *schema.Resource {
	return &schema.Resource{
		Create: resourceTrafficManagerCreate,
		Read:   resourceTrafficManagerRead,
		Update: resourceTrafficManagerUpdate,
		Delete: resourceTrafficManagerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"adminMasterXMLIP": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0",
			},
			"adminSlaveXMLIP": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0",
			},
			"authenticationServerIP": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0",
			},
			"cloud_platform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"nameip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				//				Default:  "",
			},
			"num_aptimizer_threads": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"num_children": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"numberOfCPUs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"restServerPort": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  11003,
			},
			"trafficip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"updaterIP": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0",
			},
			"gateway_ipv4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				//				Default:  "",
			},
			"gateway_ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hosts": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"if": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"ipv4_forwarding": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ipv6_forwarding": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"licence_agreed": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"manageazureroutes": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"managedpa": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"manageec2conf": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"manageiptrans": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"managereturnpath": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"managevpcconf": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"name_servers": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"ntpservers": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"routes": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"search_domains": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"shim_client_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"shim_client_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"shim_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"shim_ips": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"shim_load_balance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "round_robin",
			},
			"shim_log_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "notice",
			},
			"shim_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "portal",
			},
			"shim_portal_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"shim_proxy_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"shim_proxy_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"ssh_password_allowed": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"ssh_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  22,
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "US/Pacific",
			},
			"vlans": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"allow_update": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"bind_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "*",
			},
			"external_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  9080,
			},
			"trafficips_public_enis": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"bgp_router_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ospfv2_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ospfv2_neighbor_addrs": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"config_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"fwmark": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  320,
			},
			"iptables_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"routing_table": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  320,
			},
			"java_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  9060,
			},
			"email_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bind_ips": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"api_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  9070,
			},
			"allow": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"auth_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"snmp_bind_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "*",
			},
			"community": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "public",
			},
			"snmp_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"hash_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "md5",
			},
			"snmp_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"priv_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "noauthnopriv",
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceTrafficManagerCreate(d *schema.ResourceData, meta interface{}) error {
	err := resourceTrafficManagerSet(d, meta)
	if err != nil {
		return err
	}

	return resourceTrafficManagerRead(d, meta)
}

func resourceTrafficManagerRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client

	r, resp, err := c.GetTrafficManager(d.Get("name").(string))
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			// The resource doesn't exist anymore
			d.SetId("")

			return nil
		}

		return fmt.Errorf("Error reading resource: %s", err)
	}

	d.Set("adminMasterXMLIP", string(*r.Basic.AdminMasterXMLIP))
	d.Set("adminSlaveXMLIP", string(*r.Basic.AdminSlaveXMLIP))
	d.Set("authenticationServerIP", string(*r.Basic.AuthenticationServerIP))
	d.Set("cloud_platform", string(*r.Basic.Cloud_platform))
	d.Set("location", string(*r.Basic.Location))
	d.Set("nameip", string(*r.Basic.Nameip))
	d.Set("num_aptimizer_threads", int(*r.Basic.Num_aptimizer_threads))
	d.Set("num_children", int(*r.Basic.Num_children))
	d.Set("numberOfCPUs", int(*r.Basic.NumberOfCPUs))
	d.Set("restServerPort", int(*r.Basic.RestServerPort))
	d.Set("trafficip", []string(*r.Basic.Trafficip))
	d.Set("updaterIP", string(*r.Basic.UpdaterIP))
	d.Set("gateway_ipv4", string(*r.Appliance.Gateway_ipv4))
	d.Set("gateway_ipv6", string(*r.Appliance.Gateway_ipv6))
	d.Set("hostname", string(*r.Appliance.Hostname))
	d.Set("hosts", []string(*r.Appliance.Hosts))
	d.Set("if", []string(*r.Appliance.If_1))
	d.Set("ip", []string(*r.Appliance.Ip))
	d.Set("ipv4_forwarding", bool(*r.Appliance.Ipv4_forwarding))
	d.Set("licence_agreed", bool(*r.Appliance.Licence_agreed))
	d.Set("manageazureroutes", bool(*r.Appliance.Manageazureroutes))
	//	d.Set("managedpa", bool(*r.Appliance.Managedpa))
	d.Set("manageec2conf", bool(*r.Appliance.Manageec2conf))
	d.Set("manageiptrans", bool(*r.Appliance.Manageiptrans))
	d.Set("managereturnpath", bool(*r.Appliance.Managereturnpath))
	d.Set("managevpcconf", bool(*r.Appliance.Managevpcconf))
	d.Set("name_servers", []string(*r.Appliance.Name_servers))
	d.Set("ntpservers", []string(*r.Appliance.Ntpservers))
	d.Set("routes", []string(*r.Appliance.Routes))
	d.Set("search_domains", []string(*r.Appliance.Search_domains))
	d.Set("shim_client_id", string(*r.Appliance.Shim_client_id))
	d.Set("shim_client_key", string(*r.Appliance.Shim_client_key))
	d.Set("shim_enabled", bool(*r.Appliance.Shim_enabled))
	d.Set("shim_ips", string(*r.Appliance.Shim_ips))
	d.Set("shim_load_balance", string(*r.Appliance.Shim_load_balance))
	d.Set("shim_log_level", string(*r.Appliance.Shim_log_level))
	d.Set("shim_mode", string(*r.Appliance.Shim_mode))
	d.Set("shim_portal_url", string(*r.Appliance.Shim_portal_url))
	d.Set("shim_proxy_host", string(*r.Appliance.Shim_proxy_host))
	d.Set("Shim_proxy_port", string(*r.Appliance.Shim_proxy_port))
	d.Set("ssh_enabled", bool(*r.Appliance.Ssh_enabled))
	//	d.Set("ssh_password_allowed", bool(*r.Appliance.Ssh_password_allowed))
	d.Set("ssh_port", int(*r.Appliance.Ssh_port))
	d.Set("timezone", string(*r.Appliance.Timezone))
	d.Set("vlans", []string(*r.Appliance.Vlans))
	d.Set("allow_update", bool(*r.Cluster_comms.Allow_update))
	d.Set("bind_ip", string(*r.Cluster_comms.Bind_ip))
	d.Set("external_ip", string(*r.Cluster_comms.External_ip))
	d.Set("port", int(*r.Cluster_comms.Port))
	d.Set("trafficips_public_enis", []string(*r.Ec2.Trafficips_public_enis))
	d.Set("bgp_router_id", string(*r.Fault_tolerancec2.Bgp_router_id))
	d.Set("ospfv2_ip", string(*r.Fault_tolerancec2.Ospfv2_ip))
	d.Set("ospfv2_neighbor_addrs", []string(*r.Fault_tolerancec2.Ospfv2_neighbor_addrs))
	d.Set("config_enabled", bool(*r.Iptables.Config_enabled))
	d.Set("fwmark", int(*r.Iptrans.Fwmark))
	d.Set("iptables_enabled", bool(*r.Iptrans.Iptables_enabled))
	d.Set("routing_table", int(*r.Iptrans.Routing_table))
	d.Set("java_port", int(*r.Java.Port))
	d.Set("bind_ips", []string(*r.Rest_api.Bind_ips))
	d.Set("api_port", int(*r.Rest_api.Port))
	d.Set("allow", []string(*r.Snmp.Allow))
	d.Set("auth_password", string(*r.Snmp.Auth_password))
	d.Set("bind_ip", string(*r.Snmp.Bind_ip))
	d.Set("community", string(*r.Snmp.Community))
	d.Set("snmp_enabled", bool(*r.Snmp.Enabled))
	d.Set("hash_algorithm", string(*r.Snmp.Hash_algorithm))
	d.Set("snmp_port", string(*r.Snmp.Port))
	d.Set("priv_password", string(*r.Snmp.Priv_password))
	d.Set("security_level", string(*r.Snmp.Security_level))
	d.Set("username", string(*r.Snmp.Username))

	return nil
}
func resourceTrafficManagerUpdate(d *schema.ResourceData, meta interface{}) error {
	err := resourceTrafficManagerSet(d, meta)
	if err != nil {
		return err
	}

	return resourceTrafficManagerRead(d, meta)
}

func resourceTrafficManagerDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client
	r := stingray.NewTrafficManager(d.Id())

	_, err := c.Delete(r)
	if err != nil {
		return err
	}

	return nil
}
func resourceTrafficManagerSet(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*providerConfig).client
	r := stingray.NewTrafficManager(d.Get("name").(string))

	setString(&r.Basic.AdminMasterXMLIP, d, "adminMasterXMLIP")
	setString(&r.Basic.AdminSlaveXMLIP, d, "adminSlaveXMLIP")
	setString(&r.Basic.AuthenticationServerIP, d, "authenticationServerIP")
	setString(&r.Basic.Cloud_platform, d, "cloud_platform")
	setString(&r.Basic.Location, d, "location")
	setString(&r.Basic.Nameip, d, "nameip")
	setInt(&r.Basic.Num_aptimizer_threads, d, "num_aptimizer_threads")
	setInt(&r.Basic.Num_children, d, "num_children")
	setInt(&r.Basic.NumberOfCPUs, d, "numberOfCPUs")
	setInt(&r.Basic.RestServerPort, d, "restServerPort")
	setStringSet(&r.Basic.Trafficip, d, "trafficip")
	setString(&r.Basic.UpdaterIP, d, "updaterIP")
	setString(&r.Appliance.Gateway_ipv4, d, "gateway_ipv4")
	setString(&r.Appliance.Gateway_ipv6, d, "gateway_ipv6")
	setString(&r.Appliance.Hostname, d, "hostname")
	setStringSet(&r.Appliance.Hosts, d, "hosts")
	setStringSet(&r.Appliance.If_1, d, "if")
	setStringSet(&r.Appliance.Ip, d, "ip")
	setBool(&r.Appliance.Ipv4_forwarding, d, "ipv4_forwarding")
	setBool(&r.Appliance.Licence_agreed, d, "licence_agreed")
	setBool(&r.Appliance.Manageazureroutes, d, "manageazureroutes")
	//  setBool(&r.Appliance.Managedpa, d, "managedpa")
	setBool(&r.Appliance.Manageec2conf, d, "manageec2conf")
	setBool(&r.Appliance.Manageiptrans, d, "manageiptrans")
	setBool(&r.Appliance.Managereturnpath, d, "managereturnpath")
	setBool(&r.Appliance.Managevpcconf, d, "managevpcconf")
	setStringSet(&r.Appliance.Name_servers, d, "name_servers")
	setStringSet(&r.Appliance.Ntpservers, d, "ntpservers")
	setStringSet(&r.Appliance.Routes, d, "routes")
	setStringSet(&r.Appliance.Search_domains, d, "search_domains")
	setString(&r.Appliance.Shim_client_id, d, "shim_client_id")
	setString(&r.Appliance.Shim_client_key, d, "shim_client_key")
	setBool(&r.Appliance.Shim_enabled, d, "shim_enabled")
	setString(&r.Appliance.Shim_ips, d, "shim_ips")
	setString(&r.Appliance.Shim_load_balance, d, "shim_load_balance")
	setString(&r.Appliance.Shim_log_level, d, "shim_log_level")
	setString(&r.Appliance.Shim_mode, d, "shim_mode")
	setString(&r.Appliance.Shim_portal_url, d, "shim_portal_url")
	setString(&r.Appliance.Shim_proxy_host, d, "shim_proxy_host")
	setString(&r.Appliance.Shim_proxy_port, d, "shim_proxy_port")
	setBool(&r.Appliance.Ssh_enabled, d, "ssh_enabled")
	//	setBool(&r.Appliance.Ssh_password_allowed, d, "ssh_password_allowed")
	setInt(&r.Appliance.Ssh_port, d, "ssh_port")
	setString(&r.Appliance.Timezone, d, "timezone")
	setStringSet(&r.Appliance.Vlans, d, "vlans")
	setBool(&r.Cluster_comms.Allow_update, d, "allow_update")
	setString(&r.Cluster_comms.Bind_ip, d, "bind_ip")
	setString(&r.Cluster_comms.External_ip, d, "external_ip")
	setInt(&r.Cluster_comms.Port, d, "port")
	setStringSet(&r.Ec2.Trafficips_public_enis, d, "trafficips_public_enis")
	setString(&r.Fault_tolerancec2.Bgp_router_id, d, "bgp_router_id")
	setString(&r.Fault_tolerancec2.Ospfv2_ip, d, "ospfv2_ip")
	setStringSet(&r.Fault_tolerancec2.Ospfv2_neighbor_addrs, d, "ospfv2_neighbor_addrs")
	setBool(&r.Iptables.Config_enabled, d, "config_enabled")
	setInt(&r.Iptrans.Fwmark, d, "fwmark")
	setBool(&r.Iptrans.Iptables_enabled, d, "iptables_enabled")
	setInt(&r.Iptrans.Routing_table, d, "routing_table")
	setInt(&r.Java.Port, d, "java_port")
	setStringSet(&r.Rest_api.Bind_ips, d, "bind_ips")
	setInt(&r.Rest_api.Port, d, "api_port")
	setStringSet(&r.Snmp.Allow, d, "allow")
	setString(&r.Snmp.Auth_password, d, "auth_password")
	setString(&r.Snmp.Bind_ip, d, "bind_ip")
	setString(&r.Snmp.Community, d, "community")
	setBool(&r.Snmp.Enabled, d, "snmp_enabled")
	setString(&r.Snmp.Hash_algorithm, d, "hash_algorithm")
	setString(&r.Snmp.Port, d, "snmp_port")
	setString(&r.Snmp.Priv_password, d, "priv_password")
	setString(&r.Snmp.Security_level, d, "security_level")
	setString(&r.Snmp.Username, d, "username")

	_, err := c.Set(r)
	if err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))
	return nil
}
