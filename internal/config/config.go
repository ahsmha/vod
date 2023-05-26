package config

// Config represents the configuration settings for the Video on Demand (VOD) microservice.
type Config struct {
	// Database configuration
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string

	// Server configuration
	ServerPort int

	// Video storage configuration
	VideoStoragePath string

	// Encoding configuration
	EncodingEnabled bool
	EncodingFormat  string
	EncodingPreset  string

	// Authentication configuration
	AuthSecret string
}

// LoadConfig loads the configuration from environment variables or any other source.
func LoadConfig() *Config {
	// You can implement the logic to load the configuration from environment variables or any other source.
	// Here's an example of initializing the configuration with default values.

	return &Config{
		DBHost:           "localhost",
		DBPort:           5432,
		DBUser:           "admin",
		DBPassword:       "password",
		DBName:           "vod_db",
		ServerPort:       8080,
		VideoStoragePath: "/path/to/video/storage",
		EncodingEnabled:  true,
		EncodingFormat:   "mp4",
		EncodingPreset:   "default",
		AuthSecret:       "your-auth-secret",
	}
}
