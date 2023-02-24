package model

type ProxyCheck struct {
	Proxy string `json:"proxy"`
}

//type Shodan struct {
//	Tags map[ShodanTags]struct{} `json:"tags"`
//}
//
//type ShodanTags struct {
//	VPN   *string `json:"vpn"`
//	Proxy *string `json:"proxy"`
//}

type VPNAPI struct {
	Security VPNApiSecurity `json:"security"`
}

type VPNApiSecurity struct {
	VPN   bool `json:"vpn"`
	Proxy bool `json:"proxy"`
	Tor   bool `json:"tor"`
}
