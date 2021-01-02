package configuration

type RemoteApplicationConfiguration struct {
	application_name string `json:"name"`
	application_port int    `json:"port"`
	version          int    `version`
}

func (r *RemoteApplicationConfiguration) SetNameApplication(_name_ string) {
	r.application_name = _name_
}

func (r *RemoteApplicationConfiguration) SetPort(_port_ int) {
	r.application_port = _port_
}

func (r *RemoteApplicationConfiguration) SetVersion(_version_ int) {
	r.version = _version_
}

func (r *RemoteApplicationConfiguration) GetNameApplication() string {
	return r.application_name
}

func (r *RemoteApplicationConfiguration) GetPort() int {
	return r.application_port
}

func (r *RemoteApplicationConfiguration) GetVersion() int {
	return r.version
}
