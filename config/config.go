package config

type Config struct {
	AppConfig
	DBconfig
}

type AppConfig struct {
	AppHost string
	AppPort int
}

type DBconfig struct {
	DBdriver   string
	DBhost     string
	DBport     string
	DBuser     string
	DBname     string
	DBpassword string
}
