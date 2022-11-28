package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

const (
	MONTH = 30 * 24 * time.Hour
)

func Connect(addr, password string) *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})
	_, err := c.Ping().Result()
	if err != nil {
		fmt.Println("redis kill")
		c.Close()
		return nil
	}
	return c
}

func main() {
	var addr = "127.0.0.1:6379"
	var password = ""

	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})

	p, err := c.Ping().Result()
	if err != nil {
		fmt.Println("redis kill")
	}
	fmt.Println(p)

	_, err = c.Set("key2", "value", MONTH).Result()
	// c.Do("SET", "key", "duzhenxun2")
	rs := c.Do("GET", "key").Val()
	fmt.Println(rs)
	c.Close()
}
