package env

import "os"

// tierAddr returns the value of the environment variable. Otherwise
// the returner value will be the default dev address environment.
func tierAddr() string {
	val, ok := os.LookupEnv("REDIS_ENV_ADDRESS")
	if !ok {
		return "0.0.0.0:6379"
	}
	return val
}
