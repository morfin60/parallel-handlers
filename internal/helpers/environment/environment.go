package environment

import (
	"log"
	"os"
	"strconv"
)

// Get int32 value from environment
func GetInt32(key string, defaultValue int32) int32 {
	envValue := GetString(key, "")

	if envValue == "" {
		return defaultValue
	}

	value, err := strconv.ParseInt(envValue, 10, 32)
	if err != nil {
		log.Printf("Failed to parse value %s: %s", envValue, err.Error())

		return defaultValue
	}

	return int32(value)
}

// Get int64 value from environment
func GetInt64(key string, defaultValue int64) int64 {
	envValue := GetString(key, "")

	if envValue == "" {
		return defaultValue
	}

	value, err := strconv.ParseInt(envValue, 10, 64)
	if err != nil {
		log.Printf("Failed to parse value %s: %s", envValue, err.Error())

		return defaultValue
	}

	return value
}

// Get float32 value from environment
func GetFloat32(key string, defaultValue float32) float32 {
	envValue := GetString(key, "")

	if envValue == "" {
		return defaultValue
	}

	value, err := strconv.ParseFloat(envValue, 64)
	if err != nil {
		log.Printf("Failed to parse value %s: %s", envValue, err.Error())

		return defaultValue
	}

	return float32(value)
}

// Get float64 value from environment
func GetFloat64(key string, defaultValue float64) float64 {
	envValue := GetString(key, "")

	if envValue == "" {
		return defaultValue
	}

	value, err := strconv.ParseFloat(envValue, 64)
	if err != nil {
		log.Printf("Failed to parse value %s: %s", envValue, err.Error())

		return defaultValue
	}

	return value
}

// Get boolean value from environment
func GetBool(key string, defaultValue bool) bool {
	envValue := GetString(key, "")

	if envValue == "" {
		return defaultValue
	}

	value, err := strconv.ParseBool(envValue)
	if err != nil {
		log.Printf("Failed to parse value %s: %s", envValue, err.Error())

		return defaultValue
	}

	return value
}

// Get environment value as string or default if not present
func GetString(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}
