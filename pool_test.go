package redis_test

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/sliide/redis"

	"math/rand"
	"time"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type RedisTestSuite struct{}

var _ = Suite(
	&RedisTestSuite{},
)

func (s *RedisTestSuite) SetUpSuite(c *C) {
	redis.Init("localhost:6379")
}

func (s *RedisTestSuite) TearDownSuite(c *C) {
	redis.Close()
}

func (s *RedisTestSuite) TestIncrBy(c *C) {
	key := randSeq(32)

	c.Assert(redis.Set(key, 1), IsNil)
	c.Assert(redis.IncrBy(key, 10), IsNil)

	val, err := redis.Get(key)

	if err != nil {
		log.Println(err)
		c.Fail()
	}

	c.Assert(val, Equals, strconv.Itoa(11))
}

func (s *RedisTestSuite) TestExpire(c *C) {
	key := randSeq(32)

	c.Assert(redis.Set(key, "1"), IsNil)

	val, err := redis.Get(key)
	c.Assert(err, IsNil)
	c.Assert(val, Equals, "1")

	c.Assert(redis.Expire(key, 1), IsNil)
	time.Sleep(2 * time.Second)

	val, err = redis.Get(key)
	c.Assert(err, Not(IsNil))
	c.Assert(val, Equals, "")
}

func (s *RedisTestSuite) TestRPush(c *C) {
	key := randSeq(32)

	for i := 0; i < 2; i++ {
		err := redis.RPush(key, strconv.Itoa(i))
		c.Assert(err, Equals, nil)
	}

	vals, err := redis.LRange(key)
	c.Assert(err, IsNil)
	c.Assert(vals, DeepEquals, []string{"0", "1"})
}

func (s *RedisTestSuite) TestRedis(c *C) {

	key := randSeq(32)
	pop := randSeq(32)
	val := randSeq(32)
	val2 := randSeq(32)
	val3 := randSeq(32)

	key2 := randSeq(32)

	v, err := redis.Get(key)
	c.Assert(err, Not(Equals), nil)

	err = redis.Set(key, val)
	c.Assert(err, Equals, nil)

	v, err = redis.Get(key)
	c.Assert(err, Equals, nil)
	c.Assert(v, Equals, val)

	v, err = redis.Pop(pop)
	c.Assert(err, Not(Equals), nil)

	err = redis.LPush(pop, val)
	c.Assert(err, Equals, nil)

	err = redis.LPush(pop, val2)
	c.Assert(err, Equals, nil)

	err = redis.LPush(pop, val3)
	c.Assert(err, Equals, nil)

	v, err = redis.Pop(pop)
	c.Assert(err, Equals, nil)
	c.Assert(v, Equals, val3)

	v, err = redis.Pop(pop)
	c.Assert(err, Equals, nil)
	c.Assert(v, Equals, val2)

	v, err = redis.Pop(pop)
	c.Assert(err, Equals, nil)
	c.Assert(v, Equals, val)

	err = redis.Set(key2, "2")
	c.Assert(err, Equals, nil)

	err = redis.Incr(key2)
	c.Assert(err, Equals, nil)

	err = redis.Incr(key2)
	c.Assert(err, Equals, nil)

	err = redis.Incr(key2)
	c.Assert(err, Equals, nil)

	v, err = redis.Get(key2)
	c.Assert(err, Equals, nil)
	c.Assert(v, Equals, "5")
}

func (s *RedisTestSuite) TestHGet(c *C) {

	keys := []string{}
	for i := 0; i < 5; i++ {
		key := randSeq(10)
		redis.Set(key, fmt.Sprintf("%d", i))
		keys = append(keys, key)
	}

	values, err := redis.MGet(keys)
	c.Assert(err, IsNil)

	expectedValues := []string{}
	for _, key := range keys {
		val, err := redis.Get(key)
		c.Assert(err, IsNil)
		expectedValues = append(expectedValues, val)
	}

	c.Assert(len(values), Equals, 5)
	for i := 0; i < 5; i++ {
		c.Assert(values[i], Equals, expectedValues[i])
	}
}

// TODO: move it somewhere so we don't copy this everywhere
func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}