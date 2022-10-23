package main

import (
	"linebotim/router"
)

func main() {
	e := router.Router()
	e.Run(":8081")
}
