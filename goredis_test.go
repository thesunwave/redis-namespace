package redis_namespace

import (
	"github.com/alicebob/miniredis"
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
}

func (s *ClientTestSuite) SetupTest() {
	mr, err := miniredis.Run()
	s.Require().NoError(err)

	s.redisClient = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	s.redisClientWithNamespace = NewGoRedisWithNamespace("test_namespace", s.redisClient)
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

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
