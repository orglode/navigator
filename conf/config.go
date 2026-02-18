package conf

import (
	"os"

	"github.com/BurntSushi/toml"
)

var ConfGlobal *Config

func Init() *Config {
	cfg := &Config{}
	configPath := os.Getenv("CONFIG_FILE")
	if configPath == "" {
		configPath = "config/config.toml" // 默认路径
	}
	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		panic("config.toml is err !!")
	}
	ConfGlobal = cfg
	return cfg
}

type Config struct {
	Server *server
	Db     *Db
	Redis  *RedisConfig
}

type server struct {
	Name     string `toml:"name"`
	Addr     string `toml:"addr"`
	Env      string `toml:"env"`
	JwtToken string `toml:"jwt_token"`
	LogsPath string `toml:"logs_path"`
}

type Db struct {
	Master *MysqlConfig `toml:"master"`
	Slave  *MysqlConfig `toml:"slave"`
}

type MysqlConfig struct {
	Drive           string `toml:"drive"`
	Url             string `toml:"url"`
	MaxIdleConn     int    `toml:"max_idle_conn"`
	MaxOpenConn     int    `toml:"max_open_conn"`
	ConnMaxLifeTime int    `toml:"conn_max_life_time"`
}

type RedisConfig struct {
	Name     string `toml:"name"`
	Addr     string `json:"addr"`
	PassWord string `json:"pass_word"`
	Db       int    `json:"db"`
}
