package conf

import (
	"github.com/BurntSushi/toml"
)

func Init() *Config {
	cfg := &Config{}
	confPath := "config/config.toml"

	if _, err := toml.DecodeFile(confPath, &cfg); err != nil {
		panic("config.toml is err !!")
	}
	return cfg
}

type Config struct {
	Server *server
	Mysql  *mysqlConfig
	Redis  *redisConfig
}

type server struct {
	Name string `toml:"name"`
	Addr string `toml:"addr"`
	Env  string `toml:"env"`
}

type mysqlConfig struct {
	Name   string `toml:"name"`
	Master string `toml:"master"`
	Slave  string `toml:"slave"`
}

type redisConfig struct {
	Name     string `toml:"name"`
	Addr     string `json:"addr"`
	PassWord string `json:"pass_word"`
	Db       int    `json:"db"`
}
