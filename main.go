package main

import (
	"netdisk/router"
	"netdisk/settings"

	log "github.com/sirupsen/logrus"

	_ "netdisk/utils/yglog"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /

func main() {
	// listen
	r := router.New()

	addr := settings.ServerAddr()
	log.Println("start service base on " + addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
