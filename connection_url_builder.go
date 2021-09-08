package utilities

import (
	"fmt"
	"os"
)

const (
	// DBConnectionName const for the name of the Database connection.
	DBConnectionName string = "database"
	// RedisConnectionName const for the name of the Redis connection.
	RedisConnectionName string = "redis"
	// FiberConnectionName const for the name of the Fiber connection.
	FiberConnectionName string = "fiber"
)

// ConnectionURLBuilder func for building URL connection by given name.
// Allowed: database, redis, fiber
func ConnectionURLBuilder(name string) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch name {
	case DBConnectionName:
		// URL for Database connection.
		url = os.Getenv("DB_URL")
	case RedisConnectionName:
		// URL for Redis connection.
		url = os.Getenv("REDIS_URL")
	case FiberConnectionName:
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", name)
	}

	// Return connection URL.
	return url, nil
}
