package common

import (
	"fmt"
	log "porn/common/logger"
	"porn/library/logger"

	cfg "porn/common/config"
	"porn/library/config"

	cce "porn/common/cache"
	"porn/library/cache"

	per "porn/common/persistence"
	"porn/library/persistence"
)

func init() {
	var err error
	log.Logger, err = logger.New("./", "logger")
	if err != nil {
		panic(fmt.Sprintf("common.init logger.New error:%v", err))
	}

	cfg.Config = new(cfg.Cfg)
	if err := config.Load("./", "config", "toml", cfg.Config); err != nil {
		panic(fmt.Sprintf("common.init config.Load error:%v", err))
	}

	cce.Cache, err = cache.Connect(cfg.Config.Cache.Addr, cfg.Config.Cache.Pswd, 0)
	if err != nil {
		panic(fmt.Sprintf("common.init cache.Connect error:%v", err))
	}

	per.Persistence, err = persistence.ConnectDbMySql(cfg.Config.Mysql.Dsn)
	if err != nil {
		panic(fmt.Sprintf("common.init persistence.ConnectDbMySql error:%v", err))
	}

	log.Logger.Infof("[CONF]%+v", *cfg.Config)
}
