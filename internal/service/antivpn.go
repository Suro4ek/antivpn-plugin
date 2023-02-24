package service

import (
	"errors"
	"gorm.io/gorm"
	"hyneo-antivpn/internal/antivpn"
	"hyneo-antivpn/internal/model"
	"hyneo-antivpn/internal/source"
	"hyneo-antivpn/pkg/logging"
	"hyneo-antivpn/pkg/mysql"
)

type service struct {
	client mysql.Client
	log    logging.Logger
}

var sources []antivpn.Source

func NewService(client mysql.Client, log logging.Logger) antivpn.Service {
	sources = append(sources, source.NewProxyCheck())
	sources = append(sources, source.NewVPNApi())
	return &service{
		client: client,
		log:    log,
	}
}

func (s service) AddWhitelist(ip string) error {
	err := s.client.DB.
		Save(&model.WhitelistIP{Ip: ip}).
		Error
	return err
}

func (s service) RemoveWhitelist(ip string) error {
	err := s.client.DB.
		Model(&model.WhitelistIP{}).
		Delete(&model.WhitelistIP{Ip: ip}).
		Error
	return err
}

func (s service) AddBlackList(ip string) error {
	err := s.client.DB.
		Save(&model.BlackListIP{Ip: ip}).
		Error
	return err
}

func (s service) RemoveBlackList(ip string) error {
	err := s.client.DB.
		Model(&model.BlackListIP{}).
		Delete(&model.BlackListIP{Ip: ip}).
		Error
	return err
}

func (s service) GetResult(ip string) (bool, error) {
	var whitelistIP model.WhitelistIP
	if err := s.client.DB.
		Model(&model.WhitelistIP{}).
		Where(&model.WhitelistIP{Ip: ip}).
		First(&whitelistIP).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.log.Error(err)
		return false, err
	} else if err == nil {
		s.log.Info(ip + " is whitelist")
		return false, nil
	}
	var blackListIP model.BlackListIP
	if err := s.client.DB.
		Model(&model.BlackListIP{}).
		Where(&model.BlackListIP{Ip: ip}).
		First(&blackListIP).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.log.Error(err)
		return false, err
	} else if err == nil {
		s.log.Info(ip + " is blacklist")
		return true, nil
	}
	var ipModel model.IPModel
	if err := s.client.DB.
		Model(&model.IPModel{}).
		Where(&model.IPModel{Ip: ip}).
		First(&ipModel).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.log.Error(err)
		return false, err
	} else if err == nil {
		if ipModel.Proxy {
			s.log.Info(ip + " is ipmodel proxy")
			return true, nil
		}
		s.log.Info(ip + " is ipmodel not proxy")
		return false, nil
	}
	isVP := false
	for _, a := range sources {
		isVPN, err := a.GetResult(ip)
		if err != nil {
			s.log.Error(err)
			continue
		}
		isVP = isVPN
		break
	}

	ipModel = model.IPModel{
		Ip:    ip,
		Proxy: isVP,
	}
	s.log.Info(ip+" is proxy ", ipModel.Proxy)
	s.client.DB.Save(&ipModel)
	return true, nil
}
