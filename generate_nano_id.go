package utilities

import gonanoid "github.com/matoous/go-nanoid/v2"

const (
	// DefaultChars const for generate a new NanoID with default characters.
	DefaultChars string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz-"

	// UpperCaseChars const for generate a new NanoID with only upper case characters.
	UpperCaseChars string = "0123456789_ABCDEFGHIJKLMNOPQRSTUVWXYZ-"

	// UpperCaseWithoutDashesChars const for generate a new NanoID
	// with only upper case characters, but without dashes.
	UpperCaseWithoutDashesChars string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// UpperCaseWithoutNumbersChars const for generate a new NanoID
	// with only upper case characters, but without numbers.
	UpperCaseWithoutNumbersChars string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ_-"

	// LowerCaseChars const for generate a new NanoID with only lower case characters.
	LowerCaseChars string = "0123456789_abcdefghijklmnopqrstuvwxyz-"

	// LowerCaseWithoutDashesChars const for generate a new NanoID
	// with only lower case characters, but without dashes.
	LowerCaseWithoutDashesChars string = "0123456789abcdefghijklmnopqrstuvwxyz"

	// LowerCaseWithoutNumbersChars const for generate a new NanoID
	// with only lower case characters, but without numbers.
	LowerCaseWithoutNumbersChars string = "abcdefghijklmnopqrstuvwxyz_-"

	// OnlyNumbersChars const for generate a new NanoID with only numbers.
	OnlyNumbersChars string = "0123456789"
)

// GenerateNewNanoID func for generate a new random string with nanoID.
// If chars is empty, looking for DefaultNanoIDChars const.
//
// For example:
// ID length equals 18 characters with speed in 1kk IDs per second, ~81 years needed,
// in order to have a 1% probability of at least one collision.
// See: https://zelark.github.io/nano-id-cc/
func GenerateNewNanoID(chars string, size int) (string, error) {
	if chars == "" {
		// If not set, get default.
		chars = DefaultChars
	}
	if size == 0 {
		// If not set, get default.
		size = 21
	}
	id, err := gonanoid.Generate(chars, size)
	if err != nil {
		return "", err
	}
	return id, nil
}
