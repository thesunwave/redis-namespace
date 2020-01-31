package redis_namespace

import (
	"fmt"
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

func (suite *ClientTestSuite) SetupTest() {
	mr, err := miniredis.Run()
	suite.Require().NoError(err)

	suite.redisClient = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	suite.redisClientWithNamespace = NewGoRedisWithNamespace("test_namespace", suite.redisClient)
}

func (suite *ClientTestSuite) TestNewClient() {
	_, err := suite.redisClientWithNamespace.Ping().Result()
	suite.Require().NoError(err)

	resultSet := suite.redisClientWithNamespace.Set("key", "value", 0)
	suite.Nil(resultSet.Err())
	resultGet := suite.redisClientWithNamespace.Get("key")
	fmt.Print(resultGet)
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
