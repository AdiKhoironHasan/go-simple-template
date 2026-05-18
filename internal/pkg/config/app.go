package config

import "github.com/spf13/viper"

func AppName() string {
	return viper.GetString("APP_NAME")
}

func AppVersion() string {
	return viper.GetString("APP_VERSION")
}

func AppPort() int {
	return viper.GetInt("APP_PORT")
}

func AppDebug() bool {
	return viper.GetBool("APP_DEBUG")
}

func AppSecretKey() string {
	return viper.GetString("APP_SECRET_KEY")
}

func AppRefreshKey() string {
	return viper.GetString("APP_REFRESH_KEY")
}
