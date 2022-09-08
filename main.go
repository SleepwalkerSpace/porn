package main

import (
	_ "porn/common"
	"porn/common/config"
	"porn/common/logger"
	"porn/server/app/admin"
)

func main() {
	server := admin.NewRouter(logger.Logger)
	if err := server.Run(config.Config.ServerAddress()); err != nil {
		panic(err)
	}
}
