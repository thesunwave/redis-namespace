package redis_namespace

import (
	"github.com/go-redis/redis/v7"
	"time"
)

type Client struct {
	GoRedis
	client    GoRedis
	namespace string
}

type GoRedis interface {
	redis.Cmdable
}

func NewGoRedisWithNamespace(namespace string, client GoRedis) *Client {
	return &Client{client, client, namespace + ":"}
}

func (c *Client) Get(key string) *redis.StringCmd {
	return c.client.Get(c.namespace + key)
}

func (c *Client) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return c.client.Set(c.namespace+key, value, expiration)
}
