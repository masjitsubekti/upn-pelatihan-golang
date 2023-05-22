package infras

import (
	"fmt"

	"gitlab.com/upn-belajar-go/configs"

	"github.com/go-redis/redis"
)

//RedisNewClient create new instance of redis
func RedisNewClient(config configs.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Cache.Redis.Primary.Host, config.Cache.Redis.Primary.Port),
		Password: config.Cache.Redis.Primary.Password,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong, err)

	return client
}
