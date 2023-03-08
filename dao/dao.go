package dao

import (
	"github.com/gomodule/redigo/redis"
	"github.com/orglode/navigator/conf"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dao struct {
	conf        *conf.Config
	MySqlMaster *gorm.DB
	MySqlSlave  *gorm.DB
	Redis       redis.Conn
}

func NewDao(conf *conf.Config) *Dao {
	return &Dao{
		conf:        conf,
		MySqlMaster: initMysqlDb(conf.DbMaster),
		MySqlSlave:  initMysqlDb(conf.DbSlave),
		Redis:       initRedis(conf),
	}
}
