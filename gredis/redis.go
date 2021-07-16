package gredis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisPool *redis.Pool

const (
	RedisMaxIdle        = 10  //最大空闲连接数
	RedisIdleTimeoutSec = 240 //连接保持超时时间（秒）
)

// Setup Initialize the Redis instance
func Setup(redisURL string) {
	RedisPool = &redis.Pool{
		MaxIdle:     RedisMaxIdle,
		IdleTimeout: RedisIdleTimeoutSec,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(redisURL)
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

func Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := RedisPool.Get()
	defer conn.Close()
	return conn.Do(commandName, args...)
}

// GetLock 获取Redis锁
func GetLock(lockKey string, expiration int64) (bool, error) {
	conn := RedisPool.Get()
	defer conn.Close()
	result, err := redis.Int64(conn.Do("INCR", lockKey))
	if err != nil {
		return false, err
	}
	if result > 1 {
		return false, nil
	}
	_, err = conn.Do("EXPIRE", lockKey, expiration)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ReleaseLock 释放 Redis 锁
func ReleaseLock(lockKey string) bool {
	conn := RedisPool.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", lockKey)
	if err != nil {
		return true
	}
	return true
}

func Set(key string, value interface{}) (bool, error) {
	result, err := redis.String(Do("SET", key, value))
	if result == "OK" {
		return true, nil
	}
	return false, err
}

// Exists check a key
func Exists(key string) (bool, error) {
	return redis.Bool(Do("EXISTS", key))
}

// Get get a key
func Get(key string) (interface{}, error) {
	return Do("GET", key)
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	return redis.Bool(Do("DEL", key))
}

func GetInt64(key string) (int64, error) {
	return redis.Int64(Get(key))
}

func GetUint64(key string) (uint64, error) {
	return redis.Uint64(Get(key))
}

func GetInt(key string) (int, error) {
	return redis.Int(Get(key))
}

func GetString(key string) (string, error) {
	return redis.String(Get(key))
}

func GetBoolean(key string) (bool, error) {
	return redis.Bool(Get(key))
}

func GetFloat64(key string) (float64, error) {
	return redis.Float64(Get(key))
}
