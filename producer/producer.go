package producer

import (
	"topic"
)

type Producer struct {
	topic *topic.Topic
}

func NewProducer(name string) *Producer {
	return &Producer{
		topic : topic.NewTopic(name),
	}
}
func (p *Producer)Produce(data []byte) {
	p.topic.Write(data)
}