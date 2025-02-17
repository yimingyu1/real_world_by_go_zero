package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	// Redis config
	Redis struct {
		Host string
		Type string
	}

	DB struct {
		DataSource string
	}
}
