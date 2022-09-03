package main

import (
	"fmt"
	_ "porn/common"
	"porn/common/config"
	"porn/common/logger"
	"porn/server/app/admin"
)

func main() {
	r := admin.NewRouter(logger.Logger)
	r.Run(fmt.Sprintf("%v:%v", config.Config.Server.Host, config.Config.Server.Port))
}
