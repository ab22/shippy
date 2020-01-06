package env

import "os"

// LookupEnvOr retreives the value of the environment variable named by the key. If
// the key does not exists, then the fallback string value is returned.
func LookupEnvOr(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
