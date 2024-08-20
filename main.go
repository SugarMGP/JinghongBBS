package main

import (
	"BBS/app/midwares"
	"BBS/config/database"
	"BBS/config/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	router.Init(r)

	err := r.Run()
	if err != nil {
		log.Fatal("Server start failed: ", err)
	}
}
