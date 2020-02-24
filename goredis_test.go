package redis_namespace

import (
	"fmt"
	_ "github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ClientTestSuite struct {
	suite.Suite

	redisClient              GoRedis
	redisClientWithNamespace *Client

	namespace string
}

func (s *ClientTestSuite) SetupTest() {
	s.redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	s.namespace = "test_namespace"

	s.redisClientWithNamespace = NewGoRedisWithNamespace(s.namespace, s.redisClient)
	s.redisClient.FlushAll()
}

func (s *ClientTestSuite) TestPing() {
	_, err := s.redisClientWithNamespace.Ping().Result()
	s.Nil(err)
}

func (s *ClientTestSuite) TestGet_Set() {
	resultSet := s.redisClientWithNamespace.Set("key", "value", 0)
	s.Nil(resultSet.Err())

	keys := s.redisClientWithNamespace.Keys("*")
	s.Equal(1, len(keys.Val()))
	s.Equal("key", keys.Val()[0])

	resultGet := s.redisClientWithNamespace.Get("key")
	s.Equal(resultGet.Val(), "value")
}

func (s *ClientTestSuite) TestClient_Del() {
	resultSet := s.redisClientWithNamespace.Set("key", "value", 0)
	s.Nil(resultSet.Err())
	resultSet = s.redisClientWithNamespace.Set("key2", "value", 0)
	s.Nil(resultSet.Err())

	result := s.redisClientWithNamespace.Del("key", "key2")
	s.EqualValues(2, result.Val())
}

func (s *ClientTestSuite) TestClient_Unlink() {
	resultSet := s.redisClientWithNamespace.Set("key", "value", 0)
	s.Nil(resultSet.Err())
	resultSet = s.redisClientWithNamespace.Set("key2", "value", 0)
	s.Nil(resultSet.Err())

	result := s.redisClientWithNamespace.Unlink("key", "key2")
	s.Nil(result.Err())
	s.EqualValues(2, result.Val())
}

func (s *ClientTestSuite) TestClient_Dump() {
	resultSet := s.redisClientWithNamespace.Set("key", 10, 0)
	s.Nil(resultSet.Err())

	result := s.redisClientWithNamespace.Dump("key")
	s.Nil(result.Err())
	s.EqualValues("\u0000\xC0\n\t\u0000\xBEm\u0006\x89Z(\u0000\n", result.Val())
}

func (s *ClientTestSuite) TestClient_Exists() {
	resultSet := s.redisClientWithNamespace.Set("key", 10, 0)
	s.Nil(resultSet.Err())

	result := s.redisClientWithNamespace.Exists("key", "nosuchkey")
	s.Nil(result.Err())
	s.EqualValues(1, result.Val())
}

func (s *ClientTestSuite) TestClient_Expire() {
	resultSet := s.redisClientWithNamespace.Set("key", 10, 0)
	s.Nil(resultSet.Err())

	result := s.redisClientWithNamespace.Expire("key", 10*time.Second)
	s.Nil(result.Err())
	s.EqualValues(true, result.Val())
}

func (s *ClientTestSuite) TestClient_ExpireAt() {
	resultSet := s.redisClientWithNamespace.Set("key", 10, 0)
	s.Nil(resultSet.Err())

	result := s.redisClientWithNamespace.ExpireAt("key", time.Now().Add(10*time.Second))
	s.Nil(result.Err())
	s.EqualValues(true, result.Val())
}

func (s *ClientTestSuite) TestClient_ObjectRefCount() {
	resultSet := s.redisClientWithNamespace.LPush("mylist", "value")
	s.Nil(resultSet.Err())

	result := s.redisClientWithNamespace.ObjectRefCount("mylist")
	s.Nil(result.Err())
	s.EqualValues(1, result.Val())
}

func (s *ClientTestSuite) TestClient_ObjectEncoding() {
	resultSet := s.redisClientWithNamespace.Set("key", "value", 0)
	s.Nil(resultSet.Err())

	result := s.redisClientWithNamespace.ObjectEncoding("key")
	s.Nil(result.Err())
	s.EqualValues("embstr", result.Val())
}

func (s *ClientTestSuite) TestClient_ObjectIdleTime() {
	resultSet := s.redisClientWithNamespace.Set("key", "value", 0)
	s.Nil(resultSet.Err())

	result := s.redisClientWithNamespace.ObjectIdleTime("key")
	s.Nil(result.Err())
	s.EqualValues(0, result.Val())
}

func (s *ClientTestSuite) TestClient_Persist() {
	resultSet := s.redisClientWithNamespace.Set("key", "value", 10*time.Second)
	s.Nil(resultSet.Err())

	result := s.redisClientWithNamespace.Persist("key")
	s.Nil(result.Err())
	s.EqualValues(true, result.Val())
}

func (s *ClientTestSuite) TestClient_PExpire() {
	resultSet := s.redisClientWithNamespace.Set("key", 10, 0)
	s.Nil(resultSet.Err())

	result := s.redisClientWithNamespace.PExpire("key", 10*time.Second)
	s.Nil(result.Err())
	s.EqualValues(true, result.Val())
}

func (s *ClientTestSuite) TestClient_PExpireAt() {
	resultSet := s.redisClientWithNamespace.Set("key", 10, 0)
	s.Nil(resultSet.Err())

	result := s.redisClientWithNamespace.PExpireAt("key", time.Now().Add(10*time.Second))
	s.Nil(result.Err())
	s.EqualValues(true, result.Val())
}

func (s *ClientTestSuite) TestClient_PTTL() {
	// Need to mock that case
	//resultSet := s.redisClientWithNamespace.Set("key", 10, 1 * time.Second)
	//s.Nil(resultSet.Err())
	//
	//result := s.redisClientWithNamespace.PTTL("key")
	//s.Nil(result.Err())
	//s.EqualValues(999, result.Val()) // => we can't do that because of using real Redis connection
}

func (s *ClientTestSuite) TestClient_Rename() {
	resultSet := s.redisClientWithNamespace.Set("key", 10, 0)
	s.Nil(resultSet.Err())

	result := s.redisClientWithNamespace.Rename("key", "newkey")
	s.Nil(result.Err())
	s.EqualValues("OK", result.Val())
}

func (s *ClientTestSuite) TestClient_RenameNX() {
	resultSet := s.redisClientWithNamespace.Set("key", 10, 0)
	s.Nil(resultSet.Err())

	result := s.redisClientWithNamespace.RenameNX("key", "newkey")
	s.Nil(result.Err())
	s.EqualValues(true, result.Val())

	key := s.redisClient.Get(fmt.Sprintf("%s:%s", s.namespace, "newkey"))
	s.Nil(key.Err())
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
