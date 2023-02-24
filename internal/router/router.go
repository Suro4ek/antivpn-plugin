package router

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	antivpn2 "hyneo-antivpn/internal/antivpn"
	"hyneo-antivpn/protos/antivpn"
)

type antivpnRouter struct {
	service antivpn2.Service
	antivpn.UnimplementedAntiVPNServer
}

func NewAntiVPNRouter(service antivpn2.Service) antivpn.AntiVPNServer {
	return &antivpnRouter{
		service: service,
	}
}

func (a antivpnRouter) CheckVPN(ctx context.Context, request *antivpn.CheckVPNRequest) (*antivpn.CheckVPNResponse, error) {
	isVpn, err := a.service.GetResult(request.GetIp())
	if err != nil {
		return nil, err
	}
	return &antivpn.CheckVPNResponse{Proxy: isVpn}, nil
}

func (a antivpnRouter) AddWhitelist(_ context.Context, req *antivpn.CheckVPNRequest) (*emptypb.Empty, error) {
	err := a.service.AddWhitelist(req.GetIp())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (a antivpnRouter) RemoveWhitelist(_ context.Context, req *antivpn.CheckVPNRequest) (*emptypb.Empty, error) {
	err := a.service.RemoveWhitelist(req.GetIp())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (a antivpnRouter) AddBlackList(_ context.Context, req *antivpn.CheckVPNRequest) (*emptypb.Empty, error) {
	err := a.service.AddBlackList(req.GetIp())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (a antivpnRouter) RemoveBlackList(_ context.Context, req *antivpn.CheckVPNRequest) (*emptypb.Empty, error) {
	err := a.service.RemoveBlackList(req.GetIp())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
