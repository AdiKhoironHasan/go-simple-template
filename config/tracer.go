package config

import "github.com/spf13/viper"

// JaegerEndpoint retrieves the value of the JAEGER_ENDPOINT environment variable.
func JaegerEndpoint() string {
	return viper.GetString("JAEGER_ENDPOINT")
}
