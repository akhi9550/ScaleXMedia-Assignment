package main

import (
	"log"
	"scalexmedia-assignment/pkg/api/middleware"
	"scalexmedia-assignment/pkg/config"
	"scalexmedia-assignment/pkg/di"
	"scalexmedia-assignment/pkg/helper"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	authHelper := helper.NewAuthHelper(cfg)
	middleware.SetHelper(authHelper)

	server, err := di.InitializeAPI(cfg)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	server.Start()
}
