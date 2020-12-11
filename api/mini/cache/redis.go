package cache

import (
	"github.com/go-redis/redis/v8"
	"jjsd-go-api/api/mini/conf"
	"sync"
)

const (
	Prefix = "MINI:"
	Ver    = "1.0:"
)

var (
	cli     *redis.Client
	cliOnce sync.Once
)

func newClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     getAddr(),
		Password: "",
		DB:       0,
	})
}

func GetClient() *redis.Client {
	cliOnce.Do(func() {
		cli = newClient()
	})

	return cli
}

func getAddr() string {
	return conf.Conf.Redis.Host + ":" + conf.Conf.Redis.Port
}

func PrefixKey(key string) string {
	return Prefix + Ver + key
}
