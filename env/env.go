package env

import (
	"fmt"
	"os"
	"strconv"
)

func GetEnv[T string | bool | int | float64](key string, defaultValue *T) (*T, error) {
	value, exists := os.LookupEnv(key)

	if !exists && defaultValue != nil {
		return defaultValue, nil
	}

	if !exists && defaultValue == nil {
		return nil, fmt.Errorf("Environment variable %s not found and no default value provided", key)
	}

	var rs T
	switch p := any(&rs).(type) {
	// String types
	case *string:
		*p = value

	// Boolean type
	case *bool:
		parsed, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}
		*p = parsed

	// Numeric types
	case *int:
		parsed, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		*p = parsed
	case *float64:
		parsed, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		*p = parsed
	}

	return &rs, nil
}
