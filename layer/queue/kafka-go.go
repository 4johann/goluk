package queue

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

type KafkaGo struct {
	brokerUrls             []string
	clientId               string
	groupId                string
	topic                  string
	dialerTimeoutInSeconds int
	readTimeoutInSeconds   int
	writeTimeoutInSeconds  int
	writer                 *kafka.Writer
	reader                 *kafka.Reader
	dialer                 *kafka.Dialer
	configurationWriter    kafka.WriterConfig
	configurationReader    kafka.ReaderConfig
}

func (k *KafkaGo) AssignNewProducer() {
	k.setDialer()
	k.setProducerConfiguration()
	k.writer = kafka.NewWriter(k.configurationWriter)
}

func (k *KafkaGo) AssignNewConsumer() {
	k.setDialer()
	k.setConsumerConfiguration()
	k.reader = kafka.NewReader(k.configurationReader)
}

func (k *KafkaGo) PushAMesssage(ctx context.Context, kafkaMessage KafkaMessage) error {
	message := kafka.Message{
		Key:   kafkaMessage.Key,
		Value: kafkaMessage.Value,
		Time:  time.Now(),
	}
	return k.writer.WriteMessages(ctx, message)
}

func (k *KafkaGo) PullMessage(ctx context.Context) (KafkaMessage, error) {
	message, readError := k.reader.ReadMessage(ctx)
	return translateToKafkaMessage(message), readError
}

func (k *KafkaGo) CloseConsumer() error {
	return k.reader.Close()
}

func (k *KafkaGo) CloseProducer() error {
	return k.writer.Close()
}

func (k *KafkaGo) SetTopic(_topic string) {
	k.topic = _topic
}

func (k *KafkaGo) GetTopic() string {
	return k.topic
}

func (k *KafkaGo) SetBrokerUrls(_brokerUrls []string) {
	k.brokerUrls = _brokerUrls
}

func (k *KafkaGo) GetBrokerUrls() []string {
	return k.brokerUrls
}

func (k *KafkaGo) SetDialerTimeoutInSeconds(_dialerTimeoutInSeconds int) {
	k.dialerTimeoutInSeconds = _dialerTimeoutInSeconds
}

func (k *KafkaGo) GetDialerTimeoutInSeconds() int {
	return k.dialerTimeoutInSeconds
}

func (k *KafkaGo) SetReadTimeoutInSeconds(_readTimeoutInSeconds int) {
	k.readTimeoutInSeconds = _readTimeoutInSeconds
}

func (k *KafkaGo) GetReadTimeoutInSeconds() int {
	return k.readTimeoutInSeconds
}

func (k *KafkaGo) SetWriteTimeoutInSeconds(_writeTimeoutInSeconds int) {
	k.writeTimeoutInSeconds = _writeTimeoutInSeconds
}

func (k *KafkaGo) GetWriteTimeoutInSeconds() int {
	return k.writeTimeoutInSeconds
}

func translateToKafkaMessage(message kafka.Message) KafkaMessage {
	kafka_message := KafkaMessage{
		Topic: message.Topic,
		Key:   message.Key,
		Value: message.Value,
		Time:  message.Time,
	}
	return kafka_message
}

func (k *KafkaGo) setDialer() {
	k.dialer = &kafka.Dialer{
		Timeout:  time.Duration(k.dialerTimeoutInSeconds) * time.Second,
		ClientID: k.clientId,
	}
}

func (k *KafkaGo) setProducerConfiguration() {
	k.configurationWriter = kafka.WriterConfig{
		Brokers:          k.brokerUrls,
		Topic:            k.topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           k.dialer,
		WriteTimeout:     time.Duration(k.writeTimeoutInSeconds) * time.Second,
		ReadTimeout:      time.Duration(k.readTimeoutInSeconds) * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
}

func (k *KafkaGo) setConsumerConfiguration() {
	k.configuration_reader = kafka.ReaderConfig{
		Brokers:         k.brokerUrls,
		GroupID:         k.groupId,
		Topic:           k.topic,
		Dialer:          k.dialer,
		MaxWait:         1 * time.Second,
		ReadLagInterval: -1,
	}
}
