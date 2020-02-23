package redis

import (
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
	"sync"
)

var rd *redis.Database
var once sync.Once

func Singleton() *redis.Database {
	// redis.
	once.Do(func() {
		rd = redis.New(service.Config{
			//Network:   "tcp",
			//Addr:      config.GetRedisAddr(),
			//Timeout:   time.Duration(30) * time.Second,
			//MaxActive: 10,
			//Password:  config.GetRedisPwd(),
			//Database:  "",
			//Prefix:    "",
			//Delim:     "-",
			//Driver:    redis.Redigo(), // redis.Radix() can be used instead.
			Network:     "tcp",
			Addr:        "127.0.0.1:6379",
			Password:    "",
			Database:    "",
			MaxIdle:     0,
			MaxActive:   100,
			IdleTimeout: service.DefaultRedisIdleTimeout,// time.Duration(24) * time.Hour,
			Prefix:      "",
		})
	})
	return rd
}
