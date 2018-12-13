package main

import (
	"slinks/models"
	"slinks/routers"
)

func main() {
	models.InitStarter()
	router := routers.InitRouter()
	router.Run(":9090")
}

