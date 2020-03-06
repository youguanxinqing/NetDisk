package main

import (
	"netdisk/router"
	"netdisk/settings"

	log "github.com/sirupsen/logrus"

	_ "netdisk/utils/yglog"
)

func main() {
	// listen
	r := router.New()

	addr := settings.ServerAddr()
	log.Println("start service base on " + addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
