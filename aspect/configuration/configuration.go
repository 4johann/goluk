package configuration

type Configuration interface {
	SetNameApplication(_name_ string)
	SetPort(_port_ int)
	SetVersion(_version_ int)
	GetNameApplication()
	GetPort()
	GetVersion()
}
