package Redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Client *redis.Client

func init() {
	// Replaces with your configuration information
	Client = redis.NewClient(&redis.Options{
		Addr: "redis:6379", // Active for docker
		// Addr: ":6379", // Active to prod in local
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("redis init done")

	pong, err := Client.Ping().Result()
	fmt.Println(pong, err)
}