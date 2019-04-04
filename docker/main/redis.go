package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	//Pool Redis for reuse
	Pool *redis.Pool
)

func SetValueByKey(key string, value []byte) error {
	defer recoverRedis()
	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}
	return err
}

//GetKey will return key from redis
func GetValueByKey(key string) ([]byte, error) {
	defer recoverRedis()
	conn := Pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}

func ping() (result bool) {
	result = false
	defer recoverRedis()
	conn := Pool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		return
	}
	result = true
	return
}

func recoverRedis() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func init() {
	defer recoverRedis()
	rediHost := getEnv("REDIS_HOST", "localhost")
	redisPort := getEnv("REDIS_PORT", "6379")
	Pool = newPool(fmt.Sprintf("%s:%s", rediHost, redisPort))
}

func newPool(server string) *redis.Pool {

	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
