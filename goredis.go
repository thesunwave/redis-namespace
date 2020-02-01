package redis_namespace

import (
	"github.com/go-redis/redis/v7"
	"reflect"
	"strings"
	"time"
	"unsafe"
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

func (c *Client) Keys(pattern string) *redis.StringSliceCmd {
	var unwrappedKeys []string
	keys := c.client.Keys(c.namespace + pattern)

	for _, key := range keys.Val() {
		unwrappedKeys = append(unwrappedKeys, strings.TrimPrefix(key, c.namespace))
	}

	pointerVal := reflect.ValueOf(keys)
	val := reflect.Indirect(pointerVal)
	member := val.FieldByName("val")
	ptrToY := unsafe.Pointer(member.UnsafeAddr())
	realPtrToY := (*[]string)(ptrToY)
	*realPtrToY = unwrappedKeys

	return keys
}
