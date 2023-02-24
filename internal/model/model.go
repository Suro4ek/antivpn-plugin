package model

import "gorm.io/gorm"

type IPModel struct {
	gorm.Model
	Ip    string `gorm:"unique"`
	Proxy bool
}

type PlayerModel struct {
	Uuid    string
	MCLeaks bool
}

type BlackListIP struct {
	gorm.Model
	Ip string `gorm:"unique"`
}

type WhitelistIP struct {
	gorm.Model
	Ip string `gorm:"unique"`
}
