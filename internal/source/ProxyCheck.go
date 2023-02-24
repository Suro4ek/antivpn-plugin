package source

import (
	"encoding/json"
	"hyneo-antivpn/internal/antivpn"
	"hyneo-antivpn/internal/model"
	"hyneo-antivpn/internal/utils"
)

const PROXYCHECKURL = "https://proxycheck.io/v2/"

type ProxyCheck struct {
}

func NewProxyCheck() antivpn.Source {
	return &ProxyCheck{}
}

func (V ProxyCheck) GetResult(ip string) (bool, error) {
	b, err := utils.NewResponse(PROXYCHECKURL, ip+"?vpn=1&asn=1&key=41qi90-572b71-1rxj55-64692i")
	if err != nil {
		return false, err
	}
	var dat map[string]interface{}
	err = json.Unmarshal(b, &dat)
	js, err := json.Marshal(dat[ip])
	var VPNAPI model.ProxyCheck
	err = json.Unmarshal(js, &VPNAPI)
	if err != nil {
		return false, err
	}
	if VPNAPI.Proxy == "yes" {
		return true, nil
	}
	return false, nil
}
