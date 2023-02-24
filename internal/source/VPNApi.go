package source

import (
	"encoding/json"
	"hyneo-antivpn/internal/antivpn"
	"hyneo-antivpn/internal/model"
	"hyneo-antivpn/internal/utils"
)

const VPNURL = "https://vpnapi.io/api/"

type VPNApi struct {
}

func NewVPNApi() antivpn.Source {
	return &VPNApi{}
}

func (V VPNApi) GetResult(ip string) (bool, error) {
	b, err := utils.NewResponse(VPNURL, ip)
	if err != nil {
		return false, err
	}
	var VPNAPI model.VPNAPI
	err = json.Unmarshal(b, &VPNAPI)
	if err != nil {
		return false, err
	}
	if VPNAPI.Security.VPN || VPNAPI.Security.Proxy || VPNAPI.Security.Tor {
		return true, nil
	}
	return false, nil
}
