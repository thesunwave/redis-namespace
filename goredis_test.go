package redis_namespace

import (
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/suite"
	"testing"
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

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
