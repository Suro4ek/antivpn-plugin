package main

import (
	"context"
	"google.golang.org/grpc"
	"hyneo-antivpn/internal/config"
	"hyneo-antivpn/internal/model"
	router2 "hyneo-antivpn/internal/router"
	service2 "hyneo-antivpn/internal/service"
	"hyneo-antivpn/pkg/logging"
	"hyneo-antivpn/pkg/mysql"
	"hyneo-antivpn/protos/antivpn"
	"log"
	"net"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	cfg := config.GetConfig()
	client := mysql.NewClient(context.Background(), 5, cfg.MySQL)
	client.DB.AutoMigrate(&model.IPModel{})
	client.DB.AutoMigrate(&model.BlackListIP{})
	client.DB.AutoMigrate(&model.WhitelistIP{})
	service := service2.NewService(*client, logger)
	router := router2.NewAntiVPNRouter(service)
	addr := "0.0.0.0:" + cfg.GRPCPort
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	antivpn.RegisterAntiVPNServer(s, router)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
