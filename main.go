package main

import (
	"linebotim/router"
	"linebotim/utils"
	"log"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("load config err = ", err)
	}

	e := router.Router()
	e.Run(config.ServerAddress)
}
