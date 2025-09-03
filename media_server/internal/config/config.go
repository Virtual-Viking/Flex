package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// Config holds all configuration for the application
type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Media    MediaConfig    `mapstructure:"media"`
	External ExternalConfig `mapstructure:"external"`
	Logging  LoggingConfig  `mapstructure:"logging"`
}

// AppConfig holds application-specific configuration
type AppConfig struct {
	Name        string   `mapstructure:"name"`
	Environment string   `mapstructure:"environment"`
	Host        string   `mapstructure:"host"`
	Port        string   `mapstructure:"port"`
	Origins     []string `mapstructure:"allowed_origins"`
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host            string        `mapstructure:"host"`
	Port            string        `mapstructure:"port"`
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	Name            string        `mapstructure:"name"`
	SSLMode         string        `mapstructure:"sslmode"`
	MaxConnections  int           `mapstructure:"max_connections"`
	MaxIdleTime     time.Duration `mapstructure:"max_idle_time"`
}

// RedisConfig holds Redis configuration
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret    string        `mapstructure:"secret"`
	ExpiresIn time.Duration `mapstructure:"expires_in"`
}

// MediaConfig holds media processing configuration
type MediaConfig struct {
	RootPath      string `mapstructure:"root_path"`
	UploadPath    string `mapstructure:"upload_path"`
	PosterPath    string `mapstructure:"poster_path"`
	ThumbnailPath string `mapstructure:"thumbnail_path"`
	FFmpegPath    string `mapstructure:"ffmpeg_path"`
	MediaInfoPath string `mapstructure:"mediainfo_path"`
}

// ExternalConfig holds external API configuration
type ExternalConfig struct {
	TMDBAPIKey string `mapstructure:"tmdb_api_key"`
	OMDBAPIKey string `mapstructure:"omdb_api_key"`
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	config := &Config{
		App: AppConfig{
			Name:        getEnv("APP_NAME", "Flex Media Server"),
			Environment: getEnv("ENV", "development"),
			Host:        getEnv("HOST", "0.0.0.0"),
			Port:        getEnv("PORT", "8080"),
			Origins:     strings.Split(getEnv("ALLOWED_ORIGINS", "http://localhost:3000"), ","),
		},
		Database: DatabaseConfig{
			Host:            getEnv("DB_HOST", "localhost"),
			Port:            getEnv("DB_PORT", "5432"),
			User:            getEnv("DB_USER", "flex_user"),
			Password:        getEnv("DB_PASSWORD", "flex_password"),
			Name:            getEnv("DB_NAME", "flex_dev"),
			SSLMode:         getEnv("DB_SSLMODE", "disable"),
			MaxConnections:  getEnvAsInt("DB_MAX_CONNECTIONS", 25),
			MaxIdleTime:     getEnvAsDuration("DB_MAX_IDLE_TIME", 15*time.Minute),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
		JWT: JWTConfig{
			Secret:    getEnv("JWT_SECRET", "your-secret-key"),
			ExpiresIn: getEnvAsDuration("JWT_EXPIRES_IN", 24*time.Hour),
		},
		Media: MediaConfig{
			RootPath:      getEnv("MEDIA_ROOT_PATH", "/media/library"),
			UploadPath:    getEnv("UPLOAD_PATH", "/tmp/flex-uploads"),
			PosterPath:    getEnv("POSTER_PATH", "/tmp/flex-posters"),
			ThumbnailPath: getEnv("THUMBNAIL_PATH", "/tmp/flex-thumbnails"),
			FFmpegPath:    getEnv("FFMPEG_PATH", "ffmpeg"),
			MediaInfoPath: getEnv("MEDIAINFO_PATH", "mediainfo"),
		},
		External: ExternalConfig{
			TMDBAPIKey: getEnv("TMDB_API_KEY", ""),
			OMDBAPIKey: getEnv("OMDB_API_KEY", ""),
		},
		Logging: LoggingConfig{
			Level:  getEnv("LOG_LEVEL", "info"),
			Format: getEnv("LOG_FORMAT", "console"),
		},
	}

	return config, nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt gets an environment variable as an integer or returns a default value
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getEnvAsDuration gets an environment variable as a duration or returns a default value
func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}