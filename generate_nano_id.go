package utilities

import (
	"os"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

// GenerateNewNanoID func for generate a new random string with nanoID.
// If chars is empty, looking for NANOID_CHARS_STRING environment variable.
//
// For example:
// ID length equals 18 characters with speed in 1kk IDs per second, ~81 years needed,
// in order to have a 1% probability of at least one collision.
// See: https://zelark.github.io/nano-id-cc/
func GenerateNewNanoID(chars string, size int) (string, error) {
	if chars == "" {
		// If not set, default to environment.
		chars = os.Getenv("NANOID_CHARS_STRING")
	}
	id, err := gonanoid.Generate(chars, size)
	if err != nil {
		return "", err
	}
	return id, nil
}
