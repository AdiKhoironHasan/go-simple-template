package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	LoadEnv("../.env")

	fmt.Println("AppPort:", AppPort())
	fmt.Println("AppHost:", AppHost())

	fmt.Println("DBDriver:", DBDriver())
	fmt.Println("DBHost:", DBHost())
	fmt.Println("DBPort:", DBPort())
	fmt.Println("DBName:", DBName())
	fmt.Println("DBUser:", DBUser())
	fmt.Println("DBPassword:", DBPassword())

	fmt.Println("StorageDriver:", StorageDriver())
	fmt.Println("MinioEndpoint:", MinioEndpoint())
	fmt.Println("MinioAccessKeyID:", MinioAccessKeyID())
	fmt.Println("MinioAccessKeySecret:", MinioAccessKeySecret())
	fmt.Println("MinioBucketName:", MinioBucketName())
	fmt.Println("GCSCredentialsFile:", GCSCredentialsFile())
	fmt.Println("GCSBucketName:", GCSBucketName())

	fmt.Println("RedisHost:", RedisHost())
	fmt.Println("RedisPort:", RedisPort())
	fmt.Println("RedisDB:", RedisDB())
	fmt.Println("RedisPassword:", RedisPassword())

	fmt.Println("JaegerEndpoint:", JaegerEndpoint())
}
