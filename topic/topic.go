package topic

import (
	"gopkg.in/redis.v5"
	"redisclient"
)

type Topic struct {
	name 	string
	client	*redis.Client
}

func NewTopic(s string) *Topic {
	return &Topic{
		name:	s,
		client:	redisclient.RedisClientInstance(),
	}
}

func (t *Topic)Read(pos int64) []byte {
	return []byte(t.client.LIndex(t.name, pos).Val())
}

func (t *Topic)Write(data []byte) {
	t.client.LPush(t.name, data)
}
