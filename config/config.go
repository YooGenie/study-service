package config

import "github.com/jinzhu/configor"

var Config = struct {
	HttpPort string
	Database struct {
		Driver     string
		User       string
		Connection string
	}
}{}

func InitConfig(cfg string) {
	configor.Load(&Config, cfg)
}
