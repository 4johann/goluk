package queue

import (
	"time"
)

type KafkaMessage struct {
	Topic string
	Key   []byte
	Value []byte
	Time  time.Time
}
