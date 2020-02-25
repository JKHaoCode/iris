package redis

import (
	"encoding/json"
	"time"

	"github.com/gomodule/redigo/redis"

	config "github.com/spf13/viper"
)

var RedisConn *redis.Pool
// .DefaultRedisIdleTimeout
func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:     config.GetInt("redis.MaxIdle"),
		MaxActive:   config.GetInt("redis.MaxIdle"),
		IdleTimeout: time.Duration(5) * time.Minute,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.GetString("redis.Host"))
			if err != nil {
				return nil, err
			}
			if config.GetString("redis.Password") != "" {
				if _, err := c.Do("AUTH", config.GetString("redis.Password")); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}

func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}