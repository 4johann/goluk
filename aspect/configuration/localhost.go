package configuration

type LocalApplicationConfiguration struct {
	application_name                  string                           `json:"name"`
	application_port                  int                              `json:"port"`
	version                           int                              `json:"version"`
	remote_applications_configuration []RemoteApplicationConfiguration `json:"friends"`
	debug_mode                        bool                             `json:"debug"`
}

func (l *LocalApplicationConfiguration) SetAllRemoteAplications(applications []RemoteApplicationConfiguration) {
	l.remote_applications_configuration = applications
}

func (l *LocalApplicationConfiguration) GetRemoteApplicationByName(application_name string) RemoteApplicationConfiguration {
	var application_configuration RemoteApplicationConfiguration
	for _, remote_application_configuration := range l.remote_applications_configuration {
		if remote_application_configuration.application_name == application_name {
			application_configuration = remote_application_configuration
			break
		}
	}
	return application_configuration
}

func (l *LocalApplicationConfiguration) SetDebugMode(_debug_ bool) {
	l.debug_mode = _debug_
}

func (l *LocalApplicationConfiguration) SetNameApplication(_name_ string) {
	l.application_name = _name_
}

func (l *LocalApplicationConfiguration) SetPort(_port_ int) {
	l.application_port = _port_
}

func (l *LocalApplicationConfiguration) SetVersion(_version_ int) {
	l.version = _version_
}

func (l *LocalApplicationConfiguration) GetDebugMode() bool {
	return l.debug_mode
}

func (l *LocalApplicationConfiguration) GetNameApplication() string {
	return l.application_name
}

func (l *LocalApplicationConfiguration) GetPort() int {
	return l.application_port
}

func (l *LocalApplicationConfiguration) GetVersion() int {
	return l.version
}
