package config

import (
	"os"
	"strconv"
)

func NewConfig() *Config {
	return &Config{
		AppConfig: AppConfig{
			AppName:    getEnv("APP_NAME", "go-simple-template"),
			AppEnv:     getEnv("APP_ENV", "development"),
			AppVersion: getEnv("APP_VERSION", "v0.0.1"),
			AppHost:    getEnv("APP_HOST", "localhost"),
			AppPort:    getEnvAsInt("APP_PORT", 8000),
		},
		DBconfig: DBconfig{
			DBdriver:   getEnv("DB_DRIVER", ""),
			DBhost:     getEnv("DB_HOST", "localhost"),
			DBport:     getEnv("DB_PORT", "3306"),
			DBuser:     getEnv("DB_USER", "root"),
			DBpassword: getEnv("DB_PASSWORD", ""),
			DBname:     getEnv("DB_NAME", "go-simple-template"),
		},
		CacheConfig: CacheConfig{
			CacheDriver: getEnv("CACHE_DRIVER", "redis"),
			Redis: Redis{
				RedisHost:     getEnv("REDIS_HOST", "localhost"),
				RedisPort:     getEnvAsInt("REDIS_PORT", 6379),
				RedisDB:       getEnvAsInt("REDIS_DB", 0),
				RedisPassword: getEnv("REDIS_PASSWORD", ""),
			},
		},
		StorageConfig: StorageConfig{
			StorageDriver: getEnv("STORAGE_DRIVER", "minio"),
			Minio: Minio{
				MinioEndpoint:        getEnv("MINIO_ENDPOINT", ""),
				MinioAccessKeyID:     getEnv("MINIO_ACCESS_KEY_ID", ""),
				MinioAccessKeySecret: getEnv("MINIO_ACCESS_KEY_SECRET", ""),
				MinioBucketName:      getEnv("MINIO_BUCKET_NAME", ""),
			},
			GCS: GCS{
				CredentialsFile: getEnv("GCS_CREDENTIALS_FILE", ""),
				GCSBucketName:   getEnv("GCS_BUCKET_NAME", ""),
			},
		},
		Tracer: Tracer{
			JaegerURL: getEnv("JAEGER_URL", "http://localhost:14268/api/traces"),
		},
	}
}

type Config struct {
	AppConfig
	DBconfig
	CacheConfig
	StorageConfig
	Tracer
}

type AppConfig struct {
	AppName    string
	AppEnv     string
	AppVersion string
	AppHost    string
	AppPort    int
}

type DBconfig struct {
	DBdriver   string
	DBhost     string
	DBport     string
	DBuser     string
	DBname     string
	DBpassword string
}

type CacheConfig struct {
	CacheDriver string
	Redis       Redis
}

type StorageConfig struct {
	StorageDriver string
	Minio         Minio
	GCS           GCS
}

type Minio struct {
	MinioEndpoint        string
	MinioAccessKeyID     string
	MinioAccessKeySecret string
	MinioBucketName      string
}

type GCS struct {
	CredentialsFile string
	GCSBucketName   string
}

type Redis struct {
	RedisHost     string
	RedisPort     int
	RedisDB       int
	RedisPassword string
}

type Tracer struct {
	JaegerURL string
}

func getEnv(key string, defaultVal string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}

	if nextValue := os.Getenv(key); nextValue != "" {
		return nextValue
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
