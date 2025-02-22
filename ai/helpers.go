package ai

import (
	"runtime"

	"github.com/invopop/jsonschema"
)

// GenerateSchema generates a JSON schema for the AI
func GenerateSchema[T any]() interface{} {
	// Structured Outputs uses a subset of JSON schema
	// These flags are necessary to comply with the subset
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T
	schema := reflector.Reflect(v)
	return schema
}

func GetOS() string {
	os := runtime.GOOS
	switch os {
	case "darwin":
		return "macOS"
	case "linux":
		return "Linux"
	case "windows":
		return "Windows"
	default:
		return os
	}
}
