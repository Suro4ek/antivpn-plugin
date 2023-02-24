package storage

import "hyneo-antivpn/internal/antivpn"

type storage struct {
}

func NewStorage() antivpn.Storage {
	return &storage{}
}
