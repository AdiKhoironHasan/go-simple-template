package config

import "github.com/spf13/viper"

// StorageDriver retrieves the value of the STORAGE_DRIVER environment variable.
func StorageDriver() string {
	return viper.GetString("STORAGE_DRIVER")
}

// MinioEndpoint retrieves the value of the MINIO_ENDPOINT environment variable.
func MinioEndpoint() string {
	return viper.GetString("MINIO_ENDPOINT")
}

// MinioAccessKeyID retrieves the value of the MINIO_ACCESS_KEY_ID environment variable.
func MinioAccessKeyID() string {
	return viper.GetString("MINIO_ACCESS_KEY_ID")
}

// MinioAccessKeySecret retrieves the value of the MINIO_ACCESS_KEY_SECRET environment variable.
func MinioAccessKeySecret() string {
	return viper.GetString("MINIO_ACCESS_KEY_SECRET")
}

// MinioBucketName retrieves the value of the MINIO_BUCKET_NAME environment variable.
func MinioBucketName() string {
	return viper.GetString("MINIO_BUCKET_NAME")
}

// GCSCredentialsFile retrieves the value of the GCS_CREDENTIALS_FILE environment variable.
func GCSCredentialsFile() string {
	return viper.GetString("GCS_CREDENTIALS_FILE")
}

// GCSBucketName retrieves the value of the GCS_BUCKET_NAME environment variable.
func GCSBucketName() string {
	return viper.GetString("GCS_BUCKET_NAME")
}
