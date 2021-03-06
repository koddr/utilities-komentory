package utilities

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// TokenValidateExpireTime func for validating given JWT token with expire time.
func TokenValidateExpireTime(ctx *fiber.Ctx) (*TokenMetaData, error) {
	// Get claims from JWT.
	claims, errExtractTokenMetaData := ExtractTokenMetaData(ctx)
	if errExtractTokenMetaData != nil {
		// Return JWT parse error.
		return nil, errExtractTokenMetaData
	}

	// Checking, if now time greather than expiration from JWT.
	if time.Now().Unix() > claims.Expire {
		// Return unauthorized (permission denied) error message.
		return nil, fmt.Errorf(GenerateErrorMessage(401, "token", "was expired"))
	}

	return claims, nil
}

// TokenValidateExpireTimeAndCredentials func for validating given JWT token with expire time and credentials.
func TokenValidateExpireTimeAndCredentials(ctx *fiber.Ctx, credentials []string) (*TokenMetaData, error) {
	// Get claims from JWT.
	claims, errExtractTokenMetaData := ExtractTokenMetaData(ctx)
	if errExtractTokenMetaData != nil {
		// Return JWT parse error.
		return nil, errExtractTokenMetaData
	}

	// Checking, if now time greather than expiration from JWT.
	if time.Now().Unix() > claims.Expire {
		// Return unauthorized (permission denied) error message.
		return nil, fmt.Errorf(GenerateErrorMessage(401, "token", "was expired"))
	}

	// Checking, if list of credentials has needed credential.
	for _, credential := range credentials {
		// Return unauthorized (permission denied) error message.
		if !SearchStringInArray(credential, claims.Credentials) {
			return nil, fmt.Errorf(GenerateErrorMessage(401, "token", "no required credentials"))
		}
	}

	return claims, nil
}
