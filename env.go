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

// Hset sets field in the hash stored at key to value.
func (e Env) Hset(key, field string, value interface{}) error {
	conn := e.Pool.Get()
	defer conn.Close()

	if _, err := conn.Do("HSET", key, field, value); err != nil {
		return err
	}
	return nil
}

// Hset is a wrapper around DefaultEnv.
func Hset(key, field string, value interface{}) error {
	return DefaultEnv.Hset(key, field, value)
}

// Get returns the value of key.
func (e Env) Get(key string) (string, error) {
	conn := e.Pool.Get()
	defer conn.Close()

	return redis.String(conn.Do("GET", key))
}

// Get is a wrapper around DefaultEnv.
func Get(key string) (string, error) {
	return DefaultEnv.Get(key)
}

// Hget returns the value associated with field in the hash stored at key.
func (e Env) Hget(key, field string) (string, error) {
	conn := e.Pool.Get()
	defer conn.Close()

	return redis.String(conn.Do("HGET", key, field))
}

// Hget is a wrapper around DefaultEnv.
func Hget(key, field string) (string, error) {
	return DefaultEnv.Hget(key, field)
}

// Sadd adds the specified members to the set stored at key
func (e Env) Sadd(key string, member interface{}) error {
	conn := e.Pool.Get()
	defer conn.Close()

	if _, err := conn.Do("SADD", key, member); err != nil {
		return err
	}
	return nil
}

// Sadd is a wrapper around DefaultEnv.
func Sadd(key string, member interface{}) error {
	return DefaultEnv.Sadd(key, member)
}

// Smembers returns all the members of the set value stored at key.
func (e Env) Smembers(key string) ([]string, error) {
	conn := e.Pool.Get()
	defer conn.Close()

	return redis.Strings(conn.Do("SMEMBERS", key))
}

// Smembers is a wrapper arounf DefaultEnv.
func Smembers(key string) ([]string, error) {
	return DefaultEnv.Smembers(key)
}

// MustHget is a helper that wraps a call and log.Fatal if error.
func MustHget(key, field string) string {
	v, err := Hget(key, field)
	if err != nil {
		log.Fatalf("env.MustHget Hget %s %s, %s", key, field, err)
	}
	return v
}
