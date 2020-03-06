package main

import (
	"fmt"
	"log"
	"netdisk/router"

	_ "netdisk/settings"
)

func main() {
	// listen
	r := router.DefaultRouter()

	addr := ":8080"
	fmt.Println("start service base on " + addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
