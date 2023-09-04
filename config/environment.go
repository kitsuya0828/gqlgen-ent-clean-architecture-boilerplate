package config

import (
	"os"
)

// Application Environment name
const (
	Development = "development"
	Production  = "production"
	Test        = "test"
	E2E         = "e2e"
)

// IsDev returns APP_ENV in development mode
func IsDev() bool {
	// log.Printf("os.Getenv(APP_ENV): %s", os.Getenv("APP_ENV"))
	return os.Getenv("APP_ENV") == Development
}

// IsProd returns APP_ENV in production mode
func IsProd() bool {
	// log.Printf("os.Getenv(APP_ENV): %s", os.Getenv("APP_ENV"))
	return os.Getenv("APP_ENV") == Production
}

// IsTest returns APP_ENV in test mode
func IsTest() bool {
	return os.Getenv("APP_ENV") == Test
}

// IsE2E returns APP_ENV in e2e mode
func IsE2E() bool {
	return os.Getenv("APP_ENV") == E2E
}
