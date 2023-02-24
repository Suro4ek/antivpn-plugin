package antivpn

type Service interface {
	GetResult(ip string) (bool, error)
	AddWhitelist(ip string) error
	RemoveWhitelist(ip string) error
	AddBlackList(ip string) error
	RemoveBlackList(ip string) error
}
