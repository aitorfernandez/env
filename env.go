package env

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

// Env is a env client, it's manages the communication with Redis.
type Env struct {
	Pool *redis.Pool
}

// New returns a new Env
func New() *Env {
	return &Env{
		// Pool maintains a pool of connections.
		Pool: &redis.Pool{
			Dial: func() (redis.Conn, error) {
				conn, err := redis.Dial("tcp", tierAddr())
				if err != nil {
					log.Fatalf("env.New redis.Dial %s", err)
				}
				return conn, err
			},
			MaxActive: 12000,
			MaxIdle:   80,
		},
	}
}

// DefaultEnv is the default Env and is used by Set, Hset, Get and Hget.
var DefaultEnv = New()

// Set key to hold the string value.
func (e Env) Set(key string, value interface{}) error {
	conn := e.Pool.Get()
	defer conn.Close()

	if _, err := conn.Do("SET", key, value); err != nil {
		return err
	}
	return nil
}

// Set is a wrapper around DefaultEnv.
func Set(key string, value interface{}) error {
	return DefaultEnv.Set(key, value)
}
