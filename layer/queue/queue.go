package queue

type StreamProcessing interface {
	AssignNewConsumer()
	AssignNewProducer()
	GetBrokerUrls() []string
	GetDialerTimeoutSeconds() int
	GetReadTimeoutSeconds() int
	GetWriteTimeoutSeconds() int
	SetBrokerUrls(_broker_urls_ []string)
	SetDialerTimeoutSeconds(_dialer_timeout_in_seconds_ int)
	SetReadTimeoutSeconds(_read_timeout_in_seconds_ int)
	SetWriteTimeoutSeconds(_write_timeout_in_seconds_ int)
}
