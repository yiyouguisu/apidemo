package model

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbname"`
	PoolSize int    `mapstructure:"pool_size"`
	Slow     int    `mapstructure:"slow"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type LogConfig struct {
	Level         string `mapstructure:"level"`
	CustomLogName string `mapstructure:"custom_name"`
	AccessLogName string `mapstructure:"access_name"`
}

type AgolloConfig struct {
	AppId string `mapstructure:"appid"`
	ConfigServerURL string `mapstructure:"config_server_url"`
}

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Postgres PostgresConfig `mapstructure:"postgres"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Log      LogConfig      `mapstructure:"log"`
	AgolloConfig AgolloConfig `mapstructure:"agollo"`
}
