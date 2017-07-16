package consumer

import (
	"topic"
)

type Consumer struct {
	topic 	*topic.Topic
}

func NewConsumer(t string) *Consumer {
	return &Consumer{
		topic : topic.NewTopic(t),
	}
}
func (c *Consumer)Consume(pos int64) (data []byte) {
	data = c.topic.Read(pos)	// to do :overflow
	return
}