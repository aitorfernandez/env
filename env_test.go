package env_test

import (
	"testing"

	"github.com/aitorfernandez/env"
)

func TestGet(t *testing.T) {
	env.Set("foo", "bar")

	if got, _ := env.Get("foo"); got != "bar" {
		t.Errorf("got %v want 'bar'", got)
	}
}

func TestHget(t *testing.T) {
	env.Hset("hash", "hello", "test")

	if got, _ := env.Hget("hash", "hello"); got != "test" {
		t.Errorf("got %v want 'test'", got)
	}
}
