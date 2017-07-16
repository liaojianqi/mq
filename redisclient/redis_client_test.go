package redisclient_test

import(
	"testing"
	"redisclient"
)
func TestRedisClientInstance(t *testing.T) {
	client := redisclient.RedisClientInstance()
	_, err := client.Ping().Result()
	if err != nil {
		t.Error("RedisClientInstance error")
	}
}