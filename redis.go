package redis

var client Client = nil

func SetClient(redisClient Client) {
	client = redisClient
}

func Close() {
	client.Close()
}

func Get(key string) (string, error) {
	return client.Get(key)
}

func MGet(keys []string) ([]string, error) {
	return client.MGet(keys)
}

func Set(key string, value interface{}) error {
	return client.Set(key, value)
}

func SetEx(key string, expire int, value interface{}) error {
	return client.SetEx(key, expire, value)
}

func Expire(key string, seconds int) error {
	return client.Expire(key, seconds)
}

func Del(key string) error {
	return client.Del(key)
}

func LPush(key string, value string) error {
	return client.LPush(key, value)
}

func RPush(key string, value string) error {
	return client.RPush(key, value)
}

func LRange(key string) ([]string, error) {
	return client.LRange(key)
}

func LPop(key string) (string, error) {
	return client.LPop(key)
}

func Pop(key string) (string, error) {
	return client.Pop(key)
}

func Incr(key string) error {
	return client.Incr(key)
}

func IncrBy(key string, inc interface{}) (interface{}, error) {
	return client.IncrBy(key, inc)
}

func ZAdd(key string, score float64, value interface{}) (int, error) {
	return client.ZAdd(key, score, value)
}

func ZCount(key string, min interface{}, max interface{}) (int, error) {
	return client.ZCount(key, min, max)
}
