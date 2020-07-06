package redis_namespace

import (
	"github.com/go-redis/redis/v7"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

type Client struct {
	redis.Cmdable
	namespace string
}

func NewGoRedisWithNamespace(namespace string, client redis.Cmdable) *Client {
	return &Client{client, namespace + ":"}
}

func (c *Client) Del(keys ...string) *redis.IntCmd {
	var wrappedKeys []string

	for _, key := range keys {
		wrappedKeys = append(wrappedKeys, c.namespace+key)
	}
	return c.Cmdable.Del(wrappedKeys...)
}

func (c *Client) Unlink(keys ...string) *redis.IntCmd {
	var wrappedKeys []string

	for _, key := range keys {
		wrappedKeys = append(wrappedKeys, c.namespace+key)
	}
	return c.Cmdable.Unlink(wrappedKeys...)
}

func (c *Client) Dump(key string) *redis.StringCmd {
	return c.Cmdable.Dump(c.namespace + key)
}

func (c *Client) Exists(keys ...string) *redis.IntCmd {
	var wrappedKeys []string

	for _, key := range keys {
		wrappedKeys = append(wrappedKeys, c.namespace+key)
	}
	return c.Cmdable.Exists(wrappedKeys...)
}

func (c *Client) Expire(key string, expiration time.Duration) *redis.BoolCmd {
	return c.Cmdable.Expire(c.namespace+key, expiration)
}

func (c *Client) ExpireAt(key string, tm time.Time) *redis.BoolCmd {
	return c.Cmdable.ExpireAt(c.namespace+key, tm)
}

func (c *Client) Keys(pattern string) *redis.StringSliceCmd {
	var unwrappedKeys []string
	keys := c.Cmdable.Keys(c.namespace + pattern)

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

func (c *Client) ObjectRefCount(key string) *redis.IntCmd {
	return c.Cmdable.ObjectRefCount(c.namespace + key)
}

func (c *Client) ObjectEncoding(key string) *redis.StringCmd {
	return c.Cmdable.ObjectEncoding(c.namespace + key)
}

func (c *Client) ObjectIdleTime(key string) *redis.DurationCmd {
	return c.Cmdable.ObjectIdleTime(c.namespace + key)
}

func (c *Client) Persist(key string) *redis.BoolCmd {
	return c.Cmdable.Persist(c.namespace + key)
}

func (c *Client) PExpire(key string, expiration time.Duration) *redis.BoolCmd {
	return c.Cmdable.PExpire(c.namespace+key, expiration)
}

func (c *Client) PExpireAt(key string, tm time.Time) *redis.BoolCmd {
	return c.Cmdable.PExpireAt(c.namespace+key, tm)
}

func (c *Client) PTTL(key string) *redis.DurationCmd {
	return c.Cmdable.PTTL(c.namespace + key)
}

func (c *Client) Rename(key, newkey string) *redis.StatusCmd {
	return c.Cmdable.Rename(c.namespace+key, c.namespace+newkey)
}

func (c *Client) RenameNX(key, newkey string) *redis.BoolCmd {
	return c.Cmdable.RenameNX(c.namespace+key, c.namespace+newkey)
}

func (c *Client) Get(key string) *redis.StringCmd {
	return c.Cmdable.Get(c.namespace + key)
}

func (c *Client) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return c.Cmdable.Set(c.namespace+key, value, expiration)
}

func (c *Client) LPush(key string, values ...interface{}) *redis.IntCmd {
	return c.Cmdable.LPush(c.namespace+key, values...)
}
