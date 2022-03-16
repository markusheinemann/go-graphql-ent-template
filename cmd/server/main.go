package main

import (
	"log"
	"net/http"

	"gitlab.com/trustify/core/config"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	log.Printf("starting http server on port %s...", config.C.HttpServer.Port)
	if err := http.ListenAndServe(":"+config.C.HttpServer.Port, nil); err != nil {
		log.Fatalf("error spinning up http server %v", err)
	}

}
