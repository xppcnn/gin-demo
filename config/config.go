package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	RunMode   string    `mapstructure:"run_mode"`
	Server    Server    `mapstructure:"server"`
	LogConfig LogConfig `mapstructure:"log_config"`
	App       App       `mapstructure:"app"`
	DataBase  DataBase  `mapstructure:"data_base"`
}
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
	Version    string `mapstructure:"version"`
}
type Server struct {
	HttpPort     int `mapstructure:"http_port"`
	ReadTimeout  int `mapstructure:"read_timeout"`
	WriteTimeout int `mapstructure:"write_timeout"`
}

type App struct {
	PageSize  int    `mapstructure:"page_size"`
	JwtSecret string `mapstructure:"jwt_secret"`
}

type DataBase struct {
	Type        string `mapstructure:"type"`
	User        string `mapstructure:"user"`
	PassWord    string `mapstructure:"password"`
	Host        string `mapstructure:"host"`
	Name        string `mapstructure:"name"`
	TablePrefix string `mapstructure:"table_prefix"`
}

var config = new(Config)

func init() {
	viper.SetConfigFile("./config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if err = viper.Unmarshal(config); err != nil {
		panic(err)
	}
	viper.WatchConfig()
}

func GetConfig() Config {
	return *config
}
