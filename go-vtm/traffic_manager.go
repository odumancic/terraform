package stingray

import (
	"encoding/json"
	"net/http"
)

// A TrafficManager is a Stingray traffic manager.
type TrafficManager struct {
	jsonResource             `json:"-"`
	TrafficManagerProperties `json:"properties"`
}

type TrafficManagerProperties struct {
	Basic struct {
		AdminMasterXMLIP       *string   `json:"adminMasterXMLIP,omitempty"`
		AdminSlaveXMLIP        *string   `json:"adminSlaveXMLIP,omitempty"`
		Appliance_card         *[]string `json:"appliance_card,omitempty"`
		Appliance_sysctl       *[]string `json:"appliance_sysctl,omitempty"`
		AuthenticationServerIP *string   `json:"authenticationServerIP,omitempty"`
		Cloud_platform         *string   `json:"cloud_platform,omitempty"`
		Location               *string   `json:"location,omitempty"`
		Nameip                 *string   `json:"nameip,omitempty"`
		Num_aptimizer_threads  *int      `json:"num_aptimizer_threads,omitempty"`
		Num_children           *int      `json:"num_children,omitempty"`
		NumberOfCPUs           *int      `json:"numberOfCPUs,omitempty"`
		RestServerPort         *int      `json:"restServerPort,omitempty"`
		Trafficip              *[]string `json:"trafficip,omitempty"`
		UpdaterIP              *string   `json:"updaterIP,omitempty"`
	} `json:"basic"`
	Appliance struct {
		Gateway_ipv4         *string   `json:"gateway_ipv4,omitempty"`
		Gateway_ipv6         *string   `json:"gateway_ipv6,omitempty"`
		Hostname             *string   `json:"hostname,omitempty"`
		Hosts                *[]string `json:"hosts,omitempty"`
		If_1                 *[]string `json:"if,omitempty"`
		Ip                   *[]string `json:"ip,omitempty"`
		Ipmi_lan_access      *bool     `json:"ipmi_lan_access,omitempty"`
		Ipmi_lan_addr        *string   `json:"ipmi_lan_addr,omitempty"`
		Ipmi_lan_gateway     *string   `json:"ipmi_lan_gateway,omitempty"`
		Ipmi_lan_ipsrc       *string   `json:"ipmi_lan_ipsrc,omitempty"`
		Ipmi_lan_mask        *string   `json:"ipmi_lan_mask,omitempty"`
		Ipv4_forwarding      *bool     `json:"ipv4_forwarding,omitempty"`
		Ipv6_forwarding      *bool     `json:"ipv6_forwarding,omitempty"`
		Licence_agreed       *bool     `json:"licence_agreed,omitempty"`
		Manageazureroutes    *bool     `json:"manageazureroutes,omitempty"`
		Managedpa            *bool     `json:"managedpa,omitempty"`
		Manageec2conf        *bool     `json:"manageec2conf,omitempty"`
		Manageiptrans        *bool     `json:"manageiptrans,omitempty"`
		Managereturnpath     *bool     `json:"managereturnpath,omitempty"`
		Managevpcconf        *bool     `json:"managevpcconf,omitempty"`
		Name_servers         *[]string `json:"name_servers,omitempty"`
		Ntpservers           *[]string `json:"ntpservers,omitempty"`
		Routes               *[]string `json:"routes,omitempty"`
		Search_domains       *[]string `json:"search_domains,omitempty"`
		Shim_client_id       *string   `json:"shim_client_id,omitempty"`
		Shim_client_key      *string   `json:"shim_client_key,omitempty"`
		Shim_enabled         *bool     `json:"shim_enabled,omitempty"`
		Shim_ips             *string   `json:"shim_ips,omitempty"`
		Shim_load_balance    *string   `json:"shim_load_balance,omitempty"`
		Shim_log_level       *string   `json:"shim_log_level,omitempty"`
		Shim_mode            *string   `json:"shim_mode,omitempty"`
		Shim_portal_url      *string   `json:"shim_portal_url,omitempty"`
		Shim_proxy_host      *string   `json:"shim_proxy_host,omitempty"`
		Shim_proxy_port      *string   `json:"shim_proxy_port,omitempty"`
		Ssh_enabled          *bool     `json:"ssh_enabled,omitempty"`
		Ssh_password_allowed *bool     `json:"ssh_password_allowed,omitempty"`
		Ssh_port             *int      `json:"ssh_port,omitempty"`
		Timezone             *string   `json:"timezone,omitempty"`
		Vlans                *[]string `json:"vlans,omitempty"`
	} `json:"appliance"`
	Autodiscover  struct{} `json:"autodiscover"`
	Cluster_comms struct {
		Allow_update *bool   `json:"allow_update,omitempty"`
		Bind_ip      *string `json:"bind_ip,omitempty"`
		External_ip  *string `json:"external_ip,omitempty"`
		Port         *int    `json:"port,omitempty"`
	} `json:"cluster_comms"`
	Ec2 struct {
		Trafficips_public_enis *[]string `json:"trafficips_public_enis,omitempty"`
	} `json:"ec2"`
	Fault_tolerancec2 struct {
		Bgp_router_id         *string   `json:"bgp_router_id,omitempty"`
		Ospfv2_ip             *string   `json:"ospfv2_ip,omitempty"`
		Ospfv2_neighbor_addrs *[]string `json:"ospfv2_neighbor_addrs,omitempty"`
	} `json:"fault_tolerance"`
	Iop      struct{} `json:"iop"`
	Iptables struct {
		Config_enabled *bool `json:"config_enabled,omitempty"`
	} `json:"iptables"`
	Iptrans struct {
		Fwmark           *int  `json:"fwmark,omitempty"`
		Iptables_enabled *bool `json:"iptables_enabled,omitempty"`
		Routing_table    *int  `json:"routing_table,omitempty"`
	} `json:"iptrans"`
	Java struct {
		Port *int `json:"port,omitempty"`
	} `json:"java"`
	Kerberos         struct{} `json:"kerberos"`
	Remote_licensing struct {
		Email_address *string `json:"email_address,omitempty"`
		Message       *string `json:"message,omitempty"`
	} `json:"remote_licensing"`
	Rest_api struct {
		Bind_ips *[]string `json:"bind_ips,omitempty"`
		Port     *int      `json:"port,omitempty"`
	} `json:"rest_api"`
	Snmp struct {
		Allow          *[]string `json:"allow,omitempty"`
		Auth_password  *string   `json:"auth_password,omitempty"`
		Bind_ip        *string   `json:"bind_ip,omitempty"`
		Community      *string   `json:"community,omitempty"`
		Enabled        *bool     `json:"enabled,omitempty"`
		Hash_algorithm *string   `json:"hash_algorithm,omitempty"`
		Port           *string   `json:"port,omitempty"`
		Priv_password  *string   `json:"priv_password,omitempty"`
		Security_level *string   `json:"security_level,omitempty"`
		Username       *string   `json:"username,omitempty"`
	} `json:"snmp"`
}

func (r *TrafficManager) endpoint() string {
	return "traffic_managers"
}

func (r *TrafficManager) String() string {
	s, _ := jsonMarshal(r)
	return string(s)
}

func (r *TrafficManager) decode(data []byte) error {
	return json.Unmarshal(data, &r)
}

func NewTrafficManager(name string) *TrafficManager {
	r := new(TrafficManager)
	r.setName(name)
	return r
}

func (c *Client) GetTrafficManager(name string) (*TrafficManager, *http.Response, error) {
	r := NewTrafficManager(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

func (c *Client) ListTrafficManagers() ([]string, *http.Response, error) {
	return c.List(&TrafficManager{})
}
